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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
)

// ========================================
// Resource Int32 with Range Validators
// ========================================

func TestResourceRequiredInt32WithRange(t *testing.T) {
	attr := ResourceRequiredInt32WithRange("required int32 with range", 0, 100)
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt32WithRange should return required attribute")
	}
}

func TestResourceOptionalInt32WithRange(t *testing.T) {
	attr := ResourceOptionalInt32WithRange("optional int32 with range", 0, 50)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt32WithRange should return optional attribute")
	}
}

func TestResourceComputedInt32WithRange(t *testing.T) {
	attr := ResourceComputedInt32WithRange("computed int32 with range", 10, 99)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt32WithRange should return computed attribute")
	}
}

// ========================================
// Resource Int32 with Validators
// ========================================

func TestResourceInt32WithValidators(t *testing.T) {
	attr := ResourceInt32WithValidators("int32 with validators", int64validator.Between(0, 100))
	if !attr.IsRequired() {
		t.Fatal("ResourceInt32WithValidators should return required attribute by default")
	}
}

func TestResourceOptionalInt32WithValidators(t *testing.T) {
	attr := ResourceOptionalInt32WithValidators("optional int32 with validators", int64validator.Between(0, 100))
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt32WithValidators should return optional attribute")
	}
}

func TestResourceRequiredInt32WithValidators(t *testing.T) {
	attr := ResourceRequiredInt32WithValidators("required int32 with validators", int64validator.Between(0, 100))
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt32WithValidators should return required attribute")
	}
}

func TestResourceComputedInt32WithValidators(t *testing.T) {
	attr := ResourceComputedInt32WithValidators("computed int32 with validators", int64validator.Between(0, 100))
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt32WithValidators should return computed attribute")
	}
}

// ========================================
// Resource Float64 with Range Validators
// ========================================

func TestResourceRequiredFloat64WithRange(t *testing.T) {
	attr := ResourceRequiredFloat64WithRange("required float64 with range", 0.0, 100.0)
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredFloat64WithRange should return required attribute")
	}
}

func TestResourceOptionalFloat64WithRange(t *testing.T) {
	attr := ResourceOptionalFloat64WithRange("optional float64 with range", 0.0, 50.5)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalFloat64WithRange should return optional attribute")
	}
}

func TestResourceComputedFloat64WithRange(t *testing.T) {
	attr := ResourceComputedFloat64WithRange("computed float64 with range", 1.0, 99.9)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedFloat64WithRange should return computed attribute")
	}
}

// ========================================
// Resource Float64 with Validators
// ========================================

func TestResourceFloat64WithValidators(t *testing.T) {
	attr := ResourceFloat64WithValidators("float64 with validators", float64validator.Between(0.0, 100.0))
	if !attr.IsRequired() {
		t.Fatal("ResourceFloat64WithValidators should return required attribute by default")
	}
}

func TestResourceOptionalFloat64WithValidators(t *testing.T) {
	attr := ResourceOptionalFloat64WithValidators("optional float64 with validators", float64validator.Between(0.0, 100.0))
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalFloat64WithValidators should return optional attribute")
	}
}

func TestResourceRequiredFloat64WithValidators(t *testing.T) {
	attr := ResourceRequiredFloat64WithValidators("required float64 with validators", float64validator.Between(0.0, 100.0))
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredFloat64WithValidators should return required attribute")
	}
}

func TestResourceComputedFloat64WithValidators(t *testing.T) {
	attr := ResourceComputedFloat64WithValidators("computed float64 with validators", float64validator.Between(0.0, 100.0))
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedFloat64WithValidators should return computed attribute")
	}
}

// ========================================
// DataSource String with Validators
// ========================================

func TestDataSourceOptionalStringWithValidators(t *testing.T) {
	attr := DataSourceOptionalStringWithValidators("optional string with validators", stringvalidator.LengthAtLeast(1))
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalStringWithValidators should return optional attribute")
	}
}

func TestDataSourceRequiredStringWithValidators(t *testing.T) {
	attr := DataSourceRequiredStringWithValidators("required string with validators", stringvalidator.LengthAtLeast(1))
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredStringWithValidators should return required attribute")
	}
}

