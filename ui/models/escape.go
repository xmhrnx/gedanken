package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Escape shrinks a star of one solar mass. Escape velocity √(2GM/R)
// rises as R falls; when it passes c the star goes dark.
type Escape struct {
	radius *SliderRow
}

// Radii covered by the slider, in metres, on a log scale.
const escapeMaxR, escapeMinR = physics.SunRadius, 1e3

// NewEscape creates the model.
func NewEscape() *Escape {
	return &Escape{radius: NewSliderRow("Shrink the star", 0, func(v float32) string {
		return formatKm(escapeRadius(v))
	})}
}

// escapeRadius maps the slider onto radii from the Sun's down to 1 km.
func escapeRadius(v float32) float64 {
	return escapeMaxR * math.Pow(escapeMinR/escapeMaxR, float64(v))
}

// drawnRadius maps a real radius to dp, logarithmically.
func drawnRadius(r float64) float32 {
	f := math.Log10(r/escapeMinR) / math.Log10(escapeMaxR/escapeMinR)
	return 6 + 74*float32(f)
}

// Layout draws the star, readout and slider.
func (e *Escape) Layout(gtx giode.Context) giode.Dimensions {
	r := escapeRadius(e.radius.Value())
	v := physics.EscapeVelocity(physics.SolarMass, r)
	text := fmt.Sprintf("escape velocity %s km/s · %.3g c", groupDigits(v/1e3), v/physics.C)
	if v >= physics.C {
		text = "escape velocity ≥ c · no light gets out"
	}
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(240, func(p *giode.Painter) { e.draw(p, r, v/physics.C) }),
		readout(text),
		e.radius,
	).Layout(gtx)
}

func (e *Escape) draw(p *giode.Painter, r, frac float64) {
	cx, cy := p.Width()/2, p.Height()/2
	starfield(p)
	rd := drawnRadius(r)
	horizon := drawnRadius(physics.SchwarzschildRadius(physics.SolarMass))
	if frac >= 1 {
		p.Circle(cx, cy, horizon, "#000000")
		p.Arc(cx, cy, horizon+3, 2, 0, 360, theme.Accent)
		label(p, cx, cy+horizon+12, "event horizon · 2.95 km", giode.TextAnchorTop, theme.Accent)
		return
	}
	color := lerpColor("#ffe08a", "#ff5a3c", math.Pow(frac, 0.35))
	glow(p, cx, cy, rd, color)
	rayLen := 60 * float32(1-frac*frac)
	for i := 0; i < 12; i++ {
		a := float64(i) * math.Pi / 6
		dx, dy := float32(math.Cos(a)), float32(math.Sin(a))
		x1, y1 := cx+dx*(rd+6), cy+dy*(rd+6)
		arrow(p, x1, y1, x1+dx*rayLen, y1+dy*rayLen, 1.5, color)
	}
	p.Circle(cx, cy, rd, color)
	p.Arc(cx, cy, horizon, 1, 0, 360, theme.Faint)
	label(p, 10, p.Height()-10, "inner ring: horizon size, 2.95 km", giode.TextAnchorBottomLeft, theme.Muted)
}

// starfield scatters fixed background stars.
func starfield(p *giode.Painter) {
	for i := 0; i < 40; i++ {
		x := float32(math.Mod(float64(i)*97.3, float64(p.Width())))
		y := float32(math.Mod(float64(i)*61.7, float64(p.Height())))
		p.Circle(x, y, 1, "#3a4466")
	}
}

// formatKm prints a radius in km with digit grouping.
func formatKm(m float64) string {
	return groupDigits(m/1e3) + " km"
}

// groupDigits prints a number rounded to an integer, with commas.
func groupDigits(v float64) string {
	s := fmt.Sprintf("%.0f", v)
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}
