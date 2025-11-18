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
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// String Attribute Validators

// RequiredStringWithRegex returns a required string attribute with regex validation
func RequiredStringWithRegex(description string, pattern *regexp.Regexp, errorMsg string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
		},
	}
}

// OptionalStringWithRegex returns an optional string attribute with regex validation
func OptionalStringWithRegex(description string, pattern *regexp.Regexp, errorMsg string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
		},
	}
}

// ComputedStringWithRegex returns a computed string attribute with regex validation
func ComputedStringWithRegex(description string, pattern *regexp.Regexp, errorMsg string) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Computed:    true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
		},
	}
}

// StringWithValidators returns a required string attribute with custom validators list
func StringWithValidators(description string, validators ...validator.String) schema.StringAttribute {
	attr := schema.StringAttribute{
		Description: description,
		Required:    true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// OptionalStringWithValidators returns an optional string attribute with custom validators
func OptionalStringWithValidators(description string, validators ...validator.String) schema.StringAttribute {
	attr := schema.StringAttribute{
		Description: description,
		Optional:    true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// RequiredStringWithValidators returns a required string attribute with custom validators
func RequiredStringWithValidators(description string, validators ...validator.String) schema.StringAttribute {
	return StringWithValidators(description, validators...)
}

// ComputedStringWithValidators returns a computed string attribute with custom validators
func ComputedStringWithValidators(description string, validators ...validator.String) schema.StringAttribute {
	attr := schema.StringAttribute{
		Description: description,
		Computed:    true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// RequiredStringWithLengthBetween returns a required string attribute with length validation
func RequiredStringWithLengthBetween(description string, minLength, maxLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// OptionalStringWithLengthBetween returns an optional string attribute with length validation
func OptionalStringWithLengthBetween(description string, minLength, maxLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// RequiredStringWithLengthAtLeast returns a required string attribute with minimum length validation
func RequiredStringWithLengthAtLeast(description string, minLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// OptionalStringWithLengthAtLeast returns an optional string attribute with minimum length validation
func OptionalStringWithLengthAtLeast(description string, minLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// RequiredStringWithLengthAtMost returns a required string attribute with maximum length validation
func RequiredStringWithLengthAtMost(description string, maxLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
		Validators: []validator.String{
			stringvalidator.LengthAtMost(maxLength),
		},
	}
}

// OptionalStringWithLengthAtMost returns an optional string attribute with maximum length validation
func OptionalStringWithLengthAtMost(description string, maxLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Validators: []validator.String{
			stringvalidator.LengthAtMost(maxLength),
		},
	}
}

// OptionalSensitiveStringWithLengthAtLeast returns an optional sensitive string attribute with minimum length validation
func OptionalSensitiveStringWithLengthAtLeast(description string, minLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Sensitive:   true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// RequiredSensitiveStringWithLengthAtLeast returns a required sensitive string attribute with minimum length validation
func RequiredSensitiveStringWithLengthAtLeast(description string, minLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
		Sensitive:   true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// RequiredStringWithRegexAndLength returns a required string attribute with regex and length validation
func RequiredStringWithRegexAndLength(description string, pattern *regexp.Regexp, errorMsg string, minLength, maxLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// OptionalStringWithRegexAndLength returns an optional string attribute with regex and length validation
func OptionalStringWithRegexAndLength(description string, pattern *regexp.Regexp, errorMsg string, minLength, maxLength int) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Optional:    true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// Integer Attribute Validators

// RequiredInt64WithRange returns a required int64 attribute with range validation
func RequiredInt64WithRange(description string, minValue, maxValue int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Required:    true,
		Validators: []validator.Int64{
			int64validator.Between(minValue, maxValue),
		},
	}
}

// OptionalInt64WithRange returns an optional int64 attribute with range validation
func OptionalInt64WithRange(description string, minValue, maxValue int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Optional:    true,
		Validators: []validator.Int64{
			int64validator.Between(minValue, maxValue),
		},
	}
}

// ComputedInt64WithRange returns a computed int64 attribute with range validation
func ComputedInt64WithRange(description string, minValue, maxValue int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		Description: description,
		Computed:    true,
		Validators: []validator.Int64{
			int64validator.Between(minValue, maxValue),
		},
	}
}

// Int64WithValidators returns a required int64 attribute with custom validators
func Int64WithValidators(description string, validators ...validator.Int64) schema.Int64Attribute {
	attr := schema.Int64Attribute{
		Description: description,
		Required:    true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// OptionalInt64WithValidators returns an optional int64 attribute with custom validators
func OptionalInt64WithValidators(description string, validators ...validator.Int64) schema.Int64Attribute {
	attr := schema.Int64Attribute{
		Description: description,
		Optional:    true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// RequiredInt64WithValidators returns a required int64 attribute with custom validators
func RequiredInt64WithValidators(description string, validators ...validator.Int64) schema.Int64Attribute {
	return Int64WithValidators(description, validators...)
}

// ComputedInt64WithValidators returns a computed int64 attribute with custom validators
func ComputedInt64WithValidators(description string, validators ...validator.Int64) schema.Int64Attribute {
	attr := schema.Int64Attribute{
		Description: description,
		Computed:    true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}
