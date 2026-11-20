The previous lesson ended with a loose thread. Fusing hydrogen into helium releases plenty of energy, but only if two nuclei can be brought close enough to stick together. Nuclei are positively charged, so they repel each other, and the closer they get the harder they push apart. This lesson estimates how strong that repulsion is, shows that by the rules of classical physics it is far too strong for the Sun to overcome, and then explains how quantum mechanics lets nuclei through anyway. The answer turns out to depend on a delicate balance between two competing effects, and it decides not only how stars burn but why fusion on Earth is so hard.

## What stands in the way?

Inside a nucleus, protons and neutrons are held together by the **strong nuclear force**, which is far more powerful than electrical repulsion but acts only over a very short distance, about a femtometre (1 fm is 10⁻¹⁵ m, roughly the width of a proton). Beyond a few femtometres it is effectively zero. So two nuclei feel nothing but their electrical repulsion until they are almost touching, and then, suddenly, the strong force grabs them.

Picture the energy of the pair as a landscape. Far apart, the energy is zero. As they approach, the electrical repulsion makes the energy climb, steeply at the end. At about the point of contact the strong force takes over and the energy plunges into a deep valley, which is the bound nucleus. To fall into the valley, an incoming nucleus first has to climb the hill. That hill is called the **Coulomb barrier**, after Charles-Augustin de Coulomb, who measured the electrical force in the 1780s.

## How high is the hill?

The height of the barrier is just the electrical potential energy of the two nuclei at the distance where the strong force takes over.

::: steps The Coulomb barrier for two protons
$ E_C = \frac{k e^2 Z_1 Z_2}{r}
Two charges repel with a potential energy that grows as they approach, in proportion to the product of the charges and inversely to their separation r. Here the charges are Z₁e and Z₂e, where e is the charge of one proton and Z is the number of protons in each nucleus, and k is Coulomb's constant. A nucleus coming in from far away must have at least this much kinetic energy to reach distance r, because it gives up kinetic energy as it climbs.
---
$ k e^2 = 1.44\,\text{MeV fm}
This is a convenient value of ke², Coulomb's constant times the proton charge squared, expressed in the units nuclear physicists use: energies in MeV (mega-electronvolts) and distances in femtometres. Nothing physical has changed; the numbers simply become easy to handle at nuclear scales.
---
$ E_C \approx \frac{1.44\,\text{MeV fm} \times 1 \times 1}{1.4\,\text{fm}} \approx 1\,\text{MeV}
For two protons, Z₁ = Z₂ = 1, and the strong force takes hold when their centres are about 1.4 fm apart. The barrier comes out at about 1 MeV. For heavier nuclei it grows with the product Z₁Z₂, so carbon and a proton face a barrier several times higher.
:::

To decide whether 1 MeV is a lot, compare it with the energy that protons actually have in the Sun's core.

::: steps Thermal energy in the Sun's core
$ kT = 8.6 \times 10^{-5}\,\text{eV/K} \times 1.5 \times 10^7\,\text{K} \approx 1.3\,\text{keV}
In a gas at temperature T, the particles have typical kinetic energies of order kT, where k is Boltzmann's constant, the conversion factor between temperature and energy. The Sun's core is about 15 million kelvin, which gives typical proton energies of about 1.3 keV (thousand electronvolts).
---
$ \frac{E_C}{kT} \approx \frac{1000\,\text{keV}}{1.3\,\text{keV}} \approx 800
Dividing the barrier height by the typical energy shows how far short the protons fall. The hill is roughly 800 times higher than a typical proton can climb.
---
$ T \approx \frac{E_C}{k} \approx \frac{10^6\,\text{eV}}{8.6 \times 10^{-5}\,\text{eV/K}} \approx 10^{10}\,\text{K}
Turn the question around and ask how hot the core would need to be for typical protons to have 1 MeV. Dividing the barrier energy by Boltzmann's constant gives about ten billion kelvin, nearly a thousand times hotter than the Sun's centre.
:::

One way out might seem to be the fastest protons. A gas is not uniform: its particles share energy unevenly through collisions, and a few always have much more than average. The spread follows the **Maxwell–Boltzmann distribution**, in which the fraction of particles with energy above E falls roughly as e^(−E/kT). That falls off very fast. At E = 800 kT the fraction is about e⁻⁸⁰⁰, which is 10⁻³⁴⁷. The whole Sun contains only about 10⁵⁷ protons. Multiply the two and the expected number of protons energetic enough to climb the barrier is 10⁻²⁹⁰, which means none, not now and not ever. By classical physics, the Sun should not shine.

