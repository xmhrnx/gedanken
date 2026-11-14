package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// heroGradient fades the hero's glow into the page.
var heroGradient = giode.Linear(180, theme.Blend(theme.Night, "#8ea8ff", 0.16), "#0a0d16")

// HomeHero is the top of the home screen: a hydrogen orbital, the app
// name and overall progress, on a full-width glow.
func HomeHero(orbital giode.Element, done, total int) giode.Element {
	return giode.Stack(giode.Styles{Width: -1, Align: giode.AlignCenter, BackgroundGradient: heroGradient,
		Padding: giode.Inset{Top: 12, Bottom: 22}},
		giode.Box(giode.Styles{Width: -1, MaxWidth: theme.ContentMax}, orbital),
		giode.Stack(giode.Styles{Align: giode.AlignCenter, Gap: 8, Padding: giode.SymmetricInset(theme.Gutter, 0)},
			giode.Text("Gedanken", theme.Hero),
			giode.Text("Physics from first principles, with the mathematics and the history.",
				theme.Lead.With(giode.Styles{TextAlign: giode.TextAlignCenter, FontSize: 17})),
			giode.Box(giode.Styles{Padding: giode.SymmetricInset(12, 5), BorderRadius: 99, Background: theme.Surface,
				BorderWidth: 1, BorderColor: theme.Border, Margin: giode.Inset{Top: 6}},
				giode.Text(fmt.Sprintf("%d of %d lessons done", done, total),
					giode.Styles{FontSize: 12, Color: theme.Accent, FontFamily: "Go Mono"})),
		),
	)
}
