The equivalence principle already tells us that light must bend as it passes a massive body. What it does not tell us is how much, and that turned out to be the question on which Einstein's reputation was made. A calculation in the style of Newton, treating light as a fast particle, gives one answer. General relativity gives exactly twice as much. In 1919 a total eclipse of the Sun gave astronomers a chance to tell the two apart, and the result made Einstein the most famous scientist in the world. This lesson works out both answers, explains where the factor of two comes from, and tells the story of the measurement.

## What would Newton's physics predict?

In the 1700s many natural philosophers, following Newton, thought of light as a stream of tiny particles. If so, gravity should pull on them like on anything else. Since the acceleration of a falling body does not depend on its mass, it does not even matter how much a light particle weighs: it should be deflected by the same amount as a speeding cannonball travelling at the speed of light. The Bavarian astronomer Johann von Soldner did this calculation in 1801. We can repeat it.

Picture a particle of light travelling at speed c on a path that, if the Sun were not there, would pass the Sun's centre at a closest distance b. This distance is called the **impact parameter**. The deflection will turn out to be tiny, so to a very good approximation the particle travels in a straight line at constant speed, and all we need is the small sideways velocity that the Sun's pull adds to it during the passage. Set up a coordinate x along the straight path, with x = 0 at the point of closest approach.

::: steps Newtonian deflection of a light particle
$ a = \frac{GM}{b^2 + x^2}
When the particle is at position x along its path, its distance from the Sun's centre is √(b² + x²) by Pythagoras' theorem, so Newton's law gives an acceleration of this size, pointing toward the Sun. Far from closest approach, x is large and the pull is weak; near x = 0 it is strongest.
---
$ a_\perp = \frac{GM}{b^2 + x^2}\cdot\frac{b}{\sqrt{b^2 + x^2}} = \frac{GMb}{(b^2 + x^2)^{3/2}}
Only part of this acceleration bends the path. The part along the path speeds the particle up on approach and slows it by the same amount as it leaves, so it has no net effect. The part across the path is what deflects it, and by similar triangles it is the fraction b/√(b² + x²) of the whole.
---
$ \mathrm{d}t = \frac{\mathrm{d}x}{c}
To add up the effect of this acceleration we need to know how long the particle spends near each point. Since it moves along the nearly straight path at speed c, it takes a time dx/c to cross a short stretch of length dx. This lets us trade time, which the acceleration acts over, for distance along the path, which our formula uses.
---
$ \Delta v_\perp = \int_{-\infty}^{\infty} a_\perp\,\mathrm{d}t = \frac{GMb}{c}\int_{-\infty}^{\infty}\frac{\mathrm{d}x}{(b^2 + x^2)^{3/2}}
The total sideways velocity gained is the sideways acceleration added up over every moment of the passage, from far away on the approaching side to far away on the departing side. Replacing dt with dx/c turns this into an integral over the path, and the constants GMb/c can be taken outside it.
---
$ \int_{-\infty}^{\infty}\frac{\mathrm{d}x}{(b^2 + x^2)^{3/2}} = \frac{2}{b^2}
This is a standard integral. Substituting x = b tan θ, where θ is the angle between the line to the Sun and the line of closest approach, turns it into (1/b²) times the integral of cos θ from −90° to +90°. That integral is 2, so the whole is 2/b².
---
$ \Delta v_\perp = \frac{2GM}{bc}
Put the value of the integral into the previous line: GMb/c times 2/b² gives 2GM/(bc). This is the sideways speed the particle has picked up by the time it is far away again. It is larger for a heavier Sun and smaller for a wider pass.
---
$ \delta_{\text{Newton}} = \frac{\Delta v_\perp}{c} = \frac{2GM}{c^2 b}
Finally, convert the sideways speed into an angle. For a small deflection, the angle in radians is the sideways speed divided by the forward speed. The particle's mass appears nowhere, as we expected, because in free fall mass always cancels.
:::

Now put in the numbers for light that just grazes the edge of the Sun. The impact parameter b is then the Sun's radius, about 6.96 × 10⁵ km, and the quantity 2GM/c² for the Sun is about 2.95 km. Their ratio is about 4.2 × 10⁻⁶ radians, or 0.87 arcseconds. An arcsecond is 1/3600 of a degree, about the width of a coin seen from 4 kilometres away. So the predicted shift is minute, but not hopelessly so for a good telescope.

When Einstein calculated the bending in 1911, using only the equivalence principle and the slowing of clocks in gravity, he got essentially the same number. He did not yet have the full theory.

## Why does Einstein's theory give twice as much?

When Einstein completed general relativity in November 1915, he redid the calculation and found double the Newtonian value:

