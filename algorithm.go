package gogenetic

import (
	"errors"
	"fmt"
	"math/rand"
)

// GeneticAlgorithm is an interface which has to be implemented for GoGeneric struct
// as these methods' implementation depends on the algorithm
type GeneticAlgorithm interface {
	Fitness(solution Solution) int
	Compare(score int, otherScore int) int
}

// GoGenetic contains all the variables and information needed for computations.
// All methods from GeneticAlgorithm interface must be implemented in order to
// work as intended.
type GoGenetic struct {
	Gene            Gene
	Generations     int
	SolutionsNumber int
	SolutionLength  int
	CrossoverType   Crossover
}

// Method running computing for given parameters.
// Returns best found solution or error.
func (gogenetic *GoGenetic) Run() (Solution, error) {
	var algorithmForCheck interface{} = gogenetic
	_, ok := algorithmForCheck.(GeneticAlgorithm)

	if ok == false {
		return Solution{}, errors.New("Given GoGenetic struct does not implement GeneticAlgorithm interface.")
	}

	samples := gogenetic.randomSample()
	for i := 0; i < gogenetic.Generations; i++ {
		fmt.Println(samples)
	}
	return Solution{}, nil
}

// Method generating random sample of gens used to create first generation of solutions.
func (gogenetic GoGenetic) randomSample() []Solution {
	var solutions []Solution
	for i := 0; i < gogenetic.SolutionsNumber; i++ {
		solution := Solution{}
		for i := 0; i < gogenetic.SolutionLength; i++ {
			randomIndex := rand.Intn(len(gogenetic.Gene))
			solution = Solution{Genes: append(solution.Genes, gogenetic.Gene[randomIndex])}
		}
		solutions = append(solutions, solution)
	}
	return solutions
}

func rankByFitness(samples []Solution, f func(Solution) int) []FitnessScore {
	var scoreArr []FitnessScore
	for _, sample := range samples {
		score := f(sample)
		scoreArr = append(scoreArr, FitnessScore{
			Score: score,
			Value: sample,
		})
	}
	return scoreArr
}
