The Schwarzschild metric misbehaves at two radii. At the Schwarzschild radius, r = r_s, one of its terms becomes infinite and another vanishes; at the centre, r = 0, everything breaks down. For decades most physicists, Einstein among them, assumed that both signalled the point where the solution stops describing anything real, and that no actual star could ever be small enough for the question to arise. The resolution, worked out slowly between the 1930s and the 1960s, is that the two radii are completely different kinds of trouble. One is an artefact of the coordinates; the other is real. This lesson shows how to tell them apart, then follows, step by step, what the metric says about clocks, light and falling objects near r_s, and arrives at the event horizon: a surface that can be crossed inward but never outward.

## Is every infinity in a formula a real infinity?

An infinity in a formula may belong to the thing being described, or only to the labels chosen to describe it. Maps of the Earth offer a familiar example. In latitude and longitude, every line of longitude meets at the North Pole, and at the pole itself the question "which way is east?" has no answer: the coordinate system breaks down. Yet nothing odd happens to the ice there. Someone standing at the pole feels nothing special, and a map centred on the pole shows it as an ordinary point. A breakdown that belongs only to the labels, and that a better choice of labels removes, is called a **coordinate singularity**.

So the question about r = r_s is whether it is like the North Pole. To decide, we need a quantity that does not depend on the choice of coordinates at all, one that every observer and every coordinate system would agree on. If such a quantity blows up somewhere, the trouble is real; if it stays finite, the infinity in the metric may be an artefact.

In relativity the standard test is a **curvature invariant**: a single number built from the curvature of spacetime in a way that comes out the same whatever coordinates are used. For the Schwarzschild geometry the simplest one that is not trivially zero, called the Kretschmann scalar after the physicist Erich Kretschmann, works out to be:

::: equation Curvature invariant of the Schwarzschild geometry
K &= \frac{12\,r_s^2}{r^6} \\ \text{at } r = r_s:\quad K &= \frac{12}{r_s^4} \\ \text{as } r \to 0:\quad K &\to \infty
:::

At the Schwarzschild radius, K is perfectly finite. Nothing about the geometry itself is infinite there. At r = 0, by contrast, K grows without limit. That is a **real singularity**, sometimes called a curvature singularity: a place where the theory itself predicts infinite curvature, and so infinite tidal stretching, and where general relativity can no longer say what happens.

The formula also says something surprising about size. Since K at the Schwarzschild radius goes as 1/r_s⁴, and r_s is proportional to the mass, the more massive a black hole, the gentler the curvature at its horizon. Tidal acceleration makes this concrete. Between two points a distance h apart along the radial direction, the tidal acceleration is 2GMh/r³, the same formula as in the curved-spacetime lesson, and at r = r_s, where 2GM = c²r_s, it becomes c²h/r_s². For a black hole of ten solar masses, with r_s about 30 km, that is about 10⁸ m/s² per metre: a falling person would be stretched apart long before reaching the horizon. For the black hole at the centre of our galaxy, with r_s about 1.3 × 10¹⁰ m, it is about 5 × 10⁻⁴ m/s² per metre, far too little to feel.

::: history 1924–1958 · from Brussels to New York
The suspicion that r = r_s might be a mere coordinate effect is old, but it was confirmed piecemeal and largely ignored. In 1924 Arthur Eddington wrote down a new set of coordinates for the Schwarzschild geometry in which the infinity at r_s does not appear, but he used them for another purpose and did not draw the conclusion. In 1933 the Belgian priest and physicist Georges Lemaître showed that the trouble at r_s disappears completely in coordinates attached to a freely falling observer, and stated plainly that the singularity there is not real. Few took notice. In 1958 David Finkelstein, then a young physicist at the Stevens Institute of Technology in New Jersey, rediscovered Eddington's coordinates and saw their meaning: in them, paths falling inward cross r_s smoothly, while no path crosses it outward. The radius is a one-way surface. That paper, with work that followed it, changed minds, among them that of John Wheeler, previously a firm sceptic. The coordinates are now called Eddington–Finkelstein coordinates, and the one-way surface the **event horizon**.
:::

The rest of this lesson builds up the same conclusion directly from the Schwarzschild metric, in four stages: what happens to a clock held near r_s, how long light from there takes to get out, what a falling object itself experiences, and what happens to the meaning of r inside.

## How fast does a clock near the horizon tick?

Start with the simplest question: how does a clock held at a fixed distance from the black hole, say on a rocket hovering in place, compare with a clock far away?

