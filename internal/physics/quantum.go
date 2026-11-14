package physics

import "math"

// HydrogenLevel is the Bohr energy of level n in eV: −13.6/n².
func HydrogenLevel(n int) float64 {
	return -13.6 / float64(n*n)
}

// PhotonWavelength is the wavelength in nm of the photon emitted when
// hydrogen drops from level upper to lower: λ = hc/ΔE, hc ≈ 1240 eV·nm.
func PhotonWavelength(upper, lower int) float64 {
	return 1239.84 / (HydrogenLevel(upper) - HydrogenLevel(lower))
}

// SpectrumBand names the part of the spectrum a wavelength in nm is in.
func SpectrumBand(nm float64) string {
	switch {
	case nm < 380:
		return "ultraviolet"
	case nm > 750:
		return "infrared"
	default:
		return "visible"
	}
}

// WavelengthColor approximates the colour of visible light as "#rrggbb";
// outside the visible band it returns a pale grey.
func WavelengthColor(nm float64) string {
	var r, g, b float64
	switch {
	case nm < 380 || nm > 750:
		return "#9aa3b8"
	case nm < 440:
		r, b = (440-nm)/60, 1
	case nm < 490:
		g, b = (nm-440)/50, 1
	case nm < 510:
		g, b = 1, (510-nm)/20
	case nm < 580:
		r, g = (nm-510)/70, 1
	case nm < 645:
		r, g = 1, (645-nm)/65
	default:
		r = 1
	}
	c := func(v float64) int { return int(255 * math.Pow(v, 0.8)) }
	return "#" + hex2(c(r)) + hex2(c(g)) + hex2(c(b))
}

func hex2(v int) string {
	const digits = "0123456789abcdef"
	return string([]byte{digits[v/16], digits[v%16]})
}

// BoxWave is the particle-in-a-box wavefunction √2·sin(nπx) for x in [0, 1].
func BoxWave(n int, x float64) float64 {
	return math.Sqrt2 * math.Sin(float64(n)*math.Pi*x)
}

// Gamma is the Lorentz factor 1/√(1 − β²) for speed β = v/c.
func Gamma(beta float64) float64 {
	return 1 / math.Sqrt(1-beta*beta)
}
