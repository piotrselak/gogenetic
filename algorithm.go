package gogenetic

import (
	"math/rand"
	"sort"
)

// GoGenetic contains all the variables and information needed for computations.
type GoGenetic struct {
	gene            Gene
	generations     int
	solutionsNumber int
	solutionLength  int
	parentsLeft     int
	stopAt          int
	// Should be passed as 0 <= Mutation <= 1
	// Really small numbers are encouraged
	mutation  float32
	crossover Crossover
	fitness   func(Solution) int
}

// Method for creating GoGenetic object
func NewGoGenetic(gene Gene,
	generations int,
	solutionsNumber int,
	solutionLength int,
	parentsLeft int,
	stopAt int,
	mutation float32,
	crossover Crossover,
	fitness func(Solution) int) GoGenetic {
	return GoGenetic{
		gene:            gene,
		generations:     generations,
		solutionsNumber: solutionsNumber,
		solutionLength:  solutionLength,
		parentsLeft:     parentsLeft,
		stopAt:          stopAt,
		mutation:        mutation,
		crossover:       crossover,
		fitness:         fitness,
	}
}

// Method running computing for given parameters.
// Returns best found solution or error.
func (gogenetic *GoGenetic) Run() Solution {

	samples := gogenetic.randomSample()
	for i := 0; i < gogenetic.generations; i++ {
		sort.SliceStable(samples, func(i, j int) bool {
			return gogenetic.fitness(samples[i]) > gogenetic.fitness(samples[j]) //Change it to compare function
		})

		if gogenetic.fitness(samples[0]) >= gogenetic.stopAt {
			return samples[0]
		}

		pairs := makePairs(samples)
		chann := make(chan Solution, len(pairs)*2)
		for _, pair := range pairs {
			go gogenetic.crossover.Exchange(chann, pair[0], pair[1])
		}

		var children []Solution
		for i := 0; i < cap(chann); i++ {
			val, ok := <-chann
			if ok {
				children = append(children, val)
			}
		}
		close(chann)

		sort.SliceStable(children, func(i, j int) bool {
			return gogenetic.fitness(samples[i]) > gogenetic.fitness(samples[j]) //Change it to compare function
		})

		samples = append(samples[0:gogenetic.parentsLeft],
			children[0:gogenetic.solutionsNumber-gogenetic.parentsLeft]...)

		for i := 0; i < len(samples); i++ {
			samples[i].Mutate(gogenetic.mutation, gogenetic.gene)
		}
		// THIS LINE IS DEBUG ONLY SHOULD BE MOVED TO TEST
		/*
			if len(samples) != gogenetic.SolutionsNumber {
				println(len(samples))
				panic("wrong indexes")
			}
		*/
	}

	sort.SliceStable(samples, func(i, j int) bool {
		return gogenetic.fitness(samples[i]) > gogenetic.fitness(samples[j]) //Change it to compare function
	})
	return samples[0]
}

// Method generating random sample of gens used to create first generation of solutions.
func (gogenetic GoGenetic) randomSample() []Solution {
	var solutions []Solution
	for i := 0; i < gogenetic.solutionsNumber; i++ {
		solution := Solution{}
		for i := 0; i < gogenetic.solutionLength; i++ {
			randomIndex := rand.Intn(len(gogenetic.gene))
			solution = Solution{Genes: append(solution.Genes, gogenetic.gene[randomIndex])}
		}
		solutions = append(solutions, solution)
	}
	return solutions
}

func makePairs(arr []Solution) [][]Solution {
	var paired [][]Solution
	numberOfPairs := len(arr) / 2 // ignores last element, maybe change later
	for i := 0; i < numberOfPairs; i++ {
		paired = append(paired, []Solution{})
	}
	pairCounter := 0
	for i, el := range arr {
		paired[pairCounter] = append(paired[pairCounter], el)
		if i%2 == 1 {
			pairCounter += 1
		}
	}
	return paired
}
