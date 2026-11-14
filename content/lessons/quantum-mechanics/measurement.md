Schrödinger's equation describes how a wavefunction changes when nothing is being measured, and it does so smoothly and predictably, with no randomness at all: given ψ now, ψ at every later moment is fixed. Yet every real experiment ends with something quite different, a single definite result that could not have been predicted in advance: one dot on a screen, one click of a detector, one reading on a dial. Quantum mechanics bridges the two with a separate rule about what happens when a measurement is made. That rule works perfectly in practice, and it is also where nearly all the arguments about what quantum mechanics means take place. This lesson sets out the rule, shows it at work on a simple example, and then explains why physicists still disagree about it.

## A state with no definite energy

Take the particle in a box from the earlier lessons. It has stationary states φ₁, φ₂, φ₃ and so on, each with a definite energy E₁, E₂, E₃. Because Schrödinger's equation is linear, any sum of solutions is also a solution, so the particle can be in a combination of two levels at once, called a **superposition**:

::: equation A superposition of two energy levels
\psi = c_1\,\phi_1 + c_2\,\phi_2
:::

The coefficients c₁ and c₂ are complex numbers that say how much of each level is present. Such a state does not have a definite energy. It is tempting to read it as "the particle is really in level 1 or level 2, and we just don't know which", the way a coin under a cup is really heads or tails. That reading is wrong, and the reason is interference. The two parts of a superposition can interfere with each other, just as the waves from two slits did, and a mixture of hidden alternatives cannot do that. A superposition is a genuinely different state from either of its parts, not a statement of ignorance about them.

Yet when you measure the energy of a particle in this state, you never get anything in between. You get exactly E₁ or exactly E₂, and the odds are given by a version of the Born rule.

::: equation The Born rule for outcomes
P(E_1) = |c_1|^2, \qquad P(E_2) = |c_2|^2, \qquad |c_1|^2 + |c_2|^2 = 1
:::

The probability of each result is the squared size of its coefficient. The last condition says the two probabilities add to one, because the measurement must give one result or the other. The rule is general. For any quantity you might measure, whether energy, momentum or spin, there is a set of states in which that quantity has a definite value. Write ψ as a sum of those states, and the squared size of each coefficient is the probability of getting that value. The position version from the wavefunction lesson is the same rule applied to states of definite position.

## What a measurement leaves behind

The rule has a second half. Suppose a measurement of energy gives E₂. If you measure the energy again immediately, you get E₂ again, with certainty. That is what experiments show, and it means the state after the first measurement must be φ₂, the state with definite energy E₂. The superposition has been replaced by the one term that matches the result. This sudden change is called the **collapse** of the wavefunction.

A worked example makes the bookkeeping concrete.

::: steps One measurement, step by step
$ \psi = \sqrt{\tfrac{1}{3}}\,\phi_1 + \sqrt{\tfrac{2}{3}}\,\phi_2
Prepare a superposition of the two lowest levels. The coefficients are chosen so that their squares, ⅓ and ⅔, add to one; any other choice would not describe a normalised state, because the probabilities of the possible results must total one.
---
$ P(E_1) = \tfrac{1}{3}, \qquad P(E_2) = \tfrac{2}{3}
Measure the energy. The Born rule squares each coefficient to give the probability of each outcome. Nothing in ψ says which result this particular run will produce; the state fixes only the odds, and on any one run either result can occur.
---
$ \psi \to \phi_2
Suppose this run gives E₂. Immediately afterwards the state is φ₂ alone, the term that matches the result. We know this because a repeated energy measurement now gives E₂ every time, which is only true of φ₂. The information about the other branch, and with it the ability to interfere, is gone.
---
$ \langle E \rangle = \tfrac{1}{3}E_1 + \tfrac{2}{3}E_2
Now repeat the whole procedure on many particles, each prepared in the same superposition. Each run gives E₁ or E₂, and over many runs the average result is each value weighted by its probability. This average is called the **expectation value**, written ⟨E⟩.
---
$ \langle E \rangle = \tfrac{1}{3}E_1 + \tfrac{2}{3}\times 4E_1 = 3E_1
For a particle in a box the levels go as n², so E₂ = 4E₁. Putting that in gives an average of 3E₁. Notice that no single measurement ever gives 3E₁: the expectation value is a property of the long run, not a possible result.
:::

This explains what happened in the double-slit experiment when a detector recorded which slit each electron passed through. Recording the path is a measurement. It collapses the state ψ₁ + ψ₂ onto one of its two terms, so the interference term, which needed both terms to be present, disappears, and the fringes go with it.

## The measurement problem

Look at the two rules side by side and they do not fit comfortably. Schrödinger's equation is deterministic and smooth, and it never removes a term from a superposition. Collapse is random and abrupt, and it does exactly that. A theory with two different laws of motion has to say when each one applies, and here the answer is "during a measurement". But what counts as a measurement?

The difficulty is sharp because measuring devices are themselves made of atoms, and atoms obey Schrödinger's equation. So treat the whole apparatus quantum mechanically. Suppose a detector shows "1" if it finds energy E₁ and "2" if it finds E₂. Because the equation is linear, a particle in a superposition of E₁ and E₂ should drive the detector into a superposition of showing "1" and showing "2". Nobody has ever seen a dial in such a state; we always see one reading. Somewhere between the particle and our perception of the dial, the rule seems to switch from smooth evolution to collapse, and the theory as usually stated does not say where or why. This is the **measurement problem**.

