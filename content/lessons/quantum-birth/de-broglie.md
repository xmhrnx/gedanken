By the early 1920s, light had a split personality. Interference and diffraction showed beyond doubt that it travels as a wave. The photoelectric effect, and then Arthur Compton's experiments with X-rays, showed just as clearly that it delivers energy and momentum in particle-like lumps. Physicists had no idea how one thing could be both, but they had at least got used to the idea that light was strange. In 1924 a French graduate student, Louis de Broglie, asked the question in reverse. If waves can behave like particles, can particles behave like waves? He proposed that every electron, every atom, every moving thing has a wavelength, and he gave a formula for it. This lesson follows his reasoning from the momentum of light to the wavelength of matter, shows why the waves of everyday objects are undetectable while those of electrons are not, and describes the experiments that found them.

## What is the momentum of a photon?

De Broglie's starting point was light, and the link between its wave properties and its particle properties. A photon has no mass, yet it carries momentum: light pushes on whatever absorbs or reflects it, an effect small enough to need sensitive instruments but large enough to steer spacecraft with sails. The amount of momentum follows from two results already established.

::: steps Wavelength and momentum of a photon
$ E = pc
Special relativity relates the energy E, momentum p and mass m of anything by E² = (pc)² + (mc²)². For a photon the mass is zero, so the second term vanishes, and taking the square root leaves energy equal to momentum times the speed of light. Unlike Newton's physics, relativity allows a massless thing to carry momentum.
---
$ E = hf
From Planck and Einstein, a photon of frequency f carries energy hf, where h is Planck's constant. This is the particle side of light described in terms of a wave property, the frequency.
---
$ p = \frac{hf}{c}
Both equations give the photon's energy, so the two right-hand sides must be equal: pc = hf. Dividing both sides by c gives the momentum in terms of the frequency.
---
$ p = \frac{h}{\lambda}
For any wave, the speed equals the frequency times the wavelength, so for light c = fλ, which means f/c = 1/λ. Replacing f/c in the previous line gives a strikingly simple result: the photon's momentum is Planck's constant divided by its wavelength. A particle property, momentum, is fixed entirely by a wave property, wavelength.
:::

This was not speculation. In 1923 Compton had scattered X-rays off electrons and found that the X-rays came out with longer wavelengths, having lost exactly the energy and momentum that a collision between particles would require, with each X-ray photon carrying momentum h/λ. De Broglie's step was to read that last equation backwards and apply it to matter.

## What is the wavelength of a particle?

If a photon of wavelength λ has momentum h/λ, then perhaps a particle with momentum p has a wavelength h/p. For a particle of mass m moving at speed v much less than c, the momentum is mv, so:

::: equation The de Broglie wavelength
\lambda = \frac{h}{p} = \frac{h}{mv}
:::

The formula says that the larger a particle's momentum, the shorter its wavelength. Because h is so tiny, about 6.63 × 10⁻³⁴ J·s, the wavelength of anything you could hold in your hand is absurdly small. A baseball of mass 0.145 kg thrown at 40 m/s has momentum about 5.8 kg·m/s, and a de Broglie wavelength of about 1.1 × 10⁻³⁴ m. That is about ten billion billion times smaller than a proton. Wave effects such as diffraction show up only when a wave passes through an opening or around an obstacle not much larger than its wavelength, and nothing in the universe is small enough to diffract a baseball. That is why ordinary objects show no trace of wave behaviour, and why nobody had suspected that matter had a wavelength.

Electrons are different, because their mass is so small, about 9.11 × 10⁻³¹ kg. To see how their wavelength compares with the sizes of things they might pass through, take the usual way of giving an electron a definite speed: let it start at rest and accelerate it through a voltage V. It gains kinetic energy eV, where e is its charge.

::: steps Wavelength of an electron accelerated through V volts
$ \frac{p^2}{2m} = eV
The kinetic energy ½mv² can be written in terms of momentum, since p = mv gives ½mv² = p²/2m. Setting that equal to the energy gained from the voltage connects the momentum we need to a voltage we control. This non-relativistic form is accurate as long as eV is much less than the electron's rest energy of about 511,000 eV, which is true for the voltages used here.
---
$ p = \sqrt{2meV}
Multiply both sides by 2m and take the square root to get the electron's momentum. Doubling the voltage does not double the momentum; it raises it by a factor of √2.
---
$ \lambda = \frac{h}{\sqrt{2meV}}
Put this momentum into de Broglie's relation λ = h/p. The electron's wavelength now depends only on constants of nature and the accelerating voltage, and it shrinks as one over the square root of the voltage.
---
$ \lambda = \frac{6.63 \times 10^{-34}}{\sqrt{2 \times 9.11 \times 10^{-31} \times 1.60 \times 10^{-19} \times 54}} \approx 1.67 \times 10^{-10}\,\text{m}
Put in numbers for a voltage of 54 V, the value used in the experiment described below, with all quantities in SI units. The wavelength comes out at about 0.17 nm, and the important thing about that number is its size.
:::

A wavelength of 0.17 nm is about the same as the spacing between neighbouring atoms in a solid crystal. A crystal's atoms sit in regular rows, and a set of evenly spaced obstacles acts on a wave the way a diffraction grating, a plate ruled with many fine parallel lines, acts on light: the waves scattered from different rows reinforce one another only in particular directions, where their path differences are whole numbers of wavelengths, and cancel elsewhere. So if electrons really are waves, a beam of electrons bounced off a crystal should come off not in all directions but in sharp beams at particular angles. That was a clear, testable prediction.

::: model wave
Change the wavelength and frequency of a travelling wave. In de Broglie's relation, a particle with more momentum has a shorter wavelength, just as a photon of higher frequency does.
:::

