package curriculum

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
)

// IndexFile and LessonsDir locate the outline and lesson files in the
// content file system.
const (
	IndexFile  = "index.json"
	LessonsDir = "lessons"
)

// LoadCourse reads the outline from fsys.
func LoadCourse(fsys fs.FS) (*Course, error) {
	data, err := fs.ReadFile(fsys, IndexFile)
	if err != nil {
		return nil, fmt.Errorf("read course outline: %w", err)
	}
	var c Course
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("parse course outline: %w", err)
	}
	return &c, nil
}

// LoadLesson reads and parses a written lesson's file.
func LoadLesson(fsys fs.FS, l LessonInfo) ([]Block, error) {
	if !l.Written() {
		return nil, fmt.Errorf("lesson %q is not written yet", l.Title)
	}
	data, err := fs.ReadFile(fsys, path.Join(LessonsDir, l.File))
	if err != nil {
		return nil, fmt.Errorf("read lesson %q: %w", l.Title, err)
	}
	blocks, err := Parse(string(data))
	if err != nil {
		return nil, fmt.Errorf("lesson %s: %w", l.File, err)
	}
	return blocks, nil
}

// Chapter returns the chapter with the given id.
func (c *Course) Chapter(id string) (Chapter, bool) {
	for _, p := range c.Parts {
		for _, ch := range p.Chapters {
			if ch.ID == id {
				return ch, true
			}
		}
	}
	return Chapter{}, false
}

// Next returns the lesson after index i in the chapter, if any.
func (ch Chapter) Next(i int) (LessonInfo, int, bool) {
	if i+1 < len(ch.Lessons) {
		return ch.Lessons[i+1], i + 1, true
	}
	return LessonInfo{}, 0, false
}

// WrittenCount returns how many of the chapter's lessons have content.
func (ch Chapter) WrittenCount() int {
	n := 0
	for _, l := range ch.Lessons {
		if l.Written() {
			n++
		}
	}
	return n
}
