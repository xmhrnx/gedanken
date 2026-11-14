package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// NoteBlock highlights a lesson's key idea.
func NoteBlock(markdown string) giode.Element {
	return giode.Row(giode.Styles{
		Width: -1, Gap: 12, Wrap: giode.WrapNoWrap, Align: giode.AlignStretch,
		Background: theme.AccentSoft, BorderRadius: theme.SmallRadius, Padding: giode.UniformInset(14),
	},
		giode.Box(giode.Styles{Width: 3, BorderRadius: 2, Background: theme.Accent}),
		giode.Stack(giode.Styles{FlexGrow: 1, Gap: 4},
			giode.Text("Key idea", giode.Styles{FontSize: 12, FontWeight: giode.FontWeightSemiBold,
				Color: theme.Accent, TextTransform: giode.TextTransformUppercase}),
			giode.Markdown(markdown, theme.MarkdownBody()),
		),
	)
}
