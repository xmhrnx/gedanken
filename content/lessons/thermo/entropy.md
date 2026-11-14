A hot cup of coffee left on a table cools to room temperature. Nobody has ever seen the reverse: a lukewarm cup that grows hotter by drawing heat out of the cooler room. A drop of ink spreads through a glass of water and never gathers itself back into a drop. What is odd is that neither reverse process would break the first law. The energy the coffee would gain could come from the room, and energy would be conserved. Something else rules these processes out, and it is just as reliable as energy conservation. That something is the second law of thermodynamics, and the quantity at its centre is **entropy**. This lesson follows how entropy was first defined, as a way of measuring heat, and how it was later understood, as a way of counting.

## What does the second law say?

The law was first put in plain words. In 1850 Rudolf Clausius stated it like this: heat does not flow by itself from a colder body to a hotter one. A refrigerator does move heat from its cold interior to the warmer kitchen, but only because it is driven by work from an electric motor, and it releases more heat into the kitchen than it removes from inside. Left alone, heat runs downhill in temperature, never up.

A year later William Thomson, later Lord Kelvin, gave a version that sounds different: no engine working in a cycle can take heat from a single reservoir and turn all of it into work. It must always dump some heat at a lower temperature. The two statements are equivalent. If you had a machine that broke one of them, you could connect it to an ordinary engine or refrigerator and build a machine that broke the other. The previous lesson used exactly this kind of argument to show that no engine can beat Carnot's efficiency.

Both statements are about what never happens. Clausius wanted something stronger: a quantity that could be calculated, that would show which processes are possible and which are not.

## Where did Clausius find a new quantity?

He found it hidden inside Carnot's ideal engine. A reversible engine takes in heat Q_H at temperature T_H and gives out Q_C at T_C, and the previous lesson showed that Q_C/Q_H = T_C/T_H. Rearranging that gives Q_H/T_H = Q_C/T_C. Heat itself is not conserved by the engine, since some of it becomes work. But the quantity heat divided by temperature is: the amount of Q/T that enters at the hot side equals the amount that leaves at the cold side. Clausius saw that this ratio measures something real that a reversible process carries through unchanged, and he gave it a name.

::: equation Clausius's definition of entropy change
\Delta S = \int \frac{\mathrm{d}Q_{\text{rev}}}{T}
:::

In words: when a small amount of heat dQ enters a system reversibly, meaning by a process gentle enough that it could be run backwards, at a temperature T, the system's entropy S rises by dQ/T. Adding up these pieces gives the total change. The unit is joules per kelvin.

There is an important property hiding here. Entropy, like internal energy, is a **state function**: its value depends only on the system's present condition, not on its history. That means you can find the entropy change in any process, including a violent and irreversible one, by imagining some gentle reversible path between the same starting and ending states and calculating along that path instead. The answer is the same whichever path you choose.

## What happens to entropy when heat flows from hot to cold?

Now Clausius's quantity can be put to work on the simplest irreversible process there is: heat flowing directly from a hot body to a cold one, as when the coffee cools.

::: steps Heat flowing from hot to cold creates entropy
$ \Delta S_H = -\frac{Q}{T_H}
Take two bodies large enough that their temperatures barely change when a modest amount of heat Q passes between them. The hot body, at T_H, loses heat Q. Since its temperature stays essentially fixed, its loss could equally have happened by a slow, reversible process, so its entropy falls by Q/T_H.
---
$ \Delta S_C = +\frac{Q}{T_C}
The cold body, at T_C, gains exactly the same heat Q, because energy is conserved and nothing else is involved. By the same reasoning its entropy rises by Q/T_C. The heat is the same in both cases, but it is divided by a different temperature.
---
$ \Delta S_{\text{total}} = Q\left(\frac{1}{T_C} - \frac{1}{T_H}\right) > 0
Add the two changes to get the total for the pair. Because T_C is smaller than T_H, 1/T_C is larger than 1/T_H, so the gain on the cold side outweighs the loss on the hot side and the total is positive. If the heat went the other way, from cold to hot, every sign would flip and the total would be negative.
:::

