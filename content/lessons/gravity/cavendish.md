Newton's law of gravitation contains a constant, G, that he never knew. Orbits give the product GM for any body with something circling it, but not G or M separately. Without G, nobody could say what Earth, the Sun or any planet actually weighs. Measuring it means detecting the gravitational pull between two objects small enough to handle in a laboratory, a force so feeble it is easily swamped by a draught.

## Why it is hard

Two 1 kg masses 10 cm apart attract with a force of about 7 × 10⁻⁹ N, the weight of less than a microgram. Every object in the room, including the experimenter, pulls on the apparatus too, and air currents from a warm body push on it far harder. The measurement needs a detector that responds to tiny sideways forces while ignoring Earth's huge downward one.

## The torsion balance

The answer is a **torsion balance**: a light rod hung at its centre from a thin wire, so that it can turn in a horizontal plane. Gravity pulls the rod down, which the wire simply supports; a sideways force on the ends twists the wire, and a thin wire takes very little force to twist. Henry Cavendish hung a rod 1.8 m long with a lead ball of about 0.73 kg on each end. Next to the small balls he placed two large lead balls of about 158 kg, one on opposite sides, so their attractions twisted the rod the same way.

The twist was measured from the rod's swing. Left alone, the rod oscillates slowly back and forth about its rest position; Cavendish's took about 7 minutes per swing. The period reveals how stiff the wire is, and the shift in the rest position when the large balls are moved reveals the force.

::: steps Extracting G from the torsion balance
$ \kappa\theta = FL
In equilibrium the wire's restoring torque, κθ for a twist θ, balances the gravitational torque. Each small ball feels force F at distance L/2 from the wire, and there are two of them, so the torque is FL.
---
$ T = 2\pi\sqrt{\frac{I}{\kappa}}, \qquad I = 2m\left(\frac{L}{2}\right)^2 = \frac{mL^2}{2}
A twisted wire acts like a spring for rotation, so the rod oscillates with this period (the next chapter derives formulas of this kind). I is the moment of inertia of two small balls of mass m on arms of length L/2; the light rod adds little.
---
$ \kappa = \frac{4\pi^2 I}{T^2} = \frac{2\pi^2 mL^2}{T^2}
Solve the period equation for the wire's stiffness κ.
---
$ \frac{GMm}{r^2}\,L = \frac{2\pi^2 mL^2}{T^2}\,\theta
Substitute κ, and F = GMm/r² for the attraction between a large ball of mass M and a small one at centre distance r.
---
$ G = \frac{2\pi^2 L r^2 \theta}{M T^2}
Solve for G. The small balls' mass m cancels; what is needed is the rod length, the ball separation, the large mass, the period and the angle of twist.
:::

The force Cavendish measured was about 1.5 × 10⁻⁷ N, the weight of roughly 15 micrograms. The rod's end moved by a few millimetres. To keep his own warmth from stirring the air, he sealed the apparatus in a room and watched it from outside through telescopes, reading scales lit by lamps.

::: history 1798 · Clapham, London
The apparatus was designed by John Michell, the rector who first imagined dark stars, but he died in 1793 before using it. It passed to Cavendish, a wealthy and famously shy natural philosopher, who rebuilt much of it and housed it in an outbuilding at his home. His paper, "Experiments to Determine the Density of the Earth", gave a mean density of 5.48 times that of water; an arithmetic slip hid that his own data gave 5.45. The modern value is 5.51. Cavendish never wrote down G: he framed everything as Earth's density. His result implies G ≈ 6.74 × 10⁻¹¹ in modern units, within about 1% of today's value, and G as a named constant came into use only in the late 1800s.
:::

## Weighing the Earth

Once G is known, Earth's mass follows from the surface gravity everyone can measure.

::: steps The mass and density of the Earth
$ g = \frac{GM_E}{R^2}
Surface gravity, from the shell theorem.
---
$ M_E = \frac{gR^2}{G} = \frac{9.81 \times (6.371 \times 10^6)^2}{6.674 \times 10^{-11}} \approx 5.97 \times 10^{24}\,\text{kg}
Solve for Earth's mass and put in numbers.
---
$ \rho = \frac{M_E}{\tfrac{4}{3}\pi R^3} \approx 5{,}510\,\text{kg/m}^3
Divide by Earth's volume for the mean density.
:::

Surface rocks have densities of about 2,700 to 3,000 kg/m³, half the average. So Earth's interior must be much denser than its crust: the first hard evidence for a heavy core, now known to be mostly iron.

::: note
Measuring G between lead balls in a room turned every GM from astronomy into a mass in kilograms: the Earth, the Moon, the Sun and the planets were weighed in one stroke.
:::

The same step applies throughout the solar system. The planets' orbits give the Sun's GM = 1.327 × 10²⁰ m³/s², so the Sun's mass is about 1.99 × 10³⁰ kg, some 330,000 Earths.

## Still the least precise constant

G remains hard to measure. The accepted value is 6.67430 × 10⁻¹¹ N m² kg⁻², uncertain by about 22 parts per million, far worse than other fundamental constants, and careful modern experiments disagree with each other by more than their stated errors. Gravity is so weak, and impossible to shield, that this two-century-old measurement is still a frontier.

::: quiz
? Why could astronomers know the Sun's GM long before they knew its mass?
- The Sun's mass was measured by weighing sunlight
+ Planetary orbits fix GM directly, but separating G from M needs a laboratory measurement of G
- G was assumed to be exactly 1
= T² = 4π²a³/(GM) gives the product GM from periods and distances; only Cavendish's experiment split it.
? In the torsion balance formula G = 2π²Lr²θ/(MT²), the mass of the small balls does not appear. Why?
- It is too small to matter
+ It appears in both the gravitational torque and the moment of inertia, and cancels
- The wire carries the small balls' weight
= Heavier small balls feel a larger force but also make the rod harder to turn, by the same factor.
? Cavendish found Earth's mean density was about twice that of surface rock. What did that imply?
- Earth is hollow
+ The interior is much denser than the crust
- His value of G was wrong by a factor of two
= An average of 5.5 g/cm³ with a 2.7–3 g/cm³ crust requires much denser material deeper down.
:::
