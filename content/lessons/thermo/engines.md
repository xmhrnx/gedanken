Turning work into heat is easy. Rub your hands together, or press the brakes on a bicycle, and all of the work becomes internal energy without any effort to arrange it. Turning heat into work is a different matter. It is what a steam engine, a car engine and a power station do, and it drove the Industrial Revolution. In the early nineteenth century engineers improved engines by trial and error, and nobody knew whether there was a limit to how good they could become. There is, and it is a law of nature rather than a shortcoming of engineering: no engine, however well made, can turn all the heat it takes in into work. This lesson shows why, and finds the limit, which depends on nothing but two temperatures.

## What does every heat engine have in common?

A steam locomotive and a jet engine look nothing alike, but they share a pattern. A **heat engine** is any device that runs in a **cycle**, returning to its starting state over and over, and on each cycle takes in heat, does some work, and gets rid of the rest. The heat comes from a hot **reservoir**, meaning a source large enough that its temperature does not change noticeably: burning fuel, or a reactor core. The heat that is not turned into work goes to a cold reservoir, usually a river, the sea or the air.

Call the heat taken in from the hot reservoir in one cycle Q_H, the work done W, and the heat dumped into the cold reservoir Q_C. Because the engine ends each cycle in the same state it started, its internal energy is the same at the end as at the start. The first law then says that the energy in must equal the energy out: W = Q_H − Q_C.

What matters to the owner of an engine is how much work they get for the fuel they burn. That fraction is the **efficiency**, written η (the Greek letter eta).

::: equation Efficiency of a heat engine
\eta = \frac{W}{Q_H} = 1 - \frac{Q_C}{Q_H}
:::

The second form comes from replacing W with Q_H − Q_C. It shows that an efficiency of 100% would require Q_C = 0: an engine that dumps no heat at all. Every real engine does dump heat. A car has a radiator and a hot exhaust, and a power station has cooling towers or a warm outflow into a river, and in each case more energy leaves as waste heat than as useful work. The question that Sadi Carnot asked in 1824 was whether this waste is a defect that cleverer engineering could remove, or something unavoidable.

## What would a perfect engine look like?

Carnot's approach was to imagine the best engine that could possibly exist and find out how well it would do. If even that engine must waste heat, then the waste is a law of nature.

His ideal engine has no friction and no heat leaks. It also has a subtler property. Whenever heat flows into or out of it, the engine is at almost exactly the same temperature as the reservoir it touches, so heat crosses only a tiny temperature difference. That matters because heat flowing across a large temperature difference, from a flame straight into cold water, for instance, is a one-way process: it cannot be undone without outside help. An engine that avoids every such step could, in principle, be run backwards, with every flow of heat and work reversed. Physicists call such an engine **reversible**.

To make the idea concrete, Carnot's cycle can be run with an ideal gas in a cylinder as the working substance. It has four stages, and it helps to follow the gas through them in order.

In the first stage, the cylinder is placed against the hot reservoir at temperature T_H, and the gas is allowed to expand slowly from volume V_a to V_b. As it expands it does work on the piston, which would cool it, but heat Q_H flows in from the reservoir and keeps its temperature fixed. This is an **isothermal** expansion: one at constant temperature.

In the second stage, the cylinder is taken off the reservoir and insulated, and the gas keeps expanding, from V_b to V_c. Now no heat can enter, so the work it does comes out of its own internal energy and it cools, until it reaches the cold temperature T_C. A process with no heat flow is called **adiabatic**.

In the third stage, the cylinder is placed against the cold reservoir and the piston pushes the gas back in, from V_c to V_d. Compression would warm the gas, but heat Q_C flows out into the cold reservoir and keeps it at T_C.

In the fourth stage, the cylinder is insulated again and the gas is compressed adiabatically from V_d back to V_a. The work done on it raises its temperature back to T_H, and the gas is exactly where it began.

Over the whole cycle, the gas did more work while expanding (when it was hot and its pressure high) than was done on it while being compressed (when it was cool and its pressure low). The difference is the net work W.

## How efficient is Carnot's engine?

To find the efficiency, we need Q_H and Q_C. The heat flows happen only in the two isothermal stages, so those are where to look, and the two adiabatic stages will supply a link between the volumes.

