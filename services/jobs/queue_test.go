// Copyright 2022 Evan Hazlett
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package render

import "testing"

func TestCalculateRenderSlices4(t *testing.T) {
	expectedSlices := 4
	results, err := calculateRenderSlices(expectedSlices)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != expectedSlices {
		t.Fatalf("expected %d slices; received %d", expectedSlices, len(results))
	}
}

func TestCalculateRenderSlices5(t *testing.T) {
	expectedSlices := 6
	results, err := calculateRenderSlices(5)
	if err != nil {
		t.Fatal(err)
	}
	if v := len(results); v != expectedSlices {
		t.Fatalf("expected %d results; received %d", expectedSlices, v)
	}
}

func TestCalculateRenderSlices8(t *testing.T) {
	expectedSlices := 8
	results, err := calculateRenderSlices(expectedSlices)
	if err != nil {
		t.Fatal(err)
	}
	if v := len(results); v != expectedSlices {
		t.Fatalf("expected %d results; received %d", expectedSlices, v)
	}
}

func TestCalculateRenderSlices10(t *testing.T) {
	expectedSlices := 10
	results, err := calculateRenderSlices(expectedSlices)
	if err != nil {
		t.Fatal(err)
	}
	if v := len(results); v != expectedSlices {
		t.Fatalf("expected %d results; received %d", expectedSlices, v)
	}
}
