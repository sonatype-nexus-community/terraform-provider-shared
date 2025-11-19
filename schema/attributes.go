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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ========================================
// Resource Schema Functions
// ========================================

// ResourceStandardID returns a standard computed ID attribute
func ResourceStandardID() resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: "Internal ID of the resource",
		Computed:            true,
	}
}

// ResourceLastUpdated returns a standard computed last updated attribute for tracking last updates
func ResourceLastUpdated() resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: "String representation of the date/time the resource was last changed",
		Computed:            true,
	}
}

// ResourceRequiredString returns a required string attribute with the given description
func ResourceRequiredString(description string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		required:    true,
	})
}

// ResourceOptionalString returns an optional string attribute with the given description
func ResourceOptionalString(description string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		optional:    true,
	})
}

// ResourceComputedString returns a computed string attribute with the given description
func ResourceComputedString(description string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		computed:    true,
	})
}

// ResourceComputedStringWithDefault returns a computed string attribute with a default value
func ResourceComputedStringWithDefault(description string, defaultValue string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description:  description,
		computed:     true,
		defaultValue: stringdefault.StaticString(defaultValue),
	})
}

// ResourceSensitiveString returns an optional sensitive string attribute with the given description
func ResourceSensitiveString(description string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		optional:    true,
		sensitive:   true,
	})
}

// ResourceSensitiveRequiredString returns a required sensitive string attribute
func ResourceSensitiveRequiredString(description string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		required:    true,
		sensitive:   true,
	})
}

// ResourceComputedSensitiveString returns a computed sensitive string attribute
func ResourceComputedSensitiveString(description string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		computed:    true,
		sensitive:   true,
	})
}

// ResourceStringWithDefault returns an optional string attribute with a default value
func ResourceStringWithDefault(description string, defaultValue string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: stringdefault.StaticString(defaultValue),
	})
}

// ResourceStringEnum returns a string attribute with enum validation
func ResourceStringEnum(description string, enumValues ...string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		optional:    true,
		validators: []validator.String{
			stringvalidator.OneOf(enumValues...),
		},
	})
}

// ResourceRequiredStringEnum returns a required string attribute with enum validation
func ResourceRequiredStringEnum(description string, enumValues ...string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		required:    true,
		validators: []validator.String{
			stringvalidator.OneOf(enumValues...),
		},
	})
}

// ResourceStringEnumWithDefault returns an optional string attribute with enum validation and a default value
func ResourceStringEnumWithDefault(description string, defaultValue string, enumValues ...string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: stringdefault.StaticString(defaultValue),
		validators: []validator.String{
			stringvalidator.OneOf(enumValues...),
		},
	})
}

// ResourceComputedOptionalString returns a computed optional string (persists state for unknown)
func ResourceComputedOptionalString(description string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		optional:    true,
		computed:    true,
		planModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	})
}

// ResourceIDAttribute returns a string attribute commonly used for referencing other resources
func ResourceIDAttribute(description string) resourceschema.StringAttribute {
	return newResourceStringAttribute(stringAttributeConfig{
		description: description,
		required:    true,
		planModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
	})
}

// ========================================
// Data Source Schema Functions
// ========================================

// DataSourceComputedString returns a computed string attribute for data sources
func DataSourceComputedString(description string) datasourceschema.StringAttribute {
	return newDataSourceStringAttribute(stringAttributeConfig{
		description: description,
		computed:    true,
	})
}

// DataSourceOptionalString returns an optional string attribute for data sources
func DataSourceOptionalString(description string) datasourceschema.StringAttribute {
	return newDataSourceStringAttribute(stringAttributeConfig{
		description: description,
		optional:    true,
	})
}

// DataSourceComputedSensitiveString returns a computed sensitive string attribute for data sources
func DataSourceComputedSensitiveString(description string) datasourceschema.StringAttribute {
	return newDataSourceStringAttribute(stringAttributeConfig{
		description: description,
		computed:    true,
		sensitive:   true,
	})
}

// DataSourceSensitiveString returns an optional sensitive string attribute for data sources
func DataSourceSensitiveString(description string) datasourceschema.StringAttribute {
	return newDataSourceStringAttribute(stringAttributeConfig{
		description: description,
		optional:    true,
		sensitive:   true,
	})
}

// DataSourceStringEnum returns a string attribute with enum validation for data sources
func DataSourceStringEnum(description string, enumValues ...string) datasourceschema.StringAttribute {
	return newDataSourceStringAttribute(stringAttributeConfig{
		description: description,
		optional:    true,
		validators: []validator.String{
			stringvalidator.OneOf(enumValues...),
		},
	})
}
