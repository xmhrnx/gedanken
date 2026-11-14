A real circuit rarely contains just one battery and one resistor. A torch has a battery, a switch and a bulb; a house has dozens of lamps and appliances hanging off the same supply; a phone has billions of components. To predict what happens in such a network you need to know how components behave when they are connected together. The surprising news is that there are only two basic ways to connect two components, and two simple rules, both of them conservation laws, are enough to work out the currents and voltages in any network of resistors, however tangled.

## What rules does every circuit obey?

::: history 1845 · Königsberg, Prussia
Gustav Kirchhoff was a 21-year-old student at the university in Königsberg when he wrote down two rules for networks of wires and batteries. Others had analysed simple circuits before, but each new arrangement had been treated as a fresh puzzle. Kirchhoff saw that two general statements, applied at every junction and around every loop, always give exactly as many equations as there are unknown currents, so any network can be solved by routine algebra. He went on to a distinguished career, most famously founding the science of spectroscopy with the chemist Robert Bunsen around 1860, when they showed that each chemical element emits light of its own characteristic colours.
:::

Kirchhoff's two rules are not new laws of nature. Each is a familiar conservation law rewritten in the language of circuits, which is exactly why they can be trusted everywhere.

The first, the **current law**, says that at any junction the total current flowing in equals the total current flowing out. The previous lessons gave the reason: charge is conserved, and in a steady circuit it cannot pile up anywhere, so whatever arrives at a junction each second must leave it.

The second, the **voltage law**, says that if you go round any closed loop in a circuit and add up the voltage changes, with gains counted as positive and drops as negative, the total is zero. The reason is conservation of energy. Voltage is energy per unit charge, so the voltage changes describe what happens to the energy of a charge as it moves. A charge that travels round a loop and returns to where it started must have exactly the energy it began with; otherwise you could send charges round and round and create energy from nothing. So every volt gained in the battery is given up somewhere in the components. Walking is a fair comparison here: take any route through hilly country that ends at your own front door, and your total climb equals your total descent, whatever path you took.

## What happens when components are in series?

The first way to connect two components is one after the other, so that there is only a single path and all the charge must pass through each in turn. This is called **series** connection. Because there is only one path, and charge cannot pile up between them, the same current flows through each component. What differs between them is the voltage: the battery's push is shared out along the chain.

::: steps Resistors in series
$ V = V_1 + V_2
Go round the loop from the battery through both resistors and back. By the voltage law, the voltage V gained across the battery must equal the total dropped across the two resistors, V₁ across the first and V₂ across the second. The battery's voltage is being divided between them.
---
$ V = IR_1 + IR_2
Now apply Ohm's law to each resistor separately: the voltage across a resistor is the current through it times its resistance. The current I is the same in both, because in a single path there is nowhere else for any of it to go.
---
$ V = I(R_1 + R_2)
Since I appears in both terms, take it outside a bracket. The equation now has the same form as Ohm's law for a single resistor, V = IR, with R₁ + R₂ in the place of R.
---
$ R_{\text{series}} = R_1 + R_2
So the pair behaves, as far as the battery can tell, exactly like one resistor whose resistance is the sum of the two. Series resistances add. The same argument works for any number of resistors in a chain.
:::

This result makes sense from the previous lesson, where two identical wires joined end to end acted like one wire of double length and double resistance. It has two consequences. Adding a resistor in series always increases the total resistance, so it always reduces the current from the battery. And the battery's voltage is not split equally but in proportion to resistance: since the same current flows through each resistor, the one with the larger resistance needs the larger voltage to push that current through. An arrangement that uses this to produce a chosen fraction of a supply voltage is called a **voltage divider**, and it is how a volume knob or a light sensor in a circuit sets a voltage between zero and the full supply.

The weakness of series wiring is that single path. If any component fails by breaking the circuit, the current stops everywhere. Old strings of Christmas lights were wired in series, and when one bulb burned out, every bulb on the string went dark, leaving you to test them one at a time to find the culprit.

## What happens when components are in parallel?

The second way to connect two components is side by side, with each one joined directly across the same two points. This is called **parallel** connection. Now the situation is reversed. Each component is connected straight across the battery, so each has the same voltage across it. What differs is the current: the flow from the battery splits between the branches and joins up again afterwards.

