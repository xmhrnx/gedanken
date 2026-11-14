The uncertainty principle is one of the few results of quantum mechanics that people outside physics have heard of, and it is usually explained wrongly. The familiar story is about clumsiness: to see an electron you must bounce light off it, and the light knocks it off course, so you can never pin down both where it is and how it is moving. That story is not exactly false, but it puts the cause in the wrong place. The limit holds before anyone measures anything. It is a property that every wave has, quantum or not, and it applies to particles because quantum mechanics describes them with waves. This lesson builds the principle from that fact, then shows why it is invisible for baseballs and decisive for atoms.

## Where is a wave?

Start with an ordinary question about waves. A pure sine wave has one exact wavelength, so it has one exact wavenumber k = 2π/λ, the number of radians of wave per metre. But a pure sine wave also runs on forever, every crest the same height as the last. If you ask where it is, there is no answer: it is everywhere equally.

Suppose instead you want a wave that is concentrated in one region and nearly zero elsewhere, a short burst called a **wave packet**. A single sine wave cannot do it. You have to add together sine waves of many different wavelengths, chosen so that in one small region their crests line up and reinforce each other, while everywhere else they fall out of step and cancel. The more sharply you want the packet confined, the more quickly the waves must fall out of step away from the centre, and that requires a wider range of wavelengths.

This trade-off was known long before quantum mechanics. It is a theorem of **Fourier analysis**, the mathematics of writing any shape as a sum of sine waves, worked out in the nineteenth century. Musicians and radio engineers meet it constantly. A very short click has no definite pitch, because a brief sound must be built from many frequencies. A pure tone, with one clear pitch, must last for many cycles. You can have a sound that is sharp in time or sharp in pitch, but not both.

## From waves to particles

Now bring in quantum mechanics. De Broglie's relation says a particle's momentum is tied to its wavenumber by p = ħk, where ħ is the reduced Planck constant. So the spread of wavenumbers in a wave packet is a spread of momenta. A particle whose wavefunction is confined to a small region must contain many wavenumbers, and therefore many momenta. Here is the argument made precise, with spreads measured by the **standard deviation**, the usual statistical measure of how widely values scatter around their average.

::: steps From waves to the uncertainty principle
$ \psi(x) = \int A(k)\,e^{ikx}\,\mathrm{d}k
Any wavefunction can be written as a sum of plane waves e^{ikx}, each with its own amplitude A(k); the integral is just a sum over a continuous range of k. This is Fourier's theorem, and it is allowed for any reasonable function. The amplitudes A(k) describe the same state as ψ(x), seen from a different angle: |A(k)|² says how much of each wavenumber is present.
---
$ \Delta x\,\Delta k \geq \tfrac{1}{2}
A theorem of Fourier analysis states that a function and its spread of wavenumbers cannot both be narrow. If Δx is the standard deviation of position computed from |ψ(x)|², and Δk the standard deviation of wavenumber computed from |A(k)|², their product is at least one half. Equality holds only for a bell-shaped packet, called a Gaussian; every other shape does worse.
---
$ p = \hbar k \quad\Rightarrow\quad \Delta p = \hbar\,\Delta k
De Broglie's relation converts each wavenumber into a momentum by multiplying by ħ. Multiplying every value in a distribution by the same constant multiplies its spread by that constant too, so the spread in momentum is ħ times the spread in wavenumber. This is the step where a fact about waves becomes a fact about particles.
---
$ \Delta x\,\Delta p \geq \frac{\hbar}{2}
Multiply both sides of the Fourier inequality by ħ, which is positive, so the inequality keeps its direction. Then replace ħΔk with Δp. The result is a lower limit on the product of the spreads in position and momentum, set by nothing but Planck's constant.
:::

This is the **Heisenberg uncertainty principle**. It is worth being careful about what Δx and Δp mean. They are not errors in a single measurement. They are the spreads of results you would get by preparing many particles in exactly the same state and measuring position on half of them and momentum on the other half. The principle says that no state exists in which both spreads are small. Squeeze the position distribution and the momentum distribution must widen; narrow the momentum and the position spreads out.

The deeper point is that the particle does not secretly have an exact position and an exact momentum that we are unable to learn. A wave simply does not have both. A packet confined to a small region is made of many wavelengths, and there is no fact about which one is "really" its wavelength. The word "uncertainty" suggests ignorance, and that is misleading; physicists sometimes prefer "indeterminacy", to say that the quantities themselves are not sharply defined.

::: history 1927 · Copenhagen
Werner Heisenberg found the relation early in 1927 while working at Niels Bohr's institute in Copenhagen, deriving it from the matrix mechanics he had invented two years before. To make it vivid he described a thought experiment: a microscope that locates an electron by scattering a gamma ray off it. Gamma rays have short wavelengths, so they locate the electron sharply, but each one carries a large momentum and kicks the electron unpredictably. Bohr, returning from a skiing holiday, read the draft and objected that Heisenberg had analysed the microscope wrongly, and more importantly that the limit did not come from the disturbance at all but from the wave nature of both light and electron. The two argued heatedly. Heisenberg published, but added a postscript conceding that Bohr's points were right. Later that year Earle Kennard proved the exact form Δx Δp ≥ ħ/2, with Δx and Δp defined as standard deviations, the version derived above.
:::

