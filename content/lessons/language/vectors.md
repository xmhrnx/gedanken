Some quantities are fully described by a single number and a unit, and others are not. Knowing that a parcel weighs 3 kilograms tells you everything about its mass, but knowing that a hiker walked 5 kilometres does not tell you where she ended up. This lesson is about quantities of the second kind, which carry a direction as well as a size. Most of the important quantities in mechanics, including velocity, acceleration and force, are of this kind, so learning to describe and combine them is the first real step into the subject.

## Size alone, or size and direction?

Mass, temperature, time and energy are described completely by one number with a unit. Three kilograms is three kilograms whichever way you face, and a temperature of 20 °C does not point anywhere. Quantities like these are called **scalars**, from the idea that a single scale is enough to measure them.

Now think again about the hiker. "Walk 5 km" leaves her anywhere on a circle of radius 5 km around the start. "Walk 5 km north" puts her at one particular spot. The second instruction contains more information, and that extra information is a direction. Quantities that need a direction to be complete are called **vectors**. A vector has a **magnitude**, which is its size and is never negative, and a direction.

Many quantities in mechanics are vectors. **Displacement**, the change in an object's position, is the simplest: it says how far the object moved and in which direction. Velocity is a vector because a car doing 50 km/h northward and one doing 50 km/h southward are in very different situations. Acceleration and force are vectors for the same reason: a push to the left and a push to the right of the same strength have opposite effects.

In print, vectors are usually written in bold, **v**, or with a small arrow over the letter. The magnitude is written |**v**|, or simply v in italic or plain type. This app writes vectors in bold when the distinction matters.

## Drawing vectors as arrows

The natural picture of a vector is an arrow. Its length shows the magnitude, drawn to some chosen scale, and it points in the vector's direction. One feature of this picture is easy to overlook: where you draw the arrow does not matter. Two arrows with the same length and the same direction represent the same vector, wherever they sit on the page. A displacement of 5 km north is the same displacement whether it starts in Edinburgh or in Lisbon. This freedom to slide arrows around without changing them is what makes the rule for adding vectors work.

## Adding vectors

Suppose the hiker walks 3 km east and then 4 km north. She has walked 7 km in total, but she is not 7 km from where she started, because the two legs point in different directions and do not simply pile up. To find where she ends up, we need a rule for adding vectors.

The rule follows directly from what displacement means. Draw the first leg as an arrow. The second leg starts where the first one ended, so draw its arrow starting from the tip of the first. The total displacement is the arrow from the very start to the very end. This is called adding **tip to tail**, and because arrows can be slid around freely, it works for any two vectors of the same kind, not only displacements.

In this example the legs are at right angles, so the three arrows form a right-angled triangle, with the total displacement as its longest side. Call the length of that side d.

::: steps Adding 3 km east and 4 km north
$ d^2 = 3^2 + 4^2
The two legs meet at a right angle, so the triangle they make with the total displacement is a right triangle. Pythagoras' theorem says that in such a triangle the square of the longest side equals the sum of the squares of the other two. That gives us an equation for the one length we do not know.
---
$ d = \sqrt{25} = 5\,\text{km}
The squares are 9 and 16, which add to 25. Taking the square root, and keeping the positive answer because a distance cannot be negative, gives 5 km. So after 7 km of walking, the hiker is only 5 km from her starting point.
---
$ \tan\theta = \frac{4}{3}
The magnitude is only half the answer; we also need the direction. Measure the angle θ from the east direction. In a right triangle, the tangent of an angle is the opposite side divided by the adjacent side, and here the side opposite θ is the 4 km north leg while the adjacent side is the 3 km east leg.
---
$ \theta \approx 53^\circ
Taking the inverse tangent of 4/3 gives about 53°. The hiker's displacement is therefore 5 km in a direction 53° north of east, which is a complete description of where she ended up: a size and a direction.
:::

## Splitting a vector into components

Tip-to-tail drawing is fine for two vectors at right angles, but it becomes slow and inaccurate for three or four vectors at awkward angles. The method physicists actually use is to split each vector into **components**: the part of it that points along one chosen direction, called x (east, say), and the part that points along a second direction at right angles, called y (north). A vector can then be written as a pair of numbers, (a₁, a₂), giving its x and y components.

