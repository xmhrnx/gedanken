package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// GravityWell draws Flamm's paraboloid, a picture of the curved space
// outside a black hole, with a probe whose clock rate is shown.
type GravityWell struct {
	probe *SliderRow
}

// Radii in units of the horizon radius rs.
const wellMaxR = 6.0

// NewGravityWell creates the model.
func NewGravityWell() *GravityWell {
	return &GravityWell{probe: NewSliderRow("Probe distance", 0.3, func(v float32) string {
		return fmt.Sprintf("%.2f r_s", wellProbe(v))
	})}
}

func wellProbe(v float32) float64 { return lerp(1.02, wellMaxR, v) }

// Layout draws the well, readout and slider.
func (g *GravityWell) Layout(gtx giode.Context) giode.Dimensions {
	r := wellProbe(g.probe.Value())
	rate := physics.TimeDilation(r, 1)
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(240, func(p *giode.Painter) { g.draw(p, r) }),
		readout(fmt.Sprintf("a clock here ticks at %.0f%% of the far-away rate", rate*100)),
		g.probe,
	).Layout(gtx)
}

func (g *GravityWell) draw(p *giode.Painter, probe float64) {
	cx := p.Width() / 2
	rx := (p.Width()/2 - 16) / wellMaxR
	ry := rx * 0.3
	maxDepth := physics.FlammDepth(wellMaxR, 1)
	depthScale := (p.Height() - 40 - 2*ry*wellMaxR*0.5) / float32(maxDepth)
	ringY := func(r float64) float32 {
		return 24 + float32(maxDepth-physics.FlammDepth(r, 1))*depthScale
	}
	ring := func(r float64, w float32, c giode.Color) {
		ellipse(p, cx, ringY(r), float64(rx)*r, float64(ry)*r, 0, w, c)
	}
	for i := 0; i < 14; i++ {
		ring(1+math.Pow(float64(i)/13, 1.6)*(wellMaxR-1), 1, "#2f3d63")
	}
	for k := 0; k < 16; k++ {
		a := float64(k) * math.Pi / 8
		path := p.Path()
		for i := 0; i <= 30; i++ {
			r := 1 + float64(i)/30*(wellMaxR-1)
			x, y := cx+float32(math.Cos(a)*r)*rx, ringY(r)+float32(math.Sin(a)*r)*ry
			if i == 0 {
				path.MoveTo(x, y)
			} else {
				path.LineTo(x, y)
			}
		}
		path.Stroke(1, "#26324f")
	}
	ellipseFill(p, cx, ringY(1), rx, ry, "#000000")
	ring(1, 2, theme.Danger)
	ring(probe, 2, theme.Accent)
	p.Circle(cx+float32(probe)*rx, ringY(probe), 6, theme.Accent)
	label(p, cx, ringY(1)+ry+6, "horizon", giode.TextAnchorTop, theme.Danger)
	label(p, cx+float32(probe)*rx+10, ringY(probe), "probe", giode.TextAnchorLeft, theme.Accent)
}

// ellipseFill fills an axis-aligned ellipse.
func ellipseFill(p *giode.Painter, cx, cy, rx, ry float32, c giode.Color) {
	const n = 48
	path := p.Path()
	for i := 0; i <= n; i++ {
		a := 2 * math.Pi * float64(i) / n
		x, y := cx+rx*float32(math.Cos(a)), cy+ry*float32(math.Sin(a))
		if i == 0 {
			path.MoveTo(x, y)
		} else {
			path.LineTo(x, y)
		}
	}
	path.Close()
	path.Fill(c)
}
