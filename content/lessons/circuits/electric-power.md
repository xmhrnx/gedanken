An electricity bill does not charge you for current, or for voltage, or for charge. It charges you for energy, measured in a unit called the kilowatt-hour. To understand that unit, and why a kettle costs more to run than a phone charger, you need to know how fast each device uses energy, and that takes just one formula. The same formula answers a less obvious question with large consequences: why the cables that carry electricity across a country run at hundreds of thousands of volts, and why that fact decided the outcome of a famous commercial battle.

## How fast does a device use energy?

The rate at which energy is transferred is called **power**, and it is measured in **watts** (W), where one watt is one joule per second. To find the electrical power going into a device, we need to combine two quantities from earlier lessons. Voltage tells us how much energy each coulomb of charge gives up as it passes through the device: one volt means one joule per coulomb. Current tells us how many coulombs pass through each second: one ampere means one coulomb per second. Multiply them, and the coulombs cancel, leaving joules per second, which is power.

::: equation Electrical power
P = IV
:::

So a current of 1 A flowing through a voltage of 1 V delivers 1 W. This formula holds for any device at all: a motor, a heater, a phone charger, even a battery that is being charged, because it only uses the meanings of current and voltage and says nothing about what happens to the energy afterwards. In a resistor all the energy becomes heat, and combining P = IV with Ohm's law gives the equivalent forms P = I²R and P = V²/R, as the previous lesson showed.

Some typical figures give a feel for the scale. An LED bulb uses 5 to 10 W, a laptop 30 to 60 W, and a fridge averages about 50 W over a day, because its compressor switches on and off. A microwave oven uses about 1 kW (a thousand watts), a kettle 2 to 3 kW, and an electric car charging at home about 7 kW. The devices that use the most power are almost always the ones that heat things.

The formula also tells you the current each device draws. A 2 kW kettle on 230 V mains, as in Europe, draws I = P/V = 2000 ÷ 230 ≈ 8.7 A. The same kettle on 120 V mains, as in North America, would need about 17 A, more than ordinary household circuits there are built to supply. That is one reason kettles in North America are usually less powerful, and slower, than European ones.

## How is the energy counted?

The joule is a small unit for household purposes: a kettle uses thousands of them every second. So electricity companies use a larger unit, the **kilowatt-hour** (kWh), defined as the energy used by a device of 1 kW running for one hour.

::: equation The kilowatt-hour in joules
1\,\text{kWh} = 1000\,\text{W} \times 3600\,\text{s} = 3.6 \times 10^6\,\text{J}
:::

The calculation uses the fact that energy is power multiplied by time, with the hour converted to 3,600 seconds. One kilowatt-hour is 3.6 million joules, or 3.6 megajoules.

Because energy is just power times time, estimating a bill is straightforward once you know the power of each device and how long it runs. Boiling a 2 kW kettle for 3 minutes, which is 0.05 hours, uses 2 kW × 0.05 h = 0.1 kWh. A fridge averaging 50 W, running all 24 hours of a day, uses 0.05 kW × 24 h = 1.2 kWh. A 7 kW car charger running for 6 hours uses 42 kWh. To get the cost, multiply by the price of a kilowatt-hour, which varies widely between countries and from year to year but is often somewhere between 10 and 40 cents.

These numbers explain where a household's electricity really goes. A phone charger, which puts something under 20 watt-hours into a phone each day, barely registers, while heating, cooling and hot water usually dominate. The difference is not about how many devices you own but about how much power each draws and for how long.

## What does a voltage mean for mains AC?

There is a complication with mains electricity, because its voltage does not stay fixed. It is alternating current, so the voltage swings smoothly back and forth like a sine wave. In Europe it rises to about +325 V, falls through zero to about −325 V, and returns, fifty times a second. In North America it swings between about +170 V and −170 V, sixty times a second.

Why, then, are these supplies called 230 V and 120 V? Because quoting the peak would be misleading: the voltage only touches its peak for an instant in each cycle, and spends much of its time far lower. What matters for a heater is the average power it delivers. So the quoted figure is chosen to be the steady, unchanging voltage that would deliver the same average power to a heater. This is called the **root-mean-square** or RMS value, after the way it is calculated: square the voltage at each moment (since power goes as V²), take the average over a cycle, and then take the square root. For a sine wave the result is the peak divided by √2, so 325 V ÷ 1.41 ≈ 230 V. The convenience of RMS values is that with them, P = IV gives the average power for AC just as it does for DC.

