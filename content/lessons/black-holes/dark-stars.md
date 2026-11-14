The first argument that black holes might exist is more than two centuries old, and it uses nothing beyond Newton's law of gravity and the conservation of energy. Its picture of light turns out to be wrong. Yet it arrives at exactly the right number, the same critical radius that Einstein's theory would produce 130 years later. Following it slowly, step by step, is the best preparation for the real thing: it shows which quantities matter, why the mass of the escaping object drops out, and why the answer depends on how tightly a star's mass is packed rather than on how much there is. The following lessons then rebuild the argument on correct foundations.

## Could a star be too heavy for its own light to leave?

In the 1700s most natural philosophers followed Newton in thinking of light as a stream of tiny particles, which they called corpuscles. If light is made of particles with some mass, however small, then gravity should pull on them just as it pulls on everything else. A ball thrown upward slows down as it rises; light leaving the surface of a star should slow down too, by a tiny amount for an ordinary star.

That raises a natural question. A ball thrown up fast enough never comes back. A ball thrown too slowly does. So for any star there is a dividing speed, and if a star's gravity were strong enough, that dividing speed might exceed the speed of light. Light leaving such a star would rise, slow, stop and fall back, and the star would be invisible from afar.

To make this precise we need the dividing speed itself. It is called the **escape velocity**: the smallest launch speed with which a projectile can coast away forever, without any further push, and never fall back.

## How fast must something go to escape?

Take a star of mass M and radius R, and a projectile of mass m launched straight up from its surface at speed v. Ignore the star's atmosphere and rotation, so that gravity is the only force acting. The most direct route to the answer is through energy, because energy lets us compare the start and the end of the journey without following every moment in between.

::: steps Escape velocity from energy conservation
$ W = \int_R^\infty \frac{GMm}{r^2}\,\mathrm{d}r = \frac{GMm}{R}
First, find how much energy it takes to lift the projectile from the star's surface all the way to infinity. At distance r from the centre, gravity pulls with force GMm/r², and lifting against it through a short distance dr takes work equal to force times distance. Adding up all those small amounts from r = R to infinity gives GMm/R. The answer is finite even though the journey is infinitely long, because the pull weakens so fast with distance.
---
$ U(r) = -\frac{GMm}{r}
This is why the gravitational potential energy of the projectile at distance r is written this way. We choose to call the energy zero when the projectile is infinitely far away; then at distance r it has GMm/r less than that, since that is the work needed to get it out. The energy is negative, and more negative the closer the projectile is to the star.
---
$ E = \tfrac{1}{2}mv^2 - \frac{GMm}{R}
The total energy at launch is the kinetic energy, the energy of motion ½mv², plus the potential energy at the surface. A fast launch makes the first term large; a heavy, compact star makes the second term large and negative.
---
$ E = \tfrac{1}{2}mv_\infty^2
Now look at the far end of the journey. When the projectile is very far away, its potential energy has risen to zero, so all that remains is the kinetic energy of whatever speed v∞ it still has. If it stalls and falls back before getting there, it never reaches this stage at all.
---
$ \tfrac{1}{2}mv^2 - \frac{GMm}{R} = \tfrac{1}{2}mv_\infty^2
Gravity takes nothing away as heat or friction; it only converts kinetic energy into potential energy and back. So the total energy at launch equals the total energy far away, and we can set the two expressions equal. This single equation replaces a detailed calculation of the whole flight.
---
$ \tfrac{1}{2}mv^2 - \frac{GMm}{R} = 0
"Barely escaping" means arriving infinitely far away with no speed left over, so set v∞ = 0. A slightly slower launch makes the left side negative, and the projectile cannot reach infinity at all; a faster one leaves it with speed to spare. So this equation picks out exactly the dividing speed.
---
$ v_{\text{esc}} = \sqrt{\frac{2GM}{R}}
Move GMm/R to the right side, divide both sides by m, multiply by 2, and take the square root. The projectile's mass m has dropped out, because it appeared in both the kinetic and the potential energy. A pebble and a spacecraft need the same speed to escape, and so, on this argument, would a particle of light.
:::

Check the formula against something known. For the Earth, with M = 5.97 × 10²⁴ kg and R = 6.37 × 10⁶ m, it gives 11.2 km/s, the figure rocket engineers use. For the Sun, with M = 1.99 × 10³⁰ kg and R = 6.96 × 10⁸ m, it gives about 618 km/s. Both are far below the speed of light, about 300,000 km/s, so sunlight escapes easily.

Notice how the escape velocity depends on the star. It grows with the mass, but only as its square root, and it grows as the radius shrinks. So there are two ways to make escape harder: add mass, or squeeze the same mass into a smaller ball.

## At what size would light be trapped?

Because the projectile's mass cancelled, the same formula should apply to a corpuscle of light. Light travels at c ≈ 3.00 × 10⁸ m/s. A star traps its own light if its escape velocity is greater than c, and the boundary between visible and dark stars is where the two are exactly equal.

::: steps The radius at which light cannot escape
$ c = \sqrt{\frac{2GM}{R}}
Set the escape velocity equal to the speed of light. For a given mass M, this equation picks out the one radius R at which light would just barely escape. Any smaller radius makes the right side larger than c.
---
$ c^2 = \frac{2GM}{R}
Square both sides to remove the square root. Both sides are positive, so squaring keeps the equation true and loses nothing. Now R appears in a simple fraction.
---
$ R = \frac{2GM}{c^2}
Multiply both sides by R and divide both sides by c² to get R on its own. This is the critical radius. On Michell's argument, any star of mass M squeezed inside this radius would be dark, because even light would fall back.
:::

