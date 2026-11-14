// Gedanken teaches physics, from motion to black holes.
package main

import (
	"image"
	"log"

	"github.com/xmhrnx/giode"

	"gedanken/assets"
	"gedanken/content"
	"gedanken/internal/curriculum"
	"gedanken/ui"
	"gedanken/ui/theme"
)

func main() {
	giode.UseGofont()
	theme.Apply()
	if err := theme.RegisterFonts(); err != nil {
		log.Fatal(err)
	}
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
