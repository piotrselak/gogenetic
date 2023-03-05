package gogenetic

import (
	"math/rand"
)

// Type representing one chromosome.
type Solution struct {
	Genes []int
}

func (s *Solution) Mutate(chance float32, pool Gene) {
	for i := 0; i < len(s.Genes); i++ {
		if random := rand.Float32(); random <= chance {
			ind := rand.Intn(len(pool))
			s.Genes[i] = pool[ind]
		}
	}
}
