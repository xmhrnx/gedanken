A car rounding a bend at steady speed, a stone whirled on a string, the Moon going round the Earth: each moves at constant speed but changing direction. By the first law, a change of direction needs a force. This lesson finds out how big, and which way it points.

## Acceleration without speeding up

Velocity is a vector. In uniform circular motion its length, the speed v, stays fixed, but its direction keeps turning. A changing velocity is an acceleration, even if the speedometer never moves.

Which way does it point? The velocity is always along the circle, perpendicular to the radius. As the body goes round, the velocity vector swings round with it, and the small change Δv added in each instant points inward, toward the centre. This is **centripetal** ("centre-seeking") acceleration.

## How big: a = v²/r

Take a body moving at speed v on a circle of radius r, and follow it for a short time Δt.

::: steps Centripetal acceleration
$ \Delta\theta = \frac{v\,\Delta t}{r}
In time Δt the body travels an arc of length vΔt. An arc divided by its radius is the angle it spans, in radians, so the radius line turns through Δθ.
---
$ |\Delta\vec{v}| \approx v\,\Delta\theta
The velocity stays perpendicular to the radius, so it turns through the same angle Δθ. Drawn tail to tail, the old and new velocities form a thin triangle with two sides of length v; the short third side, Δv, is about v times the angle, just as an arc is radius times angle.
---
$ |\Delta\vec{v}| \approx \frac{v^2\,\Delta t}{r}
Substitute Δθ from the first step.
---
$ a = \lim_{\Delta t \to 0} \frac{|\Delta\vec{v}|}{\Delta t} = \frac{v^2}{r}
Acceleration is the rate of change of velocity. Divide by Δt and let it shrink, as in the rates lesson; the approximation becomes exact, and Δv points exactly at the centre.
:::

Faster motion and tighter circles both need more acceleration. Doubling the speed on the same circle needs four times as much.

It is often handier to use the **period** T, the time for one revolution. Since v = 2πr/T:

::: equation Centripetal acceleration from the period
a = \frac{v^2}{r} = \frac{4\pi^2 r}{T^2}
:::

## The force comes from something real

By the second law, a body of mass m moving in a circle needs a net inward force F = mv²/r. This is not a new kind of force. It is a job description, filled by whatever force is available: the tension in the string, the friction between tyres and road, gravity for the Moon.

Take a car on a flat bend of radius 50 m. The only inward force is static friction from the road, at most μ_s mg. Setting mv²/r ≤ μ_s mg gives v ≤ √(μ_s g r). With μ_s = 0.7 that is √(0.7 × 9.8 × 50) ≈ 18.5 m/s, about 67 km/h. Go faster and friction cannot supply enough force; the car continues in a straighter line than the road, and leaves it. On ice, with μ_s about 0.1, the limit drops to 7 m/s.

## Centrifugal force

Riders on a fast roundabout feel thrown outward. No outward force acts. Each rider's inertia tries to carry them in a straight line, tangent to the circle, and the seat or handrail must push them inward to keep them turning. What they feel is that inward push. In the rotating frame of the roundabout, which is not inertial, it is convenient to invent an outward *centrifugal force* to make the first law appear to hold; in an inertial frame it does not exist. When a string breaks, the stone flies off along the tangent, not outward along the radius.

## Falling around the Earth

Gravity can provide the centripetal force too. A cannonball fired horizontally falls toward the ground, but the ground curves away beneath it. Fire it fast enough and the surface curves away exactly as fast as it falls: it circles the Earth forever. Ignoring air, that happens when the acceleration of gravity equals v²/r at the Earth's radius R:

::: steps Speed for a circular orbit near the surface
$ g = \frac{v^2}{R}
Gravity supplies the whole centripetal acceleration.
---
$ v = \sqrt{gR} = \sqrt{9.8 \times 6.37 \times 10^6} \approx 7.9\,\text{km/s}
Solve for v.
---
$ T = \frac{2\pi R}{v} \approx 5.1 \times 10^3\,\text{s} \approx 84\,\text{min}
One lap of the Earth at that speed. The International Space Station, about 400 km up where gravity is slightly weaker, takes about 92 minutes.
:::

Astronauts aboard are not beyond gravity; at that height it is still about 90% as strong as at the surface. They float because they and their station are falling together, as the scales in a lift with a broken cable read zero.

::: model orbit
Raise the launch speed. Below about 7.9 km/s the ball falls back; at that speed it circles; faster still it traces ellipses, and eventually escapes.
:::

::: history 1659–1673 · The Hague
Christiaan Huygens derived the rule that the force needed is proportional to v²/r in 1659, in a manuscript he called *De vi centrifuga*, and published the results without proof in 1673. He thought in terms of the outward pull felt by the body, hence the name centrifugal. Newton reached the same result independently around 1665, and turned the viewpoint round: the force points inward, and something must supply it. That shift let him ask what supplies it for the Moon and the planets.
:::

::: note
Circular motion at constant speed is accelerated motion, directed toward the centre with a = v²/r. It needs a real inward force, supplied by tension, friction, gravity or a push.
:::

::: quiz
? A stone whirled in a horizontal circle on a string is released. Which way does it go?
- Straight outward, away from the centre
+ Along the tangent to the circle at the point of release
- It spirals outward
= Once the string's inward pull stops, no horizontal force acts, so the stone continues in a straight line in the direction it was moving: along the tangent.
? A car takes the same bend at twice the speed. The friction force needed is
- the same
- twice as large
+ four times as large
= F = mv²/r, so doubling v quadruples the required force.
? Why do astronauts on the space station float?
- There is no gravity at that height
- Gravity from the Moon cancels the Earth's
+ They and the station are in free fall together around the Earth
= Gravity at 400 km is about 90% of its surface value. It supplies the centripetal force for the orbit, and everything aboard falls at the same rate.
:::
