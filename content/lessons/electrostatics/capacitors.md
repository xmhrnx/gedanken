A **capacitor** is two conductors separated by an insulator. Connect it to a battery and charge piles up, positive on one conductor and negative on the other, until the voltage between them matches the battery. Disconnect it and the charge stays, ready to be released in a burst. Camera flashes, defibrillators, power supplies and every memory chip depend on this.

## Capacitance

The simplest version is two flat metal plates of area A, a small distance d apart. Put charge +Q on one and −Q on the other. The field between them is uniform and proportional to Q, so the voltage V = Ed is also proportional to Q. The constant of proportionality is the **capacitance** C.

::: equation Capacitance
Q = CV
:::

Capacitance is measured in **farads** (F): 1 F stores 1 C per volt. It depends only on the shape of the conductors and what is between them. For parallel plates in vacuum:

::: equation Parallel-plate capacitor
C = \frac{\varepsilon_0 A}{d}
:::

Bigger plates hold more charge at the same voltage, because the charge spreads out and the field from it is weaker. Closer plates help too: the same field over a shorter gap means a smaller voltage, so more charge fits before the battery's voltage is reached.

The farad is huge. With ε₀ = 8.85 × 10⁻¹² F/m and a 1 mm gap, 1 F would need plates of area A = Cd/ε₀ ≈ 1.1 × 10⁸ m², a square more than 10 km on a side. Practical capacitors range from picofarads (10⁻¹² F) to millifarads, rolled up from long strips of foil. "Supercapacitors" reach thousands of farads by using porous carbon with enormous internal surface area and a gap only molecules wide.

## Dielectrics

Filling the gap with an insulator, a **dielectric**, raises the capacitance by a factor κ (kappa), the dielectric constant: about 3 for plastics, 5–10 for glass, about 80 for water. The field polarizes the insulator's molecules, and their small displaced charges partly cancel the field between the plates. Less field means less voltage for the same charge, so C = Q/V goes up. Dielectrics also withstand stronger fields than air before breaking down, letting capacitors work at higher voltages.

## Energy stored

Charging a capacitor takes work, because each new bit of charge must be pushed onto a plate that already repels it. The voltage starts at zero and rises as charge accumulates, so the work is not simply QV.

::: steps Energy in a capacitor
$ V = \frac{q}{C}
Partway through charging, with charge q on the plates, the voltage is q/C.
---
$ \mathrm{d}W = V\,\mathrm{d}q = \frac{q}{C}\,\mathrm{d}q
Moving a further small charge dq across voltage V takes work V dq, since voltage is energy per unit charge.
---
$ W = \int_0^Q \frac{q}{C}\,\mathrm{d}q = \frac{Q^2}{2C}
Add up the work from zero charge to the final charge Q. This is the area under a straight line of V against q: a triangle of base Q and height Q/C.
---
$ U = \tfrac{1}{2}CV^2 = \tfrac{1}{2}QV
Substitute Q = CV. The average voltage during charging was half the final voltage, hence the ½.
:::

A defibrillator charges a capacitor of about 100 μF to roughly 2000 V, storing ½ × 10⁻⁴ × (2000)² = 200 J, then dumps it through the chest in a few milliseconds. A battery stores far more energy, but releases it slowly; a capacitor's strength is speed.

Where is the energy? Divide U by the volume Ad between the plates and use V = Ed and C = ε₀A/d: the energy per unit volume is ½ε₀E². The energy is stored in the field itself. This formula holds for any electric field, not just in capacitors, and it is part of why fields count as physically real.

::: note
A capacitor stores separated charge Q = CV, and with it energy ½CV², which lives in the electric field between the plates.
:::

::: history 1745–1746 · Pomerania and Leiden
Ewald von Kleist, a German cleric, and independently Pieter van Musschenbroek at Leiden, tried to store electricity in a glass jar of water held in the hand, with a wire into the water connected to a friction machine. When Musschenbroek touched the wire with his other hand he received a shock so violent he said he would not repeat it for the kingdom of France. The **Leyden jar**, later coated inside and out with metal foil, was the first capacitor. Franklin showed in the late 1740s that the charge resided on the glass between the conductors, not in the water, and wired several jars together into what he called a "battery".
:::

## Discharging

Connect the plates through a resistor and the charge flows back. The current is largest at first and dies away as the voltage falls; the charge decays exponentially, with a characteristic time RC. A 1000 μF capacitor through 1 kΩ takes about 1 second. Choosing R and C sets timing in circuits from flashing lights to radio tuners.

::: quiz
? A capacitor is charged to 10 V. You recharge it to 20 V. Its stored energy
- doubles
+ quadruples
- stays the same
= U = ½CV²; doubling V multiplies U by 4.
? You pull the plates of a charged, disconnected capacitor farther apart. The charge cannot change. What happens to the voltage?
+ It rises
- It falls
- It stays the same
= C = ε₀A/d falls as d grows, and V = Q/C, so V rises. The extra energy comes from the work you did pulling the attracting plates apart.
? Why does inserting a dielectric increase capacitance?
- It adds extra charge to the plates
+ Its polarized molecules partly cancel the field, lowering the voltage for the same charge
- It makes the plates larger
= Less field gives less voltage at the same Q, and C = Q/V.
:::
