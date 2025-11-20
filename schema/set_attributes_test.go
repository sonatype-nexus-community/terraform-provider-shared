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
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Test basic set attribute functions
func TestResourceRequiredStringSet(t *testing.T) {
	attr := ResourceRequiredStringSet("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredStringSet should return required attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceRequiredStringSet should have StringType element")
	}
}

func TestResourceOptionalStringSet(t *testing.T) {
	attr := ResourceOptionalStringSet("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalStringSet should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceOptionalStringSet should have StringType element")
	}
}

func TestResourceComputedStringSet(t *testing.T) {
	attr := ResourceComputedStringSet("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedStringSet should return computed attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceComputedStringSet should have StringType element")
	}
}

func TestResourceComputedOptionalStringSet(t *testing.T) {
	attr := ResourceComputedOptionalStringSet("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalStringSet should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalStringSet should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceComputedOptionalStringSet should have StringType element")
	}
}

// Test set attribute functions with validators
func TestResourceRequiredStringSetWithValidator(t *testing.T) {
	mockValidator := &mockSetValidator{}
	attr := ResourceRequiredStringSetWithValidator(
		"test description",
		mockValidator,
	)

	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredStringSetWithValidator should return required attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceRequiredStringSetWithValidator should have StringType element")
	}
	if len(attr.Validators) != 1 {
		t.Fatal("ResourceRequiredStringSetWithValidator should have 1 validator")
	}
}

func TestResourceOptionalStringSetWithValidator(t *testing.T) {
	mockValidator := &mockSetValidator{}
	attr := ResourceOptionalStringSetWithValidator(
		"test description",
		mockValidator,
	)

	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalStringSetWithValidator should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceOptionalStringSetWithValidator should have StringType element")
	}
	if len(attr.Validators) != 1 {
		t.Fatal("ResourceOptionalStringSetWithValidator should have 1 validator")
	}
}

func TestResourceComputedStringSetWithValidator(t *testing.T) {
	mockValidator := &mockSetValidator{}
	attr := ResourceComputedStringSetWithValidator(
		"test description",
		mockValidator,
	)

	if !attr.IsComputed() {
		t.Fatal("ResourceComputedStringSetWithValidator should return computed attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceComputedStringSetWithValidator should have StringType element")
	}
	if len(attr.Validators) != 1 {
		t.Fatal("ResourceComputedStringSetWithValidator should have 1 validator")
	}
}

func TestResourceRequiredStringSetWithMultipleValidators(t *testing.T) {
	mockValidator1 := &mockSetValidator{}
	mockValidator2 := &mockSetValidator{}
	attr := ResourceRequiredStringSetWithValidator(
		"test description",
		mockValidator1,
		mockValidator2,
	)

	if len(attr.Validators) != 2 {
		t.Fatal("ResourceRequiredStringSetWithValidator should have 2 validators")
	}
}

// Test Int64 set functions with validators
func TestResourceRequiredInt64SetWithValidator(t *testing.T) {
	mockValidator := &mockSetValidator{}
	attr := ResourceRequiredInt64SetWithValidator(
		"test description",
		mockValidator,
	)

	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt64SetWithValidator should return required attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceRequiredInt64SetWithValidator should have Int64Type element")
	}
	if len(attr.Validators) != 1 {
		t.Fatal("ResourceRequiredInt64SetWithValidator should have 1 validator")
	}
}

func TestResourceOptionalInt64SetWithValidator(t *testing.T) {
	mockValidator := &mockSetValidator{}
	attr := ResourceOptionalInt64SetWithValidator(
		"test description",
		mockValidator,
	)

	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt64SetWithValidator should return optional attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceOptionalInt64SetWithValidator should have Int64Type element")
	}
	if len(attr.Validators) != 1 {
		t.Fatal("ResourceOptionalInt64SetWithValidator should have 1 validator")
	}
}

func TestResourceComputedInt64SetWithValidator(t *testing.T) {
	mockValidator := &mockSetValidator{}
	attr := ResourceComputedInt64SetWithValidator(
		"test description",
		mockValidator,
	)

	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt64SetWithValidator should return computed attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceComputedInt64SetWithValidator should have Int64Type element")
	}
	if len(attr.Validators) != 1 {
		t.Fatal("ResourceComputedInt64SetWithValidator should have 1 validator")
	}
}

// Test Bool set functions with validators
func TestResourceRequiredBoolSetWithValidator(t *testing.T) {
	mockValidator := &mockSetValidator{}
	attr := ResourceRequiredBoolSetWithValidator(
		"test description",
		mockValidator,
	)

	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredBoolSetWithValidator should return required attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceRequiredBoolSetWithValidator should have BoolType element")
	}
	if len(attr.Validators) != 1 {
		t.Fatal("ResourceRequiredBoolSetWithValidator should have 1 validator")
	}
}

func TestResourceOptionalBoolSetWithValidator(t *testing.T) {
	mockValidator := &mockSetValidator{}
	attr := ResourceOptionalBoolSetWithValidator(
		"test description",
		mockValidator,
	)

	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalBoolSetWithValidator should return optional attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceOptionalBoolSetWithValidator should have BoolType element")
	}
	if len(attr.Validators) != 1 {
		t.Fatal("ResourceOptionalBoolSetWithValidator should have 1 validator")
	}
}

func TestResourceComputedBoolSetWithValidator(t *testing.T) {
	mockValidator := &mockSetValidator{}
	attr := ResourceComputedBoolSetWithValidator(
		"test description",
		mockValidator,
	)

	if !attr.IsComputed() {
		t.Fatal("ResourceComputedBoolSetWithValidator should return computed attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceComputedBoolSetWithValidator should have BoolType element")
	}
	if len(attr.Validators) != 1 {
		t.Fatal("ResourceComputedBoolSetWithValidator should have 1 validator")
	}
}

// Mock validator for testing
type mockSetValidator struct{}

func (m *mockSetValidator) Description(ctx context.Context) string {
	return "mock validator"
}

func (m *mockSetValidator) MarkdownDescription(ctx context.Context) string {
	return "mock validator"
}

func (m *mockSetValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	// Mock implementation
}
