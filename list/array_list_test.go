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

package list

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ArrayListTestSuite struct {
	suite.Suite
	l *ArrayList[int]
}

func (suite *ArrayListTestSuite) SetupTest() {
	suite.l = NewArrayListFrom[int]([]int{1, 2, 3, 4, 5})
}

func (suite *ArrayListTestSuite) TestImplemented() {
	var _ List[int] = suite.l
}

func (suite *ArrayListTestSuite) TestGet() {
	v, err := suite.l.Get(2)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 3, v)
}

func (suite *ArrayListTestSuite) TestAppend() {
	suite.l.Append(6)
	assert.Equal(suite.T(), []int{1, 2, 3, 4, 5, 6}, suite.l.ToSlice())
	suite.l.Append(7, 8)
	assert.Equal(suite.T(), []int{1, 2, 3, 4, 5, 6, 7, 8}, suite.l.ToSlice())
}

func (suite *ArrayListTestSuite) TestAdd() {
	err := suite.l.Add(0, 0)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []int{0, 1, 2, 3, 4, 5}, suite.l.ToSlice())
}

func (suite *ArrayListTestSuite) TestSet() {
	err := suite.l.Set(2, 33)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []int{1, 2, 33, 4, 5}, suite.l.ToSlice())
}

func (suite *ArrayListTestSuite) TestDelete() {
	v, err := suite.l.Delete(2)
	assert.Equal(suite.T(), 3, v)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []int{1, 2, 4, 5}, suite.l.ToSlice())
}

func (suite *ArrayListTestSuite) TestLen() {
	assert.Equal(suite.T(), 5, suite.l.Len())
}

func (suite *ArrayListTestSuite) TestCap() {
	assert.Equal(suite.T(), 5, suite.l.Cap())
}

func (suite *ArrayListTestSuite) TestRange() {
	sum := 0
	assert.Nil(suite.T(), suite.l.Range(func(idx int, val int) error {
		sum += val
		return nil
	}))
	assert.Equal(suite.T(), 15, sum)
}

func (suite *ArrayListTestSuite) TestToSlice() {
	assert.Equal(suite.T(), []int{1, 2, 3, 4, 5}, suite.l.ToSlice())
}

func (suite *ArrayListTestSuite) TestGenerator() {
	gen := suite.l.Generator()
	assert.Equal(suite.T(), 1, <-gen)
	assert.Equal(suite.T(), 2, <-gen)
	assert.Equal(suite.T(), 3, <-gen)
	assert.Equal(suite.T(), 4, <-gen)
	assert.Equal(suite.T(), 5, <-gen)
	assert.Equal(suite.T(), 0, <-gen)
}

func TestArrayListTestSuite(t *testing.T) {
	suite.Run(t, new(ArrayListTestSuite))
}
