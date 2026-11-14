Bohr's model put the electron of hydrogen on circular orbits of fixed radius. It got the energy levels right and almost everything else wrong. Schrödinger's equation replaces the orbits with standing waves in three dimensions, called **orbitals**, and labels each one with a short list of whole numbers.

## From orbits to orbitals

An **orbital** is a wavefunction for one electron in an atom. Following the Born rule, |ψ|² at each point gives the probability density of finding the electron there. It is not a path. The electron has no trajectory; the orbital is a cloud of probability, densest where the electron is most likely to be found.

For hydrogen the electron feels the nucleus's attraction, V = −ke²/r. Solving Schrödinger's equation for this potential is longer than for the box, but the logic is the same: only certain standing waves fit, and each is fixed by requiring ψ to behave sensibly at the nucleus, far away, and all the way around. In three dimensions there are three such conditions, so three whole numbers label each solution.

## The quantum numbers

- **n**, the principal quantum number: 1, 2, 3, … It sets the energy, E_n = −13.6 eV/n², exactly Bohr's formula, and roughly the size of the orbital.
- **l**, the orbital angular momentum number: 0, 1, …, n − 1. The electron's orbital angular momentum has size √(l(l + 1)) ħ. It sets the orbital's shape.
- **m_l**, the magnetic quantum number: whole numbers from −l to +l. It gives the component of angular momentum along a chosen axis, L_z = m_l ħ, and so the orbital's orientation.
- **m_s**, the spin quantum number: +½ or −½, for spin up or down. It was not in Schrödinger's original equation; it came from the Stern–Gerlach result.

Orbitals with l = 0, 1, 2, 3 are called s, p, d and f. The letters are left over from nineteenth-century spectroscopy, where lines were described as sharp, principal, diffuse and fundamental. So 3d means n = 3, l = 2.

::: history 1913–1926 · Copenhagen, Munich, Zürich
The numbers arrived one at a time, each forced by an experiment. Bohr's n (1913) explained hydrogen's main lines. Arnold Sommerfeld added elliptical orbits in 1916 to explain fine splittings, introducing a second number related to l. The splitting of lines in a magnetic field required a third, for orientation, now m_l. Spectra still did not fit until spin was added in 1925. Schrödinger's 1926 solution for hydrogen produced n, l and m_l together from a single equation.
:::

## Shapes

An s orbital (l = 0) is spherically symmetric: the probability depends only on distance from the nucleus. Its electron has zero orbital angular momentum, which flatly contradicts Bohr, whose ground-state electron circled with angular momentum ħ. For 1s the most likely distance from the nucleus is the Bohr radius, 0.053 nm, but the cloud has no edge.

A p orbital (l = 1) has two lobes on opposite sides of the nucleus, with zero probability on a plane between them. The three values m_l = −1, 0, +1 correspond to three orientations, often drawn along x, y and z. The five d orbitals mostly have four lobes.

The pattern of empty surfaces, called **nodes**, follows from the quantum numbers just as the nodes of the particle in a box did. An orbital has n − 1 nodes in total: l of them are planes or cones through the nucleus, and the remaining n − l − 1 are spheres around it.

::: model orbital
Real hydrogen orbitals, each dot sampled from |ψ|². Pick 2s and 3p to see the spherical and planar nodes; colour shows the sign of ψ.
:::

::: model wavefunction
A one-dimensional cousin of an orbital. Level n has n − 1 nodes, places where the electron is never found. In an atom, nodes become spheres and planes.
:::

## Counting the states

How many distinct states share each value of n? For hydrogen they all have the same energy, so this is the number of ways an electron can have energy E_n.

::: steps Number of states with principal quantum number n
$ l = 0, 1, \ldots, n - 1
For a given n, l runs from 0 up to n − 1.
---
$ m_l = -l, \ldots, +l \quad\Rightarrow\quad 2l + 1 \text{ values}
For each l there are l values below zero, l above, and zero itself.
---
$ \sum_{l=0}^{n-1}(2l + 1) = 1 + 3 + 5 + \cdots + (2n - 1) = n^2
The sum of the first n odd numbers is n². For n = 2: 1 + 3 = 4 orbitals, one 2s and three 2p.
---
$ 2n^2
Each orbital holds two spin states, m_s = ±½. So shell n has 2, 8, 18, 32 states for n = 1, 2, 3, 4.
:::

These numbers, 2, 8, 18, 32, will reappear as the lengths of the rows of the periodic table.

In hydrogen, all orbitals with the same n have the same energy, a special property of the 1/r potential. In atoms with more electrons, the inner electrons screen the nucleus, and orbitals that reach closer to it (lower l) are bound more tightly: 2s lies below 2p, 3s below 3p below 3d.

::: note
Electrons in atoms are standing waves, not orbiting particles. Each orbital is labelled by n (energy and size), l (shape), m_l (orientation), and the electron's spin m_s. Shell n holds 2n² states.
:::

::: quiz
? Which set of quantum numbers is not allowed?
- n = 2, l = 1, m_l = −1
- n = 3, l = 2, m_l = 2
+ n = 2, l = 2, m_l = 0
= l can be at most n − 1, so n = 2 allows only l = 0 or 1.
? How many orbitals (ignoring spin) are there with n = 3?
- 3
- 6
+ 9
= One 3s, three 3p and five 3d: 1 + 3 + 5 = 9 = 3².
? What is wrong with Bohr's picture of the hydrogen ground state?
- Its energy is wrong
+ It gives the electron angular momentum ħ, but the 1s state has zero orbital angular momentum
- It puts the electron too far from the nucleus
= Bohr's energies were right, but the 1s orbital has l = 0 and no orbit at all.
:::
