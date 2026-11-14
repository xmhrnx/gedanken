Roll a ball toward a hill that is too high for it, and it slows, stops partway up, and rolls back. It does this every time, because it does not have enough energy to reach the top, and energy cannot appear from nowhere. A quantum particle facing an energy barrier higher than its own energy usually bounces back too, but not always. Every so often it turns up on the far side, having crossed a region that classical physics says it can never enter. This is called **tunneling**. It is not a curiosity at the edge of the theory: it explains why some atomic nuclei fall apart, it lets the Sun shine, and it is the working principle of a microscope that can see individual atoms. This lesson shows where tunneling comes from in Schrödinger's equation and why it depends so steeply on the size of the barrier.

## A region the particle cannot enter

Set up the problem as simply as possible. A particle of mass m and energy E moves along a line toward a barrier: a stretch of width L where the potential energy jumps up to a constant value V₀, higher than E. Outside the barrier the potential energy is zero.

Classically the answer is immediate. Kinetic energy is total energy minus potential energy, so inside the barrier the particle's kinetic energy would be E − V₀, which is negative. Kinetic energy, ½mv², can never be negative, so the particle cannot be inside the barrier at all. It reaches the edge and turns back.

Quantum mechanics asks a different question. It does not track the particle's path; it asks what the wavefunction ψ looks like in each region, and the wavefunction obeys Schrödinger's equation, which has no rule against the inside of a barrier. So we solve the equation there and see what it says.

::: steps The wavefunction inside the barrier
$ -\frac{\hbar^2}{2m}\frac{\mathrm{d}^2\psi}{\mathrm{d}x^2} + V_0\,\psi = E\psi
This is the time-independent Schrödinger equation, the one that applies to a state of definite energy E, written for the inside of the barrier, where the potential energy is the constant V₀. The first term is the kinetic energy expressed through the curvature of ψ, the second the potential energy, and their sum must equal the total energy times ψ.
---
$ \frac{\mathrm{d}^2\psi}{\mathrm{d}x^2} = \frac{2m(V_0 - E)}{\hbar^2}\,\psi = \kappa^2\psi
Move the V₀ term to the right-hand side and multiply through by −2m/ħ² to isolate the second derivative. Because V₀ is larger than E, the coefficient on the right is positive, so it can be written as the square of a real number κ (kappa), with κ = √(2m(V₀ − E))/ħ. Outside the barrier, where the potential is zero, the same rearrangement gives a negative coefficient instead, which is what makes the wave oscillate there.
---
$ \psi(x) = C\,e^{-\kappa x} + D\,e^{\kappa x}
A function whose second derivative is a positive multiple of itself is not a sine or cosine but an exponential: differentiating e^{κx} twice gives κ² e^{κx}, and the same holds for e^{−κx}. So inside the barrier ψ is a mixture of a decaying and a growing exponential. For a barrier wide compared with 1/κ, the matching conditions at the two edges force the growing part to be tiny, and the wave is dominated by the decaying exponential.
---
$ \psi(x) \approx C\,e^{-\kappa x}
Keeping only the decaying term, the wavefunction falls by a factor e for every distance 1/κ it penetrates. It does not stop at the wall; it leaks in and fades steadily. That is already the essential difference from the classical picture, where the particle is simply absent.
---
$ T \approx \frac{|\psi(L)|^2}{|\psi(0)|^2} \approx e^{-2\kappa L}
The probability of crossing, called the transmission probability T, is roughly the ratio of |ψ|² at the far edge to |ψ|² at the near edge, since |ψ|² is the probability density. Squaring the exponential doubles the exponent, giving e^{−2κL}. An exact calculation multiplies this by a factor that depends on E and V₀ and is usually of order one, so the exponential sets the size of the answer.
:::

Beyond the far edge, the potential drops back to zero and the wavefunction oscillates again, as a travelling wave. It is much smaller than the incoming wave, but it is not zero, and it describes a particle moving away on the other side. That is tunneling: the wave that reaches the far side carries a small probability of finding the particle there.

It is worth being clear about what does not happen. The particle does not borrow energy and pay it back, and no law is broken. Its energy is E before, during and after. What fails is the classical assumption that a particle must have a definite position and a positive kinetic energy at every point along a path. The particle has no path; it has a wavefunction, and the wavefunction extends into the barrier.

## How steep the exponential is

Because the barrier's width sits in an exponent, tunneling is extraordinarily sensitive to it. Take an electron whose energy is 1 electronvolt (eV, the energy an electron gains crossing one volt) below the top of the barrier. Then V₀ − E = 1 eV, and κ comes out at about 5.1 × 10⁹ per metre. Now compare two widths:

::: equation Transmission for an electron 1 eV below the barrier top
L = 0.5\,\text{nm}: \; T \approx e^{-5.1} \approx 6 \times 10^{-3} \qquad L = 1\,\text{nm}: \; T \approx e^{-10.2} \approx 4 \times 10^{-5}
:::

At half a nanometre, about two atoms thick, roughly six electrons in a thousand get through. Doubling the width to one nanometre cuts that by more than a hundred times. Doubling L doubles the exponent, which squares the already small probability.

