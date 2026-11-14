package physics

import "math/rand"

// Fission is one nucleus splitting in a chain reaction.
type Fission struct {
	Nucleus    int // index into the grid
	Parent     int // index of the nucleus whose neutron caused it; -1 for the first
	Generation int
}

// ChainReaction simulates generations of fissions on a grid of nuclei.
// Each fission causes about K new ones among nearby unsplit nuclei.
type ChainReaction struct {
	Cols, Rows int
	K          float64
	Fissions   []Fission
	split      []bool
	rng        *rand.Rand
}

// NewChainReaction starts a reaction at the center of the grid.
func NewChainReaction(cols, rows int, k float64, seed int64) *ChainReaction {
	c := &ChainReaction{Cols: cols, Rows: rows, K: k, split: make([]bool, cols*rows), rng: rand.New(rand.NewSource(seed))}
	start := rows/2*cols + cols/2
	c.split[start] = true
	c.Fissions = []Fission{{Nucleus: start, Parent: -1}}
	return c
}

// Generations returns how many generations have run, the first included.
func (c *ChainReaction) Generations() int {
	return c.Fissions[len(c.Fissions)-1].Generation + 1
}

// Step runs one more generation and reports how many fissions it had.
func (c *ChainReaction) Step() int {
	last := c.Generations() - 1
	var born []Fission
	for _, f := range c.Fissions {
		if f.Generation != last {
			continue
		}
		for i := 0; i < c.offspring(); i++ {
			if n, ok := c.nearbyUnsplit(f.Nucleus); ok {
				c.split[n] = true
				born = append(born, Fission{Nucleus: n, Parent: f.Nucleus, Generation: last + 1})
			}
		}
	}
	c.Fissions = append(c.Fissions, born...)
	return len(born)
}

// offspring draws a whole number of fissions with mean K.
func (c *ChainReaction) offspring() int {
	n := int(c.K)
	if c.rng.Float64() < c.K-float64(n) {
		n++
	}
	return n
}

// nearbyUnsplit picks a random unsplit nucleus within a few cells.
func (c *ChainReaction) nearbyUnsplit(from int) (int, bool) {
	const reach, tries = 3, 12
	x0, y0 := from%c.Cols, from/c.Cols
	for i := 0; i < tries; i++ {
		x := x0 + c.rng.Intn(2*reach+1) - reach
		y := y0 + c.rng.Intn(2*reach+1) - reach
		if x < 0 || y < 0 || x >= c.Cols || y >= c.Rows {
			continue
		}
		if n := y*c.Cols + x; !c.split[n] {
			return n, true
		}
	}
	return 0, false
}

// Split reports whether nucleus n has split.
func (c *ChainReaction) Split(n int) bool { return c.split[n] }
