A fusion power plant would burn hydrogen isotopes into helium, with no long-lived radioactive waste from the fuel, no possibility of a runaway chain reaction, and fuel extracted from seawater and lithium. Physicists have worked toward one since the early 1950s. The problem has always been the same: hold a gas at over 100 million kelvin long enough for it to give back more energy than it took to heat.

## The fuel

The easiest reaction is deuterium–tritium (D–T) fusion, which releases 17.6 MeV: 3.5 MeV to the helium nucleus, called an alpha particle, and 14.1 MeV to a neutron. It has the largest reaction rate at the lowest temperature, though "lowest" means 10 to 20 keV, or about 100 to 200 million kelvin, around ten times the Sun's core. The Sun can fuse protons at lower temperatures because it has an enormous volume and billions of years; a reactor has neither.

At these temperatures the fuel is a **plasma**, a gas of bare nuclei and free electrons. Deuterium is abundant: about one hydrogen atom in 6,500 in seawater. Tritium is radioactive, with a half-life of 12.3 years, and almost none exists naturally. A power plant would have to breed its own by letting the fusion neutrons strike lithium in a blanket around the plasma.

## The Lawson criterion

::: history 1955 · Harwell, England
John Lawson, an engineer at the UK Atomic Energy Research Establishment, worked out what a fusion reactor must achieve to produce net power. His report, classified at first, was published in 1957. The condition now bears his name.
:::

A plasma loses heat constantly. The **energy confinement time** τ_E is the time it would take the plasma to lose its stored heat if heating stopped. A plasma is said to **ignite** when the alpha particles from fusion, which are charged and stay in the plasma, heat it as fast as it loses energy, so that no outside heating is needed.

::: steps The ignition condition
$ P_\alpha = \tfrac{1}{4} n^2 \langle \sigma v \rangle E_\alpha
For an equal mix of deuterium and tritium, each with density n/2, the reaction rate per volume is (n/2)²⟨σv⟩, where ⟨σv⟩ is the reaction rate coefficient, which depends on temperature. Each reaction deposits E_α = 3.5 MeV in the plasma; the neutron escapes.
---
$ W = 3 n k T
The heat stored per volume: ³⁄₂kT for each of n ions and ³⁄₂kT for each of n electrons.
---
$ P_{\mathrm{loss}} = \frac{3 n k T}{\tau_E}
Losses drain the stored heat over the confinement time.
---
$ \tfrac{1}{4} n^2 \langle \sigma v \rangle E_\alpha \geq \frac{3 n k T}{\tau_E}
Ignition: alpha heating at least matches the losses.
---
$ n T \tau_E \geq \frac{12\,k T^2}{\langle \sigma v \rangle E_\alpha}
Divide by n, multiply by T τ_E, and rearrange. Between about 10 and 20 keV, ⟨σv⟩ rises roughly as T², so the right-hand side is nearly constant there.
---
$ n T \tau_E \gtrsim 3 \times 10^{21}\,\text{keV s/m}^3
Inserting measured D–T rates gives a minimum near 15 keV. This is the **triple product** a reactor must reach.
:::

The criterion says the product of density, temperature and confinement time matters, not any one of them. That leaves two strategies. Magnetic confinement holds a thin plasma, about 10²⁰ particles per cubic metre (a few millionths of the density of air), for several seconds. Inertial confinement compresses a tiny pellet to enormous density for well under a billionth of a second, relying on the fuel's own inertia to hold it together while it burns.

::: note
Fusion power needs a plasma hot enough (about 150 million K), dense enough and confined long enough that the triple product nTτ_E exceeds about 3 × 10²¹ keV·s/m³. Then the fusion alphas keep the plasma hot on their own.
:::

## Tokamaks

Charged particles spiral around magnetic field lines rather than crossing them. Bend the field lines into closed loops around a doughnut-shaped chamber and the plasma has no ends to escape from. Many such designs were tried in the 1950s and 1960s; most leaked plasma far faster than predicted.

::: history 1968–1969 · Moscow and Culham
Soviet physicists at the Kurchatov Institute, building on a 1951 idea by Igor Tamm and Andrei Sakharov, reported that their T-3 **tokamak** (a Russian acronym for "toroidal chamber with magnetic coils") held electrons at about 10 million kelvin, far better than any Western machine. The claim was doubted until a British team flew to Moscow in 1969 with a laser system and confirmed it. Laboratories around the world switched to tokamaks.
:::

Progress since has been steady but slow. The Joint European Torus (JET) in England produced 16 MW of fusion power in 1997, briefly about two-thirds of the heating power put in. In its final experiments in 2023 it released 69 megajoules over about 5 seconds, a record for energy from a single pulse.

**ITER**, under construction at Cadarache in southern France, is a collaboration of the European Union, China, India, Japan, South Korea, Russia and the United States, agreed in 2006. It is designed to produce 500 MW of fusion power from 50 MW of heating, a gain Q of 10, in pulses of several minutes. It will not generate electricity. After long delays and cost growth, its 2024 schedule plans the start of research operations in 2034 and deuterium–tritium operation in 2039.

## Ignition at NIF

::: history 5 December 2022 · Livermore, California
At the National Ignition Facility, 192 laser beams delivered 2.05 megajoules to a millimetre-sized capsule of deuterium and tritium, compressing and heating it so that it fused before it could fly apart. It released 3.15 megajoules, the first time a controlled fusion experiment produced more energy than was delivered to the fuel. NIF has repeated ignition several times since, with higher yields.
:::

The gain counted only the laser light reaching the target. The lasers drew about 300 megajoules from the electrical grid, and NIF fires about once a day; a power plant would need several shots per second with far more efficient lasers.

Even after a gain is achieved in the plasma, a working reactor must breed its own tritium, survive years of bombardment by 14 MeV neutrons that damage the wall materials, and run continuously at a cost competitive with other sources. Fusion has been "thirty years away" for most of its history. The physics is now demonstrated; the engineering is the remaining work.

::: quiz
? Why must a fusion reactor confine its plasma for a minimum time, not just heat it?
- Fusion reactions take seconds to complete
+ The plasma loses energy continuously, and alpha heating must keep up with those losses
- The neutrons need time to breed tritium
= The Lawson criterion balances fusion heating against heat lost over the confinement time τ_E.
? Why can't the neutrons from D–T fusion help keep the plasma hot?
- They carry almost no energy
+ They have no charge, so the magnetic field does not hold them and they escape
- They are absorbed by deuterium
= Only the charged alpha particles stay in the plasma; the 14.1 MeV neutrons leave and deposit their energy in the surrounding blanket.
? NIF's 2022 shot released 3.15 MJ from 2.05 MJ of laser light. Why was it not net energy for the facility?
- Most of the fusion energy was lost as neutrinos
+ The lasers needed about 300 MJ from the grid to produce that 2.05 MJ of light
- The capsule did not actually ignite
= The gain counted only energy delivered to the target; the lasers are very inefficient.
:::
