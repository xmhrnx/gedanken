// Package components holds the reusable pieces screens are built from.
package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// PageColumn centers content in a readable column with side gutters.
func PageColumn(children ...giode.Element) giode.Element {
	return giode.Row(giode.Styles{Width: -1, Justify: giode.JustifyCenter, Wrap: giode.WrapNoWrap},
		giode.Stack(giode.Styles{
			Width: -1, MaxWidth: theme.ContentMax, Gap: 18,
			Padding: giode.Inset{Top: 8, Bottom: 40, Left: theme.Gutter, Right: theme.Gutter},
		}, children...),
	)
}
