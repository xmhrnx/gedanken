Everything in the previous chapter was about charge that sits still, or that moves only for an instant. Before 1800 the only way to make a lot of charge move was a spark: a sudden, uncontrolled rush that was over in a fraction of a second, leaving nothing to study but a crack and a smell. Then a new device made charge flow steadily, for hours, through an ordinary wire. That steady flow is an **electric current**, and almost everything electrical in modern life depends on it. This lesson asks what a current is, how to measure it, which way it goes, and how fast the charges inside a wire actually travel. The last answer is one of the surprises of physics.

## Where did steady current come from?

::: history 1800 · Pavia and Como, Italy
In the 1780s Luigi Galvani in Bologna noticed that a dissected frog's leg twitched when touched by two different metals at once, and concluded that animals contain a special "animal electricity". Alessandro Volta, a professor at Pavia, was not convinced. He suspected the electricity came from the two metals, and the frog was only acting as a sensitive detector. To prove it he left the frog out entirely: he stacked discs of zinc and copper, separated by cardboard soaked in salt water, into a column. Touch wires to the top and bottom and you felt a shock, again and again, without recharging anything. Volta described his "pile" in a letter to the Royal Society in London in 1800. It was the first battery, and within weeks others were using it to split water into hydrogen and oxygen.
:::

What mattered about Volta's pile was not its strength but its persistence. A charged jar or a friction machine could deliver a jolt, but the pile kept pushing charge for as long as its chemistry lasted. For the first time, electricity was something you could hold steady and measure at leisure, and the study of currents could begin.

## What exactly is a current?

The natural way to describe a flow of anything is to count how much of it passes a point per second. For a river you count litres per second; for traffic, cars per minute. For electricity you count charge. Choose any cross-section of a wire, an imaginary slice straight across it, and ask how much charge crosses that slice in a given time.

::: equation Current
I = \frac{Q}{t}
:::

Here Q is the charge that crosses the slice in time t. If the current is changing, the same idea applies over a very short interval, and I = dQ/dt, the rate of flow at an instant. The unit of current is the **ampere** (A), which is one coulomb of charge per second.

A coulomb is a large amount of charge: about 6.24 × 10¹⁸ times the charge of a single electron. So a current of one ampere means more than six billion billion electrons crossing the slice every second. That is roughly what flows through a phone charger's cable, which carries 1 to 3 A. A kettle draws about 10 A, a car's starter motor a few hundred, and a lightning stroke can peak at around 30,000 A, though only for some millionths of a second. The range is huge, but the quantity is always the same kind of thing: charge per second.

## Why does a current need a loop?

A battery is sometimes described as a store of electricity that empties into a circuit. That picture is wrong in an important way. A battery does not create charge, and it does not send out charge of its own that is used up somewhere along the wire. The wire is already full of charges that can move, and the battery acts like a pump that pushes them round. Charge leaves one terminal, travels through the wire and the lamp or motor, and returns to the other terminal. For this to keep happening there must be a complete path, a closed loop, which is why we call it a **circuit**. Break the loop anywhere and the flow stops everywhere.

There is a strong reason the current must be the same all the way round a simple loop. Suppose more charge entered some stretch of wire each second than left it. Charge would pile up there, and that growing pile would repel further charge arriving, slowing the flow in until it matched the flow out. In a steady circuit, then, nothing accumulates anywhere. The same idea applies where several wires meet at a junction: the current flowing in must equal the current flowing out, because charge is neither made nor destroyed at the junction and has nowhere to hide. This rule is called **Kirchhoff's current law**. It looks like a new law of circuits, but it is just conservation of charge applied to wires.

## Which way does it flow?

There is an awkward historical accident buried in every circuit diagram. In the 1740s Benjamin Franklin, studying static electricity, guessed that electricity was a single fluid. A body with extra fluid he called positive, one with too little negative, and he imagined the fluid flowing from positive to negative. He had to choose which kind of charge to call positive without any way of knowing what was really moving, and he chose the charge that collects on glass rubbed with silk.

Engineers and physicists still follow Franklin's choice. **Conventional current** is defined as the direction positive charge would move: out of the battery's positive terminal, round the circuit, and back into the negative terminal. But in metal wires the charges that actually move are electrons, and electrons carry negative charge. They drift the other way, from the negative terminal round to the positive one. Nobody could have known this until J. J. Thomson discovered the electron in 1897, a century and a half after Franklin's guess, and showed it to be a light particle carrying negative charge.

For nearly every calculation the mismatch makes no difference. Negative charge moving to the left carries charge across a slice in exactly the same way, and changes the balance of charge on each side in exactly the same way, as positive charge moving to the right. So the current is the same either way, and the conventional direction is kept. In some conductors both kinds of charge do move. In salt water, in the liquid inside a battery and in the cells of your body, the carriers are **ions**, atoms that have gained or lost electrons. Positive ions drift one way, negative ions the other, and both add to the current.

## How fast do the electrons move?

When a lamp lights the instant you flip a switch, it is tempting to imagine electrons racing from the switch to the bulb. To check, we can relate the current to the speed of the charges that carry it. Picture a wire of cross-sectional area A. Inside it are n free electrons in each cubic metre, each with charge e, and suppose they all drift along the wire at the same average speed v.

