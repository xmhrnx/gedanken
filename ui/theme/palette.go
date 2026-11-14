// Package theme holds the app's colors, sizes and shared styles.
package theme

import "github.com/xmhrnx/giode"

// Apply installs the dark "deep space" palette and makes it active.
func Apply() {
	giode.SetPalette(giode.ModeDark, giode.Palette{
		Primary:    "#8ea8ff",
		Secondary:  "#5eead4",
		Background: "#0a0d16",
		Surface:    "#131a2a",
		Text:       "#e8ecf6",
		Border:     "#25304a",
		OnPrimary:  "#0a0d16",
		Muted:      "#8d96ab",
		Accent:     "#f6bd6e",
		Success:    "#5ee39a",
		Warning:    "#f6bd6e",
		Danger:     "#f58a8a",
		Info:       "#6ee7f0",
	})
	giode.SetMode(giode.ModeDark)
}

// Theme tokens used across the app.
const (
	Background = "theme.background"
	Surface    = "theme.surface"
	Text       = "theme.text"
	Muted      = "theme.muted"
	Border     = "theme.border"
	Primary    = "theme.primary"
	OnPrimary  = "theme.onprimary"
	Accent     = "theme.accent"
	Success    = "theme.success"
	Danger     = "theme.danger"
	Info       = "theme.info"
)

// Tints mixed from theme tokens, so they follow palette changes.
var (
	SurfaceRaised = giode.Mix("theme.text", "theme.surface", 0.04)
	PrimarySoft   = giode.Mix("theme.primary", "theme.surface", 0.14)
	AccentSoft    = giode.Mix("theme.accent", "theme.surface", 0.12)
	SuccessSoft   = giode.Mix("theme.success", "theme.surface", 0.16)
	DangerSoft    = giode.Mix("theme.danger", "theme.surface", 0.16)
	Faint         = giode.Mix("theme.text", "theme.background", 0.18)
)

// partColors gives each part of the course its own hue.
var partColors = []giode.Color{
	"#8ea8ff", "#5eead4", "#fb9a5b", "#f6d365",
	"#c49bff", "#f48fc4", "#f58a8a", "#6cb8ff",
}

// PartColor returns the accent color of the i-th part.
func PartColor(i int) giode.Color {
	return partColors[i%len(partColors)]
}
