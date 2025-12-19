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
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// TestAddForbiddenDiagnostic tests the AddForbiddenDiagnostic function
func TestAddForbiddenDiagnostic(t *testing.T) {
	diags := &diag.Diagnostics{}
	AddForbiddenDiagnostic(diags, "read resource")

	if !diags.HasError() {
		t.Fatal("AddForbiddenDiagnostic should add an error to diagnostics")
	}

	if len(diags.Errors()) != 1 {
		t.Fatalf("Expected 1 error, got %d", len(diags.Errors()))
	}

	err := diags.Errors()[0]
	if err.Summary() != "Forbidden read resource" {
		t.Fatalf("Expected error summary 'Forbidden read resource', got '%s'", err.Summary())
	}

	if err.Detail() == "" {
		t.Fatal("Expected error detail to be non-empty")
	}
}

// TestAddServerErrorDiagnostic tests the AddServerErrorDiagnostic function
func TestAddServerErrorDiagnostic(t *testing.T) {
	diags := &diag.Diagnostics{}
	AddServerErrorDiagnostic(diags, "Internal server error", http.StatusInternalServerError)

	if !diags.HasError() {
		t.Fatal("AddServerErrorDiagnostic should add an error to diagnostics")
	}

	if len(diags.Errors()) != 1 {
		t.Fatalf("Expected 1 error, got %d", len(diags.Errors()))
	}

	err := diags.Errors()[0]
	if err.Summary() != "Server Error (500)" {
		t.Fatalf("Expected error summary 'Server Error (500)', got '%s'", err.Summary())
	}

	if err.Detail() == "" {
		t.Fatal("Expected error detail to be non-empty")
	}
}

// TestAddClientErrorDiagnostic tests the AddClientErrorDiagnostic function
func TestAddClientErrorDiagnostic(t *testing.T) {
	diags := &diag.Diagnostics{}
	AddClientErrorDiagnostic(diags, "Bad request format", http.StatusBadRequest)

	if !diags.HasError() {
		t.Fatal("AddClientErrorDiagnostic should add an error to diagnostics")
	}

	if len(diags.Errors()) != 1 {
		t.Fatalf("Expected 1 error, got %d", len(diags.Errors()))
	}

	err := diags.Errors()[0]
	if err.Summary() != "Client Error (400)" {
		t.Fatalf("Expected error summary 'Client Error (400)', got '%s'", err.Summary())
	}

	if err.Detail() == "" {
		t.Fatal("Expected error detail to be non-empty")
	}
}

// TestAddServerErrorDiagnosticWith503 tests the AddServerErrorDiagnostic function with 503 status code
func TestAddServerErrorDiagnosticWith503(t *testing.T) {
	diags := &diag.Diagnostics{}
	AddServerErrorDiagnostic(diags, "Service unavailable", http.StatusServiceUnavailable)

	if !diags.HasError() {
		t.Fatal("AddServerErrorDiagnostic should add an error to diagnostics")
	}

	err := diags.Errors()[0]
	if err.Summary() != "Server Error (503)" {
		t.Fatalf("Expected error summary 'Server Error (503)', got '%s'", err.Summary())
	}
}

// TestAddClientErrorDiagnosticWith404 tests the AddClientErrorDiagnostic function with 404 status code
func TestAddClientErrorDiagnosticWith404(t *testing.T) {
	diags := &diag.Diagnostics{}
	AddClientErrorDiagnostic(diags, "Resource not found", http.StatusNotFound)

	if !diags.HasError() {
		t.Fatal("AddClientErrorDiagnostic should add an error to diagnostics")
	}

	err := diags.Errors()[0]
	if err.Summary() != "Client Error (404)" {
		t.Fatalf("Expected error summary 'Client Error (404)', got '%s'", err.Summary())
	}
}

// TestAddClientErrorDiagnosticWith409 tests the AddClientErrorDiagnostic function with 409 status code
func TestAddClientErrorDiagnosticWith409(t *testing.T) {
	diags := &diag.Diagnostics{}
	AddClientErrorDiagnostic(diags, "Resource already exists", http.StatusConflict)

	if !diags.HasError() {
		t.Fatal("AddClientErrorDiagnostic should add an error to diagnostics")
	}

	err := diags.Errors()[0]
	if err.Summary() != "Client Error (409)" {
		t.Fatalf("Expected error summary 'Client Error (409)', got '%s'", err.Summary())
	}
}

// TestStatusCodeChecks tests the status code checking functions
func TestIsNotFound(t *testing.T) {
	if !IsNotFound(http.StatusNotFound) {
		t.Fatal("IsNotFound(404) should return true")
	}
	if IsNotFound(http.StatusBadRequest) {
		t.Fatal("IsNotFound(400) should return false")
	}
}

func TestIsForbidden(t *testing.T) {
	if !IsForbidden(http.StatusForbidden) {
		t.Fatal("IsForbidden(403) should return true")
	}
	if IsForbidden(http.StatusUnauthorized) {
		t.Fatal("IsForbidden(401) should return false")
	}
}

func TestIsUnauthorized(t *testing.T) {
	if !IsUnauthorized(http.StatusUnauthorized) {
		t.Fatal("IsUnauthorized(401) should return true")
	}
	if IsUnauthorized(http.StatusForbidden) {
		t.Fatal("IsUnauthorized(403) should return false")
	}
}

func TestIsConflict(t *testing.T) {
	if !IsConflict(http.StatusConflict) {
		t.Fatal("IsConflict(409) should return true")
	}
	if IsConflict(http.StatusBadRequest) {
		t.Fatal("IsConflict(400) should return false")
	}
}

func TestIsClientError(t *testing.T) {
	if !IsClientError(http.StatusBadRequest) {
		t.Fatal("IsClientError(400) should return true")
	}
	if !IsClientError(http.StatusNotFound) {
		t.Fatal("IsClientError(404) should return true")
	}
	if !IsClientError(http.StatusForbidden) {
		t.Fatal("IsClientError(403) should return true")
	}
	if IsClientError(http.StatusOK) {
		t.Fatal("IsClientError(200) should return false")
	}
	if IsClientError(http.StatusInternalServerError) {
		t.Fatal("IsClientError(500) should return false")
	}
}

func TestIsServerError(t *testing.T) {
	if !IsServerError(http.StatusInternalServerError) {
		t.Fatal("IsServerError(500) should return true")
	}
	if !IsServerError(http.StatusServiceUnavailable) {
		t.Fatal("IsServerError(503) should return true")
	}
	if IsServerError(http.StatusOK) {
		t.Fatal("IsServerError(200) should return false")
	}
	if IsServerError(http.StatusBadRequest) {
		t.Fatal("IsServerError(400) should return false")
	}
}
