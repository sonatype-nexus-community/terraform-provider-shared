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
)

// Resource String Validator Tests

func TestResourceRequiredStringWithRegex(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := ResourceRequiredStringWithRegex("test description", pattern, "must be alphanumeric")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredStringWithRegex should return a required attribute")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestResourceOptionalStringWithRegex(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := ResourceOptionalStringWithRegex("test description", pattern, "must be alphanumeric")

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalStringWithRegex should return an optional attribute")
	}
}

func TestResourceComputedStringWithRegex(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := ResourceComputedStringWithRegex("test description", pattern, "must be alphanumeric")

	if !attr.IsComputed() {
		t.Fatal("ResourceComputedStringWithRegex should return a computed attribute")
	}
}

func TestResourceStringWithValidators(t *testing.T) {
	attr := ResourceStringWithValidators("test description")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceStringWithValidators should return a required attribute by default")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestResourceOptionalStringWithValidators(t *testing.T) {
	attr := ResourceOptionalStringWithValidators("test description")

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalStringWithValidators should return an optional attribute")
	}
}

func TestResourceRequiredStringWithValidators(t *testing.T) {
	attr := ResourceRequiredStringWithValidators("test description")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredStringWithValidators should return a required attribute")
	}
}

func TestResourceComputedStringWithValidators(t *testing.T) {
	attr := ResourceComputedStringWithValidators("test description")

	if !attr.IsComputed() {
		t.Fatal("ResourceComputedStringWithValidators should return a computed attribute")
	}
}

func TestResourceRequiredStringWithLengthBetween(t *testing.T) {
	attr := ResourceRequiredStringWithLengthBetween("test description", 1, 100)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredStringWithLengthBetween should return a required attribute")
	}
}

func TestResourceOptionalStringWithLengthBetween(t *testing.T) {
	attr := ResourceOptionalStringWithLengthBetween("test description", 1, 100)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalStringWithLengthBetween should return an optional attribute")
	}
}

func TestResourceRequiredStringWithLengthAtLeast(t *testing.T) {
	attr := ResourceRequiredStringWithLengthAtLeast("test description", 1)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredStringWithLengthAtLeast should return a required attribute")
	}
}

func TestResourceOptionalStringWithLengthAtLeast(t *testing.T) {
	attr := ResourceOptionalStringWithLengthAtLeast("test description", 1)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalStringWithLengthAtLeast should return an optional attribute")
	}
}

func TestResourceRequiredStringWithLengthAtMost(t *testing.T) {
	attr := ResourceRequiredStringWithLengthAtMost("test description", 100)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredStringWithLengthAtMost should return a required attribute")
	}
}

func TestResourceOptionalStringWithLengthAtMost(t *testing.T) {
	attr := ResourceOptionalStringWithLengthAtMost("test description", 100)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalStringWithLengthAtMost should return an optional attribute")
	}
}

func TestResourceOptionalSensitiveStringWithLengthAtLeast(t *testing.T) {
	attr := ResourceOptionalSensitiveStringWithLengthAtLeast("test description", 1)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalSensitiveStringWithLengthAtLeast should return an optional attribute")
	}
}

func TestResourceRequiredSensitiveStringWithLengthAtLeast(t *testing.T) {
	attr := ResourceRequiredSensitiveStringWithLengthAtLeast("test description", 1)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredSensitiveStringWithLengthAtLeast should return a required attribute")
	}
}

func TestResourceRequiredStringWithRegexAndLength(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := ResourceRequiredStringWithRegexAndLength("test description", pattern, "must be alphanumeric", 1, 100)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredStringWithRegexAndLength should return a required attribute")
	}
}

func TestResourceOptionalStringWithRegexAndLength(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := ResourceOptionalStringWithRegexAndLength("test description", pattern, "must be alphanumeric", 1, 100)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalStringWithRegexAndLength should return an optional attribute")
	}
}

// Resource Int64 Validator Tests

func TestResourceRequiredInt64WithRange(t *testing.T) {
	attr := ResourceRequiredInt64WithRange("test description", 1, 100)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredInt64WithRange should return a required attribute")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestResourceOptionalInt64WithRange(t *testing.T) {
	attr := ResourceOptionalInt64WithRange("test description", 1, 100)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalInt64WithRange should return an optional attribute")
	}
}

func TestResourceComputedInt64WithRange(t *testing.T) {
	attr := ResourceComputedInt64WithRange("test description", 1, 100)

	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt64WithRange should return a computed attribute")
	}
}

func TestResourceInt64WithValidators(t *testing.T) {
	attr := ResourceInt64WithValidators("test description")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceInt64WithValidators should return a required attribute by default")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestResourceOptionalInt64WithValidators(t *testing.T) {
	attr := ResourceOptionalInt64WithValidators("test description")

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalInt64WithValidators should return an optional attribute")
	}
}

func TestResourceRequiredInt64WithValidators(t *testing.T) {
	attr := ResourceRequiredInt64WithValidators("test description")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredInt64WithValidators should return a required attribute")
	}
}

func TestResourceComputedInt64WithValidators(t *testing.T) {
	attr := ResourceComputedInt64WithValidators("test description")

	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt64WithValidators should return a computed attribute")
	}
}

// DataSource String Validator Tests

func TestDataSourceRequiredStringWithRegex(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := DataSourceRequiredStringWithRegex("test description", pattern, "must be alphanumeric")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("DataSourceRequiredStringWithRegex should return a required attribute")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestDataSourceOptionalStringWithRegex(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := DataSourceOptionalStringWithRegex("test description", pattern, "must be alphanumeric")

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("DataSourceOptionalStringWithRegex should return an optional attribute")
	}
}

func TestDataSourceComputedStringWithRegex(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := DataSourceComputedStringWithRegex("test description", pattern, "must be alphanumeric")

	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedStringWithRegex should return a computed attribute")
	}
}

func TestDataSourceStringWithValidators(t *testing.T) {
	attr := DataSourceStringWithValidators("test description")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("DataSourceStringWithValidators should return a required attribute by default")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestDataSourceRequiredInt64WithRange(t *testing.T) {
	attr := DataSourceRequiredInt64WithRange("test description", 1, 100)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("DataSourceRequiredInt64WithRange should return a required attribute")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestDataSourceOptionalInt64WithRange(t *testing.T) {
	attr := DataSourceOptionalInt64WithRange("test description", 1, 100)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("DataSourceOptionalInt64WithRange should return an optional attribute")
	}
}
