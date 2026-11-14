// Package art draws the generative illustrations on chapter cards and
// banners: one motif per part of the course, varied by chapter.
package art

import (
	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/styles/properties"

	"gedanken/ui/theme"
)

// motif draws one illustration into a canvas, in the given hue.
type motif interface {
	draw(p *giode.Painter, hue string)
}

// Banner is an illustration strip for a chapter. Its geometry is
// computed once, when it is created.
type Banner struct {
	motif  motif
	hue    string
	height int
	radii  giode.Corners
}

// NewBanner creates the illustration for chapter seed of part, in the
// part's hue, height dp tall, with the given corner rounding.
func NewBanner(part, seed int, hue giode.Color, height int, radii giode.Corners) *Banner {
	return &Banner{motif: motifFor(part, seed), hue: string(hue), height: height, radii: radii}
}

// motifFor picks the part's motif.
func motifFor(part, seed int) motif {
	switch part {
	case 0:
		return newGraphMotif(seed)
	case 1:
		return newOrbitMotif(seed)
	case 2:
		return newGasMotif(seed)
	case 3:
		return newFieldMotif(seed)
	case 4:
		return newSpacetimeMotif(seed)
	case 5:
		return newWavePacketMotif(seed)
	case 6:
		return newTracksMotif(seed)
	default:
		return newGalaxyMotif(seed)
	}
}

// Layout draws the banner on its tinted gradient.
func (b *Banner) Layout(gtx giode.Context) giode.Dimensions {
	return giode.Box(giode.Styles{
		Width: -1, BorderRadii: b.radii, Overflow: giode.OverflowHidden,
		BackgroundGradient: giode.Linear(135, theme.Blend(theme.Night, b.hue, 0.22), theme.Night),
	}, giode.Canvas(giode.Styles{Width: -1, Height: properties.Height(b.height)},
		func(p *giode.Painter) { b.motif.draw(p, b.hue) }),
	).Layout(gtx)
}

// tint is the hue faded toward the night backdrop: t = 1 is full hue.
func tint(hue string, t float64) giode.Color { return theme.Blend(theme.Night, hue, t) }
