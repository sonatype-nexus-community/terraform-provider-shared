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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// BaseProvider contains common provider configuration
type BaseProvider interface {
	GetAuth() interface{}
	GetBaseURL() string
	GetClient() interface{}
}

// BaseResourceConfig holds common configuration for all resources
type BaseResourceConfig struct {
	Auth    interface{}
	BaseURL string
	Client  interface{}
}

// BaseResource provides common functionality for all Terraform resources
type BaseResource struct {
	config *BaseResourceConfig
}

// NewBaseResource creates a new BaseResource with the given configuration
func NewBaseResource(config *BaseResourceConfig) *BaseResource {
	return &BaseResource{
		config: config,
	}
}

// Configure sets up the resource with provider configuration
func (r *BaseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	provider, ok := req.ProviderData.(BaseProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Type",
			fmt.Sprintf("Expected BaseProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.config = &BaseResourceConfig{
		Auth:    provider.GetAuth(),
		BaseURL: provider.GetBaseURL(),
		Client:  provider.GetClient(),
	}
}

// GetAuth returns the authentication configuration
func (r *BaseResource) GetAuth() interface{} {
	if r.config == nil {
		return nil
	}
	return r.config.Auth
}

// GetBaseURL returns the API base URL
func (r *BaseResource) GetBaseURL() string {
	if r.config == nil {
		return ""
	}
	return r.config.BaseURL
}

// GetClient returns the API client
func (r *BaseResource) GetClient() interface{} {
	if r.config == nil {
		return nil
	}
	return r.config.Client
}

// IsConfigured checks if the resource has been properly configured
func (r *BaseResource) IsConfigured() bool {
	return r.config != nil && r.config.Client != nil
}

// HasError is a helper to check if diagnostics contain errors
func HasError(diags *diag.Diagnostics) bool {
	return diags.HasError()
}

// AddErrorf adds a formatted error to diagnostics
func AddErrorf(diags *diag.Diagnostics, summary string, format string, args ...interface{}) {
	diags.AddError(summary, fmt.Sprintf(format, args...))
}

// AddErrorAndReturn adds an error to diagnostics and returns early
func AddErrorAndReturn(diags *diag.Diagnostics, summary string, details string) bool {
	diags.AddError(summary, details)
	return true
}
