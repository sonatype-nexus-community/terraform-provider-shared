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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// ========================================
// Resource Schema Functions
// ========================================

// ResourceRequiredInt64 returns a required int64 attribute
func ResourceRequiredInt64(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		required:    true,
	})
}

// ResourceOptionalInt64 returns an optional int64 attribute
func ResourceOptionalInt64(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		optional:    true,
	})
}

// ResourceComputedInt64 returns a computed int64 attribute
func ResourceComputedInt64(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		computed:    true,
	})
}

// ResourceComputedInt64WithDefault returns a computed int64 attribute with a default value
func ResourceComputedInt64WithDefault(description string, defaultValue int64) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:  description,
		computed:     true,
		defaultValue: int64default.StaticInt64(defaultValue),
	})
}

// ResourceOptionalInt64WithDefault returns an optional int64 attribute with a default value
func ResourceOptionalInt64WithDefault(description string, defaultValue int64) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: int64default.StaticInt64(defaultValue),
	})
}

// ResourceComputedOptionalInt64 returns a computed optional int64 attribute
func ResourceComputedOptionalInt64(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		optional:    true,
		computed:    true,
	})
}

// ResourceRequiredInt64WithDefault returns a required int64 attribute with a default value
func ResourceRequiredInt64WithDefault(description string, defaultValue int64) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:  description,
		required:     true,
		defaultValue: int64default.StaticInt64(defaultValue),
	})
}

// ResourceComputedOptionalInt64WithDefault returns a computed optional int64 attribute with a default value
func ResourceComputedOptionalInt64WithDefault(description string, defaultValue int64) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: int64default.StaticInt64(defaultValue),
	})
}

// ResourceOptionalInt64WithDefaultAndPlanModifier returns an optional int64 attribute with default and plan modifiers
func ResourceOptionalInt64WithDefaultAndPlanModifier(description string, defaultValue int64, planMods ...planmodifier.Int64) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:   description,
		optional:      true,
		computed:      true,
		defaultValue:  int64default.StaticInt64(defaultValue),
		planModifiers: planMods,
	})
}

// ResourceComputedInt64WithDefaultAndPlanModifier returns a computed int64 attribute with default and plan modifiers
func ResourceComputedInt64WithDefaultAndPlanModifier(description string, defaultValue int64, planMods ...planmodifier.Int64) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:   description,
		computed:      true,
		defaultValue:  int64default.StaticInt64(defaultValue),
		planModifiers: planMods,
	})
}

// ResourceComputedOptionalInt64WithDefaultAndPlanModifier returns a computed optional int64 attribute with default and plan modifiers
func ResourceComputedOptionalInt64WithDefaultAndPlanModifier(description string, defaultValue int64, planMods ...planmodifier.Int64) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:   description,
		optional:      true,
		computed:      true,
		defaultValue:  int64default.StaticInt64(defaultValue),
		planModifiers: planMods,
	})
}

// NOTE: Int32 values have dedicated helpers in int32_attributes.go.
// Terraform's plugin framework uses Int64Type for all integer attributes,
// so int32 values are wrapped as int64 internally.

// ResourceOptionalPort returns an optional int64 attribute for network ports (0-65535)
func ResourceOptionalPort(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		optional:    true,
	})
}

// ResourceRequiredPort returns a required int64 attribute for network ports (0-65535)
func ResourceRequiredPort(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		required:    true,
	})
}

// ResourcePortWithDefault returns a port attribute with a default value
func ResourcePortWithDefault(description string, defaultValue int64) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: int64default.StaticInt64(defaultValue),
	})
}

// ResourcePercentageInt returns a percentage as int64 (0-100)
func ResourcePercentageInt(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		optional:    true,
	})
}

// ResourceRequiredPercentageInt returns a required percentage as int64 (0-100)
func ResourceRequiredPercentageInt(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		required:    true,
	})
}

// ResourceDurationInt returns a duration in seconds as int64
func ResourceDurationInt(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		optional:    true,
	})
}

// ResourceRequiredDurationInt returns a required duration in seconds as int64
func ResourceRequiredDurationInt(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description: description,
		required:    true,
	})
}

// ResourceTimeoutInt returns a timeout duration in seconds as int64
func ResourceTimeoutInt(description string, defaultSeconds int64) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: int64default.StaticInt64(defaultSeconds),
	})
}

// ResourceCountInt returns a count attribute as int64
func ResourceCountInt(description string) resourceschema.Int64Attribute {
	return newResourceInt64Attribute(int64AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: int64default.StaticInt64(0),
	})
}

// ========================================
// Data Source Schema Functions
// ========================================

// DataSourceComputedInt64 returns a computed int64 attribute for data sources
func DataSourceComputedInt64(description string) datasourceschema.Int64Attribute {
	return newDataSourceInt64Attribute(int64AttributeConfig{
		description: description,
		computed:    true,
	})
}

// DataSourceOptionalInt64 returns an optional int64 attribute for data sources
func DataSourceOptionalInt64(description string) datasourceschema.Int64Attribute {
	return newDataSourceInt64Attribute(int64AttributeConfig{
		description: description,
		optional:    true,
	})
}
