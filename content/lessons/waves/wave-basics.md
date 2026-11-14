Drop a stone into a still pond and rings spread outward across the surface. Something is clearly travelling: the rings reach the far bank, and a floating leaf there will start to bob. Yet the water itself does not travel with the rings. The leaf moves up and down as each ring passes and then settles back where it was. What crosses the pond is not water but a disturbance, and with it, energy. That is what a **wave** is: a disturbance that travels through a material, carrying energy from one place to another without carrying the material along. The material a wave travels through is called its **medium**. This lesson sets out how to describe a wave, the simple relation that ties together its speed, its frequency and its wavelength, and what decides how fast a wave moves.

## How can something travel when nothing moves along?

The answer lies in how the parts of a medium are connected. Each bit of a rope, or of the water's surface, or of the air, is linked to its neighbours by forces. If you pull one bit out of place, it tugs on the next bit, which tugs on the one after, and so on. Each bit only moves a little and then returns, but the disturbance is handed on along the chain, slightly delayed at each step. The previous lessons showed that a bit of material pulled out of equilibrium by a restoring force oscillates in simple harmonic motion. A wave is that oscillation passed from bit to bit through the medium, each bit a little behind the one before it.

Waves come in two basic kinds, distinguished by the direction in which the medium moves. In a **transverse** wave the medium moves at right angles to the direction the wave travels. Tie one end of a rope to a wall and flick the other end: a hump runs along the rope toward the wall, but each piece of rope moves only up and then down as the hump passes. In a **longitudinal** wave the medium moves back and forth along the same line the wave travels. Stretch out a long, loose spring, such as a Slinky, and push and pull one end: bands where the coils are squeezed together alternate with bands where they are spread apart, and the bands run along the spring while each coil only shuffles forward and back. Sound in air is a longitudinal wave of this kind. Water waves are a mixture of the two: as a wave passes, water near the surface moves in roughly circular loops, forward at the crest and backward in the trough, which is why a floating object bobs and also sways slightly to and fro.

## What numbers describe a wave?

Consider the simplest regular wave, a smooth sine-shaped pattern travelling steadily along a rope. You can look at it in two ways. Take a photograph at one instant, and you see a shape repeating in space: crest, trough, crest, trough. Or fix your attention on a single point of the rope and watch it over time, and you see a repeating motion: up, down, up, down. A full description needs numbers for both views.

The **amplitude**, written A, is the greatest distance the medium moves from its resting position, the height of a crest above the undisturbed rope. It measures how strong the wave is: a wave with a larger amplitude carries more energy.

The **wavelength**, written λ (the Greek letter lambda), is the distance between one crest and the next. It is the length after which the shape in the photograph repeats.

The **period**, written T, is the time one point of the rope takes to go through one complete cycle of its motion. The **frequency**, written f, is the number of complete cycles per second, so f = 1/T. Frequency is measured in hertz (Hz), where one hertz is one cycle per second.

So wavelength describes the pattern in space, and period and frequency describe it in time. These are not independent, as the next section shows.

::: model wave
Change the wavelength and the frequency, and watch the speed. It is always the product of the two.
:::

## How are speed, frequency and wavelength related?

The relation follows from watching a single crest.

::: steps The wave equation v = fλ
$ \Delta x = \lambda \quad\text{in}\quad \Delta t = T
Fix your eye on one point of the rope that is at the top of a crest. In one period T, that point goes through one complete cycle and is back at the top again. But in that time the crest it was part of has moved on, and the next crest has arrived to take its place. So in one period, the pattern has moved forward by exactly the distance between neighbouring crests, one wavelength.
---
$ v = \frac{\lambda}{T}
The speed of anything moving steadily is the distance it covers divided by the time it takes. The crest covers a distance λ in a time T, so the wave's speed v is λ divided by T. This is the speed at which the pattern moves, which is not the same as the speed of any piece of the rope.
---
$ v = f\lambda
Since the frequency is the number of cycles per second, f = 1/T, and dividing by T is the same as multiplying by f. This form is the most useful one, because frequency is usually the quantity we are given or can hear. It says that the wave's speed equals the number of crests passing per second times the distance between them.
:::

The relation v = fλ looks like a formula for the speed, but in practice it is usually read the other way, because the three quantities are set by different things. The source sets the frequency. A tuning fork vibrating 440 times a second, the note A that orchestras tune to, pushes out 440 compressions into the air every second, and nothing about the air can change that number. The medium sets the speed. Sound in air at 20 °C travels at about 343 m/s, whether the sound is loud or soft, high or low. The wavelength then has no choice: it must be whatever makes the equation work, λ = v/f = 343/440 ≈ 0.78 m. A note an octave higher, at 880 Hz, has twice the frequency in the same air, so its wavelength is half as long, about 0.39 m.