::: steps Efficiency of a Carnot engine
$ Q_H = nRT_H \ln\frac{V_b}{V_a}
In the first stage the temperature is fixed, and an ideal gas's internal energy depends only on temperature, so ΔU = 0 and the first law gives Q_H = W. The work is ∫P dV, and with P = nRT_H/V from the ideal gas law it becomes nRT_H ∫dV/V. The integral of 1/V is the natural logarithm, so the work, and therefore the heat absorbed, is nRT_H ln(V_b/V_a).
---
$ Q_C = nRT_C \ln\frac{V_c}{V_d}
In the third stage the same argument applies at temperature T_C as the gas is compressed from V_c to V_d. Writing the ratio as V_c/V_d makes Q_C a positive number, the amount of heat given up to the cold reservoir.
---
$ nC_V\,\mathrm{d}T = -P\,\mathrm{d}V
Now the adiabatic stages. With no heat flow, the first law says the change in internal energy equals minus the work done. For an ideal gas, a temperature change dT changes the internal energy by nC_V dT, where C_V is the molar heat capacity at constant volume, and the work in a small expansion is P dV. So the gas's temperature falls in step with the work it does.
---
$ \frac{\mathrm{d}T}{T} = -\frac{R}{C_V}\frac{\mathrm{d}V}{V} = -(\gamma - 1)\frac{\mathrm{d}V}{V}
Replace P with nRT/V and divide both sides by nC_V T, which separates the temperature terms from the volume terms. From the previous lesson, C_P − C_V = R, so R/C_V = C_P/C_V − 1. The ratio C_P/C_V is written γ (gamma), which is 5/3 for a monatomic gas.
---
$ T V^{\gamma - 1} = \text{constant}
Integrate both sides: ln T = −(γ − 1) ln V + constant, which is the same as ln(TV^(γ−1)) = constant. So along any adiabat, the combination TV^(γ−1) stays fixed. This is the rule that tells us how much an insulated gas cools as it expands.
---
$ T_H V_b^{\gamma - 1} = T_C V_c^{\gamma - 1}, \qquad T_H V_a^{\gamma - 1} = T_C V_d^{\gamma - 1}
Apply that rule to the two adiabatic stages. The second stage runs from state b at T_H to state c at T_C, and the fourth runs from state d at T_C back to state a at T_H. Each gives one equation linking a pair of volumes through the two temperatures.
---
$ \frac{V_b}{V_a} = \frac{V_c}{V_d}
Divide the first equation by the second. The temperatures T_H and T_C appear on both sides and cancel, leaving (V_b/V_a)^(γ−1) = (V_c/V_d)^(γ−1), and taking the (γ−1)th root of both sides gives this result. The expansion ratio in the hot stage equals the compression ratio in the cold stage.
---
$ \frac{Q_C}{Q_H} = \frac{T_C}{T_H}
Divide the expression for Q_C by the one for Q_H. Because the volume ratios are equal, the two logarithms are equal and cancel, and so do n and R. The ratio of the heats depends only on the ratio of the temperatures, and every detail of the gas and the cylinder has dropped out.
---
$ \eta_{\text{Carnot}} = 1 - \frac{T_C}{T_H}
Put this ratio into η = 1 − Q_C/Q_H. The temperatures must be in kelvin, because the derivation used the ideal gas law, which measures temperature from absolute zero.
:::

The result says that the efficiency of an ideal engine depends only on the temperatures of its two reservoirs. It reaches 100% only if T_C is zero kelvin, which can never be achieved, so even a perfect engine must throw some heat away. The hotter the hot side and the colder the cold side, the better it can do.

## Why can no engine do better?

The ideal gas was only a convenient working substance. Carnot's deeper claim was that no engine of any kind, working between the same two temperatures, can be more efficient than a reversible one. The argument is short and clever.

Suppose someone claimed to have built an engine more efficient than Carnot's. Connect it to a Carnot engine running backwards. A Carnot engine run in reverse is a refrigerator: it takes in work, draws heat from the cold reservoir, and delivers heat to the hot one. Let the claimed engine supply exactly the work the refrigerator needs. Because the claimed engine is more efficient, it needs less heat from the hot reservoir to produce that work than the refrigerator pumps back into it. Now look at the pair together. No outside work goes in, the hot reservoir gains heat overall, and by the first law that heat must have come from the cold reservoir. The combined machine moves heat from cold to hot with no other effect.

