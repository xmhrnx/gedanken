The three laws of motion were first set out together in a single book: *Philosophiæ Naturalis Principia Mathematica*, "Mathematical Principles of Natural Philosophy", published by Isaac Newton in 1687. Stating the laws was only the beginning of what it did. It showed that the force that pulls an apple to the ground is the same force that holds the Moon in its path round the Earth and the planets in theirs round the Sun, and it worked out the motions of the solar system from a handful of principles. Before the *Principia*, it was widely assumed that the heavens and the Earth followed different rules: on Earth things fell and stopped, while in the sky they circled forever. After it, one set of laws covered both. This lesson tells how the book came to be written, what is in it, and the single calculation at its heart.

## What was the question everyone was asking?

By the 1680s the motions of the planets had been described with great precision, but not explained. Early in the century Johannes Kepler had found three rules that the planets obey. Each moves on an ellipse with the Sun at one focus. A line from the Sun to the planet sweeps out equal areas in equal times, so a planet moves faster when it is closer to the Sun. And the square of a planet's orbital period is proportional to the cube of its average distance from the Sun. Kepler had extracted these rules from years of observations, but they were descriptions, not consequences of anything deeper.

Several people in London suspected the deeper cause. If the Sun pulls each planet with a force that weakens as the inverse square of the distance, so that twice as far means a quarter of the pull, then something follows quickly for circular orbits. The circular motion rule a = 4π²r/T² says that the acceleration of a planet on a circle of radius r is proportional to r/T². Kepler's third law says T² is proportional to r³. Put them together and the acceleration is proportional to r/r³, which is 1/r². So an inverse-square pull fits Kepler's third law, at least if the orbits are circles.

But the orbits are not circles; they are ellipses. And no one could prove that an inverse-square force would make a planet move on an ellipse rather than on some other curve. That was the gap, and it was a hard one, because on an ellipse the distance from the Sun keeps changing, so the force keeps changing in strength and direction at the same time.

::: history 1684 · London and Cambridge
In January 1684 three members of the Royal Society met in a London coffee house: Christopher Wren, the architect of St Paul's Cathedral; Robert Hooke, the Society's curator of experiments; and Edmond Halley, a young astronomer. They agreed that an inverse-square force probably explained the planets' motions. Hooke claimed he could prove it. Wren, doubtful, offered a book worth forty shillings to whichever of them could produce a proof within two months. Nobody did. That August, Halley was in Cambridge and called on Newton, then a reclusive professor of mathematics. He asked what curve a planet would follow if the Sun attracted it with an inverse-square force. Newton replied at once: an ellipse. He said he had calculated it some years before. He could not find his papers, but promised to work it out again. In November a nine-page treatise arrived in London, *De motu corporum in gyrum*, "On the motion of bodies in orbit". Halley recognised at once what it was and went back to Cambridge to urge Newton to develop it. Over the next two and a half years, Newton worked on almost nothing else, and the nine pages grew into the *Principia*.
:::

## What is in the book?

The *Principia* opens with definitions of the basic quantities, including mass, momentum and force, followed by the three laws of motion. Then it is divided into three books.

**Book I** treats the motion of bodies under forces in empty space, as pure mathematics. Its very first proposition proves that any force directed toward a fixed centre, whatever its strength and however it varies, makes a body sweep out equal areas in equal times. That is Kepler's second law, and it turns out to need nothing about the inverse square at all, only that the force points at the centre. Later propositions prove that an inverse-square force makes a body move on a conic section: an ellipse, a parabola or a hyperbola, depending on its speed. That closed the gap that had defeated Hooke, Wren and Halley. Book I also proves a theorem that the rest of the argument depends on: a uniform sphere, or a sphere built of uniform shells, attracts any body outside it exactly as if all its mass were concentrated at its centre.

**Book II** treats motion through resisting fluids such as air and water. Much of it is a mathematical study of drag, but it had a polemical purpose. The leading theory of the solar system at the time was René Descartes', in which the planets are carried round the Sun by a huge whirlpool, or vortex, of invisible fluid matter. Newton showed that such a vortex could not reproduce Kepler's laws: the speeds at different distances came out wrong. If the planets were not being carried, something else had to be moving them.

**Book III**, titled "The System of the World", applies the mathematics to nature. Its central claim is that gravity acts between every pair of bodies in the universe, with a force proportional to each of their masses and inversely proportional to the square of the distance between them:

::: equation Newton's law of universal gravitation
F = \frac{GMm}{r^2}
:::

From this one law, together with the three laws of motion, Newton explained a remarkable range of things. Kepler's three laws followed. The tides came from the pull of the Moon and the Sun on the oceans, stronger on the near side of the Earth than at its centre and weaker on the far side. The slow wobble of the Earth's axis, which makes the position of the stars at the equinoxes drift over thousands of years, came from the Sun and Moon pulling on the Earth's equatorial bulge. That bulge itself, a slight flattening of the Earth at the poles, came from its rotation. And comets, previously regarded as unpredictable omens, turned out to move on very elongated conic sections around the Sun, like the planets.

One thing Newton could not do was find the value of the constant G. Gravity between ordinary objects is so weak that measuring it needed extraordinary care, and it was not done until Henry Cavendish's experiment of 1798, which used a delicate twisting balance to detect the attraction between lead spheres. Newton managed without G by working only with ratios: comparing one orbit with another, or the Moon's fall with an apple's.

## How did Newton check that gravity reaches the Moon?

