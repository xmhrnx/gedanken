The uncertainty principle is often explained as a clumsiness of measurement: to see an electron you must hit it with light, and the light knocks it. That picture is misleading. The limit holds before anyone measures anything. It is a property that every wave has, and quantum particles are described by waves.

## Where is a wave?

A pure sine wave has one exact wavelength, so one exact wavenumber k = 2π/λ. But it runs on forever with the same amplitude: it is not anywhere in particular. To make a wave that is concentrated in one region, a **wave packet**, you have to add together sine waves of many different wavelengths. They reinforce each other in a small region and cancel everywhere else.

The narrower you want the packet, the wider the range of wavenumbers you need. This is a mathematical fact about waves, known from the theory of Fourier analysis (writing any shape as a sum of sine waves) long before quantum mechanics. Musicians know a version of it: a very short click has no definite pitch, while a pure tone must last many cycles.

Measure the spread of positions by the standard deviation Δx and the spread of wavenumbers by Δk. Then:

::: steps From waves to the uncertainty principle
$ \psi(x) = \int A(k)\,e^{ikx}\,\mathrm{d}k
Any wavefunction can be built as a sum of plane waves e^{ikx}, each with its own amplitude A(k). |A(k)|² tells you how much of each wavenumber is present.
---
$ \Delta x\,\Delta k \geq \tfrac{1}{2}
A theorem of Fourier analysis: a function and its spectrum cannot both be narrow. The product of their standard deviations is at least ½, with equality only for a bell-shaped (Gaussian) packet.
---
$ p = \hbar k \quad\Rightarrow\quad \Delta p = \hbar\,\Delta k
De Broglie's relation turns each wavenumber into a momentum, so the spread in momentum is ħ times the spread in wavenumber.
---
$ \Delta x\,\Delta p \geq \frac{\hbar}{2}
Multiply the Fourier inequality by ħ.
:::

This is the **Heisenberg uncertainty principle**. Δx and Δp are spreads in the results you would get from many measurements on identically prepared particles: measure position on half of them, momentum on the other half. No state exists in which both spreads are small. A particle with a sharply defined position has no definite momentum, and the reverse. It is not that the electron has an exact position and momentum we fail to learn; a wave does not have both.

::: history 1927 · Copenhagen
Werner Heisenberg derived the relation from his matrix mechanics and illustrated it with a thought experiment: a microscope that locates an electron with a gamma ray, kicking it in the process. Niels Bohr objected that Heisenberg's account of the microscope was wrong in detail and that the limit came from the wave nature itself, not from disturbance. Heisenberg added a postscript conceding the point. Later that year Earle Kennard proved the exact form Δx Δp ≥ ħ/2 using standard deviations.
:::

## Why it matters

ħ is tiny, about 1.055 × 10⁻³⁴ J s, so for everyday objects the limit is invisible. A 0.145 kg baseball located to within a micrometre has a velocity spread of at least ħ/(2mΔx) ≈ 4 × 10⁻²⁸ m/s.

For an electron in an atom it decides everything. Confine an electron to about Δx ≈ 10⁻¹⁰ m, the size of an atom:

::: equation Momentum and energy of a confined electron
\Delta p \geq \frac{\hbar}{2\Delta x} \approx 5 \times 10^{-25}\,\text{kg m/s}, \qquad \frac{(\Delta p)^2}{2m} \approx 1.5 \times 10^{-19}\,\text{J} \approx 1\,\text{eV}
:::

Squeezing the electron closer to the nucleus would lower its electrical potential energy but raise this kinetic energy. The balance between the two sets the size of atoms, and it is why an electron does not spiral into the nucleus, as classical physics predicted it should. The estimate gives the right order of magnitude: hydrogen's binding energy is 13.6 eV.

The same reasoning explains the zero-point energy of the previous lesson. A particle in a box of width L has Δx below L, so its momentum cannot be exactly zero, and neither can its energy.

::: model wavefunction
Each level in the box is a standing wave, a mix of momenta +nπħ/L and −nπħ/L. Its position is confined to the box, but its momentum is not definite: it points both ways at once.
:::

There is a related relation between energy and time, ΔE Δt ≳ ħ/2, with a different meaning: a state that lasts only a time Δt has an energy spread of about ħ/Δt. Short-lived excited atoms emit spectral lines that are slightly broadened for this reason.

::: note
The uncertainty principle is not about clumsy instruments. A wave localised in space must contain many wavelengths, and with p = ħk many wavelengths means many momenta. Δx Δp ≥ ħ/2 holds for every quantum state.
:::

::: quiz
? A wave packet is squeezed to half its width. The minimum spread in its momentum
- halves
- stays the same
+ doubles
= Δp ≥ ħ/(2Δx): halving Δx doubles the lower bound on Δp.
? Why is the uncertainty principle not just a limit of measuring instruments?
- Because modern instruments are already perfect
+ Because no wave has both a definite position and a single definite wavelength
- Because measurements never disturb quantum particles
= The limit is a property of the state itself: a localised packet must be built from many wavelengths.
? Roughly why does an electron not fall into the nucleus?
- The nucleus repels it at short range
+ Confining it more tightly raises its kinetic energy more than it lowers its potential energy
- It moves too fast to be captured
= Δp ≥ ħ/(2Δx), so the kinetic energy grows as the electron is squeezed, and the total energy has a minimum at about atomic size.
:::
