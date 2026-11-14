package models

import (
	"fmt"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// ChainReaction animates generations of fission on a grid of nuclei.
// The slider sets k, how many neutrons per fission split another nucleus.
type ChainReaction struct {
	k         *SliderRow
	start     *giode.ButtonElement
	reaction  *physics.ChainReaction
	startedAt time.Time
	seed      int64
	ended     bool
}

const (
	chainCols, chainRows = 18, 11
	chainGeneration      = 0.7 // seconds per generation on screen
	chainMaxGenerations  = 12
)

// NewChainReaction creates the model.
func NewChainReaction() *ChainReaction {
	m := &ChainReaction{k: NewSliderRow("k (neutrons that cause another fission)", 0.6, func(v float32) string {
		return fmt.Sprintf("%.2f", chainK(v))
	})}
	m.start = giode.Button("Start", theme.PrimaryButton).OnClick(m.restart)
	m.reaction = physics.NewChainReaction(chainCols, chainRows, chainK(m.k.Value()), m.seed)
	return m
}

func chainK(v float32) float64 { return lerp(0.6, 2.6, v) }

// restart begins a new reaction with the current k.
func (m *ChainReaction) restart() {
	m.seed++
	m.reaction = physics.NewChainReaction(chainCols, chainRows, chainK(m.k.Value()), m.seed)
	m.startedAt = time.Now()
	m.ended = false
	m.start.Label("Restart")
}

// advance runs the generations whose time has come.
func (m *ChainReaction) advance() {
	if m.startedAt.IsZero() || m.ended {
		return
	}
	due := 1 + int(sinceSeconds(m.startedAt)/chainGeneration)
	for m.reaction.Generations() < min(due, chainMaxGenerations) {
		if m.reaction.Step() == 0 {
			m.ended = true
			return
		}
	}
	m.ended = m.reaction.Generations() >= chainMaxGenerations
}

// Layout draws the grid, readout, slider and button.
func (m *ChainReaction) Layout(gtx giode.Context) giode.Dimensions {
	m.advance()
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(220, m.draw),
		readout(m.summary()),
		m.k, m.start,
	).Layout(gtx)
}

// summary counts fissions in the latest generation and overall.
func (m *ChainReaction) summary() string {
	last := m.reaction.Generations() - 1
	latest := 0
	for _, f := range m.reaction.Fissions {
		if f.Generation == last {
			latest++
		}
	}
	return fmt.Sprintf("generation %d · %d new · %d total", last+1, latest, len(m.reaction.Fissions))
}

func (m *ChainReaction) draw(p *giode.Painter) {
	cellW, cellH := p.Width()/chainCols, p.Height()/chainRows
	pos := func(n int) (float32, float32) {
		return (float32(n%chainCols) + 0.5) * cellW, (float32(n/chainCols) + 0.5) * cellH
	}
	for n := 0; n < chainCols*chainRows; n++ {
		if !m.reaction.Split(n) {
			x, y := pos(n)
			p.Circle(x, y, 3.5, "#34406a")
		}
	}
	last := m.reaction.Generations() - 1
	for _, f := range m.reaction.Fissions {
		x, y := pos(f.Nucleus)
		if f.Parent >= 0 {
			px, py := pos(f.Parent)
			c := giode.Color("#2a3a66")
			if f.Generation == last {
				c = theme.Info
			}
			p.Line(px, py, x, y, 1.2, c)
		}
		age := float64(last-f.Generation) / 6
		c := lerpColor("#ffd36e", "#a8443a", age)
		p.Circle(x-2.5, y, 3, c)
		p.Circle(x+2.5, y, 3, c)
	}
	if m.startedAt.IsZero() {
		label(p, p.Width()/2, p.Height()/2, "press Start", giode.TextAnchorCenter, theme.Text)
	}
	if !m.startedAt.IsZero() && !m.ended {
		keepAnimating(p)
	}
}
