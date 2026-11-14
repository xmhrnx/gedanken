Two events that happen at the same time for one observer can happen at different times for another moving relative to the first. This is the least intuitive consequence of the postulates, and the one that unlocks the rest: once "at the same time" depends on the observer, time dilation and length contraction stop looking contradictory.

## What "at the same time" means

For events at the same place, simultaneity is unproblematic: two things happen together or they do not. For events far apart, you need clocks at both places, and the clocks must be synchronised. Einstein's 1905 recipe: send a flash of light from a point exactly midway between two clocks. Light travels at c in both directions, so it reaches them together, and both clocks are set to zero when it arrives. Equivalently, two distant events are simultaneous if light from them meets at the midpoint.

The recipe uses the second postulate. Applied by observers in relative motion, it gives different answers.

## Einstein's train

A train moves along an embankment at speed v. Lightning strikes both ends of the train at the instant they line up with two marks on the embankment. An observer on the embankment, standing midway between the marks, receives the two flashes together. Since she was midway and both flashes travelled at c, she concludes the strikes were simultaneous.

A passenger sits at the midpoint of the train. During the time the light is in flight, she moves toward the front flash and away from the rear one, so the front flash reaches her first. That much both observers agree on: it is a single event, two flashes arriving at one person, in a definite order. But the passenger is entitled to regard herself as at rest. In her frame she sat at the midpoint of the train, both flashes travelled at c, and the front one arrived first. So the front strike happened first. Neither observer is wrong; simultaneity is relative.

## How large is the disagreement?

Send a light flash from the centre of a carriage of proper length L₀. In the train's frame it reaches both ends at the same time. Now work in the ground frame, where the carriage moves forward at v.

::: steps Timing the two arrivals from the ground
$ L = \frac{L_0}{\gamma}
From the ground, the moving carriage is contracted.
---
$ t_{\text{back}} = \frac{L/2}{c + v}
The rear wall moves forward to meet the backward-going light, so they close at c + v from an initial separation of L/2.
---
$ t_{\text{front}} = \frac{L/2}{c - v}
The front wall runs ahead of the forward-going light, so the light gains on it at only c − v.
---
$ t_{\text{front}} - t_{\text{back}} = \frac{L}{2}\cdot\frac{(c + v) - (c - v)}{c^2 - v^2} = \frac{Lv}{c^2 - v^2}
Subtract over a common denominator.
---
$ \frac{Lv}{c^2(1 - v^2/c^2)} = \frac{L\gamma^2 v}{c^2} = \frac{\gamma L_0 v}{c^2}
Factor c² out of the denominator, recognise 1/(1 − v²/c²) as γ², and substitute L = L₀/γ.
:::

::: equation Two events simultaneous on the train, separated by L₀, as timed from the ground
\Delta t_{\text{ground}} = \frac{\gamma v L_0}{c^2}
:::

From the ground, the event at the rear of the train happens first. The disagreement grows with the distance between the events and with the speed. For a 100 m carriage at 100 m/s it is about 10⁻¹³ s, far too small to notice, which is why everyday life gives no hint of it. For a carriage moving at 0.6c it is 0.25 microseconds per 100 m.

## Order can flip, but not cause and effect

For two events far enough apart that not even light could travel from one to the other in the time between them, observers in different frames can disagree on the order: one says A came first, another says B did. But for such events, neither can affect the other. When one event can cause another, light (or something slower) can link them, and then every observer agrees which came first. The postulates leave causality intact, provided nothing travels faster than light.

This also dissolves the apparent paradox of reciprocal length contraction. To measure a moving rod, you mark its two ends at the same time. The rod's own frame disagrees that your two markings were simultaneous. In the rod's view you marked its front end first and its rear end later; in between, the ground carrying your first mark slid backward along the rod, so your two marks end up closer together than the rod is long.

::: note
Whether two separated events happen "at the same time" depends on the observer's motion. Events simultaneous in one frame are separated by γvL₀/c² in a frame moving at v along the line joining them.
:::

::: history 1905 · Bern
Einstein spent years trying to reconcile Maxwell's equations with the principle of relativity. In May 1905, after a long discussion with his friend Michele Besso, he saw the answer: time cannot be absolutely defined, and there is a link between time and signal speed. About five weeks later he submitted his paper. Henri Poincaré had argued in 1898 that simultaneity at a distance is a convention, and that clocks could be synchronised by light signals, but he kept the ether as the frame in which the "true" time was defined. Einstein abolished the privileged frame. He presented the train-and-lightning example in a popular book written in 1916.
:::

::: quiz
? In the train example, the passenger sees the front flash first. Why does she conclude the front strike happened first?
+ She sat at the midpoint, both flashes travelled at c in her frame, so the one that arrived first left first
- The front flash was brighter
- Light from the front travelled faster because the train moves toward it
= In her own frame she is at rest; equal distances at equal speed mean the earlier arrival was the earlier event.
? Two events are simultaneous on the ground and 300 km apart. Will a fast-moving observer always find them simultaneous too?
- Yes, simultaneity is absolute
+ Not if the observer moves along the line joining them; then one event comes first
- Only if the observer moves at less than half the speed of light
= Δt = γvL/c² is nonzero for any motion along the line joining the events.
? Can a moving observer see an effect happen before its cause?
- Yes, if they move fast enough
+ No; when one event can influence another, every observer agrees on their order
- Only if they move faster than light, which is allowed
= Order can differ only for events too far apart for any signal at speed c or less to link them.
:::
