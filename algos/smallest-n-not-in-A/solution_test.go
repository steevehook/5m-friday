package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type solutionSuite struct {
	suite.Suite
	input []int
}

func (s solutionSuite) Test_Solution_Success() {
	in := []int{4,3,10,11,30,-2,1,4,4,5,6}

	res := Solution(in)

	s.Equal(2, res)
}

func (s solutionSuite) Test_Solution_OnlyZero() {
	in := []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

	res := Solution(in)

	s.Equal(1, res)
}

func (s solutionSuite) Test_Solution_AllInSlice() {
	in := []int{1,2,3,4,5,6,7,8}

	res := Solution(in)

	s.Equal(9, res)
}

func (s solutionSuite) Test_Solution_OnlyNegative() {
	in := []int{-1,-20,-11,-12,-3,-1,-3,-10}

	res := Solution(in)

	s.Equal(1, res)
}

func (s solutionSuite) Test_Solution_Duplicates() {
	in := []int{-1,-20,1,1,1,1,1,1,3,1,1,3,10,4,15,6}

	res := Solution(in)

	s.Equal(2, res)
}

func TestSolution(t *testing.T) {
	suite.Run(t, new(solutionSuite))
}

func BenchmarkSolution(b *testing.B) {
	b.StartTimer()
	Solution(bigSlice())
	b.StartTimer()
}

func bigSlice() []int {
	var s []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000000; i++ {
		n := r.Intn(1000000)
		s = append(s, n)
	}
	return s
}
