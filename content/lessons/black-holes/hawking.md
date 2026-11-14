By the early 1970s a black hole looked like the simplest object in physics: a horizon that absorbs everything and emits nothing, described completely by its mass, spin and charge. Then in 1974 Stephen Hawking applied quantum theory to the space around the horizon and found that a black hole should glow.

## Why quantum theory changes the answer

General relativity treats spacetime classically. The rest of physics is quantum: every particle is an excitation of a **quantum field** that fills space, and even in its lowest-energy state, the vacuum, a field is not perfectly still but fluctuates.

What counts as "the vacuum", and what counts as a particle, depends on the observer. Hawking studied a quantum field in the spacetime of a star collapsing to a black hole. He compared the field's state long before collapse, when it was empty, with what a distant observer detects long after the horizon forms. The horizon cuts the field into a part outside and a part inside, and the outside part, on its own, is not empty. The distant observer receives a steady flux of particles with a **thermal spectrum**, the same distribution of energies as glowing matter at a definite temperature.

A common picture says pairs of "virtual particles" appear near the horizon, one falls in and the other escapes. It captures that the energy comes from the black hole and that the horizon separates partners, but it is a heuristic, not the calculation: the emitted radiation has wavelengths comparable to the whole black hole, so it is not produced at any sharp point at the horizon.

The energy carried away is paid for by the black hole's mass, which slowly decreases.

## The Hawking temperature

Hawking's result is that a black hole of mass M radiates like a body at temperature:

::: equation Hawking temperature
T = \frac{\hbar c^3}{8\pi G M k_B}
:::

Here ħ is the reduced Planck constant (quantum mechanics), c and G come from relativity and gravity, and k_B is Boltzmann's constant, which converts energy to temperature. All four corners of physics meet in one formula, which is why it is taken as a key clue for any future theory of quantum gravity.

The formula is less mysterious than it looks: the typical energy of an emitted particle, k_B T, corresponds to a wavelength set by the black hole's size (a few tens of times r_s), and that size is proportional to M. Bigger hole, longer wavelength, lower temperature.

::: steps Temperature of a solar-mass black hole
$ T = \frac{1.055 \times 10^{-34} \times (2.998 \times 10^8)^3}{8\pi \times 6.674 \times 10^{-11} \times M \times 1.381 \times 10^{-23}}
Insert the constants in SI units.
---
$ T \approx \frac{1.23 \times 10^{23}\,\text{K}\cdot\text{kg}}{M}
Evaluate the constants: numerator 2.84 × 10⁻⁹, denominator 2.31 × 10⁻³² per kilogram of M.
---
$ T \approx \frac{1.23 \times 10^{23}}{1.989 \times 10^{30}} \approx 6.2 \times 10^{-8}\,\text{K}
Use one solar mass. Sixty billionths of a degree above absolute zero.
:::

That is far colder than the 2.7 K cosmic microwave background, the afterglow of the Big Bang that fills space. A real stellar black hole therefore absorbs more radiation from the sky than it emits, and grows. Hawking radiation from such holes is hopelessly undetectable.

## Small black holes are hot

The temperature goes as 1/M. That is the opposite of ordinary objects, which get hotter when you add energy. Here, adding mass makes the black hole colder, and losing mass makes it hotter, so it radiates faster still. The process runs away.

How long does it take? Treat the hole as a hot sphere radiating by the Stefan–Boltzmann law (power = σ × area × T⁴, where σ is a constant).

::: steps How the evaporation time scales with mass
$ T \sim \frac{1}{M}
From the Hawking formula.
---
$ A \sim r_s^2 \sim M^2
The horizon's area is 4πr_s², and r_s = 2GM/c² is proportional to M.
---
$ P \sim A T^4 \sim M^2 \times M^{-4} = M^{-2}
Radiated power by Stefan–Boltzmann. Lighter holes are hotter by more than enough to make up for being smaller.
---
$ \frac{\mathrm{d}M}{\mathrm{d}t} = -\frac{P}{c^2} \sim -\frac{1}{M^2}
Radiated energy comes out of the mass through E = mc².
---
$ M^2\,\mathrm{d}M \sim -\mathrm{d}t
Rearrange so each side holds one variable.
---
$ t_{\text{evap}} \sim M^3
Integrate: M³/3 on the left falls from its initial value to zero over the lifetime. Doubling the mass makes the lifetime eight times longer.
:::

With all constants included (and only photons counted), t_evap ≈ 5120πG²M³/(ħc⁴). For one solar mass that is about 2 × 10⁶⁷ years, against a universe 1.4 × 10¹⁰ years old. A black hole would only be hotter than today's microwave background if it were lighter than about 4.5 × 10²² kg, less than the Moon. One that was light enough to finish evaporating within the age of the universe would have had to start at roughly 10¹¹ to 10¹² kg, a mountain's mass. Such holes could only have formed in the early universe, and none have been seen.

::: note
Hawking radiation follows from ordinary quantum field theory placed in the curved spacetime of a black hole. It has never been observed from a real black hole, because for any star-sized one it is fantastically faint. Its importance is theoretical: it shows black holes have a temperature and entropy, so gravity, quantum mechanics and thermodynamics must fit together.
:::

::: history 1974 · Cambridge, England
Jacob Bekenstein had argued in 1972 that a black hole should have entropy proportional to its horizon area. Hawking initially thought this was wrong, since anything with entropy and energy should have a temperature and glow. His calculation, meant in part to settle the matter, showed that it does glow. It was published in *Nature* in March 1974 under the title "Black hole explosions?"
:::

The radiation also raised a puzzle still argued about: if a black hole evaporates into purely thermal radiation, what happens to the information about what fell in? This is the black hole information problem.

::: quiz
? A black hole's mass is halved. Its Hawking temperature
- halves
+ doubles
- is unchanged
= T = ħc³/(8πGMk_B), so T ~ 1/M.
? Black hole A has ten times the mass of black hole B. How do their evaporation times compare?
- A lasts 10 times longer
- A lasts 100 times longer
+ A lasts 1000 times longer
= t_evap ~ M³, and 10³ = 1000.
? Why do stellar-mass black holes in today's universe grow rather than evaporate?
+ They are colder than the cosmic microwave background, so they absorb more than they emit
- Hawking radiation only works for small black holes
- The horizon reflects their radiation back
= At about 6 × 10⁻⁸ K they are far colder than the 2.7 K background.
:::
