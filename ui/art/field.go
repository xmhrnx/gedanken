package art

import (
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
)

// fieldMotif is the field lines between two charges: electricity and
// magnetism.
type fieldMotif struct {
	lines   [][2]float64
	starts  []int
	charges []physics.Charge
}

// Field coordinates: x in [-2, 2], y in [-0.6, 0.6] map onto the canvas.
const fieldW, fieldH = 4.0, 1.2

func newFieldMotif(seed int) *fieldMotif {
	r := rng(4, seed)
	sep := 0.6 + 0.4*r.Float64()
	q2 := -1.0
	if seed%3 == 2 {
		q2 = 1
	}
	m := &fieldMotif{charges: []physics.Charge{{X: -sep, Q: 1}, {X: sep, Q: q2}}}
	for _, c := range m.charges {
		for i := 0; i < 14; i++ {
			a := 2 * math.Pi * (float64(i) + 0.5) / 14
			line := physics.FieldLine(c.X+0.03*math.Cos(a), 0.03*math.Sin(a), m.charges, c.Q < 0, 2.2)
			m.starts = append(m.starts, len(m.lines))
			for _, pt := range line {
				m.lines = append(m.lines, [2]float64{pt.X, pt.Y})
			}
		}
	}
	m.starts = append(m.starts, len(m.lines))
	return m
}

func (m *fieldMotif) draw(p *giode.Painter, hue string) {
	w, h := p.Width(), p.Height()
	at := func(x, y float64) (float32, float32) {
		return float32((x/fieldW + 0.5)) * w, float32((0.5 - y/fieldH)) * h
	}
	for k := 0; k+1 < len(m.starts); k++ {
		path := p.Path()
		for i := m.starts[k]; i < m.starts[k+1]; i++ {
			x, y := at(m.lines[i][0], m.lines[i][1])
			if i == m.starts[k] {
				path.MoveTo(x, y)
			} else {
				path.LineTo(x, y)
			}
		}
		path.Stroke(1.2, tint(hue, 0.45+0.25*float64(k%3)/2))
	}
	for _, c := range m.charges {
		x, y := at(c.X, c.Y)
		core := "#ff8a7a"
		if c.Q < 0 {
			core = "#7fd6ff"
		}
		for i := 4; i >= 1; i-- {
			p.Circle(x, y, float32(5+3*i), tint(core, 0.12*float64(5-i)))
		}
		p.Circle(x, y, 6, giode.Color(core))
	}
}
