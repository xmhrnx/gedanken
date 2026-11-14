Roll a ball at a hill too high for it and it rolls back, every time. A quantum particle facing an energy barrier higher than its own energy usually bounces back too, but not always. Sometimes it turns up on the far side. This **tunneling** explains why some nuclei fall apart, and it lets us see individual atoms.

## Inside a forbidden region

Take a particle of energy E moving toward a barrier of height V₀ > E and width L. Classically, inside the barrier the kinetic energy E − V₀ would be negative, which is impossible, so the particle cannot enter. Schrödinger's equation says something different.

::: steps The wavefunction inside the barrier
$ -\frac{\hbar^2}{2m}\frac{\mathrm{d}^2\psi}{\mathrm{d}x^2} + V_0\,\psi = E\psi
The time-independent Schrödinger equation with constant potential V₀ inside the barrier.
---
$ \frac{\mathrm{d}^2\psi}{\mathrm{d}x^2} = \frac{2m(V_0 - E)}{\hbar^2}\,\psi = \kappa^2\psi
Rearrange. Because V₀ > E, the coefficient is positive; call it κ², with κ = √(2m(V₀ − E))/ħ. Outside the barrier the sign is the other way and the solutions oscillate.
---
$ \psi(x) = C\,e^{-\kappa x} + D\,e^{\kappa x}
A function whose second derivative is a positive multiple of itself is an exponential, not a sine. For a wide barrier the growing solution must be almost absent, or ψ would become enormous, so the wave decays as e^{−κx}.
---
$ T \approx \frac{|\psi(L)|^2}{|\psi(0)|^2} \approx e^{-2\kappa L}
The probability of getting through is roughly the ratio of |ψ|² at the far side to the near side. Past the barrier the wave oscillates again, smaller but not zero: a travelling particle.
:::

The wavefunction does not stop at the wall; it leaks in and fades exponentially. If the barrier is thin enough that some is left at the far side, there is a chance, T, that the particle is found there. Nothing is borrowed or violated: the particle's energy is E before and after.

The exponent makes tunneling extremely sensitive to the barrier. Take an electron 1 eV below the top of a barrier. Then κ ≈ 5.1 × 10⁹ m⁻¹:

::: equation Transmission for an electron 1 eV below the barrier top
L = 0.5\,\text{nm}: \; T \approx e^{-5.1} \approx 6 \times 10^{-3} \qquad L = 1\,\text{nm}: \; T \approx e^{-10.2} \approx 4 \times 10^{-5}
:::

Doubling the width cut the probability by more than a hundred times. For a proton, nearly 2000 times heavier, κ is about 43 times larger, and tunneling through the same barrier is essentially impossible. For a ball and a hill, the exponent is so huge that the answer is zero for all practical purposes.

## Alpha decay

Some heavy nuclei, such as uranium-238, emit **alpha particles**: helium nuclei of two protons and two neutrons. Inside the nucleus, the strong force holds the alpha in; just outside, the electric repulsion from the remaining protons pushes it away. Together they form a barrier around the nucleus about 30 MeV high for uranium. The alpha particles that come out carry only about 4 MeV.

::: history 1928 · Göttingen and Princeton
George Gamow, and independently Ronald Gurney and Edward Condon, explained this with the new quantum mechanics. The alpha particle bounces around inside the nucleus, about 10²¹ times a second, and on each try has a tiny chance of tunneling out. The theory also explained an empirical rule found by Hans Geiger and John Nuttall in 1911: alpha emitters with slightly higher energy have enormously shorter half-lives. Since the energy sits in an exponent, a factor of two in alpha energy can change the half-life by more than twenty powers of ten. Uranium-238 emits 4.3 MeV alphas and has a half-life of 4.5 billion years; polonium-212 emits 8.8 MeV alphas and has a half-life of 0.3 microseconds.
:::

The same physics runs in reverse in the Sun. Its core is too cool for protons to climb over each other's electric barrier classically; they fuse only because they tunnel.

## Seeing atoms

::: history 1981 · Zürich
Gerd Binnig and Heinrich Rohrer at IBM's Zürich lab built the **scanning tunneling microscope**. A metal tip, sharpened to a single atom at its end, is held a few tenths of a nanometre above a surface. The gap is a barrier for electrons; with a small voltage applied, a current tunnels across it. Because T falls exponentially with width, the current changes by roughly a factor of ten for every 0.1 nm of distance. Scanning the tip and keeping the current constant traces the surface's height to a small fraction of an atom's diameter. They shared the 1986 Nobel Prize.
:::

Tunneling is also everyday engineering. Flash memory stores bits by pushing electrons through a thin insulating layer, and the shrinking of transistors is limited partly because electrons tunnel through insulators only a few atoms thick.

::: note
Inside a barrier higher than its energy, a particle's wavefunction does not vanish but decays exponentially. If the barrier is thin, some of the wave survives on the far side, and the chance of crossing is about e^{−2κL}.
:::

::: quiz
? A barrier's width is doubled. The tunneling probability e^{−2κL}
- halves
- drops to a quarter
+ is squared, so a small probability becomes far smaller
= e^{−2κ(2L)} = (e^{−2κL})². If T was 10⁻³, it becomes 10⁻⁶.
? Why do protons tunnel through a given barrier far less readily than electrons?
- They have positive charge
+ They are heavier, so κ = √(2m(V₀ − E))/ħ is larger
- They move more slowly
= κ grows as √m, and the mass ratio is about 1836, so κ is about 43 times larger.
? Why is the scanning tunneling microscope so sensitive to height?
- The tip touches each atom
+ The tunneling current depends exponentially on the gap width
- Electrons are smaller than atoms
= Because T ≈ e^{−2κL}, a change of 0.1 nm in the gap changes the current by about ten times.
:::
