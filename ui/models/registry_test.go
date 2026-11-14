package models

import (
	"testing"

	"gedanken/internal/curriculum"
)

func TestEveryModelIDHasAModel(t *testing.T) {
	for _, id := range curriculum.ModelIDs {
		if _, ok := New(id); !ok {
			t.Errorf("no model for %q", id)
		}
	}
}
