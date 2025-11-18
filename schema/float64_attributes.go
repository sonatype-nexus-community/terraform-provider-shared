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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
)

// RequiredFloat64 returns a required float64 attribute
func RequiredFloat64(description string) schema.Float64Attribute {
	return schema.Float64Attribute{
		Description: description,
		Required:    true,
	}
}

// OptionalFloat64 returns an optional float64 attribute
func OptionalFloat64(description string) schema.Float64Attribute {
	return schema.Float64Attribute{
		Description: description,
		Optional:    true,
	}
}

// ComputedFloat64 returns a computed float64 attribute
func ComputedFloat64(description string) schema.Float64Attribute {
	return schema.Float64Attribute{
		Description: description,
		Computed:    true,
	}
}

// OptionalFloat64WithDefault returns an optional float64 attribute with a default value
func OptionalFloat64WithDefault(description string, defaultValue float64) schema.Float64Attribute {
	return schema.Float64Attribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     float64default.StaticFloat64(defaultValue),
	}
}

// RequiredFloat64WithDefault returns a required float64 attribute with a default value
func RequiredFloat64WithDefault(description string, defaultValue float64) schema.Float64Attribute {
	return schema.Float64Attribute{
		Description: description,
		Required:    true,
		Default:     float64default.StaticFloat64(defaultValue),
	}
}
