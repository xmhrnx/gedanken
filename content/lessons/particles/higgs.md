The electroweak theory of the previous lessons had a flaw at its centre. The symmetry that produces the photon, W and Z, the very thing that makes the theory work, also forbids those particles from having mass. Writing a mass into the equations by hand breaks the symmetry, and once the symmetry is broken the theory's predictions at high energy become nonsense, with probabilities above 100%, just as Fermi's old theory of beta decay had given. Yet the W and Z are among the heaviest particles known. Worse, the same symmetry forbids the electron and the quarks from having mass too, and they plainly do. This lesson explains the way out, proposed in 1964 and confirmed 48 years later: a field that fills all of space and gives particles their masses without the equations ever losing their symmetry.

## Can a symmetric law have a lopsided result?

The key idea came from an unexpected place, the physics of superconductors and magnets. In those systems the laws are perfectly symmetric, yet the state the material settles into is not. A magnet is a good example. The forces between the atoms in a piece of iron do not prefer any direction in space. But below a certain temperature the atoms' tiny magnetic moments line up with each other, and they must line up in some particular direction. The laws are symmetric; the outcome is not. Physicists call this **spontaneous symmetry breaking**: the symmetry is still in the equations, but the state the system actually occupies has hidden it.

In 1960 Yoichiro Nambu suggested that the same thing might happen in particle physics, with the vacuum, empty space itself, playing the role of the magnet. But there was a catch. Jeffrey Goldstone showed in 1961 that this kind of symmetry breaking should always produce new massless particles, and no such particles were seen. For a few years it looked as though the idea could not work. Then in 1962 Philip Anderson, a solid-state physicist, pointed out that in a superconductor the expected massless particle does not appear; instead the electromagnetic field inside the superconductor behaves as though the photon had acquired mass. That hint was what the particle physicists needed.

::: history 1964 · Brussels, Edinburgh and London
Within a few months of each other, three groups published the same mechanism: François Englert and Robert Brout in Brussels, Peter Higgs in Edinburgh, and Gerald Guralnik, Carl Hagen and Tom Kibble in London. They showed that when the symmetry belongs to a force, the would-be massless particles that Goldstone had predicted are absorbed by the force carriers, which become massive. Higgs's second paper on the subject was rejected by the first journal he sent it to, and in revising it he added a remark that the theory also predicted a new massive particle, the ripple in the new field. That remark is why the field and the particle carry his name. In 1967 and 1968 Steven Weinberg and Abdus Salam independently built the mechanism into the electroweak theory, using it to give the W and Z their masses while leaving the photon massless.
:::

## A field that is not zero

The proposal is that space is filled with a new field, now called the **Higgs field**, whose lowest-energy state is not zero. Every other field we know is zero in empty space: no electric field, no magnetic field, nothing. The Higgs field instead has the same constant value everywhere, even in a perfect vacuum, because that is the arrangement with the least energy. The equations stay symmetric, but the vacuum the universe sits in does not. A pencil balanced on its point is a fair picture of this: the situation is perfectly symmetric, but it is unstable, and once the pencil falls it must point in some particular direction, though the laws did not prefer any. The picture breaks down in one respect: the pencil falls in space, while the Higgs field "falls" to a value that is the same at every point in space.

What makes the field settle away from zero is the shape of its energy. To see how, treat the field as a single number φ. The real Higgs field has four components, but one number keeps the essential mathematics.

::: steps Where the Higgs field settles
$ V(\phi) = -\tfrac{1}{2}\mu^2\phi^2 + \tfrac{1}{4}\lambda\phi^4
This is the energy density of the field for a given value φ, with μ² and λ positive constants. For an ordinary field the φ² term would have a plus sign, making φ = 0 the bottom of a bowl. Here it has a minus sign, so near φ = 0 the energy goes down as φ grows: zero is a hilltop. The φ⁴ term, which grows faster, eventually turns the curve back up, so the shape is a hill in the middle with a valley on either side.
---
$ \frac{\mathrm{d}V}{\mathrm{d}\phi} = -\mu^2\phi + \lambda\phi^3 = 0
The field settles where its energy is lowest, and at a minimum the slope of the curve is zero. Differentiating V with respect to φ and setting the result to zero gives the condition for the flat points.
---
$ \phi = 0 \quad \text{or} \quad \phi = v = \frac{\mu}{\sqrt{\lambda}}
Factor the equation as φ(λφ² − μ²) = 0. One solution is φ = 0, which is the top of the hill and therefore unstable, like the balanced pencil. The other is φ = v = μ/√λ (with its mirror image at −v), the bottom of the valley. That is where the field settles in empty space, and measurements of the W mass and the weak force's strength give v ≈ 246 GeV.
---
$ V''(v) = -\mu^2 + 3\lambda v^2 = 2\mu^2
Now ask how sharply the valley curves at its bottom. Differentiating a second time gives −μ² + 3λφ², and at φ = v we can use λv² = μ² from the previous step to simplify it to 2μ². This matters because a small ripple of the field about its resting value behaves like a particle, and the stiffer the valley, the more energy each ripple carries.
---
$ m_H = \sqrt{2}\,\mu = \sqrt{2\lambda}\, v
In quantum field theory, in units where ħ = c = 1, the squared mass of the particle equals that curvature, so m_H² = 2μ² = 2λv². This particle is the **Higgs boson**. The theory fixes v from the W mass but says nothing about λ, so it could not predict the Higgs mass; that had to be found by experiment. The measured 125 GeV gives λ ≈ 0.13.
:::

