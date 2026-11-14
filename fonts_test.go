package main

import (
	"testing"

	"github.com/xmhrnx/giode/fonts"
)

// TestMathFontRegistered checks that formulas can find STIX Two Math
// after the app sets up its fonts. Without it, phones without a system
// math font show the italic variables of every formula as boxes.
func TestMathFontRegistered(t *testing.T) {
	if err := setupFonts(); err != nil {
		t.Fatal(err)
	}
	for _, f := range fonts.Collection() {
		if string(f.Font.Typeface) == "STIX Two Math" {
			return
		}
	}
	t.Fatal("STIX Two Math is not in the font collection")
}
