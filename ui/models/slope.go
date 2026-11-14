package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// Slope shows the tangent to f(x) = x³/3 − x at a point the reader
// moves; its slope is the derivative x² − 1.
type Slope struct {
	point *SliderRow
}

// NewSlope creates the model.
func NewSlope() *Slope {
	return &Slope{point: NewSliderRow("Point on the curve", 0.75, func(v float32) string {
		return fmt.Sprintf("x = %.2f", slopeX(v))
	})}
}

func slopeX(v float32) float64     { return lerp(-2.2, 2.2, v) }
func slopeCurve(x float64) float64 { return x*x*x/3 - x }
func slopeDeriv(x float64) float64 { return x*x - 1 }

// Layout draws the curve, tangent, readout and slider.
func (s *Slope) Layout(gtx giode.Context) giode.Dimensions {
	x := slopeX(s.point.Value())
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(220, func(p *giode.Painter) { s.draw(p, x) }),
		readout(fmt.Sprintf("f(x) = %.2f   slope = %.2f", slopeCurve(x), slopeDeriv(x))),
		s.point,
	).Layout(gtx)
}

func (s *Slope) draw(p *giode.Painter, x float64) {
	pl := newPlot(p, 14, -2.5, 2.5, -2.4, 2.4)
	pl.axes(p)
	pl.curve(p, slopeCurve, 2.5, theme.Primary)
	m, y := slopeDeriv(x), slopeCurve(x)
	c := giode.Color(theme.Success)
	switch {
	case math.Abs(m) < 0.08:
		c = theme.Accent
	case m < 0:
		c = theme.Danger
	}
	const reach = 1.1
	p.Line(pl.X(x-reach), pl.Y(y-m*reach), pl.X(x+reach), pl.Y(y+m*reach), 2, c)
	p.Circle(pl.X(x), pl.Y(y), 7, theme.Accent)
	label(p, pl.X(2.45), pl.Y(0)-4, "x", giode.TextAnchorBottomRight, theme.Muted)
	label(p, pl.X(1.9), pl.Y(slopeCurve(1.9)), "f(x)  ", giode.TextAnchorRight, theme.Primary)
	label(p, pl.X(x+reach), pl.Y(y+m*reach)-4, "tangent", giode.TextAnchorBottom, c)
}
