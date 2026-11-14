package progress

import (
	"testing"

	"gedanken/internal/curriculum"
)

func TestTrackerSavesChanges(t *testing.T) {
	var saved map[string]bool
	saves := 0
	tr := New(map[string]bool{"c/b": true}, func(m map[string]bool) error { saved, saves = m, saves+1; return nil })
	ch := curriculum.Chapter{ID: "c", Lessons: []curriculum.LessonInfo{{ID: "a"}, {ID: "b"}}}
	if err := tr.Complete("c", "a"); err != nil {
		t.Fatal(err)
	}
	if err := tr.Complete("c", "a"); err != nil {
		t.Fatal(err)
	}
	if saves != 1 || !saved["c/a"] || !saved["c/b"] {
		t.Errorf("saves=%d saved=%v", saves, saved)
	}
	if tr.ChapterDone(ch) != 2 || tr.Total() != 2 {
		t.Error("counts wrong")
	}
}
