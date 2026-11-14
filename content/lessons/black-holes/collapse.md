Schwarzschild's solution says what spacetime would look like around a mass smaller than its Schwarzschild radius. It does not say any real object gets that small. For a generation, the leading physicists were confident none could. The question became: is there any pressure in nature that can hold up a dead star against its own gravity, however heavy it is?

## The resistance

Einstein published a paper in 1939 arguing that a cluster of orbiting particles could never be compressed to its Schwarzschild radius, since the particles would need to move faster than light. Arthur Eddington, the most influential astrophysicist of the era, called the idea of a star collapsing without limit absurd and said there should be a law of nature to prevent it. Their instinct was reasonable: infinities usually mean a theory has been pushed past its range, not that nature contains them.

## What holds up a dead star

A star like the Sun is supported by the pressure of hot gas, kept hot by nuclear fusion. When the fuel runs out, the core cools and contracts. Sirius B, the first known **white dwarf** (a stellar remnant about the Sun's mass packed into roughly Earth's size), showed that something else can take over.

That something is quantum mechanics. Electrons obey the **Pauli exclusion principle**: no two can occupy the same quantum state. Squeeze electrons into a small volume and they are forced into states of higher and higher momentum, even at zero temperature. The resulting outward push is **electron degeneracy pressure**, and it does not need heat. In 1926 Ralph Fowler showed it could hold up white dwarfs indefinitely.

## Chandrasekhar's limit

In 1930 Subrahmanyan Chandrasekhar, nineteen and on a ship from India to Cambridge, noticed a problem. In a heavy white dwarf the electrons are pushed to speeds near c, and relativistic electrons push back more weakly. A scaling argument shows why that is fatal.

Take a star of N nucleons (protons and neutrons, each of mass m_p, which supply nearly all the mass) and radius R, with about one electron per two nucleons. Track the energy of a typical particle; factors of order one are dropped throughout.

::: steps Why there is a maximum mass
$ n \sim \frac{N}{R^3}
The number density of electrons is the number divided by the volume.
---
$ p \sim \hbar n^{1/3} \sim \frac{\hbar N^{1/3}}{R}
Exclusion forces each electron into its own cell of size about n^(−1/3); the uncertainty principle then gives it momentum about ħ divided by that size.
---
$ E_{\text{electron}} \sim \frac{p^2}{m_e} \sim \frac{\hbar^2 N^{2/3}}{m_e R^2}
In a light white dwarf the electrons are slow, so kinetic energy is p²/2m. It grows like 1/R² as the star shrinks.
---
$ E_{\text{gravity}} \sim -\frac{G N m_p^2}{R}
The gravitational energy per nucleon is about −GMm_p/R with M ~ N m_p. It grows like 1/R.
---
$ \frac{1}{R^2} \text{ beats } \frac{1}{R} \text{ at small } R
Non-relativistically, shrinking always raises the electrons' energy faster than it lowers gravity's, so the star finds a stable size. That is a normal white dwarf.
---
$ E_{\text{electron}} \sim pc \sim \frac{\hbar c N^{1/3}}{R}
Add mass and the star shrinks until p approaches m_e c. Then the energy is pc, not p²/2m, and grows only like 1/R.
---
$ E_{\text{total}} \sim \frac{\hbar c N^{1/3} - G N m_p^2}{R}
Now both terms scale the same way with R, so changing the size cannot rebalance them. Which term wins depends only on N.
---
$ N_{\text{max}} \sim \left(\frac{\hbar c}{G m_p^2}\right)^{3/2} \approx 2 \times 10^{57}
If the gravity term is larger, the energy keeps falling as R shrinks and nothing stops collapse. The two terms are equal at this N.
---
$ M_{\text{max}} \sim N_{\text{max}} m_p \approx 4 \times 10^{30}\,\text{kg} \approx 2\text{ solar masses}
Multiply by the nucleon mass. A careful calculation with the real star's structure gives 1.4 solar masses.
:::

The striking feature is that the answer is made of constants of nature only: ħ, c, G and the proton mass. It does not depend on what the star used to be. Above about 1.4 solar masses, electron degeneracy cannot hold up a cold star at all.

::: history 1935 · Royal Astronomical Society, London
Chandrasekhar presented his full calculation in January 1935. Eddington, who had seen the paper in advance, rose afterward and ridiculed it, saying a star must have some way to avoid such "stellar buffoonery." Most of the audience sided with Eddington, and Chandrasekhar turned to other fields. He received the Nobel Prize in 1983.
:::

## Neutron stars, and a second limit

Above the limit, electrons are squeezed into protons, making neutrons. In 1934 Walter Baade and Fritz Zwicky proposed that a supernova leaves behind a **neutron star**: a ball of neutrons about 20 km across, supported by neutron degeneracy pressure and nuclear forces. The same scaling argument applies to neutrons, so there should be a ceiling here too. In 1939 J. Robert Oppenheimer and George Volkoff computed it, finding about 0.7 solar masses for an ideal neutron gas. With modern nuclear physics the limit is roughly 2 to 2.3 solar masses; the heaviest neutron stars measured are just over 2. Neutron stars were confirmed in 1967 as pulsars.

## Collapse with nothing to stop it

Also in 1939, Oppenheimer and Hartland Snyder asked what happens to a star above every limit. They modelled it as a uniform ball of pressureless matter and solved Einstein's equations exactly. The ball shrinks through its Schwarzschild radius and on to zero size in a finite time on its own clocks. A distant observer sees it slow, redden and fade at r_s, exactly as in the previous lesson. Their paper appeared on 1 September 1939, the day Germany invaded Poland, and was largely ignored.

The obvious objection was that a perfectly round, uniform star is a fiction; perhaps real lumpiness or rotation would make the matter miss the centre and bounce. In 1965 Roger Penrose closed that door. His **singularity theorem** says, in words: if a region forms from which even outgoing light is dragged inward (a "trapped surface"), and energy is positive as for all known matter, then some paths through spacetime must end after a finite time. No symmetry is assumed. Once a horizon forms, a singularity follows. Penrose received the Nobel Prize for this in 2020.

::: note
The logic is a chain of limits. Gas pressure fails when fusion stops; electron degeneracy fails above 1.4 solar masses because relativistic electrons can't outpace gravity; neutron degeneracy fails above about 2 solar masses. Above that, general relativity predicts nothing can stop collapse, and Penrose showed this does not depend on perfect symmetry.
:::

In 1967 John Wheeler, once a sceptic himself, began using the name **black hole** in a lecture, and it stuck.

::: quiz
? Why does a white dwarf become unstable once its electrons are relativistic?
- The electrons stop obeying the exclusion principle
+ Electron energy then scales as 1/R, the same as gravity, so shrinking no longer helps
- The star gets hot and fusion restarts
= With E ~ pc the pressure term and the gravity term both scale as 1/R, and if gravity wins at one radius it wins at all.
? What did Penrose's 1965 theorem add to Oppenheimer and Snyder's result?
+ Singularities form even without perfect spherical symmetry
- Collapse stops at the neutron star stage
- Black holes emit radiation
= Oppenheimer–Snyder assumed a perfectly uniform ball; Penrose needed only a trapped surface and positive energy.
:::
