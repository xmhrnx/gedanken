package art

import (
	"math"

	"github.com/xmhrnx/giode"
)

// graphMotif is graph paper with a smooth curve, its tangent and a few
// vectors: the language of physics.
type graphMotif struct {
	curve   [][2]float32
	touch   [2]float32
	slope   float32
	vectors [][4]float32
}

func newGraphMotif(seed int) *graphMotif {
	r := rng(1, seed)
	a, b, phase := 0.16+0.08*r.Float64(), 1.5+r.Float64(), r.Float64()*6
	f := func(x float64) float64 { return 0.55 - a*math.Sin(b*x*math.Pi+phase) - 0.12*x }
	m := &graphMotif{}
	for i := 0; i <= 80; i++ {
		x := float64(i) / 80
		m.curve = append(m.curve, [2]float32{float32(x), float32(f(x))})
	}
	tx := 0.35 + 0.3*r.Float64()
	m.touch = [2]float32{float32(tx), float32(f(tx))}
	m.slope = float32((f(tx+0.001) - f(tx-0.001)) / 0.002)
	for i := 0; i < 3; i++ {
		ang := r.Float64() * 2 * math.Pi
		m.vectors = append(m.vectors, [4]float32{0.12, 0.8, 0.12 + float32(0.12*math.Cos(ang)), 0.8 - float32(0.35*math.Sin(ang))})
	}
	return m
}

func (m *graphMotif) draw(p *giode.Painter, hue string) {
	w, h := p.Width(), p.Height()
	for x := float32(0); x < w; x += 16 {
		p.Line(x, 0, x, h, 1, tint(hue, 0.09))
	}
	for y := float32(0); y < h; y += 16 {
		p.Line(0, y, w, y, 1, tint(hue, 0.09))
	}
	path := p.Path()
	strokePoints(path, m.curve, w, h)
	path.Stroke(2.5, tint(hue, 0.95))
	tx, ty := m.touch[0]*w, m.touch[1]*h
	dx := float32(70)
	dy := m.slope * dx * h / w
	p.Line(tx-dx, ty-dy, tx+dx, ty+dy, 1.5, tint(hue, 0.55))
	p.Circle(tx, ty, 4, "#ffffff")
	for _, v := range m.vectors {
		p.Line(v[0]*w, v[1]*h, v[2]*w, v[3]*h, 2, tint(hue, 0.7))
		p.Circle(v[2]*w, v[3]*h, 2.5, tint(hue, 0.7))
	}
}
