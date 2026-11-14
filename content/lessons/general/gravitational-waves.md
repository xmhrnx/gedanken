If matter curves spacetime, then matter that moves must change the curvature around it. But in relativity no influence can travel faster than light, so the change cannot appear everywhere at once. It has to spread outward from its source at a finite speed, as a wave: a ripple in the geometry of spacetime itself. Einstein predicted such waves in 1916, a few months after completing his theory, and then spent years unsure whether they were real or merely an artefact of his mathematics. Even those who believed in them thought they could never be detected. They were finally recorded on 14 September 2015. This lesson explains what a gravitational wave does, estimates why it is so weak, and describes the instrument that caught one.

## What does a gravitational wave do to the things it passes?

A gravitational wave does not push objects around the way a water wave pushes a boat. It changes the distances between them. As a wave passes, space is stretched along one direction while it is squeezed along the perpendicular direction; half a cycle later the pattern reverses, and so on at the wave's frequency. Imagine a ring of small objects floating freely in space, with a wave coming straight at the ring's face. The ring becomes an ellipse, first tall and narrow, then short and wide, then tall again.

The natural way to measure the wave's size is by the fractional change in length it causes, called the **strain** and written h. If two objects a distance L apart are moved closer or farther by ΔL, the strain is:

::: equation Strain
h = \frac{\Delta L}{L}
:::

Because the strain is a fraction, the change in length is proportional to the length: the farther apart two objects are, the larger the change between them. This is why detectors are built with long arms.

Gravitational waves travel at the speed of light, and like light their wavelength equals c divided by their frequency. The waves from merging black holes have frequencies of tens to hundreds of hertz, in the range of sound waves; a wave at 100 Hz has a wavelength of 3,000 km, about the width of the United States.

## What kind of motion makes the waves?

One might expect any vibrating mass to send out gravitational waves, just as a vibrating electric charge sends out light. It does not, and the reason is a conservation law. Light is mainly produced by a **dipole**, a separation of positive and negative charge that oscillates, like an electron bouncing back and forth in an antenna. For gravity, the corresponding quantity would be the position of the centre of mass. But momentum is conserved: whenever one part of an isolated system moves one way, another part must move the other way, and the centre of mass stays put. So there is no oscillating gravitational dipole and no dipole radiation.

The leading source is instead a changing **quadrupole**, a measure of how a mass distribution is stretched out, roughly mass times the square of its extent along some direction. A spinning dumbbell is the simplest example: seen along a fixed direction, its length grows and shrinks twice each turn, as first one end and then the other comes round. Two stars or black holes orbiting each other behave exactly like such a dumbbell, which makes them ideal sources. Their quadrupole repeats twice per orbit, so the waves they emit have twice the orbital frequency.

## Why are the waves so weak?

In 1918 Einstein worked out how strong the waves from a changing quadrupole should be. We can use his result to estimate the strain reaching Earth from a pair of black holes spiralling together far away. Factors of order one are kept only where they come easily; the aim is the size of the answer.

