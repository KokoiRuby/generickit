package slice

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DeleteTestSuite struct {
	suite.Suite
	sl []int
}

func (suite *DeleteTestSuite) SetupTest() {
	suite.sl = []int{1, 2, 3, 4, 5}
}

func (suite *DeleteTestSuite) TestDelete() {
	var err error
	suite.sl, err = Delete(suite.sl, -1)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(6, -1))

	suite.sl, err = Delete(suite.sl, 9)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(6, 9))

	suite.sl, err = Delete(suite.sl, 0)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{2, 3, 4, 5}, suite.sl)

	suite.sl, err = Delete(suite.sl, 3)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{2, 3, 4}, suite.sl)

	suite.sl, err = Delete(suite.sl, 1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{2, 4}, suite.sl)
}

func (suite *DeleteTestSuite) TestDeleteAfter() {
	var err error
	suite.sl, err = DeleteAfter(suite.sl, -1)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(6, -1))

	suite.sl, err = DeleteAfter(suite.sl, 9)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(6, 9))

	suite.sl, err = DeleteAfter(suite.sl, 4)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{1, 2, 3, 4}, suite.sl)

	suite.sl, err = DeleteAfter(suite.sl, 2)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{1, 2}, suite.sl)

	suite.sl, err = DeleteAfter(suite.sl, 0)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{}, suite.sl)
}

func (suite *DeleteTestSuite) TestDeleteRange() {
	var err error
	suite.sl, err = DeleteRange(suite.sl, 1, -1)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(6, -1))

	suite.sl, err = DeleteRange(suite.sl, 1, 9)
	assert.Error(suite.T(), err, ErrIdxOutOfRange(6, 9))

	suite.sl, err = DeleteRange(suite.sl, 3, 1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{1, 5}, suite.sl)

	suite.sl = []int{1, 2, 3, 4, 5}
	suite.sl, err = DeleteRange(suite.sl, 3, 0)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{4, 5}, suite.sl)

	suite.sl = []int{1, 2, 3, 4, 5}
	suite.sl, err = DeleteRange(suite.sl, 5, 2)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{1, 2}, suite.sl)
}

func TestDeleteTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