::: history 1926 · Cambridge, England
Eddington knew all this. In his book *The Internal Constitution of the Stars* he weighed the evidence that stars run on subatomic energy against the physicists' objection that they were not hot enough, and he decided the evidence was stronger. If physics said the stars could not do it, then physics was incomplete. "We do not argue with the critic who urges that the stars are not hot enough for this process," he wrote; "we tell him to go and find a hotter place." It was a bold position for a theorist to take, and within a few years it was vindicated by a theory he had not foreseen.
:::

## How can a particle cross a hill it cannot climb?

The hotter place turned out to be unnecessary. In quantum mechanics a particle is described by a wave, and the strength of the wave at any point gives the probability of finding the particle there. A classical ball rolling toward a hill it cannot climb simply stops and rolls back. A quantum wave meeting a barrier behaves differently: it does not stop dead at the edge but extends into the barrier, shrinking rapidly as it goes. If the barrier is thin enough, a small part of the wave is still left on the far side, and that means there is a small probability of finding the particle there. It has crossed without ever having enough energy to go over the top. This is **quantum tunnelling**.

The word suggests a particle burrowing through the hill, and that picture should not be taken literally. There is no moment at which the particle can be caught partway through with negative kinetic energy. The wave description says only that a particle arriving at one side has a calculable chance of being found on the other.

::: history 1928 · Göttingen
George Gamow, a 24-year-old physicist visiting from Leningrad, was puzzling over alpha decay, in which a heavy nucleus such as uranium-238 spits out an alpha particle (a helium nucleus). The alpha particle emerges with about 4 MeV, yet the barrier that should hold it inside the nucleus is several times higher. Classically it could never get out. Gamow saw that the new wave mechanics answered this: the alpha particle tunnels out. His calculation explained something else that had baffled everyone. Alpha emitters have half-lives ranging from fractions of a microsecond to billions of years, yet their alpha energies differ only by a factor of two or so. Tunnelling probability depends exponentially on energy, so small differences in energy produce enormous differences in decay rate. Ronald Gurney and Edward Condon at Princeton reached the same result independently. A year later, in 1929, Robert Atkinson and Fritz Houtermans turned the argument around: if alpha particles can tunnel out of nuclei, protons in a star can tunnel in.
:::

## How likely is tunnelling?

Gamow worked out the probability that two nuclei approaching each other with relative energy E tunnel through their Coulomb barrier. A good approximation is:

::: equation Tunnelling probability
P(E) \approx \exp\left(-\sqrt{\frac{E_G}{E}}\right), \quad E_G = 2 m_r c^2 \left(\pi \alpha Z_1 Z_2\right)^2
:::

The quantity E_G is called the **Gamow energy**, and it packs together everything about the pair that makes tunnelling hard. It contains α ≈ 1/137, the **fine-structure constant**, a pure number that measures the strength of electromagnetism; the charges Z₁ and Z₂; and m_r, the **reduced mass**, which is how the mass of the pair enters the motion of one relative to the other (for two equal masses it is half of either). Heavier particles have shorter quantum wavelengths, so their waves die off faster inside the barrier, and stronger charges make the barrier thicker. Both raise E_G.

For two protons, m_r c² is half of 938 MeV, and the Gamow energy comes out at about 490 keV. At a typical thermal energy of 1.3 keV the ratio E_G/E is about 380, its square root about 19.5, and the probability about e^(−19.5), or 3 × 10⁻⁹. That is small, but compare it with the classical figure of 10⁻³⁴⁷. Tunnelling has turned an impossibility into something that merely happens rarely.

## Where in the crowd does fusion happen?

Now two effects pull in opposite directions. A faster proton tunnels much more easily, because the square root in the exponent shrinks as E grows. But faster protons are exponentially rarer, because of the Maxwell–Boltzmann factor. The rate at which fusion happens at a given energy is the product of the two: how many protons have that energy, times how likely each is to get through.

The first factor falls with energy and the second rises, so their product has a peak. For two protons in the Sun's core it lies at about 6 keV, roughly five times the typical thermal energy, and it is narrow. This range is called the **Gamow window**. Almost all solar fusion happens there, among protons somewhat faster than average but not among the extremely rare very fast ones. At 6 keV the tunnelling probability is about 10⁻⁴: one approach in ten thousand gets through.

::: note
The Coulomb barrier between two protons is about 1 MeV, some 800 times the typical thermal energy in the Sun's core, so classically no fusion would ever happen. Quantum tunnelling lets protons with only a few keV pass through the barrier. Fusion happens mainly in the narrow Gamow window, where the falling number of fast protons and their rising tunnelling probability balance.
:::

