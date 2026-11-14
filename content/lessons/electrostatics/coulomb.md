Knowing that like charges repel and unlike charges attract is a start, but it is not enough to calculate anything. To predict how charges will move, or how strongly an atom holds its electrons, you need to know how big the force is, and how it changes as the charges move apart. Through the eighteenth century, as electrical experiments became more careful, people suspected that the answer would mirror Newton's law of gravity. It does. The electric force between two small charges is proportional to the product of the charges and falls off as the square of the distance between them. This lesson states that law, follows how it was found, and then compares its strength with gravity's, a comparison that turns out to explain why the world is built the way it is.

## What does the law say?

Take two **point charges**, meaning charged objects small compared with the distance between them, so that their size and shape do not matter. Call their charges q₁ and q₂ and the distance between them r. Each pushes or pulls on the other along the straight line joining them, with a force of size

::: equation Coulomb's law
F = k\frac{q_1 q_2}{r^2}, \qquad k = \frac{1}{4\pi\varepsilon_0} \approx 8.99 \times 10^9\,\text{N m}^2/\text{C}^2
:::

The formula has the same shape as Newton's law of gravitation, F = Gm₁m₂/r², with charges in place of masses. The sign of the result carries the direction. If q₁ and q₂ have the same sign, their product is positive and the force is repulsive; if they have opposite signs, the product is negative and the force is attractive. Gravity, by contrast, has only one sign: masses always attract.

The constant k sets the strength of the force in the units we use. In more advanced formulas it is usually written as 1/(4πε₀), where ε₀ ≈ 8.85 × 10⁻¹² C²/(N m²) is called the **permittivity of free space**. Writing it this way looks like a complication, but the 4π has a meaning: a sphere of radius r has surface area 4πr², and putting the 4π into the constant makes later laws about fields spreading out over spheres come out without it.

Two further facts make the law usable. First, the forces obey Newton's third law: q₁ pushes or pulls on q₂ exactly as hard as q₂ does on q₁, in the opposite direction. Second, when there are several charges, the force on any one of them is the sum of the forces from each of the others, calculated one at a time as if the rest were absent, and added as **vectors** (quantities with direction as well as size). This is the **principle of superposition**. It is not obvious that nature should work this way; the presence of a third charge might have changed how the first two interact. Experiment shows it does not, and that is what makes electrostatics calculable: any arrangement of charges, however complicated, can be handled by breaking it into pairs.

## How was the inverse-square law found?

The interesting part of the law is the exponent: why r², and not r, or r³? Two very different experiments pointed to the answer, and the first was indirect.

::: history 1766 · Warrington, England
Joseph Priestley was a dissenting minister and teacher who had become fascinated by electricity and was writing a history of it. Benjamin Franklin had mentioned to him a puzzle: a small cork ball lowered inside a charged metal can felt no electric force at all. Priestley repeated the experiment and confirmed it. Then he made a leap of reasoning. Newton had proved, nearly a century earlier, that inside a hollow shell of matter the gravitational pulls from all sides cancel exactly, so an object anywhere inside feels no force. That cancellation works only for a force that falls off as 1/r²: the nearby part of the shell pulls harder, but the farther part is bigger, and with an inverse square the two effects balance precisely. Priestley suggested, in his book published in 1767, that electric force must also follow an inverse-square law. A few years later Henry Cavendish tested the idea with great precision, but, being intensely private, he never published the result, and it lay hidden until Maxwell edited his papers in 1879.
:::

The second experiment was direct: actually measure the force at different distances. The difficulty is that electric forces between small charged objects are tiny, far too small to weigh on an ordinary balance.

::: history 1785 · Paris
Charles-Augustin de Coulomb, a French military engineer, solved that problem with a **torsion balance**. A light horizontal rod hung from a very fine silver wire, with a small charged ball on one end. A second charged ball, fixed in place, was brought near the first, and the repulsion between them swung the rod and twisted the wire. Coulomb had earlier shown that the force needed to twist a thin wire is proportional to the angle of twist, so the wire itself became a sensitive force meter. By turning a dial at the top of the wire he could twist it back until the balls sat at whatever separation he chose, and read the force from the angle. In his best-known run, halving the separation required about four times the twist, and halving it again required about sixteen times. Force multiplied by four when distance is halved is exactly what an inverse-square law predicts.
:::

Coulomb's measurements were accurate only to a few percent: charge leaked from the balls into the air, and the balls were not truly points. Modern versions of the hollow-conductor test, the descendants of Priestley's and Cavendish's idea, are far more sensitive. They show that if the exponent is not exactly 2, it differs from 2 by less than one part in 10¹⁶.

## How strong is the electric force?

