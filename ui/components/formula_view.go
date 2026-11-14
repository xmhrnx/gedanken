package components

import (
	"image"

	"gioui.org/op"
	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/components/scroll"
	"github.com/xmhrnx/giode/formula"

	"gedanken/ui/theme"
)

// FormulaView is a typeset LaTeX formula. When it fits it is placed at
// the start or in the center; when it is too wide it is set smaller,
// down to minScale of its size, and beyond that it scrolls sideways.
type FormulaView struct {
	src      string
	size     int
	centered bool
	sizes    map[int]*formula.Formula
	fitted   *formula.Formula
	scroll   *scroll.Scroll
}

const minScale = 0.6

// NewFormulaView typesets src at the given size in sp.
func NewFormulaView(src string, size int, centered bool) *FormulaView {
	v := &FormulaView{src: src, size: size, centered: centered, sizes: map[int]*formula.Formula{}}
	v.fitted = v.at(size)
	v.scroll = giode.Scroll(func() giode.Element { return v.fitted },
		giode.Styles{ScrollBar: giode.ScrollBar{Hidden: true}}).Axis(giode.Horizontal)
	return v
}

// at returns the formula typeset at size sp, made once per size.
func (v *FormulaView) at(size int) *formula.Formula {
	f, ok := v.sizes[size]
	if !ok {
		f = formula.New(v.src, giode.Styles{FontSize: theme.FontSize(size), Color: theme.Text})
		v.sizes[size] = f
	}
	return f
}

// measure lays f out unconstrained into a macro and returns it.
func measure(gtx giode.Context, f *formula.Formula) (op.CallOp, giode.Dimensions) {
	free := gtx
	free.Constraints.Min = image.Point{}
	free.Constraints.Max.X = 1 << 20
	macro := op.Record(gtx.Ops)
	dims := f.Layout(free)
	return macro.Stop(), dims
}

// Layout sets the formula at the largest size that fits, then places it.
func (v *FormulaView) Layout(gtx giode.Context) giode.Dimensions {
	width := gtx.Constraints.Max.X
	call, dims := measure(gtx, v.at(v.size))
	if dims.Size.X > width {
		scaled := int(float32(v.size) * float32(width) / float32(dims.Size.X))
		v.fitted = v.at(max(scaled, int(float32(v.size)*minScale)))
		call, dims = measure(gtx, v.fitted)
		if dims.Size.X > width {
			return v.scroll.Layout(gtx)
		}
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
