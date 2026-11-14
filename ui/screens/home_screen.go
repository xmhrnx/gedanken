// Package screens holds the app's pages.
package screens

import (
	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/components/scroll"

	"gedanken/internal/curriculum"
	"gedanken/internal/progress"
	"gedanken/ui/components"
	"gedanken/ui/models"
)

// Home lists every part and chapter of the course.
type Home struct {
	course  *curriculum.Course
	tracker *progress.Tracker
	orbital giode.Element
	cards   [][]*components.ChapterCard
	scroll  *scroll.Scroll
}

// NewHome creates the home screen; onOpen opens chapter seed of a part.
func NewHome(course *curriculum.Course, tracker *progress.Tracker, onOpen func(part, seed int, ch curriculum.Chapter)) *Home {
	h := &Home{course: course, tracker: tracker, orbital: models.NewOrbitalCloud(300, false)}
	for pi, part := range course.Parts {
		var row []*components.ChapterCard
		for ci, ch := range part.Chapters {
			row = append(row, components.NewChapterCard(ch, pi, ci,
				func() int { return tracker.ChapterDone(ch) },
				func() { onOpen(pi, ci, ch) }))
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
	total := 0
	for _, p := range h.course.Parts {
		for _, ch := range p.Chapters {
			total += len(ch.Lessons)
		}
	}
	var kids []giode.Element
	for pi, part := range h.course.Parts {
		kids = append(kids, components.PartHeader(pi, part.Title, len(part.Chapters)))
		for _, card := range h.cards[pi] {
			kids = append(kids, card)
		}
	}
	return giode.Stack(giode.Styles{Width: -1},
		components.HomeHero(h.orbital, h.tracker.Total(), total),
		components.PageColumn(kids...),
	)
}
