The equivalence principle says gravity can be made to vanish in a small falling box. It cannot be made to vanish everywhere at once: tidal effects remain. Einstein's answer to what those leftover effects are was the boldest idea in the theory. Gravity is not a force at all. It is the geometry of spacetime, and tidal effects are its curvature.

## Straight lines on a curved surface

On a flat sheet, two straight lines that start parallel stay parallel. On a sphere they do not. Two travellers leave the equator 100 km apart, both heading due north along the straightest possible paths (lines of longitude). Neither turns, yet the distance between them shrinks, and they meet at the pole. Neither felt a force; the convergence comes from the surface they walk on.

Free fall works the same way. A **geodesic** is the straightest possible path through a geometry. In Einstein's theory, a freely falling object follows a geodesic through spacetime and feels no force, which is exactly what the equivalence principle says it should feel. Two neighbouring falling objects drift together or apart, not because something pushes them, but because the geodesics they follow converge or diverge.

## Tidal acceleration is the real gravity

How big is the part of gravity that cannot be removed? Take two objects falling toward a mass M, one at distance r from its centre and one a small height h higher.

::: steps Tidal acceleration between two falling objects
$ a(r) = \frac{GM}{r^2}
Newton's gravitational acceleration at distance r. A uniform acceleration like this could be cancelled by falling with it.
---
$ a(r + h) = \frac{GM}{(r + h)^2}
The upper object is slightly farther away and is pulled slightly less.
---
$ \frac{1}{(r + h)^2} \approx \frac{1}{r^2}\left(1 - \frac{2h}{r}\right)
For h much smaller than r, use (1 + x)⁻² ≈ 1 − 2x with x = h/r.
---
$ \Delta a = a(r) - a(r + h) \approx \frac{2GMh}{r^3}
Subtract. This difference is what a person falling with the two objects would see: the lower one pulls ahead.
:::

Near Earth's surface, 2GM/r³ is about 3 × 10⁻⁶ s⁻², so two objects 1 m apart vertically separate at about 3 μm/s², too little to notice in a lift but enough to raise the oceans' tides over thousands of kilometres. In general relativity this tidal quantity is, quite literally, a measure of spacetime curvature.

## Spacetime, not just space

The curvature that matters is of spacetime, not space alone. In flat spacetime the separation between two nearby events is given by the interval of special relativity:

::: equation The interval in flat spacetime
\mathrm{d}s^2 = -c^2\,\mathrm{d}t^2 + \mathrm{d}x^2 + \mathrm{d}y^2 + \mathrm{d}z^2
:::

Curved spacetime keeps the same idea, the sum of squared steps, but the coefficients in front of each term become functions of position, collectively called the **metric**. Near a slowly moving, weakly gravitating body like Earth, the most important change is in the time term:

::: equation The weak-field metric, time part
\mathrm{d}s^2 \approx -\left(1 + \frac{2\Phi}{c^2}\right)c^2\,\mathrm{d}t^2 + \dots
:::

Here Φ is the Newtonian potential, −GM/r. At Earth's surface 2Φ/c² is only about −1.4 × 10⁻⁹. That tiny distortion of time is enough to make a dropped ball fall at 9.8 m/s²: the path of a falling object is the one that maximizes the time ticked by its own clock, and the Schwarzschild lesson later works this out in full. For everyday objects, "gravity" is almost entirely a warping of time; the warping of space matters mainly for things moving near light speed, such as light itself.

## The field equations

The other half of the theory says what produces the curvature. Einstein's field equations equate a measure of curvature, the Einstein tensor G_μν, to the density and flow of energy and momentum, T_μν:

::: equation Einstein's field equations
G_{\mu\nu} = \frac{8\pi G}{c^4}\,T_{\mu\nu}
:::

John Wheeler's summary is the standard one: spacetime tells matter how to move; matter tells spacetime how to curve. Energy, momentum and pressure all count as sources, not just mass.

::: note
In general relativity gravity is not a force. Mass and energy curve spacetime, and freely falling objects follow the straightest paths available. What we feel as weight is the floor stopping us from following that path.
:::

::: history 1912–1915 · Zürich and Berlin
Einstein knew by 1912 that he needed a geometry in which distances vary from place to place, but not the mathematics to handle it. His old classmate Marcel Grossmann, by then a mathematics professor in Zürich, pointed him to the work of Bernhard Riemann (1854) and the tensor calculus of Ricci and Levi-Civita. Their first attempt, in 1913, had the wrong field equations, and Einstein convinced himself for two years that the correct ones were impossible. In November 1915, racing David Hilbert, he found the final form, and checked it the decisive way: it explained the 43 arcseconds per century by which Mercury's orbit swings around beyond what Newton predicts, a puzzle since the 1850s.
:::

The usual rubber-sheet picture, a ball denting a stretched membrane, shows only curved space and needs gravity underneath to make the marble roll. Treat it as a picture of the geometry, not an explanation.

::: model gravity-well
The surface shows how space curves around a compact mass. Move the probe inward and watch its clock slow: that warping of time is what makes things fall.
:::

::: quiz
? In general relativity, why does a dropped apple fall?
- A gravitational force pulls it down
+ It follows a straight path through curved spacetime, and the ground is in the way
- Space is pushing it toward the Earth
= Free fall is force-free motion along a geodesic. The force you feel standing still comes from the ground pushing you off your geodesic.
? Two falling objects, one 1 m above the other, separate at about 3 μm/s². What would the separation be for objects 2 m apart?
- The same
+ About twice as large
- About four times as large
= Δa ≈ 2GMh/r³ is proportional to h for small h.
? For a slowly moving object near Earth, which part of the metric produces most of what we call gravity?
+ The time term
- The radial space term
- The angular terms
= Slow objects move a tiny distance in space per unit time, so the curvature of time dominates. Space curvature matters mainly for fast-moving things like light.
:::
