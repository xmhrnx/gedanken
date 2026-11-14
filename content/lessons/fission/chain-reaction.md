A single fission releases about 200 MeV. By nuclear standards that is a great deal, but in everyday terms it is tiny: about 3 × 10⁻¹¹ joules, far less than the energy needed to lift a grain of sand by a millimetre. For fission to matter, it has to happen in enormous numbers, around 10²⁴ nuclei for a kilogram of uranium. No laboratory could split them one at a time. The question in early 1939 was whether fission could do the work itself, each split triggering the next. This lesson explains how such a chain reaction works, why it grows so explosively when it grows at all, and why ordinary uranium cannot sustain one.

## An idea before its time

::: history 1933 · London
Leo Szilard, a Hungarian physicist who had fled Germany when the Nazis came to power, was living in a London hotel in September 1933 when he read in *The Times* of 12 September a report of a speech by Ernest Rutherford. Rutherford, the greatest experimental physicist of the day, had dismissed the idea of extracting useful energy from atoms as "moonshine". Splitting nuclei with accelerated particles used far more energy than it released, because almost all the projectiles missed. Szilard found the dismissal irritating, and as he waited for a traffic light to change on Southampton Row, he saw a way around it. Suppose there were an element that, when struck by one neutron, released two. Neutrons are not repelled by nuclei, so each of the two could strike another nucleus and release two more, and so on. The process would feed itself, and the number of reactions would multiply. He called it a **chain reaction**. He filed a patent on the idea in 1934 and later assigned it to the British Admiralty to keep it secret. What he did not have was the element. He tried beryllium and indium, without success.
:::

## Why fission releases neutrons

Five years later, fission supplied what Szilard had been missing, and the reason lies in the balance of protons and neutrons in a nucleus. Heavy nuclei need proportionally more neutrons than light ones. Their many protons repel each other strongly, and extra neutrons add strong-force attraction without adding repulsion. Uranium-235 has 143 neutrons for its 92 protons, about 1.55 neutrons per proton. Middle-sized nuclei such as barium and krypton are stable with fewer, about 1.3 to 1.5 neutrons per proton.

So when a uranium nucleus splits, the two fragments inherit the neutron-rich mix of the parent and find themselves overloaded. The quickest way to shed the surplus is to throw off a few neutrons at the moment of the split. In March 1939 groups in Paris and New York measured this independently, and found that fission does release spare neutrons. Later measurements settled the average at about **2.5 neutrons** per fission of U-235. Szilard's condition, more than one neutron out for each neutron in, was met.

Releasing neutrons is not the same as sustaining a chain, though. Not every neutron goes on to cause another fission. Some escape from the edge of the material. Some are absorbed by a nucleus that does not split, either a uranium nucleus that simply keeps the neutron or some impurity. What decides the fate of the chain is the **multiplication factor**, k: the average number of neutrons from one fission that go on to cause another fission. It is always smaller than the 2.5 released, and everything depends on whether it is smaller or larger than 1.

If k is less than 1, each generation of fissions is smaller than the one before, and the chain dies out after a few steps. The material is called **subcritical**. If k is exactly 1, each fission leads on average to exactly one more, and the rate of fission stays constant from one generation to the next. The material is **critical**, and that is the state of a reactor running steadily. If k is greater than 1, each generation is larger than the last, and the reaction grows. The material is **supercritical**.

::: model chain-reaction
Each fission releases neutrons that may strike other nuclei. Set k below 1, at 1, and above 1, and watch the population die out, hold steady, or run away.
:::

## How fast the chain grows

When k is above 1, the growth is exponential: the number of fissions is multiplied by the same factor in each generation. That kind of growth starts unremarkably and then becomes overwhelming, and working through the numbers shows how quickly.

::: steps How fast the chain grows
$ N = N_0 \cdot k^n
Start with N₀ fissions in the first generation. Each of them leads on average to k fissions in the next, so the second generation has N₀k, the third N₀k², and so on. After n generations the count has been multiplied by k a total of n times.
---
$ 2^{80} \approx 1.2 \times 10^{24}
To see what this means, take k = 2 and start from a single fission. After 10 generations there are about a thousand fissions per generation, after 20 about a million, and after 80 about 10²⁴. That is comparable to the 2.6 × 10²⁴ nuclei in a kilogram of U-235. Eighty doublings are enough to go from one nucleus to all of them.
---
$ t \approx 80 \times 10\,\text{ns} = 0.8\,\mu\text{s}
Now ask how long a generation takes. A neutron released in fission moves at about 2 × 10⁷ metres per second, roughly 7% of the speed of light. In dense, pure U-235 it zigzags between nuclei and causes the next fission after about 10 nanoseconds, or 10⁻⁸ s. Eighty generations then take under a microsecond, a millionth of a second.
---
$ 1 + 2 + 4 + \dots + 2^{n} = 2^{n+1} - 1
The total number of fissions up to generation n is the sum of all the generations, and when each is double the last, that sum is just under twice the final generation. In other words, the last generation alone contains about half of all the fissions that have happened, the one before it a quarter, and so on.
---
$ 1 - 2^{-7} \approx 0.99
It follows that the last seven generations together contain all but 1/2⁷, or 1/128, of the fissions, more than 99%. Since seven generations take about 70 nanoseconds, almost all the energy comes out in the final tenth of a microsecond. The earlier generations are like a fuse; the release is concentrated at the very end.
:::

