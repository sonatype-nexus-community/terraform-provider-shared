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

	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Test data source nested attribute functions
func TestDataSourceOptionalSingleNestedAttribute(t *testing.T) {
	attr := DataSourceOptionalSingleNestedAttribute(
		"description",
		map[string]datasourceschema.Attribute{},
	)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalSingleNestedAttribute should return optional attribute")
	}
}

func TestDataSourceComputedSingleNestedAttribute(t *testing.T) {
	attr := DataSourceComputedSingleNestedAttribute(
		"description",
		map[string]datasourceschema.Attribute{},
	)
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedSingleNestedAttribute should return computed attribute")
	}
}

func TestDataSourceComputedOptionalSingleNestedAttribute(t *testing.T) {
	attr := DataSourceComputedOptionalSingleNestedAttribute(
		"description",
		map[string]datasourceschema.Attribute{},
	)
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedOptionalSingleNestedAttribute should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("DataSourceComputedOptionalSingleNestedAttribute should return optional attribute")
	}
}

func TestDataSourceOptionalListNestedAttribute(t *testing.T) {
	// Create a minimal nested object for testing
	nestedObject := datasourceschema.NestedAttributeObject{
		Attributes: map[string]datasourceschema.Attribute{},
	}
	attr := DataSourceOptionalListNestedAttribute("description", nestedObject)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalListNestedAttribute should return optional attribute")
	}
}

func TestDataSourceComputedListNestedAttribute(t *testing.T) {
	nestedObject := datasourceschema.NestedAttributeObject{
		Attributes: map[string]datasourceschema.Attribute{},
	}
	attr := DataSourceComputedListNestedAttribute("description", nestedObject)
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedListNestedAttribute should return computed attribute")
	}
}

func TestDataSourceOptionalSetNestedAttribute(t *testing.T) {
	nestedObject := datasourceschema.NestedAttributeObject{
		Attributes: map[string]datasourceschema.Attribute{},
	}
	attr := DataSourceOptionalSetNestedAttribute("description", nestedObject)
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalSetNestedAttribute should return optional attribute")
	}
}

func TestDataSourceComputedSetNestedAttribute(t *testing.T) {
	nestedObject := datasourceschema.NestedAttributeObject{
		Attributes: map[string]datasourceschema.Attribute{},
	}
	attr := DataSourceComputedSetNestedAttribute("description", nestedObject)
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedSetNestedAttribute should return computed attribute")
	}
}
