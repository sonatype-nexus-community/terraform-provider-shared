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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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

// ResourceRequiredStringSetWithValidator returns a required set attribute with string elements and validators
func ResourceRequiredStringSetWithValidator(description string, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		required:    true,
		validators:  validators,
	})
}

// ResourceOptionalStringSetWithValidator returns an optional set attribute with string elements and validators
func ResourceOptionalStringSetWithValidator(description string, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		optional:    true,
		validators:  validators,
	})
}

// ResourceComputedStringSetWithValidator returns a computed set attribute with string elements and validators
func ResourceComputedStringSetWithValidator(description string, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		computed:    true,
		validators:  validators,
	})
}

// ResourceOptionalStringSetWithDefault returns an optional set attribute with string elements and a default value
func ResourceOptionalStringSetWithDefault(description string, defaultValue defaults.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description:  description,
		elementType:  types.StringType,
		optional:     true,
		computed:     true,
		defaultValue: defaultValue,
	})
}

// ResourceOptionalStringSetWithDefaultAndValidator returns an optional set attribute with string elements, a default value, and validators
func ResourceOptionalStringSetWithDefaultAndValidator(description string, defaultValue defaults.Set, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description:  description,
		elementType:  types.StringType,
		optional:     true,
		computed:     true,
		defaultValue: defaultValue,
		validators:   validators,
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

// ResourceRequiredInt64SetWithValidator returns a required set attribute with int64 elements and validators
func ResourceRequiredInt64SetWithValidator(description string, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		required:    true,
		validators:  validators,
	})
}

// ResourceOptionalInt64SetWithValidator returns an optional set attribute with int64 elements and validators
func ResourceOptionalInt64SetWithValidator(description string, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		optional:    true,
		validators:  validators,
	})
}

// ResourceComputedInt64SetWithValidator returns a computed set attribute with int64 elements and validators
func ResourceComputedInt64SetWithValidator(description string, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.Int64Type,
		computed:    true,
		validators:  validators,
	})
}

// ResourceOptionalInt64SetWithDefault returns an optional set attribute with int64 elements and a default value
func ResourceOptionalInt64SetWithDefault(description string, defaultValue defaults.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description:  description,
		elementType:  types.Int64Type,
		optional:     true,
		defaultValue: defaultValue,
	})
}

// ResourceOptionalInt64SetWithDefaultAndValidator returns an optional set attribute with int64 elements, a default value, and validators
func ResourceOptionalInt64SetWithDefaultAndValidator(description string, defaultValue defaults.Set, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description:  description,
		elementType:  types.Int64Type,
		optional:     true,
		defaultValue: defaultValue,
		validators:   validators,
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

// ResourceRequiredBoolSetWithValidator returns a required set attribute with bool elements and validators
func ResourceRequiredBoolSetWithValidator(description string, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		required:    true,
		validators:  validators,
	})
}

// ResourceOptionalBoolSetWithValidator returns an optional set attribute with bool elements and validators
func ResourceOptionalBoolSetWithValidator(description string, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		optional:    true,
		validators:  validators,
	})
}

// ResourceComputedBoolSetWithValidator returns a computed set attribute with bool elements and validators
func ResourceComputedBoolSetWithValidator(description string, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.BoolType,
		computed:    true,
		validators:  validators,
	})
}

// ResourceOptionalBoolSetWithDefault returns an optional set attribute with bool elements and a default value
func ResourceOptionalBoolSetWithDefault(description string, defaultValue defaults.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description:  description,
		elementType:  types.BoolType,
		optional:     true,
		defaultValue: defaultValue,
	})
}

// ResourceOptionalBoolSetWithDefaultAndValidator returns an optional set attribute with bool elements, a default value, and validators
func ResourceOptionalBoolSetWithDefaultAndValidator(description string, defaultValue defaults.Set, validators ...validator.Set) resourceschema.SetAttribute {
	return newResourceSetAttribute(collectionConfig{
		description:  description,
		elementType:  types.BoolType,
		optional:     true,
		defaultValue: defaultValue,
		validators:   validators,
	})
}

// ========================================
// Data Source Schema Functions - String Sets
// ========================================

// DataSourceRequiredStringSet returns a required set attribute with string elements for data sources
func DataSourceRequiredStringSet(description string) datasourceschema.SetAttribute {
	return newDataSourceSetAttribute(collectionConfig{
		description: description,
		elementType: types.StringType,
		required:    true,
	})
}

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
