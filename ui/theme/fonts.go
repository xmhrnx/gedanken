package theme

import (
	"fmt"

	"github.com/go-fonts/stix/stix2textbold"
	"github.com/go-fonts/stix/stix2textitalic"
	"github.com/go-fonts/stix/stix2textregular"
	"github.com/go-fonts/stix/stix2textsemibold"
	"github.com/xmhrnx/giode"
)

// Serif is the book face used for titles and reading: STIX Two Text,
// the text companion of the STIX Two Math font the equations use.
const Serif = "STIX Two Text"

// RegisterFonts adds the serif family next to the Go fonts.
func RegisterFonts() error {
	err := giode.RegisterFont(
		giode.FontFile{Family: Serif, Data: stix2textregular.TTF},
		giode.FontFile{Family: Serif, Italic: true, Data: stix2textitalic.TTF},
		giode.FontFile{Family: Serif, Weight: giode.FontWeightSemiBold, Data: stix2textsemibold.TTF},
		giode.FontFile{Family: Serif, Weight: giode.FontWeightBold, Data: stix2textbold.TTF},
	)
	if err != nil {
		return fmt.Errorf("register %s: %w", Serif, err)
	}
	return nil
}
