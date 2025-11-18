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
	"time"
)

func TestCurrentTimestamp(t *testing.T) {
	ts := CurrentTimestamp()
	if ts == "" {
		t.Fatal("CurrentTimestamp should not return empty string")
	}

	// Verify it's parseable
	_, err := ParseTimestamp(ts)
	if err != nil {
		t.Fatalf("CurrentTimestamp should return valid format: %v", err)
	}
}

func TestParseTimestamp(t *testing.T) {
	now := time.Now()
	ts := now.Format(TimestampFormat)

	parsed, err := ParseTimestamp(ts)
	if err != nil {
		t.Fatalf("ParseTimestamp failed: %v", err)
	}

	// Check that parsed time is within a second of original
	diff := parsed.Sub(now)
	if diff < 0 {
		diff = -diff
	}
	if diff > time.Second {
		t.Fatalf("Parsed time differs from original by %v", diff)
	}
}

func TestIsZeroTimestamp(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"0", true},
		{CurrentTimestamp(), false},
	}

	for _, test := range tests {
		result := IsZeroTimestamp(test.input)
		if result != test.expected {
			t.Fatalf("IsZeroTimestamp(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestTimestampNowOrEmpty(t *testing.T) {
	emptyResult := TimestampNowOrEmpty("")
	if IsZeroTimestamp(emptyResult) {
		t.Fatal("TimestampNowOrEmpty should return current timestamp for empty input")
	}

	existingTs := "Mon, 02 Jan 2006 15:04:05 MST"
	result := TimestampNowOrEmpty(existingTs)
	if result != existingTs {
		t.Fatalf("TimestampNowOrEmpty should return existing timestamp, got %s", result)
	}
}

func TestUnixTimestamp(t *testing.T) {
	unix := UnixTimestamp()
	now := time.Now().Unix()

	// Check that unix timestamp is within a second of current time
	if unix < now-1 || unix > now+1 {
		t.Fatalf("UnixTimestamp = %d, expected around %d", unix, now)
	}
}

func TestUnixTimestampConversions(t *testing.T) {
	now := time.Now()
	unix := now.Unix()

	// Convert to string
	ts := UnixTimestampToString(unix)
	if ts == "" {
		t.Fatal("UnixTimestampToString should not return empty string")
	}

	// Convert back to unix
	unix2, err := StringToUnixTimestamp(ts)
	if err != nil {
		t.Fatalf("StringToUnixTimestamp failed: %v", err)
	}

	if unix != unix2 {
		t.Fatalf("Round-trip conversion failed: %d != %d", unix, unix2)
	}
}
