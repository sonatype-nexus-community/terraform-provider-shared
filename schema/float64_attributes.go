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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
)

// ========================================
// Resource Schema Functions
// ========================================

// ResourceRequiredFloat64 returns a required float64 attribute
func ResourceRequiredFloat64(description string) resourceschema.Float64Attribute {
	return newResourceFloat64Attribute(float64AttributeConfig{
		description: description,
		required:    true,
	})
}

// ResourceOptionalFloat64 returns an optional float64 attribute
func ResourceOptionalFloat64(description string) resourceschema.Float64Attribute {
	return newResourceFloat64Attribute(float64AttributeConfig{
		description: description,
		optional:    true,
	})
}

// ResourceComputedFloat64 returns a computed float64 attribute
func ResourceComputedFloat64(description string) resourceschema.Float64Attribute {
	return newResourceFloat64Attribute(float64AttributeConfig{
		description: description,
		computed:    true,
	})
}

// ResourceComputedFloat64WithDefault returns a computed float64 attribute with a default value
func ResourceComputedFloat64WithDefault(description string, defaultValue float64) resourceschema.Float64Attribute {
	return newResourceFloat64Attribute(float64AttributeConfig{
		description:  description,
		computed:     true,
		defaultValue: float64default.StaticFloat64(defaultValue),
	})
}

// ResourceOptionalFloat64WithDefault returns an optional float64 attribute with a default value
func ResourceOptionalFloat64WithDefault(description string, defaultValue float64) resourceschema.Float64Attribute {
	return newResourceFloat64Attribute(float64AttributeConfig{
		description:  description,
		optional:     true,
		computed:     true,
		defaultValue: float64default.StaticFloat64(defaultValue),
	})
}

// ResourceComputedOptionalFloat64 returns a computed optional float64 attribute
func ResourceComputedOptionalFloat64(description string) resourceschema.Float64Attribute {
	return newResourceFloat64Attribute(float64AttributeConfig{
		description: description,
		optional:    true,
		computed:    true,
	})
}

// ResourceRequiredFloat64WithDefault returns a required float64 attribute with a default value
func ResourceRequiredFloat64WithDefault(description string, defaultValue float64) resourceschema.Float64Attribute {
	return newResourceFloat64Attribute(float64AttributeConfig{
		description:  description,
		required:     true,
		defaultValue: float64default.StaticFloat64(defaultValue),
	})
}

// ========================================
// Data Source Schema Functions
// ========================================

// DataSourceComputedFloat64 returns a computed float64 attribute for data sources
func DataSourceComputedFloat64(description string) datasourceschema.Float64Attribute {
	return newDataSourceFloat64Attribute(float64AttributeConfig{
		description: description,
		computed:    true,
	})
}

// DataSourceOptionalFloat64 returns an optional float64 attribute for data sources
func DataSourceOptionalFloat64(description string) datasourceschema.Float64Attribute {
	return newDataSourceFloat64Attribute(float64AttributeConfig{
		description: description,
		optional:    true,
	})
}
