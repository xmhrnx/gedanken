Pass sunlight through a prism and it spreads into a rainbow. Look closely, with a narrow slit and a good lens, and the rainbow turns out to be crossed by hundreds of thin dark lines, as if certain colours had been cut out. Heat a pinch of table salt in a colourless flame and you see the opposite: instead of a rainbow, a few bright lines on a dark background, dominated by a vivid yellow. These lines turned out to be fingerprints. Every chemical element produces its own set, and no two sets are alike, so reading them tells you what a substance is made of even when you cannot touch it, even when it is a star. This lesson explains where the lines come from, why hydrogen's follow a simple numerical pattern, and why the same lines appear whether a gas is giving out light or taking it in.

## Three kinds of spectrum

A **spectrum** is light spread out according to its wavelength, the distance between successive crests of the light wave. Visible light runs from about 400 nanometres (nm), violet, to about 700 nm, red. A prism or a finely ruled grating sends each wavelength in a slightly different direction, laying the colours out side by side so they can be measured.

In the middle of the nineteenth century, Gustav Kirchhoff summed up what laboratories had found as three rules. A hot, dense object, such as a glowing solid or the deep, opaque layers of a star, gives a **continuous spectrum**, an unbroken rainbow containing every wavelength. A hot, thin gas gives an **emission spectrum** instead, with bright lines at a few particular wavelengths and darkness in between. And when a cooler, thin gas sits in front of a source of continuous light, the result is an **absorption spectrum**: the full rainbow, but with dark lines at exactly the wavelengths that the same gas would emit if it were hot.

That last rule is the remarkable one. A gas absorbs precisely the colours it can emit, and no others. The dark lines in sunlight are the shadows of the gases in the Sun's cooler outer layers, absorbing light from the hotter layers beneath.

::: history 1814–1868 · Munich and Heidelberg
Joseph von Fraunhofer was an optician in Munich, trying to make better glass for telescope lenses. To measure precisely how his glass bent different colours, he needed sharp reference points in the spectrum, and in 1814, examining sunlight through a prism and a small telescope, he found them: dark lines, 574 of which he mapped. He labelled the strongest with letters, but neither he nor anyone else knew what they were. The answer came more than forty years later in Heidelberg, from the physicist Gustav Kirchhoff and the chemist Robert Bunsen, whose burner gave a hot, nearly colourless flame ideal for examining the light from heated substances. In 1859 they showed that each element, put into the flame, gives its own pattern of bright lines, and that a cool vapour of the same element absorbs light at exactly the same wavelengths. Sodium's bright yellow pair sat exactly where Fraunhofer's dark pair, labelled "D", lay in sunlight. There was sodium in the Sun. Within two years, finding lines that matched no known element, they discovered two new metals, caesium and rubidium, named from the Latin for the colours of their lines, sky blue and deep red. In 1868, astronomers observing a solar eclipse found a yellow line in light from the Sun's edge that matched no element at all. It was attributed to a new element, named helium after the Greek word for the Sun, and helium was not found on Earth until 1895.
:::

## Hydrogen's pattern

If spectral lines are fingerprints, hydrogen, the simplest atom, should have the simplest one, and it does. In visible light it shows just four lines, at 656, 486, 434 and 410 nm, red, blue-green, blue and violet. For years they were simply a list of measured numbers. Then in 1885 Johann Balmer, a Swiss mathematics teacher at a girls' school, found a formula that fitted all four to within the accuracy of the measurements. In 1888 the Swedish physicist Johannes Rydberg rewrote it in a more general form:

::: equation The Rydberg formula
\frac{1}{\lambda} = R\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right), \qquad R \approx 1.097 \times 10^{7}\,\text{m}^{-1}
:::

Here λ is the wavelength, n₁ and n₂ are whole numbers with n₂ larger than n₁, and R, now called the Rydberg constant, is a number found by fitting the measurements. Balmer's four visible lines all have n₁ = 2, with n₂ = 3, 4, 5 and 6. The formula predicted further lines, at other values of n₁, that were later found exactly where it said.

But a formula that works is not an explanation, and this one was deeply puzzling. Why should whole numbers, squared, control the colours of light from a gas? For almost thirty years nobody knew.

## Where the lines come from

The answer came in 1913 from Niels Bohr, and it rests on two ideas that have survived every later revision. First, an atom can exist only in certain states of definite energy, its **energy levels**, and not in between. Second, an atom gives out light by jumping from a higher level to a lower one, emitting a single **photon**, a quantum of light, which carries away exactly the energy difference. Bohr's picture of electrons on circular orbits turned out to be wrong, as the orbitals lesson showed, but his energy levels for hydrogen are correct; Schrödinger's equation produces exactly the same ones. Only the levels matter here.

With those two ideas, the Rydberg formula can be derived rather than fitted.

