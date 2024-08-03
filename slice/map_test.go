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

func (suite *MapTestSuite) TestFilter() {
	after := Filter(suite.sl, func(src int) (dst string, filter bool) {
		return strconv.Itoa(src), src >= 3
	})
	assert.Equal(suite.T(), []string{"3", "4", "5"}, after)
}

func (suite *MapTestSuite) TestReduce1() {
	after := Reduce(suite.sl, func(src int) int {
		return src
	})
	assert.Equal(suite.T(), 15, after)
}

func (suite *MapTestSuite) TestReduce2() {
	after := Reduce(suite.sl, func(src int) int {
		if src > 3 {
			return src
		}
		return 0
	})
	assert.Equal(suite.T(), 9, after)
}

func ExampleMap() {
	after := Map[int, string]([]int{1, 2, 3, 4, 5}, func(src int) string { return strconv.Itoa(src) })
	fmt.Println(after)
	// output:
	// [1 2 3 4 5]
}

func ExampleFilter() {
	after := Filter[int, string]([]int{1, 2, 3, 4, 5}, func(src int) (string, bool) { return strconv.Itoa(src), src >= 3 })
	fmt.Println(after)
	// output:
	// [3 4 5]
}

func ExampleReduce() {
	after := Reduce[int, int]([]int{1, 2, 3, 4, 5}, func(src int) int { return src })
	fmt.Println(after)
	// output:
	// 15
}

func TestMapTestSuite(t *testing.T) {
	suite.Run(t, new(MapTestSuite))
}
