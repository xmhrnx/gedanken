# Writing lessons

Read this with FORMAT.md before writing or editing a lesson. Good examples of the finished
standard: `lessons/black-holes/dark-stars.md`, `lessons/fission/discovery.md`,
`lessons/motion/projectiles.md`.

## Reader

An intelligent adult who reads English well. Not PhD level, not "explain like I'm 5".
They want to understand how physicists reached each idea, with the math shown and the
history alongside.

## Voice

- Concise and direct. No filler, no hype, no rhetorical warm-ups ("Have you ever wondered…"),
  no praise of the reader.
- At most one analogy, and only when it saves time. Never a stack of metaphors.
- Define an unfamiliar term in one short clause the first time it appears.
- Don't skip topics; don't go deeper than the lesson needs.

## Structure of a lesson

- 400–800 words of prose, split by `##` subheadings.
- At least one `::: equation` or `::: steps` block. Derivations go in `::: steps`: each step
  has a `$ ` LaTeX line and 1–3 sentences on why that step is allowed. Show every step a
  reader needs; skip none.
- One `::: note` with the lesson's key idea.
- `::: history` blocks where the history explains how the idea arose: who, when, what they
  saw, what they got wrong.
- A `::: model` where one of the models in FORMAT.md fits, with a one-line caption saying what
  to do or watch. Use only ids from the table.
- End with `::: quiz`: 2–3 questions, exactly one `+` each, each with an `=` explanation.
  Questions test understanding, not recall of dates.

## Accuracy

- Check every number, date and name. Round sensibly and say "about" when rounding.
- When a simple picture is wrong in detail (Bohr orbits, "virtual particle pairs"), say so in
  one sentence.

## Notation

- Equations are LaTeX (see FORMAT.md). Prose, notes, history and quizzes are plain text with
  Unicode maths (`F = ma`, `v²`, `10⁻³⁴`, `ħ`, `∇`, `∝` all render).
- Keep `\text{…}` words in equations to a minimum: equations are maths, prose explains.
