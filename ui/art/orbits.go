package art

import (
	"math"

	"github.com/xmhrnx/giode"
)

// orbitMotif is a star with planets on nested ellipses and a thrown
// body's parabola: mechanics and gravity.
type orbitMotif struct {
	star    [2]float32
	ellipse [][][2]float32
	planets [][2]float32
	arc     [][2]float32
}

func newOrbitMotif(seed int) *orbitMotif {
	r := rng(2, seed)
	m := &orbitMotif{star: [2]float32{0.68 + 0.1*float32(r.Float64()), 0.5}}
	tilt := (r.Float64() - 0.5) * 0.5
	for k := 1; k <= 4; k++ {
		a, e := 0.12*float64(k), 0.15+0.2*r.Float64()
		b := a * math.Sqrt(1-e*e)
		var ring [][2]float32
		for i := 0; i <= 90; i++ {
			t := 2 * math.Pi * float64(i) / 90
			x, y := a*math.Cos(t)-a*e, b*math.Sin(t)*2.4
			xr, yr := x*math.Cos(tilt)-y*math.Sin(tilt)*0.42, x*math.Sin(tilt)+y*math.Cos(tilt)*0.42
			ring = append(ring, [2]float32{m.star[0] + float32(xr), m.star[1] + float32(yr)})
		}
		m.ellipse = append(m.ellipse, ring)
		m.planets = append(m.planets, ring[r.Intn(len(ring))])
	}
	for i := 0; i <= 40; i++ {
		x := float64(i) / 40
		m.arc = append(m.arc, [2]float32{float32(0.04 + 0.32*x), float32(0.9 - 2.6*x*(1-x))})
	}
	return m
}

func (m *orbitMotif) draw(p *giode.Painter, hue string) {
	w, h := p.Width(), p.Height()
	for i, ring := range m.ellipse {
		path := p.Path()
		strokePoints(path, ring, w, h)
		path.Stroke(1.3, tint(hue, 0.65-0.1*float64(i)))
	}
	sx, sy := m.star[0]*w, m.star[1]*h
	for i := 5; i >= 1; i-- {
		p.Circle(sx, sy, float32(4+3*i), tint("#ffd36e", 0.12*float64(6-i)))
	}
	p.Circle(sx, sy, 5, "#fff3c4")
	for i, pl := range m.planets {
		p.Circle(pl[0]*w, pl[1]*h, float32(2.5+float64(i%2)), tint(hue, 1))
	}
	arc := p.Path()
	strokePoints(arc, m.arc, w, h)
	arc.Stroke(1.5, tint(hue, 0.4))
}
