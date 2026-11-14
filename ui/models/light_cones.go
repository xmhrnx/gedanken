package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// LightCones draws light cones near a black hole in Eddington–Finkelstein
// coordinates: time up, distance right. Cones tip toward the hole and,
// inside the horizon, every future direction leads inward.
type LightCones struct {
	probe *SliderRow
}

// Distances in units of rs.
const conesMaxR = 3.0

var coneRadii = []float64{0.35, 0.7, 1.35, 2.0, 2.7}

// NewLightCones creates the model.
func NewLightCones() *LightCones {
	return &LightCones{probe: NewSliderRow("Probe distance", 0.55, func(v float32) string {
		return fmt.Sprintf("%.2f r_s", conesProbe(v))
	})}
}

func conesProbe(v float32) float64 { return lerp(0.25, 2.9, v) }

// Layout draws the diagram, readout and slider.
func (l *LightCones) Layout(gtx giode.Context) giode.Dimensions {
	r := conesProbe(l.probe.Value())
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(230, func(p *giode.Painter) { l.draw(p, r) }),
		readout(coneVerdict(r)),
		l.probe,
	).Layout(gtx)
}

// coneVerdict describes what outgoing light does at r.
func coneVerdict(r float64) string {
	switch {
	case math.Abs(r-1) < 0.04:
		return "at the horizon: outgoing light stays put"
	case r < 1:
		return "inside: every future path leads to r = 0"
	default:
		return fmt.Sprintf("outside: light escapes, cone slope %.1f", physics.OutgoingLightSlope(r, 1))
	}
}

func (l *LightCones) draw(p *giode.Painter, probe float64) {
	const margin = 16
	x := func(r float64) float32 { return margin + float32(r/conesMaxR)*(p.Width()-2*margin) }
	baseY := p.Height() * 0.62
	zigzag(p, x(0), p.Height())
	dashed(p, x(1), 8, x(1), p.Height()-8, 1.5, theme.Danger)
	for _, r := range coneRadii {
		cone(p, x(r), baseY, r, 34, "#2f3d63")
	}
	cone(p, x(probe), baseY, probe, 60, theme.PrimarySoft)
	p.Circle(x(probe), baseY, 5, theme.Accent)
	label(p, x(0)+8, 8, "r = 0", giode.TextAnchorTopLeft, theme.Danger)
	label(p, x(1)+6, 8, "horizon", giode.TextAnchorTopLeft, theme.Danger)
	label(p, p.Width()-10, p.Height()-8, "distance r →", giode.TextAnchorBottomRight, theme.Muted)
	p.TextRotated(p.Width()-10, 40, "time →", giode.TextAnchorCenter, -90, giode.Styles{FontSize: 11, Color: theme.Muted})
}

// cone fills the future light cone at screen point (cx, cy) for radius r.
func cone(p *giode.Painter, cx, cy float32, r float64, length float32, c giode.Color) {
	inX, inY := unit(-1, 1)
	outX, outY := outgoingDirection(r)
	path := p.Path()
	path.MoveTo(cx, cy)
	path.LineTo(cx+inX*length, cy-inY*length)
	path.LineTo(cx+outX*length, cy-outY*length)
	path.Close()
	path.Fill(c)
	p.Line(cx, cy, cx+inX*length, cy-inY*length, 1.5, theme.Primary)
	p.Line(cx, cy, cx+outX*length, cy-outY*length, 1.5, theme.Accent)
}

// outgoingDirection is the unit (dr, dt) of outgoing light at r.
func outgoingDirection(r float64) (float32, float32) {
	s := physics.OutgoingLightSlope(r, 1)
	switch {
	case math.IsInf(s, 1) || math.Abs(r-1) < 0.02:
		return 0, 1
	case s > 0:
		return unit(1, s)
	default:
		return unit(-1, -s)
	}
}

func unit(x, y float64) (float32, float32) {
	n := math.Hypot(x, y)
	return float32(x / n), float32(y / n)
}

// zigzag marks the singularity at r = 0.
func zigzag(p *giode.Painter, x, height float32) {
	path := p.Path()
	path.MoveTo(x, 0)
	for y, i := float32(0), 0; y < height; y, i = y+8, i+1 {
		dx := float32(4)
		if i%2 == 1 {
			dx = -4
		}
		path.LineTo(x+dx, y+8)
	}
	path.Stroke(2, theme.Danger)
}
