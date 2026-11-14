A fusion power plant would do on Earth what the Sun does in its core: turn hydrogen into helium and collect the energy. The attractions are large. The fuel can be extracted from seawater and lithium, and counting the lithium dissolved in the oceans there is enough to last for millions of years. A fusion reactor cannot run away the way a fission reactor can, because there is no chain reaction to get out of control; if anything goes wrong, the hot fuel cools and the reactions simply stop. And the fuel burns to helium, so it leaves no long-lived radioactive waste of its own, although the walls of the machine do become radioactive. Physicists have worked toward this since the early 1950s. The problem has always been the same: hold a gas at more than 100 million kelvin long enough, and densely enough, for it to give back more energy than it took to heat. This lesson works out what "long enough" means and follows the two main ways people have tried to get there.

## Why that fuel, and why so hot?

A reactor cannot copy the Sun's own reaction. The Sun fuses ordinary protons, but the first step depends on the weak force and is so slow that an average proton waits billions of years. The Sun gets away with this because it has an enormous mass of fuel and all the time in the world. A reactor has a few grams of fuel and needs results in seconds.

So reactors use the reaction that goes fastest at the lowest temperature: the fusion of deuterium and tritium, the two heavy isotopes of hydrogen, usually shortened to D–T. It releases 17.6 MeV (million electronvolts) per reaction, shared unequally between the two products: 3.5 MeV goes to the helium nucleus, which in this context is called an **alpha particle**, and 14.1 MeV goes to a neutron. The split matters, as we will see, because the two products behave very differently afterwards.

Even D–T needs particle energies of 10 to 20 keV (thousand electronvolts) to fuse at a useful rate, which corresponds to roughly 120 to 230 million kelvin, around ten times hotter than the centre of the Sun. Nothing solid can touch matter at that temperature. At these energies every atom is stripped of its electrons, and the fuel becomes a **plasma**, a gas of bare nuclei and free electrons. Because a plasma is made of charged particles, it responds to electric and magnetic fields, and that fact is the basis of most attempts to hold one.

The two fuels have very different sources. Deuterium is plentiful: about one hydrogen atom in 6,500 in seawater is deuterium, and it can be separated cheaply. Tritium is radioactive, decaying with a half-life of 12.3 years, so almost none exists in nature. A power plant would have to make its own. The plan is to surround the plasma with a **blanket** containing lithium; when the fusion neutrons strike lithium nuclei, they split them and produce fresh tritium.

## How good is good enough?

::: history 1955 · Harwell, England
John Lawson, an engineer at the UK Atomic Energy Research Establishment, asked a plain question that the enthusiasts of the time had not answered: what must a fusion reactor actually achieve to produce more power than it consumes? He balanced the energy released by fusion against the energy the plasma loses, and found a condition that any reactor must meet whatever its design. His report was classified at first and published in 1957. The condition has carried his name ever since, and it is still the yardstick by which every fusion experiment is judged.
:::

A plasma is always losing heat: particles escape, and hot electrons radiate. The convenient measure of how well a machine holds its heat is the **energy confinement time**, written τ_E. It is the time the plasma would take to lose its stored heat if all heating were switched off. A good reactor has a long τ_E.

The goal is **ignition**. A plasma ignites when the heat from its own fusion reactions keeps it hot without any outside help, like a fire that, once lit, needs no more matches. Here the split of the 17.6 MeV becomes important. The neutron has no electric charge, so magnetic fields cannot hold it, and it flies straight out of the plasma. Only the charged alpha particle stays inside, bumping into other particles and sharing out its 3.5 MeV. So for ignition, the alpha heating alone must make up the losses.

