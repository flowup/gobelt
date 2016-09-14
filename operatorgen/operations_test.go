package operatorgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Expecteds results from tests
var (
	MapResultArray1    = TTypeSequence{&TType{3}, &TType{6}, &TType{15}, &TType{9}, &TType{12}}
	MapResultArray2    = TTypeSequence{&TType{1}, &TType{4}, &TType{25}, &TType{9}, &TType{16}}
	FilterResultArray1 = TTypeSequence{&TType{2}, &TType{4}}
	FilterResultArray2 = TTypeSequence{&TType{1}, &TType{5}, &TType{3}}
	ReverseResultArray = TTypeSequence{&TType{4}, &TType{3}, &TType{5}, &TType{2}, &TType{1}}
	SortedResultArray  = TTypeSequence{&TType{1}, &TType{2}, &TType{3}, &TType{4}, &TType{5}}
	PushResultArray    = TTypeSequence{&TType{1}, &TType{2}, &TType{5}, &TType{3}, &TType{4}, &TType{100}}
	PopResultArray     = TTypeSequence{&TType{1}, &TType{2}, &TType{5}, &TType{3}}
)

// TemplateTestSuite is testing suite
type TemplateTestSuite struct {
	suite.Suite
	array TTypeSequence
}

// SetupSuite is setting up array for all the tests
func (s *TemplateTestSuite) SetupSuite() {
	s.array = TTypeSequence{&TType{1}, &TType{2}, &TType{5}, &TType{3}, &TType{4}}

}

// TestMap is testing functionality of Map function
func (s *TemplateTestSuite) TestMap() {

	tmp := s.array.Map(func(num *TType) *TType {
		return &TType{num.code * 3}
	})
	for i, element := range tmp {
		assert.Equal(s.T(), MapResultArray1[i], element)
	}

	tmp = s.array.Map(func(num *TType) *TType {
		return &TType{num.code * num.code}
	})
	for i, element := range tmp {
		assert.Equal(s.T(), MapResultArray2[i], element)
	}
}

// TestFilter is testing functionality of Filter function
func (s *TemplateTestSuite) TestFilter() {

	tmp := s.array.Filter(func(num *TType) bool {
		if num.code%2 == 0 {
			return true
		}
		return false
	})
	for i, element := range tmp {
		assert.Equal(s.T(), FilterResultArray1[i], element)
	}

	tmp = s.array.Filter(func(num *TType) bool {
		if num.code%2 != 0 {
			return true
		}
		return false
	})
	for i, element := range tmp {
		assert.Equal(s.T(), FilterResultArray2[i], element)
	}
}

// TestReduce is testing functionality of Reduce function
func (s *TemplateTestSuite) TestReduce() {

	tmp := s.array.Reduce(func(a, b *TType) *TType {
		return &TType{a.code * b.code}
	}, &TType{1})
	assert.Equal(s.T(), &TType{120}, tmp)
}

// TestReduce is testing functionality of FindOne fuction
func (s *TemplateTestSuite) TestFindOne() {

	tmp := s.array.FindOne(func(num *TType) bool {
		if num.code%2 == 0 {
			return true
		}
		return false
	})
	assert.Equal(s.T(), &TType{2}, tmp)

	tmp = s.array.FindOne(func(num *TType) bool {
		if num.code%2 == 0 {
			return true
		}
		return false
	})
	assert.NotEqual(s.T(), FilterResultArray1, tmp)
}

// TestReverse is testing Reversing function
func (s *TemplateTestSuite) TestReverse() {

	tmp := s.array.Reverse()

	assert.Equal(s.T(), ReverseResultArray, tmp)
}

// TestSort is testing functionality of Sort function
func (s *TemplateTestSuite) TestSort() {

	tmp := s.array.Sort(func(a, b *TType) bool {
		if a.code > b.code {
			return true
		}
		return false
	})
	assert.Equal(s.T(), SortedResultArray, tmp)
}

// TestReverse is testing Reversing function
func (s *TemplateTestSuite) TestPush() {

	tmp := s.array.Push(&TType{100})
	assert.Equal(s.T(), PushResultArray, tmp)

}

// TestPop is testing Pop function
func (s *TemplateTestSuite) TestPop() {

	tmp := s.array.Pop()
	assert.Equal(s.T(), PopResultArray, tmp)

}

// TestReverse is testing Reversing function
func (s *TemplateTestSuite) TestForEach() {
	sum := 0
	s.array.ForEach(func(num *TType) {
		sum += num.code

	})
	assert.Equal(s.T(), 15, sum)

}

// Starting the test
func TestTemplateTestSuite(t *testing.T) {
	suite.Run(t, &TemplateTestSuite{})
}
