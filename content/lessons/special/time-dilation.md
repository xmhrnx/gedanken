The second postulate says that every observer measures the same speed for light, however they are moving. The first consequence to follow from it is also the most famous: observers moving relative to each other cannot agree about how much time passes. A clock that moves relative to you runs slow, as you measure it. That sounds like a claim about clocks being damaged by motion, but it is something deeper, a statement about time itself. The clearest way to see why it must be true uses a clock built from a beam of light and nothing more advanced than Pythagoras' theorem. This lesson builds that clock, derives how much it slows, shows why every other kind of clock must slow by the same amount, and looks at the evidence.

## Why use a clock made of light?

Any clock is a process that repeats regularly: a swinging pendulum, a vibrating quartz crystal, an oscillating atom. To compare clocks in different frames, we want the simplest possible process, one whose behaviour we can predict in any frame from the postulates alone. Light is ideal, because the second postulate tells us exactly how fast it moves for everyone.

So take two mirrors facing each other, a distance L apart, and let a short pulse of light bounce between them. Each round trip, from the bottom mirror to the top and back, counts as one tick. For someone at rest beside the clock, the light travels straight up and straight down, a total distance 2L, at speed c, so a tick lasts 2L/c.

Now let the whole clock move sideways, at a steady speed v, past an observer standing on the ground. For the clock's owner, riding along with it, nothing has changed: the first postulate says that physics in her frame works exactly as it would at rest, so she sees the light go straight up and down, and each tick lasts 2L/c as before. But the ground observer sees something different. While the pulse is on its way up, the top mirror moves sideways, so to hit it the light must travel along a slanting path. Coming down, the bottom mirror has moved on too, and the light follows a second slant. Each tick, seen from the ground, is a zigzag, and a zigzag is longer than a straight up-and-down path.

Here the second postulate does its work. In Newton's physics, the light would simply move faster from the ground's point of view, having picked up the clock's sideways speed, and the longer path would take the same time. But the second postulate says the ground observer measures the light moving at exactly c, no faster. A longer path at the same speed takes longer. So the ground observer measures each tick of the moving clock as lasting longer than 2L/c.

::: model light-clock
Increase the clock's speed. The light's zigzag path gets longer, so each tick takes longer for the ground observer, and γ shows by how much.
:::

## How much longer is each tick?

Call Δt₀ the time for one tick measured by the clock's owner, and Δt the time for the same tick as measured from the ground.

::: steps Time dilation from the light clock
$ \Delta t_0 = \frac{2L}{c}
In the clock's own frame, the light travels a distance L up and L down at speed c. Distance divided by speed gives the duration of one tick as the clock's owner measures it.
---
$ d = c\,\frac{\Delta t}{2}, \qquad x = v\,\frac{\Delta t}{2}
From the ground, consider just the upward half of the tick, which by symmetry lasts Δt/2. In that time the light moves along its slanting path a distance d at speed c, while the clock, and with it the top mirror, moves sideways a distance x at speed v. Both distances are written in terms of the unknown Δt we want to find.
---
$ \left(c\,\frac{\Delta t}{2}\right)^2 = L^2 + \left(v\,\frac{\Delta t}{2}\right)^2
The slanting path d is the hypotenuse of a right triangle whose other two sides are the vertical distance L between the mirrors and the sideways distance x. Pythagoras' theorem gives d² = L² + x². We are using the same L as the clock's owner, which requires that lengths at right angles to the motion are the same for both observers; the reason this is so comes just after the derivation.
---
$ \frac{\Delta t^2}{4}\left(c^2 - v^2\right) = L^2
Subtract the term with v from both sides, so that both terms containing Δt are on the left, and take out their common factor Δt²/4. The equation now says that one combination of the unknown equals a known quantity.
---
$ \Delta t = \frac{2L}{\sqrt{c^2 - v^2}}
Multiply both sides by 4, divide by c² − v², and take the positive square root, since a duration is positive. This is the ground observer's tick time. Notice that it makes sense only if v is less than c; the clock could not move at c or faster.
---
$ \Delta t = \frac{2L}{c}\,\frac{1}{\sqrt{1 - v^2/c^2}}
Take a factor c out of the square root, using √(c² − v²) = c√(1 − v²/c²). This splits the answer into two parts: one that does not depend on the motion, and one that depends only on the ratio v/c.
---
$ \Delta t = \gamma\,\Delta t_0, \qquad \gamma = \frac{1}{\sqrt{1 - v^2/c^2}}
The first factor, 2L/c, is exactly Δt₀ from the first step. The second is called the **Lorentz factor**, written γ (the Greek letter gamma). Since v is less than c, the quantity under the square root is less than 1, so γ is always at least 1. The ground observer measures every tick as lasting γ times as long as the clock's owner does.
:::

Why is the distance between the mirrors the same for both observers? Imagine two identical metre sticks, one at rest and one moving, both held at right angles to the direction of motion, and passing so close that their tips nearly touch. Each has a paintbrush at its tip. Suppose moving objects shrank sideways. Then from the resting stick's point of view, the moving stick is shorter, and its brush leaves a mark below the resting stick's tip. But by the first postulate, the moving stick's owner can equally regard herself as at rest and the other stick as moving, so she says the other stick is shorter, and her brush marks above its tip. The paint mark is a physical fact that both can inspect afterward, and it cannot be both above and below. So neither stick shrinks. Lengths perpendicular to the motion are the same for everyone.

## Does this apply to real clocks?

The light clock is a convenience, chosen because it makes the calculation easy. But the conclusion applies to every clock. Suppose the clock's owner straps her wristwatch to the light clock. If the wristwatch kept different time from the light clock while moving, she could compare the two and see them drift apart, and that drift would tell her she was moving, without looking outside. That is exactly what the first postulate forbids. So every clock moving with the light clock, mechanical, atomic or biological, must slow by the same factor γ. The process of ageing in a living body is one more clock.

