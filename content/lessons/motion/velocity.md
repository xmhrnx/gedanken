Before asking why things move, which is the job of forces in the next chapter, we need a precise way to say how they move. That turns out to need only three quantities: position, velocity and acceleration. Each one is the rate of change of the one before it, so the calculus of the previous chapter ties them together into a single chain. This lesson sets up that chain carefully, because almost every confusion in mechanics, from "a ball at the top of its flight has no acceleration" to "negative acceleration means slowing down", comes from a loose grip on these three ideas.

To keep things simple, the lesson deals with motion along a straight line, such as a train on a straight track or a ball thrown straight up. Nothing is lost by this. As the lesson on vectors showed, motion in two or three dimensions can be split into independent motions along each direction, and each of those is a straight-line problem of exactly this kind.

## Where is it? Position and displacement

To describe where something is, you first need something to measure from. So pick a reference point, called the **origin**, and choose one direction along the line to call positive. An object's **position**, written x, is then its distance from the origin with a sign attached: x = 5 m means 5 metres in the positive direction, and x = −3 m means 3 metres the other way. Where you put the origin, and which way you call positive, is entirely up to you, and a clever choice can make a problem easier. But once chosen, the choice must be kept for the whole problem, or the signs stop meaning anything.

**Displacement** is the change in position, Δx = x₂ − x₁, the final position minus the starting one. It is tempting to think of it as the distance travelled, but the two are different, and the difference matters. Walk 10 m forward and then 10 m back: you have travelled 20 m, but your displacement is zero, because you ended where you began. Displacement keeps track only of the net change, with its direction, and that is the quantity the equations of motion are built on.

## How fast, and which way? Velocity

**Velocity** is the rate at which position changes with time:

::: equation Velocity
v = \frac{\mathrm{d}x}{\mathrm{d}t}
:::

Following the previous chapter, that means velocity is the slope of the graph of position against time. A steep graph means position changes quickly, so the velocity is large; a flat graph means the object is not moving.

Because position has a sign, so does velocity. A positive velocity means the object is moving in the positive direction, and a negative one means it is moving the other way. **Speed** is the size of the velocity with the sign dropped, so it is never negative. A car reversing at 3 m/s, with forward chosen as positive, has a velocity of −3 m/s and a speed of 3 m/s. The two words are often used interchangeably in everyday speech, but in physics the sign carries information that the speed throws away.

There are two versions of velocity, just as there were two versions of a rate. The **average velocity** over a stretch of time is Δx/Δt, the displacement divided by the time taken, which is the slope of the straight line joining two points on the graph. The **instantaneous velocity** is what that average approaches as the stretch of time shrinks to nothing, which is the slope of the tangent at a single point. When physicists say "velocity" without qualification, they mean the instantaneous one.

## How fast is the velocity changing? Acceleration

**Acceleration** is the rate at which velocity changes with time:

::: equation Acceleration
a = \frac{\mathrm{d}v}{\mathrm{d}t} = \frac{\mathrm{d}^2 x}{\mathrm{d}t^2}
:::

The second form says the same thing a different way. Velocity is the derivative of position, and acceleration is the derivative of velocity, so acceleration is the derivative of the derivative of position. That is called the **second derivative**, written d²x/dt² and read "d two x by d t squared". On a position–time graph it describes how the slope itself is changing, which shows up as how sharply the graph curves.

The units follow from the definition. Velocity is in metres per second, and acceleration is a change of velocity per second, so its unit is (m/s)/s, written m/s² and said "metres per second per second". An acceleration of 2 m/s² means the velocity grows by 2 m/s in every second that passes.

## Two common confusions

The first confusion is to treat acceleration as if it meant speed. It does not: acceleration measures how quickly the velocity is changing, not how large it is. A car cruising at a steady 100 km/h on a straight motorway is moving fast but has zero acceleration, because its velocity is not changing. A ball thrown straight up is, at the very top of its flight, momentarily at rest, yet its acceleration there is as large as at any other point. A moment before, it was moving upward; a moment after, it will be moving downward. Its velocity is changing steadily from positive to negative, and passing through zero on the way does not stop that change.

The second confusion is to read negative acceleration as slowing down. A negative acceleration means only that the velocity is becoming more negative. What that does to the speed depends on which way the object is already moving. If the velocity is positive, making it more negative brings it closer to zero, so the object slows down. But if the velocity is already negative, making it more negative takes it further from zero, so the object speeds up in the negative direction. The rule that always works is this: an object slows down when its velocity and acceleration have opposite signs, and speeds up when they have the same sign. "Deceleration" is a useful everyday word, but it describes the relationship between the two signs, not the sign of a alone.

