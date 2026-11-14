package models

import (
	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/components/slider"

	"gedanken/ui/theme"
)

// SliderRow is a labelled slider showing its current value.
type SliderRow struct {
	label  string
	slider *slider.Slider
	value  func(v float32) string
}

// NewSliderRow creates a slider at initial (0–1); value formats the
// value shown beside the label.
func NewSliderRow(label string, initial float32, value func(v float32) string) *SliderRow {
	s := giode.Input.Slider(giode.Styles{Color: theme.Primary, Background: theme.Border, BorderColor: theme.Primary})
	s.SetValue(initial)
	return &SliderRow{label: label, slider: s, value: value}
}

// Value returns the slider position in [0, 1].
func (r *SliderRow) Value() float32 { return r.slider.Value() }

// Layout draws the label, value and slider.
func (r *SliderRow) Layout(gtx giode.Context) giode.Dimensions {
	return giode.Stack(giode.Styles{Width: -1, Gap: 2},
		giode.Row(giode.Styles{Width: -1, Justify: giode.JustifySpaceBetween, Wrap: giode.WrapNoWrap},
			giode.Text(r.label, theme.Caption),
			giode.Text(r.value(r.Value()), giode.Styles{FontSize: 13, Color: theme.Accent, FontFamily: "Go Mono"}),
		),
		r.slider,
	).Layout(gtx)
}
