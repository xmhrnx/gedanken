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

## What tunnelling does not explain

Tunnelling explains why fusion can happen at 15 million kelvin. It does not explain why it happens so slowly. Even after two protons have tunnelled close together, they almost always fly apart again a moment later, because a nucleus of two protons is not stable. Only if the weak force turns one of them into a neutron during that brief contact do they become a deuteron, and that is extremely rare. The barrier lets protons meet; the weak force decides how seldom they stay.

The barrier's dependence on charge also explains a pattern seen across astronomy and engineering. Because E_G grows as (Z₁Z₂)², and the tunnelling probability depends on it exponentially, each step up in nuclear charge makes fusion far harder. That is why the carbon cycle, in which protons must tunnel into carbon nuclei with six times their charge, takes over only in stars with hotter cores than the Sun's. It is also why fusion reactors on Earth use the heavy isotopes of hydrogen, deuterium and tritium, which have only a single unit of charge, and even so must run at around 100 million kelvin, because a reactor cannot afford to wait billions of years for rare events.

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
