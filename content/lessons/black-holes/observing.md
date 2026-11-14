A black hole emits essentially nothing of its own: its Hawking radiation is far too faint to detect, and no light escapes from inside its horizon. So it cannot be photographed the way a star can. What can be seen is its effect on everything around it: stars orbiting something invisible, gas heated to millions of degrees as it spirals in, ripples in spacetime when two black holes merge, and a dark silhouette against glowing gas. Michell's suggestion of 1783, to watch a visible star circling an unseen companion, turned out to be exactly the way in. This lesson works through how each kind of evidence is turned into a measurement, and why, taken together, they leave no reasonable alternative.

## How do you weigh something you cannot see?

The first strong candidate emerged in the early 1970s. A bright source of X-rays in the constellation Cygnus, named Cygnus X-1, had been known since a rocket flight in 1964, but X-ray detectors of the time could not pin down its position precisely. In 1971 the Uhuru X-ray satellite found that its brightness flickered rapidly, and the same year a radio source switched on at the same moment the X-rays changed, which fixed its position in the sky. At that spot sits a blue supergiant star catalogued as HDE 226868. Its spectral lines, the dark lines in its light that act as fingerprints of its elements, shift back and forth every 5.6 days. That periodic shift is a Doppler effect: the star is moving toward us and away from us in turn. It is orbiting something, and no telescope can see what.

The flickering already says something about the size of that unseen companion. An object cannot change its brightness as a whole faster than light can cross it, because the different parts of it cannot coordinate any faster. The X-rays from Cygnus X-1 vary in milliseconds, and in a millisecond light travels 300 km. So the X-rays come from a region no more than a few hundred kilometres across: the companion is small and compact. The remaining question was its mass.

Here there is a difficulty. We cannot see the orbit itself; the system is far too distant. All we can measure is how fast the visible star moves toward and away from us, from the Doppler shift of its lines, and how long it takes to go round. We do not know the tilt of the orbit, and we do not know the visible star's mass exactly. Remarkably, Kepler's laws still put a firm lower limit on the unseen mass.

Call the visible star's mass M₁, the unseen companion's M₂, the orbital period P and the inclination i, the tilt of the orbit relative to our line of sight, with i = 90° meaning we see it edge-on. Assume the orbit is circular, which is very nearly true for Cygnus X-1.

::: steps The mass function
$ G(M_1 + M_2)\,P^2 = 4\pi^2 a^3
Start from Kepler's third law in the form Newton derived, where a is the distance between the two bodies. It links the period to the total mass and the size of the orbit. On its own it does not help, because we know neither a nor the masses.
---
$ a_1 = \frac{M_2}{M_1 + M_2}\,a
Both bodies actually circle their common centre of mass, the balance point between them. The balance condition is M₁a₁ = M₂a₂, and the two orbit radii add up to the separation, a₁ + a₂ = a. Solving these together gives the radius of the visible star's orbit, a₁. A heavy companion pulls the visible star round in a large circle.
---
$ K = \frac{2\pi a_1}{P}\,\sin i
The visible star's orbital speed is the circumference of its circle, 2πa₁, divided by the period. But the Doppler effect only measures motion along our line of sight, so we see only the fraction sin i of that speed. The measured amplitude of the Doppler wobble, the largest speed toward or away from us, is called K.
---
$ a_1 \sin i = \frac{KP}{2\pi}
Rearrange to put the unknowns on the left. The right side contains only quantities read directly from the spectrum, K and P. So this is a measurement of the projected size of the visible star's orbit, even though the orbit itself cannot be seen.
---
$ a_1^3\sin^3 i = \frac{M_2^3}{(M_1 + M_2)^3}\,a^3\sin^3 i = \frac{M_2^3\sin^3 i}{(M_1 + M_2)^3}\cdot\frac{G(M_1 + M_2)P^2}{4\pi^2}
Cube the left side of the previous line and substitute the two earlier results into it: a₁ from the centre-of-mass condition, then a³ from Kepler's law. One factor of (M₁ + M₂) cancels between the numerator and denominator. Everything not involving the masses is now a known combination of G and P.
---
$ f = \frac{P K^3}{2\pi G} = \frac{M_2^3\,\sin^3 i}{(M_1 + M_2)^2}
Now set this equal to the cube of KP/2π and collect terms, moving G and P to the measured side. The left side holds only measured quantities; the right side holds the unknown masses and tilt. The left side, f, is called the **mass function**, and it can be calculated from the spectrum alone.
---
$ M_2 \geq f
Finally, look at the right side. The factor sin³i is at most 1, and the denominator (M₁ + M₂)² is at least M₂², so the right side can never be larger than M₂. The mass function is therefore a strict lower limit on the unseen mass, and the true mass is larger the more massive the visible star is and the more the orbit is tilted away from edge-on.
:::

