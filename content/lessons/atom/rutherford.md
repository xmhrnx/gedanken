By 1909 physicists knew that atoms contain electrons, and that electrons are light and negatively charged. Since atoms are electrically neutral and far heavier than their electrons, each atom must also contain positive charge, and that positive part must carry nearly all the mass. The open question was how the positive charge is arranged. J. J. Thomson's answer, the "plum pudding", spread it evenly through the whole atom, with electrons dotted inside. This lesson is about the experiment that overturned that picture, carried out largely by a twenty-year-old student, and the reasoning that turned a handful of strange observations into the discovery of the atomic nucleus. It showed that an atom is almost entirely empty space, with nearly all its mass packed into a point at the centre.

## A probe for the inside of an atom

You cannot look inside an atom with light: atoms are about 10⁻¹⁰ m across, several thousand times smaller than the wavelength of visible light, and no microscope can show detail much finer than the wavelength it uses. To learn what is inside, you need to throw something at them and see how it bounces.

Ernest Rutherford had the ideal projectile. Over the previous decade he had shown that some radioactive elements emit **alpha particles**, which he later identified as helium atoms stripped of their two electrons. Alpha particles are heavy, about 7,300 times the mass of an electron, and fast, around 2 × 10⁷ m/s, or about a fifteenth of the speed of light. That combination gives them a great deal of momentum, so they plough through matter in nearly straight lines. If something inside an atom could turn one aside sharply, it would have to exert a large force.

Now ask what Thomson's model predicts. Its electrons are useless for deflecting an alpha particle: a projectile 7,300 times heavier than an electron brushes it aside, the way a bowling ball brushes aside a marble. The positive charge is heavier, but it is spread evenly over the whole atom, and a spread-out charge produces only a weak electric field. Inside a uniformly charged ball the field is strongest at the surface and falls to zero at the centre, and for a ball the size of an atom it is never strong enough to do more than nudge a fast alpha. An alpha particle crossing a thin metal foil passes through a few thousand atoms, each nudging it a little in a random direction. The nudges partly cancel, and the alpha should come out within a degree or two of where it was going. Large deflections should essentially never happen.

::: history 1909–1911 · Manchester
At Manchester, Rutherford's assistant Hans Geiger was measuring how alpha particles spread out as they pass through thin metal foils. The method was painstaking. Each alpha that strikes a screen coated with zinc sulphide makes a faint flash of light, and the flashes were counted one by one through a microscope by an observer sitting in a darkened room, after half an hour of letting the eyes adjust. Ernest Marsden, an undergraduate of twenty, was learning the technique, and Rutherford suggested he check whether any alpha particles came off the front of the foil, bounced backward. Rutherford did not expect him to find any. Marsden found some: roughly 1 in 8,000 alpha particles striking a foil of platinum was turned through more than 90°. Rutherford later said it was "almost as incredible as if you fired a 15-inch shell at a piece of tissue paper and it came back and hit you." He thought about it for well over a year before publishing his explanation in 1911.
:::

## A point of positive charge

Rutherford's reasoning started from the size of the force. To turn an alpha particle around, a single encounter with a single atom must do it; many small nudges adding up could not plausibly produce such a large deflection as often as Marsden saw it. A single atom can exert a huge force only if its electric field is very strong somewhere, and the field of a charge Q at distance r is kQ/r², where k is Coulomb's constant. The only way to make the field strong is to get close to the charge, which means the charge must be packed into a tiny space. Shrink a ball of charge 10,000 times, and the field at its surface grows by 10,000², a hundred million times.

So Rutherford proposed that all the atom's positive charge, and nearly all its mass, sits in a tiny central **nucleus**, with the electrons far outside it. This picture explains both halves of the experiment. Most alpha particles pass through the vast empty space between nuclei and feel almost nothing, which is why nearly all of them go straight through. The rare alpha that happens to head almost straight at a nucleus comes so close to so concentrated a charge that the electric repulsion turns it right around.

How close does it get? We can work it out by following an alpha particle with kinetic energy K aimed straight at a gold nucleus.

::: steps Distance of closest approach
$ K = \frac{k(2e)(Ze)}{d}
As the alpha approaches, electric repulsion slows it, and at the closest point it stops for an instant before being driven back. At that moment all its kinetic energy has been turned into electric potential energy, which for two charges q₁ and q₂ a distance d apart is kq₁q₂/d. The alpha has charge 2e and the nucleus Ze, where Z is the number of protons. The gold nucleus is about fifty times heavier than the alpha, so we ignore its small recoil.
---
$ d = \frac{2kZe^2}{K}
Multiply both sides by d and divide by K, to get the distance on its own. The result already says something useful: a faster alpha, with larger K, gets closer, and a more highly charged nucleus, with larger Z, keeps it further away.
---
$ ke^2 = 1.44\,\text{MeV fm}
To put in numbers, it helps to use units suited to nuclei. Energies are measured in MeV, millions of electronvolts, where an electronvolt is the energy an electron gains crossing one volt. Distances are in femtometres, 1 fm = 10⁻¹⁵ m. In these units the combination ke², which sets the strength of the electric force between two elementary charges, is 1.44 MeV fm.
---
$ d = \frac{2 \times 79 \times 1.44\,\text{MeV fm}}{7.7\,\text{MeV}} \approx 30\,\text{fm} = 3 \times 10^{-14}\,\text{m}
For gold Z = 79, and the alpha particles Rutherford used had about 7.7 MeV. The MeV units cancel, leaving a distance of about 30 femtometres, which is about 3 × 10⁻¹⁴ m.
:::

