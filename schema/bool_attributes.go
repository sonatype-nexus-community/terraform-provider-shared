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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// RequiredBool returns a required boolean attribute
func RequiredBool(description string) schema.BoolAttribute {
	return schema.BoolAttribute{
		Description: description,
		Required:    true,
	}
}

// OptionalBool returns an optional boolean attribute
func OptionalBool(description string) schema.BoolAttribute {
	return schema.BoolAttribute{
		Description: description,
		Optional:    true,
	}
}

// ComputedBool returns a computed boolean attribute
func ComputedBool(description string) schema.BoolAttribute {
	return schema.BoolAttribute{
		Description: description,
		Computed:    true,
	}
}

// OptionalBoolWithDefault returns an optional boolean attribute with a default value
func OptionalBoolWithDefault(description string, defaultValue bool) schema.BoolAttribute {
	return schema.BoolAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     booldefault.StaticBool(defaultValue),
	}
}

// ComputedOptionalBool returns a computed optional boolean (persists state for unknown)
func ComputedOptionalBool(description string) schema.BoolAttribute {
	return schema.BoolAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	}
}

// RequiredBoolWithDefault returns a required boolean attribute with a default value
func RequiredBoolWithDefault(description string, defaultValue bool) schema.BoolAttribute {
	return schema.BoolAttribute{
		Description: description,
		Required:    true,
		Default:     booldefault.StaticBool(defaultValue),
	}
}

// ComputedBoolWithDefault returns a computed boolean attribute with a default value
func ComputedBoolWithDefault(description string, defaultValue bool) schema.BoolAttribute {
	return schema.BoolAttribute{
		Description: description,
		Computed:    true,
		Default:     booldefault.StaticBool(defaultValue),
	}
}
