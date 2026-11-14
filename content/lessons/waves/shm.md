Pull a mass on a spring aside and let go: it oscillates back and forth at a steady rhythm. The same motion appears in a swinging pendulum, a vibrating guitar string, a quartz crystal in a watch and atoms jiggling in a solid. It is called **simple harmonic motion**, and it follows whenever the force pulling something back toward equilibrium is proportional to how far it has been displaced.

## The restoring force

Stretch or compress a spring by a distance x from its relaxed length, and it pushes back with a force proportional to x:

::: equation Hooke's law
F = -kx
:::

k is the **spring constant**, measured in N/m: a stiff spring has a large k. The minus sign says the force always points back toward equilibrium, opposite to the displacement. Hooke's law holds only for small stretches; pull a spring too far and it deforms permanently.

::: history 1676–1678 · London
Robert Hooke announced his law in 1676 as an anagram, "ceiiinosssttuv", a common way to claim priority without revealing a result. Two years later he published the solution: *ut tensio, sic vis*, "as the extension, so the force". He used it in designing balance springs for watches, the key to making them keep time, and fought a bitter priority dispute with Christiaan Huygens over that invention.
:::

## The equation of motion

Put Hooke's law into Newton's second law for a mass m on a spring, on a frictionless surface, and look for the motion.

::: steps Solving for the motion of a mass on a spring
$ ma = -kx \quad\Rightarrow\quad a = -\frac{k}{m}\,x
Newton's second law with F = −kx. The acceleration is proportional to the displacement and opposite to it.
---
$ \omega^2 = \frac{k}{m}, \qquad a = -\omega^2 x
Name the constant ω² for brevity. We need a function x(t) whose second derivative (the rate of change of its rate of change) is −ω² times itself.
---
$ x = A\cos(\omega t)
Try a cosine. A is the **amplitude**, the largest displacement. At t = 0 the mass is at x = A, released from rest.
---
$ v = -A\omega\sin(\omega t), \qquad a = -A\omega^2\cos(\omega t) = -\omega^2 x
Differentiate twice. The slope of cos(ωt) is −ω sin(ωt), and the slope of sin(ωt) is ω cos(ωt). The result is −ω² times x, so the guess works for any amplitude A.
---
$ T = \frac{2\pi}{\omega} = 2\pi\sqrt{\frac{m}{k}}
The cosine repeats when ωt advances by 2π, so the **period** T, the time for one full oscillation, is 2π/ω. ω is called the angular frequency.
:::

A heavier mass oscillates more slowly, a stiffer spring faster. A 1 kg mass on a spring with k = 100 N/m has ω = 10 rad/s and T ≈ 0.63 s. The frequency f = 1/T, the number of oscillations per second, is measured in hertz (Hz).

::: note
When the restoring force is proportional to displacement, the motion is a sine wave in time, and the period, 2π√(m/k), does not depend on the amplitude.
:::

## Amplitude does not matter

The amplitude A cancelled. Pull the spring twice as far and the force is twice as large, so the mass moves twice as fast and covers twice the distance in the same time. This property, called **isochronism**, is what makes harmonic oscillators good clocks: the rhythm stays fixed as friction slowly reduces the swing.

## A shadow of circular motion

There is a neat way to see where the sine comes from. A point moving round a circle of radius A at constant angular speed ω has x-coordinate A cos(ωt). Its centripetal acceleration, ω²A pointing toward the centre, has an x-component of −ω²x. So simple harmonic motion is exactly the shadow, or projection, of uniform circular motion onto a line. The mass is fastest passing through the middle, where the shadow moves at full speed ωA, and momentarily at rest at each end.

## Energy

As the mass oscillates, energy passes back and forth between kinetic energy ½mv² and elastic potential energy ½kx², the work done stretching the spring. At the ends it is all potential; in the middle it is all kinetic. With no friction the total stays fixed:

::: equation Energy in simple harmonic motion
\tfrac{1}{2}mv^2 + \tfrac{1}{2}kx^2 = \tfrac{1}{2}kA^2
:::

## Why it is everywhere

Almost any system near a stable equilibrium behaves this way for small displacements. Near the bottom of any smooth valley, the potential energy curve looks like a parabola, ½kx², and a parabolic potential means a force −kx. That is why molecules vibrate, bridges sway and pendulums swing harmonically, as long as the motion is small. Real oscillators also lose energy to friction (damping), so their amplitude decays, but the period stays almost unchanged while it does.

::: quiz
? A mass on a spring oscillates with a period of 2 s. You double the amplitude. The new period is
- 4 s
+ 2 s
- 2√2 s
= In simple harmonic motion the period, 2π√(m/k), does not depend on amplitude.
? You replace the mass with one four times as heavy, keeping the same spring. The period
+ doubles
- quadruples
- halves
= T = 2π√(m/k), so multiplying m by 4 multiplies T by √4 = 2.
? Where is the mass moving fastest?
- At the ends of its swing
+ Passing through the equilibrium position
- Halfway between equilibrium and the end
= At equilibrium all the energy is kinetic; at the ends it is all stored in the spring and the mass is momentarily at rest.
:::
