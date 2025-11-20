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

// Test bool attribute functions
func TestResourceRequiredBool(t *testing.T) {
	attr := ResourceRequiredBool("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredBool should return required attribute")
	}
}

func TestResourceOptionalBool(t *testing.T) {
	attr := ResourceOptionalBool("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalBool should return optional attribute")
	}
}

func TestResourceComputedBool(t *testing.T) {
	attr := ResourceComputedBool("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedBool should return computed attribute")
	}
}

func TestResourceOptionalBoolWithDefault(t *testing.T) {
	attr := ResourceOptionalBoolWithDefault("test description", true)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalBoolWithDefault should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceOptionalBoolWithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceOptionalBoolWithDefault should have default value")
	}
}

func TestResourceComputedOptionalBool(t *testing.T) {
	attr := ResourceComputedOptionalBool("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalBool should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalBool should return optional attribute")
	}
	if len(attr.PlanModifiers) == 0 {
		t.Fatal("ResourceComputedOptionalBool should have plan modifiers")
	}
}

func TestResourceRequiredBoolWithDefault(t *testing.T) {
	attr := ResourceRequiredBoolWithDefault("test description", false)
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredBoolWithDefault should return required attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceRequiredBoolWithDefault should have default value")
	}
}

func TestDataSourceOptionalBool(t *testing.T) {
	attr := DataSourceOptionalBool("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalBool should return optional attribute")
	}
}

func TestDataSourceComputedBool(t *testing.T) {
	attr := DataSourceComputedBool("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedBool should return computed attribute")
	}
}
