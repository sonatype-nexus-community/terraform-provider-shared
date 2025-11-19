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

// Test attribute builder functions that create various combinations
func TestNewResourceStringAttribute(t *testing.T) {
	attr := newResourceStringAttribute(stringAttributeConfig{
		description: "test",
		required:    true,
	})
	if !attr.IsRequired() {
		t.Fatal("newResourceStringAttribute with required:true should return required attribute")
	}
}

func TestNewDataSourceStringAttribute(t *testing.T) {
	attr := newDataSourceStringAttribute(stringAttributeConfig{
		description: "test",
		computed:    true,
	})
	if !attr.IsComputed() {
		t.Fatal("newDataSourceStringAttribute with computed:true should return computed attribute")
	}
}

func TestNewResourceBoolAttribute(t *testing.T) {
	attr := newResourceBoolAttribute(boolAttributeConfig{
		description: "test",
		optional:    true,
	})
	if !attr.IsOptional() {
		t.Fatal("newResourceBoolAttribute with optional:true should return optional attribute")
	}
}

func TestNewDataSourceBoolAttribute(t *testing.T) {
	attr := newDataSourceBoolAttribute(boolAttributeConfig{
		description: "test",
		computed:    true,
	})
	if !attr.IsComputed() {
		t.Fatal("newDataSourceBoolAttribute with computed:true should return computed attribute")
	}
}

func TestNewResourceInt64Attribute(t *testing.T) {
	attr := newResourceInt64Attribute(int64AttributeConfig{
		description: "test",
		required:    true,
	})
	if !attr.IsRequired() {
		t.Fatal("newResourceInt64Attribute with required:true should return required attribute")
	}
}

func TestNewDataSourceInt64Attribute(t *testing.T) {
	attr := newDataSourceInt64Attribute(int64AttributeConfig{
		description: "test",
		optional:    true,
	})
	if !attr.IsOptional() {
		t.Fatal("newDataSourceInt64Attribute with optional:true should return optional attribute")
	}
}

func TestNewResourceFloat64Attribute(t *testing.T) {
	attr := newResourceFloat64Attribute(float64AttributeConfig{
		description: "test",
		computed:    true,
	})
	if !attr.IsComputed() {
		t.Fatal("newResourceFloat64Attribute with computed:true should return computed attribute")
	}
}

func TestNewDataSourceFloat64Attribute(t *testing.T) {
	attr := newDataSourceFloat64Attribute(float64AttributeConfig{
		description: "test",
		optional:    true,
	})
	if !attr.IsOptional() {
		t.Fatal("newDataSourceFloat64Attribute with optional:true should return optional attribute")
	}
}
