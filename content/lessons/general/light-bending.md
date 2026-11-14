The equivalence principle says light must bend in a gravitational field. The question that made Einstein famous was how much. Newtonian reasoning gives one answer, general relativity gives exactly twice as much, and in 1919 a total eclipse of the Sun gave astronomers the chance to tell them apart.

## The Newtonian answer

Treat light as a fast particle moving at c past the Sun, as Michell and others did in the 1700s. It passes at closest distance b from the Sun's centre (the **impact parameter**). The deflection is small, so the path is nearly straight, and we only need the sideways velocity the Sun's pull adds along the way.

::: steps Newtonian deflection of a light particle
$ a_\perp = \frac{GM}{b^2 + x^2}\cdot\frac{b}{\sqrt{b^2 + x^2}} = \frac{GMb}{(b^2 + x^2)^{3/2}}
At position x along the path (x = 0 at closest approach), the pull GM/(b² + x²) points toward the Sun. The fraction b/√(b² + x²) of it points sideways, across the path.
---
$ \mathrm{d}t = \frac{\mathrm{d}x}{c}
The particle moves along the nearly straight path at speed c, so the time spent near each point is dx/c.
---
$ \Delta v_\perp = \int_{-\infty}^{\infty} a_\perp\,\mathrm{d}t = \frac{GMb}{c}\int_{-\infty}^{\infty}\frac{\mathrm{d}x}{(b^2 + x^2)^{3/2}}
The total sideways velocity is the sideways acceleration added up over the whole passage, from far away on one side to far away on the other.
---
$ \int_{-\infty}^{\infty}\frac{\mathrm{d}x}{(b^2 + x^2)^{3/2}} = \frac{2}{b^2}
A standard integral: substitute x = b tan θ and it becomes (1/b²)∫cos θ dθ from −90° to 90°, which is 2/b².
---
$ \Delta v_\perp = \frac{2GM}{bc}
Combine the two lines above.
---
$ \delta_{\text{Newton}} = \frac{\Delta v_\perp}{c} = \frac{2GM}{c^2 b}
For a small angle, the deflection is the sideways speed divided by the forward speed. The particle's mass has cancelled, as it always does in free fall.
:::

For light grazing the Sun, b is the solar radius, 6.96 × 10⁵ km, and 2GM/c² is 2.95 km, so δ ≈ 4.2 × 10⁻⁶ radians, or **0.87 arcseconds**. (One arcsecond is 1/3600 of a degree, about the width of a coin seen from 4 km.) Einstein's own 1911 calculation, using only the equivalence principle, gave essentially this value.

## Einstein's answer

With the full theory in November 1915, Einstein found twice as much:

::: equation Deflection in general relativity
\delta_{\text{GR}} = \frac{4GM}{c^2 b} \approx 1.75\text{ arcseconds at the Sun's edge}
:::

The extra factor comes from the curvature of space. The weak-field metric bends both time and space by the same small amount:

::: equation The weak-field metric
\mathrm{d}s^2 \approx -\left(1 + \frac{2\Phi}{c^2}\right)c^2\,\mathrm{d}t^2 + \left(1 - \frac{2\Phi}{c^2}\right)\left(\mathrm{d}x^2 + \mathrm{d}y^2 + \mathrm{d}z^2\right)
:::

A slow object barely moves through space, so only the time term matters, and that term reproduces Newton. For light, the distance travelled dx equals c dt, so the space term counts exactly as much as the time term, and that doubles the bending. Newton's theory has no curved space at all, so it misses this half.

::: note
Half of the bending of starlight comes from the warping of time, which any theory obeying the equivalence principle must have. The other half comes from the curvature of space, which only general relativity has. Measuring the angle tests the whole theory, not just the equivalence principle.
:::

## The eclipse

Stars near the Sun are hidden by its glare except during a total eclipse. The method is to photograph the star field around the eclipsed Sun, photograph the same field at night months later when the Sun is elsewhere, and compare. Stars close to the Sun's edge should appear pushed outward, away from it.

::: history 1919 · Príncipe and Sobral
The eclipse of 29 May 1919 crossed the bright stars of the Hyades cluster. Frank Dyson, the Astronomer Royal, and Arthur Eddington organized two British expeditions: Eddington and Edwin Cottingham went to the island of Príncipe off West Africa, and Andrew Crommelin and Charles Davidson to Sobral in Brazil. Clouds covered Príncipe until the last part of totality, and only a few plates were usable; they gave 1.61 ± 0.30″. Sobral's smaller telescope gave 1.98 ± 0.12″, while its main telescope, whose images were blurred, perhaps by the Sun heating the mirror, gave a value nearer 0.9″ and was set aside. On 6 November 1919 Dyson announced that the results favoured Einstein. The next day *The Times* of London ran the headline "Revolution in Science", and Einstein became a household name. Critics have argued Eddington, a convinced relativist, chose the data he liked; a 1979 remeasurement of the Sobral plates supported the original conclusion.
:::

Einstein was lucky in his timing. A German expedition had gone to Crimea for the eclipse of 1914, when his prediction was still the Newtonian-sized 0.87″. The First World War broke out; the astronomers were interned and no measurement was made. Had it succeeded, it would have contradicted his first prediction.

The 1919 numbers were rough, with uncertainties of 10 to 20 percent. Radio telescopes now measure the bending of signals from distant quasars as they pass the Sun, and agree with general relativity to a few parts in ten thousand.

## Gravitational lenses

The same effect, on the scale of galaxies, makes the universe full of lenses. A massive galaxy or cluster between us and a more distant source bends its light into arcs, multiple images, and sometimes a complete **Einstein ring**. Lensing is now a standard tool for weighing galaxy clusters, including the dark matter in them, which emits no light of its own.

::: model gravity-well
The curved surface shows the spatial curvature that slow objects barely notice but light feels in full: the half of the bending that Newton's theory misses.
:::

::: quiz
? Why does general relativity predict twice the Newtonian deflection of light?
- Light has twice the effective mass in Einstein's theory
+ Light also responds to the curvature of space, which adds an equal amount of bending
- The Sun's gravity is stronger in general relativity
= Time curvature alone reproduces the Newtonian value; space curvature, which Newton lacks, contributes an equal extra amount.
? Light passes a star at twice the solar radius instead of grazing it. Its deflection is
- the same
+ halved
- quartered
= δ = 4GM/(c²b) is proportional to 1/b.
? Why was a solar eclipse needed to measure the bending in 1919?
- The Moon's gravity adds to the effect
+ Stars near the Sun can only be photographed when its light is blocked
- Light bends only when the Sun is eclipsed
= The bending happens all the time, but stars close to the Sun's edge are drowned by sunlight except during totality.
:::
