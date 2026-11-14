package models

import (
	"fmt"
	"math"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Pendulum swings in real time; its period depends on length, not on
// mass or (for small swings) amplitude.
type Pendulum struct {
	length *SliderRow
	start  time.Time
}

const pendulumAmplitude = 0.35 // radians

// NewPendulum creates the model.
func NewPendulum() *Pendulum {
	return &Pendulum{start: time.Now(), length: NewSliderRow("Length", 0.4, func(v float32) string {
		return fmt.Sprintf("%.2f m", pendulumLength(v))
	})}
}

func pendulumLength(v float32) float64 { return lerp(0.25, 2.5, v) }

// Layout draws the pendulum, readout and slider.
func (m *Pendulum) Layout(gtx giode.Context) giode.Dimensions {
	l := pendulumLength(m.length.Value())
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(240, func(p *giode.Painter) { m.draw(p, l) }),
		readout(fmt.Sprintf("period T = 2π√(L/g) = %.2f s", physics.PendulumPeriod(l, physics.EarthG))),
		m.length,
	).Layout(gtx)
}

func (m *Pendulum) draw(p *giode.Painter, l float64) {
	px, py := p.Width()/2, float32(18)
	scale := (p.Height() - 50) / 2.5
	theta := physics.PendulumAngle(pendulumAmplitude, l, physics.EarthG, sinceSeconds(m.start))
	r := float32(l) * scale
	bx := px + r*float32(math.Sin(theta))
	by := py + r*float32(math.Cos(theta))
	p.Line(px-40, py, px+40, py, 3, theme.Faint)
	for _, a := range []float64{-pendulumAmplitude, pendulumAmplitude} {
		ex, ey := px+r*float32(math.Sin(a)), py+r*float32(math.Cos(a))
		dashed(p, px, py, ex, ey, 1, "#26324f")
	}
	p.Line(px, py, bx, by, 2, theme.Text)
	glow(p, bx, by, 12, "#f6bd6e")
	p.Circle(bx, by, 12, theme.Accent)
	label(p, px+8, py+r/2, "L", giode.TextAnchorLeft, theme.Muted)
	keepAnimating(p)
}
