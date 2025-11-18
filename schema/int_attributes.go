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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
)

// RequiredInt64 returns a required int64 attribute
func RequiredInt64(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Required:    true,
	}
}

// OptionalInt64 returns an optional int64 attribute
func OptionalInt64(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
	}
}

// ComputedInt64 returns a computed int64 attribute
func ComputedInt64(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Computed:    true,
	}
}

// ComputedInt64WithDefault returns a computed int64 attribute with a default value
func ComputedInt64WithDefault(description string, defaultValue int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Computed:    true,
		Default:     int64default.StaticInt64(defaultValue),
	}
}

// OptionalInt64WithDefault returns an optional int64 attribute with a default value
func OptionalInt64WithDefault(description string, defaultValue int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     int64default.StaticInt64(defaultValue),
	}
}

// RequiredInt64WithDefault returns a required int64 attribute with a default value
func RequiredInt64WithDefault(description string, defaultValue int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Required:    true,
		Default:     int64default.StaticInt64(defaultValue),
	}
}

// RequiredInt32 returns a required int32 attribute
func RequiredInt32(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Required:    true,
	}
}

// OptionalInt32 returns an optional int32 attribute
func OptionalInt32(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
	}
}

// OptionalInt32WithDefault returns an optional int32 attribute with a default value
func OptionalInt32WithDefault(description string, defaultValue int32) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     int64default.StaticInt64(int64(defaultValue)),
	}
}

// RequiredInt32WithDefault returns a required int32 attribute with a default value
func RequiredInt32WithDefault(description string, defaultValue int32) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Required:    true,
		Default:     int64default.StaticInt64(int64(defaultValue)),
	}
}

// OptionalPort returns an optional int64 attribute for network ports (0-65535)
func OptionalPort(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
	}
}

// RequiredPort returns a required int64 attribute for network ports (0-65535)
func RequiredPort(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Required:    true,
	}
}

// PortWithDefault returns a port attribute with a default value
func PortWithDefault(description string, defaultValue int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     int64default.StaticInt64(defaultValue),
	}
}

// PercentageInt returns a percentage as int64 (0-100)
func PercentageInt(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
	}
}

// RequiredPercentageInt returns a required percentage as int64 (0-100)
func RequiredPercentageInt(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Required:    true,
	}
}

// DurationInt returns a duration in seconds as int64
func DurationInt(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
	}
}

// RequiredDurationInt returns a required duration in seconds as int64
func RequiredDurationInt(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Required:    true,
	}
}

// TimeoutInt returns a timeout duration in seconds as int64
func TimeoutInt(description string, defaultSeconds int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     int64default.StaticInt64(defaultSeconds),
	}
}

// CountInt returns a count attribute as int64
func CountInt(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
		Computed:    true,
		Default:     int64default.StaticInt64(0),
	}
}
