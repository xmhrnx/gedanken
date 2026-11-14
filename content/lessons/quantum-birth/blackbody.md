Quantum theory began not with atoms but with the colour of hot things. A poker in a fire glows dull red, then orange, then yellow-white as it heats. By 1900 physicists had measured the light from hot bodies precisely, and classical physics, applied carefully, predicted something absurd: that every hot object should blaze with infinite energy at short wavelengths. Max Planck fixed the formula with a trick he did not believe in, and that trick was the first quantum.

## Blackbodies

A **blackbody** is an object that absorbs all light falling on it. In 1860 Gustav Kirchhoff showed that such a body, at a given temperature, must glow with a spectrum that depends only on that temperature, not on what it is made of. A good laboratory version is a closed oven with a small hole: light entering the hole bounces around inside until it is absorbed, and the light coming out is the oven's own thermal glow.

Experiments found two simple rules. The total power radiated grows as T⁴ (the Stefan–Boltzmann law). And the wavelength of peak brightness is inversely proportional to temperature, λ_peak × T ≈ 2.90 × 10⁻³ m·K (Wien's law). The Sun's surface, at about 5800 K, peaks near 500 nm, green light; your body, at 310 K, peaks near 9 μm, in the infrared.

## The ultraviolet catastrophe

Classical physics treats the light inside the oven as standing electromagnetic waves, **modes**, that fit between the walls. Two classical results then fix the spectrum. Counting how many modes fit gives a number per unit volume, per unit frequency, of 8πf²/c³: shorter waves fit in more ways, so the count grows with frequency. And the equipartition theorem of thermal physics says that at temperature T every mode carries, on average, the same energy kT, where k = 1.38 × 10⁻²³ J/K is Boltzmann's constant.

::: equation The Rayleigh–Jeans law
u(f) = \frac{8\pi f^2}{c^3}\,kT
:::

Here u(f) is the energy per unit volume per unit frequency. It matches experiment at low frequencies. At high frequencies it fails completely: it grows without limit as f², so the total energy, summed over all frequencies, is infinite. A warm oven would pour out unlimited ultraviolet light. Paul Ehrenfest later called this the **ultraviolet catastrophe**. Real spectra rise, peak, and fall back to zero.

## Planck's quantum

Planck found the right formula first, by fitting the data, and then searched for a derivation. The one that worked required an assumption: an oscillator of frequency f can hold energy only in whole multiples of a basic unit hf.

::: steps Average energy of a mode with quantized energy
$ E_n = n\,hf, \quad n = 0, 1, 2, \dots
Planck's assumption. h is a new constant of nature; its value is fixed by matching the data.
---
$ P_n \propto \mathrm{e}^{-E_n/kT} = \mathrm{e}^{-nx}, \quad x = \frac{hf}{kT}
Boltzmann's rule from statistical physics: in thermal equilibrium the probability of a state of energy E falls off as e^(−E/kT).
---
$ Z = \sum_{n=0}^{\infty}\mathrm{e}^{-nx} = \frac{1}{1 - \mathrm{e}^{-x}}
The probabilities must add to 1, so divide each by their sum Z. The sum is a geometric series with ratio e^(−x), less than 1.
---
$ \langle n \rangle = \frac{1}{Z}\sum_{n=0}^{\infty} n\,\mathrm{e}^{-nx} = -\frac{\mathrm{d}\ln Z}{\mathrm{d}x} = \frac{1}{\mathrm{e}^{x} - 1}
The average n is each n weighted by its probability. Differentiating ln Z with respect to x pulls down a factor −n from each term, which gives the weighted sum.
---
$ \langle E \rangle = \frac{hf}{\mathrm{e}^{hf/kT} - 1}
Multiply by hf. This replaces the classical kT.
---
$ hf \ll kT: \ \langle E \rangle \approx kT \qquad hf \gg kT: \ \langle E \rangle \approx hf\,\mathrm{e}^{-hf/kT}
At low frequency e^x ≈ 1 + x and the classical answer returns. At high frequency even a single quantum hf costs far more than the typical thermal energy kT, so such modes are almost never excited and their energy dies away exponentially.
:::

Multiplying by the number of modes gives Planck's law:

::: equation Planck's radiation law
u(f) = \frac{8\pi h f^3}{c^3}\,\frac{1}{\mathrm{e}^{hf/kT} - 1}
:::

The catastrophe is cured because high-frequency modes are frozen out: they need a large chunk of energy all at once, and the thermal jostling of the walls rarely supplies it. Planck's law reproduces the measured spectra at every frequency, and Wien's and Stefan's laws follow from it.

::: note
Energy at frequency f comes in packets of size E = hf. Because the packets are tiny at low frequency, classical physics works there; at high frequency the packets are too expensive to fill, and the spectrum falls to zero.
:::

The constant is h ≈ 6.626 × 10⁻³⁴ J·s. For visible light, f ≈ 5 × 10¹⁴ Hz, one quantum is about 3 × 10⁻¹⁹ J, or 2 eV (an electronvolt, eV, is the energy an electron gains crossing 1 volt, 1.60 × 10⁻¹⁹ J). That is tiny by everyday standards, which is why nobody had noticed.

::: history 1900 · Berlin
At the Physikalisch-Technische Reichsanstalt in Berlin, Otto Lummer, Ernst Pringsheim, Heinrich Rubens and Ferdinand Kurlbaum measured blackbody spectra deep into the infrared, where they showed that Wilhelm Wien's earlier formula failed. On 7 October 1900 Rubens told Planck the latest results during a visit to his home; that evening Planck found the formula that fitted them all. He presented a derivation on 14 December 1900, and from the fit extracted values for both h and k that are within a few percent of today's. Planck, a cautious classical physicist, later called the quantization "an act of desperation" and for years regarded it as a property of the material oscillators, not of light itself. Einstein was the first to take it literally, as the next lesson shows.
:::

The best blackbody ever measured is the universe itself: the cosmic microwave background, the afterglow of the Big Bang, follows Planck's law at 2.725 K so precisely that in the COBE satellite's 1990 measurement the error bars were smaller than the thickness of the plotted curve.

::: model wave
At a fixed speed, raise the frequency and the wavelength shrinks. Shorter waves fit into a box in more ways, which is why the classical count of modes grows without limit.
:::

::: quiz
? What was the ultraviolet catastrophe?
- Ultraviolet light from the Sun was found to be dangerous
+ Classical physics predicted infinite radiated energy at high frequencies
- Blackbodies were observed to emit no ultraviolet light at all
= Giving every one of infinitely many modes an energy kT makes the total energy diverge, with most of it at short wavelengths.
? In Planck's theory, why do high-frequency modes carry almost no energy?
- They escape from the oven too quickly
- There are fewer of them than low-frequency modes
+ A single quantum hf is much larger than kT, so they are rarely excited at all
= The probability of exciting even one quantum goes as e^(−hf/kT), which is tiny when hf ≫ kT. (There are actually more high-frequency modes, not fewer.)
? A star's spectrum peaks at 290 nm, in the ultraviolet. About how hot is its surface?
- 1000 K
+ 10,000 K
- 1,000,000 K
= Wien's law: T ≈ 2.90 × 10⁻³ m·K / 2.90 × 10⁻⁷ m = 10,000 K.
:::