func TestDataSourceComputedStringWithValidators(t *testing.T) {
	attr := DataSourceComputedStringWithValidators("computed string with validators", stringvalidator.LengthAtLeast(1))
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedStringWithValidators should return computed attribute")
	}
}

// ========================================
// DataSource String with Length Validators
// ========================================

func TestDataSourceRequiredStringWithLengthBetween(t *testing.T) {
	attr := DataSourceRequiredStringWithLengthBetween("required string with length between", 1, 100)
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredStringWithLengthBetween should return required attribute")
	}
}

func TestDataSourceOptionalStringWithLengthBetween(t *testing.T) {
	attr := DataSourceOptionalStringWithLengthBetween("optional string with length between", 1, 100)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalStringWithLengthBetween should return optional attribute")
	}
}

func TestDataSourceRequiredStringWithLengthAtLeast(t *testing.T) {
	attr := DataSourceRequiredStringWithLengthAtLeast("required string with length at least", 1)
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredStringWithLengthAtLeast should return required attribute")
	}
}

func TestDataSourceOptionalStringWithLengthAtLeast(t *testing.T) {
	attr := DataSourceOptionalStringWithLengthAtLeast("optional string with length at least", 1)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalStringWithLengthAtLeast should return optional attribute")
	}
}

func TestDataSourceRequiredStringWithLengthAtMost(t *testing.T) {
	attr := DataSourceRequiredStringWithLengthAtMost("required string with length at most", 100)
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredStringWithLengthAtMost should return required attribute")
	}
}

func TestDataSourceOptionalStringWithLengthAtMost(t *testing.T) {
	attr := DataSourceOptionalStringWithLengthAtMost("optional string with length at most", 100)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalStringWithLengthAtMost should return optional attribute")
	}
}

func TestDataSourceRequiredStringWithRegexAndLength(t *testing.T) {
	pattern := regexp.MustCompile("^[a-z]+$")
	attr := DataSourceRequiredStringWithRegexAndLength("required string with regex and length", pattern, "must be lowercase letters", 1, 100)
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredStringWithRegexAndLength should return required attribute")
	}
}

func TestDataSourceOptionalStringWithRegexAndLength(t *testing.T) {
	pattern := regexp.MustCompile("^[a-z]+$")
	attr := DataSourceOptionalStringWithRegexAndLength("optional string with regex and length", pattern, "must be lowercase letters", 1, 100)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalStringWithRegexAndLength should return optional attribute")
	}
}

// ========================================
// DataSource Int32 with Range Validators
// ========================================

func TestDataSourceRequiredInt32WithRange(t *testing.T) {
	attr := DataSourceRequiredInt32WithRange("required int32 with range", 0, 100)
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredInt32WithRange should return required attribute")
	}
}

func TestDataSourceOptionalInt32WithRange(t *testing.T) {
	attr := DataSourceOptionalInt32WithRange("optional int32 with range", 0, 100)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalInt32WithRange should return optional attribute")
	}
}

func TestDataSourceComputedInt32WithRange(t *testing.T) {
	attr := DataSourceComputedInt32WithRange("computed int32 with range", 0, 100)
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedInt32WithRange should return computed attribute")
	}
}

// ========================================
// DataSource Int32 with Validators
// ========================================

func TestDataSourceInt32WithValidators(t *testing.T) {
	attr := DataSourceInt32WithValidators("int32 with validators", int64validator.Between(0, 100))
	if !attr.IsRequired() {
		t.Fatal("DataSourceInt32WithValidators should return required attribute by default")
	}
}

func TestDataSourceOptionalInt32WithValidators(t *testing.T) {
	attr := DataSourceOptionalInt32WithValidators("optional int32 with validators", int64validator.Between(0, 100))
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalInt32WithValidators should return optional attribute")
	}
}

func TestDataSourceRequiredInt32WithValidators(t *testing.T) {
	attr := DataSourceRequiredInt32WithValidators("required int32 with validators", int64validator.Between(0, 100))
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredInt32WithValidators should return required attribute")
	}
}