::: steps From energy levels to wavelengths
$ E_n = -\frac{13.6\,\text{eV}}{n^2}
These are hydrogen's energy levels, labelled by the whole number n = 1, 2, 3 and so on. The energies are in electronvolts (eV), the energy an electron gains crossing one volt. They are negative because the electron is bound: zero energy means an electron just barely freed from the atom, so 13.6 eV is what it takes to remove the electron from the lowest level.
---
$ E_{\text{photon}} = E_{n_2} - E_{n_1} = 13.6\,\text{eV}\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right)
When the electron drops from a higher level n₂ to a lower level n₁, energy is conserved, so the photon carries off exactly the difference between the two levels. Subtracting the two expressions, the minus signs rearrange so that the result is positive, since n₂ is larger than n₁. The squares of whole numbers already appear, and now we can see why: they come from the levels themselves.
---
$ E_{\text{photon}} = hf = \frac{hc}{\lambda}
A photon's energy is set by its frequency f through Planck's constant h, as Planck and Einstein had established. For any wave, frequency times wavelength equals speed, so f = c/λ, where c is the speed of light. Writing the energy in terms of wavelength links the atom's levels to the colour we observe.
---
$ \frac{1}{\lambda} = \frac{13.6\,\text{eV}}{hc}\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right)
Set the two expressions for the photon's energy equal, since they describe the same photon, and divide both sides by hc. The result has exactly the form of Rydberg's formula. Using hc ≈ 1240 eV nm, the constant in front is 13.6/1240 ≈ 0.01097 per nanometre, which is 1.097 × 10⁷ per metre: Rydberg's R, which he had found purely by fitting data, is now calculated from the atom's energy levels.
---
$ n_2 = 3 \to n_1 = 2: \quad \frac{1}{\lambda} = R\left(\frac{1}{4} - \frac{1}{9}\right) = \frac{5R}{36} \quad\Rightarrow\quad \lambda \approx 656\,\text{nm}
As a check, take the smallest jump ending on level 2, from n = 3. The bracket gives 1/4 − 1/9 = 5/36, and inverting 5R/36 gives about 656 nm, Balmer's red line. It is called hydrogen-alpha, and it gives the glowing gas clouds in astronomical photographs their characteristic pink-red colour.
:::

The derivation also explains why hydrogen's lines come in families, called **series**. All the jumps that end on the same lower level form one series. Jumps down to n = 2 make up the Balmer series, and they happen to land in visible light, which is why they were found first. Jumps down to n = 1, the Lyman series, are bigger, because the gap between level 1 and the rest is large, at least 10.2 eV, so their photons lie in the ultraviolet. Jumps down to n = 3, the Paschen series, are smaller and lie in the infrared.

::: model bohr
Pick a jump between two levels and see the wavelength and colour of the photon it produces. Bigger drops make more energetic photons with shorter wavelengths; only the drops that end on n = 2 give visible light.
:::

## Why absorption matches emission

Absorption is the same process run backwards. An electron in a lower level can absorb a photon and jump to a higher level, but only if the photon's energy exactly matches the gap between the two levels, because there are no levels in between for it to land on. A photon with the wrong energy cannot be absorbed by that atom at all, and passes straight through.

Now put a cool gas in front of a source of continuous light. Of all the photons streaming through it, the gas removes only those whose energies match its own level gaps. Those are exactly the photons it would emit when hot, since emission and absorption involve the same jumps between the same levels. So the gas cuts dark lines into the rainbow at the same wavelengths where it would otherwise shine. The absorbed energy is re-emitted soon afterwards, but in every direction, so very little of it continues toward the observer. This is Kirchhoff's third rule, and it is why Fraunhofer's dark lines and Bunsen's bright lines coincide.

## Every element its own

Atoms with more than one electron have more complicated level structures, and no simple formula like Rydberg's describes them. The electrons repel one another, and as the orbitals lesson showed, the energy of a level depends on its angular momentum l as well as on n. But the levels are fixed entirely by the nuclear charge and the number of electrons, which is to say by the identity of the element. So every element has its own unique set of lines, and its spectrum identifies it as surely as a fingerprint identifies a person.

That makes spectroscopy one of the most powerful tools in science. In 1925 Cecilia Payne, a graduate student at Harvard, used stellar spectra to show that stars are made overwhelmingly of hydrogen and helium. Astronomers then believed the Sun had roughly the same composition as the Earth, and her conclusion was at first resisted, but it was right. Spectral lines also measure motion: a source moving away from us has all its lines shifted toward longer wavelengths, by the Doppler effect, so the lines give the speeds of stars and galaxies, and with them the expansion of the universe. On Earth, spectrometers identify pollutants in air and water, check the composition of steel in foundries, and, mounted on telescopes, look for water vapour and other molecules in the atmospheres of planets around other stars.

::: note
Atoms have discrete energy levels, so they emit and absorb light only at the wavelengths whose photon energies match the gaps between those levels. Each element's levels, and therefore its pattern of lines, are unique, so a spectrum identifies the elements present, even in a distant star.
:::

::: quiz
? Why does a cool gas in front of a hot star produce dark lines at the same wavelengths it would emit if it were hot?
- The gas reflects those wavelengths back toward the star
+ Absorption and emission both involve the same jumps between the same energy levels
- Cool gases block all wavelengths equally
= A photon is absorbed only if its energy matches a gap between two levels, and these are the same gaps that set the wavelengths of the emission lines.
? Which hydrogen transition gives the shorter wavelength?
+ n = 2 to n = 1
- n = 3 to n = 2
- They are equal
= The drop from 2 to 1 releases 10.2 eV, while the drop from 3 to 2 releases only 1.9 eV, and a more energetic photon has a shorter wavelength.
? Why do the Lyman lines, which end on n = 1, lie in the ultraviolet?
- Electrons move faster near the nucleus
+ The gaps down to n = 1 are large, at least 10.2 eV, so the photons have short wavelengths
- Hydrogen only emits ultraviolet light when it is very hot
= Since λ = hc/E, the smallest Lyman energy, 13.6 × ¾ = 10.2 eV, gives the longest Lyman wavelength, about 122 nm, which is already well into the ultraviolet.
:::
