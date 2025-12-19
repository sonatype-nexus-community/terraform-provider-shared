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
)

// Test missing int64 set attribute functions
func TestResourceRequiredInt64Set(t *testing.T) {
	attr := ResourceRequiredInt64Set("description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt64Set should return required attribute")
	}
}

func TestResourceOptionalInt64Set(t *testing.T) {
	attr := ResourceOptionalInt64Set("description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt64Set should return optional attribute")
	}
}

func TestResourceComputedInt64Set(t *testing.T) {
	attr := ResourceComputedInt64Set("description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt64Set should return computed attribute")
	}
}

func TestResourceComputedOptionalInt64Set(t *testing.T) {
	attr := ResourceComputedOptionalInt64Set("description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt64Set should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt64Set should return optional attribute")
	}
}

// Test missing bool set attribute functions
func TestResourceRequiredBoolSet(t *testing.T) {
	attr := ResourceRequiredBoolSet("description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredBoolSet should return required attribute")
	}
}

func TestResourceOptionalBoolSet(t *testing.T) {
	attr := ResourceOptionalBoolSet("description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalBoolSet should return optional attribute")
	}
}

func TestResourceComputedBoolSet(t *testing.T) {
	attr := ResourceComputedBoolSet("description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedBoolSet should return computed attribute")
	}
}

func TestResourceComputedOptionalBoolSet(t *testing.T) {
	attr := ResourceComputedOptionalBoolSet("description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalBoolSet should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalBoolSet should return optional attribute")
	}
}

// Test data source set attribute functions
func TestDataSourceRequiredStringSet(t *testing.T) {
	attr := DataSourceRequiredStringSet("description")
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredStringSet should return required attribute")
	}
}

func TestDataSourceOptionalStringSet(t *testing.T) {
	attr := DataSourceOptionalStringSet("description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalStringSet should return optional attribute")
	}
}

func TestDataSourceComputedStringSet(t *testing.T) {
	attr := DataSourceComputedStringSet("description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedStringSet should return computed attribute")
	}
}

func TestDataSourceOptionalInt64Set(t *testing.T) {
	attr := DataSourceOptionalInt64Set("description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalInt64Set should return optional attribute")
	}
}

func TestDataSourceComputedInt64Set(t *testing.T) {
	attr := DataSourceComputedInt64Set("description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedInt64Set should return computed attribute")
	}
}

func TestDataSourceOptionalBoolSet(t *testing.T) {
	attr := DataSourceOptionalBoolSet("description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalBoolSet should return optional attribute")
	}
}

func TestDataSourceComputedBoolSet(t *testing.T) {
	attr := DataSourceComputedBoolSet("description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedBoolSet should return computed attribute")
	}
}
