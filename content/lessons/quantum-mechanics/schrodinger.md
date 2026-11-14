Newton's second law is a machine for predicting the future: tell it where a particle is and how it is moving, and it tells you where the particle will be a moment later. Quantum mechanics describes a particle by a wavefunction ψ instead of a position, so it needs a machine of the same kind for ψ. Given the wavefunction now, what is it a moment from now? The answer is Schrödinger's equation. It cannot be derived from classical physics, because it contains physics that classical mechanics does not. But it can be guessed from two facts about quantum waves that were already known, and following the guess shows where every piece of the equation comes from. Once we have it, we will solve it for the simplest possible case and watch energy levels appear without being put in by hand.

## What we have to work with

By 1925 two rules connected the particle properties of light and matter with their wave properties. Max Planck and Albert Einstein had found that the energy of a quantum of light is proportional to its frequency. Louis de Broglie had proposed that a particle's momentum is inversely proportional to its wavelength. Both rules involve Planck's constant h, and they look tidiest when written with three convenient quantities: the **angular frequency** ω = 2πf, which counts the cycles per second in radians rather than whole turns; the **wavenumber** k = 2π/λ, which counts radians of wave per metre; and the **reduced Planck constant** ħ = h/2π, about 1.055 × 10⁻³⁴ joule seconds. In this language the two rules read:

::: equation Energy and momentum of a quantum wave
E = \hbar\omega, \qquad p = \hbar k
:::

So a wave's rate of oscillation in time carries its energy, and its rate of oscillation in space carries its momentum. That is the bridge we need. If we can find an operation on ψ that measures how fast it oscillates in time, and another that measures how fast it oscillates in space, then the familiar statement about energy, that total energy is kinetic plus potential, E = p²/2m + V, can be turned into a statement about waves.

## Guessing the equation

Start with the simplest case: a free particle with a definite energy and a definite momentum. The wave that matches it is a **plane wave**, a single wave of one wavelength running along x forever, with the same frequency everywhere. We look for an equation that this wave satisfies for every allowed pair of k and ω, then make the bold assumption that the same equation holds for every situation.

::: steps Motivating Schrödinger's equation
$ \psi = A\,e^{i(kx - \omega t)}
This is a plane wave of amplitude A, written in complex form. By Euler's formula, e^{iθ} = cos θ + i sin θ, so it contains a cosine and a sine wave together. The complex form matters: as the next two steps show, differentiating it returns the same function times a constant, which a real cosine wave does not do in time.
---
$ \frac{\partial \psi}{\partial t} = -i\omega\,\psi
Differentiate with respect to time, holding x fixed. The exponential's derivative is itself times whatever multiplies t inside it, which is −iω. So taking a time derivative of this wave amounts to multiplying it by −iω: the derivative has read off the frequency.
---
$ i\hbar\frac{\partial \psi}{\partial t} = \hbar\omega\,\psi = E\psi
Multiply both sides by iħ. Since i × (−i) = 1, the right side becomes ħωψ, and by Planck and Einstein ħω is the energy. We now have an operation, iħ ∂/∂t, which acting on the wave gives its energy times the wave.
---
$ \frac{\partial^2 \psi}{\partial x^2} = (ik)^2\psi = -k^2\psi
Now differentiate twice with respect to position, holding t fixed. Each derivative brings down a factor ik, and (ik)² = −k². Two derivatives are used rather than one because kinetic energy depends on the square of momentum, so we want k².
---
$ -\frac{\hbar^2}{2m}\frac{\partial^2 \psi}{\partial x^2} = \frac{\hbar^2 k^2}{2m}\psi = \frac{p^2}{2m}\psi
Multiply by −ħ²/2m, where m is the particle's mass. The minus sign cancels the one from (ik)², and since p = ħk the right side is p²/2m times ψ, the kinetic energy times the wave. So a second space derivative, suitably scaled, reads off kinetic energy.
---
$ E\psi = \frac{p^2}{2m}\psi + V\psi
Take the ordinary energy relation E = p²/2m + V, where V is the potential energy, and multiply every term by ψ. Nothing quantum has happened yet; this is classical energy bookkeeping written next to the wave.
---
$ i\hbar\frac{\partial \psi}{\partial t} = -\frac{\hbar^2}{2m}\frac{\partial^2 \psi}{\partial x^2} + V(x)\,\psi
Replace Eψ and (p²/2m)ψ with the derivative expressions found above. For a plane wave in constant V this is just a rewriting. The leap, which is a guess rather than a deduction, is to assume the equation stays true when V varies from place to place and ψ is no longer a plane wave.
:::

