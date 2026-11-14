A magnetic field pushes on a charge only if the charge is moving, and it pushes in a strange direction: sideways, at right angles both to the field and to the motion. This one rule explains why charged particles spiral in magnetic fields, how the aurora forms, how particle accelerators steer their beams, and why an electric motor turns.

## The magnetic force

A charge q moving with velocity v through a magnetic field B feels a force

::: equation Magnetic force on a moving charge
\vec{F} = q\,\vec{v} \times \vec{B}, \qquad F = qvB\sin\theta
:::

The **cross product** v × B is a vector perpendicular to both v and B, with size vB sin θ, where θ is the angle between them. So:

- A charge at rest feels no magnetic force.
- A charge moving along the field lines (θ = 0) feels none either.
- The force is largest for motion straight across the field (θ = 90°), where F = qvB.

The direction comes from the **right-hand rule**: point your fingers along v, curl them toward B, and your thumb gives the direction of v × B. For a negative charge the force is reversed. This is how the tesla is defined: a field of 1 T exerts 1 N on a charge of 1 C moving across it at 1 m/s.

Combined with the electric force, the total force on a charge is F = q(E + v × B), named after Hendrik Lorentz, who set out the theory of electrons in fields in the 1890s.

## Magnetic forces do no work

Because the force is always perpendicular to the velocity, it can change a particle's direction but never its speed. It does no work, and the particle's kinetic energy stays constant. A magnetic field can steer a beam but cannot speed it up; accelerators use electric fields for that.

## Circular orbits

A force of constant size, always at right angles to the motion, is exactly what makes something move in a circle. So a charged particle moving across a uniform magnetic field travels in a circle.

::: steps Radius and period of the circular orbit
$ qvB = \frac{mv^2}{r}
The magnetic force supplies the centripetal force needed to keep mass m moving at speed v in a circle of radius r.
---
$ r = \frac{mv}{qB}
Cancel one v and solve for r. Faster or heavier particles make bigger circles; stronger fields make tighter ones.
---
$ T = \frac{2\pi r}{v} = \frac{2\pi m}{qB}
The time for one orbit is the circumference divided by the speed. The speed cancels.
:::

The last result is surprising: the orbit time does not depend on speed. A faster particle travels a bigger circle in exactly the same time (until it nears the speed of light, when relativity changes things). Ernest Lawrence exploited this in 1931 in the **cyclotron**: particles circle between two hollow D-shaped electrodes, and an alternating voltage, timed to the fixed orbital period, kicks them each half-turn. They spiral outward to high energy.

For a sense of scale, an electron moving at 10⁶ m/s across the Earth's field of 5 × 10⁻⁵ T circles with radius r = (9.1 × 10⁻³¹ × 10⁶)/(1.6 × 10⁻¹⁹ × 5 × 10⁻⁵) ≈ 0.11 m. A proton at the same speed, 1836 times heavier, circles with a radius of about 200 m.

## Spirals

If the particle's velocity has a part along the field as well as across it, the along-field part is untouched (sin 0 = 0), while the across-field part goes round in a circle. The combination is a **helix**: the particle corkscrews along the field lines.

This is what happens to charged particles streaming from the Sun. Earth's field traps many of them, spiralling along field lines and bouncing between the poles, in the **Van Allen belts**. Others are channelled down where the field lines converge near the poles, strike the upper atmosphere and make it glow: the aurora.

::: history 1897 · Cambridge, England
J. J. Thomson passed cathode rays through crossed electric and magnetic fields. He adjusted them until the beam went straight, so that qE = qvB, which gave the speed v = E/B. Then, from how much the magnetic field alone bent the beam, using r = mv/(qB), he found the charge-to-mass ratio q/m. It was more than a thousand times that of a hydrogen ion (the modern figure is about 1836), evidence for a new, very light particle: the electron. The same principle, sorting ions by how much a field bends them, is used in every **mass spectrometer**.
:::

## Force on a wire

A current in a wire is many charges moving together, so a wire carrying current across a magnetic field feels a sideways force. Adding up qvB over all the moving charges in a length L of wire gives F = BIL (for a wire at right angles to the field). This is the force that drives every electric motor and loudspeaker.

::: note
A magnetic field pushes a moving charge at right angles to both its velocity and the field. It changes direction but not speed, so a charged particle circles across the field and spirals along it.
:::

::: quiz
? A proton moves parallel to a magnetic field. The magnetic force on it is
+ zero
- along the field
- at right angles to the field
= F = qvB sin θ, and θ = 0 for motion along the field.
? A charged particle circling in a uniform magnetic field is given twice the speed. Its orbit period
- doubles
- halves
+ stays the same
= T = 2πm/(qB) does not depend on speed: the radius doubles, but so does the speed.
? Why can a magnetic field not increase a charged particle's kinetic energy?
- Magnetic fields are too weak
+ The force is always perpendicular to the velocity, so it does no work
- The force only acts on stationary charges
= Work requires a force component along the motion; the magnetic force has none.
:::
