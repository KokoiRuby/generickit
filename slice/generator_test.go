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

type GeneratorTestSuite struct {
	suite.Suite
	sl []int
}

func (suite *GeneratorTestSuite) SetupTest() {
	suite.sl = []int{1, 2, 3, 4, 5}
}

func (suite *GeneratorTestSuite) TestGenerator() {
	intGenerator := Generator(suite.sl)
	assert.Equal(suite.T(), <-intGenerator, suite.sl[0])
	assert.Equal(suite.T(), <-intGenerator, suite.sl[1])
	assert.Equal(suite.T(), <-intGenerator, suite.sl[2])
	assert.Equal(suite.T(), <-intGenerator, suite.sl[3])
	assert.Equal(suite.T(), <-intGenerator, suite.sl[4])
	assert.Equal(suite.T(), <-intGenerator, 0)
}

func ExampleGenerator() {
	intGenerator := Generator[int]([]int{1, 2, 3})
	fmt.Println(<-intGenerator)
	fmt.Println(<-intGenerator)
	fmt.Println(<-intGenerator)
	fmt.Println(<-intGenerator)
	// output:
	// 1
	// 2
	// 3
	// 0
}

func TestGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(GeneratorTestSuite))
}
