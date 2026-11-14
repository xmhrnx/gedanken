Before asking *why* things move (that is the job of forces, in the next chapter), we need a precise way to say *how* they move. Three quantities do the job: position, velocity and acceleration. Each is the rate of change of the one before it.

To keep things simple, this lesson deals with motion along a straight line. Everything carries over to two and three dimensions by treating each direction separately, as the lesson on vectors showed.

## Position

Pick a reference point, the *origin*, and a positive direction. Position x is the signed distance from the origin: x = 5 m means 5 m in the positive direction; x = −3 m means 3 m the other way. The choice of origin is arbitrary, but once made it must be kept.

*Displacement* is the change in position, Δx = x₂ − x₁. It is not the same as distance travelled. Walk 10 m forward and 10 m back: the distance is 20 m, the displacement zero.

## Velocity

*Velocity* is the rate of change of position:

::: equation Velocity
v = \frac{\mathrm{d}x}{\mathrm{d}t}
:::

It is the slope of the position–time graph. Velocity has a sign: positive means moving in the positive direction, negative means moving the other way. *Speed* is the magnitude of velocity, always positive or zero. A car reversing at 3 m/s has velocity −3 m/s and speed 3 m/s.

The average velocity over an interval is Δx/Δt, the slope of the line between two points on the graph. The instantaneous velocity is the limit as the interval shrinks: the slope of the tangent.

## Acceleration

*Acceleration* is the rate of change of velocity:

::: equation Acceleration
a = \frac{\mathrm{d}v}{\mathrm{d}t} = rac{\mathrm{d}^2 x}{\mathrm{d}t^2}
:::

The second form says acceleration is the derivative of the derivative of position; d²x/dt² is read "the second derivative of x with respect to t". Its units are (m/s)/s = m/s², metres per second, per second. An acceleration of 2 m/s² means the velocity grows by 2 m/s every second.

Two points cause most confusion:

- **Acceleration is not speed.** An object can be fast with zero acceleration (a car cruising at a steady 100 km/h) or momentarily at rest with large acceleration (a ball at the top of its flight).
- **Negative acceleration does not always mean slowing down.** It means velocity is becoming more negative. If v is positive, that slows the object. If v is already negative, it speeds it up in the negative direction. An object slows down when v and a have opposite signs.

## Going back down the chain

Derivatives take you from position to velocity to acceleration. Integrals take you back: the area under the velocity graph is the displacement, and the area under the acceleration graph is the change in velocity.

::: steps From position to acceleration: an example
$ x = 3t^2 - 12t \quad (\text{metres, } t \text{ in seconds})
Suppose a cart's position is given by this formula.
---
$ v = \frac{\mathrm{d}x}{\mathrm{d}t} = 6t - 12
Differentiate term by term with the power rule: 3t² gives 6t, and −12t gives −12.
---
$ a = \frac{\mathrm{d}v}{\mathrm{d}t} = 6\,\text{m/s}^2
Differentiate again. The −12 is constant, so it contributes nothing. The acceleration is constant.
---
$ v = 0 \text{ at } t = 2\,\text{s}, \text{ where } x = 12 - 24 = -12\,\text{m}
Set v = 0 and solve: 6t = 12. Before t = 2 s the cart moves backward (v < 0) while a > 0, so it is slowing down. At t = 2 s it stops at x = −12 m, then speeds up forward.
:::

::: note
Position, velocity, acceleration: each is the slope of the graph of the one before, and each is the area under the graph of the one after.
:::

::: history 14th century · Merton College, Oxford
Aristotle had no clear notion of speed at an instant, let alone acceleration. In the 1330s a group of Oxford scholars, the Merton Calculators, defined "instantaneous velocity" and "uniformly accelerated motion", and proved that a uniformly accelerating body covers the same distance as one moving steadily at its average speed. Nicole Oresme in Paris drew it as a graph, with the distance as an area. Galileo used the result three centuries later.
:::

::: quiz
? A ball is thrown straight up. At the top of its flight, what are its velocity and acceleration?
- Both zero
+ Velocity zero, acceleration not zero
- Velocity not zero, acceleration zero
= The ball stops for an instant, but its velocity is still changing (from up to down), so the acceleration is not zero.
? An object has v = −4 m/s and a = −2 m/s². What is happening?
- It is slowing down
+ It is speeding up in the negative direction
- It is at rest
= v and a have the same sign, so the speed increases.
? Position is x = 2t³ m. What is the acceleration at t = 1 s?
- 6 m/s²
+ 12 m/s²
- 2 m/s²
= v = 6t², a = 12t, so a = 12 m/s² at t = 1 s.
:::
