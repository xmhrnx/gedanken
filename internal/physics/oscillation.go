package physics

import "math"

// PendulumPeriod is 2π√(L/g) for small swings.
func PendulumPeriod(length, g float64) float64 {
	return 2 * math.Pi * math.Sqrt(length/g)
}

// PendulumAngle is the small-angle swing θ₀·cos(2πt/T).
func PendulumAngle(amplitude, length, g, t float64) float64 {
	return amplitude * math.Cos(2*math.Pi*t/PendulumPeriod(length, g))
}

// WaveHeight is a travelling sine wave A·sin(2π(x/λ − f·t)).
func WaveHeight(amplitude, wavelength, frequency, x, t float64) float64 {
	return amplitude * math.Sin(2*math.Pi*(x/wavelength-frequency*t))
}

// TwoSourceAmplitude is the summed amplitude at (x, y) of two equal
// sources at (±d/2, 0) emitting waves of wavelength λ, at time phase t
// (0–1 of a cycle). Bright where path lengths differ by whole λ.
func TwoSourceAmplitude(x, y, d, wavelength, t float64) float64 {
	r1 := math.Hypot(x+d/2, y)
	r2 := math.Hypot(x-d/2, y)
	w := 2 * math.Pi
	return math.Sin(w*(r1/wavelength-t)) + math.Sin(w*(r2/wavelength-t))
}

// FringeIntensity is the time-averaged brightness at angle θ for two
// slits a distance d apart: cos²(πd·sinθ/λ), from 0 to 1.
func FringeIntensity(theta, d, wavelength float64) float64 {
	c := math.Cos(math.Pi * d * math.Sin(theta) / wavelength)
	return c * c
}
