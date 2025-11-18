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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ========================================
// Resource Schema Functions
// ========================================

// ResourceRequiredStringMap returns a required map attribute with string values
func ResourceRequiredStringMap(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Required:    true,
	}
}

// ResourceOptionalStringMap returns an optional map attribute with string values
func ResourceOptionalStringMap(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
	}
}

// ResourceComputedStringMap returns a computed map attribute with string values
func ResourceComputedStringMap(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Computed:    true,
	}
}

// ResourceComputedOptionalStringMap returns a computed optional map attribute with string values
func ResourceComputedOptionalStringMap(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
	}
}

// ResourceRequiredInt64Map returns a required map attribute with int64 values
func ResourceRequiredInt64Map(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Required:    true,
	}
}

// ResourceOptionalInt64Map returns an optional map attribute with int64 values
func ResourceOptionalInt64Map(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Optional:    true,
	}
}

// ResourceComputedInt64Map returns a computed map attribute with int64 values
func ResourceComputedInt64Map(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Computed:    true,
	}
}

// ResourceRequiredBoolMap returns a required map attribute with bool values
func ResourceRequiredBoolMap(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.BoolType,
		Required:    true,
	}
}

// ResourceOptionalBoolMap returns an optional map attribute with bool values
func ResourceOptionalBoolMap(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.BoolType,
		Optional:    true,
	}
}

// ResourceComputedBoolMap returns a computed map attribute with bool values
func ResourceComputedBoolMap(description string) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		Description: description,
		ElementType: types.BoolType,
		Computed:    true,
	}
}

// ========================================
// Data Source Schema Functions
// ========================================

// DataSourceOptionalStringMap returns an optional map attribute with string values for data sources
func DataSourceOptionalStringMap(description string) datasourceschema.MapAttribute {
	return datasourceschema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
	}
}

// DataSourceComputedStringMap returns a computed map attribute with string values for data sources
func DataSourceComputedStringMap(description string) datasourceschema.MapAttribute {
	return datasourceschema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Computed:    true,
	}
}

// DataSourceOptionalInt64Map returns an optional map attribute with int64 values for data sources
func DataSourceOptionalInt64Map(description string) datasourceschema.MapAttribute {
	return datasourceschema.MapAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Optional:    true,
	}
}

// DataSourceComputedInt64Map returns a computed map attribute with int64 values for data sources
func DataSourceComputedInt64Map(description string) datasourceschema.MapAttribute {
	return datasourceschema.MapAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Computed:    true,
	}
}

// DataSourceOptionalBoolMap returns an optional map attribute with bool values for data sources
func DataSourceOptionalBoolMap(description string) datasourceschema.MapAttribute {
	return datasourceschema.MapAttribute{
		Description: description,
		ElementType: types.BoolType,
		Optional:    true,
	}
}

// DataSourceComputedBoolMap returns a computed map attribute with bool values for data sources
func DataSourceComputedBoolMap(description string) datasourceschema.MapAttribute {
	return datasourceschema.MapAttribute{
		Description: description,
		ElementType: types.BoolType,
		Computed:    true,
	}
}
