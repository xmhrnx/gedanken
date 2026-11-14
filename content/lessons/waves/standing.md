Pluck a guitar string and it does not send a wave running off into the distance. The string is held at both ends, so waves reflect back and forth, and the overlapping reflections settle into a pattern that vibrates in place. This is a **standing wave**, and it is why a string, a pipe or a drum produces particular notes rather than any frequency at all.

## Two waves, opposite directions

A wave reaching a fixed end reflects, flipped upside down, and travels back. On a string fixed at both ends you therefore have identical waves travelling in opposite directions, superposed.

::: steps Adding waves travelling in opposite directions
$ y_1 = A\sin(kx - \omega t), \qquad y_2 = A\sin(kx + \omega t)
Two waves of equal amplitude and frequency, the first moving toward +x and the second toward −x. Here k = 2π/λ and ω = 2πf, shorthand for the factors in the previous lesson's wave formula.
---
$ y = y_1 + y_2
Superposition: the displacements add.
---
$ \sin(a - b) + \sin(a + b) = 2\sin a\cos b
A standard trigonometric identity: expand both sines with the angle-sum formula, and the cos a sin b terms cancel.
---
$ y = 2A\sin(kx)\cos(\omega t)
Apply it with a = kx and b = ωt. The result no longer has the form f(x − vt): nothing travels.
:::

The result is a fixed shape, 2A sin(kx), whose size swells and shrinks in time as cos(ωt). Every point oscillates in step with the others, each with its own amplitude. Points where sin(kx) = 0 never move at all; these are **nodes**, spaced half a wavelength apart. Halfway between them are **antinodes**, where the string swings most.

## Which notes a string can play

A string of length L fixed at both ends must have a node at each end. That means a whole number of half-wavelengths must fit exactly on the string.

::: steps The allowed frequencies of a string
$ L = n\frac{\lambda_n}{2}, \qquad n = 1, 2, 3, \dots
A whole number n of half-wavelengths, each the distance between neighbouring nodes, spans the string.
---
$ \lambda_n = \frac{2L}{n}
Solve for the allowed wavelengths.
---
$ f_n = \frac{v}{\lambda_n} = \frac{nv}{2L}
Use v = fλ, where v is the wave speed on the string.
---
$ f_n = \frac{n}{2L}\sqrt{\frac{F}{\mu}}
Insert v = √(F/μ) for a string with tension F and mass per unit length μ.
:::

The lowest, n = 1, is the **fundamental**: half a wavelength fills the string, which bulges in a single loop. The others, at 2, 3, 4… times the fundamental frequency, are **harmonics** (or overtones). The formula says how to change the note: shorten the string (pressing it against a fret), tighten it (the tuning peg) or use a thicker, heavier string (the bass strings).

::: note
Fixing the ends allows only waves that fit: a whole number of half-wavelengths. So a string has a set of natural frequencies, f, 2f, 3f…, and these set the notes it can play.
:::

As an example, a guitar's vibrating length is about 0.65 m and its lowest string is tuned to 82.4 Hz. The wave speed on that string must be v = 2Lf ≈ 107 m/s. A real pluck excites the fundamental and many harmonics at once, and the mixture of their strengths gives each instrument its own tone, which is why a guitar and a piano playing the same note sound different.

Air in a pipe does the same thing with sound. A pipe open at both ends follows the same rule, f = nv/(2L). A pipe closed at one end has a node at the closed end and an antinode at the open one, so it fits odd quarter-wavelengths and plays only odd harmonics; it also sounds an octave lower than an open pipe of the same length.

::: history 1636–1701 · Paris
Pythagorean tradition held that halving a string raises its note by an octave, and that simple ratios of lengths give pleasing intervals. In *Harmonie universelle* (1636), the friar Marin Mersenne turned this into quantitative laws: frequency goes inversely with length, as the square root of tension, and inversely as the square root of the mass per length. He measured the frequency of a very long, slowly vibrating string by eye and scaled down. Around 1700 Joseph Sauveur heard that a string sounds several pitches at once, introduced the words "node" and "harmonic", and showed the harmonics are whole-number multiples of the fundamental.
:::

## Resonance

Every object that can vibrate has natural frequencies, its standing-wave patterns. Push it periodically at one of those frequencies and each push adds to the motion, so the amplitude builds up: **resonance**. A child on a swing pumped once per swing goes higher and higher. A singer can shatter a wine glass by holding its natural note loudly. A guitar's hollow body resonates with the string and makes the sound loud enough to hear; a bare string moves too little air.

Resonance is often blamed for the 1940 collapse of the Tacoma Narrows Bridge, but that was a different effect: steady wind fed energy into a twisting motion of the deck (aeroelastic flutter), not a periodic push at a natural frequency. Engineers still design buildings and bridges so that wind, footsteps and earthquakes do not drive them near their natural frequencies.

::: quiz
? A string's fundamental frequency is 100 Hz. What is the frequency of its third harmonic?
- 133 Hz
+ 300 Hz
- 900 Hz
= Harmonics are whole-number multiples: f₃ = 3 × 100 Hz.
? Pressing a guitar string against a fret halfway along its length makes the note
+ go up by an octave, doubling the frequency
- go down by an octave
- stay the same but get louder
= f = v/(2L), so halving L doubles f, which is an octave higher.
? How far apart are neighbouring nodes in a standing wave?
- One wavelength
+ Half a wavelength
- A quarter of a wavelength
= Nodes are where sin(kx) = 0, which happens every half wavelength.
:::
