An orbiting spacecraft is not beyond the reach of gravity. At the height of the International Space Station, about 400 km up, gravity is still almost 90% as strong as it is on the ground. So why doesn't the station fall? The answer is that it does fall, all the time, but it is also moving sideways so fast that the ground curves away beneath it as quickly as it drops. This lesson makes that idea precise: it works out how fast an orbiting body must move, how long it takes to go round, and why astronauts float even though gravity has hardly weakened.

## Newton's cannon

Newton explained orbits with a thought experiment. Imagine a mountain so tall that its peak rises above the air, so that nothing slows a moving object except gravity. Put a cannon on the summit and fire a ball horizontally.

A slow ball curves downward and lands near the foot of the mountain. Fire a faster one and it travels farther before it lands, because it covers more ground sideways in the time it takes to fall. So far this is ordinary projectile motion. But Earth is round, and that changes everything at high speed. As the ball travels sideways, the ground beneath it curves downward, away from it. For a fast enough ball, the ground drops away exactly as quickly as the ball falls toward it, and the ball never gets any closer to the surface. It keeps falling and keeps missing. It has gone into orbit.

Fire faster still and the ball rises away from the surface on the far side of the planet before falling back: its path is an ellipse, with the mountain at the nearest point. Faster again, and the ellipse stretches so far that the ball never returns. The same gravity acts on the ball in every case. Only its speed decides whether it lands, circles, loops round on an ellipse or leaves for good.

::: model orbit
Raise the launch speed slowly. Watch the ball fall back to the ground, then circle the planet, then trace wider and wider ellipses, and finally escape.
:::

::: history 1680s–1957 · Cambridge and Baikonur
Newton drew the cannon on a mountain in the 1680s for a more popular account of his ideas, a draft he set aside in favour of the mathematical *Principia*. It was published after his death, in 1728, as *A Treatise of the System of the World*. The point of the picture was to make his central claim feel natural: the Moon and a thrown stone obey the same law and differ only in how fast they move sideways. For 270 years the cannon stayed imaginary, because no device could reach the required speed. Then, on 4 October 1957, a Soviet rocket launched from what is now the Baikonur Cosmodrome in Kazakhstan put Sputnik 1, a polished metal sphere about 58 cm across, into orbit. It went round Earth every 96 minutes, transmitting a radio beep that amateurs around the world could pick up.
:::

## How fast must a circular orbit be?

For the simplest orbit, a circle, we can work out the speed exactly. The idea is that gravity has to supply precisely the inward pull needed to bend the satellite's path into a circle of that radius. Too little speed and gravity wins, pulling the satellite into a lower path; too much and the satellite swings out wider. There is just one speed at which the two match.

::: steps Circular orbit speed
$ \frac{GMm}{r^2} = \frac{mv^2}{r}
On the left is the gravitational force on a satellite of mass m at distance r from the centre of a planet of mass M. On the right is the force needed to keep any mass m moving at speed v on a circle of radius r: mass times the centripetal acceleration v²/r. For a circular orbit gravity is the only force acting, so it must be exactly the force required, and we can set the two equal.
---
$ v^2 = \frac{GM}{r}
Divide both sides by m, and multiply both sides by r. The satellite's mass m disappears entirely, because it appears on both sides: once as the thing gravity pulls on, once as the inertia that resists being turned. This is the same cancellation that makes all bodies fall at the same rate, and it means a feather and a battleship in the same orbit would move at the same speed.
---
$ v = \sqrt{\frac{GM}{r}}
Take the square root of both sides to get the speed itself. Because r is in the denominator, a lower orbit needs a higher speed: closer to the planet gravity is stronger, so the satellite must move faster to keep missing the ground.
:::

To use this for Earth we need the product GM. It turns out that this product is known far more precisely than either G or Earth's mass separately, for a reason the next section makes clear: it can be read directly from orbits. Its value for Earth is GM = 3.99 × 10¹⁴ m³/s². For the space station, at about 400 km above the surface, the distance from Earth's centre is r = 6.77 × 10⁶ m, and the formula gives v ≈ 7.7 km/s, around 28,000 km/h. A satellite skimming the surface itself, if there were no air to stop it, would need about 7.9 km/s. These speeds are why reaching orbit is so hard: most of a rocket's effort goes not into climbing but into going sideways fast enough.

## How long does one orbit take?

The time for one circuit is the length of the path divided by the speed. Working this out recovers Kepler's third law, and this time with the constant explained.

::: steps Orbital period
$ T = \frac{2\pi r}{v}
The satellite travels once around a circle of circumference 2πr at a steady speed v. Since speed is constant on a circular orbit, the time taken is simply distance divided by speed.
---
$ T^2 = \frac{4\pi^2 r^2}{v^2} = \frac{4\pi^2 r^2}{GM/r}
Square both sides, so that the speed appears as v², which we already know from the previous derivation. Then replace v² with GM/r. Squaring is what makes this replacement possible without having to carry a square root around.
---
$ T^2 = \frac{4\pi^2}{GM}\,r^3
Dividing by GM/r is the same as multiplying by r/GM, which turns r² into r³. This is Kepler's third law, T² ∝ r³, and now the constant of proportionality is no longer a mystery: it is 4π²/GM, set entirely by the mass of the central body. Newton also showed that the same formula holds for elliptical orbits if r is replaced by the semi-major axis a.
:::

