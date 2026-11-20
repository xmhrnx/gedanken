Before anyone could ask why the planets move as they do, someone had to find out exactly how they move. That turned out to be much harder than it sounds, because the answer everyone expected was wrong in a way that only very good measurements could reveal. Johannes Kepler found the right answer not from a theory but by fitting curves, patiently and for years, to the most accurate naked-eye observations ever made. The three rules he extracted became the facts that any theory of gravity would have to explain, and Newton's first great success was to explain all three at once.

## Why circles seemed right, and why they failed

For about two thousand years, astronomers took it for granted that the heavens move in circles. The circle was the perfect shape, the same all the way round, and steady motion around a circle seemed the natural behaviour of things that never tire or stop. When the planets refused to move steadily across the sky, sometimes speeding up, sometimes stopping and going backwards for a while, astronomers did not give up the circle. They added more circles: a planet rode on a small circle, called an **epicycle**, whose centre moved around a larger one. With enough circles, the old Earth-centred system of Ptolemy could match the observations reasonably well.

In 1543 Nicolaus Copernicus moved the Sun to the centre and set Earth moving around it with the other planets. This explained the backward loops elegantly: they are what you see when the faster-moving Earth overtakes a slower outer planet. But Copernicus kept the circles and the epicycles, and so his system predicted planetary positions no better than Ptolemy's. For Mars the errors of both systems could reach a few degrees, several times the width of the full Moon. Mars was the hardest planet to fit, and for a reason nobody yet understood: its orbit is further from a circle than that of any planet then known except Mercury, which stays so close to the Sun that it is rarely seen well.

To settle the question, astronomy needed better data. A model can only be tested as precisely as the sky is measured, and the measurements of the time were good to perhaps ten arcminutes, a sixth of a degree. An error that small could hide a lot.

::: history 1576–1601 · Hven and Prague
Tycho Brahe, a Danish nobleman, spent about twenty years on the island of Hven building the finest observatory in Europe. He had no telescope, which had not yet been invented, but he had huge sighting instruments, carefully made and checked against each other, and he measured each planet night after night, not only at the moments astronomers usually cared about. His positions were good to about two arcminutes, a thirtieth of a degree, which is close to the limit of the unaided eye. In 1599 he moved to Prague as astronomer to the Holy Roman Emperor, and in 1600 he took on a gifted, prickly young mathematician named Johannes Kepler and set him to work on Mars. Tycho guarded his data jealously and released it to Kepler only in pieces. Then, in 1601, he died suddenly, and Kepler gained access to the whole record.
:::

Kepler's work on Mars, which he called his "war" with the planet, lasted most of a decade. At one point he had a model built from circles that matched Tycho's positions to within eight arcminutes. That was better than anyone had done before, and an astronomer working with older data would have been delighted. But Kepler knew Tycho's measurements were good to about two arcminutes, so an eight-arcminute miss was four times larger than the possible error. Either the data were wrong or the circle was. He trusted Tycho, gave up the circle, and started again. He later wrote that those eight minutes had pointed the way to a complete reform of astronomy.

## The first law: ellipses

The shape that finally fitted was the **ellipse**. An ellipse is the set of points whose distances to two fixed points add up to the same total; the two points are called the **foci** (singular focus). You can draw one by putting two pins in a board, looping a string around them, and running a pencil around the loop while keeping the string taut. Move the pins together and the ellipse becomes a circle; move them apart and it stretches out.

Two numbers describe an ellipse. Half of its longest diameter is the **semi-major axis**, written a, which works as the orbit's average size. The **eccentricity**, written e, measures how stretched it is: 0 for a circle, and closer and closer to 1 for a longer and thinner ellipse. Kepler's first law says that each planet moves on an ellipse with the Sun at one focus. The Sun is not at the centre, which is the part that surprised people most: the ellipse is lopsided around it.

