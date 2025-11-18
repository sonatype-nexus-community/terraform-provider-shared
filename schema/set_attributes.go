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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// RequiredStringSet returns a required set attribute with string elements
func RequiredStringSet(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.StringType,
		Required:    true,
	}
}

// OptionalStringSet returns an optional set attribute with string elements
func OptionalStringSet(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
	}
}

// ComputedStringSet returns a computed set attribute with string elements
func ComputedStringSet(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.StringType,
		Computed:    true,
	}
}

// ComputedOptionalStringSet returns a computed optional set attribute with string elements
func ComputedOptionalStringSet(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
	}
}

// RequiredInt64Set returns a required set attribute with int64 elements
func RequiredInt64Set(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Required:    true,
	}
}

// OptionalInt64Set returns an optional set attribute with int64 elements
func OptionalInt64Set(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Optional:    true,
	}
}

// ComputedInt64Set returns a computed set attribute with int64 elements
func ComputedInt64Set(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Computed:    true,
	}
}

// RequiredBoolSet returns a required set attribute with bool elements
func RequiredBoolSet(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.BoolType,
		Required:    true,
	}
}

// OptionalBoolSet returns an optional set attribute with bool elements
func OptionalBoolSet(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.BoolType,
		Optional:    true,
	}
}

// ComputedBoolSet returns a computed set attribute with bool elements
func ComputedBoolSet(description string) schema.SetAttribute {
	return schema.SetAttribute{
		Description: description,
		ElementType: types.BoolType,
		Computed:    true,
	}
}
