package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// HomeHero is the top of the home screen: the animated atom, the app
// name and overall progress.
func HomeHero(atom giode.Element, done, written, total int) giode.Element {
	return giode.Stack(giode.Styles{Width: -1, Gap: 6, Align: giode.AlignCenter, Padding: giode.Inset{Top: 8}},
		atom,
		giode.Text("Gedanken", theme.Display.With(giode.Styles{FontSize: 34})),
		giode.Text("Physics from motion to black holes, with the math and the history.",
			theme.Caption.With(giode.Styles{FontSize: 14, TextAlign: giode.TextAlignCenter})),
		giode.Text(fmt.Sprintf("%d lessons done · %d of %d written", done, written, total),
			giode.Styles{FontSize: 12, Color: theme.Accent, FontFamily: "Go Mono"}),
	)
}
