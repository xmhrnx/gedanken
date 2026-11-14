Coulomb's law describes two charges pushing or pulling on each other across empty space, with nothing in between to carry the push. Newton had faced the same situation with gravity and disliked it. He wrote that the idea of one body acting on another at a distance, through a vacuum, without anything to convey the action, seemed to him an absurdity. Many physicists shared that unease, but for charges at rest the law worked perfectly, so for a long time the question was set aside. The **electric field** offers a different picture. Instead of saying that one charge acts directly on a distant one, it says that a charge changes the space around it, and that any other charge responds to the condition of space right where it sits. This lesson defines the field, works out the field of a single charge, shows how to draw it, and explains why physicists came to believe the field is real and not just a convenient way of talking.

## How can a force be turned into a property of space?

Suppose you have some arrangement of charges, perhaps a charged sphere or a pair of charged plates, and you want to describe their electrical influence on the space around them. A natural way is to probe it. Take a small **test charge** q, small enough that it does not disturb the other charges, and place it at some point. It feels a force F, which you can measure.

The force depends on the test charge you chose: by Coulomb's law it is proportional to q, so a test charge twice as big feels twice the force. That is inconvenient if you want to describe the arrangement itself rather than your probe. But it suggests a remedy. If force is proportional to q, then the ratio F/q is the same whatever test charge you use. It depends only on the other charges and on where you put the probe. That ratio is the electric field.

::: equation Definition of the electric field
\vec{E} = \frac{\vec{F}}{q}
:::

The arrow over E and F marks them as **vectors**: each has a direction as well as a size. The field E at a point has the direction of the force that a positive test charge would feel there, and its size is measured in newtons per coulomb (N/C). The next lesson shows that this unit is the same as volts per metre, which is how fields are more often quoted.

Defining the field this way divides one problem into two. First, the charges you are studying set up a field throughout space. Second, any charge placed at a point responds to the field at that point, feeling a force F = qE. For a positive charge, the force points along E; for a negative charge, such as an electron, it points the opposite way. Once you know the field, you can predict the force on any charge without recalculating Coulomb's law each time.

## What is the field of a single charge?

The simplest case is the field around one point charge Q. Coulomb's law gives the force on a test charge, and the definition turns it into a field.

::: steps The field around a single charge
$ F = k\frac{Qq}{r^2}
Place a test charge q at a distance r from the charge Q. Coulomb's law gives the size of the force between them, which acts along the line joining the two. This is the measurement our definition starts from.
---
$ E = \frac{F}{q} = k\frac{Q}{r^2}
Divide the force by the test charge, as the definition of the field requires. Because the force is proportional to q, the test charge cancels completely. What remains depends only on the source charge Q and the distance r, which is exactly what a description of Q's influence should depend on.
---
$ \vec{E} = k\frac{Q}{r^2}\,\hat{r}
Finally, put the direction back in. The symbol r̂ is a **unit vector**, an arrow of length one, pointing directly away from Q toward the point in question. For positive Q, the field points outward in every direction, since a positive test charge would be repelled. For negative Q, the formula gives a negative multiple of r̂, so the field points inward, toward the charge.
:::

The field of a point charge, like the force, falls off as the square of the distance. To get a feel for the size, take a charge of 1 nanocoulomb, 10⁻⁹ C, which is about what a small object might pick up by being rubbed. At a distance of 1 cm, the field is 8.99 × 10⁹ × 10⁻⁹ / (0.01)² ≈ 9 × 10⁴ N/C. An electron placed there would feel a force of about 1.4 × 10⁻¹⁴ N, which sounds small but would give something as light as an electron an enormous acceleration.

When there are several charges, each produces its own field, and by the principle of superposition the total field at any point is the vector sum of them all. The field is a way of organising Coulomb's law, so it inherits that law's simplicity.

## How can a field be drawn?

A field assigns an arrow to every point in space, which is impossible to draw directly. Michael Faraday found a way around this: draw **field lines**, curves that everywhere run in the direction of the field. A few simple rules make these drawings surprisingly informative.

Field lines begin on positive charges and end on negative ones, or run off to great distances if there is no partner charge to end on. That follows from the direction of the field, which points away from positive charges and toward negative ones. At any point along a line, the direction of the line is the direction of E there, so a positive charge released at rest would start moving along it. Where the lines crowd together the field is strong, and where they spread apart it is weak. And two field lines can never cross, because at a crossing point the field would have to point in two directions at once, and a field has only one direction at each point.

