Newton's laws can, in principle, answer every question in mechanics. In practice, following every force on a body from moment to moment is often laborious, and sometimes impossible, because the forces are complicated or unknown. A child sliding down a twisting water slide is pushed by the slide at a different angle at every point; working out the motion force by force would be a long job. This chapter introduces two quantities, energy and momentum, that let you skip much of that detail and still get the answer. The way in is a quantity called **work**, which measures what a force accomplishes as the body it acts on moves.

## How do you measure what a force achieves?

Start with an everyday sense of effort. Lift a 1 kg bag through 1 m and you have achieved something. Lift it through 2 m, or lift two bags through 1 m, and it seems fair to say you have achieved twice as much. What matters is the force you exert and the distance through which you exert it. Engineers of the eighteenth century, trying to compare the output of steam engines with that of horses and water wheels, measured it exactly this way, in pounds raised through feet. Physics keeps the idea and makes it precise. For a constant force F that pushes a body through a distance d in the direction of the force, the **work** done by that force is W = Fd.

Often the force and the motion are not in the same direction. Pull a suitcase on wheels by its handle and your pull is angled upward while the suitcase rolls horizontally. Only part of your pull helps to move it along; the upward part just takes some weight off the wheels. So the definition counts only the component of the force along the direction of motion. If the force makes an angle θ with the displacement, that component is F cos θ, and:

::: equation Work done by a constant force
W = Fd\cos\theta = \vec{F} \cdot \vec{d}
:::

The second form uses a piece of vector notation called the **dot product**. The dot product of two vectors is a single number: multiply their lengths together, then multiply by the cosine of the angle between them. Equivalently, multiply their matching components and add the results, so a·b = a₁b₁ + a₂b₂. The dot product measures how far two vectors point the same way. It is largest when they are parallel, zero when they are at right angles, and negative when they point in roughly opposite directions. Work is the dot product of force and displacement, so it measures how much the force is pushing along the motion.

Work is a **scalar**, a plain number with no direction attached, even though it is built from two vectors. Its unit is the **joule** (J), named after James Joule: one joule is the work done by a force of one newton acting through one metre, so 1 J = 1 N·m. Lifting a small apple, weighing about one newton, through one metre takes about one joule.

## Can work be negative?

The cosine in the definition means work can be positive, negative or zero, and each case has a clear meaning.

When the angle between force and motion is less than 90°, the cosine is positive and so is the work. The force is helping the motion along. Gravity does positive work on a falling stone, because it pulls down while the stone moves down.

When the angle is more than 90°, the cosine is negative and so is the work. The force is opposing the motion. Friction on a sliding box points backward while the box moves forward, so friction does negative work on it. So do a car's brakes, and so does gravity on a ball thrown upward, for as long as the ball is rising.

When the angle is exactly 90°, the cosine is zero, and the force does no work at all, however strong it is. A force at right angles to the motion can change the direction of a body but cannot help or hinder its progress. Gravity does no work on a satellite in a circular orbit, because it always points to the centre while the satellite moves along the circle. The normal force from a level floor does no work on a rolling ball, because it points up while the ball moves sideways.

That last case shows that the physics meaning of work is narrower than the everyday one. Carry a heavy suitcase along a level corridor at a steady speed. Your hand pushes up on the suitcase to hold it, while the suitcase moves horizontally, so your force is at right angles to the motion and does no work on the suitcase. Hold a heavy weight perfectly still and you do no work on it either, since it moves no distance at all. Yet in both cases you get tired. The reason is inside your body: muscles holding a load are not locked rigid but contract and relax continuously, many times a second, and each contraction uses chemical energy that ends up as heat. Work in the physics sense is done on and within your muscle fibres, but none of it is done on the suitcase.

## What if the force changes along the way?

So far the force has been constant. Often it is not. A spring pulls back harder the further you stretch it; the pull of gravity weakens as a spacecraft climbs. Then F d no longer makes sense, because there is no single F.

The way through is the same idea used in the accumulation lesson. Split the path into many short pieces, each so short that the force hardly changes across it. On each piece, the work is very nearly the force there times the length of the piece, F Δx. Add up the contributions from all the pieces, then let the pieces shrink toward zero length. The approximation becomes exact, and the sum becomes an integral:

::: equation Work done by a varying force along a line
W = \int_{x_1}^{x_2} F(x)\,\mathrm{d}x
:::

This has a simple picture. Draw a graph of force against position. Each short piece of path contributes a thin rectangle whose height is the force and whose width is the length of the piece, so its area is the work done over that piece. Adding them up means the total work is the area under the force–position curve.

