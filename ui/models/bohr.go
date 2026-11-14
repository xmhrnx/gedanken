package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Bohr draws hydrogen's energy levels. Picking a jump shows the photon
// it emits: its energy is the gap, its wavelength hc/ΔE.
type Bohr struct {
	upper, lower *SliderRow
}

const bohrLevels = 6

// NewBohr creates the model.
func NewBohr() *Bohr {
	return &Bohr{
		upper: NewSliderRow("From level", 0.2, func(v float32) string { return fmt.Sprintf("n = %d", bohrUpper(v)) }),
		lower: NewSliderRow("To level", 0.25, func(v float32) string { return fmt.Sprintf("n = %d", bohrLower(v)) }),
	}
}

func bohrUpper(v float32) int { return 2 + int(math.Round(float64(v)*(bohrLevels-2))) }
func bohrLower(v float32) int { return 1 + int(math.Round(float64(v)*(bohrLevels-2))) }

// jump returns the chosen levels with the lower one below the upper.
func (m *Bohr) jump() (int, int) {
	u, l := bohrUpper(m.upper.Value()), bohrLower(m.lower.Value())
	return u, min(l, u-1)
}

// Layout draws the levels, readout and sliders.
func (m *Bohr) Layout(gtx giode.Context) giode.Dimensions {
	u, l := m.jump()
	nm := physics.PhotonWavelength(u, l)
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(250, func(p *giode.Painter) { m.draw(p, u, l, nm) }),
		readout(fmt.Sprintf("n=%d → %d: ΔE = %.2f eV · λ = %.0f nm (%s)", u, l,
			physics.HydrogenLevel(u)-physics.HydrogenLevel(l), nm, physics.SpectrumBand(nm))),
		m.upper, m.lower,
	).Layout(gtx)
}

func (m *Bohr) draw(p *giode.Painter, u, l int, nm float64) {
	pl := newPlot(p, 18, 0, 1, -14.2, 0.8)
	x0, x1 := pl.X(0.08), pl.X(0.62)
	for n := 1; n <= bohrLevels; n++ {
		y := pl.Y(physics.HydrogenLevel(n))
		c := giode.Color("#3a4a75")
		if n == u || n == l {
			c = theme.Text
		}
		p.Line(x0, y, x1, y, 1.5, c)
		if n <= 3 {
			label(p, x1+6, y, fmt.Sprintf("n=%d  %.2f eV", n, physics.HydrogenLevel(n)), giode.TextAnchorLeft, theme.Muted)
		}
	}
	dashed(p, x0, pl.Y(0), x1, pl.Y(0), 1, "#26324f")
	label(p, x1+6, pl.Y(0), "0 eV (n → ∞): electron free", giode.TextAnchorLeft, theme.Muted)
	ax := pl.X(0.3)
	yu, yl := pl.Y(physics.HydrogenLevel(u)), pl.Y(physics.HydrogenLevel(l))
	arrow(p, ax, yu, ax, yl, 2, theme.Accent)
	color := giode.Color(physics.WavelengthColor(nm))
	wiggle := p.Path()
	mid := (yu + yl) / 2
	for i := 0; i <= 40; i++ {
		x := ax + 8 + float32(i)*2
		y := mid + 6*float32(math.Sin(float64(i)*0.8))
		if i == 0 {
			wiggle.MoveTo(x, y)
		} else {
			wiggle.LineTo(x, y)
		}
	}
	wiggle.Stroke(2, color)
	label(p, ax+92, mid, "photon", giode.TextAnchorLeft, color)
}
