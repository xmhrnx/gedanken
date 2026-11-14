A battery delivers energy steadily over hours. Sometimes what is needed instead is a large amount of energy released in a fraction of a second: the burst of light from a camera flash, or the jolt a defibrillator sends through a stopped heart. And sometimes what is needed is a tiny store of charge that can be filled and emptied billions of times a second, as in the main memory of a computer. All of these rely on the same simple device, the **capacitor**: two conductors separated by an insulator. Connect it to a battery and charge builds up, positive on one conductor and negative on the other, until the voltage between them matches the battery's. Disconnect it and the charge stays where it is, ready to be released. This lesson works out how much charge a capacitor holds, how much energy it stores, and where that energy actually lives.

## How much charge can a capacitor hold?

The simplest capacitor is two flat metal plates, each of area A, facing each other across a small gap d. Put a charge +Q on one plate and −Q on the other. The opposite charges attract, so they gather on the inner faces of the plates, facing each other across the gap, and between the plates they produce a nearly uniform field.

Here is the key observation. The field between the plates is set up by the charge on them, so if you double the charge, you double the field everywhere in the gap. The voltage between the plates is the field multiplied by the gap, V = Ed, as the previous lesson showed, so it doubles too. Charge and voltage are proportional. The constant of proportionality, which tells you how much charge the capacitor holds for each volt, is called its **capacitance**, C.

::: equation Capacitance
Q = CV
:::

Capacitance is measured in **farads** (F), named after Michael Faraday: a capacitor of 1 F holds 1 coulomb of charge on each plate for every volt across it. The capacitance does not depend on how much charge is stored or what voltage is applied. It depends only on the shape and size of the conductors and on what fills the space between them. For two parallel plates with empty space between them, it can be worked out directly.

::: steps Capacitance of two parallel plates
$ E = \frac{Q}{\varepsilon_0 A}
The charge Q is spread evenly over the inner face of a plate of area A. The field between the plates is proportional to how densely the charge is packed, Q/A, since by Gauss's law the field lines leaving the charge all cross the gap and their density measures the field. The constant in front, 1/ε₀, comes from Gauss's law applied to a flat sheet of charge; it is the same ε₀ that appears in Coulomb's law. This holds when the gap is small compared with the plates, so that the edges, where the field bulges outward, can be neglected.
---
$ V = Ed = \frac{Qd}{\varepsilon_0 A}
The field is uniform across the gap, so the voltage between the plates is the field times the distance, as the previous lesson showed. Putting in the field from the first step gives the voltage in terms of the charge and the dimensions of the plates. As expected, V is proportional to Q.
---
$ C = \frac{Q}{V} = \frac{\varepsilon_0 A}{d}
Capacitance is charge per volt, so divide Q by the voltage. The charge cancels, confirming that capacitance depends only on the geometry: the area of the plates and the distance between them.
:::

::: equation Parallel-plate capacitor
C = \frac{\varepsilon_0 A}{d}
:::

Both parts of the result make sense once you think about the field. Bigger plates hold more charge at the same voltage, because the same charge spread over a larger area is less densely packed and produces a weaker field, and so a smaller voltage. Closer plates help too: the field from a given charge is the same, but over a shorter gap it adds up to a smaller voltage, so more charge can be loaded on before the voltage reaches the battery's.

## Why is the farad so large?

A quick calculation shows that one farad is an enormous capacitance. Take plates in a vacuum separated by 1 mm. To reach 1 F, rearranging C = ε₀A/d gives an area A = Cd/ε₀ = (1 × 0.001)/(8.85 × 10⁻¹²) ≈ 1.1 × 10⁸ m². That is a square more than 10 km on a side, with the two sheets held a millimetre apart.

Practical capacitors therefore range from picofarads (10⁻¹² F) to a few millifarads. Engineers get large areas into small packages by rolling up long strips of thin metal foil separated by an insulating film, or stacking many thin layers. **Supercapacitors** go much further, reaching thousands of farads in a device the size of a drink can. They use electrodes of porous carbon, riddled with tiny holes that give each gram an internal surface area of a thousand square metres or more, and the "gap" between the charges is a layer only a few molecules thick.

