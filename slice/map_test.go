/*
 * Copyright 2024 KoKoiRuby
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type MapTestSuite struct {
	suite.Suite
	sl []int
}

func (suite *MapTestSuite) SetupTest() {
	suite.sl = []int{1, 2, 3, 4, 5}
}

func (suite *MapTestSuite) TestMap() {
	after := Map(suite.sl, func(i int) string {
		return strconv.Itoa(i)
	})
	assert.Equal(suite.T(), []string{"1", "2", "3", "4", "5"}, after)
}

func (suite *MapTestSuite) TestFilterMap() {
	after := FilterMap(suite.sl, func(src int) (dst string, filter bool) {
		return strconv.Itoa(src), src >= 3
	})
	assert.Equal(suite.T(), []string{"3", "4", "5"}, after)
}

func ExampleMap() {
	after := Map[int]([]int{1, 2, 3, 4, 5}, func(src int) string { return strconv.Itoa(src) })
	fmt.Println(after)
	// output:
	// [1 2 3 4 5]
}

func ExampleFilterMap() {
	after := FilterMap[int]([]int{1, 2, 3, 4, 5}, func(src int) (string, bool) { return strconv.Itoa(src), src >= 3 })
	fmt.Println(after)
	// output:
	// [3 4 5]
}

func TestMapTestSuite(t *testing.T) {
	suite.Run(t, new(MapTestSuite))
}
