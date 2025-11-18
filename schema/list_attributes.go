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

// ResourceRequiredStringList returns a required list attribute with string elements
func ResourceRequiredStringList(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Required:    true,
	}
}

// ResourceOptionalStringList returns an optional list attribute with string elements
func ResourceOptionalStringList(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
	}
}

// ResourceComputedStringList returns a computed list attribute with string elements
func ResourceComputedStringList(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Computed:    true,
	}
}

// ResourceComputedOptionalStringList returns a computed optional list attribute with string elements
func ResourceComputedOptionalStringList(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
	}
}

// ResourceRequiredInt64List returns a required list attribute with int64 elements
func ResourceRequiredInt64List(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Required:    true,
	}
}

// ResourceOptionalInt64List returns an optional list attribute with int64 elements
func ResourceOptionalInt64List(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Optional:    true,
	}
}

// ResourceComputedInt64List returns a computed list attribute with int64 elements
func ResourceComputedInt64List(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Computed:    true,
	}
}

// ResourceRequiredBoolList returns a required list attribute with bool elements
func ResourceRequiredBoolList(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.BoolType,
		Required:    true,
	}
}

// ResourceOptionalBoolList returns an optional list attribute with bool elements
func ResourceOptionalBoolList(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.BoolType,
		Optional:    true,
	}
}

// ResourceComputedBoolList returns a computed list attribute with bool elements
func ResourceComputedBoolList(description string) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		Description: description,
		ElementType: types.BoolType,
		Computed:    true,
	}
}

// ========================================
// Data Source Schema Functions
// ========================================

// DataSourceOptionalStringList returns an optional list attribute with string elements for data sources
func DataSourceOptionalStringList(description string) datasourceschema.ListAttribute {
	return datasourceschema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
	}
}

// DataSourceComputedStringList returns a computed list attribute with string elements for data sources
func DataSourceComputedStringList(description string) datasourceschema.ListAttribute {
	return datasourceschema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Computed:    true,
	}
}

// DataSourceOptionalInt64List returns an optional list attribute with int64 elements for data sources
func DataSourceOptionalInt64List(description string) datasourceschema.ListAttribute {
	return datasourceschema.ListAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Optional:    true,
	}
}

// DataSourceComputedInt64List returns a computed list attribute with int64 elements for data sources
func DataSourceComputedInt64List(description string) datasourceschema.ListAttribute {
	return datasourceschema.ListAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Computed:    true,
	}
}

// DataSourceOptionalBoolList returns an optional list attribute with bool elements for data sources
func DataSourceOptionalBoolList(description string) datasourceschema.ListAttribute {
	return datasourceschema.ListAttribute{
		Description: description,
		ElementType: types.BoolType,
		Optional:    true,
	}
}

// DataSourceComputedBoolList returns a computed list attribute with bool elements for data sources
func DataSourceComputedBoolList(description string) datasourceschema.ListAttribute {
	return datasourceschema.ListAttribute{
		Description: description,
		ElementType: types.BoolType,
		Computed:    true,
	}
}
