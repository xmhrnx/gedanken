A figure skater spinning slowly with her arms stretched out pulls them in tight to her body, and suddenly she is whirling several times faster. No one pushed her, and the ice is nearly frictionless, so nothing outside her has added any spin. The explanation is a third conserved quantity, the rotational counterpart of momentum, called **angular momentum**. It governs anything that turns or circles: it explains why a planet speeds up as it swings close to the Sun, why a spinning top stays upright, and why the collapsed remains of stars can spin hundreds of times a second. This lesson defines it, shows what changes it, and follows its conservation from the ice rink to the stars.

## How do you measure the amount of turning?

Momentum measures how much straight-line motion a body has. To describe motion around a point, we need a quantity that also takes into account how far from the point the motion happens. A small push at the edge of a door swings it open easily; the same push near the hinges hardly moves it. Distance from the centre clearly matters.

Take a body of mass m moving at speed v around a point, at a distance r from it. If it moves on a circle around that point, its angular momentum about the centre is:

::: equation Angular momentum of a particle on a circle
L = mvr
:::

This is the body's momentum, mv, multiplied by its distance from the centre. When the body is not moving on a circle, only the part of its velocity that is perpendicular to the radius counts, the part carrying it around the point rather than toward or away from it. So in general L = m v⊥ r, where v⊥ is that perpendicular part. A body heading straight at the point, or straight away from it, has no angular momentum about it at all, however fast it moves, because it is not going round.

Angular momentum also has a direction, which marks the axis the motion turns around. In full vector form it is written L = r × p, using the **cross product**, an operation on two vectors that gives a third vector perpendicular to both. Its direction follows a right-hand rule: curl the fingers of your right hand in the direction the body circulates, and your thumb points along L. For most of this lesson only the size matters, but the direction is what keeps a spinning top or a gyroscope pointing steadily.

## What about a spinning solid body?

A rigid body spinning about an axis, such as a wheel, a planet or a skater, can be treated as a collection of tiny particles, each moving on its own circle around the axis. They all turn together, so they share the same **angular velocity** ω (the Greek letter omega), the angle turned per second, measured in radians per second. But they do not share the same speed. A particle at distance r from the axis covers a circle of circumference 2πr in each turn, so its speed is v = ωr: particles far from the axis move faster than those near it.

Each particle's angular momentum is then m times its speed times its distance, m(ωr)r, which is mr²ω. Adding up over every particle in the body, the ω is common to all and comes out as a factor:

::: equation Angular momentum of a rigid body
L = I\omega, \qquad I = \sum m r^2
:::

The quantity I is called the **moment of inertia**. It plays the same part in rotation that mass plays in straight-line motion: it measures how hard it is to change a body's spin. But unlike mass, it depends on where the mass is, not just how much there is. Each bit of mass counts in proportion to the square of its distance from the axis. Mass twice as far out contributes four times as much. That is why a skater with outstretched arms has a much larger moment of inertia than the same skater with arms pulled in, though her mass has not changed at all.

What changes angular momentum? In straight-line motion, force changes momentum. In rotation, the corresponding quantity is **torque**, the turning effect of a force. A force has more turning effect the further from the axis it acts and the more nearly it pushes at right angles to the line from the axis; this is why door handles are placed far from the hinges. Its size is τ = rF sin θ, where r is the distance from the axis to where the force is applied and θ is the angle between that line and the force. Just as the net force on a body is the rate of change of its momentum, the net torque is the rate of change of its angular momentum, τ = dL/dt. And so, if no outside torque acts, angular momentum stays constant. That is the conservation law, and it has the same root as momentum conservation: the internal forces between the parts of a system come in equal and opposite pairs acting along the same line, so their torques cancel too.

## Why does the skater speed up?

Now the skater can be explained. While she spins on smooth ice, the forces on her are gravity, the push of the ice, and a tiny friction at the skate's tip. None of them exerts any appreciable torque about the vertical axis through her body, so her angular momentum about that axis is conserved. When she pulls her arms in, she does it with her own muscles, which are internal forces.

::: steps Pulling in the arms
$ I_1\omega_1 = I_2\omega_2
With no outside torque about the vertical axis, angular momentum is conserved. So her moment of inertia times her angular velocity before she pulls in her arms, labelled 1, equals the same product afterwards, labelled 2. Pulling her arms in changes I, and the equation says ω must change to compensate.
---
$ \omega_2 = \frac{I_1}{I_2}\,\omega_1
Divide both sides by I₂ to find the new angular velocity. The spin rate goes up by the same factor that the moment of inertia goes down. Bringing her arms and a raised leg in close to the axis can cut a skater's moment of inertia by a factor of about three, so her spin rate roughly triples: one turn per second becomes about three.
---
$ K = \tfrac{1}{2}I\omega^2
Now look at her kinetic energy. Each particle of her body has kinetic energy ½mv², with v = ωr. Adding these up gives ½(Σmr²)ω², and the sum in brackets is the moment of inertia. So rotational kinetic energy has the same form as ½mv², with I in place of m and ω in place of v.
---
$ K = \frac{L^2}{2I}
Since L = Iω, we can replace ω with L/I and simplify. Written this way, with L fixed by conservation, the kinetic energy is inversely proportional to the moment of inertia. When the skater cuts I by a factor of three, her kinetic energy rises by a factor of three.
:::

