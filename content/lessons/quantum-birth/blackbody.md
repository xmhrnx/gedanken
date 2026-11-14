Quantum theory began not with atoms or electrons but with the colour of hot things. An iron poker left in a fire glows a dull red, then orange, then a yellowish white as it gets hotter. By the late nineteenth century this glow had become an urgent practical question, because the electric lighting industry wanted to know how to make lamps that turned more heat into light, and a precise theoretical one, because the shape of the glow turned out to be universal. Physicists measured it with great care. Then, when classical physics was applied to it with equal care, it gave an answer that was not merely inaccurate but absurd. The repair, made by Max Planck in 1900, required an assumption he did not himself believe: that energy comes in indivisible lumps. This lesson explains the problem, the failed classical answer, and how Planck's lumps cured it.

## Why should every hot object glow the same way?

Everything warm gives off electromagnetic radiation, because the charged particles in it are jostling with thermal motion, and moving charges radiate. At room temperature the radiation is infrared, invisible to us; above about 500 °C enough of it is visible for the object to glow. Different materials glow somewhat differently, because a shiny surface reflects some radiation instead of absorbing and re-emitting it. So physicists considered an ideal case.

A **blackbody** is an object that absorbs all the radiation falling on it, at every wavelength, reflecting none. In 1860 Gustav Kirchhoff proved, from the laws of thermodynamics, that a blackbody at a given temperature must glow with a spectrum, a distribution of brightness across wavelengths, that depends only on its temperature: not on its size, shape or what it is made of. His argument was that if two blackbodies of different materials at the same temperature glowed differently, you could set them facing each other and heat would flow from one to the other on its own, which the second law of thermodynamics forbids. A universal spectrum was a clue to something fundamental, and finding its formula became a major goal.

A good laboratory blackbody turns out to be a hollow oven with a small hole in its wall. Light entering the hole bounces around inside, absorbed a little at each bounce, and almost none finds its way out again, so the hole absorbs nearly perfectly. When the oven is heated, the radiation streaming out of the hole is the oven's own thermal glow, with exactly the blackbody spectrum. For this reason blackbody radiation is also called cavity radiation.

Experiments with such ovens established two rules. The total power radiated grows as the fourth power of the absolute temperature, T⁴; this is the Stefan–Boltzmann law, found by Josef Stefan from measurements in 1879 and derived from thermodynamics by Ludwig Boltzmann in 1884. Doubling the temperature makes a body radiate sixteen times as much. And the wavelength at which the glow is brightest is inversely proportional to the temperature: λ_peak × T ≈ 2.90 × 10⁻³ m·K, a rule found by Wilhelm Wien and called Wien's law. That is why a heating poker shifts from red toward white: the peak moves to shorter wavelengths. The Sun's surface, at about 5800 K, peaks near 500 nm, in the green, though it looks white because it emits strongly at every visible colour. Your body, at about 310 K, peaks near 9 μm, in the infrared, which is what thermal cameras detect. What nobody had was a formula for the full shape of the curve, and a reason for it.

## What did classical physics predict?

The classical approach starts from the radiation inside the oven. Electromagnetic waves bouncing between the walls settle into standing waves, patterns that fit exactly between the walls, the way a vibrating guitar string can only hold waves that fit a whole number of half-wavelengths along its length. Each allowed pattern is called a **mode**. The calculation then needs two ingredients.

The first is a count of how many modes there are. Shorter waves can fit into the oven in many more ways than long ones, because there are more ways to arrange many small wiggles in three dimensions than a few large ones. Counting carefully gives the number of modes per unit volume in a small range of frequency as 8πf²/c³ per unit frequency, which grows with the square of the frequency f.

The second ingredient is how much energy each mode holds. Here classical thermal physics has a clear answer, the **equipartition theorem**: in thermal equilibrium at temperature T, energy is shared out equally, and every mode of vibration carries on average the same energy kT, where k ≈ 1.38 × 10⁻²³ J/K is Boltzmann's constant, the number that converts temperature into energy per particle. Multiply the two together:

::: equation The Rayleigh–Jeans law
u(f) = \frac{8\pi f^2}{c^3}\,kT
:::

