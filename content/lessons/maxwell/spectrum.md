Maxwell's equations allow electromagnetic waves of any wavelength at all. Nothing in the derivation of their speed picked out a particular size of wave, and nothing stops a wave from being a kilometre long or far smaller than an atom. Visible light, the only part our eyes can detect, covers just a narrow band of wavelengths, from about 380 to 750 nanometres, where a nanometre is a billionth of a metre. Everything else, from radio waves longer than a football pitch to gamma rays smaller than an atomic nucleus, is the same phenomenon at a different scale. This lesson surveys that whole range, called the **electromagnetic spectrum**, and asks the obvious question it raises: if all these waves are the same kind of thing, why do they behave so differently?

## How are wavelength and frequency related?

Every electromagnetic wave travels through a vacuum at the same speed, c, about 3.00 × 10⁸ m/s. That single fact ties each wave's wavelength to its frequency. In one second, f crests pass any given point, and they are spaced a wavelength λ apart, so in that second the wave advances a distance fλ. That distance per second is the wave's speed.

::: equation Wavelength and frequency
c = f\lambda
:::

Because c is fixed, knowing either the wavelength or the frequency tells you the other, and the two go in opposite directions: a longer wavelength means a lower frequency, and a shorter wavelength a higher one. So the spectrum can be described equally well by either number. An FM radio station broadcasting at 100 MHz, or 10⁸ cycles per second, sends out waves 3 m long. A microwave oven running at 2.45 GHz uses waves about 12 cm long. Green light, with a wavelength of 550 nm, oscillates about 5.5 × 10¹⁴ times every second, a frequency far beyond anything an electric circuit of Maxwell's day could produce, which is part of why it took so long to make electromagnetic waves deliberately.

::: model wave
Change the wavelength and watch the frequency respond. For a wave of fixed speed, as light has in a vacuum, the frequency rises as the wavelength falls, so that their product stays the same.
:::

## What are the bands?

The spectrum is conventionally divided into bands. It is important to understand that the boundaries between them are not sharp physical dividing lines. The names were given at different times by different people, depending on how each kind of radiation was discovered and what it was used for, and neighbouring bands blend into each other. The wavelengths below are approximate.

| Band | Wavelength | Typical sources |
|---|---|---|
| Radio | longer than about 1 m | broadcast antennas, pulsars |
| Microwave | about 1 mm to 1 m | ovens, radar, the cosmic background |
| Infrared | about 750 nm to 1 mm | warm objects, your skin |
| Visible | about 380 to 750 nm | the Sun, flames, LEDs |
| Ultraviolet | about 10 to 380 nm | the Sun, hot stars, sterilising lamps |
| X-rays | about 0.01 to 10 nm | X-ray tubes, gas around black holes |
| Gamma rays | shorter than about 0.01 nm | radioactive nuclei, supernovae |

X-rays and gamma rays in fact overlap in wavelength, which shows how loose the labels are. In practice the two names usually refer to the source rather than the wavelength: radiation from atomic nuclei or high-energy particle processes is called gamma rays, while radiation produced by fast electrons is called X-rays, even when the two have the same wavelength.

## Why do the bands behave so differently?

If radio waves and gamma rays are the same kind of thing, it is puzzling that one passes harmlessly through your body while the other can damage the DNA in your cells. There are two parts to the explanation.

The first is about how waves interact with matter at all. A wave is absorbed when something in the material can take up its energy, and that happens most readily when the material contains charges that naturally oscillate at, or can make a jump that matches, the wave's frequency. Different bands match different things. Radio waves oscillate far too slowly to stir the electrons bound in molecules, so most non-conducting materials ignore them, which is why radio passes through walls; only conductors like metals, whose free electrons can follow the slow oscillation, block them strongly. Microwaves match the rotation of water molecules, which is how a microwave oven heats food. Infrared matches the vibrations of molecules, so warm objects give it off and absorb it readily. Visible light matches the energy needed to move electrons between levels in many molecules, which is why dyes have colours, and why the molecules in your retina can detect it. X-rays are so energetic that the loosely bound outer electrons of light atoms cannot absorb them efficiently, so they pass through soft tissue; they are stopped mainly by heavy elements such as the calcium in bone, whose tightly bound inner electrons can take up their energy. That is why bones show up on an X-ray image.

The second part of the explanation could not have come from Maxwell's theory at all. Light turns out to be absorbed and emitted not continuously but in separate lumps called **photons**, each carrying an amount of energy proportional to the wave's frequency. The constant of proportionality is **Planck's constant**, h ≈ 6.63 × 10⁻³⁴ joule-seconds. Why light should come in lumps, when Maxwell's equations describe it as a smooth wave, is one of the central puzzles of the quantum chapters. For now the formula is enough, because it tells us how much energy a single photon brings to a single molecule.

