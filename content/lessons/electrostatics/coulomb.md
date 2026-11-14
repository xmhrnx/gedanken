Knowing that like charges repel is not enough to calculate anything. You need to know *how strongly*, and how the force changes with distance. The answer has the same shape as Newton's law of gravity: proportional to the product of the two charges, and falling off as the square of the distance.

## The law

Two point charges q₁ and q₂, a distance r apart, push or pull on each other along the line joining them with a force of size

::: equation Coulomb's law
F = k\frac{q_1 q_2}{r^2}, \qquad k = \frac{1}{4\pi\varepsilon_0} \approx 8.99 \times 10^9\,\text{N m}^2/\text{C}^2
:::

If the product q₁q₂ is positive (like charges) the force is repulsive; if negative, attractive. The constant ε₀ ≈ 8.85 × 10⁻¹² C²/(N m²), the **permittivity of free space**, is how k is usually written in more advanced formulas; the 4π appears because a sphere of radius r has area 4πr².

The forces obey Newton's third law: q₁ pushes on q₂ exactly as hard as q₂ pushes on q₁. And when several charges are present, the force on any one of them is the vector sum of the forces from each of the others, computed as if the rest were absent. This is the **principle of superposition**, and it is what makes electrostatics calculable.

## How it was found

::: history 1767 · Leeds, England
Joseph Priestley, following a hint from Franklin, noticed that a charged hollow metal vessel exerts no force on a charge placed inside it. He recognized that Newton had proved the same thing for gravity inside a hollow shell, which works only for an inverse-square law, and guessed that electric force also goes as 1/r². Henry Cavendish confirmed this precisely around 1773 but never published.
:::

::: history 1785 · Paris
Charles-Augustin de Coulomb, a military engineer, built a **torsion balance**: a light horizontal rod hung from a fine silver wire, with a small charged ball on one end. A second charged ball was brought near, and the repulsion twisted the wire. Coulomb had shown that the twisting force is proportional to the angle of twist, so he could measure tiny forces by turning a dial at the top until the balls sat at a chosen separation. In his best-known run, halving the separation needed about four times the twist, and halving it again about sixteen times: an inverse-square law.
:::

Coulomb's measurements were only good to a few percent. Modern versions of Priestley's hollow-conductor test show that the exponent is 2 to better than one part in 10¹⁶.

## How strong is it?

A coulomb is an enormous charge. Two charges of 1 C each, 1 m apart, would repel each other with 8.99 × 10⁹ N, roughly the weight of a million tonnes. We never see forces like that because ordinary matter is almost perfectly neutral; the positive and negative charges cancel to extraordinary precision.

The comparison with gravity makes the point. Take the electron and proton in a hydrogen atom.

::: steps Electric versus gravitational force in hydrogen
$ F_e = k\frac{e^2}{r^2} = \frac{(8.99 \times 10^9)(1.602 \times 10^{-19})^2}{(5.29 \times 10^{-11})^2} \approx 8.2 \times 10^{-8}\,\text{N}
The electron and proton each carry charge e, and their typical separation is about 5.29 × 10⁻¹¹ m (the Bohr radius).
---
$ F_g = G\frac{m_e m_p}{r^2} = \frac{(6.67 \times 10^{-11})(9.11 \times 10^{-31})(1.67 \times 10^{-27})}{(5.29 \times 10^{-11})^2} \approx 3.6 \times 10^{-47}\,\text{N}
The same calculation with Newton's law, using the masses of the electron and proton.
---
$ \frac{F_e}{F_g} = \frac{ke^2}{G m_e m_p} \approx 2.3 \times 10^{39}
Both forces go as 1/r², so the distance cancels. The ratio is the same at every separation.
:::

Electric forces outweigh gravity by about 10³⁹ between basic particles. Gravity rules the large-scale universe only because it has one sign and always adds up, while electric charges come in two signs that cancel. The electric force is what holds atoms together, makes solids solid and drives all of chemistry; the force of your hand on a table is electrical repulsion between electrons.

::: note
Coulomb's law is an inverse-square law like gravity, but about 10³⁹ times stronger between an electron and a proton. Matter hides this strength by being almost exactly neutral.
:::

## Two quick checks

Because F ∝ 1/r², doubling the distance cuts the force to a quarter, and tripling it cuts it to a ninth. Because F ∝ q₁q₂, doubling either charge doubles the force. Unlike gravity, the force can be repulsive, and it can be screened: charges in a conductor rearrange themselves to cancel fields, whereas nothing blocks gravity.

::: quiz
? Two charges repel with force F. Each charge is doubled and the distance is doubled. The new force is
- 4F
+ F
- F/4
= Doubling both charges multiplies F by 4; doubling the distance divides it by 4. The effects cancel.
? Why don't we notice the enormous electric forces between everyday objects?
- Electric forces only act inside atoms
+ Objects carry almost equal amounts of positive and negative charge, which cancel
- Gravity is stronger at human scales
= The force between individual particles is huge, but in neutral matter attractions and repulsions almost exactly cancel.
? Why does the ratio of electric to gravitational force in hydrogen not depend on the distance?
- Because the charges are equal
+ Because both forces fall off as 1/r²
- Because the masses are tiny
= Both laws have r² in the denominator, so it cancels in the ratio.
:::
