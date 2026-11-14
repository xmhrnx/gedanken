Connect the same battery first across a short, thick copper wire and then across a long, thin one, and very different currents flow. The battery's push is the same in both cases, so something about the wire itself must decide how much current that push can drive. This lesson is about that something, called resistance: how it is defined, what it depends on, why metals have it at all, and why it turns electrical energy into heat. For a great many materials, the relationship between push and flow turns out to be the simplest one possible: double the voltage and you double the current.

## What is resistance?

Recall that **voltage** is energy per unit charge: a voltage of one volt between two points means each coulomb that moves between them gains or loses one joule. A battery provides a voltage across its terminals, and that voltage drives a current through whatever is connected. The natural way to describe how hard a component is to push current through is to compare the two: how many volts does it take to drive each ampere?

::: equation Ohm's law
V = IR
:::

The quantity R is the **resistance** of the component: the voltage across it divided by the current through it. It is measured in **ohms** (Ω). A component of 1 Ω lets 1 A flow when 1 V is applied across it; a component of 100 Ω would need 100 V to drive the same current.

Here there is a subtlety worth getting clear, because it is often blurred. Written as R = V/I, the equation is only a definition: you can always divide a voltage by a current and call the answer a resistance. The real content of **Ohm's law** is an experimental claim that R does not change when you change the voltage. If that claim holds, then R is a fixed property of the component, and knowing it lets you predict the current for any voltage. For metals kept at a steady temperature it holds remarkably well, and materials that obey it are called **ohmic**.

Many useful components are not ohmic. A lamp filament's resistance rises about tenfold as it heats from room temperature to white heat, so doubling the voltage across a lamp much less than doubles the current. A diode, a device made of specially treated silicon, lets current through easily in one direction and hardly at all in the other. For components like these, V/I still has a value at each operating point, but that value changes, and you cannot treat it as a fixed number.

## How was the law found?

::: history 1825–1827 · Cologne and Berlin
Georg Ohm was a mathematics and physics teacher at a school in Cologne, with ambitions to become a professor. He set out to find how the current in a wire depends on the wire's length and thickness. His first attempts were spoiled by the voltaic piles of the day, whose push sagged as they were used, so that results drifted from minute to minute. On the suggestion of the editor Johann Poggendorff he switched to a **thermocouple**, a junction of two different metals that produces a small, very steady voltage when one end is heated and the other kept cold, an effect Thomas Seebeck had discovered a few years before. Ohm kept one junction in boiling water and the other in ice, and measured the current by how far it twisted a magnetic needle hung above the wire. The results were clean: current proportional to the driving voltage and inversely proportional to the length of the wire. He published them, with a mathematical theory of conduction, in a book called *Die galvanische Kette* in 1827. German physicists at the time distrusted theory built on mathematics rather than direct experiment, and the book was received coldly; Ohm resigned his teaching post and spent years in poorly paid work. Recognition came first from abroad, when the Royal Society in London awarded him its Copley Medal in 1841, and he finally gained a professorship in Munich in 1852, two years before his death.
:::

## What does resistance depend on?

Ohm's own measurements already hinted that resistance depends on the shape of a wire as well as its material, and the shape dependence can be worked out by a simple argument. Take a wire and join a second identical wire to its end. A charge must now pass through both, one after the other, and each takes the same voltage to drive a given current, so the total voltage needed doubles. A wire twice as long therefore has twice the resistance. Now instead lay the two wires side by side and join their ends. The current has two equal paths and divides between them, so the same voltage now drives twice the current. A wire twice as thick, with double the cross-sectional area, therefore has half the resistance.

If we take those two effects out, what is left depends only on the material. That remaining quantity is called the **resistivity**, written ρ (the Greek letter rho).

::: equation Resistance of a wire
R = \rho\frac{L}{A}
:::

Here L is the length and A the cross-sectional area. Resistivity is measured in ohm-metres, and it varies between materials more than almost any other physical property. Copper's is about 1.7 × 10⁻⁸ Ω m, so a 1-metre length of wire with a 1-square-millimetre cross-section has a resistance of about 0.017 Ω, small enough that the wiring in a house barely affects the current. Silver is slightly better than copper but expensive. Aluminium is about 60% worse, but it is light and cheap, which is why it is used for overhead power lines. Glass and rubber, at the other extreme, have resistivities of roughly 10¹⁰ to 10¹⁴ Ω m, around 10¹⁸ to 10²² times copper's. That colossal gap is why a thin layer of plastic around a copper wire can keep the current inside it.

