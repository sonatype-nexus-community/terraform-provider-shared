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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// ========================================
// Resource Schema Functions
// ========================================

// ResourceRequiredBool returns a required boolean attribute
func ResourceRequiredBool(description string) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description: description,
		required:    true,
	})
}

// ResourceOptionalBool returns an optional boolean attribute
func ResourceOptionalBool(description string) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description: description,
		optional:    true,
	})
}

// ResourceComputedBool returns a computed boolean attribute
func ResourceComputedBool(description string) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description: description,
		computed:    true,
	})
}

// ResourceOptionalBoolWithDefault returns an optional boolean attribute with a default value
func ResourceOptionalBoolWithDefault(description string, defaultValue bool) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: booldefault.StaticBool(defaultValue),
	})
}

// ResourceComputedOptionalBool returns a computed optional boolean (persists state for unknown)
func ResourceComputedOptionalBool(description string) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description:   description,
		optional:      true,
		computed:      true,
		planModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	})
}

// ResourceRequiredBoolWithDefault returns a required boolean attribute with a default value
func ResourceRequiredBoolWithDefault(description string, defaultValue bool) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description:  description,
		required:     true,
		defaultValue: booldefault.StaticBool(defaultValue),
	})
}

// ResourceComputedBoolWithDefault returns a computed boolean attribute with a default value
func ResourceComputedBoolWithDefault(description string, defaultValue bool) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description:  description,
		computed:     true,
		defaultValue: booldefault.StaticBool(defaultValue),
	})
}

// ResourceComputedOptionalBoolWithDefault returns a computed optional boolean with a default value
// (combines optional+computed flags with a static default)
func ResourceComputedOptionalBoolWithDefault(description string, defaultValue bool) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: booldefault.StaticBool(defaultValue),
	})
}

// ResourceOptionalBoolWithPlanModifier returns an optional boolean attribute with plan modifiers
// (useful for attributes that need custom plan modification logic)
func ResourceOptionalBoolWithPlanModifier(description string, planMods ...planmodifier.Bool) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description:   description,
		optional:      true,
		planModifiers: planMods,
	})
}

// ResourceOptionalBoolWithDefaultAndPlanModifier returns an optional boolean attribute with default and plan modifiers
func ResourceOptionalBoolWithDefaultAndPlanModifier(description string, defaultValue bool, planMods ...planmodifier.Bool) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description:   description,
		optional:      true,
		computed:      true,
		defaultValue:  booldefault.StaticBool(defaultValue),
		planModifiers: planMods,
	})
}

// ResourceComputedOptionalBoolWithPlanModifier returns a computed optional boolean with plan modifiers
func ResourceComputedOptionalBoolWithPlanModifier(description string, planMods ...planmodifier.Bool) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description:   description,
		optional:      true,
		computed:      true,
		planModifiers: planMods,
	})
}

// ResourceComputedOptionalBoolWithDefaultAndPlanModifier returns a computed optional boolean with default and plan modifiers
func ResourceComputedOptionalBoolWithDefaultAndPlanModifier(description string, defaultValue bool, planMods ...planmodifier.Bool) resourceschema.BoolAttribute {
	return newResourceBoolAttribute(boolAttributeConfig{
		description:   description,
		optional:      true,
		computed:      true,
		defaultValue:  booldefault.StaticBool(defaultValue),
		planModifiers: planMods,
	})
}

// ========================================
// Data Source Schema Functions
// ========================================

// DataSourceComputedBool returns a computed boolean attribute for data sources
func DataSourceComputedBool(description string) datasourceschema.BoolAttribute {
	return newDataSourceBoolAttribute(boolAttributeConfig{
		description: description,
		computed:    true,
	})
}

// DataSourceOptionalBool returns an optional boolean attribute for data sources
func DataSourceOptionalBool(description string) datasourceschema.BoolAttribute {
	return newDataSourceBoolAttribute(boolAttributeConfig{
		description: description,
		optional:    true,
	})
}
