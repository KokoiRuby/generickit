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

package setx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ContainsTestSuite struct {
	suite.Suite
	ms *MapSet[int]
}

func (suite *ContainsTestSuite) SetupTest() {
	suite.ms = NewMapSet[int]()
	suite.ms.Add(1)
	suite.ms.Add(2)
	suite.ms.Add(3)
}

func (suite *ContainsTestSuite) TestAdd() {
	suite.ms.Add(4)
	assert.True(suite.T(), suite.ms.Contains(4))
	suite.ms.Add(3)
	assert.True(suite.T(), suite.ms.Contains(3))
	assert.Equal(suite.T(), 4, suite.ms.Size())
}

func (suite *ContainsTestSuite) TestRemove() {
	suite.ms.Remove(3)
	assert.Equal(suite.T(), 2, suite.ms.Size())
	suite.ms.Remove(5)
	assert.Equal(suite.T(), 2, suite.ms.Size())
}

func (suite *ContainsTestSuite) TestContains() {
	assert.True(suite.T(), suite.ms.Contains(3))
	assert.False(suite.T(), suite.ms.Contains(4))
}

func (suite *ContainsTestSuite) TestSize() {
	assert.Equal(suite.T(), 3, suite.ms.Size())
}

func (suite *ContainsTestSuite) TestClear() {
	suite.ms.Clear()
	assert.Equal(suite.T(), 0, suite.ms.Size())
}

func TestContainsTestSuite(t *testing.T) {
	suite.Run(t, new(ContainsTestSuite))
}
