Newton's law of gravitation, F = GMm/r², contains a constant G that Newton never knew. That gap mattered more than it might seem. As the lesson on orbits showed, the motion of a satellite tells you the product GM for whatever it circles, but not G or M separately. So astronomers knew the Sun's GM and Earth's GM very well, and could say how many times more massive the Sun was than Earth, yet nobody could say what Earth actually weighed in kilograms. To find G you have to measure the gravitational pull between two objects whose masses you already know, which means objects small enough to handle and weigh. That pull turns out to be so feeble that measuring it was one of the great experimental achievements of the eighteenth century.

## Why the force is so hard to see

Start with how small the force is. Two 1 kg masses with their centres 10 cm apart attract each other with a force of about 7 × 10⁻⁹ N. That is the weight of less than a microgram, roughly a speck of dust. Every other object in the room pulls on the apparatus too, including the experimenter, and a far larger disturbance comes from air: the warmth of a human body sets up gentle convection currents, and those currents push on a delicate apparatus much harder than the gravitational pull being measured.

There is a second problem. Any laboratory object is also being pulled downward by the whole Earth, with a force billions of times larger than its pull on a neighbouring mass. So the detector has to respond to tiny sideways forces while completely ignoring the huge downward one. An ordinary balance, which compares downward forces, is no use at all.

## The torsion balance

The solution is a **torsion balance**, from the Latin for twisting. A light horizontal rod is hung at its centre from a thin vertical wire, so that it can turn freely in a horizontal plane. The wire carries the rod's weight, so Earth's downward pull does nothing except hold the wire taut. A sideways force on the ends of the rod, however, makes it turn, twisting the wire, and a long thin wire takes remarkably little force to twist. The rod settles at the angle where the wire's resistance to twisting balances the turning effect of the sideways force. The angle of twist therefore measures the force.

Henry Cavendish's version used a wooden rod 1.8 m long with a small lead ball of about 0.73 kg on each end. Next to the small balls he placed two large lead balls of about 158 kg each, one beside each end but on opposite sides of the rod, so that both attractions turned the rod in the same direction. Then he moved the large balls to the mirror-image positions, so that they turned it the other way, and measured how far the rod's resting position shifted. Measuring the shift between two positions, rather than a single angle, cancelled out errors in knowing exactly where the untwisted position was.

One quantity was still missing: how stiff the wire was, meaning how much torque, or turning effect, it takes to twist it by a given angle. Cavendish got this without applying any known force at all. Left alone, the rod does not settle at once but swings slowly back and forth about its resting position, like a pendulum turned on its side; his took about 7 minutes per swing. A stiffer wire would swing it back faster, so the period of the swing reveals the stiffness. The next chapter treats oscillations of this kind in detail; here we only need the formula for the period.

::: steps Extracting G from the torsion balance
$ \kappa\theta = FL
When the rod is at rest at its twisted angle θ, the torques on it balance. The wire resists twisting with a torque proportional to the angle, κθ, where κ (kappa) measures the wire's stiffness. On the other side, each small ball is pulled sideways by a force F at a distance L/2 from the wire, and there are two of them, so the gravitational torque is 2 × F × L/2 = FL.
---
$ T = 2\pi\sqrt{\frac{I}{\kappa}}, \qquad I = 2m\left(\frac{L}{2}\right)^2 = \frac{mL^2}{2}
A twisted wire acts like a spring for rotation: its restoring torque is proportional to the twist, just as a spring's force is proportional to its stretch. Such a system swings back and forth with this period T, where I is the **moment of inertia**, the rotational counterpart of mass that measures how hard the rod is to set turning. For two small balls of mass m at distance L/2 from the axis, I is 2m(L/2)², and the light rod adds little.
---
$ \kappa = \frac{4\pi^2 I}{T^2} = \frac{2\pi^2 mL^2}{T^2}
Square the period equation and solve it for κ. Then use I = mL²/2. The stiffness of the wire, which could not be measured directly, is now expressed through the swing period, the ball mass and the rod length, all of which can be measured.
---
$ \frac{GMm}{r^2}\,L = \frac{2\pi^2 mL^2}{T^2}\,\theta
Return to the balance of torques, κθ = FL, and fill in both sides. On the left, the force F is Newton's law of gravitation for a large ball of mass M and a small ball of mass m whose centres are a distance r apart. On the right, κ is replaced by the expression just found.
---
$ G = \frac{2\pi^2 L r^2 \theta}{M T^2}
Divide both sides by mL and multiply by r²/M to leave G on its own. The mass m of the small balls cancels, because it appears on both sides: a heavier small ball is pulled harder but is also harder to turn, by exactly the same factor. What remains needs only the rod length, the distance between ball centres, the large mass, the swing period and the angle of twist, every one of which can be measured with rulers, scales and clocks.
:::

This is a simplified version of the analysis. In the real experiment each large ball also pulls weakly on the more distant small ball at the other end of the rod, and pulls on the rod itself, and Cavendish corrected for both. But the logic is exactly as above.

The force Cavendish was measuring was about 1.5 × 10⁻⁷ N on each small ball, the weight of roughly 15 micrograms. The ends of the rod moved by only a few millimetres. To keep his own body heat from stirring the air, he enclosed the apparatus in a sealed room and never entered it during a measurement. He moved the large balls with ropes and pulleys from outside, and he read the position of the rod through telescopes set in the walls, looking at fine scales lit by lamps.

