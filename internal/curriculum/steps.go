package curriculum

import (
	"errors"
	"strings"
)

const stepSeparator = "---"

// parseSteps reads a derivation: steps separated by ---, each with a
// "$ equation" line and Markdown explanation.
func parseSteps(lines []string) ([]Step, error) {
	var steps []Step
	var cur Step
	var text []string
	flush := func() {
		cur.Explanation = strings.TrimSpace(strings.Join(text, "\n"))
		if cur.Equation != "" || cur.Explanation != "" {
			steps = append(steps, cur)
		}
		cur, text = Step{}, nil
	}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch {
		case trimmed == stepSeparator:
			flush()
		case strings.HasPrefix(trimmed, "$ ") && cur.Equation == "":
			cur.Equation = strings.TrimSpace(trimmed[2:])
		default:
			text = append(text, line)
		}
	}
	flush()
	if len(steps) == 0 {
		return nil, errors.New("steps block is empty")
	}
	return steps, nil
}
