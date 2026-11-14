package models

import (
	"fmt"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Wave is a travelling sine wave. Speed is wavelength × frequency; a
// marked point on the medium only moves up and down.
type Wave struct {
	wavelength, frequency *SliderRow
	start                 time.Time
}

// NewWave creates the model.
func NewWave() *Wave {
	return &Wave{
		start:      time.Now(),
		wavelength: NewSliderRow("Wavelength λ", 0.4, func(v float32) string { return fmt.Sprintf("%.1f m", waveLength(v)) }),
		frequency:  NewSliderRow("Frequency f", 0.3, func(v float32) string { return fmt.Sprintf("%.2f Hz", waveFreq(v)) }),
	}
}

func waveLength(v float32) float64 { return lerp(0.5, 4, v) }
func waveFreq(v float32) float64   { return lerp(0.1, 1.5, v) }

// Layout draws the wave, readout and sliders.
func (m *Wave) Layout(gtx giode.Context) giode.Dimensions {
	l, f := waveLength(m.wavelength.Value()), waveFreq(m.frequency.Value())
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(190, func(p *giode.Painter) { m.draw(p, l, f) }),
		readout(fmt.Sprintf("speed v = fλ = %.2f m/s · period %.2f s", f*l, 1/f)),
		m.wavelength, m.frequency,
	).Layout(gtx)
}

func (m *Wave) draw(p *giode.Painter, l, f float64) {
	const worldW = 8.0
	t := sinceSeconds(m.start)
	pl := newPlot(p, 14, 0, worldW, -1.4, 1.4)
	pl.axes(p)
	pl.curve(p, func(x float64) float64 { return physics.WaveHeight(1, l, f, x, t) }, 2.5, theme.Info)
	mark := worldW / 2
	y := physics.WaveHeight(1, l, f, mark, t)
	p.Circle(pl.X(mark), pl.Y(y), 6, theme.Accent)
	label(p, p.Width()-12, 10, "● one point of the medium", giode.TextAnchorTopRight, theme.Accent)
	arrow(p, pl.X(0.2), pl.Y(-1.25), pl.X(0.2+l), pl.Y(-1.25), 1.5, theme.Muted)
	label(p, pl.X(0.2+l/2), pl.Y(-1.25)-4, "λ", giode.TextAnchorBottom, theme.Muted)
	keepAnimating(p)
}
