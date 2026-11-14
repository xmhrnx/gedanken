By 1932 physicists knew what a nucleus is made of: protons and neutrons, packed at an enormous, fixed density. What they did not know was why it stays in one piece. Every known force argued that it should not. This lesson works out what the force holding nuclei together must be like, follows a young Japanese physicist's argument that its short reach required a new particle, and ends with the long detour it took to find that particle.

## Why nuclei should fly apart

Start with the size of the problem. Two protons 1 fm apart (a femtometre, 10⁻¹⁵ m, roughly the spacing of nucleons in a nucleus) repel each other electrically with a force of about 230 newtons. That is roughly the weight of a 23 kg suitcase, acting on a particle whose mass is about 1.7 × 10⁻²⁷ kg. Left to itself, that force would accelerate the two protons apart at something like 10²⁹ metres per second squared. Yet a uranium nucleus holds 92 protons inside a sphere only about 15 fm across, and it can sit there for billions of years.

Gravity cannot be the answer. Between two protons it is weaker than their electrical repulsion by a factor of about 10³⁶, so on the scale of a nucleus it is as good as absent. Neutrons do not help directly either: they carry no charge, so they cannot attract protons electrically. Something else must be at work, a third force strong enough at close quarters to beat electricity. It came to be called the **strong force**, or the nuclear force.

## What the force must be like

Nobody could look at this force directly. Its character had to be pieced together from indirect evidence: how nuclei scatter particles, how much energy binds them, which combinations of protons and neutrons exist and which do not. By the 1950s the picture had four main features, and each one explains something observed.

The first is simply that the force is **strong**. At the distances between neighbouring nucleons, its attraction is roughly a hundred times stronger than the electrical repulsion between two protons. Anything less, and no nucleus beyond hydrogen could exist.

The second is that it is **short-ranged**. Beyond about 2 fm it fades to almost nothing. The evidence for this is in an experiment already familiar: Rutherford's alpha particles scattered from gold in exactly the pattern that electrical repulsion alone predicts. That could only happen if the alphas, pushed back by the 79 protons of gold, never came close enough to the nucleus to feel anything else. They turned around tens of femtometres away, well outside the force's reach. The same short range explains why the strong force plays no part in the structure of atoms. Electrons spend their time tens of thousands of femtometres from the nucleus, where the strong force is utterly negligible and electricity has the field to itself.

The third feature is that the force is **blind to charge**. The attraction between two protons, between two neutrons, and between a proton and a neutron is very nearly the same. Electrically, protons and neutrons are as different as can be; to the strong force they look like two versions of a single particle, the nucleon. This is why neutrons are so useful in a nucleus. They add strong-force attraction without adding any electrical repulsion, which is why heavy nuclei, with many protons pushing each other apart, carry more neutrons than protons.

The fourth is that the force turns **repulsive at very short range**. When two nucleons are pushed closer than about 0.5 fm, the attraction gives way to a powerful push apart. This hard core keeps nucleons from collapsing into one another, and it is the reason the previous lesson found that nuclear matter has a fixed density: each nucleon keeps a minimum distance from its neighbours, so adding more nucleons adds more volume rather than squeezing them tighter.

## Why short range matters so much

The short range has a consequence that shapes nearly everything in nuclear physics. Because the strong force reaches only about as far as the nearest neighbours, each nucleon is attracted only by the handful of nucleons touching it. A nucleon in the middle of a uranium nucleus feels no more pull than one in the middle of a much smaller nucleus, because the extra nucleons farther away are out of reach. Making a nucleus bigger does not make each nucleon more tightly held.

Electrical repulsion behaves in exactly the opposite way. It is long-ranged, falling off only as the square of distance, so every proton repels every other proton in the nucleus, near or far. In a nucleus with Z protons there are about Z²/2 pairs pushing apart, and that number grows much faster than the number of neighbour bonds holding things together. So as nuclei get bigger, the repulsion gains on the attraction. That contest, a short-range attraction that saturates against a long-range repulsion that keeps growing, is why there is a limit to how heavy a stable nucleus can be, and ultimately why very heavy nuclei such as uranium can split.

It is also worth noticing how finely balanced the force is. The **deuteron**, the nucleus of heavy hydrogen with one proton and one neutron, is held together by only 2.2 MeV. That sounds like a lot next to chemistry, but it is small for a nucleus, about a sixth of what each pair of nucleons is worth inside helium. And no bound pair of two protons, or of two neutrons, exists at all. The strong force is strong, but only just strong enough.

## Yukawa's question: what makes a force short-ranged?

::: history 1934 · Osaka, Japan
Hideki Yukawa was a 27-year-old lecturer at Osaka Imperial University, working largely on his own, far from the centres of European physics. He asked a question that others had skirted: why should any force have a definite range at all? The best-understood force, electricity, had recently been described by quantum theory as the exchange of photons, particles of light, between charges. Photons have no mass, and the electrical force reaches to infinity, falling off gradually. Yukawa proposed that nucleons, too, attract by exchanging a particle, but one that has mass, and that the mass is precisely what limits the range. His paper appeared early in 1935 and went largely unnoticed outside Japan for two years.
:::

Yukawa's actual calculation used quantum field theory, but its essence can be captured in a rough estimate built on the energy–time version of the uncertainty principle. The estimate turns the measured range of the force into a prediction for the mass of the exchanged particle.

