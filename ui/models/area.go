package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// Area approximates the area under f(x) = 0.3x² + 0.5 on [0, 3] with
// rectangles; more rectangles approach the exact integral, 4.2.
type Area struct {
	count *SliderRow
}

const areaExact = 4.2

// NewArea creates the model.
func NewArea() *Area {
	return &Area{count: NewSliderRow("Rectangles", 0.15, func(v float32) string {
		return fmt.Sprintf("n = %d", areaN(v))
	})}
}

func areaN(v float32) int         { return 1 + int(math.Round(float64(v)*39)) }
func areaCurve(x float64) float64 { return 0.3*x*x + 0.5 }

// areaSum is the left Riemann sum with n rectangles.
func areaSum(n int) float64 {
	dx, sum := 3/float64(n), 0.0
	for i := 0; i < n; i++ {
		sum += areaCurve(float64(i)*dx) * dx
	}
	return sum
}

// Layout draws the rectangles, curve, readout and slider.
func (a *Area) Layout(gtx giode.Context) giode.Dimensions {
	n := areaN(a.count.Value())
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(220, func(p *giode.Painter) { a.draw(p, n) }),
		readout(fmt.Sprintf("sum = %.3f   exact = %.3f", areaSum(n), areaExact)),
		a.count,
	).Layout(gtx)
}

func (a *Area) draw(p *giode.Painter, n int) {
	pl := newPlot(p, 14, -0.1, 3.1, -0.15, 3.4)
	dx := 3 / float64(n)
	for i := 0; i < n; i++ {
		x := float64(i) * dx
		top := pl.Y(areaCurve(x))
		p.Rect(pl.X(x), top, pl.X(x+dx)-pl.X(x), pl.Y(0)-top, 0, theme.PrimarySoft)
		p.Line(pl.X(x), top, pl.X(x), pl.Y(0), 1, theme.Primary)
	}
	pl.axes(p)
	pl.curve(p, areaCurve, 2.5, theme.Accent)
	label(p, pl.X(0.05), pl.Y(3.3), "f(x) = 0.3x² + 0.5", giode.TextAnchorTopLeft, theme.Accent)
	label(p, pl.X(3), pl.Y(0)+4, "3", giode.TextAnchorTop, theme.Muted)
	label(p, pl.X(0), pl.Y(0)+4, "0", giode.TextAnchorTop, theme.Muted)
}
