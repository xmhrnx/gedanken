package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"github.com/xmhrnx/giode/styles/properties"

	"gedanken/internal/curriculum"
	"gedanken/ui/theme"
)

// LessonRow is a tappable lesson entry on a chapter screen. Lessons not
// written yet are shown dimmed and cannot be opened.
type LessonRow struct {
	number int
	lesson curriculum.LessonInfo
	color  giode.Color
	done   func() bool
	press  *giode.PressableElement
}

// NewLessonRow creates the row for the lesson at position number.
func NewLessonRow(number int, l curriculum.LessonInfo, color giode.Color, done func() bool, onOpen func()) *LessonRow {
	r := &LessonRow{number: number, lesson: l, color: color, done: done}
	r.press = giode.Pressable(r.content).Label(l.Title).OnClick(func(giode.ClickEvent) { onOpen() }).
		Disabled(!l.Written())
	return r
}

// Layout draws the row.
func (r *LessonRow) Layout(gtx giode.Context) giode.Dimensions {
	return r.press.Layout(gtx)
}

func (r *LessonRow) content(s giode.PressState) giode.Element {
	bg := giode.Color(theme.Surface)
	if s.Pressed || s.Hovered {
		bg = theme.SurfaceRaised
	}
	summary := r.lesson.Summary
	if !r.lesson.Written() {
		summary = "Not written yet · " + summary
	}
	return giode.Row(giode.Styles{
		Width: -1, Align: giode.AlignCenter, Gap: 14, Wrap: giode.WrapNoWrap,
		Padding: giode.SymmetricInset(14, 14), Background: bg, BorderRadius: 14,
		Opacity: r.opacity(),
	},
		r.badge(),
		giode.Stack(giode.Styles{FlexGrow: 1, Gap: 3},
			giode.Text(r.lesson.Title, theme.Heading),
			giode.Text(summary, theme.Caption),
		),
	)
}

func (r *LessonRow) opacity() properties.Opacity {
	if r.lesson.Written() {
		return 1
	}
	return 0.5
}

// badge is the lesson number, a check when done, or a clock when unwritten.
func (r *LessonRow) badge() giode.Element {
	circle := giode.Styles{Width: 32, Height: 32, BorderRadius: 16, Align: giode.AlignCenter,
		Justify: giode.JustifyCenter, Background: theme.Border}
	switch {
	case !r.lesson.Written():
		return giode.Box(circle, theme.Icon("clock", theme.Muted, 18))
	case r.done():
		circle.Background = r.color
		return giode.Box(circle, theme.Icon("check", theme.Background, 20))
	default:
		return giode.Box(circle, giode.Text(fmt.Sprint(r.number), giode.Styles{FontSize: 14,
			FontWeight: giode.FontWeightBold, Color: r.color}))
	}
}
