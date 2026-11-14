Put two protons and two neutrons on a scale, then fuse them into a helium nucleus and weigh it again. The helium is lighter, by about 0.75%. The missing mass is the key to every nuclear energy source, from the Sun to a reactor.

## Mass defect

A nucleus is a bound system: to pull it apart into free protons and neutrons you must supply energy, the **binding energy** B. By Einstein's E = mc², energy has mass, so a bound nucleus weighs less than its separate parts. The shortfall is the **mass defect**.

::: equation Binding energy from the mass defect
B = \left(Z m_p + N m_n - M\right) c^2
:::

Here Z is the number of protons, N the number of neutrons, and M the mass of the nucleus. Nuclear masses are given in atomic mass units (u, one twelfth of the mass of a carbon-12 atom), and 1 u corresponds to 931.5 MeV (a mega-electronvolt, the energy an electron gains crossing a million volts).

::: steps The binding energy of helium-4
$ \Delta m = 2 m({}^{1}\mathrm{H}) + 2 m_n - m({}^{4}\mathrm{He})
Tables list masses of neutral atoms, electrons included. Using hydrogen atoms instead of bare protons supplies the two electrons that the helium atom also contains, so the electron masses cancel.
---
$ \Delta m = 2(1.007825) + 2(1.008665) - 4.002603 = 0.030377\,\text{u}
The parts weigh 4.032980 u; the helium atom weighs 4.002603 u.
---
$ B = 0.030377\,\text{u} \times 931.5\,\text{MeV/u} \approx 28.3\,\text{MeV}
Convert mass to energy. That is about 7.1 MeV for each of the four nucleons.
:::

Chemical bonds are worth a few electronvolts. Nuclear binding is worth millions of electronvolts per particle, which is why nuclear reactions release about a million times more energy than chemical ones.

::: history 1919–1920 · Cambridge and Cardiff
Francis Aston built the mass spectrograph, which sorts ions by mass, and measured atomic masses to one part in a thousand. He found that a helium atom weighs about 0.7% less than four hydrogen atoms. Arthur Eddington seized on the result in 1920: if the Sun turned hydrogen into helium, that 0.7% of mass, converted to energy, could keep it shining for billions of years. Nobody yet knew how such a conversion could happen.
:::

## The binding energy curve

What decides whether a nuclear reaction releases energy is the binding energy *per nucleon*, B/A, where A = Z + N is the mass number.

::: model binding-curve
Slide through the nuclei. Light ones climb steeply, the peak is near iron, and heavy ones slowly lose ground.
:::

B/A rises steeply from about 1.1 MeV for the deuteron to 7.1 MeV for helium-4, reaches a broad peak of about 8.8 MeV near iron-56 and nickel-62, then declines to about 7.6 MeV for uranium-238. Any reaction that moves nuclei toward the peak releases energy: joining light nuclei (**fusion**) or splitting heavy ones (**fission**). The total number of nucleons is unchanged; they simply end up more tightly bound, and the products weigh less.

## The liquid-drop formula

In 1935 Carl Friedrich von Weizsäcker wrote a formula for B that treats the nucleus as a drop of charged liquid. Hans Bethe and Robert Bacher refined it in 1936. Each term has a simple physical reason.

::: equation Semi-empirical mass formula
B = a_V A - a_S A^{2/3} - a_C \frac{Z(Z-1)}{A^{1/3}} - a_A \frac{(A-2Z)^2}{A} + \delta
:::

- **Volume** (a_V ≈ 15.8 MeV). The strong force is short-ranged, so each nucleon bonds only with its neighbours. Binding grows in proportion to the number of nucleons, A.
- **Surface** (a_S ≈ 18.3 MeV). Nucleons at the surface have fewer neighbours. The surface area grows as R², and since R ∝ A^(1/3), as A^(2/3). This term dominates small nuclei and makes them loosely bound.
- **Coulomb** (a_C ≈ 0.71 MeV). Every proton repels every other: Z(Z − 1)/2 pairs, at a typical separation proportional to R. This term grows almost as Z² and drags down heavy nuclei.
- **Asymmetry** (a_A ≈ 23.2 MeV). The Pauli exclusion principle, which forbids two identical nucleons in the same state, makes it cheapest to have equal numbers of protons and neutrons. Any imbalance A − 2Z costs energy.
- **Pairing** (δ, about 12/√A MeV). Nucleons of the same kind bind slightly more strongly in pairs. δ is positive for even Z and even N, negative for odd Z and odd N, and zero otherwise.

The coefficients are fitted to measured masses, which is why the formula is called semi-empirical. For iron-56 (Z = 26) the terms are:

| Term | MeV |
|---|---|
| Volume | +885 |
| Surface | −268 |
| Coulomb | −121 |
| Asymmetry | −7 |
| Pairing | +2 |
| **Total** | **491** |

The measured value is 492 MeV, or 8.8 MeV per nucleon. For most nuclei the formula is good to within about 1%.

It also explains the shape of the curve. Small nuclei lose binding to the surface term, large ones to the Coulomb term, and the best compromise falls in the middle, near iron. It explains why heavy stable nuclei carry extra neutrons: uranium-238 has 146 neutrons to 92 protons, because the asymmetry penalty for adding neutrons is cheaper than the Coulomb penalty for adding protons.

::: note
A nucleus weighs less than its parts by B/c². Because binding per nucleon peaks near iron, energy is released by fusing light nuclei or splitting heavy ones, and in both cases a fraction of a percent of the mass disappears as energy.
:::

The formula misses one thing: nuclei with certain "magic" numbers of protons or neutrons (2, 8, 20, 28, 50, 82, 126) are extra stable, a sign that nucleons also fill shells, much as electrons do in atoms.

::: quiz
? Why does a helium-4 nucleus weigh less than two free protons and two free neutrons?
- Some nucleons are destroyed when it forms
+ Energy was released when it formed, and that energy carried away mass
- Neutrons lose mass inside nuclei
= The binding energy, 28.3 MeV, left as radiation when the nucleus formed; by E = mc² the nucleus is lighter by B/c².
? Which term in the formula makes binding per nucleon fall for heavy nuclei?
- The surface term
+ The Coulomb term
- The volume term
= Proton repulsion grows roughly as Z², faster than the attractive volume term, which grows as A.
? Why can both fission and fusion release energy?
- Both destroy nucleons
- Both reduce the number of protons
+ Both move nuclei toward the peak of the binding energy curve
= Light nuclei gain binding per nucleon by joining, heavy ones by splitting; the peak is near iron.
:::
