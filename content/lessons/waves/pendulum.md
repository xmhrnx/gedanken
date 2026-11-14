A weight on a string, swinging under gravity, kept the world's time for nearly three centuries. Its usefulness rests on two facts that sound too simple to be true. For small swings, the time a pendulum takes to swing back and forth does not depend on how far it swings, and it does not depend on how heavy the weight is. It depends only on the length of the string and the strength of gravity. This lesson shows where those facts come from, why the second one turned out to be a clue to something deep about gravity, and how the pendulum went from a curiosity in a cathedral to the most accurate clock in the world and a tool for measuring the shape of the Earth.

## What pulls a pendulum back?

The previous lesson showed that simple harmonic motion, the smooth back-and-forth motion with a fixed period, arises whenever the force pulling an object back toward equilibrium is proportional to its displacement. A pendulum has no spring, so the question is whether gravity can play the spring's role, and if so, what plays the role of the spring constant.

Take a small, heavy bob of mass m hanging on a light string of length L, and pull it aside so that the string makes an angle θ with the vertical. Two forces act on the bob: the tension of the string, which pulls along the string toward the pivot, and gravity, mg, which pulls straight down. The bob is constrained to move along a circular arc around the pivot, so what matters is the part of each force that points along that arc.

The tension is always at right angles to the arc, so it does nothing to speed up or slow down the bob; it only keeps it on the circle. Gravity, though, can be split into two parts. One part points along the string, away from the pivot, and is balanced by the tension. The other part points along the arc, back toward the lowest point, and has size mg sin θ. That second part is the restoring force. When the bob hangs straight down, θ = 0 and the restoring force vanishes; the farther it is pulled aside, the larger the restoring force becomes.

## Finding the period

The restoring force, mg sin θ, is not quite proportional to the displacement, because of the sine. But for small angles it very nearly is, and that is enough to give simple harmonic motion.

::: steps The period of a pendulum
$ m\,a = -mg\sin\theta
This is Newton's second law applied along the arc: mass times acceleration along the arc equals the force along the arc. The minus sign says the force points back toward the lowest point, opposite to the direction in which the bob has been displaced. The tension does not appear because it has no component along the arc.
---
$ a = -g\sin\theta
Divide both sides by m. The mass cancels, because it appears on both sides: on the left as the inertia that resists acceleration, on the right as the quantity gravity pulls on. This is why the period will not depend on the bob's mass, and the lesson returns to this cancellation below.
---
$ \sin\theta \approx \theta
For small angles measured in radians, sin θ is very nearly equal to θ itself. The reason is that on a circle of radius 1, sin θ is the height of a point and θ is the length of arc to it, and for a short arc the two are almost the same. At 0.1 radians, about 6°, they differ by less than 0.2%. Making this approximation is what turns the pendulum into a simple harmonic oscillator.
---
$ s = L\theta \quad\Rightarrow\quad a = -\frac{g}{L}\,s
The distance of the bob from the lowest point, measured along the arc, is s = Lθ, so θ = s/L. Substituting this into a = −gθ expresses the acceleration in terms of the distance along the path. The acceleration is now proportional to the displacement and points the opposite way, which is the defining condition for simple harmonic motion.
---
$ \omega^2 = \frac{g}{L}
Compare with a mass on a spring, where a = −(k/m)x and we wrote ω² = k/m. The pendulum has exactly the same form with g/L in place of k/m. So everything we learned about the spring applies, including the formula for the period, with this new value of ω².
---
$ T = \frac{2\pi}{\omega} = 2\pi\sqrt{\frac{L}{g}}
The period of simple harmonic motion is 2π/ω. With ω = √(g/L), the period is 2π√(L/g). Only the length of the string and the strength of gravity remain. The mass has cancelled, and so has the amplitude, which never appeared at all once the small-angle approximation was made.
:::

::: model pendulum
Change the length and watch the period. Making the pendulum four times as long doubles its period, because the period goes as the square root of the length.
:::

Some numbers make the formula concrete. A pendulum 1 m long has a period of about 2.0 s, so each one-way swing takes about a second. A pendulum built so that each one-way swing takes exactly one second, called a "seconds pendulum", needs to be 0.994 m long under standard gravity, which is why tall grandfather clocks are about that size inside. A playground swing a few metres long takes three or four seconds for a full back-and-forth, and that is true whoever is sitting on it, child or adult.

::: note
For small swings a pendulum's period is 2π√(L/g), set only by its length and by gravity. The mass cancels because the same mass determines both how hard gravity pulls and how strongly the bob resists being accelerated.
:::

## Why does the mass cancel?

