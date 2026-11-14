package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// ColumnRow places one list row in the centered reading column, with
// side gutters and the given space below.
func ColumnRow(child giode.Element, below int) giode.Element {
	return giode.Row(giode.Styles{Width: -1, Justify: giode.JustifyCenter, Wrap: giode.WrapNoWrap},
		giode.Box(giode.Styles{
			Width: -1, MaxWidth: theme.ContentMax,
			Padding: giode.Inset{Left: theme.Gutter, Right: theme.Gutter, Bottom: below},
		}, child),
	)
}
