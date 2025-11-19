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
// Resource Schema Functions - String Lists
// ========================================

// ResourceRequiredStringList returns a required list attribute with string elements
func ResourceRequiredStringList(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		required:    true,
	})
}

// ResourceOptionalStringList returns an optional list attribute with string elements
func ResourceOptionalStringList(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
	})
}

// ResourceComputedStringList returns a computed list attribute with string elements
func ResourceComputedStringList(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		computed:    true,
	})
}

// ResourceComputedOptionalStringList returns a computed optional list attribute with string elements
func ResourceComputedOptionalStringList(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
		computed:    true,
	})
}

// ========================================
// Resource Schema Functions - Int64 Lists
// ========================================

// ResourceRequiredInt64List returns a required list attribute with int64 elements
func ResourceRequiredInt64List(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		required:    true,
	})
}

// ResourceOptionalInt64List returns an optional list attribute with int64 elements
func ResourceOptionalInt64List(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		optional:    true,
	})
}

// ResourceComputedInt64List returns a computed list attribute with int64 elements
func ResourceComputedInt64List(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		computed:    true,
	})
}

// ========================================
// Resource Schema Functions - Bool Lists
// ========================================

// ResourceRequiredBoolList returns a required list attribute with bool elements
func ResourceRequiredBoolList(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		required:    true,
	})
}

// ResourceOptionalBoolList returns an optional list attribute with bool elements
func ResourceOptionalBoolList(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		optional:    true,
	})
}

// ResourceComputedBoolList returns a computed list attribute with bool elements
func ResourceComputedBoolList(description string) resourceschema.ListAttribute {
	return newResourceListAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		computed:    true,
	})
}

// ========================================
// Data Source Schema Functions - String Lists
// ========================================

// DataSourceOptionalStringList returns an optional list attribute with string elements for data sources
func DataSourceOptionalStringList(description string) datasourceschema.ListAttribute {
	return newDataSourceListAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
	})
}

// DataSourceComputedStringList returns a computed list attribute with string elements for data sources
func DataSourceComputedStringList(description string) datasourceschema.ListAttribute {
	return newDataSourceListAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		computed:    true,
	})
}

// ========================================
// Data Source Schema Functions - Int64 Lists
// ========================================

// DataSourceOptionalInt64List returns an optional list attribute with int64 elements for data sources
func DataSourceOptionalInt64List(description string) datasourceschema.ListAttribute {
	return newDataSourceListAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		optional:    true,
	})
}

// DataSourceComputedInt64List returns a computed list attribute with int64 elements for data sources
func DataSourceComputedInt64List(description string) datasourceschema.ListAttribute {
	return newDataSourceListAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		computed:    true,
	})
}

// ========================================
// Data Source Schema Functions - Bool Lists
// ========================================

// DataSourceOptionalBoolList returns an optional list attribute with bool elements for data sources
func DataSourceOptionalBoolList(description string) datasourceschema.ListAttribute {
	return newDataSourceListAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		optional:    true,
	})
}

// DataSourceComputedBoolList returns a computed list attribute with bool elements for data sources
func DataSourceComputedBoolList(description string) datasourceschema.ListAttribute {
	return newDataSourceListAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		computed:    true,
	})
}
