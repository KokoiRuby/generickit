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

type PoolTestSuite struct {
	suite.Suite
	pool *Pool[[]byte]
}

func (suite *PoolTestSuite) SetupTest() {
	suite.pool = NewPool[[]byte](func() []byte {
		sl := make([]byte, 1, 3)
		sl[0] = 'a'
		return sl
	})
}

func (suite *PoolTestSuite) TestGet() {
	assert.Equal(suite.T(), "a", string(suite.pool.Get()))
}

func TestPoolTestSuite(t *testing.T) {
	suite.Run(t, new(PoolTestSuite))
}
