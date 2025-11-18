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
	if attr.GetDescription() != "Internal ID of the resource" || !attr.IsComputed() {
		t.Fatal("ResourceStandardID should return a computed string attribute with correct description")
	}
}

func TestResourceTimestamp(t *testing.T) {
	attr := ResourceTimestamp()
	if !attr.IsComputed() {
		t.Fatal("ResourceTimestamp should return a computed attribute")
	}
}

func TestResourceRequiredString(t *testing.T) {
	attr := ResourceRequiredString("test description")
	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("ResourceRequiredString should return a required string attribute")
	}
	if attr.GetDescription() != "test description" {
		t.Fatal("ResourceRequiredString should have the provided description")
	}
}

func TestResourceOptionalString(t *testing.T) {
	attr := ResourceOptionalString("test description")
	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("ResourceOptionalString should return an optional string attribute")
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
