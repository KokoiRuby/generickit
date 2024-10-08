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

type ConcurrentArrayBlockingQueueTestSuite struct {
	suite.Suite
	q *ConcurrentArrayBlockingQueue[int] // bound
}

func (suite *ConcurrentArrayBlockingQueueTestSuite) SetupTest() {
	suite.q = NewConcurrentArrayBlockingQueue[int](3)
}

func (suite *ConcurrentArrayBlockingQueueTestSuite) TestEnqueue() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		for i := 0; i < 3; i++ {
			err := suite.q.Enqueue(ctx, i)
			assert.Nil(suite.T(), err)
		}
		time.Sleep(2 * time.Second)
		err := suite.q.Enqueue(ctx, 3)
		assert.Equal(suite.T(), context.DeadlineExceeded, err)
	}()

	// trigger ctxt timeout
	time.Sleep(1 * time.Second)
	v, err := suite.q.Dequeue(ctx)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), v, 0)
}

func (suite *ConcurrentArrayBlockingQueueTestSuite) TestDequeue() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := suite.q.Enqueue(ctx, 1)
	assert.Nil(suite.T(), err)
	err = suite.q.Enqueue(ctx, 2)
	assert.Nil(suite.T(), err)
	err = suite.q.Enqueue(ctx, 3)
	assert.Nil(suite.T(), err)

	go func() {
		for i := 1; i < 4; i++ {
			v, err := suite.q.Dequeue(ctx)
			assert.Nil(suite.T(), err)
			assert.Equal(suite.T(), i, v)
		}
		time.Sleep(2 * time.Second)
		_, err = suite.q.Dequeue(ctx)
		assert.Equal(suite.T(), context.DeadlineExceeded, err)
	}()

	time.Sleep(1 * time.Second)
	err = suite.q.Enqueue(ctx, 4)
	assert.Nil(suite.T(), err)
}

func (suite *ConcurrentArrayBlockingQueueTestSuite) TestLen() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ {
		err := suite.q.Enqueue(ctx, i)
		assert.Nil(suite.T(), err)
	}
	assert.Equal(suite.T(), suite.q.Len(), 3)

	v, err := suite.q.Dequeue(ctx)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, v)
	err = suite.q.Enqueue(ctx, 3)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.q.Len(), 3)
}

func (suite *ConcurrentArrayBlockingQueueTestSuite) TestToSlice() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ {
		err := suite.q.Enqueue(ctx, i)
		assert.Nil(suite.T(), err)
	}
	assert.Equal(suite.T(), []int{0, 1, 2}, suite.q.ToSlice())

	v, err := suite.q.Dequeue(ctx)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, v)
	err = suite.q.Enqueue(ctx, 3)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []int{1, 2, 3}, suite.q.ToSlice())
}

func TestConcurrentArrayBlockingQueueTestSuite(t *testing.T) {
	suite.Run(t, new(ConcurrentArrayBlockingQueueTestSuite))
}
