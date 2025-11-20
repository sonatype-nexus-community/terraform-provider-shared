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
// Resource Schema Functions - String Maps
// ========================================

// ResourceRequiredStringMap returns a required map attribute with string values
func ResourceRequiredStringMap(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		required:    true,
	})
}

// ResourceOptionalStringMap returns an optional map attribute with string values
func ResourceOptionalStringMap(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
	})
}

// ResourceComputedStringMap returns a computed map attribute with string values
func ResourceComputedStringMap(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		computed:    true,
	})
}

// ResourceComputedOptionalStringMap returns a computed optional map attribute with string values
func ResourceComputedOptionalStringMap(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
		computed:    true,
	})
}

// ========================================
// Resource Schema Functions - Int64 Maps
// ========================================

// ResourceRequiredInt64Map returns a required map attribute with int64 values
func ResourceRequiredInt64Map(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		required:    true,
	})
}

// ResourceOptionalInt64Map returns an optional map attribute with int64 values
func ResourceOptionalInt64Map(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		optional:    true,
	})
}

// ResourceComputedInt64Map returns a computed map attribute with int64 values
func ResourceComputedInt64Map(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		computed:    true,
	})
}

// ResourceComputedOptionalInt64Map returns a computed optional map attribute with int64 values
func ResourceComputedOptionalInt64Map(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		optional:    true,
		computed:    true,
	})
}

// ========================================
// Resource Schema Functions - Bool Maps
// ========================================

// ResourceRequiredBoolMap returns a required map attribute with bool values
func ResourceRequiredBoolMap(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		required:    true,
	})
}

// ResourceOptionalBoolMap returns an optional map attribute with bool values
func ResourceOptionalBoolMap(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		optional:    true,
	})
}

// ResourceComputedBoolMap returns a computed map attribute with bool values
func ResourceComputedBoolMap(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		computed:    true,
	})
}

// ResourceComputedOptionalBoolMap returns a computed optional map attribute with bool values
func ResourceComputedOptionalBoolMap(description string) resourceschema.MapAttribute {
	return newResourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		optional:    true,
		computed:    true,
	})
}

// ========================================
// Data Source Schema Functions - String Maps
// ========================================

// DataSourceOptionalStringMap returns an optional map attribute with string values for data sources
func DataSourceOptionalStringMap(description string) datasourceschema.MapAttribute {
	return newDataSourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
	})
}

// DataSourceComputedStringMap returns a computed map attribute with string values for data sources
func DataSourceComputedStringMap(description string) datasourceschema.MapAttribute {
	return newDataSourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		computed:    true,
	})
}

// ========================================
// Data Source Schema Functions - Int64 Maps
// ========================================

// DataSourceOptionalInt64Map returns an optional map attribute with int64 values for data sources
func DataSourceOptionalInt64Map(description string) datasourceschema.MapAttribute {
	return newDataSourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		optional:    true,
	})
}

// DataSourceComputedInt64Map returns a computed map attribute with int64 values for data sources
func DataSourceComputedInt64Map(description string) datasourceschema.MapAttribute {
	return newDataSourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		computed:    true,
	})
}

// ========================================
// Data Source Schema Functions - Bool Maps
// ========================================

// DataSourceOptionalBoolMap returns an optional map attribute with bool values for data sources
func DataSourceOptionalBoolMap(description string) datasourceschema.MapAttribute {
	return newDataSourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		optional:    true,
	})
}

// DataSourceComputedBoolMap returns a computed map attribute with bool values for data sources
func DataSourceComputedBoolMap(description string) datasourceschema.MapAttribute {
	return newDataSourceMapAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		computed:    true,
	})
}
