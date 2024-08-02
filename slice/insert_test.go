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
	"testing"

	"github.com/stretchr/testify/suite"
)

type InsertTestSuite struct {
	suite.Suite
	sl []int
}

func (suite *InsertTestSuite) SetupTest() {
	suite.sl = []int{1, 2, 3, 4, 5}
}

func (suite *InsertTestSuite) TestInsert() {
	var err error
	suite.sl, err = Insert(suite.sl, 0, 0)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{0, 1, 2, 3, 4, 5}, suite.sl)

	suite.sl, err = Insert(suite.sl, -1, -1)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(6, -1))

	suite.sl, err = Insert(suite.sl, 9, 9)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(6, 9))

	suite.sl, err = Insert(suite.sl, 0, 3)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{0, 1, 2, 0, 3, 4, 5}, suite.sl)

}

func (suite *InsertTestSuite) TestInsertSlice() {
	var err error
	newSl := []int{0, 1}
	suite.sl, err = InsertSlice(suite.sl, []int{}, 0)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{1, 2, 3, 4, 5}, suite.sl)

	suite.sl, err = InsertSlice(suite.sl, newSl, 0)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{0, 1, 1, 2, 3, 4, 5}, suite.sl)

	suite.sl, err = InsertSlice(suite.sl, newSl, -1)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(9, -1))

	suite.sl, err = InsertSlice(suite.sl, newSl, 9)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(8, 9))

	suite.sl, err = InsertSlice(suite.sl, newSl, 3)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{0, 1, 1, 0, 1, 2, 3, 4, 5}, suite.sl)

}

func ExampleInsert() {
	res, _ := Insert[int]([]int{1, 2, 3}, 0, 1)
	fmt.Println(res)
	// Output:
	// [1 0 2 3]
}

func ExampleInsertSlice() {
	res, _ := InsertSlice[int]([]int{1, 2, 3}, []int{4, 5}, 1)
	fmt.Println(res)
	// Output:
	// [1 4 5 2 3]
}

func TestInsertTTestSuite(t *testing.T) {
	suite.Run(t, new(InsertTestSuite))
}
