By 1869 chemists had found that the elements, sorted in order of weight, show a repeating pattern of behaviour, and they had arranged them in a table to display it. Lithium, sodium and potassium, at atomic numbers 3, 11 and 19, are soft metals that react violently with water in almost identical ways. Helium, neon, argon and krypton, at 2, 10, 18 and 36, are gases that react with almost nothing. The pattern was real and useful, but nobody knew why it existed, and in particular why the repeat lengths were 2, 8, 8, 18, 18 and 32 rather than anything else. Quantum mechanics answers that question completely, using the orbitals of the previous lesson and one additional rule. This lesson builds the periodic table from those two ingredients.

::: history 1869 · St Petersburg
Dmitri Mendeleev, a chemistry professor writing a textbook, was looking for a way to organise the sixty or so elements then known. He wrote each element's properties on a card and arranged the cards by atomic weight, in rows, so that elements with similar chemistry fell in the same column. What set his table apart from earlier attempts was his confidence in it. Where the pattern demanded an element that nobody had found, he left a gap and predicted the missing element's properties: its weight, its density, the formulas of its compounds. "Eka-aluminium", the element below aluminium, was discovered as gallium in 1875; "eka-boron" turned up as scandium in 1879, and "eka-silicon" as germanium in 1886, each close to his predictions. In a few places he also put a heavier element before a lighter one, because that was the only way to keep the columns chemically consistent. He could not say why this was needed. In 1913 Henry Moseley showed, from the X-rays emitted by different elements, that the true order is not by weight at all but by the charge on the nucleus, the **atomic number** Z, and in Moseley's ordering Mendeleev's exceptions disappear.
:::

## Why don't all the electrons fall to the bottom?

A neutral atom of atomic number Z has Z electrons around its nucleus. Physical systems tend to settle into their lowest-energy state, so you might expect every electron to drop into the lowest orbital, 1s. If that happened, atoms would differ only gradually from one to the next: adding one more proton and one more electron would make each atom a little smaller and more tightly bound than the last, and chemistry would change smoothly across the elements. There would be no repeating pattern, no sudden jump from an inert gas to a violently reactive metal between Z = 10 and Z = 11.

The jumps are there, so something stops the electrons from crowding into the same state. That something is a rule found by Wolfgang Pauli.

::: note
The Pauli exclusion principle says that no two electrons in an atom can have the same four quantum numbers n, l, m_l and m_s. So each orbital holds at most two electrons, one with spin up and one with spin down. Combined with the counting of orbitals, this single rule produces the whole structure of the periodic table.
:::

::: history 1925 · Hamburg
Pauli, then a young lecturer in Hamburg, was trying to make sense of the complicated ways in which atomic spectral lines split in magnetic fields. A 1924 paper by the English physicist Edmund Stoner had proposed a scheme for how many electrons each shell holds, based on chemistry and spectra, and Pauli saw that Stoner's numbers would follow from a simple rule if the electron had a fourth quantum number, one that could take only two values, and if no two electrons could share all four. He stated the rule because it fitted the data, without any deeper reason for it, and he had no idea what the two-valued property was. Later that year Uhlenbeck and Goudsmit proposed electron spin, which supplied it. The principle is now known to hold for every particle with half-integer spin, including protons and neutrons, and it reaches far beyond atoms: it is what holds up white dwarf stars, the dense remnants of stars like the Sun, against their own gravity, because their electrons cannot all be squeezed into the same low-energy states.
:::

## Filling the shells

Now build atoms one at a time. Each time we go to the next element, we add one proton to the nucleus and one electron, and the new electron goes into the lowest-energy orbital that still has room. The exclusion principle decides how much room there is, and counting quantum numbers tells us how many electrons each kind of orbital can take.

::: steps Capacity of subshells and periods
$ m_l = -l, \ldots, +l \quad\Rightarrow\quad 2l + 1 \text{ orbitals}
A **subshell** is a set of orbitals with the same n and l, such as 2p or 3d. Within it, m_l takes each whole value from −l to +l, so the subshell contains 2l + 1 orbitals, one for each orientation. This is the same count as in the previous lesson.
---
$ 2(2l + 1) = 2,\; 6,\; 10,\; 14 \quad \text{for s, p, d, f}
By the exclusion principle, each orbital can take two electrons, which must have opposite spins. Doubling the number of orbitals gives the capacity of each subshell: an s subshell holds 2 electrons, p holds 6, d holds 10 and f holds 14.
---
$ 1s < 2s < 2p < 3s < 3p < 4s < 3d < 4p < 5s < 4d < 5p < 6s < 4f < 5d < 6p
In atoms with many electrons, the inner electrons screen the outer ones from the full charge of the nucleus. Orbitals with low l have some probability very close to the nucleus, inside the screening, so they are held more tightly than orbitals of the same n with higher l. The effect is strong enough to push some subshells past others: 4s fills before 3d, and 6s before 4f. This list is the order in which subshells fill.
---
$ 2,\; 2 + 6 = 8,\; 2 + 6 = 8,\; 2 + 10 + 6 = 18,\; 2 + 10 + 6 = 18,\; 2 + 14 + 10 + 6 = 32
Group the list into rows. Each row of the table begins when electrons start entering a new s subshell and ends when the following p subshell is full, except the first row, which has only 1s. So the rows hold 1s; 2s 2p; 3s 3p; 4s 3d 4p; 5s 4d 5p; and 6s 4f 5d 6p. Adding the capacities of the subshells in each row gives the row lengths.
:::

