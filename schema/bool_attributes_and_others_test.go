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

// Test missing bool attribute functions
func TestResourceComputedBoolWithDefault(t *testing.T) {
	attr := ResourceComputedBoolWithDefault("description", true)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedBoolWithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedBoolWithDefault should have default value")
	}
}

func TestResourceComputedOptionalBoolWithDefault(t *testing.T) {
	attr := ResourceComputedOptionalBoolWithDefault("description", false)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalBoolWithDefault should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalBoolWithDefault should return optional attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedOptionalBoolWithDefault should have default value")
	}
}

func TestResourceOptionalBoolWithPlanModifier(t *testing.T) {
	attr := ResourceOptionalBoolWithPlanModifier("description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalBoolWithPlanModifier should return optional attribute")
	}
}

func TestResourceOptionalBoolWithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceOptionalBoolWithDefaultAndPlanModifier("description", true)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalBoolWithDefaultAndPlanModifier should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceOptionalBoolWithDefaultAndPlanModifier should return computed attribute")
	}
}

func TestResourceComputedOptionalBoolWithPlanModifier(t *testing.T) {
	attr := ResourceComputedOptionalBoolWithPlanModifier("description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalBoolWithPlanModifier should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalBoolWithPlanModifier should return optional attribute")
	}
}

func TestResourceComputedOptionalBoolWithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceComputedOptionalBoolWithDefaultAndPlanModifier("description", true)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalBoolWithDefaultAndPlanModifier should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalBoolWithDefaultAndPlanModifier should return optional attribute")
	}
}

// Test missing data source bool functions
func TestDataSourceRequiredBool(t *testing.T) {
	attr := DataSourceRequiredBool("description")
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredBool should return required attribute")
	}
}

func TestDataSourceComputedOptionalBool(t *testing.T) {
	attr := DataSourceComputedOptionalBool("description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedOptionalBool should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("DataSourceComputedOptionalBool should return optional attribute")
	}
}

// Test collection builders for data sources
func TestNewDataSourceSetAttribute(t *testing.T) {
	attr := newDataSourceSetAttribute(collectionConfig{
		description: "description",
		optional:    true,
	})
	if !attr.IsOptional() {
		t.Fatal("newDataSourceSetAttribute should return optional attribute by default")
	}
}