::: model area
Read the curve as force against position. Each rectangle is a short stretch of path over which the force is nearly constant, and its area is the work done there. As the rectangles narrow, their total area approaches the work done over the whole path.
:::

The standard example is a spring. Stretch an ideal spring a distance x beyond its natural length and it pulls back with a force proportional to the stretch, F = kx. The number k is the **spring constant**: the larger it is, the stiffer the spring and the harder it pulls for a given stretch. This proportional rule is **Hooke's law**. It holds well for most springs as long as they are not stretched so far that they deform permanently.

How much work does it take to stretch a spring? The force grows as the stretch grows, so the integral is needed.

::: steps Work to stretch a spring
$ W = \int_0^x F(x')\,\mathrm{d}x' = \int_0^x kx'\,\mathrm{d}x'
To stretch the spring slowly, you must pull with a force equal to the spring's own pull, kx', at every intermediate stretch x' along the way. (The prime on x' just distinguishes the stretch at some point during the process from the final stretch x.) The work is the integral of that force from zero stretch up to x, the area under the line F = kx'.
---
$ W = \left[\tfrac{1}{2}kx'^2\right]_0^x
The spring constant k is the same throughout, so it comes outside the integral. An antiderivative of x' is ½x'², because differentiating ½x'² gives back x'. The square brackets mean: evaluate this expression at the upper limit x and subtract its value at the lower limit 0.
---
$ W = \tfrac{1}{2}kx^2
At the upper limit the expression is ½kx²; at the lower limit it is zero. The same answer comes from geometry, without calculus: the force–position graph is a straight line through the origin, so the area under it is a triangle with base x and height kx, and half of base times height is ½kx².
:::

The square in the result has a plain meaning. Stretching a spring twice as far takes four times the work, because you pull through twice the distance and, on average, with twice the force. A spring with k = 200 N/m stretched by 0.10 m needs ½ × 200 × 0.10², which is 1 J; stretched by 0.20 m, it needs 4 J.

## What does the total work do?

When several forces act on a body, each does its own work, positive, negative or zero, and the total work done on the body is the sum of them. Because work is built from the dot product, which adds up the same way forces do, this total equals the work done by the net force. A box pushed across a floor has positive work done on it by the push, negative work by friction, and none by gravity or the normal force, which are both vertical. The total, push minus friction, is what matters for the box's motion.

What does that total work actually do to the body? The next lesson shows that it changes the body's speed, in a precise way, and that answer is what turns work into energy.

::: history 1676–1829 · London and Paris
Robert Hooke found his spring law in the 1660s, but in 1676 he published it only as a scrambled string of letters, "ceiiinosssttuv", at the end of a book on telescopes. Anagrams of this kind were a common way to claim credit for a discovery without giving it away to rivals. Two years later, in 1678, he revealed the answer: "ut tensio, sic vis", Latin for "as the extension, so the force". The idea of work as a defined quantity came much later, from engineers rather than physicists. As steam engines spread, the people who built and bought them needed a fair way to compare what different machines delivered. In 1829 the French engineer Gaspard-Gustave Coriolis, in a book on the effectiveness of machines, gave the name *travail*, "work", to force times distance and made it a central quantity of mechanics. The unit later honoured James Joule, whose experiments in the 1840s showed that work can be turned into heat at a fixed rate, the discovery at the heart of the conservation of energy.
:::

::: note
Work measures what a force achieves as its point of application moves: W = F·d, force times the distance moved along the force. Only the component of the force along the motion counts, so a force at right angles to the motion does no work, and a force opposing the motion does negative work.
:::

::: quiz
? You push a lawnmower 20 m with a force of 100 N directed 60° below the horizontal. How much work do you do on it?
- 2000 J
+ 1000 J
- 1730 J
= Only the component of your push along the motion does work, so W = 100 × 20 × cos 60° = 1000 J. The downward part of the push presses the mower into the ground and does no work.
? The Moon moves in a nearly circular orbit around the Earth. How much work does the Earth's gravity do on it in one orbit?
- A large positive amount
+ Essentially zero
- A large negative amount
= Gravity points toward the Earth, at right angles to the Moon's motion along a circle, and cos 90° = 0. So gravity changes the Moon's direction but does no work on it.
? Stretching a spring 3 cm takes 0.9 J. How much does stretching it 6 cm take?
- 1.8 J
+ 3.6 J
- 2.7 J
= The work is ½kx², so doubling the stretch multiplies it by four. You pull through twice the distance with, on average, twice the force.
:::
