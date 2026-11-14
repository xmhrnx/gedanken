// Package screens holds the app's pages.
package screens

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/internal/curriculum"
	"gedanken/internal/progress"
	"gedanken/ui/components"
	"gedanken/ui/models"
)

// homeRow is one row of the home list: the hero, a part header or a
// chapter card.
type homeRow struct {
	key     string
	build   func() giode.Element
	version func() any
}

// Home lists every part and chapter of the course. It is a lazy list:
// only rows near the screen are built, and a row is redrawn only when
// its version changes, so the animated hero is the only row rebuilt
// every frame.
type Home struct {
	rows []homeRow
	list *giode.LazyListElement
}

// NewHome creates the home screen; onOpen opens chapter seed of a part.
func NewHome(course *curriculum.Course, tracker *progress.Tracker, onOpen func(part, seed int, ch curriculum.Chapter)) *Home {
	h := &Home{}
	orbital := models.NewOrbitalCloud(300, false)
	total := lessonCount(course)
	h.rows = append(h.rows, homeRow{key: "hero",
		build:   func() giode.Element { return components.HomeHero(orbital, tracker.Total(), total) },
		version: func() any { return tracker.Total() }})
	for pi, part := range course.Parts {
		header := components.ColumnRow(components.PartHeader(pi, part.Title, len(part.Chapters)), 16)
		h.rows = append(h.rows, homeRow{key: fmt.Sprintf("part-%d", pi),
			build:   func() giode.Element { return header },
			version: func() any { return 0 }})
		for ci, ch := range part.Chapters {
			card := components.NewChapterCard(ch, pi, ci,
				func() int { return tracker.ChapterDone(ch) },
				func() { onOpen(pi, ci, ch) })
			row := components.ColumnRow(card, 18)
			h.rows = append(h.rows, homeRow{key: "chapter-" + ch.ID,
				build:   func() giode.Element { return row },
				version: func() any { return tracker.ChapterDone(ch) }})
		}
	}
	end := giode.Box(giode.Styles{Height: 24})
	h.rows = append(h.rows, homeRow{key: "end", build: func() giode.Element { return end },
		version: func() any { return 0 }})
	h.list = giode.LazyList(
		func() int { return len(h.rows) },
		func(i int) any { return h.rows[i].key },
		func(i int) giode.Element { return h.rows[i].build() },
		giode.Styles{ScrollBar: giode.ScrollBar{HiddenMaxWidth: 600}},
	).Version(func(i int) any { return h.rows[i].version() })
	return h
}

// lessonCount totals the lessons in the course.
func lessonCount(course *curriculum.Course) int {
	n := 0
	for _, p := range course.Parts {
		for _, ch := range p.Chapters {
			n += len(ch.Lessons)
		}
	}
	return n
}

// Layout draws the screen.
func (h *Home) Layout(gtx giode.Context) giode.Dimensions {
	return giode.Box(giode.Styles{Width: -1, Height: -1}, h.list).Layout(gtx)
}
