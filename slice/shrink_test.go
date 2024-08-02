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

type ShrinkTestSuite struct {
	suite.Suite
	sl1 []int
	sl2 []int
}

func (suite *ShrinkTestSuite) SetupTest() {
	suite.sl1 = make([]int, 300, 480)
	suite.sl2 = make([]int, 100, 300)
}

func (suite *ShrinkTestSuite) TestShrink() {
	suite.sl1 = Shrink(suite.sl1)
	assert.Equal(suite.T(), 384, cap(suite.sl1))
	suite.sl2 = Shrink(suite.sl2)
	assert.Equal(suite.T(), 150, cap(suite.sl2))
}

func ExampleShrink() {
	sl := make([]int, 100, 300)
	sl = Shrink[int](sl)
	fmt.Println(cap(sl))
	sl = make([]int, 300, 480)
	sl = Shrink[int](sl)
	fmt.Println(cap(sl))
	// output:
	// 150
	// 384
}

func TestShrinkTTestSuite(t *testing.T) {
	suite.Run(t, new(ShrinkTestSuite))
}