These are exactly the lengths of the rows of the periodic table: 2, 8, 8, 18, 18, 32. The numbers that puzzled chemists for half a century come out of counting standing waves and applying one rule about electrons. The filling order is a very good general guide but not an exact law, because electrons also repel one another and the energies of some subshells lie very close together. A few elements, such as chromium and copper, deviate from the simple order by one electron.

::: model orbital
Compare 3s, 3p and 3d. All three belong to the same shell, but look at how close to the nucleus each one's cloud reaches. The 3s cloud has inner parts right next to the nucleus, which is why it is held more tightly than 3d in atoms with many electrons.
:::

## Chemistry from shells

An atom's chemistry depends mainly on its outermost electrons, called the **valence electrons**. The reason is that they are the electrons least tightly held, and the ones another atom meets first when the two come close. The inner electrons, in filled shells, are bound tightly and take little part in reactions. So elements with the same arrangement of valence electrons behave alike, and that is what the columns of the table group together.

Look first at the end of each row. The **noble gases**, helium, neon, argon and the rest, have every subshell in their outer shell full. A full subshell is a particularly stable arrangement, so removing an electron is hard: it takes 24.6 eV for helium and 21.6 eV for neon, more than for any other element. Adding one is no easier, because it would have to go into the next, much higher shell. With nothing to gain from giving or taking electrons, they barely react.

Now step forward by one element, to the start of the next row. The **alkali metals**, lithium, sodium, potassium and the rest, have a single electron in a new s orbital, sitting outside a closed shell. That lone electron is far from the nucleus and is screened by all the electrons inside it, so it feels only about one unit of nuclear charge. Removing it takes only 5.1 eV for sodium, a quarter of what neon demands, although sodium has one more proton. Alkali metals give that electron away readily, which is why they react so violently.

Step back by one from the noble gases instead and you reach the **halogens**, fluorine, chlorine and the rest, which are one electron short of a closed shell. Gaining one electron completes the shell, so they take electrons eagerly from other atoms. Put sodium and chlorine together and sodium's lone electron passes to chlorine, leaving two ions held together by electric attraction, Na⁺ and Cl⁻: table salt.

In the middle of the longer rows sit the **transition metals**, such as iron, nickel and copper. Here the electrons are filling d orbitals, which lie further in than the outer s electrons. From one transition metal to the next, the outermost electrons barely change, so ten elements in a row look chemically similar, hard metals with several possible charges, quite unlike the sharp changes across the rest of a row.

So the columns of Mendeleev's table turn out to be counts of electrons outside closed shells. Sodium, with a single electron in 3s, sits under lithium, with a single electron in 2s, and above potassium, with a single electron in 4s.

## Writing it down

Chemists record the arrangement of an atom's electrons, its **electron configuration**, as a list of occupied subshells, each with a superscript giving the number of electrons in it. Neon, Z = 10, is 1s² 2s² 2p⁶: two completely filled shells, which is why it is inert. Sodium, Z = 11, is 1s² 2s² 2p⁶ 3s¹: neon's arrangement plus one more electron in a new shell, which is why it is so reactive. For heavier atoms the filled inner shells are abbreviated by the noble gas that has them, so iron, Z = 26, is written [Ar] 4s² 3d⁶, meaning argon's 18 electrons plus eight more, two in 4s and six in 3d.

::: quiz
? Why don't all of an atom's electrons sit in the 1s orbital?
- The 1s orbital is too small to hold more than one electron
+ The exclusion principle allows only two electrons, of opposite spin, in each orbital
- Electrons repel each other too strongly to share a region
= No two electrons can share all four quantum numbers, and the 1s orbital offers only two combinations, spin up and spin down, so the third electron must go into a higher orbital.
? How many electrons can a d subshell hold?
- 5
+ 10
- 18
= A d subshell has l = 2, which gives 2l + 1 = 5 orbitals, and each orbital holds two electrons of opposite spin, so the total is 10.
? Why do sodium and potassium behave alike chemically?
- They have similar atomic weights
+ Each has a single s electron outside closed shells
- They are both metals
= Chemistry is set mainly by the valence electrons, and both atoms have one loosely held electron outside a closed shell, in 3s for sodium and 4s for potassium.
:::
