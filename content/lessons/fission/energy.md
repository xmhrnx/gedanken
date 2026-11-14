When Meitner and Frisch explained fission, they could say at once that each split released about 200 MeV. That number is the starting point for everything that follows in this chapter, from power stations to the weapons of 1945. This lesson asks where the energy comes from, since nothing in fission is destroyed, and then follows the number through to a figure that alarmed physicists in 1939: how much energy a single kilogram of uranium could release if every nucleus in it split.

## Where the energy hides

A nucleus is held together by the strong force, and pulling it apart into separate protons and neutrons would cost energy. That cost is called the nucleus's **binding energy**. Run the process backwards and the same energy is released: when nucleons come together into a nucleus, the binding energy comes out, usually as radiation. So the more tightly bound a nucleus is, the more energy left it when it formed.

Energy that leaves takes mass with it. That is the content of Einstein's relation between mass and energy:

::: equation Mass–energy equivalence
E = mc^2
:::

The equation is often read as a recipe for turning matter into energy, but its real meaning is more general and more useful: any change in a system's energy is also a change in its mass, by the energy divided by c². Because c² is so large, ordinary energy changes make no measurable difference to mass. Nuclear energies are large enough that they do.

The clearest example is helium. Weigh a helium-4 nucleus and it comes out lighter than two free protons plus two free neutrons by about 0.75%. Nothing is missing; the same four particles are there. But when they came together, their binding energy left, and that energy carried away its share of mass. The shortfall is called the **mass defect**, and multiplied by c² it equals the binding energy.

Nuclear physicists measure masses in **atomic mass units** (u), defined as one twelfth of the mass of a carbon-12 atom, which makes a proton or neutron weigh just over 1 u. Energies are measured in electronvolts (eV), where one electronvolt is the energy an electron gains crossing a potential difference of one volt, and in MeV, millions of electronvolts. In these units, 1 u of mass corresponds to 931.5 MeV of energy. So a mass difference of only a thousandth of an atomic mass unit is worth nearly 1 MeV, roughly a million times the energy of a chemical bond. That ratio is the whole reason nuclear energy is different in kind from chemical energy.

## The binding energy curve

Whether a nuclear reaction releases energy depends not on the total binding energy, which simply grows with the size of the nucleus, but on the binding energy *per nucleon* (per proton or neutron). That number says how tightly each particle is held, and a reaction that leaves each nucleon more tightly held must release the difference. Plotting it against the number of nucleons gives one of the most important graphs in physics.

::: model binding-curve
Binding energy per nucleon against mass number. Find iron and nickel at the peak, then find uranium on the gently falling right-hand side, and notice where the fragments of a split uranium nucleus would land.
:::

The curve rises steeply through the light nuclei, reaches a peak near iron-56 at about 8.8 MeV per nucleon, and then falls slowly to about 7.6 MeV per nucleon for uranium. Two forces explain that shape. The strong force is short-ranged, so each nucleon is bound only to its immediate neighbours. In a small nucleus, a large share of the nucleons sit at the surface with neighbours on one side only, so small nuclei are loosely bound. The electrical repulsion between protons is long-ranged: every proton pushes on every other, however far apart they are within the nucleus. As nuclei grow, the repulsion grows faster than the attraction, and it steadily erodes the binding of heavy nuclei. The peak near iron is the best compromise between the two.

So any move toward the peak releases energy. Light nuclei climb the steep left side by joining together, which is fusion. Heavy nuclei climb the gentle right side by splitting, which is fission. A uranium nucleus that breaks into two fragments with mass numbers somewhere between about 90 and 140 moves from 7.6 MeV per nucleon to about 8.5 MeV per nucleon. Each nucleon ends up about 0.9 MeV more tightly bound.

::: note
Fission does not turn uranium into energy. The total number of protons and neutrons is unchanged. Energy is released because the nucleons are rearranged into more tightly bound nuclei, and so the products weigh slightly less than what went in.
:::

## About 200 MeV per fission

With the curve in hand, the energy of one fission can be estimated by simple bookkeeping: compare the total binding before and after.

::: steps Energy released by one U-235 fission
$ B({}^{236}\mathrm{U}) \approx 236 \times 7.6\,\text{MeV} \approx 1790\,\text{MeV}
When U-235 absorbs a neutron it becomes, for an instant, a nucleus of U-236, which is the thing that actually splits. Its total binding energy is the number of nucleons times the binding per nucleon read from the curve, about 7.6 MeV. That gives roughly 1,790 MeV, the energy it would take to pull the whole nucleus apart.
---
$ B(\text{fragments}) \approx 236 \times 8.5\,\text{MeV} \approx 2010\,\text{MeV}
The two fragments sit higher on the curve, at about 8.5 MeV per nucleon. Counting all 236 nucleons slightly overstates their total, because the two or three neutrons freed in the split are not bound to anything. But those few neutrons barely change the sum, and this is an estimate, so the simplification is reasonable.
---
$ \Delta E \approx 2010\,\text{MeV} - 1790\,\text{MeV} \approx 220\,\text{MeV}
The products are more tightly bound than the original nucleus by about 220 MeV, and that difference is the energy released. It agrees with Meitner and Frisch's estimate from electrical repulsion. Careful measurements give about 200 MeV, and the gap comes from rounding the binding per nucleon.
---
$ \Delta m = \frac{\Delta E}{c^2} \approx 200\,\text{MeV} \div 931.5\,\text{MeV/u} \approx 0.21\,\text{u}
Convert the energy back into mass using 931.5 MeV per atomic mass unit. The products weigh about a fifth of an atomic mass unit less than the uranium nucleus and neutron did, which is less than 0.1% of the original mass. This is the figure Meitner worked out from memory on the walk in December 1938.
:::

