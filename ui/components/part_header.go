package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// PartHeader introduces one part of the course on the home screen.
func PartHeader(index int, title string) giode.Element {
	return giode.Row(giode.Styles{Width: -1, Align: giode.AlignCenter, Gap: 10, Wrap: giode.WrapNoWrap,
		Padding: giode.Inset{Top: 14}},
		giode.Box(giode.Styles{Width: 4, Height: 22, BorderRadius: 2, Background: theme.PartColor(index)}),
		giode.Stack(giode.Styles{Gap: 2},
			giode.Text(fmt.Sprintf("Part %d", index+1), theme.Eyebrow),
			giode.Text(title, theme.Title),
		),
	)
}
