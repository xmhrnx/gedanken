Electrostatics is about charge sitting still. Before 1800 the only way to make charge move was a spark: a sudden, uncontrolled discharge lasting a fraction of a second. Volta's pile changed that. Connect its two ends with a wire and charge flows steadily, for hours. That steady flow is an **electric current**.

## Defining current

Current is the rate at which charge passes a given point.

::: equation Current
I = \frac{Q}{t}
:::

Here Q is the charge passing a cross-section of the wire in time t. (When the current varies, I = dQ/dt.) The unit is the **ampere** (A), one coulomb per second. Since a coulomb is about 6.24 × 10¹⁸ elementary charges, 1 A means over six billion billion electrons passing each second. A phone charger supplies 1–3 A, a kettle about 10 A, a car's starter motor a few hundred; a lightning stroke peaks around 30,000 A, but only for microseconds.

Current needs a complete loop, a **circuit**. A battery does not create charge; it pushes charge that is already in the wire around the loop, and the same current that leaves one terminal returns to the other. Charge cannot pile up anywhere in a steady circuit, so at a junction where wires meet, the current flowing in equals the current flowing out. This is **Kirchhoff's current law**, and it is just conservation of charge.

## Which way?

Franklin's sign convention left a trap. **Conventional current** is defined as the direction positive charge would move: out of the positive terminal, around the circuit, into the negative terminal. But in metal wires the moving charges are electrons, which are negative and travel the other way. For almost every calculation it does not matter, since negative charge moving left is the same current as positive charge moving right. Engineers and physicists keep the conventional direction. In salt water, batteries and living cells, positive and negative ions both move, in opposite directions, and both contribute to the current.

::: history 1897 · Cambridge, England
J. J. Thomson measured how cathode rays, the glowing beams in evacuated tubes, were bent by electric and magnetic fields. The ratio of their charge to mass was the same whatever the metal of the electrode or the gas in the tube, and over a thousand times larger than for a charged hydrogen atom. He concluded they were a universal particle, much lighter than any atom, carrying negative charge: the electron. Only then did it become clear that current in metals is a flow of electrons, moving opposite to Franklin's "positive" direction.
:::

## How fast do the electrons move?

Picture a wire of cross-sectional area A containing n free electrons per cubic metre, each with charge e, all drifting along at an average speed v.

::: steps Drift speed and current
$ \text{volume} = A\,v\,t
In time t each electron moves a distance vt. All the electrons within that distance behind a chosen cross-section pass through it: they fill a cylinder of base A and length vt.
---
$ Q = n\,e\,A\,v\,t
Multiply the volume by the number of electrons per unit volume, n, and the charge per electron, e.
---
$ I = \frac{Q}{t} = n\,e\,A\,v
Divide by t. The current is proportional to the density of carriers, their charge, the wire's area and their speed.
---
$ v = \frac{I}{neA}
Solve for the drift speed.
:::

Copper has about one free electron per atom, n ≈ 8.5 × 10²⁸ per m³. For 1 A in a wire of cross-section 1 mm² (10⁻⁶ m²):

::: equation Drift speed in a copper wire
v = \frac{1}{(8.5 \times 10^{28})(1.6 \times 10^{-19})(10^{-6})} \approx 7 \times 10^{-5}\,\text{m/s}
:::

That is less than a tenth of a millimetre per second, slower than a snail. Individual electrons actually dart about at around 10⁶ m/s, but in random directions, colliding constantly; the drift is a slow average shuffle superimposed on that motion.

Yet a lamp lights the instant you flip a switch. The wire is already full of electrons, so nothing has to travel from the switch to the lamp. What travels is the electric field that pushes them, which spreads along the circuit at a large fraction of the speed of light. A garden hose already full of water behaves the same way: open the tap and water leaves the far end at once, though no single drop has made the trip.

::: note
Current is charge per second, I = Q/t. In a wire it is a vast number of electrons drifting very slowly; the signal travels fast because the whole wire is already full of charge.
:::

## Direct and alternating

A battery gives **direct current** (DC), flowing one way. Mains electricity is **alternating current** (AC): it reverses direction 50 or 60 times a second, depending on the country. In AC the electrons merely jiggle back and forth by a tiny fraction of a millimetre, yet they still deliver energy, because energy is carried by the field, not by electrons arriving at the lamp.

::: quiz
? A current of 2 A flows for one minute. How much charge passes?
- 2 C
- 30 C
+ 120 C
= Q = It = 2 A × 60 s = 120 C.
? Three wires meet at a junction. 5 A flows in along one and 2 A flows out along another. The third wire carries
- 7 A in
+ 3 A out
- 3 A in
= Charge is conserved, so total current in equals total current out: 5 = 2 + 3.
? Why does a lamp light almost instantly, given that electrons drift at under a millimetre per second?
- Electrons speed up near the switch
- The lamp stores charge in advance
+ The wire is already full of electrons, and the field that pushes them spreads at near light speed
= Every electron in the circuit starts moving almost at once; none needs to travel from switch to lamp.
:::
