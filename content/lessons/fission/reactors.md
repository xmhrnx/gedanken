A nuclear power reactor and a fission weapon run on the same reaction. In both, neutrons split uranium or plutonium nuclei, each split releases about 200 MeV and a few more neutrons, and those neutrons go on to cause further splits. The difference lies entirely in k, the average number of neutrons from each fission that cause another. A weapon drives k well above 1 for about a microsecond. A reactor holds k at exactly 1 for months on end, so that the chain neither grows nor dies, producing a steady stream of heat that boils water and drives turbines. This lesson explains how that balance can be held at all, which at first sight seems impossible, and why reactor fuel cannot explode like a bomb even when things go badly wrong.

## Holding k at 1

Keeping k at precisely 1 sounds absurdly delicate. No real system can be set to an exact number, and any small excess compounds from one generation to the next. If k were 1.001, a 0.1% excess, the number of fissions would grow by 0.1% in every generation. Whether that is a gentle drift or a runaway depends entirely on how long a generation lasts, and working through it shows why reactors are controllable and what makes them so.

::: steps Why delayed neutrons make control possible
$ k = 1.001, \quad \text{generation time} \approx 10^{-4}\,\text{s}
Start with a reactor that is very slightly supercritical. In most reactors neutrons are deliberately slowed down before they cause fission, as the next section explains, and the slowing takes time. So a generation of neutrons, from birth to the next fission, lasts about a tenth of a millisecond. That is about ten thousand times longer than in a bomb, but still very short.
---
$ 1.001^{10\,000} \approx e^{10} \approx 22\,000
A tenth of a millisecond per generation means 10,000 generations in one second. Raising 1.001 to the power 10,000 is the same as raising e to the power of about 10, since 1.001 is very close to e^0.001. The result is about 22,000. If this were the whole story, the reactor's power would rise more than twenty-thousand-fold in one second, far too fast for any mechanism or operator to respond.
---
$ \text{fraction delayed} \approx 0.65\% \text{ for } {}^{235}\mathrm{U}
It is not the whole story, because not all fission neutrons appear at the moment of fission. About 0.65% of them, for uranium-235, come later. Some fission fragments are radioactive in a particular way: they undergo a beta decay that leaves the new nucleus with so much energy that it promptly throws off a neutron. Since the beta decay happens on its own schedule, with half-lives of seconds to about a minute, these **delayed neutrons** emerge long after the fission that produced their parent.
---
$ \text{effective generation time} \approx 0.1\,\text{s}
Now suppose k exceeds 1 by less than this delayed fraction, as it does here: 0.1% is smaller than 0.65%. Then the prompt neutrons alone are not enough to keep the chain going, since without the delayed ones k would be below 1. Each generation must wait for its share of delayed neutrons before the chain can grow. That waiting dominates, and it stretches the average generation time from a tenth of a millisecond to around a tenth of a second.
---
$ 1.001^{10} \approx 1.01
With a generation lasting a tenth of a second, a second holds only about 10 generations, and 1.001 raised to the power 10 is about 1.01. The power now rises by about 1% per second. That is slow enough for control rods and automatic systems to follow comfortably, and slow enough for a person to watch. The tiny fraction of delayed neutrons has slowed the reaction by a factor of about a thousand.
:::

::: note
Delayed neutrons are what make reactors controllable. As long as k exceeds 1 by less than the delayed-neutron fraction, the pace of the chain is set by radioactive decay, measured in seconds, not by neutron flight, measured in fractions of a millisecond. A reactor that exceeds this margin can grow on prompt neutrons alone and is called **prompt critical**, and reactor design and operation are built around never reaching that state.
:::

::: model chain-reaction
Set k to exactly 1 and watch the population of neutrons hold steady from one generation to the next. A reactor's job is to stay at this setting for months, nudged slightly above or below to raise or lower the power.
:::

## Slowing the neutrons down

Most reactors run on **thermal** neutrons, ones slowed down until they move no faster than the surrounding atoms jostle from heat. The reason is that uranium-235 is hundreds of times more likely to absorb a slow neutron, and fission, than a fast one. A slow neutron lingers near each nucleus it passes, and quantum mechanics makes the effective target much larger at low energies. Fission neutrons are born fast, at about 2 MeV, so something has to slow them by a factor of tens of millions in energy.

That something is a **moderator**: a material whose nuclei absorb energy from neutrons in collisions without absorbing the neutrons themselves. The best moderators have light nuclei, and the reason is ordinary mechanics.

::: steps Why light nuclei make good moderators
$ \frac{\Delta E}{E} = \frac{4 m M}{\left(m + M\right)^2}
When a neutron of mass m strikes a stationary nucleus of mass M head-on and bounces off elastically, the fraction of its energy that it hands over is given by this expression, which comes from applying conservation of energy and of momentum to the collision. A glancing collision transfers less, but the head-on case shows how the transfer depends on the masses. It is largest when the two masses are equal.
---
$ M = m: \quad \frac{\Delta E}{E} = 1
For a hydrogen nucleus, a single proton, M is almost exactly equal to m, and the formula gives 1: in a head-on collision the neutron stops dead and hands over all its energy, just as one billiard ball does when it strikes another squarely. Even averaged over glancing collisions, a neutron loses about half its energy per collision with hydrogen, so it takes only around twenty collisions to slow from 2 MeV to thermal energy.
---
$ M = 12m: \quad \frac{\Delta E}{E} = \frac{48}{169} \approx 0.28
For carbon, the main ingredient of graphite, M is about 12 times m, and even a head-on collision transfers only about 28%. More collisions are needed, around a hundred, but carbon absorbs very few neutrons along the way, which makes graphite a good moderator despite being heavier.
---
$ M = 238m: \quad \frac{\Delta E}{E} \approx 0.017
For uranium-238, a head-on collision transfers less than 2% of the neutron's energy. A neutron bouncing off heavy nuclei barely slows, like a ping-pong ball bouncing off a bowling ball. That is why uranium cannot moderate itself, and why a separate light material is needed.
:::

