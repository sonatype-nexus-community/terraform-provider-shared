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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// RequiredSingleNestedAttribute returns a required single nested attribute
func RequiredSingleNestedAttribute(description string, attributes map[string]schema.Attribute) schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: description,
		Required:    true,
		Attributes:  attributes,
	}
}

// OptionalSingleNestedAttribute returns an optional single nested attribute
func OptionalSingleNestedAttribute(description string, attributes map[string]schema.Attribute) schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: description,
		Optional:    true,
		Attributes:  attributes,
	}
}

// ComputedSingleNestedAttribute returns a computed single nested attribute
func ComputedSingleNestedAttribute(description string, attributes map[string]schema.Attribute) schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: description,
		Computed:    true,
		Attributes:  attributes,
	}
}

// ComputedOptionalSingleNestedAttribute returns a computed optional single nested attribute
func ComputedOptionalSingleNestedAttribute(description string, attributes map[string]schema.Attribute) schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Attributes:  attributes,
	}
}

// RequiredListNestedAttribute returns a required list nested attribute
func RequiredListNestedAttribute(description string, nestedObject schema.NestedAttributeObject) schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Description: description,
		Required:    true,
		NestedObject: nestedObject,
	}
}

// OptionalListNestedAttribute returns an optional list nested attribute
func OptionalListNestedAttribute(description string, nestedObject schema.NestedAttributeObject) schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Description: description,
		Optional:    true,
		NestedObject: nestedObject,
	}
}

// ComputedListNestedAttribute returns a computed list nested attribute
func ComputedListNestedAttribute(description string, nestedObject schema.NestedAttributeObject) schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Description: description,
		Computed:    true,
		NestedObject: nestedObject,
	}
}

// ComputedOptionalListNestedAttribute returns a computed optional list nested attribute
func ComputedOptionalListNestedAttribute(description string, nestedObject schema.NestedAttributeObject) schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		NestedObject: nestedObject,
	}
}

// RequiredSetNestedAttribute returns a required set nested attribute
func RequiredSetNestedAttribute(description string, nestedObject schema.NestedAttributeObject) schema.SetNestedAttribute {
	return schema.SetNestedAttribute{
		Description: description,
		Required:    true,
		NestedObject: nestedObject,
	}
}

// OptionalSetNestedAttribute returns an optional set nested attribute
func OptionalSetNestedAttribute(description string, nestedObject schema.NestedAttributeObject) schema.SetNestedAttribute {
	return schema.SetNestedAttribute{
		Description: description,
		Optional:    true,
		NestedObject: nestedObject,
	}
}

// ComputedSetNestedAttribute returns a computed set nested attribute
func ComputedSetNestedAttribute(description string, nestedObject schema.NestedAttributeObject) schema.SetNestedAttribute {
	return schema.SetNestedAttribute{
		Description: description,
		Computed:    true,
		NestedObject: nestedObject,
	}
}

// ComputedOptionalSetNestedAttribute returns a computed optional set nested attribute
func ComputedOptionalSetNestedAttribute(description string, nestedObject schema.NestedAttributeObject) schema.SetNestedAttribute {
	return schema.SetNestedAttribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		NestedObject: nestedObject,
	}
}
