Put two protons and two neutrons on a very precise scale, then combine them into a helium nucleus and weigh that. The helium is lighter, by about 0.75%. Nothing has been thrown away: the same four particles are there. Mass has simply gone missing. That small shortfall turns out to be the key to every nuclear energy source, from the Sun to a power station, and this lesson explains where the mass goes, how to calculate it, and why it makes both splitting heavy nuclei and joining light ones release energy.

## Why a bound nucleus weighs less

Think first about what it means for something to be bound. A nucleus holds together because the strong force pulls its nucleons toward each other. To take it apart, separating every proton and neutron and moving them far away, you would have to do work against that pull, feeding energy in. The amount needed is called the **binding energy**, B. Turn the process around and the same energy comes out: when free nucleons fall together into a nucleus, B is released, usually as gamma rays carrying it away.

Now bring in Einstein's E = mc². Its meaning is broader than "mass can be turned into energy". It says that energy itself has mass: any system that loses energy E becomes lighter by E/c². A hot cup of coffee weighs a tiny bit more than the same cup cold. For a nucleus the effect is not tiny at all. When the nucleons come together and the binding energy leaves, the nucleus is left lighter than its separate parts by B/c². That difference in mass is called the **mass defect**, and measuring it is how binding energies are found.

::: equation Binding energy from the mass defect
B = \left(Z m_p + N m_n - M\right) c^2
:::

Here Z is the number of protons, each of mass m_p; N is the number of neutrons, each of mass m_n; and M is the mass of the assembled nucleus. The bracket is the mass defect, and multiplying by c² converts it into energy.

To use this in practice, two units help. Nuclear masses are given in **atomic mass units** (u), defined as one twelfth of the mass of a carbon-12 atom, which makes a proton or neutron weigh just over 1 u. Energies are given in MeV, millions of **electronvolts**, where one electronvolt is the energy an electron gains in crossing a potential difference of one volt. The conversion between them is what makes the calculation easy: a mass of 1 u corresponds to an energy of 931.5 MeV.

::: steps The binding energy of helium-4
$ \Delta m = 2 m({}^{1}\mathrm{H}) + 2 m_n - m({}^{4}\mathrm{He})
Start from the mass defect for helium-4, which has two protons and two neutrons. There is a practical wrinkle: tables list the masses of whole neutral atoms, electrons included, not bare nuclei. So instead of two bare protons, use two hydrogen atoms, each a proton plus an electron. The helium atom also contains two electrons, so the electron masses appear on both sides and cancel, leaving the nuclear mass defect we want.
---
$ \Delta m = 2(1.007825) + 2(1.008665) - 4.002603\,\text{u}
Put in the measured masses: 1.007825 u for a hydrogen atom, 1.008665 u for a neutron, and 4.002603 u for a helium atom. Notice how close all of these are to whole numbers. The mass defect is hiding in the decimal places, which is why it took precise instruments to find it.
---
$ \Delta m = 4.032980\,\text{u} - 4.002603\,\text{u} = 0.030377\,\text{u}
The separate parts add up to 4.032980 u, and the assembled atom weighs 4.002603 u. The difference, about 0.03 u, is three quarters of one percent of the total. Small as it is, it is real, and it is the whole of the binding energy expressed as mass.
---
$ B = 0.030377\,\text{u} \times 931.5\,\text{MeV/u} \approx 28.3\,\text{MeV}
Convert mass to energy with the factor 931.5 MeV per u. Shared among four nucleons, that is about 7.1 MeV each. This is the energy that came out, as radiation, when the helium nucleus formed, and it is the energy you would have to put back in to take it apart.
:::

To see how large that is, compare it with chemistry. The bonds that hold atoms together in molecules, and that break and form when fuel burns, are worth a few electronvolts each. Nuclear binding is worth millions of electronvolts per particle. That ratio, roughly a million, is why a nuclear reaction releases about a million times more energy than a chemical reaction involving the same number of atoms, and why a large nuclear power station needs a few tens of tonnes of fuel a year where a coal station of the same output burns millions of tonnes.