## How does a field give mass?

Particles get mass by interacting with the constant field that fills space. A particle that interacts with it strongly gets a large mass; one that does not interact with it at all stays massless. The mass is not something added to the particle from outside but the energy of its interaction with the field, which is present everywhere it goes.

For the W and Z, the strength of that interaction is fixed by the strength of the weak force itself, measured by a number g ≈ 0.65. The theory gives m_W = gv/2 ≈ 0.65 × 246/2 GeV ≈ 80 GeV, exactly as measured. The photon is the particular combination of the electroweak fields that does not interact with the Higgs field, and so it stays massless. That is why electromagnetism has infinite range while the weak force is confined to 10⁻¹⁸ metres, both coming from one theory.

Quarks and charged leptons get their masses the same way, but each has its own interaction strength, called a **Yukawa coupling** and written y, and the mass is m = yv/√2. Here the theory is less satisfying. The couplings are free parameters: the top quark's is almost exactly 1, the electron's about 3 × 10⁻⁶, and nothing in the theory says why. The Higgs mechanism explains how fermions can have mass without breaking the symmetry of the equations. It does not explain why the masses have the values they do.

Popular accounts often compare the Higgs field to molasses, slowing particles as they move through it. The comparison misleads. Molasses produces drag, a force that opposes motion and brings things to rest. Mass is something different: resistance to any change in motion, including starting, stopping and turning, and a massive particle moving at steady speed through the Higgs field feels no drag at all.

It is also worth being clear about what the Higgs does not do. Most of the mass of ordinary matter does not come from it. A proton contains two up quarks and a down quark, whose Higgs-given masses add up to about 9 MeV, roughly 1% of the proton's 938 MeV. The other 99% is the energy of the gluon field binding the quarks together and the energy of the quarks' own rapid motion, counted as mass through E = mc². The Higgs gives mass to the elementary particles; the strong force gives mass to us.

## Finding the particle

Predicting a particle without a mass is a hard thing to test, because you do not know where to look. Over several decades, searches at successive colliders ruled out one range of masses after another. The final search needed a new machine.

A Higgs boson lives only about 10⁻²² seconds, far too short to leave a track, so it can be identified only by what it decays into. It is also rare: only about one proton collision in a billion makes one, and most of the ways it decays look much like the ordinary debris of other collisions. Finding it meant making billions upon billions of collisions and searching for a small excess at one particular mass.

::: history 4 July 2012 · CERN, Geneva
The Large Hadron Collider, a ring 27 kilometres around that collides protons at energies of several TeV (trillion electronvolts), had two large general-purpose detectors, ATLAS and CMS, each built and run by about 3,000 physicists. Working independently, so that each could check the other, both had been closing in on the same signal. At a seminar at CERN that morning, both reported a new boson with a mass near 125 GeV, each at the level of "five standard deviations", the threshold particle physicists require before claiming a discovery, which corresponds to a chance of less than one in a million that a random fluctuation would produce so strong a signal. Peter Higgs, then 83, was in the audience and wiped away tears. Englert and Higgs received the Nobel Prize in Physics in 2013; Brout had died in 2011.
:::

The clearest signals came from two rare kinds of decay: into two photons, and into four charged leptons, electrons or muons, by way of two Z bosons. In both cases every product can be measured precisely, so adding up their energies and momenta gives the mass of the particle that decayed, and the Higgs shows up as a sharp peak standing above a smooth background. The measured mass is about 125 GeV, roughly 133 times the mass of a proton.

Since 2012 the particle has been examined closely. Its spin is zero, as predicted, which makes it the only fundamental particle known with no spin at all. Its decays to bottom quarks, tau leptons, and pairs of W and Z bosons have been measured, and its interactions with the heaviest particles, including the top quark, have been checked. In each case the strength of the interaction is proportional to the particle's mass, exactly as the mechanism requires, and everything agrees with the Standard Model within present precision.

::: note
The Higgs field has a constant, nonzero value throughout space. Particles that interact with it gain mass in proportion to the strength of that interaction, while the equations keep their symmetry. The Higgs boson is a ripple in that field, and it was found at 125 GeV in 2012.
:::

::: quiz
? Why couldn't the W and Z masses simply be written into the electroweak theory?
- The masses were unknown in the 1960s
+ Mass terms break the symmetry the theory is built on
- Massive particles cannot carry forces
= The symmetry that produces the force carriers requires them to be massless, and breaking it by hand ruins the theory's predictions at high energy. The Higgs mechanism lets the vacuum break the symmetry instead, while the equations keep it.
? In the simple model, what sets the Higgs boson's mass?
- The value v alone
+ The curvature of the energy curve at its minimum
- The mass of the top quark
= A ripple about the bottom of the valley carries more energy the more sharply the valley curves. The squared mass equals the second derivative of V at v, which is 2μ² = 2λv².
? Where does most of the proton's mass come from?
- The Higgs field acting on its three quarks
+ The energy of the gluon field and the quarks' motion
- The electrons bound to it
= The masses the Higgs gives the up and down quarks add up to only about 1% of the proton's mass. The rest is the energy of the strong interaction inside the proton, which counts as mass through E = mc².
:::
