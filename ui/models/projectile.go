package models

import (
	"fmt"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Projectile launches a ball at a chosen angle and speed and traces its
// parabola.
type Projectile struct {
	angle, speed *SliderRow
	launch       *giode.ButtonElement
	launchedAt   time.Time
}

// World size shown, in metres: the longest possible range and height.
const projWorldW, projWorldH, projPlayback = 95.0, 48.0, 1.6

// NewProjectile creates the model.
func NewProjectile() *Projectile {
	m := &Projectile{
		angle: NewSliderRow("Launch angle", 0.47, func(v float32) string { return fmt.Sprintf("%.0f°", projAngle(v)) }),
		speed: NewSliderRow("Launch speed", 0.75, func(v float32) string { return fmt.Sprintf("%.0f m/s", projSpeed(v)) }),
	}
	m.launch = giode.Button("Launch", theme.PrimaryButton).OnClick(func() { m.launchedAt = time.Now() })
	return m
}

func projAngle(v float32) float64 { return lerp(5, 85, v) }
func projSpeed(v float32) float64 { return lerp(5, 30, v) }

func (m *Projectile) shot() physics.Projectile {
	return physics.Projectile{Speed: projSpeed(m.speed.Value()), Angle: projAngle(m.angle.Value()), G: physics.EarthG}
}

// Layout draws the flight, readout, sliders and launch button.
func (m *Projectile) Layout(gtx giode.Context) giode.Dimensions {
	s := m.shot()
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(210, func(p *giode.Painter) { m.draw(p, s) }),
		readout(fmt.Sprintf("range %.1f m · peak %.1f m · %.1f s", s.Range(), s.MaxHeight(), s.FlightTime())),
		m.angle, m.speed, m.launch,
	).Layout(gtx)
}

func (m *Projectile) draw(p *giode.Painter, s physics.Projectile) {
	const margin = 14
	scale := min((p.Width()-2*margin)/projWorldW, (p.Height()-2*margin)/projWorldH)
	ox, oy := float32(margin), p.Height()-margin
	at := func(t float64) (float32, float32) {
		x, y := s.At(t)
		return ox + float32(x)*scale, oy - float32(y)*scale
	}
	p.Line(0, oy, p.Width(), oy, 1, theme.Faint)

	total := s.FlightTime()
	elapsed := total
	if !m.launchedAt.IsZero() {
		elapsed = min(total, sinceSeconds(m.launchedAt)*projPlayback)
	}
	const steps = 80
	var px, py float32
	for i := 0; i <= steps; i++ {
		x, y := at(total * float64(i) / steps)
		if i > 0 {
			dashed(p, px, py, x, y, 1, theme.Faint)
		}
		px, py = x, y
	}
	path := p.Path()
	for i := 0; i <= steps; i++ {
		x, y := at(elapsed * float64(i) / steps)
		if i == 0 {
			path.MoveTo(x, y)
		} else {
			path.LineTo(x, y)
		}
	}
	path.Stroke(2.5, theme.Primary)
	bx, by := at(elapsed)
	p.Circle(bx, by, 7, theme.Accent)
	if elapsed >= total {
		lx, ly := at(total)
		label(p, lx, ly-10, fmt.Sprintf("%.0f m", s.Range()), giode.TextAnchorBottom, theme.Accent)
		hx, hy := at(total / 2)
		label(p, hx, hy-6, fmt.Sprintf("%.0f m high", s.MaxHeight()), giode.TextAnchorBottom, theme.Muted)
	}
	if elapsed < total {
		keepAnimating(p)
	}
}
