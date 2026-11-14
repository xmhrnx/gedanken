By 1860 electricity and magnetism were described by a collection of separate laws, each found by experiment and each covering its own territory. Coulomb's law gave the force between charges. Ampère's law related a magnetic field to the current that produced it. Faraday's law described how a changing magnetic field drives a current. Nobody had checked whether these laws, taken together, made a consistent whole. James Clerk Maxwell, a Scottish physicist, set out to do exactly that. He gathered the laws into one system, found that one of them contradicted itself, and repaired it by adding a single new term. The result is a set of four equations that describe every electric and magnetic phenomenon outside the quantum world, and the new term turned out to explain what light is. This lesson goes through the four equations one at a time, and then looks closely at Maxwell's addition.

## What language do the equations use?

Each of the four equations is a statement about a **field**: a quantity that has a size and a direction at every point in space. **E** is the electric field, which at each point gives the force per unit charge that a small test charge would feel there. **B** is the magnetic field. Two ideas, both met in earlier lessons, run through all four equations, because they are the two natural ways of asking what a field is doing.

The first is **flux** through a surface: how much of the field passes through it. To find it, take each small patch of the surface, multiply its area by the part of the field pointing straight through it, and add up the results over the whole surface. When the surface is closed, like a balloon, the flux measures how much field comes out of it overall, with field going in counted as negative.

The second is **circulation** around a loop: how much the field goes round. Walk round the loop in small steps, take the part of the field pointing along each step, multiply by the step's length, and add the results all the way round. A field that swirls around the loop has a large circulation; one that points straight across it has none.

Two of Maxwell's equations are about flux through closed surfaces, and two are about circulation round loops.

## 1. Gauss's law: how do charges make electric fields?

::: equation Gauss's law
\oint \vec{E} \cdot \mathrm{d}\vec{A} = \frac{Q_{\text{enc}}}{\varepsilon_0}
:::

The electric flux out of any closed surface equals the total charge inside it, divided by a constant ε₀, called the **permittivity of free space**, about 8.85 × 10⁻¹² farads per metre. In terms of field lines, it says that lines start on positive charges and end on negative ones, and nowhere else. Count the lines leaving a balloon and you have counted the charge inside, because every line that leaves must have started on a charge within, and lines that pass right through the balloon count once going in and once coming out, and cancel.

This law contains Coulomb's law, and seeing how is a good illustration of its power.

::: steps Coulomb's law from Gauss's law
$ \oint \vec{E} \cdot \mathrm{d}\vec{A} = E \cdot 4\pi r^2
Take a single point charge Q and surround it with an imaginary sphere of radius r, centred on the charge. By symmetry the field must point straight outward everywhere on the sphere and have the same size E at every point, since nothing distinguishes one direction from another. So the flux is simply E times the sphere's area, 4πr².
---
$ E \cdot 4\pi r^2 = \frac{Q}{\varepsilon_0}
Gauss's law says this flux equals the charge enclosed, Q, divided by ε₀. The radius of the sphere was arbitrary, so this holds at every distance.
---
$ E = \frac{Q}{4\pi\varepsilon_0 r^2}
Divide both sides by 4πr². The field falls as the inverse square of the distance, because the same total flux is spread over a sphere whose area grows as r². Multiply by a second charge and you have Coulomb's law for the force between them.
:::

::: model field-lines
Change the size and sign of the second charge. Lines leave positive charges and end on negative ones; a closed surface drawn around both charges would count only the net charge inside.
:::

## 2. Gauss's law for magnetism: are there magnetic charges?

::: equation No magnetic monopoles
\oint \vec{B} \cdot \mathrm{d}\vec{A} = 0
:::

The magnetic flux out of any closed surface is always zero. Comparing this with Gauss's law for E shows what it means: there is no magnetic equivalent of charge for field lines to start or end on. Magnetic field lines never begin or stop; they form closed loops, so every line that enters a closed surface also leaves it. This is the precise statement of a fact from the magnetism chapter: cut a bar magnet in half and you get two smaller magnets, each with a north and a south pole, never an isolated pole. Searches for such magnetic monopoles have found none.

