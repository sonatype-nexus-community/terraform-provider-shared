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

func TestStandardID(t *testing.T) {
	attr := StandardID()
	if attr.GetDescription() != "Internal ID of the resource" || !attr.IsComputed() {
		t.Fatal("StandardID should return a computed string attribute with correct description")
	}
}

func TestTimestamp(t *testing.T) {
	attr := Timestamp()
	if !attr.IsComputed() {
		t.Fatal("Timestamp should return a computed attribute")
	}
}

func TestRequiredString(t *testing.T) {
	attr := RequiredString("test description")
	if !attr.IsRequired() || attr.IsOptional() {
		t.Fatal("RequiredString should return a required string attribute")
	}
	if attr.GetDescription() != "test description" {
		t.Fatal("RequiredString should have the provided description")
	}
}

func TestOptionalString(t *testing.T) {
	attr := OptionalString("test description")
	if !attr.IsOptional() || attr.IsRequired() {
		t.Fatal("OptionalString should return an optional string attribute")
	}
}

func TestSensitiveString(t *testing.T) {
	attr := SensitiveString("password")
	if !attr.IsOptional() {
		t.Fatal("SensitiveString should return an optional attribute")
	}
	// Note: We can't easily check Sensitive property without deeper inspection
}

func TestStringEnum(t *testing.T) {
	attr := StringEnum("status", "active", "inactive", "pending")
	if !attr.IsOptional() {
		t.Fatal("StringEnum should return an optional attribute")
	}
}

func TestStringEnumWithDefault(t *testing.T) {
	attr := StringEnumWithDefault("status", "active", "active", "inactive")
	if !attr.IsOptional() || !attr.IsComputed() {
		t.Fatal("StringEnumWithDefault should return optional and computed attribute")
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
