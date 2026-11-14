package art

import "math/rand"

// rng returns a random source fixed by part and chapter, so a chapter's
// illustration is the same every time.
func rng(salt, seed int) *rand.Rand {
	return rand.New(rand.NewSource(int64(salt*1000 + seed)))
}

// stroke draws a polyline through normalized points (0–1) scaled to the
// canvas.
func strokePoints(path interface {
	MoveTo(x, y float32)
	LineTo(x, y float32)
}, pts [][2]float32, w, h float32) {
	for i, pt := range pts {
		if i == 0 {
			path.MoveTo(pt[0]*w, pt[1]*h)
		} else {
			path.LineTo(pt[0]*w, pt[1]*h)
		}
	}
}