::: steps Gravitational time dilation for a hovering clock
$ \mathrm{d}s^2 = {}&-\left(1 - \frac{r_s}{r}\right)c^2\,\mathrm{d}t^2 \\ &+ \frac{\mathrm{d}r^2}{1 - r_s/r} \\ &+ r^2\left(\mathrm{d}\theta^2 + \sin^2\theta\,\mathrm{d}\phi^2\right)
Start from the Schwarzschild metric found in the previous lesson. Recall that the coordinate t is the time kept by a clock at rest very far away, where all the factors are 1. Every other clock's reading must be worked out from the metric.
---
$ \mathrm{d}s^2 = -\left(1 - \frac{r_s}{r}\right)c^2\,\mathrm{d}t^2
Now consider a clock held at rest at one fixed radius and direction. Between two of its ticks it does not move, so dr, dθ and dφ are all zero. The radial and angular terms vanish, and only the time term survives.
---
$ \mathrm{d}\tau^2 = -\frac{\mathrm{d}s^2}{c^2} = \left(1 - \frac{r_s}{r}\right)\mathrm{d}t^2
The time a clock actually records, its proper time τ, is defined from the interval by dτ² = −ds²/c². This is the same definition used throughout relativity; the minus sign makes the square of the time positive for anything moving slower than light.
---
$ \mathrm{d}\tau = \sqrt{1 - \frac{r_s}{r}}\;\mathrm{d}t
Take the square root of both sides. During one hour of the distant clock, the hovering clock records √(1 − r_s/r) hours, which is less than one. Far away, r_s/r is tiny and the factor is close to 1, which matches the weak-field rule of the previous chapter. Close in, it is very different.
:::

Put in some values. At r = 2r_s, twice the Schwarzschild radius, the factor is √0.5 ≈ 0.71, so the hovering clock runs at about 71 percent of the distant rate. At r = 1.01 r_s, just one percent outside, it is about 0.1. And as r approaches r_s, the factor approaches zero. Seen from far away, a clock held ever closer to the horizon appears to run ever slower, without limit.

Light is affected the same way, because light is a clock of sorts: its crests are its ticks. Light sent out from the hovering clock arrives far away with its frequency lowered by the same factor, a gravitational redshift that grows without limit as the source approaches r_s. The light also arrives dimmer, since fewer crests and fewer photons arrive per second.

## How long does light from near the horizon take to get out?

Here is a second, sharper way to see what is special about r_s. Send a flash of light straight outward from just above the horizon and ask how long it takes, in the distant observer's time t, to reach some far-off radius.

::: steps Light travel time out from near r_s
$ 0 = -\left(1 - \frac{r_s}{r}\right)c^2\,\mathrm{d}t^2 + \frac{\mathrm{d}r^2}{1 - r_s/r}
For light, the interval between any two nearby points on its path is zero: ds² = 0 is what it means to travel at the speed of light. For light moving straight outward, the angles do not change, so dθ and dφ are zero, and only the time and radial terms remain.
---
$ \frac{\mathrm{d}r}{\mathrm{d}t} = \pm\, c\left(1 - \frac{r_s}{r}\right)
Move the time term to the other side, multiply by (1 − r_s/r), and take the square root. The plus sign is outgoing light, the minus sign ingoing. Far away this is ±c, as it should be. Closer in, the light seems to move more slowly in these coordinates, and at r = r_s it seems to stop. Nobody near the light measures it slowing; it is the distant observer's labels r and t that are being strained.
---
$ c\,\mathrm{d}t = \frac{\mathrm{d}r}{1 - r_s/r} = \left(1 + \frac{r_s}{r - r_s}\right)\mathrm{d}r
Turn this around to find the time taken to cover each small step outward. Then rewrite the fraction: 1/(1 − r_s/r) equals r/(r − r_s), which is the same as 1 plus r_s/(r − r_s). Splitting it this way separates an ordinary part from the part that misbehaves near r_s.
---
$ c\,t = (r_2 - r_1) + r_s\ln\frac{r_2 - r_s}{r_1 - r_s}
Add up the small steps from the starting radius r₁ to the final radius r₂. The first part integrates to the plain distance r₂ − r₁, which is what light would take in flat spacetime. The second part integrates to a natural logarithm. This is the total delay, in the distant observer's time, for light leaving r₁.
---
$ r_1 \to r_s \;\Rightarrow\; t \to \infty
Now let the starting point approach the horizon. The quantity r₁ − r_s in the denominator approaches zero, so the logarithm, and with it the travel time, grows without limit. Light emitted just above r_s takes longer and longer to escape, and light emitted exactly at r_s never gets out at all, in any finite time.
:::

