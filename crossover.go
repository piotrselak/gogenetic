package gogenetic

// Crossover interface indicates how the Gene exchange should be computed.
type Crossover interface {
	Exchange(s1 Solution, s2 Solution) (Solution, Solution)
}

type OnePoint struct{}

// Method doing a crossover exchange in the middle of the solution (for now).
func (c OnePoint) Exchange(s1 Solution, s2 Solution) (Solution, Solution) {
	for i := len(s1.Genes) / 2; i < len(s1.Genes); i++ { // ! only 50 percent
		s1.Genes[i], s2.Genes[i] = s2.Genes[i], s1.Genes[i]
	}
	return s1, s2
}