The result is worth looking at closely. The critical radius depends only on the mass, and it is directly proportional to it: double the mass and the critical radius doubles. The constant in front, 2G/c², is tiny, about 1.5 × 10⁻²⁷ metres per kilogram, because G is small and c² is enormous. So the radius is very small for any ordinary mass. With G = 6.674 × 10⁻¹¹ m³ kg⁻¹ s⁻²:

::: equation The critical radius of the Sun
R_{\text{Sun}} = \frac{2 \times 6.674 \times 10^{-11} \times 1.989 \times 10^{30}}{(2.998 \times 10^8)^2} \\ \approx 2.95\,\text{km}
:::

::: equation The critical radius of the Earth
R_{\text{Earth}} = \frac{2 \times 6.674 \times 10^{-11} \times 5.972 \times 10^{24}}{(2.998 \times 10^8)^2} \\ \approx 8.9\,\text{mm}
:::

To make the Sun dark, you would have to crush its whole mass, now spread across a ball 1.4 million kilometres wide, into a ball about 6 kilometres across, the size of a small town. The Earth would have to fit inside a marble. Nothing in everyday experience compresses matter remotely that much, which is why the idea remained a curiosity for a very long time. Whether anything in nature could do it is the question of the lesson on stellar collapse.

::: model escape
Shrink a star while keeping its mass fixed. The escape velocity grows as 1/√R; watch it pass c at the critical radius, at which point the star's light can no longer leave.
:::

::: history 1783 · Thornhill, England
John Michell was the rector of a small parish in Yorkshire and, in his spare time, one of the most inventive scientists in Britain: he founded the study of earthquakes and designed the torsion balance that Henry Cavendish later used to weigh the Earth. In 1783 he set out this argument in a long letter to Cavendish, which was read at the Royal Society that November. He framed it differently from us. He imagined a star with the same density as the Sun but a larger radius, and found that at 500 times the Sun's radius, the escape speed would reach the speed of light. (At fixed density, mass grows as R³, so the escape velocity √(2GM/R) grows in proportion to R; 500 × 618 km/s is about 3.1 × 10⁵ km/s.) Michell realised such stars, if they existed, could not be seen directly. But he also saw how they might be found anyway: if one had a visible companion star, the companion would be seen orbiting an invisible partner. That is very close to how the first black hole was actually identified nearly two centuries later. In 1796 the French mathematician Pierre-Simon Laplace independently put forward the same idea in his popular book on astronomy, then quietly removed it from later editions, probably because by then light was coming to be seen as a wave, to which the argument did not obviously apply.
:::

## What is wrong with the argument, and what is right?

The reasoning fails in detail, in several ways. First, light is not a slow cannonball. Experiments in the 1800s showed that light behaves as a wave, and relativity later established that light always travels at exactly c as measured by anyone nearby, however strong the gravity. Light climbing out of a gravitational field does not slow down; instead it loses energy by shifting to longer wavelengths, which is the gravitational redshift of the previous chapter.

Second, the picture of a dark star is wrong. In Michell's version, a corpuscle launched from a star just inside the critical radius would still rise some distance before falling back, so an observer hovering close to the star could still see it. In general relativity, as the coming lessons show, nothing from inside the critical radius ever moves outward at all, not even a short way.

Third, the formula's ingredients are themselves suspect for this job. The expression ½mv² for kinetic energy is only correct for speeds far below c, and here it is being applied at c. Newton's potential energy, −GMm/r, is only an approximation that works when gravity is weak, and here gravity is as strong as it can be.

So it is remarkable that the answer comes out right. The reason is not luck. General relativity is constructed so that it agrees with Newton's gravity wherever gravity is weak, and in the next lesson, that requirement is exactly what fixes the constant in Einstein's equations. The factor 2GM/c² is carried over from the Newtonian world into the relativistic one, and it arrives at the same place.

::: note
The Newtonian argument gets the physics wrong but the number right. Einstein's theory yields the same critical radius, 2GM/c², now called the Schwarzschild radius. This is not a coincidence: general relativity must agree with Newton where gravity is weak, and that agreement is what fixes the constant, as the next lesson shows.
:::

So the dark star is a hint, not a proof. To find out whether such objects can really exist, and what actually happens at that radius, we need a theory in which gravity and light are both treated correctly. That theory is general relativity, and the next lesson follows the first exact solution of its equations.

::: quiz
? Escape velocity does not depend on the mass of the projectile. Why not?
- Because light has no mass
+ Because m appears in both the kinetic and the potential energy and cancels
- Because gravity is weak
= Both ½mv² and GMm/R are proportional to m, so dividing the energy equation by m removes it. Every object needs the same speed to escape.
? A star of fixed mass is squeezed to one quarter of its radius. What happens to its escape velocity?
- It quadruples
+ It doubles
- It halves
= The escape velocity is √(2GM/R). Dividing R by 4 multiplies the quantity under the square root by 4, so the velocity grows by √4 = 2.
? What does general relativity say about Michell's critical radius?
- It is wrong by a factor of two
- It does not appear at all
+ The same radius, 2GM/c², appears, though the physics behind it is different
= Einstein's theory produces exactly the same radius, now called the Schwarzschild radius. But light does not slow down as Michell assumed, and nothing inside that radius moves outward even a short way.
:::