The logarithm makes the delay grow in a particular way. Each time the starting point's height above the horizon is halved, the delay increases by the same fixed amount, r_s ln 2 / c. For a ten-solar-mass black hole that amount is about 70 microseconds. So a source falling toward the horizon does not vanish abruptly; its light arrives later and later, redder and redder, and fades away within a fraction of a millisecond, though in principle a last trickle of ever-weaker signals keeps arriving forever.

## What does the falling object itself experience?

Everything so far is described from far away, using the distant observer's time t. The infinite delay might seem to say that nothing ever reaches r_s. To see whether that is true, we must follow the falling object with its own clock, its proper time τ. Take the simplest case: a probe that starts at rest very far away and falls straight in.

::: steps Free fall in the probe's own time
$ \left(1 - \frac{r_s}{r}\right)\frac{\mathrm{d}t}{\mathrm{d}\tau} = 1
Because the Schwarzschild metric is the same at every moment t, a freely falling object carries a conserved quantity, the relativistic counterpart of its energy, just as in Newtonian physics a steady gravitational field conserves energy. For this metric that quantity is (1 − r_s/r) dt/dτ. Far away, where the probe starts at rest, the factor is 1 and its clock keeps the same time as t, so the quantity equals 1 and stays 1 all the way down.
---
$ c^2\,\mathrm{d}\tau^2 = \left(1 - \frac{r_s}{r}\right)c^2\,\mathrm{d}t^2 - \frac{\mathrm{d}r^2}{1 - r_s/r}
Now write the metric for the probe's path. It falls straight in, so the angles do not change, and c²dτ² = −ds². This connects three things: the probe's own time, the distant time, and its change of radius.
---
$ c^2 = \frac{c^2}{1 - r_s/r} - \frac{1}{1 - r_s/r}\left(\frac{\mathrm{d}r}{\mathrm{d}\tau}\right)^2
Divide every term by dτ², and use the first step to replace dt/dτ with 1/(1 − r_s/r). The distant time t has now been eliminated; only the probe's radius and its own clock remain.
---
$ \left(\frac{\mathrm{d}r}{\mathrm{d}\tau}\right)^2 = c^2\,\frac{r_s}{r} = \frac{2GM}{r}
Multiply through by (1 − r_s/r) and collect terms: c² − c²r_s/r = c² − (dr/dτ)², which gives this. In the last form it is exactly Newton's escape-velocity formula, read backwards: an object falling from rest at infinity has the escape speed at every radius. Here, though, it is exact, with the probe's own clock.
---
$ \frac{\mathrm{d}r}{\mathrm{d}\tau} = -c\,\sqrt{\frac{r_s}{r}}
Take the square root, with a minus sign because r is decreasing. Look at r = r_s: the right side is simply −c, a perfectly ordinary finite number. Nothing in the probe's own description singles out the Schwarzschild radius. (r is a coordinate, so dr/dτ is not quite a speed anyone measures, but it is finite, and that is the point.)
---
$ \tau = \frac{2}{3}\,\frac{r_s}{c}\left[\left(\frac{r_1}{r_s}\right)^{3/2} - \left(\frac{r_2}{r_s}\right)^{3/2}\right]
To find how long the fall takes on the probe's own clock, add up dτ = −dr / (c√(r_s/r)) from a starting radius r₁ to a finishing radius r₂. The integral of √r is (2/3)r^(3/2), which gives this. Every term is finite for any radii, including r₂ = r_s and r₂ = 0.
---
$ \tau_{r_s \to 0} = \frac{2}{3}\,\frac{r_s}{c}
Finally set r₁ = r_s and r₂ = 0. The probe crosses the horizon after a finite time on its own clock and then reaches the centre in a further time of two-thirds of r_s/c. The infinite time in the distant observer's description belongs to the coordinate t, not to the probe.
:::

Put in numbers. For a ten-solar-mass black hole, r_s ≈ 30 km, so the trip from horizon to centre takes about 66 microseconds of the probe's time. For the black hole at the centre of our galaxy, with r_s ≈ 1.3 × 10¹⁰ m, it takes about half a minute. A probe that fires its rockets cannot stretch this indefinitely: the longest possible time from horizon to centre, achieved by starting from rest just at the horizon, is πGM/c³, about 0.15 milliseconds and about a minute respectively.

So there are two descriptions of the same fall. The distant observer sees the probe's signals arrive further and further apart and further into the red; the probe seems to slow as it nears r_s and fade, never quite seen to cross. The probe itself crosses in a finite time and notices nothing at that moment: no wall, no flash, no change in the local laws of physics. Both accounts are correct. They differ because in relativity there is no single, shared "now" across space, and the distant observer's time t is exactly the label that breaks down at r_s.

## Why does the horizon only work one way?

A useful picture of the limits of motion is the **light cone**. From any event, the paths light can take outward in every direction form a cone in a diagram of space against time. Anything with mass must travel inside that cone, more slowly than light; that is what "nothing travels faster than light" means geometrically.

