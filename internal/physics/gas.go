package physics

import (
	"math"
	"math/rand"
)

// Molecule is a particle in the unit box.
type Molecule struct{ X, Y, VX, VY float64 }

// Gas is a 2D gas of point molecules in a unit box. Speeds scale with
// √T; Pressure is measured from the momentum delivered to the walls.
type Gas struct {
	Molecules []Molecule
	impulse   float64
	elapsed   float64
}

// NewGas places n molecules with random directions and unit speed.
func NewGas(n int, seed int64) *Gas {
	rng := rand.New(rand.NewSource(seed))
	g := &Gas{}
	for i := 0; i < n; i++ {
		a := 2 * math.Pi * rng.Float64()
		speed := 0.5 + rng.Float64()
		g.Molecules = append(g.Molecules, Molecule{rng.Float64(), rng.Float64(), speed * math.Cos(a), speed * math.Sin(a)})
	}
	return g
}

// Step advances by dt at temperature scale t (1 = reference).
func (g *Gas) Step(dt, t float64) {
	s := math.Sqrt(t)
	for i := range g.Molecules {
		m := &g.Molecules[i]
		m.X += m.VX * s * dt
		m.Y += m.VY * s * dt
		if m.X < 0 || m.X > 1 {
			m.VX = -m.VX
			m.X = math.Max(0, math.Min(1, m.X))
			g.impulse += 2 * math.Abs(m.VX) * s
		}
		if m.Y < 0 || m.Y > 1 {
			m.VY = -m.VY
			m.Y = math.Max(0, math.Min(1, m.Y))
			g.impulse += 2 * math.Abs(m.VY) * s
		}
	}
	g.elapsed += dt
}

// Pressure returns the average wall impulse per unit time since the
// last call, and starts a new measurement.
func (g *Gas) Pressure() float64 {
	if g.elapsed == 0 {
		return 0
	}
	p := g.impulse / g.elapsed / 4
	g.impulse, g.elapsed = 0, 0
	return p
}
