package screens

import (
	"fmt"
	"io/fs"

	"github.com/xmhrnx/giode"

	"gedanken/internal/curriculum"
	"gedanken/internal/progress"
	"gedanken/ui/components"
	"gedanken/ui/theme"
)

// versioned is a block whose drawing depends on its own state, such as
// a derivation's revealed steps or a quiz's progress.
type versioned interface {
	Version() any
}

// LessonScreen shows one lesson and a button to the next. The lesson
// is a lazy list of its blocks: opening it lays out only what is on
// screen, the rest is laid out as the reader scrolls to it, and blocks
// that have not changed replay their recorded drawing.
type LessonScreen struct {
	part    int
	chapter curriculum.Chapter
	index   int
	tracker *progress.Tracker
	back    giode.Element
	rows    []giode.Element
	saveErr error
	list    *giode.LazyListElement
}

// NewLessonScreen loads lesson index of the chapter. onNext opens
// another lesson by index; onBack returns to the chapter.
func NewLessonScreen(content fs.FS, part int, ch curriculum.Chapter, index int, tracker *progress.Tracker,
	onBack func(), onNext func(i int)) *LessonScreen {
	s := &LessonScreen{part: part, chapter: ch, index: index, tracker: tracker, back: components.NewBackButton(onBack)}
	l := ch.Lessons[index]
	s.rows = append(s.rows, components.LessonHeader(index+1, len(ch.Lessons), l.Title, l.Summary, theme.PartColor(part)))
	if blocks, err := curriculum.LoadLesson(content, l); err != nil {
		s.rows = append(s.rows, components.ErrorNotice(err.Error()))
	} else {
		s.rows = append(s.rows, buildBlocks(blocks)...)
	}
	finish := s.finishButton(onBack, onNext)
	s.rows = append(s.rows, giode.Stack(giode.Styles{Width: -1, Gap: 12, Padding: giode.Inset{Top: 12, Bottom: 28}},
		giode.Raw(s.saveNotice), finish))
	s.list = giode.LazyList(
		func() int { return len(s.rows) },
		func(i int) any { return i },
		func(i int) giode.Element { return components.ColumnRow(s.rows[i], 18) },
		giode.Styles{ScrollBar: giode.ScrollBar{HiddenMaxWidth: 600}},
	).Version(s.version).OnReachEnd(s.complete)
	return s
}

// version changes when a row's drawing would: stateful blocks report
// their own state, and the last row shows a failed save.
func (s *LessonScreen) version(i int) any {
	if i == len(s.rows)-1 {
		return fmt.Sprint(s.saveErr)
	}
	if v, ok := s.rows[i].(versioned); ok {
		return v.Version()
	}
	return 0
}

// saveNotice shows why progress could not be saved, if it could not.
func (s *LessonScreen) saveNotice(gtx giode.Context) giode.Dimensions {
	if s.saveErr == nil {
		return giode.Dimensions{}
	}
	return components.ErrorNotice("Progress could not be saved: " + s.saveErr.Error()).Layout(gtx)
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

// readingProgress estimates how far through the lesson the reader is,
// from the last block in view.
func (s *LessonScreen) readingProgress() float32 {
	_, last := s.list.VisibleRange()
	if len(s.rows) < 2 {
		return 1
	}
	return float32(last) / float32(len(s.rows)-1)
}

// Layout draws the screen.
func (s *LessonScreen) Layout(gtx giode.Context) giode.Dimensions {
	return giode.Stack(giode.Styles{Width: -1, Height: -1},
		components.TopBar(s.back, s.chapter.Title),
		components.ReadingProgress(s.readingProgress(), theme.PartColor(s.part)),
		giode.Box(giode.Styles{Width: -1, FlexGrow: 1, Padding: giode.Inset{Top: 8}}, s.list),
	).Layout(gtx)
}
