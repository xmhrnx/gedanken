Every electron behaves like a tiny bar magnet, and the strength of that magnet is exactly the same for every electron in the universe, always. It does not come from the electron moving around a nucleus, though that motion produces magnetism too. It comes from an angular momentum, a quantity of rotation, that the electron carries simply by existing, even when it is not moving at all. This property is called **spin**. It has no faithful picture in classical physics, it takes only two values when measured, and it is the reason for the structure of the periodic table, the behaviour of magnets, and the workings of the MRI scanner. The experiment that first revealed it was looking for something else, and its discoverers misread their own result.

## A magnet in an uneven field

To see how spin was found, start with how a small magnet behaves in a magnetic field. A magnet's strength and direction are described by its **magnetic moment** μ, an arrow pointing from its south pole to its north pole. In a uniform field, one with the same strength and direction everywhere, a compass needle twists to line up with the field but is not pushed anywhere, because the forces on its two poles are equal and opposite. In a field that grows stronger in one direction, the pole sitting in the stronger part is pulled harder than the other, and the magnet as a whole feels a push. How hard, and which way, depends on which way the magnet points.

::: steps The force on a magnet in a non-uniform field
$ U = -\vec{\mu} \cdot \vec{B}
A magnetic moment μ in a magnetic field B has an energy U given by minus their dot product, which is the product of their sizes times the cosine of the angle between them. So the energy is lowest when the moment points along the field and highest when it points against it. This is why a compass needle swings to point north.
---
$ F_z = -\frac{\partial U}{\partial z}
A force always pushes toward lower energy, and its size along any direction is the rate at which the energy falls in that direction. Here we choose z as the direction in which the field changes, so the force along z is minus the rate of change of U with z.
---
$ F_z = \mu_z\,\frac{\partial B_z}{\partial z}
Suppose the field points mainly along z and its strength grows along z. Then U ≈ −μ_z B_z, and since the magnet's moment does not change as it moves, differentiating gives the force as μ_z times the field's gradient. Only the component of the moment along the field, μ_z, matters; the sideways components contribute nothing.
---
$ \Delta z \propto \mu_z
Atoms passing through such a field at the same speed spend the same time in it, so each is deflected up or down by an amount proportional to the force on it, and hence to μ_z. A plate placed after the magnet therefore records, atom by atom, the component of each atom's magnetic moment along the field.
:::

Now ask what such an experiment should show. If atomic magnets point in random directions, as classical physics expects for atoms boiling out of an oven, then μ_z takes every value between −μ and +μ, where μ is the full size of the moment. Some atoms point straight up the field, some straight down, most somewhere in between. A beam of them sent through the magnet should spread into a continuous smear on the plate.

::: history 1922 · Frankfurt
Otto Stern and Walther Gerlach set out to test a prediction of the old quantum theory of Bohr and Sommerfeld, which said that an atom's orbital angular momentum could point only in certain directions relative to a field, a rule called space quantisation. Many physicists, Stern among them, doubted it was literally true. They heated silver in an oven, let a thin beam of atoms escape through slits, sent it through a magnet with one pointed pole to make the field strongly uneven, and caught it on a glass plate. The deposit was too thin to see. By Stern's later account, it appeared only when he breathed on the plate; he was a heavy smoker of cheap cigars, and the sulphur in his breath turned the silver into black silver sulphide. The beam had split cleanly into two separate lines, with nothing in between. Stern and Gerlach took this as confirmation of space quantisation. In fact they had seen something else entirely. Silver's single outer electron has no orbital angular momentum, so the old theory's prediction did not apply to it. The two lines were the first direct sign of electron spin, a property nobody yet had in mind.
:::

## Two values only

The correct explanation arrived three years later, from a different direction. Spectroscopists had long been puzzled by lines that split into pairs, such as the famous yellow doublet of sodium, and by the way lines split in magnetic fields in patterns the orbital theory could not reproduce. Wolfgang Pauli had concluded early in 1925 that electrons needed a fourth quantum number that could take only two values, but he offered no picture of what it was.

In 1925 two Dutch graduate students in Leiden, George Uhlenbeck and Samuel Goudsmit, proposed that the electron carries its own built-in angular momentum, as if it were spinning on an axis. Measured along any direction, the component of this angular momentum is always one of exactly two values.

::: equation Electron spin along any axis
S_z = \pm\frac{\hbar}{2}, \qquad \mu_z \approx \mp\,\mu_B, \qquad \mu_B = \frac{e\hbar}{2m_e} \approx 9.27 \times 10^{-24}\,\text{J/T}
:::

