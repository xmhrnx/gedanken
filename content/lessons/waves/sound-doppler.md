An ambulance races past with its siren on, and at the moment it passes, the pitch drops: higher as it approaches, lower as it moves away. The siren itself has not changed its note; a passenger in the ambulance hears the same pitch throughout. What changes is how closely packed the sound waves are by the time they reach you. This shift, called the Doppler effect, is now used to measure the speed of everything from cars and storms to blood cells and galaxies. To understand it, we first need to be clear about what sound is and what sets its speed, and that question turned out to contain a puzzle that defeated Newton.

## What is sound?

Sound is a longitudinal wave of pressure in a medium such as air. Think of a loudspeaker cone. As it moves forward, it pushes on the air just in front of it, squeezing that air into a slightly denser layer at slightly higher pressure. As it moves back, it leaves behind a layer that is slightly thinner, at lower pressure. The squeezed layers are called **compressions** and the thinned ones **rarefactions**. Each layer pushes on the air next to it, which pushes on the next, so the pattern of compressions and rarefactions travels outward from the speaker. The air itself does not flow away from the speaker; each bit of air only jiggles back and forth a tiny distance as the wave passes, just as the lesson on wave basics described.

The pressure changes involved are remarkably small. In ordinary conversation, the pressure at your ear rises and falls by a few hundredths of a pascal, against an atmospheric pressure of about 100,000 pascals. The ear detects changes of less than one part in a million.

Our ears respond to frequencies from about 20 Hz to about 20,000 Hz, and the upper limit falls steadily with age. We perceive frequency as **pitch**: a higher frequency sounds higher. Loudness corresponds roughly to amplitude, the size of the pressure swings. Sound below the range of human hearing is called infrasound, and elephants use it to communicate over several kilometres. Sound above it is ultrasound, which bats use to locate insects by their echoes and doctors use to image the inside of the body.

## How fast does sound travel?

In air at 20 °C, sound travels at about 343 m/s, or roughly one kilometre in three seconds. This gives a simple way to judge the distance of a thunderstorm. Light from a lightning flash arrives almost instantly, while the thunder takes about three seconds per kilometre, so counting the seconds between flash and thunder and dividing by three gives the distance in kilometres.

The speed depends on the medium, and, as the lesson on wave basics explained, it is a contest between stiffness, which pushes displaced layers back, and inertia, which resists their motion. Liquids and solids are far stiffer than air, enough to outweigh their much greater density, so sound travels faster in them: about 1,480 m/s in water and between 5,000 and 6,000 m/s in steel. In air, the speed rises with temperature, by about 0.6 m/s for every degree Celsius. For a gas, the relationship can be written down exactly:

::: equation Speed of sound in a gas
v = \sqrt{\frac{\gamma P}{\rho}}
:::

Here P is the gas pressure, which measures how hard the gas pushes back when squeezed, and ρ (rho) is its density, which measures its inertia. The factor γ (gamma) is a number that depends on the kind of gas; for air it is 1.4. With P = 101,000 Pa and ρ = 1.20 kg/m³, typical of air at sea level at 20 °C, the formula gives 343 m/s, matching measurement. That factor γ has a story.

::: history 1687–1816 · Cambridge and Paris
Newton was the first to calculate the speed of sound from the properties of air, in the *Principia*. His reasoning was essentially right, but he assumed that as a layer of air is squeezed, its temperature stays the same. That gave him the formula v = √(P/ρ), without the γ, and a speed of about 298 m/s by his figures (modern recalculation with his formula gives about 290). Measurements, made by timing the flash and bang of distant cannon, gave something nearer 340 m/s. Newton could not account for the difference. In later editions he tried to close the gap with arguments that look contrived: that the solid particles of air transmit sound instantly across their own width, and that water vapour does not take part. The discrepancy stood for more than a century. Pierre-Simon Laplace found the real cause, publishing it in 1816. The compressions and rarefactions in a sound wave alternate hundreds of times a second, far too quickly for heat to flow between them. So a squeezed layer of air warms up, and warm air pushes back harder than Newton assumed; a rarefied layer cools and pushes back less. The air is effectively stiffer than Newton thought, and that extra stiffness is the factor γ. Since √1.4 ≈ 1.18, it raises Newton's speed by exactly the missing 18%.
:::

## Why does a moving source change pitch?

Now let the source move. A siren emits sound at a steady frequency f, which means it sends out one compression, one crest of the wave, every period T = 1/f. Once a crest has left the siren, it travels through the air at the speed of sound, v, regardless of how the siren is moving. The siren's motion cannot speed the sound up; it can only change where each crest starts from.

Suppose the siren is moving toward you at speed vₛ. It emits a crest, and that crest heads toward you. By the time the siren emits the next crest, one period later, it has moved forward a little, so the second crest starts from a point closer to you than the first did. The crests in front of the siren are therefore bunched closer together than they would be if it stood still. Closer crests, all travelling at the same speed, arrive at your ear more often, and you hear a higher frequency.

