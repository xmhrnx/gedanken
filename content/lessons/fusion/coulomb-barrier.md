To fuse, two nuclei must come within reach of the strong force, about a femtometre (10⁻¹⁵ m). But nuclei are positively charged, and as they approach, their electrical repulsion grows without limit. This repulsion is the **Coulomb barrier**, and estimating its height shows that, by classical physics, the Sun should not shine at all.

## How high is the barrier?

::: steps The Coulomb barrier for two protons
$ E_C = \frac{k e^2 Z_1 Z_2}{r}
The electrical potential energy of two charges Z₁e and Z₂e a distance r apart. An incoming nucleus needs at least this much kinetic energy to reach r.
---
$ k e^2 = 1.44\,\text{MeV fm}
A convenient form of Coulomb's constant times the proton charge squared, in nuclear units.
---
$ E_C \approx \frac{1.44\,\text{MeV fm} \times 1 \times 1}{1.4\,\text{fm}} \approx 1\,\text{MeV}
For two protons (Z₁ = Z₂ = 1) touching at about 1.4 fm. For heavier nuclei the barrier grows with the product Z₁Z₂.
:::

About 1 MeV, then. What energy do protons actually have in the Sun's core?

::: steps Thermal energy in the Sun's core
$ kT = 8.6 \times 10^{-5}\,\text{eV/K} \times 1.5 \times 10^7\,\text{K} \approx 1.3\,\text{keV}
The typical energy of a particle at temperature T is of order kT, where k is Boltzmann's constant. The Sun's core is about 15 million kelvin.
---
$ \frac{E_C}{kT} \approx \frac{1000\,\text{keV}}{1.3\,\text{keV}} \approx 800
The barrier is almost a thousand times higher than the typical thermal energy.
---
$ T \approx \frac{E_C}{k} \approx \frac{10^6\,\text{eV}}{8.6 \times 10^{-5}\,\text{eV/K}} \approx 10^{10}\,\text{K}
To make typical protons energetic enough to climb over, the core would need to be about ten billion kelvin, nearly a thousand times hotter than it is.
:::

Could the fastest protons make it? In a hot gas, particle energies follow the Maxwell–Boltzmann distribution, and the fraction with energy above E falls roughly as e^(−E/kT). At 800 kT that is about e⁻⁸⁰⁰, or 10⁻³⁴⁷. The Sun contains about 10⁵⁷ protons. Not a single one would ever have enough energy.

::: history 1926 · Cambridge, England
Eddington knew the difficulty. In *The Internal Constitution of the Stars* he held his ground: the evidence that stars ran on subatomic energy was strong, and if physics said the stars were not hot enough, so much the worse for current physics. "We do not argue with the critic who urges that the stars are not hot enough for this process; we tell him to go and find a hotter place."
:::

## Tunneling

The hotter place turned out to be unnecessary. Quantum mechanics describes a particle by a wave, and a wave meeting a barrier does not stop dead at its edge: it decays inside the barrier and, if the barrier is thin enough, part of it leaks through. A particle can therefore appear on the far side without ever having had enough energy to go over the top. This is **quantum tunneling**. The picture of a particle burrowing through is only a manner of speaking; there is no moment at which it can be found inside with negative kinetic energy.

::: history 1928 · Göttingen
George Gamow, a 24-year-old Soviet visitor, used tunneling to explain alpha decay. Uranium-238 emits alpha particles with about 4 MeV, though the barrier holding them in is several times higher. Gamow showed that the alpha particle tunnels out, and that small differences in energy change the tunneling probability enormously, which explained why alpha emitters have half-lives ranging from microseconds to billions of years. Ronald Gurney and Edward Condon found the same result independently. In 1929 Robert Atkinson and Fritz Houtermans turned the argument around: if alpha particles can tunnel out of nuclei, protons can tunnel in.
:::

Gamow's result for the probability that two nuclei with relative energy E tunnel through their barrier is approximately:

::: equation Tunneling probability
P(E) \approx \exp\left(-\sqrt{\frac{E_G}{E}}\right), \quad E_G = 2 m_r c^2 \left(\pi \alpha Z_1 Z_2\right)^2
:::

Here α ≈ 1/137 is the fine-structure constant, which measures the strength of electromagnetism, and m_r is the reduced mass, half the proton mass for two protons. For two protons the **Gamow energy** E_G is about 490 keV. At a typical thermal energy of 1.3 keV, P ≈ e^(−√380) ≈ e^(−19.5), about 3 × 10⁻⁹: tiny, but nowhere near 10⁻³⁴⁷.

Two effects now compete. Faster protons tunnel far more easily, but there are exponentially fewer of them. Their product peaks sharply at about 6 keV in the Sun, around five times kT, in a narrow range called the **Gamow window**. Nearly all solar fusion happens there, among protons a little faster than average, with tunneling probabilities around 10⁻⁴.

::: note
The Coulomb barrier between two protons is about 1 MeV, roughly 800 times the thermal energy in the Sun's core. Classically no fusion would occur. Tunneling lets protons with only a few keV penetrate the barrier, rarely but often enough.
:::

Tunneling explains why fusion happens at 15 million kelvin, but not why it is so slow. Most pairs of protons that tunnel together simply fly apart again; only a tiny fraction convert into a deuteron through the weak force. The higher barriers of heavier nuclei, rising with Z₁Z₂, are why the carbon cycle needs hotter stars, and why fusion on Earth uses hydrogen's heavier isotopes at temperatures around 100 million kelvin.

::: quiz
? Roughly how does the Sun's core temperature compare with what classical physics would require for proton fusion?
- About the same
- About 10 times too cool
+ About a thousand times too cool
= Climbing a 1 MeV barrier classically needs about 10¹⁰ K; the core is 1.5 × 10⁷ K.
? Why does most solar fusion happen at energies a few times kT rather than at kT?
+ Tunneling probability rises steeply with energy, outweighing the falling number of fast protons up to a point
- Protons at kT are too slow to collide
- The Coulomb barrier is lower for faster protons
= The product of the Boltzmann factor and the tunneling probability peaks at about 6 keV, the Gamow window.
? Why does fusing carbon with protons need higher temperatures than fusing two protons?
- Carbon nuclei are heavier and move more slowly
+ Carbon's six protons repel an incoming proton six times as strongly at a given distance
- Carbon cannot tunnel
= The barrier scales with Z₁Z₂, which is 6 for carbon plus a proton, and E_G grows as (Z₁Z₂)².
:::