This is the **Schrödinger equation** in one dimension. It is worth being clear about its status. The steps above are a motivation, not a proof: a plane wave fixes the equation only for the case it describes, and extending it to every potential is a hypothesis. The equation earns its place the way Newton's laws did, by agreeing with experiment, and it has agreed with every measurement on atoms, molecules and solids for a century.

Two features of the equation matter for everything that follows. First, it is **linear**: ψ appears only to the first power, never squared. That means if ψ₁ and ψ₂ are both solutions, so is any sum of them, such as ψ₁ + ψ₂. This is what allows superposition, the adding of waves that produced interference in the double slit. Second, it contains only a first derivative in time. So knowing ψ at one instant is enough to fix it at every later instant, exactly as knowing position and velocity fixes a Newtonian path. Between measurements, the evolution of the wavefunction is entirely deterministic. The randomness of quantum mechanics enters only through the Born rule, when we ask what a measurement will find.

::: history 1925–26 · Zürich and Arosa
Late in 1925 Erwin Schrödinger, then a professor in Zürich, gave a seminar on de Broglie's idea that electrons are waves. Afterwards Peter Debye, who chaired the colloquium, remarked that it was a rather childish way to talk: if there are waves, there ought to be a wave equation. Schrödinger took the point. He spent the Christmas holiday in the Alpine resort of Arosa and came back with one. Early in 1926 he published the first of four papers, in which he solved his equation for the hydrogen atom and recovered exactly the energy levels Niels Bohr had obtained in 1913 by simply postulating which orbits were allowed. In Schrödinger's version the levels came out of the mathematics, as the only standing waves that fit around the nucleus. He believed ψ was a real wave of electric charge; Max Born's probability reading came a few months later and was never to his taste.
:::

## States that do not change

Atoms sit in stable states for long periods, so it is natural to ask which wavefunctions do not change in any measurable way. The answer comes from separating time out of the equation.

::: steps Separating out time
$ \psi(x, t) = \phi(x)\,e^{-iEt/\hbar}
Guess that the wavefunction is a fixed shape φ(x) multiplied by a factor that only rotates its phase in time, at the rate set by an energy E. This is the time behaviour we found for the plane wave, where ω = E/ħ. Whether the guess works depends on whether it satisfies the equation.
---
$ i\hbar\frac{\partial \psi}{\partial t} = E\,\phi(x)\,e^{-iEt/\hbar}
Differentiate with respect to time. Only the exponential depends on t, and its derivative brings down −iE/ħ; multiplied by iħ that gives E. So the left side of Schrödinger's equation is simply E times the wavefunction.
---
$ -\frac{\hbar^2}{2m}\frac{\mathrm{d}^2\phi}{\mathrm{d}x^2} + V(x)\,\phi = E\,\phi
On the right side, the space derivatives act only on φ, so every term carries the same phase factor e^{−iEt/ħ}. Divide it out of both sides and time disappears. What remains is the **time-independent Schrödinger equation**, an equation for the shape φ alone.
:::

The guess works whenever φ solves this second equation, and such solutions are called **stationary states**. The name is apt. The probability density is |ψ|² = |φ|² |e^{−iEt/ħ}|², and a pure phase factor has size one, so the density never changes. These are the states of definite energy, and they are what atoms and molecules settle into. Finding the allowed energies of a system means finding the values of E for which the time-independent equation has sensible solutions.

## The particle in a box

The simplest test is a particle confined between two walls, at x = 0 and x = L, that it cannot pass through. Inside, there is no force, so V = 0. The walls are taken to be infinitely high, so the particle is never found outside and ψ is zero there. A wavefunction cannot jump suddenly, because its derivatives would become infinite and the equation would make no sense, so ψ must also be zero at the walls themselves. These requirements are called **boundary conditions**, and they turn out to do all the work.

