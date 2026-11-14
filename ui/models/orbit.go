package models

import (
	"fmt"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Orbit is Newton's cannon: a ball fired sideways from a mountain falls
// back, circles the planet, traces an ellipse or escapes, by speed.
type Orbit struct {
	speed   *SliderRow
	fire    *giode.ButtonElement
	firedAt time.Time
}

const orbitHeight, orbitPlayback = 0.12, 0.35 // pts per ms of replay

// NewOrbit creates the model.
func NewOrbit() *Orbit {
	m := &Orbit{speed: NewSliderRow("Launch speed", 0.5, func(v float32) string {
		return fmt.Sprintf("%.2f × circular speed", orbitSpeed(v))
	})}
	m.fire = giode.Button("Fire", theme.PrimaryButton).OnClick(func() { m.firedAt = time.Now() })
	return m
}

func orbitSpeed(v float32) float64 { return lerp(0.4, 1.6, v) }

// orbitVerdict names the outcome at a given speed.
func orbitVerdict(speed float64, o physics.OrbitOutcome) string {
	switch {
	case o == physics.Crashes:
		return "falls back to the ground"
	case o == physics.Escapes:
		return "escapes: speed ≥ √2 × circular"
	case speed > 0.97 && speed < 1.03:
		return "circular orbit"
	default:
		return "elliptical orbit"
	}
}

// Layout draws the planet, path, readout, slider and button.
func (m *Orbit) Layout(gtx giode.Context) giode.Dimensions {
	speed := orbitSpeed(m.speed.Value())
	path, outcome := physics.Cannon{Height: orbitHeight, Speed: speed}.Path()
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(260, func(p *giode.Painter) { m.draw(p, path) }),
		readout(orbitVerdict(speed, outcome)),
		m.speed, m.fire,
	).Layout(gtx)
}

func (m *Orbit) draw(p *giode.Painter, path []physics.Point) {
	cx, cy := p.Width()/2, p.Height()*0.34
	scale := min(p.Width(), p.Height()) / 6.5
	at := func(pt physics.Point) (float32, float32) { return cx + float32(pt.X)*scale, cy - float32(pt.Y)*scale }
	starfield(p)
	glow(p, cx, cy, scale, "#3b6fd1")
	p.Circle(cx, cy, scale, "#2a4f9a")
	mx, my := at(physics.Point{Y: 1})
	tri := p.Path()
	tri.MoveTo(mx-10, my+2)
	tri.LineTo(mx, my-float32(orbitHeight)*scale)
	tri.LineTo(mx+10, my+2)
	tri.Close()
	tri.Fill("#5a6b8f")

	shown := len(path)
	if !m.firedAt.IsZero() {
		shown = min(len(path), 1+int(float64(time.Since(m.firedAt).Milliseconds())*orbitPlayback))
		if shown < len(path) {
			keepAnimating(p)
		}
	}
	line := p.Path()
	for i, pt := range path[:shown] {
		x, y := at(pt)
		if i == 0 {
			line.MoveTo(x, y)
		} else {
			line.LineTo(x, y)
		}
	}
	line.Stroke(2, theme.Accent)
	bx, by := at(path[shown-1])
	p.Circle(bx, by, 5, theme.Accent)
	label(p, cx, cy, "planet", giode.TextAnchorCenter, "#c9d6f5")
}
