By the 1870s, Maxwell's theory of electricity and magnetism had shown that light is an electromagnetic wave, and his equations gave its speed: a single number, c, about 3.00 × 10⁸ m/s. But a speed has to be measured relative to something. Sound travels at about 340 m/s relative to the air; ripples move at their speed relative to the water. Every wave anyone had studied was a disturbance in some material, so physicists took it for granted that light must be too. They called its medium the **luminiferous ether**, the light-carrying ether: invisible, filling all of space, and at rest in some absolute sense. This lesson is about the attempt to detect Earth's motion through that ether, the instrument built to do it, and the result, which was nothing at all, and which turned out to matter more than almost any positive result in physics.

## What should Earth's motion through the ether do?

Earth orbits the Sun at about 30 km/s, and its direction changes through the year. Whatever the ether's state, Earth cannot be at rest in it all the time. So in a laboratory on Earth there should be an "ether wind", just as a cyclist on a still day feels a wind blowing in her face. And light, which travels at c relative to the ether, should travel at different speeds relative to the laboratory depending on whether it goes with this wind, against it, or across it.

The difficulty is how small the effect is. You cannot simply time light going one way, because you would need two clocks at different places synchronised to a billionth of a second, which was impossible at the time. You can time a round trip with one clock, by sending light out to a mirror and back. But on a round trip the effects nearly cancel: what light loses going upwind it mostly regains coming back downwind. What survives is an effect proportional not to v/c, the ratio of Earth's speed to light's speed, but to its square. With v = 30 km/s, v/c = 10⁻⁴, and (v/c)² = 10⁻⁸. Detecting the ether wind meant comparing two travel times to one part in a hundred million.

A swimmer in a river gives the idea. Swimming a certain distance upstream and back takes longer than swimming the same distance straight across the river and back, even though the river helps on the return leg; the time lost going upstream outweighs the time gained coming down. Light in the ether wind should behave the same way: a trip along the wind and back should take slightly longer than a trip across it.

## How can one compare two travel times so precisely?

Albert Michelson's answer was to let light time itself. In his **interferometer**, a beam of light strikes a half-silvered mirror, which lets half the light through and reflects the other half, splitting the beam in two. The two halves travel along arms at right angles to each other, each of length L, reflect off mirrors at the ends, and come back to the half-silvered mirror, where they are recombined and sent toward a viewer.

When the two returning beams are in step, crest meeting crest, they reinforce each other; when they are out of step by half a wave, they cancel. Because the mirrors are never perfectly aligned, the viewer sees a pattern of bright and dark bands called fringes, and the important thing is where those fringes sit. If one beam's travel time changes by a fraction of one period of the light wave, the fringes slide sideways by that same fraction of the spacing between them. A period of visible light is about 2 × 10⁻¹⁵ s. An instrument that can detect a shift of a tenth of a fringe is, in effect, a stopwatch with a resolution of a fraction of a femtosecond.

## How large a shift should there be?

Suppose the apparatus moves through the ether at speed v, along the direction of one arm. It is easiest to work in the ether's own frame, where light always moves at c, and follow each beam.

::: steps Travel times along the two arms
$ t_{\parallel} = \frac{L}{c - v} + \frac{L}{c + v}
Take the arm pointing along the motion first. On the way out, the end mirror is running away from the light, so the light closes the distance L at only c − v. On the way back, the half-silvered mirror is moving toward the light, so the gap closes at c + v. The total time is the sum of the two legs.
---
$ t_{\parallel} = \frac{2Lc}{c^2 - v^2} = \frac{2L}{c}\,\frac{1}{1 - v^2/c^2}
Add the two fractions over the common denominator (c − v)(c + v) = c² − v². The numerator becomes L(c + v) + L(c − v) = 2Lc. Then divide top and bottom by c² to show the result as the stationary round-trip time 2L/c times a correction factor that is slightly bigger than 1.
---
$ t_{\perp} = \frac{2L}{\sqrt{c^2 - v^2}} = \frac{2L}{c}\,\frac{1}{\sqrt{1 - v^2/c^2}}
Now the arm across the motion. While the light is in flight, its end mirror moves sideways, so to hit it the light must travel on a slant, and the same on the way back. Its total speed is still c, but only part of that velocity carries it along the arm: by Pythagoras, the component across the arm is √(c² − v²). Covering 2L at that speed gives this time, again 2L/c times a correction factor, but a smaller one, because it is a square root.
---
$ \Delta t = t_{\parallel} - t_{\perp} \approx \frac{2L}{c}\left[\left(1 + \frac{v^2}{c^2}\right) - \left(1 + \frac{v^2}{2c^2}\right)\right] = \frac{L}{c}\,\frac{v^2}{c^2}
Because v²/c² is so small, the correction factors can be replaced by simpler approximations: for a small number x, 1/(1 − x) ≈ 1 + x and 1/√(1 − x) ≈ 1 + x/2. Subtracting, the 1s cancel and leave a difference of half of v²/c². The trip along the motion takes longer than the trip across it, by an amount proportional to (v/c)², just as the swimmer's did.
---
$ N = \frac{2c\,\Delta t}{\lambda} = \frac{2Lv^2}{\lambda c^2}
There is no way to know where the fringes would sit if there were no ether wind, so a single reading tells you nothing. Instead the apparatus is rotated by 90°, which swaps the roles of the two arms and reverses the time difference. The fringes should therefore move by an amount corresponding to twice Δt, a path difference of 2cΔt. Dividing by the wavelength λ converts that path difference into a number N of fringe spacings.
:::

