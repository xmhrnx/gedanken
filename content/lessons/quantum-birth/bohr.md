Heat hydrogen gas, or pass a spark through it, and it glows. Spread the glow with a prism and you do not see a rainbow but a few sharp coloured lines: red at 656 nm, blue-green at 486 nm, violet at 434 and 410 nm. Every hydrogen atom in the universe emits exactly these wavelengths. In 1913 Niels Bohr explained them with a model that was wrong in its picture and right in its numbers, and that pointed the way to quantum mechanics.

## A formula without a reason

In 1885 Johann Balmer, a Swiss schoolteacher, found by trial that the visible hydrogen wavelengths fit a simple formula. Johannes Rydberg generalized it:

::: equation The Rydberg formula
\frac{1}{\lambda} = R\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right), \quad R \approx 1.097 \times 10^7\,\text{m}^{-1}
:::

Here n₁ and n₂ are whole numbers with n₂ > n₁. Balmer's visible lines have n₁ = 2; ultraviolet and infrared series with n₁ = 1 and 3 were found later. Nobody knew why whole numbers should appear.

## A stable atom is impossible classically

In 1911 Ernest Rutherford showed that an atom has a tiny, heavy, positive nucleus with electrons around it. Classically the obvious model is a miniature solar system, with the electron orbiting the nucleus. But an orbiting electron is accelerating, and Maxwell's theory says an accelerating charge radiates. The electron would lose energy and spiral into the nucleus in about 10⁻¹¹ s, emitting a continuous smear of frequencies on the way. Atoms would not exist, and certainly would not emit sharp lines.

## Bohr's rules

Bohr, who had spent 1912 working with Rutherford in Manchester, simply declared new rules:

1. The electron can move only on certain **stationary orbits**, on which it does not radiate.
2. The allowed orbits are those whose angular momentum is a whole multiple of ħ = h/2π: mvr = nħ.
3. Light is emitted or absorbed only when the electron jumps between orbits, as a single photon carrying the energy difference: hf = E_initial − E_final.

Nothing justified rule 2 except that it worked. Here is what it gives.

::: steps The energy levels of hydrogen
$ \frac{ke^2}{r^2} = \frac{mv^2}{r}
The electron, of mass m and charge −e, circles a proton of charge +e at radius r. The Coulomb attraction, with k = 1/(4πε₀) ≈ 8.99 × 10⁹ N·m²/C², supplies the centripetal force.
---
$ mv^2 = \frac{ke^2}{r}
Multiply both sides by r.
---
$ K = \tfrac{1}{2}mv^2 = \frac{ke^2}{2r}
Halve it to get the kinetic energy.
---
$ U = -\frac{ke^2}{r}
The electrical potential energy of the pair, zero when they are infinitely far apart.
---
$ E = K + U = -\frac{ke^2}{2r}
Add them. The total is negative: the electron is bound, and it would take energy ke²/2r to pull it free.
---
$ v = \frac{n\hbar}{mr}
Bohr's quantization rule mvr = nħ, solved for v.
---
$ m\left(\frac{n\hbar}{mr}\right)^2 = \frac{ke^2}{r} \;\Rightarrow\; \frac{n^2\hbar^2}{mr^2} = \frac{ke^2}{r}
Substitute this v into mv² = ke²/r from the second step, and simplify the left side.
---
$ r_n = n^2\,\frac{\hbar^2}{mke^2} = n^2 a_0
Solve for r. Only these radii are allowed; they grow as n².
---
$ a_0 = \frac{\hbar^2}{mke^2} \approx 5.29 \times 10^{-11}\,\text{m}
The smallest radius, called the Bohr radius, built entirely from known constants. It matches the measured size of atoms, about 0.1 nm across, which no classical theory could explain.
---
$ E_n = -\frac{ke^2}{2a_0}\,\frac{1}{n^2} = -\frac{mk^2e^4}{2\hbar^2}\,\frac{1}{n^2}
Put r_n into E = −ke²/2r.
---
$ E_n = -\frac{13.6\,\text{eV}}{n^2}
Evaluate: ke² ≈ 2.31 × 10⁻²⁸ J·m, divided by 2a₀, gives 2.18 × 10⁻¹⁸ J, which is 13.6 eV.
:::

So hydrogen's levels are −13.6, −3.4, −1.51, −0.85 eV, and so on, crowding together toward 0. The ground state, n = 1, takes 13.6 eV to ionize, exactly the measured ionization energy of hydrogen.

## The lines explained

An electron dropping from level n₂ to n₁ emits one photon:

::: equation Photon from a jump
hf = \frac{hc}{\lambda} = 13.6\,\text{eV}\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right)
:::

This is Rydberg's formula, with R = 13.6 eV / hc ≈ 1.097 × 10⁷ m⁻¹: Bohr derived the empirical constant from m, e, h and ε₀. The jump from 3 to 2 releases 13.6 × (1/4 − 1/9) = 1.89 eV, a wavelength of 656 nm, the red line. Jumps ending on n = 1 give the ultraviolet Lyman series, jumps ending on n = 3 the infrared Paschen series.

::: model bohr
Pick a jump between levels and see the photon's wavelength and colour. Jumps down to n = 2 give the visible lines.
:::

::: note
An atom can hold only certain energies. Light comes out when the electron drops between them, so each photon's energy, and hence its colour, is fixed by the gap. That is why spectra have lines.
:::

::: history 1913 · Manchester and Copenhagen
Bohr, 27, was thinking about atomic structure and stability when a colleague in Copenhagen, Hans Hansen, asked him how his ideas dealt with the Balmer formula. Bohr had not known it. "As soon as I saw Balmer's formula," he said later, "the whole thing was immediately clear to me." He published in July 1913. Many found the rules arbitrary, but he predicted that certain lines then attributed to hydrogen came from ionized helium, and when Alfred Fowler objected that the numbers were slightly off, Bohr showed that allowing for the motion of the nucleus removed the discrepancy almost exactly. In 1914 James Franck and Gustav Hertz showed directly that mercury atoms absorb energy from electrons only in a fixed chunk of 4.9 eV.
:::

## What is wrong with it

The orbit picture is wrong. Electrons do not travel on circular tracks; in quantum mechanics an electron in an atom has no definite path, only a probability spread around the nucleus. In the ground state its angular momentum is actually zero, not ħ. The model fails for every atom with more than one electron, even helium, and cannot say how bright each line should be. It gets hydrogen's energies right because of a special property of the 1/r Coulomb force, not because the picture is right.

What survived is the core idea: atoms have discrete energy levels, and photons carry away the differences. Why angular momentum should be quantized remained a mystery until Louis de Broglie suggested that the electron is a wave, which the next lesson explains.

::: quiz
? Why did classical physics predict that Rutherford's atom should collapse?
- The nucleus would repel the electron
+ An orbiting electron accelerates, so it should radiate energy and spiral inward
- The electron's orbit would be unstable to collisions
= Maxwell's theory says accelerating charges radiate. Losing energy continuously, the electron would fall into the nucleus in about 10⁻¹¹ s.
? How much energy is released when a hydrogen electron drops from n = 2 to n = 1?
- 3.4 eV
+ 10.2 eV
- 13.6 eV
= E₂ − E₁ = −3.4 eV − (−13.6 eV) = 10.2 eV, an ultraviolet photon of 122 nm.
? How does the radius of a Bohr orbit depend on n?
- In proportion to n
+ In proportion to n²
- In proportion to 1/n²
= r_n = n²a₀. The energy goes as 1/n² because E = −ke²/2r.
:::
