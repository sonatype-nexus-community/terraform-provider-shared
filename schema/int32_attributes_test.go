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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Test ResourceOptionalInt32WithPlanModifier
func TestResourceOptionalInt32WithPlanModifier(t *testing.T) {
	attr := ResourceOptionalInt32WithPlanModifier("description", int32planmodifier.RequiresReplace())
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt32WithPlanModifier should return optional attribute")
	}
}

// Test ResourceOptionalInt32WithDefaultAndPlanModifier
func TestResourceOptionalInt32WithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceOptionalInt32WithDefaultAndPlanModifier("description", 42, int32planmodifier.RequiresReplace())
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt32WithDefaultAndPlanModifier should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceOptionalInt32WithDefaultAndPlanModifier should return computed attribute")
	}
}

// Test ResourceComputedOptionalInt32WithPlanModifier
func TestResourceComputedOptionalInt32WithPlanModifier(t *testing.T) {
	attr := ResourceComputedOptionalInt32WithPlanModifier("description", int32planmodifier.UseStateForUnknown())
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt32WithPlanModifier should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt32WithPlanModifier should return computed attribute")
	}
}

// Test ResourceComputedOptionalInt32WithDefaultAndPlanModifier
func TestResourceComputedOptionalInt32WithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceComputedOptionalInt32WithDefaultAndPlanModifier("description", 99, int32planmodifier.UseStateForUnknown())
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt32WithDefaultAndPlanModifier should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt32WithDefaultAndPlanModifier should return computed attribute")
	}
}

// Test ResourceOptionalInt32WithValidator
func TestResourceOptionalInt32WithValidator(t *testing.T) {
	attr := ResourceOptionalInt32WithValidator("description", int32validator.Between(0, 100))
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt32WithValidator should return optional attribute")
	}
}

// Test ResourceRequiredInt32WithValidator
func TestResourceRequiredInt32WithValidator(t *testing.T) {
	attr := ResourceRequiredInt32WithValidator("description", int32validator.Between(0, 100))
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt32WithValidator should return required attribute")
	}
}

// Test ResourceOptionalInt32WithDefaultAndValidator
func TestResourceOptionalInt32WithDefaultAndValidator(t *testing.T) {
	attr := ResourceOptionalInt32WithDefaultAndValidator("description", 50, int32validator.Between(0, 100))
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt32WithDefaultAndValidator should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceOptionalInt32WithDefaultAndValidator should return computed attribute")
	}
}

// Test ResourceComputedOptionalInt32WithDefaultAndValidator
func TestResourceComputedOptionalInt32WithDefaultAndValidator(t *testing.T) {
	attr := ResourceComputedOptionalInt32WithDefaultAndValidator("description", 75, int32validator.Between(0, 100))
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt32WithDefaultAndValidator should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt32WithDefaultAndValidator should return computed attribute")
	}
}
