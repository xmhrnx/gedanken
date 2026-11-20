package content

import (
	"testing"

	"github.com/xmhrnx/giode/formula"

	"gedanken/internal/curriculum"
)

// TestEveryLessonParses checks each written lesson parses and uses
// only known models.
func TestEveryLessonParses(t *testing.T) {
	course, err := curriculum.LoadCourse(FS)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range course.Parts {
		for _, ch := range p.Chapters {
			for _, l := range ch.Lessons {
				if !l.Written() {
					continue
				}
				blocks, err := curriculum.LoadLesson(FS, l)
				if err != nil {
					t.Error(err)
					continue
				}
				checkBlocks(t, l.File, blocks)
			}
		}
	}
}

func checkBlocks(t *testing.T, file string, blocks []curriculum.Block) {
	for _, b := range blocks {
		switch b.Kind {
		case curriculum.KindModel:
			if !curriculum.KnownModel(b.Arg) {
				t.Errorf("%s: unknown model %q", file, b.Arg)
			}
		case curriculum.KindEquation:
			checkLatex(t, file, b.Text)
		case curriculum.KindSteps:
			for _, s := range b.Steps {
				checkLatex(t, file, s.Equation)
			}
		}
	}
}

// checkLatex reports formulas Giode cannot typeset as written: unknown
// commands, missing braces or arguments, unmatched \left or \begin.
func checkLatex(t *testing.T, file, src string) {
	if err := formula.Check(src); err != nil {
		t.Errorf("%s: %v in %q", file, err, src)
	}
}
