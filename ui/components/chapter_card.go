package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/internal/curriculum"
	"gedanken/ui/art"
	"gedanken/ui/theme"
)

// ChapterCard is a tappable card for one chapter: an illustration in
// the part's hue above the title, blurb and progress.
type ChapterCard struct {
	chapter curriculum.Chapter
	color   giode.Color
	banner  *art.Banner
	done    func() int
	press   *giode.PressableElement
}

// NewChapterCard creates the card for chapter seed of part; done
// reports finished lessons.
func NewChapterCard(ch curriculum.Chapter, part, seed int, done func() int, onOpen func()) *ChapterCard {
	color := theme.PartColor(part)
	c := &ChapterCard{chapter: ch, color: color, done: done,
		banner: art.NewBanner(part, seed, color, 96, giode.TopCorners(theme.Radius))}
	c.press = giode.Pressable(c.content).Label(ch.Title).OnClick(func(giode.ClickEvent) { onOpen() })
	return c
}

// Layout draws the card.
func (c *ChapterCard) Layout(gtx giode.Context) giode.Dimensions {
	return c.press.Layout(gtx)
}

func (c *ChapterCard) content(s giode.PressState) giode.Element {
	border := giode.Color(theme.Border)
	if s.Pressed || s.Hovered {
		border = c.color
	}
	total := len(c.chapter.Lessons)
	return giode.Stack(giode.Styles{Width: -1, Background: theme.Surface, BorderRadius: theme.Radius,
		BorderWidth: 1, BorderColor: border, Overflow: giode.OverflowHidden},
		c.banner,
		giode.Stack(giode.Styles{Width: -1, Gap: 8, Padding: giode.Inset{Top: 14, Left: 16, Right: 16, Bottom: 16}},
			giode.Row(giode.Styles{Width: -1, Align: giode.AlignCenter, Gap: 10, Wrap: giode.WrapNoWrap},
				giode.Text(c.chapter.Title, theme.Heading.With(giode.Styles{FlexGrow: 1})),
				theme.Icon("chevron", string(c.color), 18),
			),
			giode.Text(c.chapter.Blurb, theme.Caption.With(giode.Styles{FontSize: 14, LineHeight: 20})),
			giode.Row(giode.Styles{Width: -1, Align: giode.AlignCenter, Gap: 10, Wrap: giode.WrapNoWrap,
				Padding: giode.Inset{Top: 4}},
				giode.Box(giode.Styles{FlexGrow: 1}, giode.Progress(float32(c.done())/float32(total),
					giode.Styles{Color: c.color, Background: theme.Border, Height: 3})),
				giode.Text(fmt.Sprintf("%d/%d", c.done(), total), giode.Styles{FontSize: 12, Color: theme.Muted,
					FontFamily: "Go Mono"}),
			),
		),
	)
}
