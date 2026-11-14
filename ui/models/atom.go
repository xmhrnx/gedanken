package models

import (
	"math"
	"time"

	"github.com/xmhrnx/giode"
)

// Atom is the historical planetary picture of the atom: a nucleus with
// electrons on orbits, as Rutherford imagined it in 1911. It is labelled
// as such; OrbitalCloud shows the quantum-mechanical picture.
type Atom struct {
	start  time.Time
	height int
}

// orbit is one electron's path.
type orbit struct {
	tilt, speed, phase float64
}

var atomOrbits = []orbit{{0, 1.3, 0}, {math.Pi / 3, 1.0, 2}, {2 * math.Pi / 3, 0.8, 4}}

// nucleons are offsets (in nucleus radii) and whether each is a proton.
var nucleons = []struct {
	x, y   float32
	proton bool
}{{0, 0, true}, {0.9, 0.2, false}, {-0.8, 0.4, true}, {0.2, -0.9, false}, {-0.3, 0.9, false}, {0.7, -0.6, true}, {-0.7, -0.5, false}}

// NewAtom creates an atom drawn at the given height.
func NewAtom(height int) *Atom {
	return &Atom{start: time.Now(), height: height}
}

// Layout draws the current frame and requests the next.
func (a *Atom) Layout(gtx giode.Context) giode.Dimensions {
	return giode.Canvas(giode.Styles{Width: -1, Height: giodeHeight(a.height)}, a.draw).Layout(gtx)
}

func (a *Atom) draw(p *giode.Painter) {
	t := sinceSeconds(a.start)
	cx, cy := p.Width()/2, p.Height()/2
	r := float64(min(p.Width(), p.Height())) * 0.44
	glow(p, cx, cy, 16, "#8ea8ff")
	for _, o := range atomOrbits {
		ellipse(p, cx, cy, r, r*0.34, o.tilt, 1.5, "#2f3d63")
	}
	for _, n := range nucleons {
		c := giode.Color("#8ea8ff")
		if n.proton {
			c = "#f58a8a"
		}
		p.Circle(cx+n.x*7, cy+n.y*7, 6, c)
	}
	for _, o := range atomOrbits {
		x, y := ellipsePoint(cx, cy, r, r*0.34, o.tilt, t*o.speed+o.phase)
		p.Circle(x, y, 9, "#1d2d55")
		p.Circle(x, y, 5, "#6ee7f0")
	}
	label(p, p.Width()/2, p.Height()-8, "Rutherford's 1911 model · not to scale", giode.TextAnchorBottom, "#8d96ab")
	keepAnimating(p)
}

// ellipse strokes a rotated ellipse.
func ellipse(p *giode.Painter, cx, cy float32, a, b, tilt float64, width float32, c giode.Color) {
	const n = 72
	path := p.Path()
	for i := 0; i <= n; i++ {
		x, y := ellipsePoint(cx, cy, a, b, tilt, 2*math.Pi*float64(i)/n)
		if i == 0 {
			path.MoveTo(x, y)
		} else {
			path.LineTo(x, y)
		}
	}
	path.Stroke(width, c)
}

// ellipsePoint is the point at angle th on an ellipse rotated by tilt.
func ellipsePoint(cx, cy float32, a, b, tilt, th float64) (float32, float32) {
	x := a*math.Cos(th)*math.Cos(tilt) - b*math.Sin(th)*math.Sin(tilt)
	y := a*math.Cos(th)*math.Sin(tilt) + b*math.Sin(th)*math.Cos(tilt)
	return cx + float32(x), cy + float32(y)
}
