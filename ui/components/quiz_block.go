package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/internal/curriculum"
	"gedanken/ui/theme"
)

// QuizBlock asks multiple-choice questions one at a time. The reader
// can skip it.
type QuizBlock struct {
	questions []curriculum.Question
	options   [][]*QuizOption
	current   int
	chosen    int
	correct   int
	skipped   bool
	skip      *giode.ButtonElement
	next      *giode.ButtonElement
	retry     *giode.ButtonElement
}

// NewQuizBlock creates a quiz over the given questions.
func NewQuizBlock(questions []curriculum.Question) *QuizBlock {
	q := &QuizBlock{questions: questions, chosen: -1}
	for qi, question := range questions {
		row := make([]*QuizOption, len(question.Options))
		for oi, text := range question.Options {
			row[oi] = NewQuizOption(text, func() { q.choose(qi, oi) })
		}
		q.options = append(q.options, row)
	}
	q.skip = giode.Button("Skip quiz", theme.QuietButton).Variant(giode.VariantGhost).OnClick(func() { q.skipped = true })
	q.next = giode.Button("Next question", theme.PrimaryButton).OnClick(q.advance)
	q.retry = giode.Button("Take the quiz", theme.QuietButton).Variant(giode.VariantOutlined).OnClick(q.reset)
	return q
}

func (q *QuizBlock) finished() bool { return q.current >= len(q.questions) }

// choose records the answer to question qi.
func (q *QuizBlock) choose(qi, oi int) {
	if qi != q.current || q.chosen >= 0 {
		return
	}
	q.chosen = oi
	answer := q.questions[qi].Answer
	if oi == answer {
		q.correct++
	}
	for i, o := range q.options[qi] {
		switch {
		case i == answer:
			o.SetStatus(OptionCorrect)
		case i == oi:
			o.SetStatus(OptionWrong)
		default:
			o.SetStatus(OptionDimmed)
		}
	}
}

func (q *QuizBlock) advance() {
	q.current++
	q.chosen = -1
}

// reset starts the quiz over.
func (q *QuizBlock) reset() {
	q.current, q.chosen, q.correct, q.skipped = 0, -1, 0, false
	for _, row := range q.options {
		for _, o := range row {
			o.SetStatus(OptionIdle)
		}
	}
}

// quizState is everything the quiz's drawing depends on.
type quizState struct {
	current, chosen, correct int
	skipped                  bool
}

// Version changes whenever the quiz's drawing would, so a cached
// drawing of the block is replaced.
func (q *QuizBlock) Version() any {
	return quizState{q.current, q.chosen, q.correct, q.skipped}
}

// Layout draws the current state of the quiz.
func (q *QuizBlock) Layout(gtx giode.Context) giode.Dimensions {
	header := giode.Text("Check yourself", giode.Styles{FontSize: 12, FontWeight: giode.FontWeightSemiBold,
		Color: theme.Success, TextTransform: giode.TextTransformUppercase})
	var kids []giode.Element
	switch {
	case q.skipped:
		kids = []giode.Element{header, giode.Text("Quiz skipped.", theme.Caption), q.retry}
	case q.finished():
		q.retry.Label("Try again")
		kids = []giode.Element{header,
			giode.Text(fmt.Sprintf("%d of %d correct", q.correct, len(q.questions)), theme.Title), q.retry}
	default:
		kids = q.questionView(header)
	}
	return giode.Box(theme.Card.With(giode.Styles{Gap: 12}), kids...).Layout(gtx)
}

// questionView shows the current question, its options and feedback.
func (q *QuizBlock) questionView(header giode.Element) []giode.Element {
	question := q.questions[q.current]
	kids := []giode.Element{
		giode.Row(giode.Styles{Width: -1, Justify: giode.JustifySpaceBetween, Align: giode.AlignCenter},
			header,
			giode.Text(fmt.Sprintf("%d / %d", q.current+1, len(q.questions)), theme.Caption)),
		giode.Text(question.Prompt, theme.Heading),
	}
	for _, o := range q.options[q.current] {
		kids = append(kids, o)
	}
	if q.chosen < 0 {
		return append(kids, q.skip)
	}
	verdict, color := "Not quite.", giode.Color(theme.Danger)
	if q.chosen == question.Answer {
		verdict, color = "Correct.", theme.Success
	}
	kids = append(kids, giode.Text(verdict, giode.Styles{FontSize: 15, FontWeight: giode.FontWeightSemiBold, Color: color}))
	if question.Explanation != "" {
		kids = append(kids, giode.Text(question.Explanation, theme.Body.With(giode.Styles{FontSize: 15, Color: theme.Muted})))
	}
	if q.current == len(q.questions)-1 {
		q.next.Label("See result")
	} else {
		q.next.Label("Next question")
	}
	return append(kids, giode.Row(giode.Styles{Gap: 10, Align: giode.AlignCenter}, q.next, q.skip))
}
