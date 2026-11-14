package screens

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/internal/curriculum"
	"gedanken/ui/components"
	"gedanken/ui/models"
)

// buildBlocks creates one element per lesson block. Stateful blocks
// (derivations, quizzes, models) are created here once per lesson.
func buildBlocks(blocks []curriculum.Block) []giode.Element {
	out := make([]giode.Element, 0, len(blocks))
	for _, b := range blocks {
		out = append(out, buildBlock(b))
	}
	return out
}

func buildBlock(b curriculum.Block) giode.Element {
	switch b.Kind {
	case curriculum.KindEquation:
		return components.NewEquationBlock(b.Text, b.Arg)
	case curriculum.KindSteps:
		return components.NewStepsBlock(b.Arg, b.Steps)
	case curriculum.KindNote:
		return components.NoteBlock(b.Text)
	case curriculum.KindHistory:
		return components.HistoryBlock(b.Arg, b.Text)
	case curriculum.KindModel:
		model, ok := models.New(b.Arg)
		if !ok {
			return components.ErrorNotice(fmt.Sprintf("Unknown model %q.", b.Arg))
		}
		return components.ModelBlock(model, b.Text)
	case curriculum.KindQuiz:
		return components.NewQuizBlock(b.Questions)
	default:
		return components.ProseBlock(b.Text)
	}
}