That last result explains a fact that surprises most people: in a nuclear weapon, only a small part of the material ever fissions. As energy pours out in those final generations, the material heats to tens of millions of degrees and starts to expand. Expansion spreads the nuclei apart, so a neutron travels farther between them and is more likely to escape. Then k falls, first toward 1 and then below it, and the chain stops. The chain reaction is racing against the explosion it creates, and the explosion always wins after a very short time. Because the energy is concentrated in the last few generations, every generation that can be squeezed in before the material blows apart roughly doubles the yield. That is why a weapon needs k well above 1 at the moment the chain begins: the faster each generation multiplies, the more of them fit into the time available.

::: note
A chain reaction is exponential. The length of a generation, about 10 nanoseconds for fast neutrons in dense fissile material, sets the time scale, and k sets the rate of growth. Small changes in k change the outcome enormously, because k is raised to the power of the number of generations.
:::

## Why ordinary uranium will not do

There was a large catch. Natural uranium, as it comes out of the ground, is 99.3% uranium-238 and only 0.72% uranium-235. The two isotopes are chemically identical but behave very differently with neutrons.

Uranium-235 fissions when it absorbs a neutron of almost any energy, and it is especially likely to absorb slow ones. Uranium-238 fissions only when struck by a fast neutron with more than about 1 MeV of energy. Even then, a fast neutron hitting U-238 is far more likely to bounce off it, losing much of its energy, or to be captured without causing fission at all.

Fission neutrons are born fast, with about 2 MeV on average, so at first glance they could split U-238. But in natural uranium each neutron meets mostly U-238 nuclei, and a collision or two with them knocks it below the 1 MeV threshold. From then on it can no longer split U-238; it can only find one of the rare U-235 nuclei, and it is much more likely to be captured by U-238 first. The result is that in natural uranium, k for fast neutrons is well below 1. That is not a matter of quantity: piling up more natural uranium reduces the fraction of neutrons that escape, but it does nothing about the ones wasted inside, and a fast chain never gets going.

There were two ways around this, and both were pursued. One is to slow the neutrons down deliberately, by mixing the uranium with a **moderator**, a material such as graphite or heavy water whose light nuclei absorb energy from neutrons in collisions without capturing many of them. Very slow neutrons are hundreds of times more likely to be absorbed by U-235 than fast ones, enough to tip the balance in its favour even when it is rare. That approach works in a reactor, but a moderated chain is far too slow for an explosion, as the lesson on reactors explains.

The other way is **enrichment**: raising the proportion of U-235 in the uranium, so that a fast neutron is likely to meet a U-235 nucleus before it is wasted. Because the two isotopes are chemically identical, no chemical process can separate them; the only handle is a difference in mass of about 1%. Separating them on an industrial scale was one of the hardest tasks of the war, consuming vast factories and enormous amounts of electricity. The Hiroshima bomb used uranium enriched to an average of about 80% U-235.

A second weapon material came, unexpectedly, from U-238 itself. When a U-238 nucleus captures a neutron without fissioning, it becomes U-239, which undergoes two beta decays within a few days and turns into **plutonium-239**, an element not found in nature in usable amounts. Plutonium-239 fissions readily, like U-235, and releases about 2.9 neutrons per fission. Crucially, it is a different element from uranium, so it can be separated from it by ordinary chemistry rather than by the slow, costly route of separating isotopes by mass. A reactor running on natural uranium, where U-238 captures neutrons all the time, makes plutonium as a by-product.

::: history 3 March 1939 · New York
At Columbia University, Szilard and Walter Zinn set up an experiment to see whether fission released neutrons. They bombarded uranium and watched a screen that would show flashes if fast neutrons appeared. The flashes came. Szilard, who had imagined the chain reaction six years earlier and knew what it could mean in a world sliding toward war, later wrote: "That night, there was very little doubt in my mind that the world was headed for grief."
:::

::: quiz
? What does the multiplication factor k measure?
- The number of neutrons released per fission
+ The average number of neutrons per fission that cause another fission
- The ratio of U-235 to U-238
= About 2.5 neutrons are released per fission, but some escape and others are captured without causing fission. k counts only the ones that go on to split another nucleus, so it is always smaller than 2.5.
? With a generation time of about 10 ns and k = 2, about how long does it take to fission a kilogram of U-235?
- About a second
- About a millisecond
+ About a microsecond
= Doubling from one fission to about 10²⁴ takes about 80 generations, and 80 × 10 ns is 0.8 microseconds.
? Why can't a lump of natural uranium sustain a fast chain reaction?
+ Most neutrons are wasted on U-238 and too little U-235 is present
- Natural uranium releases no neutrons when it fissions
- U-235 only fissions with fast neutrons
= Natural uranium is only 0.72% U-235. Fast neutrons are quickly slowed below U-238's fission threshold and are then far more likely to be captured by U-238 than to find a U-235 nucleus, so k stays below 1 however much uranium there is.
:::
