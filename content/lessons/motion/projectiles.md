A **projectile** is anything that is thrown, kicked or fired and then left to move under gravity alone: a football in flight, a cannonball, the stream of water from a garden hose. Its path is a curve, and at first sight a curve looks as if it needs new mathematics, something beyond the straight-line motion of the last few lessons. It does not. The whole problem yields to one idea, which is to treat the horizontal and vertical parts of the motion as two separate straight-line problems. This lesson shows why that is allowed, uses it to find the shape of the path, and works out how far a projectile goes and which launch angle sends it furthest.

## Why the motion splits in two

Ignore air resistance for now. Then once a projectile has left the hand, the gun or the boot, the only force acting on it is gravity, and gravity pulls straight down. That single fact settles the behaviour of each direction separately.

Horizontally, nothing pushes or pulls the projectile, so there is nothing to change its horizontal velocity. Whatever horizontal speed it had at launch, it keeps for the whole flight. Vertically, the projectile has a downward acceleration of g, exactly as a dropped stone does, because the same gravity acts on it in the same way.

The two motions do not interfere with each other, for the reason the vectors lesson gave: moving sideways does not change how fast you are moving up or down, and gravity, pulling straight down, has no sideways part with which to affect horizontal motion. This leads to a result that surprises most people the first time they meet it. Fire a bullet horizontally from a rifle and, at the same instant and from the same height, simply drop another bullet. On flat ground and without air, the two bullets hit the ground at the same moment. The fired bullet travels hundreds of metres sideways while the dropped one goes nowhere, but their vertical motions are identical: both start with zero vertical velocity and both fall with acceleration g. The fired bullet is not held up by its speed; it simply covers a lot of horizontal distance while it falls.

## Setting up the problem

Suppose a projectile is launched from the origin with speed v₀ at an angle θ above the horizontal. To use the split, we break the launch velocity into its horizontal and vertical components, exactly as in the vectors lesson, with the launch velocity as the longest side of a right triangle:

::: equation Launch velocity components
u &= v_0\cos\theta \quad \text{(horizontal)} \\ w &= v_0\sin\theta \quad \text{(vertical, upward)}
:::

Let x measure horizontal distance and y measure height, with upward as positive. Now apply the constant-acceleration equations to each direction on its own. Horizontally the acceleration is zero, so the position is simply speed times time. Vertically the acceleration is −g, so the position equation from the earlier lesson applies with v₀ replaced by w:

::: equation Position at time t
x &= ut \\ y &= wt - \tfrac{1}{2}gt^2
:::

These two equations are a complete description of the flight. The only thing connecting them is the time t, which is the same for both, because they describe the same object at the same moment.

::: model projectile
Change the launch angle and speed and watch the path. Find the angle that gives the longest range, and notice that two angles adding up to 90°, such as 30° and 60°, land at the same spot.
:::

## What shape is the path?

The equations above give x and y separately in terms of time. To see the shape of the path itself, we want a single equation linking height to horizontal distance, with time removed.

::: steps The trajectory is a parabola
$ t = \frac{x}{u}
Start with the horizontal equation, x = ut, and divide both sides by u to get the time. Because the horizontal speed never changes, horizontal distance is a direct measure of time elapsed: twice as far across means twice as long in the air. This is the link we need between the two motions.
---
$ y = w\left(\frac{x}{u}\right) - \tfrac{1}{2}g\left(\frac{x}{u}\right)^2
Now replace t in the vertical equation with x/u. The two expressions are equal at every moment of the flight, so the equation stays true. Height is now written in terms of horizontal distance, with time gone.
---
$ \frac{w}{u} = \frac{v_0\sin\theta}{v_0\cos\theta} = \tan\theta, \qquad u^2 = v_0^2\cos^2\theta
To tidy the result, we use the components. In the ratio w/u the launch speed cancels, leaving sin θ / cos θ, which is the tangent of the launch angle. And squaring u gives the expression on the right, which we need for the second term.
---
$ y = x\tan\theta - \frac{g}{2v_0^2\cos^2\theta}\,x^2
Putting those into the previous equation gives height as a function of horizontal distance. For a given launch, tan θ and the fraction in front of x² are just fixed numbers, so the equation has the form y = (a number) × x − (a positive number) × x². That is the equation of a parabola that opens downward.
:::

So without air resistance, every projectile follows a parabola. The first term on the right, x tan θ, is the straight line the projectile would follow if there were no gravity, aimed along the launch direction. The second term is how far gravity has pulled it below that line, and it grows as the square of the distance, because the distance fallen grows as the square of the time.

## How long, how far, how high?

On level ground, the projectile lands when its height returns to zero. That condition is enough to find both how long it is in the air and how far it travels.

