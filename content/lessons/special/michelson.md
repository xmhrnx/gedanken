Maxwell's equations give one speed for light, c. Nineteenth-century physicists took for granted that this must be its speed relative to some medium, as sound's speed is relative to air. They called the medium the **luminiferous ether**: invisible, filling all space, at rest in some absolute sense. Earth orbits the Sun at about 30 km/s, so it should be moving through the ether, and an instrument on Earth should feel an "ether wind". Albert Michelson built one sensitive enough to detect it. It found nothing.

## The idea

Light sent upwind and back should take a different time from light sent across the wind and back, just as a swimmer takes longer to cross a river and back upstream than across it. The effect depends on (v/c)², and for v = 30 km/s, v/c = 10⁻⁴, so (v/c)² = 10⁻⁸. Measuring a timing difference of one part in a hundred million needed a new instrument.

## The interferometer

Michelson's **interferometer** splits a beam of light in two with a half-silvered mirror. The two halves travel along perpendicular arms, each of length L, bounce off mirrors at the ends, and return to recombine. If they come back in step they reinforce; out of step they cancel. The recombined light shows a pattern of bright and dark fringes, and a change in the travel time of one beam by a fraction of a light period shifts the fringes sideways by that fraction of a fringe.

## The expected shift

Suppose the apparatus moves through the ether at speed v along one arm. Work in the ether's frame, where light always moves at c.

::: steps Travel times along the two arms
$ t_{\parallel} = \frac{L}{c - v} + \frac{L}{c + v} = \frac{2L}{c}\,\frac{1}{1 - v^2/c^2}
Along the arm parallel to the motion, the end mirror runs away from the light on the way out (closing speed c − v), and the start comes toward it on the way back (c + v). Add the times over a common denominator.
---
$ t_{\perp} = \frac{2L}{\sqrt{c^2 - v^2}} = \frac{2L}{c}\,\frac{1}{\sqrt{1 - v^2/c^2}}
Along the perpendicular arm, the mirror moves sideways while the light is in flight, so the light must travel on a slant. Its speed c is the hypotenuse; its component across the arm is √(c² − v²) by Pythagoras.
---
$ \Delta t = t_{\parallel} - t_{\perp} \approx \frac{2L}{c}\left[\left(1 + \frac{v^2}{c^2}\right) - \left(1 + \frac{v^2}{2c^2}\right)\right] = \frac{L}{c}\,\frac{v^2}{c^2}
For small x, 1/(1 − x) ≈ 1 + x and 1/√(1 − x) ≈ 1 + x/2. The parallel trip takes slightly longer.
---
$ N = \frac{2c\,\Delta t}{\lambda} = \frac{2L v^2}{\lambda c^2}
Nobody knows where the fringes would be with no ether wind, so the apparatus is rotated by 90°, swapping the roles of the arms. That doubles the change. Dividing the path difference 2cΔt by the wavelength λ gives the expected shift N in fringes.
:::

For the 1887 apparatus, with an effective arm length of about 11 m (folded by repeated mirror reflections), λ ≈ 500 nm and v/c = 10⁻⁴:

::: equation The expected fringe shift in 1887
N = \frac{2 \times 11\,\text{m} \times 10^{-8}}{5 \times 10^{-7}\,\text{m}} \approx 0.4
:::

Four-tenths of a fringe is easy to see.

::: history 1881–1887 · Potsdam and Cleveland
Michelson, a young US Navy officer studying in Europe, built his first interferometer in 1881 near Berlin, moving it to Potsdam to escape traffic vibrations. Its arms were 1.2 m and the expected shift only 0.04 fringe; he saw nothing, but the result was not convincing. In 1887, at the Case School in Cleveland, he and the chemist Edward Morley mounted a much larger version on a stone slab floating in a trough of mercury, so it could be turned smoothly. They watched as it rotated, at different times of day and in different months. The shift was certainly less than a twentieth of the expected 0.4 fringe, and probably less than a fortieth. Michelson considered the experiment a failure.
:::

## Explaining nothing

The null result was a crisis. Perhaps Earth dragged the ether along with it, so there was no wind at the surface. But that contradicted **stellar aberration**, the small yearly tilt in the apparent direction of stars that Bradley had explained in 1729 by Earth moving through undisturbed light. Perhaps Earth happened to be at rest in the ether. But six months later it moves the opposite way at 60 km/s relative to its earlier motion, and the result was still null.

In 1889 George FitzGerald, and in 1892 Hendrik Lorentz, proposed something stranger: that objects moving through the ether shrink along their motion by exactly the factor √(1 − v²/c²). That would shorten the parallel arm just enough to cancel the timing difference. Lorentz built an elaborate theory around it. It worked, but it looked like a patch designed to hide the ether from every experiment.

In 1905 Einstein took the opposite route: if no experiment can detect the ether, drop it, and take the constancy of the speed of light as a starting principle. How much the Michelson–Morley result itself influenced him is debated; he later said it played only a small role compared with Maxwell's theory.

::: note
If light moved through an ether, Earth's motion should have shifted Michelson and Morley's fringes by about 0.4. They saw less than a twentieth of that. Light's speed did not depend on the direction of Earth's motion.
:::

::: quiz
? Why did Michelson rotate the interferometer?
- To average out vibrations
+ Because the shift only shows up as a change when the arms swap roles relative to the ether wind
- To measure the speed of light directly
= There is no reference position without ether wind, so the experiment looks for fringes moving as the apparatus turns.
? How does the expected effect scale with Earth's speed v through the ether?
- With v
+ With v²
- With 1/v
= Upwind and downwind effects nearly cancel; only the second-order term, proportional to (v/c)², survives.
? What did the FitzGerald–Lorentz contraction propose?
- That Earth drags the ether along
- That light speeds up in the parallel arm
+ That the arm moving along the ether wind shrinks just enough to cancel the timing difference
= A contraction by √(1 − v²/c²) exactly offsets the extra time on the parallel arm.
:::
