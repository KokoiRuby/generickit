package slice

import (
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

func TestShrinkTTestSuite(t *testing.T) {
	suite.Run(t, new(ShrinkTestSuite))
}
