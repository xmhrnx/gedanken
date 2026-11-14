At the start of the nineteenth century, most physicists followed Newton in thinking that light is a stream of tiny particles. The rival idea, that light is a wave, had supporters, but no experiment had forced a choice. Thomas Young supplied one. He showed that two beams of light, overlapping, can add up to darkness. Two streams of particles cannot do that: adding more particles to a spot can only make it brighter. Waves can, because a crest meeting a trough cancels. Young's experiment did more than settle a dispute. It measured the wavelength of light, a length less than a thousandth of a millimetre, with nothing more than a ruler. This lesson explains how.

## How can two waves cancel?

When two waves pass through the same place, the disturbance there is the sum of the two separate disturbances. For water waves that means the heights of the water add; for light, the electric fields add. This rule, called superposition, has a dramatic consequence. Where two waves arrive with crest on crest, the result is a larger wave: **constructive interference**. Where they arrive with crest on trough, the two disturbances are equal and opposite and add to nothing: **destructive interference**.

Which of these happens at a given spot depends on how far each wave has travelled to get there. Suppose two waves start out in step from two sources. If one has travelled exactly one wavelength farther than the other, it has gone through one whole extra cycle and is back in step: the crests coincide. The same is true for two extra wavelengths, or any whole number. But if the extra distance is half a wavelength, one wave arrives at a crest just as the other arrives at a trough, and they cancel. The extra distance one wave travels compared with the other is called the **path difference**, and it decides everything. A path difference of 0, λ, 2λ and so on gives brightness; a path difference of λ/2, 3λ/2, 5λ/2 and so on gives darkness.

There is a practical catch, and it is the reason this experiment was not done earlier. The argument above assumes the two sources start in step and stay that way. Sources that keep a fixed relation between their crests are called **coherent**. An ordinary lamp is nothing like that: its light comes from countless atoms emitting short bursts at random, so the crests of one lamp's light shift relative to another's millions of times a second. Two separate lamps pointed at the same wall do produce interference, but the bright and dark places move around far too fast to see, and the eye records a smooth average. Young's insight was to take light from a single small source and split it in two. Then whatever random changes the source goes through, both halves go through them together, and the relation between them stays fixed.

## What did Young's experiment look like?

In its classic form, light from one narrow source falls on a barrier with two narrow parallel slits, a distance d apart. Each slit, being narrow, diffracts the light that passes through it, spreading it out as the previous lesson described. So beyond the barrier, two spreading waves overlap, and because they came from the same source they are coherent. On a screen a distance L away appears a row of evenly spaced bright and dark bands, called **fringes**.

The pattern says something no particle theory could explain. Cover one slit and the screen is evenly lit by the diffracted light from the other. Uncover it, letting in more light, and some places on the screen get darker. Adding light has removed light there.

::: model interference
Change the separation between the two slits. Bright fringes form where the path difference is a whole number of wavelengths, and moving the slits farther apart packs the fringes closer together.
:::

## Where do the fringes fall?

To locate the fringes, measure positions on the screen by the distance y from its central point, the point directly opposite the middle of the two slits, or by the angle θ between the line to a point on the screen and the straight-ahead direction. The question is then: for each point, what is the path difference?

::: steps Deriving the fringe spacing
$ \Delta = r_2 - r_1 \approx d\sin\theta
Let r₁ and r₂ be the distances from the two slits to a point on the screen. The screen is far away compared with the slit separation d, so the two lines to it are nearly parallel, both at angle θ. Draw a short line from the nearer slit perpendicular to these rays; it cuts off a small right triangle whose hypotenuse is d and whose short side, d sin θ, is the extra distance travelled by the light from the farther slit.
---
$ d\sin\theta = m\lambda \quad (m = 0, \pm 1, \pm 2, \ldots)
A point on the screen is bright when the path difference is a whole number of wavelengths, so that the two waves arrive in step. The integer m counts which fringe it is. The value m = 0 is the central bright fringe, where the paths are equal; positive and negative m count fringes on either side.
---
$ d\sin\theta = \left(m + \tfrac{1}{2}\right)\lambda
A point is dark when the path difference is a whole number of wavelengths plus a half, so that crest meets trough. These dark fringes lie halfway between neighbouring bright ones.
---
$ \sin\theta \approx \tan\theta = \frac{y}{L}
The slit separation d is normally hundreds or thousands of wavelengths, so the angles at which the first several fringes appear are small. For small angles the sine and the tangent are almost equal, and the tangent is just the distance y on the screen divided by the distance L to the screen. This trades an angle, which is hard to measure, for a distance on the screen, which is easy.
---
$ y_m = \frac{m\lambda L}{d}
Replace sin θ in the bright-fringe condition with y/L, and solve for y. The m-th bright fringe sits this far from the centre of the screen.
---
$ \Delta y = y_{m+1} - y_m = \frac{\lambda L}{d}
Subtract the positions of two neighbouring bright fringes. The m cancels, so the spacing is the same everywhere in the pattern: the fringes are evenly spaced. The dark fringes, lying halfway between, are spaced the same way.
:::