The prediction can be checked with the numbers of the decisive experiment of 1887. Its arms were made effectively about 11 m long by bouncing the light back and forth between several mirrors, the light had a wavelength of about 500 nm, and v/c was taken as 10⁻⁴ from Earth's orbital speed.

::: equation The expected fringe shift in 1887
N = \frac{2 \times 11\,\text{m} \times 10^{-8}}{5 \times 10^{-7}\,\text{m}} \approx 0.4
:::

Four-tenths of a fringe is a large, easy shift to see. The instrument could detect a hundredth of a fringe. If the ether wind existed, the experiment could not miss it.

::: history 1881–1887 · Berlin, Potsdam and Cleveland
Michelson was a young officer of the United States Navy who had already made a name measuring the speed of light. On leave to study in Europe, he built his first interferometer in 1881 in Berlin, in Hermann von Helmholtz's laboratory, but the vibrations of city traffic shook the fringes so badly that he moved it to the astrophysical observatory at Potsdam. Its arms were only 1.2 m long. He expected a shift of about 0.04 fringe, saw none, and concluded that there was no ether wind. Others were unconvinced: the expected shift was close to what the instrument could detect, and first Alfred Potier and then Hendrik Lorentz pointed out that Michelson had overlooked the effect on the cross arm, which halved the true prediction. Back in America, at the Case School of Applied Science in Cleveland, Michelson teamed up with Edward Morley, a chemist at the neighbouring Western Reserve University and a meticulous experimenter. They mounted a much larger instrument on a heavy sandstone slab that floated on a trough of mercury, so it could be turned slowly and smoothly without stopping, and watched the fringes as it rotated, at noon and in the evening, over several days in July 1887. Whatever shift there was, it was certainly less than a twentieth of the expected 0.4 fringe, and probably less than a fortieth. Michelson regarded the experiment as a failure, because it had not measured what he set out to measure.
:::

## How do you explain a null result?

The failure to find the ether wind was a genuine crisis, because each obvious way out ran into other evidence.

Perhaps Earth drags the ether along with it near its surface, as a moving ball drags a thin layer of air, so that there is no wind in the laboratory. But that contradicts **stellar aberration**, a small yearly back-and-forth shift in the apparent positions of all stars, which James Bradley had discovered and explained in 1729. Bradley's explanation was that the telescope must be tilted slightly forward to catch light falling from a star, in the way you tilt an umbrella forward when walking through vertical rain; the tilt reverses as Earth's motion reverses during the year. That explanation works only if the light reaching the telescope is not being carried along with Earth. Perhaps, then, Earth just happened to be at rest in the ether in July 1887. But six months later Earth moves in the opposite direction, 60 km/s different from before, and later repetitions at other seasons, by Morley and others, still found no wind of the expected size.

In 1889 the Irish physicist George FitzGerald, and independently in 1892 the Dutch physicist Hendrik Lorentz, proposed something stranger. Suppose that any object moving through the ether is shortened along its direction of motion, by exactly the factor √(1 − v²/c²). Then the arm pointing into the ether wind is slightly shorter than the cross arm, by just enough to make up for the extra time light spends on it, and the two travel times come out equal. Lorentz argued that this might be real: if the forces holding the atoms of a solid together are electromagnetic, they might change with motion through the ether in just this way. He went on to build an elaborate theory in which several such effects conspired to hide the ether from every possible experiment. It worked, but it had an uncomfortable look about it. An ether that exists but that nature conspires to make undetectable is hard to tell apart from no ether at all.

In 1905 Albert Einstein took that thought seriously. If no experiment can detect motion through the ether, then perhaps there is nothing to detect. He dropped the ether, and took as a starting principle that light travels at the same speed c for every observer moving steadily, whatever their motion. The shortening that FitzGerald and Lorentz had proposed reappeared in his theory, but as a consequence of how space and time are measured, not as a squeezing caused by an ether wind. How much the Michelson–Morley result itself guided Einstein is still debated by historians. He later said that it had played only a small part in his thinking, compared with the structure of Maxwell's theory itself. But for most physicists, it was the experiment that made the new theory necessary.

::: note
If light travelled through an ether, Earth's motion should have shifted Michelson and Morley's fringes by about 0.4 of a fringe as their instrument turned. They saw less than a twentieth of that. The speed of light did not depend on the direction of Earth's motion, and the attempts to save the ether ended by making it undetectable in principle.
:::

::: quiz
? Why did Michelson and Morley rotate their interferometer?
- To average out vibrations from the building
+ Because the expected shift shows up only as a change when the two arms swap roles relative to the ether wind
- To measure the speed of light directly
= There is no reference reading taken without an ether wind, so a single position of the fringes means nothing. Turning the instrument by 90° should reverse the time difference between the arms and make the fringes move.
? How does the expected time difference depend on Earth's speed v through the ether?
- In proportion to v
+ In proportion to v²
- In proportion to 1/v
= On a round trip, the time lost against the wind is almost exactly regained with it. Only an effect of order (v/c)² survives, which is why the measurement needed precision of one part in a hundred million.
? What did FitzGerald and Lorentz propose to explain the null result?
- That Earth drags the ether along with it
- That light speeds up in the arm pointing along the motion
+ That the arm moving along the ether wind shrinks by just enough to cancel the timing difference
= A contraction of the parallel arm by the factor √(1 − v²/c²) exactly makes up for the extra time light spends on it, so both round trips take the same time.
:::