::: steps Photon energy from wavelength
$ E = hf
The energy of one photon is Planck's constant times the frequency of the light. This is a statement about each lump separately: a brighter beam of the same light has more photons, not more energetic ones.
---
$ E = \frac{hc}{\lambda}
Wavelengths are usually easier to quote than frequencies, so rewrite the formula in terms of λ. From c = fλ, the frequency is f = c/λ, and replacing f with this gives the energy directly from the wavelength. Shorter wavelengths mean more energy per photon.
---
$ hc \approx 1240\,\text{eV} \cdot \text{nm}
For atoms and molecules the joule is an inconveniently large unit, so physicists use the **electronvolt** (eV): the energy an electron gains when it crosses a voltage of one volt, about 1.60 × 10⁻¹⁹ J. Multiplying h by c and converting to electronvolts and nanometres gives this handy number. Chemical bonds, for comparison, hold atoms together with energies of a few electronvolts.
---
$ E \approx \frac{1240\,\text{eV nm}}{550\,\text{nm}} \approx 2.3\,\text{eV}
Now apply it to green light at 550 nm: divide 1240 by the wavelength in nanometres. A photon of visible light carries about two electronvolts, comparable to the energy of a chemical bond, which is why light can drive chemistry in your retina and in plant leaves.
---
$ E \approx \frac{1240\,\text{eV nm}}{0.1\,\text{nm}} \approx 12{,}000\,\text{eV}
The same formula for an X-ray of wavelength 0.1 nm gives about 12,000 electronvolts, thousands of times more than visible light. A photon like this can knock electrons out of any atom it hits.
:::

This explains the difference between radiation that merely warms you and radiation that harms you. What matters for damage is the energy delivered by each individual photon to each individual molecule, not the total amount of energy arriving. Radio photons carry around a millionth of an electronvolt or less, far too little to break any bond, so a powerful radio transmitter can only warm you, by shaking many molecules gently. Ultraviolet photons carry a few electronvolts, enough to break or rearrange bonds in DNA directly, which is what causes sunburn. Above roughly 10 eV, in the far ultraviolet and beyond, single photons can **ionise** atoms, tearing electrons out of them altogether and breaking molecules apart. So a weak beam of X-rays can do chemical damage that no amount of radio power could, because each X-ray photon arrives carrying a large lump of energy at once.

## What can we see from the ground?

The atmosphere absorbs most of the spectrum, and for the same reasons as above. Ozone high in the atmosphere absorbs most of the Sun's ultraviolet, whose photons are energetic enough to break its molecules. Water vapour and carbon dioxide absorb much of the infrared, because it matches their molecular vibrations. X-rays and gamma rays are stopped high up by the sheer quantity of air above us. Only two broad ranges of wavelength get through to the ground largely unhindered: visible light, and a wide band of radio waves.

That simple fact shaped the history of astronomy. Astronomy began with eyes, because visible light reaches the ground. Radio telescopes followed in the twentieth century, through the second window. But astronomers who wanted to study the universe in X-rays or gamma rays had to wait for rockets and satellites that could carry instruments above the atmosphere, and when they did, they found a sky full of objects, from exploding stars to hot gas around black holes, that had been invisible from the ground.

::: note
Radio waves, microwaves, infrared, visible light, ultraviolet, X-rays and gamma rays are all electromagnetic waves travelling at the same speed c. They differ only in wavelength, and therefore in frequency and in the energy carried by each photon, E = hf, which largely determines how each band interacts with matter.
:::

## How was the spectrum discovered?

::: history 1800–1914 · Europe
The spectrum was found piece by piece, and for a long time nobody knew the pieces belonged together. In 1800 William Herschel, the astronomer who had discovered the planet Uranus, was trying to find out which colours of sunlight carried the most heat. He spread sunlight into a rainbow with a prism and placed thermometers in each colour, and then, as a check, one just beyond the red end, where there was no visible light at all. That thermometer warmed fastest. Something invisible lay beyond the red: infrared. The next year Johann Ritter in Germany, reasoning that if there was something beyond one end of the spectrum there might be something beyond the other, used silver chloride, a chemical that darkens in light, and found that it darkened fastest just beyond the violet: ultraviolet. For most of the century these were thought of simply as invisible kinds of light. Hertz made radio waves deliberately in 1887, and showed they behaved like light. Wilhelm Röntgen discovered X-rays in 1895 when a screen across his laboratory glowed while he was operating a cathode-ray tube, and in 1900 Paul Villard found gamma rays coming from radium. At first nobody knew whether X-rays and gamma rays were waves at all. That was settled when they were shown to diffract, spreading and interfering like waves, from the regularly spaced atoms of crystals: X-rays in 1912, in experiments suggested by Max von Laue, and gamma rays in 1914, by Ernest Rutherford and Edward Andrade.
:::

::: quiz
? Compared with red light, blue light has
- a longer wavelength and a higher frequency
+ a shorter wavelength and a higher frequency
- the same frequency but a higher speed
= In a vacuum all colours travel at the same speed c, so c = fλ means a shorter wavelength must go with a higher frequency. Blue light has the shorter wavelength.
? Why can ultraviolet light cause sunburn while far more intense radio waves cannot?
+ Each ultraviolet photon carries enough energy to break chemical bonds, and radio photons do not
- Ultraviolet travels faster than radio
- Radio waves are not electromagnetic
= Chemical damage depends on the energy each photon delivers to a single molecule, E = hf, not on the total intensity. An ultraviolet photon carries a few electronvolts, enough to break a bond in DNA, while a radio photon carries only a tiny fraction of that.
? A radio wave has a wavelength of 300 m. What is its frequency?
- 300 Hz
- 1 kHz
+ 1 MHz
- 1 GHz
= Rearranging c = fλ gives f = c/λ = (3 × 10⁸ m/s) ÷ (300 m) = 10⁶ Hz, which is 1 MHz.
:::
