The sea rises and falls about twice a day, and the rhythm follows the Moon: high tide comes roughly 50 minutes later each day, just as moonrise does. People have connected tides with the Moon since antiquity. But two features of the tides resist the obvious explanation. If the Moon simply pulls the water toward itself, there should be one bulge of water facing the Moon and one high tide a day, yet there are two. And the Sun, whose gravitational pull on Earth is about 180 times stronger than the Moon's, raises tides less than half as large. Both puzzles have the same answer: tides are caused not by gravity itself but by the difference in gravity from one side of a body to the other.

## Why the average pull cannot be felt

Start with an observation that seems unrelated. Earth and the Moon orbit their common centre of mass, the balance point between them, which lies inside Earth about three quarters of the way from the centre to the surface. That means Earth as a whole is constantly accelerating toward the Moon, in free fall, just as the Moon is constantly falling toward Earth.

Everything on Earth shares that fall: the ground, the oceans, the people. As the previous lessons showed with astronauts in orbit, when you and your surroundings fall together, the pull that causes the fall cannot be felt. So the Moon's average pull on Earth has no visible effect on the oceans at all. Only one thing can produce an effect: the extent to which the Moon's pull differs from place to place across Earth.

And it does differ, because gravity weakens with distance. The side of Earth facing the Moon is about 6,400 km closer to the Moon than Earth's centre, so it is pulled slightly harder than the centre. The far side is 6,400 km farther away, so it is pulled slightly less.

## Why there are two bulges

Now look at the oceans from the point of view of Earth's centre, which is falling toward the Moon at the average rate. The water on the near side is pulled a little harder than the centre, so it is pulled ahead, toward the Moon, relative to the solid Earth. The water on the far side is pulled a little less than the centre, so the centre is pulled away from it and it is left behind. Relative to Earth, the far-side water moves away from the Moon.

The result is two bulges, one on the side facing the Moon and one on the side facing away. In between, around the middle, the water is slightly lower. A body that feels a stronger pull on one side than the other is stretched along the line joining it to the source of the pull, and this stretching effect is called a **tidal force**.

Earth spins under these two bulges, which stay roughly lined up with the Moon. A point on the coast is carried through one bulge, then through the other, giving two high tides and two low tides in each rotation. The Moon does not stay still while this happens: it moves about 13° along its orbit each day, in the same direction as Earth spins, so Earth has to turn a little more than once to bring a coastline back under the same bulge. That extra turning takes about 50 minutes, which is why a full cycle lasts about 24 hours 50 minutes and successive high tides are about 12 hours 25 minutes apart.

## How strong is the stretch?

To compare the Moon with the Sun we need a formula for the size of the difference in pull. The idea is simple: work out the pull on the near side, subtract the pull at the centre, and see how the remainder depends on the distance and mass of the body doing the pulling. Because Earth is small compared with the distance to the Moon, an approximation makes the answer clear.

::: steps The tidal acceleration
$ \Delta a = \frac{GM}{(d-R)^2} - \frac{GM}{d^2}
Here M is the Moon's mass, d the distance between the centres of Earth and the Moon, and R Earth's radius. The first term is the acceleration the Moon gives to water on the near side, which is a distance d − R from the Moon; the second is the acceleration it gives to Earth's centre. Their difference is the extra pull on the near side relative to the centre, which is what raises the tide.
---
$ \frac{GM}{(d-R)^2} = \frac{GM}{d^2}\left(1 - \frac{R}{d}\right)^{-2}
Rewrite the first term by taking d out of the bracket: d − R is the same as d(1 − R/d), so its square is d²(1 − R/d)². This puts the whole difference in terms of the small ratio R/d, which for the Moon is about 1/60. The reason for doing this is that expressions in a small quantity can be simplified with a reliable approximation.
---
$ \left(1 - \frac{R}{d}\right)^{-2} \approx 1 + \frac{2R}{d}
For any small number x, (1 − x)⁻² is very nearly 1 + 2x, which you can check by multiplying (1 + 2x)(1 − x)² out and seeing that only terms in x² and smaller remain beyond 1. With x = 1/60 the approximation changes the tidal result by only about 3%, and for the much more distant Sun the error is far smaller still.
---
$ \Delta a \approx \frac{GM}{d^2}\left(1 + \frac{2R}{d}\right) - \frac{GM}{d^2}
Put the approximation back into the first term. Written out this way, the near-side pull is the pull at the centre, GM/d², plus a small extra piece. Subtracting the pull at the centre is now straightforward.
---
$ \Delta a \approx \frac{2GMR}{d^3}
The two copies of GM/d² cancel exactly, which is the mathematical version of saying that the average pull has no effect. What is left is the small extra piece, GM/d² multiplied by 2R/d. The result says the tidal acceleration grows with the mass of the pulling body and with the size of the body being stretched, and it falls off as 1/d³.
:::

That cube is the key to the second puzzle. Gravity itself falls off as 1/d², but the difference in gravity across a body falls off as 1/d³, one power faster. So for tides, distance matters much more than it does for the pull itself. A very massive body that is far away can exert a strong pull but only a weak stretch, while a smaller, nearby body can exert a larger stretch.

