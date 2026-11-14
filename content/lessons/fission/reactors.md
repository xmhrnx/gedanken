A nuclear power reactor and a fission weapon run on the same reaction: neutrons split uranium or plutonium nuclei, releasing about 200 MeV each and more neutrons. The difference lies in k. A bomb drives k well above 1 for a microsecond. A reactor holds k at exactly 1 for months, producing a steady stream of heat that boils water and drives turbines.

## Holding k at 1

Keeping k at precisely 1 sounds impossibly delicate. If k were 1.001, a 0.1% excess, the reaction would grow by 0.1% every generation. Whether that is manageable depends entirely on how long a generation lasts.

::: steps Why delayed neutrons make control possible
$ k = 1.001, \quad \text{generation time} \approx 10^{-4}\,\text{s}
In a reactor, neutrons are slowed down before they cause fission, so a generation of *prompt* neutrons takes about a tenth of a millisecond, far longer than in a bomb but still very short.
---
$ 1.001^{10\,000} \approx e^{10} \approx 22\,000
In one second that is 10,000 generations. If all neutrons were prompt, the power would rise more than twenty-thousand-fold in a second, much too fast for any mechanism or operator to respond.
---
$ \text{fraction delayed} \approx 0.65\% \text{ for } {}^{235}\mathrm{U}
About 0.65% of fission neutrons are not emitted at the moment of fission. They come from the radioactive decay of certain fission fragments, typically seconds to about a minute later.
---
$ \text{effective generation time} \approx 0.1\,\text{s}
If k exceeds 1 by less than the delayed fraction, the chain cannot grow on prompt neutrons alone; each generation has to wait for its delayed neutrons. That stretches the average generation to around a tenth of a second.
---
$ 1.001^{10} \approx 1.01
Now one second holds about 10 generations, and the power rises by about 1% per second. That is slow enough for control rods and automatic systems to follow comfortably.
:::

::: note
Delayed neutrons are the reason reactors are controllable. As long as the excess of k above 1 stays smaller than the delayed-neutron fraction, the reaction's pace is set by radioactive decay measured in seconds, not by neutron flight measured in microseconds. A reactor that exceeds this margin is called **prompt critical**, and reactor design and operation are built around never reaching it.
:::

## Moderators

Most reactors use **thermal** neutrons, slowed down to the energies of the surrounding atoms, because U-235 is hundreds of times more likely to fission when struck by a slow neutron than by a fast one. A **moderator** slows them: a material with light nuclei that absorbs few neutrons. A neutron colliding with a light nucleus loses a large share of its energy in each collision, as a billiard ball does when it hits another of the same mass. Ordinary water, heavy water and graphite are the usual moderators.

Because slow neutrons are so effective, a moderated reactor can run on uranium with only a small proportion of U-235. Fermi's first pile in 1942 used natural uranium and graphite. Most power reactors today use **low-enriched uranium**, about 3–5% U-235, with water serving as both moderator and coolant.

## Control rods

**Control rods** contain elements that absorb neutrons strongly, such as boron or cadmium. Inserting them lowers k; withdrawing them raises it. Operators use them, along with dissolved absorbers in the coolant, to start up, adjust power and shut down. In an emergency, all rods drop in at once (a *scram*) and the chain stops within seconds, though the fission products continue to produce decay heat that must be removed for a long time afterwards.

## Why reactor fuel cannot explode like a bomb

A nuclear explosion needs a fast-neutron chain that is strongly supercritical, so that many generations, each lasting about 10 nanoseconds, pass before the material flies apart. Reactor fuel cannot do this:

- In fuel that is 3–5% U-235, the U-238 captures or slows so many neutrons that a fast chain cannot be sustained at all, whatever the quantity. The fuel reaches k = 1 only with a moderator.
- A moderated chain is slow: each generation takes at least a tenth of a millisecond. If power surges, the heat boils the water moderator or breaks apart the core long before a significant fraction of the fuel could fission. Losing the moderator stops the chain.

A reactor can still suffer a severe accident, with overheating, melting of the fuel, steam or hydrogen explosions and release of radioactive material. It cannot produce a nuclear detonation.

::: history 26 April 1986 · Chernobyl, Ukraine
During a poorly conducted safety test at night, operators of Unit 4, a graphite-moderated, water-cooled reactor of the Soviet RBMK type, ran it at low power with most control rods withdrawn and safety systems disabled. In this design, when the cooling water turned to steam, k *increased*, so heating led to more power, which led to more steam. When the operators tried to shut down, the rods' design briefly added reactivity. Power surged to many times its rated level within seconds, the fuel shattered, and a steam explosion blew the heavy lid off the reactor; a second explosion followed. The burning graphite spread radioactive material across Europe. Two workers died that night and 28 more of acute radiation sickness within months. It was a runaway and a steam explosion, not a nuclear detonation: the core destroyed itself and the chain stopped.
:::

::: history 2 December 1942 · Chicago
When Fermi brought Chicago Pile-1 to criticality, he withdrew the last control rod in steps, calculating each time where the neutron count should level off, and took a lunch break midway. He could do this calmly because he understood delayed neutrons: the pile's power rose by a factor of two only about every two minutes.
:::

::: quiz
? What makes a reactor's chain reaction slow enough to control?
- Control rods slow the neutrons down
+ A small fraction of neutrons is emitted seconds after fission
- The moderator absorbs most of the energy
= Delayed neutrons, about 0.65% for U-235, stretch the effective generation time to around 0.1 s.
? Why do most reactors use a moderator?
+ Slow neutrons are far more likely to cause U-235 to fission
- Fast neutrons are too dangerous to the operators
- The moderator produces extra neutrons
= Slowing neutrons raises the chance of fission so much that low-enriched fuel can reach k = 1.
? What was the Chernobyl explosion?
- A small nuclear detonation
+ A power surge that caused a steam explosion and destroyed the core
- A hydrogen bomb test that went wrong
= The reactor ran away and destroyed itself, but low-enriched moderated fuel cannot detonate like a bomb.
:::
