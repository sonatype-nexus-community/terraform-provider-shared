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

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ============================================================================
// Resource Validators - String Attributes
// ============================================================================

// ResourceRequiredStringWithRegex returns a required string attribute with regex validation
func ResourceRequiredStringWithRegex(description string, pattern *regexp.Regexp, errorMsg string) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
		},
	}
}

// ResourceOptionalStringWithRegex returns an optional string attribute with regex validation
func ResourceOptionalStringWithRegex(description string, pattern *regexp.Regexp, errorMsg string) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
		},
	}
}

// ResourceComputedStringWithRegex returns a computed string attribute with regex validation
func ResourceComputedStringWithRegex(description string, pattern *regexp.Regexp, errorMsg string) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
		},
	}
}

// ResourceStringWithValidators returns a required string attribute with custom validators list
func ResourceStringWithValidators(description string, validators ...validator.String) resourceschema.StringAttribute {
	attr := resourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceOptionalStringWithValidators returns an optional string attribute with custom validators
func ResourceOptionalStringWithValidators(description string, validators ...validator.String) resourceschema.StringAttribute {
	attr := resourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceRequiredStringWithValidators returns a required string attribute with custom validators
func ResourceRequiredStringWithValidators(description string, validators ...validator.String) resourceschema.StringAttribute {
	return ResourceStringWithValidators(description, validators...)
}

// ResourceComputedStringWithValidators returns a computed string attribute with custom validators
func ResourceComputedStringWithValidators(description string, validators ...validator.String) resourceschema.StringAttribute {
	attr := resourceschema.StringAttribute{
		MarkdownDescription: description,
		Computed:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceRequiredStringWithLengthBetween returns a required string attribute with length validation
func ResourceRequiredStringWithLengthBetween(description string, minLength, maxLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// ResourceOptionalStringWithLengthBetween returns an optional string attribute with length validation
func ResourceOptionalStringWithLengthBetween(description string, minLength, maxLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// ResourceRequiredStringWithLengthAtLeast returns a required string attribute with minimum length validation
func ResourceRequiredStringWithLengthAtLeast(description string, minLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// ResourceOptionalStringWithLengthAtLeast returns an optional string attribute with minimum length validation
func ResourceOptionalStringWithLengthAtLeast(description string, minLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// ResourceRequiredStringWithLengthAtMost returns a required string attribute with maximum length validation
func ResourceRequiredStringWithLengthAtMost(description string, maxLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.LengthAtMost(maxLength),
		},
	}
}

// ResourceOptionalStringWithLengthAtMost returns an optional string attribute with maximum length validation
func ResourceOptionalStringWithLengthAtMost(description string, maxLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.LengthAtMost(maxLength),
		},
	}
}

// ResourceOptionalSensitiveStringWithLengthAtLeast returns an optional sensitive string attribute with minimum length validation
func ResourceOptionalSensitiveStringWithLengthAtLeast(description string, minLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Sensitive:           true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// ResourceRequiredSensitiveStringWithLengthAtLeast returns a required sensitive string attribute with minimum length validation
func ResourceRequiredSensitiveStringWithLengthAtLeast(description string, minLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Sensitive:           true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// ResourceRequiredStringWithRegexAndLength returns a required string attribute with regex and length validation
func ResourceRequiredStringWithRegexAndLength(description string, pattern *regexp.Regexp, errorMsg string, minLength, maxLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// ResourceOptionalStringWithRegexAndLength returns an optional string attribute with regex and length validation
func ResourceOptionalStringWithRegexAndLength(description string, pattern *regexp.Regexp, errorMsg string, minLength, maxLength int) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// ============================================================================
// Resource Validators - Integer Attributes
// ============================================================================

// ResourceRequiredInt32WithRange returns a required int32 attribute with range validation
func ResourceRequiredInt32WithRange(description string, minValue, maxValue int32) resourceschema.Int32Attribute {
	return resourceschema.Int32Attribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.Int32{
			int32validator.Between(minValue, maxValue),
		},
	}
}

// ResourceOptionalInt32WithRange returns an optional int32 attribute with range validation
func ResourceOptionalInt32WithRange(description string, minValue, maxValue int32) resourceschema.Int32Attribute {
	return resourceschema.Int32Attribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.Int32{
			int32validator.Between(minValue, maxValue),
		},
	}
}

// ResourceComputedInt32WithRange returns a computed int32 attribute with range validation
func ResourceComputedInt32WithRange(description string, minValue, maxValue int32) resourceschema.Int32Attribute {
	return resourceschema.Int32Attribute{
		MarkdownDescription: description,
		Computed:            true,
		Validators: []validator.Int32{
			int32validator.Between(minValue, maxValue),
		},
	}
}

// ResourceInt32WithValidators returns a required int32 attribute with custom validators
func ResourceInt32WithValidators(description string, validators ...validator.Int64) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Required:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceOptionalInt32WithValidators returns an optional int32 attribute with custom validators
func ResourceOptionalInt32WithValidators(description string, validators ...validator.Int64) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Optional:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceRequiredInt32WithValidators returns a required int32 attribute with custom validators
func ResourceRequiredInt32WithValidators(description string, validators ...validator.Int64) resourceschema.Int64Attribute {
	return ResourceInt32WithValidators(description, validators...)
}

// ResourceComputedInt32WithValidators returns a computed int32 attribute with custom validators
func ResourceComputedInt32WithValidators(description string, validators ...validator.Int64) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Computed:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceRequiredInt64WithRange returns a required int64 attribute with range validation
func ResourceRequiredInt64WithRange(description string, minValue, maxValue int64) resourceschema.Int64Attribute {
	return resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.Int64{
			int64validator.Between(minValue, maxValue),
		},
	}
}

// ResourceOptionalInt64WithRange returns an optional int64 attribute with range validation
func ResourceOptionalInt64WithRange(description string, minValue, maxValue int64) resourceschema.Int64Attribute {
	return resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.Int64{
			int64validator.Between(minValue, maxValue),
		},
	}
}

// ResourceComputedInt64WithRange returns a computed int64 attribute with range validation
func ResourceComputedInt64WithRange(description string, minValue, maxValue int64) resourceschema.Int64Attribute {
	return resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.Between(minValue, maxValue),
		},
	}
}

// ResourceInt64WithValidators returns a required int64 attribute with custom validators
func ResourceInt64WithValidators(description string, validators ...validator.Int64) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Required:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceOptionalInt64WithValidators returns an optional int64 attribute with custom validators
func ResourceOptionalInt64WithValidators(description string, validators ...validator.Int64) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Optional:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceRequiredInt64WithValidators returns a required int64 attribute with custom validators
func ResourceRequiredInt64WithValidators(description string, validators ...validator.Int64) resourceschema.Int64Attribute {
	return ResourceInt64WithValidators(description, validators...)
}

// ResourceComputedInt64WithValidators returns a computed int64 attribute with custom validators
func ResourceComputedInt64WithValidators(description string, validators ...validator.Int64) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Computed:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceOptionalInt64WithDefaultAndValidators returns an optional int64 attribute with default value and validators
func ResourceOptionalInt64WithDefaultAndValidators(description string, defaultValue int64, validators ...validator.Int64) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(defaultValue),
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceRequiredInt64WithDefaultAndValidators returns a required int64 attribute with default value and validators
func ResourceRequiredInt64WithDefaultAndValidators(description string, defaultValue int64, validators ...validator.Int64) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Required:            true,
		Default:             int64default.StaticInt64(defaultValue),
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceComputedInt64WithDefaultAndValidators returns a computed int64 attribute with default value and validators
func ResourceComputedInt64WithDefaultAndValidators(description string, defaultValue int64, validators ...validator.Int64) resourceschema.Int64Attribute {
	attr := resourceschema.Int64Attribute{
		MarkdownDescription: description,
		Computed:            true,
		Default:             int64default.StaticInt64(defaultValue),
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ============================================================================
// Resource Validators - Float64 Attributes
// ============================================================================

// ResourceRequiredFloat64WithRange returns a required float64 attribute with range validation
func ResourceRequiredFloat64WithRange(description string, minValue, maxValue float64) resourceschema.Float64Attribute {
	return resourceschema.Float64Attribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.Float64{
			float64validator.Between(minValue, maxValue),
		},
	}
}

// ResourceOptionalFloat64WithRange returns an optional float64 attribute with range validation
func ResourceOptionalFloat64WithRange(description string, minValue, maxValue float64) resourceschema.Float64Attribute {
	return resourceschema.Float64Attribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.Float64{
			float64validator.Between(minValue, maxValue),
		},
	}
}

// ResourceComputedFloat64WithRange returns a computed float64 attribute with range validation
func ResourceComputedFloat64WithRange(description string, minValue, maxValue float64) resourceschema.Float64Attribute {
	return resourceschema.Float64Attribute{
		MarkdownDescription: description,
		Computed:            true,
		Validators: []validator.Float64{
			float64validator.Between(minValue, maxValue),
		},
	}
}

// ResourceFloat64WithValidators returns a required float64 attribute with custom validators
func ResourceFloat64WithValidators(description string, validators ...validator.Float64) resourceschema.Float64Attribute {
	attr := resourceschema.Float64Attribute{
		MarkdownDescription: description,
		Required:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceOptionalFloat64WithValidators returns an optional float64 attribute with custom validators
func ResourceOptionalFloat64WithValidators(description string, validators ...validator.Float64) resourceschema.Float64Attribute {
	attr := resourceschema.Float64Attribute{
		MarkdownDescription: description,
		Optional:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ResourceRequiredFloat64WithValidators returns a required float64 attribute with custom validators
func ResourceRequiredFloat64WithValidators(description string, validators ...validator.Float64) resourceschema.Float64Attribute {
	return ResourceFloat64WithValidators(description, validators...)
}

// ResourceComputedFloat64WithValidators returns a computed float64 attribute with custom validators
func ResourceComputedFloat64WithValidators(description string, validators ...validator.Float64) resourceschema.Float64Attribute {
	attr := resourceschema.Float64Attribute{
		MarkdownDescription: description,
		Computed:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ============================================================================
// Data Source Validators - String Attributes
// ============================================================================

// DataSourceRequiredStringWithRegex returns a required string attribute with regex validation for data sources
func DataSourceRequiredStringWithRegex(description string, pattern *regexp.Regexp, errorMsg string) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
		},
	}
}

// DataSourceOptionalStringWithRegex returns an optional string attribute with regex validation for data sources
func DataSourceOptionalStringWithRegex(description string, pattern *regexp.Regexp, errorMsg string) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
		},
	}
}

// DataSourceComputedStringWithRegex returns a computed string attribute with regex validation for data sources
func DataSourceComputedStringWithRegex(description string, pattern *regexp.Regexp, errorMsg string) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
		},
	}
}

