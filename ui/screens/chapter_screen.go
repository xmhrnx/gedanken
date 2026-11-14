package screens

import (
	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/components/scroll"

	"gedanken/internal/curriculum"
	"gedanken/internal/progress"
	"gedanken/ui/components"
	"gedanken/ui/theme"
)

// ChapterScreen lists a chapter's lessons.
type ChapterScreen struct {
	part    int
	chapter curriculum.Chapter
	tracker *progress.Tracker
	back    giode.Element
	banner  *components.ChapterBanner
	rows    []*components.LessonRow
	scroll  *scroll.Scroll
}

// NewChapterScreen creates the screen; onLesson opens the i-th lesson.
func NewChapterScreen(part, seed int, ch curriculum.Chapter, tracker *progress.Tracker, onBack func(), onLesson func(i int)) *ChapterScreen {
	s := &ChapterScreen{part: part, chapter: ch, tracker: tracker, back: components.NewBackButton(onBack),
		banner: components.NewChapterBanner(part, seed, ch, func() int { return tracker.ChapterDone(ch) })}
	for i, l := range ch.Lessons {
		s.rows = append(s.rows, components.NewLessonRow(i+1, l, theme.PartColor(part),
			func() bool { return tracker.Done(ch.ID, l.ID) },
			func() { onLesson(i) }))
	}
	s.scroll = giode.Scroll(s.content, giode.Styles{ScrollBar: giode.ScrollBar{HiddenMaxWidth: 600}})
	return s
}

// Layout draws the screen.
func (s *ChapterScreen) Layout(gtx giode.Context) giode.Dimensions {
	return giode.Stack(giode.Styles{Width: -1, Height: -1},
		components.TopBar(s.back, "All chapters"),
		giode.Box(giode.Styles{Width: -1, FlexGrow: 1}, s.scroll),
	).Layout(gtx)
}

func (s *ChapterScreen) content() giode.Element {
	kids := []giode.Element{s.banner}
	for _, r := range s.rows {
		kids = append(kids, r)
	}
	return components.PageColumn(kids...)
}
