By the early 1970s a black hole looked like the simplest object in physics. Its horizon swallows anything that crosses it and lets nothing out, and theorists had proved that once a black hole settles down, it is described completely by just three numbers: its mass, its spin and its electric charge. Everything else about the star that formed it is lost from view. Then, in 1974, Stephen Hawking applied quantum theory to the space just outside the horizon and found that this picture is incomplete. A black hole should glow, with a definite temperature, and in doing so it should slowly lose mass. This lesson explains where the glow comes from, works out how hot a black hole is, and follows the surprising consequence that small black holes get hotter as they shrink.

## Why should quantum theory change the answer?

General relativity is a classical theory: it treats spacetime as a smooth geometry with definite properties, and particles as things that follow definite paths through it. The rest of physics is quantum. In quantum field theory, the framework that describes all known particles, every kind of particle is an excitation of a **quantum field** that fills all of space; a photon, for example, is a single excitation of the electromagnetic field. Even in its lowest possible energy state, which we call the vacuum, a quantum field is not perfectly still. The uncertainty principle forbids it from having exactly zero value everywhere, so it constantly fluctuates.

A subtle consequence is that what counts as "empty" can depend on who is looking. The vacuum is defined as the state with no particles, and particles are defined by how the field oscillates in time. Observers whose clocks run differently, or who cannot see the same regions of spacetime, can divide the same field into oscillations in different ways, and so disagree about whether particles are present.

Hawking studied a quantum field in the spacetime of a star collapsing to form a black hole. Long before the collapse, he took the field to be in its vacuum state, with no particles anywhere. He then asked what a distant observer would detect long after the horizon had formed. The key is that the horizon divides the field into two parts, one outside and one inside, and the outside observer has no access to the inside part. The fluctuations of the field on the two sides are correlated with each other. When the inside part is hidden, the outside part, considered on its own, is no longer empty. The distant observer receives a steady stream of particles flowing outward from the vicinity of the black hole, and their energies follow a **thermal spectrum**: exactly the distribution of energies emitted by an object glowing at one definite temperature.

A popular picture describes this with pairs of "virtual particles" that pop into existence near the horizon; one falls in, the other escapes, and the escaping one is the radiation. The picture captures two true features: the energy carried away comes from the black hole, and the horizon separates partners that are correlated. But it is a heuristic, not the calculation, and it is misleading in an important way. The emitted radiation has wavelengths much larger than the black hole itself, so it is not produced at any sharp point at the horizon. It emerges from a broad region around the hole, and the particle-pair picture says nothing about why the spectrum should be thermal. The real calculation compares how the field oscillates before and after the collapse and finds the thermal spectrum directly.

The energy carried away by the radiation must come from somewhere, and it comes from the black hole's mass, which slowly decreases.

## How hot is a black hole?

Hawking's result is that a black hole of mass M radiates like an object at temperature:

::: equation Hawking temperature
T = \frac{\hbar c^3}{8\pi G M k_B}
:::

Each constant in it comes from a different part of physics. The reduced Planck constant ħ comes from quantum mechanics; the speed of light c from relativity; G from gravity; and Boltzmann's constant k_B, which converts energy to temperature, from thermodynamics. All four meet in one formula, and if any one of them is set aside, the effect disappears: without ħ, for instance, the temperature is zero, which is the classical answer. That is why the formula is treated as a vital clue to any future theory that unites gravity with quantum mechanics.

Its form is less mysterious than it looks. Radiation from an object tends to have wavelengths comparable to the size of the object, and the size of a black hole is set by its Schwarzschild radius, r_s = 2GM/c², which is proportional to the mass. The energy of a photon is inversely proportional to its wavelength, and temperature measures the typical energy of the radiation. So a bigger hole emits longer wavelengths, which means lower-energy photons and a lower temperature: T is proportional to 1/M. The detailed calculation supplies the numerical factor, and the peak of the spectrum turns out to sit at a wavelength about 16 times the Schwarzschild radius.

What does the formula give for a real black hole?

