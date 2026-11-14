package curriculum

import (
	"fmt"
	"strings"
)

// parseQuiz reads "? prompt", "- wrong", "+ right" and "= why" lines.
func parseQuiz(lines []string) ([]Question, error) {
	var qs []Question
	var cur *Question
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if strings.HasPrefix(trimmed, "? ") {
			qs = append(qs, Question{Prompt: strings.TrimSpace(trimmed[2:]), Answer: -1})
			cur = &qs[len(qs)-1]
			continue
		}
		if cur == nil {
			return nil, fmt.Errorf("quiz line %q comes before any question", trimmed)
		}
		switch {
		case strings.HasPrefix(trimmed, "+ "):
			if cur.Answer >= 0 {
				return nil, fmt.Errorf("question %q has more than one right answer", cur.Prompt)
			}
			cur.Answer = len(cur.Options)
			cur.Options = append(cur.Options, strings.TrimSpace(trimmed[2:]))
		case strings.HasPrefix(trimmed, "- "):
			cur.Options = append(cur.Options, strings.TrimSpace(trimmed[2:]))
		case strings.HasPrefix(trimmed, "= "):
			cur.Explanation = strings.TrimSpace(trimmed[2:])
		default:
			cur.Prompt += " " + trimmed
		}
	}
	for _, q := range qs {
		if q.Answer < 0 || len(q.Options) < 2 {
			return nil, fmt.Errorf("question %q needs two or more options and one right answer", q.Prompt)
		}
	}
	if len(qs) == 0 {
		return nil, fmt.Errorf("quiz block is empty")
	}
	return qs, nil
}
