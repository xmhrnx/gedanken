package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// LessonHeader shows the lesson's position, title and summary.
func LessonHeader(number, count int, title, summary string, color giode.Color) giode.Element {
	return giode.Stack(giode.Styles{Width: -1, Gap: 10, Padding: giode.Inset{Top: 10, Bottom: 10}},
		giode.Text(fmt.Sprintf("Lesson %d of %d", number, count), theme.Eyebrow.With(giode.Styles{Color: color})),
		giode.Text(title, theme.Display),
		giode.Text(summary, theme.Lead),
		giode.Box(giode.Styles{Width: 48, Height: 2, Background: color, Margin: giode.Inset{Top: 6}}),
	)
}
