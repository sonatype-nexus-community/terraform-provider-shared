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

	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Test nested attribute functions
func TestResourceRequiredSingleNestedAttribute(t *testing.T) {
	attrs := map[string]resourceschema.Attribute{
		"test": ResourceRequiredString("test attr"),
	}
	attr := ResourceRequiredSingleNestedAttribute("test description", attrs)
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredSingleNestedAttribute should return required attribute")
	}
}

func TestResourceOptionalSingleNestedAttribute(t *testing.T) {
	attrs := map[string]resourceschema.Attribute{
		"test": ResourceOptionalString("test attr"),
	}
	attr := ResourceOptionalSingleNestedAttribute("test description", attrs)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalSingleNestedAttribute should return optional attribute")
	}
}

func TestResourceComputedSingleNestedAttribute(t *testing.T) {
	attrs := map[string]resourceschema.Attribute{
		"test": ResourceComputedString("test attr"),
	}
	attr := ResourceComputedSingleNestedAttribute("test description", attrs)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedSingleNestedAttribute should return computed attribute")
	}
}

func TestResourceComputedOptionalSingleNestedAttribute(t *testing.T) {
	attrs := map[string]resourceschema.Attribute{
		"test": ResourceComputedString("test attr"),
	}
	attr := ResourceComputedOptionalSingleNestedAttribute("test description", attrs)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalSingleNestedAttribute should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalSingleNestedAttribute should return optional attribute")
	}
}

func TestResourceRequiredListNestedAttribute(t *testing.T) {
	nestedObject := resourceschema.NestedAttributeObject{
		Attributes: map[string]resourceschema.Attribute{
			"test": ResourceRequiredString("test attr"),
		},
	}
	attr := ResourceRequiredListNestedAttribute("test description", nestedObject)
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredListNestedAttribute should return required attribute")
	}
}

func TestResourceOptionalListNestedAttribute(t *testing.T) {
	nestedObject := resourceschema.NestedAttributeObject{
		Attributes: map[string]resourceschema.Attribute{
			"test": ResourceOptionalString("test attr"),
		},
	}
	attr := ResourceOptionalListNestedAttribute("test description", nestedObject)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalListNestedAttribute should return optional attribute")
	}
}

func TestResourceComputedListNestedAttribute(t *testing.T) {
	nestedObject := resourceschema.NestedAttributeObject{
		Attributes: map[string]resourceschema.Attribute{
			"test": ResourceComputedString("test attr"),
		},
	}
	attr := ResourceComputedListNestedAttribute("test description", nestedObject)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedListNestedAttribute should return computed attribute")
	}
}

func TestResourceComputedOptionalListNestedAttribute(t *testing.T) {
	nestedObject := resourceschema.NestedAttributeObject{
		Attributes: map[string]resourceschema.Attribute{
			"test": ResourceComputedString("test attr"),
		},
	}
	attr := ResourceComputedOptionalListNestedAttribute("test description", nestedObject)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalListNestedAttribute should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalListNestedAttribute should return optional attribute")
	}
}

func TestResourceRequiredSetNestedAttribute(t *testing.T) {
	nestedObject := resourceschema.NestedAttributeObject{
		Attributes: map[string]resourceschema.Attribute{
			"test": ResourceRequiredString("test attr"),
		},
	}
	attr := ResourceRequiredSetNestedAttribute("test description", nestedObject)
	if !attr.IsRequired() {
		t.Fatal("ResourceRequiredSetNestedAttribute should return required attribute")
	}
}

func TestResourceOptionalSetNestedAttribute(t *testing.T) {
	nestedObject := resourceschema.NestedAttributeObject{
		Attributes: map[string]resourceschema.Attribute{
			"test": ResourceOptionalString("test attr"),
		},
	}
	attr := ResourceOptionalSetNestedAttribute("test description", nestedObject)
	if !attr.IsOptional() {
		t.Fatal("ResourceOptionalSetNestedAttribute should return optional attribute")
	}
}

func TestResourceComputedSetNestedAttribute(t *testing.T) {
	nestedObject := resourceschema.NestedAttributeObject{
		Attributes: map[string]resourceschema.Attribute{
			"test": ResourceComputedString("test attr"),
		},
	}
	attr := ResourceComputedSetNestedAttribute("test description", nestedObject)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedSetNestedAttribute should return computed attribute")
	}
}

func TestResourceComputedOptionalSetNestedAttribute(t *testing.T) {
	nestedObject := resourceschema.NestedAttributeObject{
		Attributes: map[string]resourceschema.Attribute{
			"test": ResourceComputedString("test attr"),
		},
	}
	attr := ResourceComputedOptionalSetNestedAttribute("test description", nestedObject)
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedOptionalSetNestedAttribute should return computed attribute")
	}
	if !attr.IsOptional() {
		t.Fatal("ResourceComputedOptionalSetNestedAttribute should return optional attribute")
	}
}
