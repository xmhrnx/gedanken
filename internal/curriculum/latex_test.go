package curriculum

import "testing"

func TestCheckLatex(t *testing.T) {
	good := []string{`v = \sqrt{\frac{2GM}{r}}`, `E &= \tfrac{1}{2}mv^2 \\ &= x`, `c = 3\,\text{m/s}`, `\left( a \right)`, `0.65\%`, `a\;b\!c`}
	for _, s := range good {
		if err := CheckLatex(s); err != nil {
			t.Errorf("%s: %v", s, err)
		}
	}
	bad := []string{`\frac{a}{b`, `a}`, `\fraction{a}`, `\mathcal{L}`, `50%`, `\@`}
	for _, s := range bad {
		if CheckLatex(s) == nil {
			t.Errorf("%s: expected an error", s)
		}
	}
}
