Drop two stones into a pond a little way apart, and two sets of ripples spread out and cross. Where they meet, something strange happens. The ripples do not bounce off each other or break up; each set passes straight through the other and carries on exactly as before. But while they overlap, the water does something neither set of ripples does alone: in some places it heaves up twice as high, and in others it lies almost flat, as if the ripples had cancelled. This behaviour is called **superposition**, and it is the most characteristic thing waves do. Particles cannot pass through each other, and two streams of particles cannot cancel to leave nothing. Waves can. That is why, for over a century, the question of whether light is a wave was settled by looking for exactly this effect.

## What happens when two waves meet?

For most waves of moderate size, the rule is remarkably simple. When two waves overlap, the displacement of the medium at each point is the sum of the displacements that each wave would produce there on its own:

::: equation The principle of superposition
y = y_1 + y_2
:::

Displacements can be positive or negative, up or down, so adding them can make the total either larger or smaller. And once the waves have passed through each other, each carries on as if the other had never been there, with its shape and speed unchanged.

Why should waves simply add? The rule holds when the medium responds in proportion to the push it receives, so that doubling the disturbance doubles the response. A medium that behaves this way is called **linear**. Then the effect of two disturbances together is just the sum of their separate effects, because neither one changes how the medium responds to the other. Sound at ordinary loudness, light, ripples on a pond and small vibrations on a string all behave this way to a very good approximation. Very violent waves do not: a breaking wave on a beach or the shock wave from an explosion changes the medium it passes through, and then superposition fails.

## When do waves reinforce, and when do they cancel?

Consider two waves with the same frequency and the same amplitude A arriving at the same place. What happens depends on how their oscillations line up there.

If the two waves arrive **in phase**, meaning crest arrives with crest and trough with trough, their displacements are always equal and add to double. The result is a wave of amplitude 2A. This is called **constructive interference**. If they arrive half a cycle out of phase, so that the crest of one arrives with the trough of the other, their displacements are always equal and opposite and add to zero. The medium does not move at all. This is called **destructive interference**. For any alignment in between, the result is an oscillation with an amplitude somewhere between 0 and 2A.

What decides whether the waves arrive in phase? Often it is simply how far each has travelled. Suppose two sources vibrate in step, emitting crests at the same moments. Each wave goes through one full cycle for every wavelength it travels. If the distances from the two sources to some point differ by exactly one wavelength, or two, or any whole number, then one wave has gone through a whole number more cycles than the other, which brings it back into step: crest meets crest. If the distances differ by half a wavelength, or one and a half, or any whole number plus a half, one wave is half a cycle behind the other: crest meets trough.

::: equation Conditions for two sources in step (n = 0, 1, 2, …)
\text{constructive: } \Delta &= n\lambda \\ \text{destructive: } \Delta &= \left(n + \tfrac{1}{2}\right)\lambda
:::

Here Δ is the **path difference**, the difference between the two distances, and λ is the wavelength. Since the path difference changes smoothly from place to place, the region where two waves overlap is divided into alternating zones of strong and weak motion: an **interference pattern**.

If waves can cancel, does energy vanish at the quiet places? It does not. Cancellation moves energy rather than destroying it. The energy a wave carries goes as the square of its amplitude. Where the waves reinforce, the amplitude is 2A, so the energy there is four times that of a single wave, twice what you would get by simply adding the energies of two separate waves. The extra energy at the bright places exactly makes up for the missing energy at the dark ones. Averaged over the whole pattern, the total is just the sum of what the two waves bring.

::: note
Overlapping waves add point by point. Where the path difference is a whole number of wavelengths they reinforce, and where it is a whole number plus a half they cancel. A pattern of alternating strong and weak bands is the unmistakable signature of waves.
:::

## Young's double slit

The most famous interference experiment uses light. Shine light of a single colour onto a barrier with two narrow, parallel slits a small distance d apart, and look at a screen a distance L beyond it. If light were a stream of particles travelling in straight lines, you would expect to see two bright lines on the screen, one behind each slit. Instead you see a row of evenly spaced bright and dark bands, called **fringes**.

The explanation is interference. Each slit acts as a new source of light waves spreading out on the far side. The two sources are in step, because the same light feeds both slits. At each point on the screen, the waves from the two slits have travelled slightly different distances, and where the path difference is a whole number of wavelengths, there is a bright fringe; where it is a whole number plus a half, a dark one.

::: model interference
Change the separation of the two sources and watch the fringes. Moving the sources closer together spreads the bright bands farther apart.
:::

The spacing of the fringes can be worked out from the geometry, and doing so turns a pattern on a screen into a measurement of the wavelength of light.

