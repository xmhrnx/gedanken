Pull a mass on a spring to one side and let go, and it moves back and forth at a steady rhythm. The same kind of motion turns up in a swinging pendulum, a vibrating guitar string, the quartz crystal in a watch and the atoms jiggling inside a solid. It has a name, **simple harmonic motion**, and it is worth studying carefully, because it is the starting point for everything else in this chapter: waves are simple harmonic motion passed along from place to place. The motion arises whenever the force pulling something back toward its resting position is proportional to how far it has been moved, and this lesson shows why that one condition produces a smooth, regular rhythm whose timing does not depend on how big the swing is.

## What pulls the mass back?

Any oscillation needs two ingredients. First, a resting position, called **equilibrium**, where the object would stay if left alone. Second, a force that pushes the object back toward equilibrium whenever it is moved away. Without the first there is nothing to oscillate about; without the second the object would simply drift off.

A spring provides both. At its natural length it exerts no force, so a mass attached to it sits still. Stretch it, and it pulls back; compress it, and it pushes back. Robert Hooke found that, for moderate stretches, the force is proportional to the distance x by which the spring is stretched or compressed:

::: equation Hooke's law
F = -kx
:::

The constant k is called the **spring constant** and is measured in newtons per metre. It says how stiff the spring is: a stiff spring has a large k and needs a large force to stretch it a little. The minus sign carries the important information about direction. It says that the force always points opposite to the displacement, back toward equilibrium: stretch the spring to the right (positive x) and it pulls left (negative F).

Hooke's law is not exact. It holds well for small stretches, but pull a spring too far and the force stops being proportional to the stretch, and eventually the spring deforms permanently. The law describes the region near equilibrium, and as we will see at the end of this lesson, that is exactly why it is so widely useful.

::: history 1660–1678 · London
Robert Hooke was curator of experiments at the newly founded Royal Society, a prolific and quarrelsome inventor who worked on everything from microscopes to the architecture of London after the Great Fire. One of his concerns was the balance spring, a coiled spring that could regulate a pocket watch the way a pendulum regulates a clock. In 1676 he announced the law behind it in an odd form: the anagram "ceiiinosssttuv", a common way at the time of claiming priority for a discovery without revealing it. Two years later he gave the solution, the Latin phrase *ut tensio, sic vis*, meaning "as the extension, so the force". By then he was locked in a bitter dispute with the Dutch scientist Christiaan Huygens, who had built a balance-spring watch in 1675 and had a strong claim to the invention.
:::

## Turning the force into a motion

Knowing the force is not the same as knowing the motion. To find how the mass actually moves, we combine Hooke's law with Newton's second law, F = ma, which links force to acceleration, and then look for a description of position over time that satisfies the result. Take a mass m on a spring lying on a frictionless surface, so that the spring's force is the only horizontal force acting.

::: steps Solving for the motion of a mass on a spring
$ ma = -kx \quad\Rightarrow\quad a = -\frac{k}{m}\,x
Newton's second law says the force on the mass equals its mass times its acceleration, and Hooke's law gives that force as −kx. Setting the two equal and dividing by m gives the acceleration. It is proportional to the displacement and always points the opposite way: the farther the mass is from equilibrium, the harder it is accelerated back.
---
$ \omega^2 = \frac{k}{m}, \qquad a = -\omega^2 x
The ratio k/m appears so often that it is worth giving it a name. We write it as ω² (omega squared), choosing a square because it will turn out that ω itself has a direct meaning. The problem now reads: find a position x(t) whose acceleration, its second derivative or rate of change of its rate of change, is always −ω² times the position itself.
---
$ x = A\cos(\omega t)
We guess a cosine, because a cosine is a function that, after two rounds of taking slopes, comes back as its own negative. A is the **amplitude**, the largest distance the mass moves from equilibrium. At t = 0 the cosine equals 1, so this guess describes a mass pulled out to x = A and released from rest.
---
$ v = -A\omega\sin(\omega t), \qquad a = -A\omega^2\cos(\omega t) = -\omega^2 x
Now check the guess. The velocity is the slope of position over time, and the slope of cos(ωt) is −ω sin(ωt). The acceleration is the slope of the velocity, and the slope of sin(ωt) is ω cos(ωt). The result is −Aω² cos(ωt), which is −ω² times the original x. The guess satisfies the equation, and it does so whatever the value of A.
---
$ T = \frac{2\pi}{\omega} = 2\pi\sqrt{\frac{m}{k}}
The cosine repeats itself every time its argument ωt increases by 2π, so one complete oscillation takes a time T with ωT = 2π. This T is the **period**. It also shows what ω means: it is the **angular frequency**, the number of radians the cosine advances per second. Replacing ω with √(k/m) gives the period in terms of the mass and the spring.
:::

The formula agrees with intuition about each ingredient. A heavier mass has more inertia, so the spring takes longer to turn it round: the period grows with m. A stiffer spring pulls harder for the same stretch, so it turns the mass round faster: the period shrinks as k grows. For a 1 kg mass on a spring with k = 100 N/m, ω = √100 = 10 radians per second and T = 2π/10 ≈ 0.63 s. The **frequency** f = 1/T, the number of complete oscillations per second, is measured in hertz (Hz); here it is about 1.6 Hz.