::: steps Estimating the strain from a merging binary
$ h \approx \frac{2G}{c^4 r}\,\ddot{Q}
Start from Einstein's quadrupole formula. It says the strain at distance r from the source is set by how rapidly the source's quadrupole Q is changing; the two dots mean the second derivative with respect to time, the rate of change of the rate of change. The factor G/c⁴ is about 8 × 10⁻⁴⁵ in SI units, and that tiny number is the fundamental reason gravitational waves are so weak: spacetime is extraordinarily stiff.
---
$ \mu = \frac{m_1 m_2}{m_1 + m_2}
Next describe the source. Two masses m₁ and m₂ in a circular orbit both circle their common centre of mass. A standard result of mechanics says their relative motion is the same as that of a single body of this **reduced mass** μ circling at a radius equal to their separation a. Replacing two bodies by one makes the quadrupole easy to write down.
---
$ Q = \mu a^2\cos^2\omega t = \tfrac{1}{2}\mu a^2\left(1 + \cos 2\omega t\right)
The single body's position along one fixed axis is a cos ωt, where ω is the orbital angular speed, so its quadrupole along that axis is μ times that position squared. The trigonometric identity cos²θ = ½(1 + cos 2θ) shows that this consists of a constant, which produces no waves, plus a part oscillating at 2ω, twice the orbital frequency, as the dumbbell picture promised.
---
$ \ddot{Q} = -2\mu a^2\omega^2\cos 2\omega t
Differentiate twice with respect to time. The constant disappears, and each derivative of cos 2ωt brings down a factor of 2ω, giving 4ω² overall, which with the ½ in front leaves 2μa²ω². This amplitude, the largest value of the oscillating part, is what sets the size of the wave.
---
$ h \approx \frac{4G\mu a^2\omega^2}{c^4 r} = \frac{4G\mu v^2}{c^4 r}
Put the amplitude into the quadrupole formula: 2G/(c⁴r) times 2μa²ω². Then notice that aω is the speed at which the two bodies move relative to each other, v, since a body circling at radius a with angular speed ω moves at aω. The strain grows with the reduced mass and with the square of the orbital speed.
---
$ h \approx 2\,\frac{r_s}{r}\,\frac{v^2}{c^2}, \quad r_s = \frac{2G\mu}{c^2}
Finally, regroup the constants to make the meaning clear. Writing 4Gμ/c² as twice the Schwarzschild radius r_s of the reduced mass, the strain becomes that radius divided by our distance, times the square of the orbital speed as a fraction of c. Since r_s is a few kilometres per solar mass and r is astronomically large, only masses of many suns moving at a good fraction of light speed produce anything we could hope to detect.
:::

Now put in the numbers for the first event detected. The two black holes had about 36 and 29 times the mass of the Sun, so the reduced mass was about 16 solar masses and its Schwarzschild radius about 47 km. Just before merging they were circling each other at roughly half the speed of light, so v²/c² ≈ 0.25. They were about 1.3 billion light-years away, which is r ≈ 1.2 × 10²⁵ m. The formula gives h ≈ 2 × 10⁻²¹. The measured peak strain was about 1.0 × 10⁻²¹, which is as close as an estimate this rough can expect to come.

Pause on what that number means. Two black holes, each tens of times heavier than the Sun, collided at half the speed of light and released more power in a fraction of a second than all the stars in the visible universe combined. And when the ripple reached Earth, it changed the length of a metre stick by one part in 10²¹.

::: note
Spacetime is extremely stiff. Even the most violent events in the universe, black holes colliding at half the speed of light, shake it here on Earth by only about one part in 10²¹.
:::

## How do you measure one part in 10²¹?

The answer is the Laser Interferometer Gravitational-Wave Observatory, LIGO. It has two detectors in the United States, at Hanford in Washington state and Livingston in Louisiana, about 3,000 km apart. Each is shaped like an L, with two arms 4 km long at right angles to each other.

A laser beam is split in two at the corner of the L, and each half travels down one arm, bounces off a mirror at the far end, and returns. The two halves are then recombined. The arms are adjusted so that when the beams recombine, their wave crests and troughs line up exactly out of step and cancel, leaving the output dark. A passing gravitational wave stretches one arm while squeezing the other, so the two beams travel slightly different distances, the cancellation is no longer perfect, and a little light leaks out. The amount of light measures the strain.

With a strain of 10⁻²¹ and 4 km arms, the arms change length by about 4 × 10⁻¹⁸ m, a few hundred times smaller than a proton. Reaching that sensitivity took decades of engineering. The light is bounced back and forth about 300 times in each arm before being recombined, which multiplies the effective arm length. The beams travel in one of the largest ultra-high vacuum systems in the world, so that stray gas molecules do not disturb them. The mirrors hang as pendulums from elaborate suspensions that isolate them from ground vibrations, passing trucks and ocean waves breaking hundreds of kilometres away. And the laser is powerful enough that the random arrival of individual photons does not drown the signal.

