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

// RequiredStringList returns a required list attribute with string elements
func RequiredStringList(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Required:    true,
	}
}

// OptionalStringList returns an optional list attribute with string elements
func OptionalStringList(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
	}
}

// ComputedStringList returns a computed list attribute with string elements
func ComputedStringList(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Computed:    true,
	}
}

// ComputedOptionalStringList returns a computed optional list attribute with string elements
func ComputedOptionalStringList(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
	}
}

// RequiredInt64List returns a required list attribute with int64 elements
func RequiredInt64List(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Required:    true,
	}
}

// OptionalInt64List returns an optional list attribute with int64 elements
func OptionalInt64List(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Optional:    true,
	}
}

// ComputedInt64List returns a computed list attribute with int64 elements
func ComputedInt64List(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Computed:    true,
	}
}

// RequiredBoolList returns a required list attribute with bool elements
func RequiredBoolList(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.BoolType,
		Required:    true,
	}
}

// OptionalBoolList returns an optional list attribute with bool elements
func OptionalBoolList(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.BoolType,
		Optional:    true,
	}
}

// ComputedBoolList returns a computed list attribute with bool elements
func ComputedBoolList(description string) schema.ListAttribute {
	return schema.ListAttribute{
		Description: description,
		ElementType: types.BoolType,
		Computed:    true,
	}
}
