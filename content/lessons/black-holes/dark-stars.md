The first argument for black holes is more than two centuries old and uses nothing beyond Newton's gravity and high-school algebra. It turns out to be wrong in its picture of light, yet it lands on exactly the right number. Following it carefully is the best warm-up for the real thing.

## Light as a projectile

In the 1700s most natural philosophers followed Newton in treating light as a stream of tiny particles ("corpuscles"). If light has mass, gravity should pull on it like on anything else. A ball thrown upward slows down; so should light leaving a star. That raises an obvious question: could a star be so strong that its light never gets away?

To answer it you need the **escape velocity**: the minimum launch speed that lets a projectile coast away forever, never falling back.

## Deriving escape velocity

Take a star of mass M and radius R, and a projectile of mass m launched straight up from the surface with speed v. Ignore air and the star's rotation.

::: steps Escape velocity from energy conservation
$ E = \tfrac{1}{2}mv^2 - \frac{GMm}{R}
The total energy at launch is kinetic energy plus gravitational potential energy. The potential energy of two masses a distance r apart is −GMm/r: zero when they are infinitely far apart, more negative the closer they are.
---
$ E = \tfrac{1}{2}mv_\infty^2 - \frac{GMm}{\infty} = \tfrac{1}{2}mv_\infty^2
Far away the potential energy has gone to zero, leaving only the kinetic energy of whatever speed v∞ the projectile still has.
---
$ \tfrac{1}{2}mv^2 - \frac{GMm}{R} = \tfrac{1}{2}mv_\infty^2
Gravity is conservative (no friction, nothing lost), so total energy is the same at launch and at infinity.
---
$ \tfrac{1}{2}mv^2 - \frac{GMm}{R} = 0
"Barely escaping" means arriving at infinity with zero speed left over: v∞ = 0. Any slower and the projectile falls back; any faster and it escapes with speed to spare.
---
$ v_{\text{esc}} = \sqrt{\frac{2GM}{R}}
Add GMm/R to both sides, cancel m, multiply by 2, take the square root. The projectile's mass has dropped out.
:::

Check it: for Earth (M = 5.97 × 10²⁴ kg, R = 6.37 × 10⁶ m) this gives 11.2 km/s, the number rocket engineers use. For the Sun (M = 1.99 × 10³⁰ kg, R = 6.96 × 10⁸ m) it gives about 618 km/s.

## Setting the escape velocity equal to c

Because m cancelled, the same formula should hold for a light corpuscle. Light moves at c ≈ 3.00 × 10⁸ m/s. A star traps its light if its escape velocity exceeds c, and the boundary case is where they are equal.

::: steps The radius at which light cannot escape
$ c = \sqrt{\frac{2GM}{R}}
Set the escape velocity equal to the speed of light.
---
$ c^2 = \frac{2GM}{R}
Square both sides.
---
$ R = \frac{2GM}{c^2}
Solve for R. Any star of mass M squeezed inside this radius would, on this argument, be dark.
:::

The radius depends only on the mass and grows in proportion to it. Plugging in numbers (G = 6.674 × 10⁻¹¹ m³ kg⁻¹ s⁻²):

::: equation Two examples
\text{Sun: } R &= \frac{2 \times 6.674 \times 10^{-11} \times 1.989 \times 10^{30}}{(2.998 \times 10^8)^2} \approx 2.95\,\text{km} \\ \text{Earth: } R &= \frac{2 \times 6.674 \times 10^{-11} \times 5.972 \times 10^{24}}{(2.998 \times 10^8)^2} \approx 8.9\,\text{mm}
:::

To make the Sun dark you would have to crush its entire mass, now spread across 1.4 million kilometres, into a ball about 6 km across. Earth would need to fit inside a marble. Nothing in everyday physics does this, which is why the idea stayed a curiosity for a long time.

::: model escape
Shrink a star of fixed mass. Escape velocity grows as 1/√R; when it passes c, the star's light can no longer leave.
:::

::: history 1783 · Thornhill, England
John Michell, a Yorkshire rector and one of the best scientists of his day, sent this argument to the Royal Society. He phrased it differently: a star with the Sun's density but 500 times its radius would have an escape speed of about c (at fixed density, escape velocity grows in proportion to R, and 500 × 618 km/s ≈ 3.1 × 10⁵ km/s). He also suggested how to find such a star: watch for a visible star orbiting an unseen companion. That is close to how the first black hole was found, nearly two centuries later. In 1796 Pierre-Simon Laplace independently put the same idea in his *Exposition du système du monde*, then quietly dropped it from later editions.
:::

## What is wrong, and what is right

The reasoning fails in detail. Light is not a slowed cannonball: experiments from the 1800s on showed that light behaves as a wave, and relativity later established that it always travels at c locally, however strong the gravity. Light climbing out of a gravitational field does not slow down; it loses energy by shifting to longer wavelengths. Also, in Michell's picture a corpuscle launched from just inside the radius would rise some distance and fall back, so a nearby observer could still see the star. In general relativity nothing from inside crosses outward at all.

Even the formula's ingredients are suspect: ½mv² is the wrong kinetic energy for anything moving near c, and Newton's potential is only an approximation when gravity is strong.

::: note
The Newtonian argument gets the physics wrong but the number right. Einstein's theory yields the same critical radius, 2GM/c², now called the Schwarzschild radius. This is not quite luck: general relativity is built to agree with Newton in weak gravity, and that agreement is what fixes the constant, as the next lesson shows.
:::

So the dark star is a hint, not a proof. To find out whether such objects can really exist, and what happens at that radius, you need a theory in which gravity and light are treated correctly.

::: quiz
? Escape velocity does not depend on the projectile's mass. Why?
- Because light has no mass
+ Because m appears in both the kinetic and potential energy and cancels
- Because gravity is weak
= Both ½mv² and GMm/R are proportional to m, so dividing through removes it.
? A star of fixed mass is squeezed to one quarter of its radius. Its escape velocity
- quadruples
+ doubles
- halves
= v_esc = √(2GM/R); dividing R by 4 multiplies v by √4 = 2.
? What does general relativity say about Michell's critical radius?
- It is wrong by a factor of two
- It does not appear at all
+ The same radius, 2GM/c², appears, though the physics behind it is different
= The Schwarzschild radius has the same formula, even though light does not slow down as Michell assumed.
:::