## What does an insulator between the plates do?

Filling the gap with an insulator, called a **dielectric**, raises the capacitance by a factor κ (the Greek letter kappa), known as the **dielectric constant** of the material. It is between about 2 and 4 for common plastics, between about 5 and 10 for glass, and about 80 for water.

The reason is polarization, the same effect that lets a charged balloon stick to a wall. The field between the plates pulls the positive part of each molecule in the dielectric slightly one way and the negative part slightly the other. No charge leaves the molecules, but the effect is that a thin layer of negative charge appears on the dielectric's surface facing the positive plate, and positive charge on the surface facing the negative plate. These induced charges produce a field of their own, pointing opposite to the field from the plates, and partly cancel it. With a weaker field inside, the voltage for a given charge is smaller, and since C = Q/V, the capacitance is larger by the factor κ.

Dielectrics have a second advantage. Most solid insulators can withstand much stronger fields than air before they break down and spark, so a capacitor with a good dielectric can be charged to higher voltages without failing.

## How much energy does a capacitor store?

Charging a capacitor takes work. Every bit of charge added to the positive plate is repelled by the positive charge already there and attracted back by the negative charge on the other plate, so it has to be pushed. The work done is stored as energy, and it comes back when the capacitor discharges.

You might guess the energy is simply charge times voltage, QV. It is not, because the voltage is not constant during charging. It starts at zero, when the plates are empty, and rises steadily as charge accumulates. The first bit of charge is moved across almost no voltage and costs almost nothing; the last bit is moved across the full voltage. The work has to be added up piece by piece.

::: steps Energy in a capacitor
$ V = \frac{q}{C}
Imagine the capacitor partway through charging, when the charge on the plates has reached some value q, smaller than the final Q. Since Q = CV holds at every stage, the voltage at that moment is q/C. Using a lowercase q keeps this changing amount separate from the final charge.
---
$ \mathrm{d}W = V\,\mathrm{d}q = \frac{q}{C}\,\mathrm{d}q
Now move a further small charge dq from one plate to the other. Voltage is energy per unit charge, so moving charge dq across voltage V takes work V dq. The amount dq is small enough that the voltage hardly changes while it is moved.
---
$ W = \int_0^Q \frac{q}{C}\,\mathrm{d}q = \frac{Q^2}{2C}
Add up the work for every small step, from an empty capacitor to the final charge Q. The integral of q is q²/2, so the total is Q²/2C. Geometrically, this is the area under the straight line of voltage plotted against charge: a triangle with base Q and height Q/C, whose area is half of base times height.
---
$ U = \tfrac{1}{2}CV^2 = \tfrac{1}{2}QV
Use Q = CV to write the stored energy in terms of whichever quantities are convenient. The factor of a half is the answer to the puzzle about QV: the charge was moved, on average, across half the final voltage, because the voltage rose steadily from zero.
:::

A defibrillator shows the numbers involved. A typical one charges a capacitor of about 100 microfarads (10⁻⁴ F) to around 2000 V, storing U = ½ × 10⁻⁴ × (2000)² = 200 J. That is about the energy needed to lift a 20 kg suitcase one metre, not much in itself. What matters is that the capacitor delivers it through the patient's chest in a few milliseconds, a power of tens of kilowatts. A battery of the same size stores far more energy, but its chemical reactions can release it only slowly. A capacitor's strength is speed.

## Where is the energy stored?

The energy in a capacitor has to be somewhere. One answer is that it belongs to the charges, which have been pushed into an arrangement they resist. But there is another way to describe it that turns out to be more fundamental.

