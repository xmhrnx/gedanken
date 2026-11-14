Ørsted showed that electricity makes magnetism. The obvious question was whether magnetism could make electricity. For a decade experimenters put magnets next to wires and found nothing. The answer, when it came, was that a *steady* magnetic field does nothing; a *changing* one drives a current. Almost every watt of electricity generated today depends on that fact.

## Faraday's ring

::: history 1831 · London
On 29 August 1831, Michael Faraday at the Royal Institution wound two separate coils of insulated wire on opposite sides of an iron ring. One coil was connected to a battery, the other to a galvanometer, a sensitive current meter. With current flowing steadily in the first coil, the galvanometer showed nothing. But at the instant he connected the battery the needle jumped, and at the instant he disconnected it the needle jumped the other way. In the following weeks he found that simply pushing a bar magnet into a coil, or pulling it out, also produced a current, but holding it still did not. Joseph Henry in America had seen similar effects at about the same time but published later.
:::

Faraday's key insight was that what matters is change. A current is induced whenever the magnetic field passing through a circuit increases or decreases, however the change is caused: by switching a current, moving a magnet, moving the circuit, or turning it.

## Magnetic flux

To make "the field passing through a circuit" precise, define the **magnetic flux** Φ (phi) through a flat loop of area A in a uniform field B:

::: equation Magnetic flux
\Phi = BA\cos\theta
:::

θ is the angle between the field and the line perpendicular to the loop. Flux is largest when the field passes straight through the loop and zero when the field lies along its plane. In field-line terms, flux counts the lines threading the loop. Its unit is the **weber** (Wb), one tesla times one square metre.

## Faraday's law

The induced voltage, called an **electromotive force** or EMF (ε), equals the rate of change of the flux. For a coil of N turns, each turn contributes, so:

::: equation Faraday's law of induction
\varepsilon = -N\frac{\mathrm{d}\Phi}{\mathrm{d}t}
:::

A faster change gives a larger voltage, which is why Faraday's galvanometer jumped only at the moment of switching. The minus sign is **Lenz's law** (Heinrich Lenz, 1834): the induced current always flows in the direction whose own magnetic field opposes the change in flux. Push a magnet's north pole into a coil and the coil's induced current makes its near end a north pole, repelling the magnet. You must do work to push it in, and that work becomes the electrical energy. If the sign were the other way, the induced current would pull the magnet in faster, inducing more current, without limit: a perpetual motion machine. Lenz's law is energy conservation.

## A moving rod

The simplest case to work out connects induction with the magnetic force of the previous lessons. A metal rod of length L slides at speed v along two parallel rails, which are joined at one end, forming a rectangle whose area grows. A uniform field B points straight through the rectangle.

::: steps EMF of a sliding rod
$ \Phi = BA = BLx
The circuit is a rectangle of width L and length x, where x is the rod's distance from the joined end. The field passes straight through, so cos θ = 1.
---
$ \frac{\mathrm{d}\Phi}{\mathrm{d}t} = BL\frac{\mathrm{d}x}{\mathrm{d}t} = BLv
B and L are fixed; only x changes, at the rod's speed v.
---
$ \varepsilon = BLv
Faraday's law with N = 1, taking the size and leaving the direction to Lenz's law.
---
$ qvB \cdot L = q\,\varepsilon
Check from the other direction: each charge in the rod moves at v across B and feels a force qvB along the rod. Carried the length of the rod, that force does work qvBL, which is q times an EMF of BLv. The two views agree.
:::

A rod 0.5 m long moving at 2 m/s across 0.1 T gives 0.1 V. Small, but a coil of many turns spinning quickly in a strong field gives hundreds of volts, as the next lesson shows.

Not every case reduces to moving charges. When a coil sits still and the magnet moves, or the current in a nearby coil changes, no charge in the coil is moving through a field. Faraday's law still holds, because a changing magnetic field creates an electric field that circles around it. Maxwell made this part of his equations, and it is half of how light propagates.

::: note
A changing magnetic flux through a circuit induces a voltage equal to its rate of change. The induced current always opposes the change that caused it, which is energy conservation in disguise.
:::

## Transformers and other uses

Faraday's ring was the first **transformer**. Feed alternating current into one coil and the flux in the iron changes continuously, inducing an alternating voltage in the second coil. Both coils share the same changing flux, so the voltage per turn is the same in each, giving V₂/V₁ = N₂/N₁. A coil with ten times the turns gives ten times the voltage (and, conserving power, a tenth of the current). This is how grids step voltage up for transmission and down for homes. It works only with changing current, which is why the grid is AC.

Induction also powers induction cooktops, which drive currents in the pan itself; wireless phone chargers; the pickups on electric guitars; metal detectors; and the magnetic brakes on some trains and roller coasters, where induced currents in a moving metal fin oppose its motion.

::: quiz
? A magnet sits motionless inside a coil connected to a meter. The meter reads
+ zero
- a steady current
- a current that grows with time
= Only a changing flux induces a voltage; a stationary magnet gives a constant flux.
? A magnet's north pole is pushed toward a coil. What does the induced current do?
- Attracts the magnet, pulling it in
+ Makes the coil's near end a north pole, repelling the magnet
- Nothing, because the coil was not connected to a battery
= Lenz's law: the induced current opposes the increase in flux, so it pushes back on the approaching pole.
? A transformer has 100 turns on its input coil and 2000 on its output coil. With 12 V AC in, the output is
- 0.6 V AC
+ 240 V AC
- 240 V DC
= V₂/V₁ = N₂/N₁ = 20, so the output is 20 × 12 = 240 V, alternating like the input.
:::
