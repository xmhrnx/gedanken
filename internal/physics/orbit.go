package physics

import "math"

// OrbitOutcome is what happens to a cannonball fired horizontally.
type OrbitOutcome int

// Outcomes of Newton's cannon.
const (
	Crashes OrbitOutcome = iota
	Orbits
	Escapes
)

// Point is a position in the plane.
type Point struct{ X, Y float64 }

// Cannon fires a ball horizontally from a mountain at height h above a
// planet of radius 1 with GM = 1. Speed is in units of the circular
// orbit speed at the launch height.
type Cannon struct {
	Height float64
	Speed  float64
}

// Path integrates the flight with small steps and returns the points
// until the ball lands, completes an orbit or leaves the view.
func (c Cannon) Path() ([]Point, OrbitOutcome) {
	r0 := 1 + c.Height
	v := c.Speed * math.Sqrt(1/r0)
	x, y, vx, vy := 0.0, r0, v, 0.0
	const dt, maxSteps, viewRadius = 0.002, 40000, 4.0
	pts := []Point{{x, y}}
	angle := 0.0
	for i := 0; i < maxSteps; i++ {
		r := math.Hypot(x, y)
		ax, ay := -x/(r*r*r), -y/(r*r*r)
		vx += ax * dt
		vy += ay * dt
		px, py := x, y
		x += vx * dt
		y += vy * dt
		angle += math.Abs(math.Atan2(px*y-py*x, px*x+py*y))
		if i%10 == 0 {
			pts = append(pts, Point{x, y})
		}
		switch r = math.Hypot(x, y); {
		case r < 1:
			return append(pts, Point{x, y}), Crashes
		case r > viewRadius:
			return pts, Escapes
		case angle >= 2*math.Pi:
			return append(pts, Point{x, y}), Orbits
		}
	}
	return pts, Orbits
}

// EscapeSpeedRatio is escape speed over circular speed: √2.
var EscapeSpeedRatio = math.Sqrt2
