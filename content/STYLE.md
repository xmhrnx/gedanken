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

A lesson is a full treatment of its topic, not a summary: about 2,500–3,500 words of prose,
in 7–10 `##` sections that read like a sequence of questions answered. Depth comes from
explaining more of the physics, not from repeating it.

- **Opening:** a paragraph that says what the lesson is about, why it matters, and what the
  reader will be able to explain by the end.
- **The puzzle:** what people knew, what they expected, and what did not fit.
- **The core idea, built up carefully:** the main argument or derivation, in full.
- **Worked examples:** at least two, with real numbers and units, each followed by a sentence
  on what the number means. Choose examples that show different sides of the idea.
- **Going deeper:** at least one section that takes the idea further than an introduction
  would: a more general form, a second derivation from a different angle, a subtlety, or the
  condition under which it holds.
- **Common misconceptions:** where readers often go wrong, and why the wrong picture fails.
- **Limits:** where the idea stops working and what replaces it there.
- **Where it shows up:** applications, in nature and technology, and how this lesson connects
  to earlier and later ones (name them).
- **History:** `::: history` blocks told as stories, where they explain how the idea arose.
- Derivations go in `::: steps`. Each step has a `$ ` LaTeX line and an explanation of 2–5
  sentences: what we did, why it is allowed, and what it achieved. Never "rearrange" or
  "substitute" alone; say what is being rearranged and why.
- One `::: note` with the lesson's key idea, in two or three full sentences.
- A `::: model` where one of the models in FORMAT.md fits, with a caption saying what to do
  and what to notice.
- End with `::: quiz`: 4–5 questions, exactly one `+` each, each with an `=` explanation of
  one or two full sentences. Questions test understanding and the ability to apply the idea
  to a new case, not recall of dates.

## Accuracy

- Check every number, date and name. Round sensibly and say "about" when rounding.
- When a simple picture is wrong in detail (Bohr orbits, "virtual particle pairs"), say so and
  say what replaces it.

## Notation

- Equations are LaTeX (see FORMAT.md). Prose, notes, history and quizzes are plain text with
  Unicode maths (`F = ma`, `v²`, `10⁻³⁴`, `ħ`, `∇`, `∝` all render).
- Keep `\text{…}` words in equations to a minimum: equations are maths, prose explains.
