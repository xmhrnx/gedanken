Every fission weapon faces the same problem. The material must be kept subcritical until the moment of use, then made strongly supercritical so fast that the chain reaction runs through many generations before the material blows itself apart. The 1945 weapons solved this problem in two different ways, and the reason there were two comes down to a single isotope, plutonium-240.

## The race against predetonation

A chain reaction can be started by any stray neutron. There are always a few around: some nuclei fission spontaneously, without being struck, and cosmic rays produce neutrons too.

While the material is being brought together, it passes through a stage where it is only slightly supercritical. If a stray neutron starts a chain at that moment, the reaction grows slowly (k barely above 1), heats the material, and pushes it apart before the assembly is complete. This is **predetonation**, and the result is a **fizzle**: an explosion far smaller than intended, though still dangerous.

::: equation Generations available before the material disperses
N = N_0 \cdot k^n
:::

The outcome depends on k at the time the chain gets going. With k only slightly above 1, kⁿ grows slowly and the material disperses after relatively few doublings. The design goal is therefore to pass through the marginal stage faster than a stray neutron is likely to appear, and then start the chain deliberately at the moment of maximum supercriticality.

::: note
How fast the assembly must be depends on how many stray neutrons the material produces. A material with few spontaneous fissions tolerates a slow assembly; a material with many requires a fast one.
:::

## Gun-type: Little Boy

The simplest approach is to keep two subcritical pieces apart and then bring them together quickly. In a **gun-type** weapon, a conventional propellant charge drives one subcritical piece of uranium down a barrel into another. Neither piece alone can sustain a chain; together they form a supercritical mass.

By the standards of a chain reaction, which completes in about a microsecond, this assembly is slow. It works for uranium-235 because U-235 has a very low rate of spontaneous fission, so a stray neutron is unlikely to appear during the assembly.

The Hiroshima bomb, **Little Boy**, used this method with about 64 kg of uranium enriched to an average of about 80% U-235. Only about a kilogram of it fissioned, a little over 1%. It was inefficient, but its designers were so confident of it that it was **never tested before it was used**. Its yield, about 15 kilotons, was determined afterwards from measurements and damage surveys.

## Why plutonium could not use a gun

The original plan was a plutonium gun, codenamed Thin Man. It assumed that reactor-made plutonium would behave like the tiny samples previously made in particle accelerators.

::: history 1944 · Los Alamos
In spring and summer 1944 Emilio Segrè's group measured the first plutonium from a production reactor and found its rate of spontaneous fission was far higher than in the earlier samples. The cause was **plutonium-240**, made when a plutonium-239 nucleus captures a neutron instead of fissioning. Reactor plutonium inevitably contains some Pu-240, which fissions spontaneously at a high rate and fills the material with stray neutrons. A gun assembly would almost certainly predetonate. In July 1944 Thin Man was abandoned and Los Alamos was reorganized around a much harder idea.
:::

Plutonium-240 is chemically identical to plutonium-239, so it could not be removed by chemistry, and separating isotopes of plutonium was out of the question. The assembly had to be made much faster instead.

## Implosion: Fat Man

The alternative relied on the density result from the critical-mass lesson: critical mass falls as density rises. In an **implosion** weapon, a subcritical sphere of plutonium is surrounded by conventional high explosive. When the explosive is detonated, it drives a converging shock wave inward that compresses the sphere. Its density rises, its critical mass falls below the amount present, and it becomes supercritical.

Implosion is much faster than a gun, because the material only has to move a short distance inward at very high speed, so it outruns the stray neutrons from Pu-240. It also achieves a higher degree of supercriticality, which made it more efficient.

The difficulty was symmetry. If the inward push is uneven, the sphere squirts out sideways rather than compressing, and the weapon fails. Making the compression smooth and uniform was the central technical problem at Los Alamos from 1944 onward, and it was uncertain enough that this design was tested at Trinity before use.

The Nagasaki bomb, **Fat Man**, contained about 6.2 kg of plutonium, about a tenth of the uranium in Little Boy, and produced about 21 kilotons. Roughly a kilogram of its plutonium fissioned, roughly 15 to 20% of the material.

::: history 1943 · Los Alamos
Implosion was first proposed by Seth Neddermeyer, a physicist whose idea was initially received with scepticism. It was taken up seriously after John von Neumann showed in late 1943 that compression could be much greater than first thought. Once the Pu-240 problem appeared, it became the only route to a plutonium weapon.
:::

Implosion became the basis of almost all later fission weapons, because it uses less material more efficiently. The gun method was simpler but wasteful, and it works only with uranium.

::: quiz
? What is predetonation?
- Detonating the weapon too high above the ground
+ A stray neutron starting the chain before assembly is complete, giving a fizzle
- The conventional explosive failing to fire
= A chain started while the material is only slightly supercritical blows it apart early, so the yield is small.
? Why could plutonium not be used in a gun-type weapon?
- Plutonium cannot sustain a chain reaction
- Plutonium is too dense to fit in a barrel
+ Pu-240 fissions spontaneously so often that a gun is too slow to avoid predetonation
= Reactor plutonium always contains Pu-240, whose stray neutrons demand a much faster assembly.
? How does implosion make a subcritical sphere supercritical?
+ It compresses the sphere so its density rises and its critical mass falls
- It adds a second piece of plutonium
- It converts U-238 into plutonium during the explosion
= Critical mass scales as 1/ρ², so squeezing the same material makes it supercritical.
:::
