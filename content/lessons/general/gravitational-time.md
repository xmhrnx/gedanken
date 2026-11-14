Of all the predictions of general relativity, the one with the largest everyday consequences comes straight from the equivalence principle, with no curved geometry needed: clocks higher up in a gravitational field run faster than clocks lower down. Your head ages faster than your feet. The effect is tiny, but the satellite navigation in your phone would drift off by kilometres a day without correcting for it.

## The argument with a rocket

Use the equivalence principle in the usual way: work out what happens in an accelerating rocket, where only special relativity is needed, then conclude the same happens in gravity.

A rocket of height h accelerates upward at g in empty space. A clock at the top emits light of frequency f_top toward a receiver at the floor. Assume gh/c² is very small, as it is for any real rocket or building.

::: steps Frequency shift in an accelerating rocket
$ t = \frac{h}{c}
The light takes about this long to reach the floor. (The floor moves a little toward it meanwhile, but that changes t only by a fraction of order gh/c², which we neglect.)
---
$ v = g\,t = \frac{gh}{c}
While the light travels, the floor gains this much speed, compared with the rocket's speed when the light left the top.
---
$ f_{\text{bottom}} = f_{\text{top}}\left(1 + \frac{v}{c}\right)
The receiver is moving toward the source at speed v relative to the frame of emission, so it sees a Doppler blueshift. For v much less than c the shift is v/c.
---
$ f_{\text{bottom}} = f_{\text{top}}\left(1 + \frac{gh}{c^2}\right)
Substitute v = gh/c. Light sent down is received blueshifted; light sent up would be received redshifted by the same fraction.
---
$ f_{\text{bottom}} = f_{\text{top}}\left(1 + \frac{gh}{c^2}\right) \text{ in a gravitational field}
By the equivalence principle, a rocket resting on a planet with surface gravity g must show the same shift. Here nothing is moving, so there is no Doppler effect to blame.
:::

Now think about what the shift means in a static field. Nothing in the setup changes with time, so wave crests cannot pile up between top and bottom: every crest sent from the top arrives at the bottom. If the bottom clock counts more crests per second than the top clock sent per second, then a "second" of the bottom clock must be longer. The bottom clock runs slow.

::: equation Weak-field clock rate
\frac{\mathrm{d}\tau_{\text{top}}}{\mathrm{d}\tau_{\text{bottom}}} = 1 + \frac{gh}{c^2}
:::

More generally gh is the difference in gravitational potential ΔΦ between the two clocks, so the rule is: a clock gains a fraction ΔΦ/c² over a clock lower in the potential. At Earth's surface g/c² ≈ 1.1 × 10⁻¹⁶ per metre. A clock 1 m higher gains about 10⁻¹⁶ of a second every second, or about a quarter of a microsecond over an 80-year lifetime.

::: note
Gravity changes the rate of time itself, not just the clocks. Every process, atomic, chemical or biological, runs slower deeper in a gravitational well, by the fraction ΔΦ/c².
:::

## Testing it in a tower

::: history 1959 · Harvard
Robert Pound and his student Glen Rebka measured the shift over the 22.5 m height of Harvard's Jefferson Laboratory tower. The predicted fractional shift, gh/c², was 2.5 × 10⁻¹⁵. They could only see something so small because of the Mössbauer effect, discovered in 1958: atoms of iron-57 locked in a crystal emit and absorb gamma rays of an extraordinarily sharp frequency. Gamma rays sent down the tower arrived slightly blueshifted, and no longer matched the absorber at the bottom; moving the source at a carefully controlled speed of under a micrometre per second added a Doppler shift that restored the match and so measured the gravitational one. Their result, published in 1960, agreed with Einstein within 10 percent, and Pound and Joseph Snider improved this to 1 percent in 1964.
:::

Modern optical clocks have made the tower unnecessary. In 2010 a group at NIST in Colorado detected the faster ticking of an aluminium-ion clock raised by just 33 cm.

## GPS

A GPS receiver finds its position by timing signals from satellites: each signal says when it was sent, and the delay, times c, gives the distance. A timing error of 1 μs means a ranging error of 300 m, so the satellite clocks must be right to a few nanoseconds.

The satellites orbit about 20,200 km up, where two relativistic effects compete:

- **Gravity**: they sit higher in Earth's potential, so their clocks run fast by ΔΦ/c² ≈ 5.3 × 10⁻¹⁰, about **45 μs per day**. (At this height g varies a lot, so the full potential difference GM(1/R_Earth − 1/r) is used, not gh.)
- **Speed**: they move at about 3.9 km/s, so special-relativistic time dilation slows them by v²/2c² ≈ 8 × 10⁻¹¹, about **7 μs per day**.

::: equation Net drift of GPS clocks
45\,\mu\text{s} - 7\,\mu\text{s} \approx 38\,\mu\text{s per day}
:::

Left uncorrected, 38 μs a day would build into a position error of about 11 km per day. The correction is built in before launch: the satellite clocks are set to tick at 10.22999999543 MHz instead of the nominal 10.23 MHz, so that once in orbit they run at the right rate as seen from the ground.

## Looking ahead

The rule ΔΦ/c² is the weak-field limit. Near a black hole the exact Schwarzschild factor √(1 − r_s/r) takes over, and at the horizon a distant observer sees a falling clock slow to a stop. The black-holes chapter picks up the story there.

::: model gravity-well
Move the probe deeper into the well and watch its clock rate fall compared with a clock far away. Far out, the rate follows the weak-field rule 1 + Φ/c².
:::

::: quiz
? In the rocket argument, why is the light received at the floor blueshifted?
- The light speeds up as it falls
+ The floor has gained speed toward the source while the light was in transit
- The top clock runs slow
= The receiver moves toward the source at v = gh/c by the time the light arrives, giving a Doppler shift v/c = gh/c².
? Which effect makes GPS satellite clocks run fast compared with clocks on the ground?
- Their orbital speed
+ Their height in Earth's gravitational potential
- The thinner atmosphere
= Being higher in the potential gains about 45 μs per day; their speed loses about 7, leaving a net gain of about 38.
? A clock is moved from the floor to a shelf 2 m higher. By about what fraction does its rate change?
- 10⁻⁹
+ 2 × 10⁻¹⁶
- 2 × 10⁻²¹
= gh/c² = 9.8 × 2 / (9 × 10¹⁶) ≈ 2 × 10⁻¹⁶, and the higher clock runs faster.
:::
