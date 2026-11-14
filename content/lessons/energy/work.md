Newton's laws answer every question in mechanics in principle, but following forces moment by moment is often laborious. This chapter introduces quantities that let you skip the details: energy and momentum. The way in is **work**, which measures what a force accomplishes as its point of application moves.

## Force times distance

Lift a 1 kg bag 1 m and you have done something; lift it 2 m, or lift two bags, and you have done twice as much. Engineers of the 1700s measured the output of steam engines this way, in pounds raised through feet. Physics makes it precise. For a constant force F pushing a body a distance d along the force's direction, the work done is W = Fd.

If the force is at an angle θ to the displacement, only its component along the motion counts:

::: equation Work done by a constant force
W = Fd\cos\theta = \vec{F} \cdot \vec{d}
:::

The second form uses the **dot product** of two vectors: multiply their lengths and the cosine of the angle between them, or equivalently multiply matching components and add, a·b = a₁b₁ + a₂b₂. It measures how much two vectors point the same way.

Work is a scalar, a single number with no direction. Its unit is the **joule** (J): 1 J = 1 N·m, about the work done lifting a small apple 1 m.

## Positive, negative and zero

The sign of cos θ gives the sign of the work.

- **Positive** (θ < 90°): the force helps the motion. Gravity does positive work on a falling stone.
- **Negative** (θ > 90°): the force opposes it. Friction on a sliding box does negative work, as do brakes.
- **Zero** (θ = 90°): the force is perpendicular to the motion. Gravity does no work on a satellite in a circular orbit, and the normal force from a level floor does no work on a rolling ball.

The physics meaning differs from the everyday one. Carry a heavy suitcase along a level corridor at steady speed: your upward force is perpendicular to the motion, so you do no work on the suitcase. You still get tired, because muscles holding a load are constantly contracting and relaxing, using chemical energy internally. And holding a weight still does no work at all on the weight, since d = 0.

## Forces that vary

Often the force changes as the body moves. Then split the path into short pieces, each so small that the force is nearly constant over it, add F·Δx over all the pieces, and let them shrink. That is exactly the integral from the accumulation lesson:

::: equation Work done by a varying force along a line
W = \int_{x_1}^{x_2} F(x)\,\mathrm{d}x
:::

On a graph of force against position, the work is the area under the curve.

::: model area
Read the curve as force against position. The rectangles are small stretches of path at nearly constant force; their total area approaches the work done.
:::

The standard example is a spring. Stretch an ideal spring a distance x from its natural length and it pulls back with a force proportional to the stretch, F = kx, where the *spring constant* k says how stiff it is. This is **Hooke's law**, good for small stretches.

::: steps Work to stretch a spring
$ W = \int_0^x F(x')\,\mathrm{d}x' = \int_0^x kx'\,\mathrm{d}x'
You must pull with a force equal to the spring's, kx', at each stretch x' along the way, from 0 up to the final stretch x.
---
$ W = \left[\tfrac{1}{2}kx'^2\right]_0^x = \tfrac{1}{2}kx^2
The antiderivative of x' is ½x'². Evaluate at the limits. Geometrically, the force–position graph is a straight line and this is the area of the triangle under it: half of base x times height kx.
:::

Doubling the stretch takes four times the work: twice the distance, at twice the average force. A spring with k = 200 N/m stretched by 0.10 m needs ½ × 200 × 0.01 = 1 J.

## Work by several forces

When several forces act, each does its own work, and the total work is their sum, which equals the work done by the net force. The next lesson shows what that total does: it changes the body's speed, in a precise way that defines kinetic energy.

::: history 1678 and 1829 · London and Paris
Robert Hooke published his spring law in 1678 as a Latin anagram, "ut tensio, sic vis" ("as the extension, so the force"), having first announced it in 1676 scrambled as ceiiinosssttuv so that he could claim priority without revealing it. The word *work* as a defined quantity came much later: Gaspard-Gustave Coriolis, an engineer studying the efficiency of machines, introduced *travail* for force times distance in 1829. The unit honours James Joule, whose experiments in the 1840s showed that work and heat are interchangeable.
:::

::: note
Work is force times distance moved along the force: W = F·d. Only the component of force along the motion counts, so perpendicular forces do no work.
:::

::: quiz
? You push a lawnmower 20 m with a force of 100 N directed 60° below the horizontal. How much work do you do on it?
- 2000 J
+ 1000 J
- 1730 J
= Only the horizontal component does work: W = 100 × 20 × cos 60° = 1000 J.
? The Moon moves in a nearly circular orbit around the Earth. How much work does the Earth's gravity do on it in one orbit?
- A large positive amount
+ Essentially zero
- A large negative amount
= Gravity points toward the Earth, perpendicular to the Moon's motion on a circle, so cos 90° = 0 and no work is done.
? Stretching a spring 3 cm takes 0.9 J. How much does stretching it 6 cm take?
- 1.8 J
+ 3.6 J
- 2.7 J
= W = ½kx², so doubling x multiplies the work by 4.
:::
