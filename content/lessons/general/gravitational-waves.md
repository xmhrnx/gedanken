If matter curves spacetime, then moving matter should change the curvature, and the change cannot spread faster than light. It travels outward as a wave: a ripple in the geometry of spacetime itself. Einstein predicted these waves in 1916 and doubted for years whether they were real or could ever be seen. They were finally detected on 14 September 2015.

## What a gravitational wave does

A passing gravitational wave stretches space in one direction while squeezing it in the perpendicular direction, then reverses, over and over at the wave's frequency. A ring of floating particles facing the wave becomes an ellipse, first tall and then wide. The size of the effect is measured by the **strain** h, the fractional change in length:

::: equation Strain
h = \frac{\Delta L}{L}
:::

Gravitational waves travel at the speed of light, and like light their wavelength is c divided by the frequency. A wave at 100 Hz has a wavelength of 3000 km.

Unlike light, they cannot come from a single object oscillating back and forth, a "dipole". Conservation of momentum forbids that: when one mass moves one way, something else moves the other, and the shifts cancel. The leading source is a changing **quadrupole**, a measure of how lopsided a mass distribution is. Two stars orbiting each other are ideal: their quadrupole repeats twice per orbit.

## How small?

Einstein worked out in 1918 how strong the waves from a moving mass should be. We can use his formula to estimate what Earth receives from a pair of black holes merging far away.

::: steps Estimating the strain from a merging binary
$ h \approx \frac{2G}{c^4 r}\,\ddot{Q}
Einstein's quadrupole formula: the strain at distance r is set by the second time derivative of the source's quadrupole Q (mass times distance squared). The factor G/c⁴ is about 8 × 10⁻⁴⁵ in SI units, which is why the waves are so weak.
---
$ Q = \mu a^2\cos^2\omega t = \tfrac{1}{2}\mu a^2\left(1 + \cos 2\omega t\right)
For two masses m₁ and m₂ in a circular orbit of separation a and angular speed ω, the motion is equivalent to one body of reduced mass μ = m₁m₂/(m₁ + m₂) circling at radius a. Along one axis its quadrupole varies like this, at twice the orbital frequency.
---
$ \ddot{Q} = -2\mu a^2\omega^2\cos 2\omega t
Differentiate twice: each derivative of cos 2ωt brings down a factor 2ω. The amplitude is 2μa²ω².
---
$ h \approx \frac{4G\mu v^2}{c^4 r}
Substitute the amplitude and write aω = v, the relative orbital speed.
---
$ h \approx 2\,\frac{r_s}{r}\,\frac{v^2}{c^2}, \quad r_s = \frac{2G\mu}{c^2}
Rearranged: the strain is roughly the Schwarzschild radius of the reduced mass divided by our distance, times (v/c)². Only nearly light-speed motion of black-hole-sized masses gives anything detectable.
:::

For the first event detected, the black holes had about 36 and 29 solar masses, so μ ≈ 16 solar masses and r_s ≈ 47 km. Just before merging they circled at about half the speed of light, so v²/c² ≈ 0.25, and they were about 1.3 billion light years away, r ≈ 1.2 × 10²⁵ m. The estimate gives h ≈ 2 × 10⁻²¹. The measured peak was about 1.0 × 10⁻²¹, as close as a calculation this rough can hope for.

::: note
Spacetime is extremely stiff. Even the most violent events in the universe, black holes colliding at half the speed of light, shake it on Earth by about one part in 10²¹.
:::

## LIGO

The Laser Interferometer Gravitational-Wave Observatory has two L-shaped detectors, at Hanford, Washington and Livingston, Louisiana, about 3000 km apart. Each has two arms 4 km long at right angles. A laser beam is split, sent down both arms, reflected back by mirrors hanging as pendulums, and recombined. The beams are set to cancel at the output. A passing wave lengthens one arm and shortens the other, the cancellation becomes imperfect, and light leaks out.

With h ≈ 10⁻²¹ and L = 4 km, the arms change length by about 4 × 10⁻¹⁸ m, a few hundred times smaller than a proton. Reaching that needed the light to bounce back and forth about 300 times in each arm, ultra-high vacuum, mirrors isolated from ground vibration, and enough laser power that the random arrival of photons does not swamp the signal. Two sites far apart rule out local disturbances: a real wave reaches both within 10 ms, the light travel time between them.

::: history 2015 · Hanford and Livingston
On 14 September 2015, days before the upgraded Advanced LIGO's first official run, both detectors recorded the same "chirp": a signal sweeping from about 35 to 250 Hz in two tenths of a second, arriving at Livingston 7 ms before Hanford. It matched the computed waveform of two black holes spiralling together into one of about 62 solar masses. The missing three solar masses had been radiated as gravitational waves, about 5 × 10⁴⁷ J in a fraction of a second. The collaboration spent five months checking it, including ruling out a deliberately injected fake, before announcing on 11 February 2016. Rainer Weiss, Kip Thorne and Barry Barish received the 2017 Nobel Prize.
:::

The waves had been detected indirectly long before. In 1974 Russell Hulse and Joseph Taylor discovered a pulsar orbiting another neutron star. Over the following years its orbit shrank exactly as fast as general relativity says it should, from the energy carried off by gravitational waves. That earned the 1993 Nobel Prize.

Since 2015, detectors have recorded hundreds of mergers. In August 2017, waves from two colliding neutron stars arrived 1.7 seconds before a burst of gamma rays from the same event, after 130 million years of travel, showing that gravitational waves move at the speed of light to about one part in 10¹⁵.

The first LIGO signal was also the first direct evidence that black holes exist and merge. The next chapter explains what black holes are and how general relativity predicted them.

::: model wave
Gravitational waves travel at c, like light. Set the frequency to LIGO's band, tens to hundreds of hertz, and note how long the wavelength becomes.
:::

::: quiz
? Why can't gravitational waves come from a single mass vibrating back and forth?
- The mass would have to move faster than light
+ Momentum conservation means something else moves the opposite way, cancelling the effect; the leading source is a changing quadrupole
- Single masses do not curve spacetime
= A dipole wave would need the centre of mass to oscillate, which momentum conservation forbids. Orbiting pairs have a changing quadrupole.
? A gravitational wave of strain 10⁻²¹ passes a 4 km LIGO arm. The arm changes length by about
- 4 × 10⁻²¹ m
+ 4 × 10⁻¹⁸ m
- 4 mm
= ΔL = hL = 10⁻²¹ × 4000 m.
? The same black-hole merger happened ten times farther away. The strain at Earth would be
- 100 times smaller
+ 10 times smaller
- unchanged
= h falls as 1/r. (The energy flux falls as 1/r², but detectors measure strain, an amplitude.)
:::
