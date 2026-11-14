// Package models holds the interactive graphics shown in lessons.
// Each model is a stateful element created once per lesson block.
package models

import (
	"fmt"
	"math"
	"time"

	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/components/kit"
	"github.com/xmhrnx/giode/styles/properties"

	"gedanken/ui/theme"
)

// plot maps data coordinates onto a rectangle of the canvas.
type plot struct {
	left, top, width, height float32
	xMin, xMax, yMin, yMax   float64
}

// stageColor is the backdrop of every model.
var stageGradient = giode.Linear(160, "#141c33", "#0b0f1c")

// stage frames a canvas of the given height on the dark backdrop.
func stage(height int, draw func(p *giode.Painter)) giode.Element {
	return giode.Box(giode.Styles{
		Width: -1, BorderRadius: theme.SmallRadius, Overflow: giode.OverflowHidden,
		BackgroundGradient: stageGradient,
	}, giode.Canvas(giode.Styles{Width: -1, Height: properties.Height(height)}, draw))
}

// newPlot fits a data range into the canvas with a margin.
func newPlot(p *giode.Painter, margin float32, xMin, xMax, yMin, yMax float64) plot {
	return plot{margin, margin, p.Width() - 2*margin, p.Height() - 2*margin, xMin, xMax, yMin, yMax}
}

// X maps a data x to canvas dp.
func (pl plot) X(x float64) float32 {
	return pl.left + float32((x-pl.xMin)/(pl.xMax-pl.xMin))*pl.width
}

// Y maps a data y to canvas dp, upward.
func (pl plot) Y(y float64) float32 {
	return pl.top + pl.height - float32((y-pl.yMin)/(pl.yMax-pl.yMin))*pl.height
}

// axes draws the x and y axes through data zero, clamped to the plot.
func (pl plot) axes(p *giode.Painter) {
	x0 := pl.X(math.Max(pl.xMin, math.Min(0, pl.xMax)))
	y0 := pl.Y(math.Max(pl.yMin, math.Min(0, pl.yMax)))
	p.Line(pl.left, y0, pl.left+pl.width, y0, 1, theme.Faint)
	p.Line(x0, pl.top, x0, pl.top+pl.height, 1, theme.Faint)
}

// curve strokes f sampled across the plot's x range.
func (pl plot) curve(p *giode.Painter, f func(float64) float64, width float32, c giode.Color) {
	const samples = 120
	path := p.Path()
	for i := 0; i <= samples; i++ {
		x := pl.xMin + (pl.xMax-pl.xMin)*float64(i)/samples
		if i == 0 {
			path.MoveTo(pl.X(x), pl.Y(f(x)))
		} else {
			path.LineTo(pl.X(x), pl.Y(f(x)))
		}
	}
	path.Stroke(width, c)
}

// dashed draws a dashed straight line.
func dashed(p *giode.Painter, x1, y1, x2, y2, width float32, c giode.Color) {
	const dash, gap = 6, 5
	length := float32(math.Hypot(float64(x2-x1), float64(y2-y1)))
	if length == 0 {
		return
	}
	dx, dy := (x2-x1)/length, (y2-y1)/length
	for s := float32(0); s < length; s += dash + gap {
		e := min(s+dash, length)
		p.Line(x1+dx*s, y1+dy*s, x1+dx*e, y1+dy*e, width, c)
	}
}

// arrow draws a line with a head at (x2, y2).
func arrow(p *giode.Painter, x1, y1, x2, y2, width float32, c giode.Color) {
	p.Line(x1, y1, x2, y2, width, c)
	a := math.Atan2(float64(y2-y1), float64(x2-x1))
	const head = 9
	for _, side := range []float64{-0.45, 0.45} {
		hx := x2 - head*float32(math.Cos(a+side))
		hy := y2 - head*float32(math.Sin(a+side))
		p.Line(x2, y2, hx, hy, width, c)
	}
}

// glow paints soft concentric rings fading into the backdrop.
func glow(p *giode.Painter, cx, cy, r float32, c giode.Color) {
	const rings = 6
	for i := rings; i >= 1; i-- {
		t := float64(rings-i+1) / rings * 0.35
		p.Circle(cx, cy, r+float32(i)*min(6, r*0.25), giode.Color(giode.Mix(string(c), "#0f1528", t)))
	}
}

// lerpColor blends two "#rrggbb" colors.
func lerpColor(a, b string, t float64) giode.Color {
	t = math.Max(0, math.Min(1, t))
	var ar, ag, ab, br, bg, bb int
	fmt.Sscanf(a, "#%02x%02x%02x", &ar, &ag, &ab)
	fmt.Sscanf(b, "#%02x%02x%02x", &br, &bg, &bb)
	mix := func(x, y int) int { return int(float64(x) + (float64(y)-float64(x))*t) }
	return giode.Color(fmt.Sprintf("#%02x%02x%02x", mix(ar, br), mix(ag, bg), mix(ab, bb)))
}

// keepAnimating asks for another frame while an animation runs.
func keepAnimating(p *giode.Painter) {
	kit.RequestFrame(p.Context())
}

// readout is a line of live numbers under a model.
func readout(text string) giode.Element {
	return giode.Text(text, theme.Readout)
}

// sinceSeconds returns the seconds elapsed since t.
func sinceSeconds(t time.Time) float64 {
	return time.Since(t).Seconds()
}

// lerp maps t in [0, 1] onto [a, b].
func lerp(a, b float64, t float32) float64 {
	return a + (b-a)*float64(t)
}

// label draws a small caption on a canvas.
func label(p *giode.Painter, x, y float32, text string, anchor giode.TextAnchor, color giode.Color) {
	p.Text(x, y, text, anchor, giode.Styles{FontSize: 11, Color: color})
}