The mass matters just as much. κ grows as the square root of the mass, and a proton is about 1836 times heavier than an electron, so its κ is about 43 times larger. Through the same 1 nm barrier, the exponent becomes roughly 440 instead of 10, and a probability of e⁻⁴⁴⁰ is zero for every practical purpose. For a ball and a hill the exponent is larger by a factor too big to write sensibly. That is why tunneling is a phenomenon of electrons, light nuclei and atomic distances, and why nobody has ever seen a ball pass through a hill.

## Alpha decay

The first great success of tunneling solved a puzzle in radioactivity. Some heavy nuclei, such as uranium-238, emit **alpha particles**, which are helium nuclei made of two protons and two neutrons. The puzzle was how they get out.

Picture the forces on an alpha particle. Inside the nucleus, the strong nuclear force, which acts only over very short distances, holds it in. Just outside, the strong force has faded and the electric repulsion between the alpha's two protons and the ninety remaining in the nucleus pushes it away. Combined, these make a barrier around the nucleus, like the rim of a crater. For uranium, the top of that rim is about 30 million electronvolts (MeV) high. Yet the alpha particles that come out carry only about 4 MeV. Classically, an alpha with 4 MeV could never climb a 30 MeV wall. Scattering experiments made the puzzle sharper: alpha particles fired at uranium with twice that energy were simply turned back by the electric repulsion, never getting in.

::: history 1928 · Göttingen and Princeton
George Gamow, a young Soviet physicist visiting Göttingen, and independently Ronald Gurney and Edward Condon at Princeton, saw that the new quantum mechanics answered the riddle. The alpha particle rattles around inside the nucleus, striking the barrier about 10²¹ times a second, and on each strike it has a tiny probability of tunneling through. The theory explained more than the escape. In 1911 Hans Geiger and John Nuttall had found an empirical rule: alpha emitters that release slightly more energetic alphas have enormously shorter half-lives. With tunneling the reason was clear. A more energetic alpha faces a thinner and lower stretch of barrier, and since the barrier enters through an exponent, a modest change in energy changes the escape rate by huge factors. Uranium-238 emits 4.3 MeV alphas and has a half-life of about 4.5 billion years. Polonium-212 emits 8.8 MeV alphas and has a half-life of about 0.3 microseconds. A factor of two in energy corresponds to a factor of more than 10²³ in lifetime.
:::

Tunneling also runs the other way, letting particles get into a nucleus rather than out. In the core of the Sun, at about 15 million kelvin, protons move far too slowly to climb over each other's electric repulsion. Classically they would almost never get close enough for the strong force to fuse them. They fuse because, very occasionally, they tunnel through the repulsion, and with the vast number of protons in the core, those rare events are enough to power the Sun.

## Seeing atoms

The steepness that makes tunneling rare is also what makes it useful as a measuring tool, because a tiny change in distance produces a large change in current.

::: history 1981 · Zürich
Gerd Binnig and Heinrich Rohrer at IBM's research laboratory near Zürich built the first **scanning tunneling microscope**. Its tip is a metal needle so sharp that its end is a single atom, held a few tenths of a nanometre above a conducting surface. The gap between tip and surface is a barrier for electrons: they do not have the energy to leave the metal and cross empty space. Apply a small voltage, though, and a current flows by tunneling. Because the probability falls exponentially with width, the current changes by roughly a factor of ten for every 0.1 nm change in the gap, a fraction of an atom's diameter. Scanning the tip across the surface while a feedback circuit moves it up and down to keep the current constant traces out the surface's shape, atom by atom. Binnig and Rohrer shared half of the 1986 Nobel Prize in Physics for it.
:::

Tunneling is also everyday engineering. Flash memory, in phones and USB drives, stores each bit as charge trapped on a tiny isolated electrode, and it writes and erases by forcing electrons to tunnel through a thin insulating layer. Tunneling is also one of the limits on how small transistors can be made: when an insulating layer is only a few atoms thick, electrons leak through it, wasting power and blurring the difference between on and off.

::: note
Inside a barrier higher than its energy, a particle's wavefunction does not vanish but decays exponentially. If the barrier is thin enough, some of the wave survives on the far side, and the chance of crossing is about e^{−2κL}, which depends extremely steeply on the barrier's width and on the particle's mass.
:::

::: quiz
? A barrier's width is doubled. What happens to the tunneling probability e^{−2κL}?
- It halves
- It drops to a quarter
+ It is squared, so a small probability becomes far smaller
= Doubling L doubles the exponent, and e^{−2κ(2L)} = (e^{−2κL})², so a probability of 10⁻³ becomes 10⁻⁶.
? Why do protons tunnel through a given barrier far less readily than electrons?
- Because they have positive charge
+ Because they are heavier, so κ = √(2m(V₀ − E))/ħ is larger
- Because they move more slowly
= κ grows as the square root of the mass, and a proton is about 1836 times heavier than an electron, so its κ is about 43 times larger and the exponent is 43 times bigger.
? Why is the scanning tunneling microscope so sensitive to height?
- The tip touches each atom in turn
+ The tunneling current depends exponentially on the width of the gap
- Electrons are smaller than atoms
= Because T ≈ e^{−2κL}, changing the gap by just 0.1 nm changes the current by about a factor of ten, so tiny bumps produce large signals.
:::
