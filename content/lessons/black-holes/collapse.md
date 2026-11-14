Schwarzschild's solution describes the spacetime around a mass that is smaller than its own Schwarzschild radius, and the previous lesson showed that such a spacetime would contain an event horizon. But the solution does not say that any real object ever gets that small. The Sun would have to be crushed to 6 km across, the Earth to the size of a marble, and for a generation the leading physicists were confident that nothing in nature could do it. The question that decided the matter was not about relativity at all, but about matter: when a star dies, is there any pressure in nature that can hold it up against its own gravity, however heavy it is? This lesson follows the answer through a chain of limits, each discovered by someone who was not believed at first, until nothing is left to stop the collapse.

## Why were physicists so sure it could not happen?

The resistance was not mere stubbornness. Infinities in a physical theory usually mean that the theory has been pushed beyond the range where it applies, not that nature really contains them. Newton's gravity predicts infinite force between two point masses that touch, but nobody thinks real masses do that; something else, the solidity of matter, intervenes first. It was natural to expect that something would likewise intervene before any star reached its Schwarzschild radius.

Two of the most authoritative voices said so explicitly. Arthur Eddington, the most influential astrophysicist of the 1920s and 1930s, declared that a star collapsing without limit was absurd and that there should be a law of nature to prevent it. And in 1939 Einstein himself published a paper studying a cluster of particles orbiting one another under their own gravity. He showed that such a cluster could never be compressed to its Schwarzschild radius, because the particles would need to orbit faster than light to keep from falling inward, and he concluded that the Schwarzschild singularity does not occur in physical reality. His calculation was correct, but it assumed the particles keep orbiting in balance. It did not consider what happens if they simply fall.

## What holds up a star once its fuel runs out?

A star like the Sun is held up by the pressure of hot gas. Its core is kept at millions of degrees by nuclear fusion, and the gas pushes outward against gravity's inward pull. When the fuel runs out, fusion stops, the core cools, the gas pressure drops, and gravity squeezes the core smaller. If heat were the only source of pressure, the squeezing would never stop.

The first sign that something else can take over came from the faint companion of Sirius, the brightest star in the night sky. By 1915 astronomers knew that this companion, Sirius B, had roughly the mass of the Sun, and measurements of its colour showed that it was hot. A hot star that is so faint must be very small, and Sirius B turned out to be about the size of the Earth. A teaspoon of its matter would weigh several tonnes. Objects like it are called **white dwarfs**: stellar remnants about as massive as the Sun but packed into roughly the volume of the Earth. They are cooling, with no fuel left to burn, yet they do not collapse. Something other than heat holds them up.

That something is quantum mechanics. Electrons obey the **Pauli exclusion principle**: no two electrons can occupy the same quantum state. In ordinary matter this hardly matters, because there are far more available states than electrons. But squeeze electrons into a very small volume, and the low-momentum states fill up, so each additional electron is forced into a state of higher momentum than the last. Even at zero temperature, the electrons are moving fast, and fast-moving particles push on their surroundings. This outward push is called **electron degeneracy pressure**, and it needs no heat at all. In 1926 the British physicist Ralph Fowler showed that it could hold up a white dwarf indefinitely, however cold the star became. The problem seemed solved.

## Is there a limit to what electrons can hold up?

In 1930 Subrahmanyan Chandrasekhar, nineteen years old and on a ship from India to begin research at Cambridge, worked through Fowler's theory and noticed something Fowler had left out. In a heavier white dwarf, gravity squeezes harder, so the electrons are packed more tightly and pushed to higher momenta. In a heavy enough star, they approach the speed of light. And special relativity changes how a particle's energy depends on its momentum once it moves that fast. A scaling argument, which tracks how each kind of energy changes as the star shrinks and ignores numerical factors of order one, shows why that change is fatal.

Consider a star made of N nucleons, protons and neutrons, which supply nearly all its mass, each with mass m_p. There is about one electron for every two nucleons, a factor we drop along with the others. The star has radius R. We ask how the energy of a typical particle changes if R changes. If shrinking the star raises the energy, the star resists shrinking and can find a stable size; if shrinking lowers the energy, it will collapse.

