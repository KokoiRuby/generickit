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

package syncx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MapTestSuite struct {
	suite.Suite
	m *Map[string, int]
}

func (suite *MapTestSuite) SetupTest() {
	suite.m = NewMap[string, int]()
	suite.m.Store("a", 1)
	suite.m.Store("b", 2)
	suite.m.Store("c", 3)
}

func (suite *MapTestSuite) TestLoad() {
	v, ok := suite.m.Load("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 1, v)
	v, ok = suite.m.Load("b")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 2, v)
	v, ok = suite.m.Load("c")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 3, v)
	v, ok = suite.m.Load("d")
	assert.False(suite.T(), ok)
	assert.Equal(suite.T(), 0, v)
}

func (suite *MapTestSuite) TestDelete() {
	suite.m.Delete("a")
	v, ok := suite.m.Load("a")
	assert.False(suite.T(), ok)
	assert.Equal(suite.T(), 0, v)
}

func (suite *MapTestSuite) TestCompareAndDelete() {
	ok := suite.m.CompareAndDelete("a", 11)
	assert.False(suite.T(), ok)
	ok = suite.m.CompareAndDelete("a", 1)
	assert.True(suite.T(), ok)
}

func (suite *MapTestSuite) TestCompareAndSwap() {
	ok := suite.m.CompareAndSwap("a", 11, 11)
	assert.False(suite.T(), ok)
	ok = suite.m.CompareAndSwap("a", 1, 11)
	assert.True(suite.T(), ok)
	v, ok := suite.m.Load("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 11, v)
}

func (suite *MapTestSuite) TestLoadAndDelete() {
	v, ok := suite.m.LoadAndDelete("d")
	assert.False(suite.T(), ok)
	assert.Equal(suite.T(), 0, v)
	v, ok = suite.m.LoadAndDelete("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 1, v)
	v, ok = suite.m.Load("a")
	assert.False(suite.T(), ok)
	assert.Equal(suite.T(), 0, v)
}

func (suite *MapTestSuite) TestLoadOrStore() {
	v, ok := suite.m.LoadOrStore("d", 4)
	assert.False(suite.T(), ok)
	assert.Equal(suite.T(), 4, v)
	v, ok = suite.m.LoadOrStore("a", 11)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 1, v)
}

func (suite *MapTestSuite) TestRange() {
	suite.m.Range(func(k string, v int) bool {
		fmt.Printf("%v: %v\n", k, v)
		return true
	})
}

func (suite *MapTestSuite) TestSwap() {
	v, ok := suite.m.Swap("a", 11)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 1, v)
	v, ok = suite.m.Load("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 11, v)
}

func TestMapTestSuite(t *testing.T) {
	suite.Run(t, new(MapTestSuite))
}
