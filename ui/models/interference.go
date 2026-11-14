package models

import (
	"fmt"
	"math"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Interference shows waves from two sources overlapping, and the
// bright and dark fringes they make on a screen.
type Interference struct {
	separation *SliderRow
	start      time.Time
	shades     []giode.Color
}

const (
	interWorldW = 9.0 // wavelengths across the stage
	interCell   = 6   // dp per drawn cell
	interShades = 24
)

// NewInterference creates the model.
func NewInterference() *Interference {
	m := &Interference{start: time.Now(), separation: NewSliderRow("Source separation d", 0.35, func(v float32) string {
		return fmt.Sprintf("%.1f λ", interSeparation(v))
	})}
	for i := 0; i < interShades; i++ {
		m.shades = append(m.shades, lerpColor("#070b18", "#a8ecff", float64(i)/(interShades-1)))
	}
	return m
}

func interSeparation(v float32) float64 { return lerp(0.6, 5, v) }

// Layout draws the wave field, readout and slider.
func (m *Interference) Layout(gtx giode.Context) giode.Dimensions {
	d := interSeparation(m.separation.Value())
	fringes := int(math.Floor(d)) * 2
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(260, func(p *giode.Painter) { m.draw(p, d) }),
		readout(fmt.Sprintf("%d dark fringes on each side · wider d, narrower fringes", fringes/2)),
		m.separation,
	).Layout(gtx)
}

func (m *Interference) draw(p *giode.Painter, d float64) {
	phase := math.Mod(sinceSeconds(m.start)*0.8, 1)
	screenH := float32(26)
	fieldH := p.Height() - screenH
	unit := float64(p.Width()) / interWorldW
	for py := float32(0); py < fieldH; py += interCell {
		for px := float32(0); px < p.Width(); px += interCell {
			x := (float64(px+interCell/2) - float64(p.Width())/2) / unit
			y := float64(py+interCell/2) / unit
			a := physics.TwoSourceAmplitude(x, y, d, 1, phase)
			i := int((a + 2) / 4 * (interShades - 1))
			p.Rect(px, py, interCell, interCell, 0, m.shades[max(0, min(interShades-1, i))])
		}
	}
	for px := float32(0); px < p.Width(); px += 3 {
		x := (float64(px) - float64(p.Width())/2) / unit
		theta := math.Atan2(x, float64(fieldH)/unit)
		in := physics.FringeIntensity(theta, d, 1)
		p.Rect(px, fieldH+4, 3, screenH-8, 0, lerpColor("#0b0f1c", "#f6e7a6", in))
	}
	for _, s := range []float64{-d / 2, d / 2} {
		p.Circle(float32(p.Width()/2)+float32(s*unit), 3, 5, theme.Accent)
	}
	label(p, 8, fieldH+screenH/2, "screen", giode.TextAnchorLeft, theme.Muted)
	keepAnimating(p)
}
