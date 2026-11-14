Of all the predictions of general relativity, the one with the largest practical consequences needs no curved geometry at all. It follows from the equivalence principle and a little special relativity: a clock higher up in a gravitational field runs faster than an identical clock lower down. Your head ages slightly faster than your feet. The difference is far too small to feel, but the navigation system in your phone would drift off by kilometres every day if its engineers ignored it. This lesson derives the effect, tells how it was first measured in a university tower, and works through the numbers for satellite navigation.

## How can a rocket tell us about gravity?

The strategy is the one Einstein used in 1907. Gravity is hard to reason about directly, because we do not yet have a theory of it. An accelerating rocket in empty space is easy, because special relativity handles it completely. The equivalence principle says that nothing done inside a small closed room can tell the two situations apart. So we work out what happens in the rocket and conclude that the same thing must happen in gravity.

Here is the setup. A rocket of height h accelerates upward at a steady rate g, far from any planet. At the top sits a clock that emits light of a definite frequency f_top, a definite number of wave crests per second, and the light travels down to a receiver on the floor. We assume gh/c² is very small, which is true for any real rocket or building: for a 100 m tower on Earth it is about 10⁻¹⁴.

::: steps Frequency shift in an accelerating rocket
$ t = \frac{h}{c}
First, find how long the light takes to reach the floor. It covers the height h at speed c, so the travel time is h/c. The floor does move up a little to meet the light during the trip, but that shortens the time only by a fraction of order gh/c², which we have agreed is negligible.
---
$ v = g\,t = \frac{gh}{c}
While the light is on its way, the rocket keeps accelerating, so by the time the light arrives the floor is moving faster than the rocket was at the moment the light left the top. The extra speed is acceleration times time, g times h/c. This is the key fact: the receiver is moving toward the place the light came from.
---
$ f_{\text{bottom}} = f_{\text{top}}\left(1 + \frac{v}{c}\right)
A receiver moving toward a source meets the wave crests more often than they were sent, so it measures a higher frequency. This is the Doppler effect, the same effect that raises the pitch of an approaching siren. For speeds much less than c, the frequency goes up by the fraction v/c.
---
$ f_{\text{bottom}} = f_{\text{top}}\left(1 + \frac{gh}{c^2}\right)
Replace v with gh/c from the second step. The light arrives at the floor shifted to a higher frequency, a **blueshift**, by the fraction gh/c². Running the argument the other way, light sent up from floor to ceiling would arrive at a lower frequency, a **redshift**, by the same fraction.
---
$ f_{\text{bottom}} = f_{\text{top}}\left(1 + \frac{gh}{c^2}\right) \text{ at rest in gravity}
Now apply the equivalence principle. A room resting on a planet with surface gravity g cannot be distinguished from the rocket, so the floor of that room must also receive the light blueshifted by gh/c². But here nothing is moving, so the Doppler effect cannot be the explanation. Something else must be responsible, and the next section says what.
:::

## What does a frequency shift mean when nothing moves?

Think carefully about the room resting on the planet. Nothing in it changes with time: the clocks are bolted in place, and the planet's field is steady. Wave crests leave the top clock and travel down to the bottom one. They cannot pile up somewhere in between, because a steady situation cannot keep accumulating crests; nor can any be lost. So every crest that leaves the top arrives at the bottom, and over any long stretch, the number sent equals the number received.

Yet the bottom clock measures more crests per second than the top clock sent per second. The only way both statements can be true is if the two clocks disagree about how long a second is. A second on the bottom clock must contain more of the crests because it is a longer stretch of time, as judged by the top clock. In other words, the bottom clock is running slow.

::: equation Weak-field clock rate
\frac{\mathrm{d}\tau_{\text{top}}}{\mathrm{d}\tau_{\text{bottom}}} = 1 + \frac{gh}{c^2}
:::

Here τ means the time each clock actually records, its proper time. The quantity gh is the difference in gravitational potential energy per kilogram between the two clocks, written ΔΦ. So the general rule, valid whenever gravity is weak, is that a clock gains a fraction ΔΦ/c² over a clock lower in the potential.

How big is this on Earth? The ratio g/c² is about 1.1 × 10⁻¹⁶ per metre. So a clock one metre higher than another gains about 10⁻¹⁶ of a second every second. Over an 80-year lifetime, about 2.5 × 10⁹ seconds, that adds up to roughly a quarter of a microsecond. That is the sense in which your head ages faster than your feet.

It is important that this is not a defect of clocks. The argument never used anything about how the clock works. Any process can serve as a clock: a vibrating atom, a chemical reaction, a beating heart. If one of them ran slow at the bottom and another did not, comparing the two would tell you whether you were in a rocket or on a planet, and the equivalence principle says no experiment can do that.

::: note
Gravity changes the rate at which time itself passes, not just the behaviour of clocks. Every process, whether atomic, chemical or biological, runs slower deeper in a gravitational field, by the fraction ΔΦ/c² compared with a process higher up.
:::

## Could anyone measure one part in a thousand trillion?

For decades after Einstein's prediction, the answer seemed to be no, at least on Earth. Astronomers looked for the redshift in light leaving the Sun and from the dense white dwarf star Sirius B, but other effects in those stars' atmospheres muddied the results. A clean test needed a laboratory measurement of a frequency shift of a few parts in 10¹⁵, and no source of light had a frequency that sharp.

