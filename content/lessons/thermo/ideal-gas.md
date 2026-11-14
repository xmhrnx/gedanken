Gases are the simplest form of matter to describe. Squeeze one, heat it or add more of it, and its pressure, volume and temperature change in a way that is almost the same for air, helium, carbon dioxide or steam. Three experimental laws, found over 150 years, combine into one equation.

## Three experimental laws

**Boyle's law (1662).** At fixed temperature, the pressure of a fixed amount of gas is inversely proportional to its volume: halve the volume and the pressure doubles. PV = constant.

**Charles's law (1787, published by Gay-Lussac in 1802).** At fixed pressure, volume grows in proportion to temperature. Measured in Celsius, the graph of volume against temperature is a straight line, and for every gas the line, extended backwards, reaches zero volume at about −273 °C. That point is absolute zero, and measuring from it, in kelvin, gives V/T = constant.

**Avogadro's hypothesis (1811).** Equal volumes of any gases, at the same temperature and pressure, contain the same number of molecules. So at fixed P and T, volume is proportional to the amount of gas.

## Combining them

Amount of gas is counted in *moles*: one mole is 6.022 × 10²³ molecules (Avogadro's number, N_A), chosen so that a mole of carbon-12 weighs 12 grams.

::: steps From three laws to PV = nRT
$ P_1 V_1 = P_2 V'
Take a fixed amount of gas from state (P₁, V₁, T₁) to (P₂, V₂, T₂) in two stages. First change the pressure to P₂ at constant temperature T₁; by Boyle's law the volume becomes some V'.
---
$ \frac{V'}{T_1} = \frac{V_2}{T_2}
Then heat or cool it to T₂ at constant pressure P₂; by Charles's law the volume goes from V' to V₂.
---
$ \frac{P_1 V_1}{T_1} = \frac{P_2 V'}{T_1} = \frac{P_2 V_2}{T_2}
Divide the first equation by T₁, then multiply the second by P₂ and substitute. The intermediate volume V' has dropped out.
---
$ \frac{PV}{T} = nR
Since the two states were arbitrary, PV/T is the same in every state of a given amount of gas. By Avogadro, doubling the amount doubles V at fixed P and T, so the constant is proportional to the number of moles n. Call the proportionality constant R.
:::

::: equation The ideal gas law
PV = nRT
:::

- **P** is pressure, in pascals.
- **V** is volume, in cubic metres.
- **n** is the amount of gas, in moles.
- **T** is the absolute temperature, in kelvin. Celsius does not work: 0 °C is not "no temperature".
- **R** = 8.314 J/(mol·K), the *gas constant*, the same for every gas.

Physicists often count molecules instead of moles. With N = nN_A molecules, the law reads PV = NkT, where k = R/N_A = 1.381 × 10⁻²³ J/K is *Boltzmann's constant*. It will reappear in kinetic theory and entropy.

## Using it

One mole of any gas at 0 °C (273.15 K) and 1 atmosphere (101 325 Pa) takes up V = nRT/P = 8.314 × 273.15 / 101 325 ≈ 0.0224 m³, or 22.4 litres. That is a box about 28 cm on a side, whether it holds 2 grams of hydrogen or 44 grams of carbon dioxide.

A car tyre pumped to 300 kPa absolute on a cold 0 °C morning warms to 30 °C on the motorway. Its volume barely changes, so P/T is constant: P = 300 × 303/273 ≈ 333 kPa.

::: model gas
Raise the temperature at fixed volume. The molecules speed up and the pressure rises in proportion to the kelvin temperature.
:::

## Why "ideal"

No real gas obeys the law exactly. It describes a gas whose molecules are points that never attract or repel except in brief collisions. Real molecules have size and pull weakly on one another, which matters when they are crowded together: at high pressure, or near the temperature where the gas would condense to a liquid. Air at room conditions departs from the ideal law by less than 0.1%. Steam close to boiling, or any gas at hundreds of atmospheres, departs much more.

::: history 1662 · Oxford
Robert Boyle, with his assistant Robert Hooke, trapped air in the short closed arm of a J-shaped glass tube and poured mercury into the long arm. When the added column of mercury matched the atmosphere's own 29 inches, doubling the total pressure, the trapped air shrank to half its volume. Boyle explained it with a picture of air as tiny coiled springs, which was wrong, but his measurements were good. Jacques Charles, a balloonist, studied the effect of temperature in the 1780s without publishing; Joseph Louis Gay-Lussac measured it carefully in 1802 and credited him.
:::

::: note
At low density, every gas obeys the same law, PV = nRT. What the gas is made of does not matter, only how many molecules there are. The next lesson explains why.
:::

::: quiz
? A sealed rigid container of gas is heated from 27 °C to 327 °C. Its pressure
- increases by a factor of about 12
+ doubles
- stays the same
= In kelvin the temperature goes from 300 K to 600 K. At fixed volume P ∝ T, so pressure doubles.
? Two identical balloons at the same temperature and pressure hold helium and carbon dioxide. Which contains more molecules?
- Helium, because its molecules are lighter
- Carbon dioxide, because its molecules are bigger
+ They contain the same number
= PV = NkT does not depend on the type of gas: equal P, V and T mean equal N.
? A gas is compressed to a third of its volume while its kelvin temperature doubles. Its pressure changes by a factor of
- 2/3
- 3/2
+ 6
= P = nRT/V: doubling T doubles P, and dividing V by 3 triples it again.
:::
