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

package validators

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func TestStringOneOfValidator(t *testing.T) {
	validators := StringOneOfValidator("value1", "value2", "value3")

	if validators == nil {
		t.Fatal("StringOneOfValidator should not return nil")
	}

	if len(validators) != 1 {
		t.Fatalf("StringOneOfValidator should return a slice with 1 validator, got %d", len(validators))
	}

	if validators[0] == nil {
		t.Fatal("StringOneOfValidator should return a non-nil validator")
	}

	// Verify it implements validator.String interface
	var _ validator.String = validators[0]
}

func TestStringOneOfValidator_MultipleValues(t *testing.T) {
	values := []string{"one", "two", "three", "four", "five"}
	validators := StringOneOfValidator(values...)

	if len(validators) != 1 {
		t.Fatalf("StringOneOfValidator should return a slice with 1 validator, got %d", len(validators))
	}
}

func TestStringOneOfValidator_SingleValue(t *testing.T) {
	validators := StringOneOfValidator("single")

	if len(validators) != 1 {
		t.Fatalf("StringOneOfValidator should return a slice with 1 validator, got %d", len(validators))
	}

	if validators[0] == nil {
		t.Fatal("StringOneOfValidator should return a non-nil validator")
	}
}