Because the Sun sits at a focus, the planet's distance from it changes around the orbit. The closest point, the **perihelion**, is at a distance a(1 − e), and the farthest, the **aphelion**, is at a(1 + e). For Earth, e is only 0.017, so its orbit is very nearly a circle, merely offset a little from the Sun. That is part of why the circle had survived so long: for most planets the difference is small. Mars has e ≈ 0.093, and its distance from the Sun swings by about 20% between perihelion and aphelion, which was just enough for Tycho's measurements to catch. Comets show what high eccentricity looks like: many have e above 0.99, plunging close to the Sun and then retreating far beyond the planets.

## The second law: equal areas

Kepler actually found his second law before the first. Imagine a line drawn from the Sun to the planet. As the planet moves, this line sweeps across the inside of the orbit like the hand of a clock, and Kepler found that it sweeps out equal areas in equal times.

That rule fixes how the planet's speed changes. Near perihelion the line is short, so to sweep out a given area in a day it has to swing through a large angle, which means the planet must move fast. Near aphelion the line is long, and a small swing covers the same area, so the planet moves slowly. The planet speeds up as it falls toward the Sun and slows down as it climbs away.

Earth shows this, though gently. It is closest to the Sun in early January, when it moves at about 30.3 km/s, and farthest in early July, at about 29.3 km/s. The difference is visible in the calendar. In the northern hemisphere, the half of the year from the September equinox to the March equinox, when Earth is nearer the Sun and moving faster, is about a week shorter than the other half.

## The third law: periods and sizes

The first two laws describe each orbit separately. Kepler was also convinced, for reasons that were partly mystical, that the planets must form a single ordered system, and he searched for years for a rule linking them. He found it in 1618. If T is a planet's period, the time for one orbit, and a its semi-major axis, then T² is proportional to a³, with the same constant of proportionality for every planet.

The constant can be made to disappear by choosing units well. Measure T in years and a in **astronomical units** (AU), where one AU is Earth's average distance from the Sun, about 1.50 × 10⁸ km. Then for Earth T = 1 and a = 1, so the constant is 1, and since it is the same for all the planets, it is 1 for every one of them:

::: equation Kepler's third law, in years and AU
T^2 = a^3
:::

::: steps Predicting the periods of Mars and Jupiter
$ T = a^{3/2}
Take the square root of both sides of T² = a³. Both sides are positive, so this is allowed, and it puts the law in the form we want: give it an orbit size and it returns a period. The exponent 3/2 is larger than 1, which means a planet twice as far out takes nearly three times as long, because it has a longer path to cover and also moves along it more slowly.
---
$ T_{\text{Mars}} = 1.524^{3/2} \approx 1.88
Mars's semi-major axis is 1.524 AU. Raising this to the power 3/2 gives 1.88 years, or 687 days, which is exactly Mars's observed period. The law is reproducing a number that was measured independently, from the time Mars takes to return to the same place among the stars.
---
$ T_{\text{Jupiter}} = 5.203^{3/2} \approx 11.9
Jupiter is 5.203 AU from the Sun, and the formula gives about 11.9 years, again matching observation, and Saturn, at 9.54 AU, comes out at about 29.5 years, as observed. A rule that works across a factor of more than twenty in distance is not a coincidence. After telescopes arrived, astronomers found that Jupiter's own moons obey the same law around Jupiter, with a different constant, which was an early hint that the constant depends on the body at the centre.
:::

The third law was also useful in reverse. Periods are easy to measure: you only need to watch a planet return to the same place relative to the Sun and stars, and correct for Earth's own motion. Distances in kilometres are very hard to measure. But T² = a³ turns every period into a distance in AU, so the law gave the relative sizes of all the planetary orbits, a complete scale map of the solar system, more than a century before anyone knew how many kilometres an AU was. A single measured distance would fix the scale of the whole map. That distance came first from observations of Mars in 1672 and more precisely from the transits of Venus across the Sun in the 1760s.

## Two worked examples: a comet and a voyage to Mars

The laws are worth trying on cases Kepler never saw, because they apply to anything that orbits the Sun, not only to the planets he studied.

