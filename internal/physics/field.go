package physics

import "math"

// Charge is a point charge in the plane.
type Charge struct {
	X, Y, Q float64
}

// FieldAt is the electric field (Coulomb, k = 1) of the charges at (x, y).
func FieldAt(x, y float64, charges []Charge) (float64, float64) {
	var ex, ey float64
	for _, c := range charges {
		dx, dy := x-c.X, y-c.Y
		r2 := dx*dx + dy*dy
		if r2 < 1e-9 {
			continue
		}
		f := c.Q / (r2 * math.Sqrt(r2))
		ex += f * dx
		ey += f * dy
	}
	return ex, ey
}

// FieldLine traces a line of force from (x, y) along the field (or
// against it when backward), stopping near a charge or outside bounds.
func FieldLine(x, y float64, charges []Charge, backward bool, bound float64) []Point {
	const step, maxSteps, stopNear = 0.01, 600, 0.03
	pts := []Point{{x, y}}
	for i := 0; i < maxSteps; i++ {
		ex, ey := FieldAt(x, y, charges)
		n := math.Hypot(ex, ey)
		if n == 0 {
			break
		}
		if backward {
			ex, ey = -ex, -ey
		}
		x += step * ex / n
		y += step * ey / n
		pts = append(pts, Point{x, y})
		if math.Abs(x) > bound || math.Abs(y) > bound {
			break
		}
		for _, c := range charges {
			if math.Hypot(x-c.X, y-c.Y) < stopNear {
				return pts
			}
		}
	}
	return pts
}