::: equation Deflection in general relativity
\delta_{\text{GR}} = \frac{4GM}{c^2 b} \approx 1.75''\ \text{at the Sun's edge}
:::

The extra half comes from the curvature of space. In the weak gravity of the Sun, the metric (the rule converting coordinate steps into real distances and times) changes both the time part and the space part by the same small amount, with opposite signs:

::: equation The weak-field metric
\mathrm{d}s^2 \approx {}&-\left(1 + \frac{2\Phi}{c^2}\right)c^2\,\mathrm{d}t^2 \\ &+ \left(1 - \frac{2\Phi}{c^2}\right)\left(\mathrm{d}x^2 + \mathrm{d}y^2 + \mathrm{d}z^2\right)
:::

Here Φ = −GM/r is the Newtonian potential. Now compare how a slow object and a ray of light respond. A slow object moves only a small distance through space in each second, so the space term, multiplied by its tiny dx², hardly affects it. Only the time term matters, and the previous lessons showed that the time term alone reproduces Newton's gravity. That is why Newton's theory works so well for planets.

Light is different. For a light ray, the distance covered in space is always c times the time taken, so dx² is as large as c²dt². The space term therefore affects light just as strongly as the time term does. The time term's contribution is the Newtonian 0.87 arcseconds; the space term, being the same size, adds another 0.87 arcseconds. Newton's theory has no curvature of space at all, so it captures only the first half.

::: note
Half of the bending of starlight comes from the warping of time, which any theory that obeys the equivalence principle must contain. The other half comes from the curvature of space, which only general relativity has. Measuring the angle therefore tests the whole theory, not just the equivalence principle.
:::

## How do you see a star next to the Sun?

The difficulty is obvious: stars near the Sun's edge in the sky are drowned by its glare. The only time they can be photographed is during a total solar eclipse, when the Moon blocks the Sun's disc for a few minutes. The method is to photograph the field of stars around the eclipsed Sun, then photograph the same field months later at night, when the Sun is in another part of the sky, and compare the two. If light is bent, stars near the Sun's edge appear pushed slightly outward, away from the Sun, in the eclipse photographs. The predicted shift falls off as 1/b, so stars farther out move less.

::: history 1919 · Príncipe and Sobral
The eclipse of 29 May 1919 had an unusual advantage: the Sun would be in front of the Hyades, a cluster of bright stars, so there would be plenty to measure. Frank Dyson, the Astronomer Royal, and Arthur Eddington, who had championed Einstein's theory in Britain during the war, organised two expeditions. Eddington and Edwin Cottingham went to the island of Príncipe off the west coast of Africa; Andrew Crommelin and Charles Davidson went to Sobral in northern Brazil. On Príncipe the morning brought heavy cloud and rain, and the sky cleared only partway through totality, so only a few plates showed enough stars. They gave a deflection of 1.61 ± 0.30 arcseconds. At Sobral the weather was good. A smaller telescope there gave 1.98 ± 0.12 arcseconds, while the main telescope gave a value near the Newtonian 0.9; its star images were blurred, probably because sunlight had warmed and distorted its mirror, and its results were set aside. On 6 November 1919 Dyson announced at a joint meeting of the Royal Society and the Royal Astronomical Society that the results favoured Einstein. The next day The Times of London carried the headline "Revolution in Science", and Einstein became a household name almost overnight. Critics later argued that Eddington, already convinced of relativity, had chosen the data he liked. When the Sobral plates were remeasured with modern equipment in 1979, they supported the original conclusion.
:::

Einstein was lucky in his timing. In 1914 the German astronomer Erwin Freundlich had led an expedition to Crimea to observe an eclipse there, at a time when Einstein's prediction was still the Newtonian-sized value from 1911. The First World War broke out just before the eclipse, the astronomers were interned as enemy aliens, and no measurement was made. Had it succeeded, it would have contradicted his prediction, which he corrected only in 1915.

The 1919 measurements were rough, with uncertainties of 10 to 20 percent, and later eclipse expeditions did little better. The decisive tests came with radio astronomy. Radio telescopes can track distant quasars, very bright cores of far-off galaxies, as the Sun passes in front of them, without waiting for an eclipse. These measurements now agree with general relativity to a few parts in ten thousand.

## What happens when the lens is a galaxy?

The same effect operates on a far larger scale. When a massive galaxy, or a cluster of galaxies, lies almost exactly between us and a more distant galaxy, its gravity bends the distant light around it, acting like a crude lens. Depending on the alignment, the background galaxy can appear as several separate images, as stretched arcs, or, when the alignment is nearly perfect, as a complete circle of light called an **Einstein ring**. Einstein himself considered the idea in 1936 and thought it would never be observed; the first lensed quasar was found in 1979. Because the amount of bending depends on the mass doing the bending, not on whether that mass shines, lensing is now a standard way of weighing galaxy clusters, including the dark matter in them, which emits no light of its own.

::: model gravity-well
The curved surface shows the curvature of space around a compact mass. Slow objects barely notice it, but light feels it in full, and it supplies the half of the bending that Newton's theory misses.
:::

::: quiz
? Why does general relativity predict twice the Newtonian deflection of light?
- Light has twice the effective mass in Einstein's theory
+ Light also responds to the curvature of space, which adds an equal amount of bending
- The Sun's gravity is stronger in general relativity
= The curvature of time alone reproduces the Newtonian value. Light moves as far through space as through time, so the curvature of space, which Newton's theory lacks, adds an equal amount.
? A ray passes a star at twice the star's radius instead of grazing its surface. How does its deflection compare?
- It is the same
+ It is halved
- It is quartered
= The deflection is δ = 4GM/(c²b), which is proportional to 1/b. Doubling the impact parameter b halves the angle.
? Why was a solar eclipse needed to measure the bending in 1919?
- The Moon's gravity adds to the effect
+ Stars near the Sun can only be photographed when its light is blocked
- Light bends only when the Sun is eclipsed
= The bending happens all the time, but stars close to the Sun's edge are lost in its glare except during totality. Radio telescopes later avoided this problem by tracking quasars instead.
:::
