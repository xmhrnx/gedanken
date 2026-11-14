package art

import (
	"math"

	"github.com/xmhrnx/giode"
)

// galaxyMotif is a two-armed spiral galaxy in a field of stars:
// particles and the cosmos.
type galaxyMotif struct {
	centre [2]float32
	stars  [][3]float32
	field  [][3]float32
}

func newGalaxyMotif(seed int) *galaxyMotif {
	r := rng(8, seed)
	m := &galaxyMotif{centre: [2]float32{0.55 + 0.2*float32(r.Float64()), 0.5}}
	twist := 0.5 + 0.3*r.Float64()
	for i := 0; i < 600; i++ {
		arm := float64(i%2) * math.Pi
		t := r.Float64() * 4
		rad := 0.02 * math.Exp(twist*t)
		a := arm + t + (r.Float64()-0.5)*0.6
		spread := (r.Float64() - 0.5) * 0.04 * (1 + t)
		x := (rad + spread) * math.Cos(a)
		y := (rad + spread) * math.Sin(a) * 1.9
		m.stars = append(m.stars, [3]float32{float32(x), float32(y), float32(r.Float64())})
	}
	for i := 0; i < 70; i++ {
		m.field = append(m.field, [3]float32{float32(r.Float64()), float32(r.Float64()), float32(r.Float64())})
	}
	return m
}

func (m *galaxyMotif) draw(p *giode.Painter, hue string) {
	w, h := p.Width(), p.Height()
	for _, s := range m.field {
		p.Circle(s[0]*w, s[1]*h, 0.6+s[2], tint("#ffffff", 0.2+0.5*float64(s[2])))
	}
	cx, cy := m.centre[0]*w, m.centre[1]*h
	for i := 6; i >= 1; i-- {
		p.Circle(cx, cy, float32(4+5*i), tint(hue, 0.07*float64(7-i)))
	}
	for _, s := range m.stars {
		x, y := cx+s[0]*w*0.7, cy+s[1]*h*0.62
		p.Circle(x, y, 0.8+s[2]*0.8, tint(hue, 0.35+0.65*float64(s[2])))
	}
	p.Circle(cx, cy, 5, "#fff3c4")
}