For Cygnus X-1, with P = 5.6 days and K ≈ 75 km/s, the mass function is about 0.25 solar masses. That alone proves little; a small star would satisfy it. But the visible star is a supergiant, with a spectrum showing it to be many times heavier than the Sun, and putting a realistic M₁ into the equation forces M₂ to be much larger than f. Early estimates in 1972 put the companion above about 6 solar masses; a 2021 analysis, with an improved distance, gives about 21. That is far above the roughly 2 to 2.3 solar mass limit for neutron stars from the previous lesson. Nothing known except a black hole is that heavy, that small and that dark.

::: history 1974 · a bet
In 1974, with the evidence for Cygnus X-1 strong but not conclusive, Stephen Hawking made a bet with his friend Kip Thorne that it did not contain a black hole. Hawking explained that this was an insurance policy. He had invested much of his career in black holes, and if they turned out not to exist, at least he would have the consolation of winning the bet; the prize was a subscription to a magazine. By 1990 the evidence had piled up so far that Hawking conceded, breaking into Thorne's office to sign the framed bet.
:::

## What lies at the centre of our galaxy?

The strongest evidence of all comes from much closer to home. Since the early 1990s two teams, one led by Reinhard Genzel in Germany and one by Andrea Ghez in California, have tracked individual stars in the innermost region of the Milky Way, within a few light-weeks of its centre. The centre is hidden from ordinary telescopes by dust, so they observed in infrared light, which passes through it, and used techniques that correct for the blurring of the atmosphere to see individual stars. Over years, the stars were seen swinging around an invisible point that coincides with a compact radio source called Sagittarius A*, written Sgr A*.

One star, named S2, completes an orbit in about 16 years, so by now it has been followed around more than once. Here, unlike Cygnus X-1, the full ellipse of the orbit is visible on the sky, so its size and tilt can be measured directly, and Kepler's third law gives the mass with no lower-limit ambiguity.

::: steps Weighing Sgr A* with one star
$ M = \frac{a^3}{P^2}
Kepler's third law takes this simple form if distances are measured in astronomical units (AU, the average Earth–Sun distance), times in years and masses in solar masses. The reason is that the Earth's orbit has a = 1 AU, P = 1 year and a central mass of 1 solar mass, so the constant in the law is 1 in these units. The star's own mass is negligible next to the central mass, so M is just the mass of Sgr A*.
---
$ a \approx 0.125'' \times 8{,}200\,\text{pc} \approx 1{,}030\,\text{AU}
Next find S2's semi-major axis, half the long axis of its ellipse. On the sky it spans about 0.125 arcseconds. An object 1 AU across seen from a distance of 1 parsec spans one arcsecond, and the galactic centre is about 8,200 parsecs (about 27,000 light-years) away. So the angle times the distance gives about 1,030 AU, about 34 times the distance from the Sun to Neptune.
---
$ M \approx \frac{1030^3}{16.0^2} \approx 4.3 \times 10^6\ M_{\text{Sun}}
Put the semi-major axis and the period of 16.0 years into Kepler's law. The cube of 1,030 is about 1.09 × 10⁹, and dividing by 256 gives over four million solar masses, all inside the region that S2 comes closest to, about 120 AU from the centre at the nearest point of its orbit.
:::

Pause on that result. Four million suns' worth of mass sits inside a region smaller than our own solar system, and it gives out almost no light. Could it be something other than a black hole? A cluster of four million ordinary stars packed that tightly would blaze brightly. A cluster of dead stars, such as white dwarfs or neutron stars, would be dark but would collide and scatter within a time far shorter than the age of the galaxy, so it could not last. Nothing but a black hole can pack that much mass into that volume, stay dark and remain stable. Genzel and Ghez shared the 2020 Nobel Prize in Physics for this work, alongside Roger Penrose.

## What did the first gravitational wave reveal?

On 14 September 2015 the two LIGO detectors in the United States recorded a **gravitational wave**, a ripple in spacetime that stretches and squeezes distances as it passes. The event was named GW150914, after its date. The wave changed the length of each 4 km detector arm by about 4 × 10⁻¹⁸ m, a few hundred times smaller than a proton. Its shape matched the prediction of general relativity for two black holes of about 36 and 29 solar masses spiralling together 1.3 billion light-years away and merging into a single black hole of about 62 solar masses. The missing 3 solar masses had been converted into gravitational-wave energy in a fraction of a second.

What made this signal persuasive was its final part. After the merger, the waveform showed a brief "ringdown", a rapidly fading oscillation, and its frequency and decay rate matched those calculated for a single rotating black hole, described by Roy Kerr's 1963 solution of Einstein's equations, settling into its final shape. No other kind of object of that mass could orbit so closely before merging. Rainer Weiss, Barry Barish and Kip Thorne received the 2017 Nobel Prize in Physics, and hundreds of mergers have been detected since.

## Can we see a black hole's shadow?

Light passing close to a black hole is bent strongly, and close enough in, it can even go round in circles. That leads to the most direct evidence of all: a picture. To understand what such a picture should show, follow a ray of light coming in from far away, aimed so that it would miss the centre by a distance b, its impact parameter.