The rule about crowding is not just a convenient convention; for a point charge it matches the inverse-square law exactly, and the reason is geometric. Draw some number of lines, say N, leaving a charge evenly in all directions. At distance r they all pass through an imaginary sphere of surface area 4πr², so the number of lines crossing each square metre is N/(4πr²). That density falls as 1/r², just like the field. The match works because space has three dimensions, so the area over which the lines spread grows as r². This is the idea behind **Gauss's law**, one of the basic laws of electricity: the number of field lines leaving any closed surface is proportional to the total charge inside it. It is the same reasoning Priestley used, in reverse, when he argued from the hollow can that the force must be inverse-square.

::: model field-lines
Change the second charge's sign and size. With opposite signs, the lines run from the positive charge to the negative one. With like signs, they bend away from each other, and there is a point between the charges where the field is zero.
:::

Two patterns come up again and again. A pair of equal and opposite charges close together is called a **dipole**. Its field lines loop from the positive charge around to the negative one, and this is the pattern of a polarized atom, of a water molecule, and of a radio antenna. Two equal charges of the same sign produce a field that is exactly zero at the midpoint between them, where their pushes on a test charge are equal and opposite and cancel.

::: history 1830s–1860s · London and Cambridge
Michael Faraday, the son of a blacksmith, left school at thirteen and became one of the greatest experimental physicists of his century, though he had almost no mathematical training. Instead of equations, he thought in pictures. Sprinkling iron filings around magnets and studying charged objects, he came to see the space around them as filled with "lines of force", and he took these lines to be physically real: under tension along their length, like stretched elastic, and pushing sideways against one another. The mathematicians of his day mostly treated the lines as a picture, useful for intuition but not part of the real physics.

James Clerk Maxwell, a generation younger and a superb mathematician, took Faraday's ideas seriously. Through the 1850s and 1860s he turned the lines of force into precise equations for the electric and magnetic fields. The equations made a prediction nobody had expected: a disturbance in the fields should travel through space as a wave, at a speed that could be calculated from electrical measurements. The speed came out equal to the measured speed of light. Light, Maxwell concluded, is an electromagnetic wave.
:::

## Is the field real, or just bookkeeping?

For charges at rest, the field adds nothing new. Every prediction it makes could be obtained directly from Coulomb's law, and the field is simply a more convenient way of organising the calculation. If that were all, it would be a matter of taste whether to call it real.

The difference appears when charges move. Suppose you suddenly shake a charge. Coulomb's law, taken literally, would say that a distant charge feels the change instantly. It does not. The distant charge feels nothing until a time r/c later, where c is the speed of light: the time it takes for the disturbance in the field to travel outward and arrive. During that interval, energy and momentum have left the region around the shaken charge but have not yet reached the distant one, so something in between must be carrying them. That something is the field. Light itself is exactly such a travelling disturbance of electric and magnetic fields, carrying energy across space from the Sun to Earth. A thing that carries energy and momentum and travels at a definite speed is hard to dismiss as bookkeeping, and physicists now treat fields as being as real as particles.

## What happens to the field inside a conductor?

Fields also explain an important property of metals. In a conductor that has settled down, with no current flowing, the electric field inside must be zero. The reason is that a conductor contains free electrons. If there were any field inside, it would push on them, and they would move. They keep moving, rearranging themselves, until their own field exactly cancels whatever field was there. Only then does the motion stop.

Two consequences follow. Any extra charge put on a conductor ends up on its outer surface, since charge inside would create a field inside. And the field just outside the surface points straight out from it, at right angles; if it had any component along the surface, charges would slide along until it vanished. A hollow metal box shields its interior from outside electric fields in the same way, because the charges on its surface rearrange to cancel them. Such a box is called a **Faraday cage**, and it is why the occupants of a car or an aircraft struck by lightning are protected: the metal skin carries the charge around them.

::: note
A charge sets up an electric field throughout the space around it, and any other charge feels a force qE from the field at the place where it sits. The field of a point charge falls off as 1/r², which is why the density of field lines, spreading over spheres of area 4πr², measures its strength.
:::

::: quiz
? At some point the electric field is 500 N/C pointing east. What force acts on an electron placed there?
- 500 N east
+ About 8 × 10⁻¹⁷ N west
- About 8 × 10⁻¹⁷ N east
= The size of the force is F = qE = (1.6 × 10⁻¹⁹ C)(500 N/C) = 8 × 10⁻¹⁷ N. The electron's charge is negative, so the force points opposite to the field, toward the west.
? You double the test charge placed at a point. What happens to the electric field there?
- It doubles
- It halves
+ It stays the same
= Doubling the test charge doubles the force on it, so the ratio E = F/q is unchanged. The field is set by the other charges, not by the probe used to measure it.
? Why can two field lines never cross?
- Because they repel each other
+ Because the field has a single direction at each point
- Because charges sit at the crossing points
= A field line runs along the direction of the field. If two lines crossed, the field at the crossing would have two different directions at once, which is impossible.
:::
