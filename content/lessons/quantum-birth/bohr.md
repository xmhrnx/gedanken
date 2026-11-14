Heat hydrogen gas, or pass an electric spark through it, and it glows a pinkish purple. Spread that glow out with a prism and you do not see a continuous rainbow, as you would from a hot solid. You see a few sharp, separate lines of colour on a dark background: a red line at 656 nm, a blue-green one at 486 nm, and violet ones at 434 and 410 nm. Every hydrogen atom anywhere, in a laboratory tube or in a star billions of light years away, emits exactly these wavelengths and no others. For decades nobody could say why. In 1913 Niels Bohr explained them with a model of the atom that was wrong in its picture and right in its numbers, and that pointed the way to quantum mechanics. This lesson follows his reasoning, derives the energy levels of hydrogen, and then says plainly what was wrong with it and what survived.

## Was there a pattern in the lines?

There was, and it was found long before anyone could explain it. In 1885 Johann Balmer, a mathematics teacher at a girls' school in Basel, set himself the puzzle of finding a formula for the four visible hydrogen wavelengths, which had been measured precisely. By trial and error he found one that fitted all four to within the accuracy of the measurements, involving nothing but whole numbers and a single constant. A few years later the Swedish physicist Johannes Rydberg put it in a more general form:

::: equation The Rydberg formula
\frac{1}{\lambda} = R\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right), \quad R \approx 1.097 \times 10^7\,\text{m}^{-1}
:::

Here n₁ and n₂ are whole numbers with n₂ larger than n₁, and R is a constant fitted to the data, now called the Rydberg constant. Balmer's visible lines all have n₁ = 2, with n₂ = 3, 4, 5 and 6. The formula suggested other families of lines, and they were found: a series in the ultraviolet with n₁ = 1, discovered by Theodore Lyman, and one in the infrared with n₁ = 3, discovered by Friedrich Paschen. The formula worked perfectly. But it was a pattern without a reason. Nothing in physics said why whole numbers should appear in the light of an atom, still less why they should appear as one over a square.

## Why could a classical atom not exist?

In 1911 Ernest Rutherford, firing alpha particles at thin gold foil, showed that an atom has nearly all its mass and all its positive charge packed into a tiny central nucleus, with the electrons somewhere around it. The obvious picture is a miniature solar system: the electron orbits the nucleus, held by electric attraction as a planet is held by gravity.

Classical physics has a fatal objection to this picture. An electron moving in a circle is constantly changing direction, so it is accelerating, and Maxwell's theory says that an accelerating charge radiates electromagnetic waves, carrying energy away. The orbiting electron would therefore lose energy continuously, spiral inward, and crash into the nucleus. The calculation gives a time of about 10⁻¹¹ seconds. On the way in, its orbital frequency would rise smoothly, so it would emit a continuous smear of frequencies, not sharp lines. If classical physics were right, atoms would collapse almost instantly, and they would never produce a line spectrum. Yet atoms are stable for billions of years, and they produce nothing but lines.

## What rules did Bohr propose?

Bohr, a young Danish physicist, had spent several months of 1912 working with Rutherford in Manchester, and was convinced the nuclear atom was right. His response to the classical objection was not to explain it away but to overrule it. He proposed three new rules and asked only whether they gave the right answers.

The first rule says that the electron can move only on certain special orbits, called **stationary states**, and while it is on one of them it does not radiate, whatever Maxwell's theory says. The second rule picks out which orbits are allowed: those on which the electron's **angular momentum**, its mass times its speed times the radius of the orbit, mvr, is a whole-number multiple of a basic amount ħ, pronounced "h-bar", which is Planck's constant divided by 2π. So mvr = nħ, where n = 1, 2, 3 and so on. The third rule says that light is emitted or absorbed only when the electron jumps from one allowed orbit to another, and then as a single photon whose energy hf equals the difference between the energies of the two orbits.

Bohr had reasons for the form of the second rule, including the requirement that his atom should behave classically for very large orbits, but at bottom it was justified by one thing: it worked. Here is what it gives.

