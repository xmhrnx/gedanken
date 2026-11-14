package models

import (
	"fmt"
	"math"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// Decay shows a sample of radioactive nuclei. Each decays at a random
// moment, yet the count halves every half-life.
type Decay struct {
	start     *giode.ButtonElement
	sample    physics.DecaySample
	startedAt time.Time
	seed      int64
}

const (
	decayCols, decayRows = 20, 10
	decayHalfLife        = 1.6 // seconds on screen
	decayMax             = 5.0 // half-lives shown
)

// NewDecay creates the model.
func NewDecay() *Decay {
	m := &Decay{seed: 1, sample: physics.NewDecaySample(decayCols*decayRows, 1)}
	m.start = giode.Button("Start", theme.PrimaryButton).OnClick(m.restart)
	return m
}

func (m *Decay) restart() {
	m.seed++
	m.sample = physics.NewDecaySample(decayCols*decayRows, m.seed)
	m.startedAt = time.Now()
	m.start.Label("Restart")
}

// elapsed is the time since the start in half-lives.
func (m *Decay) elapsed() float64 {
	if m.startedAt.IsZero() {
		return 0
	}
	return math.Min(decayMax, sinceSeconds(m.startedAt)/decayHalfLife)
}

// Layout draws the nuclei and graph, readout and button.
func (m *Decay) Layout(gtx giode.Context) giode.Dimensions {
	t := m.elapsed()
	total := float64(decayCols * decayRows)
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(250, func(p *giode.Painter) { m.draw(p, t) }),
		readout(fmt.Sprintf("%.1f half-lives · %d remain · expected %.0f", t, m.sample.Remaining(t), total/math.Pow(2, t))),
		m.start,
	).Layout(gtx)
}

func (m *Decay) draw(p *giode.Painter, t float64) {
	gridH := p.Height() * 0.55
	cw, ch := p.Width()/decayCols, gridH/decayRows
	for i, d := range m.sample.Times {
		x, y := (float32(i%decayCols)+0.5)*cw, (float32(i/decayCols)+0.5)*ch
		if d > t {
			p.Circle(x, y, 4, theme.Accent)
		} else {
			p.Circle(x, y, 2, "#34406a")
		}
	}
	total := float64(decayCols * decayRows)
	g := plot{left: 30, top: gridH + 12, width: p.Width() - 44, height: p.Height() - gridH - 30,
		xMin: 0, xMax: decayMax, yMin: 0, yMax: total}
	g.axes(p)
	for h := 1.0; h <= decayMax; h++ {
		dashed(p, g.X(h), g.Y(0), g.X(h), g.Y(total), 1, "#1d2741")
	}
	g.curve(p, func(x float64) float64 { return total / math.Pow(2, x) }, 1.5, theme.Muted)
	if t > 0 {
		path := p.Path()
		for i := 0; i <= 100; i++ {
			x := t * float64(i) / 100
			if i == 0 {
				path.MoveTo(g.X(x), g.Y(float64(m.sample.Remaining(x))))
			} else {
				path.LineTo(g.X(x), g.Y(float64(m.sample.Remaining(x))))
			}
		}
		path.Stroke(2, theme.Accent)
	}
	label(p, g.left+g.width, g.top+g.height+4, "half-lives", giode.TextAnchorTopRight, theme.Muted)
	if !m.startedAt.IsZero() && t < decayMax {
		keepAnimating(p)
	}
}
