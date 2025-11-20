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
// Resource Schema Functions - String Sets
// ========================================

// ResourceRequiredStringSet returns a required set attribute with string elements
func ResourceRequiredStringSet(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		required:    true,
	})
}

// ResourceOptionalStringSet returns an optional set attribute with string elements
func ResourceOptionalStringSet(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
	})
}

// ResourceComputedStringSet returns a computed set attribute with string elements
func ResourceComputedStringSet(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		computed:    true,
	})
}

// ResourceComputedOptionalStringSet returns a computed optional set attribute with string elements
func ResourceComputedOptionalStringSet(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
		computed:    true,
	})
}

// ========================================
// Resource Schema Functions - Int64 Sets
// ========================================

// ResourceRequiredInt64Set returns a required set attribute with int64 elements
func ResourceRequiredInt64Set(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		required:    true,
	})
}

// ResourceOptionalInt64Set returns an optional set attribute with int64 elements
func ResourceOptionalInt64Set(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		optional:    true,
	})
}

// ResourceComputedInt64Set returns a computed set attribute with int64 elements
func ResourceComputedInt64Set(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		computed:    true,
	})
}

// ResourceComputedOptionalInt64Set returns a computed optional set attribute with int64 elements
func ResourceComputedOptionalInt64Set(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		optional:    true,
		computed:    true,
	})
}

// ========================================
// Resource Schema Functions - Bool Sets
// ========================================

// ResourceRequiredBoolSet returns a required set attribute with bool elements
func ResourceRequiredBoolSet(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		required:    true,
	})
}

// ResourceOptionalBoolSet returns an optional set attribute with bool elements
func ResourceOptionalBoolSet(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		optional:    true,
	})
}

// ResourceComputedBoolSet returns a computed set attribute with bool elements
func ResourceComputedBoolSet(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		computed:    true,
	})
}

// ResourceComputedOptionalBoolSet returns a computed optional set attribute with bool elements
func ResourceComputedOptionalBoolSet(description string) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		optional:    true,
		computed:    true,
	})
}

// ========================================
// Data Source Schema Functions - String Sets
// ========================================

// DataSourceOptionalStringSet returns an optional set attribute with string elements for data sources
func DataSourceOptionalStringSet(description string) datasourceschema.SetAttribute {
	return newDataSourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
	})
}

// DataSourceComputedStringSet returns a computed set attribute with string elements for data sources
func DataSourceComputedStringSet(description string) datasourceschema.SetAttribute {
	return newDataSourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		computed:    true,
	})
}

// ========================================
// Data Source Schema Functions - Int64 Sets
// ========================================

// DataSourceOptionalInt64Set returns an optional set attribute with int64 elements for data sources
func DataSourceOptionalInt64Set(description string) datasourceschema.SetAttribute {
	return newDataSourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		optional:    true,
	})
}

// DataSourceComputedInt64Set returns a computed set attribute with int64 elements for data sources
func DataSourceComputedInt64Set(description string) datasourceschema.SetAttribute {
	return newDataSourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		computed:    true,
	})
}

// ========================================
// Data Source Schema Functions - Bool Sets
// ========================================

// DataSourceOptionalBoolSet returns an optional set attribute with bool elements for data sources
func DataSourceOptionalBoolSet(description string) datasourceschema.SetAttribute {
	return newDataSourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		optional:    true,
	})
}

// DataSourceComputedBoolSet returns a computed set attribute with bool elements for data sources
func DataSourceComputedBoolSet(description string) datasourceschema.SetAttribute {
	return newDataSourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		computed:    true,
	})
}
