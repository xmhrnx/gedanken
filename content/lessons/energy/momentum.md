Fire a rifle and it slams back into your shoulder. Nothing pushed the rifle backward except the bullet it launched. This lesson introduces the quantity that explains the kick, and that is conserved in every collision, explosion and push.

## Momentum and impulse

The **momentum** of a body is its mass times its velocity:

::: equation Momentum
\vec{p} = m\vec{v}
:::

It is a vector, pointing along the velocity, with units kg·m/s. A slow truck and a fast bullet can have the same momentum.

Newton's own second law was stated in terms of momentum. For constant mass, ma = m dv/dt = d(mv)/dt, so the second law can be written F = dp/dt: net force is the rate of change of momentum. Multiply by a time interval and the effect of a force acting over time Δt is a change in momentum:

::: equation Impulse
\vec{F}\,\Delta t = \Delta\vec{p}
:::

The product FΔt (or, for a varying force, the area under a force–time graph) is called the **impulse**. It explains every device that cushions a blow. A car occupant of 70 kg moving at 15 m/s must lose 1050 kg·m/s of momentum in a crash, whatever happens. Stopped by an airbag over about 0.1 s, the average force is about 10 kN. Stopped by a dashboard in 0.01 s, it is ten times larger. Bending your knees on landing, or drawing back your hands to catch an egg, works the same way: same change of momentum, longer time, smaller force.

## Conservation of momentum

Now two bodies interact, by a collision, a spring, a rope or anything else, with no outside forces acting.

::: steps Momentum is conserved in an isolated system
$ \vec{F}_{12} = -\vec{F}_{21}
The force on body 1 from body 2 is equal and opposite to the force on body 2 from body 1. This is Newton's third law.
---
$ \frac{\mathrm{d}\vec{p}_1}{\mathrm{d}t} = \vec{F}_{12}, \qquad \frac{\mathrm{d}\vec{p}_2}{\mathrm{d}t} = \vec{F}_{21}
The second law in momentum form, applied to each body. These are the only forces acting.
---
$ \frac{\mathrm{d}}{\mathrm{d}t}\left(\vec{p}_1 + \vec{p}_2\right) = \vec{F}_{12} + \vec{F}_{21} = 0
Add the two equations. The forces cancel by the first step.
---
$ \vec{p}_1 + \vec{p}_2 = \text{constant}
Zero rate of change means the total never changes, however complicated or violent the forces between the bodies.
:::

The argument extends to any number of bodies, since internal forces always cancel in pairs. The total momentum of an isolated system is constant.

## The rifle

Before firing, rifle and bullet are at rest, so the total momentum is zero. It must still be zero afterwards. Take a 10 g bullet leaving at 900 m/s from a 4.0 kg rifle:

::: steps Recoil velocity
$ 0 = m v + M V
Total momentum after firing equals total before, which was zero. m, v are the bullet's mass and velocity; M, V the rifle's.
---
$ V = -\frac{m v}{M} = -\frac{0.010 \times 900}{4.0} \approx -2.3\,\text{m/s}
The minus sign means the rifle moves backward.
:::

The momenta are equal and opposite, but the energies are not. The bullet carries ½ × 0.010 × 900² ≈ 4000 J; the rifle about 10 J. Kinetic energy is ½mv² = p²/(2m), so for equal momentum the lighter body gets nearly all the energy. That is why the bullet is lethal and the kick merely bruises.

## Collisions

Momentum is conserved in every collision. Kinetic energy is not, in general, since some goes into heat, sound and crumpled metal.

- In an **elastic** collision kinetic energy is also conserved. Steel balls and billiard balls come close. When two equal masses collide head-on elastically, they exchange velocities: the moving ball stops dead and the struck one moves off at the original speed, as in Newton's cradle.
- In a **perfectly inelastic** collision the bodies stick together, and the most kinetic energy consistent with momentum conservation is lost. A 1000 kg car at 20 m/s hitting an identical stationary car and locking bumpers: the pair moves off at 10 m/s, conserving momentum, but the kinetic energy drops from 200 kJ to 100 kJ. Half went into damage and heat.

Momentum conservation holds far beyond mechanics. Light carries momentum, and in relativity and quantum mechanics the law survives intact, while the formula mv does not.

::: history 1644–1669 · Amsterdam and London
Descartes proposed that God keeps the total "quantity of motion", mass times speed, constant in the universe. He ignored direction, and several of his rules for collisions were plainly wrong. In 1668 the Royal Society invited papers on the laws of collision. John Wallis treated inelastic collisions; Christopher Wren and Christiaan Huygens independently treated elastic ones. Huygens showed that the conserved quantity is mass times velocity counted with its sign, positive one way and negative the other, and that in elastic collisions the sum of mv² is conserved as well: momentum and energy, side by side.
:::

::: note
Momentum p = mv is conserved in any isolated system, because internal forces come in equal and opposite pairs. Kinetic energy is conserved only in elastic collisions.
:::

::: quiz
? A skater of 60 kg standing on ice throws a 3 kg ball forward at 10 m/s. What is her recoil speed?
+ 0.5 m/s
- 10 m/s
- 2 m/s
= Total momentum stays zero: 3 × 10 = 60 × V, so V = 0.5 m/s backward.
? Why does an airbag reduce injury?
- It reduces the passenger's change in momentum
+ It spreads the same change in momentum over a longer time, so the force is smaller
- It absorbs momentum so that total momentum is not conserved
= The passenger must stop either way, so Δp is fixed. Since FΔt = Δp, a longer stopping time means a smaller force.
? Two cars of equal mass collide and stick together. Which is conserved?
- Kinetic energy only
+ Momentum only
- Both momentum and kinetic energy
= Momentum is conserved in all collisions. In a sticking collision some kinetic energy always turns into heat and deformation.
:::
