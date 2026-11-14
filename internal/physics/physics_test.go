package physics

import (
	"math"
	"testing"
)

func near(a, b, tol float64) bool { return math.Abs(a-b) <= tol*math.Abs(b) }

func TestProjectile45Degrees(t *testing.T) {
	p := Projectile{Speed: 20, Angle: 45, G: 9.8}
	if !near(p.Range(), 400/9.8, 1e-9) {
		t.Errorf("range %v", p.Range())
	}
	if x, y := p.At(p.FlightTime()); !near(x, p.Range(), 1e-9) || math.Abs(y) > 1e-9 {
		t.Errorf("lands at %v,%v", x, y)
	}
}

func TestSunHorizonIsAboutThreeKilometres(t *testing.T) {
	if rs := SchwarzschildRadius(SolarMass); !near(rs, 2953, 0.01) {
		t.Errorf("rs = %v m", rs)
	}
	if v := EscapeVelocity(SolarMass, SchwarzschildRadius(SolarMass)); !near(v, C, 1e-9) {
		t.Errorf("escape at horizon = %v", v)
	}
}

func TestIronPeak(t *testing.T) {
	if n := MostBound(); n.A < 56 || n.A > 62 {
		t.Errorf("peak at %s", n.Name)
	}
}

func TestChainGrowsWithKAboveOne(t *testing.T) {
	c := NewChainReaction(40, 30, 2, 1)
	for i := 0; i < 5; i++ {
		c.Step()
	}
	if len(c.Fissions) < 20 {
		t.Errorf("only %d fissions after 6 generations at k=2", len(c.Fissions))
	}
}

func TestBiggerSphereLeaksLess(t *testing.T) {
	small := SampleSphere(1, 20000, 0, 1).EscapeFrac
	big := SampleSphere(6, 20000, 0, 1).EscapeFrac
	if !(small > big) || !near(big, 3.0/(4*6), 0.25) {
		t.Errorf("escape small=%v big=%v", small, big)
	}
}
