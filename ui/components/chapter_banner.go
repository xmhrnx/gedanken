package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/internal/curriculum"
	"gedanken/ui/art"
	"gedanken/ui/theme"
)

// ChapterBanner is the head of a chapter screen: a large illustration
// with the chapter's title, blurb and progress beneath.
type ChapterBanner struct {
	part    int
	chapter curriculum.Chapter
	art     *art.Banner
	done    func() int
}

// NewChapterBanner creates the banner for chapter seed of part; done
// reports finished lessons.
func NewChapterBanner(part, seed int, ch curriculum.Chapter, done func() int) *ChapterBanner {
	return &ChapterBanner{part: part, chapter: ch, done: done,
		art: art.NewBanner(part, seed, theme.PartColor(part), 170, giode.UniformCorners(theme.Radius))}
}

// Layout draws the banner.
func (b *ChapterBanner) Layout(gtx giode.Context) giode.Dimensions {
	color := theme.PartColor(b.part)
	return giode.Stack(giode.Styles{Width: -1, Gap: 10, Padding: giode.Inset{Bottom: 8}},
		b.art,
		giode.Text(fmt.Sprintf("Part %d · %s", b.part+1, plural(len(b.chapter.Lessons), "lesson")),
			theme.Eyebrow.With(giode.Styles{Color: color, Margin: giode.Inset{Top: 8}})),
		giode.Text(b.chapter.Title, theme.Display),
		giode.Text(b.chapter.Blurb, theme.Lead),
		giode.Progress(float32(b.done())/float32(len(b.chapter.Lessons)),
			giode.Styles{Color: color, Background: theme.Border, Height: 3}),
	).Layout(gtx)
}
