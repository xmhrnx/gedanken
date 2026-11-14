// Package curriculum loads the course outline and parses lesson files.
// It has no UI dependencies.
package curriculum

// Course is the whole outline: parts, their chapters and lessons.
type Course struct {
	Parts []Part `json:"parts"`
}

// Part groups related chapters ("Classical Mechanics").
type Part struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Chapters []Chapter `json:"chapters"`
}

// Chapter is one topic made of several lessons.
type Chapter struct {
	ID      string       `json:"id"`
	Title   string       `json:"title"`
	Blurb   string       `json:"blurb"`
	Lessons []LessonInfo `json:"lessons"`
}

// LessonInfo is a lesson's outline entry. File is empty for lessons
// that are not written yet.
type LessonInfo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	File    string `json:"file"`
}

// BlockKind says how a lesson block is presented.
type BlockKind int

// The kinds of lesson blocks; see content/FORMAT.md.
const (
	KindProse BlockKind = iota
	KindEquation
	KindSteps
	KindNote
	KindHistory
	KindModel
	KindQuiz
)

// Block is one piece of a lesson. Arg holds the caption, label or
// model id; Text the Markdown or equation body.
type Block struct {
	Kind      BlockKind
	Arg       string
	Text      string
	Steps     []Step
	Questions []Question
}

// Step is one line of a derivation and why it holds.
type Step struct {
	Equation    string
	Explanation string
}

// Question is one multiple-choice quiz question.
type Question struct {
	Prompt      string
	Options     []string
	Answer      int
	Explanation string
}

// Written reports whether the lesson has content.
func (l LessonInfo) Written() bool { return l.File != "" }
