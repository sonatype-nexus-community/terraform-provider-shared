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

package schema

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ========================================
// String Attribute Configuration
// ========================================

// stringAttributeConfig holds configuration for string attribute builders
type stringAttributeConfig struct {
	description   string
	required      bool
	optional      bool
	computed      bool
	sensitive     bool
	defaultValue  defaults.String
	validators    []validator.String
	planModifiers []planmodifier.String
}

// newResourceStringAttribute creates a resource string attribute from config
func newResourceStringAttribute(config stringAttributeConfig) resourceschema.StringAttribute {
	attr := resourceschema.StringAttribute{
		MarkdownDescription: config.description,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
		Sensitive:           config.sensitive,
	}
	if config.defaultValue != nil {
		attr.Default = config.defaultValue
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	if len(config.planModifiers) > 0 {
		attr.PlanModifiers = config.planModifiers
	}
	return attr
}

// newDataSourceStringAttribute creates a datasource string attribute from config
func newDataSourceStringAttribute(config stringAttributeConfig) datasourceschema.StringAttribute {
	attr := datasourceschema.StringAttribute{
		MarkdownDescription: config.description,
		Optional:            config.optional,
		Computed:            config.computed,
		Sensitive:           config.sensitive,
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	return attr
}

// ========================================
// Bool Attribute Configuration
// ========================================

// boolAttributeConfig holds configuration for bool attribute builders
type boolAttributeConfig struct {
	description   string
	required      bool
	optional      bool
	computed      bool
	defaultValue  defaults.Bool
	planModifiers []planmodifier.Bool
}

// newResourceBoolAttribute creates a resource bool attribute from config
func newResourceBoolAttribute(config boolAttributeConfig) resourceschema.BoolAttribute {
	attr := resourceschema.BoolAttribute{
		MarkdownDescription: config.description,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
	if config.defaultValue != nil {
		attr.Default = config.defaultValue
	}
	if len(config.planModifiers) > 0 {
		attr.PlanModifiers = config.planModifiers
	}
	return attr
}

// newDataSourceBoolAttribute creates a datasource bool attribute from config
func newDataSourceBoolAttribute(config boolAttributeConfig) datasourceschema.BoolAttribute {
	return datasourceschema.BoolAttribute{
		MarkdownDescription: config.description,
		Optional:            config.optional,
		Computed:            config.computed,
	}
}

// ========================================
// Int64 Attribute Configuration
// ========================================

// int64AttributeConfig holds configuration for int64 attribute builders
type int64AttributeConfig struct {
	description  string
	required     bool
	optional     bool
	computed     bool
	defaultValue defaults.Int64
	validators   []validator.Int64
}

// newResourceInt64Attribute creates a resource int64 attribute from config
func newResourceInt64Attribute(config int64AttributeConfig) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: config.description,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
	if config.defaultValue != nil {
		attr.Default = config.defaultValue
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	return attr
}

// newDataSourceInt64Attribute creates a datasource int64 attribute from config
func newDataSourceInt64Attribute(config int64AttributeConfig) datasourceschema.Int64Attribute {
	attr := datasourceschema.Int64Attribute{
		MarkdownDescription: config.description,
		Optional:            config.optional,
		Computed:            config.computed,
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	return attr
}

// ========================================
// Int32 Attribute Configuration
// ========================================

// int32AttributeConfig holds configuration for int32 attribute builders
// Note: int32 attributes use Int64Attribute internally since Terraform doesn't have a native Int32Type
type int32AttributeConfig struct {
	description  string
	required     bool
	optional     bool
	computed     bool
	defaultValue defaults.Int64
	validators   []validator.Int64
}

// newResourceInt32Attribute creates a resource int32 attribute from config
func newResourceInt32Attribute(config int32AttributeConfig) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: config.description,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
	if config.defaultValue != nil {
		attr.Default = config.defaultValue
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	return attr
}

// newDataSourceInt32Attribute creates a datasource int32 attribute from config
func newDataSourceInt32Attribute(config int32AttributeConfig) datasourceschema.Int64Attribute {
	attr := datasourceschema.Int64Attribute{
		MarkdownDescription: config.description,
		Optional:            config.optional,
		Computed:            config.computed,
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	return attr
}

// ========================================
// Float64 Attribute Configuration
// ========================================

// float64AttributeConfig holds configuration for float64 attribute builders
type float64AttributeConfig struct {
	description  string
	required     bool
	optional     bool
	computed     bool
	defaultValue defaults.Float64
	validators   []validator.Float64
}

// newResourceFloat64Attribute creates a resource float64 attribute from config
func newResourceFloat64Attribute(config float64AttributeConfig) resourceschema.Float64Attribute {
	attr := resourceschema.Float64Attribute{
		MarkdownDescription: config.description,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
	if config.defaultValue != nil {
		attr.Default = config.defaultValue
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	return attr
}

// newDataSourceFloat64Attribute creates a datasource float64 attribute from config
func newDataSourceFloat64Attribute(config float64AttributeConfig) datasourceschema.Float64Attribute {
	attr := datasourceschema.Float64Attribute{
		MarkdownDescription: config.description,
		Optional:            config.optional,
		Computed:            config.computed,
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	return attr
}
