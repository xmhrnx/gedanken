Niels Bohr's model of 1913 put hydrogen's electron on circular orbits of fixed radius, like a tiny planet. It gave exactly the right energy levels, and that success made it famous, but almost everything else about it was wrong: the shapes, the angular momentum, the very idea of a path. Schrödinger's equation replaced Bohr's orbits with something stranger and more accurate, three-dimensional standing waves called **orbitals**, each labelled by a short list of whole numbers. This lesson explains what an orbital is, where its labels come from, what the orbitals look like, and how counting them leads directly to the shape of the periodic table.

## From orbits to orbitals

An **orbital** is a wavefunction for a single electron in an atom. Following the Born rule, the square of its size, |ψ|², at each point gives the probability density of finding the electron there. That is the whole of what an orbital tells you, and it is important to see what it does not say. It is not a path. The electron does not travel around the orbital the way a planet travels around the Sun; quantum mechanics gives it no trajectory at all. The orbital is better pictured as a cloud of probability, dense where the electron is likely to be found and thin where it rarely is.

For hydrogen, the electron is attracted to the single proton of the nucleus with a potential energy V = −ke²/r, where k is Coulomb's constant, e the elementary charge and r the distance from the nucleus. Solving Schrödinger's equation for this potential is longer than for the particle in a box, but the logic is the same. Only certain standing waves are allowed, because the wavefunction must behave sensibly in every direction: it must not blow up at the nucleus, it must fade to zero far away so the electron is somewhere, and it must join smoothly onto itself when you go all the way around the nucleus. In one dimension the box gave one such condition and one whole number, n. In three dimensions there are three independent conditions, one for the distance from the nucleus and two for direction, so each solution carries three whole numbers.

## The quantum numbers

The first is the **principal quantum number** n, which can be 1, 2, 3 and so on. It sets the energy, through E_n = −13.6 eV/n², which is exactly the formula Bohr had found. (An electronvolt, eV, is the energy an electron gains crossing one volt; the minus sign means the electron is bound, and 13.6 eV is the energy needed to free it from the lowest level.) It also sets the orbital's rough size: larger n means a cloud that extends further from the nucleus.

The second is the **orbital angular momentum quantum number** l, which for a given n can take the values 0, 1, and so on up to n − 1. It fixes how much angular momentum the electron has from its motion around the nucleus: the size of that angular momentum is √(l(l + 1)) ħ, where ħ is the reduced Planck constant. Angular momentum shapes the cloud. An electron with none spreads evenly in every direction, while more angular momentum pushes the probability away from the axis and into lobes.

The third is the **magnetic quantum number** m_l, which for a given l takes the whole-number values from −l to +l. Angular momentum has a direction as well as a size, and m_l fixes its component along a chosen axis, L_z = m_l ħ. In effect, it sets the orientation of the orbital. The name comes from the fact that in a magnetic field, states with different m_l have slightly different energies, which splits spectral lines.

A fourth number does not come from Schrödinger's equation at all. The **spin quantum number** m_s is +½ or −½, for the electron's spin up or down along the chosen axis. It is a property of the electron itself, not of the orbital, so each orbital can be occupied by an electron with either spin.

For historical reasons orbitals with l = 0, 1, 2 and 3 are called s, p, d and f. The letters are left over from nineteenth-century spectroscopy, where series of spectral lines were described as sharp, principal, diffuse and fundamental. An orbital is named by its n followed by its letter, so 3d means n = 3 and l = 2.

::: history 1913–1926 · Copenhagen, Munich, Leiden, Zürich
The numbers arrived one at a time, each forced on physicists by an experiment the previous set could not explain. Bohr's n, in 1913, accounted for hydrogen's main spectral lines. Closer inspection showed that many of those lines were actually clusters of lines very close together, and in 1916 Arnold Sommerfeld in Munich explained this fine structure by allowing elliptical orbits as well as circles, which introduced a second number related to today's l. The splitting of lines in a magnetic field demanded a third, for orientation, the ancestor of m_l. Even then, spectra did not quite fit: lines that should have been single came in pairs, and magnetic splittings had the wrong patterns. The missing piece, spin, was proposed in Leiden in 1925 by George Uhlenbeck and Samuel Goudsmit to explain exactly these anomalies, and it also accounted, in hindsight, for the two beams Otto Stern and Walther Gerlach had seen in 1922. Then, in 1926, Schrödinger solved his equation for hydrogen in Zürich, and n, l and m_l all appeared together from a single calculation, no longer as rules imposed one by one but as the natural labels of standing waves around a nucleus.
:::

## The shapes of orbitals

An s orbital, with l = 0, is spherically symmetric: the probability of finding the electron depends only on the distance from the nucleus, not the direction. An s electron has zero orbital angular momentum. This flatly contradicts Bohr, whose electron in the lowest orbit circled the nucleus with angular momentum ħ. Experiment sides with Schrödinger. For the 1s orbital, the lowest of all, the single most likely distance from the nucleus is 0.053 nm, which is exactly the radius of Bohr's first orbit, called the **Bohr radius**. But the cloud has no edge. The probability thins out smoothly with distance and never quite reaches zero, so the electron is occasionally found much further out, and occasionally right at the nucleus.