::: steps Energy levels of a particle in a box
$ -\frac{\hbar^2}{2m}\frac{\mathrm{d}^2\phi}{\mathrm{d}x^2} = E\phi
Inside the box V = 0, so the time-independent equation keeps only the kinetic term. It says that the second derivative of φ must be a negative multiple of φ itself, since E is positive.
---
$ \phi(x) = A\sin kx + B\cos kx, \qquad k = \frac{\sqrt{2mE}}{\hbar}
Sines and cosines are exactly the functions whose second derivative is a negative multiple of themselves: twice differentiating sin kx gives −k² sin kx. Putting either into the equation shows it works provided ħ²k²/2m = E, which fixes k in terms of the energy. The most general solution is any mixture of the two, with constants A and B still to be found.
---
$ \phi(0) = 0 \quad\Rightarrow\quad B = 0
Apply the left wall. At x = 0 the sine is zero but the cosine equals one, so φ(0) = B. For φ to vanish there, B must be zero, and only the sine survives.
---
$ \phi(L) = A\sin kL = 0 \quad\Rightarrow\quad kL = n\pi, \quad n = 1, 2, 3, \ldots
Apply the right wall. We cannot set A = 0, because then ψ would be zero everywhere and there would be no particle. So sin kL must vanish, which happens only when kL is a whole multiple of π. The choice n = 0 also gives a wave that is zero everywhere, so the counting starts at n = 1. This is the step where quantisation happens: only certain wavenumbers fit.
---
$ E_n = \frac{\hbar^2 k^2}{2m} = \frac{n^2\pi^2\hbar^2}{2mL^2} = \frac{n^2 h^2}{8mL^2}
Put k = nπ/L back into the relation between k and E. Because k can take only certain values, so can the energy. The last form uses ħ = h/2π, which turns π²ħ² into h²/4.
---
$ \int_0^L A^2\sin^2\frac{n\pi x}{L}\,\mathrm{d}x = \frac{A^2 L}{2} = 1 \quad\Rightarrow\quad A = \sqrt{\frac{2}{L}}
Finally, the particle must be somewhere in the box, so the total probability must be one. The average value of sin² over whole half-cycles is ½, so the integral is A²L/2. Setting it equal to one fixes the size of the wave. The amplitude does not affect the energies; it only makes the probabilities add up.
:::

Pause on what has happened. Nobody told the particle which energies it may have. The rule came out of a simple demand: the wave must fit in the box, with a node at each wall, so only whole numbers of half-wavelengths are allowed, just as on a guitar string pinned at both ends. Discrete energy levels, which Bohr had to postulate, are here a consequence of confining a wave.

The levels also have a definite pattern. They grow as n², so the second level is four times the first, the third nine times, the fourth sixteen times. And the lowest level, n = 1, is not zero. A confined particle can never be completely at rest. This minimum is called the **zero-point energy**, and it has real consequences: it is why liquid helium does not freeze at ordinary pressure, even at absolute zero.

Whether any of this matters depends on the scale. For an electron in a box 1 nanometre wide, about the size of a small molecule, the lowest level is E₁ = h²/8mL² ≈ 6.0 × 10⁻²⁰ J, or about 0.38 electronvolts (eV, the energy an electron gains crossing one volt). That is comparable to the energies of chemical bonds, so the quantisation shapes chemistry. For a speck of dust in a box a millimetre wide, the mass is some twenty powers of ten larger and L² is twelve powers larger, so the spacing between levels is unmeasurably small. The levels are still there, but so close together that energy appears continuous, which is why the everyday world shows no sign of them.

::: model wavefunction
Step through n. Each level fits one more half-wavelength into the box and gains one more node, and its energy grows as n²: watch how quickly the levels spread apart.
:::

::: note
Schrödinger's equation turns the energy relation E = p²/2m + V into an equation for waves, using E = ħω and p = ħk. When a wave is confined, only the standing waves that fit survive, so energy comes in discrete levels, and the lowest of them is never zero.
:::

::: quiz
? A particle in a box is in its lowest level. Which change multiplies that energy by four?
- Doubling the width of the box
+ Halving the width of the box
- Doubling the particle's mass
= The lowest energy is E₁ = h²/8mL², which goes as 1/L², so halving L multiplies it by four; doubling the mass would halve it instead.
? Why can a particle in a box not have zero energy?
- The walls give it energy every time it hits them
+ The n = 0 solution is zero everywhere, which describes no particle at all
- Energy is always measured relative to the height of the walls
= Setting kL = nπ with n = 0 makes the sine vanish across the whole box, so the lowest state that actually contains a particle is n = 1, with energy above zero.
? What is the ratio E₃/E₂ for a particle in a box?
- 3/2
+ 9/4
- 1
= The energies grow as n², so E₃/E₂ = 3²/2² = 9/4.
:::
