An orbiting spacecraft is not beyond gravity. At the height of the International Space Station gravity is still almost 90% as strong as on the ground. The station stays up because it is falling all the time, and moving sideways so fast that the ground curves away beneath it as quickly as it falls.

## Newton's cannon

Newton illustrated this with a thought experiment. Put a cannon on a mountain so high that air does not matter, and fire horizontally. A slow ball curves down and lands nearby. A faster one lands farther away. Because Earth is round, the ground beneath a fast enough ball drops away as fast as the ball falls toward it: the ball never lands. It has gone into orbit. Fire faster still and the path stretches into an ellipse; faster again and the ball escapes for good.

::: model orbit
Raise the launch speed slowly. Watch the ball fall back, then circle, then trace wider ellipses, and finally escape.
:::

::: history 1680s · Cambridge
Newton drew the cannon on a mountain in the mid-1680s for a more popular account of his system, published after his death as *A Treatise of the System of the World* (1728). The picture makes vivid his central claim: the Moon and a thrown stone follow the same physics and differ only in speed. The first artificial satellite, Sputnik 1, went up on 4 October 1957, 270 years after the *Principia*, and circled Earth every 96 minutes.
:::

## Speed for a circular orbit

For a circle, gravity must provide exactly the centripetal force needed at that radius.

::: steps Circular orbit speed
$ \frac{GMm}{r^2} = \frac{mv^2}{r}
Gravitational force (left) equals mass times centripetal acceleration (right), for a satellite of mass m at distance r from the centre of a planet of mass M.
---
$ v^2 = \frac{GM}{r}
Cancel m from both sides and one factor of r. The satellite's mass has dropped out, as it does for anything in free fall.
---
$ v = \sqrt{\frac{GM}{r}}
Take the square root. Lower orbits need higher speeds.
:::

For Earth, GM = 3.99 × 10¹⁴ m³/s² (this product is known far more precisely than G or M separately). At 400 km up, r = 6.77 × 10⁶ m, giving v ≈ 7.7 km/s. Skimming the surface, if air allowed it, would take about 7.9 km/s.

## Period and Kepler's third law

The time for one circuit is the circumference divided by the speed, and this gives Kepler's third law with its constant now explained.

::: steps Orbital period
$ T = \frac{2\pi r}{v}
One circumference at speed v.
---
$ T^2 = \frac{4\pi^2 r^2}{v^2} = \frac{4\pi^2 r^2}{GM/r}
Square both sides and substitute v² = GM/r.
---
$ T^2 = \frac{4\pi^2}{GM}\,r^3
Kepler's T² ∝ r³, with the constant 4π²/GM set by the central mass alone. For ellipses the same formula holds with r replaced by the semi-major axis a.
:::

This makes orbits a scale for weighing. Measure any satellite's period and distance, and you have GM for whatever it circles: the Moon gives Earth's, the planets give the Sun's, moons give Jupiter's. Turning GM into a mass in kilograms needs G, which is the subject of the last lesson in this chapter.

Some useful cases:

| Orbit | Radius | Speed | Period |
|---|---|---|---|
| Low Earth orbit (ISS, 400 km up) | 6,770 km | 7.7 km/s | about 92 min |
| Geostationary | 42,160 km | 3.1 km/s | 23 h 56 min |
| The Moon | 384,400 km | 1.0 km/s | 27.3 days |

A **geostationary** satellite orbits once per sidereal day (Earth's rotation period relative to the stars, 23 h 56 min) above the equator, so it hangs over one spot; that is why satellite-TV dishes point at a fixed patch of sky. Solving T² = 4π²r³/GM for that T fixes the radius at about 42,160 km, or 35,800 km above the surface.

## Weightless but not gravity-free

Astronauts float because they and their spacecraft fall together with the same acceleration, so nothing presses them against the floor. You would feel the same in a lift whose cable had snapped. "Zero gravity" is a misnomer: at 400 km gravity is about 89% of its surface value.

::: note
An orbit is free fall with enough sideways speed that the ground curves away as fast as you fall. The required speed, √(GM/r), does not depend on the satellite's mass.
:::

## Changing orbits

Speed and orbit shape go together. A spacecraft in a circular orbit that fires its engine forward, speeding up, rises onto an ellipse whose far point is higher. Counterintuitively, it then moves more slowly there, because it has climbed against gravity. To catch a craft ahead of you in the same orbit, you briefly slow down: you drop to a lower, faster orbit, gain on the target, then rise again. At √2 times the circular speed, about 41% more, the ellipse opens out completely and the craft escapes; that is the next lesson.

::: quiz
? Why does an astronaut on the ISS float?
- Gravity is negligible at that height
+ The astronaut and the station fall together with the same acceleration
- The station's rotation cancels gravity
= Gravity at 400 km is still about 89% of the surface value; everything in the station is in free fall together.
? Satellite A orbits at four times the radius of satellite B around the same planet. A's speed is
- four times B's
- one quarter of B's
+ half of B's
= v = √(GM/r), so multiplying r by 4 divides v by √4 = 2.
? A heavier satellite placed in the same circular orbit as a lighter one has
- a shorter period
- a longer period
+ the same period
= The satellite's mass cancels in GMm/r² = mv²/r, so speed and period depend only on M and r.
:::