The cancellation of the mass may look like an algebraic convenience, but it hides an assumption that turned out to be one of the deepest questions in physics. Mass plays two quite different roles in the derivation. In F = ma it measures **inertia**, how strongly an object resists being accelerated; this is called **inertial mass**. In the force of gravity, mg, it measures how strongly gravity pulls on the object; this is called **gravitational mass**. Nothing in Newton's laws requires these two to be the same. One could imagine a world in which gold was pulled by gravity a little more strongly, relative to its inertia, than wood. In that world, a gold pendulum and a wooden pendulum of the same length would have different periods.

Newton saw this, and tested it. He hung equal-weight boxes from strings of equal length, filled them with gold, silver, lead, glass, sand, salt, wood, water and wheat, and compared their swings. The periods matched to about one part in a thousand, so for all those materials the two kinds of mass were equal to that precision. Modern experiments, using torsion balances and satellites in orbit, confirm the equality to better than one part in 10¹⁴. Einstein took this equality not as a coincidence but as a clue, and made it the starting point of general relativity: if gravity accelerates everything in exactly the same way, perhaps gravity is not a force acting on objects at all but a property of space and time.

## What happens on large swings?

The approximation sin θ ≈ θ is excellent for small swings but fails for big ones. For larger angles sin θ is noticeably smaller than θ, so the restoring force is weaker than the simple formula assumes, the bob is turned round more gently, and the period grows. A useful correction for moderate amplitudes θ₀ (measured in radians) is T ≈ T₀(1 + θ₀²/16), where T₀ is the small-angle period 2π√(L/g). At an amplitude of 20°, the period is about 0.8% longer than T₀. At 90°, where the bob starts level with the pivot, the exact period is about 18% longer; the correction formula, meant for moderate angles, slightly underestimates this.

A fraction of a percent may sound negligible, but for a clock it is large. A day has 86,400 seconds, and 0.8% of that is about 11 minutes. So a clock adjusted to keep time with a 20° swing would run about 11 minutes a day fast if its swing shrank to a few degrees, as it might when the mechanism became dirty or the driving weight ran low. Clockmakers therefore designed pendulums to swing through only a few degrees, where the correction is tiny and changes very little even if the swing varies.

::: history 1583–1673 · Pisa, The Hague and Cayenne
According to his first biographer, the young Galileo, sitting in Pisa cathedral, timed a swinging lamp against his own pulse and noticed that wide and narrow swings took the same time. The story was written down decades later and may have been polished in the telling, and Galileo went too far with the idea, claiming that even very large swings take the same time, which is not quite true. But the insight was real, and he later proposed using a pendulum to regulate a clock. It was Christiaan Huygens, in The Hague, who built the first working pendulum clock, in 1656. Before it, a very good mechanical clock might drift by a quarter of an hour a day; Huygens's clocks kept time to within perhaps fifteen seconds. Huygens also worried about the large-swing problem and solved it in principle: he proved that a bob guided along a curve called a cycloid, rather than a circle, takes exactly the same time to swing whatever the amplitude. Then came a surprise from far away. In 1672 the French astronomer Jean Richer took a carefully regulated pendulum clock to Cayenne in French Guiana, near the equator, and found that it lost about two and a half minutes a day. Gravity there was slightly weaker. Newton and Huygens explained why: Earth's rotation partly offsets gravity at the equator, and Earth bulges there, so the equator is farther from the centre. It was the first evidence that Earth is not a perfect sphere.
:::

## A pendulum as a gravity meter

Richer's discovery points to another use for the formula. Turn it around, g = 4π²L/T², and the pendulum becomes a way to measure gravity. Measure the length of the string with a ruler, time a hundred swings with a stopwatch to reduce the effect of your reaction time, and you can find g to a fraction of a percent in an afternoon.

For about two centuries, refined versions of this method were the main way to map gravity across the Earth. Expeditions carried precision pendulums to the equator, to high latitudes and up mountains. Their measurements showed how gravity grows from the equator to the poles, which confirmed and measured Earth's flattening. They also revealed small local variations caused by dense or light rock beneath the surface, which later became a tool for finding mineral deposits. In the twentieth century, instruments that time a freely falling mass, or measure the stretch of an extremely sensitive spring, replaced the pendulum. But the idea is the one in this lesson: a simple oscillator whose period depends on gravity, and on almost nothing else.

::: quiz
? Two pendulums of equal length have bobs of 1 kg and 5 kg. For small swings, which has the longer period?
- The 5 kg one
- The 1 kg one
+ Neither: the periods are equal
= The mass cancels, because it sets both the pull of gravity and the resistance to acceleration. The period 2π√(L/g) depends only on length and gravity.
? To double a pendulum's period, you should make it
- twice as long
+ four times as long
- half as long
= The period goes as the square root of the length. To double the square root, you have to multiply the length by four.
? A pendulum clock is taken to the Moon, where gravity is about one sixth of Earth's. Its pendulum swings
- six times faster
+ about 2.4 times more slowly
- at the same rate
= The period goes as 1/√g. Gravity is six times weaker, so the period is longer by √6, which is about 2.45.
:::
