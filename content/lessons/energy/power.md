Climb a flight of stairs slowly, or run up it two steps at a time. Either way you raise your body through the same height, so you do the same work against gravity, and by the end you have gained the same potential energy. Yet running leaves you out of breath and walking does not. The difference is not how much work you did but how quickly you did it. The rate of doing work, or more generally of transferring energy from one place or form to another, is called **power**. It is the quantity that decides how big an engine needs to be, why a car uses so much more fuel at high speed, and why a horse was once the unit of comparison for machines.

## What is power?

Power is work divided by the time taken to do it:

::: equation Power
P = \frac{W}{t} \quad \text{or, at an instant,} \quad P = \frac{\mathrm{d}W}{\mathrm{d}t}
:::

The first form gives the average power over a time t. If the work is not done at a steady rate, the average hides the variation, just as an average speed hides the moments a car stopped at traffic lights. The second form gives the power at a single instant: the derivative of work with respect to time, the rate at which work is being done at that moment, in the sense of the rates lesson.

The unit of power is the **watt** (W), named after James Watt. One watt is one joule of work, or of energy transferred, per second: 1 W = 1 J/s. A thousand watts is a kilowatt (kW), and a million is a megawatt (MW).

To see the definition at work, go back to the stairs. A 70 kg person climbs a 3 m flight. The work done against gravity is mgh = 70 × 9.8 × 3, about 2,060 J. Running up in 4 s takes an average power of 2,060 / 4, about 510 W. Walking up in 12 s takes about 170 W. The work is the same in both cases; the power differs by a factor of three because the time does. Half a kilowatt is close to what a fit person can manage for a few seconds, and it is why a sprint up the stairs feels hard.

## How does power relate to force and speed?

Often what we know about a machine is the force it exerts and how fast it moves, not the work done over some interval. A car engine pushes the car along a road; a crane hauls up a load; a cyclist drives the pedals round. For these, power has a simple and useful form.

::: steps Power delivered by a force
$ \mathrm{d}W = F\,\mathrm{d}x
Consider a force F pushing a body along the direction of the force. In a very short time dt, the body moves a short distance dx. Over so short a distance the force is effectively constant, so by the definition of work, the force does work F dx in that moment.
---
$ P = \frac{\mathrm{d}W}{\mathrm{d}t} = F\frac{\mathrm{d}x}{\mathrm{d}t}
Power is the rate of doing work, so divide the small amount of work by the small time it took. Since F is constant over the tiny interval, it stays outside as a factor, and what is divided by dt is just the distance moved.
---
$ P = Fv
The distance moved per unit time, dx/dt, is the body's velocity v. So the power a force delivers is the force times the speed of the body it pushes. If the force acts at an angle to the motion, only its component along the velocity does work, and the general form is the dot product P = F·v.
:::

This short result explains why cars and bicycles have gears. An engine can deliver only up to a certain maximum power, set by how fast it can burn fuel. Since P = Fv, a fixed power can be spent as a large force at low speed or a small force at high speed, but not both at once. On a steep hill the car needs a large force to climb, so a low gear makes the wheels turn slowly with a large force. On a level motorway little force is needed but the speed is high, so a high gear makes the wheels turn fast with a small force. The gearbox does not create power; it chooses how to divide the same power between force and speed.

The same result explains why driving fast is expensive. At a steady speed on a level road the car is not accelerating, so the engine's driving force must exactly balance the forces resisting the motion, mainly air drag and the rolling resistance of the tyres. At motorway speeds drag dominates. The friction lesson showed that drag grows as the square of speed. The power needed to overcome it is drag force times speed, so it grows as the cube of speed.

Numbers show how steep that is. For a typical car with a drag coefficient of 0.3 and a frontal area of 2.2 m², moving at 30 m/s (about 110 km/h) through air of density 1.2 kg/m³, the drag force is ½ × 1.2 × 0.3 × 2.2 × 30², about 360 N. The power needed just to push through the air is 360 × 30, about 11 kW. At half the speed, 15 m/s, the drag force is a quarter as large and the speed is half, so the power is an eighth: about 1.3 kW. Doubling the speed multiplies the power by eight. Because the car also covers each kilometre in half the time, the energy used per kilometre, and so the fuel, goes up by a factor of four.

