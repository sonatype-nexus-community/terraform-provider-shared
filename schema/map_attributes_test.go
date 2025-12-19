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

// Test DataSourceRequiredStringMap
func TestDataSourceRequiredStringMap(t *testing.T) {
	attr := DataSourceRequiredStringMap("required string map description")
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredStringMap should return required attribute")
	}
	if attr.GetMarkdownDescription() != "required string map description" {
		t.Fatal("DataSourceRequiredStringMap should have the provided description")
	}
}

// Test DataSourceRequiredInt64Map
func TestDataSourceRequiredInt64Map(t *testing.T) {
	attr := DataSourceRequiredInt64Map("required int64 map description")
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredInt64Map should return required attribute")
	}
	if attr.GetMarkdownDescription() != "required int64 map description" {
		t.Fatal("DataSourceRequiredInt64Map should have the provided description")
	}
}

// Test DataSourceRequiredBoolMap
func TestDataSourceRequiredBoolMap(t *testing.T) {
	attr := DataSourceRequiredBoolMap("required bool map description")
	if !attr.IsRequired() {
		t.Fatal("DataSourceRequiredBoolMap should return required attribute")
	}
	if attr.GetMarkdownDescription() != "required bool map description" {
		t.Fatal("DataSourceRequiredBoolMap should have the provided description")
	}
}
