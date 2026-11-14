Energy is the capacity to do work. A moving hammer can drive a nail; a raised weight can turn a clock as it descends; a stretched spring can launch a dart. This lesson defines the two forms that mechanics deals with: energy of motion and energy of position.

## Kinetic energy

What does work do to a body? Apply a constant net force F to a mass m over a distance d, and combine the definition of work with the constant-acceleration equations.

::: steps The work–energy theorem
$ W = Fd = mad
The work done by the net force, which points along the motion. Then use the second law, F = ma.
---
$ v^2 = u^2 + 2ad
The constant-acceleration equation linking speeds u and v at the start and end of the distance d.
---
$ ad = \tfrac{1}{2}\left(v^2 - u^2\right)
Rearrange it for the product ad, which appears in the work.
---
$ W = \tfrac{1}{2}mv^2 - \tfrac{1}{2}mu^2
Substitute. The net work equals the change in the quantity ½mv².
:::

That quantity is the **kinetic energy**:

::: equation Kinetic energy
K = \tfrac{1}{2}mv^2
:::

The theorem holds for varying forces too: splitting the path into short pieces and adding them up gives the same result. Kinetic energy is measured in joules, like work, and depends on speed but not direction, so it is never negative.

Because it goes as v², speed dominates. A 1000 kg car at 20 m/s (72 km/h) carries ½ × 1000 × 400 = 200 kJ. At 40 m/s it carries 800 kJ, and the brakes must remove four times as much energy, which is why braking distance grows as the square of speed.

## Potential energy

Lift a mass m through height h at steady speed. You push up with force mg, so you do work mgh. That work has not vanished: let the mass fall back and gravity does work mgh on it, returning that much kinetic energy. The work is stored in the arrangement of mass and Earth. This stored energy is **potential energy**:

::: equation Gravitational potential energy near the Earth's surface
U = mgh
:::

To match the car's 200 kJ, you would have to lift it h = 200 000 / (1000 × 9.8) ≈ 20 m, a seven-storey drop. A crash at 72 km/h delivers about the same energy as driving off the roof of such a building.

Only *changes* in potential energy matter, so the zero level is a free choice: floor, table top or sea level, whichever is convenient. The formula mgh is valid only near the surface, where g is constant. For large distances the general form is U = −GMm/r, which the dark stars lesson used to find escape velocity; its zero is at infinite separation, so it is negative everywhere else.

A stretched spring stores potential energy too. The work to stretch it, found in the last lesson, is recovered when it is released:

::: equation Spring potential energy
U = \tfrac{1}{2}kx^2
:::

## Which forces have a potential energy?

Lifting a book to a shelf takes the same work mgh whether you raise it straight up or carry it by a winding staircase; only the change in height matters. Forces whose work depends only on the start and end points, not the path, are called **conservative**. Gravity and ideal springs are conservative, and each has a potential energy defined by U = −(work done by the force), so that the energy lost by one form is gained by the other.

Friction is not conservative. Push a box across the floor and back, and friction opposes you both ways: the work done against it depends on the length of the path, and you get none of it back as motion. There is no "friction potential energy". The energy does not disappear, though; it turns into heat, as the next lesson describes.

::: history 1686–1740 · Leipzig and Paris
Descartes held that the true measure of a body's "force" was its quantity of motion, mass times speed. In 1686 Gottfried Leibniz argued it was mass times speed squared, which he called *vis viva*, "living force". The dispute ran for half a century. Willem 's Gravesande dropped brass balls into soft clay in the 1720s and found that a ball dropped from four times the height, arriving at twice the speed, made a dent four times as deep. Émilie du Châtelet used that result in her *Institutions de physique* (1740) to back Leibniz. Both sides were partly right: mv measures momentum and mv² measures energy, and each is conserved in its own way. Coriolis added the factor ½ in 1829 so that the quantity would equal the work done.
:::

::: note
Net work changes kinetic energy: W = ΔK, with K = ½mv². Work done against a conservative force like gravity is stored as potential energy and can be fully recovered.
:::

::: quiz
? A cyclist doubles her speed. Her kinetic energy
- doubles
+ quadruples
- stays the same, since mass is unchanged
= K = ½mv², so doubling v multiplies K by 4.
? A 2 kg book is lifted from the floor to a 1.5 m shelf, once straight up and once by a curving path. How do the changes in potential energy compare?
- The curving path gives more, as it is longer
+ They are equal, about 29 J each
- The straight path gives more
= Gravity is conservative: ΔU = mgh = 2 × 9.8 × 1.5 ≈ 29 J whatever the route.
? Why can friction not be given a potential energy?
- It is too weak
+ The work it does depends on the path, not just the endpoints
- It acts only on moving bodies
= A potential energy needs work that depends only on where a body starts and ends. Friction always opposes motion, so a longer path means more work against it.
:::
