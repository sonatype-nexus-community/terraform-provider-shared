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
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"syscall"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TestAPIError(t *testing.T) {
	title, message := APIError("creating", "User", "invalid email")
	if title != "Error creating User" {
		t.Fatalf("APIError title = %s, expected 'Error creating User'", title)
	}
	if message != "Could not creating User: invalid email" {
		t.Fatalf("APIError message = %s, expected 'Could not creating User: invalid email'", message)
	}
}

func TestNotFoundError(t *testing.T) {
	title, message := NotFoundError("User", "user123")
	if title != "User Not Found" {
		t.Fatalf("NotFoundError title = %s", title)
	}
	if message == "" {
		t.Fatal("NotFoundError message should not be empty")
	}
}

func TestValidationError(t *testing.T) {
	title, message := ValidationError("email", "must be a valid email address")
	if title != "Invalid email" {
		t.Fatalf("ValidationError title = %s", title)
	}
	if message == "" {
		t.Fatal("ValidationError message should not be empty")
	}
}

func TestConflictError(t *testing.T) {
	title, message := ConflictError("User", "user already exists")
	if title != "Conflict creating User" {
		t.Fatalf("ConflictError title = %s", title)
	}
	if message == "" {
		t.Fatal("ConflictError message should not be empty")
	}
}

func TestUnauthorizedError(t *testing.T) {
	title, message := UnauthorizedError("create user")
	if title != "Unauthorized create user" {
		t.Fatalf("UnauthorizedError title = %s", title)
	}
	if message == "" {
		t.Fatal("UnauthorizedError message should not be empty")
	}
}

func TestTimeoutError(t *testing.T) {
	title, message := TimeoutError("creating", "User")
	if title != "Timeout creating" {
		t.Fatalf("TimeoutError title = %s", title)
	}
	if message == "" {
		t.Fatal("TimeoutError message should not be empty")
	}
}

func TestIsNotFound(t *testing.T) {
	if !IsNotFound(404) {
		t.Fatal("IsNotFound(404) should return true")
	}
	if IsNotFound(200) {
		t.Fatal("IsNotFound(200) should return false")
	}
}

func TestIsForbidden(t *testing.T) {
	if !IsForbidden(403) {
		t.Fatal("IsForbidden(403) should return true")
	}
	if IsForbidden(401) {
		t.Fatal("IsForbidden(401) should return false")
	}
}

func TestIsUnauthorized(t *testing.T) {
	if !IsUnauthorized(401) {
		t.Fatal("IsUnauthorized(401) should return true")
	}
	if IsUnauthorized(403) {
		t.Fatal("IsUnauthorized(403) should return false")
	}
}

func TestIsConflict(t *testing.T) {
	if !IsConflict(409) {
		t.Fatal("IsConflict(409) should return true")
	}
	if IsConflict(400) {
		t.Fatal("IsConflict(400) should return false")
	}
}

func TestIsClientError(t *testing.T) {
	tests := []struct {
		code     int
		expected bool
	}{
		{400, true},
		{404, true},
		{499, true},
		{200, false},
		{500, false},
	}

	for _, test := range tests {
		result := IsClientError(test.code)
		if result != test.expected {
			t.Fatalf("IsClientError(%d) = %v, expected %v", test.code, result, test.expected)
		}
	}
}

func TestIsServerError(t *testing.T) {
	tests := []struct {
		code     int
		expected bool
	}{
		{500, true},
		{502, true},
		{599, true},
		{400, false},
		{200, false},
	}

	for _, test := range tests {
		result := IsServerError(test.code)
		if result != test.expected {
			t.Fatalf("IsServerError(%d) = %v, expected %v", test.code, result, test.expected)
		}
	}
}

func TestParseAPIError(t *testing.T) {
	// Test with nil response
	result := ParseAPIError(nil)
	if result != "Unknown error" {
		t.Fatalf("ParseAPIError(nil) should return 'Unknown error', got %s", result)
	}

	// Test with HTTP response without body
	resp := &http.Response{
		Status: "400 Bad Request",
		Body:   io.NopCloser(bytes.NewReader([]byte(""))),
	}
	result = ParseAPIError(resp)
	if result != "400 Bad Request" {
		t.Fatalf("ParseAPIError failed: got %s", result)
	}

	// Test with HTTP response with body
	resp = &http.Response{
		Status: "500 Internal Server Error",
		Body:   io.NopCloser(bytes.NewReader([]byte("Server error details"))),
	}
	result = ParseAPIError(resp)
	if result != "500 Internal Server Error: Server error details" {
		t.Fatalf("ParseAPIError with body failed: got %s", result)
	}
}

func TestAddAPIErrorDiagnostic(t *testing.T) {
	diags := diag.Diagnostics{}
	resp := &http.Response{
		Status: "400 Bad Request",
		Body:   io.NopCloser(bytes.NewReader([]byte(""))),
	}

	AddAPIErrorDiagnostic(&diags, "creating", "User", resp, nil)
	if !diags.HasError() {
		t.Fatal("AddAPIErrorDiagnostic should add an error to diagnostics")
	}

	// Test with error
	diags = diag.Diagnostics{}
	testErr := &net.OpError{Op: "dial", Net: "tcp"}
	AddAPIErrorDiagnostic(&diags, "creating", "User", resp, testErr)
	if !diags.HasError() {
		t.Fatal("AddAPIErrorDiagnostic should add an error to diagnostics")
	}
}