## 3. Faraday's law: how does a changing magnetic field make an electric field?

::: equation Faraday's law of induction
\oint \vec{E} \cdot \mathrm{d}\vec{l} = -\frac{\mathrm{d}\Phi_B}{\mathrm{d}t}
:::

Here Φ_B is the magnetic flux through a loop. When that flux changes, an electric field circulates around the loop, and if a wire lies along the loop, this field drives a current through it. Notice what the equation does not say: it does not mention a wire at all. The circulating electric field exists whether or not any conductor is there to feel it. The minus sign is Lenz's law, which says the induced current opposes the change that caused it. Every generator and transformer works by this law.

## 4. The Ampère–Maxwell law: what makes magnetic fields?

::: equation Ampère's law with Maxwell's correction
\oint \vec{B} \cdot \mathrm{d}\vec{l} = \mu_0 I_{\text{enc}} + \mu_0 \varepsilon_0 \frac{\mathrm{d}\Phi_E}{\mathrm{d}t}
:::

The first term on the right is Ampère's law from the magnetism chapter: a current I passing through a loop produces a magnetic field circulating around it. The constant μ₀ is the **permeability of free space**, about 1.26 × 10⁻⁶ newtons per ampere squared. The second term is Maxwell's addition. It says that a changing electric flux Φ_E through the loop also produces a circulating magnetic field, exactly as if a current were flowing. Maxwell called this term the **displacement current**, though no charge actually moves.

## Why was the extra term needed?

Maxwell did not add the term on a whim. Ampère's original law, with only the current term, contradicts itself, and the contradiction appears in a very ordinary situation: charging a **capacitor**, a pair of metal plates separated by a small gap. While a battery charges it, current flows along the wire onto one plate and off the other, but no charge crosses the gap.

Draw a loop around the wire leading to one plate. Ampère's law is stated in terms of the current passing through the loop, which means the current crossing a surface whose edge is the loop, and an important part of the law is that any such surface must give the same answer. The obvious choice is a flat disc spanning the loop, which the wire pierces, so the current through it is I. But you could equally choose a surface shaped like a bag, with the same loop as its rim, that bulges out to pass through the gap between the plates. No wire crosses that surface, so the current through it is zero. The same loop now has two different circulations, depending on which surface you happened to choose, and that is impossible. Something is missing, and Maxwell saw what it was: in the gap, where there is no current, the electric field is growing as the plates charge up.

::: steps The displacement current in a charging capacitor
$ E = \frac{Q}{\varepsilon_0 A}
Between two parallel plates of area A carrying charges +Q and −Q, the electric field is very nearly uniform, apart from small effects at the edges. Gauss's law, applied to a small box enclosing part of one plate, gives its size: the charge per unit area divided by ε₀.
---
$ \Phi_E = EA = \frac{Q}{\varepsilon_0}
The bag-shaped surface passes through the gap, where all of this field crosses it. The electric flux through it is the field times the plate area, and the area cancels, leaving the charge on the plate divided by ε₀.
---
$ \frac{\mathrm{d}\Phi_E}{\mathrm{d}t} = \frac{1}{\varepsilon_0}\frac{\mathrm{d}Q}{\mathrm{d}t}
As the capacitor charges, Q grows, so the flux grows too. Taking the rate of change of both sides, the constant 1/ε₀ simply comes along.
---
$ \frac{\mathrm{d}\Phi_E}{\mathrm{d}t} = \frac{I}{\varepsilon_0}
The charge on the plate grows at exactly the rate the current in the wire delivers it, so dQ/dt is the current I. The changing flux in the gap is tied directly to the current in the wire.
---
$ \mu_0 \varepsilon_0 \frac{\mathrm{d}\Phi_E}{\mathrm{d}t} = \mu_0 I
Multiply both sides by μ₀ε₀. The left side is Maxwell's new term evaluated on the bag-shaped surface, and the right side is the old current term evaluated on the flat surface. They are equal, so both surfaces now give the same circulation, and the contradiction is gone.
:::

