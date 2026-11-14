Solids and liquids are complicated: each substance expands, compresses and stores energy in its own way. Gases are different. Squeeze a gas, heat it or pump more of it into a container, and its pressure, volume and temperature change in almost exactly the same way whether it is air, helium, carbon dioxide or steam. That sameness is a clue that something simple is going on underneath. This lesson follows how three separate experimental rules, found over a century and a half, were combined into one equation, the ideal gas law, and what each part of that equation means.

## What did the experiments find?

The three quantities that describe a sample of gas are its **pressure** P, the force it exerts on each square metre of its container's walls; its **volume** V, the space it fills; and its **temperature** T. The early experimenters held one of these fixed and watched how the other two were related.

The first rule came from Robert Boyle in 1662. He kept the temperature of a trapped sample of air fixed and changed the pressure on it. The volume responded in inverse proportion: double the pressure and the air shrinks to half its volume; triple it and the volume falls to a third. In other words, the product PV stays the same. This is **Boyle's law**. It already says something surprising, because it holds with no adjustment for what the gas is.

The second rule concerns temperature. In the 1780s the French balloonist Jacques Charles found that if a gas is kept at constant pressure and warmed, its volume grows steadily. Joseph Louis Gay-Lussac measured the effect carefully and published it in 1802, crediting Charles. If you plot volume against temperature in degrees Celsius, the points fall on a straight line. The striking part is what happens when you extend those lines backwards to colder temperatures than anyone could reach. For every gas, the line hits zero volume at the same place, about −273 °C. A gas cannot have less than zero volume, so this point marks a floor: it is absolute zero, the starting point of the kelvin scale. Measured in kelvin, the straight line passes through the origin, and the rule takes the simple form that V/T is constant. This is **Charles's law**.

The third idea was a bold guess. In 1811 Amedeo Avogadro, trying to make sense of the simple whole-number ratios in which gases combine chemically, proposed that equal volumes of any gases, at the same temperature and pressure, contain the same number of molecules. If that is true, then at fixed P and T, the volume of a gas depends only on how many molecules it contains, and is proportional to that number. **Avogadro's hypothesis** was ignored for nearly fifty years and then became one of the foundations of chemistry.

## How do we count molecules?

Molecules are far too small and numerous to count one by one, so chemists count them in bundles called **moles**. One mole is exactly 6.02214076 × 10²³ particles, a number called **Avogadro's number**, N_A. The size of the bundle was chosen so that a mole of carbon-12 atoms has a mass of 12 grams (since 2019 the number is fixed by definition, and the 12 grams is very nearly but no longer exactly true). The practical result is that a mole of any substance has a mass in grams roughly equal to the total of its atoms' mass numbers: about 2 g for hydrogen gas, 28 g for nitrogen, 44 g for carbon dioxide. The amount of gas, in moles, is written n.

## How do three laws become one?

Each experimental law holds one quantity fixed. A real change of state, such as a tyre heating up on the road, usually changes everything at once. The trick is to break any such change into stages in which only two quantities vary at a time, so that one of the known laws applies to each stage.

