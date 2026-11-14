An electron behaves like a tiny magnet, and the strength of that magnet is the same for every electron, always. It comes from an angular momentum the electron carries even when it is not moving or orbiting anything. This **spin** has no good classical picture, and the experiment that first revealed it was looking for something else.

## A magnet in an uneven field

A small magnet is characterised by its **magnetic moment** μ, a vector pointing from its south pole to its north pole. In a uniform magnetic field it only twists. In a field that grows stronger in one direction, it is also pushed.

::: steps The force on a magnet in a non-uniform field
$ U = -\vec{\mu} \cdot \vec{B}
The energy of a magnetic moment in a field B is lowest when the moment lines up with the field.
---
$ F_z = -\frac{\partial U}{\partial z} = \mu_z\,\frac{\partial B_z}{\partial z}
A force pushes toward lower energy. If the field points along z and grows along z, only the z-component of μ matters.
---
$ \Delta z \propto \mu_z
Atoms passing through such a field are deflected up or down by an amount proportional to μ_z, the component of their moment along the field.
:::

If atomic magnets pointed in random directions, as classical physics expected, μ_z would take every value from −μ to +μ. A beam of atoms sent through the field would spread into a continuous smear.

::: history 1922 · Frankfurt
Otto Stern and Walther Gerlach sent a beam of silver atoms from an oven through a carefully shaped magnet and onto a glass plate. They were testing a prediction of the old quantum theory that an atom's orbital angular momentum could point only in certain directions. The silver deposit was too thin to see until it was exposed to fumes, by one account the sulphurous smoke of Stern's cheap cigars, which blackened it. The beam had split into two separate lines, with nothing in between. They took it as confirmation of the old theory. In fact silver's outer electron has no orbital angular momentum at all; the two lines were the first sign of electron spin, which nobody yet had in mind.
:::

## Two values only

In 1925 two Dutch students, George Uhlenbeck and Samuel Goudsmit, proposed that the electron has its own intrinsic angular momentum. Measured along any axis, its component is always one of two values.

::: equation Electron spin along any axis
S_z = \pm\frac{\hbar}{2}, \qquad \mu_z \approx \mp\,\mu_B, \qquad \mu_B = \frac{e\hbar}{2m_e} \approx 9.27 \times 10^{-24}\,\text{J/T}
:::

The two states are called **spin up** and **spin down**. The magnetic moment points opposite to the spin because the electron's charge is negative, and its size is very nearly one Bohr magneton μ_B. A silver atom's other 46 electrons pair off with opposite spins and cancel, so the one outer electron sets the atom's moment: two values, two lines.

Spin is not a little ball rotating. Hendrik Lorentz pointed out to Uhlenbeck that a spinning electron of its estimated size would need a surface moving faster than light to produce that angular momentum; experiments since show no size for the electron at all. The name stayed, but spin is simply a property particles have, like charge. Protons and neutrons also have spin ½ in units of ħ. Photons have spin 1.

## Measuring along different axes

Spin's quantum strangeness shows up when Stern–Gerlach magnets are chained together.

- Pass atoms through a magnet along z and keep only the up beam. A second z magnet sends all of them up. The state is definite.
- Send the z-up beam through a magnet along x instead. Half go each way.
- Now take the x-up atoms and measure z again. Half come out down. Measuring x erased the earlier z result.

S_x and S_z cannot both be definite at once, just as position and momentum cannot. An electron with spin up along z is in a superposition of up and down along x. For a magnet tilted at angle θ from the direction the spin was prepared in:

::: equation Probability of spin up along an axis at angle θ
P(\uparrow) = \cos^2\frac{\theta}{2}, \qquad P(\downarrow) = \sin^2\frac{\theta}{2}
:::

At θ = 0 the result is certain. At θ = 90° it is 50–50. At θ = 180° the spin is certain to come out "down", which is the same direction. This rule returns in the next lesson.

::: note
Spin is angular momentum built into a particle, not produced by any motion. Measured along any axis, an electron's spin is always +ħ/2 or −ħ/2, and measuring one axis scrambles the others.
:::

::: quiz
? Classically, what did Stern and Gerlach expect to see on their plate?
- Two separate lines
+ A continuous smear, because magnets could point in any direction
- One line, undeflected
= A continuous range of μ_z gives a continuous range of deflections.
? Atoms selected as spin up along z go through a magnet along x. What fraction are deflected up?
- All of them
- None of them
+ Half of them
= At θ = 90°, P(↑) = cos²(45°) = ½.
? Why is spin not literally a spinning sphere?
- Electrons do not have angular momentum
+ To carry ħ/2 at any plausible size, its surface would have to move faster than light
- Spinning objects cannot have magnetic moments
= Lorentz's objection: the required rotation speed exceeds c, and the electron shows no measurable size.
:::
