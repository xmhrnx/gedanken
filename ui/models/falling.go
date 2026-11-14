package models

import (
	"fmt"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Falling drops a ball in slow motion and leaves a mark every 0.1 s,
// so the growing gaps (1 : 3 : 5 : 7 ...) show the acceleration.
type Falling struct {
	drop      *giode.ButtonElement
	droppedAt time.Time
}

const (
	fallDuration = 1.0 // simulated seconds
	fallInterval = 0.1 // seconds between marks
	fallSlowMo   = 0.4 // playback speed
)

// NewFalling creates the model.
func NewFalling() *Falling {
	m := &Falling{}
	m.drop = giode.Button("Drop", theme.PrimaryButton).OnClick(func() { m.droppedAt = time.Now() })
	return m
}

// elapsed is the simulated time since the drop.
func (m *Falling) elapsed() float64 {
	if m.droppedAt.IsZero() {
		return 0
	}
	return min(fallDuration, sinceSeconds(m.droppedAt)*fallSlowMo)
}

// Layout draws the drop, readout and button.
func (m *Falling) Layout(gtx giode.Context) giode.Dimensions {
	t := m.elapsed()
	d := 0.5 * physics.EarthG * t * t
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(300, func(p *giode.Painter) { m.draw(p, t) }),
		readout(fmt.Sprintf("t %.2f s · fallen %.2f m · %.1f m/s", t, d, physics.EarthG*t)),
		m.drop,
	).Layout(gtx)
}

func (m *Falling) draw(p *giode.Painter, t float64) {
	const top, bottomPad = 20, 20
	maxDrop := 0.5 * physics.EarthG * fallDuration * fallDuration
	scale := (p.Height() - top - bottomPad) / float32(maxDrop)
	cx := p.Width() / 2
	y := func(tt float64) float32 { return top + float32(0.5*physics.EarthG*tt*tt)*scale }

	for metre := 0; metre <= int(maxDrop); metre++ {
		ty := top + float32(metre)*scale
		p.Line(36, ty, 50, ty, 1.5, theme.Faint)
		label(p, 30, ty, fmt.Sprintf("%d m", metre), giode.TextAnchorRight, theme.Muted)
	}
	p.Line(43, top, 43, y(fallDuration), 1, theme.Faint)
	label(p, cx+46, y(fallInterval), "every 0.1 s", giode.TextAnchorLeft, theme.Primary)

	for k := 1; float64(k)*fallInterval <= t+1e-9; k++ {
		gy := y(float64(k) * fallInterval)
		p.Circle(cx, gy, 9, theme.PrimarySoft)
		p.Line(cx+18, gy, cx+40, gy, 1, theme.Primary)
	}
	p.Circle(cx, y(0), 9, theme.PrimarySoft)
	p.Circle(cx, y(t), 10, theme.Accent)
	if !m.droppedAt.IsZero() && t < fallDuration {
		keepAnimating(p)
	}
}
