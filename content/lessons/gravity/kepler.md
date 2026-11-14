Before anyone knew *why* planets move as they do, someone had to find out exactly *how* they move. Johannes Kepler did this not from theory but by fitting curves to the most accurate naked-eye observations ever made. The three rules he extracted became the facts that Newton's gravity had to explain.

## The problem with circles

For two thousand years astronomers assumed planets move on circles, or on circles riding on circles (epicycles), at steady speeds. Copernicus put the Sun at the centre in 1543 but kept the circles, and his system predicted planetary positions no better than the old Earth-centred one. Predicted positions of Mars, whose orbit is the least circular of the planets then known, could be off by a few degrees.

Kepler inherited Tycho Brahe's records of Mars, accurate to about 2 arcminutes (a thirtieth of a degree). After years of work his best circular model still missed Tycho's positions by up to 8 arcminutes. A lesser astronomer would have blamed the data. Kepler trusted it and abandoned the circle.

::: history 1600–1619 · Prague and Linz
Tycho Brahe built an observatory on the Danish island of Hven and spent two decades measuring planetary positions with giant sighting instruments. In 1600 he hired the young Kepler in Prague, then died in 1601, leaving him the data. Kepler's "war with Mars" lasted most of a decade. He published the first two laws in *Astronomia Nova* (1609), having found the area law before the ellipse, and the third law in *Harmonices Mundi* (1619). He got the cause wrong: he imagined a force from the rotating Sun sweeping the planets along, perhaps magnetic, and he also hunted for musical harmonies among the orbits.
:::

## The first law: ellipses

An **ellipse** is the set of points whose distances to two fixed points, the foci, add up to the same total. Draw one with a loop of string around two pins. Half the longest diameter is the **semi-major axis** a, which serves as the orbit's average size. The **eccentricity** e measures how squashed it is: 0 for a circle, approaching 1 for a long thin ellipse.

::: note
Each planet moves on an ellipse with the Sun at one focus, not at the centre. The planet sweeps out equal areas in equal times, and the square of its period is proportional to the cube of its semi-major axis.
:::

The closest point to the Sun (perihelion) is at distance a(1 − e) and the farthest (aphelion) at a(1 + e). Earth's eccentricity is only 0.017, so its orbit is very nearly a circle offset slightly from the Sun. Mars has e ≈ 0.093, enough that its distance from the Sun swings by about 20%. Comets can have e above 0.99.

## The second law: equal areas

Draw a line from the Sun to a planet. As the planet moves, the line sweeps out equal areas in equal times. Near perihelion the line is short, so the planet must move faster to cover the same area; near aphelion it moves slower. Earth is closest to the Sun in early January and moves at about 30.3 km/s then, against about 29.3 km/s in early July. That is why the stretch from the September equinox to the March equinox is about a week shorter than the other half of the year.

## The third law: periods and sizes

Kepler searched for a link between a planet's period T and its orbit size a, and found that T² is proportional to a³ for every planet. Measuring T in years and a in astronomical units (AU, the average Earth–Sun distance, about 1.50 × 10⁸ km) makes the constant exactly 1 for Earth, and therefore for all planets:

::: equation Kepler's third law, in years and AU
T^2 = a^3
:::

::: steps Predicting the periods of Mars and Jupiter
$ T = a^{3/2}
Take the square root of both sides of T² = a³. A planet farther out takes disproportionately longer: it has a bigger circuit and also moves more slowly.
---
$ T_{\text{Mars}} = 1.524^{3/2} \approx 1.88
Mars's semi-major axis is 1.524 AU. Its observed period is 1.88 years (687 days).
---
$ T_{\text{Jupiter}} = 5.203^{3/2} \approx 11.86
Jupiter, at 5.203 AU, takes 11.86 years, matching observation. The law also holds for Jupiter's moons around Jupiter, with a different constant.
:::

The third law also runs backwards. Periods are easy to measure, so the law gave the relative sizes of all the planetary orbits long before anyone could measure a distance in kilometres. One absolute distance, found later from transits of Venus, then fixed the whole scale.

## What the laws do not say

Kepler's laws are descriptions. They say nothing about why ellipses, why equal areas, or why the power 3/2. They are also slightly wrong: planets tug on each other, so real orbits wobble and slowly turn. Newton showed that all three laws follow from one force law, and that the small deviations follow from the same law too. That is the next lesson.

::: quiz
? Where in its orbit does a planet move fastest?
- At aphelion, the farthest point
+ At perihelion, the closest point
- It moves at constant speed everywhere
= Equal areas in equal times: where the Sun–planet line is shortest, the planet must cover more distance to sweep the same area.
? An asteroid has a semi-major axis of 4 AU. What is its period?
- 4 years
+ 8 years
- 16 years
= T = a^(3/2) = 4^(3/2) = 8 years.
? Why did Kepler abandon circular orbits for Mars?
- Circles could not put the Sun at the centre
+ His best circular model disagreed with Tycho's data by 8 arcminutes, more than the data's error
- Newton's gravity required ellipses
= Tycho's positions were good to about 2 arcminutes, so an 8-arcminute miss meant the model, not the data, was wrong.
:::
