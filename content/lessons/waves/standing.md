Pluck a guitar string and you might expect a wave to run along it and away, like a hump flicked along a rope. That is not what happens. The string is held fixed at both ends, so any wave on it reflects at each end and travels back, over and over. The reflections overlap, and within a fraction of a second they settle into a pattern that vibrates in place, with some points of the string swinging widely and others not moving at all. This is a **standing wave**. It explains a fact so familiar that it rarely seems to need explaining: a string, a pipe or a drum produces particular notes, not any frequency at all. This lesson shows how two travelling waves combine into a standing one, why only certain wavelengths fit, and how that sets the notes an instrument can play.

## What happens when a wave meets a fixed end?

When a wave travelling along a string reaches a point that is held still, the string there cannot move, so it cannot pass the disturbance on. Instead the wave is reflected and travels back along the string. It comes back upside down: a hump going in returns as a dip. The reason is that the fixed end has to stay at zero displacement, and the only way the incoming and reflected waves can always add to zero at that point is if the reflected wave is the inverted mirror image of the incoming one.

On a string fixed at both ends, waves reflect at both ends again and again. Once things have settled, the string carries two waves of the same frequency and amplitude travelling in opposite directions. By the principle of superposition from the previous lesson, the actual displacement of the string is the sum of the two. The question is what that sum looks like.

## Adding two opposite waves

To find out, we write the two waves as formulas and add them, using one piece of trigonometry to see what the sum really is.

::: steps Adding waves travelling in opposite directions
$ y_1 = A\sin(kx - \omega t), \qquad y_2 = A\sin(kx + \omega t)
These are two sine waves with the same amplitude A and frequency, the first moving toward +x and the second toward −x, as the sign in front of ωt shows. The symbols are shorthand for the factors in the wave formula from the lesson on wave basics: k = 2π/λ is called the **wave number**, and ω = 2πf is the angular frequency. Using them keeps the expressions short.
---
$ y = A\sin(kx - \omega t) + A\sin(kx + \omega t)
By superposition, the displacement of the string at each point and time is the sum of the two waves' displacements. As it stands, this is just two travelling waves written side by side. To see what the string actually does, we need to combine them into a single expression.
---
$ \sin(a \pm b) = \sin a\cos b \pm \cos a\sin b
This is the standard formula for the sine of a sum or difference of two angles. Applied to our two waves, with a = kx and b = ωt, it splits each sine into a part that depends on position times a part that depends on time.
---
$ \sin(a - b) + \sin(a + b) = 2\sin a\cos b
Add the two versions of the formula. The terms cos a sin b appear once with a minus sign and once with a plus sign, so they cancel, while the terms sin a cos b appear twice. The sum of two sines with opposite shifts collapses into a single product.
---
$ y = 2A\sin(kx)\cos(\omega t)
Apply this identity to the two waves. The result is a product of one factor that depends only on position, sin(kx), and another that depends only on time, cos(ωt). A travelling wave always has position and time mixed together in the form x − vt; this has no such mixture, so nothing travels.
:::

The meaning of the result is easiest to see by reading the two factors separately. The factor 2A sin(kx) is a fixed shape along the string, a sine curve that does not move. The factor cos(ωt) makes the whole shape swell, shrink to flat, invert, and swell back again, over and over. Every point on the string oscillates up and down in step with every other, but each with its own amplitude, set by where it sits on the fixed shape.

At some points sin(kx) is zero, and there the string never moves at all, however long you watch. These points are called **nodes**. Since sin(kx) is zero every time kx increases by π, and k = 2π/λ, the nodes are spaced half a wavelength apart. Halfway between neighbouring nodes are the points where sin(kx) is largest, and the string swings with the greatest amplitude, 2A. These are called **antinodes**. On a vibrating string you can see the pattern: the string blurs into one or more loops, each stretching from one node to the next.

## Which notes can a string play?

So far nothing has limited which wavelengths are possible. The fixed ends do that. A string of length L held at both ends cannot move at either end, so both ends must be nodes. Since nodes are half a wavelength apart, a whole number of half-wavelengths must fit exactly between the ends. A wave of any other wavelength would require one end to move, and its reflections would not settle into a steady pattern; they would interfere destructively and die away.

::: steps The allowed frequencies of a string
$ L = n\frac{\lambda_n}{2}, \qquad n = 1, 2, 3, \dots
The string, from one fixed end to the other, must span exactly n half-wavelengths, where n is a whole number. Each half-wavelength is one loop from node to node, so n counts the loops. This is the condition the fixed ends impose.
---
$ \lambda_n = \frac{2L}{n}
Multiply both sides by 2 and divide by n to find the wavelengths that fit. The longest is 2L, twice the length of the string; the others are a half, a third, a quarter of that, and so on. No wavelength in between is allowed.
---
$ f_n = \frac{v}{\lambda_n} = \frac{nv}{2L}
The wave speed v is fixed by the string, as the lesson on wave basics showed, so each allowed wavelength corresponds to one allowed frequency through v = fλ, or f = v/λ. Substituting the allowed wavelengths gives the allowed frequencies. They are whole-number multiples of the lowest one, v/2L.
---
$ f_n = \frac{n}{2L}\sqrt{\frac{F}{\mu}}
For a stretched string, the wave speed is v = √(F/μ), where F is the tension and μ the mass per unit length. Putting this in shows how each property of the string affects its notes. It is the complete recipe for the frequencies a string can produce.
:::

