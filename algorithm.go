package gogenetic

import (
	"math/rand"
	"sort"
)

// GoGenetic contains all the variables and information needed for computations.
type GoGenetic struct {
	Gene            Gene
	Generations     int
	SolutionsNumber int
	SolutionLength  int
	ParentsLeft     int
	// Should be passed as 0 <= Mutation <= 1
	Mutation  float32
	Crossover Crossover
	Fitness   func(Solution) int
}

// Method running computing for given parameters.
// Returns best found solution or error.
func (gogenetic *GoGenetic) Run() (Solution, error) {

	samples := gogenetic.randomSample()
	for i := 0; i < gogenetic.Generations; i++ {
		sort.SliceStable(samples, func(i, j int) bool {
			return gogenetic.Fitness(samples[i]) > gogenetic.Fitness(samples[j]) //Change it to compare function
		})

		pairs := makePairs(samples)
		chann := make(chan Solution, len(pairs)*2)
		for _, pair := range pairs {
			go gogenetic.Crossover.Exchange(chann, pair[0], pair[1])
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
			return gogenetic.Fitness(samples[i]) > gogenetic.Fitness(samples[j]) //Change it to compare function
		})

		samples = append(samples[0:gogenetic.ParentsLeft],
			children[0:gogenetic.SolutionsNumber-gogenetic.ParentsLeft]...)

		// THIS LINE IS DEBUG ONLY
		/*
			if len(samples) != gogenetic.SolutionsNumber {
				println(len(samples))
				panic("wrong indexes")
			}
		*/
	}
	sort.SliceStable(samples, func(i, j int) bool {
		return gogenetic.Fitness(samples[i]) > gogenetic.Fitness(samples[j]) //Change it to compare function
	})
	return samples[0], nil
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