::: history 1959 · Harvard
In 1958 the German physicist Rudolf Mössbauer discovered that atoms locked firmly into a crystal can emit and absorb gamma rays without recoiling, and as a result the gamma rays have an astonishingly precise frequency. The emitted gamma rays from one piece of iron-57 are absorbed strongly by another piece only if the two frequencies match almost perfectly. Robert Pound at Harvard and his graduate student Glen Rebka realised this was the tool they needed. They put a gamma-ray source at the top of the 22.5 m tower of Harvard's Jefferson Laboratory and an absorber at the bottom. The predicted fractional shift, gh/c², was 2.5 × 10⁻¹⁵. Because the arriving gamma rays were blueshifted, they no longer quite matched the absorber, and fewer were absorbed. Pound and Rebka then moved the source with a carefully controlled speed of well under a micrometre per second, adding a small Doppler shift of their own, and found the speed at which the match was restored. That speed measured the gravitational shift. Their result, published in 1960, agreed with Einstein's prediction to within 10 percent. In 1964 Pound and Joseph Snider improved the agreement to about 1 percent.
:::

Modern atomic clocks have made the tower unnecessary. Optical clocks, which count the oscillations of light absorbed by single trapped atoms, are now so stable that in 2010 a team at the National Institute of Standards and Technology in Colorado detected the faster ticking of an aluminium-ion clock raised by just 33 centimetres. Today the best clocks can see the difference over a single millimetre.

## Why does GPS need relativity?

The Global Positioning System is the place where this effect stops being a curiosity. A GPS receiver works out where it is by timing signals from several satellites. Each signal carries the time at which it was sent; the receiver notes when it arrived, and the delay multiplied by the speed of light gives the distance to that satellite. With distances to four satellites, the receiver can solve for its position and for its own clock error. The catch is the speed of light: a timing error of one microsecond means a distance error of 300 metres. To locate you within a few metres, the satellites' clocks must be correct to a few nanoseconds.

The satellites orbit about 20,200 km above the ground, and two relativistic effects act on their clocks in opposite directions.

The first is the gravitational effect of this lesson. The satellites sit much higher in the Earth's gravitational potential than receivers on the ground, so their clocks run fast. At this height gravity is much weaker than at the surface, so the simple product gh will not do; instead one uses the full potential difference, GM(1/R − 1/r), where R is the Earth's radius and r the radius of the orbit. Divided by c², it comes to about 5.3 × 10⁻¹⁰. Over a day of 86,400 seconds, that is about 45 microseconds gained.

The second is the time dilation of special relativity. The satellites move at about 3.9 km/s, and a moving clock runs slow by a fraction of about v²/2c² when v is much less than c. That comes to about 8 × 10⁻¹¹, or about 7 microseconds lost per day.

::: equation Net drift of GPS clocks
45\,\mu\text{s} - 7\,\mu\text{s} \approx 38\,\mu\text{s per day}
:::

The gravitational effect wins. Left uncorrected, 38 microseconds a day multiplied by the speed of light would build into a ranging error of about 11 kilometres per day, and the positions it gave would be useless within hours. The correction is built into the satellites before launch: their clocks are set to tick at 10.22999999543 MHz instead of the nominal 10.23 MHz, so that once in orbit, running fast by exactly the predicted amount, they keep the right time as seen from the ground. Every time a phone finds its position, it relies on Einstein's 1907 argument being correct.

## Where does the simple rule stop working?

The rule ΔΦ/c² is an approximation for weak gravity, where the shift is a tiny fraction. Near a very compact object it is replaced by the exact result from the Schwarzschild solution: a clock held at distance r from the centre runs slow, compared with a distant clock, by the factor √(1 − r_s/r), where r_s = 2GM/c² is the Schwarzschild radius. Far from the mass this reduces to the weak-field rule. Close to r_s it does something the weak-field rule never could: it falls to zero, so that to a distant observer a clock at that radius appears to stop. The black-holes chapter picks up the story there.

::: model gravity-well
Move the probe deeper into the well and compare its clock with one far away. Far out, the rate follows the weak-field rule 1 + Φ/c², with Φ negative; close in, it falls much faster.
:::

::: quiz
? In the rocket argument, why does the light arrive at the floor blueshifted?
- The light speeds up as it falls
+ The floor has gained speed toward the source while the light was travelling
- The top clock runs slow
= By the time the light arrives, the floor is moving toward the source at v = gh/c. A receiver approaching the source meets crests more often, a Doppler shift of v/c = gh/c².
? Which effect makes GPS satellite clocks run fast compared with clocks on the ground?
- Their orbital speed
+ Their height in the Earth's gravitational potential
- The thinner atmosphere
= Being higher in the potential makes them gain about 45 μs per day. Their orbital speed makes them lose about 7 μs, leaving a net gain of about 38 μs per day.
? A clock is moved from the floor to a shelf 2 m higher. By about what fraction does its rate change?
- 10⁻⁹
+ 2 × 10⁻¹⁶
- 2 × 10⁻²¹
= The fraction is gh/c² = (9.8 × 2) / (9 × 10¹⁶) ≈ 2 × 10⁻¹⁶. The clock on the shelf runs faster.
:::
