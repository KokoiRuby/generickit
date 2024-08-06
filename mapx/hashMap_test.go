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
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"hash/fnv"
	"testing"
)

type User struct {
	id   uint
	name string
	age  uint8
}

func (u User) Hash() uint64 {
	h := fnv.New64a()
	_, err := h.Write([]byte(fmt.Sprintf("%d", u.id)))
	if err != nil {
		return 0
	}
	return h.Sum64()
}

func (u User) Equal(k any) bool {
	// type assertion
	if ku, ok := k.(User); ok {
		if ku.id == u.id && ku.name == u.name && ku.age == u.age {
			return true
		}
	}
	return false
}

type HashMapTestSuite struct {
	suite.Suite
	m *HashMap[User, int]
}

func (suite *HashMapTestSuite) SetupTest() {
	suite.m = NewHashMap[User, int](5)
	_ = suite.m.Put(User{id: 0, name: "Alice", age: 10}, 0)
	_ = suite.m.Put(User{id: 1, name: "Bob", age: 20}, 1)
}

func (suite *HashMapTestSuite) TestPut() {
	err := suite.m.Put(User{id: 2, name: "Cindy", age: 30}, 2)
	assert.Nil(suite.T(), err)
}

func (suite *HashMapTestSuite) TestGet() {
	v, isFound := suite.m.Get(User{id: 0, name: "Alice", age: 10})
	assert.True(suite.T(), isFound)
	assert.Equal(suite.T(), 0, v)
	v, isFound = suite.m.Get(User{id: 0, name: "Alice", age: 100})
	assert.False(suite.T(), isFound)
	assert.Equal(suite.T(), 0, v)
	v, isFound = suite.m.Get(User{id: 1, name: "Bob", age: 20})
	assert.True(suite.T(), isFound)
	assert.Equal(suite.T(), 1, v)
	v, isFound = suite.m.Get(User{id: 2, name: "Cindy", age: 30})
	assert.False(suite.T(), isFound)
	assert.Equal(suite.T(), 0, v)

}

func (suite *HashMapTestSuite) TestDelete() {
	v, isDeleted := suite.m.Delete(User{id: 0, name: "Alice", age: 10})
	assert.True(suite.T(), isDeleted)
	assert.Equal(suite.T(), 0, v)
	v, isFound := suite.m.Get(User{id: 0, name: "Alice", age: 10})
	assert.False(suite.T(), isFound)
	assert.Equal(suite.T(), 0, v)
}

func (suite *HashMapTestSuite) TestKeys() {
	users := suite.m.Keys()
	assert.Equal(suite.T(), 2, len(users))
	for _, u := range users {
		suite.T().Logf("%+v", u)
	}
}

func (suite *HashMapTestSuite) TestValues() {
	vals := suite.m.Values()
	assert.Equal(suite.T(), 2, len(vals))
	for _, val := range vals {
		suite.T().Logf("%+v", val)
	}
}

func (suite *HashMapTestSuite) TestLen() {
	assert.Equal(suite.T(), uint64(2), suite.m.Len())
}

func TestHashMapTestSuite(t *testing.T) {
	suite.Run(t, new(HashMapTestSuite))
}