::: steps The photon sphere and the shadow
$ \left(\frac{\mathrm{d}r}{\mathrm{d}\lambda}\right)^2 = \frac{1}{b^2} - \frac{1}{r^2}\left(1 - \frac{r_s}{r}\right)
For light in the Schwarzschild metric, two quantities are conserved along its path, because the geometry does not change with time or with direction around the hole: an energy and an angular momentum. Their ratio is the impact parameter b. Combining these with ds² = 0 gives this equation for how the ray's distance r changes along its path, where λ is a parameter that marks progress along the ray. The left side cannot be negative, so the ray can only travel where 1/b² is at least as large as the second term.
---
$ V(r) = \frac{1}{r^2}\left(1 - \frac{r_s}{r}\right)
Call that second term V(r). It acts like a hill the light must climb. Far away it is small; it rises to a peak as r decreases, then falls to zero at r = r_s. An incoming ray turns back where V(r) reaches 1/b². If 1/b² is larger than the peak of the hill, the ray never turns back and plunges into the black hole.
---
$ \frac{\mathrm{d}V}{\mathrm{d}r} = -\frac{2}{r^3} + \frac{3r_s}{r^4} = 0 \;\Rightarrow\; r = \tfrac{3}{2}\,r_s
To find the top of the hill, write V as 1/r² − r_s/r³, take its derivative with respect to r and set it to zero. Multiplying through by r⁴ leaves −2r + 3r_s = 0, so the peak is at r = 1.5 r_s. Light at exactly this radius, moving sideways, can circle the black hole; this radius is called the **photon sphere**. The orbit is unstable, so the slightest nudge sends the light inward or outward.
---
$ V\!\left(\tfrac{3}{2}r_s\right) = \frac{4}{9r_s^2}\cdot\frac{1}{3} = \frac{4}{27\,r_s^2}
Evaluate the height of the hill at its peak. At r = 1.5 r_s, 1/r² is 4/(9r_s²) and 1 − r_s/r is 1 − 2/3 = 1/3. Their product is 4/(27r_s²). This is the largest value of 1/b² for which an incoming ray can still turn back.
---
$ b_c = \sqrt{\frac{27}{4}}\,r_s = \frac{3\sqrt{3}}{2}\,r_s \approx 2.6\,r_s
Set 1/b² equal to the peak height and solve for b. Rays aimed with an impact parameter smaller than this critical value b_c clear the hill and fall in; rays aimed farther out swing round and escape. So, seen from far away, the black hole blocks a disc of radius b_c in the sky.
:::

Against a background of bright gas, then, a black hole should appear as a dark disc called its **shadow**, with an apparent radius of about 2.6 r_s. That is noticeably larger than the horizon itself, because the hole's gravity bends light toward us from all around, magnifying its own image. The edge of the shadow is outlined by a thin bright ring of light that has skimmed close to the photon sphere. For a spinning black hole the shadow is slightly flattened and shifted, but its size changes by only a few percent, so the size is a good measure of the mass.

Even the largest black holes are tiny on the sky, because they are so far away. Seeing one needed a telescope the size of the Earth. The Event Horizon Telescope achieved this by linking radio dishes on several continents, recording signals simultaneously and combining them, so that together they act as a single instrument with the resolving power of a dish as wide as the planet. In April 2019 it released an image of M87*, the black hole at the centre of the galaxy M87, about 55 million light-years away: a ring of glowing gas about 42 microarcseconds across surrounding a dark centre. The size matches the shadow expected for a black hole of about 6.5 billion solar masses, a mass consistent with measurements from the motions of stars in that galaxy. In May 2022 it released an image of Sgr A*, a ring about 52 microarcseconds across, exactly as expected for the 4.3 million solar masses found from the orbits of stars like S2.

::: note
Every method infers the black hole from its gravity. Orbits give the mass; rapid flickering or the size of the shadow puts an upper limit on the size. Once enough mass is shown to sit inside its own Schwarzschild radius, and no known form of matter can support it there, general relativity leaves no alternative.
:::

::: quiz
? Why is the mass function only a lower limit on the unseen companion's mass?
- Because Doppler shifts are imprecise
+ Because sin i is at most 1 and the visible star's mass adds to the denominator
- Because the orbit is not circular
= The mass function is f = M₂³ sin³i / (M₁ + M₂)². The tilt factor can only shrink it and the visible star's mass can only enlarge the denominator, so f can never exceed M₂.
? Why can S2's orbit give Sgr A*'s mass directly, without a mass function?
+ The orbit is seen on the sky, so its size and tilt are measured
- S2 is much heavier than Sgr A*
- Sgr A* emits X-rays
= Because the whole ellipse is resolved, both the semi-major axis a and the tilt are known. Kepler's third law then gives M = a³/P² with no unknown factor left.
? The shadow of a non-rotating black hole has a radius of about
- r_s
- 1.5 r_s
+ 2.6 r_s
= Rays with an impact parameter below b = (3√3/2) r_s ≈ 2.6 r_s fall in, which sets the shadow's size. The value 1.5 r_s is the radius of the photon sphere, where light can orbit.
:::