// DataSourceStringWithValidators returns a required string attribute with custom validators for data sources
func DataSourceStringWithValidators(description string, validators ...validator.String) datasourceschema.StringAttribute {
	attr := datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceOptionalStringWithValidators returns an optional string attribute with custom validators for data sources
func DataSourceOptionalStringWithValidators(description string, validators ...validator.String) datasourceschema.StringAttribute {
	attr := datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceRequiredStringWithValidators returns a required string attribute with custom validators for data sources
func DataSourceRequiredStringWithValidators(description string, validators ...validator.String) datasourceschema.StringAttribute {
	return DataSourceStringWithValidators(description, validators...)
}

// DataSourceComputedStringWithValidators returns a computed string attribute with custom validators for data sources
func DataSourceComputedStringWithValidators(description string, validators ...validator.String) datasourceschema.StringAttribute {
	attr := datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Computed:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceRequiredStringWithLengthBetween returns a required string attribute with length validation for data sources
func DataSourceRequiredStringWithLengthBetween(description string, minLength, maxLength int) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// DataSourceOptionalStringWithLengthBetween returns an optional string attribute with length validation for data sources
func DataSourceOptionalStringWithLengthBetween(description string, minLength, maxLength int) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// DataSourceRequiredStringWithLengthAtLeast returns a required string attribute with minimum length validation for data sources
func DataSourceRequiredStringWithLengthAtLeast(description string, minLength int) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// DataSourceOptionalStringWithLengthAtLeast returns an optional string attribute with minimum length validation for data sources
func DataSourceOptionalStringWithLengthAtLeast(description string, minLength int) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(minLength),
		},
	}
}

