// Package progress records which lessons the reader has finished.
package progress

import (
	"maps"

	"gedanken/internal/curriculum"
)

// SaveFunc persists the finished-lesson set.
type SaveFunc func(done map[string]bool) error

// Tracker holds finished lessons and saves them on every change.
type Tracker struct {
	done map[string]bool
	save SaveFunc
}

// New starts from previously saved progress; save is called with a
// copy of the set whenever it changes.
func New(saved map[string]bool, save SaveFunc) *Tracker {
	done := maps.Clone(saved)
	if done == nil {
		done = map[string]bool{}
	}
	return &Tracker{done: done, save: save}
}

func key(chapter, lesson string) string { return chapter + "/" + lesson }

// Complete marks a lesson finished and saves the change.
func (t *Tracker) Complete(chapter, lesson string) error {
	k := key(chapter, lesson)
	if t.done[k] {
		return nil
	}
	t.done[k] = true
	return t.save(maps.Clone(t.done))
}

// Done reports whether a lesson is finished.
func (t *Tracker) Done(chapter, lesson string) bool {
	return t.done[key(chapter, lesson)]
}

// ChapterDone counts the finished lessons of a chapter.
func (t *Tracker) ChapterDone(ch curriculum.Chapter) int {
	n := 0
	for _, l := range ch.Lessons {
		if t.Done(ch.ID, l.ID) {
			n++
		}
	}
	return n
}

// Total counts all finished lessons.
func (t *Tracker) Total() int { return len(t.done) }
