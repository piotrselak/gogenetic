package main

import (
	"fmt"
	"math"

	"github.com/piotrselak/gogenetic"
)

var numbers []int = []int{1, 2, 3, 6, 10, 17, 25, 29, 30, 41, 51, 60, 70, 79, 80}

func Fitness(solution gogenetic.Solution) int {
	genes := solution.Genes
	a := 0
	b := 0
	for i, gene := range genes {
		if gene == 0 {
			a += numbers[i]
		} else {
			b += numbers[i]
		}
	}
	return int(math.Abs(float64(a)-float64(b))) * -1
}

func main() {
	genetics := gogenetic.GoGenetic{
		Gene:            []int{0, 1},
		Generations:     15,
		SolutionsNumber: 50,
		SolutionLength:  len(numbers),
		ParentsLeft:     3,
		Mutation:        0.00001,
		Crossover:       gogenetic.OnePoint{},
		Fitness:         Fitness,
	}
	solution := genetics.Run()

	result := Fitness(solution)
	fmt.Println(solution)
	fmt.Println(result)
}
