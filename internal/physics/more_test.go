package physics

import (
	"math"
	"testing"
)

func TestCannonOutcomes(t *testing.T) {
	cases := map[float64]OrbitOutcome{0.5: Crashes, 1.0: Orbits, 1.2: Orbits, 1.5: Escapes}
	for speed, want := range cases {
		if _, got := (Cannon{Height: 0.1, Speed: speed}).Path(); got != want {
			t.Errorf("speed %v: outcome %v, want %v", speed, got, want)
		}
	}
}

func TestBalmerAlpha(t *testing.T) {
	if nm := PhotonWavelength(3, 2); math.Abs(nm-656) > 2 {
		t.Errorf("H-alpha = %v nm", nm)
	}
	if SpectrumBand(PhotonWavelength(2, 1)) != "ultraviolet" {
		t.Error("Lyman-alpha should be ultraviolet")
	}
}

func TestDecayHalves(t *testing.T) {
	s := NewDecaySample(20000, 1)
	if r := float64(s.Remaining(1)) / 20000; math.Abs(r-0.5) > 0.02 {
		t.Errorf("after one half-life %.3f remain", r)
	}
}

func TestGasPressureGrowsWithTemperature(t *testing.T) {
	g := NewGas(200, 1)
	for i := 0; i < 2000; i++ {
		g.Step(0.005, 1)
	}
	cold := g.Pressure()
	for i := 0; i < 2000; i++ {
		g.Step(0.005, 4)
	}
	if hot := g.Pressure(); hot < 1.5*cold {
		t.Errorf("pressure cold %v hot %v", cold, hot)
	}
}

func TestFieldPointsAwayFromPositive(t *testing.T) {
	ex, _ := FieldAt(1, 0, []Charge{{0, 0, 1}})
	if ex <= 0 {
		t.Error("field should point away from a positive charge")
	}
	if g := Gamma(0.6); math.Abs(g-1.25) > 1e-9 {
		t.Errorf("gamma(0.6) = %v", g)
	}
}
