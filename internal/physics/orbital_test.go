package physics

import (
	"math"
	"testing"
)

func TestOneSMeanRadius(t *testing.T) {
	pts := SampleOrbital(Orbitals[0], 20000, 1)
	sum := 0.0
	for _, p := range pts {
		sum += math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
	}
	if mean := sum / float64(len(pts)); math.Abs(mean-1.5) > 0.05 {
		t.Errorf("1s mean radius %.3f, want 1.5 Bohr radii", mean)
	}
}

func TestTwoPHasNodalPlane(t *testing.T) {
	near, up := 0, 0
	for _, p := range SampleOrbital(Orbitals[2], 20000, 1) {
		if math.Abs(p.Z) < 0.2 {
			near++
		}
		if p.Positive == (p.Z > 0) {
			up++
		}
	}
	if near > 400 || up != 20000 {
		t.Errorf("2p: %d points near z=0, %d with sign matching z", near, up)
	}
}

func TestThreeSHasTwoSphericalNodes(t *testing.T) {
	// 27 - 18r + 2r² = 0 at r = (9 ± 3√3)/2: about 1.90 and 7.10 Bohr radii.
	var threeS Orbital
	for _, o := range Orbitals {
		if o.Name == "3s" {
			threeS = o
		}
	}
	near := 0
	for _, p := range SampleOrbital(threeS, 20000, 1) {
		r := math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
		if math.Abs(r-1.902) < 0.05 || math.Abs(r-7.098) < 0.05 {
			near++
		}
	}
	if near > 60 {
		t.Errorf("%d points sit on the 3s nodes", near)
	}
}