::: steps Energy density of an electric field
$ U = \tfrac{1}{2}CV^2 = \tfrac{1}{2}\left(\frac{\varepsilon_0 A}{d}\right)(Ed)^2
Start from the stored energy and express everything in terms of the field between the plates. The capacitance is ε₀A/d, and the voltage is the field times the gap, V = Ed. This is allowed because both relations hold for any parallel-plate capacitor with a uniform field.
---
$ U = \tfrac{1}{2}\varepsilon_0 E^2 \,(Ad)
Multiply out. One factor of d cancels with the d in the denominator, and what remains is ½ε₀E² multiplied by Ad. But Ad is just the volume of the space between the plates, where the field is.
---
$ u = \frac{U}{Ad} = \tfrac{1}{2}\varepsilon_0 E^2
Divide by that volume to get the energy per cubic metre, called the **energy density** u. It depends only on the strength of the field at that place. The plates, their area and their spacing have all disappeared.
:::

That result suggests a striking interpretation: the energy is stored in the electric field itself, at an amount of ½ε₀E² in every cubic metre of space where the field exists. The formula, derived here for a capacitor, turns out to hold for any electric field whatever. It is also one of the reasons physicists treat fields as real. Light carries energy across empty space from the Sun to Earth, with no charges along the way; on this view, the energy travels in the electric and magnetic fields of the light wave.

::: note
A capacitor stores separated charge Q = CV, with capacitance set only by its shape and what fills the gap. Charging it stores energy ½CV², and that energy lives in the electric field between the plates, at ½ε₀E² per cubic metre.
:::

::: history 1745–1746 · Pomerania and Leiden
The first capacitor was found by accident, twice, in the space of a few months. In 1745 Ewald von Kleist, a German cleric in Pomerania, tried to store electricity in a glass bottle of water, with a nail pushed through the cork into the water and connected to a friction machine that generated charge. When he touched the nail while holding the bottle, he received a startling shock. Early in 1746, at the University of Leiden in the Netherlands, Pieter van Musschenbroek and his associates tried a similar experiment, holding a glass jar of water in one hand while it was charged through a wire. When the experimenter touched the wire with his other hand, the shock was so violent that Musschenbroek wrote he would not repeat it for the whole kingdom of France.

The device became known as the **Leyden jar**, and it was soon improved by coating the inside and outside of the glass with metal foil, making two conductors separated by glass: exactly the plan of a capacitor. Benjamin Franklin investigated it in the late 1740s. By taking a charged jar apart, he showed that the charge was held by the glass between the conductors, not by the water, and he connected several jars together to store more charge, calling the arrangement a "battery", the first use of the word in electricity.
:::

## What happens when a capacitor discharges?

Connect the two plates of a charged capacitor through a resistor, a component that limits how fast current can flow, and the charge flows back from one plate to the other. At first the voltage is at its highest, so the current is largest. As charge drains away, the voltage falls, so the current falls, so the charge drains more slowly. The result is a smooth decay that never quite reaches zero: the charge falls by the same fraction in each equal interval of time, a pattern called exponential decay.

The pace of that decay is set by the product of the resistance R and the capacitance C, called the **time constant** RC. After one time constant, the charge has fallen to about 37% of its starting value; after five, it is less than 1%. A 1000 μF capacitor discharging through a 1 kΩ resistor has a time constant of 10⁻³ F × 10³ Ω = 1 second. By choosing R and C, circuit designers can set timings from microseconds to minutes, in everything from blinking lights to the circuits that tune a radio to a station.

::: quiz
? A capacitor is charged to 10 V. You then charge it to 20 V instead. What happens to its stored energy?
- It doubles
+ It quadruples
- It stays the same
= The stored energy is U = ½CV², which depends on the square of the voltage. Doubling the voltage therefore multiplies the energy by 2² = 4.
? You pull apart the plates of a charged capacitor that is not connected to anything, so its charge cannot change. What happens to the voltage?
+ It rises
- It falls
- It stays the same
= Moving the plates apart increases d, so the capacitance C = ε₀A/d falls, and with Q fixed the voltage V = Q/C rises. The extra stored energy comes from the work you did pulling the attracting plates apart.
? Why does inserting a dielectric between the plates increase the capacitance?
- It adds extra charge to the plates
+ Its polarized molecules partly cancel the field, lowering the voltage for the same charge
- It makes the plates larger
= The field between the plates polarizes the dielectric, and the induced charges partly cancel that field. A weaker field means a smaller voltage for the same charge, and since C = Q/V, the capacitance goes up.
:::
