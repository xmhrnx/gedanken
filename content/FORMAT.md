# Lesson file format

Each lesson is one Markdown file under `content/lessons/<chapter>/<lesson>.md`, listed in
`content/index.json` by its `file` field. Lessons without a `file` show as "not written yet".

A file is ordinary Markdown (CommonMark: paragraphs, **bold**, *italic*, lists, tables) with
fenced blocks that start with `::: kind [argument]` on their own line and end with `:::`.
The lesson title comes from `index.json`, so the file has no `#` heading.

## Blocks

### Prose
Any Markdown outside a fenced block. `##` subheadings are allowed.

### Equation
A displayed equation, one or more lines. The optional argument is a caption.

```
::: equation Newton's second law
F = ma
:::
```

### Steps (a derivation)
A titled derivation. Steps are separated by a line holding only `---`. In each step the
first line starting with `$ ` is the equation; the rest is Markdown explaining it.

```
::: steps Deriving escape velocity
$ ½mv² = GMm/r
Kinetic energy at launch must equal the energy needed to climb out of the well.
---
$ v = √(2GM/r)
Cancel m, multiply by 2, take the square root.
:::
```

### Note
A highlighted key idea.

```
::: note
Mass cancels: the escape velocity is the same for a pebble and a planet.
:::
```

### History
A historical aside. The argument is a short label (year, place).

```
::: history 1783 · Thornhill, England
John Michell, a country rector, ...
:::
```

### Model
An interactive graphic. The argument is the model id; the body is an optional caption.

```
::: model escape
Drag the slider to shrink the star.
:::
```

Model ids and what they show:

| id | Shows |
|---|---|
| `slope` | tangent to a curve at a draggable point |
| `area` | rectangles approximating the area under a curve |
| `projectile` | launch angle and speed, trajectory |
| `falling` | ball dropped, marked every 0.1 s |
| `escape` | a star shrinking until its escape velocity passes c |
| `gravity-well` | curved space around a black hole, clock rate at a probe |
| `light-cones` | light cones tipping over near a horizon |
| `binding-curve` | binding energy per nucleon against mass number |
| `chain-reaction` | generations of fission for a chosen k |
| `critical-mass` | neutrons escaping a sphere of chosen size |
| `atom` | animated nucleus with orbiting electrons |
| `orbit` | Newton's cannon: launch speed slider; falls back, circles, ellipses, escapes |
| `pendulum` | swinging pendulum; length slider, period shown |
| `wave` | travelling wave; wavelength and frequency sliders, speed shown |
| `interference` | two sources or slits; separation slider, bright and dark fringes |
| `gas` | molecules bouncing in a box; temperature slider, pressure shown |
| `field-lines` | electric field of two charges; second charge's sign and size slider |
| `light-clock` | a moving light clock; speed slider, time dilation factor γ |
| `bohr` | hydrogen energy levels; pick a jump, see the photon's wavelength and colour |
| `wavefunction` | particle in a box; level n slider, ψ and probability |ψ|² |
| `decay` | radioactive nuclei decaying; count halves every half-life |

### Quiz
One or more questions. Each starts with `? `, options are `- ` (wrong) or `+ ` (right,
exactly one), and an optional `= ` line explains the answer. Readers can skip quizzes.

```
::: quiz
? What happens to escape velocity if the radius halves at the same mass?
- It halves
+ It grows by √2
- It doubles
= v ~ 1/√r, so halving r multiplies v by √2.
:::
```

## Writing equations

Equation bodies and the `$ ` line of each derivation step are **LaTeX**, typeset by Giode's
`formula` package. Write them as you would inside `$$ … $$`, without the dollar signs:

```
::: equation Escape velocity
v = \sqrt{\frac{2GM}{r}}
:::

::: equation Two lines, aligned at the equals signs
E &= \tfrac{1}{2}mv^2 \\ &= \frac{p^2}{2m}
:::
```

- One formula per block. Break lines with `\\` and align them with `&`.
- Letters are italic variables. Use `\mathrm{d}x` for an upright d, `\text{ km/s}` for units
  and words, `\,` for a thin space before a unit: `c = 3.00 \times 10^8\,\text{m/s}`.
- Supported: `\frac \dfrac \tfrac \sqrt \sqrt[n]`, `^ _`, `\sum \int \oint \lim`, `\left( \right)`,
  `\vec \hat \bar \dot`, `\text \mathrm \mathbf`, Greek letters, and operators such as
  `\nabla \partial \hbar \propto \Rightarrow \approx \leq \pm \times \cdot \infty \to`.
  Anything else is drawn as written, so `go test ./content/` checks commands and braces.
- Prose, notes, history and quizzes stay plain text: write short inline maths in Unicode
  (`v = √(2GM/r)`, `r_s`, `10²⁴`). Every symbol renders, ∇ ℏ ∝ included.
