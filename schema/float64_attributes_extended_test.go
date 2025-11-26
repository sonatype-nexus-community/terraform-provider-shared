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

// Test float64 attribute functions

func TestResourceRequiredFloat64(t *testing.T) {
	attr := ResourceRequiredFloat64("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredFloat64 should return required attribute")
	}
}

func TestResourceOptionalFloat64(t *testing.T) {
	attr := ResourceOptionalFloat64("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalFloat64 should return optional attribute")
	}
}

func TestResourceComputedFloat64(t *testing.T) {
	attr := ResourceComputedFloat64("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedFloat64 should return computed attribute")
	}
}

func TestResourceComputedFloat64WithDefault(t *testing.T) {
	attr := ResourceComputedFloat64WithDefault("test description", 3.14)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedFloat64WithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedFloat64WithDefault should have default value")
	}
}

func TestResourceOptionalFloat64WithDefault(t *testing.T) {
	attr := ResourceOptionalFloat64WithDefault("test description", 3.14)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalFloat64WithDefault should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceOptionalFloat64WithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceOptionalFloat64WithDefault should have default value")
	}
}

func TestResourceComputedOptionalFloat64(t *testing.T) {
	attr := ResourceComputedOptionalFloat64("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalFloat64 should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalFloat64 should return optional attribute")
	}
}

func TestResourceRequiredFloat64WithDefault(t *testing.T) {
	attr := ResourceRequiredFloat64WithDefault("test description", 3.14)
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredFloat64WithDefault should return required attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceRequiredFloat64WithDefault should have default value")
	}
}

func TestResourceComputedOptionalFloat64WithDefault(t *testing.T) {
	attr := ResourceComputedOptionalFloat64WithDefault("test description", 3.14)
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalFloat64WithDefault should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalFloat64WithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedOptionalFloat64WithDefault should have default value")
	}
}

func TestResourceOptionalFloat64WithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceOptionalFloat64WithDefaultAndPlanModifier("test description", 3.14)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalFloat64WithDefaultAndPlanModifier should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceOptionalFloat64WithDefaultAndPlanModifier should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceOptionalFloat64WithDefaultAndPlanModifier should have default value")
	}
}

func TestResourceComputedFloat64WithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceComputedFloat64WithDefaultAndPlanModifier("test description", 3.14)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedFloat64WithDefaultAndPlanModifier should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedFloat64WithDefaultAndPlanModifier should have default value")
	}
}

func TestResourceComputedOptionalFloat64WithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceComputedOptionalFloat64WithDefaultAndPlanModifier("test description", 3.14)
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalFloat64WithDefaultAndPlanModifier should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalFloat64WithDefaultAndPlanModifier should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedOptionalFloat64WithDefaultAndPlanModifier should have default value")
	}
}

func TestDataSourceOptionalFloat64(t *testing.T) {
	attr := DataSourceOptionalFloat64("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalFloat64 should return optional attribute")
	}
}

func TestDataSourceComputedFloat64(t *testing.T) {
	attr := DataSourceComputedFloat64("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedFloat64 should return computed attribute")
	}
}
