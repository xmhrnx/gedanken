The two halves of this chapter meet in a pair of machines. A **motor** uses the force on a current in a magnetic field to turn electrical energy into motion. A **generator** uses induction to turn motion into electrical energy. They look alike because they are alike: a motor and a generator are essentially the same device run in opposite directions. Between them they handle most of the world's electricity. Generators produce almost all of it, apart from the growing share that comes from solar panels, and motors, in everything from factory pumps to fridges, use roughly half. This lesson shows how each works, why each one resists the other, and why that resistance is exactly what keeps the books balanced on energy.

## How does a current make something turn?

The lesson on moving charges showed that a straight wire of length L, carrying current I at right angles to a field B, feels a sideways force F = BIL. A force on a straight wire pushes it along, but a motor needs something that turns. The trick is to bend the wire into a rectangular loop and mount it on an axle so that it can rotate, in a uniform magnetic field.

Follow the current round the loop. Along one long side it flows one way; along the opposite long side it flows back the other way. Both sides cut across the field, so both feel a force, but because their currents are opposite, so are their forces: one side is pushed up and the other down. Two equal and opposite forces acting on opposite sides of an axle do not move the loop anywhere, but they do twist it. That twisting effect is called a **torque**: the turning equivalent of a force, measured as a force multiplied by its distance from the axis.

::: steps Torque on a rectangular coil
$ F = N B I a
Take a rectangular coil of N turns with sides a and b, where the sides of length a lie at right angles to the field, and the axle runs through the middle, parallel to them. Each of these sides carries current I in N separate wires across the field, so the force on it is N times BIa, from F = BIL. The other two sides contribute no turning, because their forces push along the axle rather than round it.
---
$ \tau = 2 \times F \times \frac{b}{2}\sin\theta
Each force acts at a distance b/2 from the axle, but only the part of that distance at right angles to the force produces turning, which brings in a factor of sin θ. Here θ is the angle between the field and the line sticking straight out of the face of the coil. The two forces twist the coil in the same rotational direction, so their torques add, giving the factor of 2.
---
$ \tau = NBIab\sin\theta
Multiply out: the 2 and the half cancel, leaving N, B, I and the two side lengths.
---
$ \tau = NIAB\sin\theta
The product ab is just the area A of the coil. So the torque depends on the coil's area and number of turns, the current and the field, and not on the coil's particular shape; a more careful calculation shows this is true for a flat coil of any shape.
:::

The sin θ factor says the torque varies as the coil turns. It is greatest when θ = 90°, which is when the plane of the coil lies along the field, so that the field sweeps across its face. It falls to zero when θ = 0, when the coil faces the field squarely: then the two forces pull directly away from each other through the axle and produce no twist. So a coil left to itself would swing round until it faced the field, overshoot a little, swing back, and settle there. That is not yet a motor.

To keep the coil turning, the current in it must reverse every half-turn, just as it passes the dead position, so that the forces flip and keep pushing it the same way round. In a simple DC motor this is done by a **commutator**: a metal ring split into two halves, fixed to the axle and connected to the two ends of the coil. Two stationary contacts called **brushes** press against the ring and carry current in. As the coil turns past the dead point, each brush slides from one half of the ring to the other, swapping the connections and reversing the current in the coil. Real motors use several coils set at different angles, so that some coil is always near its position of greatest torque and the turning is smooth, and they wind their coils on iron cores to strengthen the field.

::: history 1821 · London
Within a year of Ørsted's discovery, everyone was trying to understand the strange circular force between a current and a magnet. In September 1821 Michael Faraday, then a young assistant at the Royal Institution, built a device to show it directly. He stood a bar magnet upright in a cup of mercury and hung a wire from above so that its lower end dipped into the mercury, free to move. When he passed a current through the wire and the mercury, the wire's lower end swept round and round the magnet for as long as the current flowed. It was the first conversion of electrical energy into continuous motion, the ancestor of every electric motor, though it was a demonstration of a principle rather than a useful machine.
:::

## How does turning make a current?

Now take away the battery and turn the same coil yourself, steadily, in the same field. The amount of field passing through the coil, its magnetic flux, now changes continuously as the coil turns, from a maximum when it faces the field, through zero when it is edge-on, to a maximum in the opposite sense when it has turned half-way round. By Faraday's law, a changing flux induces an EMF, so a voltage appears across the coil's ends. To find how large it is, describe the rotation by its **angular speed** ω (the Greek letter omega), the angle turned per second, measured in radians, where a full turn is 2π radians.

::: steps EMF of a rotating coil
$ \theta = \omega t
At steady rotation, the angle between the field and the line sticking straight out of the coil grows at a constant rate. If we start the clock when the coil faces the field squarely, the angle after a time t is just the angular speed times the time.
---
$ \Phi = BA\cos\theta = BA\cos(\omega t)
The flux through one turn of the coil is BA cos θ, from the definition of flux in the previous lesson. Writing in θ = ωt shows how it changes with time: it swings smoothly between +BA and −BA, once round each full turn.
---
$ \varepsilon = -N\frac{\mathrm{d}\Phi}{\mathrm{d}t}
Faraday's law says the EMF across a coil of N turns is minus N times the rate of change of the flux through each turn. So we need the rate at which BA cos(ωt) changes.
---
$ \varepsilon = -NBA \cdot \left(-\omega\sin(\omega t)\right) = NBA\omega\sin(\omega t)
The rate of change of cos(ωt) is −ω sin(ωt): the cosine curve turns into a sine curve, scaled up by ω because a faster rotation makes everything happen faster. The two minus signs cancel. The EMF itself therefore swings smoothly back and forth as a sine wave.
---
$ \varepsilon_{\max} = NBA\omega
The sine never exceeds 1, so the peak voltage is NBAω. More turns, a stronger field, a larger coil or faster turning all raise it, and each raises it in direct proportion.
:::

