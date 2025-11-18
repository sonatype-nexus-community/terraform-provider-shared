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
	return resourceschema.BoolAttribute{
		Description: description,
		Required:    true,
	}
}

// ResourceOptionalBool returns an optional boolean attribute
func ResourceOptionalBool(description string) resourceschema.BoolAttribute {
	return resourceschema.BoolAttribute{
		Description: description,
		Optional:    true,
	}
}

// ResourceComputedBool returns a computed boolean attribute
func ResourceComputedBool(description string) resourceschema.BoolAttribute {
	return resourceschema.BoolAttribute{
		Description: description,
		Computed:    true,
	}
}

// ResourceOptionalBoolWithDefault returns an optional boolean attribute with a default value
func ResourceOptionalBoolWithDefault(description string, defaultValue bool) resourceschema.BoolAttribute {
	return resourceschema.BoolAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     booldefault.StaticBool(defaultValue),
	}
}

// ResourceComputedOptionalBool returns a computed optional boolean (persists state for unknown)
func ResourceComputedOptionalBool(description string) resourceschema.BoolAttribute {
	return resourceschema.BoolAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	}
}

// ResourceRequiredBoolWithDefault returns a required boolean attribute with a default value
func ResourceRequiredBoolWithDefault(description string, defaultValue bool) resourceschema.BoolAttribute {
	return resourceschema.BoolAttribute{
		Description: description,
		Required:    true,
		Default:     booldefault.StaticBool(defaultValue),
	}
}

// ResourceComputedBoolWithDefault returns a computed boolean attribute with a default value
func ResourceComputedBoolWithDefault(description string, defaultValue bool) resourceschema.BoolAttribute {
	return resourceschema.BoolAttribute{
		Description: description,
		Computed:    true,
		Default:     booldefault.StaticBool(defaultValue),
	}
}

// ========================================
// Data Source Schema Functions
// ========================================

// DataSourceComputedBool returns a computed boolean attribute for data sources
func DataSourceComputedBool(description string) datasourceschema.BoolAttribute {
	return datasourceschema.BoolAttribute{
		Description: description,
		Computed:    true,
	}
}

// DataSourceOptionalBool returns an optional boolean attribute for data sources
func DataSourceOptionalBool(description string) datasourceschema.BoolAttribute {
	return datasourceschema.BoolAttribute{
		Description: description,
		Optional:    true,
	}
}