A p orbital, with l = 1, has two lobes on opposite sides of the nucleus, with zero probability on the flat plane between them. There are three p orbitals for each n, one for each of the values m_l = −1, 0 and +1. They are usually drawn as three identical dumbbells pointing along the x, y and z axes, a convenient set formed by combining the m_l states. The five d orbitals, with l = 2, are more elaborate, and most of them have four lobes.

The surfaces where the probability is exactly zero are called **nodes**, and they follow a pattern set by the quantum numbers, just as the nodes of the particle in a box did. An orbital has n − 1 nodes in total. Of these, l are planes or cones passing through the nucleus, which is why an s orbital has none of that kind and a p orbital has one plane. The remaining n − l − 1 are spheres centred on the nucleus. So 2s has one spherical node, a shell at a certain radius where the electron is never found, while 2p has none of those but one flat plane instead.

::: model orbital
Real hydrogen orbitals, with each dot placed at random according to |ψ|², so the density of dots shows where the electron is likely to be. Pick 2s and look for the spherical node, a hollow shell inside the cloud; then pick 3p and find both kinds of node. Colour shows the sign of ψ, which is opposite on either side of a node.
:::

::: model wavefunction
A one-dimensional cousin of an orbital: the particle in a box. Level n has n − 1 nodes, places where the electron is never found. In an atom the same rule holds, but the nodes become spheres, planes and cones.
:::

## Counting the states

How many different states share the same value of n? In hydrogen they all have the same energy, so this is the number of different ways the electron can have the energy E_n. The answer comes from counting the allowed combinations of quantum numbers.

::: steps Number of states with principal quantum number n
$ l = 0, 1, \ldots, n - 1
For a given n, the angular momentum number l can take any whole value from 0 up to n − 1, which is n choices in all. A larger orbit, in this sense, has room for more angular momentum. For n = 1 only l = 0 is allowed; for n = 2, both 0 and 1.
---
$ m_l = -l, \ldots, +l \quad\Rightarrow\quad 2l + 1 \text{ values}
For each l, the magnetic number m_l runs through the whole numbers from −l to +l. That is l negative values, l positive values and zero itself, giving 2l + 1 orientations. So each s subshell has one orbital, each p subshell three, and each d subshell five.
---
$ \sum_{l=0}^{n-1}(2l + 1) = 1 + 3 + 5 + \cdots + (2n - 1) = n^2
Add up the orbitals over all allowed l. The terms are the odd numbers from 1 up to 2n − 1, and the sum of the first n odd numbers is always n², as you can check: 1, 1 + 3 = 4, 1 + 3 + 5 = 9. For n = 2 this gives four orbitals, one 2s and three 2p.
---
$ 2n^2
Finally, each orbital can hold an electron with spin up or spin down, two states for each orbital. So the total number of states with principal quantum number n is twice n²: 2, 8, 18 and 32 for n = 1, 2, 3 and 4.
:::

These numbers, 2, 8, 18 and 32, will reappear in the next lesson as the lengths of the rows of the periodic table. That is not a coincidence: the table's structure is the structure of these standing waves.

There is one more subtlety. In hydrogen, all the orbitals with the same n have exactly the same energy, a special property of the pure 1/r attraction of a single proton. In atoms with more than one electron, this no longer holds. The inner electrons partly shield the outer ones from the nucleus, a process called **screening**, so an outer electron feels a weaker attraction than the full nuclear charge would give. Orbitals with low l, especially s orbitals, have some probability very close to the nucleus, inside the screening electrons, where they feel the full attraction, so they are bound more tightly. That is why, in every atom heavier than hydrogen, 2s lies below 2p, and 3s below 3p below 3d.

::: note
Electrons in atoms are standing waves, not orbiting particles. Each orbital is labelled by n, which sets its energy and size, l, which sets its shape, and m_l, which sets its orientation, and each can hold an electron of either spin. Counting the combinations shows that shell n holds 2n² states.
:::

::: quiz
? Which set of quantum numbers is not allowed?
- n = 2, l = 1, m_l = −1
- n = 3, l = 2, m_l = 2
+ n = 2, l = 2, m_l = 0
= The angular momentum number l can be at most n − 1, so for n = 2 only l = 0 or l = 1 is allowed, and l = 2 is impossible.
? How many orbitals are there with n = 3, ignoring spin?
- 3
- 6
+ 9
= There is one 3s orbital, three 3p orbitals and five 3d orbitals, and 1 + 3 + 5 = 9, which is 3², as the general rule says.
? What is wrong with Bohr's picture of the hydrogen ground state?
- Its energy is wrong
+ It gives the electron angular momentum ħ, but the 1s state has zero orbital angular momentum
- It puts the electron too far from the nucleus
= Bohr's energy for the ground state is correct, but the 1s orbital has l = 0, so the electron has no orbital angular momentum and follows no orbit at all.
:::
