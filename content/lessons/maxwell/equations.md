By 1860 electricity and magnetism were described by a handful of separate laws, each found by experiment: Coulomb's law for charges, Ampère's law for currents, Faraday's law of induction. James Clerk Maxwell gathered them into one consistent set, found a gap, and filled it. The result is four equations that describe every electric and magnetic phenomenon outside the quantum world.

Each equation is about a **field**: a quantity with a value and direction at every point in space. **E** is the electric field (force per unit charge), **B** the magnetic field. Two ideas recur:

- **Flux** through a surface: how much field passes through it, the field's perpendicular component times area, added up over the surface.
- **Circulation** around a loop: the field's component along the loop, added up all the way round.

## 1. Gauss's law: charges are sources of E

::: equation Gauss's law
\oint \vec{E} \cdot \mathrm{d}\vec{A} = \frac{Q_{\text{enc}}}{\varepsilon_0}
:::

The electric flux out of any closed surface equals the charge inside, divided by a constant ε₀ (the permittivity of free space, about 8.85 × 10⁻¹² F/m). Field lines start on positive charges and end on negative ones. For a single point charge, the flux spreads over a sphere of area 4πr², which is where Coulomb's inverse-square law comes from.

::: model field-lines
Change the second charge. Lines leave positive charges and end on negative ones; a closed surface around both counts only the net charge inside.
:::

## 2. Gauss's law for magnetism: no magnetic charges

::: equation No magnetic monopoles
\oint \vec{B} \cdot \mathrm{d}\vec{A} = 0
:::

The magnetic flux out of any closed surface is zero. Magnetic field lines never start or stop; they form closed loops. Cut a bar magnet in half and you get two smaller magnets, each with a north and a south pole, never an isolated pole. Searches for such "magnetic monopoles" have found none.

## 3. Faraday's law: changing B makes E

::: equation Faraday's law of induction
\oint \vec{E} \cdot \mathrm{d}\vec{l} = -\frac{\mathrm{d}\Phi_B}{\mathrm{d}t}
:::

Φ_B is the magnetic flux through a loop. When it changes, an electric field circulates around the loop and drives current through any wire lying along it. The minus sign (Lenz's law) says the induced current opposes the change that caused it. Every generator and transformer works this way.

## 4. Ampère–Maxwell law: currents and changing E make B

::: equation Ampère's law with Maxwell's correction
\oint \vec{B} \cdot \mathrm{d}\vec{l} = \mu_0 I_{\text{enc}} + \mu_0 \varepsilon_0 \frac{\mathrm{d}\Phi_E}{\mathrm{d}t}
:::

A current I through a loop produces a magnetic field circulating around it; μ₀ is the permeability of free space, about 1.26 × 10⁻⁶ N/A². The second term is Maxwell's addition: a changing electric flux Φ_E acts like a current, which he called the **displacement current**.

## Why the extra term is needed

Ampère's original law had only the current term, and it contradicts itself. Charge a capacitor, two metal plates with a gap. A loop around the wire leading to it has current I passing through it. But the law lets you choose any surface bounded by the loop, and a surface that bulges out to pass between the plates has no current through it at all. The same loop would have two different circulations.

::: steps The displacement current in a charging capacitor
$ E = \frac{Q}{\varepsilon_0 A}
Between two plates of area A carrying charges +Q and −Q, the field is uniform; this follows from Gauss's law.
---
$ \Phi_E = EA = \frac{Q}{\varepsilon_0}
The electric flux through the bulging surface, which passes between the plates, is the field times the plate area.
---
$ \frac{\mathrm{d}\Phi_E}{\mathrm{d}t} = \frac{1}{\varepsilon_0}\frac{\mathrm{d}Q}{\mathrm{d}t} = \frac{I}{\varepsilon_0}
The charge on the plates grows at the rate the current delivers it.
---
$ \mu_0 \varepsilon_0 \frac{\mathrm{d}\Phi_E}{\mathrm{d}t} = \mu_0 I
So the new term gives, through the gap, exactly what the current term gives through the wire. Both surfaces now agree.
:::

The correction is tiny in everyday circuits, which is why Ampère never saw it. Its importance is elsewhere: it means a changing E makes B, just as Faraday's law says a changing B makes E. That symmetry is what allows light to exist, as the next lesson shows.

## The differential form

The same laws can be written at a single point, using the divergence ∇· (how much a field spreads out from a point) and the curl ∇× (how much it circulates around a point). ρ is charge density and **J** current density.

::: equation Maxwell's equations in differential form
\nabla \cdot \vec{E} &= \frac{\rho}{\varepsilon_0} \\ \nabla \cdot \vec{B} &= 0 \\ \nabla \times \vec{E} &= -\frac{\partial \vec{B}}{\partial t} \\ \nabla \times \vec{B} &= \mu_0 \vec{J} + \mu_0 \varepsilon_0 \frac{\partial \vec{E}}{\partial t}
:::

::: note
Charges make electric fields; there are no magnetic charges; a changing magnetic field makes an electric field; currents and changing electric fields make magnetic fields. Maxwell's added term completes the symmetry between E and B.
:::

::: history 1861–1865 · London
Maxwell introduced the displacement current in "On Physical Lines of Force" (1861–62), where he pictured space filled with spinning molecular vortices and idle wheels between them. The machinery was wrong, but the equations survived. In "A Dynamical Theory of the Electromagnetic Field" (1865) he dropped the model and wrote the theory as about twenty equations in twenty variables, component by component. The compact four-equation vector form is due to Oliver Heaviside and, independently, Heinrich Hertz, between 1884 and 1890.
:::

::: quiz
? Which equation explains why an electric generator works?
- Gauss's law
- Gauss's law for magnetism
+ Faraday's law
- The Ampère–Maxwell law
= A rotating coil sees a changing magnetic flux, and Faraday's law says that drives a circulating electric field, and so a current.
? What does the zero on the right of Gauss's law for magnetism mean physically?
+ Magnetic field lines have no beginning or end, so there are no isolated magnetic poles
- Magnetic fields are always zero inside closed surfaces
- Magnets have no effect on charges
= Every line entering a closed surface also leaves it; there is no magnetic equivalent of a lone charge.
? Why did Maxwell add the displacement current term?
- To explain magnets
- To make light faster
+ Without it, Ampère's law gives different answers for different surfaces bounded by the same loop
= Around a charging capacitor, one surface catches the current and another does not; the changing electric flux in the gap supplies the missing term.
:::