## Why do wires resist at all?

The obvious question is why a metal, full of free electrons, resists current in the first place. An electric field inside the metal pushes on each free electron. In empty space an electron pushed by a steady field would accelerate without limit, going faster and faster. Inside a metal it gains speed only briefly before it scatters off something, and loses the extra velocity it had picked up in the direction of the field. Then the field accelerates it again, and it scatters again.

The result of many such cycles is a steady average drift. The stronger the field, the more speed each electron gains between collisions, so the average drift is proportional to the field. The field inside a wire is proportional to the voltage across it, and the drift speed sets the current, as the previous lesson showed. Put those together and the current is proportional to the voltage. That is Ohm's law, explained.

What do the electrons scatter off? The simple picture says atoms, but quantum mechanics shows that electrons pass freely through a perfectly regular arrangement of atoms. What scatters them is anything that spoils the regularity: impurity atoms, flaws in the crystal, and above all the vibrations of the atoms caused by heat. This explains why metals resist more when they are hot. Hotter atoms vibrate more strongly, the regular pattern is disturbed more, and electrons are scattered more often. Some materials, cooled close to absolute zero, lose their resistance entirely. This **superconductivity** cannot be explained by the scattering picture at all; it needs a fully quantum treatment in which electrons pair up and move together without losing energy.

## Why do wires get hot?

Each time an electron scatters, it hands some of the energy it gained from the field to the metal, making the atoms vibrate harder. Vibrating atoms are what we mean by heat. So a resistor is a device that steadily turns electrical energy into thermal energy, and we can work out at what rate.

::: steps Power dissipated in a resistor
$ \Delta U = QV
Voltage is defined as energy per unit charge, so a charge Q that crosses a voltage V gives up energy QV. In a resistor nothing else takes this energy, so all of it ends up as heat in the material.
---
$ P = \frac{\Delta U}{t} = \frac{Q}{t}V = IV
Power is the rate at which energy is transferred, so divide the energy by the time t taken. The voltage stays as it is, and the charge divided by the time is the current. So the heating rate is simply current times voltage, a result that holds for any component, not only resistors.
---
$ P = (IR)\,I = I^2 R
For a resistor we also know V = IR, so we can replace V with IR. This form is useful when the current is fixed, as in wires carrying a known current. It shows the heating grows with the square of the current: double the current and the heat output goes up four times.
---
$ P = V \cdot \frac{V}{R} = \frac{V^2}{R}
Alternatively, replace I with V/R. This form is useful when the voltage is fixed, as with appliances plugged into the mains. It shows something that surprises many people: at a fixed voltage, a smaller resistance produces more heat, because it lets more current flow.
:::

This heating is called **Joule heating**, after James Joule, who measured it carefully in 1840 and found the I²R dependence. It is the whole point of toasters, kettles, electric heaters and old-fashioned incandescent bulbs, which glow because their filament is hot. Everywhere else it is a nuisance: it is why phone chargers get warm and why power lines lose some of the energy they carry.

It is also why a short circuit is dangerous. A short circuit means the two sides of a supply are joined by almost zero resistance, perhaps by a damaged wire. Then the V²/R form says the power is enormous, because R is tiny while V stays at the full supply voltage. In practice the current becomes huge and heats the household wiring, which was never meant to carry it, until insulation melts or a fire starts. Fuses and circuit breakers exist to cut the current before that happens.

::: note
Resistance is voltage divided by current, and Ohm's law is the experimental fact that for metals at a fixed temperature this ratio stays constant, so V = IR. Resistance arises because electrons are repeatedly scattered as they drift, and each scattering hands energy to the metal, which appears as heat at a rate I²R.
:::

::: quiz
? A wire is replaced by one of the same material, twice as long and with twice the cross-sectional area. Its resistance
- doubles
+ stays the same
- quadruples
= Resistance is R = ρL/A. Doubling the length doubles R, and doubling the area halves it, so the two changes cancel exactly.
? A heater draws 10 A from a 230 V supply. What is its resistance?
- 2300 Ω
+ 23 Ω
- 0.043 Ω
= Resistance is voltage divided by current, so R = V/I = 230 V ÷ 10 A = 23 Ω.
? The current through a fixed resistor is tripled. What happens to its heat output?
- It triples
- It stays the same
+ It rises ninefold
= Heating goes as P = I²R. With R fixed, tripling the current multiplies the power by 3², which is 9.
:::