::: steps Temperature of a solar-mass black hole
$ T = \frac{1.055 \times 10^{-34} \times (2.998 \times 10^8)^3}{8\pi \times 6.674 \times 10^{-11} \times 1.381 \times 10^{-23} \times M}
Put the values of the constants into the formula, all in SI units: ħ in joule-seconds, c in metres per second, G in its SI units and k_B in joules per kelvin. The mass M stays as a symbol for now, so the result can be reused for any black hole.
---
$ T \approx \frac{2.84 \times 10^{-9}}{2.32 \times 10^{-32}\,M} \approx \frac{1.23 \times 10^{23}\,\text{K kg}}{M}
Evaluate the numerator and denominator separately. The numerator, ħc³, is about 2.84 × 10⁻⁹; the denominator, without the M, is about 2.32 × 10⁻³². Dividing gives a single constant: a black hole's temperature in kelvin is 1.23 × 10²³ divided by its mass in kilograms.
---
$ T \approx \frac{1.23 \times 10^{23}}{1.989 \times 10^{30}} \approx 6.2 \times 10^{-8}\,\text{K}
Now set M equal to one solar mass, 1.989 × 10³⁰ kg. The result is about sixty billionths of a degree above absolute zero. Black holes formed from collapsing stars are at least a few solar masses, so they are colder still.
:::

That is extraordinarily cold. It is far colder than the 2.7 K cosmic microwave background, the faint afterglow of the Big Bang that fills all of space. Any object colder than its surroundings absorbs more radiation than it emits, so a real stellar black hole today takes in more energy from the sky than it radiates away, and grows. Its own Hawking radiation is hopelessly undetectable, swamped by the background and by every other source of light near it.

## Why do small black holes get hotter?

The formula T ∝ 1/M has a strange consequence. For ordinary objects, adding energy raises the temperature: put a kettle on the stove and it heats up. A black hole does the opposite. Adding mass makes it colder; losing mass makes it hotter. So a black hole that is radiating more than it absorbs loses mass, gets hotter, radiates faster, loses mass faster still, and the process runs away. In the language of thermodynamics it has a negative heat capacity.

How long does the whole process take? A rough answer comes from treating the black hole as a hot sphere the size of its horizon, radiating according to the **Stefan–Boltzmann law**, which says that the power radiated by a hot body is σAT⁴, where A is the surface area, T the temperature and σ a constant. We only want to know how the lifetime depends on the mass, so constants are dropped.

::: steps How the evaporation time scales with mass
$ T \propto \frac{1}{M}
Start from the Hawking formula. Everything in it apart from M is a constant of nature, so the temperature is simply proportional to 1/M: halve the mass and the temperature doubles.
---
$ A \propto r_s^2 \propto M^2
The radiating surface is the horizon, a sphere of area 4πr_s². Since r_s = 2GM/c² is proportional to M, the area is proportional to M². A lighter black hole is smaller and has less surface to radiate from.
---
$ P \propto A\,T^4 \propto M^2 \times M^{-4} = M^{-2}
Combine these in the Stefan–Boltzmann law. The area brings in M², the fourth power of the temperature brings in M⁻⁴, and together they give M⁻². So a lighter hole radiates more power, not less: its higher temperature more than makes up for its smaller surface.
---
$ \frac{\mathrm{d}M}{\mathrm{d}t} = -\frac{P}{c^2} \propto -\frac{1}{M^2}
The energy carried off by the radiation comes out of the black hole's mass, through E = mc². So the rate at which the mass decreases is the radiated power divided by c². This is a differential equation: it relates the rate of change of M to M itself.
---
$ M^2\,\mathrm{d}M \propto -\mathrm{d}t
Multiply both sides by M² and by dt, so that everything involving M is on the left and everything involving time is on the right. This separation is what allows each side to be added up, or integrated, on its own.
---
$ \frac{M_0^3}{3} \propto t_{\text{evap}}
Integrate both sides. On the left, M² integrates to M³/3, evaluated from the starting mass M₀ down to zero, which gives M₀³/3. On the right, the time adds up from zero to the moment the hole vanishes, t_evap. So the total lifetime is proportional to the cube of the starting mass.
---
$ t_{\text{evap}} \propto M_0^3
Dropping the constant factor, this is the result. Its meaning is that lifetime rises steeply with mass: doubling the mass makes the lifetime eight times longer, and a hole ten times heavier lasts a thousand times longer. Almost all of a black hole's life is spent large and cold; the end, when it is small and hot, comes very quickly.
:::

