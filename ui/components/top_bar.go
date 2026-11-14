package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// TopBar is a screen header with a back button and a title.
func TopBar(back giode.Element, title string) giode.Element {
	return giode.Row(giode.Styles{
		Width: -1, Align: giode.AlignCenter, Gap: 4, Wrap: giode.WrapNoWrap,
		Padding:    giode.Inset{Top: 6, Bottom: 6, Left: 6, Right: theme.Gutter},
		Background: theme.Background,
	},
		back,
		giode.Text(title, giode.Styles{FontSize: 15, FontWeight: giode.FontWeightSemiBold, Color: theme.Muted,
			TextOverflow: giode.TextOverflowEllipsis}),
	)
}

// NewBackButton creates the round back button used in top bars.
func NewBackButton(onBack func()) giode.Element {
	return giode.IconButton(theme.Icon("back", theme.Text, 22), giode.Styles{Color: theme.Text}).
		Variant(giode.IconButtonGhost).OnClick(onBack)
}
