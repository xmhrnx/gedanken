At the start of the 1800s most physicists followed Newton and thought light was a stream of particles. Thomas Young showed that two beams of light can add up to darkness. Particles cannot do that; waves can. The experiment also measured the wavelength of light, a length far too small to see.

## Interference

When two waves overlap, their displacements add. Where crest meets crest, they reinforce: **constructive interference**. Where crest meets trough, they cancel: **destructive interference**. Whether two waves arrive in step depends on the difference in the distances they have travelled, the **path difference**.

- Path difference a whole number of wavelengths (0, λ, 2λ, …): in step, bright.
- Path difference a whole number plus a half (λ/2, 3λ/2, …): out of step, dark.

For a steady pattern the two sources must keep a fixed phase relationship (they must be **coherent**). Two separate lamps flicker randomly relative to each other and their pattern washes out. Young's trick was to split light from one source into two.

## The experiment

Light from a single narrow source falls on a barrier with two narrow slits a distance d apart. Each slit diffracts the light, so the two spread-out waves overlap on a screen a distance L away. The screen shows a row of evenly spaced bright and dark bands, called **fringes**.

::: model interference
Change the slit separation. Bright fringes are where the path difference is a whole number of wavelengths; wider separation packs them closer together.
:::

## Where the fringes fall

Measure position on the screen by the distance y from the central point, directly opposite the midpoint of the slits, or by the angle θ from that line.

::: steps Deriving the fringe spacing
$ \Delta = r_2 - r_1 \approx d\sin\theta
The screen is far away compared with d, so the rays from the two slits to a point on the screen are nearly parallel. The extra path for the farther slit is the short side of a small right triangle with hypotenuse d and angle θ.
---
$ d\sin\theta = m\lambda \quad (m = 0, \pm 1, \pm 2, \ldots)
Bright fringes are where the path difference is a whole number of wavelengths. m = 0 is the central bright fringe, where the paths are equal.
---
$ d\sin\theta = \left(m + \tfrac{1}{2}\right)\lambda
Dark fringes are halfway between, where the waves arrive half a wavelength out of step.
---
$ \sin\theta \approx \tan\theta = \frac{y}{L}
For small angles, which is the usual case since d is much larger than λ, the sine and tangent are nearly equal, and the tangent is the screen distance over L.
---
$ y_m = \frac{m\lambda L}{d}
Combine the last two lines: the m-th bright fringe sits at this distance from the centre.
---
$ \Delta y = y_{m+1} - y_m = \frac{\lambda L}{d}
Neighbouring bright fringes are evenly spaced, by this amount. So are the dark ones.
:::

The formula can be turned round to measure λ = d Δy / L. With a red laser (λ = 633 nm), slits 0.25 mm apart and a screen 2 m away, the fringes are about 5 mm apart: easy to see and measure with a ruler, though the wavelength is less than a thousandth of a millimetre. Light of a longer wavelength gives wider fringes. With white light, each colour makes its own pattern, so the central fringe is white and the others are edged with colour.

The fringes do not go on forever with equal brightness. Each slit's own diffraction pattern (previous lesson) sets an overall envelope: the fringes are brightest in the middle and fade out.

::: note
Two coherent waves add to bright fringes where their paths differ by whole wavelengths and to dark ones in between. The fringe spacing λL/d makes a microscopic wavelength into a measurable distance.
:::

::: history 1801–1807 · London
Thomas Young, a physician who also helped decipher the Rosetta Stone, announced the principle of interference in his Bakerian Lecture to the Royal Society in November 1801. He demonstrated it by splitting a sunbeam with a thin card and later showed the two-slit arrangement in his 1807 lectures. From the fringes he estimated wavelengths of about 1/36,000 inch (0.7 μm) for red light and 1/60,000 inch (0.42 μm) for violet, close to modern values. British physicists, loyal to Newton, mostly ignored him, and Henry Brougham attacked his work anonymously in the Edinburgh Review. The wave theory was accepted only after Fresnel's work around 1820.
:::

## A coda: one photon at a time

A century later, light turned out to arrive in lumps, photons. Dim the source until only one photon is in the apparatus at a time and each one lands at a single spot on the screen. Yet after many thousands, the spots build up Young's fringes. Electrons, atoms and even large molecules do the same. Making sense of that is one of the central puzzles of quantum mechanics, taken up in a later chapter.

::: quiz
? In a double-slit experiment, the slits are moved closer together. What happens to the fringes?
+ They spread further apart
- They move closer together
- They stay the same
= Δy = λL/d, so a smaller d gives a larger spacing.
? At a dark fringe, what is the path difference from the two slits?
- Zero
- A whole number of wavelengths
+ An odd number of half-wavelengths
= The waves arrive half a cycle out of step and cancel: (m + ½)λ.
? Why did Young need both slits to be lit by one source rather than by two lamps?
- Two lamps would be too bright
+ Two independent sources have no fixed phase relationship, so their fringes wash out
- One source has a single wavelength
= Interference fringes stay put only if the two waves keep a steady phase difference, which splitting one source guarantees.
:::