::: steps Range on level ground
$ wT - \tfrac{1}{2}gT^2 = 0
Let T be the time at which the projectile lands. At that moment its height is zero, so we set y = 0 in the vertical equation, with t = T. This gives one equation with one unknown, T.
---
$ T\left(w - \tfrac{1}{2}gT\right) = 0
Both terms contain T, so we take it out as a common factor. A product is zero only when one of its factors is zero, so either T = 0 or the bracket is zero. The solution T = 0 is the moment of launch, when the projectile is also at ground level; it is correct but not the one we want.
---
$ T = \frac{2w}{g} = \frac{2v_0\sin\theta}{g}
Setting the bracket to zero and solving gives the landing time. It depends only on the vertical part of the launch velocity, as it should: the time in the air is decided entirely by how fast the projectile was thrown upward, and the horizontal motion plays no part.
---
$ R = uT = \frac{2v_0^2\sin\theta\cos\theta}{g}
The range R is the horizontal distance covered during the flight. Since the horizontal speed u is constant, the range is that speed times the flight time. Multiplying u = v₀ cos θ by T gives this expression.
---
$ R = \frac{v_0^2\sin(2\theta)}{g}
A standard trigonometric identity says that 2 sin θ cos θ is equal to sin(2θ). Using it collapses the angle dependence into a single sine, which makes it easy to see how the range changes with the launch angle.
:::

The last form answers the gunner's oldest question. The sine of an angle is largest, equal to 1, when the angle is 90°. So the range is greatest when 2θ = 90°, that is, when θ = 45°, and that maximum range is v₀²/g. The reason is a trade-off. A steep launch puts most of the speed into the vertical direction, giving a long flight but little horizontal speed; a shallow launch gives plenty of horizontal speed but a short flight. At 45° the speed is split equally, and the product of the two is as large as it can be.

The formula shows something else. Since sin(2θ) equals sin(180° − 2θ), the angles θ and 90° − θ give the same range. A ball launched at 30° and one launched at 60°, at the same speed, land at the same place. The 60° shot gets there by a higher, slower route, spending longer in the air but moving more slowly across, while the 30° shot takes a flatter, quicker path. Doubling the launch speed, meanwhile, quadruples the range, because the range depends on the square of the speed.

The greatest height comes from the vertical motion alone. At the top of the path, the vertical velocity is momentarily zero, and the horizontal velocity carries on unchanged. Using the time-free equation v² = v₀² + 2aΔy for the vertical direction, with starting vertical speed w, final vertical speed zero and acceleration −g, gives 0 = w² − 2gH. So the maximum height is H = w²/(2g), which again depends only on how fast the projectile was launched upward.

## What air changes

Everything so far has ignored the air, and for real projectiles that matters. Air exerts **drag**, a force that opposes the motion and grows quickly with speed, roughly as the square of the speed for most everyday objects. Drag steadily robs the projectile of horizontal speed, so it covers less ground in the second half of its flight than in the first. The path becomes lopsided, rising at a shallower slope than it falls, and the angle for greatest range drops below 45°.

How much this matters depends on the projectile. For a heavy, slow shot put, the parabola is a good approximation. For a golf ball it is poor, because drag is large and the ball's backspin also gives it lift, which can keep it in the air much longer than a parabola predicts. For an artillery shell travelling at hundreds of metres per second, the parabola is badly wrong, and gunners rely on detailed tables computed with drag included.

::: note
A projectile's horizontal and vertical motions are independent. Horizontally its velocity never changes; vertically it is in free fall. Solve each motion on its own, and link them only through the time.
:::

## How the parabola was found

The idea that horizontal and vertical motion can simply be combined seems natural now, but it took a long time to arrive, because the older theory of motion said something quite different.

::: history 1537–1638 · Venice and Leiden
Medieval scholars explained a thrown object's motion by "impetus", a kind of driving power given to it at launch, which gradually wore out. On this view, a cannonball flew in a straight line along its launch direction until its impetus ran low, then curved round and dropped straight down. In 1537 the Italian mathematician Niccolò Tartaglia, writing for gunners, argued that the path must actually be curved all the way, and he stated that a gun aimed at 45° reaches furthest, but he could not give a satisfactory reason. The answer came from Galileo. In *Two New Sciences*, published in Leiden in 1638 because his work could no longer be printed in Italy, he combined two results of his own: a body moving horizontally keeps its speed unless something acts on it, and a falling body covers distances that grow as the square of the time. Put together, a steady sideways motion and a square-law fall trace out a parabola. From that he proved that 45° gives the greatest range, and that angles equally far above and below 45° give equal ranges, a fact gunners had noticed in practice without knowing why.
:::

::: quiz
? At the top of its path, what is a projectile's velocity?
- Zero
+ Horizontal, equal to v₀ cos θ
- Vertical, equal to v₀ sin θ
= At the top the vertical component is momentarily zero, but the horizontal component never changes during the flight. So the projectile is still moving sideways at its launch horizontal speed, v₀ cos θ.
? A ball launched at 20° lands 30 m away on level ground. At what other angle, with the same speed, would it land at the same spot?
- 40°
- 45°
+ 70°
= Because sin(2θ) = sin(180° − 2θ), the angles θ and 90° − θ give the same range. The partner of 20° is 90° − 20° = 70°, a higher and slower path to the same point.
? Doubling the launch speed, with the same angle, changes the range by what factor?
- 2
+ 4
- √2
= The range is R = v₀² sin(2θ)/g, so it is proportional to the square of the launch speed. Doubling the speed doubles both the flight time and the horizontal speed, which together multiply the range by four.
:::