## Where does the Gamow energy come from?

The formula for P(E) was stated above, not derived, and it is worth seeing where it comes from, because the derivation shows what the exponent really measures. The starting point is the rule from the lesson on tunnelling. Inside a barrier of height V, a particle of energy E has a wave that does not oscillate but dies away, falling by a factor e for every distance 1/κ, where κ = √(2m(V − E))/ħ. A taller barrier, or a heavier particle, makes κ larger and the decay faster. When the barrier's height changes from place to place, as the Coulomb hill does, the wave decays a little in each thin slice, at the rate set by the local height, and the total effect is found by adding up the decay across all the slices.

::: steps The tunnelling exponent for a Coulomb barrier
$ P \approx \exp\left(-\frac{2}{\hbar}\int_{R}^{r_c} \sqrt{2m_r\left(V(r)-E\right)}\,\mathrm{d}r\right)
The probability is the square of the wave's amplitude, which is why the factor 2 appears in front: the amplitude falls by e^(−∫κ dr) across the barrier, and the probability by the square of that. The mass is the reduced mass m_r, because what tunnels is the separation between the two nuclei. The integral runs from R, where the strong force takes over, out to r_c, the point where an incoming nucleus would classically stop and turn back.
---
$ V(r) = \frac{Z_1 Z_2 k e^2}{r}, \quad r_c = \frac{Z_1 Z_2 k e^2}{E}
The barrier is the electrical potential energy, and the turning point is where it equals the incoming energy. For two protons at 6 keV, r_c is 1.44 MeV fm divided by 0.006 MeV, or 240 fm, while R is only about 1.4 fm. The barrier is therefore about two hundred times wider than a proton. Most of the decay happens across that long, low outer stretch, not near the peak, so very little is lost by letting the integral start at zero instead of R.
---
$ \int_0^{r_c} \sqrt{\frac{Z_1 Z_2 k e^2}{r} - E}\,\mathrm{d}r = \frac{\pi}{2}\,\frac{Z_1 Z_2 k e^2}{\sqrt{E}}
This integral can be done exactly by changing variable to an angle (r = r_c sin²θ), and the answer is surprisingly tidy. It grows in proportion to the product of the charges, because a stronger repulsion makes the barrier both taller and wider, and it falls as 1/√E, because a faster nucleus starts its tunnelling closer in and has less barrier left to cross.
---
$ P \approx \exp\left(-2\pi Z_1 Z_2 \frac{\alpha c}{v}\right)
Put the integral back into the exponent and use two substitutions: ke² = αħc, which is the definition of the fine-structure constant, and E = ½m_r v², where v is the speed at which the nuclei approach each other. Planck's constant cancels out of the exponent, and what is left is strikingly simple. Tunnelling depends on how fast the nuclei approach compared with αc, which is c/137, or about 2,200 km/s.
---
$ 2\pi Z_1 Z_2 \frac{\alpha c}{v} = \sqrt{\frac{2 m_r c^2 \left(\pi \alpha Z_1 Z_2\right)^2}{E}} = \sqrt{\frac{E_G}{E}}
Writing v back in terms of E and gathering the constants under one square root gives exactly the Gamow energy defined earlier. For two protons at 6 keV, v is about 0.5% of the speed of light, and 2παc/v is about 9, the same exponent found before from E_G.
:::

The form 2πZ₁Z₂αc/v makes the physics easy to read. Nuclei that meet slowly, compared with a speed of a couple of thousand kilometres per second, have almost no chance of getting through; nuclei that meet fast pass easily. Charge makes things worse in direct proportion inside the exponent, which is why the effect of adding protons is so brutal.

The derivation also shows why fusion rates are so sensitive to temperature. Combining the tunnelling factor with the Maxwell–Boltzmann factor and finding where their product peaks gives the centre of the Gamow window:

::: equation The centre of the Gamow window
E_0 = \left(\frac{E_G\,(kT)^2}{4}\right)^{1/3}
:::

For two protons in the Sun, with E_G = 490 keV and kT = 1.3 keV, this gives about 6 keV, as quoted above. At the peak the two exponents together come to 3E₀/kT, about 14 for two protons. A small rise in temperature lowers that combined exponent, and because the rate depends on it exponentially, the rate climbs far faster than the temperature. Near the Sun's central temperature the proton–proton rate goes up roughly as the fourth power of T: a core 10% hotter fuses about 50% faster. That steepness acts as a thermostat. If the core heats a little, the burning speeds up, the extra pressure makes the core expand and cool, and the rate falls back.

