package curriculum

import (
	"fmt"
	"strings"
)

// fence is an open ::: block while parsing.
type fence struct {
	kind  string
	arg   string
	line  int
	lines []string
}

const fenceMark = ":::"

// Parse splits a lesson file into blocks. Markdown outside fences
// becomes prose blocks.
func Parse(src string) ([]Block, error) {
	var blocks []Block
	var prose []string
	var open *fence

	flushProse := func() {
		if text := strings.TrimSpace(strings.Join(prose, "\n")); text != "" {
			blocks = append(blocks, Block{Kind: KindProse, Text: text})
		}
		prose = nil
	}

	for i, line := range strings.Split(strings.ReplaceAll(src, "\r\n", "\n"), "\n") {
		trimmed := strings.TrimSpace(line)
		switch {
		case open == nil && strings.HasPrefix(trimmed, fenceMark) && trimmed != fenceMark:
			flushProse()
			kind, arg, _ := strings.Cut(strings.TrimSpace(strings.TrimPrefix(trimmed, fenceMark)), " ")
			open = &fence{kind: kind, arg: strings.TrimSpace(arg), line: i + 1}
		case open != nil && trimmed == fenceMark:
			b, err := buildBlock(open)
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, b)
			open = nil
		case open != nil:
			open.lines = append(open.lines, line)
		case open == nil && strings.HasPrefix(trimmed, "## "):
			// Each section is its own block, so long lessons can be laid
			// out a section at a time as the reader scrolls.
			flushProse()
			prose = append(prose, line)
		default:
			prose = append(prose, line)
		}
	}
	if open != nil {
		return nil, fmt.Errorf("line %d: ::: %s is never closed", open.line, open.kind)
	}
	flushProse()
	return blocks, nil
}

// buildBlock turns a closed fence into a block.
func buildBlock(f *fence) (Block, error) {
	body := strings.TrimSpace(strings.Join(f.lines, "\n"))
	b := Block{Arg: f.arg, Text: body}
	var err error
	switch f.kind {
	case "equation":
		b.Kind = KindEquation
	case "note":
		b.Kind = KindNote
	case "history":
		b.Kind = KindHistory
	case "model":
		b.Kind = KindModel
		if f.arg == "" {
			err = fmt.Errorf("model needs an id")
		}
	case "steps":
		b.Kind = KindSteps
		b.Steps, err = parseSteps(f.lines)
	case "quiz":
		b.Kind = KindQuiz
		b.Questions, err = parseQuiz(f.lines)
	default:
		err = fmt.Errorf("unknown block kind %q", f.kind)
	}
	if err != nil {
		return Block{}, fmt.Errorf("line %d: %w", f.line, err)
	}
	return b, nil
}
