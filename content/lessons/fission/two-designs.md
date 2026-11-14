Every fission weapon faces the same basic problem, and it follows directly from the last two lessons. The material must be kept subcritical, with k below 1, until the moment of use, or it would start reacting on its own. Then it must be made strongly supercritical, with k well above 1, and made so quickly that the chain reaction runs through many generations before the material blows itself apart. The two weapons used in 1945 solved this problem in two different ways. This lesson explains the shared difficulty, the two solutions at the level of their basic principle, and why there had to be two: the reason comes down to a single isotope, plutonium-240, and a discovery in 1944 that nearly derailed the project.

## The race against stray neutrons

A chain reaction does not need to be lit deliberately. Any free neutron wandering into a supercritical mass can start one, and there are always a few about. Some nuclei undergo **spontaneous fission**: every so often, without being struck by anything, a heavy nucleus simply splits, releasing neutrons like any other fission. Cosmic rays from space also knock neutrons out of nuclei. In a lump of fissile material, these stray neutrons appear at random moments, at an average rate that depends on the material.

That randomness creates the central danger. While the pieces of material are being brought together, the arrangement passes through an in-between stage where it is only slightly supercritical, with k barely above 1. If a stray neutron happens to start a chain at that moment, the chain grows, but slowly, because each generation is only a little larger than the last. It still produces enough heat to push the material apart, and it does so before the assembly is complete, so the material never reaches the strongly supercritical state it was meant to. This is called **predetonation**, and the result is a **fizzle**: an explosion far smaller than intended, though still a dangerous one.

::: equation Growth of the chain after n generations
N = N_0 \cdot k^n
:::

The equation from the chain-reaction lesson shows why the starting moment matters so much. The number of generations available before the material disperses is limited, so what counts is how much the chain multiplies in each one. With k only slightly above 1, kⁿ grows slowly, and the material disperses after relatively few effective doublings. With k well above 1, the same number of generations multiplies the chain enormously. So the yield depends heavily on how supercritical the material is at the instant the chain begins.

::: model chain-reaction
Compare a value of k just above 1 with one well above 1, over the same number of generations. The first is a chain started too early, during assembly; the second is one started at the right moment.
:::

The aim of any design, then, is to pass through the marginal stage faster than a stray neutron is likely to turn up, and then to start the chain at the moment of greatest supercriticality. How fast is fast enough depends on how often stray neutrons appear, and that can be put in a simple statistical form.

::: steps Why the stray-neutron rate sets the speed of assembly
$ S\,T
Suppose stray neutrons appear in the material at random, at an average rate S per unit time, and suppose the dangerous in-between stage lasts a time T. Then S times T is the average number of stray neutrons expected during that stage. Neither number needs to be known exactly to follow the argument; what matters is their product.
---
$ P(\text{no stray neutron}) = e^{-S T}
Events that occur independently at random, at a steady average rate, follow what is called a Poisson distribution, the same statistics as radioactive decay. Its key result is that the chance of seeing none at all in a time T is e raised to the power −ST. This is exactly the form of the decay law, and for the same reason: each short moment carries the same small, independent chance.
---
$ e^{-S T} \approx 1 \text{ only when } S T \ll 1
For the assembly to succeed reliably, this probability must be close to 1, which happens only when ST is much smaller than 1. If ST is around 1 or larger, a predetonation becomes likely rather than unlucky. So the speed of assembly needed is set by the stray-neutron rate: the time T must be short compared with 1/S, the average waiting time between stray neutrons.
---
$ S \uparrow \;\Rightarrow\; T \downarrow
The consequence is the whole story of this lesson. A material that produces few stray neutrons, with small S, tolerates a slow assembly. A material that produces many, with large S, demands a much faster one, and if S rises enough, a method that was fast enough becomes hopeless.
:::

::: note
How fast the assembly must be depends on how many stray neutrons the material produces. A material with few spontaneous fissions tolerates a slow assembly; a material with many requires a fast one.
:::

## Gun-type: Little Boy

The simplest approach is to keep two pieces of material apart, each too small to be critical on its own, and then bring them together quickly. In a **gun-type** weapon, a conventional explosive charge drives one subcritical piece of uranium down a barrel into another. Neither piece alone can sustain a chain; together they form a supercritical mass.

By the standards of a chain reaction, which runs its course in about a microsecond, this assembly is slow. A projectile, however fast, takes a long time to travel any distance compared with the nanoseconds of a neutron generation, so the arrangement spends a comparatively long time in the in-between stage. It works for uranium-235 because U-235 has a very low rate of spontaneous fission. Its S is small, so even a slow assembly is unlikely to meet a stray neutron at the wrong moment.

The Hiroshima bomb, **Little Boy**, used this method. It contained about 64 kg of uranium, enriched to an average of about 80% U-235. Only about a kilogram of it fissioned, a little over 1% of the total; the rest was scattered by the explosion before it could react. By later standards that was very inefficient, but the principle was so simple that its designers were confident of it without a test, and it was **never tested before it was used**. Its yield, about 15 kilotons, had to be worked out afterwards from measurements and surveys of the damage.

