The sea rises and falls about twice a day, and the timing follows the Moon: high tide comes roughly 50 minutes later each day, just as moonrise does. The surprise is that there are two high tides, not one, and that the Sun, whose pull on Earth is about 180 times stronger than the Moon's, raises tides less than half as large. Both facts come from one idea: tides are caused not by gravity itself but by the *difference* in gravity across a body.

## Pull that varies across the Earth

Earth and Moon orbit their common centre of mass, so Earth as a whole is in free fall toward the Moon. Everything on Earth shares that fall, so the average pull of the Moon cannot be felt at all. What is left is the variation. The side of Earth facing the Moon is about 6,400 km closer than the centre, so it is pulled slightly harder than the centre; the far side is pulled slightly less.

Relative to Earth's centre, then, the near-side ocean is pulled toward the Moon and the far-side ocean is left behind, pulled away from the Moon. Both sides bulge outward. As Earth turns under the two bulges, a coast passes through both each day, which gives two high tides. Because the Moon moves along its orbit in the meantime, a full cycle takes about 24 h 50 min, so successive high tides are about 12 h 25 min apart.

## How large is the effect?

Compare the Moon's pull at Earth's centre, a distance d from the Moon, with its pull at the near surface, a distance d − R.

::: steps The tidal acceleration
$ \Delta a = \frac{GM}{(d-R)^2} - \frac{GM}{d^2}
The extra acceleration toward the Moon on the near side compared with the centre. M is the Moon's mass, R Earth's radius.
---
$ \frac{GM}{(d-R)^2} = \frac{GM}{d^2}\left(1 - \frac{R}{d}\right)^{-2}
Factor out d² from the first term. For the Moon, R/d is about 1/60, a small number.
---
$ \left(1 - \frac{R}{d}\right)^{-2} \approx 1 + \frac{2R}{d}
For small x, (1 − x)⁻² ≈ 1 + 2x. With x = 1/60 the error is under 3%.
---
$ \Delta a \approx \frac{GM}{d^2}\cdot\frac{2R}{d} = \frac{2GMR}{d^3}
Subtract GM/d². The leading terms cancel and the difference is what remains.
:::

The tidal effect falls off as 1/d³, one power faster than gravity itself. Distance matters more than mass.

::: note
Tides come from the difference in gravity between one side of a body and the other. That difference goes as M/d³, so the nearby Moon beats the massive but distant Sun.
:::

Putting in numbers: for the Moon (GM = 4.90 × 10¹² m³/s², d = 3.84 × 10⁸ m) the tidal acceleration is about 1.1 × 10⁻⁶ m/s², roughly a ten-millionth of g. For the Sun (GM = 1.33 × 10²⁰ m³/s², d = 1.50 × 10¹¹ m) it is about 5 × 10⁻⁷ m/s². The Sun's tidal effect is about 46% of the Moon's.

## Spring and neap tides

When Sun, Moon and Earth line up, at new and full Moon, the solar and lunar bulges add, giving larger **spring tides** (the name has nothing to do with the season). At first and last quarter the Sun pulls at right angles to the Moon and partly cancels it, giving smaller **neap tides**. The ratio of the two is roughly (1 + 0.46)/(1 − 0.46), close to 3.

The two-bulge picture is an idealisation. It predicts an open-ocean tide of about half a metre, but continents block the bulges from sweeping round, and each ocean basin sloshes with its own natural period. Real tides are the basins' response to this rhythmic forcing: almost nothing in parts of the Mediterranean, more than 15 metres in the Bay of Fundy, where the basin's natural period is close to the tidal period.

::: history 1616–1687 · Rome and Cambridge
Galileo thought tides were the oceans sloshing as Earth spun and orbited, and offered them as proof that Earth moves. He dismissed Kepler's suggestion that the Moon attracts the water as an occult fancy. His theory predicted one tide a day, not two, and he waved the difficulty away. Newton, in the *Principia* (1687), gave the correct explanation from differential gravity, accounted for spring and neap tides, and even used the tides to estimate the Moon's mass, though his value came out about twice too large.
:::

## Tides elsewhere

Tidal forces stretch everything, not only oceans. The solid ground itself rises and falls twice a day, by up to a few tens of centimetres. Friction in the moving water slowly drains energy from Earth's spin, lengthening the day by roughly 2 milliseconds per century, and the bulges tug the Moon forward in its orbit, so it moves about 3.8 cm farther away each year, measured by bouncing lasers off reflectors the Apollo astronauts left behind. Earth's much stronger tides on the Moon long ago stopped the Moon's own spin relative to Earth, which is why we always see the same face. Near a black hole, the same 1/d³ difference becomes strong enough to stretch a falling body apart.

::: quiz
? Why is there a high tide on the side of Earth facing away from the Moon?
- The Moon's gravity reflects off Earth's core
+ The far side is pulled less than Earth's centre, so relative to the centre it is left behind
- Earth's rotation flings the water outward on that side only
= Tides depend on the difference from the pull at the centre: more on the near side, less on the far side, so both bulge.
? The Sun's gravitational pull on Earth is far stronger than the Moon's. Why are solar tides smaller?
- The Sun is made of gas
+ Tidal effects go as M/d³, and the Sun's greater distance outweighs its greater mass
- Sunlight heats the ocean and flattens the bulge
= Gravity goes as M/d², but its variation across Earth goes as M/d³, which favours the nearby Moon.
? Spring tides occur when
- it is spring in the northern hemisphere
+ the Sun and Moon are lined up with Earth, at new or full Moon
- the Moon is at first or last quarter
= Lined up, the solar and lunar tides add; at quarter Moon they partly cancel.
:::