With all the constants included, and counting only photons as the emitted particles, the lifetime is t_evap ≈ 5120πG²M³/(ħc⁴). For one solar mass that comes to about 2 × 10⁶⁷ years. The universe is about 1.4 × 10¹⁰ years old, so a solar-mass black hole has not lost a measurable fraction of its mass, and in any case, as we saw, it is growing.

Turn the question around. How small would a black hole have to be to be hotter than the microwave background today, so that it shrinks instead of grows? Setting T = 2.7 K in the formula gives M ≈ 4.5 × 10²² kg, a little less than the mass of the Moon. And how small would one have to have been to finish evaporating within the age of the universe? The answer is roughly 10¹¹ to 10¹² kg, about the mass of a rocky asteroid half a kilometre to a kilometre across, packed into a region no bigger than a proton. No star could make such a hole; the only conceivable origin is the extreme density of the very early universe. Astronomers have searched for the final bursts of radiation such "primordial" black holes would emit as they finish evaporating, and have not found any.

::: note
Hawking radiation follows from ordinary quantum field theory applied in the curved spacetime around a black hole. It has never been observed from a real black hole, because for any hole formed from a star it is fantastically faint. Its importance is theoretical: it shows that black holes have a temperature and an entropy, so gravity, quantum mechanics and thermodynamics must fit together in one framework.
:::

::: history 1972–1974 · Princeton and Cambridge
In 1972 Jacob Bekenstein, a graduate student at Princeton working with John Wheeler, was troubled by a thought experiment Wheeler had posed: throw a cup of hot tea into a black hole, and its entropy, the measure of its disorder, seems to vanish from the universe, violating the second law of thermodynamics. Bekenstein proposed that the black hole itself must carry entropy, proportional to the area of its horizon. He was inspired by a result Hawking had proved shortly before: in any classical process, the total area of black-hole horizons never decreases, just as entropy never decreases. Hawking thought Bekenstein's proposal was wrong, and said so firmly. Anything that has both energy and entropy, he pointed out, must have a temperature, and anything with a temperature must radiate; but black holes, by definition, emit nothing. He set out to do the calculation properly, applying quantum field theory near a horizon, partly to settle the matter. To his own surprise he found that black holes do radiate, with exactly the thermal spectrum needed to make Bekenstein's entropy consistent. He published the result in Nature in March 1974, under the title "Black hole explosions?", the question mark reflecting how startling it seemed.
:::

The discovery also raised a puzzle that is still argued about. If a black hole eventually evaporates completely into radiation that is purely thermal, the radiation carries no record of what fell in: a black hole made from encyclopedias and one made from the same mass of hydrogen would leave identical glows. But quantum mechanics, as normally understood, never destroys information; in principle the past can always be reconstructed from the present. Either the radiation is not exactly thermal and somehow carries the information out, or quantum mechanics must be modified near black holes. This is the **black hole information problem**, and resolving it is one of the main goals of research into quantum gravity.

::: quiz
? A black hole's mass is halved. What happens to its Hawking temperature?
- It halves
+ It doubles
- It is unchanged
= The temperature is T = ħc³/(8πGMk_B), which is proportional to 1/M. Halving the mass therefore doubles the temperature.
? Black hole A has ten times the mass of black hole B. How do their evaporation times compare?
- A lasts 10 times longer
- A lasts 100 times longer
+ A lasts 1,000 times longer
= The evaporation time is proportional to M³, and 10³ = 1,000. The cube comes from the power falling as 1/M² while the mass to be radiated grows as M.
? Why do black holes formed from stars grow rather than evaporate in today's universe?
+ They are colder than the cosmic microwave background, so they absorb more than they emit
- Hawking radiation only happens for small black holes
- The horizon reflects their radiation back inside
= A solar-mass black hole has a temperature of about 6 × 10⁻⁸ K, far below the 2.7 K of the background radiation. Like any object colder than its surroundings, it takes in more energy than it gives out.
:::