## Did it explain anything already known?

It did, and immediately. Bohr's model of hydrogen, eleven years earlier, had rested on a rule with no justification: the electron's angular momentum, mvr, must be a whole-number multiple of ħ, Planck's constant divided by 2π. De Broglie offered a reason.

Picture the electron's wave wrapped around a circular orbit. As the wave travels around and comes back to its starting point, it must join smoothly onto itself, crest meeting crest. If it did not, then on each lap it would arrive slightly out of step with itself, and after many laps the overlapping waves would cancel out. So the orbit's circumference must hold a whole number of wavelengths:

::: equation Standing waves on a circle
2\pi r = n\lambda = n\,\frac{h}{mv} \;\Rightarrow\; mvr = n\,\frac{h}{2\pi} = n\hbar
:::

Multiplying both sides by mv/2π turns the condition on wavelengths into a condition on angular momentum, and it is exactly Bohr's rule. What had been a bare assumption became a consequence of the electron being a wave that has to fit into a closed space. The same thing happens on a guitar string, which can vibrate only at frequencies whose waves fit a whole number of half-wavelengths between its fixed ends; that is why it plays definite notes. The comparison breaks down in one important respect: the real electron wave in an atom is not a wave running around a circle but a three-dimensional pattern filling the space around the nucleus, as Erwin Schrödinger's wave equation showed two years later. But the central lesson survived. Quantization, the restriction of a system to definite values, comes from waves being confined.

::: note
Every particle has a wavelength λ = h/p, set by its momentum. Wave behaviour shows up only when that wavelength is comparable to the size of whatever the particle meets, which happens for electrons passing through crystals but never for baseballs passing through doorways.
:::

::: history 1925–1927 · New York and Aberdeen
At the Bell Telephone Laboratories in New York, Clinton Davisson and his assistant Lester Germer were studying how electrons bounce off metal surfaces, a question that mattered for the design of vacuum tubes. Their target was a block of nickel in an evacuated glass bulb. In 1925 a bottle of liquid air used in the apparatus cracked, air rushed into the bulb, and the hot nickel oxidized. To clean it they heated it strongly for a long time, which, without their knowing it, turned the nickel from a mass of tiny randomly oriented crystals into a few large ones. When they resumed, the electrons no longer scattered smoothly in all directions but came off in sharp peaks at particular angles. Davisson did not at first know what to make of this. In the summer of 1926 he attended a meeting of the British Association in Oxford and, to his surprise, heard Max Born present some of Davisson's own earlier data as possible evidence for de Broglie's electron waves, an interpretation first suggested by the young physicist Walter Elsasser. Back in New York, Davisson and Germer searched deliberately, and in 1927 found a strong peak at an angle of 50° for electrons of 54 eV. The spacing between rows of atoms on the nickel surface was known from X-ray measurements to be 0.215 nm, and the grating condition gave a wavelength of 0.215 nm × sin 50° ≈ 0.165 nm, against de Broglie's prediction of 0.167 nm. In the same year George Paget Thomson in Aberdeen fired faster electrons through very thin metal films and photographed rings of diffraction, just like those X-rays make. Davisson and Thomson shared the Nobel Prize in 1937; de Broglie had received his in 1929. There was a family irony in Thomson's prize. His father, J. J. Thomson, had won his in 1906 for showing that the electron is a particle; the son won his for showing that it is a wave.
:::

## A wave of what?

Electron waves are now used every day. An electron microscope uses electrons with wavelengths tens of thousands of times shorter than those of visible light, and so avoids the diffraction limit that stops optical microscopes from resolving anything smaller than a few hundred nanometres; the best electron microscopes show individual atoms. And the wave behaviour is not special to electrons. Neutrons, whole helium atoms, and in 1999 even molecules of C₆₀, footballs of sixty carbon atoms, have been made to produce interference patterns.

The most revealing experiment was done in 1989 by Akira Tonomura's group at Hitachi in Japan. They sent electrons through a two-path arrangement equivalent to Young's double slit, so slowly that only one electron was in the apparatus at a time, and recorded where each arrived. Each electron landed at a single point on the detector, as a particle would. But as thousands of them accumulated, the points built up a pattern of bright and dark fringes, as a wave would, with almost no electrons landing in the dark bands.

That experiment sharpens the question de Broglie left unanswered: if the electron is a wave, what is waving? It is not a wave of electron substance spread out in space, because each electron always arrives whole, at one point. The answer that physicists came to accept was given by Max Born in 1926. The wave describes probability: where it is large, the electron is likely to be found, and where it cancels to zero, the electron is never found. That interpretation is the starting point of quantum mechanics, the subject of the next chapter.

::: quiz
? A proton and an electron move at the same speed. Which has the longer de Broglie wavelength?
+ The electron, because its smaller mass gives it less momentum
- The proton, because it is heavier
- Neither, because equal speeds give equal wavelengths
= The wavelength is h/mv, so at equal speeds it depends inversely on mass. The electron's momentum is about 1800 times smaller than the proton's, so its wavelength is about 1800 times longer.
? The voltage used to accelerate an electron is increased fourfold. What happens to its wavelength?
- It doubles
+ It halves
- It falls to a quarter
= The wavelength is h/√(2meV), which is proportional to 1/√V. Four times the voltage gives twice the momentum, and so half the wavelength.
? Why do we never see a thrown ball diffract as it passes through a doorway?
- Large objects have no wavelength
+ Its wavelength, around 10⁻³⁴ m, is far too small compared with the doorway
- Air molecules destroy the wave
= Diffraction is noticeable only when the wavelength is comparable to the size of the opening. A ball's wavelength is some 34 powers of ten smaller than a door, so its spreading is far too small ever to detect.
:::
