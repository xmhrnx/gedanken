Turning work into heat is easy: rub two sticks together. Turning heat into work is what drove the Industrial Revolution, and it turns out to have a hard limit. No engine, however well made, can convert all the heat it takes in into work. The limit depends only on temperatures.

## What a heat engine does

A *heat engine* runs in a cycle, returning to its starting state again and again. Each cycle it takes heat Q_H from a hot reservoir (burning fuel, a reactor core), does work W, and dumps the rest, Q_C, into a cold reservoir (a river, the air). Because the engine returns to its starting state, its internal energy does not change over a cycle, and the first law gives W = Q_H − Q_C.

::: equation Efficiency of a heat engine
\eta = \frac{W}{Q_H} = 1 - \frac{Q_C}{Q_H}
:::

The *efficiency* η is the fraction of the heat taken in that comes out as work. To reach 100% the engine would have to dump no heat at all. Every real engine does dump heat: a car's radiator and exhaust, a power station's cooling towers. The question is whether that is a flaw of engineering or a law of nature.

## The Carnot cycle

Sadi Carnot answered it by imagining the most perfect engine possible. It has no friction and no heat leaks, and it only lets heat flow across tiny temperature differences, so every step could be run backwards. Such an engine is called *reversible*. With an ideal gas in a cylinder as the working substance, the cycle has four stages:

1. **Isothermal expansion** at T_H: the gas touches the hot reservoir, absorbs Q_H and expands from volume V_a to V_b.
2. **Adiabatic expansion**: insulated, the gas keeps expanding to V_c, doing work and cooling to T_C.
3. **Isothermal compression** at T_C: touching the cold reservoir, the gas is compressed to V_d and gives up Q_C.
4. **Adiabatic compression**: insulated again, it is compressed back to V_a, warming to T_H.

## Carnot's efficiency

::: steps Efficiency of a Carnot engine
$ Q_H = nRT_H \ln\frac{V_b}{V_a}
In stage 1 the temperature is fixed, so the ideal gas's internal energy does not change and Q_H equals the work done. Using P = nRT/V in W = ∫P dV gives nRT ∫dV/V = nRT ln(V_b/V_a).
---
$ Q_C = nRT_C \ln\frac{V_c}{V_d}
In stage 3 the same reasoning gives the heat released as the gas is compressed from V_c to V_d.
---
$ T_H V_b^{\gamma - 1} = T_C V_c^{\gamma - 1}, \quad T_H V_a^{\gamma - 1} = T_C V_d^{\gamma - 1}
On an adiabat, applying the first law with Q = 0 shows that TV^(γ−1) stays constant, where γ = C_P/C_V. Stage 2 links state b to c, and stage 4 links d back to a.
---
$ \frac{V_b}{V_a} = \frac{V_c}{V_d}
Divide the first adiabat equation by the second: the temperatures cancel, leaving (V_b/V_a)^(γ−1) = (V_c/V_d)^(γ−1).
---
$ \frac{Q_C}{Q_H} = \frac{T_C}{T_H}
Divide the two heats; the equal logarithms cancel, as do n and R.
---
$ \eta_{\text{Carnot}} = 1 - \frac{T_C}{T_H}
Substitute into η = 1 − Q_C/Q_H. Temperatures must be in kelvin.
:::

The ideal gas was only a convenient working substance. Carnot's deeper point was that **no engine working between the same two temperatures can beat a reversible one**, whatever it is made of. Suppose some engine did. Let it drive a Carnot engine running backwards as a refrigerator, supplying exactly the work the refrigerator needs. Being more efficient, it takes less heat from the hot reservoir to make that work than the refrigerator delivers back to it. Taken together, the pair would move heat from the cold reservoir to the hot one with no outside work and nothing else changing. Heat never flows spontaneously from cold to hot, so no such engine exists. That principle is the second law, the subject of the next lesson.

::: note
No heat engine can be more efficient than 1 − T_C/T_H. The limit is set by the temperatures alone, and it reaches 100% only if the cold side is at absolute zero.
:::

## Real engines

A modern coal or nuclear power station heats steam to about 840 K (570 °C) and condenses it near 300 K. Its Carnot limit is 1 − 300/840 ≈ 64%; real plants achieve about 35–45%, losing the rest to friction, heat leaks and steps that are far from reversible. A car engine manages roughly 25–35%. To do better, engineers raise the hot temperature, limited by what materials can stand, or stack one cycle above another, as combined-cycle gas plants do to exceed 60%.

The same theory runs backwards. A refrigerator or heat pump uses work to move heat from cold to hot. A heat pump warming a house to 293 K from outside air at 273 K could, in the ideal limit, deliver T_H/(T_H − T_C) ≈ 15 joules of heat for each joule of electricity. Real ones deliver 3 or 4, still much better than an electric heater's 1.

::: history 1824 · Paris
Sadi Carnot, a young French army engineer, published *Reflections on the Motive Power of Fire* in 1824, convinced that steam power lay behind Britain's industrial and military strength. He compared heat falling from hot to cold with water falling through a mill wheel, and believed, like most of his contemporaries, that heat was a conserved fluid, caloric. That was wrong: an engine turns some of the heat into work, so less heat reaches the cold side than left the hot. Yet his conclusion, that the best efficiency depends only on the two temperatures, survived. The book was almost unread until Émile Clapeyron restated it with pressure–volume diagrams in 1834. Carnot died of cholera in 1832, aged 36. Around 1850 Kelvin and Clausius reconciled his argument with energy conservation.
:::

::: quiz
? An engine takes in heat at 600 K and rejects it at 300 K. Its greatest possible efficiency is
+ 50%
- 100%
- 200%
= η = 1 − 300/600 = 0.5. Real engines between these temperatures do worse.
? Which change would raise the Carnot efficiency of an engine most?
- Using helium instead of air as the working gas
+ Raising the temperature of the hot reservoir
- Making the cylinder larger
= Carnot's limit depends only on T_H and T_C, not on the working substance or the size of the engine.
? Why must every heat engine reject some heat to a cold reservoir?
- Because of friction, which better engineering could eliminate
+ Because even a perfect reversible engine has Q_C/Q_H = T_C/T_H, which is zero only at absolute zero
- Because the first law forbids converting heat to work
= The first law allows full conversion; the limit comes from Carnot's argument, which rests on the second law.
:::