## Two worked cases

The same formulas handle any pair of nuclei, and two contrasting examples show their reach.

First, a proton meeting a carbon-12 nucleus in the Sun's core, the first step of the CNO cycle. The reduced mass is 12/13 of a proton mass, about 860 MeV, and Z₁Z₂ = 6. The Gamow energy is 2 × 860 MeV × (π × 6/137)², about 33 MeV, nearly seventy times the proton–proton value. At kT = 1.3 keV the Gamow window centres on about 24 keV, where the tunnelling probability is e^(−37), about 10⁻¹⁶. Including the scarcity of protons that fast, the combined exponent at the peak is about 55, against 14 for two protons. Per encounter, then, carbon is harder to reach by a factor of about e⁴¹, roughly 10¹⁸. That the CNO cycle supplies even 1% of the Sun's power is due to the proton–proton chain's other handicap, the weak-force step, which carbon does not need during the collision. And because its combined exponent is four times larger, the CNO rate rises much more steeply with temperature, roughly as T¹⁸ near the Sun's centre, which is why it wins in only slightly hotter stars.

Second, deuterium meeting tritium in a fusion reactor at kT = 10 keV, about 116 million kelvin. The reduced mass is about 1.2 atomic mass units, or 1,125 MeV, and Z₁Z₂ = 1, so E_G is about 1.18 MeV. The Gamow window centres on about 31 keV, where the tunnelling probability is about e^(−6.2), or 2 × 10⁻³: one approach in five hundred gets through, compared with one in ten thousand for protons in the Sun. More important, deuterium and tritium need no weak-force conversion once they touch. The strong force rearranges them directly into helium-4 and a neutron, helped by an unstable state of helium-5 at about 64 keV that makes them especially eager to combine. Together these make the deuterium–tritium reaction roughly 10²⁴ times faster, pair for pair, than proton–proton fusion at comparable conditions, which is the only reason a reactor on Earth is thinkable at all.

## Common misconceptions

The most common mistake is to think that fusion needs a plasma hot enough for nuclei to go over the barrier. Neither stars nor reactors come close. The Sun's protons have about a thousandth of the barrier energy, and a reactor's deuterium and tritium, at 10 keV, have a few percent of their barrier of a few hundred keV. Every fusion reaction in the universe at ordinary stellar temperatures happens by tunnelling.

A second picture, popular in books, says that a tunnelling particle "borrows" the energy it lacks for a brief moment, as the uncertainty principle allows, and pays it back on the far side. It sounds plausible, but it is not the theory and it does not give the right numbers. The barrier here is hundreds of femtometres wide, and a proton could not cross it in the time such a loan would allow. The real statement is the one used above: the wave decays inside the barrier at a calculable rate, and whatever remains on the far side is the chance of finding the particle there.

A third mistake is to think that hotter is always better. The fusion rate does rise steeply with temperature, but not without end. Once the Gamow window has moved up past the barrier's most troublesome region, extra energy helps less, and for deuterium–tritium the reaction rate peaks at a temperature of roughly 65 keV, about 750 million kelvin, and then slowly falls. Meanwhile a hotter plasma loses energy faster as radiation and costs more to confine. Reactors therefore aim at 10–20 keV, where the balance between fusion gained and energy lost is most favourable.

Finally, it is tempting to treat the barrier height as the measure of difficulty. Two protons and two deuterons face exactly the same barrier height, since each nucleus has one unit of charge. But the deuterons are twice as heavy, so their waves decay faster and their Gamow energy is twice as large; by tunnelling alone they fuse less readily. In practice deuterium fuses vastly more easily than ordinary hydrogen, because it needs no weak-force step. The barrier decides how often nuclei touch; what happens when they touch is a separate question.

## What tunnelling does not explain, and where the formula fails

Tunnelling explains why fusion can happen at 15 million kelvin. It does not explain why it happens so slowly. Even after two protons have tunnelled close together, they almost always fly apart again a moment later, because a nucleus of two protons is not stable. Only if the weak force turns one of them into a neutron during that brief contact do they become a deuteron, and that is extremely rare. The barrier lets protons meet; the weak force decides how seldom they stay.

The barrier's dependence on charge also explains a pattern seen across astronomy and engineering. Because E_G grows as (Z₁Z₂)², and the tunnelling probability depends on it exponentially, each step up in nuclear charge makes fusion far harder. That is why the carbon cycle, in which protons must tunnel into carbon nuclei with six times their charge, takes over only in stars with hotter cores than the Sun's. It is also why fusion reactors on Earth use the heavy isotopes of hydrogen, deuterium and tritium, which have only a single unit of charge, and even so must run at around 100 million kelvin, because a reactor cannot afford to wait billions of years for rare events.

