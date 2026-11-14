package screens

import (
	"io/fs"

	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/components/scroll"

	"gedanken/internal/curriculum"
	"gedanken/internal/progress"
	"gedanken/ui/components"
	"gedanken/ui/theme"
)

// LessonScreen shows one lesson and a button to the next.
type LessonScreen struct {
	part    int
	chapter curriculum.Chapter
	index   int
	tracker *progress.Tracker
	back    giode.Element
	blocks  []giode.Element
	finish  *giode.ButtonElement
	saveErr error
	scroll  *scroll.Scroll
}

// NewLessonScreen loads lesson index of the chapter. onNext opens
// another lesson by index; onBack returns to the chapter.
func NewLessonScreen(content fs.FS, part int, ch curriculum.Chapter, index int, tracker *progress.Tracker,
	onBack func(), onNext func(i int)) *LessonScreen {
	s := &LessonScreen{part: part, chapter: ch, index: index, tracker: tracker, back: components.NewBackButton(onBack)}
	lesson := ch.Lessons[index]
	if blocks, err := curriculum.LoadLesson(content, lesson); err != nil {
		s.blocks = []giode.Element{components.ErrorNotice(err.Error())}
	} else {
		s.blocks = buildBlocks(blocks)
	}
	s.finish = s.finishButton(onBack, onNext)
	s.scroll = giode.Scroll(s.content, giode.Styles{ScrollBar: giode.ScrollBar{HiddenMaxWidth: 600}}).
		OnReachBottom(s.complete)
	return s
}

// finishButton leads to the next written lesson, or back to the chapter.
func (s *LessonScreen) finishButton(onBack func(), onNext func(int)) *giode.ButtonElement {
	if next, i, ok := s.chapter.Next(s.index); ok && next.Written() {
		return giode.Button("Next: "+next.Title, theme.PrimaryButton).
			TrailingIcon(theme.Icon("chevron", theme.OnPrimary, 18)).
			OnClick(func() { s.complete(); onNext(i) })
	}
	return giode.Button("Finish chapter", theme.PrimaryButton).OnClick(func() { s.complete(); onBack() })
}

// complete marks the lesson finished; a failed save is shown on screen.
func (s *LessonScreen) complete() {
	s.saveErr = s.tracker.Complete(s.chapter.ID, s.chapter.Lessons[s.index].ID)
}

// Layout draws the screen.
func (s *LessonScreen) Layout(gtx giode.Context) giode.Dimensions {
	return giode.Stack(giode.Styles{Width: -1, Height: -1},
		components.TopBar(s.back, s.chapter.Title),
		giode.Box(giode.Styles{Width: -1, FlexGrow: 1}, s.scroll),
	).Layout(gtx)
}

func (s *LessonScreen) content() giode.Element {
	l := s.chapter.Lessons[s.index]
	kids := []giode.Element{
		components.LessonHeader(s.index+1, len(s.chapter.Lessons), l.Title, l.Summary, theme.PartColor(s.part)),
	}
	kids = append(kids, s.blocks...)
	if s.saveErr != nil {
		kids = append(kids, components.ErrorNotice("Progress could not be saved: "+s.saveErr.Error()))
	}
	kids = append(kids, giode.Box(giode.Styles{Width: -1, Padding: giode.Inset{Top: 12}}, s.finish))
	return components.PageColumn(kids...)
}