::: steps The ignition condition
$ P_\alpha = \tfrac{1}{4} n^2 \langle \sigma v \rangle E_\alpha
Take an equal mix of deuterium and tritium ions with total density n, so each kind has density n/2. The number of reactions per cubic metre per second is the product of the two densities times ⟨σv⟩, the **reaction rate coefficient**, which measures how readily a pair fuses and depends strongly on temperature. That gives (n/2)²⟨σv⟩ reactions, and each leaves E_α = 3.5 MeV in the plasma. The result is the alpha heating power per cubic metre.
---
$ W = 3 n k T
Next, the heat stored in each cubic metre. A particle in a gas at temperature T carries on average 3/2 kT of kinetic energy, where k is Boltzmann's constant. There are n ions and, because each hydrogen ion brings one electron, also n electrons, so the total is 2n × 3/2 kT = 3nkT.
---
$ P_{\mathrm{loss}} = \frac{3 n k T}{\tau_E}
By the definition of the confinement time, the plasma loses its stored heat W over a time τ_E. So the power lost per cubic metre is W divided by τ_E. A longer confinement time means slower losses.
---
$ \tfrac{1}{4} n^2 \langle \sigma v \rangle E_\alpha \geq \frac{3 n k T}{\tau_E}
Ignition means the alpha heating at least matches the losses, so we set the first expression greater than or equal to the third. Everything in this inequality is either a property of the fuel or something the machine controls: density, temperature and confinement time.
---
$ n \tau_E \geq \frac{12\,k T}{\langle \sigma v \rangle E_\alpha}
Multiply both sides by τ_E and by 4, and divide by n and by ⟨σv⟩E_α, so that the quantities the machine controls collect on the left. One factor of n cancels, which is why the density appears only once. This is Lawson's condition in its original form: the product of density and confinement time must exceed a number that depends on temperature.
---
$ n T \tau_E \geq \frac{12\,k T^2}{\langle \sigma v \rangle E_\alpha}
Multiply both sides by T. This seems an odd thing to do, but it pays off: between about 10 and 20 keV, ⟨σv⟩ rises roughly in proportion to T², so the right-hand side becomes nearly constant across the whole range where a reactor would operate. The condition then becomes a single target for the product of three quantities.
---
$ n T \tau_E \gtrsim 3 \times 10^{21}\,\text{keV s/m}^3
Inserting measured D–T reaction rates gives a minimum value near a temperature of 15 keV, with the temperature measured in energy units (kT in keV). This number is the **triple product** that a reactor must reach to ignite.
:::

What the result says is that no single quantity is enough. A plasma can be extremely hot and still fail if it is thin or leaks quickly. What counts is the product of density, temperature and confinement time. Since the best temperature is fixed by the physics of D–T fusion, at around 150 million kelvin, the real choice is between density and time, and that choice has divided the field into two strategies.

**Magnetic confinement** holds a thin plasma, about 10²⁰ particles per cubic metre, a few millionths of the density of the air around us, and tries to keep it hot for seconds. Plugging in the numbers, n = 10²⁰ m⁻³ at 15 keV needs τ_E of about two seconds. **Inertial confinement** goes to the other extreme: it compresses a tiny pellet of fuel to enormous density for well under a billionth of a second. Nothing holds the fuel in place; it burns in the instant before its own inertia lets it fly apart, which is where the name comes from.

::: note
A fusion reactor must reach a triple product of density, temperature and confinement time of about 3 × 10²¹ keV·s/m³, at a temperature near 150 million kelvin. Beyond that point the charged alpha particles from fusion keep the plasma hot on their own, while the neutrons escape and carry the useful energy out.
:::

## Bottling a plasma with magnets

A charged particle in a magnetic field cannot move straight across the field lines; the magnetic force pushes it sideways, so it spirals around a line instead. Along the line, though, it moves freely. A straight tube of field would therefore hold a plasma in from the sides but let it pour out of the ends. The natural fix is to bend the tube round into a ring, a doughnut shape called a **torus**, so that the field lines close on themselves and there are no ends.

In the early 1950s, in secret programmes in Britain, the United States and the Soviet Union, physicists built many such machines. In Princeton Lyman Spitzer designed the **stellarator**, whose coils are twisted to stop particles drifting out of the ring. Nearly all the early machines disappointed. Plasmas proved to be restless: small ripples grew into instabilities that threw the plasma against the walls far faster than the simple theory predicted. In 1958 the British machine ZETA was announced to the press as having produced fusion neutrons, and within months the claim had to be withdrawn; the neutrons came from instabilities, not from a hot plasma. That same year the programmes were declassified, as each country realised that the others were equally stuck.

