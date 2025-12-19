/*
 * Copyright (c) 2019-present Sonatype, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package errors

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"reflect"
	"syscall"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// ErrorMessage is a standardized error message structure
type ErrorMessage struct {
	Operation   string
	ResourceID  string
	ErrorDetail string
}

// APIError creates a standardized error for API operations
func APIError(operation string, resourceType string, details string) (string, string) {
	title := fmt.Sprintf("Error %s %s", operation, resourceType)
	message := fmt.Sprintf("Could not %s %s: %s", operation, resourceType, details)
	return title, message
}

// NotFoundError creates a standardized not found error
func NotFoundError(resourceType string, resourceID string) (string, string) {
	title := fmt.Sprintf("%s Not Found", resourceType)
	message := fmt.Sprintf("The %s with ID '%s' was not found. It may have been deleted outside of Terraform.", resourceType, resourceID)
	return title, message
}

// ValidationError creates a standardized validation error
func ValidationError(field string, reason string) (string, string) {
	title := fmt.Sprintf("Invalid %s", field)
	message := fmt.Sprintf("The %s value is invalid: %s", field, reason)
	return title, message
}

// ConflictError creates a standardized conflict error
func ConflictError(resourceType string, details string) (string, string) {
	title := fmt.Sprintf("Conflict creating %s", resourceType)
	message := fmt.Sprintf("A conflict occurred: %s", details)
	return title, message
}

// UnauthorizedError creates a standardized unauthorized error
func UnauthorizedError(operation string) (string, string) {
	title := fmt.Sprintf("Unauthorized %s", operation)
	message := fmt.Sprintf("You do not have permission to %s. Please check your credentials and permissions.", operation)
	return title, message
}

// TimeoutError creates a standardized timeout error
func TimeoutError(operation string, resourceType string) (string, string) {
	title := fmt.Sprintf("Timeout %s", operation)
	message := fmt.Sprintf("Operation timed out while %s %s. The resource may have been created but confirmation could not be received.", operation, resourceType)
	return title, message
}

// ParseAPIError extracts error information from an HTTP response
func ParseAPIError(response *http.Response) string {
	if response == nil {
		return "Unknown error"
	}

	// Try to read the response body for error details
	if response.Body != nil {
		body, err := io.ReadAll(response.Body)
		if err == nil && len(body) > 0 {
			return fmt.Sprintf("%s: %s", response.Status, string(body))
		}
	}

	return response.Status
}

// AddAPIErrorDiagnostic adds a standardized API error to diagnostics
func AddAPIErrorDiagnostic(diags *diag.Diagnostics, operation string, resourceType string, response *http.Response, originalErr error) {
	title, baseMessage := APIError(operation, resourceType, "")
	details := ParseAPIError(response)
	if originalErr != nil {
		details = fmt.Sprintf("%s: %v", details, originalErr)
	}
	diags.AddError(title, fmt.Sprintf("%s %s", baseMessage, details))
}

// AddNotFoundDiagnostic adds a standardized not found error to diagnostics
func AddNotFoundDiagnostic(diags *diag.Diagnostics, resourceType string, resourceID string) {
	title, message := NotFoundError(resourceType, resourceID)
	diags.AddError(title, message)
}

// AddValidationDiagnostic adds a standardized validation error to diagnostics
func AddValidationDiagnostic(diags *diag.Diagnostics, field string, reason string) {
	title, message := ValidationError(field, reason)
	diags.AddError(title, message)
}

// AddConflictDiagnostic adds a standardized conflict error to diagnostics
func AddConflictDiagnostic(diags *diag.Diagnostics, resourceType string, details string) {
	title, message := ConflictError(resourceType, details)
	diags.AddError(title, message)
}

// AddUnauthorizedDiagnostic adds a standardized unauthorized error to diagnostics
func AddUnauthorizedDiagnostic(diags *diag.Diagnostics, operation string) {
	title, message := UnauthorizedError(operation)
	diags.AddError(title, message)
}

// AddTimeoutDiagnostic adds a standardized timeout error to diagnostics
func AddTimeoutDiagnostic(diags *diag.Diagnostics, operation string, resourceType string) {
	title, message := TimeoutError(operation, resourceType)
	diags.AddError(title, message)
}

// AddForbiddenDiagnostic adds a standardized 403 Forbidden error to diagnostics
func AddForbiddenDiagnostic(diags *diag.Diagnostics, operation string) {
	title := fmt.Sprintf("Forbidden %s", operation)
	message := fmt.Sprintf("You do not have permission to %s (HTTP 403). Please check your access credentials and permissions.", operation)
	diags.AddError(title, message)
}

// AddServerErrorDiagnostic adds a standardized 5xx error to diagnostics
func AddServerErrorDiagnostic(diags *diag.Diagnostics, message string, statusCode int) {
	title := fmt.Sprintf("Server Error (%d)", statusCode)
	details := fmt.Sprintf("The server returned an error (HTTP %d). Details: %s", statusCode, message)
	diags.AddError(title, details)
}

// AddClientErrorDiagnostic adds a standardized 4xx error to diagnostics
func AddClientErrorDiagnostic(diags *diag.Diagnostics, message string, statusCode int) {
	title := fmt.Sprintf("Client Error (%d)", statusCode)
	details := fmt.Sprintf("The request could not be processed (HTTP %d). Details: %s", statusCode, message)
	diags.AddError(title, details)
}

// IsNotFound checks if an HTTP response is a 404
func IsNotFound(statusCode int) bool {
	return statusCode == http.StatusNotFound
}

// IsForbidden checks if an HTTP response is a 403
func IsForbidden(statusCode int) bool {
	return statusCode == http.StatusForbidden
}

// IsUnauthorized checks if an HTTP response is a 401
func IsUnauthorized(statusCode int) bool {
	return statusCode == http.StatusUnauthorized
}

// IsConflict checks if an HTTP response is a 409
func IsConflict(statusCode int) bool {
	return statusCode == http.StatusConflict
}

// IsClientError checks if the status code is a 4xx error
func IsClientError(statusCode int) bool {
	return statusCode >= 400 && statusCode < 500
}

// IsServerError checks if the status code is a 5xx error
func IsServerError(statusCode int) bool {
	return statusCode >= 500 && statusCode < 600
}

// HandleAPIError processes API errors and detects network-related errors.
// It adds the error to diagnostics with appropriate context.
// This is useful for providers that need to distinguish between network errors
// and API errors during resource operations.
func HandleAPIError(message string, err *error, httpResponse *http.Response, respDiags *diag.Diagnostics) {
	networkError, errorMessage := detectNetworkError(*err)
	if networkError {
		respDiags.AddError(
			errorMessage,
			fmt.Sprintf("Networking Error: %s (%v)", errorMessage, *err),
		)
	} else {
		if httpResponse != nil {
			respDiags.AddError(
				message,
				fmt.Sprintf("%s: %s: %s", *err, httpResponse.Status, extractResponseBody(httpResponse)),
			)
		} else {
			respDiags.AddError(
				message,
				fmt.Sprintf("Unexpected Error: %v ('%s'): ", *err, reflect.TypeOf(*err)),
			)
		}
	}
}

// HandleAPIWarning processes API warnings and adds them to diagnostics.
// Use this for non-critical API issues that should be reported to the user.
func HandleAPIWarning(message string, err *error, httpResponse *http.Response, respDiags *diag.Diagnostics) {
	status := ""
	if httpResponse != nil {
		status = httpResponse.Status
	}
	respDiags.AddWarning(
		message,
		fmt.Sprintf("%s: %s: %s", *err, status, extractResponseBody(httpResponse)),
	)
}

// detectNetworkError identifies if an error is network-related (as opposed to API-related)
// It checks for specific error types like DNS errors, connection errors, and timeouts.
func detectNetworkError(err error) (bool, string) {
	// Check for specific error types
	if opErr, ok := err.(*net.OpError); ok {
		// Network operation error (dial, read, write, etc.)
		return true, fmt.Sprintf("OpError: %s, %s", opErr.Op, opErr.Net)
	}

	if dnsErr, ok := err.(*net.DNSError); ok {
		// DNS resolution error
		return true, fmt.Sprintf("DNS Error: %v", dnsErr)
	}

	if errors.Is(err, context.DeadlineExceeded) {
		// Timeout error
		return true, fmt.Sprintf("Connection timed out: %v", err)
	}

	// General network error check
	if errors.Is(err, syscall.ECONNREFUSED) {
		return true, fmt.Sprintf("Connection refused: %v", err)
	}

	return false, ""
}

// extractResponseBody reads and returns the body from an HTTP response.
// This is a helper function for error reporting.
func extractResponseBody(httpResponse *http.Response) []byte {
	if httpResponse == nil || httpResponse.Body == nil {
		return []byte{}
	}
	body, _ := io.ReadAll(httpResponse.Body)
	err := httpResponse.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	return body
}