## How is power different from energy?

Power and energy are easily confused, partly because the words are used loosely in everyday speech. Energy is an amount: how much work can be done, or has been done. Power is a rate: how fast that energy is delivered. The difference shows up in electricity bills, which charge for energy, measured in **kilowatt-hours**. A kilowatt-hour is the energy delivered by a power of one kilowatt running for one hour:

::: equation The kilowatt-hour
1\,\text{kWh} = 1000\,\text{W} \times 3600\,\text{s} = 3.6 \times 10^6\,\text{J}
:::

So a 2 kW kettle running for 3 minutes, a twentieth of an hour, uses 2 × 0.05 = 0.1 kWh. A 10 W light bulb left on for ten hours uses the same amount. One has high power for a short time, the other low power for a long time, and the energy, which is what you pay for, is the same.

Some typical powers give a sense of scale:

| | Power |
|---|---|
| Human body at rest | about 100 W |
| Trained cyclist, sustained for an hour | about 250–300 W |
| Sunlight on 1 m² above the atmosphere | about 1.36 kW |
| Family car at full throttle | about 100 kW |
| Large power station | about 1 GW |

The first line follows from the food we eat. A diet of 2,000 kilocalories a day supplies about 8.4 million joules. Spread over the 86,400 seconds in a day, that is close to 100 W, the power of an old-fashioned incandescent light bulb. Nearly all of it ends up as heat, which is why a crowded room warms up.

## Why does some power always go to waste?

No machine turns all the power it takes in into useful output. Some always ends up as heat, through friction in moving parts, electrical resistance in wires, or, in an engine, simply because hot exhaust gas must be thrown away. The **efficiency** of a machine is the fraction of the input power that comes out as useful output, the useful power out divided by the power in.

Efficiencies vary widely. A car's petrol engine turns about 30% of the chemical energy in its fuel into work at the wheels at best, and less in town driving. A large electric motor manages over 90%. Human muscle converts roughly 25% of the chemical energy it uses into mechanical work, the rest becoming the heat that makes exercise warm. In every case, as the last lesson showed, the lost energy is not destroyed. It still exists, but spread out as heat at a low temperature it can no longer be gathered up to do useful work.

::: history 1782 · Birmingham
James Watt's improved steam engine, far more economical with coal than the engines before it, was sold mainly to the owners of mines and mills, many of whom had been using horses to drive pumps and machinery. To sell the engine, Watt had to tell them how many horses it would replace, so he needed a figure for the work a horse could do. From observing horses turning mill wheels, he settled in 1782 on a value of 33,000 foot-pounds per minute: one **horsepower**, the work of raising 33,000 pounds through one foot each minute, about 746 W in modern units. The figure is generous. A horse can work at that rate for a while but not for a whole working day, while a horse at full gallop can briefly exceed ten times as much. As a sales tool the unit worked, and it is still used for car engines. The SI unit of power was named the watt in his honour a century later, in 1882.
:::

::: note
Power is the rate at which energy is transferred: P = W/t, or P = Fv for a force pushing a body at speed v. Doing the same work faster takes more power, and overcoming drag takes power that grows as the cube of speed.
:::

::: quiz
? Two people of equal mass climb the same staircase, one in 10 s and one in 20 s. Compared with the slower one, the faster one does
- twice the work at the same power
+ the same work at twice the power
- twice the work at twice the power
= The work depends only on their weight and the height climbed, so it is the same for both. Doing it in half the time takes twice the power.
? A car's engine delivers 30 kW to the wheels while it cruises at a steady 25 m/s. What is the total resistive force on the car?
+ 1200 N
- 750 N
- 30 000 N
= At steady speed the driving force exactly balances the resistance, and P = Fv gives F = 30,000 W ÷ 25 m/s = 1200 N.
? If air drag dominates, how does the power needed to cruise change when speed doubles?
- It doubles
- It quadruples
+ It rises eightfold
= The drag force grows as v², and the power needed is force times speed, so it grows as v³. Doubling the speed multiplies the power by 2³ = 8.
:::
