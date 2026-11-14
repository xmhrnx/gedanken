A black hole emits essentially nothing, so it cannot be seen directly. What can be seen is its effect on everything nearby: stars orbiting it, gas heated as it falls in, and ripples in spacetime when two of them merge. Michell's 1783 suggestion, to watch a visible star circling an unseen companion, turned out to be the way in.

## Cygnus X-1: weighing an invisible companion

In 1971 the Uhuru X-ray satellite and radio observations pinned down a bright, rapidly flickering X-ray source in Cygnus. It coincided with a blue supergiant star, HDE 226868, whose spectral lines shift back and forth every 5.6 days. The star is orbiting something no telescope can see. X-rays that flicker in milliseconds must come from a region no bigger than light crosses in that time, a few hundred kilometres, so the companion is small and compact. The question was its mass.

You cannot measure the full orbit, only the star's speed toward and away from us, from the Doppler shift of its spectral lines. Combined with Kepler's third law, that is enough to put a floor under the unseen mass.

Call the visible star's mass M₁, the unseen companion's M₂, the orbital period P, and the inclination i (the tilt of the orbit, with i = 90° seen edge-on). Assume a circular orbit.

::: steps The mass function
$ G(M_1 + M_2) P^2 = 4\pi^2 a^3
Kepler's third law in Newton's form, where a is the separation of the two bodies. Both orbit their common centre of mass.
---
$ a_1 = \frac{a M_2}{M_1 + M_2}
The centre of mass sits where M₁a₁ = M₂a₂ with a₁ + a₂ = a, so the visible star's orbit has this radius.
---
$ K = \frac{2\pi a_1}{P} \sin i
The star's orbital speed is circumference over period. We see only the part along our line of sight, which is the speed times sin i. K is the measured amplitude of the Doppler wobble.
---
$ a_1 \sin i = \frac{KP}{2\pi}
Rearrange the previous line.
---
$ f = \frac{PK^3}{2\pi G} = \frac{M_2^3 \sin^3 i}{(M_1 + M_2)^2}
Cube the last line, replace a₁ using the second line and a³ using Kepler's law, and collect terms. The left side holds only measured quantities; the right side holds the unknown masses. f is called the mass function.
---
$ M_2 \geq f
Since sin i ≤ 1 and M₁ + M₂ ≥ M₂, the right side can never exceed M₂. The mass function is a strict lower limit on the unseen mass, and the real value is larger the heavier the visible star is.
:::

For Cygnus X-1, P = 5.6 days and K ≈ 75 km/s give f ≈ 0.25 solar masses. That alone proves little, but the visible star is a supergiant of many solar masses, and putting a realistic M₁ into the equation forces M₂ well above any neutron star limit. Early estimates put the companion above about 6 solar masses; a 2021 analysis gives about 21. Nothing known except a black hole is that heavy, that small and that dark.

::: history 1974 · a bet
Kip Thorne and Stephen Hawking bet on whether Cygnus X-1 contained a black hole, Hawking taking "no" as insurance in case his life's work proved wrong. He conceded in 1990, as the evidence piled up.
:::

## The centre of our galaxy

Since the early 1990s two teams, led by Reinhard Genzel and Andrea Ghez, have tracked individual stars in the innermost light-weeks of the Milky Way, using infrared light to see through dust. The stars swing around a point coinciding with the radio source Sagittarius A* (Sgr A*). The star S2 completes an orbit in about 16 years. Here the full orbit is visible, so Kepler's law gives the mass directly.

::: steps Weighing Sgr A* with one star
$ M = \frac{a^3}{P^2}
In units of AU (the Earth–Sun distance), years and solar masses, Kepler's third law takes this form, because Earth's orbit has a = 1, P = 1, M = 1. The star's own mass is negligible.
---
$ a \approx 1030\,\text{AU}, \quad P \approx 16.0\text{ years}
S2's semi-major axis, from its angular size on the sky (0.125 arcseconds) and the distance to the galactic centre (about 26,000 light-years).
---
$ M \approx \frac{1030^3}{16.0^2} \approx 4.3 \times 10^6\text{ solar masses}
Over four million Suns, inside a region S2 approaches to within about 120 AU.
:::

Nothing but a black hole can pack that mass into that volume and stay dark; a cluster of stars or dead stars would be bright or would have broken up long ago. Genzel and Ghez shared the 2020 Nobel Prize in Physics for this, alongside Penrose.

## Hearing a merger

On 14 September 2015 the two LIGO detectors in the United States recorded GW150914, a **gravitational wave**: a ripple in spacetime that stretches and squeezes distances as it passes, here by less than a thousandth of a proton's width over 4 km. The waveform matched relativity's prediction for two black holes of about 36 and 29 solar masses spiralling together 1.3 billion light-years away, merging into one of about 62. The missing 3 solar masses left as gravitational-wave energy in a fraction of a second. The signal's final "ringdown" matched the vibration of a single Kerr (rotating) black hole settling down. Rainer Weiss, Barry Barish and Kip Thorne received the 2017 Nobel Prize; hundreds of mergers have been detected since.

## Seeing the shadow

Light passing close to a black hole is bent strongly. At r = 1.5 r_s, the **photon sphere**, light can orbit the hole. Light rays aimed within a certain distance of the centre spiral in; those aimed further out escape. Against bright surrounding gas, this produces a dark disc, the **shadow**:

::: equation Apparent radius of the shadow (non-rotating hole)
b = \frac{3\sqrt{3}}{2} r_s \approx 2.6\, r_s
:::

The shadow looks larger than the horizon itself, because the hole bends light toward us from all around. Spin changes the size by only a few percent.

The Event Horizon Telescope linked radio dishes across the planet into an Earth-sized instrument. In April 2019 it released an image of M87*, at the centre of the galaxy M87, 55 million light-years away: a ring of glowing gas about 42 microarcseconds across around a dark centre. Its size matches a shadow of diameter 5.2 r_s for 6.5 × 10⁹ solar masses. In May 2022 it released an image of Sgr A*, a ring about 52 microarcseconds across, consistent with the mass found from stellar orbits.

::: note
Every method infers the black hole from its gravity. Orbits give the mass, X-ray flicker or the shadow size gives an upper limit on the size, and once enough mass sits inside its own Schwarzschild radius, general relativity leaves no alternative.
:::

::: quiz
? Why is the mass function only a lower limit on the unseen mass?
- Because Doppler shifts are imprecise
+ Because sin i ≤ 1 and the visible star's mass adds to the denominator
- Because the orbit is not circular
= f = M₂³ sin³i/(M₁ + M₂)², which can never exceed M₂.
? Why can S2's orbit give Sgr A*'s mass directly, without a mass function?
+ The orbit is resolved on the sky, so its size and tilt are measured
- S2 is much heavier than Sgr A*
- Sgr A* emits X-rays
= Measuring the full ellipse gives a and i, so Kepler's law yields M directly.
? The shadow of a non-rotating black hole has radius about
- r_s
- 1.5 r_s
+ 2.6 r_s
= b = (3√3/2) r_s ≈ 2.6 r_s; 1.5 r_s is the photon sphere's radius.
:::
