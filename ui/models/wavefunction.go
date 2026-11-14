package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Wavefunction shows a particle in a box: the standing wave ψ for level
// n and the probability |ψ|² of finding the particle at each place.
type Wavefunction struct {
	level *SliderRow
}

const boxLevels = 6

// NewWavefunction creates the model.
func NewWavefunction() *Wavefunction {
	return &Wavefunction{level: NewSliderRow("Energy level", 0, func(v float32) string {
		return fmt.Sprintf("n = %d", boxLevel(v))
	})}
}

func boxLevel(v float32) int { return 1 + int(math.Round(float64(v)*(boxLevels-1))) }

// Layout draws ψ and |ψ|², the readout and the slider.
func (m *Wavefunction) Layout(gtx giode.Context) giode.Dimensions {
	n := boxLevel(m.level.Value())
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(250, func(p *giode.Painter) { m.draw(p, n) }),
		readout(fmt.Sprintf("E = %d × E₁ · %d nodes inside the box, where it is never found", n*n, n-1)),
		m.level,
	).Layout(gtx)
}

func (m *Wavefunction) draw(p *giode.Painter, n int) {
	half := p.Height() / 2
	top := plot{left: 24, top: 10, width: p.Width() - 48, height: half - 20, xMin: 0, xMax: 1, yMin: -1.6, yMax: 1.6}
	bottom := plot{left: 24, top: half + 8, width: p.Width() - 48, height: half - 18, xMin: 0, xMax: 1, yMin: 0, yMax: 2.1}
	for _, x := range []float32{top.left, top.left + top.width} {
		p.Rect(x-3, 6, 6, p.Height()-12, 2, theme.Faint)
	}
	top.axes(p)
	top.curve(p, func(x float64) float64 { return physics.BoxWave(n, x) }, 2.5, theme.Primary)
	area := p.Path()
	area.MoveTo(bottom.X(0), bottom.Y(0))
	for i := 0; i <= 120; i++ {
		x := float64(i) / 120
		w := physics.BoxWave(n, x)
		area.LineTo(bottom.X(x), bottom.Y(w*w))
	}
	area.LineTo(bottom.X(1), bottom.Y(0))
	area.Close()
	area.Fill(theme.AccentSoft)
	bottom.curve(p, func(x float64) float64 { w := physics.BoxWave(n, x); return w * w }, 2, theme.Accent)
	label(p, top.left+6, top.top+2, "ψ", giode.TextAnchorTopLeft, theme.Primary)
	label(p, bottom.left+6, bottom.top+2, "|ψ|²: probability", giode.TextAnchorTopLeft, theme.Accent)
}
