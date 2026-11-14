package components

import (
	"image"

	"gioui.org/op"
	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/components/scroll"
	"github.com/xmhrnx/giode/formula"

	"gedanken/ui/theme"
)

// FormulaView is a typeset LaTeX formula. When it fits it is placed
// at the start or in the center; when it is wider than the screen it
// scrolls sideways.
type FormulaView struct {
	formula  *formula.Formula
	scroll   *scroll.Scroll
	centered bool
}

// NewFormulaView typesets src at the given size in sp.
func NewFormulaView(src string, size int, centered bool) *FormulaView {
	f := formula.New(src, giode.Styles{FontSize: theme.FontSize(size), Color: theme.Text})
	v := &FormulaView{formula: f, centered: centered}
	v.scroll = giode.Scroll(func() giode.Element { return f },
		giode.Styles{ScrollBar: giode.ScrollBar{Hidden: true}}).Axis(giode.Horizontal)
	return v
}

// Layout measures the formula, then centers it or lets it scroll.
func (v *FormulaView) Layout(gtx giode.Context) giode.Dimensions {
	measure := gtx
	measure.Constraints.Min = image.Point{}
	measure.Constraints.Max.X = 1 << 20
	macro := op.Record(gtx.Ops)
	dims := v.formula.Layout(measure)
	call := macro.Stop()
	width := gtx.Constraints.Max.X
	if dims.Size.X > width {
		return v.scroll.Layout(gtx)
	}
	x := 0
	if v.centered {
		x = (width - dims.Size.X) / 2
	}
	offset := op.Offset(image.Pt(x, 0)).Push(gtx.Ops)
	call.Add(gtx.Ops)
	offset.Pop()
	return giode.Dimensions{Size: image.Pt(width, dims.Size.Y), Baseline: dims.Baseline}
}
