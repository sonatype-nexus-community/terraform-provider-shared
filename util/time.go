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
	"time"
)

// TimestampFormat is the standard format for Terraform timestamps
const TimestampFormat = time.RFC850

// CurrentTimestamp returns the current time formatted as a standard timestamp
func CurrentTimestamp() string {
	return time.Now().Format(TimestampFormat)
}

// ParseTimestamp parses a timestamp string using the standard format
func ParseTimestamp(ts string) (time.Time, error) {
	return time.Parse(TimestampFormat, ts)
}

// IsZeroTimestamp checks if a timestamp string represents zero/unset time
func IsZeroTimestamp(ts string) bool {
	return ts == "" || ts == "0"
}

// FormatTimestamp formats a time.Time using the standard format
func FormatTimestamp(t time.Time) string {
	return t.Format(TimestampFormat)
}

// TimestampNowOrEmpty returns current timestamp or empty string if provided timestamp is empty
func TimestampNowOrEmpty(currentTs string) string {
	if IsZeroTimestamp(currentTs) {
		return CurrentTimestamp()
	}
	return currentTs
}

// UnixTimestamp returns the current Unix timestamp
func UnixTimestamp() int64 {
	return time.Now().Unix()
}

// UnixTimestampToString converts Unix timestamp to RFC850 format
func UnixTimestampToString(unix int64) string {
	return time.Unix(unix, 0).Format(TimestampFormat)
}

// StringToUnixTimestamp converts RFC850 formatted string to Unix timestamp
func StringToUnixTimestamp(ts string) (int64, error) {
	t, err := ParseTimestamp(ts)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