::: steps The spacing of the fringes
$ \Delta = d\sin\theta
Consider a point on the screen in a direction making an angle θ with the line straight ahead from the slits. Because the screen is far away compared with the slit separation, the two paths to that point are nearly parallel. Then, by drawing a small right-angled triangle at the slits, one path is longer than the other by d sin θ.
---
$ d\sin\theta = n\lambda
A bright fringe appears wherever the path difference is a whole number of wavelengths, so set d sin θ equal to nλ. Each whole number n gives the direction of one bright fringe: n = 0 straight ahead, n = 1 the first fringe on either side, and so on.
---
$ \sin\theta \approx \tan\theta = \frac{y}{L}
We want positions on the screen rather than angles. If a fringe is a height y above the centre of the screen, then tan θ = y/L from the right-angled triangle formed by the slits, the centre of the screen and the fringe. For the small angles in a typical experiment, sin θ and tan θ are almost equal, so we can use y/L for sin θ too.
---
$ y_n = \frac{n\lambda L}{d}
Put y/L in place of sin θ in the condition d sin θ = nλ, and multiply both sides by L/d. This gives the height of the n-th bright fringe above the centre of the screen.
---
$ \Delta y = \frac{\lambda L}{d}
The gap between neighbouring bright fringes is the difference between y for n + 1 and y for n, which is λL/d. It does not depend on n, so near the centre of the pattern the fringes are evenly spaced. The spacing grows with the wavelength and the distance to the screen, and shrinks as the slits are moved apart.
:::

The numbers show why this was such a powerful experiment. Take red light with a wavelength of 650 nm, slits 0.25 mm apart, and a screen 1 m away. The fringe spacing is 650 × 10⁻⁹ × 1 / (0.25 × 10⁻³) = 2.6 × 10⁻³ m, or 2.6 mm: easy to see and measure with a ruler. Run the formula backwards and the measured spacing gives the wavelength. The slit geometry acts as a magnifier, turning a length of less than a thousandth of a millimetre, far too small to measure directly, into a spacing of millimetres.

::: history 1801–1820 · London and Paris
For a century after Newton, most scientists held that light was a stream of tiny particles, largely because Newton had thought so. Thomas Young was an unlikely challenger: a physician by training, a gifted linguist who would later help decipher the Egyptian hieroglyphs on the Rosetta Stone, and a man interested in nearly everything. His studies of sound, where interference was already familiar, led him to suspect that light was a wave too, and that interference was the way to prove it. In 1803 he showed the Royal Society fringes produced by splitting a narrow beam of sunlight with the edge of a thin card, so that light passing on either side of the card overlapped; his lectures of 1807 describe the version with two slits. From the fringe spacing he worked out the wavelengths of different colours, and got them right, all less than a thousandth of a millimetre. The response was hostile. Henry Brougham, writing anonymously in the influential *Edinburgh Review*, mocked Young's papers as lacking any merit, and the wave theory made little headway in Britain. It was accepted only after Augustin Fresnel in France, working independently, gave it a full mathematical treatment in the late 1810s that predicted diffraction patterns in detail, and those predictions were confirmed.
:::

## Where else does interference show up?

Once you know what to look for, interference is everywhere. Noise-cancelling headphones contain a microphone that picks up incoming sound and a circuit that plays back its mirror image, a wave with the same shape but upside down, so that the two cancel at your ear. The shimmering colours on a soap bubble or a film of oil on a wet road come from light reflecting off both the front and back surfaces of the thin film. The light from the back surface has travelled a little farther, and depending on the film's thickness, that extra path makes some wavelengths reinforce and others cancel. Which colours survive changes as the thickness varies across the film, producing bands of colour.

Sound gives another example. Two tuning forks of slightly different frequencies, sounded together, produce **beats**: the combined sound swells and fades regularly. The two waves drift in and out of step, reinforcing when they are in phase and cancelling when they are out of phase, and the number of beats per second equals the difference between the two frequencies. Musicians tune by listening for beats against a reference note and adjusting until the beats slow down and disappear.

Interference does need one condition: the two waves must keep a fixed phase relationship, so that the pattern stays in place long enough to be seen. Two separate light bulbs do not. Each emits light in short, random bursts from countless atoms, and the phase of one relative to the other changes unpredictably far faster than any eye or instrument can follow. Their interference pattern shifts so rapidly that it averages out to an even glow. That is why Young fed both of his slits from a single source: the light reaching the two slits came from the same place at the same moment, so it stayed in step.

::: quiz
? Two speakers play the same tone in step. At a point 3.0 m from one and 3.5 m from the other, the sound is very quiet. What could the wavelength be?
- 0.5 m
+ 1.0 m
- 3.5 m
= Very quiet means destructive interference, so the 0.5 m path difference must be a whole number of wavelengths plus a half. With a wavelength of 1.0 m, the path difference is exactly half a wavelength.
? In a double-slit experiment you move the slits closer together. The fringes
+ spread farther apart
- move closer together
- stay the same
= The fringe spacing is λL/d, so making the slit separation d smaller makes the spacing larger.
? When two waves cancel at a point, what happens to their energy?
- It is destroyed
+ It is redistributed to places where the waves reinforce
- It turns into heat at that point
= Interference rearranges energy rather than destroying it. The dark places are balanced by bright places where the energy is more than the sum of the two separate waves.
:::
