package physics

import "math"

// EscapeVelocity is √(2GM/r): the launch speed that never falls back.
func EscapeVelocity(mass, radius float64) float64 {
	return math.Sqrt(2 * G * mass / radius)
}

// SchwarzschildRadius is 2GM/c², the horizon radius of a mass.
func SchwarzschildRadius(mass float64) float64 {
	return 2 * G * mass / (C * C)
}

// FlammDepth is the depth of Flamm's paraboloid, the embedding surface
// that pictures the curved space around a mass: z = 2√(rs(r − rs)).
// It is zero at the horizon and r below rs has no embedding.
func FlammDepth(r, rs float64) float64 {
	if r <= rs {
		return 0
	}
	return 2 * math.Sqrt(rs*(r-rs))
}

// OutgoingLightSlope is dt/dr of an outgoing light ray in
// Eddington–Finkelstein coordinates (c = 1): (1 + rs/r)/(1 − rs/r).
// Positive outside the horizon, infinite at it, negative inside: there
// "outgoing" light also moves inward. Ingoing light always has slope −1.
func OutgoingLightSlope(r, rs float64) float64 {
	if r == rs {
		return math.Inf(1)
	}
	return (1 + rs/r) / (1 - rs/r)
}

// TimeDilation is how fast a clock at r ticks relative to one far
// away: √(1 − rs/r). Zero at and inside the horizon.
func TimeDilation(r, rs float64) float64 {
	if r <= rs {
		return 0
	}
	return math.Sqrt(1 - rs/r)
}