The boldest claim in Book III is that the force making an apple fall is the same force that holds the Moon in orbit. Newton tested this with a calculation that can now be repeated with modern figures. The test rests on two facts: the Moon is about 60 Earth radii from the Earth's centre, and, by the theorem from Book I, the Earth pulls on outside bodies as if all its mass were at its centre. So an apple at the surface is one Earth radius from the centre of attraction, and the Moon is sixty.

::: steps Does gravity at the Moon fall off as the inverse square?
$ a_{\text{predicted}} = \frac{g}{60^2} = \frac{9.8}{3600} \approx 2.72 \times 10^{-3}\,\text{m/s}^2
Start with the prediction. At the surface, one Earth radius from the centre, gravity gives an acceleration g of 9.8 m/s². If gravity weakens as the inverse square of distance, then at 60 times that distance it should be 60², or 3,600, times weaker. Dividing g by 3,600 gives the acceleration the Moon should have if the same gravity reaches it.
---
$ a_{\text{observed}} = \frac{4\pi^2 r}{T^2}
Now the observation, which uses nothing about gravity at all. The Moon moves on a nearly circular orbit, so whatever force holds it there must give it the centripetal acceleration from the circular motion lesson. That acceleration depends only on the size of the orbit and the time taken to go round it, both of which astronomers could measure.
---
$ a_{\text{observed}} = \frac{4\pi^2 \times 3.84 \times 10^8}{(2.36 \times 10^6)^2} \approx 2.72 \times 10^{-3}\,\text{m/s}^2
Put in the Moon's distance, 3.84 × 10⁸ m, and its orbital period relative to the stars, 27.3 days, which is 2.36 × 10⁶ s. The result matches the prediction from the first step. Two numbers obtained in completely different ways, one from dropping objects on Earth and one from watching the sky, agree.
:::

The agreement says something striking. The Moon is falling toward the Earth, continually, with exactly the acceleration an apple would have if it were lifted to the Moon's distance. It never arrives because it is also moving sideways, fast enough that the Earth's surface keeps curving away beneath it, just like the cannonball in the circular motion lesson. Newton later recalled that he had first tried a rougher version of this comparison around 1666 and found that the results "answer pretty nearly". The early mismatch was probably due partly to a poor figure for the size of the Earth; with the better measurements available by the 1680s, the agreement became close. In a popular account of the system of the world published shortly after his death, he illustrated the idea with exactly that cannon on a mountaintop, firing ever faster until the ball no longer comes down.

## How was the book written and received?

The *Principia* is written in Latin and argued almost entirely in geometry. Newton had already invented his version of calculus, which he called the method of fluxions, but he did not use its notation in the book. Instead he reasoned with diagrams and with ratios of quantities as they shrink toward zero, a geometric form of the same ideas, written in the mathematical language his readers knew and could check. The result is extremely demanding. Even well-trained readers struggled, and few read it all.

Publishing it was a drama of its own. The Royal Society agreed to publish the book but had recently spent its funds on a lavishly illustrated *History of Fishes*, so Halley paid for the printing himself and edited the text through the press. Then Hooke demanded acknowledgement that he had suggested the inverse-square law first. Newton, furious, threatened to withhold Book III altogether, and in the end removed several references to Hooke from the text instead. The book appeared in July 1687.

Not everyone was persuaded. Followers of Descartes on the Continent, and Gottfried Leibniz, objected that Newton's gravity acted across empty space with no mechanism to carry it. To them that was a step backward, toward the "occult qualities" that the new science was supposed to have abandoned. Newton agreed that he did not know what gravity was. In the second edition of 1713 he wrote "hypotheses non fingo", "I frame no hypotheses": he would describe precisely how gravity behaves and leave its cause open.

The predictions settled the matter over the following decades. Using Newton's theory, Halley studied the records of bright comets and concluded that those seen in 1531, 1607 and 1682 were one object, returning about every 76 years. He predicted it would return in 1758. Halley died in 1742. The comet was sighted on Christmas night of 1758, and it has carried his name ever since.

::: note
The *Principia* showed that one set of laws, the three laws of motion plus universal gravitation, governs falling apples, the Moon, the planets, comets and the tides. The Moon test made the case concrete: the Moon falls toward the Earth with exactly the acceleration an inverse-square weakening of ordinary gravity predicts.
:::

Newton's mechanics stood without serious challenge for more than two centuries. Einstein's relativity eventually replaced it for very high speeds and very strong gravity, and the gravity chapter follows that story. But for almost everything else, from designing bridges to steering spacecraft to other planets, the laws of the *Principia* are still the ones engineers use.

::: quiz
? In the Moon test, why is the predicted acceleration g/3600?
- Because the Moon is 3600 times lighter than the Earth
+ Because the Moon is about 60 Earth radii away, and gravity weakens as the inverse square
- Because a month is about 3600 times longer than a minute
= An inverse-square force at 60 times the distance from the Earth's centre is 60² = 3600 times weaker than at the surface. The Moon's mass plays no part, because it cancels from its own acceleration.
? The Moon test compares an apple, one Earth radius from the Earth's centre, with the Moon at 60 radii. Which result from Book I justifies measuring from the Earth's centre?
- Kepler's second law
+ A uniform sphere attracts outside bodies as if all its mass were at its centre
- Comets move on conic sections
= Without that theorem it would be unclear how far an apple is from the Earth, whose mass is spread out all around and beneath it. The theorem lets the whole Earth be treated as a point at its centre.
? The Moon has a sizeable acceleration toward the Earth. Why does it not crash into it?
- The Sun's pull balances the Earth's
- Gravity switches off beyond the atmosphere
+ Its sideways velocity carries it past as fast as it falls
= Like Newton's cannonball, the Moon is falling all the time, but it moves sideways fast enough that the Earth's surface curves away beneath it, so it never gets closer.
:::