::: steps Frequency heard from a moving source
$ \lambda = vT
Start with a stationary source. In one period T, a crest travels a distance vT through the air before the next crest is emitted from the same spot. So the distance between crests, the wavelength, is vT. This is just v = fλ written in another form.
---
$ \lambda' = vT - v_s T = (v - v_s)T
Now let the source move toward you at speed vₛ. During one period, the first crest still travels vT, but the source chases after it and covers a distance vₛT. The next crest therefore starts vₛT closer to the first one, and the wavelength ahead of the source is shortened to (v − vₛ)T.
---
$ f' = \frac{v}{\lambda'} = \frac{v}{(v - v_s)T}
The sound still travels through the air at speed v, so the frequency you hear is the number of crests reaching you per second: the speed divided by the distance between crests. Because that distance is now shorter, more crests arrive each second.
---
$ f' = f\,\frac{v}{v - v_s}
Since 1/T is the source's own frequency f, the heard frequency is f multiplied by v/(v − vₛ), a factor greater than 1. For a source moving away from you, the crests behind it are spread out instead, by vₛT each, so the same argument gives v/(v + vₛ), a factor less than 1, and a lower pitch.
:::

::: note
A moving source squeezes the waves ahead of it closer together and stretches those behind it farther apart. An approaching source therefore sounds higher and a receding one lower, by the factor v/(v ∓ vₛ), even though the source's own frequency never changes.
:::

A worked example shows how large the effect is in everyday life. Take an ambulance siren at 700 Hz moving at 25 m/s, which is 90 km/h. As it approaches, you hear 700 × 343/(343 − 25) = 700 × 343/318 ≈ 755 Hz. As it recedes, you hear 700 × 343/(343 + 25) = 700 × 343/368 ≈ 652 Hz. The drop as it passes is about 14%, between two and three semitones on a piano, which is easily heard. Notice also that the pitch does not slide continuously while the ambulance approaches: at constant speed it is steady at 755 Hz, then changes quickly as the ambulance goes by, then is steady at 652 Hz.

The listener can move too, and that also changes the pitch, for a slightly different reason. If you move toward a stationary siren, the spacing of the crests in the air is unchanged, but you run into them more often than you would standing still. The frequency you hear is then f(v + vₒ)/v, where vₒ is your speed toward the source. For speeds much smaller than the speed of sound, the moving-source and moving-listener formulas give almost the same shift. They are not identical, because sound travels through air, so it matters whether it is the source or the listener that is moving relative to the air.

::: history 1842–1845 · Prague and Utrecht
Christian Doppler, a mathematics professor in Prague, proposed the effect in 1842. He applied it to both sound and light, and his real interest was in the stars: he hoped it would explain why some double stars appear coloured, one star bluer and the other redder, as they circle each other and move toward or away from us. That application was wrong. The colours of stars come mostly from their temperatures, and stars move far too slowly compared with light for the effect to change their colour visibly. But the principle was right, and it was soon tested with sound. In 1845 the Dutch scientist Christophorus Buys Ballot arranged for a group of horn players to ride on an open railway carriage pulled by a locomotive past musicians standing beside the line near Utrecht. The musicians, who had perfect pitch, heard the horns' notes rise as the train approached and fall as it moved away, by the amounts the formula predicted.
:::

## What happens at the speed of sound?

The formula for an approaching source has v − vₛ in the denominator, and as the source's speed approaches the speed of sound, that difference shrinks toward zero. Physically, the source is almost keeping up with its own crests, so each new crest is emitted nearly on top of the previous one, and the crests pile up into a single intense front of compressed air.

A source moving faster than sound outruns its waves entirely. Each crest still spreads out as a sphere from the point where it was emitted, but the source is always ahead of all of them. The spheres overlap along the surface of a cone trailing behind the source, and along that cone the pressure jumps abruptly. This is a **shock wave**, heard on the ground as a sharp double bang called a **sonic boom**. A common misunderstanding is that the boom happens only at the moment an aircraft breaks the sound barrier. In fact a supersonic aircraft drags its cone along for as long as it flies faster than sound, and the boom sweeps across the ground beneath its whole flight path. The crack of a whip is a small sonic boom: a wave travelling along the whip speeds up as the whip tapers, and its tip ends up moving faster than sound.

## The Doppler effect at work

The Doppler effect has become one of the most widely used tools in science and engineering. Police radar guns send out radio waves and measure the shift in frequency of the waves reflected from a moving car, which gives its speed. Weather radar does the same with raindrops, revealing winds inside a storm. Medical ultrasound scanners measure the shift of sound reflected from moving blood cells, showing how fast blood flows through the heart and arteries. And astronomers measure the shift in the light of stars and galaxies to find how fast they are moving toward or away from us, which is how the expansion of the universe was discovered. For light, the formula has to be modified, because light needs no medium and its speed is the same for every observer; working out the correct version requires relativity, the subject of a later chapter.

::: quiz
? A siren approaches you, passes, and moves away at constant speed. What do you hear?
- A pitch that rises steadily the whole time
+ A steady higher pitch, a drop as it passes, then a steady lower pitch
- The same pitch throughout, getting louder then quieter
= At constant speed the shifted frequency stays constant while the siren approaches, and a different constant value while it recedes. The switch from one to the other happens as it passes.
? Why does the pitch of an approaching source rise?
- The source vibrates faster when moving
- The sound travels faster than 343 m/s through the air
+ The crests are emitted closer together, so they arrive more often
= The speed of sound through the air does not change. The source's motion shortens the distance between crests ahead of it, so more crests reach the listener each second.
? Newton's predicted speed of sound was about 15% too low. What did he miss?
- The weight of the air
+ Air heats up when it is compressed quickly, which makes it effectively stiffer
- The Earth's rotation
= Compressions in a sound wave happen too fast for heat to escape, so the pressure rises more than Newton assumed. Laplace's factor γ accounts for this extra stiffness.
:::
