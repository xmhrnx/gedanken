package physics

import (
	"math"
	"math/rand"
)

// NeutronTrack is one neutron's straight path inside a sphere, in units
// of the mean free path, projected onto a plane through the center.
type NeutronTrack struct {
	X0, Y0, X1, Y1 float64
	Escaped        bool
}

// CriticalSample is a Monte Carlo sample of neutrons born uniformly in
// a sphere of the given radius (in mean free paths) and flying in
// random directions for an exponentially distributed distance.
type CriticalSample struct {
	Radius     float64
	Tracks     []NeutronTrack
	EscapeFrac float64
}

// SampleSphere runs n neutrons; tracks keeps the first few for drawing.
func SampleSphere(radius float64, n, tracks int, seed int64) CriticalSample {
	rng := rand.New(rand.NewSource(seed))
	s := CriticalSample{Radius: radius}
	escaped := 0
	for i := 0; i < n; i++ {
		t := neutron(rng, radius)
		if t.Escaped {
			escaped++
		}
		if i < tracks {
			s.Tracks = append(s.Tracks, t)
		}
	}
	s.EscapeFrac = float64(escaped) / float64(n)
	return s
}

// K is the multiplication factor: neutrons per fission that go on to
// cause another fission.
func (s CriticalSample) K() float64 {
	return Neutrons * FissionHit * (1 - s.EscapeFrac)
}

// neutron flies one neutron from a random point in the sphere.
func neutron(rng *rand.Rand, radius float64) NeutronTrack {
	var x, y, z float64
	for {
		x, y, z = 2*rng.Float64()-1, 2*rng.Float64()-1, 2*rng.Float64()-1
		if x*x+y*y+z*z <= 1 {
			break
		}
	}
	x, y, z = x*radius, y*radius, z*radius
	cosT := 2*rng.Float64() - 1
	sinT := math.Sqrt(1 - cosT*cosT)
	phi := 2 * math.Pi * rng.Float64()
	dx, dy, dz := sinT*math.Cos(phi), sinT*math.Sin(phi), cosT
	d := rng.ExpFloat64()
	ex, ey, ez := x+dx*d, y+dy*d, z+dz*d
	return NeutronTrack{X0: x, Y0: y, X1: ex, Y1: ey, Escaped: ex*ex+ey*ey+ez*ez > radius*radius}
}