To find the components of a vector with magnitude A, pointing at an angle θ above the x direction, picture the right triangle it makes with the axes. The vector is the longest side. Its x component is the side lying along the x axis, next to the angle, and its y component is the side opposite the angle. The definitions of cosine (adjacent over longest side) and sine (opposite over longest side) then give the components directly:

::: equation Components of a vector
a_1 &= A\cos\theta \\ a_2 &= A\sin\theta
:::

For example, a force of 10 N acting at 30° above the horizontal has a horizontal component of 10 cos 30° ≈ 8.7 N and a vertical component of 10 sin 30° = 5 N. Those two numbers say how hard the force pushes sideways and how hard it lifts, which is often exactly what a problem needs to know.

The same right triangle lets you go back the other way. Pythagoras gives the magnitude from the components, and the tangent gives the angle, just as in the hiker example:

::: equation Magnitude and direction from components
A &= \sqrt{a_1^2 + a_2^2} \\ \tan\theta &= \frac{a_2}{a_1}
:::

## Why components make addition easy

The reward for splitting vectors up is that adding them becomes ordinary arithmetic. To add two vectors, add their x components together and their y components together, separately:

::: equation
(a_1, a_2) + (b_1, b_2) = (a_1 + b_1,\ a_2 + b_2)
:::

This works because the two directions are independent of each other. Walking east does not change how far north you are, and walking north does not change how far east you are, so the eastward parts of any journey simply add up among themselves, and so do the northward parts. Check it on the hiker: her first leg is (3, 0) and her second is (0, 4), so the total is (3, 4), and the formulas above turn that back into 5 km at 53°. The same independence will turn up again when we study projectiles, where the horizontal and vertical parts of a thrown ball's motion can be treated as two separate, simple problems.

Two smaller operations complete the toolkit. Multiplying a vector by an ordinary number k stretches it by that factor, which in components means k(a₁, a₂) = (ka₁, ka₂). A negative k reverses the direction as well, so −**b** is **b** turned around. Subtracting a vector then needs no new rule: **a** − **b** is just **a** + (−**b**).

::: note
A vector has a size and a direction, and vectors add by their components. Split each vector into x and y parts, add the parts in each direction separately, and recombine with Pythagoras' theorem if you need the overall size.
:::

## Where vectors came from

It is surprising how recently this neat system arrived. For two centuries after Newton, physicists wrote out every equation in separate x, y and z versions, three times over.

::: history 1843–1890s · Dublin, New Haven and London
In 1843 the Irish mathematician William Rowan Hamilton, after years of trying to extend complex numbers to three dimensions, realised while walking along a canal in Dublin that the trick needed four parts, not three. He called these new numbers quaternions and carved their key rule into a bridge on the spot. Quaternions could describe rotations elegantly, and James Clerk Maxwell used some of their notation when he wrote up his theory of electricity and magnetism in 1873. But physicists found them awkward: the fourth part rarely meant anything physical, and the rules for multiplying them were strange. In the 1880s Josiah Willard Gibbs at Yale and Oliver Heaviside in England, working independently, kept the useful three-part piece of a quaternion and threw the rest away, inventing the vectors and vector products taught today. Hamilton's followers, led by Peter Guthrie Tait, attacked the new system fiercely in print through the 1890s; Tait called it "a sort of hermaphrodite monster". The vector system won anyway, because it said what physicists needed to say with less machinery.
:::

::: quiz
? A plane flies at 200 km/h north through a wind blowing 50 km/h east. What is its speed over the ground?
- 250 km/h
+ About 206 km/h
- 150 km/h
= The two velocities are at right angles, so they add like the legs of a right triangle: √(200² + 50²) = √42 500 ≈ 206 km/h. Simply adding 200 and 50 would only be right if the wind blew in the same direction as the plane.
? A force of 10 N acts at 30° above the horizontal. What is its horizontal component?
- 5 N
+ About 8.7 N
- 10 N
= The horizontal component is the side next to the angle, so it is 10 cos 30° ≈ 10 × 0.866 ≈ 8.7 N. The 5 N option is the vertical component, 10 sin 30°.
? Which of these is a scalar?
- Velocity
- Force
+ Temperature
= Temperature is fully described by a single number with a unit and has no direction. Velocity and force both need a direction to be complete, so they are vectors.
:::