## Why baseballs do not notice

The constant on the right is tiny, ħ ≈ 1.055 × 10⁻³⁴ joule seconds, and that is why nobody noticed the principle before the twentieth century. Take a baseball of mass 0.145 kg and suppose its position is known to within a micrometre, far better than anyone could actually manage. The minimum spread in its momentum is ħ/(2Δx), and dividing by the mass gives the minimum spread in its velocity: about 4 × 10⁻²⁸ metres per second. At that speed the ball would take about nine billion years to drift the width of a single atom. For any object you can see, the uncertainty principle is irrelevant in practice.

## Why atoms do

For an electron confined inside an atom, the same arithmetic gives an answer that decides how matter is built. An atom is about 10⁻¹⁰ m across, so an electron bound to one has Δx of roughly that size. Then:

::: equation Momentum and energy of a confined electron
\Delta p \geq \frac{\hbar}{2\Delta x} \approx 5 \times 10^{-25}\,\text{kg m/s}, \qquad \frac{(\Delta p)^2}{2m} \approx 1.5 \times 10^{-19}\,\text{J} \approx 1\,\text{eV}
:::

The second expression is a rough estimate of the electron's kinetic energy, using the momentum spread as a stand-in for a typical momentum. It comes to about one electronvolt (eV), the energy an electron gains crossing one volt. That is the same scale as the energies of atoms: hydrogen's electron is bound by 13.6 eV, and the estimate, crude as it is, lands within a factor of about ten.

Now think about what this means for the stability of atoms. Classical physics had a serious problem here. An electron attracted to the nucleus should be able to lower its energy without limit by moving closer, and a circling electron should radiate away its energy and spiral inward in a fraction of a nanosecond. Matter, on that account, should collapse.

The uncertainty principle prevents the collapse. Squeezing the electron closer to the nucleus does lower its electrical potential energy, but it also shrinks Δx, which forces Δp up and raises the kinetic energy. The kinetic energy grows as 1/Δx², while the potential energy falls only as 1/Δx, so at small enough sizes the kinetic cost always wins. The total energy therefore has a minimum at some particular size, and that size, worked out properly, is the size of the hydrogen atom. Atoms are as big as they are because electrons resist being confined.

The same reasoning explains the zero-point energy of a particle in a box, from the previous lesson. A particle in a box of width L has a position spread smaller than L, so its momentum spread cannot be zero, and so neither can its kinetic energy. The box calculation and the uncertainty principle are two views of one fact.

::: model wavefunction
Each level of the particle in a box is a standing wave, which is a mix of two travelling waves with momenta +nπħ/L and −nπħ/L. Its position is confined to the box, but its momentum is not definite: a measurement could find it moving either way. Higher levels have larger momenta, as well as more nodes.
:::

## Energy and time

There is a related relation between energy and time, often written ΔE Δt ≳ ħ/2, but it means something different and has to be read with care. Time in quantum mechanics is not a measurable property of a particle the way position is; it is the clock against which things change. The relation says that a state which lasts only a short time Δt cannot have a sharply defined energy: its energy is spread over a range of about ħ/Δt. This has measurable effects. An excited atom typically survives for about ten nanoseconds before emitting light, so its upper energy level is slightly blurred, and the spectral line it emits has a small but definite width, called its natural linewidth.

::: note
The uncertainty principle is not about clumsy instruments. A wave confined in space must be built from many wavelengths, and since p = ħk, many wavelengths means many momenta. So Δx Δp ≥ ħ/2 holds for every quantum state, and it is what keeps atoms from collapsing.
:::

::: quiz
? A wave packet is squeezed to half its width. What happens to the minimum possible spread in its momentum?
- It halves
- It stays the same
+ It doubles
= The bound is Δp ≥ ħ/(2Δx), so halving Δx doubles the smallest momentum spread the state can have.
? Why is the uncertainty principle not simply a limit of our measuring instruments?
- Because modern instruments are already perfect
+ Because no wave has both a definite position and a single definite wavelength
- Because measurements never disturb quantum particles
= The limit is a property of the state itself: a localised packet must be built from many wavelengths, so there is no exact momentum for a better instrument to find.
? Roughly why does an electron not fall into the nucleus?
- The nucleus repels it at short distances
+ Confining it more tightly raises its kinetic energy faster than it lowers its potential energy
- It moves too fast to be captured
= Squeezing the electron shrinks Δx and forces up Δp, so its kinetic energy grows as 1/Δx² while the potential energy falls only as 1/Δx, and the total has a minimum at about the size of an atom.
:::
