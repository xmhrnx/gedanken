In 1905 Albert Einstein, a 26-year-old clerk at the Swiss patent office in Bern, resolved the conflict between Newton's mechanics and Maxwell's electromagnetism by taking two simple statements seriously and following them wherever they led. Each statement on its own was widely accepted. Together they seem to contradict each other, and the contradiction forced a new understanding of space and time.

## Inertial frames

A **frame of reference** is a point of view from which to measure positions and times: a lab, a train, a spacecraft, each with its own rulers and clocks. An **inertial frame** is one that is not accelerating, where an object with no forces on it moves in a straight line at constant speed. A smoothly moving train is (very nearly) one; a braking train is not.

## Postulate 1: the principle of relativity

**The laws of physics are the same in every inertial frame.**

No experiment done inside a closed lab can tell you whether the lab is at rest or moving at constant velocity. Galileo made this point in 1632 with a ship: below decks, butterflies fly, fish swim and drops fall just the same whether the ship is moored or sailing smoothly. Only relative motion has meaning. Einstein extended this from mechanics to all of physics, including electromagnetism and light.

## Postulate 2: the constancy of the speed of light

**Light in vacuum travels at the same speed, c ≈ 3.00 × 10⁸ m/s, in every inertial frame.**

Einstein stated it as: light in empty space travels at a definite speed c, whatever the motion of the body emitting it. Combined with the first postulate, this means every inertial observer measures the same c, whatever their own motion. Maxwell's equations predict a single speed c; the first postulate says those equations hold in every inertial frame; so every frame finds the same c.

## The conflict

Everyday experience says velocities add. Walk forward at 1 m/s on a train moving at 30 m/s, and you move at 31 m/s relative to the ground. In Newton's physics this is exact:

::: equation Galilean velocity addition
u = u' + v
:::

where u′ is your velocity measured on the train and v is the train's velocity. Apply this to a light beam sent forward from the train, u′ = c, and the ground observer should see c + v. The second postulate says the ground observer also sees c. Both cannot be true unless something in our ideas of distance and time is wrong. Einstein concluded that the culprit is the assumption that time and distance are the same for every observer.

## What replaces it

Working through the consequences, which the following lessons do, Einstein found that velocities combine according to a different rule. It is derived from the Lorentz transformation, the relativistic rule for converting positions and times between frames; here we only check that it does what it must.

::: steps Checking the relativistic addition rule
$ u = \frac{u' + v}{1 + u'v/c^2}
The relativistic velocity addition formula, for motion along one line.
---
$ u \approx u' + v \quad \text{when } u', v \ll c
For everyday speeds u′v/c² is tiny: for walking on a train, about 3 × 10⁻¹⁶. The denominator is 1 to great precision and the rule reduces to Newton's.
---
$ u = \frac{c + v}{1 + cv/c^2} = \frac{c + v}{1 + v/c}
Now send a light beam forward from the train: set u′ = c.
---
$ u = \frac{c + v}{(c + v)/c} = c
Write 1 + v/c as (c + v)/c and cancel. The ground observer measures exactly c, as the second postulate requires, whatever v is.
:::

Adding 0.5c to 0.5c gives (0.5 + 0.5)/(1 + 0.25) = 0.8c, not c. No combination of speeds below c ever reaches c.

The consequences come one after another, each from the two postulates with nothing else assumed: moving clocks run slow, moving objects are shorter, events simultaneous for one observer are not for another, and mass is a form of energy. Each sounds absurd, and each has been measured.

::: note
Two postulates: physics is the same in every inertial frame, and light's speed is c in every inertial frame. Together they force time and distance to depend on the observer.
:::

::: history 1895–1905 · Aarau and Bern
At 16, Einstein wondered what he would see if he chased a light beam at speed c. He would see an electromagnetic wave frozen in place, a pattern of fields that Maxwell's equations do not allow. Years later his 1905 paper "On the Electrodynamics of Moving Bodies" opened with a different puzzle: a magnet moving past a wire and a wire moving past a magnet produce the same current, yet the theory of the time explained the two cases in completely different ways. Einstein took this as a sign that only relative motion matters. Lorentz and Henri Poincaré had already found much of the mathematics, but treated it as describing how matter behaved in the ether. Einstein made it a statement about space and time themselves and called the ether "superfluous".
:::

The second postulate has been tested directly. In 1964 at CERN, gamma rays emitted by particles moving at 99.975% of the speed of light were timed and found to travel at c, not at nearly 2c, to within about one part in ten thousand.

::: quiz
? A spaceship moving at 0.9c relative to Earth fires a laser forward. How fast does an Earth observer measure the laser light to move?
- 1.9c
- 0.1c
+ c
= The second postulate: light's speed is c in every inertial frame, regardless of the source's motion.
? According to the first postulate, what can a passenger in a windowless train moving smoothly at constant velocity determine by experiments inside?
+ Nothing about the train's speed
- The speed, using a pendulum
- The speed, using light beams
= Physics is identical in every inertial frame, so no internal experiment distinguishes constant motion from rest.
? Two spaceships approach Earth from opposite sides, each at 0.5c relative to Earth. How fast does one see the other approach?
- c
+ 0.8c
- 0.5c
= Using u = (u′ + v)/(1 + u′v/c²): (0.5 + 0.5)c/(1 + 0.25) = 0.8c.
:::
