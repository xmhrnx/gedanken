A car rounding a bend at a steady speed, a stone whirled round on a string, the Moon going round the Earth: each of these moves at a constant speed but in a constantly changing direction. The first law says that a body left alone moves in a straight line, so any curved path is evidence of a force. This lesson finds out which way that force must point and how large it has to be. The answer is short, a = v²/r, and it turns out to reach from a car skidding off a road to the reason astronauts float.

## How can there be acceleration without speeding up?

In everyday speech, accelerating means going faster. In physics it means something broader: **acceleration** is any change of velocity, and velocity is a vector, with a direction as well as a size. So a body accelerates if its speed changes, if its direction changes, or both.

In **uniform circular motion**, motion around a circle at constant speed, the size of the velocity, the speed v, stays fixed. But its direction does not. At every moment the velocity points along the circle, at right angles to the radius, the line from the centre to the body. As the body goes round, the radius swings round with it, and the velocity swings round too. A velocity whose direction keeps changing is changing, so the body is accelerating, even though a speedometer would show a constant reading.

Which way does this acceleration point? Look at the velocity at two nearby moments. The new velocity is the old one turned slightly toward the inside of the circle. The small change Δv that must be added to the old velocity to get the new one points inward, toward the centre. This holds at every point of the circle, so the acceleration always points at the centre. It is called **centripetal** acceleration, from the Latin for "centre-seeking".

## How large is it?

The direction is now settled. To find the size, take a body moving at speed v on a circle of radius r, and follow it for a short time Δt. The idea is to compare two triangles: one made of positions and one made of velocities.

::: steps Centripetal acceleration
$ \Delta\theta = \frac{v\,\Delta t}{r}
In the time Δt the body travels a distance vΔt along the circle. An angle in radians is defined as the length of arc it spans divided by the radius, so the radius line turns through this angle Δθ. This step turns a time into an angle.
---
$ |\Delta\vec{v}| \approx v\,\Delta\theta
The velocity always stays at right angles to the radius, so when the radius turns through Δθ, the velocity turns through exactly the same angle. Draw the old and new velocities from the same starting point. They form a thin triangle with two long sides of length v and a narrow angle Δθ between them, and the short third side is the change Δv. For a small angle, that short side is very nearly the long side times the angle, just as a short arc is the radius times the angle.
---
$ |\Delta\vec{v}| \approx \frac{v^2\,\Delta t}{r}
Now replace Δθ with the expression from the first step, v Δt / r. This gives the change in velocity in terms of quantities we know: the speed, the radius and the time elapsed. One factor of v came from the size of the velocity vector, the other from how fast it turns.
---
$ a = \lim_{\Delta t \to 0} \frac{|\Delta\vec{v}|}{\Delta t} = \frac{v^2}{r}
Acceleration is the rate of change of velocity, so divide the change by Δt. Then let Δt shrink toward zero, as in the rates lesson. The thin triangle becomes ever thinner, the approximations of the second step become exact, and the direction of Δv becomes exactly perpendicular to the velocity, pointing straight at the centre.
:::

The result says that both faster motion and tighter turning need more acceleration, but not equally. Halving the radius doubles the acceleration, while doubling the speed quadruples it. The square arises because speed enters twice: a faster body swings its velocity round the circle more quickly, and it also has a longer velocity vector to swing.

It is often more convenient to describe circular motion by its **period** T, the time for one complete revolution. In one period the body covers the circumference 2πr, so v = 2πr/T. Putting that into a = v²/r gives a second form:

::: equation Centripetal acceleration from the period
a = \frac{v^2}{r} = \frac{4\pi^2 r}{T^2}
:::

## Where does the force come from?

By the second law, a body of mass m accelerating at v²/r must have a net force on it of F = mv²/r, pointing toward the centre. This is often called the **centripetal force**, but the name can mislead. It is not a new kind of force, alongside gravity and friction. It is a job description: whatever circular motion is happening, some real force or combination of forces must be supplying an inward pull of exactly mv²/r. For a stone on a string, the job is done by the tension in the string. For the Moon, it is done by the Earth's gravity. For a car on a bend, it is done by friction between the tyres and the road.

The car shows what happens when the available force is not enough. Take a flat bend of radius 50 m. The only horizontal force on the car is static friction from the road, pointing toward the centre of the bend, and that can be at most μ_s mg, as the friction lesson showed. So the car can follow the bend only if mv²/r ≤ μ_s mg. The mass cancels, leaving v ≤ √(μ_s g r). With μ_s = 0.7 for dry asphalt, the limit is √(0.7 × 9.8 × 50), about 18.5 m/s or 67 km/h. Above that speed, friction cannot supply enough inward force. The car does not get flung outward; it simply turns less sharply than the road does and so drifts toward the outside, following a straighter line than the bend, until it leaves the road. On ice, with μ_s about 0.1, the same bend can be taken safely only below about 7 m/s, which is why roads that are fine in summer become dangerous in winter.

## Is there an outward force?

Riders on a fast fairground roundabout feel themselves thrown outward, and it is natural to think some outward force is acting on them. In an inertial frame, there is none.