// DataSourceRequiredStringWithLengthAtMost returns a required string attribute with maximum length validation for data sources
func DataSourceRequiredStringWithLengthAtMost(description string, maxLength int) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.LengthAtMost(maxLength),
		},
	}
}

// DataSourceOptionalStringWithLengthAtMost returns an optional string attribute with maximum length validation for data sources
func DataSourceOptionalStringWithLengthAtMost(description string, maxLength int) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.LengthAtMost(maxLength),
		},
	}
}

// DataSourceRequiredStringWithRegexAndLength returns a required string attribute with regex and length validation for data sources
func DataSourceRequiredStringWithRegexAndLength(description string, pattern *regexp.Regexp, errorMsg string, minLength, maxLength int) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// DataSourceOptionalStringWithRegexAndLength returns an optional string attribute with regex and length validation for data sources
func DataSourceOptionalStringWithRegexAndLength(description string, pattern *regexp.Regexp, errorMsg string, minLength, maxLength int) datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(pattern, errorMsg),
			stringvalidator.LengthBetween(minLength, maxLength),
		},
	}
}

// ============================================================================
// Data Source Validators - Integer Attributes
// ============================================================================

// DataSourceRequiredInt32WithRange returns a required int32 attribute with range validation for data sources
func DataSourceRequiredInt32WithRange(description string, minValue, maxValue int32) datasourceschema.Int32Attribute {
	return datasourceschema.Int32Attribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.Int32{
			int32validator.Between(minValue, maxValue),
		},
	}
}