That separation points to the main limit of Gamow's formula. It describes only the electrical part of the problem. The full fusion rate is the tunnelling probability multiplied by a factor that measures how readily the nuclei react once they are in contact, which depends on the details of nuclear structure. Physicists call it the **astrophysical S-factor**, and it can differ between reactions by more than twenty powers of ten: tiny for proton–proton, because of the weak force, and large for deuterium–tritium, because of the helium-5 state. It has to be measured. At stellar energies the rates are so small that laboratory measurements are usually made at higher energies and extrapolated down using the tunnelling formula. Since 1991 the LUNA experiment, deep under the Gran Sasso mountain in Italy, where the rock blocks the cosmic rays that would swamp the signal, has measured several solar reactions inside or near their Gamow windows, including the fusion of two helium-3 nuclei.

The formula also treats the nuclei as bare charges in empty space. In a star the surrounding electrons gather slightly around each nucleus and partly hide its charge, which lowers the barrier a little. In the Sun this **electron screening** raises the proton–proton rate by a few percent. At the enormous densities inside white dwarfs and the crusts of neutron stars, the picture changes completely. There nuclei are packed into a crystal, and they can fuse even near absolute zero, because their quantum jiggling about fixed positions brings neighbours close enough to tunnel. This is called **pycnonuclear** fusion, from the Greek for "dense", and temperature plays almost no part in it.

There is one more way to shrink the barrier that has actually been demonstrated. A **muon** is a heavier cousin of the electron, about 207 times as massive. If a muon replaces the electron in a molecule of deuterium and tritium, the molecule shrinks by about the same factor, holding the nuclei so close that they fuse within about a trillionth of a second, with no heat at all. This **muon-catalysed fusion** was seen in 1956 by Luis Alvarez's group at Berkeley. It cannot produce power, though. A muon lives only about 2.2 microseconds, and it sometimes sticks to the helium nucleus it has just helped to make, so each muon catalyses at most about a hundred to a hundred and fifty fusions before it is lost, which releases less energy than making the muon costs.

## Where it shows up

The Coulomb barrier and tunnelling through it run through several lessons in this course. The same calculation, run in reverse, is Gamow's explanation of alpha decay from the lesson on Radioactivity: an alpha particle escaping a nucleus tunnels out through the same kind of barrier, and the exponential sensitivity to energy is why the half-lives in Half-Life range from microseconds to billions of years. The general mechanism is the one in Tunneling, where it also explains the scanning tunnelling microscope and the flash memory in phones, both of which depend on electrons crossing thin barriers.

In this chapter, the barrier is why How Stars Shine needs temperatures of millions of kelvin and why heavier stars switch to the CNO cycle. It is why every design for Fusion Power uses deuterium and tritium and runs at around 100 million kelvin, and why the fuel in Thermonuclear Weapons must be heated by a fission explosion before it will burn. Further on, it explains a fact about the early universe in The Big Bang: as the universe cooled, nuclei with more charge could no longer tunnel into one another fast enough, which is part of the reason the first few minutes made hydrogen and helium but almost nothing heavier. In older stars, as later stages of burning fuse helium into carbon and carbon into heavier elements, each step needs a hotter core because each step faces a higher barrier.

::: quiz
? Roughly how does the Sun's core temperature compare with what classical physics would require for proton fusion?
- About the same
- About 10 times too cool
+ About a thousand times too cool
= Climbing a 1 MeV barrier classically needs typical energies corresponding to about 10¹⁰ K. The core is about 1.5 × 10⁷ K, nearly a thousand times cooler.
? Why does most solar fusion happen at energies a few times kT rather than at kT?
+ Tunnelling probability rises steeply with energy, outweighing the falling number of fast protons up to a point
- Protons at kT are too slow to collide
- The Coulomb barrier is lower for faster protons
= Protons near kT are plentiful but almost never get through the barrier, while very fast protons get through easily but are vanishingly rare. The product of the two peaks at about 6 keV, in the Gamow window.
? Why does fusing carbon with protons need higher temperatures than fusing two protons?
- Carbon nuclei are heavier and move more slowly
+ Carbon's six protons repel an incoming proton six times as strongly at a given distance
- Carbon cannot tunnel
= The barrier height scales with Z₁Z₂, which is 6 for carbon and a proton, and the Gamow energy grows as (Z₁Z₂)². Since the tunnelling probability depends on it exponentially, much higher energies, and so temperatures, are needed.
:::
