package gogenetic

// Crossover interface indicates how the genome exchange should be computed.
type Crossover interface {
	Exchange(s1 Solution, s2 Solution) (Solution, Solution)
}

type OnePoint struct{}

// Method doing a crossover exchange in the middle of the solution (for now).
func (c OnePoint) Exchange(s1 Solution, s2 Solution) (Solution, Solution) {
	for i := 0; i < len(s1.Genomes); i++ { // ! only 50 percent
		s1.Genomes[i], s2.Genomes[i] = s2.Genomes[i], s1.Genomes[i]
	}
	return s1, s2
}