::: history 1968–1969 · Moscow and Culham
Soviet physicists at the Kurchatov Institute in Moscow, building on a 1951 idea of Igor Tamm and Andrei Sakharov, had developed a design they called the **tokamak**, a Russian acronym for "toroidal chamber with magnetic coils". Besides an external field, a tokamak drives a large electric current around the ring through the plasma itself, and that current's own magnetic field helps hold the plasma steady. In 1968 Lev Artsimovich, the head of the Soviet fusion programme, reported that the T-3 tokamak had reached electron temperatures of about 10 million kelvin and held its heat far better than any Western machine. Western physicists suspected a mistake in the measurement. So in 1969, at the height of the Cold War, a British team from the Culham laboratory flew to Moscow with a laser system that could measure the electron temperature directly, by scattering light off the electrons. It confirmed the Soviet claim. Within a few years laboratories around the world had switched to tokamaks.
:::

Progress since then has been steady but slow, because each step in performance needs a bigger machine. The Joint European Torus (JET) at Culham in England, the largest tokamak of its generation, produced 16 megawatts of fusion power in 1997, briefly about two-thirds of the heating power put in. In its final experiments in 2023, before it was shut down, it released 69 megajoules of fusion energy over about 5 seconds, the most energy from a single pulse so far.

The next step is **ITER**, under construction at Cadarache in southern France. It is a collaboration of the European Union, China, India, Japan, South Korea, Russia and the United States, formally agreed in 2006. ITER is designed to produce 500 megawatts of fusion power from 50 megawatts of heating, a **gain**, written Q, of 10, in pulses lasting several minutes. It will not generate electricity; its job is to show that a burning plasma, one heated mainly by its own alphas, can be controlled. The project has suffered long delays and large increases in cost. Its schedule, revised in 2024, plans the start of research operations in 2034 and deuterium–tritium operation in 2039.

## Ignition with lasers

::: history 5 December 2022 · Livermore, California
At the National Ignition Facility (NIF), 192 laser beams delivered 2.05 megajoules of ultraviolet light into a small gold cylinder holding a capsule of frozen deuterium and tritium about two millimetres across. The light heated the cylinder, which bathed the capsule in X-rays, and the capsule's outer layer blew off so violently that the fuel inside was driven inward, compressed and heated until it fused before it could fly apart. The capsule released 3.15 megajoules, the first time a controlled fusion experiment had produced more energy than was delivered to the fuel. NIF has repeated ignition several times since, with higher yields.
:::

The achievement needs careful reading. The gain counted only the laser light that reached the target. To make that 2.05 megajoules of light, the lasers drew about 300 megajoules from the electrical grid, because they convert electricity into light very inefficiently. NIF can fire only about once a day. A power plant on this principle would need targets ignited several times per second, lasers far more efficient than NIF's, and cheap targets made by the million.

Even when a plasma's gain is high enough, a working power plant must solve problems that the physics experiments mostly avoid. It must breed its own tritium in the blanket and recover it. Its walls must survive years of bombardment by 14 MeV neutrons, which knock atoms out of place in the metal and make it brittle and radioactive. It must turn the neutrons' heat into electricity, run continuously for months, and do all this at a cost that competes with other ways of making power. Fusion has been jokingly described as "thirty years away" for most of its history. What has changed is that the physics has now been demonstrated: plasmas have been confined and ignited. The remaining work is mostly engineering, and it is still substantial.

::: quiz
? Why must a fusion reactor confine its plasma for a minimum time, not just heat it?
- Fusion reactions take seconds to complete
+ The plasma loses energy continuously, and alpha heating must keep up with those losses
- The neutrons need time to breed tritium
= Heat leaks out of a plasma all the time, at a rate set by the confinement time τ_E. The Lawson criterion says the fusion heating must at least match that loss, which requires density, temperature and confinement time together to be large enough.
? Why can't the neutrons from D–T fusion help keep the plasma hot?
- They carry almost no energy
+ They have no charge, so the magnetic field does not hold them and they escape
- They are absorbed by deuterium
= Only the charged alpha particles are held in the plasma and share their 3.5 MeV with it. The 14.1 MeV neutrons leave at once and deposit their energy in the surrounding blanket, where it can be collected as heat.
? NIF's 2022 shot released 3.15 MJ from 2.05 MJ of laser light. Why was it not net energy for the facility?
- Most of the fusion energy was lost as neutrinos
+ The lasers needed about 300 MJ from the grid to produce that 2.05 MJ of light
- The capsule did not actually ignite
= The gain compared the fusion output only with the light delivered to the target. Because the lasers turn electricity into light very inefficiently, the facility as a whole used about a hundred times more energy than the capsule released.
:::
