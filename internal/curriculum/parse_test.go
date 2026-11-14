package curriculum

import "testing"

const sample = `Intro with **bold**.

::: equation Newton
F = ma
:::

::: steps Escape
$ ½mv² = GMm/r
Energy balance.
---
$ v = √(2GM/r)
Solve for v.
:::

::: history 1783 · England
Michell.
:::

::: model escape
Drag it.
:::

::: quiz
? Pick the right one
- wrong
+ right
= because
? Second
+ yes
- no
:::

Outro.`

func TestParseAllKinds(t *testing.T) {
	blocks, err := Parse(sample)
	if err != nil {
		t.Fatal(err)
	}
	want := []BlockKind{KindProse, KindEquation, KindSteps, KindHistory, KindModel, KindQuiz, KindProse}
	if len(blocks) != len(want) {
		t.Fatalf("got %d blocks, want %d", len(blocks), len(want))
	}
	for i, k := range want {
		if blocks[i].Kind != k {
			t.Errorf("block %d: kind %d, want %d", i, blocks[i].Kind, k)
		}
	}
	steps := blocks[2].Steps
	if len(steps) != 2 || steps[1].Equation != "v = √(2GM/r)" || steps[0].Explanation != "Energy balance." {
		t.Errorf("steps parsed wrong: %+v", steps)
	}
	q := blocks[5].Questions
	if len(q) != 2 || q[0].Answer != 1 || q[0].Explanation != "because" || q[1].Answer != 0 {
		t.Errorf("quiz parsed wrong: %+v", q)
	}
	if blocks[4].Arg != "escape" || blocks[3].Arg != "1783 · England" {
		t.Errorf("args parsed wrong: %q %q", blocks[4].Arg, blocks[3].Arg)
	}
}

func TestParseErrors(t *testing.T) {
	cases := map[string]string{
		"unclosed":     "::: note\nhello",
		"unknown kind": "::: banana\n:::",
		"two answers":  "::: quiz\n? q\n+ a\n+ b\n:::",
		"no answer":    "::: quiz\n? q\n- a\n- b\n:::",
		"empty steps":  "::: steps T\n:::",
		"model no id":  "::: model\n:::",
	}
	for name, src := range cases {
		if _, err := Parse(src); err == nil {
			t.Errorf("%s: expected an error", name)
		}
	}
}