The result dr/dt = ±c(1 − r_s/r) says how these cones look near a black hole. Far away the outgoing and ingoing edges move at ±c, and the cone is the ordinary symmetric shape. Closer in, in the distant observer's coordinates, both edges seem to slow and the cone narrows. That picture is distorted by the bad coordinates. In coordinates that follow infalling light, such as Eddington and Finkelstein's, the picture is clean: the cones tilt progressively inward as r decreases. At r = r_s the outward edge of the cone points straight up the time axis, so even light aimed directly outward only hovers at r_s. Inside, the entire cone, both edges, points toward smaller r.

::: model light-cones
Light cones along a line into a black hole. Watch the outer edge of each cone: at the horizon it becomes vertical, and inside, every possible future direction leads inward.
:::

## Why does r become time inside the horizon?

The tipping of the cones can also be read straight from the signs in the metric, and the argument shows why nothing inside can stop falling.

::: steps No standing still inside the horizon
$ r < r_s \;\Rightarrow\; 1 - \frac{r_s}{r} < 0
Inside the horizon r is smaller than r_s, so r_s/r is greater than 1, and the factor 1 − r_s/r that appears in the metric is negative. This single sign change is what reorganises everything.
---
$ \mathrm{d}s^2 = {}&+\left(\frac{r_s}{r} - 1\right)c^2\,\mathrm{d}t^2 \\ &- \frac{\mathrm{d}r^2}{r_s/r - 1} \\ &+ r^2\,\mathrm{d}\Omega^2
Rewrite the metric with the signs made explicit, using dΩ² for the angular part. The time term, normally negative, is now positive; the radial term, normally positive, is now negative. In the interval, the term with the minus sign is the one that behaves like time. So inside, r plays the role of time, and t the role of a direction in space.
---
$ \frac{\mathrm{d}r^2}{r_s/r - 1} &> \left(\frac{r_s}{r} - 1\right)c^2\,\mathrm{d}t^2 + r^2\,\mathrm{d}\Omega^2 \\ &\geq 0
Any object with mass moves slower than light, which means the interval along its path is always negative, ds² < 0. Rearranged, the negative radial term must outweigh the other two terms combined. Those two are each zero or positive, so the radial term can never be zero: dr can never be zero. An object inside cannot hold its radius fixed, any more than an object outside can hold its time fixed.
---
$ \frac{\mathrm{d}r}{\mathrm{d}\tau} < 0
The inequality says r must keep changing, but not yet in which direction. That is settled by continuity: an object that crossed the horizon moving inward, as every infalling object does, cannot reverse without passing through dr = 0, which is forbidden. So r decreases, always, for everything inside, light included, until it reaches r = 0.
:::

That is the content of the horizon. Outside, you can choose to move toward the centre or away from it, as you can choose to walk north or south. Inside, decreasing r is no more avoidable than next Tuesday. Firing rockets changes only how you reach r = 0, not whether, and no amount of thrust can stretch the remaining time beyond the limit of πGM/c³ found above. The singularity is not a place in front of you that might be dodged; it is a moment in your future.

::: note
The event horizon is not a material surface, and nothing local marks it. It is the boundary of the region from which no path into the future, not even light's, reaches the outside. The real singularity at r = 0 is where the curvature becomes infinite and general relativity stops working.
:::

::: model gravity-well
The well deepens toward the centre, and the probe's clock slows as it moves in. At the horizon the clock rate seen from far away falls to zero, and inside, no path leads back out.
:::

::: quiz
? How can you tell that r = r_s is only a coordinate singularity?
- Light slows down there
+ A curvature invariant, which every coordinate system agrees on, stays finite there
- The metric has no terms at r = r_s
= The Kretschmann scalar K = 12r_s²/r⁶ is finite at r = r_s and becomes infinite only at r = 0. A falling probe's own description, dr/dτ = −c√(r_s/r), is also finite at r_s.
? A clock hovers at r = 4r_s/3. Compared with a distant clock, at what rate does it run?
- 3/4 the rate
+ 1/2 the rate
- 1/4 the rate
= The rate factor is √(1 − r_s/r) = √(1 − 3/4) = √(1/4) = 1/2. The hovering clock records half an hour for every hour of the distant clock.
? Why can't an object inside the horizon avoid reaching r = 0?
- Its rockets are not powerful enough
- The tidal forces pull it in
+ Inside, decreasing r is the future direction of time
= For r < r_s the metric's signs swap, so dr can never be zero along any path slower than light, and continuity fixes its sign as decreasing. Every path into the future leads to smaller r.
:::
