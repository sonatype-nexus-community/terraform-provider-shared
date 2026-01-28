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

package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type MockProvider struct {
	auth    interface{}
	baseURL string
	client  interface{}
}

func (m *MockProvider) GetAuth() interface{} {
	return m.auth
}

func (m *MockProvider) GetBaseURL() string {
	return m.baseURL
}

func (m *MockProvider) GetClient() interface{} {
	return m.client
}

func TestNewBaseResource(t *testing.T) {
	config := &BaseResourceConfig{
		Auth:    "test-auth",
		BaseURL: "http://example.com",
		Client:  "test-client",
	}

	res := NewBaseResource(config)
	if res == nil {
		t.Fatal("NewBaseResource should not return nil")
	}

	if res.GetAuth() != "test-auth" {
		t.Fatalf("GetAuth() = %v, expected 'test-auth'", res.GetAuth())
	}
	if res.GetBaseURL() != "http://example.com" {
		t.Fatalf("GetBaseURL() = %s, expected 'http://example.com'", res.GetBaseURL())
	}
	if res.GetClient() != "test-client" {
		t.Fatalf("GetClient() = %v, expected 'test-client'", res.GetClient())
	}
}

func TestConfigure_ValidProvider(t *testing.T) {
	config := &BaseResourceConfig{}
	res := &BaseResource{config: config}

	provider := &MockProvider{
		auth:    "configured-auth",
		baseURL: "http://configured.example.com",
		client:  "configured-client",
	}

	req := resource.ConfigureRequest{
		ProviderData: provider,
	}
	resp := &resource.ConfigureResponse{}

	res.Configure(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		if res.GetAuth() != "configured-auth" {
			t.Fatalf("GetAuth() = %v, expected 'configured-auth'", res.GetAuth())
		}
		if res.GetBaseURL() != "http://configured.example.com" {
			t.Fatalf("GetBaseURL() = %s, expected 'http://configured.example.com'", res.GetBaseURL())
		}
		if res.GetClient() != "configured-client" {
			t.Fatalf("GetClient() = %v, expected 'configured-client'", res.GetClient())
		}
	}
}

func TestConfigure_InvalidProvider(t *testing.T) {
	config := &BaseResourceConfig{}
	res := &BaseResource{config: config}

	req := resource.ConfigureRequest{
		ProviderData: "invalid-provider",
	}
	resp := &resource.ConfigureResponse{}

	res.Configure(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("Configure should add an error when provider data is invalid")
	}
}

func TestConfigure_NilProviderData(t *testing.T) {
	config := &BaseResourceConfig{}
	res := &BaseResource{config: config}

	req := resource.ConfigureRequest{
		ProviderData: nil,
	}
	resp := &resource.ConfigureResponse{}

	res.Configure(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatal("Configure should not add an error when provider data is nil")
	}
}

func TestGetAuth(t *testing.T) {
	// Test with config
	config := &BaseResourceConfig{Auth: "test-auth"}
	resource := &BaseResource{config: config}

	if resource.GetAuth() != "test-auth" {
		t.Fatalf("GetAuth() = %v, expected 'test-auth'", resource.GetAuth())
	}

	// Test with nil config
	resource = &BaseResource{config: nil}
	if resource.GetAuth() != nil {
		t.Fatalf("GetAuth() = %v, expected nil", resource.GetAuth())
	}
}

func TestGetBaseURL(t *testing.T) {
	// Test with config
	config := &BaseResourceConfig{BaseURL: "http://example.com"}
	resource := &BaseResource{config: config}

	if resource.GetBaseURL() != "http://example.com" {
		t.Fatalf("GetBaseURL() = %s, expected 'http://example.com'", resource.GetBaseURL())
	}

	// Test with nil config
	resource = &BaseResource{config: nil}
	if resource.GetBaseURL() != "" {
		t.Fatalf("GetBaseURL() = %s, expected ''", resource.GetBaseURL())
	}
}

func TestGetClient(t *testing.T) {
	// Test with config
	config := &BaseResourceConfig{Client: "test-client"}
	resource := &BaseResource{config: config}

	if resource.GetClient() != "test-client" {
		t.Fatalf("GetClient() = %v, expected 'test-client'", resource.GetClient())
	}

	// Test with nil config
	resource = &BaseResource{config: nil}
	if resource.GetClient() != nil {
		t.Fatalf("GetClient() = %v, expected nil", resource.GetClient())
	}
}

func TestIsConfigured(t *testing.T) {
	// Test with configured resource
	config := &BaseResourceConfig{
		Auth:    "auth",
		BaseURL: "http://example.com",
		Client:  "client",
	}
	resource := &BaseResource{config: config}

	if !resource.IsConfigured() {
		t.Fatal("IsConfigured() should return true when config and client are set")
	}

	// Test with nil config
	resource = &BaseResource{config: nil}
	if resource.IsConfigured() {
		t.Fatal("IsConfigured() should return false when config is nil")
	}

	// Test with config but nil client
	config = &BaseResourceConfig{
		Auth:    "auth",
		BaseURL: "http://example.com",
		Client:  nil,
	}
	resource = &BaseResource{config: config}

	if resource.IsConfigured() {
		t.Fatal("IsConfigured() should return false when client is nil")
	}
}

func TestHasError(t *testing.T) {
	// Test with no errors
	diags := diag.Diagnostics{}
	if HasError(&diags) {
		t.Fatal("HasError should return false when there are no errors")
	}

	// Test with errors
	diags.AddError("Test Error", "This is a test error")
	if !HasError(&diags) {
		t.Fatal("HasError should return true when there are errors")
	}
}

func TestAddErrorf(t *testing.T) {
	diags := diag.Diagnostics{}
	AddErrorf(&diags, "Error Summary", "This is a %s error", "formatted")

	if !diags.HasError() {
		t.Fatal("AddErrorf should add an error to diagnostics")
	}

	errors := diags.Errors()
	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, got %d", len(errors))
	}

	if errors[0].Detail() != "This is a formatted error" {
		t.Fatalf("Error detail = %s, expected 'This is a formatted error'", errors[0].Detail())
	}
}

func TestAddErrorAndReturn(t *testing.T) {
	diags := diag.Diagnostics{}
	result := AddErrorAndReturn(&diags, "Error Summary", "Error Details")

	if !result {
		t.Fatal("AddErrorAndReturn should return true")
	}

	if !diags.HasError() {
		t.Fatal("AddErrorAndReturn should add an error to diagnostics")
	}

	errors := diags.Errors()
	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, got %d", len(errors))
	}
}
