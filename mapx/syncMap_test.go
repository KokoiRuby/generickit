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

package mapx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SyncMapTestSuite struct {
	suite.Suite
	m *SyncMap[string, int]
}

func (suite *SyncMapTestSuite) SetupTest() {
	suite.m = NewSyncMap[string, int]()
	err := suite.m.Put("a", 1)
	if err != nil {
		return
	}
	err = suite.m.Put("b", 2)
	if err != nil {
		return
	}
	err = suite.m.Put("c", 3)
	if err != nil {
		return
	}
}

func (suite *SyncMapTestSuite) TestGet() {
	v, ok := suite.m.Get("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 1, v)
	v, ok = suite.m.Get("b")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 2, v)
	v, ok = suite.m.Get("c")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 3, v)
}

func (suite *SyncMapTestSuite) TestDelete() {
	v, ok := suite.m.Delete("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 1, v)
	v, ok = suite.m.Delete("d")
	assert.False(suite.T(), ok)
	assert.Equal(suite.T(), 0, v)
}

func (suite *SyncMapTestSuite) TestKeys() {
	suite.T().Log(suite.m.Keys())
}

func (suite *SyncMapTestSuite) TestValues() {
	suite.T().Log(suite.m.Values())
}

func (suite *SyncMapTestSuite) TestLen() {
	assert.Equal(suite.T(), uint64(3), suite.m.Len())
}

func TestSyncMapTestSuite(t *testing.T) {
	suite.Run(t, new(SyncMapTestSuite))
}