Read the equation piece by piece. S_z is the spin's component along the chosen axis, and it is either plus or minus half of ħ, the reduced Planck constant. The two states are called **spin up** and **spin down**. The electron's magnetic moment along the same axis is very nearly one **Bohr magneton**, μ_B, a natural unit of atomic magnetism built from the electron's charge e and mass m_e. The moment points opposite to the spin because the electron's charge is negative: a negative charge rotating one way produces the same magnetism as a positive charge rotating the other.

This explains Stern and Gerlach's plate. A silver atom has 47 electrons. Forty-six of them pair off, each pair with one spin up and one spin down, so their magnetism cancels. The one remaining outer electron sets the atom's whole magnetic moment, and it has only two possible values of μ_z. Two values, two lines.

## Not a spinning ball

The name suggests a tiny ball rotating on its axis, and that picture is wrong. Uhlenbeck and Goudsmit took their idea to Hendrik Lorentz, the most respected theorist in the Netherlands, who worked out what it would take. To carry angular momentum ħ/2 with the size then estimated for the electron, its surface would have to move many times faster than light. The students asked their supervisor, Paul Ehrenfest, to withdraw the paper; he told them it had already been sent, and that they were young enough to afford a stupidity. As it turned out, they were right about the angular momentum and wrong only about the picture. A year earlier Ralph Kronig had had the same idea, been talked out of it by Pauli, and never published.

Experiments since then have found no size for the electron at all, down to distances far smaller than any atomic nucleus, so there is nothing to rotate. The name stuck, but spin is best thought of as a basic property a particle has, like its charge or mass, which happens to behave mathematically like angular momentum and to carry a magnetic moment. Other particles have it too. Protons and neutrons, like electrons, have spin ½ in units of ħ. Photons have spin 1. The Higgs boson has spin 0.

## Measuring along different axes

Spin's quantum strangeness shows up most clearly when Stern–Gerlach magnets are chained together, each one passing on only one of its two beams.

Start by sending atoms through a magnet whose field points along z, and keep only the beam deflected up. Send those atoms through a second magnet, also along z, and every one of them goes up again. Nothing surprising: the first magnet selected atoms with spin up along z, and they still have it.

Now send the z-up atoms through a magnet turned to point along x, at right angles. Half go one way and half the other, at random. That might seem to mean the atoms had a definite spin along x all along that we had not yet looked at. The third stage rules that out. Take only the atoms that came out x-up and measure them along z again. Half now come out z-down, although every one of them was selected as z-up two magnets earlier. Measuring along x has erased the earlier z result.

The conclusion is that spin along x and spin along z cannot both be definite at once, just as position and momentum cannot. An electron with spin up along z is not secretly up or down along x; it is in a superposition of the two, and measuring x forces one result and destroys the z information. For a magnet tilted at any angle θ from the direction along which the spin was prepared, quantum mechanics gives the probabilities exactly:

::: equation Probability of spin up along an axis at angle θ
P(\uparrow) = \cos^2\frac{\theta}{2}, \qquad P(\downarrow) = \sin^2\frac{\theta}{2}
:::

Check the rule against the cases above. At θ = 0, the second magnet is aligned with the first, cos² 0 = 1, and the result is certain. At θ = 90°, the x magnet, cos² 45° = ½, and the result is an even split. At θ = 180°, the magnet is turned upside down, P(↑) = 0, and every electron comes out "down" relative to that magnet, which is the same physical direction as its original "up". The angle is halved inside the cosine, a hint that spin is not an ordinary arrow in space; this rule returns in the next lesson, where it sets the strength of the correlations between entangled particles.

::: note
Spin is angular momentum built into a particle, not produced by any motion. Measured along any axis, an electron's spin is always +ħ/2 or −ħ/2, and measuring it along one axis scrambles its value along the others.
:::

::: quiz
? Classically, what did Stern and Gerlach expect to see on their plate?
- Two separate lines
+ A continuous smear, because atomic magnets could point in any direction
- A single undeflected line
= If the moments point every which way, μ_z takes every value from −μ to +μ, and since deflection is proportional to μ_z the atoms should land in a continuous band.
? Atoms selected as spin up along z pass through a magnet along x. What fraction are deflected up?
- All of them
- None of them
+ Half of them
= The x axis is at θ = 90° to z, so P(↑) = cos²(45°) = ½: the z-up state is an equal superposition of up and down along x.
? Why is spin not literally a spinning sphere?
- Because electrons do not have angular momentum
+ Because to carry ħ/2 at any plausible size, its surface would have to move faster than light
- Because spinning objects cannot have magnetic moments
= This was Lorentz's objection: the required rotation speed exceeds the speed of light, and experiments have since found no measurable size for the electron to rotate.
:::
