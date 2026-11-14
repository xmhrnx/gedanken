package physics

import (
	"math"
	"math/rand"
)

// Orbital is a hydrogen orbital, with its wavefunction ψ in units of
// the Bohr radius (normalisation does not matter for sampling).
type Orbital struct {
	Name   string
	Extent float64 // radius that holds nearly all of the probability
	Psi    func(r, cosTheta float64) float64
}

// Orbitals lists the hydrogen orbitals the home screen cycles through.
var Orbitals = []Orbital{
	{"1s", 5, func(r, _ float64) float64 { return math.Exp(-r) }},
	{"2s", 13, func(r, _ float64) float64 { return (2 - r) * math.Exp(-r/2) }},
	{"2p", 12, func(r, c float64) float64 { return r * c * math.Exp(-r/2) }},
	{"3s", 27, func(r, _ float64) float64 { return (27 - 18*r + 2*r*r) * math.Exp(-r/3) }},
	{"3p", 22, func(r, c float64) float64 { return r * (6 - r) * c * math.Exp(-r/3) }},
	{"3d", 20, func(r, c float64) float64 { return r * r * (3*c*c - 1) * math.Exp(-r/3) }},
}

// CloudPoint is one sampled electron position; Positive is the sign
// of ψ there, which distinguishes the lobes.
type CloudPoint struct {
	X, Y, Z  float64
	Positive bool
}

// SampleOrbital draws n positions distributed as |ψ|², by rejection
// sampling inside a cube of half-width Extent.
func SampleOrbital(o Orbital, n int, seed int64) []CloudPoint {
	rng := rand.New(rand.NewSource(seed))
	peak := orbitalPeak(o)
	pts := make([]CloudPoint, 0, n)
	for len(pts) < n {
		x := (2*rng.Float64() - 1) * o.Extent
		y := (2*rng.Float64() - 1) * o.Extent
		z := (2*rng.Float64() - 1) * o.Extent
		r := math.Sqrt(x*x + y*y + z*z)
		if r > o.Extent {
			continue
		}
		cos := 1.0
		if r > 0 {
			cos = z / r
		}
		psi := o.Psi(r, cos)
		if rng.Float64()*peak < psi*psi {
			pts = append(pts, CloudPoint{x, y, z, psi >= 0})
		}
	}
	return pts
}

// orbitalPeak finds the largest |ψ|² on a grid, the bound for sampling.
func orbitalPeak(o Orbital) float64 {
	peak := 0.0
	for i := 0; i <= 400; i++ {
		r := o.Extent * float64(i) / 400
		for _, c := range []float64{-1, -0.5, 0, 0.5, 1} {
			if p := o.Psi(r, c); p*p > peak {
				peak = p * p
			}
		}
	}
	return peak
}
