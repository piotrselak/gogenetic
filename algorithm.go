package gogenetic

import (
	"errors"
	"math/rand"
)

// GeneticAlgorithm is an interface which has to be implemented for GoGeneric struct
// as these methods' implementation depends on the algorithm
type GeneticAlgorithm interface {
	Fitness(solution []int64) int64
	Compare(score int64, otherScore int64) int64
}

// GoGenetic contains all the variables and information needed for computations.
// All methods from GeneticAlgorithm interface must be implemented in order to
// work as intended.
type GoGenetic struct {
	Genome          Genome
	Generations     int
	SolutionsNumber int
	SolutionLength  int
	CrossoverType   Crossover
	solutions       []Solution
}

// Method running computing for given parameters.
// Returns best found solution or error.
func (gogenetic GoGenetic) Run() (Solution, error) {
	var algorithmForCheck interface{} = gogenetic
	algorithm, ok := algorithmForCheck.(GeneticAlgorithm)

	if ok == false {
		return Solution{}, errors.New("Given GoGenetic struct does not implement GeneticAlgorithm interface.")
	}

	return Solution{}, nil
}

// Method generating random sample of gens used to create first generation of solutions.
func (gogenetic *GoGenetic) randomSample() {
	var solutions []Solution
	for i := 0; i < gogenetic.SolutionsNumber; i++ {
		solution := Solution{}
		for i := 0; i < gogenetic.SolutionLength; i++ {
			randomIndex := rand.Intn(len(gogenetic.Genome))
			solution = Solution{Genomes: append(solution.Genomes, gogenetic.Genome[randomIndex])}
		}
		solutions = append(solutions, solution)
	}
	gogenetic.solutions = solutions
}
