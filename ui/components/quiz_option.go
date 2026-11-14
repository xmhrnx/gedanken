package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// OptionStatus is how a quiz option is drawn after an answer.
type OptionStatus int

// Option states.
const (
	OptionIdle OptionStatus = iota
	OptionCorrect
	OptionWrong
	OptionDimmed
)

// QuizOption is one tappable answer.
type QuizOption struct {
	text   string
	status OptionStatus
	press  *giode.PressableElement
}

// NewQuizOption creates an option that calls onChoose when tapped.
func NewQuizOption(text string, onChoose func()) *QuizOption {
	o := &QuizOption{text: text}
	o.press = giode.Pressable(o.content).Label(text).OnClick(func(giode.ClickEvent) { onChoose() })
	return o
}

// SetStatus changes how the option is drawn.
func (o *QuizOption) SetStatus(s OptionStatus) { o.status = s }

// Layout draws the option.
func (o *QuizOption) Layout(gtx giode.Context) giode.Dimensions {
	return o.press.Layout(gtx)
}

func (o *QuizOption) content(s giode.PressState) giode.Element {
	st := giode.Styles{Width: -1, Padding: giode.SymmetricInset(14, 12), BorderRadius: 12,
		BorderWidth: 1, BorderColor: theme.Border, Background: theme.SurfaceRaised}
	text := giode.Styles{FontSize: 15, Color: theme.Text, LineHeight: 21}
	switch o.status {
	case OptionCorrect:
		st.Background, st.BorderColor = theme.SuccessSoft, theme.Success
	case OptionWrong:
		st.Background, st.BorderColor = theme.DangerSoft, theme.Danger
	case OptionDimmed:
		text.Color = theme.Muted
	default:
		if s.Pressed || s.Hovered {
			st.BorderColor = theme.Primary
		}
	}
	return giode.Box(st, giode.Text(o.text, text))
}
