package models

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"

	"gedanken/internal/physics"
	"gedanken/ui/theme"
)

// FieldLines draws the electric field of a +1 charge and a second
// charge whose sign and size the reader sets.
type FieldLines struct {
	charge  *SliderRow
	forStep int
	lines   [][]physics.Point
}

const fieldSteps, fieldBound = 16, 1.6

// NewFieldLines creates the model.
func NewFieldLines() *FieldLines {
	return &FieldLines{forStep: -1, charge: NewSliderRow("Second charge", 0.25, func(v float32) string {
		return fmt.Sprintf("%+.1f", fieldCharge(fieldStep(v)))
	})}
}

func fieldStep(v float32) int      { return int(math.Round(float64(v) * fieldSteps)) }
func fieldCharge(step int) float64 { return -2 + 4*float64(step)/fieldSteps }
func (m *FieldLines) charges() []physics.Charge {
	return []physics.Charge{{X: -0.5, Q: 1}, {X: 0.5, Q: fieldCharge(m.forStep)}}
}

// trace recomputes the lines when the charge changes.
func (m *FieldLines) trace(step int) {
	if step == m.forStep {
		return
	}
	m.forStep, m.lines = step, nil
	for _, c := range m.charges() {
		if c.Q == 0 {
			continue
		}
		n := int(math.Round(math.Abs(c.Q) * 12))
		for i := 0; i < n; i++ {
			a := 2 * math.Pi * (float64(i) + 0.5) / float64(n)
			x, y := c.X+0.04*math.Cos(a), c.Y+0.04*math.Sin(a)
			m.lines = append(m.lines, physics.FieldLine(x, y, m.charges(), c.Q < 0, fieldBound))
		}
	}
}

// Layout draws the field, readout and slider.
func (m *FieldLines) Layout(gtx giode.Context) giode.Dimensions {
	m.trace(fieldStep(m.charge.Value()))
	q := fieldCharge(m.forStep)
	note := "opposite charges: lines run from + to −"
	if q > 0 {
		note = "like charges: lines push apart"
	} else if q == 0 {
		note = "a single charge: lines spread out evenly"
	}
	return giode.Stack(giode.Styles{Width: -1, Gap: 12},
		stage(250, m.draw),
		readout(note),
		m.charge,
	).Layout(gtx)
}

func (m *FieldLines) draw(p *giode.Painter) {
	cx, cy := p.Width()/2, p.Height()/2
	scale := min(p.Width(), p.Height()) / (2 * 1.2)
	at := func(pt physics.Point) (float32, float32) { return cx + float32(pt.X)*scale, cy - float32(pt.Y)*scale }
	for _, line := range m.lines {
		path := p.Path()
		for i, pt := range line {
			x, y := at(pt)
			if i == 0 {
				path.MoveTo(x, y)
			} else {
				path.LineTo(x, y)
			}
		}
		path.Stroke(1.2, "#4a6cc4")
	}
	for _, c := range m.charges() {
		x, y := at(physics.Point{X: c.X, Y: c.Y})
		color, sign := giode.Color(theme.Danger), "+"
		if c.Q < 0 {
			color, sign = theme.Info, "−"
		}
		if c.Q == 0 {
			p.Circle(x, y, 9, "#34406a")
			continue
		}
		p.Circle(x, y, 11, color)
		label(p, x, y, sign, giode.TextAnchorCenter, "#0a0d16")
	}
}