The same relation holds for every kind of wave, which is what makes it so useful. Light travels at 3.00 × 10⁸ m/s. Green light has a wavelength of about 530 nanometres (530 billionths of a metre), so its frequency is 3.00 × 10⁸ divided by 530 × 10⁻⁹, about 5.7 × 10¹⁴ Hz: nearly six hundred million million oscillations every second. At the other end of the range, an FM radio station broadcasting at 100 MHz sends out waves 3 m long.

::: note
A wave carries energy, not matter: each bit of the medium oscillates in place while the disturbance moves on. Its speed, frequency and wavelength are tied by v = fλ, with the frequency set by the source and the speed set by the medium.
:::

## What decides how fast a wave travels?

If the medium sets the speed, what is it about the medium that matters? Think back to how a wave is passed along. A displaced bit of medium is pulled back by its connection to its neighbours, and resists being moved because it has mass. The stronger the pull back, the faster each bit snaps into motion and the sooner the disturbance is handed on, so a stiffer medium carries waves faster. The more mass each bit has, the more sluggishly it responds, so a heavier medium carries waves more slowly. Wave speed is always a contest between stiffness and inertia.

For a stretched string, the stiffness comes from the tension F, the force with which the string is pulled tight, and the inertia from its mass per unit length μ (the Greek letter mu), measured in kilograms per metre. The two combine like this:

::: equation Wave speed on a string
v = \sqrt{\frac{F}{\mu}}
:::

Tighten the string and the waves go faster; use a thicker, heavier string at the same tension and they go slower. That is exactly how stringed instruments are tuned and built, as the lesson on standing waves will show. Sound follows the same rule. Water and steel are far denser than air, which on its own would slow sound down, but they are stiffer still in proportion, so sound travels about 1,480 m/s in water and around 5,000 m/s or more in steel, many times faster than in air.

## How can a wave be written as a formula?

It is often useful to describe a wave not in words but as a formula giving the displacement y of the medium at every position x and every time t. For a sine wave of amplitude A moving in the +x direction, the formula is:

::: equation A travelling sine wave
y(x, t) = A\sin\left(\frac{2\pi}{\lambda}x - 2\pi f t\right)
:::

The two views of a wave are both built into this expression. Freeze time at a fixed t, and what remains is a sine curve in x that repeats every time x increases by λ, because that adds 2π to the angle: this is the photograph. Fix your position at a particular x, and what remains is a sine oscillation in t that repeats every time t increases by 1/f: this is the motion of one point. The minus sign between the two terms is what makes the pattern move toward +x. As t increases, the second term grows, so to keep the angle the same, and so to stay on the same crest, x must grow as well. The crest moves forward. With a plus sign instead, the wave would travel toward −x.

::: history 1747–1822 · Paris and Berlin
The mathematics of waves began with the vibrating string. In 1747 Jean le Rond d'Alembert, a Parisian mathematician, sent a paper to the Berlin Academy in which he wrote down the equation of motion for a stretched string and solved it. His solution had a surprising generality: any shape at all travelling to the right at speed v, added to any shape travelling to the left, satisfies the equation, which in modern notation reads y = f(x − vt) + g(x + vt). A wave, in other words, does not have to be a sine; on an ideal string, any pulse keeps its shape as it travels. This opened a long argument. Daniel Bernoulli claimed that every motion of a string could be built by adding together sine waves, and Leonhard Euler and d'Alembert both doubted it, on different grounds. The question was settled only in the next century, when Joseph Fourier showed, in work on the flow of heat begun in 1807 and published in 1822, that very general shapes can indeed be built from sines. That is one reason the sine wave is the standard example: understand sine waves, and you can build almost any other wave from them.
:::

## Real waves

The picture in this lesson is idealised. Real waves lose amplitude as they spread out, because the same energy is shared over a larger region, and as friction within the medium turns some of their energy into heat. In many media the speed depends a little on the frequency, so a pulse made of many frequencies gradually spreads out as it travels; this is called **dispersion**, and it is why a prism splits white light into colours. But the ideal picture is the right starting point for sound, light, earthquake waves and, much later in this course, the matter waves of quantum mechanics.

::: quiz
? A violinist plays a note of twice the frequency. Compared with the first note, the sound wave in the same air has
- twice the speed
+ half the wavelength
- twice the wavelength
= The speed is set by the air, not the violin, so it stays the same. If the frequency doubles, the wavelength λ = v/f must halve.
? As an ocean wave passes a floating buoy, the buoy
- travels with the wave to the shore
+ moves up and down, and slightly back and forth, but stays roughly in place
- stays perfectly still
= The wave carries energy past the buoy, but the water only oscillates about its resting position, and the buoy moves with the water.
? A wave on a string has frequency 50 Hz and wavelength 0.4 m. Its speed is
- 125 m/s
+ 20 m/s
- 0.008 m/s
= The speed is frequency times wavelength: 50 crests per second, each 0.4 m apart, give 50 × 0.4 = 20 m/s.
:::