The coulomb is an enormous amount of charge, and a single calculation shows how enormous. Two charges of 1 C each, held 1 m apart, would push on each other with a force of 8.99 × 10⁹ newtons. That is close to the weight of a million tonnes. We never see forces anything like this in daily life, and the reason is that ordinary matter is almost perfectly neutral. A sugar cube contains around two hundred thousand coulombs of positive charge in its protons, but an equal amount of negative charge in its electrons, and the two cancel to extraordinary precision.

The best way to see the true strength of the electric force is to compare it with gravity in a case where both act on the same pair of particles. The simplest such pair is the electron and proton in a hydrogen atom.

::: steps Electric versus gravitational force in hydrogen
$ F_e = k\frac{e^2}{r^2} = \frac{(8.99 \times 10^9)(1.602 \times 10^{-19})^2}{(5.29 \times 10^{-11})^2} \approx 8.2 \times 10^{-8}\,\text{N}
The proton has charge +e and the electron −e, so the product of the charges has size e². For the distance, use 5.29 × 10⁻¹¹ m, the Bohr radius, which is the most likely distance between the two in a hydrogen atom in its lowest-energy state. Coulomb's law then gives an attractive force of about 8 × 10⁻⁸ newtons, tiny by everyday standards but large for particles this light.
---
$ F_g = G\frac{m_e m_p}{r^2} = \frac{(6.67 \times 10^{-11})(9.11 \times 10^{-31})(1.67 \times 10^{-27})}{(5.29 \times 10^{-11})^2} \approx 3.6 \times 10^{-47}\,\text{N}
Now do the same calculation with Newton's law of gravitation, using the gravitational constant G and the masses of the electron and proton, at the same distance. The two particles are so light that their gravitational attraction is almost unimaginably small.
---
$ \frac{F_e}{F_g} = \frac{ke^2}{G m_e m_p} \approx 2.3 \times 10^{39}
Divide the first force by the second. Both laws have r² in the denominator, so the distance cancels: the ratio is the same whether the particles are inside an atom or a kilometre apart. What is left depends only on fundamental constants, which makes it a fair comparison of the two forces themselves.
:::

Between an electron and a proton, the electric force is about 10³⁹ times stronger than gravity. That is a number with thirty-nine zeros, and nothing else in physics prepares you for a ratio that large.

Why, then, does gravity rule the motion of planets, stars and galaxies? Because gravity has only one sign. Every bit of mass attracts every other, so in a planet or a star the contributions all add up. Electric charges come in two signs that cancel, and large objects are almost exactly neutral, so their electric effects on each other nearly vanish. On the scale of atoms and everyday objects, though, the electric force does almost everything. It holds electrons to nuclei, binds atoms into molecules, drives all of chemistry and biology, and makes solid objects solid. When you press your hand on a table, the push you feel comes from the electrons in your hand and in the table repelling each other, with help from a rule of quantum mechanics, the exclusion principle, that prevents electrons from crowding into the same state.

::: note
Coulomb's law is an inverse-square law like gravity, but between an electron and a proton the electric force is about 10³⁹ times stronger. Ordinary matter hides that strength by being almost exactly neutral, which is why gravity dominates at large scales and electricity at small ones.
:::

## How does the force change with charge and distance?

Because the force goes as 1/r², distance has a strong effect. Double the distance between two charges and the force falls to a quarter; triple it and the force falls to a ninth. Because the force is proportional to q₁q₂, doubling either charge doubles the force, and doubling both multiplies it by four.

One more difference from gravity matters in practice. Electric forces can be screened. Put a charge near a metal object and the free electrons in the metal rearrange themselves to cancel the electric effects inside it. That is why the cork ball in Franklin's can felt nothing. Nothing similar is known for gravity: there is no material that blocks it, because there is no negative mass to do the cancelling.

::: quiz
? Two charges repel with force F. Each charge is doubled and the distance between them is doubled. What is the new force?
- 4F
+ F
- F/4
= Doubling both charges multiplies the product q₁q₂, and so the force, by 4. Doubling the distance divides the force by 2² = 4, so the two effects cancel and the force is F again.
? Why don't we notice the enormous electric forces between everyday objects?
- Electric forces only act inside atoms
+ Objects carry almost equal amounts of positive and negative charge, which cancel
- Gravity is stronger at human scales
= The forces between individual charged particles are huge, but everyday objects contain almost exactly equal numbers of protons and electrons. The attractions and repulsions very nearly cancel, leaving almost no net force.
? Why does the ratio of electric to gravitational force in hydrogen not depend on the distance?
- Because the charges are equal
+ Because both forces fall off as 1/r²
- Because the masses are tiny
= Both Coulomb's law and Newton's law of gravitation have r² in the denominator. When one force is divided by the other, the r² cancels, so the ratio is the same at every distance.
:::