::: steps Estimating the mass of the exchanged particle
$ \Delta E\,\Delta t \sim \hbar
The energy–time uncertainty relation says that an energy can only be pinned down to within ΔE over a time interval Δt if their product is at least about ħ, the reduced Planck constant. Turned around, this means a system can depart from strict energy bookkeeping by ΔE for a time of about ħ/ΔE without any measurement being able to catch it. This is the loophole the argument uses.
---
$ \Delta t \sim \frac{\hbar}{mc^2}
To emit a particle of mass m, a nucleon must find at least the particle's rest energy, ΔE = mc², which it does not have to spare. Substituting this ΔE into the relation gives the length of time the particle can exist before it must be absorbed again by a neighbouring nucleon. The heavier the particle, the shorter its permitted life.
---
$ r \sim c\,\Delta t \sim \frac{\hbar}{mc}
Nothing travels faster than light, so in its short life the particle can cover a distance of at most about c times Δt. That distance is the farthest apart two nucleons can be and still exchange it, which makes it the range of the force. Now the connection is explicit: range is inversely proportional to the mass of the carrier, and a massless carrier, like the photon, gives unlimited range.
---
$ mc^2 \sim \frac{\hbar c}{r}
Multiply both sides by mc² and divide by r to turn the relation around, so that it gives the carrier's rest energy from the force's range. Writing it with ħc is convenient, because ħc has the tidy value of about 197 MeV·fm, which is exactly the right units for nuclear distances and energies.
---
$ mc^2 \sim \frac{197\,\text{MeV fm}}{1.4\,\text{fm}} \approx 140\,\text{MeV}
Put in a range of about 1.4 fm, the distance over which the force between nucleons is strongest, and the carrier should have a rest energy of roughly 140 MeV. That is about 270 times the electron's mass and about one seventh of the proton's, a weight between the two. Yukawa called such a particle a **meson**, from the Greek for "middle".
:::

The result was bold: a prediction of a particle nobody had seen, with a mass fixed by the size of nuclei. The "borrowing" picture should not be taken literally. The exchanged particle is not a real object briefly smuggled across; it is a convenient way of describing what quantum field theory calculates. But the full calculation gives the same relation between range and mass, and it also gives the shape of the force. The energy of interaction between two nucleons, in Yukawa's theory, falls off faster than any inverse-square law:

::: equation Yukawa potential
V(r) = -g^2 \frac{e^{-r/a}}{r}, \quad a = \frac{\hbar}{mc}
:::

Here g measures the strength of the coupling, playing the part that charge plays in electricity, and a is the range. At distances much shorter than a, the exponential factor is close to 1 and the potential looks like the 1/r of electricity. Beyond a, the exponential takes over and cuts the force off sharply: at three times the range it is already down to about 5% of what a 1/r law would give.

::: note
A force carried by a particle of mass m reaches only about ħ/(mc). The strong force between nucleons reaches about a femtometre, so its carrier should have a mass of about 140 MeV/c², and that particle, the pion, was found in 1947.
:::

## The wrong particle, then the right one

::: history 1936–1947 · Pasadena and Bristol
In 1936 Carl Anderson and Seth Neddermeyer, studying cosmic rays at Caltech, found a charged particle about 200 times as heavy as the electron. That was close enough to Yukawa's estimate that everyone assumed it was his meson, and Yukawa's paper suddenly became famous. But the new particle behaved wrongly. A carrier of the strong force should be grabbed almost instantly by any nucleus it meets, and in the 1940s experiments in Rome showed that this particle passed through nuclei as if they were barely there. It could not be the glue of the nucleus. The real carrier was found in 1947 by Cecil Powell, César Lattes and Giuseppe Occhialini in Bristol, who exposed photographic emulsions on mountaintops to catch cosmic rays. In the tracks they found a particle with a mass of about 140 MeV/c² that decayed into the lighter one. They called it the **pion**. The earlier particle, now called the **muon**, turned out to have nothing to do with the strong force; it is a heavier cousin of the electron. Yukawa received the Nobel Prize in Physics in 1949, the first Japanese laureate, and Powell received it in 1950.
:::

## The modern view

Yukawa's picture turned out to be true, but not the whole truth. Since the 1970s, protons, neutrons and pions have been understood as composite particles made of **quarks**, held together by a deeper force carried by massless particles called **gluons**. That deeper interaction is the fundamental strong force. The force between nucleons is a leftover effect of it, much as the forces that hold neutral molecules together are a leftover of the electrical forces inside them.

Pion exchange still describes the longer-range part of the force between nucleons very well, at distances beyond about 1.5 fm. Closer in, heavier mesons take over, and the repulsive core at half a femtometre can only really be understood in terms of quarks. But the central lesson of 1934 stands: a force's range tells you the mass of whatever carries it.

::: quiz
? Why didn't Rutherford's gold-foil experiment reveal the strong force?
- The strong force does not act on alpha particles
+ The alpha particles never got close enough to the gold nucleus to feel it
- The strong force is weaker than electricity at all distances
= Repelled by the 79 protons of gold, the alpha particles turned back tens of femtometres from the nucleus. The strong force reaches only a couple of femtometres, so they never felt it and scattered exactly as electricity alone predicts.
? If the strong force's carrier were twice as heavy, what would happen to the range?
+ It would halve
- It would double
- It would not change
= The range is about ħ/(mc), inversely proportional to the carrier's mass. A heavier carrier can exist for a shorter time and so travels a shorter distance before it must be absorbed.
? Why does each nucleon in a large nucleus feel attraction only from a few neighbours?
- Neutrons shield the force
+ The force is short-ranged, so distant nucleons contribute almost nothing
- The Pauli principle cancels the force from distant nucleons
= Beyond about 2 fm the strong force falls off exponentially, so only the nucleons immediately around a given nucleon contribute. That is why adding nucleons does not hold each one more tightly, while proton repulsion keeps growing.
:::
