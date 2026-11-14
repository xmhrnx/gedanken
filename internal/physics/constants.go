// Package physics computes the quantities the interactive models show.
// It has no UI dependencies.
package physics

// Physical constants in SI units.
const (
	G          = 6.674e-11 // gravitational constant, m³/(kg·s²)
	C          = 2.998e8   // speed of light, m/s
	SolarMass  = 1.989e30  // kg
	SunRadius  = 6.957e8   // m
	EarthG     = 9.81      // surface gravity, m/s²
	Neutrons   = 2.5       // average neutrons released per U-235 fission
	FissionHit = 0.5       // chance an absorbed neutron causes another fission
)
