Ørsted had shown that electricity makes magnetism: a current in a wire produces a magnetic field around it. The obvious next question was whether it worked the other way. Could magnetism make electricity? For a decade after 1820, experimenters put magnets next to wires and coils, and wires next to other current-carrying wires, and found nothing. The answer, when it finally came, explained why they had failed. A *steady* magnetic field does nothing to a circuit; a *changing* one drives a current through it. This effect, called electromagnetic induction, is the basis of almost every watt of electricity generated today, and of the transformers that carry it to our homes.

## Why was the effect so hard to find?

The experimenters of the 1820s were reasoning from a natural but misleading analogy. A steady current makes a steady magnetic field, so they expected a steady magnetic field to make a steady current. They set up strong magnets beside coils of wire, connected the coils to sensitive meters, and waited. Nothing happened, because nothing could. What they were missing was that the effect only appears at moments of change, and those moments are brief.

::: history 1831 · London
Michael Faraday had started life as a bookbinder's apprentice with almost no formal schooling, and had risen to become the leading experimenter at the Royal Institution in London. On 29 August 1831 he wound two separate coils of insulated wire on opposite sides of an iron ring, so that they did not touch. One coil could be connected to a battery; the other was connected to a galvanometer, a sensitive current meter. With current flowing steadily in the first coil, the galvanometer in the second showed nothing, just as everyone before him had found. But at the instant he connected the battery the needle jumped, then settled back to zero, and at the instant he disconnected it the needle jumped the other way. Over the following weeks he found that pushing a bar magnet into a coil produced a current, and pulling it out produced a current in the opposite direction, but holding it still inside the coil produced nothing. Joseph Henry in America had seen similar effects at about the same time, but Faraday published first, and his systematic study of the effect is what convinced others.
:::

Faraday's key insight was that what matters is change. A current is induced in a circuit whenever the magnetic field passing through it increases or decreases, and it does not matter how the change is brought about. Switching a nearby current on or off, moving a magnet, moving the circuit, or turning it round in a field all produce the same kind of effect. To turn this insight into a law, he needed a way to say precisely how much magnetic field passes through a circuit.

## How much field passes through a loop?

Think of the field lines of the previous lessons and imagine a flat loop of wire placed among them. Some lines thread through the loop, and the more of them do so, the more field passes through it. Three things affect the number: how strong the field is, which makes the lines more crowded; how large the loop is; and how the loop is tilted, since a loop facing the field squarely catches the most lines, while a loop turned edge-on catches none. These are combined in a quantity called the **magnetic flux**, written Φ (the Greek letter phi). For a flat loop of area A in a uniform field B, it is

::: equation Magnetic flux
\Phi = BA\cos\theta
:::

where θ is the angle between the field and the line sticking straight out of the face of the loop. When the field passes straight through, θ = 0 and cos θ = 1, so the flux is simply BA, as large as it can be. When the field lies along the plane of the loop, θ = 90° and cos θ = 0, so no field passes through and the flux is zero. The unit of flux is the **weber** (Wb), equal to one tesla multiplied by one square metre.

## What does Faraday's law say?

A changing flux drives a current round a circuit, and to drive a current something must supply energy to the charges, just as a battery does. The push that does this is measured like a battery's voltage, in volts, and is called an **electromotive force**, or EMF, written ε. (Despite the name, it is not a force; it is energy per unit charge.) Faraday found that the EMF depends on how fast the flux changes. If the circuit is a coil of N turns, the same changing flux passes through every turn, and each turn contributes the same EMF, so they add.

::: equation Faraday's law of induction
\varepsilon = -N\frac{\mathrm{d}\Phi}{\mathrm{d}t}
:::

The EMF is the rate of change of flux, multiplied by the number of turns. This explains every detail of Faraday's ring. While the current in the first coil was steady, the flux through the second coil was large but unchanging, so there was no EMF. Switching on made the flux rise from zero to its full value in a fraction of a second, a rapid change, so the needle jumped. Switching off made it fall just as quickly, and the needle jumped the other way, because a falling flux gives an EMF of the opposite sign. A faster change gives a larger voltage, which is why plunging a magnet quickly into a coil gives a bigger kick than moving it slowly.

## Which way does the induced current flow?

The minus sign in Faraday's law carries a physical principle of its own, worked out in 1834 by Heinrich Lenz, a physicist working in St Petersburg, and called **Lenz's law**: the induced current always flows in the direction whose own magnetic field opposes the change in flux that produced it.

Here is what that means in practice. Push the north pole of a magnet toward a coil. The flux through the coil increases. The induced current flows in the direction that makes the near end of the coil a north pole too, so that it repels the approaching magnet and pushes back against the increase. Pull the magnet away, and the current reverses, making the near end a south pole that attracts the retreating magnet and tries to hold it back. Either way, you have to do work to move the magnet, and that work is where the electrical energy in the coil comes from.

There is a powerful argument that the sign could not be otherwise. Suppose the induced current helped the change instead of opposing it. Then pushing a magnet toward a coil would make the coil attract it, pulling it in faster, which would increase the flux faster, which would induce a larger current, which would pull harder still. The magnet would accelerate and the current would grow without limit, with no energy supplied: a perpetual motion machine. Since energy cannot be created from nothing, the induced current must oppose the change. Lenz's law is the conservation of energy, appearing in electromagnetism.

## Can induction be explained by the magnetic force?

