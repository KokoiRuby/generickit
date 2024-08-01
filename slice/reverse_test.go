package slice

import (
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

func TestReverseTestSuite(t *testing.T) {
	suite.Run(t, new(ReverseTestSuite))
}