::: steps Drift speed and current
$ \text{volume} = A\,v\,t
Fix one cross-section of the wire and wait for a time t. In that time each electron moves a distance vt along the wire. So every electron that started within a distance vt upstream of the cross-section gets through it, and those electrons fill a cylinder with base area A and length vt.
---
$ Q = n\,e\,A\,v\,t
To find the charge in that cylinder, multiply its volume by the number of free electrons per cubic metre, n, which gives the number of electrons, and then by the charge each carries, e. This is the total charge that crosses the cross-section in time t.
---
$ I = \frac{Q}{t} = n\,e\,A\,v
Current is charge per unit time, so divide by t, which cancels. The result says the current grows with the number of carriers, the charge each carries, the size of the wire and the speed of the drift, which matches intuition: more carriers, or faster ones, deliver more charge per second.
---
$ v = \frac{I}{neA}
Divide both sides by neA to isolate the drift speed. Everything on the right can be measured or looked up, so we can now find how fast the electrons really move for a given current.
:::

Copper has roughly one free electron per atom, which works out to n ≈ 8.5 × 10²⁸ per cubic metre. Take a current of 1 A in a household-sized wire with a cross-section of 1 square millimetre, which is 10⁻⁶ m².

::: equation Drift speed in a copper wire
v = \frac{1\,\text{A}}{(8.5 \times 10^{28}\,\text{m}^{-3})(1.6 \times 10^{-19}\,\text{C})(10^{-6}\,\text{m}^2)} \approx 7 \times 10^{-5}\,\text{m/s}
:::

That is less than a tenth of a millimetre per second. A snail is faster. At this rate an electron would take hours to travel a metre of wire. The reason the number is so small is that n is so enormous: there are so many free electrons in copper that only a gentle drift is needed to carry an ampere.

The drift is not the whole motion. Each electron is also darting about at around 10⁶ m/s, in random directions, constantly bouncing off the vibrating atoms of the metal. Because the directions are random, this frantic motion carries no net charge anywhere. The drift is a slow average shuffle laid on top of it, like a crowd milling about in a square while very slowly edging toward one exit.

## Then why does the lamp light at once?

The answer is that nothing has to travel from the switch to the lamp. The wire, including the filament of the lamp, is already full of free electrons before you touch the switch. Closing the switch completes the circuit, and the battery's push, carried by an electric field in and around the wires, spreads along the whole loop at a large fraction of the speed of light. Within a tiny fraction of a second every free electron in the circuit, including those already inside the lamp, starts drifting at once. The electrons that warm the filament were there all along.

A garden hose that is already full of water behaves in a similar way: open the tap and water leaves the far end immediately, though no drop has made the whole trip. The comparison is good for the timing, but it fails on what carries the energy. In a hose the water itself carries the push, while in a circuit the energy travels mainly through the electric and magnetic fields around the wire, as later chapters will show.

::: note
Current is the rate at which charge passes a point, I = Q/t, measured in amperes. In a metal wire it is a vast number of electrons drifting very slowly, a fraction of a millimetre per second. A lamp lights at once because the wire is already full of electrons and the push that sets them moving spreads at nearly the speed of light.
:::

## One way, or back and forth?

A battery pushes charge in one direction only, giving **direct current**, or DC. The electricity from a wall socket is different. It is **alternating current**, AC: the push reverses direction 50 or 60 times every second, depending on the country, so the current flows one way, stops, flows the other way, and stops again, over and over.

Given how slowly electrons drift, this means that in an AC circuit they hardly go anywhere. Each one rocks back and forth by less than a thousandth of a millimetre and ends each cycle where it began. Yet the lamp still lights and the kettle still boils. That only makes sense because, as the hose comparison suggested, the energy does not ride along with the electrons from the power station to your kettle. It is carried by the fields, and the electrons in the kettle's element simply pass it on to the metal as heat, whichever way they happen to be moving. Why the world settled on AC for its power grids is a story told later in this chapter and the next.

::: quiz
? A current of 2 A flows for one minute. How much charge passes any point in the wire?
- 2 C
- 30 C
+ 120 C
= Charge is current multiplied by time, so Q = It = 2 A × 60 s = 120 C. One minute must be converted to 60 seconds because an ampere is a coulomb per second.
? Three wires meet at a junction. 5 A flows in along one and 2 A flows out along another. What does the third wire carry?
- 7 A in
+ 3 A out
- 3 A in
= Charge cannot pile up at the junction, so the total current in must equal the total current out. With 5 A coming in and only 2 A leaving through the second wire, the remaining 3 A must leave through the third.
? Why does a lamp light almost instantly, when electrons drift at under a millimetre per second?
- Electrons speed up greatly near the switch
- The lamp stores charge in advance
+ The wire is already full of electrons, and the push that moves them spreads at nearly light speed
= Every free electron in the circuit, including those already in the lamp, starts drifting almost at once when the circuit closes. None of them needs to travel from the switch to the lamp.
:::