That never happens. Heat flows by itself from hot to cold, never the other way, and every attempt to build a machine that does otherwise has failed. So the claimed engine cannot exist, and every reversible engine between the same temperatures must have exactly the Carnot efficiency, whatever it is made of. The principle that heat does not flow unaided from cold to hot is the second law of thermodynamics, the subject of the next lesson.

::: note
No heat engine working between two temperatures can be more efficient than 1 − T_C/T_H. The limit depends only on those temperatures, not on the engine's design or working substance, and it reaches 100% only if the cold side is at absolute zero.
:::

## How close do real engines come?

A modern coal-fired power station heats steam to about 840 K (570 °C) and condenses it at around 300 K. Its Carnot limit is 1 − 300/840 ≈ 64%, but the best such plants achieve about 45%, losing the rest to friction, heat leaks and steps that are far from reversible. Nuclear plants of the common pressurised-water design run their steam much cooler, around 560 K, because of limits on the reactor; their Carnot limit is about 46% and they achieve about 33%. A car engine reaches roughly 25–35% at its best operating point, and less in stop-and-go traffic.

The limit tells engineers where to push. Raising the hot temperature helps, but only as far as turbine blades and boiler tubes can withstand. Another approach is to stack two cycles: a combined-cycle gas plant burns gas in a turbine at a very high temperature, then uses the turbine's hot exhaust to raise steam for a second cycle. Together they exceed 60%.

The same theory works in reverse. A refrigerator or a **heat pump** uses work to move heat from a cold place to a warm one. A heat pump warming a house to 293 K by drawing heat from outside air at 273 K could, in the ideal limit, deliver T_H/(T_H − T_C) = 293/20 ≈ 15 joules of heat for every joule of electricity it uses. That is not a violation of energy conservation: most of the heat is moved from outside, and the electricity only pays for the moving. Real heat pumps deliver 3 or 4 joules per joule, still far better than an electric heater, which delivers exactly 1.

::: history 1824 · Paris
Sadi Carnot was a young French army engineer who believed that steam power lay behind Britain's industrial and military strength, and that France needed to understand it. In 1824 he published a short book, *Reflections on the Motive Power of Fire*. He compared heat falling from a hot body to a cold one with water falling through a mill wheel: the wheel's output depends on how far the water falls, and he argued that an engine's output depends on how far the heat "falls" in temperature.

Carnot, like most of his contemporaries, thought heat was a conserved fluid, caloric. That was wrong, since an engine turns some heat into work, so less heat reaches the cold side than left the hot one; water leaving a mill wheel is not used up, but heat leaving an engine partly is. Yet his central conclusion, that the best possible efficiency depends only on the two temperatures, survived the correction. The book was barely read until Émile Clapeyron restated its argument with pressure–volume diagrams in 1834. By then Carnot was dead, killed by cholera in 1832 at the age of 36. Around 1850 William Thomson (later Lord Kelvin) and Rudolf Clausius reconciled Carnot's reasoning with the new law of energy conservation, and in doing so arrived at the second law.
:::

::: quiz
? An engine takes in heat at 600 K and rejects it at 300 K. What is its greatest possible efficiency?
+ 50%
- 100%
- 200%
= The Carnot limit is η = 1 − T_C/T_H = 1 − 300/600 = 0.5. Any real engine running between these temperatures does worse.
? Which change would raise the Carnot efficiency of an engine?
- Using helium instead of air as the working gas
+ Raising the temperature of the hot reservoir
- Making the cylinder larger
= Carnot's limit depends only on the two reservoir temperatures, T_H and T_C. The working substance and the size of the engine do not appear in it at all.
? Why must every heat engine reject some heat to a cold reservoir?
- Because of friction, which better engineering could eliminate
+ Because even a perfect reversible engine has Q_C/Q_H = T_C/T_H, which is zero only at absolute zero
- Because the first law forbids converting heat to work
= The first law alone would allow complete conversion of heat into work. The limit comes from Carnot's argument, which rests on the second law: heat does not flow by itself from cold to hot.
:::
