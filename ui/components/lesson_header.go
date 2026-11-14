package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// LessonHeader shows the lesson's position, title and summary.
func LessonHeader(number, count int, title, summary string, color giode.Color) giode.Element {
	return giode.Stack(giode.Styles{Width: -1, Gap: 8, Padding: giode.Inset{Top: 4, Bottom: 6}},
		giode.Text(fmt.Sprintf("Lesson %d of %d", number, count), theme.Eyebrow.With(giode.Styles{Color: color})),
		giode.Text(title, theme.Display),
		giode.Text(summary, giode.Styles{FontSize: 16, Color: theme.Muted, LineHeight: 23}),
	)
}
