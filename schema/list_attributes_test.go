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

// Test basic list attribute functions - String Lists
func TestResourceRequiredStringList(t *testing.T) {
	attr := ResourceRequiredStringList("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredStringList should return required attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceRequiredStringList should have StringType element")
	}
}

func TestResourceOptionalStringList(t *testing.T) {
	attr := ResourceOptionalStringList("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalStringList should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceOptionalStringList should have StringType element")
	}
}

func TestResourceComputedStringList(t *testing.T) {
	attr := ResourceComputedStringList("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedStringList should return computed attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceComputedStringList should have StringType element")
	}
}

func TestResourceComputedOptionalStringList(t *testing.T) {
	attr := ResourceComputedOptionalStringList("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalStringList should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalStringList should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("ResourceComputedOptionalStringList should have StringType element")
	}
}

// Test basic list attribute functions - Int64 Lists
func TestResourceRequiredInt64List(t *testing.T) {
	attr := ResourceRequiredInt64List("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt64List should return required attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceRequiredInt64List should have Int64Type element")
	}
}

func TestResourceOptionalInt64List(t *testing.T) {
	attr := ResourceOptionalInt64List("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt64List should return optional attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceOptionalInt64List should have Int64Type element")
	}
}

func TestResourceComputedInt64List(t *testing.T) {
	attr := ResourceComputedInt64List("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt64List should return computed attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceComputedInt64List should have Int64Type element")
	}
}

func TestResourceComputedOptionalInt64List(t *testing.T) {
	attr := ResourceComputedOptionalInt64List("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt64List should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt64List should return optional attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("ResourceComputedOptionalInt64List should have Int64Type element")
	}
}

// Test basic list attribute functions - Bool Lists
func TestResourceRequiredBoolList(t *testing.T) {
	attr := ResourceRequiredBoolList("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredBoolList should return required attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceRequiredBoolList should have BoolType element")
	}
}

func TestResourceOptionalBoolList(t *testing.T) {
	attr := ResourceOptionalBoolList("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalBoolList should return optional attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceOptionalBoolList should have BoolType element")
	}
}

func TestResourceComputedBoolList(t *testing.T) {
	attr := ResourceComputedBoolList("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedBoolList should return computed attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceComputedBoolList should have BoolType element")
	}
}

func TestResourceComputedOptionalBoolList(t *testing.T) {
	attr := ResourceComputedOptionalBoolList("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalBoolList should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalBoolList should return optional attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("ResourceComputedOptionalBoolList should have BoolType element")
	}
}

// Test data source list attribute functions - String Lists
func TestDataSourceOptionalStringList(t *testing.T) {
	attr := DataSourceOptionalStringList("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalStringList should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("DataSourceOptionalStringList should have StringType element")
	}
}

func TestDataSourceComputedStringList(t *testing.T) {
	attr := DataSourceComputedStringList("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedStringList should return computed attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("DataSourceComputedStringList should have StringType element")
	}
}

// Test data source list attribute functions - Int64 Lists
func TestDataSourceOptionalInt64List(t *testing.T) {
	attr := DataSourceOptionalInt64List("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalInt64List should return optional attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("DataSourceOptionalInt64List should have Int64Type element")
	}
}

func TestDataSourceComputedInt64List(t *testing.T) {
	attr := DataSourceComputedInt64List("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedInt64List should return computed attribute")
	}
	if attr.ElementType != types.Int64Type {
		t.Fatal("DataSourceComputedInt64List should have Int64Type element")
	}
}

// Test data source list attribute functions - Bool Lists
func TestDataSourceOptionalBoolList(t *testing.T) {
	attr := DataSourceOptionalBoolList("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalBoolList should return optional attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("DataSourceOptionalBoolList should have BoolType element")
	}
}

func TestDataSourceComputedBoolList(t *testing.T) {
	attr := DataSourceComputedBoolList("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedBoolList should return computed attribute")
	}
	if attr.ElementType != types.BoolType {
		t.Fatal("DataSourceComputedBoolList should have BoolType element")
	}
}
