The two halves of this chapter meet in a pair of machines. A **motor** uses the force on a current in a magnetic field to turn electrical energy into motion. A **generator** uses induction to turn motion into electrical energy. They are the same device run in opposite directions, and between them they handle most of the world's electricity: generators produce nearly all of it, and motors use roughly half.

## The motor

A wire carrying current I across a field B feels a force F = BIL. Make the wire into a rectangular loop, free to rotate about an axis, in a uniform field. The two sides of the loop that cut across the field carry current in opposite directions, so they are pushed in opposite directions: one up, one down. Together the two forces make a **torque**, a turning effect.

::: equation Torque on a coil of N turns and area A
\tau = NIAB\sin\theta
:::

Here θ is the angle between the field and the line perpendicular to the coil. The torque is greatest when the coil's plane lies along the field and zero when the coil faces the field squarely. Left alone, the coil would turn until it faced the field and then stop.

To keep it turning, the current must reverse every half-turn. In a DC motor a **commutator** does this: a split metal ring on the axle, touched by sliding contacts called brushes, swaps the connections each time the coil passes the dead point. Real motors use many coils at different angles so the torque is smooth, and iron cores to strengthen the field.

::: history 1821 · London
In September 1821, prompted by Ørsted's discovery the year before, Michael Faraday built a device in which a wire hanging into a cup of mercury, with a magnet standing in the middle, rotated continuously around the magnet whenever current flowed. It was the first conversion of electrical energy into continuous motion, the ancestor of every motor, though it was a demonstration rather than a useful machine.
:::

## The generator

Now turn the coil yourself, at a steady angular speed ω (in radians per second), in the same field, with no battery. The flux through the coil changes continuously, so by Faraday's law an EMF appears.

::: steps EMF of a rotating coil
$ \theta = \omega t
At steady rotation the angle between the field and the coil's perpendicular grows at a constant rate.
---
$ \Phi = BA\cos(\omega t)
The flux through one turn is BA cos θ. It swings between +BA and −BA as the coil turns.
---
$ \varepsilon = -N\frac{\mathrm{d}\Phi}{\mathrm{d}t} = NBA\omega\sin(\omega t)
Faraday's law. The derivative of cos(ωt) is −ω sin(ωt), and the two minus signs cancel.
---
$ \varepsilon_{\max} = NBA\omega
The peak voltage. More turns, a stronger field, a larger coil or faster turning all raise it.
:::

The output is a sine wave: alternating voltage, produced naturally by rotation. The voltage is largest when the coil's plane lies along the field, which is when the flux is zero but changing fastest. A generator with sliding rings instead of a commutator delivers this AC directly. A coil of 100 turns, 10 cm square (A = 0.01 m²), in a 0.5 T field, turning 50 times a second (ω = 2π × 50 ≈ 314 rad/s), peaks at 100 × 0.5 × 0.01 × 314 ≈ 160 V.

::: history 1831–1832 · London and Paris
In October 1831, weeks after discovering induction, Faraday spun a copper disc between the poles of a magnet and drew a steady current from contacts at its axle and rim: the first generator. In 1832 Hippolyte Pixii in Paris built a hand-cranked machine that spun a magnet past coils, producing alternating current; a commutator was soon added to give DC. Generators became practical in the 1860s, when Werner von Siemens and others used electromagnets powered by the machine's own output in place of weak permanent magnets.
:::

## Back EMF and energy

A generator does not give electricity for free. When current flows from it, that current sits in the magnetic field and feels the motor force, which by Lenz's law opposes the rotation. The harder the electrical load, the harder the coil is to turn. Mechanical work in, electrical energy out.

The same applies in reverse. A spinning motor is also a generator: its rotating coils induce a **back EMF** opposing the supply voltage. At full speed the back EMF nearly cancels the supply, and only a small current flows. When a motor starts, or is jammed, there is no back EMF and the current is large, which is why lights can dim briefly when a big motor switches on, and why a stalled motor can burn out. Electric cars use the symmetry directly: in **regenerative braking**, the motor becomes a generator, turning the car's kinetic energy back into charge in the battery.

::: note
A motor and a generator are one machine. Current in a field makes a coil turn; turning a coil in a field makes a current. Each direction resists the other, which is how energy is conserved.
:::

## Powering civilization

In a power station, steam from burning fuel or nuclear fission, or falling water, or wind, turns a turbine attached to a generator. Large generators spin electromagnets inside fixed coils, turning at a rate locked to the grid frequency of 50 or 60 Hz. Nikola Tesla's AC induction motor of 1888, which needs no brushes or commutator, made alternating current useful for factories as well as lighting, and Westinghouse's AC plant at Niagara Falls in the 1890s showed hydroelectric power could be sent to a distant city. Edison's DC systems lost that contest, though DC has returned for some long-distance links.

::: quiz
? A generator coil is turned twice as fast. Its peak voltage
- stays the same
+ doubles
- quadruples
= The peak EMF is NBAω, proportional to the angular speed.
? Why does a generator become harder to turn when a load draws current from it?
- Friction increases with current
+ The current in the coil feels a magnetic force that opposes the rotation
- The magnetic field gets weaker
= Lenz's law: the induced current resists the motion causing it, so the mechanical work you do becomes the electrical energy delivered.
? When is a motor most likely to draw a dangerously large current?
- At full speed with no load
+ When it is stalled or just starting
- When the supply voltage is low
= With the coil not turning, there is no back EMF to oppose the supply, so the current is limited only by the coil's small resistance.
:::
