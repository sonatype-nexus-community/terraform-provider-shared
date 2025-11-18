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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// StandardID returns a standard computed ID attribute
func StandardID() schema.StringAttribute {
	return schema.StringAttribute{
		Description: "Internal ID of the resource",
		Computed:    true,
	}
}

// Timestamp returns a standard computed timestamp attribute for tracking last updates
func Timestamp() schema.StringAttribute {
	return schema.StringAttribute{
		Description: "String representation of the date/time the resource was last changed",
		Computed:    true,
	}
}

// RequiredString returns a required string attribute with the given description
func RequiredString(description string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
	}
}

// OptionalString returns an optional string attribute with the given description
func OptionalString(description string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
	}
}

// ComputedString returns a computed string attribute with the given description
func ComputedString(description string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Computed:    true,
	}
}

// ComputedStringWithDefault returns a computed string attribute with a default value
func ComputedStringWithDefault(description string, defaultValue string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Computed:    true,
		Default:     stringdefault.StaticString(defaultValue),
	}
}

// SensitiveString returns an optional sensitive string attribute with the given description
func SensitiveString(description string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Sensitive:   true,
	}
}

// SensitiveRequiredString returns a required sensitive string attribute
func SensitiveRequiredString(description string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
		Sensitive:   true,
	}
}

// ComputedSensitiveString returns a computed sensitive string attribute
func ComputedSensitiveString(description string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Computed:    true,
		Sensitive:   true,
	}
}

// StringWithDefault returns an optional string attribute with a default value
func StringWithDefault(description string, defaultValue string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString(defaultValue),
	}
}

// StringEnum returns a string attribute with enum validation
func StringEnum(description string, enumValues ...string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Validators: []validator.String{
			stringvalidator.OneOf(enumValues...),
		},
	}
}

// RequiredStringEnum returns a required string attribute with enum validation
func RequiredStringEnum(description string, enumValues ...string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
		Validators: []validator.String{
			stringvalidator.OneOf(enumValues...),
		},
	}
}

// StringEnumWithDefault returns an optional string attribute with enum validation and a default value
func StringEnumWithDefault(description string, defaultValue string, enumValues ...string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString(defaultValue),
		Validators: []validator.String{
			stringvalidator.OneOf(enumValues...),
		},
	}
}

// ComputedOptionalString returns a computed optional string (persists state for unknown)
func ComputedOptionalString(description string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
}

// ResourceIDAttribute returns a string attribute commonly used for referencing other resources
func ResourceIDAttribute(description string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
	}
}
