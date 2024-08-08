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
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ConcurrentLinkedBlockingQueueTestSuite struct {
	suite.Suite
	q1 *ConcurrentLinkedBlockingQueue[int] // bound
	q2 *ConcurrentLinkedBlockingQueue[int] // unbound
}

func (suite *ConcurrentLinkedBlockingQueueTestSuite) SetupTest() {
	suite.q1 = NewConcurrentLinkedBlockingQueue[int](3)
	suite.q2 = NewConcurrentLinkedBlockingQueue[int](-1)
}

func (suite *ConcurrentLinkedBlockingQueueTestSuite) TestEnqueue() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := suite.q1.Enqueue(ctx, 1)
	assert.Nil(suite.T(), err)
	err = suite.q1.Enqueue(ctx, 2)
	assert.Nil(suite.T(), err)
	err = suite.q1.Enqueue(ctx, 3)
	assert.Nil(suite.T(), err)
	err = suite.q1.Enqueue(ctx, 4)
	assert.Equal(suite.T(), context.DeadlineExceeded, err)
}

func (suite *ConcurrentLinkedBlockingQueueTestSuite) TestDequeue() {

}

func TestConcurrentLinkedBlockingQueueTestSuite(t *testing.T) {
	suite.Run(t, new(ConcurrentLinkedBlockingQueueTestSuite))
}
