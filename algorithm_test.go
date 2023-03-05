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

func TestRankByFitness(t *testing.T) {
	gogenetic := GoGenetic{
		Gene:            Gene{0, 1},
		Generations:     1,
		SolutionsNumber: 4,
		SolutionLength:  10,
		Crossover:       OnePoint{},
		Fitness:         Fitness,
	}
	solutions := gogenetic.randomSample()
	scoreMap := rankByFitness(solutions, gogenetic.Fitness)
	if len(scoreMap) != 4 {
		t.Errorf("Could not rank by fitness.")
	}
	t.Log(scoreMap)
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
	scoreMap := rankByFitness(solutions, gogenetic.Fitness)
	pairs := makePairs(scoreMap)
	t.Log(pairs)
	if len(pairs) != len(scoreMap)/2 {
		t.Errorf("There is no right ammount of pairs")
	}
}
