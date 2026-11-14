Two numbers sit inside Maxwell's equations, and both come from tabletop experiments that have nothing to do with light. One, ε₀, is found by measuring the force between charges or the charge a capacitor holds. The other, μ₀, is found from the force between two wires carrying current. Nobody measuring them was thinking about light at all. Yet when Maxwell worked out how fast his electromagnetic waves should travel, these two numbers combined into a speed, and the speed was the speed of light. It was one of the great moments in physics: optics, the study of light, which had been a separate science for two thousand years, turned out to be a branch of electromagnetism. This lesson derives that speed and tells how the match was first noticed.

## Where do the two constants come from?

The first constant, **ε₀**, called the permittivity of free space, measures how strong an electric field a given charge produces in empty space. It appears in Coulomb's law for the force between two charges q₁ and q₂ a distance r apart, F = q₁q₂/(4πε₀r²). You can find it by measuring the force between charged objects, or, more practically, by measuring how much charge a pair of metal plates of known size and separation holds at a given voltage. Its value is about 8.85 × 10⁻¹² in SI units.

The second, **μ₀**, called the permeability of free space, plays the same role for magnetism: it measures how strong a magnetic field a given current produces. It appears in the force between two long parallel wires a distance r apart, carrying currents I₁ and I₂, which attract each other with a force of μ₀I₁I₂/(2πr) on each metre of their length. Its value is about 1.26 × 10⁻⁶ in SI units.

Look at what these experiments involve: charged plates, wires, batteries, springs and balances to measure forces. Neither mentions light, or waves, or anything moving fast. That is what makes the result so striking.

## How fast does an electromagnetic wave travel?

The previous lesson set up a plane wave travelling along x, with its electric field along y and its magnetic field along z, both varying as sine waves: E_y = E₀ sin(kx − ωt) and B_z = B₀ sin(kx − ωt), where the wave's speed is v = ω/k. Applying Faraday's law to a thin rectangle gave one link between the two amplitudes: E₀ = vB₀. That link involves the speed, but it cannot tell us what the speed is, because one equation cannot fix both the speed and the ratio of the fields. We need a second, independent link, and the Ampère–Maxwell law provides it. In empty space there are no currents, so only Maxwell's displacement-current term survives, and that is where the constants enter.

