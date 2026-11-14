Two numbers in Maxwell's equations come from tabletop experiments with no light in them. ε₀ is measured by charging a capacitor; μ₀ is set by the force between two wires carrying current. Maxwell found that his equations combine them into a speed, and that the speed is the speed of light.

## The two constants

- **ε₀**, the permittivity of free space, sets how strong an electric field a given charge makes. It appears in Coulomb's law, F = q₁q₂/(4πε₀r²).
- **μ₀**, the permeability of free space, sets how strong a magnetic field a given current makes. Two parallel wires a distance r apart carrying currents I₁ and I₂ attract with a force μ₀I₁I₂/(2πr) per metre of length.

Neither definition mentions light, or motion, or waves.

## The derivation

Take the plane wave from the previous lesson, moving along x with E along y and B along z:

E_y = E₀ sin(kx − ωt), B_z = B₀ sin(kx − ωt), with speed v = ω/k.

Faraday's law already gave one link between the amplitudes. The Ampère–Maxwell law gives another. There are no currents in empty space, so only the displacement-current term survives.

::: steps Wave speed from Maxwell's equations
$ E_0 = vB_0
From Faraday's law applied to a thin rectangle in the xy-plane, as shown in the previous lesson.
---
$ \oint \vec{B} \cdot \mathrm{d}\vec{l} = \mu_0 \varepsilon_0 \frac{\mathrm{d}\Phi_E}{\mathrm{d}t}
The Ampère–Maxwell law with no current. Apply it to a thin rectangle in the xz-plane, of width dx along the wave and height h along B, so that E passes through it.
---
$ -\frac{\partial B_z}{\partial x}\,\mathrm{d}x\,h = \mu_0 \varepsilon_0 \frac{\partial E_y}{\partial t}\,h\,\mathrm{d}x
Going round the rectangle the right way for its normal (along +y), the two sides parallel to B contribute the difference of B across dx, with a minus sign. The electric flux is E_y times the area h dx.
---
$ -kB_0 \cos(kx - \omega t) = -\mu_0 \varepsilon_0\, \omega E_0 \cos(kx - \omega t)
Cancel h dx and differentiate the sine waves.
---
$ B_0 = \mu_0 \varepsilon_0 \frac{\omega}{k} E_0 = \mu_0 \varepsilon_0\, v E_0
Cancel the cosine and the minus signs, and use ω/k = v.
---
$ B_0 = \mu_0 \varepsilon_0\, v (vB_0) = \mu_0 \varepsilon_0 v^2 B_0
Substitute E₀ = vB₀ from the first step.
---
$ v = \frac{1}{\sqrt{\mu_0 \varepsilon_0}}
Cancel B₀ and take the square root. The speed depends on nothing but the two constants: not on the wavelength, not on the amplitude, not on the source.
:::

## The numbers

::: equation Putting in the measured constants
v &= \frac{1}{\sqrt{(1.2566 \times 10^{-6}\,\text{N/A}^2)(8.8542 \times 10^{-12}\,\text{F/m})}} \\ &= \frac{1}{\sqrt{1.1127 \times 10^{-17}\,\text{s}^2/\text{m}^2}} \\ &\approx 2.998 \times 10^8\,\text{m/s}
:::

The units work out: N/A² × F/m reduces to s²/m², so the result is a speed. And it is the measured speed of light, about 300,000 km/s. Electromagnetic waves in vacuum all travel at this speed, whatever their wavelength. The conclusion is hard to avoid: light is an electromagnetic wave.

::: note
Constants measured with capacitors and current-carrying wires predict the speed of light. c = 1/√(μ₀ε₀) unified optics with electricity and magnetism.
:::

::: history 1862 · London
Maxwell did not yet have the clean derivation above. In "On Physical Lines of Force" he calculated the speed of transverse waves in his imagined mechanical medium of vortices and found it equal to the ratio of electrical units that Wilhelm Weber and Rudolf Kohlrausch had measured in 1856: about 310,700 km/s. Hippolyte Fizeau had measured the speed of light in 1849 as about 315,000 km/s. Maxwell, working at his family estate in Scotland away from his reference books, had to wait until he returned to London to compare the numbers. He wrote, in italics, that "we can scarcely avoid the inference that light consists in the transverse undulations of the same medium which is the cause of electric and magnetic phenomena."
:::

::: history 1865 · London
In "A Dynamical Theory of the Electromagnetic Field" Maxwell derived the waves from the field equations alone, with no vortices. He still believed the waves needed a medium, the luminiferous ether, as almost everyone did for another forty years. Hertz produced and detected such waves in 1887–88, eight years after Maxwell's death.
:::

## What the speed is relative to

The derivation gives one speed, c, with no mention of who measures it. For Maxwell's generation this meant c relative to the ether, the medium the waves were assumed to travel in. Someone moving through the ether should then measure light going faster or slower. Testing that is the story of the special relativity chapter, and the answer was that everyone measures the same c.

Today the logic runs the other way. Since 1983 the metre has been defined so that c is exactly 299,792,458 m/s, and since 2019 μ₀ is a measured quantity, very close to 4π × 10⁻⁷ N/A², with ε₀ = 1/(μ₀c²).

::: quiz
? According to Maxwell's derivation, how does the speed of an electromagnetic wave in vacuum depend on its wavelength?
- Longer waves travel faster
- Shorter waves travel faster
+ It does not depend on wavelength at all
= v = 1/√(μ₀ε₀) contains only the two constants. Radio waves and X-rays travel at the same speed in vacuum.
? If ε₀ were four times larger, with μ₀ unchanged, the speed of light would be
+ half as large
- twice as large
- four times smaller
= c ∝ 1/√ε₀, and 1/√4 = 1/2.
? Why was Maxwell's calculation so persuasive that light is electromagnetic?
- He measured the speed of light himself
+ A speed built only from electrical and magnetic measurements matched the independently measured speed of light
- It explained the colours of the rainbow
= ε₀ and μ₀ come from experiments with no light in them, yet they combine to give light's speed.
:::