## Why the amplitude does not matter

One feature of the result is easy to miss: the amplitude A does not appear in the period. Pull the mass out 1 cm or 5 cm, and it takes exactly the same time to complete a swing. That may seem surprising, since the mass has farther to go on a big swing.

The reason lies in Hooke's law. Pull the spring twice as far and the force is twice as large at every corresponding point of the swing, so the mass is accelerated twice as hard and reaches twice the speed. It has twice as far to go, but it moves twice as fast, and the two effects cancel exactly. This property, called **isochronism** (from the Greek for "equal time"), is what makes harmonic oscillators good clocks. Any real oscillator loses energy gradually to friction, so its swings grow smaller, but if the period does not depend on the size of the swing, the clock keeps the same rhythm as the swings die away.

::: note
When the restoring force is proportional to displacement, the position traces a sine wave in time. The period, 2π√(m/k), depends only on the mass and the stiffness, not on the amplitude, which is why such oscillators keep good time.
:::

## A shadow of circular motion

There is a way to see why a cosine appears, without any calculus. Imagine a point moving steadily around a circle of radius A, turning through ω radians every second. Its position along a horizontal line through the circle's centre, its x-coordinate, is A cos(ωt). So if you shine a light from above and watch the point's shadow on a line below, the shadow moves back and forth exactly like the mass on the spring.

The match goes deeper than position. A point moving steadily around a circle has a centripetal acceleration of ω²A directed toward the centre, as the lessons on gravity showed. The horizontal part of that acceleration is −ω² times the point's x-coordinate. That is precisely the acceleration of the mass on the spring. So simple harmonic motion is exactly the projection, or shadow, of uniform circular motion onto a straight line. The picture also explains how the speed varies. When the circling point is at the top or bottom of the circle, it is moving parallel to the line, so its shadow moves at the full speed ωA: that corresponds to the mass passing through equilibrium. When the point is at the far left or right, it is moving straight up or down, and its shadow is momentarily still: that corresponds to the mass at the end of its swing, about to turn back.

## Where the energy goes

Energy gives another view of the same motion. A moving mass has kinetic energy ½mv². A stretched or compressed spring stores elastic potential energy ½kx², which is the work done in stretching it against its own force. As the mass oscillates, energy passes back and forth between these two forms. At the ends of the swing the mass is momentarily at rest, so all the energy is stored in the spring. As the mass passes through equilibrium the spring is at its natural length and stores nothing, so all the energy is kinetic, and the mass is moving at its fastest. With no friction, no energy is lost, and the total stays fixed at the value it had at the start, when the mass was held still at x = A:

::: equation Energy in simple harmonic motion
\tfrac{1}{2}mv^2 + \tfrac{1}{2}kx^2 = \tfrac{1}{2}kA^2
:::

This equation gives the speed at any position without solving for the motion over time. At equilibrium, x = 0, it gives ½mv² = ½kA², so the top speed is A√(k/m), which is ωA, the same answer as the circle picture.

## Why it is everywhere

Springs are not common in nature, yet simple harmonic motion is. The reason is that almost any system resting at a stable equilibrium behaves like a spring when it is disturbed only slightly. Picture the potential energy of the system as a landscape, with the equilibrium at the bottom of a valley. Near the very bottom, any smooth valley is shaped like a parabola, and a parabolic potential energy, ½kx², is exactly the potential energy of a spring. Its slope gives a force of −kx. So for small displacements, the restoring force is proportional to the displacement, whatever the system is made of.

That is why molecules vibrate harmonically, why bridges and buildings sway at definite periods, and why a pendulum, the subject of the next lesson, swings with a steady beat, as long as the motion stays small. For larger disturbances the valley stops looking like a parabola, the force stops being proportional to displacement, and the period starts to depend on the amplitude.

Real oscillators also lose energy, to friction or air resistance or internal heating, and this loss is called **damping**. A damped oscillator's amplitude shrinks gradually from swing to swing. As long as the damping is light, though, the period barely changes while the swings die away, which is exactly the property a good clock needs.

::: quiz
? A mass on a spring oscillates with a period of 2 s. You double the amplitude. The new period is
- 4 s
+ 2 s
- 2√2 s
= In simple harmonic motion the period, 2π√(m/k), does not depend on amplitude. The mass has twice as far to go but moves twice as fast.
? You replace the mass with one four times as heavy, keeping the same spring. The period
+ doubles
- quadruples
- halves
= The period is 2π√(m/k), so multiplying m by 4 multiplies the period by √4, which is 2.
? Where is the mass moving fastest?
- At the ends of its swing
+ Passing through the equilibrium position
- Halfway between equilibrium and the end
= At equilibrium the spring stores no energy, so all the energy is kinetic and the speed is greatest. At the ends all the energy is stored in the spring and the mass is momentarily at rest.
:::
