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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ========================================
// Resource Schema Functions
// ========================================

// ResourceRequiredInt32 returns a required int32 attribute
func ResourceRequiredInt32(description string) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description: description,
		required:    true,
	})
}

// ResourceOptionalInt32 returns an optional int32 attribute
func ResourceOptionalInt32(description string) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description: description,
		optional:    true,
	})
}

// ResourceComputedInt32 returns a computed int32 attribute
func ResourceComputedInt32(description string) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description: description,
		computed:    true,
	})
}

// ResourceComputedInt32WithDefault returns a computed int32 attribute with a default value
func ResourceComputedInt32WithDefault(description string, defaultValue int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:  description,
		computed:     true,
		defaultValue: int32default.StaticInt32(defaultValue),
	})
}

// ResourceOptionalInt32WithDefault returns an optional int32 attribute with a default value
func ResourceOptionalInt32WithDefault(description string, defaultValue int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: int32default.StaticInt32(defaultValue),
	})
}

// ResourceComputedOptionalInt32 returns a computed optional int32 attribute
func ResourceComputedOptionalInt32(description string) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description: description,
		optional:    true,
		computed:    true,
	})
}

// ResourceComputedOptionalInt32WithDefault returns a computed optional int32 attribute with a default value
func ResourceComputedOptionalInt32WithDefault(description string, defaultValue int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: int32default.StaticInt32(defaultValue),
	})
}

// ResourceRequiredInt32WithDefault returns a required int32 attribute with a default value
func ResourceRequiredInt32WithDefault(description string, defaultValue int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:  description,
		required:     true,
		defaultValue: int32default.StaticInt32(defaultValue),
	})
}

// ResourceOptionalInt32WithPlanModifier returns an optional int32 attribute with plan modifiers
func ResourceOptionalInt32WithPlanModifier(description string, planMods ...planmodifier.Int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:   description,
		optional:      true,
		planModifiers: planMods,
	})
}

// ResourceOptionalInt32WithDefaultAndPlanModifier returns an optional int32 attribute with default and plan modifiers
func ResourceOptionalInt32WithDefaultAndPlanModifier(description string, defaultValue int32, planMods ...planmodifier.Int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:   description,
		optional:      true,
		computed:      true,
		defaultValue:  int32default.StaticInt32(defaultValue),
		planModifiers: planMods,
	})
}

// ResourceComputedOptionalInt32WithPlanModifier returns a computed optional int32 attribute with plan modifiers
func ResourceComputedOptionalInt32WithPlanModifier(description string, planMods ...planmodifier.Int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:   description,
		optional:      true,
		computed:      true,
		planModifiers: planMods,
	})
}

// ResourceComputedOptionalInt32WithDefaultAndPlanModifier returns a computed optional int32 attribute with default and plan modifiers
func ResourceComputedOptionalInt32WithDefaultAndPlanModifier(description string, defaultValue int32, planMods ...planmodifier.Int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:   description,
		optional:      true,
		computed:      true,
		defaultValue:  int32default.StaticInt32(defaultValue),
		planModifiers: planMods,
	})
}

// ResourceOptionalInt32WithValidator returns an optional int32 attribute with validators
func ResourceOptionalInt32WithValidator(description string, validators ...validator.Int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description: description,
		optional:    true,
		validators:  validators,
	})
}

// ResourceRequiredInt32WithValidator returns a required int32 attribute with validators
func ResourceRequiredInt32WithValidator(description string, validators ...validator.Int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description: description,
		required:    true,
		validators:  validators,
	})
}

// ResourceOptionalInt32WithDefaultAndValidator returns an optional int32 attribute with default and validators
func ResourceOptionalInt32WithDefaultAndValidator(description string, defaultValue int32, validators ...validator.Int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: int32default.StaticInt32(defaultValue),
		validators:   validators,
	})
}

// ResourceComputedOptionalInt32WithDefaultAndValidator returns a computed optional int32 attribute with default and validators
func ResourceComputedOptionalInt32WithDefaultAndValidator(description string, defaultValue int32, validators ...validator.Int32) resourceschema.Int32Attribute {
	return newResourceInt32Attribute(int32AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: int32default.StaticInt32(defaultValue),
		validators:   validators,
	})
}

// ========================================
// Data Source Schema Functions
// ========================================

// DataSourceComputedInt32 returns a computed int32 attribute for data sources
func DataSourceComputedInt32(description string) datasourceschema.Int32Attribute {
	return newDataSourceInt32Attribute(int32AttributeConfig{
		description: description,
		computed:    true,
	})
}

// DataSourceOptionalInt32 returns an optional int32 attribute for data sources
func DataSourceOptionalInt32(description string) datasourceschema.Int32Attribute {
	return newDataSourceInt32Attribute(int32AttributeConfig{
		description: description,
		optional:    true,
	})
}
