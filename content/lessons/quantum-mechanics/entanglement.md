When two particles interact, quantum mechanics allows them to end up sharing a single state, so that neither particle has a definite state of its own. Measure one, and you learn at once what a measurement on the other will show, however far away it has travelled. Einstein thought this meant quantum mechanics must be incomplete: there had to be some hidden facts, carried by each particle, that the wavefunction leaves out. For thirty years the question looked like philosophy. Then John Bell showed it could be tested, and experiments have since settled it against Einstein's view, though in a way that still does not allow any message to travel faster than light. This lesson follows that argument, which many physicists regard as the deepest result about what quantum mechanics says the world is like.

## A pair with one state between them

Some processes produce two particles whose spins must add to zero. An atom with no spin can emit two particles in opposite directions, for example, and since angular momentum is conserved, if one has spin up along some axis, the other must have spin down along it. That much would be true in classical physics too. The quantum twist is that neither particle is in a definite spin state before it is measured. The pair is in a superposition of both possibilities at once, called the **singlet** state:

::: equation The spin singlet
\psi = \frac{1}{\sqrt{2}}\left(\uparrow_A\downarrow_B - \downarrow_A\uparrow_B\right)
:::

Read this as: "A up and B down" plus, with a minus sign, "A down and B up", each with equal weight. The factor 1/√2 makes the probabilities add to one, since each term then has probability ½. The minus sign matters: it is what makes the pair's total spin zero along every axis, not just along z.

A pair in a state like this is called **entangled**. The defining feature is that the state of the whole cannot be written as a state of A multiplied by a state of B. There is no description of particle A on its own that captures what is going on, only a description of the pair.

Now give the particles to two experimenters, conventionally called Alice and Bob, who carry them far apart. Alice measures her particle's spin along z. She gets up or down, at random, with even odds. Bob measures his particle along the same axis, and he always gets the opposite of Alice. The same is true for any axis they choose, provided both choose the same one.

::: history 1935 · Princeton
Albert Einstein, with Boris Podolsky and Nathan Rosen, argued that this situation proves quantum mechanics is incomplete. Their reasoning went like this. Bob can predict his own result with certainty, without touching his particle, simply by asking Alice what she found. If a quantity can be predicted with certainty without disturbing a system, they said, then it must correspond to some real property of that system, an "element of reality". Since the wavefunction does not specify Bob's result in advance, it leaves out something real. The only alternative seemed to be that Alice's measurement instantly changes a distant particle, which Einstein later called "spooky action at a distance" and could not accept. Their paper used position and momentum; in 1951 David Bohm recast it in terms of spins, the form used here, which made the argument much easier to test.
:::

## Hidden instructions?

The perfect anticorrelation, on its own, is easy to explain with nothing spooky at all. Put a red card in one envelope and a blue card in another, mix them up, and post one to each of two cities. Whoever opens an envelope learns instantly what colour the other city has, and nothing travelled between them; the answer was fixed when the envelopes were sealed. The analogy has one important limit, and it is exactly where Bell found his opening: a card has only one colour to reveal, but a spin can be asked about along infinitely many axes.

Still, the idea generalises naturally. Perhaps each particle leaves the source carrying a hidden set of instructions, one answer for every axis it might be measured along, arranged so that the partner's answers are always opposite. Theories of this kind are called **local hidden-variable** theories: "hidden variable" because the instructions are not in the wavefunction, and "local" because nothing done to one particle can affect the other instantly. For decades it seemed impossible to distinguish such a theory from quantum mechanics, so the choice between them looked like a matter of taste.

## Bell's theorem

In 1964 John Bell, a physicist from Belfast working at CERN, showed that it is not. Hidden instructions and quantum mechanics agree when Alice and Bob measure along the same axis, but they make different predictions when the two measure along different axes. The clearest form of his argument was worked out in 1969 by John Clauser, Michael Horne, Abner Shimony and Richard Holt, and it is known by their initials as CHSH.

Here is the setup. On each run, Alice chooses at random between two axes, a and a′, and Bob between two others, b and b′. Each records +1 for up and −1 for down. For any pair of axes, E is the average over many runs of the product of their two results. If they always agree, E = +1; if they always disagree, E = −1; if their results are unrelated, E = 0.

