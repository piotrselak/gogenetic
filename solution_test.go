package gogenetic

import (
	"testing"

	"github.com/piotrselak/gogenetic/internal/testutils"
)

func TestOnePointCrossover(t *testing.T) {
	solution1 := Solution{
		Genes: []int{0, 1, 1, 0},
	}
	solution2 := Solution{
		Genes: []int{1, 0, 0, 1},
	}
	s1, s2 := solution1.CrossoverExchange(solution2, OnePoint{})

	if !testutils.Equal(s1.Genes, []int{0, 1, 0, 1}) {
		t.Error("Crossover exchange did not work for solution1")
	}

	if !testutils.Equal(s2.Genes, []int{1, 0, 1, 0}) {
		t.Error("Crossover exchange did not work for solution2")
	}
}
