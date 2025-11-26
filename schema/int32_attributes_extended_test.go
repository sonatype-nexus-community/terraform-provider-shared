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

// Test int32 attribute functions

func TestResourceRequiredInt32(t *testing.T) {
	attr := ResourceRequiredInt32("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt32 should return required attribute")
	}
}

func TestResourceOptionalInt32(t *testing.T) {
	attr := ResourceOptionalInt32("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt32 should return optional attribute")
	}
}

func TestResourceComputedInt32(t *testing.T) {
	attr := ResourceComputedInt32("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt32 should return computed attribute")
	}
}

func TestResourceComputedInt32WithDefault(t *testing.T) {
	attr := ResourceComputedInt32WithDefault("test description", 42)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt32WithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedInt32WithDefault should have default value")
	}
}

func TestResourceOptionalInt32WithDefault(t *testing.T) {
	attr := ResourceOptionalInt32WithDefault("test description", 42)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt32WithDefault should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceOptionalInt32WithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceOptionalInt32WithDefault should have default value")
	}
}

func TestResourceComputedOptionalInt32(t *testing.T) {
	attr := ResourceComputedOptionalInt32("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt32 should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt32 should return optional attribute")
	}
}

func TestResourceComputedOptionalInt32WithDefault(t *testing.T) {
	attr := ResourceComputedOptionalInt32WithDefault("test description", 42)
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt32WithDefault should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt32WithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedOptionalInt32WithDefault should have default value")
	}
}

func TestResourceRequiredInt32WithDefault(t *testing.T) {
	attr := ResourceRequiredInt32WithDefault("test description", 42)
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt32WithDefault should return required attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceRequiredInt32WithDefault should have default value")
	}
}

func TestDataSourceOptionalInt32(t *testing.T) {
	attr := DataSourceOptionalInt32("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalInt32 should return optional attribute")
	}
}

func TestDataSourceComputedInt32(t *testing.T) {
	attr := DataSourceComputedInt32("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedInt32 should return computed attribute")
	}
}