::: steps Wave speed from Maxwell's equations
$ E_0 = vB_0
Start from the result of the previous lesson, which came from applying Faraday's law to a thin rectangle in the xy-plane. It says that a changing magnetic field must accompany a varying electric field, with amplitudes in the ratio of the wave's speed.
---
$ \oint \vec{B} \cdot \mathrm{d}\vec{l} = \mu_0 \varepsilon_0 \frac{\mathrm{d}\Phi_E}{\mathrm{d}t}
Now write down the Ampère–Maxwell law with the current term removed, since there is no current in empty space. To use it, we need a loop with the magnetic field running along two of its sides and the electric field passing through it. So take a thin rectangle in the xz-plane, of width dx along the direction of travel and height h along the magnetic field. The electric field, pointing along y, passes straight through it.
---
$ \oint \vec{B} \cdot \mathrm{d}\vec{l} = \left[B_z(x) - B_z(x + \mathrm{d}x)\right] h = -\frac{\partial B_z}{\partial x}\,\mathrm{d}x\,h
Go round the rectangle in the direction that pairs with flux counted along +y; by the right-hand rule that means going up the side at x and down the side at x + dx. The two short sides run along x, at right angles to B, and contribute nothing. The two long sides contribute B_z(x) times h and minus B_z(x + dx) times h. Their difference is minus the rate of change of B_z with x, times dx.
---
$ \mu_0 \varepsilon_0 \frac{\mathrm{d}\Phi_E}{\mathrm{d}t} = \mu_0 \varepsilon_0 \frac{\partial E_y}{\partial t}\,h\,\mathrm{d}x
On the other side of the law, the electric flux through the thin rectangle is E_y times its area, h dx. Its rate of change is the rate at which E_y changes with time at that place, times the area, all multiplied by μ₀ε₀.
---
$ -\frac{\partial B_z}{\partial x} = \mu_0 \varepsilon_0 \frac{\partial E_y}{\partial t}
Set the two sides equal, as the law requires. The factor h dx appears on both sides and cancels, leaving a statement about the fields at a single place: a magnetic field that varies along the wave must come with an electric field that varies in time.
---
$ -kB_0 \cos(kx - \omega t) = -\mu_0 \varepsilon_0\, \omega E_0 \cos(kx - \omega t)
Put in the sine waves. The rate of change of sin(kx − ωt) with x is k cos(kx − ωt), and with t it is −ω cos(kx − ωt). So both sides pick up a cosine and a minus sign.
---
$ B_0 = \mu_0 \varepsilon_0 \frac{\omega}{k} E_0 = \mu_0 \varepsilon_0\, v E_0
The same cosine appears on both sides, so for the equation to hold at every place and time the amplitudes must match: the cosines cancel and so do the minus signs. Rearranging and writing v for ω/k gives the second link between the amplitudes. Unlike the first, this one contains the two constants.
---
$ B_0 = \mu_0 \varepsilon_0\, v\, (vB_0) = \mu_0 \varepsilon_0 v^2 B_0
Now combine the two links. Replace E₀ in the previous line with vB₀, from the first step. The magnetic amplitude B₀ now appears on both sides.
---
$ v^2 = \frac{1}{\mu_0 \varepsilon_0}
Divide both sides by B₀, which is allowed because a real wave has a non-zero amplitude. The amplitude disappears completely, so the result applies to strong waves and weak ones alike. Then divide by μ₀ε₀ to leave v² on its own.
---
$ v = \frac{1}{\sqrt{\mu_0 \varepsilon_0}}
Take the square root. The speed depends on nothing except the two constants: not on the wavelength, not on the amplitude, and not on what produced the wave. Every electromagnetic wave in a vacuum must travel at this one speed.
:::

## What number comes out?

Now put in the measured values of the two constants.

::: equation Putting in the measured constants
v &= \frac{1}{\sqrt{(1.2566 \times 10^{-6}\,\text{N/A}^2)(8.8542 \times 10^{-12}\,\text{F/m})}} \\ &= \frac{1}{\sqrt{1.1127 \times 10^{-17}\,\text{s}^2/\text{m}^2}} \\ &\approx 2.998 \times 10^8\,\text{m/s}
:::

The first thing to check is the units, because a formula that claims to give a speed had better produce metres per second. A farad, the unit of capacitance, is a coulomb per volt, which works out to A² s²/(N m). So the product of N/A² and F/m becomes s²/m², the units cancel to leave seconds squared per metre squared, and one over the square root of that is metres per second. It really is a speed.

And it is not just any speed. It is about 300,000 kilometres per second, the measured speed of light. Electromagnetic waves of every wavelength travel at this one speed through a vacuum, and light travels at exactly that speed too. The conclusion is very hard to avoid: light is an electromagnetic wave. This speed is now given its own symbol, c, and the formula is written c = 1/√(μ₀ε₀).

It is worth pausing on how strange this is. ε₀ describes how charges push on each other when they sit still. μ₀ describes how currents push on each other. Neither experiment involves anything travelling, and yet their combination is a speed, and the speed at which light crosses the universe. Electricity, magnetism and light, three subjects that had been studied for centuries by different methods, were one subject.

::: note
Constants measured with charged plates and current-carrying wires, with no light involved, combine to predict a wave speed of c = 1/√(μ₀ε₀) ≈ 3.00 × 10⁸ m/s. That is the measured speed of light, which showed that light is an electromagnetic wave and joined optics to electricity and magnetism.
:::

## How did Maxwell first notice the match?

