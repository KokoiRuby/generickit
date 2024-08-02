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
	"testing"
)

type FindTestSuite struct {
	suite.Suite
	sl []int
}

func (suite *FindTestSuite) SetupTest() {
	suite.sl = []int{1, 2, 3, 4, 5, 5}
}

func (suite *FindTestSuite) TestFind() {
	idx, isFound := Find(suite.sl, 3)
	assert.Equal(suite.T(), 2, idx)
	assert.True(suite.T(), isFound)

	idx, isFound = Find(suite.sl, 0)
	assert.Equal(suite.T(), -1, idx)
	assert.False(suite.T(), isFound)
}

func (suite *FindTestSuite) TestFindLast() {
	idx, isFound := FindLast(suite.sl, 5)
	assert.Equal(suite.T(), 5, idx)
	assert.True(suite.T(), isFound)

	idx, isFound = FindLast(suite.sl, 0)
	assert.Equal(suite.T(), -1, idx)
	assert.False(suite.T(), isFound)
}

func (suite *FindTestSuite) TestFindAll() {
	idxSlice, isFound := FindAll(suite.sl, 5)
	assert.Equal(suite.T(), []int{4, 5}, idxSlice)
	assert.True(suite.T(), isFound)

	idxSlice, isFound = FindAll(suite.sl, 0)
	assert.Equal(suite.T(), []int{}, idxSlice)
	assert.False(suite.T(), isFound)
}

func ExampleFind() {
	idx, isFound := Find[int]([]int{1, 2, 3, 4, 5, 5}, 3)
	fmt.Println(idx)
	fmt.Println(isFound)
	// output:
	// 2
	// true
}

func ExampleFindLast() {
	idx, isFound := Find[int]([]int{1, 2, 3, 4, 5, 5}, 5)
	fmt.Println(idx)
	fmt.Println(isFound)
	// output:
	// 5
	// true
}

func ExampleFindAll() {
	idxSlice, isFound := FindAll[int]([]int{1, 2, 3, 4, 5, 5}, 5)
	fmt.Println(idxSlice)
	fmt.Println(isFound)
	// output:
	// [4 5]
	// true
}

func TestFindTestSuite(t *testing.T) {
	suite.Run(t, new(FindTestSuite))
}
