package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// ModelBlock frames an interactive model with its caption.
func ModelBlock(model giode.Element, caption string) giode.Element {
	kids := []giode.Element{
		giode.Text("Interactive", giode.Styles{FontSize: 12, FontWeight: giode.FontWeightSemiBold,
			Color: theme.Info, TextTransform: giode.TextTransformUppercase}),
	}
	if caption != "" {
		kids = append(kids, giode.Text(caption, theme.Caption))
	}
	kids = append(kids, model)
	return giode.Box(theme.Card, kids...)
}
