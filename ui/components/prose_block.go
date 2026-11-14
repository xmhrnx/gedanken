package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// ProseBlock renders lesson Markdown.
func ProseBlock(markdown string) giode.Element {
	return giode.Markdown(markdown, theme.MarkdownBody())
}
