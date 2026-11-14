A weight on a string, swinging under gravity, kept the world's time for nearly three centuries. Its usefulness rests on two facts that seem too simple to be true: for small swings, the period does not depend on how far it swings, nor on how heavy the weight is. It depends only on the length of the string and the strength of gravity.

## Forces on a pendulum

Take a small bob of mass m on a light string of length L, pulled aside by an angle θ. Two forces act: the string's tension along the string, and gravity mg straight down. Split gravity into a part along the string, which the tension balances, and a part along the bob's arc, mg sin θ, pointing back toward the bottom. That second part is the restoring force.

::: steps The period of a pendulum
$ m\,a = -mg\sin\theta
Newton's second law along the arc. The minus sign means the force points back toward the lowest point.
---
$ a = -g\sin\theta
The mass cancels. It appears once as the inertia resisting acceleration and once in the weight providing the force.
---
$ \sin\theta \approx \theta
For small angles in radians, sin θ is very nearly θ itself. At 0.1 rad (about 6°) the difference is under 0.2%.
---
$ s = L\theta \quad\Rightarrow\quad a = -\frac{g}{L}\,s
The distance along the arc is s = Lθ, so θ = s/L. The acceleration is now proportional to the displacement and opposite to it.
---
$ \omega^2 = \frac{g}{L}
This is simple harmonic motion, a = −ω²s, as in the previous lesson, with ω² = g/L in place of k/m.
---
$ T = 2\pi\sqrt{\frac{L}{g}}
The period is 2π/ω. Only length and gravity remain.
:::

::: model pendulum
Change the length and watch the period. Four times the length doubles the period.
:::

A pendulum 1 m long has a period of about 2.0 s, so each one-way swing takes about a second. A "seconds pendulum", with exactly one second per swing, is 0.994 m long at standard gravity. A playground swing a few metres long takes three or four seconds per round trip, whoever is sitting on it.

::: note
For small swings a pendulum's period is 2π√(L/g). The mass cancels because the same mass sets both the pull of gravity and the resistance to being accelerated.
:::

## Why the mass cancels

The cancellation is not a minor detail. Mass plays two roles here: **inertial mass**, which resists acceleration in F = ma, and **gravitational mass**, which sets how hard gravity pulls. Nothing in Newton's laws says they must be equal. Newton tested it by swinging pendulums with bobs of gold, silver, lead, glass, sand, salt, wood, water and wheat, and found equal periods to about one part in a thousand. Modern experiments confirm the equality to better than one part in 10¹⁴. Einstein later made it the starting point of general relativity.

## Large swings

The approximation sin θ ≈ θ fails for big swings, and the period grows. A good correction for moderate amplitudes θ₀ is T ≈ T₀(1 + θ₀²/16), with θ₀ in radians. At a 20° amplitude the period is about 0.8% longer than the small-angle value; at 90° it is 18% longer. For a clock this matters: a clock set right with a 20° swing would run about 11 minutes a day fast if its swing shrank to a few degrees. Clock pendulums therefore swing only a few degrees, where the correction is tiny and barely changes as the swing varies.

::: history 1583–1673 · Pisa, The Hague and Cayenne
According to his first biographer, the young Galileo timed a swinging lamp in Pisa cathedral against his pulse and found that wide and narrow swings took the same time. The story may be embellished, and Galileo overstated the result, claiming equal periods even for very large swings. In 1656 Christiaan Huygens built the first pendulum clock, cutting the error of good clocks from perhaps fifteen minutes a day to about fifteen seconds. He also found that a bob following a cycloid, rather than a circle, swings in exactly the same time at any amplitude. In 1672 Jean Richer took a pendulum clock to Cayenne, near the equator, and found it lost about two and a half minutes a day: gravity there is slightly weaker, the first evidence that Earth is not a perfect sphere.
:::

## Measuring gravity

Turn the formula around, g = 4π²L/T², and a pendulum becomes a gravity meter. Timing 100 swings with a stopwatch and measuring the string with a ruler gives g to a fraction of a percent. Refined pendulums mapped variations in g across the Earth for two centuries, revealing its equatorial bulge and dense rocks beneath the surface, before falling-mass and spring instruments replaced them.

::: quiz
? Two pendulums of equal length have bobs of 1 kg and 5 kg. For small swings, which has the longer period?
- The 5 kg one
- The 1 kg one
+ Neither: the periods are equal
= Mass cancels: T = 2π√(L/g) depends only on length and gravity.
? To double a pendulum's period, you should make it
- twice as long
+ four times as long
- half as long
= T ∝ √L, so doubling T needs four times the length.
? A pendulum clock is taken to the Moon, where gravity is about one sixth of Earth's. Its pendulum swings
- six times faster
+ about 2.4 times more slowly
- at the same rate
= T ∝ 1/√g, and √6 ≈ 2.45, so each swing takes about 2.4 times longer.
:::
