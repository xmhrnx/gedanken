A hot cup of coffee cools to room temperature; nobody has seen one warm up by drawing heat from the room. Perfume spreads through a room and never gathers back into the bottle. Neither reverse process breaks energy conservation. Something else forbids them, and that something is the second law of thermodynamics. Its central quantity is entropy.

## The second law, in words

Rudolf Clausius stated it in 1850: heat does not flow spontaneously from a colder body to a hotter one. A refrigerator can push heat uphill, but only by using work, and it warms its surroundings more than it cools its interior. A year later William Thomson (later Lord Kelvin) gave an equivalent form: no engine working in a cycle can turn heat from a single reservoir entirely into work. Each statement can be proved from the other, and the previous lesson used them to show that no engine beats Carnot's.

## Clausius's entropy

Carnot's reversible engine has a hidden regularity. It takes Q_H in at T_H and gives Q_C out at T_C, with Q_C/Q_H = T_C/T_H. Rearranged, Q_H/T_H = Q_C/T_C: the quantity Q/T that enters the engine equals the quantity that leaves it. Clausius saw that Q/T measures something that a reversible process conserves, and he named it.

::: equation Clausius's definition of entropy change
\Delta S = \int \frac{\mathrm{d}Q_{\text{rev}}}{T}
:::

When heat dQ enters a system reversibly at temperature T, the system's *entropy* S rises by dQ/T. Like internal energy, entropy depends only on the state, so to find the change in an irreversible process you calculate it along any reversible path between the same start and end. Its unit is J/K.

Now let heat flow irreversibly, straight from a hot body to a cold one.

::: steps Heat flowing from hot to cold creates entropy
$ \Delta S_H = -\frac{Q}{T_H}
The hot body, large enough that its temperature barely changes, loses heat Q at temperature T_H.
---
$ \Delta S_C = +\frac{Q}{T_C}
The cold body gains the same Q at its own temperature T_C. Energy is conserved: what one loses the other gains.
---
$ \Delta S_{\text{total}} = Q\left(\frac{1}{T_C} - \frac{1}{T_H}\right) > 0
Add them. Because T_C < T_H, 1/T_C is larger than 1/T_H, so the total is positive. Heat flowing the other way would make it negative.
:::

Moving 1000 J from a body at 400 K to one at 300 K raises the total entropy by 1000 × (1/300 − 1/400) ≈ 0.83 J/K. This gives the second law its general form: **in an isolated system, total entropy never decreases.** Reversible processes leave it unchanged; every real process increases it.

Entropy can fall locally. A freezer lowers the entropy of the water it turns to ice, and living things build ordered structures, but each exports more entropy to its surroundings, as heat, than it removes.

## Boltzmann's entropy

Clausius's definition says how to measure entropy, not what it is. Ludwig Boltzmann answered that with molecules. A *macrostate* is what you can measure: pressure, volume, temperature. A *microstate* is a complete specification of every molecule's position and velocity. Many microstates look the same from outside, and W counts how many correspond to a given macrostate.

::: equation Boltzmann's entropy
S = k \ln W
:::

Here k = 1.381 × 10⁻²³ J/K is Boltzmann's constant. The logarithm makes entropy add up properly: two separate systems have W₁W₂ joint microstates, and ln(W₁W₂) = ln W₁ + ln W₂.

The second law becomes a statement about probability. Systems wander at random among their microstates, so they end up in the macrostates with overwhelmingly the most microstates. Whether the two definitions agree can be tested on one case.

::: steps A gas doubling its volume, counted two ways
$ W_2 = 2^N W_1
A box holds N molecules of an ideal gas on one side of a partition. Remove the partition and the gas fills twice the volume at the same temperature. Each molecule now has twice as many places to be, and the choices are independent, so the count of microstates grows by 2 for each molecule.
---
$ \Delta S = k\ln W_2 - k\ln W_1 = k\ln 2^N = Nk\ln 2
Boltzmann's formula, using ln(ab) = ln a + ln b and ln 2^N = N ln 2.
---
$ Q = NkT\ln 2
Clausius's formula needs a reversible path to the same end state: let the gas push a piston slowly outward to twice its volume at constant temperature. From the previous lesson, the heat absorbed is NkT ln(V₂/V₁).
---
$ \Delta S = \frac{Q}{T} = Nk\ln 2
Divide by T. The two definitions give the same answer.
:::

The reverse, all the molecules gathering back into one half, is not forbidden; it is just improbable. Each molecule is in the left half with probability ½, so the chance for all N at once is (½)^N. For 100 molecules that is about 10⁻³⁰. For the roughly 10²² molecules in a litre of air, it will never be seen.

::: note
Entropy counts the microscopic arrangements consistent with what you observe, S = k ln W. Heat flows from hot to cold, and gases spread out, because the result has vastly more arrangements than the start.
:::

## The arrow of time

The laws governing individual molecular collisions work equally well forwards and backwards. The direction of time in everyday life comes from entropy: the past is the direction of lower entropy. That in turn requires the universe to have started in a state of very low entropy, which remains a question for cosmology.

::: history 1865–1877 · Zurich and Graz
Clausius coined "entropy" in 1865 from the Greek for "transformation", choosing a word that sounded like "energy". He summed up: the energy of the universe is constant; its entropy tends to a maximum. In 1877 Boltzmann, in Graz, connected entropy to the counting of molecular arrangements. Many physicists then doubted that atoms existed at all, and objections to his statistical reading of the second law dogged him for years. He died in 1906, two years before Jean Perrin's experiments on Brownian motion settled the existence of atoms. S = k log W is carved on his tombstone in Vienna, though Boltzmann never wrote it in quite that form: Max Planck did, around 1900, and was the first to calculate k.
:::

::: quiz
? 600 J of heat flows from a body at 600 K to one at 300 K. The total entropy change is
- zero, since energy is conserved
+ +1 J/K
- −1 J/K
= ΔS = 600/300 − 600/600 = 2 − 1 = +1 J/K. Spontaneous heat flow always increases total entropy.
? According to Boltzmann, why does a gas spread to fill its container?
- A force pushes the molecules apart
+ Spread-out arrangements vastly outnumber gathered ones, so random motion almost certainly ends there
- Entropy is a fluid that flows into empty space
= The spread-out macrostate has 2^N times more microstates; the gathered state is allowed but overwhelmingly improbable.
? A refrigerator makes ice, lowering the water's entropy. How is this consistent with the second law?
- The second law does not apply to machines
- Ice has more microstates than water
+ The refrigerator releases heat to the room, raising the room's entropy by more than the water loses
= The law forbids a decrease of total entropy in an isolated system, not a local decrease paid for elsewhere.
:::
