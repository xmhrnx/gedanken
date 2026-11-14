A figure skater spinning with arms outstretched pulls them in and suddenly whirls several times faster. No one pushed her. The explanation is a third conserved quantity, the rotational counterpart of momentum. It also explains why planets speed up near the Sun and why collapsed stars spin hundreds of times a second.

## Angular momentum of a point mass

Take a body of mass m moving with speed v around a point, at distance r. For motion on a circle, its **angular momentum** about the centre is

::: equation Angular momentum of a particle on a circle
L = mvr
:::

More generally, only the part of the velocity perpendicular to the radius counts: L = m v⊥ r. A body heading straight toward or away from the point has none. In full vector form L = r × p, using the cross product, a vector perpendicular to both r and p; its direction marks the axis of rotation by a right-hand rule (curl the fingers of the right hand the way the body circulates, and the thumb points along L).

## Spinning bodies

A rigid body spinning about an axis is a collection of particles, each moving on its own circle. If the body turns at angular velocity ω (radians per second), a particle at distance r from the axis has speed v = ωr, so its angular momentum is m(ωr)r = mr²ω. Add them all:

::: equation Angular momentum of a rigid body
L = I\omega, \qquad I = \sum m r^2
:::

I is the **moment of inertia**, the rotational counterpart of mass. It depends not just on how much mass there is but on how far from the axis it sits, and the distance counts squared. Mass far out contributes much more than mass near the axis.

What changes angular momentum is **torque**, the turning effect of a force: force times its perpendicular distance from the axis, τ = rF sin θ. In the same way that force is the rate of change of momentum, torque is the rate of change of angular momentum, τ = dL/dt. So if no outside torque acts, L stays constant.

## The skater

::: steps Pulling in the arms
$ I_1\omega_1 = I_2\omega_2
On smooth ice there is almost no external torque about the vertical axis, so angular momentum before equals angular momentum after.
---
$ \omega_2 = \frac{I_1}{I_2}\,\omega_1
Solve for the new angular velocity. Pulling the arms and a leg in to the axis can cut I by a factor of about three, so ω roughly triples: one turn per second becomes three.
---
$ K = \tfrac{1}{2}I\omega^2 = \frac{L^2}{2I}
Rotational kinetic energy follows from adding ½mv² over every particle, with v = ωr. Written in terms of L, which is fixed, it is inversely proportional to I.
:::

So when I drops by three, the kinetic energy triples. The extra comes from the skater's muscles, which must pull her arms inward against their tendency to fly outward along the tangent. Angular momentum is conserved; kinetic energy is not.

## Equal areas

Kepler found in 1609 that a line from the Sun to a planet sweeps out equal areas in equal times, so a planet moves fastest when nearest the Sun. Angular momentum explains why.

::: steps Kepler's second law from angular momentum
$ \mathrm{d}A = \tfrac{1}{2}\,r\,(v_{\perp}\,\mathrm{d}t)
In a short time dt the planet moves sideways a distance v⊥ dt. The swept region is a thin triangle with height r and that base, so its area is half their product.
---
$ \frac{\mathrm{d}A}{\mathrm{d}t} = \tfrac{1}{2}r v_{\perp} = \frac{L}{2m}
Divide by dt and use L = m v⊥ r.
---
$ \frac{\mathrm{d}A}{\mathrm{d}t} = \text{constant}
The Sun's gravity points straight at the Sun, along r, so it exerts no torque about the Sun. L is constant, and so is the rate of sweeping area.
:::

The argument works for any force directed toward a fixed centre, not just inverse-square gravity. Newton proved it this way as the very first proposition of the *Principia*.

::: model orbit
Choose a launch speed that gives an ellipse. The ball moves fastest at its closest approach to the Earth and slowest at its farthest point, keeping its angular momentum fixed.
:::

## Collapsing stars

The same principle operates on a huge scale. If the Sun, which turns about once every 25 days, were squeezed from its radius of 700 000 km down to 10 km while keeping its angular momentum, its moment of inertia would fall by a factor of about 5 × 10⁹, and it would spin in under a millisecond. Real collapsing stellar cores shed much of their angular momentum, but the neutron stars they leave behind, about 20 km across, spin anywhere from once every few seconds to hundreds of times a second; the fastest known turns 716 times a second.

::: history 1609–1746 · Prague, Cambridge, Basel and Paris
Kepler found the equal-area law in Mars's orbit and published it in *Astronomia nova* (1609), with no physical explanation beyond a vague force from the Sun. Newton derived it from central forces in 1687. The general principle, that a system free of outside torques keeps its total angular momentum, then called the "conservation of areas", was stated independently in 1746 by Leonhard Euler, Daniel Bernoulli and Patrick d'Arcy.
:::

In quantum mechanics angular momentum survives as a conserved quantity but comes only in multiples of a basic unit, ħ/2, a fact central to atoms and particle spin.

::: note
With no outside torque, angular momentum L = Iω is conserved. Bring mass closer to the axis and the spin speeds up.
:::

::: quiz
? A spinning skater extends her arms, doubling her moment of inertia. Her rate of spin
- doubles
+ halves
- stays the same
= L = Iω is conserved, so doubling I halves ω.
? A comet on a long elliptical orbit is ten times closer to the Sun at its nearest point than at its farthest. How does its speed at the nearest point compare with that at the farthest?
- The same
- Ten times slower
+ Ten times faster
= At those two points the velocity is perpendicular to the radius, so L = mvr. Constant L with r ten times smaller means v ten times larger.
? Why is a planet's angular momentum about the Sun constant?
- The planet moves in a circle
- Gravity is weak at large distances
+ The Sun's gravity points directly at the Sun, so it exerts no torque about it
= A force along the line to the centre has zero perpendicular distance from it, so its torque is zero and L cannot change.
:::
