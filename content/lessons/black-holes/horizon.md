The Schwarzschild metric misbehaves at two radii: at r = r_s a term blows up, and at r = 0 everything does. For decades many physicists, Einstein included, assumed both signalled that the solution stops being physical. The resolution is that the two are completely different kinds of trouble.

## Coordinate singularity versus real singularity

An infinity in a formula may belong to the geometry or only to the labels you chose. The North Pole is a familiar example: in latitude and longitude, every longitude meets there and "which direction is east" has no answer, yet nothing odd happens to the ground. A different map makes the pole an ordinary point. That is a **coordinate singularity**.

To tell the two apart, compute something that does not depend on coordinates. In relativity the standard test is a **curvature invariant**, a number built from the curvature that every observer and every coordinate system agrees on. For the Schwarzschild geometry the simplest one, the Kretschmann scalar, is:

::: equation Curvature invariant of the Schwarzschild geometry
K &= \frac{12 r_s^2}{r^6} \\ \text{At } r = r_s\text{:}\quad K &= \frac{12}{r_s^4} \quad \text{(finite)} \\ \text{At } r \to 0\text{:}\quad K &\to \infty
:::

At the Schwarzschild radius the curvature is finite, and for a big black hole it is small: since K falls as 1/r_s⁴, the more massive the hole, the gentler its horizon. At r = 0 the curvature is truly infinite. That is a **real singularity**, a place where the theory itself predicts unbounded tidal forces and stops making sense.

In 1933 Georges Lemaître showed that the trouble at r_s disappears in coordinates attached to a freely falling observer. In 1958 David Finkelstein rediscovered coordinates Arthur Eddington had written down in 1924 without noticing their meaning (now Eddington–Finkelstein coordinates) and showed that infalling paths cross r_s smoothly. Nothing stops an object there. The radius is crossable, but it is a one-way surface: the **event horizon**.

## Time dilation, read off the metric

What makes r_s special shows up in how clocks behave near it.

::: steps Gravitational time dilation
$ ds^2 = {}&-\left(1 - \frac{r_s}{r}\right)c^2\,\mathrm{d}t^2 \\ &+ \frac{\mathrm{d}r^2}{1 - \frac{r_s}{r}} \\ &+ r^2\left(\mathrm{d}\theta^2 + \sin^2\theta\,\mathrm{d}\phi^2\right)
Start from the Schwarzschild metric. Recall t is the time on a clock at rest far away.
---
$ ds^2 = -\left(1 - \frac{r_s}{r}\right)c^2\,\mathrm{d}t^2
Consider a clock held at rest at fixed r, θ, φ (on a rocket hovering in place). Then dr = dθ = dφ = 0, and only the time term survives.
---
$ \mathrm{d}\tau^2 = -\frac{ds^2}{c^2} = \left(1 - \frac{r_s}{r}\right)\mathrm{d}t^2
The time a clock actually reads, its proper time τ, is defined by dτ² = −ds²/c².
---
$ \mathrm{d}\tau = \sqrt{1 - \frac{r_s}{r}}\,\mathrm{d}t
Take the square root. During a distant hour dt, the hovering clock ticks off less than an hour.
:::

At r = 2r_s the factor is √0.5 ≈ 0.71. At r = 1.01 r_s it is about 0.1. As r approaches r_s it goes to zero: from far away, a clock near the horizon appears to stop. Light from it is stretched by the same factor, so it is redshifted to ever longer wavelengths and dimmer.

## Two descriptions of one fall

Suppose a probe falls in. A distant observer watches its signals arrive further and further apart and further into the red. The probe seems to slow as it nears r_s and fade, never quite seen to cross. In t, crossing takes infinitely long.

The probe's own clock tells a different story. Its proper time to reach r_s is finite, and it notices nothing local as it crosses: no wall, no flash. For a ten-solar-mass black hole, the time from horizon to singularity is at most about πGM/c³ ≈ 0.15 milliseconds of its own time; for the four-million-solar-mass hole at the centre of our galaxy, about a minute.

Both accounts are correct. They disagree because "when" is not shared between observers in relativity, and the distant observer's coordinate t is exactly the one that breaks at r_s.

## Light cones tip over

A **light cone** marks the paths light can take from an event; anything with mass must stay inside it, which is what "nothing travels faster than light" means geometrically. For light moving radially, ds² = 0, which gives:

::: equation Radial light speed in Schwarzschild coordinates
\frac{\mathrm{d}r}{\mathrm{d}t} = \pm c\left(1 - \frac{r_s}{r}\right)
:::

Far away this is ±c. Closer in, both outgoing and ingoing light seem to slow in the distant observer's coordinates, and the cones narrow and tilt inward. In infalling coordinates the picture is clean: at r_s the outward edge of the cone points straight up in time, so even outgoing light stays at r_s. Inside, the whole cone points toward smaller r.

::: model light-cones
Light cones along a line into the black hole. At the horizon the outer edge becomes vertical; inside, every future direction leads inward.
:::

## Why inside, r becomes time

Look at the signs in the metric. Outside, the dt² term is negative and the dr² term positive: t behaves like time, r like space. For r < r_s, 1 − r_s/r is negative, so both signs flip. Now r plays the role of time.

That is the content of the horizon. Outside, you can choose to move toward or away from the centre, as you can choose to walk north or south. Inside, decreasing r is no more avoidable than next Tuesday. Firing rockets only changes how you reach r = 0, not whether. The singularity is not a place in front of you; it is a moment in your future.

::: note
The event horizon is not a material surface and nothing local marks it. It is the boundary of the region from which no future-directed path, even light's, reaches the outside. The real singularity at r = 0 is where the theory gives infinite curvature and stops working.
:::

::: model gravity-well
The well deepens toward the centre. The horizon marks where escape would require going faster than light.
:::

::: quiz
? How can you tell r = r_s is only a coordinate singularity?
- Light slows down there
+ A coordinate-independent curvature invariant stays finite there
- The metric has no terms at r = r_s
= K = 12r_s²/r⁶ is finite at r = r_s and infinite only at r = 0.
? A clock hovers at r = 4r_s/3. Compared with a distant clock, it runs at
- 3/4 the rate
+ 1/2 the rate
- 1/4 the rate
= √(1 − r_s/r) = √(1 − 3/4) = √(1/4) = 1/2.
? Why can't an object inside the horizon avoid r = 0?
- Its rockets are not powerful enough
- The tidal forces pull it in
+ Inside, decreasing r is the future direction of time
= For r < r_s the metric's signs swap, and every future-pointing path has r decreasing.
:::
