Mechanics already has a law of energy conservation: kinetic plus potential energy stays constant, unless friction takes some away. Friction does not destroy that energy. It ends up as random motion of molecules, the internal energy of warm brakes or warm hands. Once heat is counted as a way of moving energy, conservation of energy holds without exception. That is the first law of thermodynamics.

## The law

A *system* is whatever we choose to study: a gas in a cylinder, an engine, a cup of coffee. Its *internal energy* U is the total energy of its molecules. Energy can cross the boundary in two ways: as heat Q, driven by a temperature difference, or as work W, done when a force moves something, such as a piston.

::: equation The first law of thermodynamics
\Delta U = Q - W
:::

Here Q is heat added *to* the system and W is work done *by* the system on its surroundings. Heat in raises the internal energy; work out lowers it. (Chemists and many modern texts write ΔU = Q + W with W the work done *on* the system. The physics is the same; only the sign convention differs.)

Internal energy depends only on the system's present state: its temperature, volume, pressure. Heat and work do not. The same change in U can be reached by heating, by stirring, or by a mixture, so it makes no sense to ask how much heat a body "has".

## The work done by a gas

Gas in a cylinder pushes a piston of area A outward a small distance dx.

::: steps Work done by an expanding gas
$ \mathrm{d}W = F\,\mathrm{d}x
Work is force times distance moved in the direction of the force.
---
$ F = PA
The gas pushes on the piston with pressure P over its area A.
---
$ \mathrm{d}W = PA\,\mathrm{d}x = P\,\mathrm{d}V
A times dx is the volume the piston sweeps out, the increase dV in the gas's volume.
---
$ W = \int_{V_1}^{V_2} P\,\mathrm{d}V
For a larger change the pressure may vary, so add up the small pieces: the work is the area under the curve of P against V. At constant pressure it is simply PΔV.
:::

A gas that expands does positive work; a gas that is compressed has work done on it, and W is negative.

## Internal energy of an ideal gas

For a gas of single atoms such as helium or argon, kinetic theory says each atom has average kinetic energy (3/2)kT, and with no forces between atoms there is no potential energy. So U = (3/2)NkT = (3/2)nRT. The internal energy of an ideal gas depends only on its temperature. Molecules with two or more atoms can also rotate and vibrate, which stores extra energy, but U still depends on T alone.

## Two ways to heat a gas

How much heat does it take to warm one mole of a monatomic gas by one kelvin? The answer depends on what else happens while you heat it.

::: steps Heat capacities of a monatomic ideal gas
$ Q_V = \Delta U = \tfrac{3}{2} n R\,\Delta T
At constant volume the piston does not move, so W = 0 and all the heat goes into internal energy.
---
$ C_V = \frac{Q_V}{n\,\Delta T} = \tfrac{3}{2} R
The molar heat capacity at constant volume, about 12.5 J/(mol·K).
---
$ Q_P = \Delta U + P\,\Delta V
At constant pressure the gas expands as it warms, so heat must also supply the work PΔV.
---
$ P\,\Delta V = nR\,\Delta T
From PV = nRT at constant P, a change ΔT changes V by ΔV = nRΔT/P.
---
$ C_P = \frac{Q_P}{n\,\Delta T} = \tfrac{3}{2}R + R = \tfrac{5}{2} R
ΔU is the same as before, since U depends only on T. The extra R is the cost of pushing back the surroundings.
:::

The ratio C_P/C_V = 5/3 ≈ 1.67 is just what is measured for helium, neon and argon, a striking check on kinetic theory.

## Special processes

- **Constant volume**: W = 0, so ΔU = Q.
- **Constant temperature** (isothermal) for an ideal gas: ΔU = 0, so Q = W. Every joule of heat absorbed leaves as work.
- **No heat exchange** (adiabatic): Q = 0, so ΔU = −W. A gas that expands without heat input cools; one that is compressed warms. A bicycle pump gets hot for this reason.

::: model gas
Heating at fixed volume: no work is done, so all the heat raises the molecules' kinetic energy, and the pressure climbs.
:::

::: history 1842–1850 · Heilbronn, Manchester, Berlin
Three people reached energy conservation in the 1840s by different routes. Julius Robert Mayer, a German ship's doctor, noticed that sailors' venous blood was unusually red in the tropics, and reasoned that the body burned less fuel to keep warm there; in 1842 he estimated the mechanical equivalent of heat from the difference between C_P and C_V of air. James Joule, a Manchester brewer's son, measured it directly with paddle wheels over the 1840s. Hermann von Helmholtz stated the general principle in 1847. Mayer's work was ignored for years and Joule's at first met scepticism; in 1850 Rudolf Clausius put the law in the form used here.
:::

::: note
Energy is conserved once heat is counted: ΔU = Q − W. Internal energy is a property of a state; heat and work are two ways of changing it.
:::

::: quiz
? A gas absorbs 500 J of heat and does 200 J of work on a piston. Its internal energy
+ rises by 300 J
- rises by 700 J
- falls by 300 J
= ΔU = Q − W = 500 − 200 = 300 J.
? Why does it take more heat to warm a gas by 1 K at constant pressure than at constant volume?
- The molecules move faster at constant pressure
+ At constant pressure the gas expands and does work, which the heat must also supply
- Pressure stores energy that volume does not
= The rise in internal energy is the same either way; the extra heat pays for the work PΔV.
? Air in a bicycle pump is compressed quickly, too fast for heat to escape. Its temperature
- falls, because Q = 0
- stays the same, because no heat is added
+ rises, because work done on the gas increases its internal energy
= Adiabatic: Q = 0 and W is negative, so ΔU = −W is positive and the gas warms.
:::
