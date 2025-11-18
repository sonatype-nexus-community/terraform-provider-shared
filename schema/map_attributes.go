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

// RequiredStringMap returns a required map attribute with string values
func RequiredStringMap(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Required:    true,
	}
}

// OptionalStringMap returns an optional map attribute with string values
func OptionalStringMap(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
	}
}

// ComputedStringMap returns a computed map attribute with string values
func ComputedStringMap(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Computed:    true,
	}
}

// ComputedOptionalStringMap returns a computed optional map attribute with string values
func ComputedOptionalStringMap(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
	}
}

// RequiredInt64Map returns a required map attribute with int64 values
func RequiredInt64Map(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Required:    true,
	}
}

// OptionalInt64Map returns an optional map attribute with int64 values
func OptionalInt64Map(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Optional:    true,
	}
}

// ComputedInt64Map returns a computed map attribute with int64 values
func ComputedInt64Map(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.Int64Type,
		Computed:    true,
	}
}

// RequiredBoolMap returns a required map attribute with bool values
func RequiredBoolMap(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.BoolType,
		Required:    true,
	}
}

// OptionalBoolMap returns an optional map attribute with bool values
func OptionalBoolMap(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.BoolType,
		Optional:    true,
	}
}

// ComputedBoolMap returns a computed map attribute with bool values
func ComputedBoolMap(description string) schema.MapAttribute {
	return schema.MapAttribute{
		Description: description,
		ElementType: types.BoolType,
		Computed:    true,
	}
}