::: history 1919–1920 · Cambridge and Cardiff
The mass defect was discovered before anyone could explain it. At Cambridge, Francis Aston built the **mass spectrograph**, an instrument that sends charged atoms through electric and magnetic fields so that atoms of different mass land in different places, and measured atomic masses to about one part in a thousand. Among his results was that a helium atom weighs about 0.7% less than four hydrogen atoms. In 1920 the astronomer Arthur Eddington, addressing the British Association in Cardiff, saw what this could mean. The Sun had been shining for far longer than any chemical fuel could last. If, somewhere in its core, hydrogen were being built up into helium, that 0.7% of the mass, released as energy, could keep it shining for billions of years. Nobody yet knew how hydrogen nuclei could be made to combine, and the details would take another two decades, but Eddington's guess about the Sun's fuel was right.
:::

## The binding energy curve

Total binding energy grows with the size of a nucleus, simply because there are more nucleons. That makes it a poor guide on its own. What decides whether a nuclear reaction releases energy is how tightly each nucleon is held, the binding energy *per nucleon*, B/A, where A = Z + N is the mass number. If a reaction rearranges nucleons so that, on average, each ends up more tightly bound, the difference comes out as energy.

::: model binding-curve
Slide through the nuclei from hydrogen to uranium. Notice how steeply the light ones climb, where the curve peaks, and how gently it falls for the heavy ones.
:::

The curve has a shape worth memorising. It rises steeply from about 1.1 MeV per nucleon for the deuteron to 7.1 MeV for helium-4, which sits noticeably above its neighbours because it is an especially compact arrangement. It then climbs more slowly to a broad peak of about 8.8 MeV per nucleon around iron-56 and nickel-62, the most tightly bound nuclei in nature. After that it declines gradually, reaching about 7.6 MeV per nucleon for uranium-238.

Any reaction that moves nuclei toward the peak releases energy. Light nuclei can climb the steep left side by joining together, which is **fusion**, the process that powers the Sun. Heavy nuclei can climb the gentle right side by splitting into two middle-sized pieces, which is **fission**, the process behind nuclear reactors and the first atomic bombs. In both cases no nucleons are created or destroyed. The same protons and neutrons simply end up in more tightly bound arrangements, the products weigh slightly less than what went in, and the missing mass appears as energy.

## Why the curve has this shape

A curve this regular invites an explanation, and one was found by treating the nucleus as a drop of liquid. The idea fits what earlier lessons established: nucleons pack at a fixed density, like molecules in a liquid, and each one is attracted only by its immediate neighbours, just as a molecule in a drop is. In 1935 Carl Friedrich von Weizsäcker wrote down a formula for the binding energy on this basis, and Hans Bethe and Robert Bacher refined it in 1936. It is a sum of five terms, each with a clear physical reason.

::: equation Semi-empirical mass formula
B = a_V A - a_S A^{2/3} - a_C \frac{Z(Z-1)}{A^{1/3}} - a_A \frac{(A-2Z)^2}{A} + \delta
:::

The **volume term**, a_V A with a_V ≈ 15.8 MeV, is the starting point. Because the strong force is short-ranged, each nucleon bonds only with the few neighbours around it, so each contributes about the same amount of binding. Total binding should then be proportional to the number of nucleons, A. If this were the only term, every nucleus would have the same binding per nucleon and the curve would be flat.

The **surface term**, with a_S ≈ 18.3 MeV, corrects for the nucleons at the edge. They have neighbours on one side only, so they are less tightly bound, just as molecules at the surface of a liquid are, which is the origin of surface tension. The number of surface nucleons is proportional to the surface area, which grows as R², and since R is proportional to A^(1/3), as A^(2/3). In a small nucleus a large fraction of the nucleons are on the surface, so this term explains why light nuclei are loosely bound and why the curve rises steeply at the left.

The **Coulomb term**, with a_C ≈ 0.71 MeV, accounts for the electrical repulsion of the protons. Every proton repels every other, and there are Z(Z − 1)/2 pairs of them, at a typical separation proportional to the radius R, which is why A^(1/3) appears in the denominator. The term grows almost as Z², much faster than the volume term grows with A, so it is what drags the curve down for heavy nuclei.

