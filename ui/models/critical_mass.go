package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// CriticalMass shows neutrons born inside a sphere. Their typical path
// length is fixed, so a bigger sphere loses a smaller share through its
// surface and k rises past 1.
type CriticalMass struct {
	radius  *SliderRow
	cached  physics.CriticalSample
	forStep int
}

// Sphere radius range, in mean free paths.
const critMinR, critMaxR, critSteps = 0.5, 8.0, 60

// NewCriticalMass creates the model.
func NewCriticalMass() *CriticalMass {
	return &CriticalMass{forStep: -1, radius: NewSliderRow("Sphere radius", 0.35, func(v float32) string {
		return fmt.Sprintf("%.1f mean free paths", critRadius(critStep(v)))
	})}
}

func critStep(v float32) int      { return int(math.Round(float64(v) * critSteps)) }
func critRadius(step int) float64 { return critMinR + (critMaxR-critMinR)*float64(step)/critSteps }

// sample reruns the simulation only when the radius changes.
func (c *CriticalMass) sample() physics.CriticalSample {
	if step := critStep(c.radius.Value()); step != c.forStep {
		c.forStep = step
		c.cached = physics.SampleSphere(critRadius(step), 4000, 80, 7)
	}
	return c.cached
}

// Layout draws the sphere, readout and slider.
func (c *CriticalMass) Layout(gtx giode.Context) giode.Dimensions {
	s := c.sample()
	k := s.K()
	state := "subcritical"
	switch {
	case math.Abs(k-1) < 0.02:
		state = "critical"
	case k > 1:
		state = "supercritical"
	}
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(250, func(p *giode.Painter) { c.draw(p, s) }),
		readout(fmt.Sprintf("%.0f%% escape · k = %.2f · %s", s.EscapeFrac*100, k, state)),
		c.radius,
	).Layout(gtx)
}

func (c *CriticalMass) draw(p *giode.Painter, s physics.CriticalSample) {
	cx, cy := p.Width()/2, p.Height()/2
	scale := (min(p.Width(), p.Height())/2 - 12) / critMaxR
	r := float32(s.Radius) * scale
	p.Circle(cx, cy, r, "#1b2747")
	p.Arc(cx, cy, r, 2, 0, 360, theme.Primary)
	for _, t := range s.Tracks {
		x0, y0 := cx+float32(t.X0)*scale, cy+float32(t.Y0)*scale
		x1, y1 := cx+float32(t.X1)*scale, cy+float32(t.Y1)*scale
		col := giode.Color(theme.Success)
		if t.Escaped {
			col = theme.Danger
		}
		p.Line(x0, y0, x1, y1, 1.2, col)
		p.Circle(x1, y1, 2.2, col)
	}
	label(p, 10, 10, "● absorbed", giode.TextAnchorTopLeft, theme.Success)
	label(p, 10, 26, "● escaped", giode.TextAnchorTopLeft, theme.Danger)
	label(p, p.Width()-10, p.Height()-10, "path length ~ 1 mean free path", giode.TextAnchorBottomRight, theme.Muted)
}
