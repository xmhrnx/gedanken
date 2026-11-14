Some quantities are fully described by one number and a unit. Mass, temperature, time and energy are like this: 3 kg is 3 kg whichever way you face. These are *scalars*.

Others need a direction as well. "Walk 5 km" does not tell you where you end up; "walk 5 km north" does. Displacement (the change in position), velocity, acceleration and force all have direction. These are *vectors*. A vector has a *magnitude* (its size, always positive) and a direction.

In print, vectors are often written in bold, **v**, or with an arrow over the letter. The magnitude is written |**v**| or just v. This app writes vectors in bold when the distinction matters.

## Drawing vectors

A vector is drawn as an arrow: its length shows the magnitude and it points in the direction. Where you draw the arrow does not matter. Two arrows of the same length and direction are the same vector.

## Adding vectors

Walk 3 km east, then 4 km north. You have walked 7 km, but you are not 7 km from where you started. Displacements add *tip to tail*: draw the second arrow starting where the first ends, and the sum is the arrow from the start of the first to the end of the second.

Here the two legs are at right angles, so the sum is the hypotenuse of a right triangle. Call its length d:

::: steps Adding 3 km east and 4 km north
$ d^2 = 3^2 + 4^2
The two legs and the total displacement form a right triangle, so Pythagoras' theorem applies.
---
$ d = \sqrt{25} = 5\,\text{km}
Take the square root. You are 5 km from the start.
---
$ \tan\theta = \frac{4}{3}, \quad \theta \approx 53^\circ
The angle θ measured from east satisfies tan θ = opposite/adjacent. So the displacement points 53° north of east.
:::

## Components

Tip-to-tail drawing works but is slow for anything beyond two vectors. The efficient method is to split each vector into *components*: how much of it points along x (say, east) and how much along y (north). Write a vector as a pair (a₁, a₂).

A vector of magnitude A at angle θ from the x-axis has components:

::: equation Components of a vector
a_1 &= A\cos\theta \\ a_2 &= A\sin\theta
:::

These come from the right triangle the vector makes with the axes: cos θ is adjacent over hypotenuse, sin θ is opposite over hypotenuse.

Going back the other way:

::: equation Magnitude and direction from components
A &= \sqrt{a_1^2 + a_2^2} \\ \tan\theta &= \frac{a_2}{a_1}
:::

With components, adding vectors is just adding numbers, separately for each direction:

::: equation
(a_1, a_2) + (b_1, b_2) = (a_1 + b_1,\ a_2 + b_2)
:::

This works because x and y are independent: moving east does not change how far north you are. The same independence is what makes projectile motion simple later on.

Multiplying a vector by a scalar k stretches it: k(a₁, a₂) = (ka₁, ka₂). A negative k flips its direction. Subtracting **b** means adding −**b**.

::: note
Vectors add by components. Split each one into x and y parts, add the parts separately, then recombine with Pythagoras if you need the magnitude.
:::

::: history 1880s · New Haven and London
Vectors as we use them are surprisingly recent. In the 1840s William Rowan Hamilton invented quaternions, four-part numbers that could describe rotations, and James Clerk Maxwell first wrote his electromagnetism partly in that language. In the 1880s Josiah Willard Gibbs at Yale and Oliver Heaviside in England independently stripped quaternions down to the three-component vectors and dot and cross products taught today. Quaternion enthusiasts fought this for years, but the simpler system won.
:::

::: quiz
? A plane flies at 200 km/h north through a wind blowing 50 km/h east. What is its speed over the ground?
- 250 km/h
+ About 206 km/h
- 150 km/h
= The velocities are at right angles, so the speed is √(200² + 50²) = √42 500 ≈ 206 km/h.
? A force of 10 N acts at 30° above the horizontal. What is its horizontal component?
- 5 N
+ About 8.7 N
- 10 N
= Horizontal component = 10 cos 30° ≈ 10 × 0.866 = 8.66 N. The vertical component is 10 sin 30° = 5 N.
? Which of these is a scalar?
- Velocity
- Force
+ Temperature
= Temperature has no direction. Velocity and force do.
:::