Start with Halley's comet. Its orbit has a semi-major axis of about 17.8 AU and an eccentricity of about 0.967, so it is a long, thin ellipse. The third law gives its period at once: 17.8 to the power 3/2 is about 75 years, which is the interval between its returns, last seen in 1986 and due again in 2061. The first law gives its range of distances. At perihelion it is a(1 − e) = 17.8 × 0.033, about 0.59 AU from the Sun, inside the orbit of Venus. At aphelion it is a(1 + e), about 35 AU, beyond Neptune. The second law then tells us how its speed changes between these points. At perihelion and at aphelion the comet is moving exactly at right angles to the line from the Sun, so in a short time it sweeps out a thin triangle whose area is half of its distance times the length it travels. Equal areas in equal times then means distance times speed is the same at the two ends.

::: equation Speeds at the two ends of an orbit
\frac{v_{\text{peri}}}{v_{\text{aph}}} = \frac{r_{\text{aph}}}{r_{\text{peri}}} = \frac{1+e}{1-e}
:::

For Earth, with e = 0.017, the ratio is about 1.035, which is the 30.3 against 29.3 km/s quoted above. For Halley's comet it is 1.967/0.033, about 60. The comet rushes past the Sun at about 55 km/s and then crawls through its distant aphelion at less than 1 km/s. So it spends most of its 75 years far out and slow, and only a few months close enough to grow a bright tail, which is why each return is brief.

The second example is a spacecraft travelling from Earth to Mars. Once its engine is off, a spacecraft is just another small body orbiting the Sun, and Kepler's laws apply to it. The cheapest route, worked out by Walter Hohmann in 1925, is half of an ellipse that touches Earth's orbit at perihelion and Mars's orbit at aphelion. Its semi-major axis is the average of the two distances, (1 + 1.524)/2 ≈ 1.26 AU, so its full period would be 1.26 to the power 3/2, about 1.42 years. The craft flies only half the ellipse, which takes about 0.71 years, or roughly 260 days. That is close to the eight or nine months that real missions to Mars take on such paths.

There is a catch the third law also explains. Mars must arrive at the meeting point when the spacecraft does, so the launch has to happen when the two planets are in the right positions. Because Earth goes round in one year and Mars in 1.88, Earth gains one full lap on Mars every 1/(1 − 1/1.88) ≈ 2.1 years, about 26 months. That is why launch windows to Mars open only once every two years or so, and why space agencies plan missions around them.

## Going deeper: why equal areas?

Of the three laws, the second has the most general explanation, and seeing it shows what the law is really telling us. It is the same idea that appears in Angular Momentum: a spinning skater speeds up as she pulls her arms in.

::: steps Equal areas from a force aimed at the Sun
$ \mathrm{d}A = \tfrac{1}{2}\, r \cdot r\,\mathrm{d}\theta
In a short time the line from the Sun to the planet turns through a small angle dθ, measured in radians. The region it sweeps is a thin triangle with two long sides of length r and a short side of length r dθ, the small arc the planet moves along. The area of a triangle is half its base times its height, so the area swept is ½ r² dθ.
---
$ \frac{\mathrm{d}A}{\mathrm{d}t} = \tfrac{1}{2}\, r^2 \frac{\mathrm{d}\theta}{\mathrm{d}t}
Divide by the short time dt to get the rate at which area is swept. The fraction dθ/dt is how fast the planet's direction from the Sun is turning. Kepler's second law says this whole rate is constant, so wherever r is small, the turning rate must be large, which is the speeding up near perihelion.
---
$ L = m r^2 \frac{\mathrm{d}\theta}{\mathrm{d}t}
Now compare a quantity from mechanics: the **angular momentum** L of a body of mass m going round a centre, which measures how much turning motion it carries. For a body at distance r turning at rate dθ/dt it has exactly this form, and it changes only if some force pushes the body sideways, across the line to the centre.
---
$ \frac{\mathrm{d}A}{\mathrm{d}t} = \frac{L}{2m}
Comparing the two expressions, the rate of sweeping area is the angular momentum divided by twice the mass. The Sun's pull points straight along the line to the Sun, so it has no sideways part and cannot change L. So L is constant, and the area swept per second is constant: the second law is the conservation of angular momentum written as geometry.
:::