Here u(f) is the energy per unit volume of the oven, per unit range of frequency. At low frequencies this formula matches the measurements well. At high frequencies it fails completely, and not by a little. It says the energy keeps rising as f², without limit. Add up the energy over all frequencies and the total is infinite. A warm oven, or a cup of tea, should be pouring out unlimited amounts of ultraviolet light and X-rays. The real spectrum does nothing of the kind: it rises, reaches a peak, and falls back toward zero at high frequency. In 1911 Paul Ehrenfest named the failure the **ultraviolet catastrophe**.

The reasoning behind the catastrophe was not careless. The mode count is simple geometry, and equipartition was a well-tested result of the kinetic theory of gases. That is what made it serious: two of the most trusted parts of classical physics, combined in the most straightforward way, gave nonsense.

## How did Planck fix the formula?

Planck worked the other way round. Rather than starting from classical principles, he looked for a formula that matched the measurements, found one in October 1900, and then searched for a derivation that would justify it. The only derivation he could make work required an assumption with no basis in classical physics: an oscillator of frequency f can hold energy only in whole-number multiples of a basic unit hf, where h is a new constant. Planck applied this to the charged oscillators he imagined in the oven walls; today the same calculation is applied directly to each mode of the radiation, and that is the version followed here.

::: steps Average energy of a mode whose energy comes in lumps
$ E_n = n\,hf, \quad n = 0, 1, 2, \dots
This is Planck's assumption. A mode of frequency f can hold no energy, one lump hf, two lumps, and so on, but nothing in between. The constant h is new, and its value has to be found by matching the formula to experiment.
---
$ P_n \propto \mathrm{e}^{-E_n/kT} = \mathrm{e}^{-nx}, \quad x = \frac{hf}{kT}
In thermal equilibrium, the chance of finding a system in a state of energy E falls off as e^(−E/kT); this is Boltzmann's rule from statistical physics, and it holds classically too. Higher-energy states are exponentially less likely. Writing x for the ratio of one lump hf to the typical thermal energy kT keeps the algebra short.
---
$ Z = \sum_{n=0}^{\infty}\mathrm{e}^{-nx} = \frac{1}{1 - \mathrm{e}^{-x}}
The probabilities have to add up to 1, so each must be divided by the sum of all of them, called Z. The sum is a geometric series, 1 + r + r² + …, with ratio r = e^(−x), which is less than 1, so it adds up to 1/(1 − r).
---
$ \langle n \rangle = \frac{1}{Z}\sum_{n=0}^{\infty} n\,\mathrm{e}^{-nx} = -\frac{\mathrm{d}\ln Z}{\mathrm{d}x}
The average number of lumps is each possible n weighted by its probability. There is a neat way to find this sum: differentiating e^(−nx) with respect to x brings down a factor −n, so differentiating Z term by term gives minus the weighted sum we want. Dividing by Z turns this into minus the derivative of ln Z.
---
$ \langle n \rangle = \frac{1}{\mathrm{e}^{x} - 1}
Since ln Z = −ln(1 − e^(−x)), its derivative is −e^(−x)/(1 − e^(−x)). Changing the sign and multiplying top and bottom by e^x gives this compact result for the average number of lumps in a mode.
---
$ \langle E \rangle = \frac{hf}{\mathrm{e}^{hf/kT} - 1}
Each lump carries energy hf, so multiplying the average number of lumps by hf gives the average energy of a mode. This is what replaces the classical kT, which gave every mode the same energy regardless of its frequency.
---
$ hf \ll kT: \ \langle E \rangle \approx kT \qquad hf \gg kT: \ \langle E \rangle \approx hf\,\mathrm{e}^{-hf/kT}
Look at the two extremes. At low frequency, x is small and e^x ≈ 1 + x, so the average energy is about kT, the classical answer. At high frequency, a single lump hf costs far more than the typical thermal energy kT, so the mode is almost never excited at all, and its average energy dies away exponentially.
:::

Multiplying the average energy of a mode by the number of modes gives Planck's law:

::: equation Planck's radiation law
u(f) = \frac{8\pi h f^3}{c^3}\,\frac{1}{\mathrm{e}^{hf/kT} - 1}
:::

