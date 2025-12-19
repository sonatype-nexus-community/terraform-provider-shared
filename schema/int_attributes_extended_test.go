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

// Test int64 attribute functions
func TestResourceRequiredInt64(t *testing.T) {
	attr := ResourceRequiredInt64("test description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt64 should return required attribute")
	}
}

func TestResourceOptionalInt64(t *testing.T) {
	attr := ResourceOptionalInt64("test description")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt64 should return optional attribute")
	}
}

func TestResourceComputedInt64(t *testing.T) {
	attr := ResourceComputedInt64("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt64 should return computed attribute")
	}
}

func TestResourceComputedInt64WithDefault(t *testing.T) {
	attr := ResourceComputedInt64WithDefault("test description", 42)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt64WithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedInt64WithDefault should have default value")
	}
}

func TestResourceOptionalInt64WithDefault(t *testing.T) {
	attr := ResourceOptionalInt64WithDefault("test description", 42)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt64WithDefault should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceOptionalInt64WithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceOptionalInt64WithDefault should have default value")
	}
}

func TestResourceComputedOptionalInt64(t *testing.T) {
	attr := ResourceComputedOptionalInt64("test description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt64 should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt64 should return optional attribute")
	}
}

func TestResourceRequiredInt64WithDefault(t *testing.T) {
	attr := ResourceRequiredInt64WithDefault("test description", 42)
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredInt64WithDefault should return required attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceRequiredInt64WithDefault should have default value")
	}
}

func TestResourceComputedOptionalInt64WithDefault(t *testing.T) {
	attr := ResourceComputedOptionalInt64WithDefault("test description", 42)
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt64WithDefault should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt64WithDefault should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedOptionalInt64WithDefault should have default value")
	}
}

func TestResourceOptionalInt64WithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceOptionalInt64WithDefaultAndPlanModifier("test description", 42)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalInt64WithDefaultAndPlanModifier should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceOptionalInt64WithDefaultAndPlanModifier should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceOptionalInt64WithDefaultAndPlanModifier should have default value")
	}
}

func TestResourceComputedInt64WithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceComputedInt64WithDefaultAndPlanModifier("test description", 42)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedInt64WithDefaultAndPlanModifier should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedInt64WithDefaultAndPlanModifier should have default value")
	}
}

func TestResourceComputedOptionalInt64WithDefaultAndPlanModifier(t *testing.T) {
	attr := ResourceComputedOptionalInt64WithDefaultAndPlanModifier("test description", 42)
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalInt64WithDefaultAndPlanModifier should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalInt64WithDefaultAndPlanModifier should return computed attribute")
	}
	if attr.Default == nil {
		t.Fatal("ResourceComputedOptionalInt64WithDefaultAndPlanModifier should have default value")
	}
}

func TestResourceOptionalPort(t *testing.T) {
	attr := ResourceOptionalPort("test port")
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalPort should return optional attribute")
	}
}

func TestDataSourceOptionalInt64(t *testing.T) {
	attr := DataSourceOptionalInt64("test description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalInt64 should return optional attribute")
	}
}

func TestDataSourceComputedInt64(t *testing.T) {
	attr := DataSourceComputedInt64("test description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedInt64 should return computed attribute")
	}
}

// Test missing int64 attribute functions
func TestResourceRequiredPort(t *testing.T) {
	attr := ResourceRequiredPort("port number")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredPort should return required attribute")
	}
}

func TestResourcePortWithDefault(t *testing.T) {
	attr := ResourcePortWithDefault("port description", 8080)
	if !attr.IsOptional() {
		t.Fatal("ResourcePortWithDefault should return optional attribute")
	}
	if !attr.IsComputed() {
		t.Fatal("ResourcePortWithDefault should return computed attribute")
	}
}

func TestResourcePercentageInt(t *testing.T) {
	attr := ResourcePercentageInt("percentage description")
	if !attr.IsOptional() {
		t.Fatal("ResourcePercentageInt should return optional attribute")
	}
}

func TestResourceRequiredPercentageInt(t *testing.T) {
	attr := ResourceRequiredPercentageInt("required percentage description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredPercentageInt should return required attribute")
	}
}

func TestResourceDurationInt(t *testing.T) {
	attr := ResourceDurationInt("duration description")
	if !attr.IsOptional() {
		t.Fatal("ResourceDurationInt should return optional attribute")
	}
}

func TestResourceRequiredDurationInt(t *testing.T) {
	attr := ResourceRequiredDurationInt("required duration description")
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredDurationInt should return required attribute")
	}
}

func TestResourceTimeoutInt(t *testing.T) {
	attr := ResourceTimeoutInt("timeout description", 300)
	if !attr.IsOptional() {
		t.Fatal("ResourceTimeoutInt should return optional attribute")
	}
}

func TestResourceCountInt(t *testing.T) {
	attr := ResourceCountInt("count description")
	if !attr.IsOptional() {
		t.Fatal("ResourceCountInt should return optional attribute")
	}
}
