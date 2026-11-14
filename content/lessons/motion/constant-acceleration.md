Many real motions have an acceleration that stays constant, or close enough to constant that the difference does not matter: a stone dropped from a bridge, a car pulling away at steady throttle, a train braking smoothly into a station. In this one case, the calculus of the previous lessons collapses into four short algebraic equations, and those four equations solve most of the problems a first course in mechanics can pose. This lesson derives each of them from the ideas of slope and area, so that you know where they come from, what each one is for, and, just as important, when they stop being true.

## Setting up the problem

To write equations we need names for things. Let t be the time that has passed since we started the clock. At that starting moment, t = 0, the object is at position x₀ and has velocity v₀; the small zero is a reminder that these are the starting values. At a later time t, its position is x and its velocity is v. Finally, a is the acceleration, which we assume has the same value throughout.

That last assumption is the price of admission. Everything in this lesson depends on it, and the equations give wrong answers, sometimes badly wrong, when it fails. A car whose driver eases off the brake halfway through a stop, or a skydiver slowed more and more by air resistance, does not have a constant acceleration, and needs the fuller methods of calculus.

## How does the velocity change?

Acceleration is the rate at which velocity changes. If that rate is constant, then the velocity changes by the same amount every second: by a in the first second, another a in the next, and so on. After t seconds, the velocity has changed by a times t in total, and adding that to the starting velocity gives the first equation.

::: equation 1. Velocity over time
v = v_0 + at
:::

On a graph of velocity against time, this is a straight line. It starts at height v₀ when t = 0, and its slope is a, which fits what we learned earlier: the slope of a velocity–time graph is the acceleration, and a constant slope means a straight line.

## Where is the object? Position over time

Finding the position takes one more idea. In the lesson on accumulation, we saw that the displacement over some time interval is the area under the velocity–time graph for that interval. Here the graph is a straight line, so the region under it, from time 0 to time t, is a simple shape: a rectangle with a triangle sitting on top. We can find its area with nothing more than school geometry.

::: steps Deriving position over time
$ x - x_0 = \text{area under } v \text{ from } 0 \text{ to } t
The displacement, x − x₀, is the total change in position, and position changes at a rate equal to the velocity. So the displacement is the velocity accumulated over time, which is the area under the velocity–time graph. This turns a motion problem into a geometry problem.
---
$ \text{rectangle} = v_0 t
Draw a horizontal line at height v₀, the starting velocity. The part of the region below it is a rectangle, v₀ high and t wide, so its area is v₀t. This is the distance the object would cover if it simply kept its starting velocity and never accelerated.
---
$ \text{triangle} = \tfrac{1}{2} \cdot t \cdot (v - v_0)
Above the rectangle sits a right triangle. Its base runs along the time axis for a length t, and its height is the gain in velocity, v − v₀. A triangle's area is half its base times its height, which gives this expression. This piece is the extra distance gained by speeding up.
---
$ \text{triangle} = \tfrac{1}{2} \cdot t \cdot at = \tfrac{1}{2}at^2
The triangle's height, v − v₀, is still written in terms of the final velocity, which we would rather not need. Equation 1 says that v − v₀ = at, so we replace it. The triangle's area becomes ½at², which depends only on the acceleration and the time.
---
$ x = x_0 + v_0 t + \tfrac{1}{2}at^2
The total displacement is the sum of the two areas, so x − x₀ = v₀t + ½at². Adding x₀ to both sides gives the position itself, which is usually what a problem asks for.
:::

::: equation 2. Position over time
x = x_0 + v_0 t + \tfrac{1}{2}at^2
:::

The t² term is the signature of acceleration. It means that position, plotted against time, is a parabola rather than a straight line, and that the effect of acceleration grows quickly: twice as long a time gives four times the extra distance.

We can check this result by going down the chain instead of up it. Differentiating equation 2 with the power rule gives dx/dt = v₀ + at, which is exactly equation 1, and differentiating once more gives a, the constant we started with. The geometry and the calculus agree, which is a good sign that nothing has gone wrong.

## A short cut through the average

There is a quicker way to the displacement when you know both the starting and final velocities. Because the velocity rises in a straight line, it spends just as much time above its midpoint value as below it, so the average velocity over the whole interval is exactly halfway between v₀ and v. Displacement is average velocity times time, which gives the third equation.

::: equation 3. Displacement from average velocity
x - x_0 = \tfrac{1}{2}(v_0 + v)t
:::

