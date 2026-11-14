package models

import (
	"fmt"
	"math"
	"time"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// LightClock compares a clock at rest with one moving sideways. Light
// in the moving clock travels a longer diagonal at the same speed, so
// each tick takes γ times longer.
type LightClock struct {
	speed *SliderRow
	start time.Time
}

const lightTick = 1.4 // seconds per tick of the rest clock, on screen

// NewLightClock creates the model.
func NewLightClock() *LightClock {
	return &LightClock{start: time.Now(), speed: NewSliderRow("Speed of the moving clock", 0.6, func(v float32) string {
		return fmt.Sprintf("%.2f c", clockBeta(v))
	})}
}

func clockBeta(v float32) float64 { return lerp(0, 0.95, v) }

// Layout draws both clocks, the readout and the slider.
func (m *LightClock) Layout(gtx giode.Context) giode.Dimensions {
	beta := clockBeta(m.speed.Value())
	g := physics.Gamma(beta)
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(230, func(p *giode.Painter) { m.draw(p, beta, g) }),
		readout(fmt.Sprintf("γ = %.3f · the moving clock ticks %.2f× slower", g, g)),
		m.speed,
	).Layout(gtx)
}

// bounce is the photon's height (0 bottom to 1 top) after t seconds
// for a tick of the given period.
func bounce(t, period float64) float32 {
	f := math.Mod(t/period, 1)
	return float32(1 - math.Abs(1-2*f))
}

func (m *LightClock) draw(p *giode.Painter, beta, g float64) {
	t := sinceSeconds(m.start)
	top, bottom := float32(34), p.Height()-26
	h := bottom - top
	restX := float32(40)
	mirror := func(x float32) {
		p.Line(x-14, top, x+14, top, 3, theme.Text)
		p.Line(x-14, bottom, x+14, bottom, 3, theme.Text)
	}
	mirror(restX)
	p.Circle(restX, bottom-bounce(t, lightTick)*h, 5, theme.Accent)
	label(p, restX, 12, "at rest", giode.TextAnchorCenter, theme.Muted)

	left, right := float32(100), p.Width()-20
	span := right - left
	// On screen light covers h per half tick; the moving clock covers
	// β times that distance sideways.
	cScreen := float64(h) / (lightTick / 2)
	period := lightTick * g
	travel := math.Mod(t*beta*cScreen, float64(span))
	x := left + float32(travel)
	mirror(x)
	const trail = 60
	path := p.Path()
	started := false
	for i := 0; i <= trail; i++ {
		tt := t - float64(trail-i)*0.02
		tx := float32(float64(left) + math.Mod(tt*beta*cScreen, float64(span)))
		if tx > x {
			continue // before the clock last wrapped around
		}
		ty := bottom - bounce(tt, period)*h
		if !started {
			path.MoveTo(tx, ty)
			started = true
		} else {
			path.LineTo(tx, ty)
		}
	}
	path.Stroke(1.5, theme.PrimarySoft)
	p.Circle(x, bottom-bounce(t, period)*h, 5, theme.Accent)
	label(p, (left+right)/2, 12, "moving →", giode.TextAnchorCenter, theme.Muted)
	keepAnimating(p)
}
