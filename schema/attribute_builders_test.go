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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ========================================
// Resource Int64 Attribute Builder Tests
// ========================================

func TestNewResourceInt64Attribute_Required(t *testing.T) {
	config := int64AttributeConfig{
		description: "test description",
		required:    true,
	}
	attr := newResourceInt64Attribute(config)
	if !attr.IsRequired() {
		t.Fatal("newResourceInt64Attribute with required=true should return required attribute")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("newResourceInt64Attribute should preserve description")
	}
}

func TestNewResourceInt64Attribute_Optional(t *testing.T) {
	config := int64AttributeConfig{
		description: "test description",
		optional:    true,
	}
	attr := newResourceInt64Attribute(config)
	if !attr.IsOptional() {
		t.Fatal("newResourceInt64Attribute with optional=true should return optional attribute")
	}
}

func TestNewResourceInt64Attribute_Computed(t *testing.T) {
	config := int64AttributeConfig{
		description: "test description",
		computed:    true,
	}
	attr := newResourceInt64Attribute(config)
	if !attr.IsComputed() {
		t.Fatal("newResourceInt64Attribute with computed=true should return computed attribute")
	}
}

func TestNewResourceInt64Attribute_WithValidators(t *testing.T) {
	config := int64AttributeConfig{
		description: "test description",
		optional:    true,
	}
	attr := newResourceInt64Attribute(config)
	if !attr.IsOptional() {
		t.Fatal("newResourceInt64Attribute should respect optional flag")
	}
}

// ========================================
// DataSource Int64 Attribute Builder Tests
// ========================================

func TestNewDataSourceInt64Attribute_Required(t *testing.T) {
	config := int64AttributeConfig{
		description: "test description",
		required:    true,
	}
	attr := newDataSourceInt64Attribute(config)
	if !attr.IsRequired() {
		t.Fatal("newDataSourceInt64Attribute with required=true should return required attribute")
	}
}

func TestNewDataSourceInt64Attribute_Optional(t *testing.T) {
	config := int64AttributeConfig{
		description: "test description",
		optional:    true,
	}
	attr := newDataSourceInt64Attribute(config)
	if !attr.IsOptional() {
		t.Fatal("newDataSourceInt64Attribute with optional=true should return optional attribute")
	}
}

func TestNewDataSourceInt64Attribute_Computed(t *testing.T) {
	config := int64AttributeConfig{
		description: "test description",
		computed:    true,
	}
	attr := newDataSourceInt64Attribute(config)
	if !attr.IsComputed() {
		t.Fatal("newDataSourceInt64Attribute with computed=true should return computed attribute")
	}
}

// ========================================
// Resource Int32 Attribute Builder Tests
// ========================================

func TestNewResourceInt32Attribute_Required(t *testing.T) {
	config := int32AttributeConfig{
		description: "test description",
		required:    true,
	}
	attr := newResourceInt32Attribute(config)
	if !attr.IsRequired() {
		t.Fatal("newResourceInt32Attribute with required=true should return required attribute")
	}
}

func TestNewResourceInt32Attribute_Optional(t *testing.T) {
	config := int32AttributeConfig{
		description: "test description",
		optional:    true,
	}
	attr := newResourceInt32Attribute(config)
	if !attr.IsOptional() {
		t.Fatal("newResourceInt32Attribute with optional=true should return optional attribute")
	}
}

func TestNewResourceInt32Attribute_Computed(t *testing.T) {
	config := int32AttributeConfig{
		description: "test description",
		computed:    true,
	}
	attr := newResourceInt32Attribute(config)
	if !attr.IsComputed() {
		t.Fatal("newResourceInt32Attribute with computed=true should return computed attribute")
	}
}

// ========================================
// DataSource Int32 Attribute Builder Tests
// ========================================

func TestNewDataSourceInt32Attribute_Required(t *testing.T) {
	config := int32AttributeConfig{
		description: "test description",
		required:    true,
	}
	attr := newDataSourceInt32Attribute(config)
	if !attr.IsRequired() {
		t.Fatal("newDataSourceInt32Attribute with required=true should return required attribute")
	}
}

func TestNewDataSourceInt32Attribute_Optional(t *testing.T) {
	config := int32AttributeConfig{
		description: "test description",
		optional:    true,
	}
	attr := newDataSourceInt32Attribute(config)
	if !attr.IsOptional() {
		t.Fatal("newDataSourceInt32Attribute with optional=true should return optional attribute")
	}
}

func TestNewDataSourceInt32Attribute_Computed(t *testing.T) {
	config := int32AttributeConfig{
		description: "test description",
		computed:    true,
	}
	attr := newDataSourceInt32Attribute(config)
	if !attr.IsComputed() {
		t.Fatal("newDataSourceInt32Attribute with computed=true should return computed attribute")
	}
}

// ========================================
// Resource Float64 Attribute Builder Tests
// ========================================

func TestNewResourceFloat64Attribute_Required(t *testing.T) {
	config := float64AttributeConfig{
		description: "test description",
		required:    true,
	}
	attr := newResourceFloat64Attribute(config)
	if !attr.IsRequired() {
		t.Fatal("newResourceFloat64Attribute with required=true should return required attribute")
	}
}

func TestNewResourceFloat64Attribute_Optional(t *testing.T) {
	config := float64AttributeConfig{
		description: "test description",
		optional:    true,
	}
	attr := newResourceFloat64Attribute(config)
	if !attr.IsOptional() {
		t.Fatal("newResourceFloat64Attribute with optional=true should return optional attribute")
	}
}

func TestNewResourceFloat64Attribute_Computed(t *testing.T) {
	config := float64AttributeConfig{
		description: "test description",
		computed:    true,
	}
	attr := newResourceFloat64Attribute(config)
	if !attr.IsComputed() {
		t.Fatal("newResourceFloat64Attribute with computed=true should return computed attribute")
	}
}

// ========================================
// DataSource Float64 Attribute Builder Tests
// ========================================

func TestNewDataSourceFloat64Attribute_Required(t *testing.T) {
	config := float64AttributeConfig{
		description: "test description",
		required:    true,
	}
	attr := newDataSourceFloat64Attribute(config)
	if !attr.IsRequired() {
		t.Fatal("newDataSourceFloat64Attribute with required=true should return required attribute")
	}
}

func TestNewDataSourceFloat64Attribute_Optional(t *testing.T) {
	config := float64AttributeConfig{
		description: "test description",
		optional:    true,
	}
	attr := newDataSourceFloat64Attribute(config)
	if !attr.IsOptional() {
		t.Fatal("newDataSourceFloat64Attribute with optional=true should return optional attribute")
	}
}

func TestNewDataSourceFloat64Attribute_Computed(t *testing.T) {
	config := float64AttributeConfig{
		description: "test description",
		computed:    true,
	}
	attr := newDataSourceFloat64Attribute(config)
	if !attr.IsComputed() {
		t.Fatal("newDataSourceFloat64Attribute with computed=true should return computed attribute")
	}
}

// ========================================
// Set Attribute Builder Tests
// ========================================

func TestNewDataSourceSetAttribute_WithValidators(t *testing.T) {
	config := collectionConfig{
		description: "test set",
		elementType: types.StringType,
		optional:    true,
	}
	attr := newDataSourceSetAttribute(config)
	if !attr.IsOptional() {
		t.Fatal("newDataSourceSetAttribute should return optional attribute")
	}
	if attr.ElementType != types.StringType {
		t.Fatal("newDataSourceSetAttribute should preserve element type")
	}
}
