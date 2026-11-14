// Package screens holds the app's pages.
package screens

import (
	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/components/scroll"

	"gedanken/internal/curriculum"
	"gedanken/internal/progress"
	"gedanken/ui/components"
	"gedanken/ui/models"
	"gedanken/ui/theme"
)

// Home lists every part and chapter of the course.
type Home struct {
	course  *curriculum.Course
	tracker *progress.Tracker
	atom    giode.Element
	cards   [][]*components.ChapterCard
	scroll  *scroll.Scroll
}

// NewHome creates the home screen; onOpen opens a chapter of a part.
func NewHome(course *curriculum.Course, tracker *progress.Tracker, onOpen func(part int, ch curriculum.Chapter)) *Home {
	h := &Home{course: course, tracker: tracker, atom: models.NewAtom(200)}
	for pi, part := range course.Parts {
		var row []*components.ChapterCard
		for _, ch := range part.Chapters {
			row = append(row, components.NewChapterCard(ch, theme.PartColor(pi),
				func() int { return tracker.ChapterDone(ch) },
				func() { onOpen(pi, ch) }))
		}
		h.cards = append(h.cards, row)
	}
	h.scroll = giode.Scroll(h.content, giode.Styles{ScrollBar: giode.ScrollBar{HiddenMaxWidth: 600}})
	return h
}

// Layout draws the screen.
func (h *Home) Layout(gtx giode.Context) giode.Dimensions {
	return giode.Box(giode.Styles{Width: -1, Height: -1}, h.scroll).Layout(gtx)
}

func (h *Home) content() giode.Element {
	written, total := 0, 0
	for _, p := range h.course.Parts {
		for _, ch := range p.Chapters {
			written += ch.WrittenCount()
			total += len(ch.Lessons)
		}
	}
	kids := []giode.Element{components.HomeHero(h.atom, h.tracker.Total(), written, total)}
	for pi, part := range h.course.Parts {
		kids = append(kids, components.PartHeader(pi, part.Title))
		for _, card := range h.cards[pi] {
			kids = append(kids, card)
		}
	}
	return components.PageColumn(kids...)
}