::: steps Why there is a maximum mass
$ n \sim \frac{N}{R^3}
Start with the number of electrons per unit volume, n. It is the number of electrons, which is proportional to N, divided by the star's volume, which is proportional to R³. The symbol ∼ means "roughly equal to, ignoring numerical factors".
---
$ p \sim \hbar\, n^{1/3} \sim \frac{\hbar N^{1/3}}{R}
The exclusion principle forces each electron into its own small region of space, of size about n^(−1/3), the typical spacing between electrons. The uncertainty principle says that a particle confined to a region of size Δx must have momentum of at least about ħ/Δx, where ħ is the reduced Planck constant. So the typical electron momentum is ħ divided by the spacing, and it grows as the star shrinks.
---
$ E_{\text{electron}} \sim \frac{p^2}{m_e} \sim \frac{\hbar^2 N^{2/3}}{m_e R^2}
In a light white dwarf the electrons move well below the speed of light, so their kinetic energy is the familiar p²/2m, with m_e the electron mass. Substituting the momentum from the previous step shows that this energy grows as 1/R²: halving the star's radius quadruples it.
---
$ E_{\text{gravity}} \sim -\frac{G N m_p^2}{R}
Now the other side of the balance. Each nucleon sits in the gravitational field of a star of mass M ≈ Nm_p, at a typical distance R from the rest, so its gravitational potential energy is about −GMm_p/R. This energy becomes more negative as the star shrinks, but only as 1/R.
---
$ \frac{1}{R^2} \text{ beats } \frac{1}{R} \text{ at small } R
Compare the two. As R gets smaller, 1/R² grows faster than 1/R, so the rising electron energy always overtakes the falling gravitational energy eventually. The star shrinks until the two are balanced, then stops. That stable configuration is an ordinary white dwarf, which is why Fowler's theory worked.
---
$ E_{\text{electron}} \sim pc \sim \frac{\hbar c N^{1/3}}{R}
Now add more mass. A heavier star squeezes harder and settles at a smaller radius, so the electron momentum climbs until it approaches m_e c. At that point the electrons are relativistic, and special relativity says their energy is no longer p²/2m but approximately pc, momentum times the speed of light. With the same momentum as before, this energy grows only as 1/R.
---
$ E_{\text{total}} \sim \frac{\hbar c N^{1/3} - G N m_p^2}{R}
Add the electron and gravitational energies. Both now go as 1/R, so they can be written over a common denominator. This is the crucial change: changing the size of the star scales both terms by the same factor, so shrinking can no longer tip the balance toward the electrons. Which term wins depends only on the numerator, and so only on the number of nucleons N.
---
$ \hbar c\, N^{1/3} = G N m_p^2 \;\Rightarrow\; N_{\text{max}} \sim \left(\frac{\hbar c}{G m_p^2}\right)^{3/2}
The numerator changes sign when the two terms are equal. For smaller N the electron term wins, the total energy is positive, and the star can expand a little until its electrons are no longer fully relativistic and a stable size is found. For larger N gravity wins, the total energy is negative and keeps falling as R shrinks, and nothing stops the collapse. Dividing both sides by N^(1/3) and solving gives the dividing value.
---
$ \frac{\hbar c}{G m_p^2} \approx 1.7 \times 10^{38} \;\Rightarrow\; N_{\text{max}} \approx 2 \times 10^{57}
Now evaluate. The ratio ħc/(Gm_p²) compares the strength of quantum effects to the strength of gravity between two protons, and because gravity is so weak, it is enormous. Raising it to the power 3/2 gives the critical number of nucleons.
---
$ M_{\text{max}} \sim N_{\text{max}}\, m_p \approx 4 \times 10^{30}\,\text{kg} \approx 2\,M_\text{Sun}
Multiply by the mass of a nucleon, about 1.67 × 10⁻²⁷ kg, to get the maximum mass, about twice the mass of the Sun. A careful calculation that follows how density and pressure vary through a real star, and that counts the electrons correctly, lowers this to about 1.4 solar masses.
:::

The striking feature of this result is what it is made of: only constants of nature, the reduced Planck constant ħ, the speed of light c, the gravitational constant G and the proton mass. It does not depend on what the star was made of, how it formed or how hot it once was. Above about 1.4 solar masses, now called the **Chandrasekhar limit**, electron degeneracy pressure cannot hold up a cold star at all. So the theory that had seemed to rescue dead stars turned out to have a ceiling built into it.

::: history 1935 · Royal Astronomical Society, London
Chandrasekhar spent the next few years working out the full theory, and in January 1935 he presented it at a meeting of the Royal Astronomical Society in London. Eddington, who had seen the work beforehand and had been friendly and encouraging, was scheduled to speak immediately afterward. To Chandrasekhar's astonishment, Eddington used his time to ridicule the result. He did not dispute the mathematics; he argued that combining relativity with quantum mechanics in this way must be wrong, because the conclusion was intolerable, and that a star must have some way to avoid what he called "stellar buffoonery." Eddington's prestige was enormous, and most of the audience sided with him. Several leading physicists privately thought Chandrasekhar was right but did not say so publicly. Chandrasekhar, then twenty-four, eventually turned to other problems in astrophysics, and moved to the University of Chicago. He was proved right, and nearly half a century later, in 1983, he received the Nobel Prize in Physics.
:::

## What happens above the limit?

Above the Chandrasekhar limit, gravity squeezes the electrons until they are forced into the atomic nuclei, where they combine with protons to make neutrons. In 1934, only two years after the neutron was discovered, the astronomers Walter Baade and Fritz Zwicky proposed that this is what happens in a **supernova**, the explosion of a massive star: the core collapses, and what remains is a **neutron star**, a ball made almost entirely of neutrons, with the mass of the Sun or more in a sphere about 20 km across. Neutrons are also subject to the exclusion principle, so a neutron star is held up by neutron degeneracy pressure, helped by the strong nuclear forces between neutrons when they are packed this tightly.

