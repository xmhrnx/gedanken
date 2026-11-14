package theme

import (
	"fmt"
	"math"

	"github.com/xmhrnx/giode"
)

// Blend mixes two "#rrggbb" colors: t = 0 gives a, t = 1 gives b. It
// returns an opaque color, which composites predictably on the GPU.
func Blend(a, b string, t float64) giode.Color {
	t = math.Max(0, math.Min(1, t))
	var ar, ag, ab, br, bg, bb int
	fmt.Sscanf(a, "#%02x%02x%02x", &ar, &ag, &ab)
	fmt.Sscanf(b, "#%02x%02x%02x", &br, &bg, &bb)
	mix := func(x, y int) int { return int(float64(x) + (float64(y)-float64(x))*t) }
	return giode.Color(fmt.Sprintf("#%02x%02x%02x", mix(ar, br), mix(ag, bg), mix(ab, bb)))
}

// Night is the darkest backdrop color, the base of drawings.
const Night = "#0b0f1c"
