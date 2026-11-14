A light bulb's atoms emit independently, at random moments, in random directions, with a spread of wavelengths. A laser's light is the opposite: one wavelength, one direction, all the waves in step. The difference comes from a process Einstein identified in 1917, decades before anyone could use it.

## Three ways to meet a photon

Take an atom with two energy levels, a lower level 1 and an upper level 2, separated by ΔE. A photon whose energy hf matches ΔE can interact with it in three ways.

- **Absorption**: an atom in level 1 takes in the photon and jumps to level 2.
- **Spontaneous emission**: an atom in level 2 drops to level 1 on its own after a short, random time, emitting a photon in a random direction. This is how ordinary light sources shine.
- **Stimulated emission**: a passing photon of the right energy makes an atom in level 2 drop to level 1 and emit a second photon. The new photon is a copy of the first: the same wavelength, direction, phase and polarisation.

Stimulated emission is the key. One photon in, two identical photons out. Those two can stimulate two more, and so on: light amplifies itself.

::: history 1917 · Berlin
Einstein was looking for a simple derivation of Planck's formula for thermal radiation. He found that atoms in equilibrium with radiation could only reproduce it if, alongside absorption and spontaneous emission, there was a third process: emission triggered by radiation already present. He also showed that, for each atom, the rates of absorption and stimulated emission must be governed by the same coefficient. Nobody saw a practical use for decades.
:::

## Why amplification is hard

Whether a beam grows or fades as it passes through a material depends on a competition. The same photons cause absorption in atoms in level 1 and stimulated emission in atoms in level 2.

::: steps The condition for amplification
$ R_{\text{abs}} = B\,\rho\,N_1
The rate of absorption is proportional to the number of atoms in the lower level, N₁, and to the density of light at the right frequency, ρ. B measures how strongly the atom couples to light.
---
$ R_{\text{stim}} = B\,\rho\,N_2
Stimulated emission has the same form with N₂, the number in the upper level. Einstein showed the coefficient B is the same for both, for levels of equal degeneracy (the number of states sharing that energy).
---
$ R_{\text{stim}} - R_{\text{abs}} = B\,\rho\,(N_2 - N_1)
Each stimulated emission adds a photon to the beam; each absorption removes one. The beam grows only if N₂ > N₁.
---
$ \frac{N_2}{N_1} = e^{-\Delta E/kT}
In thermal equilibrium at temperature T, atoms follow Boltzmann's law: the upper level is always less populated than the lower. Here k is Boltzmann's constant.
:::

For visible light, ΔE is about 2 eV. At room temperature kT ≈ 0.026 eV, so N₂/N₁ ≈ e^{−77}, about 10⁻³⁴. Essentially every atom is in the lower level, and light is absorbed, not amplified. Heating does not help: no temperature makes N₂ larger than N₁.

Amplification requires a **population inversion**: more atoms in the upper level than the lower, a state that never occurs in equilibrium and must be maintained by continually pumping energy in.

## Making an inversion

With only two levels, pumping with light cannot produce an inversion. As N₂ approaches N₁, the pump light stimulates emission as often as it is absorbed, and the best you can reach is equal populations. Practical lasers use three or four levels.

In a three-level laser, the pump lifts atoms from the ground state to a high level. They drop quickly to an intermediate **metastable** level, one where atoms linger for a relatively long time, often a millisecond rather than the usual ten nanoseconds or so. Atoms pile up there faster than they leave, and an inversion builds between the metastable level and the ground state. In a four-level laser the lasing transition ends on a level above the ground state that empties quickly, so even a few excited atoms make an inversion.

The material sits between two mirrors, one of them letting through a percent or so of the light. Photons travelling along the axis bounce back and forth, stimulating more emission on every pass, while those going sideways escape. The light that leaks through the partial mirror is the beam.

::: history 1954–1960 · New York and Malibu
Charles Townes and his students at Columbia built the first such amplifier in 1954, the **maser**, using ammonia molecules and microwaves; Nikolai Basov and Alexander Prokhorov in Moscow developed the idea independently. In 1958 Townes and Arthur Schawlow showed how to do the same with visible light using two mirrors. On 16 May 1960 Theodore Maiman at Hughes Research Laboratories got the first laser working: a ruby crystal, in which chromium ions form a three-level system, pumped by a photographer's flash lamp, giving red pulses at 694 nm. The name, coined by Gordon Gould, stands for Light Amplification by Stimulated Emission of Radiation. Townes, Basov and Prokhorov shared the 1964 Nobel Prize.
:::

::: model bohr
Choose a jump between levels. In a laser, a photon at exactly that wavelength passing an excited atom makes it emit an identical twin.
:::

## Uses

The properties follow from stimulated emission: every photon copies the one that triggered it, so the light is a single colour and coherent (the waves stay in step), and the cavity makes it a narrow beam. Lasers carry data through optical fibres, read barcodes, cut steel, reshape corneas, cool atoms to within millionths of a degree of absolute zero, and measure the length of LIGO's 4 km arms to a fraction of a proton's width.

::: note
Stimulated emission copies photons, so light passing through excited atoms can grow. It grows only if more atoms are in the upper level than the lower, a population inversion that never happens in thermal equilibrium and must be maintained by pumping.
:::

::: quiz
? Why can't you make a laser simply by heating a gas?
- Hot atoms move too fast
+ In equilibrium the upper level always has fewer atoms than the lower, so absorption wins
- Heating destroys the energy levels
= N₂/N₁ = e^{−ΔE/kT} is less than 1 at every temperature.
? What is special about the photon produced by stimulated emission?
- It has twice the energy of the incoming photon
+ It matches the incoming photon in wavelength, direction and phase
- It travels in a random direction
= Stimulated emission produces an exact copy, which is why laser light is coherent.
? Why do practical lasers use a metastable level?
- It emits light without being pumped
+ Atoms linger there long enough to build up a population inversion
- It absorbs the pump light more strongly
= Atoms arrive faster than they leave, so the upper lasing level fills beyond the lower one.
:::