::: steps Resistors in parallel
$ I = I_1 + I_2
At the junction where the wire from the battery divides, the current law says the current I arriving must equal the sum of the currents I₁ and I₂ leaving along the two branches. The same thing happens in reverse where the branches rejoin.
---
$ I = \frac{V}{R_1} + \frac{V}{R_2}
Apply Ohm's law to each branch, written as I = V/R. Each branch has the full battery voltage V across it, since each is connected directly between the same two points. So each branch draws its own current independently of the other.
---
$ I = V\left(\frac{1}{R_1} + \frac{1}{R_2}\right)
Take the common factor V outside a bracket. Compare this with Ohm's law for a single resistor, written as I = V × (1/R): the bracket plays the role of 1/R.
---
$ \frac{1}{R_{\text{parallel}}} = \frac{1}{R_1} + \frac{1}{R_2}
So the pair behaves like a single resistor whose reciprocal is the sum of the reciprocals. In parallel it is not the resistances that add but their reciprocals, which measure how easily each branch lets current through.
---
$ R_{\text{parallel}} = \frac{R_1 R_2}{R_1 + R_2}
For exactly two resistors, combine the two fractions over a common denominator, R₁R₂, and turn the result upside down. This gives a handy shortcut: product over sum. Two equal resistors R in parallel give R²/2R, which is R/2.
:::

The result matches the earlier picture of two identical wires laid side by side, which acted like one wire of double thickness and half the resistance. Adding a branch in parallel always *decreases* the total resistance, because it opens another route for current without taking anything away from the routes already there. The combined resistance is always less than that of the smallest single branch. And the current divides in the opposite way to voltage in series: the branch with the lower resistance carries the larger current, because it has the same voltage across it but lets charge through more easily.

Household wiring is parallel, and you can see why. Every appliance is connected across the mains, so every one gets the full mains voltage and works independently of the others; switching off the lamp does not switch off the fridge. But each appliance switched on adds another branch, lowering the total resistance and so drawing more total current from the supply through the shared wiring. That is why plugging too many heaters into one circuit trips the breaker: each heater on its own is fine, but together they draw more current than the circuit's wiring can safely carry.

::: note
Components in series share the same current, and their resistances add, so the total is larger than any one of them. Components in parallel share the same voltage, and the reciprocals of their resistances add, so the total is smaller than any single branch.
:::

## How do the rules combine in practice?

Most real circuits mix the two kinds of connection, and the method is to replace groups of resistors by their single equivalents, one step at a time, until only one resistance is left. Take a 12 V battery driving a 6 Ω resistor in series with a parallel pair of resistors, one of 4 Ω and one of 12 Ω.

Start with the parallel pair, since it can be replaced by one equivalent resistor. Product over sum gives 4 × 12 / (4 + 12) = 48/16 = 3 Ω, less than either branch as expected. That 3 Ω is in series with the 6 Ω resistor, so the whole circuit behaves like 6 + 3 = 9 Ω. The battery therefore supplies a current of 12 V ÷ 9 Ω ≈ 1.33 A.

Now work back outward. All of that 1.33 A flows through the 6 Ω resistor, which therefore has 1.33 × 6 = 8 V across it. By the voltage law the remaining 4 V must appear across the parallel pair, and each branch of the pair has that same 4 V. The 4 Ω branch carries 4 ÷ 4 = 1 A, and the 12 Ω branch carries 4 ÷ 12 ≈ 0.33 A. As a check, 1 + 0.33 = 1.33 A, which is exactly the current arriving at the junction, as the current law requires. Notice that the lower-resistance branch took three times as much current as the other, because its resistance is three times smaller.

The same two rules apply to other components, though the results can come out differently. Capacitors, for instance, combine the other way round: in parallel their capacitances add, because side-by-side plates simply act like bigger plates, while in series the reciprocals add.

::: quiz
? Two 10 Ω resistors are connected in parallel. What is their combined resistance?
- 20 Ω
+ 5 Ω
- 10 Ω
= Using product over sum, 10 × 10 / (10 + 10) = 5 Ω. Two equal resistors in parallel always give half of one, because the current has two equally easy paths instead of one.
? A third resistor is added in parallel to two others connected across a battery. What happens to the current drawn from the battery?
+ It increases
- It decreases
- It stays the same
= The new branch adds a path for current, which lowers the total resistance, so more current flows at the same voltage. The currents in the original branches are unchanged, because each still has the full battery voltage across it.
? A 2 Ω and an 8 Ω resistor are in series across a 10 V supply. What is the voltage across the 8 Ω resistor?
- 2 V
- 5 V
+ 8 V
= The total resistance is 10 Ω, so the current is 10 V ÷ 10 Ω = 1 A, and the 8 Ω resistor has 1 A × 8 Ω = 8 V across it. In series the voltage divides in proportion to resistance.
:::
