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
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// StringToPtr converts a string to a pointer to string
func StringToPtr(s string) *string {
	return &s
}

// StringPtrToValue converts a pointer to string to a types.String value
func StringPtrToValue(s *string) types.String {
	if s == nil {
		return types.StringNull()
	}
	return types.StringValue(*s)
}

// BoolToPtr converts a bool to a pointer to bool
func BoolToPtr(b bool) *bool {
	return &b
}

// BoolPtrToValue converts a pointer to bool to a types.Bool value
func BoolPtrToValue(b *bool) types.Bool {
	if b == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*b)
}

// Int64ToPtr converts an int64 to a pointer to int64
func Int64ToPtr(i int64) *int64 {
	return &i
}

// Int64PtrToValue converts a pointer to int64 to a types.Int64 value
func Int64PtrToValue(i *int64) types.Int64 {
	if i == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*i)
}

// Int32ToPtr converts an int32 to a pointer to int32
func Int32ToPtr(i int32) *int32 {
	return &i
}

// Int32PtrToValue converts a pointer to int32 to a types.Int64 value (int32 stored as int64)
func Int32PtrToValue(i *int32) types.Int64 {
	if i == nil {
		return types.Int64Null()
	}
	return types.Int64Value(int64(*i))
}

// StringToInt64 converts a string to int64
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// StringToInt32 converts a string to int32
func StringToInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

// Int64ToString converts an int64 to a string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Int32ToString converts an int32 to a string
func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// StringToFloat converts a string to float64
func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// FloatToString converts a float64 to a string with precision
func FloatToString(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}

// StringToBool converts a string to bool (case-insensitive)
func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// BoolToString converts a bool to "true" or "false"
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// StringSliceToValue converts a slice of strings to types.Set
func StringSliceToValue(ss []string) types.Set {
	val, _ := types.SetValueFrom(context.Background(), types.StringType, ss)
	return val
}

// StringPtrSliceToValue converts a slice of string pointers to types.Set
func StringPtrSliceToValue(ss []*string) types.Set {
	var values []string
	for _, s := range ss {
		if s != nil {
			values = append(values, *s)
		}
	}
	return StringSliceToValue(values)
}

// Int64SliceToValue converts a slice of int64s to types.Set
func Int64SliceToValue(ii []int64) types.Set {
	val, _ := types.SetValueFrom(context.Background(), types.Int64Type, ii)
	return val
}

// SafeString ensures we have a valid string, returning empty string if nil
func SafeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// SafeInt64 ensures we have a valid int64, returning 0 if nil
func SafeInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// SafeInt32 ensures we have a valid int32, returning 0 if nil
func SafeInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// SafeBool ensures we have a valid bool, returning false if nil
func SafeBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// NilIfEmpty returns nil if the string is empty, otherwise returns a pointer to the string
func NilIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// EmptyIfNil returns an empty string if the pointer is nil, otherwise returns the dereferenced value
func EmptyIfNil(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// FormatValue formats any value to a string representation
func FormatValue(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