The same scaling argument applies to neutrons, with one change: the particles providing the pressure are now the same particles providing the mass. So there should be a ceiling here too. In 1939 J. Robert Oppenheimer and his student George Volkoff calculated it using general relativity, finding about 0.7 solar masses for an idealised gas of neutrons that ignored nuclear forces. With modern nuclear physics the limit is thought to be roughly 2 to 2.3 solar masses, though the exact value is still uncertain because matter at such densities cannot be studied in a laboratory. The heaviest neutron stars measured so far are a little over 2 solar masses, consistent with that range. Neutron stars stopped being speculation in 1967, when Jocelyn Bell Burnell and Antony Hewish discovered **pulsars**, regularly flashing radio sources that were soon identified as rapidly spinning neutron stars.

## What if nothing stops the collapse?

Also in 1939, Oppenheimer and another student, Hartland Snyder, asked the obvious next question: what happens to a star that is above every limit, so that no pressure can halt it? To make the problem solvable they idealised the star as a perfectly round, uniform ball of matter with no pressure at all, and solved Einstein's equations for it exactly.

Their answer was the first description of a black hole forming. Outside the ball, the spacetime is exactly Schwarzschild's at every moment, as Birkhoff's theorem requires. The ball shrinks through its Schwarzschild radius and on to zero size in a finite time as measured by clocks riding on its surface; for a star of a few solar masses, the final plunge takes well under a millisecond of the star's own time. A distant observer, however, sees the surface slow down, redden and fade as it approaches r_s, exactly as the previous lesson described for a falling probe. Light from the surface takes longer and longer to escape, and the star effectively disappears from view, leaving behind the empty Schwarzschild geometry with a horizon. Their paper appeared in the journal Physical Review on 1 September 1939, the day Germany invaded Poland. Within a few years Oppenheimer was running the Manhattan Project, and the work was largely ignored for two decades.

The natural objection to Oppenheimer and Snyder was that a perfectly round, perfectly uniform star is a fiction. Real stars are lumpy and they rotate. Perhaps, many thought, the infalling matter would miss the exact centre, swirl past itself and bounce back out, so that no singularity would form. Soviet physicists in the early 1960s even published arguments that this was what generically happened.

In 1965 Roger Penrose closed that door. He introduced the idea of a **trapped surface**: a closed surface, such as a sphere, from which even light emitted outward is dragged inward, so that the surface's area shrinks no matter which way the light goes. Inside a Schwarzschild horizon every sphere is a trapped surface. Penrose then proved his **singularity theorem**, which in words says: if a trapped surface forms, and energy is never negative, as is true for all known forms of matter, then some paths through spacetime must come to an end after a finite time. They cannot be continued; spacetime runs out. The theorem assumes no symmetry at all. Lumps and rotation do not help. Once collapse has gone far enough to form a trapped surface, the theory itself guarantees that it breaks down somewhere inside. Penrose received a share of the Nobel Prize in Physics in 2020 for this work.

::: note
The case for black holes is a chain of failed supports. Gas pressure fails when fusion stops; electron degeneracy pressure fails above about 1.4 solar masses because relativistic electrons can no longer outpace gravity; neutron degeneracy pressure fails above about 2 solar masses. Beyond that, general relativity says nothing can stop the collapse, and Penrose showed that this does not depend on the star being perfectly round.
:::

The name came last. In late 1967, John Wheeler, who a decade earlier had argued against Oppenheimer that real collapse must somehow avoid the singularity, used the phrase **black hole** in a public lecture, and it stuck at once. The objects had a name before anyone had found one; how they were found is the subject of the final lesson of this chapter.

::: quiz
? Why does a white dwarf become unstable once its electrons are moving close to the speed of light?
- The electrons stop obeying the exclusion principle
+ The electrons' energy then grows only as 1/R, the same as gravity's, so shrinking no longer helps
- The star heats up and fusion restarts
= With relativistic electrons, E ∼ pc, so the electron term and the gravity term both scale as 1/R. If gravity wins at one radius, it wins at every radius, and the star collapses.
? Why does the Chandrasekhar limit not depend on what the star was originally made of?
- Because all stars have the same composition
+ Because the limit comes out in terms of constants of nature only: ħ, c, G and the proton mass
- Because white dwarfs have no electrons
= The scaling argument gives N_max ∼ (ħc/Gm_p²)^(3/2), with no property of the particular star in it. Only the mix of electrons per nucleon adjusts the number slightly.
? What did Penrose's 1965 theorem add to Oppenheimer and Snyder's result?
+ It showed that singularities form even without perfect spherical symmetry
- It showed that collapse stops at the neutron-star stage
- It showed that black holes emit radiation
= Oppenheimer and Snyder assumed a perfectly round, uniform ball. Penrose needed only a trapped surface and positive energy, so lumpiness and rotation cannot prevent the singularity.
:::