Putting in numbers makes this concrete. For the Moon, GM = 4.90 × 10¹² m³/s² and d = 3.84 × 10⁸ m, and the tidal acceleration comes out at about 1.1 × 10⁻⁶ m/s², roughly a ten-millionth of g. For the Sun, GM = 1.33 × 10²⁰ m³/s², about 27 million times the Moon's, but d = 1.50 × 10¹¹ m, about 390 times farther. Cubing 390 gives about 60 million, which more than cancels the mass advantage, and the Sun's tidal acceleration comes out at about 5 × 10⁻⁷ m/s², or about 46% of the Moon's. The Sun pulls on Earth 180 times harder than the Moon does, but it stretches Earth less than half as much.

::: note
Tides come from the difference in gravity between one side of a body and the other, not from the pull itself. That difference goes as M/d³, so the nearby Moon raises larger tides than the far more massive but distant Sun.
:::

## Spring and neap tides

Since the Sun also raises two bulges, the tide we see is the combination of the Moon's and the Sun's. When the Sun, Moon and Earth are lined up, as they are at new Moon and full Moon, the solar and lunar bulges point the same way and add together, giving larger tides called **spring tides**. The name has nothing to do with the season; it comes from the sense of water springing up. At first and last quarter, when the Moon appears half lit, the Sun and Moon are at right angles as seen from Earth. Then the Sun's bulges fall where the Moon's low water is, partly filling it in and partly cancelling the Moon's effect, giving smaller tides called **neap tides**. With the Sun's effect at 46% of the Moon's, the ratio between them should be roughly (1 + 0.46)/(1 − 0.46), about 2.7.

The two-bulge picture is an idealisation, and real tides depart from it a great deal. On an Earth covered entirely by deep ocean, it predicts a range between high and low water of about half a metre. But the continents block the bulges from sweeping freely around the globe, and each ocean basin and bay has its own natural period of sloshing, like water in a bath. The tidal forces act as a rhythmic push on each basin, and how the basin responds depends on how close its own period is to the push's 12 hours 25 minutes. In parts of the Mediterranean the tidal range is only a few centimetres. In the Bay of Fundy in Canada, whose natural period is close to the tidal period, it exceeds 15 metres. Predicting the tide at a particular harbour therefore relies on long records of past tides as well as the theory.

::: history 1616–1687 · Rome and Cambridge
Galileo had his own theory of the tides, and he cared about it a great deal, because he believed it proved that Earth moves. In his account, the combination of Earth's spin and its motion around the Sun made each point on the surface speed up and slow down during the day, and the oceans sloshed in their basins as a result, like water in a carried bowl. Kepler had suggested instead that the Moon attracts the water, and Galileo dismissed this as an appeal to occult influences. His own theory predicted one high tide a day, not two, a difficulty he explained away with the shapes of the basins. Newton, in the *Principia* in 1687, gave the correct explanation from the difference in gravity across Earth. He accounted for the two daily tides, for spring and neap tides, and for the tide's dependence on the Moon's distance, and he even turned the argument round and used the relative sizes of lunar and solar tides to estimate the Moon's mass. His estimate came out about twice the true value, largely because local effects distort the measured tides, but the method was sound.
:::

## Tides elsewhere

Tidal forces stretch everything, not only oceans. The solid ground itself rises and falls twice a day, by up to a few tens of centimetres, though nobody notices because everything around rises with it; precise instruments, and the particle accelerators at CERN, must correct for it.

Tides also slowly reshape the Earth–Moon system. Water moving in the tides rubs against the sea floor, turning some of the energy of Earth's rotation into heat, so Earth's spin is gradually slowing and the day is lengthening by roughly 2 milliseconds per century. Because Earth spins faster than the Moon orbits, friction drags the tidal bulges slightly ahead of the line to the Moon, and the nearer bulge tugs the Moon forward along its orbit. That raises the Moon into a higher orbit: it moves about 3.8 cm farther from Earth each year, a figure measured by timing laser pulses bounced off reflectors that the Apollo astronauts left on the lunar surface.

Earth's tidal effect on the Moon is much stronger than the Moon's on Earth, because Earth is about 81 times more massive. Long ago, friction from these tides slowed the Moon's spin until it matched its orbit, so it now turns exactly once for each trip around Earth. That is why we always see the same face. Farther afield, the same 1/d³ effect becomes enormous near a black hole, where the difference in gravity between an infalling body's head and feet would be strong enough to stretch it apart.

::: quiz
? Why is there a high tide on the side of Earth facing away from the Moon?
- The Moon's gravity reflects off Earth's core
+ The far side is pulled less than Earth's centre, so relative to the centre it is left behind
- Earth's rotation flings the water outward on that side only
= Tides depend on how the pull at each point differs from the pull at Earth's centre. The near side is pulled more and the far side less, so the water bulges outward on both sides.
? The Sun's pull on Earth is far stronger than the Moon's. Why are solar tides smaller?
- The Sun is made of gas
+ Tidal effects go as M/d³, and the Sun's greater distance outweighs its greater mass
- Sunlight heats the ocean and flattens the bulge
= The pull itself goes as M/d², but its variation across Earth goes as M/d³. The extra power of distance favours the nearby Moon.
? Spring tides occur when
- it is spring in the northern hemisphere
+ the Sun and Moon are lined up with Earth, at new or full Moon
- the Moon is at first or last quarter
= When the Sun and Moon are lined up, their tidal bulges point the same way and add. At quarter Moon they are at right angles and partly cancel, giving neap tides.
:::
