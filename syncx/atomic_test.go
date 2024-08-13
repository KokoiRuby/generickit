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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ValueSuite struct {
	suite.Suite
	m *Value[int]
}

func (suite *ValueSuite) SetupTest() {
	suite.m = NewValue(0)
}

func (suite *ValueSuite) TestLoad() {
	assert.Equal(suite.T(), suite.m.Load(), 0)
}
func (suite *ValueSuite) TestStore() {
	suite.m.Store(1)
	assert.Equal(suite.T(), suite.m.Load(), 1)
}

func (suite *ValueSuite) TestSwap() {
	assert.Equal(suite.T(), 0, suite.m.Swap(1))
	assert.Equal(suite.T(), suite.m.Load(), 1)
}
func (suite *ValueSuite) TestCompareAndSwap() {
	swapped := suite.m.CompareAndSwap(0, 1)
	assert.True(suite.T(), swapped)
	swapped = suite.m.CompareAndSwap(0, 1)
	assert.Equal(suite.T(), suite.m.Load(), 1)
	assert.False(suite.T(), swapped)

}

func TestValueSuite(t *testing.T) {
	suite.Run(t, new(ValueSuite))
}
