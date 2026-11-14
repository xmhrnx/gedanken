An ambulance races past and its siren drops in pitch at the moment it goes by: higher as it approaches, lower as it recedes. The siren itself has not changed. What changes is how closely its sound waves are packed when they reach you. Understanding that effect starts with what sound is.

## What sound is

Sound is a longitudinal wave of pressure. A vibrating loudspeaker cone pushes on the air in front of it, squeezing it into a slightly denser, higher-pressure layer, then pulls back and leaves a thinner one. These **compressions** and **rarefactions** pass from molecule to molecule outward from the source. The air itself only jiggles back and forth. The pressure changes are tiny: ordinary conversation varies the pressure by a few hundredths of a pascal, against an atmospheric pressure of about 100,000 pascals.

Humans hear frequencies from about 20 Hz to 20,000 Hz, the upper limit falling with age. Pitch is how we perceive frequency; loudness corresponds roughly to amplitude. Below 20 Hz is infrasound, which elephants use; above 20 kHz is ultrasound, which bats use for echolocation and doctors for imaging.

## The speed of sound

In air at 20 °C sound travels at about 343 m/s, or roughly 1 km in 3 seconds; counting seconds from a lightning flash to the thunder and dividing by three gives the storm's distance in kilometres. The speed rises with temperature, by about 0.6 m/s per degree Celsius. Liquids and solids are much stiffer than air, so sound is faster in them: about 1,480 m/s in water and around 5,000 to 6,000 m/s in steel.

::: equation Speed of sound in a gas
v = \sqrt{\frac{\gamma P}{\rho}}
:::

Here P is the pressure, ρ the density and γ a number that for air is 1.4. With P = 101,000 Pa and ρ = 1.20 kg/m³, this gives 343 m/s.

::: history 1687–1816 · Cambridge and Paris
Newton calculated the speed of sound in the *Principia* using v = √(P/ρ), assuming the air's temperature stays constant as it is squeezed. He got about 298 m/s (modern recalculations give about 290), while measurements gave around 340, and he massaged his argument with dubious corrections to close the gap. The error stood for over a century until Pierre-Simon Laplace saw the reason in 1816: compressions happen too fast for heat to flow away, so squeezed air warms up and pushes back harder. That introduces the factor γ, and √1.4 ≈ 1.18 is exactly the missing 18%.
:::

## The Doppler effect

Now let the source move. A siren emits wave crests at a steady frequency f, one every period T = 1/f. If the siren is moving toward you at speed vₛ, each crest is emitted from a little closer to you than the last, so the crests bunch up ahead of it.

::: steps Frequency heard from a moving source
$ \lambda = vT
For a stationary source, crests are separated by the distance sound travels in one period.
---
$ \lambda' = vT - v_s T = (v - v_s)T
A source approaching at speed vₛ moves forward by vₛT between emitting one crest and the next, so in front of it the crests are closer together by that amount.
---
$ f' = \frac{v}{\lambda'} = \frac{v}{(v - v_s)T}
The sound still travels at v through the air; the frequency you hear is the speed divided by the shortened wavelength.
---
$ f' = f\,\frac{v}{v - v_s}
Use 1/T = f. For a receding source, the crests behind it are stretched instead: replace −vₛ with +vₛ.
:::

::: note
Motion of the source squeezes the waves ahead of it and stretches those behind. Approaching sources sound higher, receding ones lower, by a factor v/(v ∓ vₛ).
:::

Take a 700 Hz siren on an ambulance moving at 25 m/s (90 km/h). Approaching, you hear 700 × 343/318 ≈ 755 Hz; receding, 700 × 343/368 ≈ 652 Hz. The drop as it passes is about 14%, between two and three semitones on a piano.

A moving listener also hears a shift, for a slightly different reason: moving toward the source, you meet crests more often. The frequency heard is then f(v + vₒ)/v for a listener approaching at speed vₒ. For speeds much smaller than the speed of sound the two formulas give almost the same shift; they differ because sound needs air, and it matters which one is moving through it.

::: history 1842–1845 · Prague and Utrecht
Christian Doppler, a mathematics professor in Prague, proposed the effect in 1842 for both sound and light, hoping it explained the colours of double stars. That application was wrong: stellar colours come mostly from temperature, and stars move far too slowly to change colour visibly. The effect itself was right. In 1845 Christophorus Buys Ballot tested it in the Netherlands with horn players on an open railway carriage pulled past musicians with perfect pitch, who heard the notes rise and fall as predicted.
:::

## Faster than sound

As vₛ approaches v, the crests ahead of the source pile up on top of one another. A source moving faster than sound outruns its own waves entirely; they build into a cone-shaped shock wave trailing behind it, heard on the ground as a **sonic boom**. A supersonic jet drags this boom along its whole flight path, not just the moment it passes the speed of sound. A whip's crack is a small sonic boom from its tip.

The Doppler effect is now a workhorse tool: police radar and weather radar measure speeds from the shift in reflected radio waves, ultrasound scanners measure blood flow, and astronomers measure the motion of stars and galaxies from shifts in their light. For light the formula must be modified by relativity, a later chapter's subject.

::: quiz
? A siren approaches you, passes, and moves away at constant speed. What do you hear?
- A pitch that rises steadily the whole time
+ A steady higher pitch, a drop as it passes, then a steady lower pitch
- The same pitch throughout, getting louder then quieter
= At constant speed the shifted frequency is constant on approach and constant on recession; the switch happens as it passes.
? Why does the pitch of an approaching source rise?
- The source vibrates faster when moving
- The sound travels faster than 343 m/s through the air
+ The crests are emitted closer together, so they arrive more often
= The speed of sound through the air is unchanged; the source's motion shortens the wavelength ahead of it.
? Newton's predicted speed of sound was about 15% too low. What did he miss?
- The weight of the air
+ Air heats up when it is compressed quickly, which makes it stiffer
- The Earth's rotation
= Compressions in a sound wave are too fast for heat to escape, so pressure rises more than Newton assumed; Laplace's factor γ corrects this.
:::