func TestAddNotFoundDiagnostic(t *testing.T) {
	diags := diag.Diagnostics{}
	AddNotFoundDiagnostic(&diags, "User", "user123")
	if !diags.HasError() {
		t.Fatal("AddNotFoundDiagnostic should add an error to diagnostics")
	}
}

func TestAddValidationDiagnostic(t *testing.T) {
	diags := diag.Diagnostics{}
	AddValidationDiagnostic(&diags, "email", "must be valid")
	if !diags.HasError() {
		t.Fatal("AddValidationDiagnostic should add an error to diagnostics")
	}
}

func TestAddConflictDiagnostic(t *testing.T) {
	diags := diag.Diagnostics{}
	AddConflictDiagnostic(&diags, "User", "already exists")
	if !diags.HasError() {
		t.Fatal("AddConflictDiagnostic should add an error to diagnostics")
	}
}

func TestAddUnauthorizedDiagnostic(t *testing.T) {
	diags := diag.Diagnostics{}
	AddUnauthorizedDiagnostic(&diags, "create user")
	if !diags.HasError() {
		t.Fatal("AddUnauthorizedDiagnostic should add an error to diagnostics")
	}
}

func TestAddTimeoutDiagnostic(t *testing.T) {
	diags := diag.Diagnostics{}
	AddTimeoutDiagnostic(&diags, "creating", "User")
	if !diags.HasError() {
		t.Fatal("AddTimeoutDiagnostic should add an error to diagnostics")
	}
}

func TestDetectNetworkError_OpError(t *testing.T) {
	opErr := &net.OpError{Op: "dial", Net: "tcp"}
	isNet, msg := detectNetworkError(opErr)
	if !isNet {
		t.Fatal("detectNetworkError should return true for OpError")
	}
	if msg == "" {
		t.Fatal("detectNetworkError should return a non-empty message for OpError")
	}
}

func TestDetectNetworkError_DNSError(t *testing.T) {
	dnsErr := &net.DNSError{Err: "name resolution failed"}
	isNet, msg := detectNetworkError(dnsErr)
	if !isNet {
		t.Fatal("detectNetworkError should return true for DNSError")
	}
	if msg == "" {
		t.Fatal("detectNetworkError should return a non-empty message for DNSError")
	}
}

func TestDetectNetworkError_DeadlineExceeded(t *testing.T) {
	isNet, msg := detectNetworkError(context.DeadlineExceeded)
	if !isNet {
		t.Fatal("detectNetworkError should return true for DeadlineExceeded")
	}
	if msg == "" {
		t.Fatal("detectNetworkError should return a non-empty message for DeadlineExceeded")
	}
}

func TestDetectNetworkError_ECONNREFUSED(t *testing.T) {
	isNet, msg := detectNetworkError(syscall.ECONNREFUSED)
	if !isNet {
		t.Fatal("detectNetworkError should return true for ECONNREFUSED")
	}
	if msg == "" {
		t.Fatal("detectNetworkError should return a non-empty message for ECONNREFUSED")
	}
}

func TestDetectNetworkError_NonNetworkError(t *testing.T) {
	isNet, msg := detectNetworkError(new(struct{ error }))
	if isNet {
		t.Fatal("detectNetworkError should return false for non-network errors")
	}
	if msg != "" {
		t.Fatalf("detectNetworkError should return empty string for non-network error, got %s", msg)
	}
}

func TestHandleAPIError_NetworkError(t *testing.T) {
	diags := diag.Diagnostics{}
	var err error = &net.OpError{Op: "dial", Net: "tcp"}
	resp := &http.Response{
		Status: "500 Internal Server Error",
		Body:   io.NopCloser(bytes.NewReader([]byte(""))),
	}

	HandleAPIError("Test Error", &err, resp, &diags)
	if !diags.HasError() {
		t.Fatal("HandleAPIError should add an error to diagnostics")
	}
}

func TestHandleAPIError_APIError(t *testing.T) {
	diags := diag.Diagnostics{}
	var err error
	resp := &http.Response{
		Status: "400 Bad Request",
		Body:   io.NopCloser(bytes.NewReader([]byte("Invalid request"))),
	}

	HandleAPIError("Test Error", &err, resp, &diags)
	if !diags.HasError() {
		t.Fatal("HandleAPIError should add an error to diagnostics")
	}
}

func TestHandleAPIError_NoResponse(t *testing.T) {
	diags := diag.Diagnostics{}
	var err error

	HandleAPIError("Test Error", &err, nil, &diags)
	if !diags.HasError() {
		t.Fatal("HandleAPIError should add an error to diagnostics")
	}
}

func TestHandleAPIWarning(t *testing.T) {
	diags := diag.Diagnostics{}
	testErr := new(struct {
		msg string
	})
	var err error
	resp := &http.Response{
		Status: "200 OK",
		Body:   io.NopCloser(bytes.NewReader([]byte("Warning message"))),
	}

	HandleAPIWarning("Test Warning", &err, resp, &diags)
	// The function should have been called without panicking
	_ = testErr
}
