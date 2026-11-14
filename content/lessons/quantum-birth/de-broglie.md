By the early 1920s light had a split personality. Interference showed it was a wave; the photoelectric effect and Compton scattering showed it delivered energy and momentum in particle-like lumps. In 1924 a French graduate student, Louis de Broglie, asked the question in reverse: if waves can behave like particles, can particles behave like waves? He proposed that every electron, every atom, every moving thing has a wavelength.

## The momentum of a photon

The starting point is light. Special relativity says a massless particle carries momentum p = E/c. Combine that with Planck's E = hf:

::: steps Wavelength and momentum of a photon
$ E = pc
For anything with zero mass, the relativistic energy–momentum relation E² = (pc)² + (mc²)² reduces to E = pc.
---
$ E = hf
Planck and Einstein: a photon of frequency f carries energy hf.
---
$ p = \frac{hf}{c}
Set the two expressions for E equal and divide by c.
---
$ p = \frac{h}{\lambda}
For light, c = fλ, so f/c = 1/λ. The photon's momentum is set by its wavelength alone.
:::

Compton's X-ray experiment in 1923 had confirmed exactly this momentum. De Broglie's step was to read the last equation backwards and apply it to matter.

## Matter waves

::: equation The de Broglie wavelength
\lambda = \frac{h}{p} = \frac{h}{mv}
:::

For anything everyday, the wavelength is absurdly small. A baseball of 0.145 kg at 40 m/s has p ≈ 5.8 kg·m/s and λ ≈ 1.1 × 10⁻³⁴ m, far smaller than a proton. No experiment could reveal waves that fine, which is why ordinary objects show no wave behaviour.

Electrons are different, because their mass is tiny. An electron accelerated from rest through a voltage V gains kinetic energy eV.

::: steps Wavelength of an electron accelerated through V volts
$ \frac{p^2}{2m} = eV
Kinetic energy ½mv², written in terms of momentum p = mv, equals the energy gained from the voltage. This is accurate when eV is much less than the electron's rest energy, 511 keV.
---
$ p = \sqrt{2meV}
Solve for p.
---
$ \lambda = \frac{h}{\sqrt{2meV}}
Use λ = h/p.
---
$ \lambda = \frac{6.63 \times 10^{-34}}{\sqrt{2 \times 9.11 \times 10^{-31} \times 1.60 \times 10^{-19} \times 54}} \approx 1.67 \times 10^{-10}\,\text{m}
For 54 volts, the value in the experiment below: about 0.17 nm.
:::

That is comparable to the spacing between atoms in a crystal. A crystal's regular rows of atoms should therefore act on electrons as a diffraction grating, a set of evenly spaced obstacles, acts on light: waves scattered from different rows reinforce each other only in particular directions.

::: model wave
Change the wavelength and frequency of a travelling wave. In de Broglie's relation, a faster particle has more momentum and so a shorter wavelength.
:::

## Bohr's rule explained

De Broglie's idea gave an immediate payoff. Picture the electron's wave wrapped around a Bohr orbit. For the wave to join smoothly onto itself, rather than cancelling out after many laps, the circumference must hold a whole number of wavelengths:

::: equation Standing waves on a circle
2\pi r = n\lambda = n\,\frac{h}{mv} \;\Rightarrow\; mvr = n\,\frac{h}{2\pi} = n\hbar
:::

This is exactly Bohr's quantization rule, which had been pure assumption. The picture is still too simple: the real electron wave in an atom spreads in three dimensions, as Schrödinger's equation showed two years later. But it made clear that quantization comes from waves fitting into a confined space, like the notes of a guitar string.

::: note
Every particle has a wavelength λ = h/p. Wave behaviour shows up only when this wavelength is comparable to the size of whatever the particle meets, which happens for electrons and atoms, never for baseballs.
:::

::: history 1925–1927 · New York and Aberdeen
At Bell Telephone Laboratories, Clinton Davisson and Lester Germer were studying how electrons bounce off nickel, for reasons to do with vacuum tubes. In 1925 a flask of liquid air broke, air rushed into the hot apparatus, and the nickel target oxidized. To clean it they heated it for a long time, which, unknown to them, turned it from many tiny crystals into a few large ones. Afterwards the scattered electrons came off in sharp peaks at particular angles. Davisson learned of de Broglie's idea in 1926 at a meeting in Oxford, where Max Born showed Davisson's own earlier data as possible evidence for electron waves. Back in New York, he and Germer searched deliberately and in 1927 found a strong peak at 50° for 54 eV electrons, where the spacing between nickel's surface rows, 0.215 nm, gives a wavelength of 0.215 nm × sin 50° = 0.165 nm, matching de Broglie's 0.167 nm. The same year George Paget Thomson in Aberdeen saw diffraction rings from electrons passing through thin metal films. Davisson and Thomson shared the 1937 Nobel Prize; de Broglie had received his in 1929. Thomson's father, J. J. Thomson, had won his in 1906 for showing that the electron is a particle.
:::

## Waves of what?

Electron diffraction is now routine: electron microscopes use wavelengths far shorter than light's to resolve individual atoms. Neutrons, helium atoms, and in 1999 even C₆₀ molecules of 60 carbon atoms have been made to interfere. In 1989 Akira Tonomura's group at Hitachi sent electrons through a double-slit arrangement one at a time and watched an interference pattern build up from single dots.

That last experiment shows the puzzle. Each electron arrives at one spot, like a particle, but where they land follows a wave pattern. De Broglie did not say what was waving. The answer, given by Max Born in 1926, is that the wave measures probability, and it is the starting point of quantum mechanics in the next chapter.

::: quiz
? A proton and an electron move at the same speed. Which has the longer de Broglie wavelength?
+ The electron, because its smaller mass gives it less momentum
- The proton, because it is heavier
- They are equal, because the speeds are equal
= λ = h/mv. At equal speed the electron's momentum is about 1800 times smaller, so its wavelength is about 1800 times longer.
? If the accelerating voltage on an electron is increased fourfold, its wavelength
- doubles
+ halves
- falls to a quarter
= λ = h/√(2meV) ∝ 1/√V, so four times the voltage gives half the wavelength.
? Why do we never see a thrown ball diffract through a doorway?
- Large objects have no wavelength
+ Its wavelength, around 10⁻³⁴ m, is far too small compared with the doorway
- Air molecules destroy the wave
= Diffraction is noticeable only when the wavelength is comparable to the opening. A ball's wavelength is some 34 orders of magnitude smaller than a door.
:::
