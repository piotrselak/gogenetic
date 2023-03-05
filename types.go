package gogenetic

// Type representing one chromosome.
type Solution struct {
	Genes []int
}

type FitnessScore struct {
	Score int
	Value Solution
}
