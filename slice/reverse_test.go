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

type ReverseTestSuite struct {
	suite.Suite
	sl1 []int
	sl2 []int
}

func (suite *ReverseTestSuite) SetupTest() {
	suite.sl1 = []int{1, 2, 3, 4, 5}
	suite.sl2 = []int{1, 2, 3, 4, 5, 6}
}

func (suite *ReverseTestSuite) TestReverse() {
	assert.Equal(suite.T(), []int{5, 4, 3, 2, 1}, Reverse(suite.sl1))
	assert.Equal(suite.T(), []int{6, 5, 4, 3, 2, 1}, Reverse(suite.sl2))
}

func ExampleReverse() {
	res := Reverse[int]([]int{5, 4, 3, 2, 1})
	fmt.Println(res)
	// output:
	// [1 2 3 4 5]
}

func TestReverseTestSuite(t *testing.T) {
	suite.Run(t, new(ReverseTestSuite))
}
