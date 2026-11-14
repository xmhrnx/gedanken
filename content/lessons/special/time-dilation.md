If every observer measures the same speed for light, then observers moving relative to each other cannot agree about time. The cleanest way to see this uses a clock made of light, and needs nothing more than Pythagoras' theorem.

## The light clock

Take two mirrors facing each other, a distance L apart, with a pulse of light bouncing between them. Each round trip is one tick. For someone at rest beside the clock, the light goes straight up and straight down, a distance 2L, at speed c.

Now let the clock move sideways at speed v past an observer on the ground. For the clock's owner, riding along, nothing has changed: the light still goes straight up and down. But the ground observer sees the mirrors move while the light is in flight, so the light must travel along a diagonal to hit the top mirror, and another diagonal to come back down. The diagonal path is longer. By the second postulate, the ground observer also measures the light moving at c. A longer path at the same speed takes longer.

::: model light-clock
Increase the clock's speed. The light's path gets longer and each tick takes longer for the ground observer; γ shows by how much.
:::

## Deriving the factor

Call Δt₀ the time for one tick measured by the clock's owner, and Δt the time for the same tick measured from the ground.

::: steps Time dilation from the light clock
$ \Delta t_0 = \frac{2L}{c}
In the clock's own frame, light travels L up and L down at speed c.
---
$ d = c\,\frac{\Delta t}{2}, \qquad x = v\,\frac{\Delta t}{2}
From the ground, take the upward half of the tick, which lasts Δt/2. In that time the light travels a distance d along its diagonal, while the clock moves a distance x sideways.
---
$ \left(c\,\frac{\Delta t}{2}\right)^2 = L^2 + \left(v\,\frac{\Delta t}{2}\right)^2
The diagonal d is the hypotenuse of a right triangle whose other sides are the mirror separation L and the sideways distance x, so d² = L² + x². The separation is the same L for both observers, because it is perpendicular to the motion (see below).
---
$ \frac{\Delta t^2}{4}\left(c^2 - v^2\right) = L^2
Move the v term to the left and factor out Δt²/4.
---
$ \Delta t = \frac{2L}{\sqrt{c^2 - v^2}}
Multiply by 4, divide by c² − v², take the square root.
---
$ \Delta t = \frac{2L}{c}\,\frac{1}{\sqrt{1 - v^2/c^2}}
Take c out of the square root: √(c² − v²) = c√(1 − v²/c²).
---
$ \Delta t = \gamma\,\Delta t_0, \qquad \gamma = \frac{1}{\sqrt{1 - v^2/c^2}}
The first factor is Δt₀. The second is called the **Lorentz factor**, γ. Since v < c, γ is always at least 1: the ground observer measures each tick as longer.
:::

Why is L the same for both? Imagine two identical metre sticks passing each other at right angles to their length, each with a paintbrush at its tip. If moving sticks shrank sideways, the moving one's brush would mark the other below its tip; but from the other stick's point of view it is the other one moving, and the mark should be above. The marks cannot be both, so neither shrinks. Lengths perpendicular to the motion are unchanged.

## Not just light clocks

The light clock is a convenience. Every clock moving with it, mechanical, atomic or biological, must slow by the same factor. If a wristwatch strapped to the light clock kept different time, the owner could compare the two and detect their own motion, violating the first postulate. **Moving clocks run slow**: time itself, as measured by any process, passes more slowly in a frame moving relative to you. The time Δt₀ measured by a clock in its own rest frame is called **proper time**.

How big is γ?

| v | γ |
|---|---|
| 250 m/s (airliner) | 1.000 000 000 000 35 |
| 0.1c | 1.005 |
| 0.6c | 1.25 |
| 0.9c | 2.29 |
| 0.99c | 7.09 |

At everyday speeds the effect is about v²/(2c²), a fraction of a nanosecond per hour on an airliner. Near c it grows without limit.

## Evidence: muons

Cosmic rays striking the upper atmosphere make **muons**, heavy cousins of the electron that decay with a mean lifetime of 2.2 microseconds. Even at nearly c, a muon would cover only about 660 m in 2.2 μs on average, so almost none should reach the ground from 15 km up. Yet they arrive in large numbers. At v = 0.995c, γ ≈ 10: from the ground, the muon's internal clock runs ten times slow, and it lasts long enough to cover about 6.6 km on average. In 1963 David Frisch and James Smith counted muons on Mount Washington, 1910 m up, and again at sea level. Without time dilation, only a few dozen of every 563 should have survived the descent; they counted about 400.

The effect is also built into everyday technology. GPS satellites move at about 3.9 km/s, and their clocks lose about 7 microseconds a day from motion alone (gravity makes them gain more, as general relativity explains). Uncorrected, positions would drift by kilometres within a day.

::: note
A moving clock ticks slowly by the factor γ = 1/√(1 − v²/c²). The light clock shows why: same light speed, longer path, longer tick.
:::

::: history 1905 · Bern
Lorentz and Joseph Larmor had already written down the factor γ in the late 1890s as part of the mathematics relating frames in the ether, with a "local time" Lorentz regarded as a calculating device. Einstein's 1905 paper said the slowing is real: a clock carried away and brought back would be behind one that stayed. He also predicted that a clock at the equator, carried around by Earth's rotation, would run slow compared with one at a pole. That turned out to be wrong, for a reason he could not see in 1905: Earth's equatorial bulge puts the equator clock higher in the gravitational field, and the two effects cancel exactly on the sea-level surface.
:::

Time dilation is symmetric: each of two observers in relative motion sees the other's clock run slow. That is not a contradiction, because comparing the clocks directly requires bringing them together, and at least one must turn around, changing frames. The one that turns around ends up younger: the "twin paradox", which is not a paradox.

::: quiz
? In the light clock, why does the ground observer measure a longer tick?
- The light slows down when the clock moves
+ The light travels a longer, diagonal path at the same speed c
- The mirrors move further apart
= The second postulate fixes the speed at c; the path is longer, so the time is longer.
? A spaceship travels at 0.6c. An hour passes on the ship's clock. How much time passes for Earth observers?
- 48 minutes
- 1 hour
+ 75 minutes
= γ = 1/√(1 − 0.36) = 1/0.8 = 1.25, so 1 hour of proper time is 1.25 hours from Earth.
? Why must a mechanical clock on the spaceship slow by the same factor as the light clock?
+ Otherwise comparing the two would reveal the ship's motion, violating the first postulate
- Mechanical clocks contain light
- Because the ship's engines affect them
= All physics in the ship's frame must be the same as at rest, so every clock there keeps the same time.
:::
