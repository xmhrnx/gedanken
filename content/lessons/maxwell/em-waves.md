Faraday's law says a changing magnetic field makes an electric field. Maxwell's correction to Ampère's law says a changing electric field makes a magnetic field. Put those together in empty space, with no charges or currents anywhere, and something remarkable becomes possible: a disturbance in which each field keeps regenerating the other, moving off on its own.

## How the fields sustain each other

Shake a charge up and down. Its electric field nearby changes, and that changing E produces a magnetic field circling around it. That magnetic field is also changing, so it produces an electric field a little further out. Each field's change creates the other one slightly further from the source. The pattern detaches from the charge and travels outward. This is an **electromagnetic wave**, and it needs no medium: the fields themselves are what waves.

## The shape of the wave

Far from the source, over a small region, the wave looks like a **plane wave**: the fields are the same everywhere on each flat surface perpendicular to the direction of travel. Take the wave moving along x, with E pointing along y and B along z.

::: equation A plane electromagnetic wave travelling along x
E_y &= E_0 \sin(kx - \omega t) \\ B_z &= B_0 \sin(kx - \omega t)
:::

Here k = 2π/λ is the wavenumber (λ the wavelength) and ω = 2πf the angular frequency (f the frequency). A crest sits where kx − ωt is constant, so it moves at speed v = ω/k = fλ.

Maxwell's equations fix several features of this wave:

- It is **transverse**: E and B both point perpendicular to the direction of travel, and perpendicular to each other. Gauss's laws forbid a field component along the direction of travel in a plane wave.
- E and B rise and fall **in step**, peaking at the same places and times.
- The direction of travel is the direction of **E** × **B**.
- The sizes of E and B are locked together, as the next derivation shows.

::: model wave
Change the wavelength and frequency. Their product is the speed; for an electromagnetic wave in vacuum it is always c.
:::

## Faraday's law links E to B

Apply Faraday's law to a thin rectangle in the xy-plane, of width dx along the wave and height h along E.

::: steps E = vB from Faraday's law
$ \oint \vec{E} \cdot \mathrm{d}\vec{l} = \left[E_y(x + \mathrm{d}x) - E_y(x)\right] h = \frac{\partial E_y}{\partial x}\,\mathrm{d}x\,h
Going round the rectangle, only the two sides parallel to E contribute, and they are traversed in opposite directions. What survives is the difference of E across the width dx.
---
$ -\frac{\mathrm{d}\Phi_B}{\mathrm{d}t} = -\frac{\partial B_z}{\partial t}\,h\,\mathrm{d}x
The magnetic flux through the rectangle is B_z times its area h dx.
---
$ \frac{\partial E_y}{\partial x} = -\frac{\partial B_z}{\partial t}
Faraday's law sets the two equal; cancel h dx. A field that varies in space must be accompanied by a field that varies in time.
---
$ kE_0 \cos(kx - \omega t) = \omega B_0 \cos(kx - \omega t)
Differentiate the sine waves: ∂/∂x brings out k, and ∂/∂t brings out −ω, which cancels the minus sign.
---
$ E_0 = \frac{\omega}{k} B_0 = vB_0
Cancel the cosine. The ratio of the field strengths is the wave's speed.
:::

The same argument applied to the Ampère–Maxwell law gives a second relation between E₀ and B₀, and the two together fix the speed. That is the subject of the next lesson. The answer is about 3.00 × 10⁸ m/s, so the peak electric field in sunlight, about 1000 V/m, comes with a magnetic field of only about 3 × 10⁻⁶ T, less than a tenth of Earth's own field.

## Energy and momentum

The wave carries energy, shared equally between its electric and magnetic fields. The power crossing each square metre is the **intensity**. Sunlight at Earth's distance delivers about 1360 W/m², enough to heat the ground and drive photosynthesis across a hundred and fifty million kilometres of empty space. The wave also carries momentum: light pushes on what it hits, with a pressure equal to intensity divided by c, about 4.5 micropascals for sunlight fully absorbed. Solar sails exploit this.

::: note
An electromagnetic wave is a self-sustaining pair of fields: changing E makes B, changing B makes E. It is transverse, E and B are in phase and perpendicular, and E = cB. No medium is needed.
:::

::: history 1886–1888 · Karlsruhe
Heinrich Hertz made a spark jump across a gap in a circuit and detected the result as a tiny spark in a separate loop of wire across the room. He showed the disturbance had a wavelength of a few metres, formed standing waves when reflected off a metal wall, could be focused and polarised, and travelled at a speed comparable to light's. Asked what use it was, he reportedly said none. Within a decade Marconi was sending radio signals across kilometres.
:::

::: quiz
? A plane electromagnetic wave travels north and its electric field points east at some instant. Which way can its magnetic field point?
- North
+ Up or down
- East
= B must be perpendicular to both E and the direction of travel; the direction of E × B gives which one.
? Why can an electromagnetic wave cross empty space when sound cannot?
- Light is made of particles
+ The wave is in the fields themselves, which each regenerate the other
- Space contains a thin gas
= Sound is a vibration of matter. An electromagnetic wave needs only electric and magnetic fields, which exist in a vacuum.
? Doubling the electric field amplitude of a wave in vacuum does what to its magnetic field amplitude?
- Halves it
- Leaves it unchanged
+ Doubles it
= E₀ = cB₀ with c fixed, so the amplitudes are proportional.
:::
