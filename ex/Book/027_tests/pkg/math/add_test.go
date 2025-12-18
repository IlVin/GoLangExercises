package math

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteMath struct {
	suite.Suite
}

func (s *SuiteMath) SetupTest() {
}

func (s *SuiteMath) TestAddPositive() {
	sum, err := Add(1, 2)
	s.Assert().Equal(err, nil, "Err is nil")
	s.Assert().Equal(sum, 3, "Sum eq 3")
}

func (s *SuiteMath) TestAddNegative() {
	s.Run("Add negative and positive", func() {
		_, err := Add(-1, 2)
		s.Assert().NotEqual(err, nil)
	})

	s.Run("Add positive and negative", func() {
		_, err := Add(1, -2)
		s.Assert().NotEqual(err, nil)
	})

	s.Run("Add negative and negative", func() {
		_, err := Add(-1, -2)
		s.Assert().NotEqual(err, nil)
	})

	s.Run("Add zero and zero", func() {
		_, err := Add(0, 0)
		s.Assert().NotEqual(err, nil)
	})

	s.Run("Add positive and zero", func() {
		_, err := Add(1, 0)
		s.Assert().NotEqual(err, nil)
	})

	s.Run("Add zero and positive", func() {
		_, err := Add(0, 2)
		s.Assert().NotEqual(err, nil)
	})
}

func (s *SuiteMath) TestEstimateValue() {
	s.Run("lt 10", func() {
		res := EstimateValue(9)
		s.Assert().Equal(res, "small")
	})

	s.Run("eq 10", func() {
		res := EstimateValue(10)
		s.Assert().Equal(res, "medium")
	})

	s.Run("gt 10 and lt 100", func() {
		res := EstimateValue(11)
		s.Assert().Equal(res, "medium")
	})

	s.Run("eq 100", func() {
		res := EstimateValue(100)
		s.Assert().Equal(res, "big")
	})

	s.Run("gt 100", func() {
		res := EstimateValue(101)
		s.Assert().Equal(res, "big")
	})
}

var tests = []struct {
	name string
	arg  int
	want string
}{
	{
		name: "-10",
		arg:  -10,
		want: "small",
	},
	{
		name: "0",
		arg:  0,
		want: "small",
	},
	{
		name: "9",
		arg:  9,
		want: "small",
	},
	{
		name: "10",
		arg:  10,
		want: "medium",
	},
	{
		name: "11",
		arg:  11,
		want: "medium",
	},
	{
		name: "99",
		arg:  99,
		want: "medium",
	},
	{
		name: "100",
		arg:  100,
		want: "big",
	},
	{
		name: "101",
		arg:  101,
		want: "big",
	},
	{
		name: "1000",
		arg:  1000,
		want: "big",
	},
}

func (s *SuiteMath) TestEstimateValueTableDriven() {
	for _, tt := range tests {
		s.Run(tt.name, func() {
			res := EstimateValue(tt.arg)
			s.Assert().Equal(res, tt.want)
		})

	}
}

func TestSuiteMath(t *testing.T) {
	suite.Run(t, new(SuiteMath))
}
