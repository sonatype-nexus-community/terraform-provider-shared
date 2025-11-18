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
	if result == "" {
		t.Fatal("ParseAPIError(nil) should return a string")
	}

	// Test with HTTP response
	resp := &http.Response{
		Status: "400 Bad Request",
	}
	result = ParseAPIError(resp)
	if result != "400 Bad Request" {
		t.Fatalf("ParseAPIError failed: got %s", result)
	}
}
