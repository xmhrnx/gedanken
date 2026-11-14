The first law says a net force changes velocity. The second law says by how much. It is the most useful equation in classical physics, and nearly every problem in mechanics starts by writing it down.

## The law

Push a trolley and it accelerates. Push twice as hard and the acceleration doubles. Load it with twice the mass and the same push gives half the acceleration. Experiments of this kind give:

::: equation Newton's second law
\sum \vec{F} = m\vec{a}
:::

The net force on a body equals its mass times its acceleration. Three points matter.

- **Net force.** The left side is the vector sum of every force acting on the body, not any single one of them.
- **Vectors.** The acceleration points in the direction of the net force. Because vectors add by components, the law holds separately in each direction: ΣFₓ = maₓ and ΣF_y = ma_y, which is why horizontal and vertical motion of a projectile could be treated apart.
- **Mass as inertia.** Rearranged as a = F/m, the law says mass measures resistance to acceleration. This is the number the last lesson called inertia.

The SI unit of force follows from the equation. One **newton** (N) is the force that gives 1 kg an acceleration of 1 m/s²: 1 N = 1 kg·m/s². That is roughly the weight of a small apple, about 100 g.

## Weight, and why everything falls alike

The force of gravity on a body near the Earth's surface is its **weight**, W = mg, directed downward. Weight is a force, measured in newtons; mass is an amount of inertia, measured in kilograms. A 70 kg person weighs about 690 N on Earth and about 110 N on the Moon, but has the same mass on both.

The second law now explains a puzzle from the free-fall lesson. If gravity is the only force, then:

::: steps Every body falls with the same acceleration
$ \sum F = W = mg
In free fall the only force is the body's weight. Take downward as positive.
---
$ ma = mg
Apply the second law: the net force equals mass times acceleration.
---
$ a = g
Divide by m. A heavier body feels a larger force, but it has proportionally more inertia, and the two cancel exactly.
:::

This cancellation is not obvious. The m in W = mg measures how strongly gravity pulls (*gravitational mass*); the m in F = ma measures resistance to acceleration (*inertial mass*). Experiments show they are equal to better than one part in 10¹⁴. Newton noticed the coincidence; Einstein built general relativity on it.

## Example: a lift

You stand on bathroom scales in a lift. The scales read the force with which they push up on you, call it N. Two forces act on you: N upward and your weight mg downward.

::: steps What the scales read
$ N - mg = ma
Take upward as positive. The net force is N − mg, and the second law sets it equal to ma, where a is the lift's acceleration.
---
$ N = m(g + a)
Add mg to both sides and factor out m.
---
$ N = 70 \times (9.8 + 2.0) \approx 830\,\text{N}
For a 70 kg person in a lift accelerating upward at 2.0 m/s², the scales read about 830 N, as if you weighed 85 kg.
:::

Accelerating downward, a is negative and the reading drops. If the cable snaps, a = −g and N = 0: the scales read nothing. Nothing about gravity has changed; the floor has simply stopped pushing. That is what astronauts in orbit experience, as the circular motion lesson will show.

## Solving problems

The method is always the same. Pick the body. List every force acting on it, ideally in a sketch with an arrow for each (a *free-body diagram*). Choose axes. Write ΣF = ma along each axis and solve. Most errors come from a missing force or one counted twice, not from the algebra.

::: history 1687–1750 · Cambridge and Berlin
Newton never wrote F = ma. His second law said the *change of motion* is proportional to the force impressed and happens along the line of that force, where "motion" meant mass times velocity, now called momentum. He worked with sudden impulses and with geometry, not with equations of this form. Leonhard Euler wrote the law as force equals mass times acceleration, applied component by component, in a paper of 1750, and showed it could be used for any mechanical system. The form in textbooks today is his.
:::

::: note
Net force equals mass times acceleration. Force decides how velocity changes; mass decides how hard it is to change.
:::

::: quiz
? A 2 kg box is pushed with 10 N to the right while friction pulls 4 N to the left. What is its acceleration?
- 5 m/s²
+ 3 m/s²
- 7 m/s²
= Net force is 10 − 4 = 6 N to the right; a = 6 N ÷ 2 kg = 3 m/s².
? A lift moves upward at a constant 3 m/s. Compared with your weight, the scales read
- more
- less
+ the same
= Constant velocity means a = 0, so N = mg. Only acceleration changes the reading.
? Why does a bowling ball not fall faster than a tennis ball, ignoring air?
- Gravity pulls equally on both
+ Gravity pulls harder on the bowling ball, but its greater mass resists acceleration equally more
- The Earth's acceleration cancels the difference
= Weight is mg and inertia is m, so a = mg/m = g for both.
:::
