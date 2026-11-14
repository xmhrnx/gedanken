A single fission releases 200 MeV, a tiny amount in everyday terms: about 3 × 10⁻¹¹ joules. To matter, fissions must happen in enormous numbers. The question in early 1939 was whether fission could sustain itself, each event triggering the next.

::: history 1933 · London
Leo Szilard, a Hungarian physicist newly arrived as a refugee, read in *The Times* of 12 September that Ernest Rutherford had dismissed talk of extracting energy from atoms as "moonshine". Waiting at a traffic light on Southampton Row, he realized that if some element absorbed one neutron and emitted two, the process could multiply into a **chain reaction**. He filed a patent on the idea in 1934 and later assigned it to the British Admiralty to keep it secret. He did not yet know which element, if any, would work.
:::

## Neutrons from fission

Fission answered Szilard's question. Heavy nuclei contain proportionally more neutrons than medium-sized ones: uranium-235 has about 1.55 neutrons per proton, while barium and krypton are stable at about 1.3 to 1.5. When uranium splits, the fragments are overloaded with neutrons and shed a few immediately. In March 1939 groups in Paris and New York independently measured this: on average about **2.5 neutrons** are released per U-235 fission.

Not all of them cause another fission. Some escape from the material, and some are captured by nuclei without causing fission. What matters is the **multiplication factor** k: the average number of neutrons from one fission that go on to cause another fission.

- k < 1: each generation is smaller than the last and the reaction dies out (**subcritical**).
- k = 1: the number of fissions per generation stays constant (**critical**). This is a reactor running steadily.
- k > 1: the reaction grows (**supercritical**).

::: model chain-reaction
Each fission releases neutrons that may strike other nuclei. Change k and watch the population die out, hold steady, or explode.
:::

## Exponential growth

If each generation multiplies the number of fissions by k, then after n generations:

::: steps How fast the chain grows
$ N = N_0 \cdot k^n
Start with N₀ fissions. Each generation multiplies the count by k, so after n generations it has been multiplied n times.
---
$ 2^{80} \approx 1.2 \times 10^{24}
Take k = 2 and a single starting fission. Eighty generations give about 10²⁴ fissions, comparable to the 2.6 × 10²⁴ nuclei in a kilogram of U-235.
---
$ t \approx 80 \times 10\,\text{ns} = 0.8\,\mu\text{s}
In pure, dense U-235 a neutron travels a few centimetres in about 10 nanoseconds (10⁻⁸ s) before causing the next fission. Eighty generations take under a microsecond.
---
$ 1 - 2^{-7} \approx 0.99
Because each generation is double the one before, the last seven generations contain 99% of all fissions. Almost all the energy comes out in the final tenth of a microsecond.
:::

That last point explains why so little of the material in a bomb actually fissions. As energy pours out, the material heats to millions of degrees and begins to expand. Expansion spreads the nuclei apart, more neutrons escape, and k falls below 1. The chain is racing against the explosion it creates, and the explosion wins after a short time. A weapon works only if k is well above 1 at the moment the chain starts, so that many generations occur before the material blows itself apart.

::: note
A chain reaction is exponential. The length of a generation, about 10 ns for fast neutrons in a bomb, sets the time scale, and k sets the growth rate. Small changes in k change the outcome enormously, because they are raised to the power of the number of generations.
:::

## U-235 and U-238

Natural uranium is 99.3% uranium-238 and only 0.72% uranium-235. The two behave very differently.

- **U-235** fissions when it absorbs a neutron of any energy, and especially readily with slow ones.
- **U-238** fissions only with fast neutrons above about 1 MeV, and even then it far more often just scatters them, robbing their energy, or captures them without fissioning.

Fission neutrons are born fast, with about 2 MeV on average, but in natural uranium they quickly lose energy in collisions with U-238 and drop below its fission threshold. Each neutron is far more likely to be wasted on U-238 than to find one of the rare U-235 nuclei. The result is k < 1 for fast neutrons in natural uranium, however much of it you pile together.

There are two ways around this. One is to slow the neutrons deliberately with a **moderator**, such as graphite or heavy water, so they are captured mostly by U-235 at low energy. That works in a reactor but is far too slow for an explosion, as the reactors lesson explains. The other is **enrichment**: raising the proportion of U-235. Because U-235 and U-238 are chemically identical and differ in mass by about 1%, separating them was one of the hardest industrial tasks of the war. The Hiroshima bomb used uranium enriched to about 80% U-235.

The second path to a bomb material came from U-238 itself. When it captures a neutron it decays, within days, into **plutonium-239**, a new element that fissions readily, like U-235, and releases about 2.9 neutrons per fission. Plutonium is chemically different from uranium, so it can be separated by chemistry rather than by mass.

::: history 1939 · New York
On 3 March 1939 Szilard and Walter Zinn at Columbia saw flashes on an oscilloscope showing that fission released extra neutrons. "That night," Szilard later wrote, "there was very little doubt in my mind that the world was headed for grief."
:::

::: quiz
? What does the multiplication factor k measure?
- The number of neutrons released per fission
+ The average number of neutrons per fission that cause another fission
- The ratio of U-235 to U-238
= About 2.5 neutrons are released, but some escape or are captured, so k is smaller than that.
? With a generation time of about 10 ns and k = 2, about how long does it take to fission a kilogram of U-235?
- About a second
- About a millisecond
+ About a microsecond
= About 80 generations are needed, and 80 × 10 ns is under a microsecond.
? Why can't a lump of natural uranium sustain a fast chain reaction?
+ Most neutrons are wasted on U-238 and too little U-235 is present
- Natural uranium releases no neutrons when it fissions
- U-235 only fissions with fast neutrons
= Natural uranium is only 0.72% U-235, and U-238 absorbs or slows most neutrons without fissioning.
:::
