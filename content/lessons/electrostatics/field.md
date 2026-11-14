Coulomb's law says two charges push on each other across empty space, with nothing in between. Newton was uneasy about this kind of action at a distance for gravity, and many physicists shared the unease. The **electric field** replaces it with a different picture: a charge changes the space around it, and other charges respond to the space where they are.

## Force per unit charge

Put a small test charge q at some point near other charges. It feels a force F. Double the test charge and the force doubles, because Coulomb's law is proportional to q. So the ratio F/q does not depend on the test charge at all. It depends only on the other charges and on where you are. That ratio is the electric field.

::: equation Definition of the electric field
\vec{E} = \frac{\vec{F}}{q}
:::

E is a vector at every point in space: its direction is the direction of the force on a positive charge, and its size is measured in newtons per coulomb (N/C), which the next lesson shows is the same as volts per metre. Once you know E at a point, the force on any charge placed there is simply F = qE. A negative charge is pushed opposite to E.

## Field of a point charge

::: steps The field around a single charge
$ F = k\frac{Qq}{r^2}
Coulomb's law for a test charge q a distance r from a charge Q.
---
$ E = \frac{F}{q} = k\frac{Q}{r^2}
Divide by q. The test charge drops out, as promised.
---
$ \vec{E} = k\frac{Q}{r^2}\,\hat{r}
Add the direction: r̂ is a unit vector pointing away from Q. For positive Q the field points outward; for negative Q, inward.
:::

For example, 1 nanocoulomb (10⁻⁹ C) produces a field of 8.99 × 10⁹ × 10⁻⁹ / (0.01)² ≈ 9 × 10⁴ N/C at a distance of 1 cm. Fields from several charges add as vectors, by superposition, just as forces do.

## Field lines

A vector at every point is hard to draw, so we draw **field lines**: curves that run along the direction of E everywhere. The rules:

- Lines start on positive charges and end on negative charges (or run off to infinity).
- The direction of a line at any point is the direction of E there.
- Where lines crowd together the field is strong; where they spread out it is weak.
- Lines never cross, because E has only one direction at each point.

The third rule is not just a convention; for a point charge it matches the inverse-square law exactly. Draw N lines leaving a charge. At distance r they pierce a sphere of area 4πr², so their density is N/(4πr²), which falls as 1/r² just like E. This works because field lines in three dimensions spread over area that grows as r², and it is the idea behind **Gauss's law**: the number of lines leaving any closed surface is proportional to the charge inside.

::: model field-lines
Change the second charge's sign and size. With opposite signs the lines run from one charge to the other; with like signs they bend apart, leaving a point between where the field is zero.
:::

Two equal and opposite charges close together form a **dipole**, the pattern of a polarized atom or a water molecule. Two like charges produce a field that vanishes at the midpoint, where their pushes cancel.

::: history 1830s–1850s · London
Michael Faraday, with almost no formal mathematics, thought about electricity and magnetism in terms of "lines of force" filling space. He took them to be physically real: stretched along their length and pushing sideways on each other. Most mathematicians treated them as a picture. In the 1850s and 1860s James Clerk Maxwell turned Faraday's lines into equations for the field, and found that changes in the field travel at the speed of light.
:::

## Is the field real?

For charges at rest, the field is just a way of reorganizing Coulomb's law; it predicts nothing new. The difference shows up when charges move. If you suddenly shake a charge, a distant charge does not feel the change instantly; it feels it a time r/c later, when the disturbance in the field arrives. In the meantime the field carries energy and momentum. Light itself is a travelling disturbance of electric and magnetic fields. So physicists treat fields as real things, not bookkeeping.

## Conductors

Inside a conductor in equilibrium the field is zero. If it were not, the free electrons would move until it was. Any excess charge therefore sits on the surface, and the field just outside is perpendicular to the surface. This is why a metal car or aircraft shields its occupants from lightning, and why a hollow metal box, a **Faraday cage**, blocks outside electric fields.

::: note
A charge sets up a field throughout space, and any other charge feels a force qE from the field where it sits. The field of a point charge falls as 1/r², which is why field line density captures its strength.
:::

::: quiz
? At some point the electric field is 500 N/C pointing east. What force acts on an electron placed there?
- 500 N east
+ About 8 × 10⁻¹⁷ N west
- About 8 × 10⁻¹⁷ N east
= F = qE = (1.6 × 10⁻¹⁹ C)(500 N/C) = 8 × 10⁻¹⁷ N; the electron's charge is negative, so the force points opposite to E.
? You double the test charge at a point. The electric field there
- doubles
- halves
+ stays the same
= E = F/q; doubling q doubles F, so the ratio is unchanged. The field is set by the other charges.
? Why can two field lines never cross?
- Because they repel each other
+ Because the field has a single direction at each point
- Because charges sit at the crossing points
= A crossing would mean two different field directions at the same place.
:::
