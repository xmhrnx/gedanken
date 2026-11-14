package components

import (
	"fmt"

	"github.com/xmhrnx/giode"

	"gedanken/internal/curriculum"
	"gedanken/ui/theme"
)

// StepsBlock walks through a derivation one step at a time.
type StepsBlock struct {
	title    string
	steps    []curriculum.Step
	formulas []*FormulaView
	shown    int
	next     *giode.ButtonElement
	all      *giode.ButtonElement
}

// NewStepsBlock creates a derivation showing its first step.
func NewStepsBlock(title string, steps []curriculum.Step) *StepsBlock {
	b := &StepsBlock{title: title, steps: steps, shown: 1}
	for _, s := range steps {
		b.formulas = append(b.formulas, NewFormulaView(s.Equation, 19, false))
	}
	b.next = giode.Button("Next step", theme.PrimaryButton).OnClick(func() { b.shown = min(b.shown+1, len(b.steps)) })
	b.all = giode.Button("Show all", theme.QuietButton).Variant(giode.VariantOutlined).
		OnClick(func() { b.shown = len(b.steps) })
	return b
}

// Version changes when another step is revealed, so a cached drawing
// of the block is replaced.
func (b *StepsBlock) Version() any { return b.shown }

// Layout draws the revealed steps and the controls.
func (b *StepsBlock) Layout(gtx giode.Context) giode.Dimensions {
	kids := []giode.Element{
		giode.Text(fmt.Sprintf("Derivation · %d steps", len(b.steps)), giode.Styles{FontSize: 12,
			FontWeight: giode.FontWeightSemiBold, Color: theme.Primary, TextTransform: giode.TextTransformUppercase}),
	}
	if b.title != "" {
		kids = append(kids, giode.Text(b.title, theme.Heading))
	}
	for i := 0; i < b.shown; i++ {
		kids = append(kids, stepView(i+1, b.formulas[i], b.steps[i]))
	}
	if b.shown < len(b.steps) {
		kids = append(kids, giode.Row(giode.Styles{Gap: 10, Padding: giode.Inset{Top: 4}}, b.next, b.all))
	} else if len(b.steps) > 1 {
		kids = append(kids, giode.Text("End of derivation.", theme.Caption))
	}
	return giode.Box(theme.Card.With(giode.Styles{Gap: 14}), kids...).Layout(gtx)
}

// stepView is one numbered step: a badge beside the equation and its
// explanation. (A connecting line between badges would need the row to
// stretch its children, and Giode's AlignStretch currently mis-measures
// wrapped text, cutting off the explanation's last line.)
func stepView(n int, f *FormulaView, s curriculum.Step) giode.Element {
	badge := giode.Box(giode.Styles{Width: 26, Height: 26, BorderRadius: 13, Background: theme.PrimarySoft,
		Align: giode.AlignCenter, Justify: giode.JustifyCenter},
		giode.Text(fmt.Sprint(n), giode.Styles{FontSize: 13, FontWeight: giode.FontWeightBold, Color: theme.Primary}))
	body := []giode.Element{}
	if s.Equation != "" {
		body = append(body, f)
	}
	if s.Explanation != "" {
		opts := theme.MarkdownBody()
		opts.Styles.FontSize = 15
		opts.Styles.Color = theme.Muted
		body = append(body, giode.Markdown(s.Explanation, opts))
	}
	return giode.Row(giode.Styles{Width: -1, Gap: 12, Wrap: giode.WrapNoWrap, Align: giode.AlignStart},
		badge,
		giode.Stack(giode.Styles{FlexGrow: 1, Gap: 6, Padding: giode.Inset{Bottom: 8}}, body...),
	)
}
