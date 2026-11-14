// Package assets embeds the app's icons.
package assets

import "embed"

// Icons holds the SVG icons under icons/.
//
//go:embed icons
var Icons embed.FS