A number makes this concrete. Move 1000 J from a body at 400 K to one at 300 K. The cold body gains 1000/300 ≈ 3.33 J/K, the hot body loses 1000/400 = 2.50 J/K, and the total entropy rises by about 0.83 J/K. The process that happens is the one that increases total entropy; the one that never happens is the one that would decrease it.

This turned out to be general, and it gives the second law its most useful form: **in an isolated system, total entropy never decreases.** An isolated system is one that exchanges no energy with anything outside it. A perfectly reversible process inside it leaves the total entropy unchanged, and every real process, which always involves some friction, some heat flowing across a temperature difference, or some mixing, increases it.

That statement is about totals. Entropy can fall in one part of a system if it rises by more somewhere else. A freezer lowers the entropy of the water it turns into ice, and living things build highly ordered structures out of disordered raw materials. Neither breaks the law, because each releases heat into its surroundings, and the entropy that heat carries away exceeds what was removed locally.

## What is entropy, physically?

Clausius's definition tells you how to measure entropy but not what it is. Why should Q/T, of all things, always increase? The answer came from Ludwig Boltzmann, who approached the question through molecules.

Begin with a distinction. A **macrostate** is a description of a system in terms of what you can measure from outside: its pressure, volume and temperature. A **microstate** is a complete description of every molecule: exactly where each one is and how it is moving. A single macrostate can be produced by an immense number of different microstates, because the molecules can be shuffled around in countless ways without changing anything you can measure. Write W for the number of microstates that correspond to a given macrostate. (This W has nothing to do with work; the letter comes from the German *Wahrscheinlichkeit*, probability.)

::: equation Boltzmann's entropy
S = k \ln W
:::

Here k ≈ 1.381 × 10⁻²³ J/K is Boltzmann's constant, the same one that appears in PV = NkT. The logarithm is there for a good reason. Take two separate systems, one with W₁ microstates and one with W₂. Every microstate of the first can be paired with every microstate of the second, so together they have W₁ × W₂ microstates. But entropy, like energy or mass, should simply add when you put two systems side by side. The logarithm turns the multiplication into addition, since ln(W₁W₂) = ln W₁ + ln W₂.

With this definition the second law becomes a statement about probability. Molecules are constantly colliding and moving, so a system wanders at random from one microstate to another. It will spend almost all its time in the macrostates that have the most microstates, simply because there are so many more of them. Entropy increases because a system moving at random almost always drifts from macrostates with few arrangements into macrostates with vastly more.

## Do the two definitions agree?

Clausius defined entropy through heat and temperature; Boltzmann defined it by counting. They look nothing alike, so it is worth checking that they give the same answer in a case where both can be worked out.

::: steps A gas doubling its volume, counted two ways
$ W_2 = 2^N W_1
A box holds N molecules of an ideal gas, all confined to one half by a partition. Remove the partition and the gas spreads to fill the whole box at the same temperature. Because the temperature is unchanged, the molecules' speeds are distributed just as before, so the only difference is in where they can be. Each molecule now has twice as much room, and since the molecules do not interact, each one's position can be chosen independently. So the number of microstates doubles once for every molecule: N doublings in all.
---
$ \Delta S = k\ln W_2 - k\ln W_1 = k\ln 2^N = Nk\ln 2
Apply Boltzmann's formula to the states before and after and subtract. The difference of two logarithms is the logarithm of the ratio, which is 2^N, and the logarithm of a power brings the exponent down in front: ln 2^N = N ln 2. That is Boltzmann's prediction for the entropy gained.
---
$ Q = NkT\ln 2
Clausius's formula needs heat flowing reversibly, but the gas rushing into empty space involved no heat at all. So use the fact that entropy is a state function, and find a gentle path to the same end state: let the gas push a piston slowly outward from the original volume to twice that volume, while in contact with a reservoir that keeps its temperature fixed. The previous lesson showed that such an isothermal expansion absorbs heat nRT ln(V₂/V₁), which is NkT ln 2 here, since nR = Nk and V₂/V₁ = 2.
---
$ \Delta S = \frac{Q}{T} = Nk\ln 2
The heat entered at the single temperature T, so Clausius's entropy change is just Q divided by T. It matches Boltzmann's result exactly. Two definitions arrived at from completely different directions, one from steam engines and one from counting molecular arrangements, describe the same quantity.
:::

