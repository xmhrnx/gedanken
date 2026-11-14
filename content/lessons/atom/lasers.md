The atoms in the filament of an old light bulb each emit light on their own schedule: at random moments, in random directions, over a broad spread of wavelengths, with no relation between one atom's light and the next. A laser's light is the opposite in every respect. It has a single, very precise wavelength, it travels in one narrow direction, and all its waves rise and fall in step with one another. The difference comes from a process Albert Einstein identified in 1917, in a paper about something else entirely, more than forty years before anyone built a device that used it. This lesson explains that process, why it is so hard to exploit, and the trick that makes it work.

## Three ways for an atom to meet light

Start with the simplest possible atom for the purpose: one with just two energy levels, a lower level 1 and an upper level 2, separated by an energy gap ΔE. As the lesson on spectra showed, such an atom interacts only with photons whose energy hf matches the gap, where h is Planck's constant and f the light's frequency. There are three ways that interaction can go.

The first is **absorption**. An atom in the lower level takes in a passing photon and jumps up to the upper level. The photon disappears, its energy stored in the atom.

The second is **spontaneous emission**. An atom in the upper level does not stay there. After a short time, which is random for each atom but averages to a definite lifetime, typically around ten nanoseconds, it drops back to the lower level on its own and emits a photon. Nothing decides the direction of that photon or the exact moment it leaves, so a crowd of atoms emitting this way produces light that goes everywhere, with no relation between one photon and the next. This is how flames, light bulbs and the Sun shine.

The third is the one that matters here, **stimulated emission**. A photon of the right energy passes an atom that is already in the upper level. The photon is not absorbed, since the atom has nowhere higher to go, but its presence triggers the atom to drop to the lower level at once and emit a second photon. The remarkable thing is that the new photon is an exact copy of the one that triggered it: the same wavelength, the same direction of travel, the same phase, meaning its wave crests line up with the original's, and the same polarisation, the direction in which its electric field oscillates.

Stimulated emission is the key to everything that follows. One photon goes in, and two identical photons come out. Each of those can trigger another atom, giving four, then eight, and so on. In the right conditions, light amplifies itself, and every photon in the growing beam is a copy of the first. The name **laser** is an acronym for exactly this: light amplification by stimulated emission of radiation.

::: history 1917 · Berlin
Einstein was not thinking about amplifying light. He wanted a simple, clean derivation of Max Planck's formula for the radiation inside a hot oven, which Planck had found in 1900 by a mixture of inspired guessing and difficult argument. Einstein imagined atoms sitting in such radiation, absorbing and emitting, until everything settled into equilibrium. He found that with only absorption and spontaneous emission the balance came out wrong: it did not reproduce Planck's formula. To make it work, he had to add a third process, emission triggered by the radiation already present. He also found that the coefficients for absorption and stimulated emission must be equal, apart from a factor for levels that contain different numbers of states. The idea was published, accepted, and largely left alone. It seemed a fine point of theory with no practical consequence.
:::

## Why amplification is hard

If stimulated emission makes light grow, why does light passing through ordinary matter not grow? Because the same photons that can cause stimulated emission can also be absorbed, and in ordinary matter absorption wins. Whether a beam grows or fades depends on a competition between the two, and the competition is decided by how many atoms are in each level.

::: steps The condition for amplification
$ R_{\text{abs}} = B\,\rho\,N_1
The rate of absorption is proportional to the number of atoms in the lower level, N₁, since only those atoms can absorb. It is also proportional to ρ, the density of light energy at the right frequency, since more photons mean more chances. The constant B, called an Einstein coefficient, measures how strongly the atom couples to light.
---
$ R_{\text{stim}} = B\,\rho\,N_2
Stimulated emission has exactly the same form, except that it depends on N₂, the number of atoms in the upper level, since only those can be stimulated. Einstein showed that the coefficient B is the same as for absorption, provided the two levels have equal **degeneracy**, the number of distinct quantum states that share each energy. So, atom for atom, the two processes are equally likely.
---
$ R_{\text{stim}} - R_{\text{abs}} = B\,\rho\,(N_2 - N_1)
Each stimulated emission adds a photon to the beam, and each absorption removes one, so the net rate at which photons are added is the difference between the two rates. Because B and ρ are common to both, they factor out, and the sign of the result depends only on N₂ − N₁. The beam grows if there are more atoms in the upper level than the lower, and fades otherwise.
---
$ \frac{N_2}{N_1} = e^{-\Delta E/kT}
Now ask how atoms are actually distributed between the levels in a material left alone at temperature T. In thermal equilibrium they follow Boltzmann's law, a result of statistical physics, where k is Boltzmann's constant and kT is a measure of the typical thermal energy available. Since ΔE and kT are both positive, the exponential is always less than one, so the upper level always holds fewer atoms than the lower.
:::

Put in numbers to see how lopsided this is. For visible light the gap ΔE is about 2 eV, where an electronvolt is the energy an electron gains crossing one volt. At room temperature kT is about 0.026 eV, so the ratio is e⁻⁷⁷, about 10⁻³⁴. For every atom in the upper level there are some 10³⁴ in the lower. Light passing through is absorbed, not amplified, which is why ordinary materials are dark or transparent, never glowing brighter than the light that enters them.

Heating does not help. Raising T makes the exponent smaller and N₂/N₁ larger, but the ratio approaches one only as the temperature goes to infinity, and it never exceeds one. No temperature puts more atoms in the upper level than the lower.

