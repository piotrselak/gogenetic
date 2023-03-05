package gogenetic

import "testing"

func Fitness(solution Solution) int {
	return 5
}

func TestSolutionGeneration(t *testing.T) {
	gogenetic := GoGenetic{
		Gene:            Gene{0, 1},
		Generations:     1,
		SolutionsNumber: 4,
		SolutionLength:  10,
		Crossover:       OnePoint{},
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
		Crossover:       OnePoint{},
		Fitness:         Fitness,
	}
	gogenetic.Run()
}

func TestMakePairs(t *testing.T) {
	gogenetic := GoGenetic{
		Gene:            Gene{0, 1},
		Generations:     1,
		SolutionsNumber: 4,
		SolutionLength:  10,
		Crossover:       OnePoint{},
		Fitness:         Fitness,
	}
	solutions := gogenetic.randomSample()
	pairs := makePairs(solutions)
	t.Log(pairs)
	if len(pairs) != len(solutions)/2 {
		t.Errorf("There is no right ammount of pairs")
	}
}
