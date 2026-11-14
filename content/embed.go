// Package content embeds the course outline and lesson files, so they
// ship inside the binary on every platform.
package content

import "embed"

// FS holds index.json and lessons/.
//
//go:embed index.json lessons
var FS embed.FS