## Going up and down the chain

Derivatives carry you down the chain, from position to velocity to acceleration. Integrals carry you back up. Because the area under a rate graph is the total accumulated change, the area under a velocity–time graph is the displacement, and the area under an acceleration–time graph is the change in velocity. So if you know any one of the three quantities at every moment, together with where the object started and how fast it was moving, you can find the others.

Here is the downward direction worked through for a cart whose position is given by a formula.

::: steps From position to acceleration: an example
$ x = 3t^2 - 12t \quad (\text{metres, } t \text{ in seconds})
Suppose a cart on a straight track has this position at time t. At t = 0 it is at the origin. The formula is a complete description of the motion, so everything else about it must be contained in this one expression; the task is to extract it.
---
$ v = \frac{\mathrm{d}x}{\mathrm{d}t} = 6t - 12
Velocity is the derivative of position, so we differentiate each term with the power rule. The term 3t² becomes 3 × 2t = 6t, and −12t, which is t to the first power, becomes −12. At t = 0 this gives v = −12 m/s, so the cart starts out moving backward.
---
$ a = \frac{\mathrm{d}v}{\mathrm{d}t} = 6\,\text{m/s}^2
Acceleration is the derivative of velocity, so we differentiate again. The 6t becomes 6, and the −12 is a constant, which has zero derivative because it does not change. The acceleration is therefore the same at every moment: 6 m/s² in the positive direction.
---
$ 6t - 12 = 0 \quad \Rightarrow \quad t = 2\,\text{s}
To find when the cart stops, we set the velocity to zero and solve. Adding 12 to both sides and dividing by 6 gives t = 2 s. This is the moment the cart turns around.
---
$ x(2) = 3(2)^2 - 12(2) = -12\,\text{m}
Putting t = 2 s back into the position formula tells us where the turn happens: 12 − 24 = −12 m, twelve metres behind the origin. Before this, the velocity was negative while the acceleration was positive, opposite signs, so the cart was slowing down. After it, both are positive, so the cart speeds up as it heads forward.
:::

That example shows the second confusion from the other side. The acceleration was positive the whole time, yet for the first two seconds the cart was slowing down, simply because it was moving the other way.

::: note
Position, velocity and acceleration form a chain. Each is the slope of the graph of the one before it, and each is found from the area under the graph of the one after it.
:::

## How "speed at an instant" became an idea

These definitions now seem obvious, but for most of history they were not available. Aristotle, whose physics dominated European thought for nearly two thousand years, described motion in terms of whole journeys and their causes, and had no clear way to talk about how fast something was moving at a single moment, let alone how quickly that speed was changing.

::: history 1330s–1350s · Oxford and Paris
The first real progress came from a small group of scholars at Merton College, Oxford, in the 1330s, who later became known as the Merton Calculators. Thinkers such as William Heytesbury and Richard Swineshead were interested in how qualities like heat or speed could increase by degrees, and in working that out they defined instantaneous velocity and "uniformly accelerated" motion, meaning motion whose velocity grows by equal amounts in equal times. They proved what is now called the mean speed theorem: a body that accelerates uniformly covers the same distance as one moving steadily at the average of its starting and final speeds. About twenty years later in Paris, Nicole Oresme gave a geometric proof by drawing speed as height above a line of time, so that the distance covered appeared as an area, the same picture used in the previous chapter. Their work was pure reasoning, not tied to any measurement of real falling bodies. Nearly three centuries later Galileo used the same result in his study of falling bodies, and it became one of the foundations of mechanics.
:::

::: quiz
? A ball is thrown straight up. At the top of its flight, what are its velocity and acceleration?
- Both zero
+ Velocity zero, acceleration not zero
- Velocity not zero, acceleration zero
= The ball is at rest for an instant, but its velocity is still changing from upward to downward. Since acceleration measures how quickly velocity changes, it is not zero there.
? An object has v = −4 m/s and a = −2 m/s². What is happening?
- It is slowing down
+ It is speeding up in the negative direction
- It is at rest
= Velocity and acceleration have the same sign, so the velocity is moving further from zero. The object is getting faster while moving in the negative direction.
? An object's position is x = 2t³ m. What is its acceleration at t = 1 s?
- 6 m/s²
+ 12 m/s²
- 2 m/s²
= Differentiating once gives v = 6t², and differentiating again gives a = 12t. At t = 1 s that is 12 m/s²; the 6 m/s² option is the velocity at that moment, not the acceleration.
:::
