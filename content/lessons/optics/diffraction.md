You can hear someone talking in the next room through an open door, even when you cannot see them. Sound reaches you around the edge of the doorway; light, apparently, does not. For a long time this was taken as evidence that light is not a wave at all, since waves on water and in air plainly spread around obstacles. The answer turned out to be that light does spread around edges, a behaviour called **diffraction**, but on a scale so small that everyday life hides it. This lesson explains why waves spread when they pass an edge, works out how much, and shows that the same spreading sets a hard limit on how sharp any telescope, microscope or eye can see.

## Why do waves spread at all?

Watch ripples on a pond pass through a gap between two stones. On the far side they do not continue as a neat strip the width of the gap. They curve outward into the regions that the stones should have shaded. The narrower the gap, the more nearly the ripples spread out as semicircles, as if the gap itself were a new source.

That observation is the basis of a way of predicting how any wave moves, proposed by Christiaan Huygens in 1678. Treat every point on a wavefront, the line joining neighbouring crests, as a small source sending out circular wavelets, and the wavefront a moment later is the line along which those wavelets join up. On a long, unobstructed wavefront, every point has neighbours on both sides, and the wavelets combine into a front that moves straight ahead. At the edge of an obstacle, the points next to the edge have no neighbours on the shaded side, and the wavelets they send there are not matched by any others. So the wave leaks into the shadow.

Huygens himself used his construction only to explain reflection and refraction; he did not think of his wavelets as interfering, adding where crests meet and cancelling where crest meets trough. That step came from Augustin Fresnel in the 1810s, who added the wavelets with their crests and troughs taken into account. The combination, now called the Huygens–Fresnel principle, is what predicts diffraction patterns correctly, and it is the version used below.

How much a wave spreads depends on one ratio: its wavelength compared with the size of the opening. Through a gap many wavelengths wide, the edges affect only a thin region near the shadow line, and the wave goes almost straight through with a sharp shadow. Through a gap about one wavelength wide, the whole gap acts like a single source and the wave fans out in all directions. Sound waves of ordinary speech are around a metre long, comparable to a doorway, so they spread freely through it. Visible light has a wavelength of about half a micrometre, two million times shorter. A doorway is millions of wavelengths wide for light, and the spreading is far too small to notice. That is the whole difference between hearing and seeing around corners.

## What does light look like through a narrow slit?

To see light diffract you need an opening not much wider than its wavelength. Shine light of a single wavelength λ, from a laser for instance, through a narrow slit of width a, and let it fall on a screen a distance L away. The patch on the screen is not a sharp image of the slit. It is a bright central band, much wider than the slit, with a row of fainter bands on each side, separated by dark gaps.

The dark gaps are the interesting part, because they show light adding to darkness. Every point across the slit sends wavelets toward every point on the screen. At most places on the screen they arrive with a mixture of timings, and some light is left over. But at certain angles the wavelets from the slit cancel completely. The way to find those angles is to pair the wavelets off.

::: steps Where the dark bands fall
$ \Delta = \frac{a}{2}\sin\theta
Divide the slit into a top half and a bottom half, and pair each point in the top half with the point exactly a/2 below it. Consider light leaving at an angle θ to the straight-ahead direction. The screen is far away compared with the slit, so the two rays from a pair are very nearly parallel, and the lower one has farther to go by the short side of a small right triangle with hypotenuse a/2: a path difference of (a/2) sin θ.
---
$ \frac{a}{2}\sin\theta = \frac{\lambda}{2}
Now ask when a pair cancels. If the lower wavelet has half a wavelength farther to travel, it arrives exactly out of step with the upper one, its crest meeting the other's trough, and the two cancel. This equation picks out the angle at which that happens.
---
$ a\sin\theta = \lambda
At that angle, every point in the top half has a partner in the bottom half that cancels it, and the pairing uses up the whole slit. So no light at all reaches the screen in this direction. Multiplying both sides by 2 gives the first dark band in its usual form.
---
$ a\sin\theta = m\lambda \quad (m = \pm 1, \pm 2, \ldots)
The same argument works with the slit cut into four strips, or six, or any even number. If each strip's wavelets are half a wavelength behind the strip above, neighbouring strips cancel in pairs, and that happens whenever the path difference across the whole slit is a whole number of wavelengths. These are all the dark bands; the bright bands lie roughly halfway between them, where the cancellation is incomplete.
---
$ y_1 \approx \frac{\lambda L}{a}
For small angles, which is the usual case, sin θ is very nearly equal to tan θ, and tan θ is the distance y from the centre of the screen divided by L. Putting that into a sin θ = λ gives the position of the first dark band. The central bright band runs from −y₁ to +y₁, so its full width is 2λL/a.
:::

The striking feature of this result is that a sits in the denominator. Make the slit **narrower** and the central band gets **wider**. Squeezing light through a smaller opening spreads it more, not less, which is the opposite of what the idea of light as a stream of straight-moving particles would suggest. Put in numbers: red light with λ = 650 nm passing through a slit 0.1 mm wide onto a screen 2 m away gives y₁ = (650 × 10⁻⁹ m × 2 m)/(10⁻⁴ m) = 1.3 cm. The central band is about 2.6 cm across, more than 200 times wider than the slit that made it.