The output is a sine wave: a voltage that alternates in direction twice in every turn. So alternating current is not an engineering choice imposed on generators; it is what a rotating coil naturally produces. Notice also when the voltage peaks: when sin(ωt) = 1, which is when the coil lies edge-on to the field. At that moment no field passes through the coil at all, yet the voltage is greatest, because Faraday's law cares about the rate of change of flux, and that is when the flux is changing fastest. A generator that collects current through two separate continuous rings, called slip rings, delivers this AC directly. One with a commutator instead flips the connections each half-turn and delivers a current that always flows one way, though it pulses.

To see the size of the effect, take a coil of 100 turns, 10 cm square so that A = 0.01 m², in a field of 0.5 T, turning 50 times a second, which means ω = 2π × 50 ≈ 314 radians per second. The peak voltage is 100 × 0.5 × 0.01 × 314 ≈ 160 V, from a coil small enough to hold in one hand.

::: history 1831–1832 · London and Paris
In October 1831, only weeks after discovering induction, Faraday mounted a copper disc so that it could spin between the poles of a large magnet. Contacts touching the disc's axle and its rim drew off a steady current whenever the disc turned. It was the first generator. The next year Hippolyte Pixii, an instrument maker in Paris, built a hand-cranked machine that spun a horseshoe magnet past two coils and produced alternating current; at the suggestion of Ampère he soon added a commutator to turn it into direct current, which was what experimenters of the time wanted. For decades generators remained weak, limited by the strength of their permanent magnets. They became practical in the 1860s, when Werner von Siemens and others realized that a generator could use part of its own output to power electromagnets in place of the permanent ones, producing far stronger fields.
:::

## Why is a generator hard to turn?

A generator does not give electricity for nothing, and Lenz's law shows exactly where the price is paid. When a lamp or any other load is connected and current flows in the generator's coil, that current is sitting in a magnetic field, so it feels the motor force. Lenz's law says the induced current always opposes the change that produced it, which here means the force opposes the rotation. The more current the load draws, the larger this opposing force, and the harder you must work to keep the coil turning. Mechanical work goes in, electrical energy comes out, and the two match, apart from small losses as heat.

The same thing happens the other way round. A motor's coil, once spinning in a magnetic field, is also a generator, and it induces an EMF of its own. By Lenz's law this EMF opposes the supply voltage that drives the motor, so it is called a **back EMF**. At full speed the back EMF nearly cancels the supply, and only a small current flows: just enough to make up for friction and the work the motor is doing. When the motor is starting, or is jammed so that it cannot turn, there is no back EMF, and the current is limited only by the small resistance of the coil, so it is very large. This is why the lights in a house can dim for a moment when a big motor switches on, and why a motor that is stalled for long can overheat and burn out.

Electric cars use this symmetry directly. In **regenerative braking**, the car's motor is switched to act as a generator. The car's motion turns the coil, the coil drives current back into the battery, and the opposing force from Lenz's law slows the car down, turning kinetic energy back into stored charge instead of wasting it as heat in the brakes.

::: note
A motor and a generator are one machine. Current in a coil in a magnetic field makes it turn, and turning a coil in a magnetic field makes a current. Each effect opposes the other, so the work done to turn a generator equals the electrical energy it delivers, and energy is conserved.
:::

## How do generators power the world?

In a power station, something turns a turbine attached to a generator: steam from water boiled by burning fuel or by nuclear fission, or falling water, or the wind. Large generators reverse the arrangement described above, spinning powerful electromagnets inside fixed coils, which avoids passing huge currents through sliding contacts. Their speed of rotation is locked to the frequency of the grid, 50 or 60 cycles per second. On the other side, Nikola Tesla's AC induction motor of 1888, which needs no brushes or commutator at all, made alternating current useful for driving factory machinery as well as for lighting. Together with the transformer, it is why AC won its contest with Edison's direct current, as the lesson on electric power described, and why the plant at Niagara Falls in the 1890s could send hydroelectric power to a distant city.

::: quiz
? A generator coil is turned twice as fast. What happens to its peak voltage?
- It stays the same
+ It doubles
- It quadruples
= The peak EMF is NBAω, which is proportional to the angular speed ω. Turning twice as fast makes the flux change twice as quickly, so the peak voltage doubles.
? Why does a generator become harder to turn when a load draws current from it?
- Friction increases with current
+ The current in the coil feels a magnetic force that opposes the rotation
- The magnetic field gets weaker
= By Lenz's law the induced current resists the motion that causes it. The extra mechanical work you must do against that force is exactly what becomes the electrical energy delivered to the load.
? When is a motor most likely to draw a dangerously large current?
- At full speed with no load
+ When it is stalled or just starting
- When the supply voltage is low
= A coil that is not turning produces no back EMF to oppose the supply. The current is then limited only by the coil's small resistance, so it is very large.
:::
