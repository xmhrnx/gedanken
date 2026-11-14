package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// ErrorNotice reports a problem in place of content.
func ErrorNotice(message string) giode.Element {
	return giode.Box(theme.Card.With(giode.Styles{Background: theme.DangerSoft, BorderColor: theme.Danger}),
		giode.Text("Something went wrong", theme.Heading),
		giode.Text(message, theme.Caption),
	)
}