Ordinary water, heavy water (water in which the hydrogen is the heavier isotope deuterium, which absorbs fewer neutrons than ordinary hydrogen) and graphite are the usual moderators. Because slow neutrons are so effective at causing fission in U-235, a moderated reactor can run on uranium with only a small proportion of U-235. Fermi's first pile in 1942 used natural uranium, with its 0.72% U-235, and graphite. Most power reactors today use **low-enriched uranium**, about 3 to 5% U-235, with ordinary water serving both as moderator and as coolant, the fluid that carries heat away from the fuel.

## Control rods

To adjust k, a reactor uses **control rods** made of materials that absorb neutrons strongly, such as boron or cadmium. Pushing the rods further into the core soaks up more neutrons and lowers k; pulling them out raises it. Because delayed neutrons make the chain respond slowly, small movements of the rods give smooth, gradual changes in power. Operators use the rods, along with neutron absorbers dissolved in the coolant in many designs, to start the reactor up, to adjust its power and to shut it down.

In an emergency, all the rods are driven into the core at once, which is called a *scram*. The chain reaction stops within seconds. The reactor does not go cold, though. The fission products built up in the fuel keep decaying, and that **decay heat** amounts to several percent of full power just after shutdown, falling over hours and days. It must be carried away for a long time afterwards, or the fuel will overheat even with the chain stopped. The accident at Fukushima in 2011 was caused by a failure to remove decay heat after the reactors had shut down.

## Why reactor fuel cannot explode like a bomb

A nuclear explosion needs a fast-neutron chain that is strongly supercritical, so that many generations, each lasting about 10 nanoseconds, pass before the material flies apart. Reactor fuel cannot meet that requirement, for two reasons that reinforce each other.

The first is the fuel itself. In uranium that is only 3 to 5% U-235, most of the nuclei are U-238, which captures or slows so many fast neutrons that a fast chain cannot be sustained at all, whatever the quantity of fuel. The same reasoning that rules out a fast chain in natural uranium rules it out here. The fuel reaches k = 1 only with a moderator present to slow the neutrons.

The second is speed. A moderated chain is slow: each generation of prompt neutrons takes at least a tenth of a millisecond, ten thousand times longer than in a bomb. If the power surges, the heat boils the water moderator or breaks apart the core long before a significant fraction of the fuel could fission. And in a water-moderated reactor, losing the water removes the moderator, which stops the chain.

None of this makes reactors harmless. A reactor can still suffer a severe accident, with overheating, melting of the fuel, explosions of steam or of hydrogen gas released from hot metal, and the release of radioactive material into the environment. What it cannot produce is a nuclear detonation.

::: history 26 April 1986 · Chernobyl, Ukraine
The worst reactor accident in history shows both what can go wrong and what cannot. Unit 4 at the Chernobyl power station was a Soviet RBMK reactor, moderated by graphite and cooled by water. During a poorly planned safety test in the early hours of the morning, its operators ran it at low power, in an unstable state, with most of the control rods withdrawn and several safety systems disabled. The design had a dangerous property. In most reactors, if the cooling water boils, k falls. In the RBMK, the graphite did the moderating while the water mainly absorbed neutrons, so when water turned to steam, fewer neutrons were absorbed and k *increased*. Heating made more power, which made more steam, which made more power. When the operators pressed the button to shut down, a flaw in the design of the rods meant that their first movement into the core briefly added reactivity instead of removing it. Within seconds the power surged to many times the reactor's rated level, the fuel shattered, and a steam explosion blew the massive lid off the reactor; a second explosion followed. The exposed graphite burned for days, carrying radioactive material across Europe. Two workers died that night, and 28 more, plant staff and firefighters among them, died of acute radiation sickness within months. Yet it was a runaway and a steam explosion, not a nuclear detonation. The core destroyed itself, and in destroying itself it stopped the chain.
:::

::: history 2 December 1942 · Chicago
The first reactor showed the other side of the story, how calmly a chain reaction can be handled when its timing is understood. When Fermi brought Chicago Pile-1 to criticality, he had the last control rod withdrawn in steps, calculating each time where the neutron count should level off, and checking that it did. Partway through, with the pile still just short of critical, he even called a break for lunch. He could proceed so deliberately because he understood delayed neutrons: once the pile was critical, its power doubled only about every two minutes, slow enough for him to watch it climb, judge it, and order the rods back in.
:::

::: quiz
? What makes a reactor's chain reaction slow enough to control?
- Control rods slow the neutrons down
+ A small fraction of neutrons is emitted seconds after fission
- The moderator absorbs most of the energy
= Delayed neutrons, about 0.65% of the total for U-235, come from the decay of fission fragments seconds later. As long as the reactor relies on them, the effective generation time stretches to around 0.1 s, slow enough for control.
? Why do most reactors use a moderator?
+ Slow neutrons are far more likely to cause U-235 to fission
- Fast neutrons are too dangerous to the operators
- The moderator produces extra neutrons
= U-235 is hundreds of times more likely to fission with a slow neutron than a fast one. Slowing the neutrons raises the chance of fission so much that low-enriched fuel can reach k = 1.
? What was the Chernobyl explosion?
- A small nuclear detonation
+ A power surge that caused a steam explosion and destroyed the core
- A hydrogen bomb test that went wrong
= The reactor's power ran away within seconds, shattering the fuel and causing a steam explosion. Low-enriched, moderated fuel cannot detonate like a bomb, and the destruction of the core stopped the chain.
:::
