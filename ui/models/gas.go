package models

import (
	"fmt"
	"math"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Gas shows molecules bouncing in a box. Raising the temperature makes
// them faster, so they hit the walls harder and more often: pressure.
type Gas struct {
	temperature *SliderRow
	gas         *physics.Gas
	last        time.Time
	sampled     time.Time
	pressure    float64
}

const gasMolecules = 70

// NewGas creates the model.
func NewGas() *Gas {
	return &Gas{gas: physics.NewGas(gasMolecules, 3), temperature: NewSliderRow("Temperature", 0.3, func(v float32) string {
		return fmt.Sprintf("%.0f K", gasKelvin(v))
	})}
}

func gasKelvin(v float32) float64 { return lerp(50, 900, v) }

// advance steps the simulation by the time since the last frame.
func (m *Gas) advance(scale float64) {
	now := time.Now()
	if !m.last.IsZero() {
		m.gas.Step(math.Min(now.Sub(m.last).Seconds(), 0.05)*0.35, scale)
	}
	m.last = now
	if now.Sub(m.sampled) > time.Second {
		m.pressure, m.sampled = m.gas.Pressure(), now
	}
}

// Layout draws the box, readout and slider.
func (m *Gas) Layout(gtx giode.Context) giode.Dimensions {
	k := gasKelvin(m.temperature.Value())
	m.advance(k / 300)
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(240, func(p *giode.Painter) { m.draw(p, k) }),
		readout(m.summary()),
		m.temperature,
	).Layout(gtx)
}

// summary reports the measured wall pressure.
func (m *Gas) summary() string {
	if m.pressure == 0 {
		return "measuring pressure…"
	}
	return fmt.Sprintf("pressure %.1f · at fixed volume P ∝ T", m.pressure)
}

func (m *Gas) draw(p *giode.Painter, kelvin float64) {
	side := min(p.Width(), p.Height()) - 24
	ox, oy := (p.Width()-side)/2, (p.Height()-side)/2
	p.Rect(ox, oy, side, side, 4, "#101830")
	box := p.Path()
	box.MoveTo(ox, oy)
	box.LineTo(ox+side, oy)
	box.LineTo(ox+side, oy+side)
	box.LineTo(ox, oy+side)
	box.Close()
	box.Stroke(2, theme.Primary)
	color := lerpColor("#6cb8ff", "#ff7a59", (kelvin-50)/850)
	for _, mol := range m.gas.Molecules {
		p.Circle(ox+float32(mol.X)*side, oy+float32(mol.Y)*side, 3.5, color)
	}
	keepAnimating(p)
}