::: history 1862 · London and Glenlair, Scotland
Maxwell did not have the clean derivation above when he first saw the connection. In 1861 he was working on "On Physical Lines of Force", in which he pictured space as filled with a mechanical medium of tiny spinning vortices. He calculated how fast a sideways wobble would travel through this imagined medium, and found that the speed came out equal to a ratio between electrical and magnetic units. That ratio had been measured in 1856 by Wilhelm Weber and Rudolf Kohlrausch in Germany, by comparing the same charge measured in two different ways, and from their result the speed came to about 310,700 kilometres per second. Maxwell was spending the summer at his family's estate at Glenlair in Scotland, far from his reference books, and could not check the number against the speed of light until he returned to London in the autumn. When he did, he found that Hippolyte Fizeau in Paris had measured the speed of light in 1849, with a rapidly spinning toothed wheel, at about 315,000 kilometres per second. The two figures agreed within the accuracy of the experiments. Maxwell wrote, putting the words in italics: "we can scarcely avoid the inference that light consists in the transverse undulations of the same medium which is the cause of electric and magnetic phenomena."
:::

The vortices were a scaffold, and Maxwell took them down once the building stood.

::: history 1865 · London
In "A Dynamical Theory of the Electromagnetic Field", Maxwell derived the waves from the field equations alone, with no vortices or wheels at all. He still believed the waves needed a medium to travel in, the "luminiferous ether", as almost every physicist did for another forty years. He did not live to see his prediction confirmed in the laboratory: he died in 1879, and it was eight years later, in 1887 and 1888, that Heinrich Hertz produced and detected electromagnetic waves and showed they travelled at roughly the speed of light.
:::

## Relative to what?

The derivation produced one speed, c, and said nothing about who measures it. That raises an awkward question. Speeds are normally measured relative to something: a car's speed relative to the road, a boat's relative to the water. So relative to what does light travel at c?

For Maxwell's generation, the answer seemed obvious: relative to the ether, the medium the waves were assumed to travel in, just as sound travels at a fixed speed relative to the air. Then someone moving through the ether should measure light going faster when moving toward it and slower when moving away, just as you would measure sound differently while riding in a fast car. Testing that prediction is the story of the chapter on special relativity, and the answer astonished everyone: everyone measures the same speed c, however they are moving. Maxwell's equations had been right in a deeper sense than their author knew.

Today the logic runs the other way round. Since 1983 the metre has been defined by the distance light travels in a set fraction of a second, which makes c exactly 299,792,458 metres per second by definition, no longer a measured number. And since 2019, μ₀ has been a measured quantity, very close to the old value of 4π × 10⁻⁷ N/A², with ε₀ then fixed by the relation ε₀ = 1/(μ₀c²). The formula that once predicted the speed of light is now used to pin down one of the constants it was built from.

::: quiz
? According to Maxwell's derivation, how does the speed of an electromagnetic wave in a vacuum depend on its wavelength?
- Longer waves travel faster
- Shorter waves travel faster
+ It does not depend on wavelength at all
= The speed v = 1/√(μ₀ε₀) contains only the two constants, and the wavelength cancelled out in the derivation. Radio waves and X-rays travel at the same speed in a vacuum.
? If ε₀ were four times larger, with μ₀ unchanged, what would the speed of light be?
+ Half as large
- Twice as large
- A quarter as large
= The speed goes as 1/√ε₀. Multiplying ε₀ by 4 multiplies the speed by 1/√4, which is one half.
? Why was Maxwell's calculation such persuasive evidence that light is electromagnetic?
- He measured the speed of light himself
+ A speed built only from electrical and magnetic measurements matched the independently measured speed of light
- It explained the colours of the rainbow
= The constants ε₀ and μ₀ come from experiments with charges and currents that involve no light at all. That they combine to give exactly the speed of light would be an extraordinary coincidence unless light is an electromagnetic wave.
:::