That last result may seem to break the conservation of energy, but it does not. The extra kinetic energy comes from the skater's muscles. To pull her arms inward, she has to supply the inward force that bends each hand's path into a tighter circle, against its tendency to carry on along the tangent, and in doing so her muscles do work. That work becomes the extra energy of the faster spin. When she spreads her arms again, her spin slows and the energy goes back into her body, as heat. Angular momentum is conserved throughout; kinetic energy is not, because internal forces are doing work.

## Why do planets sweep out equal areas?

Early in the seventeenth century, Johannes Kepler found from careful observations of Mars that a line from the Sun to a planet sweeps out equal areas in equal times. So a planet must move faster when it is close to the Sun, where a short line has to sweep out its share of area quickly, and slower when it is far away. Kepler had no explanation for this rule. Angular momentum provides one.

::: steps Kepler's second law from angular momentum
$ \mathrm{d}A = \tfrac{1}{2}\,r\,(v_{\perp}\,\mathrm{d}t)
In a short time dt, the planet moves a short way along its orbit. Its sideways movement, at right angles to the line from the Sun, is v⊥ dt. The region swept out by the line is then a thin triangle with its point at the Sun, a height of r and a base of v⊥ dt, and the area of a triangle is half its base times its height. Any motion directly toward or away from the Sun lengthens or shortens the triangle without adding area, which is why only v⊥ appears.
---
$ \frac{\mathrm{d}A}{\mathrm{d}t} = \tfrac{1}{2}r v_{\perp} = \frac{L}{2m}
Divide by dt to get the rate at which area is swept out. The result, ½ r v⊥, is almost the planet's angular momentum about the Sun, m v⊥ r, which differs only by the factor m. Dividing L by 2m makes the two equal. So the rate of sweeping area is just the angular momentum per unit mass, halved.
---
$ \frac{\mathrm{d}A}{\mathrm{d}t} = \text{constant}
The Sun's gravity on the planet always points straight at the Sun, along the line r. A force along that line has no turning effect about the Sun, because the angle θ in τ = rF sin θ is zero, so the torque is zero. With no torque, the planet's angular momentum about the Sun is constant, and therefore so is the rate at which it sweeps out area.
:::

The argument uses only one fact about gravity: that it points toward a fixed centre. It does not depend on the inverse-square law at all, so equal areas would hold for any force directed at the centre. Newton saw this, and the very first proposition of the *Principia* proves the same result with a geometric argument built on short successive pulls toward the centre.

::: model orbit
Choose a launch speed that gives an ellipse. The ball moves fastest at its closest approach to the Earth and slowest at its farthest point, so that its angular momentum stays the same all the way round.
:::

## What happens when a star collapses?

The same principle works on an enormous scale. The Sun turns on its axis about once every 25 days. Imagine squeezing it, from its radius of about 700,000 km down to just 10 km, while keeping its angular momentum fixed. Its moment of inertia depends on the square of its size, so it would fall by a factor of about (700,000 / 10)², roughly 5 × 10⁹. The spin rate would rise by the same factor, and one turn every 25 days would become one turn in less than a millisecond.

Nature does something like this. When a massive star exhausts its fuel, its core collapses in a fraction of a second into a **neutron star**, an object about 20 km across made mostly of neutrons. Real collapsing cores shed much of their angular momentum along the way, so they do not spin up as much as the simple calculation suggests. Even so, neutron stars spin anywhere from once every few seconds to hundreds of times a second. The fastest known turns 716 times a second.

::: history 1609–1746 · Prague, Cambridge, Basel and Paris
Kepler found the equal-area rule while working out the orbit of Mars, and published it in his *Astronomia nova*, "New Astronomy", in 1609. He tried to explain it with a force from the Sun that swept the planets round like the spokes of a turning wheel, an idea that did not survive. Newton derived the rule from the nature of central forces in 1687, showing that it follows from the direction of the force alone. The general principle took longer to emerge. That any system free of outside torques keeps its total angular momentum, then called the "conservation of areas", was stated independently in 1746 by three people: the mathematicians Leonhard Euler and Daniel Bernoulli, and the French engineer Patrick d'Arcy. Euler went on to build a full mechanics of rotating rigid bodies on it.
:::

Angular momentum keeps its importance in quantum mechanics, where it is still conserved but can no longer take any value. Any measured component of a particle's angular momentum is a whole-number multiple of ħ/2, where ħ is the reduced Planck constant. That restriction shapes the structure of atoms and gives particles such as the electron their intrinsic spin.

::: note
With no outside torque, angular momentum L = Iω is conserved. Bring mass closer to the axis and the moment of inertia drops, so the spin speeds up; the same principle makes planets move fastest when nearest the Sun.
:::

::: quiz
? A spinning skater extends her arms, doubling her moment of inertia. Her rate of spin
- doubles
+ halves
- stays the same
= Angular momentum L = Iω is conserved because no outside torque acts, so doubling I must halve ω.
? A comet on a long elliptical orbit is ten times closer to the Sun at its nearest point than at its farthest. How does its speed at the nearest point compare with that at the farthest?
- The same
- Ten times slower
+ Ten times faster
= At those two points the velocity is at right angles to the line from the Sun, so L = mvr. With L constant and r ten times smaller, v must be ten times larger.
? Why is a planet's angular momentum about the Sun constant?
- The planet moves in a circle
- Gravity is weak at large distances
+ The Sun's gravity points directly at the Sun, so it exerts no torque about it
= A force directed along the line to the centre has no turning effect about that centre, so its torque is zero and the angular momentum cannot change.
:::