Most of the 200 MeV, about 170 MeV, appears as kinetic energy of the two fragments, which fly apart under their electrical repulsion. Being heavy and highly charged, they slam into the surrounding atoms and stop within a small fraction of a millimetre, turning their energy into heat in the material. That is how a reactor makes its heat. The rest of the energy is carried by the free neutrons, by gamma rays released at the moment of fission, and by the later radioactive decay of the fragments, which are overloaded with neutrons and decay in steps toward stability.

## Comparison with chemistry

Chemical reactions rearrange the outer electrons of atoms, which are held by a few electronvolts. Burning one carbon atom to carbon dioxide releases about 4 eV. A molecule of TNT, the standard high explosive, releases about 10 eV when it detonates; it is powerful not because each molecule releases much, but because it releases it so quickly.

::: equation Nuclear versus chemical energy per reaction
200\,\text{MeV} \div {\sim}4\,\text{eV} \approx 5 \times 10^7
:::

Fission releases tens of millions of times more energy per event than burning. A uranium atom and a carbon atom differ in mass by a factor of about 20, so one might wonder whether the comparison per kilogram is less dramatic. The next calculation shows that it is still enormous.

::: steps Fully fissioning one kilogram of U-235
$ N = \left(1000\,\text{g} \div 235\,\text{g/mol}\right) \times 6.02 \times 10^{23}\,\text{/mol} \approx 2.56 \times 10^{24}
First count the nuclei. One mole of U-235 has a mass of 235 grams, so a kilogram is about 4.26 moles. Each mole contains Avogadro's number of atoms, 6.02 × 10²³. Multiplying gives about 2.56 × 10²⁴ nuclei in a kilogram.
---
$ E = 2.56 \times 10^{24} \times 200\,\text{MeV} \approx 5.1 \times 10^{26}\,\text{MeV}
If every one of those nuclei split, each releasing 200 MeV, the total energy is simply the number of nuclei times the energy per fission. The answer in MeV is correct but hard to picture, so the next step converts it to the everyday unit of energy.
---
$ E \approx 5.1 \times 10^{26}\,\text{MeV} \times 1.60 \times 10^{-13}\,\text{J/MeV} \approx 8.2 \times 10^{13}\,\text{J}
One MeV is 1.60 × 10⁻¹³ joules. Multiplying gives about 8 × 10¹³ joules, enough to supply a large city with electricity for about a day, from a lump of metal small enough to hold in one hand.
---
$ 8.2 \times 10^{13}\,\text{J} \div 4.18 \times 10^{12}\,\text{J/kt} \approx 20\,\text{kt}
Explosive energy is conventionally measured in tonnes of TNT. A kiloton of TNT is defined as 4.184 × 10¹² J, so the kilogram of uranium is equivalent to about 20 kilotons. In practice not all of the 200 MeV counts: about 10 MeV per fission is carried off by neutrinos, which escape without depositing anything, and some fragment decay energy is released only slowly, long after the event. The usable figure is about 17 to 20 kilotons.
:::

A kilogram of TNT releases about 4.2 million joules, so a kilogram of fully fissioned uranium is equivalent to roughly twenty million kilograms of high explosive. This was the number that alarmed physicists in 1939. It meant that a weapon small enough to be carried by a single aircraft could match the explosive energy carried by thousands of heavy bombers. In practice the weapons of 1945 fissioned only a small fraction of their material, for reasons that the coming lessons explain: the material blows itself apart within a microsecond, long before most of it can split.

::: history 1905–1932 · Bern and Cambridge
Einstein derived E = mc² in 1905 as a consequence of special relativity, in a three-page paper that ended by suggesting that salts of radium, which release energy continuously, might be used to test it. For decades it seemed of theoretical interest only. Precise mass measurements, above all by Francis Aston with his mass spectrograph in the 1920s, showed that nuclei weigh less than their parts. In 1932 John Cockcroft and Ernest Walton at Cambridge split lithium nuclei with artificially accelerated protons, producing two alpha particles, and the energy of the alphas matched the mass lost in the reaction. It was the first direct check of E = mc² in a nuclear reaction. Fission, six years later, made the equation famous.
:::

::: quiz
? Why does splitting uranium release energy?
+ The fragments are more tightly bound per nucleon than uranium
- Some protons are destroyed and turned into energy
- The neutron that causes fission carries the energy in with it
= The fragments sit higher on the binding energy curve, so the same nucleons are held more tightly after the split. The number of nucleons does not change; the products are simply lighter, and the lost mass appears as energy.
? Roughly how much energy does one U-235 fission release?
- 200 eV
- 2 MeV
+ 200 MeV
= About 236 nucleons each become about 0.9 MeV more tightly bound, which adds up to roughly 200 MeV. That is tens of millions of times the energy of a chemical reaction.
? About how much TNT is equivalent to fully fissioning 1 kg of U-235?
- 20 tonnes
+ 20 kilotons
- 20 megatons
= A kilogram holds 2.56 × 10²⁴ nuclei, and at 200 MeV each they release about 8 × 10¹³ J. Dividing by 4.18 × 10¹² J per kiloton gives about 20 kilotons.
:::