::: steps The energy levels of hydrogen
$ \frac{ke^2}{r^2} = \frac{mv^2}{r}
Take an electron, of mass m and charge −e, moving in a circle of radius r around a proton of charge +e. Something must pull it toward the centre to keep it on a circle, and the force needed for mass m at speed v on radius r is mv²/r. Here the electric attraction between the two charges supplies it. By Coulomb's law that attraction is ke²/r², where k = 1/(4πε₀) ≈ 8.99 × 10⁹ N·m²/C² is the constant of electric force.
---
$ mv^2 = \frac{ke^2}{r}
Multiply both sides by r. This tidier form will be used twice: once to find the energy and once to find the allowed radii.
---
$ K = \tfrac{1}{2}mv^2 = \frac{ke^2}{2r}
The electron's kinetic energy is half of mv². Halving both sides of the previous line expresses it in terms of the radius alone: a smaller orbit means a faster electron.
---
$ U = -\frac{ke^2}{r}
The electric potential energy of the electron and proton is this, negative because they attract. It is measured from zero when they are infinitely far apart; as they come closer the energy falls, which is what attraction means.
---
$ E = K + U = -\frac{ke^2}{2r}
The total energy is the sum. The potential energy is twice as large as the kinetic energy and opposite in sign, so the total is negative. A negative total means the electron is bound: to pull it free and leave it at rest far away, where the energy is zero, you would have to supply ke²/2r.
---
$ v = \frac{n\hbar}{mr}
Now bring in Bohr's second rule, mvr = nħ, and solve it for the speed. This is the only new ingredient; everything so far was classical.
---
$ m\left(\frac{n\hbar}{mr}\right)^2 = \frac{ke^2}{r} \;\Rightarrow\; \frac{n^2\hbar^2}{mr^2} = \frac{ke^2}{r}
Put this expression for v into mv² = ke²/r from the second step. Squaring the bracket and cancelling one factor of m gives the form on the right. The speed has now disappeared, leaving an equation in the radius alone.
---
$ r_n = n^2\,\frac{\hbar^2}{mke^2} = n^2 a_0
Multiply both sides by r² and divide by ke² to solve for r. Only these radii are allowed, one for each whole number n, and they grow as the square of n: the second orbit is four times as large as the first, the third nine times.
---
$ a_0 = \frac{\hbar^2}{mke^2} \approx 5.29 \times 10^{-11}\,\text{m}
The smallest radius, for n = 1, is called the Bohr radius, a₀. It is built entirely from known constants of nature, and it gives a hydrogen atom a diameter of about 0.1 nm, matching the measured size of atoms. No classical theory could predict the size of an atom at all, because classical physics has no constant from which to build a length on that scale; Planck's h supplied one.
---
$ E_n = -\frac{ke^2}{2a_0}\,\frac{1}{n^2} = -\frac{mk^2e^4}{2\hbar^2}\,\frac{1}{n^2}
Put the allowed radii r_n = n²a₀ into the energy E = −ke²/2r from the fifth step. Each allowed orbit therefore has a definite energy, falling off as one over n². Writing a₀ out in full shows that this energy, too, is built only from constants of nature.
---
$ E_n = -\frac{13.6\,\text{eV}}{n^2}
Evaluate the constant. With ke² ≈ 2.31 × 10⁻²⁸ J·m and 2a₀ ≈ 1.06 × 10⁻¹⁰ m, the ratio is about 2.18 × 10⁻¹⁸ J. Dividing by 1.60 × 10⁻¹⁹ J per electronvolt gives 13.6 eV.
:::

So hydrogen's allowed energies are −13.6 eV for n = 1, −3.4 eV for n = 2, −1.51 eV for n = 3, −0.85 eV for n = 4, and so on, crowding closer and closer together as they approach zero. The lowest level, n = 1, is called the **ground state**, and it is where a hydrogen atom normally sits. To remove the electron from it entirely takes 13.6 eV, and that is exactly the measured ionization energy of hydrogen, the energy needed to strip its electron away. This was the first time anyone had calculated a property of an atom from fundamental constants and got the right answer.

## Where do the lines come from?

Now apply the third rule. When the electron drops from a higher level n₂ to a lower level n₁, the atom emits one photon carrying away the difference in energy:

::: equation Photon from a jump
hf = \frac{hc}{\lambda} = 13.6\,\text{eV}\left(\frac{1}{n_1^2} - \frac{1}{n_2^2}\right)
:::

This is Rydberg's formula, with whole numbers that now have a meaning: they label the orbits the electron jumps between. Dividing 13.6 eV by hc gives R ≈ 1.097 × 10⁷ m⁻¹, so Bohr had derived the empirical constant from the mass and charge of the electron, Planck's constant, and the constant of electric force. A jump from n = 3 to n = 2 releases 13.6 × (1/4 − 1/9) ≈ 1.89 eV, which corresponds to a wavelength of 656 nm: the red line. The other visible lines are jumps from 4, 5 and 6 down to 2. Jumps ending on n = 1 involve larger energy gaps and give the ultraviolet Lyman series; jumps ending on n = 3 involve smaller gaps and give the infrared Paschen series. Absorption works the same way in reverse: hydrogen gas in front of a bright source absorbs exactly the photons whose energies match an upward jump, which is why the same wavelengths appear as dark lines in the spectra of stars.

