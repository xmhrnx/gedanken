The equivalence principle says that gravity can be made to vanish inside a small, freely falling room. It cannot be made to vanish everywhere at once: over larger distances, falling objects drift together or apart, and no choice of frame removes that. Einstein's answer to what those leftover tidal effects are was the boldest idea in his theory. Gravity, he proposed, is not a force at all. It is the geometry of spacetime, and the tidal effects are that geometry's curvature. This lesson explains what that means, why it fits the facts, and what equation decides how spacetime curves.

## How can paths converge without any force?

Start with a question about geometry rather than physics. On a flat sheet of paper, two straight lines that start out parallel stay parallel forever. That is one of the defining facts of flat, or Euclidean, geometry. On the surface of a sphere it fails.

Picture two travellers who stand on the equator 100 km apart and both set off due north. Each walks along the straightest path the surface allows, which on a sphere is a line of longitude. Neither ever turns left or right. Yet the distance between them steadily shrinks, and at the North Pole they meet. If they did not know they were on a sphere, they might conclude that some force was pulling them together. There is no force. The convergence comes entirely from the shape of the surface they walk on. (The comparison is limited: the sphere is curved space, a surface, whereas Einstein's curvature involves time too, as we will see. But the logic of straight paths converging carries over.)

The straightest possible path in any geometry, flat or curved, is called a **geodesic**. On a flat sheet geodesics are ordinary straight lines; on a sphere they are great circles, such as lines of longitude.

Einstein's proposal is that a freely falling object follows a geodesic through spacetime. This explains at once why a falling object feels no force: it is not being pushed off its path, it is simply going as straight as the geometry allows, which is exactly what the equivalence principle says a falling observer experiences. And it explains tidal effects: two neighbouring falling objects drift together or apart not because anything acts on them, but because the geodesics they follow converge or diverge, like the travellers' lines of longitude.

## How big is the part of gravity that cannot be removed?

If tidal effects are the true signature of curvature, it is worth knowing how large they are. Take two small objects falling straight toward a planet or star of mass M. One is at distance r from the centre; the other is directly above it, a small height h higher. Someone falling along with them sees only the difference between their accelerations, because the shared part is exactly what falling cancels.

::: steps Tidal acceleration between two falling objects
$ a(r) = \frac{GM}{r^2}
Begin with Newton's gravitational acceleration at distance r from the centre. Newton's law is accurate here because the field is weak. On its own, this acceleration can be cancelled completely by falling along with it, so it is not yet the part we are after.
---
$ a(r + h) = \frac{GM}{(r + h)^2}
The upper object is at distance r + h, so its acceleration is given by the same law with r replaced by r + h. Because it is farther away, it is pulled slightly less. The tidal effect is the small difference between this and the line above.
---
$ \frac{1}{(r + h)^2} = \frac{1}{r^2}\left(1 + \frac{h}{r}\right)^{-2}
To compare the two, factor r² out of the bracket in the denominator. This rewrites the awkward expression as the simple 1/r² times a correction that depends only on the ratio h/r, which is a small number when the objects are close together.
---
$ \left(1 + \frac{h}{r}\right)^{-2} \approx 1 - \frac{2h}{r}
For any small number x, (1 + x)⁻² is very close to 1 − 2x; the terms left out involve x², which is tiny when x is small. With x = h/r, a metre compared with the 6,400 km radius of the Earth, the approximation is excellent. It turns the correction into something we can subtract easily.
---
$ \Delta a = a(r) - a(r + h) \approx \frac{GM}{r^2}\cdot\frac{2h}{r} = \frac{2GMh}{r^3}
Subtract the upper acceleration from the lower one. The GM/r² terms cancel, leaving only the correction. This difference is what a person falling with the two objects actually sees: the lower object pulls ahead of the upper one, at a rate proportional to their separation h.
:::

Put in numbers for the Earth. With r equal to the Earth's radius, 2GM/r³ comes to about 3 × 10⁻⁶ per second squared. So two objects one metre apart, one above the other, separate at about 3 micrometres per second squared. That is far too little to notice in a falling lift. But the same formula applied across the Earth's whole diameter, with the Moon or the Sun as the mass M, gives the stretching that raises the ocean tides twice a day. In general relativity this tidal quantity is not just a consequence of curvature; it is, quite literally, a measurement of it.

## Why spacetime, and not just space?

The sphere analogy might suggest that gravity is curved space. It is not mainly that. What curves is spacetime, the four-dimensional combination of space and time that special relativity introduced, and for most everyday purposes the important part is the time.

To see what this means we need a way of describing geometry with numbers. In special relativity, the separation between two nearby events, say a firecracker going off at one place and time and another at a nearby place and slightly later time, is measured by the **interval** ds:

::: equation The interval in flat spacetime
\mathrm{d}s^2 = -c^2\,\mathrm{d}t^2 + \mathrm{d}x^2 + \mathrm{d}y^2 + \mathrm{d}z^2
:::

It works like Pythagoras' theorem, a sum of squared steps, except that the time step enters with a minus sign. Every observer, however they are moving, agrees on ds², even though they disagree about dt and dx separately. For an object that moves from one event to the other, −ds²/c² is the square of the time that passes on the object's own clock, called its **proper time**.

Curved spacetime keeps the same idea, a sum of squared steps, but allows the numbers in front of each term to vary from place to place. The full set of these position-dependent coefficients is called the **metric**: it is the rule that says how much distance and how much time a given step in the coordinates really corresponds to. Near a slowly moving, weakly gravitating body such as the Earth, the most important change is to the coefficient of the time term:

::: equation The weak-field metric, time part
\mathrm{d}s^2 \approx -\left(1 + \frac{2\Phi}{c^2}\right)c^2\,\mathrm{d}t^2 + \dots
:::

Here Φ is the familiar Newtonian gravitational potential, −GM/r, which is negative and gets more negative closer to the mass. So the coefficient in front of the time term is slightly less than 1 near a mass, and a clock there ticks off slightly less proper time than a distant one: time runs slower deeper down.

How large is the change? At the Earth's surface 2Φ/c² is only about −1.4 × 10⁻⁹, roughly one part in a billion. It seems absurd that so tiny a distortion could have anything to do with something as obvious as a falling apple. But it does, and the reason is the rule for how free objects move. In relativity, a freely falling object follows the path between two events along which its own clock records the greatest possible time. A ball thrown upward and caught again rises into the region where clocks run slightly faster, and the path that gains the most time by doing so, without wasting time by moving too fast, is exactly the familiar parabola with acceleration 9.8 m/s². The Schwarzschild lesson in the next chapter carries out this calculation in full and shows that the time term reproduces Newton's gravity exactly.

The surprising conclusion is that for slow, everyday objects, what we call gravity is almost entirely a warping of time. The warping of space is also there, by a similar amount, but a slow object moves through very little space in each second, so it barely notices. Space curvature matters mainly for things moving near the speed of light, such as light itself, which is why it shows up in the bending of starlight.

## What tells spacetime how to curve?

So far we have half a theory: given a curved spacetime, objects move along geodesics. The other half must say what produces the curvature. Einstein's **field equations** answer that. On the left is the Einstein tensor G_μν, a combination of derivatives of the metric that measures curvature. On the right is the stress–energy tensor T_μν, which lists the density and flow of energy and momentum at each point. A **tensor** here is simply a table of numbers, with one entry for each pair of directions; the labels μ and ν each run over time and the three directions of space.

::: equation Einstein's field equations
G_{\mu\nu} = \frac{8\pi G}{c^4}\,T_{\mu\nu}
:::

The physicist John Wheeler gave the standard summary: spacetime tells matter how to move, and matter tells spacetime how to curve. One feature has no counterpart in Newton's theory: mass is not the only source. Energy of any kind, momentum, and even pressure all curve spacetime. The factor 8πG/c⁴ is extremely small, about 2 × 10⁻⁴³ in SI units, which is why it takes the mass of a planet or star to produce curvature anyone can notice.

::: note
In general relativity gravity is not a force. Mass and energy curve spacetime, and freely falling objects follow the straightest paths that curved spacetime allows. What we feel as weight is the floor pushing us off that path.
:::

::: history 1912–1915 · Zürich and Berlin
By 1912 Einstein knew that he needed a geometry in which distances and time intervals could vary from place to place, but he did not know the mathematics for handling one. He turned to his old university friend Marcel Grossmann, by then a professor of mathematics in Zürich, and asked for help. Grossmann pointed him to the geometry of curved spaces developed by Bernhard Riemann in 1854 and to the tensor calculus built on it by Gregorio Ricci-Curbastro and Tullio Levi-Civita. Their joint paper of 1913 got the geometry right but the field equations wrong, and Einstein then talked himself into believing, for more than two years, that equations of the correct general kind could not work. In the autumn of 1915, aware that the mathematician David Hilbert was now working on the same problem, he returned to them in a burst of effort. In November he found the final form and checked it in the most convincing way available. For decades astronomers had known that the long axis of Mercury's elliptical orbit swings around the Sun 43 arcseconds per century faster than Newton's theory allows. Einstein's equations gave exactly that amount, with nothing adjusted to fit. He later wrote that the result left him beside himself with excitement for days.
:::

You will often see curved spacetime drawn as a heavy ball denting a stretched rubber sheet, with a marble rolling around the dent. The picture is useful for the shape of curved space, but as an explanation it is circular: the marble rolls into the dent only because real gravity pulls it down. It also shows only space, leaving out the curvature of time that does most of the work for slow objects. Treat it as a picture of the geometry, not of the cause.

::: model gravity-well
The surface shows how space curves around a compact mass. Move the probe inward and watch its clock slow compared with a distant clock: that warping of time is what makes things fall.
:::

::: quiz
? According to general relativity, why does a dropped apple fall?
- A gravitational force pulls it down
+ It follows the straightest path through curved spacetime, and the ground is in the way
- Space pushes it toward the Earth
= Free fall is motion along a geodesic with no force acting. The force you feel when standing still is the ground pushing you off your geodesic.
? Two falling objects, one 1 m above the other, separate at about 3 μm/s². What would the rate be for objects 2 m apart?
- The same
+ About twice as large
- About four times as large
= The tidal acceleration is Δa ≈ 2GMh/r³, which is proportional to the separation h when h is small. Doubling h doubles the rate.
? For a slowly moving object near the Earth, which part of the metric produces most of what we call gravity?
+ The time term
- The radial space term
- The angular terms
= A slow object moves only a tiny distance through space in each second, so the curvature of time dominates its motion. The curvature of space matters mainly for fast-moving things such as light.
:::