The **asymmetry term**, with a_A ≈ 23.2 MeV, is a quantum effect. The **Pauli exclusion principle**, which forbids two identical particles from occupying the same quantum state, means protons and neutrons each fill their own ladder of energy levels from the bottom up. If a nucleus has many more neutrons than protons, the extra neutrons must sit on high rungs of their ladder while lower rungs on the proton ladder stand empty. That costs energy, so the formula charges a penalty for any imbalance, measured by A − 2Z, which is the same as N − Z.

Finally, the **pairing term** δ, of size about 12/√A MeV, reflects the fact that two protons, or two neutrons, bind slightly more strongly when they pair up. It is positive when both Z and N are even, negative when both are odd, and zero otherwise, which is why nuclei with even numbers of both are the most common stable ones.

The coefficients are not derived from first principles; they are adjusted so the formula matches measured masses, which is why it is called **semi-empirical**. The test of the idea is how well five terms do. For iron-56, with Z = 26 and A = 56, they give:

| Term | MeV |
|---|---|
| Volume | +885 |
| Surface | −268 |
| Coulomb | −121 |
| Asymmetry | −7 |
| Pairing | +2 |
| **Total** | **491** |

The measured binding energy is 492 MeV, or 8.8 MeV per nucleon. For most nuclei the formula agrees with experiment to within about 1%, which is remarkable for a model that treats the nucleus as a droplet of fluid.

The formula also explains the shape of the curve, rather than merely fitting it. Small nuclei lose binding to the surface term, because so many of their nucleons are on the outside. Large nuclei lose it to the Coulomb term, because proton repulsion grows faster than the attraction. The best compromise between the two falls in the middle, near iron. And the formula explains why heavy stable nuclei carry extra neutrons: uranium-238 has 146 neutrons to 92 protons. Adding protons would increase the Coulomb penalty sharply, while adding neutrons costs only the milder asymmetry penalty, so heavy nuclei settle with a neutron surplus. That surplus will matter a great deal when uranium splits, because the fragments inherit more neutrons than they can comfortably hold.

::: note
A nucleus weighs less than its separate parts by B/c², the binding energy divided by c². Because binding per nucleon peaks near iron, energy is released both by fusing light nuclei and by splitting heavy ones, and in each case a fraction of a percent of the mass disappears as energy.
:::

## What the drop leaves out

The liquid drop gets the broad trend right, but not every detail. Nuclei with certain numbers of protons or neutrons, 2, 8, 20, 28, 50, 82 and 126, are noticeably more tightly bound than the formula predicts. These became known as **magic numbers**. Helium-4, with 2 protons and 2 neutrons, is doubly magic, which is why it stands out above the curve. The explanation, worked out around 1949 by Maria Goeppert Mayer and independently by J. Hans D. Jensen and colleagues, is that nucleons, like electrons in an atom, fill shells of energy levels, and a nucleus with a completely filled shell is especially stable. A nucleus is in some ways a liquid drop and in other ways an orderly quantum structure, and both pictures are needed.

::: quiz
? Why does a helium-4 nucleus weigh less than two free protons and two free neutrons?
- Some nucleons are destroyed when it forms
+ Energy was released when it formed, and that energy carried away mass
- Neutrons lose mass inside nuclei
= When the nucleons came together, the binding energy of 28.3 MeV left as radiation. Since energy has mass, by E = mc², the nucleus left behind is lighter by B/c².
? Which term in the formula makes binding per nucleon fall for heavy nuclei?
- The surface term
+ The Coulomb term
- The volume term
= Every proton repels every other, so the repulsion grows roughly as Z², faster than the attractive volume term, which grows only as A. In heavy nuclei it wins enough ground to pull the curve down.
? Why can both fission and fusion release energy?
- Both destroy nucleons
- Both reduce the number of protons
+ Both move nuclei toward the peak of the binding energy curve
= Light nuclei gain binding per nucleon by joining and heavy ones by splitting, because the curve peaks near iron. Either way the products are more tightly bound and lighter, and the lost mass comes out as energy.
:::
