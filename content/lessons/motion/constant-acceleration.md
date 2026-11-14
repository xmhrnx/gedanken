Many real motions have an acceleration that is constant, or close enough: a dropped stone, a car pulling away at steady throttle, a train braking. For this case the calculus of the previous lessons reduces to four algebraic equations that solve most textbook problems. This lesson derives them, so that you know where they come from and when they apply.

## Notation

- t: time elapsed since we started the clock
- x₀ and v₀: position and velocity at t = 0
- x and v: position and velocity at time t
- a: the acceleration, assumed constant throughout

## Velocity

If acceleration is constant, velocity changes by the same amount every second. After t seconds it has changed by a·t:

::: equation 1. Velocity over time
v = v_0 + at
:::

On a velocity–time graph this is a straight line with slope a, starting at height v₀.

## Position

Displacement is the area under the velocity–time graph. That graph is a straight line, so the area under it is a trapezoid, which splits into a rectangle and a triangle.

::: steps Deriving position over time
$ x - x_0 = \text{area under } v \text{ from } 0 \text{ to } t
Displacement is the accumulation of velocity, which is the area under its graph.
---
$ \text{rectangle} = v_0 t
The bottom part of the region is a rectangle of height v₀ and width t.
---
$ \text{triangle} = \tfrac{1}{2} \cdot t \cdot (v - v_0) = \tfrac{1}{2} \cdot t \cdot at = \tfrac{1}{2}at^2
Above the rectangle is a triangle of base t and height v − v₀. From equation 1, v − v₀ = at.
---
$ x = x_0 + v_0 t + \tfrac{1}{2}at^2
Add the two areas and move x₀ to the right-hand side.
:::

::: equation 2. Position over time
x = x_0 + v_0 t + \tfrac{1}{2}at^2
:::

You can check this with the derivative: dx/dt = v₀ + at, which is equation 1, and differentiating again gives a. The equations are consistent.

## Average velocity

Because velocity changes linearly, the average velocity over the interval is simply the midpoint of the starting and ending values. This is the Merton rule from the previous lesson's history note.

::: equation 3. Displacement from average velocity
x - x_0 = \tfrac{1}{2}(v_0 + v)t
:::

This only holds when acceleration is constant. If the velocity curve bends, the average is not the midpoint.

## Eliminating time

Often a problem gives distances and speeds but no time: how fast is a car going after braking for 40 m? Combine equations 1 and 3 to get rid of t.

::: steps Deriving the time-free equation
$ t = \frac{v - v_0}{a}
Solve equation 1 for t. This needs a ≠ 0; with zero acceleration the question is trivial anyway.
---
$ x - x_0 = \tfrac{1}{2}(v_0 + v) \cdot \frac{v - v_0}{a}
Substitute this t into equation 3.
---
$ 2a(x - x_0) = (v + v_0)(v - v_0) = v^2 - v_0^2
Multiply both sides by 2a. The right side is a difference of squares.
---
$ v^2 = v_0^2 + 2a(x - x_0)
Rearrange.
:::

::: equation 4. Velocity and displacement
v^2 = v_0^2 + 2a(x - x_0)
:::

## Using them

Each equation leaves out one quantity: equation 1 has no position, 2 has no final velocity, 3 has no acceleration, 4 has no time. Identify which quantity a problem neither gives nor asks for, and use the equation without it.

Example: a car at 30 m/s brakes at 6 m/s². How far does it travel before stopping? Time is not given or asked, so use equation 4 with v = 0 and a = −6 m/s² (negative because it opposes the motion): 0 = 30² + 2(−6)Δx, so Δx = 900/12 = 75 m.

Note that stopping distance grows with the *square* of the speed. At 15 m/s the same brakes stop the car in under 19 m, a quarter of the distance.

::: note
With constant acceleration, velocity is a straight line in time and position is a parabola. Four equations follow; pick the one that leaves out the quantity you neither know nor need.
:::

::: quiz
? A cyclist starts from rest and accelerates at 2 m/s² for 5 s. How far does she travel?
- 10 m
+ 25 m
- 50 m
= Equation 2 with x₀ = 0, v₀ = 0: x = ½ × 2 × 5² = 25 m.
? If a car's speed doubles, how does its braking distance change (same deceleration)?
- It doubles
+ It quadruples
- It stays the same
= From v² = 2aΔx with final speed zero, Δx is proportional to v₀², so doubling v₀ quadruples Δx.
? Why does x − x₀ = ½(v₀ + v)t fail when acceleration varies?
+ The average velocity is no longer the midpoint of v₀ and v
- Velocity cannot change if acceleration varies
- The equation has the wrong units
= The midpoint is the average only when velocity changes linearly with time.
:::
