Schrödinger's equation says how a wavefunction changes when no one is looking: smoothly, predictably, with no randomness at all. Yet every experiment ends with a definite, random result: a dot on a screen, a click in a detector. The rules of quantum mechanics bridge the two with a separate postulate about measurement, and that postulate is where most of the arguments about the theory live.

## Superposition and outcomes

A particle in a box can be in level 1, level 2, or a **superposition**, a sum of both:

ψ = c₁φ₁ + c₂φ₂

Here φ₁ and φ₂ are the states of definite energy E₁ and E₂, and c₁, c₂ are complex numbers. Such a state does not have a definite energy. It is not secretly in one level with the other a matter of ignorance; the two parts can interfere, as the two slits' waves did.

Measure its energy, though, and you never get anything in between. You get E₁ or E₂, and the Born rule gives the odds.

::: equation The Born rule for outcomes
P(E_1) = |c_1|^2, \qquad P(E_2) = |c_2|^2, \qquad |c_1|^2 + |c_2|^2 = 1
:::

The same rule holds for any quantity: expand ψ in the states that have definite values of that quantity, and the squared size of each coefficient is the probability of that value.

## Collapse

The textbook rule adds one more step. Right after a measurement gives E₂, the state is φ₂. Measure the energy again at once and you get E₂ with certainty. The superposition has been replaced by the single term matching the result. This is called the **collapse** of the wavefunction.

::: steps One measurement, step by step
$ \psi = \sqrt{\tfrac{1}{3}}\,\phi_1 + \sqrt{\tfrac{2}{3}}\,\phi_2
Prepare a superposition. The coefficients are chosen so the probabilities add to 1: ⅓ + ⅔ = 1.
---
$ P(E_1) = \tfrac{1}{3}, \qquad P(E_2) = \tfrac{2}{3}
Measure the energy. The Born rule gives each outcome's probability. Nothing in ψ says which one will happen on a given run.
---
$ \psi \to \phi_2
Suppose the result is E₂. The state is now φ₂ alone: a repeated energy measurement will give E₂ every time.
---
$ \langle E \rangle = \tfrac{1}{3}E_1 + \tfrac{2}{3}E_2
Over many identically prepared particles, the average result is the probability-weighted sum, called the expectation value. With E₂ = 4E₁ it is 3E₁, a value no single measurement ever gives.
:::

This is what happened in the double-slit experiment with a which-slit detector. Recording the path collapses ψ₁ + ψ₂ to one term, the interference term disappears, and so do the fringes.

## The measurement problem

Collapse sits uneasily with the rest of the theory. Schrödinger's equation is deterministic and linear; collapse is random and abrupt. A measuring device is made of atoms, which obey Schrödinger's equation. Treat the device quantum mechanically and, by linearity, a particle in superposition should put the device itself into a superposition of showing E₁ and showing E₂. Nobody sees that. Where, and why, does the rule switch? The theory as usually stated does not say. This is the **measurement problem**.

::: history 1932–1935 · Princeton and Oxford
John von Neumann gave collapse its formal place in 1932, as a process separate from Schrödinger evolution. In 1935 Schrödinger, unhappy with it, imagined a cat in a sealed box whose life depends on whether a radioactive atom has decayed. Applied literally, the rules seem to leave the cat in a superposition of alive and dead until someone opens the box. He meant the example as a reductio, not a description of cats.
:::

Part of the puzzle is now well understood. A large object interacts constantly with its surroundings: air molecules, light, heat. These interactions spread the phase information that interference needs into the environment, where it is lost in practice. This process, **decoherence**, studied from the 1970s by H. Dieter Zeh, Wojciech Zurek and others, explains why big things never show interference and behave classically. It does not by itself explain why one particular outcome occurs.

## Interpretations

Physicists agree on how to use the rules and on every prediction they make. They disagree about what the rules mean. Briefly, the main positions are:

- **Copenhagen** views, associated with Bohr and Heisenberg: ψ is a tool for predicting measurement results, and asking what happens between measurements goes beyond what physics can say.
- **Many worlds** (Hugh Everett, 1957): there is no collapse. Every outcome happens, each in its own branch, and decoherence keeps the branches from interfering.
- **Pilot waves** (Louis de Broglie, 1927; David Bohm, 1952): particles always have definite positions, guided by the wavefunction. Randomness reflects our ignorance of the starting positions.
- **Objective collapse** (Ghirardi, Rimini and Weber, 1986, and others): collapse is a real physical process that adds a small random term to Schrödinger's equation. These theories predict tiny deviations that experiments are searching for.

Other views, such as QBism, treat ψ as an agent's degrees of belief. None of the interpretations so far has an experiment that singles it out.

::: note
Quantum mechanics gives exact probabilities for measurement results, and after a measurement the state matches the result. How a deterministic equation produces a single random outcome is the measurement problem, and it remains open.
:::

::: quiz
? A state is ψ = 0.6φ₁ + 0.8φ₂. What is the probability of measuring E₂?
- 0.8
+ 0.64
- 0.36
= P(E₂) = |c₂|² = 0.8² = 0.64, and 0.36 + 0.64 = 1.
? An energy measurement gives E₁. An immediate second measurement gives
+ E₁ with certainty
- E₁ or E₂ with the original probabilities
- an average of E₁ and E₂
= After the first measurement the state is φ₁, which has the definite energy E₁.
? What does decoherence explain, and what does it leave open?
- It explains which outcome occurs, but not why interference fades
+ It explains why large objects show no interference, but not why one particular outcome occurs
- It shows that Schrödinger's equation is wrong for large objects
= Interaction with the environment suppresses interference, but the state still contains all outcomes.
:::