In everyday circuits the new term is tiny compared with ordinary currents, which is why Ampère, working with steady currents in wires, never noticed its absence. Its importance lies elsewhere. It means that a changing electric field makes a magnetic field, just as Faraday's law says that a changing magnetic field makes an electric field. With that symmetry in place, the two fields can sustain each other in empty space, with no charges or currents anywhere, and travel off as a wave. That is what light turns out to be, as the next lesson shows.

## Can the laws be written at a single point?

The four equations above describe loops and surfaces of any size. They can also be shrunk down to describe what happens at a single point, which is often more convenient. Two new operations are needed. The **divergence** of a field, written ∇·, measures how much the field spreads outward from a point, as if the point were a source. The **curl**, written ∇×, measures how much the field circulates around a point. In this form, ρ is the charge density, the charge per unit volume, and **J** is the current density, the current per unit area.

::: equation Maxwell's equations in differential form
\nabla \cdot \vec{E} &= \frac{\rho}{\varepsilon_0} \\ \nabla \cdot \vec{B} &= 0 \\ \nabla \times \vec{E} &= -\frac{\partial \vec{B}}{\partial t} \\ \nabla \times \vec{B} &= \mu_0 \vec{J} + \mu_0 \varepsilon_0 \frac{\partial \vec{E}}{\partial t}
:::

Read line by line, these say the same four things as before: charge is a source of electric field; magnetic field has no sources; a changing magnetic field makes the electric field curl around it; and currents and changing electric fields make the magnetic field curl around them. Together with the Lorentz force, which says how the fields push on charges, they are the whole of classical electromagnetism.

::: note
Charges make electric fields; there are no magnetic charges; a changing magnetic field makes an electric field; and currents and changing electric fields make magnetic fields. Maxwell's added term, the displacement current, removed a contradiction in Ampère's law and completed the symmetry between E and B that allows electromagnetic waves to exist.
:::

::: history 1861–1865 · London
Maxwell reached the displacement current by an unlikely route. In a paper called "On Physical Lines of Force", published in parts in 1861 and 1862, he tried to picture what space must be like to carry electric and magnetic forces. He imagined it filled with tiny spinning vortices, with small "idle wheels" between them, like ball bearings, which could be displaced slightly when an electric force acted. That displacement of the wheels is where the name "displacement current" comes from. The machinery was entirely wrong, and Maxwell himself did not take it literally, but the equations he extracted from it survived. In "A Dynamical Theory of the Electromagnetic Field" in 1865 he dropped the mechanical model altogether and wrote the theory directly in terms of fields, as about twenty equations in twenty unknowns, each component written out separately. The compact four-equation form used today was worked out between 1884 and 1890 by Oliver Heaviside, a self-taught English telegraph engineer, and independently by Heinrich Hertz in Germany, both using the new notation of vectors.
:::

::: quiz
? Which of Maxwell's equations explains why an electric generator works?
- Gauss's law
- Gauss's law for magnetism
+ Faraday's law
- The Ampère–Maxwell law
= A rotating coil in a magnet's field sees a changing magnetic flux. Faraday's law says a changing flux produces a circulating electric field, which drives a current round the coil.
? What does the zero on the right of Gauss's law for magnetism mean physically?
+ Magnetic field lines have no beginning or end, so there are no isolated magnetic poles
- Magnetic fields are always zero inside closed surfaces
- Magnets have no effect on charges
= A zero net flux means every field line that enters a closed surface also leaves it. So nothing inside can be a source or sink of magnetic field, which is to say there is no magnetic equivalent of a lone electric charge.
? Why did Maxwell add the displacement current term?
- To explain permanent magnets
- To make light travel faster
+ Without it, Ampère's law gives different answers for different surfaces bounded by the same loop
= Around a charging capacitor, one surface is crossed by the current in the wire and another passes through the gap where there is no current. The changing electric flux in the gap supplies exactly the missing term, so both surfaces agree.
:::
