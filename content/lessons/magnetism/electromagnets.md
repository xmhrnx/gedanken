For two centuries after Gilbert, electricity and magnetism were studied side by side without anyone finding a link between them. Static charges ignored magnets; magnets ignored static charges. The link appeared only once Volta's pile made steady currents possible, and when it came, it took a few months to turn a single observation into a quantitative law.

## Ørsted's compass

::: history 1820 · Copenhagen
Hans Christian Ørsted, a professor who believed on philosophical grounds that the forces of nature were connected, placed a wire carrying current from a voltaic pile near a compass during a lecture in the spring of 1820. The needle moved. Over the following months he found that it swung not toward or away from the wire but *across* it, at right angles, and that reversing the current reversed the swing. He announced the result in a four-page Latin pamphlet in July 1820. Within weeks it was being repeated across Europe.
:::

The direction was the surprise. Every force known until then, gravitational, electric or magnetic, pulled along the line between two objects. Ørsted's force went round. The magnetic field lines around a straight wire are circles centred on the wire. A **right-hand rule** gives their direction: grip the wire with your right thumb along the current, and your fingers curl the way the field points.

::: history 1820 · Paris
André-Marie Ampère heard of Ørsted's result at the French Academy in September 1820 and within a few weeks showed something new: two parallel wires carrying currents in the same direction attract each other, and in opposite directions repel. No magnets were involved at all. Ampère proposed that all magnetism is caused by electric currents, including tiny circulating currents inside magnetized iron. Jean-Baptiste Biot and Félix Savart measured the field around a wire in the same months and found it fell as 1/r.
:::

Ampère's guess about atomic currents was close to the truth, though the "currents" turned out to be electron spin and orbital motion, which only quantum mechanics describes properly.

## Ampère's law

Ampère's work led to a general rule linking a magnetic field to the current that causes it. Draw any closed loop in space. Go round it in small steps dl, and at each step multiply the step length by the part of B along it. The total, written ∮B·dl, depends only on the current passing through the loop.

::: equation Ampère's law
\oint \vec{B} \cdot \mathrm{d}\vec{l} = \mu_0 I_{\text{enclosed}}
:::

When the arrangement is symmetric, this gives fields almost for free.

::: steps Field of a long straight wire from Ampère's law
$ \oint \vec{B} \cdot \mathrm{d}\vec{l} = B \oint \mathrm{d}l
Choose a circular loop of radius r centred on the wire. By symmetry B has the same size all round the circle and points along it, so B comes out of the sum.
---
$ B \cdot 2\pi r = \mu_0 I
The loop's total length is its circumference, 2πr, and the current through it is the wire's current I.
---
$ B = \frac{\mu_0 I}{2\pi r}
Solve for B. This is the 1/r field Biot and Savart measured.
:::

## Coils and electromagnets

Bend a wire into a loop and the field lines from every part of it pass through the middle in the same direction, reinforcing one another. Wind many loops into a long coil, a **solenoid**, and the field inside becomes strong and nearly uniform, while outside it is weak and resembles that of a bar magnet. Applying Ampère's law to a rectangular loop half inside the coil gives:

::: equation Field inside a long solenoid
B = \mu_0 n I
:::

Here n is the number of turns per metre. A coil with 1000 turns per metre carrying 2 A gives B = 4π × 10⁻⁷ × 1000 × 2 ≈ 2.5 mT, about fifty times Earth's field.

Put an iron core inside and the field grows by a factor of hundreds or thousands, because the coil's field aligns the iron's magnetic domains, and their fields add to the coil's. That is an **electromagnet**: a magnet that can be switched on and off and whose strength is set by the current. William Sturgeon built the first in 1825, with an iron bar varnished for insulation. Joseph Henry, in Albany, New York, wound many turns of silk-insulated wire around the iron, and in 1831 built one at Yale that lifted about a tonne.

Electromagnets are everywhere: in scrapyard cranes, relays, door locks, loudspeakers, hard drives, and MRI scanners, whose superconducting coils carry hundreds of amperes with no resistance at all.

::: note
A current is surrounded by circular magnetic field lines, and Ampère's law relates the field round any loop to the current through it. Coiling the wire concentrates the field; an iron core multiplies it.
:::

## Defining the ampere

Ampère's parallel wires gave, for over 70 years, the official definition of the ampere: the current that, flowing in two long parallel wires 1 m apart, makes them attract with a force of 2 × 10⁻⁷ N per metre of length. That is why μ₀ was exactly 4π × 10⁻⁷ T m/A. Since 2019 the ampere is defined instead by fixing the elementary charge, and μ₀ is a measured quantity, equal to the old value to about one part in 10¹⁰.

::: quiz
? Ørsted's compass needle, placed under a wire running north to south, swung toward east or west. Why was this surprising?
- Compasses were not expected to respond to anything but the Earth
+ The force pointed across the wire, not along a line toward or away from it
- The needle should have pointed at the battery
= Every known force acted along the line between objects; the magnetic field of a current circles around the wire.
? Two parallel wires carry currents in opposite directions. They
- attract
+ repel
- feel no force
= Parallel currents attract and antiparallel currents repel, as Ampère found.
? A solenoid's current is doubled and its number of turns per metre is halved. The field inside
- doubles
+ stays the same
- halves
= B = μ₀nI; the two changes cancel.
:::
