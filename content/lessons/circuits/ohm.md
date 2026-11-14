Connect the same battery across a short thick copper wire and a long thin one, and very different currents flow. The material and shape of a conductor determine how much current a given voltage drives through it. For many materials the relationship is the simplest one possible: double the voltage, double the current.

## Ohm's law

The **resistance** R of a component is the voltage across it divided by the current through it.

::: equation Ohm's law
V = IR
:::

Resistance is measured in **ohms** (Ω): 1 Ω lets 1 A flow when 1 V is applied. Strictly, V = IR is just the definition of R. **Ohm's law** is the experimental claim that for many conductors, especially metals at steady temperature, R stays the same whatever the voltage. Such conductors are called **ohmic**.

Many useful components are not ohmic. A lamp filament's resistance rises about tenfold as it heats from room temperature to white heat. A diode conducts easily one way and hardly at all the other. For these, R = V/I still has a value at each operating point, but it changes.

::: history 1825–1827 · Cologne and Berlin
Georg Ohm, a schoolteacher, measured the current through wires of different lengths and thicknesses. Voltaic piles gave unsteady results, so he switched to a **thermocouple**, a junction of two metals that gives a small, very steady voltage when one end is heated, recently discovered by Thomas Seebeck. He measured current by the twist of a magnetic needle. His results, published in *Die galvanische Kette*, showed current proportional to the driving voltage and inversely proportional to the wire's length. German physicists, then suspicious of mathematical theories, received the work coldly, and Ohm struggled for years to get a university post. Recognition came from abroad: the Royal Society gave him its Copley Medal in 1841.
:::

## Resistivity

Resistance depends on shape as well as material. A wire twice as long has twice the resistance, like two wires end to end. A wire with twice the cross-sectional area has half the resistance, like two wires side by side. Pulling the material's own contribution out gives the **resistivity** ρ (rho).

::: equation Resistance of a wire
R = \rho\frac{L}{A}
:::

Copper's resistivity is about 1.7 × 10⁻⁸ Ω m, so 1 m of wire with a 1 mm² cross-section has a resistance of about 0.017 Ω. Silver is slightly better; aluminium, used in power lines because it is light and cheap, about 60% worse. Glass and rubber have resistivities around 10¹⁰ to 10¹⁴ Ω m, more than 10¹⁸ times copper's, which is why they make good insulators.

## Why wires resist

An electric field in a metal pushes on its free electrons. In empty space they would accelerate indefinitely. In a metal they keep scattering off vibrating atoms and impurities, losing the speed they gained, so they settle into a steady average drift proportional to the field. Since the field is proportional to the voltage and the drift sets the current, the current is proportional to the voltage: Ohm's law.

This picture also explains why metals resist more when hot. Hotter atoms vibrate more, scattering electrons more often. (Some materials, cooled close to absolute zero, lose all resistance; this **superconductivity** needs quantum mechanics to explain.)

## Why wires heat up

Each collision hands energy from the electron to the atoms, making them vibrate harder. That is heat. Every coulomb that crosses a resistor loses V joules of electrical energy, and all of it turns into thermal energy.

::: steps Power dissipated in a resistor
$ \Delta U = QV
Charge Q crossing a voltage V loses energy QV, by the definition of voltage.
---
$ P = \frac{\Delta U}{t} = \frac{Q}{t}V = IV
Power is energy per unit time, and Q/t is the current.
---
$ P = I^2 R
Substitute V = IR. For a fixed resistance, heating grows as the square of the current.
---
$ P = \frac{V^2}{R}
Or substitute I = V/R. At a fixed voltage, a smaller resistance draws more current and dissipates more power.
:::

This **Joule heating**, named for James Joule, who measured it in 1840, is the whole point of toasters, kettles and incandescent bulbs. Elsewhere it is waste: it is why phone chargers get warm and why power lines lose energy. It is also why a short circuit is dangerous: almost zero resistance across a fixed voltage means a huge current, and I²R heating in the wiring can start a fire. Fuses and circuit breakers cut the current before that happens.

::: note
Resistance is voltage divided by current, and for metals at fixed temperature it is constant: V = IR. The energy lost by charge crossing a resistance becomes heat at a rate I²R.
:::

::: quiz
? A wire is replaced by one of the same material, twice as long and with twice the cross-sectional area. Its resistance
- doubles
+ stays the same
- quadruples
= R = ρL/A; doubling L doubles R and doubling A halves it.
? A heater draws 10 A at 230 V. What is its resistance?
- 2300 Ω
+ 23 Ω
- 0.043 Ω
= R = V/I = 230/10 = 23 Ω.
? The current through a fixed resistor is tripled. Its heat output
- triples
- stays the same
+ rises ninefold
= P = I²R, so tripling I multiplies P by 9.
:::