Amplification therefore requires a **population inversion**, more atoms in the upper level than in the lower one. That is a state that never occurs in thermal equilibrium. It has to be created and kept going by continually pumping energy into the material from outside, and the art of the laser lies in finding a way to do that.

## Making an inversion

The obvious approach is to shine intense light on the material at the frequency of the transition, lifting atoms from level 1 to level 2. With only two levels, this cannot work. As N₂ rises toward N₁, the pump light causes stimulated emission from the upper level almost as often as it causes absorption from the lower, because the two processes have the same coefficient. The best the pump can do is bring the two populations to equal, at which point the material becomes transparent and stops absorbing. It never gets past equal.

The way round is to use three or four levels. In a **three-level laser**, the pump lifts atoms from the ground state to a high level, well above the one that will do the lasing. From there they drop very quickly, without emitting useful light, into an intermediate level that is **metastable**, meaning atoms stay in it for an unusually long time, often around a millisecond rather than the usual ten nanoseconds or so. Because atoms arrive in the metastable level much faster than they leave it, they pile up there, and eventually it holds more atoms than the ground state. Now there is an inversion between the metastable level and the ground state, and stimulated emission between the two amplifies light. The pump light, at a different frequency, does not stimulate this transition, so it no longer works against itself. The drawback is that more than half of all the atoms must be lifted out of the ground state before any inversion appears, which takes a powerful pump.

A **four-level laser** avoids this. Its lasing transition ends not on the ground state but on a level slightly above it, which empties very quickly into the ground state. That lower lasing level is therefore almost always nearly empty, so even a small number of atoms in the upper level makes an inversion. Most lasers in use today work this way.

An inversion alone gives an amplifier. To make a laser, the material is placed between two mirrors facing each other, one reflecting nearly all the light and the other letting through about one percent. Photons that happen to travel along the axis between the mirrors bounce back and forth through the material hundreds of times, stimulating more emission on every pass, so the light along the axis builds up enormously. Photons travelling in any other direction leave through the sides after a single pass and are lost. The light that leaks through the partly transparent mirror is the laser beam.

::: history 1954–1960 · New York and Malibu
The first working amplifier of this kind used microwaves, not light. In 1954 Charles Townes and his students at Columbia University in New York built the **maser**, short for microwave amplification by stimulated emission of radiation, by separating ammonia molecules in an upper energy state from those in a lower one and sending the excited ones into a cavity. Nikolai Basov and Alexander Prokhorov in Moscow developed the same idea independently. Extending it to visible light, whose wavelength is tens of thousands of times shorter than the maser's, seemed very difficult, because no one could build a closed cavity that small. In 1958 Townes and Arthur Schawlow showed how to do without one, using just two parallel mirrors. Gordon Gould, a graduate student at Columbia, had reached similar ideas and coined the word "laser" in his notebooks. On 16 May 1960 Theodore Maiman at Hughes Research Laboratories in Malibu got the first laser working. It was a small rod of synthetic ruby, in which chromium ions form a three-level system, surrounded by a coiled flash lamp of the kind used in photography, and it gave pulses of deep red light at 694 nm. Townes, Basov and Prokhorov shared the 1964 Nobel Prize in Physics.
:::

::: model bohr
Choose a jump between two levels and note the wavelength of the photon it produces. In a laser, a photon of exactly that wavelength, passing an atom already in the upper level, makes it emit an identical copy.
:::

## What laser light is good for

Every special property of laser light follows from stimulated emission and the mirrors. Each photon copies the one that triggered it, so the light has a single, sharply defined wavelength and is **coherent**: its waves stay in step with each other over long distances. The mirrors select only light travelling along their axis, so the beam is narrow and spreads very little.

Those properties put lasers almost everywhere. Pulses of laser light carry data through optical fibres across oceans. Lasers read barcodes, cut and weld steel, reshape the cornea of the eye to correct vision, and, tuned carefully, slow and cool clouds of atoms to within millionths of a degree of absolute zero. And because coherent light from one laser can be split and recombined to produce interference fringes, it can measure distances with extraordinary precision: the LIGO detectors, which sense gravitational waves, use lasers to detect changes in the length of their 4 km arms smaller than a ten-thousandth of the width of a proton.

::: note
Stimulated emission copies photons, so light passing through excited atoms can grow. It grows only if more atoms are in the upper level than the lower, a population inversion that never occurs in thermal equilibrium and must be maintained by pumping energy in, usually through a third or fourth level.
:::

::: quiz
? Why can you not make a laser simply by heating a gas?
- Hot atoms move too fast to emit light
+ In equilibrium the upper level always has fewer atoms than the lower, so absorption wins
- Heating destroys the atoms' energy levels
= By Boltzmann's law N₂/N₁ = e^{−ΔE/kT}, which is less than one at every temperature, so there are always more atoms ready to absorb than to be stimulated.
? What is special about the photon produced by stimulated emission?
- It has twice the energy of the incoming photon
+ It matches the incoming photon in wavelength, direction and phase
- It travels in a random direction
= Stimulated emission produces an exact copy of the triggering photon, and a beam built from copies is coherent and travels in one direction.
? Why do practical lasers use a metastable level?
- It emits light without being pumped
+ Atoms linger there long enough to build up a population inversion
- It absorbs the pump light more strongly
= Because atoms arrive in the metastable level much faster than they leave it, they accumulate until the upper lasing level holds more atoms than the lower one.
:::