::: model bohr
Pick a jump between two energy levels and see the wavelength and colour of the photon it emits. Jumps that end on n = 2 give the visible lines; those that end on n = 1 are all in the ultraviolet.
:::

::: note
An atom can hold only certain energies. Light is emitted when an electron drops from one allowed energy to a lower one, so each photon's energy, and therefore its colour, is fixed by the gap between two levels. That is why the light of an element comes in sharp lines rather than a continuous rainbow.
:::

::: history 1913–1914 · Copenhagen, Manchester and Berlin
Back in Copenhagen early in 1913, Bohr, then 27, was working on the stability of atoms and had not been thinking about spectra, which seemed too complicated to be useful. A colleague, the spectroscopist Hans Hansen, asked him how his ideas dealt with Balmer's formula. Bohr did not know it. When he looked it up, he said later, "the whole thing was immediately clear to me." He published his theory in three papers that year, the first in July. Many physicists found the rules arbitrary, and some found them outrageous. But Bohr made a striking prediction. Certain lines seen in the spectra of stars, and attributed to hydrogen, did not fit the Rydberg formula, and Bohr argued that they came instead from helium atoms that had lost one of their two electrons, whose single remaining electron orbits a nucleus with twice the charge. When the British spectroscopist Alfred Fowler objected that the predicted wavelengths were slightly off, Bohr showed that the discrepancy came from treating the nucleus as fixed; allowing for the fact that electron and nucleus both circle their common centre of mass removed it almost exactly. In 1914 James Franck and Gustav Hertz in Berlin fired electrons through mercury vapour and found that the atoms absorbed energy from them only in a fixed amount of 4.9 eV, direct evidence that atoms have discrete energy levels, though Franck and Hertz did not at first interpret it that way.
:::

## What is wrong with Bohr's model?

The picture of electrons on circular orbits is wrong. In quantum mechanics, which replaced Bohr's model in the mid-1920s, an electron in an atom has no definite path at all. It is described by a wavefunction, which says only how likely the electron is to be found at each place, and in the ground state of hydrogen that likelihood forms a spherical cloud around the nucleus, densest at the centre, not a ring. The ground state also has zero angular momentum, not ħ as Bohr's rule requires. The model fails outright for any atom with more than one electron, even helium, and it cannot say how bright each spectral line should be. It gets hydrogen's energies exactly right because of a special mathematical property of the Coulomb force, which falls off as one over the square of distance, not because its picture is right.

::: model orbital
Choose an orbital and see where the electron is likely to be found. These clouds, not circular orbits, are what quantum mechanics puts in place of Bohr's picture; the 1s cloud is hydrogen's ground state.
:::

What survived is the core idea: atoms have discrete energy levels, and photons carry away the differences between them. That idea is entirely correct, and it explains every line spectrum ever measured. But Bohr's second rule remained a mystery. Why should angular momentum come only in whole multiples of ħ? The first hint of an answer came in 1924, from Louis de Broglie's suggestion that the electron is itself a kind of wave, which the next lesson explains.

::: quiz
? Why did classical physics predict that Rutherford's nuclear atom should collapse?
- The nucleus would repel the electron
+ An orbiting electron is accelerating, so it should radiate energy and spiral inward
- The electron's orbit would be disrupted by collisions with other atoms
= Maxwell's theory says that any accelerating charge radiates, and an electron moving in a circle is accelerating. Losing energy continuously, it would spiral into the nucleus in about 10⁻¹¹ seconds.
? How much energy is released when a hydrogen electron drops from n = 2 to n = 1?
- 3.4 eV
+ 10.2 eV
- 13.6 eV
= The photon carries the difference between the two levels: E₂ − E₁ = −3.4 eV − (−13.6 eV) = 10.2 eV. That is an ultraviolet photon with a wavelength of about 122 nm.
? How does the radius of an allowed Bohr orbit depend on n?
- In proportion to n
+ In proportion to n²
- In proportion to 1/n²
= Combining the force balance with mvr = nħ gives r_n = n²a₀. The energy, by contrast, goes as 1/n², because E = −ke²/2r and r grows as n².
:::
