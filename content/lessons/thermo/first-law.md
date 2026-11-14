Mechanics already has a law of energy conservation. A swinging pendulum trades kinetic energy for potential energy and back, and the total stays fixed. But a real pendulum slows down and stops, a sliding box comes to rest, and a car braking to a halt loses all its kinetic energy. For a long time these looked like genuine losses, exceptions that made energy conservation a useful rule for idealised machines rather than a law of nature. The first law of thermodynamics removes the exceptions. The "lost" energy has not vanished: it has become the random motion of molecules, the internal energy of warm brakes and warm air. Once heat is counted as one of the ways energy moves, energy is conserved in every process ever observed. This lesson states that law, works out what it says about gases, and uses it to explain a surprising measurement.

## What exactly is being conserved?

To apply a conservation law you must first say what you are keeping track of. Physicists call it the **system**: whatever you choose to draw a boundary around, such as the gas in a cylinder, a whole engine, or a cup of coffee. Everything outside the boundary is the surroundings.

The system has an **internal energy**, written U: the total energy of all its molecules, including the kinetic energy of their motion and the potential energy of the forces between them. Energy can cross the boundary in only two ways. It can flow as **heat**, written Q, which happens whenever there is a temperature difference between the system and its surroundings. Or it can be transferred as **work**, written W, which happens when a force pushes something through a distance, as when a gas pushes a piston outward or a paddle stirs a liquid. If those are the only ways in or out, then any change in the internal energy must be accounted for by them.

::: equation The first law of thermodynamics
\Delta U = Q - W
:::

The signs need care. Here Q is the heat added *to* the system, and W is the work done *by* the system on its surroundings. Heat flowing in raises the internal energy; work done on the outside world uses some of it up, so it is subtracted. The convention comes from the study of engines, where the work an engine delivers is what you want. Chemists and many modern textbooks write ΔU = Q + W, with W meaning work done *on* the system. Both say the same thing; only the bookkeeping differs, so it is worth checking which convention a source uses.

## Why can't a body "contain" heat?

There is an important difference between U on one side of the law and Q and W on the other. The internal energy depends only on the system's present condition, its **state**, which for a gas is fixed by its pressure, volume and temperature. If a gas returns to the same pressure, volume and temperature, it has the same internal energy, whatever happened to it in between. Physicists call such a quantity a **state function**.

Heat and work are not like that. They describe a process, not a state. Suppose you want to raise the internal energy of a bottle of water by 1000 J. You could heat it on a stove, or you could shake it vigorously and do 1000 J of work on it, or you could do some of each. The end state is identical in every case, but the split between heat and work is entirely different. So it makes no sense to ask how much heat the water "has": heat is a name for energy while it is crossing the boundary, and once inside, it is simply internal energy.

## How much work does a gas do when it expands?

The most important kind of work in thermodynamics is done by a gas pushing a piston, since that is how steam engines, car engines and turbines turn heat into motion. Picture gas in a cylinder closed by a piston of area A, and let the gas push the piston outward by a small distance dx.

::: steps Work done by an expanding gas
$ \mathrm{d}W = F\,\mathrm{d}x
Work is defined as force times the distance moved in the direction of the force. We take a small distance dx so that the force can be treated as constant while the piston moves, even if it changes over a longer stroke.
---
$ F = PA
Pressure is force per unit area, so the total force the gas exerts on the piston is its pressure multiplied by the piston's area. This force acts straight outward, along the direction the piston moves, so all of it does work.
---
$ \mathrm{d}W = PA\,\mathrm{d}x = P\,\mathrm{d}V
Put the force into the work formula. The product A dx is the volume of the thin slice the piston sweeps out as it moves, which is exactly the increase dV in the gas's volume. The piston's area and position have disappeared, and the work is written entirely in terms of properties of the gas.
---
$ W = \int_{V_1}^{V_2} P\,\mathrm{d}V
Over a larger expansion, from volume V₁ to V₂, the pressure usually changes, so add up the contributions P dV from every small slice. That sum is the integral, which is the area under the curve of pressure plotted against volume. When the pressure stays constant it reduces to W = PΔV.
:::

The shape of the result is general: it does not depend on the cylinder being round or the piston being flat. A gas that expands does positive work on its surroundings. A gas that is compressed has work done on it, and in this convention W is negative, so the internal energy rises.

## What is the internal energy of an ideal gas?

To use the first law, we need to know U. For a gas made of single atoms, such as helium or argon, kinetic theory gives the answer directly. Each atom has an average kinetic energy of (3/2)kT, and in an ideal gas the atoms exert no forces on each other, so there is no potential energy to add. The total for N atoms is U = (3/2)NkT, which can also be written (3/2)nRT for n moles.

The striking feature is that U depends only on temperature. Squeeze an ideal gas into half the volume at the same temperature and its internal energy is unchanged, because the molecules move just as fast as before and have no forces between them to store energy. Molecules with two or more atoms, such as nitrogen and oxygen, can also spin and vibrate, which stores extra energy and makes U larger, but for an ideal gas it still depends on T alone.

## Why does it matter how you heat a gas?

Now a question that the first law answers neatly. How much heat does it take to warm one mole of helium by one kelvin? It sounds as though there should be a single answer, but experiment gives two, depending on whether the gas is allowed to expand.