So the slowing is not a mechanical effect on light clocks; it is a property of time. **Moving clocks run slow**: any process, measured from a frame in which it is moving, takes longer than it does in its own rest frame. The time Δt₀ measured by a clock in its own rest frame, the frame in which it is not moving, is called its **proper time**. Every other observer measures a longer interval for the same process.

How large is the effect? The value of γ depends only on v/c:

| v | γ |
|---|---|
| 250 m/s (airliner) | 1.000 000 000 000 35 |
| 0.1c | 1.005 |
| 0.6c | 1.25 |
| 0.9c | 2.29 |
| 0.99c | 7.09 |

At everyday speeds, γ exceeds 1 by about v²/(2c²), a fraction so small that for an airliner it amounts to roughly a nanosecond lost in an hour of flight. That is why the effect went unnoticed for all of human history. Near c, by contrast, γ grows rapidly and without limit.

## Is there evidence?

Nature provides a test. When cosmic rays, fast particles from space, strike the upper atmosphere, they create **muons**, particles like heavy electrons, about 207 times as massive, which are unstable. A muon at rest decays, on average, after 2.2 microseconds. Even moving at nearly the speed of light, a muon would on average cover only about 660 m in that time (3 × 10⁸ m/s × 2.2 × 10⁻⁶ s). Muons are made about 15 km up, so without time dilation, only a tiny fraction should survive to reach the ground. Yet detectors at sea level count them in large numbers.

Time dilation explains it. A muon moving at 0.995c has γ ≈ 10. Seen from the ground, its internal clock, the clock that decides when it decays, runs ten times slow, so on average it lasts about 22 microseconds and covers about 6.6 km. A good fraction survives the trip. In 1963 David Frisch and James Smith measured this carefully. They counted muons of a known speed range on the summit of Mount Washington in New Hampshire, 1910 m up, about 563 per hour, and then counted again at sea level. Without time dilation, only a few dozen of those 563 should have survived the descent. They counted about 400 per hour, in good agreement with relativity.

The effect is also built into everyday technology. Satellites of the Global Positioning System orbit at about 3.9 km/s, and their atomic clocks, seen from the ground, lose about 7 microseconds a day because of their motion. They actually gain more than that overall, because clocks higher in Earth's gravity run faster, an effect explained by general relativity. The system corrects for both. Light travels about 2 km in 7 microseconds, so without the correction, computed positions would drift by kilometres within a day.

::: note
A clock moving at speed v ticks slowly, as measured by an observer it moves past, by the factor γ = 1/√(1 − v²/c²). The light clock shows why: its light follows a longer path at the same speed c, so each tick takes longer, and the first postulate forces every other kind of clock to keep in step with it.
:::

::: history 1897–1905 · Leiden, Cambridge and Bern
The factor γ was not new in 1905. Hendrik Lorentz and Joseph Larmor had both written it down in the late 1890s, as part of the mathematics needed to explain why experiments could not detect motion through the ether. Lorentz introduced a quantity he called "local time" for moving systems, but regarded it as a calculating device, not the time a real clock would show; true time, in his view, was the time of the ether's frame. Einstein's 1905 paper said the slowing was real. A clock carried away from another and brought back, he wrote, would be found behind the one that stayed at home. He also predicted that a clock at the equator, carried around by Earth's rotation at about 460 m/s, would run slightly slow compared with an identical clock at a pole. That particular prediction turned out to be wrong, for a reason he could not have seen in 1905: Earth's rotation makes it bulge at the equator, which places an equatorial clock farther from Earth's centre, where gravity makes clocks run faster. The two effects cancel exactly, so all clocks at sea level keep the same time.
:::

## If each sees the other slow, who is right?

Time dilation is symmetric. The clock's owner, by the first postulate, can regard herself as at rest and the ground as moving, so she measures the ground's clocks running slow by the same factor γ. That looks like a contradiction: how can each clock run slower than the other?

It is not, because the comparison each observer makes involves clocks at different places, and comparing distant clocks depends on how they are synchronised, which the lesson on simultaneity shows observers disagree about. To compare two clocks without any such ambiguity, they must be brought back together at one place, side by side. For that to happen, at least one of them must turn around, and turning around means accelerating, which means leaving its inertial frame. The situation is then no longer symmetric. The clock that turned around is the one that ends up behind. This is the so-called twin paradox: a twin who travels to a distant star and back returns younger than the twin who stayed home. It is not a paradox, only a surprise, and it has been checked with atomic clocks flown around the world on airliners.

::: quiz
? In the light clock, why does the ground observer measure a longer tick?
- The light slows down when the clock moves
+ The light travels a longer, slanting path at the same speed c
- The mirrors move farther apart
= The second postulate fixes the light's speed at c for the ground observer too, so it cannot speed up to make up for the longer path. A longer path at the same speed takes more time.
? A spaceship travels at 0.6c relative to Earth. One hour passes on the ship's clock. How much time passes for observers on Earth?
- 48 minutes
- 1 hour
+ 75 minutes
= γ = 1/√(1 − 0.36) = 1/0.8 = 1.25, so one hour of proper time on the ship is measured from Earth as 1.25 hours, or 75 minutes.
? Why must a mechanical clock on the spaceship slow by the same factor as the light clock?
+ Otherwise comparing the two clocks would reveal the ship's motion, which the first postulate forbids
- Because mechanical clocks contain light
- Because the ship's engines disturb them
= All physics on the ship must work exactly as it would at rest, so no two clocks there can drift apart. Every process on the ship, including ageing, keeps the same time as the light clock.
:::
