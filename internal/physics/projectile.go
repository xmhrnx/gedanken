package physics

import "math"

// Projectile is a launch on flat ground with no air resistance.
type Projectile struct {
	Speed float64 // m/s
	Angle float64 // degrees above horizontal
	G     float64 // m/s²
}

// components returns the horizontal and vertical launch speeds.
func (p Projectile) components() (vx, vy float64) {
	rad := p.Angle * math.Pi / 180
	return p.Speed * math.Cos(rad), p.Speed * math.Sin(rad)
}

// FlightTime is how long until it lands: t = 2v·sinθ/g.
func (p Projectile) FlightTime() float64 {
	_, vy := p.components()
	return 2 * vy / p.G
}

// Range is the horizontal distance travelled: R = v²·sin2θ/g.
func (p Projectile) Range() float64 {
	vx, _ := p.components()
	return vx * p.FlightTime()
}

// MaxHeight is the peak height: H = (v·sinθ)²/2g.
func (p Projectile) MaxHeight() float64 {
	_, vy := p.components()
	return vy * vy / (2 * p.G)
}

// At returns the position after t seconds.
func (p Projectile) At(t float64) (x, y float64) {
	vx, vy := p.components()
	return vx * t, vy*t - 0.5*p.G*t*t
}