Compare that with the size of the atom, about 10⁻¹⁰ m. The alpha particle got within about a three-thousandth of the atom's width of the centre, and all the way in, it still obeyed Coulomb's law, the simple inverse-square law for point charges. That means it never touched the nucleus, so the nucleus must be smaller still. Later experiments with faster particles, which can get closer, put the radius of the gold nucleus at about 7 fm, some 20,000 times smaller than the atom. To picture the scale: if the nucleus were a marble 1 cm across, the atom would be about 200 m across, a large sports stadium with a marble at its centre and a few electrons somewhere in the stands.

## The scattering law

A good story is not yet physics. Rutherford's real achievement was to make the idea predict numbers. He treated the alpha particle as a charged body moving past a point charge under an inverse-square repulsion, the same mathematics as a comet swinging past the Sun except that the force pushes instead of pulls. The path is a hyperbola, and how sharply it bends depends on how closely the alpha was aimed at the nucleus. Since the aim is random, he could work out what fraction of alpha particles should be scattered through each angle θ:

::: equation Rutherford's scattering formula
N(\theta) \propto \frac{Z^2}{K^2\,\sin^4(\theta/2)}
:::

This says three separate things. The number scattered through a given angle grows as the square of the nuclear charge Z, because a stronger charge deflects more strongly. It falls as the square of the alpha's energy K, because a faster alpha is harder to turn. And it falls very steeply with angle, as 1/sin⁴(θ/2), because large deflections require nearly head-on aim, which is rare.

In 1913 Geiger and Marsden tested all three. They measured scattering at angles from 5° to 150°, across which the formula predicts the count should fall by a factor of about 250,000, and they varied the foil material and the alpha energy. The agreement held throughout. That was the real proof: not that a few alphas bounced back, but that the precise pattern of scattering matched what a point-like charged nucleus would produce.

The same measurements gave a rough value for the nuclear charge: about half the atomic weight, in units of the electron's charge. That same year, Henry Moseley, measuring the X-rays emitted by different elements, pinned it down exactly. The nuclear charge Z of an element is its position in the periodic table, its **atomic number**. Hydrogen has 1, helium 2, and gold 79.

## A new problem

Rutherford's atom answered one question and opened a worse one. The electrons, attracted to the nucleus, cannot simply sit still outside it, or they would fall in. They must be moving, presumably orbiting, as planets orbit the Sun. But an orbiting electron is constantly changing direction, which means it is accelerating, and the theory of electromagnetism says an accelerating charge radiates light and loses energy. The calculation is merciless: an orbiting electron should spiral into the nucleus in about 10⁻¹¹ seconds. Every atom in the universe should have collapsed almost instantly. Atoms are plainly stable.

Niels Bohr, a young Danish physicist who spent several months working with Rutherford in Manchester in 1912, took this as a sign that classical physics itself fails inside atoms. His model of 1913, with electrons allowed only in certain orbits where they do not radiate, was the first answer, and it began the long road to quantum mechanics.

::: model atom
The nuclear atom as Rutherford pictured it in 1911. It is not to scale: at the size of this drawing, the real nucleus would be far too small to see. Electrons do not follow orbits either, as later lessons show; treat this as a historical picture, not a modern one.
:::

::: note
A few alpha particles bounce straight back from a thin gold foil, which only a tiny, concentrated charge can do. So the atom's positive charge and nearly all its mass are in a nucleus at most a few times 10⁻¹⁴ m across, while the rest of the atom, about 10⁻¹⁰ m across, is nearly empty space occupied by electrons.
:::

::: quiz
? Why can Thomson's plum pudding atom not send an alpha particle backward?
- Its electrons absorb the alpha particle
+ Its positive charge is spread out, so its electric field is weak everywhere
- It has no positive charge
= A charge spread over the whole atom never produces a strong field, and only a strong field, which requires concentrated charge, can exert enough force to turn a fast, heavy alpha around.
? Alpha particles of twice the energy are fired at the same foil. What happens to the distance of closest approach?
+ It halves
- It doubles
- It stays the same
= The closest approach is d = 2kZe²/K, so doubling the energy K halves the distance: the faster alpha climbs twice as far up the electric hill.
? Why did most alpha particles pass through the foil almost undeflected?
- They were too fast to be affected by electric charge
+ Most of the atom is empty space far from the nucleus
- The electrons shielded them completely
= Only alpha particles passing very close to a nucleus feel a strong force, and since the nucleus is tiny compared with the atom, very few pass that close.
:::