func TestDataSourceComputedInt32WithValidators(t *testing.T) {
	attr := DataSourceComputedInt32WithValidators("computed int32 with validators", int64validator.Between(0, 100))
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedInt32WithValidators should return computed attribute")
	}
}

// ========================================
// DataSource Int64 with Range Validators
// ========================================

func TestDataSourceRequiredInt64WithRange(t *testing.T) {
	attr := DataSourceRequiredInt64WithRange("required int64 with range", 0, 100)
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredInt64WithRange should return required attribute")
	}
}

func TestDataSourceOptionalInt64WithRange(t *testing.T) {
	attr := DataSourceOptionalInt64WithRange("optional int64 with range", 0, 100)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalInt64WithRange should return optional attribute")
	}
}

func TestDataSourceComputedInt64WithRange(t *testing.T) {
	attr := DataSourceComputedInt64WithRange("computed int64 with range", 0, 100)
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedInt64WithRange should return computed attribute")
	}
}

// ========================================
// DataSource Int64 with Validators
// ========================================

func TestDataSourceInt64WithValidators(t *testing.T) {
	attr := DataSourceInt64WithValidators("int64 with validators", int64validator.Between(0, 100))
	if !attr.IsRequired() {
		t.Fatal("DataSourceInt64WithValidators should return required attribute by default")
	}
}

func TestDataSourceOptionalInt64WithValidators(t *testing.T) {
	attr := DataSourceOptionalInt64WithValidators("optional int64 with validators", int64validator.Between(0, 100))
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalInt64WithValidators should return optional attribute")
	}
}

func TestDataSourceRequiredInt64WithValidators(t *testing.T) {
	attr := DataSourceRequiredInt64WithValidators("required int64 with validators", int64validator.Between(0, 100))
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredInt64WithValidators should return required attribute")
	}
}

func TestDataSourceComputedInt64WithValidators(t *testing.T) {
	attr := DataSourceComputedInt64WithValidators("computed int64 with validators", int64validator.Between(0, 100))
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedInt64WithValidators should return computed attribute")
	}
}

// ========================================
// DataSource Float64 with Range Validators
// ========================================

func TestDataSourceRequiredFloat64WithRange(t *testing.T) {
	attr := DataSourceRequiredFloat64WithRange("required float64 with range", 0.0, 100.0)
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredFloat64WithRange should return required attribute")
	}
}

func TestDataSourceOptionalFloat64WithRange(t *testing.T) {
	attr := DataSourceOptionalFloat64WithRange("optional float64 with range", 0.0, 100.0)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalFloat64WithRange should return optional attribute")
	}
}

func TestDataSourceComputedFloat64WithRange(t *testing.T) {
	attr := DataSourceComputedFloat64WithRange("computed float64 with range", 0.0, 100.0)
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedFloat64WithRange should return computed attribute")
	}
}

// ========================================
// DataSource Float64 with Validators
// ========================================

func TestDataSourceFloat64WithValidators(t *testing.T) {
	attr := DataSourceFloat64WithValidators("float64 with validators", float64validator.Between(0.0, 100.0))
	if !attr.IsRequired() {
		t.Fatal("DataSourceFloat64WithValidators should return required attribute by default")
	}
}

func TestDataSourceOptionalFloat64WithValidators(t *testing.T) {
	attr := DataSourceOptionalFloat64WithValidators("optional float64 with validators", float64validator.Between(0.0, 100.0))
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalFloat64WithValidators should return optional attribute")
	}
}

func TestDataSourceRequiredFloat64WithValidators(t *testing.T) {
	attr := DataSourceRequiredFloat64WithValidators("required float64 with validators", float64validator.Between(0.0, 100.0))
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredFloat64WithValidators should return required attribute")
	}
}

func TestDataSourceComputedFloat64WithValidators(t *testing.T) {
	attr := DataSourceComputedFloat64WithValidators("computed float64 with validators", float64validator.Between(0.0, 100.0))
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedFloat64WithValidators should return computed attribute")
	}
}
