Faraday's law says that a changing magnetic field makes an electric field. Maxwell's correction to Ampère's law says that a changing electric field makes a magnetic field. Each statement on its own describes one field being produced by another. Put them together, in empty space with no charges or currents anywhere, and something new becomes possible: a disturbance in which each field keeps regenerating the other, so that the pair can break free of whatever started it and travel off on its own. Maxwell saw this in his equations in the 1860s. This lesson describes what such a wave looks like, why its electric and magnetic parts must be arranged the way they are, and what it carries with it.

## How can fields sustain each other?

Start with a single electric charge and shake it up and down. While it sat still, its electric field was steady, but now the field near it keeps changing. By the Ampère–Maxwell law, a changing electric field produces a magnetic field circling around the region where the change happens. That magnetic field is also changing, since it was produced by something that keeps changing, and by Faraday's law a changing magnetic field produces an electric field, a little further out. That new electric field changes in its turn and produces a magnetic field further out still.

Each field's change creates the other one slightly further from the source, and so the disturbance moves outward, step by step, away from the shaking charge. Once it has left, it no longer depends on the charge at all: stop shaking, and the part already on its way carries on. This travelling disturbance is an **electromagnetic wave**.

There is a deep difference between this and the waves met earlier in the course. A sound wave is a vibration of air; a water wave is a motion of water. Both need a material to travel through, and without one they cannot exist. An electromagnetic wave needs nothing of the kind. What waves is the electric and magnetic field itself, and fields can exist in a perfect vacuum. Physicists of the nineteenth century found this hard to accept, and almost all of them assumed space must be filled with some invisible medium, called the ether, for the waves to travel in. That assumption, and how it was abandoned, belongs to the chapter on relativity.

## What shape does the wave have?

Close to the shaking charge the pattern of fields is complicated. But far away, over any region small compared with the distance from the source, the wave looks much simpler, just as the surface of the Earth looks flat over a small field even though the planet is round. There it becomes a **plane wave**: on every flat surface at right angles to the direction of travel, the fields are the same at every point. To describe it, set up axes with the wave travelling along x, and suppose the electric field points along y and the magnetic field along z.

::: equation A plane electromagnetic wave travelling along x
E_y &= E_0 \sin(kx - \omega t) \\ B_z &= B_0 \sin(kx - \omega t)
:::

These are the same travelling sine waves met in the chapter on waves. E₀ and B₀ are the **amplitudes**, the largest values each field reaches. The quantity k = 2π/λ is called the **wavenumber**, where λ is the wavelength, the distance from one crest to the next. The quantity ω = 2πf is the **angular frequency**, where f is the frequency, the number of crests passing a point each second. A particular crest sits wherever kx − ωt has a particular value. As time t increases, x must increase to keep that value the same, so the crest moves forward, at a speed v = ω/k, which is the same as fλ.

Maxwell's equations do not allow just any pair of fields to travel like this. They fix several features of the wave, and each has a reason.

First, the wave is **transverse**: both fields point at right angles to the direction of travel, never along it. The reason lies in the two Gauss's laws. In empty space there are no charges, so neither field can have any sources, which means the flux of each field out of any closed surface is zero. Now picture a flat box with its faces at right angles to the direction of travel. If the field had a part pointing along x, and that part rose and fell along the wave, more field would leave through one face of the box than entered through the other, and the box would seem to contain a source. Since it cannot, the fields have no part along the direction of travel.

Second, the electric and magnetic fields are at right angles to each other as well, and they rise and fall **in step**, reaching their peaks at the same places at the same times. The direction of travel is the direction of the cross product **E** × **B**: point the fingers of your right hand along E, curl them toward B, and your thumb points the way the wave goes.

Third, the sizes of the two fields are locked together, and the ratio between them is set by the wave's speed. That is the next thing to derive.

::: model wave
Change the wavelength and the frequency. Their product is the wave's speed; for an electromagnetic wave in a vacuum that product always comes out the same, whatever wavelength you choose.
:::

## How does Faraday's law link E to B?

To find the link between the two fields, apply Faraday's law to a small loop placed in the wave. The loop should be chosen so that the electric field runs along two of its sides and the magnetic field passes through it. So take a thin rectangle lying in the xy-plane, with width dx along the direction of travel and height h along the electric field. The magnetic field, pointing along z, passes straight through it.

