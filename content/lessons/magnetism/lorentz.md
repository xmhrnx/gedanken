The previous lesson described magnetic fields and where they come from. This one asks what a magnetic field does to an electric charge, and the answer is unlike any other force in everyday physics. A magnetic field pushes on a charge only if the charge is moving, and it pushes in an odd direction: sideways, at right angles both to the field and to the motion. That single rule explains why charged particles circle and spiral in magnetic fields, how the northern and southern lights form, how particle accelerators steer their beams, and why an electric motor turns.

## What force does a magnetic field exert?

Gravity and the electric force both act along the line joining two objects, and both act on things at rest. Experiments with beams of charged particles and with current-carrying wires showed that the magnetic force is different on both counts. A charge sitting still in a magnetic field feels nothing at all. A moving charge feels a force that grows with its speed, depends on the angle between its motion and the field, and points neither along the motion nor along the field but at right angles to both. All of this is captured in one formula.

::: equation Magnetic force on a moving charge
\vec{F} = q\,\vec{v} \times \vec{B}, \qquad F = qvB\sin\theta
:::

Here q is the charge, v its velocity and B the magnetic field. The symbol × denotes the **cross product** of two vectors, an operation that produces a third vector at right angles to both of them. Its size is vB sin θ, where θ is the angle between the velocity and the field. The second form of the equation gives just the size of the force.

The sin θ factor tells us how the force depends on direction. A charge moving straight along the field lines has θ = 0, and since sin 0 = 0, it feels no force at all, however fast it goes. A charge moving straight across the field has θ = 90°, where sin θ = 1, and feels the largest possible force, F = qvB. And of course a charge at rest has v = 0 and feels nothing.

To find which of the two possible perpendicular directions the force takes, there is a **right-hand rule**. Point the fingers of your right hand along the velocity, curl them toward the direction of the field, and your thumb points along v × B, which is the direction of the force on a positive charge. On a negative charge such as an electron, the force points the opposite way.

This formula is also how the unit of magnetic field is defined. A field of one tesla is one that exerts a force of one newton on a charge of one coulomb moving straight across it at one metre per second. When a charge feels both an electric field E and a magnetic field B, the two forces simply add, giving a total force F = q(E + v × B). This combined expression is called the **Lorentz force**, after the Dutch physicist Hendrik Lorentz, who in the 1890s built a theory of matter in which electrons respond to fields in exactly this way.

## Can a magnetic field speed a particle up?

Here is the first consequence of the sideways direction. A force changes a particle's kinetic energy only by doing work on it, and a force does work only if part of it points along the direction of motion: pushing a cart forward speeds it up, but pushing straight down on it does not. The magnetic force always points at right angles to the velocity, so it never has any part along the motion. It therefore does no work, and the particle's speed, and so its kinetic energy, never changes.

What the magnetic force can do is change the particle's direction. A magnetic field can steer a beam of particles but cannot accelerate it. That is why particle accelerators use magnets to bend their beams round a ring and electric fields to give them energy.

## Why do charged particles move in circles?

Now think about a particle moving straight across a uniform magnetic field. The force on it has a constant size, qvB, because neither its speed nor the field changes. And the force always points at right angles to the motion, turning the particle's direction without speeding it up. A force of constant size that always points sideways to the motion is exactly what keeps an object going round a circle, like the pull of a string on a whirled stone. So the particle travels in a circle, and we can work out how big the circle is and how long one lap takes.

::: steps Radius and period of the circular orbit
$ qvB = \frac{mv^2}{r}
Any object of mass m moving at speed v round a circle of radius r needs a force of mv²/r directed toward the centre to keep it on that path. Here the only force acting is the magnetic one, which always points toward the centre of the circle, so it must supply exactly this amount.
---
$ r = \frac{mv}{qB}
Divide both sides by v, which cancels one factor of v, and then rearrange to put r on its own by multiplying both sides by r and dividing by qB. The result says faster or heavier particles make larger circles, because they are harder to turn, while a larger charge or a stronger field makes the circle tighter.
---
$ T = \frac{2\pi r}{v}
To find the time T for one lap, divide the distance round the circle, its circumference 2πr, by the speed at which the particle travels it.
---
$ T = \frac{2\pi}{v} \cdot \frac{mv}{qB} = \frac{2\pi m}{qB}
Replace r with mv/qB from the second step. The speed v now appears on the top and the bottom and cancels. The time for one lap depends only on the particle's mass and charge and on the field.
:::

That last result is genuinely surprising. A faster particle travels on a bigger circle, but it covers the extra distance at exactly the rate needed to finish the lap in the same time. Double the speed and the radius doubles too, so the time stays the same. This holds until the particle approaches the speed of light, when relativity changes the details.

The American physicist Ernest Lawrence saw how to exploit this in 1931, when he built the first **cyclotron**. Particles circle inside two hollow D-shaped metal chambers placed back to back with a gap between them. An alternating voltage across the gap gives the particles a small kick each time they cross it. Because the lap time does not depend on speed, the voltage can flip at one fixed rate and still be timed right for every particle at every stage. With each kick the particles move faster and swing out onto a wider circle, spiralling outward until they emerge at high energy.

