Two particles that have interacted can share a single quantum state, so that neither has a definite state of its own. Measure one and you learn what the other will do, however far away it is. Einstein thought this showed quantum mechanics was incomplete. Experiments have since settled the question against him, though not in a way that allows messages faster than light.

## An entangled pair

Some processes produce two particles whose total spin is zero. If one has spin up along z, the other has spin down. But neither is in a definite state before measurement; the pair is in a superposition of both possibilities:

::: equation The spin singlet
\psi = \frac{1}{\sqrt{2}}\left(\uparrow_A\downarrow_B - \downarrow_A\uparrow_B\right)
:::

A pair in such a state is **entangled**: the state of the whole cannot be written as a state of A times a state of B. Alice measures particle A along z and gets up or down, 50–50. Bob, measuring B along the same axis, always gets the opposite. The same holds for any axis, as long as both use the same one.

::: history 1935 · Princeton
Albert Einstein, Boris Podolsky and Nathan Rosen argued that this proves quantum mechanics incomplete. Bob can predict his result with certainty by asking Alice, without disturbing his particle. So, they reasoned, the result must have been fixed all along by some element of reality that the wavefunction leaves out. The alternative, that Alice's measurement instantly affects a distant particle, Einstein later called "spooky action at a distance". Their paper used position and momentum; David Bohm recast it with spins in 1951.
:::

Perfect anticorrelation on its own is easy to explain without anything spooky. Put one red and one blue card in two envelopes and mail them apart: opening one tells you the other. Perhaps each particle carries hidden instructions for every axis. For thirty years this looked like a question of taste.

## Bell's theorem

In 1964 John Bell, a physicist at CERN, showed that it is not. Hidden instructions and quantum mechanics make different predictions when Alice and Bob measure along *different* axes. The cleanest form, due to John Clauser, Michael Horne, Abner Shimony and Richard Holt in 1969, goes as follows. Alice chooses at random between axes a and a′, Bob between b and b′. Each result is recorded as +1 (up) or −1 (down), and E(a, b) is the average of the product of their results.

::: steps The local limit
$ A, A', B, B' = \pm 1
Assume each particle carries definite answers for every axis it might be asked about, fixed at the source, and that Alice's choice cannot affect Bob's answer. This is local realism.
---
$ AB - AB' + A'B + A'B' = A(B - B') + A'(B + B')
Group the terms for one pair of particles.
---
$ A(B - B') + A'(B + B') = \pm 2
B and B′ are each ±1, so either they are equal, making B − B′ = 0 and B + B′ = ±2, or opposite, making B + B′ = 0 and B − B′ = ±2. Either way one bracket vanishes and the other gives ±2.
---
$ S = E(a,b) - E(a,b') + E(a',b) + E(a',b'), \qquad |S| \leq 2
Averaging over many pairs, a quantity that is always +2 or −2 has an average between −2 and +2.
:::

Quantum mechanics predicts something else. From the spin rule of the previous lesson, if Alice's axis and Bob's are at angle θ, the singlet gives E = −cos θ.

::: steps The quantum prediction
$ E(\theta) = P(\text{same}) - P(\text{opposite}) = \sin^2\frac{\theta}{2} - \cos^2\frac{\theta}{2} = -\cos\theta
Once Alice finds up along her axis, Bob's particle acts as spin down along it. At angle θ, it comes out the same as Alice's with probability sin²(θ/2).
---
$ a = 0^\circ,\; a' = 90^\circ,\; b = 45^\circ,\; b' = 135^\circ
Choose axes so that three of the pairs are 45° apart and one, a and b′, is 135° apart.
---
$ S = -\tfrac{\sqrt{2}}{2} - \tfrac{\sqrt{2}}{2} - \tfrac{\sqrt{2}}{2} - \tfrac{\sqrt{2}}{2} = -2\sqrt{2}
cos 45° = √2/2 and cos 135° = −√2/2, so E(a, b′) = +√2/2, and it enters S with a minus sign. |S| = 2√2 ≈ 2.83, beyond the local limit of 2.
:::

No assignment of hidden answers can reproduce this. If experiments find |S| > 2, then either the results were not fixed in advance or a measurement here affects outcomes there. Bell's theorem rules out local hidden variables; it leaves open non-local ones, such as Bohm's pilot waves.

::: history 1972–2022 · Berkeley, Orsay, Innsbruck, Delft
Stuart Freedman and John Clauser made the first test in 1972 with photon pairs from calcium atoms, and found the quantum prediction. In 1982 Alain Aspect's team in Orsay switched the measurement settings while the photons were in flight, so no signal at light speed could coordinate the two sides. Anton Zeilinger's group in Innsbruck closed this gap more strictly in 1998 with random settings and stations 400 m apart. In 2015 teams in Delft, Vienna and Boulder closed the remaining loopholes together, including detectors missing too many particles. Every test has violated Bell's inequality. Aspect, Clauser and Zeilinger shared the 2022 Nobel Prize.
:::

## No signals faster than light

Entanglement cannot carry a message. Whatever axis Alice picks, Bob on his own sees up and down 50–50, a random sequence. The correlations appear only when the two lists are brought together and compared, which needs an ordinary signal at or below light speed. This is guaranteed by the theory (the no-communication theorem), so entanglement and relativity coexist.

The correlations are nonetheless useful. They underlie quantum key distribution, which makes eavesdropping detectable, and quantum teleportation, which transfers a quantum state using entanglement plus an ordinary message.

::: note
Entangled particles share one state. Bell showed that their correlations at different angles cannot come from instructions fixed in advance: local hidden variables give |S| ≤ 2, quantum mechanics gives 2√2, and experiments agree with quantum mechanics. No usable signal travels between the particles.
:::

::: quiz
? Alice and Bob measure an entangled singlet along the same axis. What do they find?
- The same result every time
+ Opposite results every time
- Unrelated results
= At θ = 0, E = −cos 0 = −1: the results are perfectly anticorrelated.
? Why can't Alice use entanglement to send Bob a message instantly?
- The particles lose entanglement when separated
+ Bob's results alone are 50–50 random whatever Alice does
- Measurements take longer than light travel time
= The correlation shows up only when both lists are compared, which needs an ordinary signal.
? An experiment finds S = 2.7. What does this rule out?
- Quantum mechanics
+ Any theory in which each particle carries fixed answers and distant choices cannot influence results
- Relativity
= Local hidden variables require |S| ≤ 2; quantum mechanics allows up to 2√2 ≈ 2.83.
:::
