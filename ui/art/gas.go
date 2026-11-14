package art

import (
	"math"

	"github.com/xmhrnx/giode"
)

// gasMotif is molecules with motion streaks over faint isotherms:
// matter and heat.
type gasMotif struct {
	molecules [][4]float32
	isotherms [][][2]float32
}

func newGasMotif(seed int) *gasMotif {
	r := rng(3, seed)
	m := &gasMotif{}
	for i := 0; i < 70; i++ {
		a := r.Float64() * 2 * math.Pi
		l := 0.01 + 0.04*r.Float64()
		x, y := r.Float64(), r.Float64()
		m.molecules = append(m.molecules, [4]float32{float32(x), float32(y), float32(l * math.Cos(a)), float32(l * math.Sin(a) * 3)})
	}
	for k := 1; k <= 3; k++ {
		var curve [][2]float32
		for i := 0; i <= 60; i++ {
			v := 0.08 + 0.92*float64(i)/60
			curve = append(curve, [2]float32{float32(v), float32(1 - 0.06*float64(k)/v)})
		}
		m.isotherms = append(m.isotherms, curve)
	}
	return m
}

func (m *gasMotif) draw(p *giode.Painter, hue string) {
	w, h := p.Width(), p.Height()
	for i, c := range m.isotherms {
		path := p.Path()
		strokePoints(path, c, w, h)
		path.Stroke(1.5, tint(hue, 0.25+0.12*float64(i)))
	}
	for i, mol := range m.molecules {
		x, y := mol[0]*w, mol[1]*h
		p.Line(x, y, x-mol[2]*w, y-mol[3]*h, 1.5, tint(hue, 0.3))
		p.Circle(x, y, float32(2+i%3), tint(hue, 0.6+0.4*float64(i%4)/3))
	}
}
