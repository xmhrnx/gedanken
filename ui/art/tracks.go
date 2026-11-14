package art

import (
	"math"

	"github.com/xmhrnx/giode"
)

// tracksMotif is a nucleus with particle tracks curling away, as in a
// bubble chamber: atoms and nuclei.
type tracksMotif struct {
	nucleus  [2]float32
	nucleons [][3]float32
	tracks   [][][2]float32
}

func newTracksMotif(seed int) *tracksMotif {
	r := rng(7, seed)
	m := &tracksMotif{nucleus: [2]float32{0.3 + 0.4*float32(r.Float64()), 0.5}}
	for i := 0; i < 22; i++ {
		a, d := r.Float64()*2*math.Pi, 0.035*math.Sqrt(r.Float64())
		m.nucleons = append(m.nucleons, [3]float32{float32(d * math.Cos(a)), float32(d * math.Sin(a) * 3.2), float32(i % 2)})
	}
	for k := 0; k < 7; k++ {
		a, curl := r.Float64()*2*math.Pi, (r.Float64()-0.5)*14
		var track [][2]float32
		x, y := float64(m.nucleus[0]), float64(m.nucleus[1])
		for i := 0; i < 70; i++ {
			a += curl * 0.004 * float64(i) / 20
			x += 0.008 * math.Cos(a)
			y += 0.008 * math.Sin(a) * 3.2
			track = append(track, [2]float32{float32(x), float32(y)})
		}
		m.tracks = append(m.tracks, track)
	}
	return m
}

func (m *tracksMotif) draw(p *giode.Painter, hue string) {
	w, h := p.Width(), p.Height()
	for i, t := range m.tracks {
		path := p.Path()
		strokePoints(path, t, w, h)
		path.Stroke(1.4, tint(hue, 0.4+0.1*float64(i%4)))
	}
	nx, ny := m.nucleus[0]*w, m.nucleus[1]*h
	for i := 5; i >= 1; i-- {
		p.Circle(nx, ny, float32(10+4*i), tint(hue, 0.08*float64(6-i)))
	}
	for _, n := range m.nucleons {
		c := giode.Color("#ff8a7a")
		if n[2] == 1 {
			c = "#9db4ff"
		}
		p.Circle(nx+n[0]*w, ny+n[1]*h, 5, c)
	}
}