::: history 1783–1798 · Thornhill and Clapham
The apparatus was designed by John Michell, the Yorkshire rector who had argued in 1783 that some stars might be so compact that light could not escape them. Michell built the torsion balance to weigh the Earth but died in 1793 before he could use it. It passed to Henry Cavendish, a wealthy and famously shy natural philosopher who avoided conversation and communicated with his household servants by notes. Cavendish rebuilt much of the apparatus and housed it in an outbuilding at his home in Clapham, then on the edge of London. His paper of 1798 was titled "Experiments to Determine the Density of the Earth", and that is how he framed the result: Earth's mean density came out at 5.48 times that of water. A later reader found an arithmetic slip, and his own data actually gave 5.45; the modern value is 5.51. Cavendish never wrote down G at all, because the language of a universal gravitational constant came into common use only near the end of the nineteenth century. Translated into modern terms, his result corresponds to G ≈ 6.74 × 10⁻¹¹, within about 1% of today's value.
:::

## Weighing the Earth

Cavendish's description of his result as "the density of the Earth" makes sense once you see how directly it follows. Surface gravity, g = 9.81 m/s², had been measured precisely for a long time, and Earth's radius was known. The only thing standing between those numbers and Earth's mass was G.

::: steps The mass and density of the Earth
$ g = \frac{GM_E}{R^2}
This is surface gravity on a round planet of mass M_E and radius R. It comes from Newton's law combined with the shell theorem, which says that a spherical body pulls on things outside it as if all its mass were at the centre. Before Cavendish, this equation had two unknowns, G and M_E, tangled together.
---
$ M_E = \frac{gR^2}{G} = \frac{9.81 \times (6.371 \times 10^6)^2}{6.674 \times 10^{-11}} \approx 5.97 \times 10^{24}\,\text{kg}
Multiply both sides by R² and divide by G to get Earth's mass on its own. With G now measured, every quantity on the right is known, and the result is about six million million million million kilograms. This is the sense in which Cavendish weighed the Earth: he never put it on a scale, but he supplied the one missing number.
---
$ \rho = \frac{M_E}{\tfrac{4}{3}\pi R^3} \approx 5{,}510\,\text{kg/m}^3
Density is mass per volume. Earth's volume is that of a sphere, (4/3)πR³, so dividing the mass by it gives Earth's average density. The result is about 5.5 times the density of water.
:::

That number was a discovery in itself. The rocks at Earth's surface have densities of about 2,700 to 3,000 kg/m³, roughly half the average. If the whole planet were made of surface rock, its density would match theirs. Since it is much higher, the interior must contain material far denser than anything at the surface. This was the first firm evidence for a heavy core. Earthquake waves later showed that Earth does have a core, and that it is made mostly of iron, with some nickel.

The same step weighs everything else. Once G is known, every GM that astronomers had measured from orbits becomes a mass in kilograms. The planets' orbits give the Sun's GM = 1.327 × 10²⁰ m³/s², so the Sun's mass is about 1.99 × 10³⁰ kg, some 330,000 times Earth's. The Moon, Jupiter and every planet with a moon were weighed the same way, all at once.

::: note
Orbits measure only the product GM. Measuring G directly, from the pull between lead balls in a sealed room, turned every GM into a mass in kilograms, so the Earth, the Sun and the planets were weighed in one stroke.
:::

## Still the least precise constant

More than two centuries later, G remains surprisingly hard to measure. The accepted value is 6.67430 × 10⁻¹¹ N m² kg⁻², uncertain by about 22 parts per million. That sounds precise, but it is far worse than for other fundamental constants: the charge of the electron, for example, is now fixed exactly by definition, and many other constants are known to a few parts per billion. Worse, careful modern experiments using different methods disagree with one another by more than their stated uncertainties, which suggests that some hidden source of error has not yet been understood.

The reasons are the same ones Cavendish faced. Gravity is extraordinarily weak compared with other forces, and it cannot be shielded: there is no material that blocks gravity the way metal blocks electric fields, so every mass nearby, from the building to the water table under it, contributes a pull that has to be accounted for. Weighing the Earth with lead balls is still, in its modern form, a frontier experiment.

::: quiz
? Why could astronomers know the Sun's GM long before they knew its mass?
- The Sun's mass was measured by weighing sunlight
+ Planetary orbits fix GM directly, but separating G from M needs a laboratory measurement of G
- G was assumed to be exactly 1
= The formula T² = 4π²a³/(GM) gives the product GM from a planet's period and distance. Splitting that product into G and M needed an independent measurement of G, which Cavendish's experiment supplied.
? In the torsion balance formula G = 2π²Lr²θ/(MT²), the mass of the small balls does not appear. Why?
- It is too small to matter
+ It appears in both the gravitational torque and the moment of inertia, and cancels
- The wire carries the small balls' weight
= Heavier small balls feel a larger gravitational force, but they also make the rod harder to turn, by the same factor. The two effects cancel exactly.
? Cavendish found that Earth's mean density is about twice that of surface rock. What did that imply?
- Earth is hollow
+ The interior is much denser than the crust
- His value of G was wrong by a factor of two
= An average of about 5.5 g/cm³, with surface rock at only 2.7 to 3 g/cm³, means there must be much denser material deeper down. That was the first evidence for Earth's iron core.
:::
