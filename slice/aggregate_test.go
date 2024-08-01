package slice

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AggregateTestSuite struct {
	suite.Suite
	sl []int
}

func (suite *AggregateTestSuite) SetupTest() {
	suite.sl = []int{1, 2, 3, 4, 5}
}

func (suite *AggregateTestSuite) TestMax() {
	assert.Equal(suite.T(), 5, Max(suite.sl))
}

func (suite *AggregateTestSuite) TestMin() {
	assert.Equal(suite.T(), 1, Min(suite.sl))
}

func (suite *AggregateTestSuite) TestSum() {
	assert.Equal(suite.T(), 15, Sum(suite.sl))
}

func TestAggregateTestSuite(t *testing.T) {
	suite.Run(t, new(AggregateTestSuite))
}