::: steps The local limit
$ A, A', B, B' = \pm 1
Assume local hidden variables. Then each particle leaves the source with a definite answer for every axis: A is what Alice would get along a, A′ what she would get along a′, and likewise B and B′ for Bob. Each answer is +1 or −1, and Bob's answers cannot depend on which axis Alice chooses, because the choice is made far away. This assumption is called local realism.
---
$ AB - AB' + A'B + A'B' = A(B - B') + A'(B + B')
For one pair of particles, combine the four possible products in this particular way, then factor out A from the first two terms and A′ from the last two. This is ordinary algebra, allowed because all four answers exist at once under our assumption, even though in any real run only one pair of them is measured.
---
$ A(B - B') + A'(B + B') = \pm 2
B and B′ are each +1 or −1, so either they are equal or opposite. If equal, B − B′ = 0 and B + B′ is ±2. If opposite, B + B′ = 0 and B − B′ is ±2. Either way one bracket vanishes and the other is ±2, and multiplying by A or A′, which are ±1, leaves ±2. So for every single pair, the combination is exactly +2 or −2.
---
$ S = E(a,b) - E(a,b') + E(a',b) + E(a',b'), \qquad |S| \leq 2
Now average over many pairs. A quantity that is +2 or −2 on every pair must have an average somewhere between −2 and +2. The average of the combination is the same combination of averages, which defines S. So any local hidden-variable theory, whatever its instructions, must give |S| no larger than 2.
:::

That bound, |S| ≤ 2, is a **Bell inequality**. Notice how little went into it: only that the answers exist before measurement and that one side's choice cannot affect the other side's answer. Nothing about quantum mechanics was used.

## What quantum mechanics predicts

Now compute the same quantity from quantum mechanics, using the spin rule from the previous lesson: a spin prepared along one axis and measured along another at angle θ comes out up with probability cos²(θ/2).

::: steps The quantum prediction
$ E(\theta) = P(\text{same}) - P(\text{opposite})
Suppose Alice's and Bob's axes are at angle θ. The product of their results is +1 when they get the same answer and −1 when they get opposite answers, so its average is the probability of agreeing minus the probability of disagreeing.
---
$ E(\theta) = \sin^2\frac{\theta}{2} - \cos^2\frac{\theta}{2} = -\cos\theta
Once Alice finds up along her axis, the singlet guarantees that Bob's particle now behaves as spin down along Alice's axis. Measured along Bob's axis, at angle θ, it comes out up, the same as Alice, with probability sin²(θ/2), and opposite with probability cos²(θ/2). The same holds if Alice found down. The trigonometric identity cos²(θ/2) − sin²(θ/2) = cos θ turns the difference into −cos θ.
---
$ a = 0^\circ,\; a' = 90^\circ,\; b = 45^\circ,\; b' = 135^\circ
Choose the axes to make the quantum result as different from the local one as possible. With these choices, three of the four pairs used in S, namely a with b, a′ with b, and a′ with b′, are 45° apart, while a and b′ are 135° apart.
---
$ S = -\tfrac{\sqrt{2}}{2} - \tfrac{\sqrt{2}}{2} - \tfrac{\sqrt{2}}{2} - \tfrac{\sqrt{2}}{2} = -2\sqrt{2}
Each pair 45° apart gives E = −cos 45° = −√2/2. The pair 135° apart gives E = −cos 135° = +√2/2, but it enters S with a minus sign, so it also contributes −√2/2. The four terms add to −2√2, so |S| = 2√2, about 2.83, comfortably beyond the local limit of 2.
:::

Pause on what this means. No set of instructions carried by the particles, however cleverly arranged, can produce correlations this strong. If experiments find |S| greater than 2, then at least one of the two assumptions behind the Bell inequality is false: either the results were not fixed before measurement, or a choice made on one side can affect what happens on the other. Bell's theorem does not rule out every hidden-variable theory, only local ones. Bohm's pilot-wave theory, for example, survives precisely because its hidden variables are non-local, linked instantly across any distance.

::: history 1972–2022 · Berkeley, Orsay, Innsbruck, Delft
The first test came in 1972, when Stuart Freedman and John Clauser at Berkeley measured pairs of photons emitted by calcium atoms and found the quantum prediction, not the local limit. Critics pointed to loopholes. In that experiment the settings were fixed in advance, so in principle some unknown signal could have told each side what the other was measuring. In 1982 Alain Aspect's team at Orsay, near Paris, changed the settings while the photons were in flight, too late for any signal at light speed to coordinate the two sides. In 1998 Anton Zeilinger's group in Innsbruck tightened this further, with settings chosen by random-number generators at stations 400 metres apart. The last major loophole was that detectors missed many particles, so the ones caught might not be representative. In 2015 groups in Delft, Vienna and Boulder each closed this together with the light-speed loophole in a single experiment. Every test has violated Bell's inequality, by the amount quantum mechanics predicts. Aspect, Clauser and Zeilinger shared the 2022 Nobel Prize in Physics.
:::

## No signals faster than light

If Alice's measurement seems to fix Bob's result instantly, can she use it to send him a message? She cannot, and the reason is worth seeing. Whatever axis Alice chooses, and whatever result she gets, Bob looking only at his own particles sees up and down with even odds, a random sequence indistinguishable from coin flips. Nothing Alice does changes those odds. The correlations exist, but they only become visible when the two lists of results are brought together and compared, and bringing them together requires an ordinary signal, travelling at or below the speed of light. Quantum mechanics guarantees this in general, in a result called the no-communication theorem, which is why entanglement and Einstein's relativity can coexist.

The correlations are nonetheless useful. They underlie quantum key distribution, in which two parties use entangled particles to build a shared secret key, and any eavesdropper unavoidably disturbs the correlations and so reveals the intrusion. They also allow quantum teleportation, which transfers the exact quantum state of one particle onto another far away, using entanglement together with an ordinary message, so that nothing travels faster than light.

::: note
Entangled particles share one state, and neither has a state of its own. Bell showed that their correlations at different angles cannot come from instructions fixed in advance: local hidden variables give |S| ≤ 2, quantum mechanics gives 2√2, and experiments agree with quantum mechanics. Even so, no usable signal passes between the particles.
:::

::: quiz
? Alice and Bob measure the two particles of a singlet pair along the same axis. What do they find?
- The same result every time
+ Opposite results every time
- Results with no relation to each other
= At θ = 0 the quantum prediction is E = −cos 0 = −1, which means the products of their results are always −1: whenever one finds up, the other finds down.
? Why can Alice not use entanglement to send Bob a message instantly?
- The particles stop being entangled once they are separated
+ Bob's results on their own are 50–50 random, whatever Alice does
- Measurements always take longer than light takes to travel between them
= Alice's choices do not change the odds Bob sees, so the correlation shows up only when both lists are compared, and that comparison needs an ordinary signal.
? An experiment finds S = 2.7. What does this rule out?
- Quantum mechanics
+ Any theory in which each particle carries fixed answers and distant choices cannot affect results
- Einstein's theory of relativity
= Local hidden variables require |S| ≤ 2, so a value of 2.7 rules them out, while it fits quantum mechanics, which allows values up to 2√2 ≈ 2.83.
:::
