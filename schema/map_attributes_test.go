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

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Test basic map attribute functions - String Maps
func TestResourceRequiredStringMap(t *testing.T) {
	attr := ResourceRequiredStringMap("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredStringMap should return required attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceRequiredStringMap should have StringType element")
	}
}

func TestResourceOptionalStringMap(t *testing.T) {
	attr := ResourceOptionalStringMap("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalStringMap should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceOptionalStringMap should have StringType element")
	}
}

func TestResourceComputedStringMap(t *testing.T) {
	attr := ResourceComputedStringMap("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedStringMap should return computed attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceComputedStringMap should have StringType element")
	}
}

func TestResourceComputedOptionalStringMap(t *testing.T) {
	attr := ResourceComputedOptionalStringMap("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalStringMap should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalStringMap should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceComputedOptionalStringMap should have StringType element")
	}
}

// Test basic map attribute functions - Int64 Maps
func TestResourceRequiredInt64Map(t *testing.T) {
	attr := ResourceRequiredInt64Map("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt64Map should return required attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceRequiredInt64Map should have Int64Type element")
	}
}

func TestResourceOptionalInt64Map(t *testing.T) {
	attr := ResourceOptionalInt64Map("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt64Map should return optional attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceOptionalInt64Map should have Int64Type element")
	}
}

func TestResourceComputedInt64Map(t *testing.T) {
	attr := ResourceComputedInt64Map("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt64Map should return computed attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceComputedInt64Map should have Int64Type element")
	}
}

func TestResourceComputedOptionalInt64Map(t *testing.T) {
	attr := ResourceComputedOptionalInt64Map("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt64Map should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt64Map should return optional attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceComputedOptionalInt64Map should have Int64Type element")
	}
}

// Test basic map attribute functions - Bool Maps
func TestResourceRequiredBoolMap(t *testing.T) {
	attr := ResourceRequiredBoolMap("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredBoolMap should return required attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceRequiredBoolMap should have BoolType element")
	}
}

func TestResourceOptionalBoolMap(t *testing.T) {
	attr := ResourceOptionalBoolMap("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalBoolMap should return optional attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceOptionalBoolMap should have BoolType element")
	}
}

func TestResourceComputedBoolMap(t *testing.T) {
	attr := ResourceComputedBoolMap("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedBoolMap should return computed attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceComputedBoolMap should have BoolType element")
	}
}

func TestResourceComputedOptionalBoolMap(t *testing.T) {
	attr := ResourceComputedOptionalBoolMap("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalBoolMap should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalBoolMap should return optional attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceComputedOptionalBoolMap should have BoolType element")
	}
}

// Test data source map attribute functions - String Maps
func TestDataSourceOptionalStringMap(t *testing.T) {
	attr := DataSourceOptionalStringMap("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalStringMap should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("DataSourceOptionalStringMap should have StringType element")
	}
}

func TestDataSourceComputedStringMap(t *testing.T) {
	attr := DataSourceComputedStringMap("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedStringMap should return computed attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("DataSourceComputedStringMap should have StringType element")
	}
}

// Test data source map attribute functions - Int64 Maps
func TestDataSourceOptionalInt64Map(t *testing.T) {
	attr := DataSourceOptionalInt64Map("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalInt64Map should return optional attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("DataSourceOptionalInt64Map should have Int64Type element")
	}
}

func TestDataSourceComputedInt64Map(t *testing.T) {
	attr := DataSourceComputedInt64Map("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedInt64Map should return computed attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("DataSourceComputedInt64Map should have Int64Type element")
	}
}

// Test data source map attribute functions - Bool Maps
func TestDataSourceOptionalBoolMap(t *testing.T) {
	attr := DataSourceOptionalBoolMap("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalBoolMap should return optional attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("DataSourceOptionalBoolMap should have BoolType element")
	}
}

func TestDataSourceComputedBoolMap(t *testing.T) {
	attr := DataSourceComputedBoolMap("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedBoolMap should return computed attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("DataSourceComputedBoolMap should have BoolType element")
	}
}
