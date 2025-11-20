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
)

// ========================================
// Resource Schema Functions
// ========================================

// ResourceRequiredSingleNestedAttribute returns a required single nested attribute
func ResourceRequiredSingleNestedAttribute(description string, attributes map[string]resourceschema.Attribute) resourceschema.SingleNestedAttribute {
	return resourceschema.SingleNestedAttribute{
		Description: description,
		Required:    true,
		Attributes:  attributes,
	}
}

// ResourceOptionalSingleNestedAttribute returns an optional single nested attribute
func ResourceOptionalSingleNestedAttribute(description string, attributes map[string]resourceschema.Attribute) resourceschema.SingleNestedAttribute {
	return resourceschema.SingleNestedAttribute{
		Description: description,
		Optional:    true,
		Attributes:  attributes,
	}
}

// ResourceComputedSingleNestedAttribute returns a computed single nested attribute
func ResourceComputedSingleNestedAttribute(description string, attributes map[string]resourceschema.Attribute) resourceschema.SingleNestedAttribute {
	return resourceschema.SingleNestedAttribute{
		Description: description,
		Computed:    true,
		Attributes:  attributes,
	}
}

// ResourceComputedOptionalSingleNestedAttribute returns a computed optional single nested attribute
func ResourceComputedOptionalSingleNestedAttribute(description string, attributes map[string]resourceschema.Attribute) resourceschema.SingleNestedAttribute {
	return resourceschema.SingleNestedAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Attributes:  attributes,
	}
}

// ResourceRequiredListNestedAttribute returns a required list nested attribute
func ResourceRequiredListNestedAttribute(description string, nestedObject resourceschema.NestedAttributeObject) resourceschema.ListNestedAttribute {
	return resourceschema.ListNestedAttribute{
		Description: description,
		Required:    true,
		NestedObject: nestedObject,
	}
}

// ResourceOptionalListNestedAttribute returns an optional list nested attribute
func ResourceOptionalListNestedAttribute(description string, nestedObject resourceschema.NestedAttributeObject) resourceschema.ListNestedAttribute {
	return resourceschema.ListNestedAttribute{
		Description: description,
		Optional:    true,
		NestedObject: nestedObject,
	}
}

// ResourceComputedListNestedAttribute returns a computed list nested attribute
func ResourceComputedListNestedAttribute(description string, nestedObject resourceschema.NestedAttributeObject) resourceschema.ListNestedAttribute {
	return resourceschema.ListNestedAttribute{
		Description: description,
		Computed:    true,
		NestedObject: nestedObject,
	}
}

// ResourceComputedOptionalListNestedAttribute returns a computed optional list nested attribute
func ResourceComputedOptionalListNestedAttribute(description string, nestedObject resourceschema.NestedAttributeObject) resourceschema.ListNestedAttribute {
	return resourceschema.ListNestedAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		NestedObject: nestedObject,
	}
}

// ResourceRequiredSetNestedAttribute returns a required set nested attribute
func ResourceRequiredSetNestedAttribute(description string, nestedObject resourceschema.NestedAttributeObject) resourceschema.SetNestedAttribute {
	return resourceschema.SetNestedAttribute{
		Description: description,
		Required:    true,
		NestedObject: nestedObject,
	}
}

// ResourceOptionalSetNestedAttribute returns an optional set nested attribute
func ResourceOptionalSetNestedAttribute(description string, nestedObject resourceschema.NestedAttributeObject) resourceschema.SetNestedAttribute {
	return resourceschema.SetNestedAttribute{
		Description: description,
		Optional:    true,
		NestedObject: nestedObject,
	}
}

// ResourceComputedSetNestedAttribute returns a computed set nested attribute
func ResourceComputedSetNestedAttribute(description string, nestedObject resourceschema.NestedAttributeObject) resourceschema.SetNestedAttribute {
	return resourceschema.SetNestedAttribute{
		Description: description,
		Computed:    true,
		NestedObject: nestedObject,
	}
}

// ResourceComputedOptionalSetNestedAttribute returns a computed optional set nested attribute
func ResourceComputedOptionalSetNestedAttribute(description string, nestedObject resourceschema.NestedAttributeObject) resourceschema.SetNestedAttribute {
	return resourceschema.SetNestedAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		NestedObject: nestedObject,
	}
}

// ========================================
// Data Source Schema Functions
// ========================================

// DataSourceOptionalSingleNestedAttribute returns an optional single nested attribute for data sources
func DataSourceOptionalSingleNestedAttribute(description string, attributes map[string]datasourceschema.Attribute) datasourceschema.SingleNestedAttribute {
	return datasourceschema.SingleNestedAttribute{
		Description: description,
		Optional:    true,
		Attributes:  attributes,
	}
}

// DataSourceComputedSingleNestedAttribute returns a computed single nested attribute for data sources
func DataSourceComputedSingleNestedAttribute(description string, attributes map[string]datasourceschema.Attribute) datasourceschema.SingleNestedAttribute {
	return datasourceschema.SingleNestedAttribute{
		Description: description,
		Computed:    true,
		Attributes:  attributes,
	}
}

// DataSourceComputedOptionalSingleNestedAttribute returns a computed optional single nested attribute for data sources
func DataSourceComputedOptionalSingleNestedAttribute(description string, attributes map[string]datasourceschema.Attribute) datasourceschema.SingleNestedAttribute {
	return datasourceschema.SingleNestedAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Attributes:  attributes,
	}
}

// DataSourceOptionalListNestedAttribute returns an optional list nested attribute for data sources
func DataSourceOptionalListNestedAttribute(description string, nestedObject datasourceschema.NestedAttributeObject) datasourceschema.ListNestedAttribute {
	return datasourceschema.ListNestedAttribute{
		Description: description,
		Optional:    true,
		NestedObject: nestedObject,
	}
}

// DataSourceComputedListNestedAttribute returns a computed list nested attribute for data sources
func DataSourceComputedListNestedAttribute(description string, nestedObject datasourceschema.NestedAttributeObject) datasourceschema.ListNestedAttribute {
	return datasourceschema.ListNestedAttribute{
		Description: description,
		Computed:    true,
		NestedObject: nestedObject,
	}
}

// DataSourceOptionalSetNestedAttribute returns an optional set nested attribute for data sources
func DataSourceOptionalSetNestedAttribute(description string, nestedObject datasourceschema.NestedAttributeObject) datasourceschema.SetNestedAttribute {
	return datasourceschema.SetNestedAttribute{
		Description: description,
		Optional:    true,
		NestedObject: nestedObject,
	}
}

// DataSourceComputedSetNestedAttribute returns a computed set nested attribute for data sources
func DataSourceComputedSetNestedAttribute(description string, nestedObject datasourceschema.NestedAttributeObject) datasourceschema.SetNestedAttribute {
	return datasourceschema.SetNestedAttribute{
		Description: description,
		Computed:    true,
		NestedObject: nestedObject,
	}
}
