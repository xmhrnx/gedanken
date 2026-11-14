Climb a flight of stairs slowly or run up it: either way you raise your body through the same height and do the same work against gravity. Running leaves you out of breath because you do it faster. The rate of doing work, or of transferring energy, is **power**.

## Definition

::: equation Power
P = \frac{W}{t} \quad \text{or, at an instant,} \quad P = \frac{\mathrm{d}W}{\mathrm{d}t}
:::

The first form is the average over a time t; the second, the instantaneous rate, is the derivative of work with respect to time, from the rates lesson. The unit is the **watt** (W): 1 W = 1 J/s.

Run up a 3 m flight of stairs in 4 s with a mass of 70 kg. The work against gravity is mgh = 70 × 9.8 × 3 ≈ 2060 J, so your average power is about 510 W. Walk it in 12 s and the power is 170 W, for the same work.

## Power, force and speed

When a force moves something steadily, its power has a simple form.

::: steps Power delivered by a force
$ \mathrm{d}W = F\,\mathrm{d}x
In a short time dt the body moves dx along the force, and the force does work F dx.
---
$ P = \frac{\mathrm{d}W}{\mathrm{d}t} = F\frac{\mathrm{d}x}{\mathrm{d}t}
Divide by dt to get the rate of doing work. F is treated as constant over the tiny interval.
---
$ P = Fv
dx/dt is the velocity. If the force is at an angle to the motion, only its component along the velocity counts: P = F·v.
:::

This explains gears. An engine delivers a roughly fixed maximum power, so P = Fv means it can supply a large force at low speed or a small one at high speed. Low gear trades speed for force on a hill; high gear does the reverse on a motorway.

It also explains why fast driving is expensive. At steady speed on a level road the engine must supply power equal to the force of drag times speed. From the friction lesson, drag grows as v², so the power needed to overcome it grows as v³. For a typical car (drag coefficient 0.3, frontal area 2.2 m²) at 30 m/s, about 110 km/h, drag is about 360 N and the power to beat it is about 11 kW. At half that speed it is one-eighth as much, about 1.3 kW. Doubling speed costs eight times the power, and four times the energy per kilometre.

## Energy is not power

Electricity bills are in **kilowatt-hours**, a unit of energy: a power of 1 kW for one hour.

::: equation The kilowatt-hour
1\,\text{kWh} = 1000\,\text{W} \times 3600\,\text{s} = 3.6 \times 10^6\,\text{J}
:::

A 2 kW kettle running for 3 minutes uses 0.1 kWh. Power says how fast; energy says how much. Some scales for comparison:

| | Power |
|---|---|
| Human body at rest | about 100 W |
| Trained cyclist, sustained for an hour | about 250–300 W |
| Sunlight on 1 m² above the atmosphere | about 1.36 kW |
| Family car at full throttle | about 100 kW |
| Large power station | about 1 GW |

The first line comes from food. A diet of 2000 kilocalories a day is about 8.4 million joules spread over 86 400 seconds: close to 100 W, the output of an old-fashioned light bulb.

## Efficiency

No machine turns all its input into useful output; some always goes to heat. The **efficiency** is the useful output power divided by the input power. A petrol engine manages about 30%, a large electric motor over 90%, human muscle about 25%. The rest is not destroyed, as the last lesson showed, but it is no longer useful.

::: history 1782 · Birmingham
James Watt's improved steam engines were sold to mine and mill owners who had been using horses to drive pumps and mills, so he rated them in horses. From observing mill horses he settled on a value of 33 000 foot-pounds per minute for one **horsepower**, about 746 W. It is generous: a horse can work at that rate for a while but not all day, and a sprinting horse can briefly exceed ten times that. The unit sold engines, and is still used for cars. The SI unit was named after Watt in 1882.
:::

::: note
Power is the rate of energy transfer: P = W/t, or P = Fv for a force on a moving body. The same work done faster takes more power.
:::

::: quiz
? Two people of equal mass climb the same staircase, one in 10 s and one in 20 s. Compared with the slower one, the faster one does
- twice the work at the same power
+ the same work at twice the power
- twice the work at twice the power
= Work depends on weight and height only. Doing it in half the time doubles the power.
? A car's engine delivers 30 kW while it cruises at 25 m/s. What is the total resistive force on the car?
+ 1200 N
- 750 N
- 30 000 N
= At steady speed the engine's force balances resistance, and P = Fv gives F = 30 000 / 25 = 1200 N.
? If air drag dominates, how does the power needed to cruise change when speed doubles?
- It doubles
- It quadruples
+ It rises eightfold
= Drag force goes as v², and power is force times speed, so power goes as v³.
:::