## Why plutonium could not use a gun

Plutonium was the project's other material, and the original plan was to use it in a gun as well. This weapon, codenamed Thin Man, was designed on the assumption that plutonium from a reactor would behave like the tiny samples previously made in particle accelerators, whose rate of spontaneous fission was low.

::: history 1944 · Los Alamos
In the spring and summer of 1944, Emilio Segrè's group at Los Alamos measured the spontaneous fission rate of the first plutonium made in a reactor, a pilot reactor at Oak Ridge. It was far higher than in the accelerator-made samples. The cause was **plutonium-240**. In a reactor, plutonium-239 is bathed in neutrons for as long as the fuel stays in the reactor, and some of its nuclei capture a neutron without fissioning, becoming plutonium-240. Accelerator samples had been too small and too briefly irradiated for this to happen much. Plutonium-240 fissions spontaneously at a high rate, and even a modest admixture fills the material with stray neutrons. The calculation was clear: a gun assembly of reactor plutonium would almost certainly predetonate. In July 1944 Thin Man was abandoned, and Oppenheimer reorganised the Los Alamos laboratory around a much harder idea.
:::

The obvious remedy, removing the plutonium-240, was not available. It is the same element as plutonium-239, so it behaves identically in every chemical process, and separating isotopes of plutonium, which differ in mass even less in proportion than those of uranium, was out of the question in wartime. In the language of the steps above, S could not be reduced, so T had to be. The assembly had to be made far faster than any gun could manage.

## Implosion: Fat Man

The alternative relied on the density result from the critical-mass lesson: critical mass falls as the square of the density rises. A piece of material that is subcritical at its normal density becomes supercritical if it is compressed, without anything being added. In an **implosion** weapon, a subcritical sphere of plutonium is surrounded by conventional high explosive. When the explosive is detonated, it drives a shock wave inward from all sides that compresses the sphere. Its density rises, its critical mass falls below the amount actually present, and it becomes supercritical.

Implosion is much faster than a gun, and the reason is geometric. A gun has to move a piece of material across a gap of some length. Implosion moves every part of the sphere only a short distance inward, and at very high speed, so the whole passage through the in-between stage is over far sooner. That brought the time T down enough to outrun the stray neutrons from plutonium-240. Compression also achieves a greater degree of supercriticality than simply joining two pieces, so more of the material fissions before the chain stops, which made the design more efficient.

The difficulty was symmetry. If the inward push is even slightly uneven, the material is squeezed more on one side than another and squirts out sideways instead of compressing. Making the compression smooth and uniform was the central technical problem at Los Alamos from 1944 onward. It was uncertain enough that this design, unlike the gun, was tested at Trinity before it was used.

The Nagasaki bomb, **Fat Man**, contained about 6.2 kg of plutonium, about a tenth of the mass of uranium in Little Boy, and produced about 21 kilotons. Roughly a kilogram of its plutonium fissioned, about 15 to 20% of the material, more than ten times the efficiency of the gun.

::: history 1943 · Los Alamos
Implosion was first proposed in the spring of 1943 by Seth Neddermeyer, a young experimental physicist. Many senior scientists were sceptical: it seemed a crude and unreliable way to assemble material compared with the straightforward gun. Neddermeyer was allowed to pursue it with a small group, and progress was slow. It was taken up seriously after John von Neumann, visiting Los Alamos in late 1943, showed that a converging shock could compress the material much more strongly than had been assumed, which would make the method both faster and more efficient. Then, in 1944, the plutonium-240 problem turned implosion from an interesting option into the only route to a plutonium weapon, and the laboratory staked its effort on it.
:::

Implosion became the basis of almost all later fission weapons, because it uses less material and uses it more efficiently. The gun method was simpler but wasteful, and it works only with uranium, the one weapon material whose low spontaneous fission rate forgives a slow assembly.

::: quiz
? What is predetonation?
- Detonating the weapon too high above the ground
+ A stray neutron starting the chain before assembly is complete, giving a fizzle
- The conventional explosive failing to fire
= A chain started while the material is only slightly supercritical grows slowly but still heats the material and blows it apart early. The assembly never reaches full supercriticality, so the yield is small.
? Why could plutonium not be used in a gun-type weapon?
- Plutonium cannot sustain a chain reaction
- Plutonium is too dense to fit in a barrel
+ Pu-240 fissions spontaneously so often that a gun is too slow to avoid predetonation
= Reactor plutonium always contains some Pu-240, whose spontaneous fissions supply stray neutrons at a high rate. A gun assembly lasts long enough that one would very likely start the chain too early.
? How does implosion make a subcritical sphere supercritical?
+ It compresses the sphere so its density rises and its critical mass falls
- It adds a second piece of plutonium
- It converts U-238 into plutonium during the explosion
= Critical mass scales as 1/ρ², so squeezing the same material to a higher density lowers its critical mass below the amount present. Nothing is added; the material simply becomes supercritical.
:::