// DataSourceOptionalInt32WithRange returns an optional int32 attribute with range validation for data sources
func DataSourceOptionalInt32WithRange(description string, minValue, maxValue int32) datasourceschema.Int32Attribute {
	return datasourceschema.Int32Attribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.Int32{
			int32validator.Between(minValue, maxValue),
		},
	}
}

// DataSourceComputedInt32WithRange returns a computed int32 attribute with range validation for data sources
func DataSourceComputedInt32WithRange(description string, minValue, maxValue int32) datasourceschema.Int32Attribute {
	return datasourceschema.Int32Attribute{
		MarkdownDescription: description,
		Computed:            true,
		Validators: []validator.Int32{
			int32validator.Between(minValue, maxValue),
		},
	}
}

// DataSourceInt32WithValidators returns a required int32 attribute with custom validators for data sources
func DataSourceInt32WithValidators(description string, validators ...validator.Int64) datasourceschema.Int64Attribute {
	attr := datasourceschema.Int64Attribute{
		MarkdownDescription: description,
		Required:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceOptionalInt32WithValidators returns an optional int32 attribute with custom validators for data sources
func DataSourceOptionalInt32WithValidators(description string, validators ...validator.Int64) datasourceschema.Int64Attribute {
	attr := datasourceschema.Int64Attribute{
		MarkdownDescription: description,
		Optional:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceRequiredInt32WithValidators returns a required int32 attribute with custom validators for data sources
func DataSourceRequiredInt32WithValidators(description string, validators ...validator.Int64) datasourceschema.Int64Attribute {
	return DataSourceInt32WithValidators(description, validators...)
}

// DataSourceComputedInt32WithValidators returns a computed int32 attribute with custom validators for data sources
func DataSourceComputedInt32WithValidators(description string, validators ...validator.Int64) datasourceschema.Int64Attribute {
	attr := datasourceschema.Int64Attribute{
		MarkdownDescription: description,
		Computed:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceRequiredInt64WithRange returns a required int64 attribute with range validation for data sources
func DataSourceRequiredInt64WithRange(description string, minValue, maxValue int64) datasourceschema.Int64Attribute {
	return datasourceschema.Int64Attribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.Int64{
			int64validator.Between(minValue, maxValue),
		},
	}
}

// DataSourceOptionalInt64WithRange returns an optional int64 attribute with range validation for data sources
func DataSourceOptionalInt64WithRange(description string, minValue, maxValue int64) datasourceschema.Int64Attribute {
	return datasourceschema.Int64Attribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.Int64{
			int64validator.Between(minValue, maxValue),
		},
	}
}

// DataSourceComputedInt64WithRange returns a computed int64 attribute with range validation for data sources
func DataSourceComputedInt64WithRange(description string, minValue, maxValue int64) datasourceschema.Int64Attribute {
	return datasourceschema.Int64Attribute{
		MarkdownDescription: description,
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.Between(minValue, maxValue),
		},
	}
}

// DataSourceInt64WithValidators returns a required int64 attribute with custom validators for data sources
func DataSourceInt64WithValidators(description string, validators ...validator.Int64) datasourceschema.Int64Attribute {
	attr := datasourceschema.Int64Attribute{
		MarkdownDescription: description,
		Required:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceOptionalInt64WithValidators returns an optional int64 attribute with custom validators for data sources
func DataSourceOptionalInt64WithValidators(description string, validators ...validator.Int64) datasourceschema.Int64Attribute {
	attr := datasourceschema.Int64Attribute{
		MarkdownDescription: description,
		Optional:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceRequiredInt64WithValidators returns a required int64 attribute with custom validators for data sources
func DataSourceRequiredInt64WithValidators(description string, validators ...validator.Int64) datasourceschema.Int64Attribute {
	return DataSourceInt64WithValidators(description, validators...)
}

// DataSourceComputedInt64WithValidators returns a computed int64 attribute with custom validators for data sources
func DataSourceComputedInt64WithValidators(description string, validators ...validator.Int64) datasourceschema.Int64Attribute {
	attr := datasourceschema.Int64Attribute{
		MarkdownDescription: description,
		Computed:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// ============================================================================
// Data Source Validators - Float64 Attributes
// ============================================================================

// DataSourceRequiredFloat64WithRange returns a required float64 attribute with range validation for data sources
func DataSourceRequiredFloat64WithRange(description string, minValue, maxValue float64) datasourceschema.Float64Attribute {
	return datasourceschema.Float64Attribute{
		MarkdownDescription: description,
		Required:            true,
		Validators: []validator.Float64{
			float64validator.Between(minValue, maxValue),
		},
	}
}

// DataSourceOptionalFloat64WithRange returns an optional float64 attribute with range validation for data sources
func DataSourceOptionalFloat64WithRange(description string, minValue, maxValue float64) datasourceschema.Float64Attribute {
	return datasourceschema.Float64Attribute{
		MarkdownDescription: description,
		Optional:            true,
		Validators: []validator.Float64{
			float64validator.Between(minValue, maxValue),
		},
	}
}

// DataSourceComputedFloat64WithRange returns a computed float64 attribute with range validation for data sources
func DataSourceComputedFloat64WithRange(description string, minValue, maxValue float64) datasourceschema.Float64Attribute {
	return datasourceschema.Float64Attribute{
		MarkdownDescription: description,
		Computed:            true,
		Validators: []validator.Float64{
			float64validator.Between(minValue, maxValue),
		},
	}
}

// DataSourceFloat64WithValidators returns a required float64 attribute with custom validators for data sources
func DataSourceFloat64WithValidators(description string, validators ...validator.Float64) datasourceschema.Float64Attribute {
	attr := datasourceschema.Float64Attribute{
		MarkdownDescription: description,
		Required:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceOptionalFloat64WithValidators returns an optional float64 attribute with custom validators for data sources
func DataSourceOptionalFloat64WithValidators(description string, validators ...validator.Float64) datasourceschema.Float64Attribute {
	attr := datasourceschema.Float64Attribute{
		MarkdownDescription: description,
		Optional:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}

// DataSourceRequiredFloat64WithValidators returns a required float64 attribute with custom validators for data sources
func DataSourceRequiredFloat64WithValidators(description string, validators ...validator.Float64) datasourceschema.Float64Attribute {
	return DataSourceFloat64WithValidators(description, validators...)
}

// DataSourceComputedFloat64WithValidators returns a computed float64 attribute with custom validators for data sources
func DataSourceComputedFloat64WithValidators(description string, validators ...validator.Float64) datasourceschema.Float64Attribute {
	attr := datasourceschema.Float64Attribute{
		MarkdownDescription: description,
		Computed:            true,
	}
	if len(validators) > 0 {
		attr.Validators = validators
	}
	return attr
}