::: steps Heat capacities of a monatomic ideal gas
$ Q_V = \Delta U = \tfrac{3}{2} n R\,\Delta T
First, hold the volume fixed by locking the piston. Nothing moves, so the gas does no work, W = 0, and the first law says all the heat goes into internal energy. Since U = (3/2)nRT, a temperature rise ΔT raises U by (3/2)nRΔT, and that is the heat needed.
---
$ C_V = \frac{Q_V}{n\,\Delta T} = \tfrac{3}{2} R
The **molar heat capacity** is the heat needed per mole per kelvin, so divide by nΔT. At constant volume it is (3/2)R, about 12.5 J/(mol·K), a value fixed entirely by kinetic theory with nothing about helium in particular.
---
$ Q_P = \Delta U + P\,\Delta V
Now let the piston move freely so the pressure stays constant. As the gas warms it expands and pushes the piston out, doing work PΔV on the surroundings. The first law, rearranged as Q = ΔU + W, says the heat must now pay for two things: the rise in internal energy and that work.
---
$ P\,\Delta V = nR\,\Delta T
To find the work, apply the ideal gas law PV = nRT before and after the warming. With P fixed, the only things that change are V and T, so a change ΔT in temperature produces a change PΔV = nRΔT. The work done depends only on the temperature rise, not on the size of the container.
---
$ C_P = \frac{Q_P}{n\,\Delta T} = \tfrac{3}{2}R + R = \tfrac{5}{2} R
Add the two parts and divide by nΔT. The rise in internal energy is exactly the same as at constant volume, because for an ideal gas U depends only on temperature. The extra R, about 8.3 J/(mol·K), is the cost of pushing back the surroundings as the gas expands.
:::

The ratio of the two heat capacities, C_P/C_V, is (5/2)/(3/2) = 5/3 ≈ 1.67. Measurements on helium, neon and argon give 1.67 to within the accuracy of the experiment. That is a strong check, because the prediction combines the first law with the kinetic-theory picture of atoms as tiny moving points, and neither idea was built to fit this number. For air, whose molecules can also rotate, the ratio is about 1.40, which kinetic theory also explains once rotation is included.

## What happens in the simplest processes?

Three special cases come up repeatedly, and the first law makes each of them easy to reason about.

If the volume is held fixed, no work is done, and every joule of heat added goes into internal energy: ΔU = Q. This is the case shown in the model below.

If the temperature is held fixed, which is called an **isothermal** process, then for an ideal gas the internal energy does not change, because it depends only on temperature. The first law then says Q = W. A gas expanding slowly while kept in contact with a warm reservoir turns every joule of heat it absorbs into work. This is the heart of the ideal engine described in the next lesson.

If no heat crosses the boundary, either because the system is insulated or because the change is too quick for heat to flow, the process is called **adiabatic**. Then Q = 0 and ΔU = −W. A gas that expands adiabatically does work using its own internal energy, so it cools. That is why the air rushing out of a pressurised can feels cold, and why rising air in the atmosphere cools as it expands, forming clouds. A gas that is compressed adiabatically has work done on it, so it warms. Pump up a bicycle tyre quickly and the bottom of the pump gets hot, mostly for this reason, with friction adding a little.

::: model gas
Heat the gas at fixed volume. No work is done, so all of the heat raises the molecules' kinetic energy, and the pressure climbs with the temperature.
:::

## Who discovered it?

::: history 1842–1850 · Heilbronn, Manchester, Berlin
Energy conservation was reached in the 1840s by several people working independently and by very different routes. Julius Robert Mayer, a German doctor on a ship to Java, noticed that the blood he drew from sailors' veins was unusually bright red in the tropics. He reasoned that in a hot climate the body needs to burn less of its fuel to keep warm, so less oxygen is taken from the blood. From this he arrived at the idea that heat and mechanical work are interchangeable forms of one thing. In 1842 he estimated the exchange rate using the difference between C_P and C_V of air, the same difference derived in this lesson. His paper was largely ignored, and he struggled for years for recognition.

James Joule, the son of a wealthy brewer near Manchester, reached the same conclusion by measurement. Through the 1840s he showed, with paddle wheels churning water and with electric currents heating wires, that a given amount of work always produced the same amount of heat. Senior scientists were at first sceptical of the precision he claimed. In 1847 Hermann von Helmholtz, then a young army surgeon in Berlin, set out conservation of energy as a general principle covering mechanics, heat, electricity and living things. In 1850 Rudolf Clausius, also in Berlin, wrote the law in the form used here, with internal energy as a property of the system and heat and work as two ways of changing it.
:::

::: note
Energy is conserved once heat is counted: ΔU = Q − W. Internal energy is a property of the system's state, while heat and work are two different ways of changing it, so a body has internal energy but does not contain heat or work.
:::

::: quiz
? A gas absorbs 500 J of heat and does 200 J of work on a piston. What happens to its internal energy?
+ It rises by 300 J
- It rises by 700 J
- It falls by 300 J
= By the first law, ΔU = Q − W = 500 J − 200 J = 300 J. Of the energy that came in as heat, 200 J left again as work and the rest stayed.
? Why does it take more heat to warm a gas by 1 K at constant pressure than at constant volume?
- The molecules move faster at constant pressure
+ At constant pressure the gas expands and does work, which the heat must also supply
- Pressure stores energy that volume does not
= The rise in internal energy is the same in both cases, because it depends only on temperature. At constant pressure the heat must also pay for the work PΔV done in pushing the surroundings back.
? Air in a bicycle pump is compressed quickly, too fast for heat to escape. What happens to its temperature?
- It falls, because Q = 0
- It stays the same, because no heat is added
+ It rises, because work done on the gas increases its internal energy
= The process is adiabatic, so Q = 0, and work is done on the gas, so W is negative. The first law gives ΔU = −W, which is positive, so the gas warms.
:::