To get a feel for the sizes, take an electron moving at 10⁶ m/s straight across the Earth's magnetic field, about 5 × 10⁻⁵ T. Its radius is r = mv/qB = (9.1 × 10⁻³¹ × 10⁶) ÷ (1.6 × 10⁻¹⁹ × 5 × 10⁻⁵) ≈ 0.11 m, a circle about the size of a dinner plate. A proton at the same speed has the same size of charge but is 1,836 times heavier, so its circle is 1,836 times wider, with a radius of about 200 m. The weak field of the Earth is enough to bend fast electrons sharply.

## What if the particle moves partly along the field?

Most particles do not move exactly across the field. Any velocity can be split into two parts: one along the field lines and one across them. The part along the field feels no force, since for that part sin θ = 0, so it carries on unchanged. The part across the field feels the full sideways force and goes round in a circle, exactly as before. Put the two together and the particle moves steadily along the field while circling around it: a corkscrew path called a **helix**. In effect, charged particles are tied to magnetic field lines, free to slide along them but not to cross them.

This is what happens to charged particles near the Earth. Many are trapped by the Earth's field, spiralling along field lines from one hemisphere to the other and bouncing back where the lines crowd together near the poles; these trapped particles form the **Van Allen belts**, two great doughnut-shaped regions around the planet. Others travel down the field lines where they dip into the atmosphere near the poles, collide with oxygen and nitrogen high up, and make them glow green and red: the aurora. The popular picture says these particles stream straight from the Sun down the field lines. The real process is less direct: the solar wind stretches the Earth's field into a long tail on the night side, particles are stored and accelerated there, and it is these particles that then rain down along the field lines into the polar skies.

::: history 1897 · Cambridge, England
The magnetic force was the key to one of the great discoveries of physics. J. J. Thomson sent a beam of cathode rays, the mysterious glow-producing rays in evacuated glass tubes, through electric and magnetic fields arranged to push in opposite directions. He adjusted the fields until the beam went straight, so that the electric force qE balanced the magnetic force qvB, which told him the speed of the particles, v = E/B. Then, from how tightly the magnetic field alone bent the beam, using r = mv/(qB), he worked out the ratio of their charge to their mass. It was more than a thousand times larger than for a charged hydrogen atom, the lightest charged object then known; the modern figure is about 1,836 times. The rays were made of a new, extremely light particle: the electron. The same principle, sorting charged particles by how sharply a magnetic field bends them, is the basis of the **mass spectrometer**, which chemists use to weigh molecules.
:::

## What is the force on a wire?

A current in a wire is a great many charges moving together, so a wire carrying a current across a magnetic field must feel a force too: the sum of the tiny sideways pushes on all of its moving charges. We can add them up.

::: steps Force on a current-carrying wire
$ N = n A L
Take a straight length L of wire, with cross-sectional area A, lying at right angles to a field B. If the wire contains n moving charge carriers per cubic metre, then this length contains nAL of them, since AL is its volume.
---
$ F = (nAL)\,q v B
Each carrier, with charge q, drifts along the wire at speed v, straight across the field, so each feels a sideways force qvB. All the carriers are pushed the same way, so the total force is the number of carriers times the force on each.
---
$ I = n q A v
From the lesson on current, the current in a wire is related to its carriers by I = nqAv: the number per cubic metre, times the charge on each, times the area, times the drift speed.
---
$ F = BIL
Regroup the force from the second step as (nqAv) × L × B and recognise the bracket as the current I. The force depends only on things easy to measure: the field, the current and the length of wire in the field. This is the force that drives every electric motor and loudspeaker.
:::

::: note
A magnetic field pushes a moving charge at right angles to both its velocity and the field, with force qvB sin θ. Because the push is always sideways, it changes the particle's direction but never its speed, so a charged particle circles across the field and corkscrews along it.
:::

::: quiz
? A proton moves parallel to a magnetic field. What is the magnetic force on it?
+ Zero
- A force along the field
- A force at right angles to the field
= The size of the force is qvB sin θ, and for motion along the field θ = 0, so sin θ = 0. Only motion across the field lines produces a magnetic force.
? A charged particle circling in a uniform magnetic field is given twice the speed. What happens to the time it takes for one orbit?
- It doubles
- It halves
+ It stays the same
= The period is T = 2πm/(qB), which does not contain the speed. The radius doubles, but the particle also moves twice as fast, so it completes the larger circle in the same time.
? Why can a magnetic field not increase a charged particle's kinetic energy?
- Magnetic fields are too weak
+ The force is always perpendicular to the velocity, so it does no work
- The force only acts on stationary charges
= A force changes kinetic energy only by doing work, which requires part of the force to point along the motion. The magnetic force always points sideways, so it can turn the particle but never speed it up or slow it down.
:::
