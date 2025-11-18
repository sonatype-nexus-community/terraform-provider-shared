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
	"fmt"
	"io"
	"net/http"

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