This is the mean speed theorem of the Merton scholars, which appeared in the history of the previous lesson, now written as an equation. It carries a warning with it. The average velocity equals the midpoint only when velocity changes in a straight line, that is, only when acceleration is constant. If the velocity graph curves, the object spends more time near one end of its range than the other, and the midpoint is no longer the average.

## Answering questions that do not mention time

Many practical questions involve speeds and distances but say nothing about time. How fast is a car still going after braking over 40 metres? How long a runway does a plane need to reach take-off speed? Equations 1 to 3 all contain t, so to answer these questions directly we want an equation without it. We can get one by combining two of the equations we already have and eliminating t between them.

::: steps Deriving the time-free equation
$ t = \frac{v - v_0}{a}
Start from equation 1, v = v₀ + at, and solve it for the time: subtract v₀ from both sides, then divide by a. This expresses t in terms of velocities and the acceleration. Dividing by a needs a to be non-zero, but that is no loss, because with zero acceleration the velocity never changes and the question has an obvious answer.
---
$ x - x_0 = \tfrac{1}{2}(v_0 + v) \cdot \frac{v - v_0}{a}
Now put this expression for t into equation 3 in place of the time. The two sides are still equal, because we have only replaced t with something equal to it. The equation now contains no time at all, which was the aim; what remains is to tidy it up.
---
$ 2a(x - x_0) = (v + v_0)(v - v_0)
Multiply both sides by 2a to clear the fraction and the factor of ½. Doing the same thing to both sides of an equation keeps it true. The right-hand side is now a product of a sum and a difference of the same two velocities.
---
$ 2a(x - x_0) = v^2 - v_0^2
A sum times a difference of the same two quantities is always the difference of their squares, because when (v + v₀)(v − v₀) is multiplied out, the cross terms vv₀ and −v₀v cancel. That leaves v² − v₀².
---
$ v^2 = v_0^2 + 2a(x - x_0)
Finally, add v₀² to both sides so that the final velocity stands alone. The result links the change in speed directly to the distance covered, with no mention of how long it took.
:::

::: equation 4. Velocity and displacement
v^2 = v_0^2 + 2a(x - x_0)
:::

## Choosing the right equation

There are five quantities in play: displacement, starting velocity, final velocity, acceleration and time. Each of the four equations contains four of them and leaves one out. Equation 1 has no position, equation 2 has no final velocity, equation 3 has no acceleration and equation 4 has no time. That observation gives a simple strategy. Find the one quantity that a problem neither gives you nor asks for, and use the equation that leaves it out; then every symbol in that equation is either known or wanted.

Try it on a braking car. A car travelling at 30 m/s brakes with a steady deceleration of 6 m/s². How far does it travel before it stops? The problem says nothing about time and does not ask for it, so equation 4 is the one to use. The final velocity is v = 0, since the car stops. Taking the direction of travel as positive, the acceleration is a = −6 m/s², negative because it opposes the motion. Equation 4 becomes 0 = 30² + 2 × (−6) × Δx, where Δx is the stopping distance. That gives 12Δx = 900, so Δx = 75 m.

The structure of equation 4 holds a lesson for every driver. With the final speed zero, it says the stopping distance equals v₀²/2a, so it grows with the square of the starting speed. Halve the speed to 15 m/s, with the same brakes, and the car stops in about 19 m, only a quarter of the distance. Double the speed and the stopping distance quadruples. This is why a modest increase in speed makes such a large difference to whether a car can stop in time.

::: note
When acceleration is constant, velocity changes in a straight line with time and position follows a parabola. Four equations follow from the ideas of slope and area; choose the one that leaves out the quantity you neither know nor need.
:::

::: quiz
? A cyclist starts from rest and accelerates at 2 m/s² for 5 s. How far does she travel?
- 10 m
+ 25 m
- 50 m
= Equation 2 with x₀ = 0 and v₀ = 0 gives x = ½ × 2 × 5² = 25 m. The 50 m option forgets the factor of ½ that comes from the triangle's area.
? If a car's speed doubles, how does its braking distance change, assuming the same deceleration?
- It doubles
+ It quadruples
- It stays the same
= With the final speed zero, equation 4 gives a stopping distance of v₀²/2a, which is proportional to the square of the starting speed. Doubling v₀ therefore multiplies the distance by four.
? Why does x − x₀ = ½(v₀ + v)t fail when the acceleration varies?
+ The average velocity is no longer the midpoint of v₀ and v
- Velocity cannot change if acceleration varies
- The equation has the wrong units
= The average velocity equals the midpoint of the starting and final values only when velocity changes in a straight line with time. If the acceleration varies, the velocity graph curves and the object spends more time near one end of its range.
:::
