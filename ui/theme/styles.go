package theme

import (
	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/styles/properties"
)

// Sizes in dp.
const (
	Gutter      = 18
	Radius      = 16
	SmallRadius = 10
	ContentMax  = 720
)

// Text styles. Titles, headings and summaries are set in the serif.
// Lesson prose stays in the sans: it is full of symbols (≈ √ → ≤) the
// serif lacks, and Giode's fallback for missing glyphs currently draws
// neighbouring serif text in the wrong weight.
var (
	Hero    = giode.Styles{FontSize: 44, FontWeight: giode.FontWeightSemiBold, Color: Text, LineHeight: 50, FontFamily: Serif}
	Display = giode.Styles{FontSize: 32, FontWeight: giode.FontWeightSemiBold, Color: Text, LineHeight: 38, FontFamily: Serif}
	Title   = giode.Styles{FontSize: 25, FontWeight: giode.FontWeightSemiBold, Color: Text, LineHeight: 30, FontFamily: Serif}
	Heading = giode.Styles{FontSize: 19, FontWeight: giode.FontWeightSemiBold, Color: Text, LineHeight: 24, FontFamily: Serif}
	Body    = giode.Styles{FontSize: 16, Color: Text, LineHeight: 26}
	Lead    = giode.Styles{FontSize: 18, Color: Muted, LineHeight: 26, FontFamily: Serif, FontStyle: giode.FontStyleItalic}
	Caption = giode.Styles{FontSize: 13, Color: Muted, LineHeight: 18}
	Eyebrow = giode.Styles{FontSize: 12, FontWeight: giode.FontWeightSemiBold, Color: Muted,
		TextTransform: giode.TextTransformUppercase}
	Readout = giode.Styles{FontSize: 14, Color: Text, FontFamily: "Go Mono", LineHeight: 20}
)

// Card is the base style of a raised panel.
var Card = giode.Styles{
	FlexDirection: giode.FlexDirectionColumn,
	Width:         -1,
	Background:    Surface,
	BorderRadius:  Radius,
	BorderWidth:   1,
	BorderColor:   Border,
	Padding:       giode.UniformInset(16),
	Gap:           10,
}

// MarkdownBody renders lesson prose.
func MarkdownBody() giode.MarkdownOptions {
	return giode.MarkdownOptions{Styles: Body, Gap: 14}
}

// PrimaryButton is a filled call-to-action button style.
var PrimaryButton = giode.Styles{
	Background:   Primary,
	Color:        OnPrimary,
	BorderRadius: 12,
	FontSize:     15,
	FontWeight:   giode.FontWeightSemiBold,
	Padding:      giode.SymmetricInset(18, 12),
}

// QuietButton is an outlined secondary button style.
var QuietButton = giode.Styles{
	BorderColor:  Border,
	Color:        Muted,
	BorderRadius: 12,
	FontSize:     14,
	Padding:      giode.SymmetricInset(16, 10),
}

// Icon returns a registered icon in the given color and size.
func Icon(name, color string, size int) giode.Element {
	return giode.Icon(name, giode.Styles{Color: giode.Color(color), Width: properties.Width(size)})
}

// FontSize converts a size in sp to a style value.
func FontSize(sp int) properties.FontSize { return properties.FontSize(sp) }
