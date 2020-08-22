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

package tools_test

import (
	"fmt"
	"testing"
	"tools"

	"github.com/stretchr/testify/assert"
)

func TestUnixIntToTimeStr(t *testing.T) {
	timeStr := tools.Date(1598094878, "YYYY-MM-DD HH:mm:ss")
	fmt.Printf("time: %s\n", timeStr)
	assert.Equal(t, "2020-08-22 19:14:38", timeStr)
}

func TestUnixStrToTimeStr(t *testing.T) {
	timeStr := tools.DateS("1598094878", "YYYY-MM-DD HH:mm:ss")
	fmt.Printf("time: %s\n", timeStr)
	assert.Equal(t, "2020-08-22 19:14:38", timeStr)
}

func TestParseStrToDateTime(t *testing.T) {
	value, err := tools.ParseStrToDateTime("2019-01-02 15:22:05", "Asia/Shanghai")
	assert.Equal(t, nil, err)
	zoneName, _ := value.Zone()
	assert.Equal(t, 2019, value.Year())
	assert.Equal(t, 2, value.Day())
	assert.Equal(t, "CST", zoneName)
	fmt.Printf("value: %v\n", value)
}

func TestParseStrToDate(t *testing.T) {
	value, err := tools.ParseStrToDate("2020-10-22", "Asia/Shanghai")
	assert.Equal(t, nil, err)
	zoneName, _ := value.Zone()
	assert.Equal(t, 2020, value.Year())
	assert.Equal(t, 22, value.Day())
	assert.Equal(t, "CST", zoneName)
	fmt.Printf("value: %v\n", value)
}