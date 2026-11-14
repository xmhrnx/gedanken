package models

import "github.com/xmhrnx/giode"

// constructors builds each model by the id lessons use.
var constructors = map[string]func() giode.Element{
	"slope":          func() giode.Element { return NewSlope() },
	"area":           func() giode.Element { return NewArea() },
	"projectile":     func() giode.Element { return NewProjectile() },
	"falling":        func() giode.Element { return NewFalling() },
	"escape":         func() giode.Element { return NewEscape() },
	"gravity-well":   func() giode.Element { return NewGravityWell() },
	"light-cones":    func() giode.Element { return NewLightCones() },
	"binding-curve":  func() giode.Element { return NewBindingCurve() },
	"chain-reaction": func() giode.Element { return NewChainReaction() },
	"critical-mass":  func() giode.Element { return NewCriticalMass() },
	"atom":           func() giode.Element { return NewAtom(220) },
	"orbit":          func() giode.Element { return NewOrbit() },
	"pendulum":       func() giode.Element { return NewPendulum() },
	"wave":           func() giode.Element { return NewWave() },
	"interference":   func() giode.Element { return NewInterference() },
	"gas":            func() giode.Element { return NewGas() },
	"field-lines":    func() giode.Element { return NewFieldLines() },
	"light-clock":    func() giode.Element { return NewLightClock() },
	"bohr":           func() giode.Element { return NewBohr() },
	"wavefunction":   func() giode.Element { return NewWavefunction() },
	"decay":          func() giode.Element { return NewDecay() },
	"orbital":        func() giode.Element { return NewOrbitalCloud(260, true) },
}

// New creates the model with the given id.
func New(id string) (giode.Element, bool) {
	build, ok := constructors[id]
	if !ok {
		return nil, false
	}
	return build(), true
}
