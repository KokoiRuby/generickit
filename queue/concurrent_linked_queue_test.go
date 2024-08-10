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

package queue

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConcurrentLinkedQueueTestSuite struct {
	suite.Suite
	q *ConcurrentLinkedQueue[int] // unbound
}

func (suite *ConcurrentLinkedQueueTestSuite) SetupTest() {
	suite.q = NewConcurrentLinkedQueue[int]()
}

func (suite *ConcurrentLinkedQueueTestSuite) TestEnqueue() {
	err := suite.q.Enqueue(1)
	assert.Nil(suite.T(), err)
	err = suite.q.Enqueue(2)
	assert.Nil(suite.T(), err)
	err = suite.q.Enqueue(3)
	assert.Nil(suite.T(), err)
}

func (suite *ConcurrentLinkedQueueTestSuite) TestDequeue() {
	v, err := suite.q.Dequeue()
	assert.Errorf(suite.T(), err, "empty queue")
	assert.Equal(suite.T(), 0, v)

	go func() {
		for i := 0; i < 5; i++ {
			_ = suite.q.Enqueue(i)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			v, err = suite.q.Dequeue()
			assert.Nil(suite.T(), err)
			assert.Equal(suite.T(), i, v)
		}
	}()

}

func (suite *ConcurrentLinkedQueueTestSuite) TestToSlice() {
	_ = suite.q.Enqueue(1)
	_ = suite.q.Enqueue(2)
	_ = suite.q.Enqueue(3)
	assert.Equal(suite.T(), []int{1, 2, 3}, suite.q.ToSlice())
}

func TestConcurrentLinkedQueueTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(ConcurrentLinkedQueueTestSuite))
}