This argument tells us something Kepler could not know: the second law would hold for any force that points straight toward the Sun, whatever its strength and however it changed with distance. Newton proved exactly this, with a geometrical version of the same argument, in the opening pages of the Principia. The second law therefore tells us only the direction of the force. The first and third laws are what reveal how strong it is: only a force that weakens as the square of the distance gives closed ellipses with the Sun at a focus, and only that force gives T² proportional to a³. Universal Gravitation shows how Newton drew those conclusions.

::: note
Each planet moves on an ellipse with the Sun at one focus, sweeping out equal areas in equal times, so it moves fastest when closest to the Sun. The square of the period is proportional to the cube of the orbit's size, with the same constant for every planet, which links all the orbits into one system.
:::

## Common misconceptions

The most widespread mistake is to think that the seasons come from Earth's changing distance from the Sun. If they did, both hemispheres would have summer at the same time, and they do not. Earth is in fact closest to the Sun in early January, in the middle of the northern winter. The distance does change the sunlight Earth receives, by about 7% between perihelion and aphelion, because brightness falls as the square of the distance and 1.034² ≈ 1.07, but this is much smaller than the effect of the tilt of Earth's axis, which changes how high the Sun climbs in the sky and how long it stays up. The seasons come from the tilt.

A second mistake comes from diagrams. Textbooks draw planetary orbits as visibly squashed ellipses, so that the shape is clear, and readers come away thinking the planets swing wildly in and out. Drawn to scale, Earth's orbit is indistinguishable from a circle by eye. Its shortest diameter is less than its longest by about 0.015%, a fraction of a millimetre on a page-sized drawing. What does show is that the Sun is off-centre, by 1.7% of the radius. Even Mars's orbit is round to within half a percent; its eccentricity shows mainly as the Sun sitting 9% away from the centre. An eccentric orbit is not so much a stretched circle as a shifted one.

A third mistake is to put the Sun at the centre of the ellipse. The centre of an ellipse is an empty point in space with nothing special about it; the Sun is at a focus, and the second focus is also empty. This matters because all the orbit's behaviour, the closest approach, the farthest point and the changes of speed, is measured from the focus.

Finally, people sometimes read the third law as saying that heavier planets take longer to go round. The law contains no mass of the planet at all. Jupiter is more than 300 times as massive as Earth, and a pebble at Jupiter's distance would go round in the same 11.9 years. The constant in T² ∝ a³ depends only on the body at the centre, a fact that Newton's theory explains and that the lessons on Orbits and Universal Gravitation make use of.

## What the laws do not say

Kepler's laws describe; they do not explain. They say nothing about why orbits should be ellipses rather than some other shape, why equal areas, or why the exponent is 3/2 and not something else. Kepler did look for a cause. He imagined that the rotating Sun sent out some influence, perhaps magnetic, that swept the planets around, weaker farther out, which is why outer planets move more slowly. It was a serious attempt to treat the solar system as a physical machine rather than a geometry exercise, and in that spirit it was ahead of its time, but the mechanism was wrong. He also spent much effort on the idea that the planets' speeds formed musical harmonies, and the third law appeared in a book devoted largely to that theme.

The laws are also not exactly right. The planets pull on each other as well as being pulled by the Sun, so real orbits wobble slightly and slowly turn. Kepler's data were not good enough to show this, and that was fortunate: if the deviations had been large, the clean pattern might never have been found.

The explanation came about seventy years later. Newton showed that all three laws follow from a single force, a pull toward the Sun that weakens with the square of the distance, and that the same force also accounts for the small departures from them. That is the next lesson, Universal Gravitation.

## Where the laws stop being exact

Kepler's laws are exact for one small body going round one large body that does not move. The real solar system breaks both conditions, and each break became a discovery.

