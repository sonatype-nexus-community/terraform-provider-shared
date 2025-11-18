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

// String Validator Tests

func TestRequiredStringWithRegex(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := RequiredStringWithRegex("test description", pattern, "must be alphanumeric")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredStringWithRegex should return a required attribute")
	}
	if attr.GetDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestOptionalStringWithRegex(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := OptionalStringWithRegex("test description", pattern, "must be alphanumeric")

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalStringWithRegex should return an optional attribute")
	}
}

func TestComputedStringWithRegex(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := ComputedStringWithRegex("test description", pattern, "must be alphanumeric")

	if !attr.IsComputed() {
		t.Fatal("ComputedStringWithRegex should return a computed attribute")
	}
}

func TestStringWithValidators(t *testing.T) {
	attr := StringWithValidators("test description")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("StringWithValidators should return a required attribute by default")
	}
	if attr.GetDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestOptionalStringWithValidators(t *testing.T) {
	attr := OptionalStringWithValidators("test description")

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalStringWithValidators should return an optional attribute")
	}
}

func TestRequiredStringWithValidators(t *testing.T) {
	attr := RequiredStringWithValidators("test description")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredStringWithValidators should return a required attribute")
	}
}

func TestComputedStringWithValidators(t *testing.T) {
	attr := ComputedStringWithValidators("test description")

	if !attr.IsComputed() {
		t.Fatal("ComputedStringWithValidators should return a computed attribute")
	}
}

func TestRequiredStringWithLengthBetween(t *testing.T) {
	attr := RequiredStringWithLengthBetween("test description", 1, 100)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredStringWithLengthBetween should return a required attribute")
	}
}

func TestOptionalStringWithLengthBetween(t *testing.T) {
	attr := OptionalStringWithLengthBetween("test description", 1, 100)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalStringWithLengthBetween should return an optional attribute")
	}
}

func TestRequiredStringWithLengthAtLeast(t *testing.T) {
	attr := RequiredStringWithLengthAtLeast("test description", 1)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredStringWithLengthAtLeast should return a required attribute")
	}
}

func TestOptionalStringWithLengthAtLeast(t *testing.T) {
	attr := OptionalStringWithLengthAtLeast("test description", 1)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalStringWithLengthAtLeast should return an optional attribute")
	}
}

func TestRequiredStringWithLengthAtMost(t *testing.T) {
	attr := RequiredStringWithLengthAtMost("test description", 100)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredStringWithLengthAtMost should return a required attribute")
	}
}

func TestOptionalStringWithLengthAtMost(t *testing.T) {
	attr := OptionalStringWithLengthAtMost("test description", 100)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalStringWithLengthAtMost should return an optional attribute")
	}
}

func TestOptionalSensitiveStringWithLengthAtLeast(t *testing.T) {
	attr := OptionalSensitiveStringWithLengthAtLeast("test description", 1)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalSensitiveStringWithLengthAtLeast should return an optional attribute")
	}
}

func TestRequiredSensitiveStringWithLengthAtLeast(t *testing.T) {
	attr := RequiredSensitiveStringWithLengthAtLeast("test description", 1)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredSensitiveStringWithLengthAtLeast should return a required attribute")
	}
}

func TestRequiredStringWithRegexAndLength(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := RequiredStringWithRegexAndLength("test description", pattern, "must be alphanumeric", 1, 100)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredStringWithRegexAndLength should return a required attribute")
	}
}

func TestOptionalStringWithRegexAndLength(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	attr := OptionalStringWithRegexAndLength("test description", pattern, "must be alphanumeric", 1, 100)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalStringWithRegexAndLength should return an optional attribute")
	}
}

// Int64 Validator Tests

func TestRequiredInt64WithRange(t *testing.T) {
	attr := RequiredInt64WithRange("test description", 1, 100)

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredInt64WithRange should return a required attribute")
	}
	if attr.GetDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestOptionalInt64WithRange(t *testing.T) {
	attr := OptionalInt64WithRange("test description", 1, 100)

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalInt64WithRange should return an optional attribute")
	}
}

func TestComputedInt64WithRange(t *testing.T) {
	attr := ComputedInt64WithRange("test description", 1, 100)

	if !attr.IsComputed() {
		t.Fatal("ComputedInt64WithRange should return a computed attribute")
	}
}

func TestInt64WithValidators(t *testing.T) {
	attr := Int64WithValidators("test description")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("Int64WithValidators should return a required attribute by default")
	}
	if attr.GetDescription() != "test description" {
		t.Fatal("description mismatch")
	}
}

func TestOptionalInt64WithValidators(t *testing.T) {
	attr := OptionalInt64WithValidators("test description")

	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalInt64WithValidators should return an optional attribute")
	}
}

func TestRequiredInt64WithValidators(t *testing.T) {
	attr := RequiredInt64WithValidators("test description")

	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredInt64WithValidators should return a required attribute")
	}
}

func TestComputedInt64WithValidators(t *testing.T) {
	attr := ComputedInt64WithValidators("test description")

	if !attr.IsComputed() {
		t.Fatal("ComputedInt64WithValidators should return a computed attribute")
	}
}