This turns orbits into a set of scales. Measure the period and distance of anything orbiting a body, and the formula gives that body's GM. The Moon's orbit weighs Earth; the planets' orbits weigh the Sun; the moons of Jupiter, which Galileo first saw, weigh Jupiter. That is why GM is so well known: periods and distances can be measured very precisely. What orbits cannot do is separate G from M. To convert GM into a mass in kilograms you need G on its own, and that took a delicate laboratory experiment, the subject of the last lesson in this chapter.

Three orbits around Earth show the range:

| Orbit | Radius | Speed | Period |
|---|---|---|---|
| Low Earth orbit (ISS, 400 km up) | 6,770 km | 7.7 km/s | about 92 min |
| Geostationary | 42,160 km | 3.1 km/s | 23 h 56 min |
| The Moon | 384,400 km | 1.0 km/s | 27.3 days |

The middle row is a special case worth explaining. Earth turns once relative to the stars every 23 hours 56 minutes, a period called the **sidereal day**. (The 24-hour solar day is slightly longer because Earth also moves along its orbit around the Sun.) A satellite placed on a circular orbit above the equator, moving in the same direction as Earth turns and with exactly that period, goes round at the same rate as the ground beneath it, so it seems to hang motionless over one spot. Such an orbit is called **geostationary**. Putting T = 23 h 56 min into T² = 4π²r³/GM and solving for r gives about 42,160 km from Earth's centre, or about 35,800 km above the surface. There is exactly one such orbit, which is why it is crowded with communications and weather satellites, and why a satellite-television dish can be bolted in place pointing at a fixed patch of sky.

## Weightless, but not without gravity

Astronauts on the space station float, and this is often described as "zero gravity". That description is wrong, as we have seen: at 400 km, gravity is about 89% of its strength at the surface. The astronauts float because they and the station are falling together. Both are in orbit, both accelerate toward Earth at exactly the same rate, so nothing presses the astronauts against the floor. What we normally feel as weight is not gravity itself but the floor pushing up on us to stop us falling. Remove the push, by letting the floor fall too, and the feeling of weight disappears even though gravity is still there.

You would feel the same thing, briefly, in a lift whose cable had snapped, or at the top of a fast drop on a roller coaster. The space station is in that state continuously, for months at a time. Physicists call it **free fall**: motion under gravity alone, with no other force acting.

::: note
An orbit is free fall with enough sideways speed that the ground curves away as fast as you fall toward it. The speed needed for a circular orbit, √(GM/r), depends on the planet's mass and the orbit's radius but not on the satellite's own mass.
:::

## Changing orbits

Because speed and orbit go together, a spacecraft changes its orbit by changing its speed, and the results are often the opposite of what intuition suggests.

Suppose a spacecraft in a circular orbit fires its engine briefly forward along its direction of motion, speeding up. It now has more speed than a circle at that height requires, so gravity can no longer bend its path tightly enough to hold it on the circle. It swings outward onto an ellipse, whose nearest point is where the engine fired and whose farthest point is on the opposite side of the planet, higher up. And when it reaches that high point it is moving more slowly than before, because it has climbed against gravity and traded speed for height. Speeding up has made it slower, at least on the far side.

This is why manoeuvring in orbit feels backwards. To catch up with another craft ahead of you in the same orbit, you do not speed up, which would lift you into a higher and slower path and make you fall behind. Instead you briefly slow down. You drop to a lower orbit, which is faster and shorter, gain on the target, and then speed up to rise back to its height at the right moment.

Keep increasing the speed and the ellipse grows longer and longer. At √2 times the circular speed, about 41% more, it stops closing at all: the spacecraft leaves and never returns. That speed is called the escape velocity, and it is the subject of the next lesson.

::: quiz
? Why does an astronaut on the International Space Station float?
- Gravity is negligible at that height
+ The astronaut and the station fall together with the same acceleration
- The station's rotation cancels gravity
= Gravity at 400 km is still about 89% of its value at the surface. Everything in the station is in free fall together, so nothing pushes the astronaut against the floor.
? Satellite A orbits at four times the radius of satellite B around the same planet. A's speed is
- four times B's
- one quarter of B's
+ half of B's
= Circular orbit speed is √(GM/r). Multiplying r by 4 divides the speed by √4, which is 2.
? A heavier satellite placed in the same circular orbit as a lighter one has
- a shorter period
- a longer period
+ the same period
= The satellite's mass cancels from GMm/r² = mv²/r, so its speed and period depend only on the planet's mass and the orbit's radius.
:::
