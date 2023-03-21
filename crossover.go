package gogenetic

import "math/rand"

// Crossover interface indicates how the Gene exchange should be computed.
type Crossover interface {
	Exchange(channel chan Solution, s1 Solution, s2 Solution)
}

type OnePoint struct{}

// Starting point for crossover is being randomised.
func (c OnePoint) Exchange(channel chan Solution, s1 Solution, s2 Solution) {
	startPoint := rand.Intn(len(s1.Genes))
	for i := startPoint; i < len(s1.Genes); i++ {
		s1.Genes[i], s2.Genes[i] = s2.Genes[i], s1.Genes[i]
	}
	channel <- s1
	channel <- s2
}

type KPoint struct {
	K int
}

func (c KPoint) Exchange(channel chan Solution, s1, s2 Solution) {

}
