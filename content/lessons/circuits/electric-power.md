Your electricity bill does not charge you for current, or voltage, or charge. It charges you for energy, measured in a unit called the kilowatt-hour. Understanding that unit, and how fast each appliance uses energy, takes one formula.

## Power is current times voltage

Voltage is energy per coulomb; current is coulombs per second. Multiply them and the coulombs cancel, leaving energy per second, which is power.

::: equation Electrical power
P = IV
:::

Power is measured in **watts** (W): 1 W is 1 J/s, and 1 A × 1 V = 1 W. This holds for any device: a motor, a heater, a phone charger, a battery being charged. For a resistor, where all the energy becomes heat, Ohm's law gives the equivalent forms P = I²R and P = V²/R, as the previous lesson showed.

Typical figures: an LED bulb 5–10 W, a laptop 30–60 W, a fridge averaging about 50 W over a day as its compressor cycles on and off, a microwave about 1 kW, a kettle 2–3 kW, an electric car charging at home about 7 kW. A 2 kW kettle on 230 V mains draws I = P/V ≈ 8.7 A; on 120 V, the same kettle would need about 17 A, which is one reason kettles in North America are usually less powerful and slower.

## The kilowatt-hour

The joule is too small for household use. Electricity companies use the **kilowatt-hour** (kWh): the energy used by a 1 kW device running for one hour.

::: equation The kilowatt-hour in joules
1\,\text{kWh} = 1000\,\text{W} \times 3600\,\text{s} = 3.6 \times 10^6\,\text{J}
:::

Energy is power times time, so a bill is easy to estimate. Boiling a 2 kW kettle for 3 minutes uses 2 kW × 0.05 h = 0.1 kWh. A fridge at an average 50 W uses about 1.2 kWh a day. A 7 kW car charger running 6 hours uses 42 kWh. Multiply by the price per kWh, which varies widely by country and year but is often between 10 and 40 cents, and you have the cost. Heating and cooling usually dominate a home's use; a phone charger, using roughly 5 Wh a day, barely registers.

## Mains AC and RMS

Mains voltage alternates. In Europe it swings between about +325 V and −325 V fifty times a second; in North America between ±170 V sixty times a second. The quoted figures, 230 V and 120 V, are **root-mean-square** (RMS) values: the steady DC voltage that would deliver the same average power to a heater. For a sine wave the RMS value is the peak divided by √2. Using RMS values, P = IV gives average power just as for DC.

## Why power lines use high voltage

A power station must deliver a fixed power P to a distant city through cables with resistance R. The cables waste some of it as heat.

::: steps Power lost in transmission
$ I = \frac{P}{V}
To deliver power P at transmission voltage V, the line must carry current I = P/V.
---
$ P_{\text{loss}} = I^2 R
The cables, with resistance R, dissipate heat at the rate I²R, like any resistor.
---
$ P_{\text{loss}} = \frac{P^2 R}{V^2}
Substitute for I. For fixed delivered power and fixed cables, the loss falls as the square of the voltage.
---
$ \frac{P_{\text{loss}}}{P} = \frac{PR}{V^2}
As a fraction of the power sent, the loss also falls as 1/V². Raising the voltage tenfold cuts the fractional loss a hundredfold.
:::

So long-distance lines run at hundreds of thousands of volts, typically 110 kV to 765 kV, and overall transmission and distribution losses are usually 5–10%. But nobody wants 400,000 V in their kitchen. The voltage must be stepped up at the power station and down again near homes. **Transformers**, explained in the magnetism chapter, do this efficiently, but they work only with alternating current. That fact decided a famous commercial fight.

::: history 1880s–1890s · United States
Thomas Edison opened the Pearl Street Station in New York in 1882, supplying direct current at about 110 V. At that voltage the losses limited each station to customers within a mile or so. George Westinghouse bought the patents of Nikola Tesla's AC induction motor in 1888 and backed alternating current, which transformers could step up for long-distance lines. Edison's camp campaigned on the dangers of AC, but its economics won: the Niagara Falls power plant, built by Westinghouse, began sending AC to Buffalo in 1896. Very long high-voltage DC lines have since returned for special cases, using modern electronics to convert.
:::

::: note
Power is current times voltage, P = IV. Energy is power times time, billed in kilowatt-hours (3.6 MJ each). Transmitting power at high voltage keeps the current, and so the I²R loss, small.
:::

::: quiz
? A 1.5 kW heater runs for 4 hours. How much energy does it use?
- 0.375 kWh
+ 6 kWh
- 6000 kWh
= Energy = power × time = 1.5 kW × 4 h = 6 kWh.
? A power line's voltage is doubled while the delivered power stays the same. The heat lost in the line
- doubles
- halves
+ falls to a quarter
= The current halves, and the loss goes as I², so it falls by a factor of 4.
? A device draws 0.5 A from a 12 V supply. Its power is
+ 6 W
- 24 W
- 0.04 W
= P = IV = 0.5 A × 12 V = 6 W.
:::
