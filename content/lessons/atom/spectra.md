Pass sunlight through a prism and you get a rainbow. Look closely and the rainbow is crossed by hundreds of thin dark lines. Heat a pinch of salt in a flame and the light comes out in a few bright lines instead of a rainbow. These lines are fingerprints: every element has its own set, and reading them tells you what something is made of, even a star.

## Lines in the light

A **spectrum** is light spread out by wavelength. Three kinds appear in the laboratory, summarised by Gustav Kirchhoff:

- A hot, dense object, such as a glowing solid or the deep layers of a star, gives a **continuous spectrum**: all wavelengths.
- A hot, thin gas gives an **emission spectrum**: bright lines at a few particular wavelengths.
- A cooler, thin gas in front of a continuous source gives an **absorption spectrum**: the continuous rainbow with dark lines at exactly the wavelengths that gas would emit.

::: history 1814–1860 · Munich and Heidelberg
Joseph von Fraunhofer, an optician, mapped 574 dark lines in the solar spectrum in 1814 while testing glass for lenses. Nobody knew what they were. In 1859 Kirchhoff and the chemist Robert Bunsen, whose burner gave a nearly colourless flame, showed that each element emits its own pattern of bright lines, and that a cool vapour of the same element absorbs the same wavelengths. Sodium's yellow pair matched Fraunhofer's dark "D" lines: there was sodium in the Sun. Within two years they discovered two new elements, caesium and rubidium, from unknown lines. In 1868 a yellow line in the Sun's spectrum matched no known element; it was named helium and found on Earth only in 1895.
:::

## Hydrogen's pattern

Hydrogen, the simplest atom, has the simplest pattern. Its visible lines are at 656, 486, 434 and 410 nm. In 1885 Johann Balmer, a Swiss schoolteacher, found a formula that fitted them, and in 1888 Johannes Rydberg generalised it:

::: equation The Rydberg formula
\frac{1}{\lambda} = R\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right), \qquad R \approx 1.097 \times 10^{7}\,\text{m}^{-1}
:::

Here n₁ and n₂ are whole numbers with n₂ > n₁. Balmer's visible lines are n₁ = 2. Nobody knew why whole numbers should appear until Bohr explained it in 1913: an atom has discrete energy levels, and a line is a photon emitted as the electron drops from one level to another. Schrödinger's equation later gave the same levels without Bohr's orbits.

::: steps From energy levels to wavelengths
$ E_n = -\frac{13.6\,\text{eV}}{n^2}
The energy levels of hydrogen. Negative energy means the electron is bound; 13.6 eV is what it takes to remove it from the lowest level.
---
$ E_{\text{photon}} = E_{n_2} - E_{n_1} = 13.6\,\text{eV}\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right)
Energy is conserved: when the electron drops from level n₂ to a lower level n₁, the photon carries off the difference.
---
$ E_{\text{photon}} = hf = \frac{hc}{\lambda}
A photon's energy is set by its frequency, and frequency times wavelength is c.
---
$ \frac{1}{\lambda} = \frac{13.6\,\text{eV}}{hc}\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right)
Equate the two expressions and divide by hc. With hc = 1240 eV nm, the constant is 13.6/1240 nm⁻¹ = 1.097 × 10⁷ m⁻¹: Rydberg's R.
---
$ n_2 = 3 \to n_1 = 2: \quad \frac{1}{\lambda} = R\left(\frac{1}{4} - \frac{1}{9}\right) = \frac{5R}{36} \quad\Rightarrow\quad \lambda \approx 656\,\text{nm}
The first Balmer line, red, called hydrogen-alpha. It is the colour of glowing hydrogen clouds in astronomical photographs.
:::

Other values of n₁ give other series. Drops to n = 1 (the Lyman series) release more energy and lie in the ultraviolet; drops to n = 3 (Paschen) lie in the infrared. Only the Balmer series falls in visible light, which is why it was found first.

::: model bohr
Pick a jump between levels. Bigger drops make shorter-wavelength photons; drops ending on n = 2 are the visible ones.
:::

Absorption is the same process in reverse. An electron in level n₁ can absorb a photon only if its energy exactly matches the gap to some higher level. Light at other wavelengths passes straight through. That is why a cool gas removes just its own lines from a continuous spectrum.

Bohr's circular orbits are wrong, as the orbitals lesson showed, but his energy levels for hydrogen are right, and only the levels matter here.

## Every element its own

Atoms with more electrons have more complicated level structures, with no simple formula, because electrons repel each other and levels depend on l as well as n. But the levels are fixed by the nuclear charge and the electrons, so each element's set of lines is unique.

That makes spectra one of the most powerful tools in science. In 1925 Cecilia Payne showed from stellar spectra that stars are made mostly of hydrogen and helium, against the belief that they resembled Earth. Lines shifted by the Doppler effect give the speeds of stars and the expansion of the universe. Spectrometers identify pollutants, check the composition of steel, and look for water vapour in the atmospheres of other planets.

::: note
Atoms have discrete energy levels, so they emit and absorb light only at wavelengths whose photon energies match the gaps. Each element's levels, and so its lines, are unique: a spectrum identifies the elements present.
:::

::: quiz
? Why does a cool gas in front of a hot star produce dark lines at the same wavelengths it would emit when hot?
- The gas reflects those wavelengths back toward the star
+ Absorption and emission both involve the same jumps between the same energy levels
- Cool gases block all light equally
= A photon is absorbed only if its energy matches a gap between levels, the same gaps that set the emission lines.
? Which hydrogen transition gives the shorter wavelength?
+ n = 2 to n = 1
- n = 3 to n = 2
- They are equal
= 2 → 1 releases 10.2 eV, 3 → 2 only 1.9 eV. More energy means shorter wavelength.
? Why do the Lyman lines (ending on n = 1) lie in the ultraviolet?
- Electrons move faster near the nucleus
+ The gaps down to n = 1 are large, at least 10.2 eV, so the photons have short wavelengths
- Hydrogen only emits ultraviolet light when it is hot
= λ = hc/E, and E ≥ 13.6 × ¾ = 10.2 eV gives λ ≤ 122 nm.
:::
