package art

import (
	"math"

	"github.com/xmhrnx/giode"
)

// wavePacketMotif is a wave packet, its envelope and its probability:
// quantum physics.
type wavePacketMotif struct {
	wave, envelope, prob [][2]float32
}

func newWavePacketMotif(seed int) *wavePacketMotif {
	r := rng(6, seed)
	centre, width, k := 0.4+0.2*r.Float64(), 0.1+0.06*r.Float64(), 40+30*r.Float64()
	m := &wavePacketMotif{}
	for i := 0; i <= 200; i++ {
		x := float64(i) / 200
		env := math.Exp(-(x - centre) * (x - centre) / (2 * width * width))
		m.wave = append(m.wave, [2]float32{float32(x), float32(0.45 - 0.33*env*math.Cos(k*x))})
		m.envelope = append(m.envelope, [2]float32{float32(x), float32(0.45 - 0.33*env)})
		m.prob = append(m.prob, [2]float32{float32(x), float32(0.98 - 0.3*env*env)})
	}
	return m
}

func (m *wavePacketMotif) draw(p *giode.Painter, hue string) {
	w, h := p.Width(), p.Height()
	area := p.Path()
	area.MoveTo(0, h)
	for _, pt := range m.prob {
		area.LineTo(pt[0]*w, pt[1]*h)
	}
	area.LineTo(w, h)
	area.Close()
	area.Fill(tint(hue, 0.3))
	env := p.Path()
	strokePoints(env, m.envelope, w, h)
	env.Stroke(1, tint(hue, 0.4))
	wave := p.Path()
	strokePoints(wave, m.wave, w, h)
	wave.Stroke(2.2, tint(hue, 1))
}