This example also shows what the second law really forbids. The reverse process, every molecule drifting back into the original half by chance, is not impossible; it is just extraordinarily unlikely. Each molecule spends half its time on the left, so the chance that all N are on the left at the same moment is (½)^N. For 100 molecules, that is about 10⁻³⁰: you could watch for the age of the universe many times over and never see it. A litre of air contains around 10²² molecules, and for that many the chance is so small that "never" is the only honest word. The second law is a law of statistics, but for everyday amounts of matter its predictions are as certain as anything in physics.

::: note
Entropy counts the microscopic arrangements consistent with what you observe, S = k ln W. Heat flows from hot to cold, and gases spread out, because the end result has vastly more arrangements than the start, so random molecular motion almost always leads there.
:::

## Why does time have a direction?

There is a deep puzzle here. The laws governing individual molecular collisions work equally well forwards and backwards: a film of two molecules bouncing off each other looks just as reasonable run in reverse. Yet a film of coffee cooling, or ink spreading, looks obviously wrong when reversed. The difference between past and future in everyday life comes not from the laws of motion but from entropy. The future is the direction in which entropy increases.

That answer raises another question. If entropy always increases, it must have been lower in the past, and lower still before that. The direction of time we experience therefore depends on the universe having begun in a state of remarkably low entropy. Why it did is not explained by thermodynamics, and it remains an open question in cosmology.

::: history 1865–1877 · Zurich and Graz
Clausius, by then a professor in Zurich, coined the word "entropy" in 1865 from the Greek for "transformation", deliberately choosing a word that sounded like "energy" because he saw the two as partners. He summed up thermodynamics in two sentences: the energy of the universe is constant, and its entropy tends to a maximum. What entropy was, he did not say.

In 1877 Ludwig Boltzmann, working in Graz, gave the answer by connecting entropy to the number of ways molecules can be arranged. It was a bold move, because many leading physicists and chemists of the day doubted that atoms existed at all, and saw no reason to believe a law of nature could rest on probabilities. Critics pointed out that molecular motion is reversible, so how could it produce an irreversible law? Boltzmann's answer, that irreversibility is overwhelming probability rather than strict necessity, took decades to be widely accepted, and the long arguments wore on him. He died in 1906. Two years later Jean Perrin's experiments on Brownian motion, the jittering of tiny particles in water as they are struck by molecules, convinced nearly everyone that atoms were real. The formula S = k log W is carved on Boltzmann's tombstone in Vienna, although he never wrote it in quite that form: Max Planck did, around 1900, and was the first to calculate the value of k.
:::

::: quiz
? 600 J of heat flows from a body at 600 K to one at 300 K. What is the total entropy change?
- Zero, since energy is conserved
+ +1 J/K
- −1 J/K
= The cold body gains 600/300 = 2 J/K and the hot body loses 600/600 = 1 J/K, so the total rises by 1 J/K. Spontaneous heat flow from hot to cold always increases total entropy.
? According to Boltzmann, why does a gas spread out to fill its container?
- A force pushes the molecules apart
+ Spread-out arrangements vastly outnumber gathered ones, so random motion almost certainly ends there
- Entropy is a fluid that flows into empty space
= The spread-out macrostate has 2^N times as many microstates as the gathered one. Gathering back into one half is allowed, but for any real number of molecules it is overwhelmingly improbable.
? A refrigerator makes ice, lowering the water's entropy. How is this consistent with the second law?
- The second law does not apply to machines
- Ice has more microstates than water
+ The refrigerator releases heat to the room, raising the room's entropy by more than the water loses
= The second law forbids a decrease in the total entropy of an isolated system. A local decrease is allowed as long as a larger increase happens somewhere else, which here is the warm kitchen.
:::
