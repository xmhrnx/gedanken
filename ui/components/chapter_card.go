package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/internal/curriculum"
	"gedanken/ui/theme"
)

// ChapterCard is a tappable card for one chapter on the home screen.
type ChapterCard struct {
	chapter curriculum.Chapter
	color   giode.Color
	done    func() int
	press   *giode.PressableElement
}

// NewChapterCard creates a card; done reports finished lessons.
func NewChapterCard(ch curriculum.Chapter, color giode.Color, done func() int, onOpen func()) *ChapterCard {
	c := &ChapterCard{chapter: ch, color: color, done: done}
	c.press = giode.Pressable(c.content).Label(ch.Title).OnClick(func(giode.ClickEvent) { onOpen() })
	return c
}

// Layout draws the card.
func (c *ChapterCard) Layout(gtx giode.Context) giode.Dimensions {
	return c.press.Layout(gtx)
}

func (c *ChapterCard) content(s giode.PressState) giode.Element {
	st := theme.Card
	if s.Pressed || s.Hovered {
		st.Background = theme.SurfaceRaised
		st.BorderColor = c.color
	}
	total, written := len(c.chapter.Lessons), c.chapter.WrittenCount()
	status := fmt.Sprintf("%d of %d lessons done", c.done(), total)
	if written < total {
		status += fmt.Sprintf(" · %d written", written)
	}
	return giode.Box(st,
		giode.Row(giode.Styles{Width: -1, Align: giode.AlignCenter, Gap: 10, Wrap: giode.WrapNoWrap},
			giode.Stack(giode.Styles{FlexGrow: 1, Gap: 4},
				giode.Text(c.chapter.Title, theme.Heading),
				giode.Text(c.chapter.Blurb, theme.Caption),
			),
			theme.Icon("chevron", string(c.color), 20),
		),
		giode.Progress(float32(c.done())/float32(total), giode.Styles{Color: c.color, Background: theme.Border, Height: 4}),
		giode.Text(status, giode.Styles{FontSize: 12, Color: theme.Muted}),
	)
}