Consider what each rider's body would do if left alone. By the first law it would travel in a straight line, along the tangent to the circle, which is the direction it is moving at that instant. The seat or handrail stops that from happening by pushing the rider inward, continually bending their path away from the straight line and into a circle. What the rider feels is that inward push, pressing on their back or side. Because the push comes from the outside of the circle, the sensation is of being pressed outward against the seat.

In the rotating frame of the roundabout itself, which is not an inertial frame, the rider is at rest, so it is often convenient to invent an outward **centrifugal force** to balance the inward push and make the first law appear to hold. That bookkeeping device is useful for engineers designing spinning machinery. But it has no physical source, and it disappears when you describe the same motion from the ground. The clearest test is to cut the string on a whirling stone. If an outward force were acting, the stone would fly straight outward along the radius. In fact it flies off along the tangent, in the direction it was already moving, exactly as the first law predicts once the inward pull stops.

## How does gravity hold an orbit?

Gravity can do the centripetal job too, and that idea is the key to orbits. Imagine a cannon on a very high mountain firing a ball horizontally. The ball falls toward the ground as it travels, and lands some distance away. Fire it faster and it lands further away. But the Earth is round, so the ground curves downward away from the ball as it travels. Fire it fast enough, and the ground curves away beneath the ball exactly as fast as the ball falls toward it. The ball keeps falling but never gets any closer to the ground: it circles the Earth forever.

How fast is fast enough? Ignore the air, and suppose the ball skims round just above the surface. Then gravity must supply exactly the centripetal acceleration for a circle of the Earth's radius.

::: steps Speed for a circular orbit near the surface
$ g = \frac{v^2}{R}
The only force on the ball is gravity, which gives it an acceleration g toward the Earth's centre. For a circular orbit of radius R, the Earth's radius, that acceleration must equal the centripetal acceleration v²/R. Any slower and gravity pulls the ball into a tighter curve than the Earth's surface, so it hits the ground.
---
$ v = \sqrt{gR} = \sqrt{9.8 \times 6.37 \times 10^6} \approx 7.9\,\text{km/s}
Multiply both sides by R and take the square root to find the speed. With the Earth's radius of 6.37 × 10⁶ m, this comes to about 7,900 metres per second, roughly 28,000 km/h, more than twenty times the speed of sound in air.
---
$ T = \frac{2\pi R}{v} \approx 5.1 \times 10^3\,\text{s} \approx 84\,\text{min}
The time for one orbit is the distance round the Earth divided by the speed. It comes to about 5,100 seconds, or 84 minutes. The International Space Station orbits about 400 km higher, where gravity is a little weaker and the circle a little larger, and takes about 92 minutes.
:::

That last point explains a common puzzle. Astronauts aboard the space station float freely, so it is tempting to say they are beyond the reach of gravity. They are not. At 400 km up, gravity is still about 90% as strong as at the surface, and it is precisely that gravity that keeps the station in orbit. The astronauts float because they and the station are falling together, all the time, toward the Earth, while their sideways speed carries them round it. Nothing needs to push them against the floor, for the same reason that the scales in the lift with the broken cable read zero.

::: model orbit
Raise the launch speed and watch the path change. Below about 7.9 km/s the ball falls back to the ground; at that speed it circles; faster still it traces wider and wider ellipses, and above about 11.2 km/s it escapes and never returns.
:::

::: history 1659–1673 · The Hague and Cambridge
Christiaan Huygens, the Dutch physicist who invented the pendulum clock, worked out in 1659 that the force needed to keep a body moving in a circle is proportional to v²/r. He wrote the derivation in a manuscript he titled *De vi centrifuga*, "On centrifugal force", and published the results, without proofs, in his book on pendulum clocks in 1673; the manuscript itself appeared only after his death. Huygens thought about the problem from the body's own point of view, in terms of the outward pull a weight on a string seems to exert, hence the name. Newton reached the same result independently around 1665, as a young man working at home while plague closed Cambridge. He then turned the viewpoint around. The force points inward, toward the centre, and something real must supply it. That shift made a new question possible: if something pulls the Moon toward the Earth and the planets toward the Sun, what is it? That question led to the *Principia*.
:::

::: note
Motion in a circle at constant speed is accelerated motion, because the direction of the velocity keeps changing. The acceleration points toward the centre and has size v²/r, so it needs a real inward force of mv²/r, supplied by tension, friction, gravity or a push.
:::

::: quiz
? A stone whirled in a horizontal circle on a string is released. Which way does it go?
- Straight outward, away from the centre
+ Along the tangent to the circle at the point of release
- It spirals outward
= Once the string's inward pull stops, no horizontal force acts, so by the first law the stone continues in a straight line in the direction it was moving at that moment, which is along the tangent.
? A car takes the same bend at twice the speed. The friction force needed is
- the same
- twice as large
+ four times as large
= The required inward force is mv²/r, so doubling v on the same bend multiplies it by four. If friction cannot supply that much, the car drifts to the outside of the bend.
? Why do astronauts on the space station float?
- There is no gravity at that height
- Gravity from the Moon cancels the Earth's
+ They and the station are in free fall together around the Earth
= Gravity at 400 km is about 90% of its surface value, and it supplies the centripetal force that keeps the station in orbit. Everything aboard falls at the same rate, so nothing presses against the floor.
:::
