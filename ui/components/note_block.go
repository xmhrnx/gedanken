package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// NoteBlock highlights a lesson's key idea. A short accent line marks
// it rather than a full-height side bar: a side bar needs the row to
// stretch its children, and Giode's AlignStretch currently mis-measures
// wrapped text, cutting off its last line.
func NoteBlock(markdown string) giode.Element {
	return giode.Stack(giode.Styles{
		Width: -1, Gap: 6, Background: theme.AccentSoft, BorderRadius: theme.SmallRadius,
		Padding: giode.UniformInset(16),
	},
		giode.Box(giode.Styles{Width: 28, Height: 3, BorderRadius: 2, Background: theme.Accent}),
		giode.Text("Key idea", giode.Styles{FontSize: 12, FontWeight: giode.FontWeightSemiBold,
			Color: theme.Accent, TextTransform: giode.TextTransformUppercase, Margin: giode.Inset{Top: 4}}),
		giode.Markdown(markdown, theme.MarkdownBody()),
	)
}