Two detectors far apart are essential. Any local disturbance, such as a small earthquake or a falling tree, affects only one site. A real gravitational wave, travelling at the speed of light, reaches both within 10 milliseconds, the time light takes to cover the distance between them, and produces the same signal in each.

::: history 2015 · Hanford and Livingston
On 14 September 2015 the upgraded Advanced LIGO detectors were running in a final engineering test, days before their first official observing run was due to begin. At 09:50 universal time both detectors recorded the same signal: a "chirp" that swept upward from about 35 to 250 Hz in two tenths of a second, rising in strength, then stopped. It reached Livingston 7 milliseconds before Hanford. The shape matched, almost perfectly, the waveform computed from general relativity for two black holes spiralling together and merging into one of about 62 solar masses. Since the two had weighed about 65 solar masses in total, about 3 solar masses had been converted into gravitational-wave energy, roughly 5 × 10⁴⁷ joules, in a fraction of a second. The signal was so clean that many in the collaboration first suspected it was a test: the team had a procedure for secretly injecting fake signals to check its own analysis. It took five months of checks, including confirming that no injection had been made, before the discovery was announced on 11 February 2016. Rainer Weiss, Kip Thorne and Barry Barish, who had led the design and construction of LIGO over four decades, received the Nobel Prize in Physics in 2017.
:::

The waves had in fact been detected indirectly long before. In 1974 Russell Hulse and Joseph Taylor discovered a **pulsar**, a rapidly spinning neutron star that sends out a radio pulse with every turn, in a tight orbit around another neutron star. The pulses act as an extremely precise clock, so the orbit could be tracked for years. If gravitational waves carry energy away, the two stars should slowly spiral closer and their orbital period should shrink. It did, at exactly the rate general relativity predicts, and Hulse and Taylor received the 1993 Nobel Prize for the discovery.

Since 2015 the detectors, now joined by Virgo in Italy and KAGRA in Japan, have recorded hundreds of mergers. One stands out. In August 2017, waves from two colliding neutron stars were followed 1.7 seconds later by a burst of gamma rays from the same spot in the sky. After a journey of about 130 million years, the two signals arrived within two seconds of each other, which shows that gravitational waves travel at the speed of light to within about one part in 10¹⁵.

The first LIGO signal was also the first direct evidence that black holes exist as objects that can orbit and merge. The next chapter explains what black holes are and how general relativity came to predict them.

::: model wave
Gravitational waves travel at c, like light. Set the frequency to LIGO's band, tens to hundreds of hertz, and notice how long the wavelength becomes: thousands of kilometres.
:::

::: quiz
? Why can't gravitational waves come from a single mass vibrating back and forth on its own?
- The mass would have to move faster than light
+ Momentum conservation means something else must move the opposite way, so the centre of mass never oscillates; the leading source is a changing quadrupole
- A single mass does not curve spacetime
= Dipole radiation would need the centre of mass of an isolated system to oscillate, which momentum conservation forbids. Orbiting pairs have a changing quadrupole, and that is what radiates.
? A gravitational wave of strain 10⁻²¹ passes a 4 km LIGO arm. By about how much does the arm's length change?
- 4 × 10⁻²¹ m
+ 4 × 10⁻¹⁸ m
- 4 mm
= The change in length is the strain times the length: ΔL = hL = 10⁻²¹ × 4,000 m = 4 × 10⁻¹⁸ m, a few hundred times smaller than a proton.
? If the same black-hole merger had happened ten times farther away, how would the strain at Earth compare?
- It would be 100 times smaller
+ It would be 10 times smaller
- It would be unchanged
= The strain falls as 1/r. The energy carried by the wave falls as 1/r², but detectors measure the strain, which is the wave's amplitude.
:::
