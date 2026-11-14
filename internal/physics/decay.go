package physics

import (
	"math"
	"math/rand"
)

// DecaySample gives each of n nuclei a random decay time, measured in
// half-lives, drawn from the exponential distribution.
type DecaySample struct {
	Times []float64
}

// NewDecaySample draws n decay times.
func NewDecaySample(n int, seed int64) DecaySample {
	rng := rand.New(rand.NewSource(seed))
	s := DecaySample{Times: make([]float64, n)}
	for i := range s.Times {
		s.Times[i] = rng.ExpFloat64() / math.Ln2
	}
	return s
}

// Remaining counts nuclei not yet decayed after t half-lives.
func (s DecaySample) Remaining(t float64) int {
	n := 0
	for _, d := range s.Times {
		if d > t {
			n++
		}
	}
	return n
}
