Kepler's laws describe how the planets move, but they do not say why. By the 1680s several people in London suspected that the Sun pulls on the planets with a force that weakens with distance, but nobody could show that such a force would produce Kepler's ellipses, or say exactly how fast it should weaken. Isaac Newton did both, and went further. He proposed that the force holding the planets to the Sun is the same force that makes an apple fall, and he tested that claim with numbers anyone could check. This lesson follows his reasoning: first from Kepler's third law to the inverse-square law, then to the test that showed gravity on Earth and gravity in the sky are one thing.

## What does a curved path require?

Start with what an orbit needs. Newton's first law says that a body left alone moves in a straight line at steady speed. A planet does not do that: it curves around the Sun. So something must be acting on it, and the question is what kind of force, and how strong.

For motion in a circle the answer was already known. A body moving at speed v around a circle of radius r is constantly changing direction, and the acceleration needed to do that points toward the centre and has size v²/r. This is called **centripetal acceleration**, from the Latin for "centre-seeking", and it was worked out in the lesson on circular motion. Most planetary orbits are close to circles (Earth's eccentricity is 0.017), so as a first approximation we can treat them as circles and ask what acceleration the Sun must supply to each planet. Kepler's third law then tells us how that acceleration must change with distance.

::: steps The inverse-square law from Kepler's third law
$ a = \frac{v^2}{r}
This is the centripetal acceleration of a planet moving at speed v on a circular orbit of radius r around the Sun. Whatever the Sun does to the planet, it must produce exactly this acceleration, pointing inward, or the planet would not stay on the circle. The trouble is that v is not one of the things Kepler's law mentions, so we want to rewrite it.
---
$ v = \frac{2\pi r}{T}
In one period T the planet travels once around the circle, a distance 2πr. Since it moves at steady speed on a circle, its speed is this distance divided by the time. This step expresses the speed through the period and the radius, the two quantities Kepler's third law connects.
---
$ a = \frac{4\pi^2 r}{T^2}
Put this expression for v into a = v²/r. Squaring 2πr/T gives 4π²r²/T², and dividing by r leaves one power of r on top. The acceleration is now written entirely in terms of the orbit's radius and its period.
---
$ T^2 = Cr^3
Kepler's third law says T² is proportional to r³ (for a circle, the semi-major axis is just the radius). C is the constant of proportionality, and the important point is that it is the same for every planet. This is the observed fact that will decide how the force behaves.
---
$ a = \frac{4\pi^2}{C} \cdot \frac{1}{r^2}
Replace T² in the previous expression with Cr³. On top there is one power of r, below there are three, so two powers remain in the denominator. The acceleration toward the Sun falls off as the inverse square of the distance: twice as far, a quarter as much. And the factor in front, 4π²/C, is the same for every planet, so the acceleration depends only on where a planet is, not on which planet it is.
:::

That last point deserves a pause. Mercury and Jupiter are very different bodies, yet if you placed either of them at the same distance from the Sun, it would accelerate toward the Sun at the same rate. That is exactly how bodies fall on Earth: a heavy stone and a light one fall together, as Galileo had argued. The Sun's pull behaves like ordinary falling.

## From an acceleration to a force

An acceleration that is the same for every planet means a force that is proportional to each planet's mass, because force equals mass times acceleration. A planet twice as massive needs twice the force to be given the same acceleration, and gets it. So the Sun's pull on a planet of mass m goes as m/r².

Newton then turned the argument around using his third law, which says that forces come in equal and opposite pairs. If the Sun pulls the planet, the planet pulls the Sun just as hard. By the same reasoning that made the force proportional to the planet's mass, it should also be proportional to the Sun's mass M: neither body is special. Putting the two together gives a force proportional to the product of the masses and inversely proportional to the square of the distance between them:

::: equation Newton's law of universal gravitation
F = \frac{GMm}{r^2}
:::

G is a constant of nature, the same everywhere, which sets how strong gravity is. Its modern value is 6.674 × 10⁻¹¹ N m² kg⁻². That is an extraordinarily small number. Two people standing a metre apart attract each other with a force of a few ten-millionths of a newton, far too small to notice, and you feel no pull at all from a large building next to you. Gravity seems strong only because Earth is enormous. Newton could not measure G, and nobody did for more than a century, as a later lesson in this chapter describes.

## Is it really the same force? The Moon test

The word "universal" makes a bold claim: that the force keeping the Moon in its orbit is the same gravity that pulls an apple to the ground, merely weakened by distance. Before Newton, most people assumed the heavens and the Earth followed different rules. The claim could be tested, because every number needed was already known.

The Moon is about 60 Earth radii from Earth's centre: about 384,400 km, against an Earth radius of 6,371 km. An apple at the surface is one Earth radius from the centre and falls with acceleration g = 9.81 m/s². If gravity weakens as the inverse square of distance, then at 60 times the distance it should be 60² = 3,600 times weaker. Meanwhile, the Moon's actual acceleration can be worked out from its orbit alone, with no theory of gravity at all, just from how far away it is and how long it takes to go round. If the two numbers agree, the Moon is falling exactly as an apple would at that height.

::: steps Newton's Moon test
$ a_{\text{pred}} = \frac{g}{60^2} = \frac{9.81}{3600} \approx 2.72 \times 10^{-3}\,\text{m/s}^2
This is the prediction. Take the measured acceleration of falling bodies at Earth's surface and divide it by 60², because the inverse-square law says the pull is 3,600 times weaker at 60 times the distance. It assumes that Earth pulls as if its mass were concentrated at its centre, a point we return to below.
---
$ r = 3.844 \times 10^8\,\text{m}, \quad T = 27.32\,\text{days} = 2.361 \times 10^6\,\text{s}
These are the Moon's measured orbital radius and its period. The period used is the **sidereal month**, the time to return to the same place among the stars, rather than the 29.5 days between full Moons, which is longer because Earth has moved along its own orbit in the meantime. Converting days to seconds keeps the units consistent with g.
---
$ a_{\text{Moon}} = \frac{4\pi^2 r}{T^2} = \frac{4\pi^2 \times 3.844 \times 10^8}{(2.361 \times 10^6)^2}
This is the Moon's actual centripetal acceleration, from the formula derived above for any circular orbit. Nothing about gravity goes into it: it is pure geometry and timing, the acceleration anything must have to move round that circle in that time.
---
$ a_{\text{Moon}} \approx 2.72 \times 10^{-3}\,\text{m/s}^2
Working out the numbers gives the same value as the prediction, to better than 1%. The Moon is being pulled toward Earth exactly as hard as inverse-square gravity, starting from the apple, says it should be.
:::

It helps to picture what this number means. In one second, the Moon falls about 1.4 mm away from the straight line it would otherwise follow. That is a tiny drop compared with the kilometre or so it travels sideways in the same second, which is why its path curves so gently. It is also exactly how far an apple would fall in its first second if you could hold it 60 Earth radii up and let go. The Moon is falling all the time; it simply moves sideways fast enough that it keeps missing.

::: history 1666–1687 · Woolsthorpe and Cambridge
In 1665 plague closed the University of Cambridge, and the young Newton went home to his family's farm at Woolsthorpe in Lincolnshire. There, he later said, he compared the force needed to hold the Moon in its orbit with the force of gravity at Earth's surface and "found them answer pretty nearly". The match was rough, possibly because he used a poor value for Earth's size, and he left the problem for years. Meanwhile Robert Hooke, in letters to Newton around 1680, suggested that the planets might be held by an attraction falling off as the inverse square of distance, but he could not work out what orbit it would produce. In 1684 Edmond Halley visited Newton and asked exactly that question. Newton answered at once that he had calculated it: an ellipse. He could not find the calculation, so he redid it, and the work grew over the next two and a half years into the *Principia*, published in 1687 at Halley's expense. Hooke claimed credit for the inverse square; Newton, furious, removed references to him. The story of the falling apple comes from Newton himself, told in conversation near the end of his life, and he presented it as the moment the question occurred to him, not the answer.
:::

## Why measure from Earth's centre?

The Moon test quietly assumed something that is not at all obvious. It treated the apple as being "one Earth radius" from Earth, as if all of Earth's mass were concentrated at its centre. But Earth is not a point. The ground under the apple is a metre away, while the far side of the planet is nearly 13,000 km away, and every piece pulls with a different strength in a different direction. Why should all those pulls add up to the pull of a single point mass at the centre?

Newton proved that they do. A body whose mass is arranged in spherically symmetric layers, the same in every direction from the centre, attracts anything outside it exactly as if all its mass were at the centre. This result is called the **shell theorem**, and it depends on the force being inverse-square: for most other force laws it would fail. Proving it was one reason the *Principia* took as long as it did. It justifies the Moon test, and it gives a simple formula for gravity at the surface of any round planet of mass M and radius R:

::: equation Surface gravity of a planet of mass M and radius R
g = \frac{GM}{R^2}
:::

The formula says something worth noticing: the mass of the falling object is nowhere in it. Every object at a planet's surface falls with the same acceleration, which is Galileo's result again, now explained. A larger mass is pulled harder, in exact proportion, and resists acceleration more, in exact proportion, so the two effects cancel.

## Worked examples: the Moon, the Sun and two people

The law is most useful when it is put to work on numbers, and three examples show three different things it can do: predict, weigh and compare.

First, prediction. The Moon has a mass of about 7.35 × 10²² kg and a radius of about 1,737 km. Putting these into g = GM/R² gives 6.674 × 10⁻¹¹ × 7.35 × 10²² divided by (1.737 × 10⁶)², which is about 1.62 m/s², roughly one sixth of Earth's surface gravity. That is the value the Apollo astronauts experienced, and it is why they could bound across the surface with heavy backpacks. Notice how the two factors work against each other. The Moon has about 1/81 of Earth's mass, which on its own would make its gravity 81 times weaker, but its radius is only about 0.27 of Earth's, and being that much closer to the centre multiplies the pull by about 13. Mass and size both matter, and neither alone tells you how heavy you will feel.

Second, weighing. Turn the circular-orbit argument around: if the Sun's pull GMm/r² supplies a planet's centripetal force m·4π²r/T², then the Sun's mass is M = 4π²r³/(GT²). Earth's orbit has r = 1.496 × 10¹¹ m and T = 1 year = 3.156 × 10⁷ s. Working this out gives M ≈ 2.0 × 10³⁰ kg, about 330,000 times Earth's mass. The planet's own mass has cancelled, so a single orbit, of anything, weighs the body at the centre. This is how the masses of the Sun, of Jupiter from its moons, and of distant stars from their companions are found. There is one catch: you need G, and until G was measured, as described in Weighing the Earth, astronomers knew only the product GM, not M itself.

::: steps Weighing the Sun with Earth's orbit
$ \frac{GMm}{r^2} = \frac{4\pi^2 m r}{T^2}
Gravity is the only force on Earth in its orbit, so it must supply exactly the centripetal force a circular orbit needs. On the left is the Sun's pull; on the right is Earth's mass times the centripetal acceleration derived earlier.
---
$ M = \frac{4\pi^2 r^3}{G T^2}
Earth's mass m appears on both sides and cancels. Multiplying both sides by r² and dividing by G leaves the Sun's mass in terms of the orbit's radius and period and the constant G. This is Kepler's third law with its constant explained: C = 4π²/(GM).
---
$ M = \frac{4\pi^2 (1.496 \times 10^{11})^3}{6.674 \times 10^{-11} \times (3.156 \times 10^7)^2} \\ \approx 1.99 \times 10^{30}\,\text{kg}
With the numbers in, the top is about 1.32 × 10³⁵ and the bottom about 6.65 × 10⁴, and their ratio is the mass of the Sun. The whole calculation needs only a distance, a time and G.
:::

Third, comparison. Two people of 70 kg each, standing 1 m apart, attract each other with F = 6.674 × 10⁻¹¹ × 70 × 70 / 1², about 3 × 10⁻⁷ N, roughly the weight of a grain of fine sand. Earth and Moon, by contrast, pull on each other with about 2 × 10²⁰ N. Both come from the same formula. Gravity is the weakest force we deal with, yet it dominates astronomy, because it adds up: every kilogram attracts every other, there is no negative mass to cancel it, and planets and stars contain enormous numbers of kilograms.

## Going deeper: gravity inside the Earth

The shell theorem has a second half, and it leads somewhere surprising. Newton also proved that a uniform spherical shell exerts no net force at all on anything inside it. The pulls from the near part of the shell, which is close but small, exactly cancel those from the far part, which is distant but large. That is another result that works only for an inverse-square force: the area of a patch of shell seen at a given angle grows as the square of its distance, and the pull from each piece falls as the square, so the two match.

This tells us how gravity behaves if you go down a deep shaft. Everything above you is a shell, and it pulls you nowhere. Only the sphere of material below you counts.

::: steps Gravity at depth in a uniform planet
$ M(r) = M \frac{r^3}{R^3}
Suppose, for simplicity, that the planet has the same density all the way through. A ball of radius r inside it then holds a fraction r³/R³ of the total mass, because volume goes as the cube of the radius. The rest of the planet, the shell outside radius r, contributes nothing.
---
$ g(r) = \frac{G M(r)}{r^2} = \frac{GM}{R^3}\, r
Apply the outside form of the law to the inner ball alone. One power of r³ on top and r² below leave a single r: inside a uniform planet, gravity grows in direct proportion to the distance from the centre. It is zero at the centre, where everything pulls equally in all directions, and reaches its full value g at the surface.
:::

A force that grows in proportion to the distance from a centre, pulling back toward it, is exactly the kind that produces simple harmonic motion, as the lesson Simple Harmonic Motion explains. Drop a stone down an imaginary airless tunnel straight through a uniform Earth and it would fall through the centre, rise to the far side and fall back, oscillating forever. Each one-way trip would take π√(R/g), which for Earth is about 42 minutes. Remarkably, the same 42 minutes holds for a straight tunnel between any two points on the surface, not only one through the centre.

The real Earth is not uniform. Its iron core is far denser than its rocky mantle, so as you descend through the mantle you get closer to that dense core faster than you lose the mantle above you, and gravity actually increases slightly at first, to about 10.7 m/s² at the boundary of the core 2,900 km down, before falling to zero at the centre. Seismologists use measurements of this kind, together with earthquake waves, to work out how density changes inside the planet.

::: note
One force law, F = GMm/r², explains a falling apple, the Moon's orbit and all three of Kepler's laws. Gravity on Earth and gravity in the heavens are the same force, weakening with the square of the distance.
:::

## What the law explained, and what it left open

The circular-orbit argument above only used Kepler's third law. Newton went much further. He showed that an inverse-square force produces orbits that are exactly ellipses with the Sun at a focus, Kepler's first law, and that any force directed toward a fixed centre produces equal areas in equal times, the second. The three separate rules that Kepler had found by fitting data became consequences of one force.

The law also explained where Kepler was slightly wrong. Since every mass attracts every other, the planets pull on one another, and their orbits wobble a little away from perfect ellipses. Astronomers spent the following century and a half calculating these wobbles, and the calculations matched the observations. The most dramatic success came in 1846. Uranus had been drifting from its predicted path, and Urbain Le Verrier in Paris worked out where an unseen planet would have to be to cause the drift. Johann Galle in Berlin pointed a telescope there and found Neptune within about a degree of the predicted spot.

Newton's law is not the final word, and Newton knew it had a gap. It says how strong gravity is, but not how the pull reaches across empty space, and Newton admitted he had no explanation for that. Einstein's general relativity, more than two centuries later, replaced the idea of a force with curved space and time. But in weak gravity, which covers everything in the solar system except a few very precise measurements, the two theories agree so closely that spacecraft are still navigated with Newton's formula.

::: quiz
? If you moved to twice your present distance from Earth's centre, your weight would be
- half as large
+ a quarter as large
- unchanged
= Gravity goes as 1/r², so doubling the distance divides the force by 2² = 4.
? What does the Moon test compare?
- The Moon's mass with Earth's mass
+ The Moon's centripetal acceleration with g reduced by a factor of 60²
- The time light takes to reach the Moon with the length of a month
= The Moon is about 60 Earth radii away, so inverse-square gravity predicts an acceleration of g/3600 there. That matches 4π²r/T², the acceleration worked out from the Moon's orbit alone.
? Kepler's third law (T² ∝ r³) implies that the Sun's pull on a planet goes as
- 1/r
+ 1/r²
- 1/r³
= The centripetal acceleration is 4π²r/T². Replacing T² with a constant times r³ leaves r/r³, which is 1/r².
:::