The formula Δy = λL/d is the heart of the experiment. Everything in it except λ is a length you can measure with a ruler: the slit separation, the distance to the screen, and the gap between fringes. So it can be turned around to give the wavelength, λ = d Δy / L. And the arrangement acts as a magnifier for wavelengths, because L is huge compared with d. With a red laser of wavelength 633 nm, slits 0.25 mm apart and a screen 2 m away, Δy = (633 × 10⁻⁹ m × 2 m)/(2.5 × 10⁻⁴ m) ≈ 5 mm. A wavelength less than a thousandth of a millimetre becomes a fringe spacing of five millimetres, easy to measure.

The formula also explains the colours Young saw. A longer wavelength gives wider fringes, so in red light the bands are farther apart than in blue. With white light, every colour makes its own pattern. At the centre, where the path difference is zero for every wavelength, all the colours are bright together and the fringe is white. Farther out the patterns fall out of step with one another, and the fringes become edged with colour, blue on the inside and red on the outside, before blurring together.

The fringes are not equally bright all the way out. Each slit, having a finite width, makes its own single-slit diffraction pattern, and the two-slit fringes are seen only where that pattern sends light. So the fringes are brightest near the centre and fade away to the sides, under an envelope set by the width of each slit.

::: note
Two coherent waves add to brightness where their paths differ by a whole number of wavelengths and cancel to darkness halfway between. Because the fringe spacing is λL/d, a long distance to the screen turns a microscopic wavelength into a distance you can measure with a ruler.
:::

::: history 1801–1820 · London and Paris
Thomas Young trained as a physician, read a dozen languages, and later made some of the first progress in deciphering Egyptian hieroglyphs from the Rosetta Stone. He came to light through sound: he was struck that two musical notes of slightly different pitch make a throbbing beat, loud and soft in turn, and wondered whether light might combine in the same way. In his Bakerian Lecture to the Royal Society in November 1801 he set out the principle of interference and used it to explain the colours of soap films. In 1803 he demonstrated it directly, splitting a narrow sunbeam with the edge of a thin card held in it, so that light passing either side of the card overlapped and made fringes. The two-slit arrangement appeared in the published version of his lectures in 1807. From fringe measurements he estimated the wavelength of red light at about 1/36,000 of an inch (0.7 μm) and violet at about 1/60,000 of an inch (0.42 μm), close to modern values. The reception was hostile. Newton's particle theory carried enormous authority in Britain, and the lawyer and politician Henry Brougham attacked Young's papers anonymously in the Edinburgh Review as "destitute of every species of merit". Young was largely ignored until Augustin Fresnel in France, around 1815 to 1820, developed the wave theory independently and in mathematical detail, and with it explained diffraction in a way nobody could dismiss.
:::

## What happens if the light is very dim?

A century after Young, light turned out to have a second character. Einstein's explanation of the photoelectric effect in 1905, taken up in a later chapter, showed that light delivers its energy in indivisible lumps called photons. That raises a hard question about Young's experiment. What if the source is made so dim that only one photon is in the apparatus at a time?

The experiment has been done, with increasing care, since 1909. Each photon arrives at a single point on the screen, like a particle, never as a faint spread-out wave. Yet the points are not scattered at random. After thousands of photons have landed one by one, they have built up exactly Young's pattern of fringes, with no photons at all at the dark places. Each photon, it seems, is affected by both slits. The same thing has since been seen with electrons, with whole atoms, and with molecules made of dozens of atoms. Making sense of this is one of the central problems of quantum mechanics, and it is taken up in a later chapter.

::: quiz
? In a double-slit experiment, the slits are moved closer together. What happens to the fringes?
+ They spread farther apart
- They move closer together
- They stay exactly the same
= The fringe spacing is λL/d, so a smaller slit separation d gives a larger spacing. Closer slits produce a smaller path difference at a given angle, so you must go to a larger angle to reach the next fringe.
? At a dark fringe, what is the path difference between the waves from the two slits?
- Zero
- A whole number of wavelengths
+ An odd number of half-wavelengths
= At a dark fringe the two waves arrive half a cycle out of step, crest on trough, so they cancel. That happens when the path difference is (m + ½)λ.
? Why did Young light both slits from a single source rather than placing a lamp behind each one?
- Two lamps would make the pattern too bright to see
+ Two independent sources have no fixed relation between their crests, so their fringes shift too fast to see
- Only a single source has a definite wavelength
= Interference fringes stay still only if the two waves keep a steady phase relation. Splitting the light of one source guarantees that, because both halves share every random change of the source.
:::
