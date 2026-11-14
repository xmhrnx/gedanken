# Writing lessons

Read this with FORMAT.md before writing or editing a lesson. The reference for the voice is
`lessons/atom/electron.md`: read it first, and match its pace.

## Reader

An intelligent adult who reads English well and wants to understand how physicists reached
each idea. Not a specialist, not a child. They will follow a careful argument gladly; what
they dislike is being rushed, talked down to, or buried in metaphors.

## Voice: a good lecturer explaining to a smart friend

The lesson should read as if someone who loves the subject were sitting with the reader and
explaining it properly, with time to spare. That means:

- **Set up the problem before solving it.** Say what was puzzling, what people expected, and
  why the obvious approach fails. A result lands only when the reader feels the question.
- **Explain the reasoning, not only the result.** Every claim that matters gets its "because".
  If a force acts sideways, say why that bends a path into a circle. If a quantity cancels,
  say why that is useful.
- **Write in connected paragraphs.** Sentences lead into each other ("So…", "That is why…",
  "Now suppose…"). No lists of facts, no clipped sentences, no telegraphic fragments.
- **Pause on results.** After an equation or a number, say what it means, what it implies,
  and whether it surprised anyone.
- **Tell history as a story of ideas.** Who was stuck on what, what they tried, what they got
  wrong, what changed their minds. Dates and names serve the story; never list them.
- **Keep the language plain.** Define every technical term in a short clause the first time.
  Use everyday words where they are accurate.
- **Use at most one analogy per idea**, only where it truly helps, and say where it breaks.
- **No filler and no condescension.** No "Have you ever wondered…", no "Simply put", no
  praise of the reader, no "as we all know". Warmth comes from care in explaining, not from
  chatter.

## Structure of a lesson

- About 1,200–1,800 words of prose, split by `##` subheadings that read like a sequence of
  questions answered.
- Open with a paragraph that says what the lesson is about and why it matters.
- Derivations go in `::: steps`. Each step has a `$ ` LaTeX line and an explanation of 2–5
  sentences: what we did, why it is allowed, and what it achieved. Never "rearrange" or
  "substitute" alone; say what is being rearranged and why.
- Worked numbers where they make an idea concrete, with units and a sentence on what the
  number means.
- One `::: note` with the lesson's key idea, in two or three full sentences.
- `::: history` blocks where the history explains how the idea arose: written as a short
  story, not a list.
- A `::: model` where one of the models in FORMAT.md fits, with a caption saying what to do
  and what to notice.
- End with `::: quiz`: 2–3 questions, exactly one `+` each, each with an `=` explanation of
  one or two full sentences. Questions test understanding, not recall of dates.

## Accuracy

- Check every number, date and name. Round sensibly and say "about" when rounding.
- When a simple picture is wrong in detail (Bohr orbits, "virtual particle pairs"), say so and
  say what replaces it.

## Notation

- Equations are LaTeX (see FORMAT.md). Prose, notes, history and quizzes are plain text with
  Unicode maths (`F = ma`, `v²`, `10⁻³⁴`, `ħ`, `∇`, `∝` all render).
- Keep `\text{…}` words in equations to a minimum: equations are maths, prose explains.
