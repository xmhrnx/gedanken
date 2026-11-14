package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// PartHeader introduces one part of the course: a large numeral in the
// part's hue beside its title.
func PartHeader(index int, title string, chapters int) giode.Element {
	hue := theme.PartColor(index)
	return giode.Row(giode.Styles{Width: -1, Align: giode.AlignCenter, Gap: 14, Wrap: giode.WrapNoWrap,
		Padding: giode.Inset{Top: 22, Bottom: 2}},
		giode.Text(fmt.Sprintf("%02d", index+1), giode.Styles{FontSize: 46, FontFamily: theme.Serif,
			Color: hue, LineHeight: 48}),
		giode.Stack(giode.Styles{Gap: 2, FlexGrow: 1},
			giode.Text(fmt.Sprintf("Part %d · %s", index+1, plural(chapters, "chapter")), theme.Eyebrow),
			giode.Text(title, theme.Title),
		),
	)
}