First, the Sun moves. Newton's third law says the planet pulls the Sun as hard as the Sun pulls the planet, so both go round their common centre of mass. For Jupiter, about 1/1047 of the Sun's mass, that centre lies about 740,000 km from the Sun's centre, just outside its surface. The full form of the third law, which the Orbits lesson derives, has the sum of the two masses in it rather than the Sun's mass alone, so the constant is very slightly different for each planet. For Jupiter the correction is about one part in a thousand, too small for Tycho's instruments but easy for a modern telescope. The same wobble of a star, measured through the Doppler shift of its light, is now one of the main ways planets around other stars are found.

Second, the planets pull on each other. Their tugs make each orbit's ellipse slowly turn and change shape. Once Newton's theory could calculate these departures, they became a tool for discovery: the unexplained drift of Uranus led to the finding of Neptune in 1846, a story told in Universal Gravitation.

Third, and most subtly, the inverse-square law itself is not quite exact. Mercury's ellipse turns by about 574 arcseconds per century. The pulls of the other planets account for about 531 of them; the remaining 43 were unexplained for more than half a century, until Einstein's general relativity accounted for exactly that amount in 1915. Near a very massive body, or for very precise measurements, Kepler's ellipses are replaced by orbits that slowly rotate, as described in Curved Spacetime.

## Where Kepler's laws show up

Every orbiting object obeys these laws to a good first approximation, which makes them among the most used results in astronomy. The third law, in the general form Newton gave it, is how astronomers weigh things they cannot touch: measure the period and size of a moon's orbit and you learn the mass of its planet, measure a pair of stars going round each other and you learn the stars' masses. Applied to stars orbiting the centre of our galaxy, it reveals a black hole of about four million solar masses, as described in Seeing the Invisible. Applied to stars and gas in the outskirts of galaxies, it failed: they move too fast for the visible matter, which was one of the first clues to Dark Matter.

Closer to home, the third law sets the height of geostationary satellites, and the ellipses and equal areas govern every spacecraft trajectory, from the Hohmann transfer above to the slingshot passes that carried the Voyager probes past the outer planets. The NASA telescope that found thousands of planets around other stars by watching them pass in front of their stars was named Kepler, and it used his third law to turn each planet's period into the size of its orbit.

The lessons that follow build on these laws directly. Universal Gravitation shows how Newton derived them from one force, Orbits uses them to understand satellites and spacecraft, and Escape Velocity asks what happens when an orbit stops being an ellipse at all. Angular Momentum gives the general principle behind the second law.

::: quiz
? Where in its orbit does a planet move fastest?
- At aphelion, the farthest point from the Sun
+ At perihelion, the closest point to the Sun
- It moves at the same speed everywhere
= The line from the Sun must sweep equal areas in equal times. Where that line is shortest, the planet has to travel farther along its orbit each day to cover the same area, so it moves fastest at perihelion.
? An asteroid has a semi-major axis of 4 AU. What is its period?
- 4 years
+ 8 years
- 16 years
= In years and AU, T = a^(3/2). The square root of 4 is 2, and 2 cubed is 8, so the asteroid takes 8 years to go round the Sun.
? Why did Kepler abandon circular orbits for Mars?
- Circles could not put the Sun at the centre
+ His best circular model missed Tycho's positions by 8 arcminutes, four times the measurements' likely error
- Newton's gravity required ellipses
= Tycho's positions were good to about 2 arcminutes, so an 8-arcminute miss could not be blamed on the data. It meant the model itself was wrong.
? A comet's closest approach to the Sun is 0.5 AU and its farthest is 31.5 AU. What is its period?
- About 32 years
+ About 64 years
- About 180 years
= The semi-major axis is the average of the two distances, (0.5 + 31.5)/2 = 16 AU. Then T = 16^(3/2) = 4³ = 64 years.
? Earth is closest to the Sun in early January. Why is that not a contradiction with northern winter?
- Because Earth moves fastest in January and so receives less heat
+ Because the seasons are set by the tilt of Earth's axis, and the 3% change in distance matters much less
- Because the Sun is at the centre of the ellipse, so the distance does not really change
= The tilt changes how high the Sun climbs and how long the day lasts, which outweighs the roughly 7% change in sunlight from distance. The southern hemisphere has its summer in January.
:::
