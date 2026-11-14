package components

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// EquationBlock is a displayed formula with an optional caption.
type EquationBlock struct {
	formula *FormulaView
	caption string
}

// NewEquationBlock typesets the LaTeX source.
func NewEquationBlock(latex, caption string) *EquationBlock {
	return &EquationBlock{formula: NewFormulaView(latex, 21, true), caption: caption}
}

// Layout draws the formula on its tinted panel.
func (b *EquationBlock) Layout(gtx giode.Context) giode.Dimensions {
	kids := []giode.Element{b.formula}
	if b.caption != "" {
		kids = append(kids, giode.Text(b.caption, giode.Styles{FontSize: 12, Color: theme.Muted,
			TextAlign: giode.TextAlignCenter}))
	}
	return giode.Stack(giode.Styles{
		Width: -1, Align: giode.AlignCenter, Gap: 8, Padding: giode.SymmetricInset(14, 16),
		Background: theme.PrimarySoft, BorderRadius: theme.SmallRadius,
	}, kids...).Layout(gtx)
}
