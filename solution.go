package gogenetic

// Type representing one chromosome.
type Solution struct {
	Genes []int
}

func (s Solution) CrossoverExchange(other Solution, crossover Crossover) (Solution, Solution) {
	s1, s2 := crossover.Exchange(s, other)
	return s1, s2
}