The catastrophe is cured, and the reason is worth putting in words. There are still more and more modes at higher frequencies, just as classical physics said. But each high-frequency mode can only take energy in a large lump, all at once, and the random thermal jostling of the oven rarely supplies that much in one go. Those modes are, in effect, frozen out. Low-frequency modes, with tiny lumps, behave classically. The peak of the spectrum falls where the two effects balance, where hf is a few times kT. Planck's law matches the measured spectra at every frequency, and both the Stefan–Boltzmann law and Wien's law follow from it as consequences.

The constant has the value h ≈ 6.626 × 10⁻³⁴ J·s, now called Planck's constant. For visible light, with f ≈ 5 × 10¹⁴ Hz, one lump is about 3 × 10⁻¹⁹ J. Physicists usually express such small energies in electronvolts: one **electronvolt** (eV) is the energy an electron gains in crossing a potential difference of one volt, 1.60 × 10⁻¹⁹ J. A lump of visible light is about 2 eV. That is minute by everyday standards, which is why nobody had noticed that energy came in lumps: a light bulb emits about 10²⁰ of them every second, and their graininess is invisible.

::: note
Radiation of frequency f exchanges energy in lumps of size hf. At low frequencies the lumps are small compared with the thermal energy kT and classical physics works, but at high frequencies a single lump costs far more than the heat can supply, so those modes stay nearly empty and the spectrum falls to zero instead of rising forever.
:::

::: history 1900 · Berlin
The decisive measurements were made at the Physikalisch-Technische Reichsanstalt, Germany's national standards laboratory in Berlin, partly to set standards for the lighting industry. Otto Lummer and Ernst Pringsheim, and then Heinrich Rubens and Ferdinand Kurlbaum, measured blackbody spectra ever deeper into the infrared, and found that a formula Wien had proposed in 1896, which had looked correct, failed at long wavelengths. On Sunday 7 October 1900 Rubens and his wife visited the Plancks for lunch, and Rubens mentioned the newest results. That evening Planck found a formula that joined Wien's, correct at short wavelengths, to the long-wavelength behaviour Rubens had seen, and sent it to Rubens on a postcard. It fitted every measurement. Planck then spent what he later called the most strenuous weeks of his life looking for a reason, and presented his derivation on 14 December 1900. From the fit he extracted values of both h and k within a few per cent of today's. He was a cautious, classically minded physicist, and years later described the assumption of energy lumps as "an act of desperation", which he expected someone to remove; for a long time he thought of it as a property of the oscillators in the walls, not of light itself. The ultraviolet catastrophe, despite the way the story is often told, was not what drove him: its full force was spelled out after 1900, by Lord Rayleigh, James Jeans and Einstein. Einstein was also the first to take the lumps literally, as the next lesson shows.
:::

The best blackbody ever measured is the universe itself. The cosmic microwave background, the faint radiation left over from the hot early universe, fills all of space and follows Planck's law at a temperature of about 2.725 K. When the COBE satellite measured its spectrum in 1990, the data points fitted Planck's curve so well that the error bars were smaller than the width of the line drawn through them.

::: model wave
At a fixed wave speed, raise the frequency and the wavelength shrinks. Shorter waves can fit into a box in more ways, which is why the classical count of modes keeps growing at high frequency.
:::

::: quiz
? What was the ultraviolet catastrophe?
- The discovery that ultraviolet light from the Sun is harmful
+ The classical prediction that a hot body radiates infinite energy at high frequencies
- The observation that blackbodies emit no ultraviolet light at all
= Classical physics gave each of infinitely many modes the same average energy kT, and the number of modes grows with frequency. The total energy then comes out infinite, with most of it at short wavelengths.
? In Planck's theory, why do high-frequency modes carry almost no energy?
- They escape from the oven too quickly
- There are fewer of them than low-frequency modes
+ A single lump of energy hf is much larger than kT, so they are rarely excited at all
= The chance of exciting even one lump goes as e^(−hf/kT), which is tiny when hf is much larger than kT. There are in fact more high-frequency modes than low-frequency ones; they are simply almost all empty.
? A star's spectrum peaks at 290 nm, in the ultraviolet. About how hot is its surface?
- 1,000 K
+ 10,000 K
- 1,000,000 K
= Wien's law gives T ≈ (2.90 × 10⁻³ m·K)/(2.90 × 10⁻⁷ m) = 10,000 K. A peak at a shorter wavelength than the Sun's means a hotter surface than the Sun's 5800 K.
:::
