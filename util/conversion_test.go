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

package util

import (
	"testing"
)

func TestStringToPtr(t *testing.T) {
	s := "test"
	ptr := StringToPtr(s)
	if ptr == nil || *ptr != s {
		t.Fatal("StringToPtr failed")
	}
}

func TestStringPtrToValue(t *testing.T) {
	s := "test"
	val := StringPtrToValue(&s)
	if val.IsNull() || val.ValueString() != s {
		t.Fatal("StringPtrToValue failed")
	}

	val = StringPtrToValue(nil)
	if !val.IsNull() {
		t.Fatal("StringPtrToValue(nil) should return null value")
	}
}

func TestBoolToPtr(t *testing.T) {
	b := true
	ptr := BoolToPtr(b)
	if ptr == nil || *ptr != b {
		t.Fatal("BoolToPtr failed")
	}
}

func TestBoolPtrToValue(t *testing.T) {
	b := true
	val := BoolPtrToValue(&b)
	if val.IsNull() || val.ValueBool() != b {
		t.Fatal("BoolPtrToValue failed")
	}

	val = BoolPtrToValue(nil)
	if !val.IsNull() {
		t.Fatal("BoolPtrToValue(nil) should return null value")
	}
}

func TestInt64ToPtr(t *testing.T) {
	i := int64(42)
	ptr := Int64ToPtr(i)
	if ptr == nil || *ptr != i {
		t.Fatal("Int64ToPtr failed")
	}
}

func TestInt64PtrToValue(t *testing.T) {
	i := int64(42)
	val := Int64PtrToValue(&i)
	if val.IsNull() || val.ValueInt64() != i {
		t.Fatal("Int64PtrToValue failed")
	}

	val = Int64PtrToValue(nil)
	if !val.IsNull() {
		t.Fatal("Int64PtrToValue(nil) should return null value")
	}
}

func TestStringToInt64(t *testing.T) {
	result, err := StringToInt64("42")
	if err != nil || result != 42 {
		t.Fatal("StringToInt64 failed")
	}

	_, err = StringToInt64("invalid")
	if err == nil {
		t.Fatal("StringToInt64 should error on invalid input")
	}
}

func TestInt64ToString(t *testing.T) {
	result := Int64ToString(42)
	if result != "42" {
		t.Fatalf("Int64ToString failed: got %s, expected 42", result)
	}
}

func TestStringToBool(t *testing.T) {
	result, err := StringToBool("true")
	if err != nil || !result {
		t.Fatal("StringToBool('true') failed")
	}

	result, err = StringToBool("false")
	if err != nil || result {
		t.Fatal("StringToBool('false') failed")
	}
}

func TestBoolToString(t *testing.T) {
	if BoolToString(true) != "true" {
		t.Fatal("BoolToString(true) failed")
	}
	if BoolToString(false) != "false" {
		t.Fatal("BoolToString(false) failed")
	}
}

func TestSafeString(t *testing.T) {
	s := "test"
	if SafeString(&s) != s {
		t.Fatal("SafeString failed")
	}

	if SafeString(nil) != "" {
		t.Fatal("SafeString(nil) should return empty string")
	}
}

func TestSafeInt64(t *testing.T) {
	i := int64(42)
	if SafeInt64(&i) != i {
		t.Fatal("SafeInt64 failed")
	}

	if SafeInt64(nil) != 0 {
		t.Fatal("SafeInt64(nil) should return 0")
	}
}

func TestSafeBool(t *testing.T) {
	b := true
	if !SafeBool(&b) {
		t.Fatal("SafeBool(&true) failed")
	}

	if SafeBool(nil) {
		t.Fatal("SafeBool(nil) should return false")
	}
}

func TestNilIfEmpty(t *testing.T) {
	result := NilIfEmpty("test")
	if result == nil || *result != "test" {
		t.Fatal("NilIfEmpty('test') failed")
	}

	result = NilIfEmpty("")
	if result != nil {
		t.Fatal("NilIfEmpty('') should return nil")
	}
}

func TestEmptyIfNil(t *testing.T) {
	s := "test"
	if EmptyIfNil(&s) != s {
		t.Fatal("EmptyIfNil failed")
	}

	if EmptyIfNil(nil) != "" {
		t.Fatal("EmptyIfNil(nil) should return empty string")
	}
}
