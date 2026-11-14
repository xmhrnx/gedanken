package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// HistoryBlock is a historical aside labelled with a year and place.
func HistoryBlock(label, markdown string) giode.Element {
	body := theme.MarkdownBody()
	body.Styles.FontSize = 15
	body.Styles.LineHeight = 23
	return giode.Box(theme.Card.With(giode.Styles{BorderColor: giode.Mix("theme.accent", "theme.surface", 0.35)}),
		giode.Row(giode.Styles{Gap: 8, Align: giode.AlignCenter},
			giode.Text("History", giode.Styles{FontSize: 12, FontWeight: giode.FontWeightSemiBold,
				Color: theme.Accent, TextTransform: giode.TextTransformUppercase}),
			giode.Text(label, theme.Caption),
		),
		giode.Markdown(markdown, body),
	)
}