::: steps E = vB from Faraday's law
$ \oint \vec{E} \cdot \mathrm{d}\vec{l} = \left[E_y(x + \mathrm{d}x) - E_y(x)\right] h
Go round the rectangle anticlockwise as seen from the +z side, which is the direction that Faraday's law pairs with flux counted along +z. The two short sides run along x, at right angles to the electric field, so they contribute nothing. The two long sides run along y, the direction of the field: going up the side at x + dx contributes E_y(x + dx) times h, and coming down the side at x contributes minus E_y(x) times h.
---
$ \oint \vec{E} \cdot \mathrm{d}\vec{l} = \frac{\partial E_y}{\partial x}\,\mathrm{d}x\,h
The difference between the field at two points a tiny distance dx apart is the rate at which the field changes with x, multiplied by dx. The symbol ∂ marks a **partial derivative**, a rate of change with respect to one variable while the others are held fixed; here we hold the time fixed and look along x.
---
$ -\frac{\mathrm{d}\Phi_B}{\mathrm{d}t} = -\frac{\partial B_z}{\partial t}\,h\,\mathrm{d}x
Now the other side of Faraday's law. The rectangle is so thin that the magnetic field is effectively the same all over it, so the flux through it is B_z times its area, h dx. The rate of change of the flux is therefore the rate of change of B_z with time, at that place, times the area.
---
$ \frac{\partial E_y}{\partial x} = -\frac{\partial B_z}{\partial t}
Faraday's law sets the two sides equal. The factor h dx appears on both sides and cancels, so the size of the rectangle does not matter, which is a good sign: we now have a statement about the fields at a single place. It says an electric field that varies from place to place must come with a magnetic field that varies in time.
---
$ kE_0 \cos(kx - \omega t) = \omega B_0 \cos(kx - \omega t)
Put in the sine waves. The rate of change of sin(kx − ωt) with x is k cos(kx − ωt), and its rate of change with t is −ω cos(kx − ωt). That second minus sign cancels the minus sign already on the right-hand side.
---
$ E_0 = \frac{\omega}{k} B_0 = vB_0
Both sides contain the same cosine, which cancels, so the equation holds at every place and time only if kE₀ = ωB₀. Dividing by k and recognising ω/k as the wave's speed v gives the result: the electric field amplitude equals the wave speed times the magnetic field amplitude.
:::

So the two fields are not independent: once you know the strength of one, the speed of the wave tells you the strength of the other. The same argument applied to the Ampère–Maxwell law, using a rectangle in the xz-plane instead, gives a second relation between E₀ and B₀, and the two relations together fix the speed itself. That calculation is the subject of the next lesson. Its answer is about 3.00 × 10⁸ m/s.

That large number tells us something about the balance between the two fields. In the units we use, the magnetic part of an electromagnetic wave looks small. Sunlight at the Earth, for instance, has a peak electric field of about 1,000 volts per metre, so its magnetic field is only about 1,000 ÷ (3 × 10⁸) ≈ 3 × 10⁻⁶ T, less than a tenth of the Earth's own magnetic field. Yet the two parts are equal partners, as the next section shows.

## What does the wave carry?

A wave that can heat the ground must carry energy, and an electromagnetic wave carries it in its fields. Both electric and magnetic fields store energy, and in an electromagnetic wave the energy is shared exactly equally between the two, despite the small-looking number for B. The energy flowing across each square metre at right angles to the wave, per second, is called the wave's **intensity**, measured in watts per square metre. Sunlight arriving at the Earth's distance from the Sun delivers about 1,360 W/m². That energy has crossed about 150 million kilometres of nearly empty space, with nothing to carry it but the fields, and it heats the ground and drives photosynthesis in every green plant.

The wave also carries momentum, which means that light pushes on whatever absorbs it. The pressure equals the intensity divided by the speed of light. For sunlight that is completely absorbed, this comes to 1,360 ÷ (3 × 10⁸) ≈ 4.5 × 10⁻⁶ pascals, less than a ten-billionth of the pressure of the atmosphere. It is far too small to feel, but in space, with no air resistance and no friction, it can be put to use: a **solar sail**, a huge, thin, reflective sheet attached to a spacecraft, is pushed steadily by sunlight and can slowly build up great speed without carrying any fuel.

::: note
An electromagnetic wave is a self-sustaining pair of fields: a changing electric field makes a magnetic field, and a changing magnetic field makes an electric field. The wave is transverse, its two fields are perpendicular and in step, and their strengths are locked together by E = cB. It needs no medium, because what waves is the field itself.
:::

::: history 1886–1888 · Karlsruhe, Germany
Maxwell's prediction remained untested for more than twenty years, because nobody knew how to make or detect electromagnetic waves other than light. Heinrich Hertz, a young professor at the technical college in Karlsruhe, managed both. He made sparks jump across a small gap in a circuit; each spark sent current surging rapidly back and forth, which, by Maxwell's theory, should send out waves. To detect them he used a simple loop of wire with its own tiny gap, placed across the room, and in a darkened laboratory he saw minute sparks jump across it each time the first circuit sparked. He then showed that the disturbance behaved exactly like a wave: it reflected off a metal wall to form standing waves, with a wavelength he could measure, of several metres in his first experiments; it could be focused by curved mirrors and polarised; and it travelled at a speed comparable to that of light. Asked what use these waves might have, Hertz is reported to have said that they had none. Within a decade Guglielmo Marconi was sending radio signals across kilometres, and soon across the Atlantic.
:::

::: quiz
? A plane electromagnetic wave travels north, and at some instant its electric field points east. Which way can its magnetic field point?
- North
+ Up or down
- East
= The magnetic field must be at right angles both to the electric field and to the direction of travel, which leaves only up or down. The rule that the wave travels along E × B decides which of the two it is.
? Why can an electromagnetic wave cross empty space when sound cannot?
- Light is made of particles
+ The wave is in the fields themselves, which each regenerate the other
- Space contains a thin gas that carries light
= Sound is a vibration of matter, so it needs air or some other material. An electromagnetic wave consists of electric and magnetic fields, which can exist in a vacuum and keep producing each other as they travel.
? The electric field amplitude of a wave in a vacuum is doubled. What happens to its magnetic field amplitude?
- It halves
- It is unchanged
+ It doubles
= Faraday's law requires E₀ = cB₀, and the speed c is fixed. So the two amplitudes are proportional, and doubling one doubles the other.
:::
