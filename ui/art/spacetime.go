package art

import (
	"math"

	"github.com/xmhrnx/giode"
)

// spacetimeMotif is a grid pulled in by a mass, with a light ray
// bending past it: relativity.
type spacetimeMotif struct {
	mass  [2]float64
	depth float64
	ray   [][2]float32
}

func newSpacetimeMotif(seed int) *spacetimeMotif {
	r := rng(5, seed)
	m := &spacetimeMotif{mass: [2]float64{0.45 + 0.25*r.Float64(), 0.55}, depth: 0.05 + 0.03*r.Float64()}
	for i := 0; i <= 60; i++ {
		x := float64(i) / 60
		dx := x - m.mass[0]
		y := 0.18 + 0.12*math.Exp(-dx*dx*40)
		m.ray = append(m.ray, [2]float32{float32(x), float32(y)})
	}
	return m
}

// warp pulls a point toward the mass, more strongly nearby.
func (m *spacetimeMotif) warp(x, y, aspect float64) (float64, float64) {
	dx, dy := x-m.mass[0], (y-m.mass[1])/aspect
	d := math.Hypot(dx, dy) + 0.05
	pull := m.depth / d
	pull = math.Min(pull, d*0.9)
	return x - pull*dx/d, y - pull*dy/d*aspect
}

func (m *spacetimeMotif) draw(p *giode.Painter, hue string) {
	w, h := p.Width(), p.Height()
	aspect := float64(w / h)
	line := func(x0, y0, x1, y1 float64, c giode.Color) {
		path := p.Path()
		for i := 0; i <= 50; i++ {
			t := float64(i) / 50
			x, y := m.warp(x0+(x1-x0)*t, y0+(y1-y0)*t, aspect)
			if i == 0 {
				path.MoveTo(float32(x)*w, float32(y)*h)
			} else {
				path.LineTo(float32(x)*w, float32(y)*h)
			}
		}
		path.Stroke(1, c)
	}
	for k := 0; k <= 22; k++ {
		line(float64(k)/22, 0, float64(k)/22, 1, tint(hue, 0.35))
	}
	for k := 0; k <= 7; k++ {
		line(0, float64(k)/7, 1, float64(k)/7, tint(hue, 0.35))
	}
	mx, my := float32(m.mass[0])*w, float32(m.mass[1])*h
	for i := 5; i >= 1; i-- {
		p.Circle(mx, my, float32(3+3*i), tint(hue, 0.1*float64(6-i)))
	}
	p.Circle(mx, my, 7, "#000000")
	path := p.Path()
	strokePoints(path, m.ray, w, h)
	path.Stroke(2, "#fff3c4")
}
