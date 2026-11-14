package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// ReadingProgress is a thin bar under the top bar that fills as the
// reader scrolls through a lesson.
func ReadingProgress(fraction float32, color giode.Color) giode.Element {
	return giode.Progress(fraction, giode.Styles{Color: color, Background: theme.Background, Height: 2})
}