::: model interference
Change the separation and watch the waves spreading from each opening overlap. Every opening diffracts the light that passes through it, and where the spread-out waves meet, they add or cancel.
:::

## Why can no lens see infinitely fine detail?

Most optical instruments have round openings rather than slits: the pupil of an eye, the lens of a camera, the mirror of a telescope. Light passing through a circular opening of diameter D diffracts too. The pattern is a bright central spot, called the Airy disc after the astronomer George Airy who first calculated it, surrounded by faint rings. The calculation is harder than for a slit because the opening is round, and the result has an extra factor: the first dark ring falls at an angle where sin θ ≈ 1.22λ/D. The 1.22 comes purely from the circular shape.

This has a consequence that no amount of careful grinding can avoid. A perfect lens does not focus a point of light, such as a distant star, to a point. It focuses it to an Airy disc. Two stars close together in the sky make two overlapping discs, and if they are too close the discs merge into a single blur. The usual rule, due to Lord Rayleigh, is that two points can just be told apart when the centre of one disc falls on the first dark ring of the other. That is the **diffraction limit**:

::: equation The diffraction limit of a circular opening
\theta_{\min} \approx 1.22\,\frac{\lambda}{D}
:::

Here θ_min is the smallest angle between two points that can still be seen as separate, in radians. The formula says two things. Shorter wavelengths give finer detail, and so do larger openings.

Take the human eye. In bright light the pupil is about 3 mm across, and the eye is most sensitive to green light of about 550 nm. Then θ_min ≈ 1.22 × (550 × 10⁻⁹ m)/(3 × 10⁻³ m) ≈ 2 × 10⁻⁴ radians, about 0.013 degrees. At that angle, the two headlights of a car, 1.5 m apart, merge into one at a distance of about 7 km. Real eyes do a little worse than this, because the eye's own lens is not perfect, but not much worse; evolution has brought the eye close to the physical limit.

The Hubble Space Telescope has a mirror 2.4 m across, 800 times wider than a pupil, and resolves about 0.05 arcseconds (an arcsecond is 1/3600 of a degree). That is the real reason astronomers build ever larger telescopes: not only to collect more light from faint objects, but because a wider mirror diffracts less and shows finer detail. The same limit applies to microscopes, which cannot resolve features much smaller than about half the wavelength of the light they use, around 200 nm for visible light. To see smaller things, such as viruses or the arrangement of atoms, you need waves with much shorter wavelengths, which is why electron microscopes exist.

::: note
Waves spread when they pass through an opening, by an angle of roughly the wavelength divided by the width of the opening. Narrower openings spread light more, and because every lens and mirror is itself an opening, diffraction sets a hard limit, about 1.22λ/D, on how fine a detail any optical instrument can show.
:::

::: history 1665–1819 · Bologna and Paris
The Jesuit astronomer Francesco Grimaldi, working in Bologna, noticed that the shadow of a thin rod in a narrow beam of sunlight was wider than straight-line rays allowed and was edged with coloured fringes. He coined the word "diffraction" for this breaking of light, and his book describing it appeared in 1665, two years after his death. Newton knew of the fringes and studied them himself, but explained them with his particle theory, supposing that forces near the edge of an object deflect passing light corpuscles. With Newton's authority behind it, the particle view dominated for a century. In 1818 the French Academy of Sciences set a prize competition on diffraction, and most of its judges expected particles to win. Augustin Fresnel, a civil engineer, entered a full wave theory that predicted the fringes with high precision. One judge, Siméon Poisson, worked through Fresnel's mathematics and found what he thought was a fatal flaw: the theory predicted that in the very centre of the shadow of a small round disc there should be a bright spot, as bright as if the disc were not there. That seemed absurd. Another judge, François Arago, tried the experiment and found the spot. Fresnel won the prize in 1819, and the bright spot is now named after Arago or, with some irony, after Poisson.
:::

::: quiz
? Light passes through a slit, and the slit is then made narrower. What happens to the central bright band on the screen?
- It gets narrower
+ It gets wider
- It keeps the same width but becomes dimmer
= The first dark band falls where sin θ = λ/a, so a smaller slit width a means a larger angle. Squeezing light through a narrower opening spreads it more.
? Why can we hear around corners but not see around them?
+ Sound wavelengths are comparable to a doorway, while light wavelengths are millions of times smaller
- Sound is a wave and light is not
- Light travels too fast to diffract
= Diffraction is strong only when the opening is not much wider than the wavelength. Speech has wavelengths of about a metre, close to a doorway's width, but light's half-micrometre wavelength makes the spreading at a doorway negligible.
? Which change lets a telescope separate two stars that are closer together in the sky?
- Observing in red light instead of blue
+ Using a larger mirror
- Using an eyepiece with a longer focal length
= The smallest resolvable angle is about 1.22λ/D, so a larger diameter D, or a shorter wavelength, separates closer stars. Red light has a longer wavelength and makes the limit worse.
:::
