Newton's second law tells you how a particle's position changes in time. Quantum mechanics needs the equivalent for the wavefunction: given ψ now, what is ψ a moment later? The answer, Schrödinger's equation, cannot be derived from classical physics. It can be guessed, though, from two facts you already have, and the guess is short.

## Two quantum rules and a wave

Planck and Einstein tied a quantum's energy to its frequency; de Broglie tied momentum to wavelength. Written with the angular frequency ω = 2πf, the wavenumber k = 2π/λ and the reduced Planck constant ħ = h/2π ≈ 1.055 × 10⁻³⁴ J s, they read:

::: equation Energy and momentum of a quantum wave
E = \hbar\omega, \qquad p = \hbar k
:::

The simplest wave for a free particle with definite energy and momentum is a plane wave, a single sine wave running along x forever. In complex form it is ψ = A e^{i(kx − ωt)}. Now look for an equation that this wave satisfies for every choice of k and ω, as long as they are related as energy and momentum are for a particle: E = p²/2m + V.

::: steps Motivating Schrödinger's equation
$ \psi = A\,e^{i(kx - \omega t)}
A plane wave of amplitude A. The complex form is needed: a real wave such as cos(kx − ωt) does not work in the steps below.
---
$ \frac{\partial \psi}{\partial t} = -i\omega\,\psi \quad\Rightarrow\quad i\hbar\frac{\partial \psi}{\partial t} = \hbar\omega\,\psi = E\psi
Differentiating with respect to time brings down −iω. Multiplying by iħ (and using i² = −1) turns this into the energy times ψ.
---
$ \frac{\partial^2 \psi}{\partial x^2} = -k^2\psi \quad\Rightarrow\quad -\frac{\hbar^2}{2m}\frac{\partial^2 \psi}{\partial x^2} = \frac{\hbar^2 k^2}{2m}\psi = \frac{p^2}{2m}\psi
Differentiating twice with respect to x brings down (ik)² = −k². Multiplying by −ħ²/2m gives the kinetic energy p²/2m times ψ.
---
$ E\psi = \frac{p^2}{2m}\psi + V\psi
Multiply the energy relation E = p²/2m + V through by ψ. Here V is the potential energy.
---
$ i\hbar\frac{\partial \psi}{\partial t} = -\frac{\hbar^2}{2m}\frac{\partial^2 \psi}{\partial x^2} + V(x)\,\psi
Replace Eψ and (p²/2m)ψ with the derivatives found above. The leap is to assume this holds even when V depends on x and ψ is no longer a plane wave.
:::

This is the **Schrödinger equation** in one dimension. It is a guess, not a theorem; it earns its place because its predictions agree with experiment. It is linear: if ψ₁ and ψ₂ are solutions, so is any sum of them, which is what makes superposition and interference possible. It is also first order in time, so ψ now fixes ψ at every later moment.

::: history 1925–26 · Zürich and Arosa
At a seminar in Zürich in late 1925 Schrödinger presented de Broglie's matter waves. Peter Debye remarked that a wave ought to have a wave equation. Schrödinger spent the Christmas holiday in the Alpine resort of Arosa and came back with one. In the first of four papers in 1926 he solved it for hydrogen and recovered the energy levels Bohr had found by fiat. He thought ψ was a real wave of charge; Born's probability reading came months later.
:::

## Stationary states

For a state of definite energy E, the time dependence is just a rotating phase, ψ(x, t) = φ(x) e^{−iEt/ħ}. Substituting this removes time and leaves the **time-independent Schrödinger equation**, −(ħ²/2m) φ″ + Vφ = Eφ. Since |e^{−iEt/ħ}|² = 1, the probability density of such a state does not change: these are the stable states of atoms.

## The particle in a box

The simplest test is a particle confined between two impenetrable walls at x = 0 and x = L. Inside, V = 0. The walls are infinitely high, so ψ must vanish outside and, to be continuous, at the walls too.

::: steps Energy levels of a particle in a box
$ -\frac{\hbar^2}{2m}\frac{\mathrm{d}^2\phi}{\mathrm{d}x^2} = E\phi
Inside the box V = 0, so only the kinetic term remains.
---
$ \phi(x) = A\sin kx + B\cos kx, \qquad k = \frac{\sqrt{2mE}}{\hbar}
Sines and cosines are the functions whose second derivative is a negative multiple of themselves. Substituting either gives E = ħ²k²/2m.
---
$ \phi(0) = 0 \quad\Rightarrow\quad B = 0
At x = 0 the sine is zero but the cosine is 1, so the cosine part must be absent.
---
$ \phi(L) = A\sin kL = 0 \quad\Rightarrow\quad kL = n\pi, \quad n = 1, 2, 3, \ldots
A sine is zero only at whole multiples of π. A = 0 or n = 0 would give ψ = 0 everywhere, meaning no particle at all, so n starts at 1.
---
$ E_n = \frac{\hbar^2 k^2}{2m} = \frac{n^2\pi^2\hbar^2}{2mL^2} = \frac{n^2 h^2}{8mL^2}
Substitute k = nπ/L into E = ħ²k²/2m. Requiring A = √(2/L) makes the total probability 1.
:::

Only whole-number multiples of half a wavelength fit in the box, as on a guitar string, and so only certain energies are allowed. Quantisation is not added by hand; it follows from the boundary conditions. The levels grow as n²: 1, 4, 9, 16 times the lowest. And the lowest is not zero. A confined particle can never be completely at rest; this minimum is called the **zero-point energy**.

The numbers matter only at small scales. For an electron in a box 1 nm wide, about the size of a small molecule, E₁ = h²/8mL² ≈ 6.0 × 10⁻²⁰ J ≈ 0.38 eV, comparable to chemical energies. For a grain of dust in a box a millimetre wide the spacing is unmeasurably small, which is why the classical world looks continuous.

::: model wavefunction
Step through n. Each level adds one half-wavelength and one node, and its energy grows as n².
:::

::: note
Schrödinger's equation turns E = p²/2m + V into an equation for waves, using E = ħω and p = ħk. Confine the wave and only standing waves survive, so energy comes in discrete levels.
:::

::: quiz
? A particle in a box is in its lowest level. Which change raises that energy by a factor of 4?
- Doubling the width of the box
+ Halving the width of the box
- Doubling the particle's mass
= E₁ = h²/8mL² goes as 1/L², so halving L multiplies it by 4.
? Why can't a particle in a box have zero energy?
- The walls give it energy when it hits them
+ The n = 0 solution is ψ = 0 everywhere, which means no particle at all
- Energy is always measured relative to the walls
= kL = nπ with n = 0 makes the sine vanish everywhere, so the lowest real state is n = 1.
? What is the ratio E₃ / E₂ for a particle in a box?
- 3/2
+ 9/4
- 1
= E_n ∝ n², so E₃/E₂ = 9/4.
:::
