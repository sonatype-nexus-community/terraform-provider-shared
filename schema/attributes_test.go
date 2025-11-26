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

func TestResourceStandardID(t *testing.T) {
	attr := ResourceStandardID()
	if attr.GetMarkdownDescription() != "Internal ID of the resource" || !attr.IsComputed() {
		t.Fatal("ResourceStandardID should return a computed string attribute with correct description")
	}
}

func TestResourceLastUpdated(t *testing.T) {
	attr := ResourceLastUpdated()
	if !attr.IsComputed() {
		t.Fatal("ResourceLastUpdated should return a computed attribute")
	}
}

func TestResourceRequiredString(t *testing.T) {
	attr := ResourceRequiredString("test description")
	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredString should return a required string attribute")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("ResourceRequiredString should have the provided description")
	}
}

func TestResourceOptionalString(t *testing.T) {
	attr := ResourceOptionalString("test description")
	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalString should return an optional string attribute")
	}
}

func TestResourceOptionalStringWithDefault(t *testing.T) {
	attr := ResourceOptionalStringWithDefault("test description", "default value")
	if !attr.IsOptional() || !attr.IsComputed() {
		t.Fatal("ResourceOptionalStringWithDefault should return an optional and computed string attribute")
	}
	if attr.GetMarkdownDescription() != "test description" {
		t.Fatal("ResourceOptionalStringWithDefault should have the provided description")
	}
}

func TestResourceSensitiveString(t *testing.T) {
	attr := ResourceSensitiveString("password")
	if !attr.IsOptional() {
		t.Fatal("ResourceSensitiveString should return an optional attribute")
	}
	// Note: We can't easily check Sensitive property without deeper inspection
}

func TestResourceStringEnum(t *testing.T) {
	attr := ResourceStringEnum("status", "active", "inactive", "pending")
	if !attr.IsOptional() {
		t.Fatal("ResourceStringEnum should return an optional attribute")
	}
}

func TestResourceOptionalStringEnum(t *testing.T) {
	attr := ResourceOptionalStringEnum("status", "active", "inactive", "pending")
	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalStringEnum should return an optional attribute")
	}
	if attr.GetMarkdownDescription() != "status" {
		t.Fatal("ResourceOptionalStringEnum should have the provided description")
	}
}

func TestResourceRequiredStringEnum(t *testing.T) {
	attr := ResourceRequiredStringEnum("status", "active", "inactive", "pending")
	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredStringEnum should return a required attribute")
	}
	if attr.GetMarkdownDescription() != "status" {
		t.Fatal("ResourceRequiredStringEnum should have the provided description")
	}
}

func TestResourceStringEnumWithDefault(t *testing.T) {
	attr := ResourceStringEnumWithDefault("status", "active", "active", "inactive")
	if !attr.IsOptional() || !attr.IsComputed() {
		t.Fatal("ResourceStringEnumWithDefault should return optional and computed attribute")
	}
}

func TestResourceIDAttribute(t *testing.T) {
	attr := ResourceIDAttribute("organization_id")
	if !attr.IsRequired() {
		t.Fatal("ResourceIDAttribute should return a required attribute")
	}
}

func TestStandardResourceAttributes(t *testing.T) {
	attrs := StandardResourceAttributes()
	if len(attrs) != 2 {
		t.Fatalf("StandardResourceAttributes should return 2 attributes, got %d", len(attrs))
	}
	if _, exists := attrs["id"]; !exists {
		t.Fatal("StandardResourceAttributes should include 'id'")
	}
	if _, exists := attrs["last_updated"]; !exists {
		t.Fatal("StandardResourceAttributes should include 'last_updated'")
	}
}

func TestResourceComputedStringFunc(t *testing.T) {
	attr := ResourceComputedString("computed description")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedString should return a computed attribute")
	}
}

func TestResourceComputedStringWithDefaultFunc(t *testing.T) {
	attr := ResourceComputedStringWithDefault("description", "default")
	if !attr.IsComputed() {
		t.Fatal("ResourceComputedStringWithDefault should return a computed attribute")
	}
}

func TestDataSourceComputedStringFunc(t *testing.T) {
	attr := DataSourceComputedString("description")
	if !attr.IsComputed() {
		t.Fatal("DataSourceComputedString should return a computed attribute")
	}
}

func TestDataSourceOptionalStringFunc(t *testing.T) {
	attr := DataSourceOptionalString("optional description")
	if !attr.IsOptional() {
		t.Fatal("DataSourceOptionalString should return an optional attribute")
	}
}

func TestDataSourceRequiredStringFunc(t *testing.T) {
	attr := DataSourceRequiredString("required description")
	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("DataSourceRequiredString should return a required attribute")
	}
	if attr.GetMarkdownDescription() != "required description" {
		t.Fatal("DataSourceRequiredString should have the provided description")
	}
}

func TestNamedResourceAttributesFunc(t *testing.T) {
	attrs := NamedResourceAttributes()
	if len(attrs) == 0 {
		t.Fatal("NamedResourceAttributes should return at least one attribute")
	}
}

func TestIdentityResourceAttributesFunc(t *testing.T) {
	attrs := IdentityResourceAttributes()
	if len(attrs) == 0 {
		t.Fatal("IdentityResourceAttributes should return at least one attribute")
	}
}

func TestOwnershipResourceAttributesFunc(t *testing.T) {
	attrs := OwnershipResourceAttributes()
	if len(attrs) == 0 {
		t.Fatal("OwnershipResourceAttributes should return at least one attribute")
	}
}

func TestAuditableResourceAttributesFunc(t *testing.T) {
	attrs := AuditableResourceAttributes()
	if len(attrs) == 0 {
		t.Fatal("AuditableResourceAttributes should return at least one attribute")
	}
}

func TestDataSourceIDAttributeFunc(t *testing.T) {
	attr := DataSourceIDAttribute()
	// DataSourceIDAttribute returns a single StringAttribute, not a map
	if attr.GetMarkdownDescription() == "" && !attr.IsComputed() {
		t.Fatal("DataSourceIDAttribute should return a valid attribute")
	}
}

func TestSourceAndReadOnlyAttributesFunc(t *testing.T) {
	attrs := SourceAndReadOnlyAttributes()
	if len(attrs) == 0 {
		t.Fatal("SourceAndReadOnlyAttributes should return at least one attribute")
	}
}