The lowest frequency, with n = 1, is called the **fundamental**. Half a wavelength fills the string, which vibrates in a single loop, and this sets the pitch we hear. The higher ones, at 2, 3, 4 and more times the fundamental frequency, are called **harmonics** or **overtones**: the second harmonic has two loops with a node in the middle, the third has three, and so on.

The formula also reads as a set of instructions for playing music. Shortening the string raises every frequency: that is what a guitarist does by pressing the string against a fret, so that only the part between the fret and the bridge vibrates. Tightening the string raises the tension and so the pitch: that is what the tuning pegs do. And making the string heavier, at the same tension, lowers the pitch: that is why the bass strings of a guitar are thicker, often wound with wire to add mass without making them too stiff.

::: note
Fixing the ends of a string allows only waves that fit a whole number of half-wavelengths between them. So a string has a set of natural frequencies, f, 2f, 3f and so on, and these determine the notes it can play.
:::

A worked example. A guitar's vibrating length is about 0.65 m, and its lowest string is tuned to a fundamental of 82.4 Hz. The fundamental has a wavelength of twice the string length, 1.3 m, so the wave speed on the string must be v = fλ = 82.4 × 1.3 ≈ 107 m/s. The tension and the string's weight are chosen to produce exactly that speed.

A real pluck does not produce only the fundamental. It sets off the fundamental and many harmonics at once, each with its own strength, and the ear hears them blended into a single note with a particular quality, or **timbre**. The mixture depends on how and where the string is plucked or struck, and on how the instrument's body responds. That is why a guitar and a piano playing the same note, with the same fundamental, still sound different: they mix the harmonics differently.

Air in a pipe does the same thing with sound. In a pipe open at both ends, the air at each open end is free to move, so the ends are antinodes of the air's motion, and the same counting argument gives the same rule, f = nv/2L, with v now the speed of sound. A pipe closed at one end is different. The closed end is a node, since the air there cannot move, and the open end is an antinode. The distance from a node to the next antinode is a quarter of a wavelength, so the pipe must hold an odd number of quarter-wavelengths. Its fundamental wavelength is 4L, twice as long as for an open pipe of the same length, so it sounds an octave lower, and it can produce only the odd harmonics: 1, 3, 5 and so on times the fundamental. This is part of why a clarinet, which behaves roughly like a closed pipe, sounds so different from a flute, which behaves like an open one.

::: history 1636–1701 · Paris
That string length and pitch are linked was known in antiquity: the Pythagoreans had found that halving a string raises its note by an octave, and that simple ratios of lengths give pleasing musical intervals. But nobody could measure the frequencies involved, which were far too fast to count. Marin Mersenne, a French friar who also served as a clearing house for the scientific correspondence of Europe, found a way round this. He strung very long, heavy strings that vibrated slowly enough to count by eye, and then used the relations he found to work out the frequencies of ordinary short strings. In *Harmonie universelle*, published in 1636, he set out the laws that bear his name: frequency goes inversely with length, as the square root of the tension, and inversely as the square root of the mass per unit length. These are exactly the dependences in the formula above. Around 1700 Joseph Sauveur, studying sound in Paris, recognised that a vibrating string sounds several pitches at once. He introduced the words "node" and "harmonic", and showed that the harmonics are whole-number multiples of the fundamental.
:::

## Resonance

Every object that can vibrate, from a guitar string to a bridge, has its own natural frequencies, the frequencies of its standing-wave patterns. If something pushes on the object periodically at one of those frequencies, each push arrives at just the right moment in the vibration to add to it, and the amplitude builds up from push to push. This is **resonance**. A child on a swing, pushed once each time the swing comes back, goes higher and higher. A singer holding a loud note at exactly the natural frequency of a wine glass can make it vibrate so violently that it shatters. A guitar's hollow body is designed to resonate with its strings over a wide range of frequencies: a bare string is thin and moves very little air, and without the body the note would be barely audible.

Resonance is often blamed for the collapse of the Tacoma Narrows Bridge in 1940, but that is not quite right. The wind that day was not gusting at the bridge's natural frequency. It blew fairly steadily, and the motion of the deck itself caused the airflow to push on it in a way that fed energy into a twisting oscillation, which grew until the bridge tore apart. Engineers call this aeroelastic flutter. Still, ordinary resonance is a real danger, and engineers design buildings and bridges so that wind, footsteps and earthquakes do not push them near their natural frequencies.

::: quiz
? A string's fundamental frequency is 100 Hz. What is the frequency of its third harmonic?
- 133 Hz
+ 300 Hz
- 900 Hz
= The harmonics are whole-number multiples of the fundamental, so the third harmonic is 3 × 100 Hz = 300 Hz.
? Pressing a guitar string against a fret halfway along its length makes the note
+ go up by an octave, doubling the frequency
- go down by an octave
- stay the same but get louder
= The fundamental frequency is v/2L. Halving the vibrating length doubles the frequency, which is one octave higher.
? How far apart are neighbouring nodes in a standing wave?
- One wavelength
+ Half a wavelength
- A quarter of a wavelength
= Nodes are where sin(kx) is zero, which happens every time kx increases by π, that is, every half wavelength.
:::
