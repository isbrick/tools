// Copyright 2020 tools authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSliceContainsStr(t *testing.T) {
	result := IsSliceContainsStr([]string{"hello", "world"}, "worLd")
	assert.Equal(t, true, result)
}

func TestIsSliceContainsInt64(t *testing.T) {
	result := IsSliceContainsInt64([]int64{12, 13}, 12)
	assert.Equal(t, true, result)
}

func TestIsSliceContainsInt(t *testing.T) {
	result := IsSliceContainsInt([]int{12, 13}, 12)
	assert.Equal(t, true, result)
}

func TestIntSliceUnion(t *testing.T) {
	result := IntSliceUnion([]int{12, 13, 1}, []int{12, 13, 2})
	assert.Equal(t, 4, len(result))
}

func TestIntSliceDifference(t *testing.T) {
	result := IntSliceDifference([]int{12, 13, 1}, []int{12, 13, 2})
	assert.Equal(t, 1, result[0])
}

func TestIntSliceIntersect(t *testing.T) {
	result := IntSliceIntersect([]int{12, 13, 1}, []int{12, 13, 2})
	assert.Equal(t, 12, result[0])
	assert.Equal(t, 13, result[1])
}
