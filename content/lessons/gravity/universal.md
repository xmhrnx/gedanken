Kepler's laws describe the planets; they do not explain them. Newton's step was to propose that one force, the same one that drops an apple, holds the Moon in orbit and the planets around the Sun, and to work out exactly how that force must weaken with distance.

## From Kepler's third law to the inverse square

A body moving in a circle needs a centripetal acceleration (an acceleration toward the centre) of v²/r, as the circular motion lesson showed. Most planetary orbits are close to circles, so treat them as circles and ask what acceleration the Sun must supply.

::: steps The inverse-square law from Kepler's third law
$ a = \frac{v^2}{r}
Centripetal acceleration for a circular orbit of radius r at speed v.
---
$ v = \frac{2\pi r}{T}
In one period T the planet covers one circumference.
---
$ a = \frac{4\pi^2 r}{T^2}
Substitute v into a = v²/r.
---
$ T^2 = Cr^3
Kepler's third law: T² is proportional to r³, with the same constant C for every planet.
---
$ a = \frac{4\pi^2}{C} \cdot \frac{1}{r^2}
Substitute T² = Cr³. The acceleration toward the Sun falls off as the inverse square of the distance, and the factor in front is the same for every planet.
:::

So the Sun pulls each planet with an acceleration ∝ 1/r², whatever the planet. The force is mass times acceleration, so the force on a planet is proportional to its mass m. By Newton's third law the planet pulls the Sun equally hard, so the force should be proportional to the Sun's mass M as well. Put together:

::: equation Newton's law of universal gravitation
F = \frac{GMm}{r^2}
:::

G is a universal constant, now known to be 6.674 × 10⁻¹¹ N m² kg⁻². Its smallness is why you feel no pull from a nearby building. Measuring it took another century, as a later lesson describes.

## The Moon test

The bold claim is *universal*: the force that keeps the Moon in orbit is ordinary terrestrial gravity, weakened by distance. That can be checked with numbers available in Newton's time. The Moon is about 60 Earth radii away (384,400 km against 6,371 km). Gravity at Earth's surface, one Earth radius from the centre, gives g = 9.81 m/s². If gravity is inverse-square, at 60 times the distance it should be 60² = 3,600 times weaker.

::: steps Newton's Moon test
$ a_{\text{pred}} = \frac{g}{60^2} = \frac{9.81}{3600} \approx 2.72 \times 10^{-3}\,\text{m/s}^2
The acceleration the inverse-square law predicts at the Moon's distance.
---
$ r = 3.844 \times 10^8\,\text{m}, \quad T = 27.32\,\text{days} = 2.361 \times 10^6\,\text{s}
The Moon's orbital radius and its period relative to the stars (the sidereal month).
---
$ a_{\text{Moon}} = \frac{4\pi^2 r}{T^2} = \frac{4\pi^2 \times 3.844 \times 10^8}{(2.361 \times 10^6)^2}
The Moon's actual centripetal acceleration, from its observed orbit alone.
---
$ a_{\text{Moon}} \approx 2.72 \times 10^{-3}\,\text{m/s}^2
The two agree to better than 1%. In one second the Moon falls about 1.4 mm away from a straight line toward Earth, exactly as much as an apple 60 Earth radii up would fall.
:::

::: note
One force law, F = GMm/r², explains a falling apple, the Moon's orbit and Kepler's laws. Gravity on Earth and gravity in the heavens are the same thing.
:::

::: history 1666–1687 · Woolsthorpe and Cambridge
Newton later said that in 1666, at home in Lincolnshire while plague closed Cambridge, he compared the force holding the Moon with gravity at Earth's surface and "found them answer pretty nearly". His early numbers matched only roughly, perhaps because he used a poor value for Earth's size, and he set the problem aside. In 1684 Edmond Halley asked him what orbit an inverse-square force would produce; Newton replied that he had calculated it: an ellipse. Halley pressed him to publish, and the result was the *Principia* (1687). Robert Hooke had suggested an inverse-square attraction in letters to Newton and claimed credit; Newton, furious, cut references to him from the book. The apple story comes from Newton himself, told decades later in his old age.
:::

## Measuring from the centre

The Moon test quietly assumes that Earth pulls as if all its mass sat at its centre, so that an apple on the surface is "one Earth radius" away. That is not obvious: the ground under your feet is much closer than the far side of the planet. Newton proved it: a spherically symmetric body attracts objects outside it exactly as a point mass at its centre would. This **shell theorem** is part of why the *Principia* took so long. It gives gravity at the surface of any round planet:

::: equation Surface gravity of a planet of mass M and radius R
g = \frac{GM}{R^2}
:::

Newton also showed that an inverse-square force produces exactly Kepler's ellipses and his equal-area law, and that the planets' small tugs on one another explain the departures from Kepler that astronomers measured. In 1846 the same arithmetic, applied to wobbles in the orbit of Uranus, located an unseen planet: Neptune was found within about a degree of the predicted position.

Newton's law is not the last word. It does not say how the pull is transmitted, a gap Newton himself admitted. General relativity replaces it, but in weak gravity the two agree so closely that spacecraft are still navigated with Newton's formula.

::: quiz
? If you moved to twice your current distance from Earth's centre, your weight would be
- half as large
+ a quarter as large
- unchanged
= F ∝ 1/r²: doubling r divides the force by 4.
? What does the Moon test compare?
- The Moon's mass with Earth's mass
+ The Moon's centripetal acceleration with g reduced by 60²
- The time light takes to reach the Moon with the length of a month
= The Moon is about 60 Earth radii away, so inverse-square gravity predicts g/3600 there, which matches 4π²r/T² for the Moon's orbit.
? Kepler's third law (T² ∝ r³) implies the Sun's pull on a planet goes as
- 1/r
+ 1/r²
- 1/r³
= a = 4π²r/T² and T² ∝ r³ give a ∝ r/r³ = 1/r².
:::
