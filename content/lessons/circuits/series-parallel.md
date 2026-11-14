Real circuits contain many components. There are only two basic ways to connect two of them: one after the other, so the same current passes through both (**series**), or side by side, so they share the same voltage (**parallel**). Any network of resistors can be analysed by combining these two rules, plus two conservation laws.

## Kirchhoff's two laws

::: history 1845 · Königsberg
Gustav Kirchhoff, then a 21-year-old student, stated two rules that let anyone work out the currents in a network of wires and batteries, however tangled. They are not new physics, just conservation of charge and of energy written for circuits, but they turned circuit analysis into routine algebra. Kirchhoff later co-founded spectroscopy with Robert Bunsen.
:::

- **Current law:** at any junction, the total current in equals the total current out. Charge cannot pile up.
- **Voltage law:** around any closed loop, the voltages add to zero. A charge that goes round a loop and returns to its start must end with the energy it began with, so the voltage gained in the battery equals the voltage lost in the components.

The second law has a familiar counterpart: walk any route through hilly country and back to your front door, and your total climb equals your total descent.

## Series

In series, components form a single path, so the same current I flows through each. The battery's voltage is shared among them.

::: steps Resistors in series
$ V = V_1 + V_2
By the voltage law, the battery's voltage V is used up across the two resistors in turn.
---
$ V = IR_1 + IR_2 = I(R_1 + R_2)
Apply Ohm's law to each resistor. The current I is the same in both, since there is nowhere else for it to go.
---
$ R_{\text{series}} = R_1 + R_2
The pair behaves like a single resistor of this value. Series resistances add.
:::

Adding a resistor in series always increases the total resistance and reduces the current. The voltage divides in proportion to resistance: the bigger resistor gets the bigger share. This **voltage divider** is how a volume knob or a sensor circuit produces a chosen fraction of a supply voltage.

The drawback of series wiring is that the current has only one path. Old strings of Christmas lights were wired this way, and when one bulb burned out, the circuit was broken and every bulb went dark.

## Parallel

In parallel, each component connects directly across the same two points, so each has the same voltage V. The current splits among the branches.

::: steps Resistors in parallel
$ I = I_1 + I_2
By the current law, the current from the battery divides between the two branches and recombines afterwards.
---
$ I = \frac{V}{R_1} + \frac{V}{R_2} = V\left(\frac{1}{R_1} + \frac{1}{R_2}\right)
Apply Ohm's law to each branch. Both have the full voltage V across them.
---
$ \frac{1}{R_{\text{parallel}}} = \frac{1}{R_1} + \frac{1}{R_2}
Compare with I = V/R for a single equivalent resistor. In parallel, the reciprocals add.
---
$ R_{\text{parallel}} = \frac{R_1 R_2}{R_1 + R_2}
For two resistors only, combine the fractions. Two equal resistors R in parallel give R/2.
:::

Adding a branch in parallel always *decreases* the total resistance, because it opens another path for current. The combined resistance is less than the smallest branch. More current flows through the branch with lower resistance.

Household wiring is parallel. Every appliance gets the full mains voltage and works independently; switching off the lamp does not switch off the fridge. But each appliance switched on adds a path and draws more total current from the supply, which is why plugging too many heaters into one circuit trips the breaker.

::: note
Series components share the same current and their resistances add. Parallel components share the same voltage and their reciprocal resistances add, so the total is less than any single branch.
:::

## A worked example

A 12 V battery drives a 6 Ω resistor in series with a parallel pair of 4 Ω and 12 Ω.

The parallel pair: 4 × 12 / (4 + 12) = 3 Ω. Total: 6 + 3 = 9 Ω. Current from the battery: 12/9 ≈ 1.33 A. Voltage across the 6 Ω resistor: 1.33 × 6 = 8 V, leaving 4 V across the pair. Through the 4 Ω branch flows 4/4 = 1 A; through the 12 Ω branch, 4/12 ≈ 0.33 A. Check: 1 + 0.33 = 1.33 A, as the current law requires.

Capacitors combine the other way round: in parallel their capacitances add, and in series their reciprocals add.

::: quiz
? Two 10 Ω resistors are connected in parallel. Their combined resistance is
- 20 Ω
+ 5 Ω
- 10 Ω
= Equal resistors in parallel give half of one: 10 × 10 / 20 = 5 Ω.
? A third resistor is added in parallel to two others across a battery. The current from the battery
+ increases
- decreases
- stays the same
= Another path lowers the total resistance, so more current flows at the same voltage. The currents in the original branches are unchanged.
? Two resistors of 2 Ω and 8 Ω are in series across 10 V. What is the voltage across the 8 Ω resistor?
- 2 V
- 5 V
+ 8 V
= The current is 10/(2 + 8) = 1 A, so the 8 Ω resistor has 8 V across it. Voltage divides in proportion to resistance.
:::
