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
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuxiliaryTestSuite struct {
	suite.Suite
	m     map[string]int
	empty map[string]int
	kSl   []string
	vSl   []int
}

func (suite *AuxiliaryTestSuite) SetupTest() {
	suite.m = make(map[string]int, 3)
	suite.m["a"] = 1
	suite.m["b"] = 2
	suite.m["c"] = 3

	suite.kSl = []string{"a", "b", "c"}
	suite.vSl = []int{1, 2, 3}
}

func (suite *AuxiliaryTestSuite) TestKeys() {
	kSl := Keys(suite.m)
	suite.T().Log(kSl)
}

func (suite *AuxiliaryTestSuite) TestValues() {
	vSl := Values(suite.m)
	suite.T().Log(vSl)
}

func (suite *AuxiliaryTestSuite) TestKeysValues() {
	kSl, vSl := KeysValues(suite.m)
	suite.T().Log(kSl)
	suite.T().Log(vSl)
}

func (suite *AuxiliaryTestSuite) TestToMap() {
	m, _ := ToMap(suite.kSl, suite.vSl)
	suite.T().Log(m)
}

func TestAuxiliaryTestSuite(t *testing.T) {
	suite.Run(t, new(AuxiliaryTestSuite))
}
