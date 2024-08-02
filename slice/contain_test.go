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

type ContainsTestSuite struct {
	suite.Suite
	sl []int
}

func (suite *ContainsTestSuite) SetupTest() {
	suite.sl = []int{1, 2, 3, 4, 5}
}

func (suite *ContainsTestSuite) TestContains() {
	assert.Equal(suite.T(), true, Contains(suite.sl, 3))
	assert.Equal(suite.T(), false, Contains(suite.sl, 6))
}

func ExampleContains() {
	fmt.Println(Contains([]int{1, 2, 3, 4, 5}, 5))
	// output:
	// true
}

func TestContainsTestSuite(t *testing.T) {
	suite.Run(t, new(ContainsTestSuite))
}
