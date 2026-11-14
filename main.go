// Gedanken teaches physics, from motion to black holes.
package main

import (
	"image"
	"log"

	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/formula"

	"gedanken/assets"
	"gedanken/content"
	"gedanken/internal/curriculum"
	"gedanken/ui"
	"gedanken/ui/theme"
)

func main() {
	if err := setupFonts(); err != nil {
		log.Fatal(err)
	}
	theme.Apply()
	if err := giode.RegisterIconsFS(assets.Icons, "icons"); err != nil {
		log.Fatalf("register icons: %v", err)
	}
	course, err := curriculum.LoadCourse(content.FS)
	if err != nil {
		log.Fatal(err)
	}
	root := ui.NewApp(content.FS, course)

	app := giode.New(giode.Config{
		Title:      "Gedanken",
		Size:       image.Pt(400, 840),
		Resizable:  true,
		Background: "theme.background",
	})
	if err := app.Run(root.View); err != nil {
		log.Fatal(err)
	}
}

// setupFonts installs the Go fonts, the serif and the math font.
// UseGofont replaces the font collection, which drops the STIX Two Math
// font the formula package registered when it was imported, so it is
// registered again; without it, phones show formula variables as boxes.
func setupFonts() error {
	giode.UseGofont()
	formula.Register()
	return theme.RegisterFonts()
}