::: steps From three laws to PV = nRT
$ P_1 V_1 = P_2 V'
Take a fixed amount of gas from a starting state (P₁, V₁, T₁) to any final state (P₂, V₂, T₂). Do it in two stages, beginning by changing the pressure from P₁ to P₂ while keeping the temperature at T₁. Boyle's law applies to this stage because the temperature is fixed, so PV is unchanged, and the volume ends at some intermediate value V' that we do not yet know.
---
$ \frac{V'}{T_1} = \frac{V_2}{T_2}
Now keep the pressure at P₂ and heat or cool the gas from T₁ to T₂. Charles's law applies because the pressure is fixed, so V/T stays constant, and the volume moves from V' to the final V₂. The gas has reached the final state, and every step was covered by a measured law.
---
$ \frac{P_1 V_1}{T_1} = \frac{P_2 V'}{T_1}
Divide both sides of the first equation by T₁. That is allowed because we do the same thing to both sides, and it is useful because it creates the combination V'/T₁ on the right, which the second equation tells us about.
---
$ \frac{P_2 V'}{T_1} = \frac{P_2 V_2}{T_2}
Multiply both sides of the second equation by P₂. Its left side is now exactly the right side of the previous line, so the two equations can be chained together.
---
$ \frac{P_1 V_1}{T_1} = \frac{P_2 V_2}{T_2}
Join the chain. The unknown intermediate volume V' has disappeared, which is the purpose of the two-stage route. Since the starting and final states were arbitrary, the quantity PV/T must have the same value in every state of a given amount of gas.
---
$ \frac{PV}{T} = nR
Finally, ask what the constant depends on. By Avogadro's hypothesis, doubling the amount of gas at fixed pressure and temperature doubles the volume, so PV/T doubles too. The constant is therefore proportional to the number of moles n, and the constant of proportionality, called R, turns out to be the same for every gas.
:::

Multiplying through by T gives the result in its familiar form.

::: equation The ideal gas law
PV = nRT
:::

Each letter carries a unit, and the equation only works if they are consistent. Pressure is in pascals (newtons per square metre), volume in cubic metres, n in moles and T in kelvin. Temperature must be absolute: 0 °C is not "no temperature", and a formula that says pressure is proportional to temperature would predict nonsense if fed Celsius values. The **gas constant** R is 8.314 J/(mol·K).

Physicists often prefer to count individual molecules rather than moles. If N is the number of molecules, then N = nN_A, and the law becomes PV = NkT, where k = R/N_A ≈ 1.381 × 10⁻²³ J/K is **Boltzmann's constant**. It is the gas constant for a single molecule, and it will turn up again when we connect temperature to molecular motion and when we define entropy.

## What does the law predict?

Here is the first consequence. One mole of any gas at 0 °C (273.15 K) and normal atmospheric pressure (101 325 Pa) occupies V = nRT/P = 8.314 × 273.15 / 101 325 ≈ 0.0224 m³, or 22.4 litres. That is a cube about 28 cm on a side. The striking part is that the gas does not matter. The cube holds 2 grams of hydrogen or 44 grams of carbon dioxide, and the volume is the same, because the law counts only molecules, not what they are.

A second example shows how to use the law when only some things change. A car tyre is pumped to 300 kPa (absolute pressure, which includes the atmosphere's share) on a cold morning at 0 °C, and warms to 30 °C on a long drive. The rubber barely stretches, so the volume and the amount of air are fixed, and PV = nRT reduces to P/T staying constant. In kelvin, the temperature rises from 273 K to 303 K, so the pressure becomes 300 × 303/273 ≈ 333 kPa. A 30-degree warming raises the pressure by about 11%, which is why tyre pressures should be checked cold.

::: model gas
Raise the temperature while the volume stays fixed. The molecules speed up and the pressure rises in proportion to the kelvin temperature, just as P/T = constant predicts.
:::

## Why is it called "ideal"?

No real gas obeys this law exactly. The law describes a gas whose molecules behave as points that take up no room and exert no forces on each other except during brief collisions. Real molecules have a small but definite size, and they attract one another weakly when they are close. Neither effect matters much when the molecules are far apart, which is the usual situation: at room temperature and pressure, the molecules in air are separated by about ten times their own diameter, and air departs from the ideal law by less than 0.1%.

The effects grow when molecules are crowded together. At hundreds of atmospheres, the molecules' own volume becomes a noticeable fraction of the container, and the gas is harder to compress than the law predicts. Near the temperature where a gas would condense into a liquid, the attractions between molecules pull them together and the pressure falls below the ideal value. Steam just above its boiling point at atmospheric pressure, for example, deviates by one or two percent, and much more at the high pressures inside a power-station boiler. The ideal gas law is best understood as the limit that every real gas approaches as it becomes thin, and the next lesson explains why that limit is the same for all of them.

::: history 1662 · Oxford
Robert Boyle, working in Oxford with his young assistant Robert Hooke, trapped a pocket of air in the short, sealed arm of a J-shaped glass tube, then poured mercury into the long, open arm. The extra mercury pressed on the trapped air, and Boyle measured how much it shrank. When the added column was about 29 inches tall, matching the pressure of the atmosphere itself and so doubling the total pressure on the air, the air was squeezed to half its original volume. Boyle explained the springiness of air by imagining it made of tiny coiled particles like springs or fleece. That picture was wrong; air resists compression because its molecules are moving and bombarding the walls, as the next lesson shows. But his measurements were careful and his law has lasted.
:::

::: note
At low density every gas obeys the same law, PV = nRT. The type of molecule does not matter, only how many there are, which suggests the law comes from something all molecules share. The next lesson shows that this shared property is their motion.
:::

::: quiz
? A sealed rigid container of gas is heated from 27 °C to 327 °C. What happens to its pressure?
- It increases by a factor of about 12
+ It doubles
- It stays the same
= In kelvin the temperature rises from 300 K to 600 K, a factor of 2. At fixed volume P is proportional to the absolute temperature, so the pressure doubles; using Celsius would wrongly suggest a factor of 12.
? Two identical balloons at the same temperature and pressure hold helium and carbon dioxide. Which contains more molecules?
- Helium, because its molecules are lighter
- Carbon dioxide, because its molecules are bigger
+ They contain the same number
= The law PV = NkT contains no property of the gas itself. Equal pressure, volume and temperature therefore mean an equal number of molecules, which is Avogadro's hypothesis.
? A gas is compressed to a third of its volume while its kelvin temperature doubles. By what factor does its pressure change?
- 2/3
- 3/2
+ 6
= From P = nRT/V, doubling T doubles the pressure, and dividing V by 3 triples it again, giving a factor of 2 × 3 = 6.
:::
