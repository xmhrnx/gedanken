De Broglie gave every particle a wavelength, λ = h/p, and electron diffraction confirmed it. That left an awkward question: a wave of *what*? An electron always arrives at a detector as a single, whole particle at one spot. Whatever is waving, it is not the electron's charge smeared out in space.

## The wavefunction

Quantum mechanics describes a particle by a **wavefunction** ψ(x, t): a number attached to every point of space at every moment. Unlike a water wave, ψ is complex. At each point it has a size |ψ| and a phase, an angle that says where in its cycle the wave is. Neither can be seen directly. What connects ψ to experiment is one rule.

::: equation The Born rule
P(x \text{ to } x + \mathrm{d}x) = |\psi(x,t)|^2\,\mathrm{d}x
:::

The probability of finding the particle in a small interval dx is |ψ|² times the width of that interval. So |ψ|² is a **probability density**: probability per unit length (per unit volume in three dimensions). Where |ψ|² is large the particle is likely to turn up; where ψ = 0 it never does.

The particle must be found somewhere, so the probabilities over all space add to one:

::: equation Normalisation
\int_{-\infty}^{\infty} |\psi(x,t)|^2\,\mathrm{d}x = 1
:::

The rule gives probabilities, not outcomes. Two electrons prepared with identical wavefunctions can land in different places. Repeat the experiment many times and the pattern of hits traces out |ψ|².

::: history 1926 · Göttingen
Erwin Schrödinger, who had just written down the equation for ψ, believed it described a real spread-out distribution of the electron's charge. Max Born, analysing how electrons scatter off atoms, saw that this could not work: the scattered wave spreads in all directions, but each electron is detected in one place. In his paper he proposed that ψ determines the probability of an outcome. In the proofs he added a footnote correcting himself: the probability is proportional to the *square* of ψ. Schrödinger never fully accepted the interpretation. Born received the Nobel Prize for it in 1954.
:::

## Why square it: interference

Squaring is what lets probabilities interfere. Send electrons at a screen with two slits. Let ψ₁ be the wave that comes through slit 1 and ψ₂ the wave through slit 2. Waves add, so the total is ψ₁ + ψ₂. Write each in terms of its size and phase, ψ = |ψ|e^{iφ}.

::: steps Probability with two open slits
$ \psi = \psi_1 + \psi_2
With both slits open, the waves from each slit superpose. It is the amplitudes that add, not the probabilities.
---
$ |\psi|^2 = (\psi_1 + \psi_2)^*(\psi_1 + \psi_2)
For a complex number z, |z|² = z*z, where z* is the complex conjugate (the same size, opposite phase).
---
$ |\psi|^2 = |\psi_1|^2 + |\psi_2|^2 + \psi_1^*\psi_2 + \psi_1\psi_2^*
Multiply out the brackets.
---
$ |\psi|^2 = |\psi_1|^2 + |\psi_2|^2 + 2|\psi_1||\psi_2|\cos(\phi_2 - \phi_1)
The last two terms are complex conjugates of each other, so their sum is twice the real part of ψ₁*ψ₂, which is |ψ₁||ψ₂|cos(φ₂ − φ₁).
:::

The first two terms are what you would get for bullets: the chance through slit 1 plus the chance through slit 2. The third, the **interference term**, can be positive or negative depending on the phase difference, which changes across the screen. Where the waves arrive in step, the probability is up to four times that from one slit; where they arrive out of step, it falls to zero. Opening a second slit can make some places on the screen *less* likely to be hit.

::: model interference
Change the slit separation and watch the fringes move. For electrons, the bright bands are where |ψ|² is large: places many electrons land, one at a time.
:::

The pattern builds up even when electrons pass through the apparatus one at a time, hours apart. Each makes a single dot; after thousands of dots, the fringes appear. Each electron's ψ went through both slits, and its own |ψ|² sets where it can land.

::: history 1961–1989 · Tübingen, Bologna, Tokyo
Claus Jönsson first showed electron interference with microscopic slits, up to five of them, in 1961. In 1974 Pier Giorgio Merli, Gian Franco Missiroli and Giulio Pozzi in Bologna recorded the pattern forming from single electrons. In 1989 Akira Tonomura's team at Hitachi filmed it: scattered dots that slowly assemble into fringes.
:::

If you add a detector that records which slit each electron used, the fringes vanish and the hits follow |ψ₁|² + |ψ₂|². Measuring the path destroys the interference term; the lesson on measurement returns to this.

::: model wavefunction
A particle trapped in a box. The wave ψ can be positive or negative; the probability |ψ|² never is. At the nodes, where ψ crosses zero, the particle is never found.
:::

::: note
The wavefunction is not a physical wave you can see. It is an amplitude: its square gives the probability of finding the particle at each place. Because amplitudes add before they are squared, quantum probabilities can interfere.
:::

::: quiz
? In a region where ψ is twice as large as at another place, the particle is
- twice as likely to be found, per unit length
+ four times as likely to be found, per unit length
- equally likely to be found, since only the phase matters
= Probability density is |ψ|², so doubling |ψ| multiplies it by 4.
? With both slits open, some points on the screen get fewer electrons than with one slit open. Why?
- Electrons through the two slits collide and knock each other away
+ Amplitudes from the two slits add and can cancel before being squared
- The second slit blocks part of the beam
= The interference term 2|ψ₁||ψ₂|cos(φ₂ − φ₁) is negative where the waves arrive out of step.
? Electrons are sent one at a time through a double slit. What builds up on the screen?
- Two bright bands behind the slits
+ An interference pattern, made of individual dots
- A smooth blur with no structure
= Each electron lands at one point, with probability set by its own |ψ|², which includes interference.
:::
