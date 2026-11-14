package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// BindingCurve plots binding energy per nucleon against mass number.
// Light nuclei release energy by fusing, heavy ones by splitting; both
// move toward the iron peak.
type BindingCurve struct {
	pick *SliderRow
}

// NewBindingCurve creates the model.
func NewBindingCurve() *BindingCurve {
	return &BindingCurve{pick: NewSliderRow("Nucleus", 1, func(v float32) string {
		return bindingPick(v).Name
	})}
}

// bindingPick maps the slider onto a nuclide.
func bindingPick(v float32) physics.Nuclide {
	n := len(physics.BindingCurve)
	return physics.BindingCurve[int(math.Round(float64(v)*float64(n-1)))]
}

// Layout draws the curve, legend, readout and slider.
func (b *BindingCurve) Layout(gtx giode.Context) giode.Dimensions {
	n := bindingPick(b.pick.Value())
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(230, func(p *giode.Painter) { b.draw(p, n) }),
		giode.Row(giode.Styles{Width: -1, Justify: giode.JustifySpaceBetween},
			giode.Text("← fusion releases energy", giode.Styles{FontSize: 12, Color: theme.Info}),
			giode.Text("fission releases energy →", giode.Styles{FontSize: 12, Color: theme.Danger}),
		),
		readout(fmt.Sprintf("A = %d · %.2f MeV per nucleon", n.A, n.Binding)),
		b.pick,
	).Layout(gtx)
}

func (b *BindingCurve) draw(p *giode.Painter, pick physics.Nuclide) {
	pl := newPlot(p, 16, 0, 245, 0, 10.8)
	for e := 2.0; e <= 8; e += 2 {
		dashed(p, pl.X(0), pl.Y(e), pl.X(245), pl.Y(e), 1, "#1d2741")
	}
	pl.axes(p)
	path := p.Path()
	for i, n := range physics.BindingCurve {
		if i == 0 {
			path.MoveTo(pl.X(float64(n.A)), pl.Y(n.Binding))
		} else {
			path.LineTo(pl.X(float64(n.A)), pl.Y(n.Binding))
		}
	}
	path.Stroke(2.5, theme.Primary)
	for _, n := range physics.BindingCurve {
		p.Circle(pl.X(float64(n.A)), pl.Y(n.Binding), 2.5, theme.Primary)
	}
	peak := physics.MostBound()
	p.Circle(pl.X(float64(peak.A)), pl.Y(peak.Binding), 5, theme.Success)
	arrow(p, pl.X(10), pl.Y(3.5), pl.X(40), pl.Y(6.8), 2, theme.Info)
	arrow(p, pl.X(225), pl.Y(6.6), pl.X(130), pl.Y(7.6), 2, theme.Danger)
	p.Circle(pl.X(float64(pick.A)), pl.Y(pick.Binding), 7, theme.Accent)
	label(p, pl.X(float64(peak.A)), pl.Y(peak.Binding)-8, "iron peak", giode.TextAnchorBottom, theme.Success)
	label(p, pl.X(4)+8, pl.Y(7.07)+6, "He-4", giode.TextAnchorTopLeft, theme.Muted)
	label(p, pl.X(235), pl.Y(7.59)+8, "U-235", giode.TextAnchorTop, theme.Muted)
	label(p, pl.X(245), pl.Y(0)-4, "mass number A", giode.TextAnchorBottomRight, theme.Muted)
	label(p, pl.X(4), pl.Y(10.8), "MeV per nucleon", giode.TextAnchorTopLeft, theme.Muted)
}
