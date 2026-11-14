package screens

import (
	"fmt"

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
	rows    []*components.LessonRow
	scroll  *scroll.Scroll
}

// NewChapterScreen creates the screen; onLesson opens the i-th lesson.
func NewChapterScreen(part int, ch curriculum.Chapter, tracker *progress.Tracker, onBack func(), onLesson func(i int)) *ChapterScreen {
	s := &ChapterScreen{part: part, chapter: ch, tracker: tracker, back: components.NewBackButton(onBack)}
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
	color := theme.PartColor(s.part)
	done := s.tracker.ChapterDone(s.chapter)
	kids := []giode.Element{
		giode.Stack(giode.Styles{Width: -1, Gap: 8, Padding: giode.Inset{Bottom: 6}},
			giode.Text(fmt.Sprintf("Part %d", s.part+1), theme.Eyebrow.With(giode.Styles{Color: color})),
			giode.Text(s.chapter.Title, theme.Display),
			giode.Text(s.chapter.Blurb, giode.Styles{FontSize: 16, Color: theme.Muted, LineHeight: 23}),
			giode.Progress(float32(done)/float32(len(s.chapter.Lessons)),
				giode.Styles{Color: color, Background: theme.Border, Height: 4}),
		),
	}
	for _, r := range s.rows {
		kids = append(kids, r)
	}
	return components.PageColumn(kids...)
}
