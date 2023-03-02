package gogenetic

import "testing"

func (gogenetic GoGenetic) Fitness(solution []bool) int64 {
	return 5
}

func (gogenetic GoGenetic) Compare(score int64, otherScore int64) int64 {
	return 5
}

func TestGenetic(t *testing.T) {
	gogenetic := GoGenetic{}
	var test_probe interface{} = gogenetic
	_, ok := test_probe.(GeneticAlgorithm)
	if ok != true {
		t.Errorf("GoGenetic does not implement GeneticAlgorithmTrait")
	}
}

func TestSolutionGeneration(t *testing.T) {
	gogenetic := GoGenetic{
		Genome:          Genome{0, 1},
		Generations:     1,
		SolutionsNumber: 4,
		SolutionLength:  10,
		CrossoverType:   OnePoint{},
	}

	gogenetic.randomSample()
	solutions := gogenetic.solutions
	if len(solutions) != 4 {
		t.Errorf("randomSample does not generate samples.")
	}
	t.Log(solutions)
}
