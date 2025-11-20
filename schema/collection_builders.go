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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ========================================
// Internal Configuration Types
// ========================================

// collectionConfig holds common configuration for collection attributes (list, map, set)
type collectionConfig struct {
	description string
	elementType attr.Type
	required    bool
	optional    bool
	computed    bool
	validators  []validator.Set
}

// ========================================
// Resource Collection Builders
// ========================================

// newResourceListAttribute creates a list attribute from the given configuration
func newResourceListAttribute(config collectionConfig) resourceschema.ListAttribute {
	return resourceschema.ListAttribute{
		MarkdownDescription: config.description,
		ElementType:         config.elementType,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
}

// newResourceMapAttribute creates a map attribute from the given configuration
func newResourceMapAttribute(config collectionConfig) resourceschema.MapAttribute {
	return resourceschema.MapAttribute{
		MarkdownDescription: config.description,
		ElementType:         config.elementType,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
}

// newResourceSetAttribute creates a set attribute from the given configuration
func newResourceSetAttribute(config collectionConfig) resourceschema.SetAttribute {
	attr := resourceschema.SetAttribute{
		MarkdownDescription: config.description,
		ElementType:         config.elementType,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	return attr
}

// ========================================
// Data Source Collection Builders
// ========================================

// newDataSourceListAttribute creates a list attribute from the given configuration for data sources
func newDataSourceListAttribute(config collectionConfig) datasourceschema.ListAttribute {
	return datasourceschema.ListAttribute{
		MarkdownDescription: config.description,
		ElementType:         config.elementType,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
}

// newDataSourceMapAttribute creates a map attribute from the given configuration for data sources
func newDataSourceMapAttribute(config collectionConfig) datasourceschema.MapAttribute {
	return datasourceschema.MapAttribute{
		MarkdownDescription: config.description,
		ElementType:         config.elementType,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
}

// newDataSourceSetAttribute creates a set attribute from the given configuration for data sources
func newDataSourceSetAttribute(config collectionConfig) datasourceschema.SetAttribute {
	attr := datasourceschema.SetAttribute{
		MarkdownDescription: config.description,
		ElementType:         config.elementType,
		Required:            config.required,
		Optional:            config.optional,
		Computed:            config.computed,
	}
	if len(config.validators) > 0 {
		attr.Validators = config.validators
	}
	return attr
}
