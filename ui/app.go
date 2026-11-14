// Package ui is the app's interface: the root view and navigation.
package ui

import (
	"io/fs"

	"github.com/xmhrnx/giode"
	"github.com/xmhrnx/giode/settings"

	"gedanken/internal/curriculum"
	"gedanken/internal/progress"
	"gedanken/ui/screens"
	"gedanken/ui/theme"
)

// progressKey is where finished lessons are saved.
const progressKey = "progress"

// Route paths.
const (
	routeHome    = "/"
	routeChapter = "/chapter"
	routeLesson  = "/lesson"
)

// App owns the course, the reader's progress and the open screens.
type App struct {
	content fs.FS
	course  *curriculum.Course
	tracker *progress.Tracker
	router  *giode.Router
	home    *screens.Home
	chapter *screens.ChapterScreen
	lesson  *screens.LessonScreen
	part    int
	current curriculum.Chapter
}

// NewApp builds the app over the course stored in content.
func NewApp(content fs.FS, course *curriculum.Course) *App {
	saved := settings.Value(progressKey, map[string]bool{})
	tracker := progress.New(saved, func(done map[string]bool) error { return settings.Save(progressKey, done) })
	a := &App{content: content, course: course, tracker: tracker}
	a.home = screens.NewHome(course, a.tracker, a.openChapter)
	a.router = giode.NewRouter(
		giode.Route{Path: routeHome, View: func() giode.Element { return a.home }},
		giode.Route{Path: routeChapter, View: func() giode.Element { return a.chapter }},
		giode.Route{Path: routeLesson, View: func() giode.Element { return a.lesson }},
	)
	return a
}

// View is the root element, built every frame.
func (a *App) View() giode.Element {
	return giode.Box(giode.Styles{Width: -1, Height: -1, Background: theme.Background}, a.router)
}

func (a *App) openChapter(part int, ch curriculum.Chapter) {
	a.part, a.current = part, ch
	a.chapter = screens.NewChapterScreen(part, ch, a.tracker, a.back, a.openLesson)
	a.router.Push(routeChapter)
}

func (a *App) openLesson(i int) {
	a.lesson = screens.NewLessonScreen(a.content, a.part, a.current, i, a.tracker, a.back, a.nextLesson)
	a.router.Push(routeLesson)
}

// nextLesson replaces the open lesson, so Back still leads to the chapter.
func (a *App) nextLesson(i int) {
	a.lesson = screens.NewLessonScreen(a.content, a.part, a.current, i, a.tracker, a.back, a.nextLesson)
}

func (a *App) back() { a.router.Back() }