In some cases, induction can be understood using only the magnetic force from the lesson on moving charges, and checking that the two explanations agree is a good test of Faraday's law. Picture two parallel metal rails, a distance L apart, joined at one end by a fixed bar. A metal rod lies across the rails and slides along them at a steady speed v, so the rod, rails and end bar form a rectangle that grows longer as the rod moves. A uniform magnetic field B points straight through the rectangle.

::: steps EMF of a sliding rod
$ \Phi = BA = BLx
The circuit is a rectangle of width L and length x, where x is the rod's current distance from the joined end, so its area is Lx. The field passes straight through the rectangle, so cos θ = 1, and the flux is the field times the area.
---
$ \frac{\mathrm{d}\Phi}{\mathrm{d}t} = BL\frac{\mathrm{d}x}{\mathrm{d}t} = BLv
To use Faraday's law we need the rate at which the flux changes. The field and the width of the rails stay fixed, so the only thing changing is x, and the rate at which x grows is just the rod's speed v. So the flux grows at the steady rate BLv.
---
$ \varepsilon = BLv
Faraday's law, for a single loop with N = 1, says the EMF equals this rate of change. Here we take the size of the EMF and leave its direction to Lenz's law. A faster rod, a stronger field or wider rails all give a larger voltage.
---
$ F = qvB
Now look at the same situation another way. Each free charge q in the rod is carried along with it at speed v, straight across the field, so by the rule from the lesson on moving charges it feels a magnetic force qvB. That force points along the length of the rod, so it pushes charges along the rod and round the circuit.
---
$ W = qvB \cdot L = q\,(BLv)
As a charge is pushed the full length L of the rod, this force does work qvBL on it. An EMF is defined as work done per unit charge, so the EMF is qvBL ÷ q = BLv. This is exactly the answer Faraday's law gave, which shows that, for a moving conductor, induction is the magnetic force on the charges it carries.
:::

A rod half a metre long moving at 2 m/s across a field of 0.1 T gives an EMF of 0.5 × 2 × 0.1 = 0.1 V. That is small, but nothing stops us using a coil of many turns spinning quickly in a strong field instead, and then the voltages reach hundreds of volts, as the next lesson shows.

However, not every case of induction reduces to a moving conductor. When a coil sits still and a magnet moves past it, or when the current in a nearby coil is switched on, no charge in the coil is moving through any field, so there is no magnetic force on it at all. Yet Faraday's law still holds, and current still flows. Something else must push the charges. The answer is that a changing magnetic field creates an electric field around itself, with field lines that circle round the region where the magnetic field is changing, and that electric field pushes on the charges in the wire. A wire is not even needed: the circling electric field exists in empty space. James Clerk Maxwell made this idea one of his four equations, and it is half of the explanation of how light travels, as the next chapter shows.

::: note
A changing magnetic flux through a circuit induces an EMF equal to the rate of change of the flux, multiplied by the number of turns. The induced current always flows so as to oppose the change that caused it, which is the conservation of energy: you must do work to change the flux, and that work becomes electrical energy.
:::

## What is induction used for?

Faraday's iron ring was, without his knowing it, the first **transformer**. Feed an alternating current into the first coil, and the flux it drives round the iron rises and falls continuously, so it induces an alternating EMF in the second coil all the time, not just at the moments of switching. Because the iron guides almost all of the flux through both coils, every turn of either coil sees the same changing flux, so the voltage per turn is the same in each. A coil with more turns therefore has a proportionally larger voltage: V₂/V₁ = N₂/N₁. A secondary coil with ten times as many turns as the primary gives ten times the voltage. Energy is conserved, so a good transformer that raises the voltage tenfold delivers about a tenth of the current. This is how power grids raise the voltage for long-distance transmission and lower it again for homes, as the lesson on electric power explained. And because a transformer needs changing flux, it works only with alternating current, which is why the grid is AC.

Induction turns up in many other places. Induction cooktops drive alternating currents in the base of the pan itself, which heats up while the glass surface stays cool. Wireless phone chargers pass energy across a small gap from one coil to another. The pickups on electric guitars sense the vibrating steel strings as changing flux through a coil. Metal detectors notice the currents induced in buried metal. And magnetic brakes on some trains and roller coasters use the currents induced in a moving metal fin; by Lenz's law, those currents oppose the fin's motion and slow it smoothly, with nothing touching.

::: quiz
? A magnet sits motionless inside a coil connected to a meter. What does the meter read?
+ Zero
- A steady current
- A current that grows with time
= Only a changing flux induces an EMF. A stationary magnet produces a large but constant flux through the coil, so nothing drives a current.
? A magnet's north pole is pushed toward a coil. What does the induced current do?
- It attracts the magnet, pulling it in
+ It makes the coil's near end a north pole, repelling the magnet
- Nothing, because the coil is not connected to a battery
= By Lenz's law the induced current opposes the increase in flux, so it makes the near end of the coil a north pole that pushes back on the approaching magnet. The work done against that push becomes the electrical energy.
? A transformer has 100 turns on its input coil and 2,000 on its output coil. With 12 V AC in, what is the output?
- 0.6 V AC
+ 240 V AC
- 240 V DC
= Every turn sees the same changing flux, so V₂/V₁ = N₂/N₁ = 2000/100 = 20, and the output is 20 × 12 V = 240 V. It alternates like the input, because it is induced by the input's changing flux.
:::