## Why do power lines use such high voltages?

Now to the large-scale question. A power station must deliver a certain power P to a city many kilometres away. The cables between them, however thick, have some resistance R, and like any resistor they turn some of the electrical energy passing through them into heat. That heat is simply lost. The engineer's question is how to keep the loss small.

::: steps Power lost in transmission
$ I = \frac{P}{V}
To deliver power P at a transmission voltage V, the line must carry a current I chosen so that IV = P. Dividing both sides by V gives the current. For a fixed power, a higher voltage means a smaller current.
---
$ P_{\text{loss}} = I^2 R
The cables carry this current through their resistance R, so they heat up at the rate I²R, like any resistor. This is the form to use here because the current in the cable is what we know; the voltage across the cable itself is only a small part of V.
---
$ P_{\text{loss}} = \left(\frac{P}{V}\right)^2 R = \frac{P^2 R}{V^2}
Replace I with P/V from the first step and square it. The delivered power P and the cable resistance R are fixed by the job and the cables. What is left for the engineer to choose is V, and the loss falls with the square of it.
---
$ \frac{P_{\text{loss}}}{P} = \frac{PR}{V^2}
Divide both sides by P to get the loss as a fraction of the power sent. This is the number that matters economically, and it too falls as 1/V². Raising the voltage tenfold cuts the fraction lost a hundredfold.
:::

That square is the whole story. Doubling the voltage halves the current needed, and since the heating goes as the square of the current, it cuts the loss to a quarter. So long-distance lines run at very high voltages, typically 110,000 V to 765,000 V, and the total losses in transmission and local distribution are usually 5 to 10% of the electricity generated.

But nobody wants 400,000 V in their kitchen. The voltage has to be raised at the power station and lowered again in stages before it reaches homes. The device that does this efficiently is the **transformer**, explained in the magnetism chapter, and it has one crucial limitation: it works only with alternating current. That single fact decided one of the most famous commercial contests of the nineteenth century.

::: history 1880s–1890s · United States
In 1882 Thomas Edison opened the Pearl Street Station in New York, the first central power station in the United States, supplying direct current at about 110 V to light the offices of lower Manhattan. At such a low voltage the currents were large and the cable losses severe, so each station could only serve customers within a mile or so; lighting a whole city with DC would have meant a power station every few blocks. George Westinghouse saw that alternating current, which transformers could raise to high voltage for transmission and lower again for use, did not have this limit. In 1888 he bought the patents for Nikola Tesla's AC induction motor, which made AC useful for machinery as well as lighting. Edison's camp fought back with a public campaign about the dangers of high-voltage AC, but the economics were against them. The decisive demonstration was the hydroelectric plant at Niagara Falls, equipped by Westinghouse, which began sending AC to Buffalo, about 40 kilometres away, in 1896. AC became the standard for grids worldwide. Very long high-voltage DC lines have returned in recent decades for special cases, such as undersea cables, now that modern electronics can convert between AC and DC efficiently.
:::

::: note
Electrical power is current times voltage, P = IV, because voltage is energy per coulomb and current is coulombs per second. Energy is power multiplied by time and is billed in kilowatt-hours of 3.6 MJ each. Power is sent over long distances at high voltage because that keeps the current small, and the heat lost in the cables goes as the square of the current.
:::

::: quiz
? A 1.5 kW heater runs for 4 hours. How much energy does it use?
- 0.375 kWh
+ 6 kWh
- 6000 kWh
= Energy is power multiplied by time, so 1.5 kW × 4 h = 6 kWh. Keeping the power in kilowatts and the time in hours gives the answer directly in kilowatt-hours.
? A power line's voltage is doubled while the power it delivers stays the same. What happens to the heat lost in the line?
- It doubles
- It halves
+ It falls to a quarter
= To carry the same power at twice the voltage, the line needs only half the current. The heat lost goes as I²R, so halving the current divides the loss by four.
? A device draws 0.5 A from a 12 V supply. What is its power?
+ 6 W
- 24 W
- 0.04 W
= Power is current times voltage, so P = 0.5 A × 12 V = 6 W.
:::
