package gogenetic

import "testing"

func (gogenetic GoGenetic) Fitness(solution Solution) int {
	return 5
}

func (gogenetic GoGenetic) Compare(score int, otherScore int) int {
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
		Gene:            Gene{0, 1},
		Generations:     1,
		SolutionsNumber: 4,
		SolutionLength:  10,
		CrossoverType:   OnePoint{},
	}

	solutions := gogenetic.randomSample()
	if len(solutions) != 4 {
		t.Errorf("randomSample does not generate samples.")
	}
}

func TestRunningGoGenetic(t *testing.T) {
	gogenetic := GoGenetic{
		Gene:            Gene{0, 1},
		Generations:     1,
		SolutionsNumber: 4,
		SolutionLength:  10,
		CrossoverType:   OnePoint{},
	}
	gogenetic.Run()
}

func TestRankByFitness(t *testing.T) {
	gogenetic := GoGenetic{
		Gene:            Gene{0, 1},
		Generations:     1,
		SolutionsNumber: 4,
		SolutionLength:  10,
		CrossoverType:   OnePoint{},
	}
	solutions := gogenetic.randomSample()
	scoreMap := rankByFitness(solutions, gogenetic.Fitness)
	if len(scoreMap) != 4 {
		t.Errorf("Could not rank by fitness.")
	}
	t.Log(scoreMap)
}