::: history 1932–1935 · Princeton and Oxford
John von Neumann, a Hungarian mathematician then at Princeton, gave quantum mechanics its rigorous mathematical form in a 1932 book. He was candid about the two kinds of change: one governed by Schrödinger's equation, the other a discontinuous jump on measurement, which he wrote in as a separate postulate. He also showed that the line between the measured system and the measuring device could be moved almost anywhere without changing any prediction, which made it hard to say where collapse really happens. In 1935 Erwin Schrödinger, by then in Oxford and increasingly unhappy with the direction the theory had taken, pushed the problem to an absurd extreme. He imagined a cat sealed in a box with a radioactive atom, a detector and a flask of poison, rigged so that the poison is released if the atom decays. Applied literally, the rules seem to leave the cat in a superposition of alive and dead until someone opens the box. Schrödinger meant this as a reductio, an argument that something in the rules must be missing, not as a claim about cats.
:::

## Why big things look classical

Part of the puzzle is now well understood. A cat, a dial or a dust grain is never isolated. It is constantly struck by air molecules and light, and it gives off heat. Each of these interactions leaves a record of the object's state in its surroundings: a photon that bounced off the dial's needle carries away information about where the needle is. Interference between two possibilities requires their phases to stay in a definite relation, and once that phase information has leaked into billions of untraceable particles in the environment, it can never be gathered back. For all practical purposes the interference term vanishes.

This process is called **decoherence**. It was studied from the 1970s onward by H. Dieter Zeh, Wojciech Zurek and others, and it is fast: for an object the size of a dust grain in air, decoherence is far quicker than any measurement could follow. It explains why large objects never show interference and why the world we see behaves classically. It has been watched happening, in experiments that let carefully isolated systems slowly interact with their surroundings.

What decoherence does not do is pick an outcome. After decoherence, the state of particle, detector and environment together still contains both results, now unable to interfere. It explains why we never see a blur of "1" and "2", but not why we see "1" on this run rather than "2". That last step is where the interpretations of quantum mechanics part company.

## What the rules might mean

Physicists agree completely on how to use the rules and on every prediction they make. What they disagree about is what the rules describe.

The oldest family of views is associated with Niels Bohr and Werner Heisenberg in Copenhagen. On these views the wavefunction is a tool for predicting the results of measurements, and asking what the particle is doing between measurements goes beyond what physics can meaningfully say. Collapse is then not a physical event but an update in what we can predict.

In 1957 Hugh Everett, a graduate student at Princeton, proposed taking Schrödinger's equation at its word and dropping collapse altogether. If the detector and the observer end up in a superposition of seeing "1" and seeing "2", then both happen: the universe divides into branches, each containing one result and one version of the observer, and decoherence keeps the branches from ever interfering again. This became known as the **many-worlds** interpretation. Its difficulty is to explain why the Born rule's probabilities apply when every outcome occurs.

A third approach keeps definite particles. Louis de Broglie suggested in 1927, and David Bohm developed fully in 1952, that particles always have exact positions and are steered by the wavefunction, which acts as a **pilot wave**. The theory reproduces every prediction of standard quantum mechanics. Its randomness comes only from our ignorance of the starting positions, but the price is that the guiding influence acts instantly across any distance.

A fourth approach treats collapse as real physics. In 1986 Giancarlo Ghirardi, Alberto Rimini and Tullio Weber proposed adding a small random term to Schrödinger's equation that occasionally localises particles. For a single particle the effect is far too rare to notice, but in a large object containing many particles it acts almost at once. Unlike the others, such **objective collapse** theories predict small departures from standard quantum mechanics, and experiments are searching for them. Other views exist too, including QBism, which treats the wavefunction as a measure of an agent's degrees of belief. So far no experiment has singled out any one interpretation.

::: note
Quantum mechanics gives exact probabilities for the results of measurements, and after a measurement the state matches the result obtained. Decoherence explains why large objects never show interference, but how a deterministic equation yields one particular random outcome is the measurement problem, and it remains open.
:::

::: quiz
? A particle is in the state ψ = 0.6φ₁ + 0.8φ₂. What is the probability of measuring the energy E₂?
- 0.8
+ 0.64
- 0.36
= The probability is the squared size of the coefficient, 0.8² = 0.64, and together with 0.6² = 0.36 for E₁ the total is one, as it must be.
? An energy measurement gives E₁. What does an immediate second energy measurement give?
+ E₁ with certainty
- E₁ or E₂ with the original probabilities
- An average of E₁ and E₂
= After the first measurement the state has collapsed to φ₁, which has the definite energy E₁, so repeating the measurement gives the same result every time.
? What does decoherence explain, and what does it leave open?
- It explains which outcome occurs, but not why interference fades
+ It explains why large objects show no interference, but not why one particular outcome occurs
- It shows that Schrödinger's equation is wrong for large objects
= Interaction with the environment carries away the phase information that interference needs, but the overall state still contains every outcome, so the choice of one result is left unexplained.
:::
