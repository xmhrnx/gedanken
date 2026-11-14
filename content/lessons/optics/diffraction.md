You can hear someone talking around a corner but you cannot see them. Both sound and light are waves, and all waves spread when they pass an edge or squeeze through a gap. This spreading is **diffraction**. The difference is scale: sound waves are about a metre long, comparable to a doorway, while visible light waves are about half a micrometre, two million times smaller.

## Huygens' principle

In 1678 Christiaan Huygens proposed a way to follow a wave forward: treat every point on a wavefront as a source of small spherical wavelets, and the new wavefront is where those wavelets add up. On an open wavefront, the sideways parts of neighbouring wavelets cancel and the wave moves straight on. At the edge of an obstacle there are no neighbours on one side to cancel, so the wave spreads into the shadow.

How much it spreads depends on the ratio of wavelength to the size of the opening. Through a gap many wavelengths wide, a wave goes almost straight through, with a sharp shadow. Through a gap about one wavelength wide, it fans out in all directions.

## The single slit

Shine light of wavelength λ through a narrow slit of width a onto a screen a distance L away. The pattern is a bright central band, with dimmer bands on either side, separated by dark gaps. The dark gaps come from wavelets from different parts of the slit cancelling.

::: steps Where the first dark band falls
$ \Delta = \frac{a}{2}\sin\theta
Split the slit into a top half and a bottom half. Pair each point in the top half with the point a distance a/2 below it. For light heading to the screen at angle θ, the lower point's path is longer by (a/2) sin θ, since the screen is far enough away that the two rays are nearly parallel.
---
$ \frac{a}{2}\sin\theta = \frac{\lambda}{2}
If that path difference is half a wavelength, the two wavelets arrive exactly out of step, crest on trough, and cancel.
---
$ a\sin\theta = \lambda
Every point in the top half has such a partner, so all the light cancels in pairs. The first dark band is at this angle.
---
$ a\sin\theta = m\lambda \quad (m = \pm 1, \pm 2, \ldots)
Splitting the slit into four, six or more strips repeats the argument: whenever a sin θ is a whole number of wavelengths, adjacent strips cancel in pairs.
---
$ y_1 \approx \frac{\lambda L}{a}
For small angles, sin θ ≈ tan θ = y/L, where y is the distance from the centre of the screen. The central bright band runs from −y₁ to +y₁, so its full width is 2λL/a.
:::

The striking result is that the central band gets **wider** as the slit gets **narrower**. Squeezing light through a smaller opening spreads it more. With red light (λ = 650 nm) through a slit 0.1 mm wide onto a screen 2 m away, the central band is about 2.6 cm across, hundreds of times wider than the slit.

::: model interference
Watch how waves spreading from openings overlap. Each slit diffracts the light; where the spread waves meet, they add or cancel.
:::

## Round holes and sharp images

Light through a round opening of diameter D makes a bright central disc, the Airy disc, surrounded by faint rings. Its first dark ring is at an angle given by sin θ ≈ 1.22λ/D; the 1.22 comes from the round geometry. Two point sources closer together than about this angle blur into one. This is the **diffraction limit**, and no lens or mirror of diameter D can beat it.

::: equation The diffraction limit of a circular aperture
\theta_{\min} \approx 1.22\,\frac{\lambda}{D}
:::

For your eye, with a 3 mm pupil and green light, θ_min ≈ 2 × 10⁻⁴ radians, about 0.013 degrees: enough to separate car headlights 1.5 m apart at about 7 km. The Hubble Space Telescope, with a 2.4 m mirror, resolves about 0.05 arcseconds. Bigger telescopes see finer detail not only because they collect more light, but because they diffract less. The same limit is why microscopes cannot resolve detail much smaller than about half the wavelength of the light they use.

::: note
Waves spread as they pass through openings, by an angle of about λ/a. Narrow openings spread light more, and diffraction sets a hard limit on how sharp any optical image can be.
:::

::: history 1665–1819 · Bologna and Paris
Francesco Grimaldi noticed coloured fringes at the edges of shadows and coined the word "diffraction"; his book was published in 1665, two years after his death. Newton explained the fringes as forces bending light corpuscles near edges. In 1818 Augustin Fresnel entered a competition of the French Academy with a full wave theory of diffraction. Siméon Poisson, a judge who favoured particles, showed that Fresnel's theory predicted something absurd: a bright spot at the very centre of the shadow of a round disc. François Arago did the experiment and found the spot. Fresnel won, and the bright spot is now named after Arago or, ironically, Poisson.
:::

::: quiz
? Light passes through a slit, and the slit is made narrower. What happens to the central bright band on the screen?
- It gets narrower
+ It gets wider
- It stays the same width but dims
= The first dark band is at sin θ = λ/a, so a smaller a means a larger angle.
? Why do we hear around corners but not see around them?
+ Sound wavelengths are comparable to doorways; light wavelengths are far smaller
- Sound is a wave and light is not
- Light travels too fast to diffract
= Diffraction is strong when the opening is about one wavelength across. Doorways are about a metre; light waves are about half a micrometre.
? Which change lets a telescope separate two stars that are closer together?
- Using red light instead of blue
+ Using a larger mirror
- Using a longer eyepiece
= The diffraction limit is about 1.22λ/D: a larger D, or a shorter λ, gives a smaller resolvable angle.
:::
