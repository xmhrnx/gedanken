package models

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/components/segmented"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// OrbitalCloud draws a hydrogen orbital as quantum mechanics describes
// it: a cloud of places the electron may be found, distributed as |ψ|²,
// coloured by the sign of ψ, slowly turning. On the home screen it
// cycles through orbitals; in a lesson the reader picks one.
type OrbitalCloud struct {
	height int
	picker *segmented.Control
	clouds [][]physics.CloudPoint
	reach  []float64
	shades map[bool][]giode.Color
	start  time.Time
}

const (
	cloudPoints = 1100
	cloudShades = 12
	cloudCycle  = 7.0 // seconds per orbital when cycling
	cloudFade   = 1.0 // seconds of cross-fade
	cloudFirst  = 2   // start the cycle on 2p, the most recognisable shape
)

// NewOrbitalCloud creates the model; interactive adds an orbital picker.
func NewOrbitalCloud(height int, interactive bool) *OrbitalCloud {
	m := &OrbitalCloud{height: height, start: time.Now(), shades: map[bool][]giode.Color{}}
	for i, o := range physics.Orbitals {
		pts := physics.SampleOrbital(o, cloudPoints, int64(i+1))
		m.clouds = append(m.clouds, pts)
		m.reach = append(m.reach, percentileRadius(pts, 0.94))
	}
	for sign, c := range map[bool]string{true: "#8db8ff", false: "#ffb070"} {
		for i := 0; i < cloudShades; i++ {
			m.shades[sign] = append(m.shades[sign], lerpColor("#0b0f1c", c, 0.25+0.75*float64(i)/(cloudShades-1)))
		}
	}
	if interactive {
		var names []string
		for _, o := range physics.Orbitals {
			names = append(names, o.Name)
		}
		m.picker = giode.Input.Segmented(names...)
		m.picker.SetSelected(2)
		m.picker.Styles(giode.Styles{Width: -1})
	}
	return m
}

// percentileRadius is the radius holding the given share of points.
func percentileRadius(pts []physics.CloudPoint, share float64) float64 {
	rs := make([]float64, len(pts))
	for i, p := range pts {
		rs[i] = math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
	}
	sort.Float64s(rs)
	return rs[int(share*float64(len(rs)-1))]
}

// showing returns the orbital to draw and its fade (0–1).
func (m *OrbitalCloud) showing(t float64) (int, float64) {
	if m.picker != nil {
		return m.picker.Selected(), 1
	}
	n := len(m.clouds)
	i := (int(t/cloudCycle) + cloudFirst) % n
	into := math.Mod(t, cloudCycle)
	fade := math.Min(1, math.Min(into/cloudFade, (cloudCycle-into)/cloudFade))
	return i, fade
}

// Layout draws the cloud and, in a lesson, the picker.
func (m *OrbitalCloud) Layout(gtx giode.Context) giode.Dimensions {
	canvas := giode.Canvas(giode.Styles{Width: -1, Height: giodeHeight(m.height)}, m.draw)
	if m.picker == nil {
		return canvas.Layout(gtx)
	}
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		giode.Box(giode.Styles{Width: -1, BorderRadius: theme.SmallRadius, Overflow: giode.OverflowHidden,
			BackgroundGradient: stageGradient}, canvas),
		m.picker,
	).Layout(gtx)
}

func (m *OrbitalCloud) draw(p *giode.Painter) {
	t := sinceSeconds(m.start)
	i, fade := m.showing(t)
	cx, cy := p.Width()/2, p.Height()/2-6
	scale := float64(min(p.Width(), p.Height())) * 0.40 / m.reach[i]
	spin, tilt := t*0.3, 0.38
	cs, ss, ct, st := math.Cos(spin), math.Sin(spin), math.Cos(tilt), math.Sin(tilt)
	for _, pt := range m.clouds[i] {
		x := pt.X*cs + pt.Y*ss
		depth := -pt.X*ss + pt.Y*cs
		y := pt.Z*ct - depth*st
		near := (pt.Z*st + depth*ct) / m.reach[i]
		level := (0.55 + 0.45*math.Max(-1, math.Min(1, near))) * fade
		shade := m.shades[pt.Positive][int(level*(cloudShades-1))]
		// A 1.7 dp square reads as a dot and is far cheaper than a circle.
		p.Rect(cx+float32(x*scale)-0.85, cy-float32(y*scale)-0.85, 1.7, 1.7, 0, shade)
	}
	p.Circle(cx, cy, 2.2, "#ffffff")
	o := physics.Orbitals[i]
	label(p, 12, 10, fmt.Sprintf("hydrogen · %s orbital", o.Name), giode.TextAnchorTopLeft, theme.Text)
	label(p, p.Width()/2, p.Height()-8, "each dot: a place the electron may be found · colour: sign of ψ",
		giode.TextAnchorBottom, theme.Muted)
	keepAnimating(p)
}
