The ideal gas law is the same for every gas, which suggests it does not depend on what the molecules are, only on how many there are and how they move. Kinetic theory makes that precise. It pictures a gas as a huge number of tiny molecules flying about and bouncing off the walls, and derives the pressure from their collisions. Comparing the result with PV = NkT then reveals what temperature is.

## The model

Assume:

- The gas is N identical molecules, each of mass m, moving in random directions.
- The molecules are tiny compared with the distances between them, and exert no forces except during collisions.
- Collisions with the walls are elastic: a molecule bounces off with the same speed.

Collisions between molecules do happen, constantly, but they conserve momentum and energy, so on average they only shuffle speeds among the molecules without changing the totals. We can ignore them.

## Pressure from collisions

Put the gas in a cubical box of side L, so its volume is V = L³. Fix attention first on one molecule and one wall, the wall perpendicular to the x-axis.

::: steps Deriving the pressure of a gas
$ \Delta p_x = m v_x - (-m v_x) = 2 m v_x
A molecule with x-velocity v_x hits the wall and bounces back with −v_x. Its y- and z-velocities are untouched. The change in its momentum is 2mv_x, and by Newton's third law the wall receives an equal and opposite impulse.
---
$ \Delta t = \frac{2L}{v_x}
After bouncing, the molecule must cross to the far wall and back, a distance 2L at speed v_x, before it hits this wall again. Hits on the side walls do not change v_x.
---
$ F_1 = \frac{\Delta p_x}{\Delta t} = \frac{2mv_x}{2L/v_x} = \frac{m v_x^2}{L}
Force is rate of change of momentum. Averaged over many bounces, this molecule pushes on the wall with force mv_x²/L.
---
$ F = \frac{m}{L}\left(v_{x1}^2 + v_{x2}^2 + \cdots + v_{xN}^2\right) = \frac{N m \langle v_x^2 \rangle}{L}
Add up all N molecules. The sum of the v_x² values is N times their average, written ⟨v_x²⟩.
---
$ P = \frac{F}{L^2} = \frac{N m \langle v_x^2 \rangle}{L^3} = \frac{N m \langle v_x^2 \rangle}{V}
Pressure is force per area, and the wall's area is L². L² times L is the volume.
---
$ \langle v^2 \rangle = \langle v_x^2 \rangle + \langle v_y^2 \rangle + \langle v_z^2 \rangle = 3\langle v_x^2 \rangle
For each molecule, v² = v_x² + v_y² + v_z² by Pythagoras in three dimensions. The motion is random, with no preferred direction, so the three averages are equal.
---
$ PV = \tfrac{1}{3} N m \langle v^2 \rangle = \tfrac{2}{3} N \left( \tfrac{1}{2} m \langle v^2 \rangle \right)
Substitute ⟨v_x²⟩ = ⟨v²⟩/3 and multiply by V. Writing it the second way picks out ½m⟨v²⟩, the average kinetic energy per molecule.
:::

The cube was only for convenience. The result holds for a container of any shape, because pressure is the same on every wall of a gas at rest.

## What temperature is

The derivation gives PV in terms of molecular motion. Experiment, through the ideal gas law, gives PV = NkT. Both describe the same gas.

::: steps Temperature as molecular energy
$ \tfrac{2}{3} N \left( \tfrac{1}{2} m \langle v^2 \rangle \right) = NkT
Set the kinetic-theory expression for PV equal to the ideal gas law.
---
$ \tfrac{1}{2} m \langle v^2 \rangle = \tfrac{3}{2} k T
Cancel N and multiply both sides by 3/2.
:::

::: equation Average kinetic energy of a gas molecule
\tfrac{1}{2} m \langle v^2 \rangle = \tfrac{3}{2} k T
:::

The absolute temperature of a gas is a direct measure of the average kinetic energy of its molecules' random motion. At absolute zero, in this classical picture, that motion would stop. (Quantum mechanics leaves some unavoidable motion even there, but the classical result holds well for gases at ordinary temperatures.) The formula does not contain m on the right: at the same temperature, a heavy molecule and a light one have the same average kinetic energy, so the light one moves faster.

## How fast are the molecules?

Solve for the root-mean-square speed, v_rms = √⟨v²⟩ = √(3kT/m). For nitrogen, the main gas in air, m = 28 × 1.66 × 10⁻²⁷ kg = 4.65 × 10⁻²⁶ kg. At 300 K this gives about 520 m/s, faster than sound in air (which is about 340 m/s; sound is carried by these molecules, so it is a bit slower than they are). Hydrogen molecules, 14 times lighter, move √14 ≈ 3.7 times faster, about 1900 m/s.

::: model gas
Raise the temperature and watch the molecules speed up. Their average kinetic energy, and with it the pressure, grows in proportion to T.
:::

::: history 1738–1860 · Basel, London, Zurich, Aberdeen
Daniel Bernoulli derived Boyle's law from bouncing particles in *Hydrodynamica* in 1738, but the idea went nowhere while heat was thought to be a fluid. In 1845 John James Waterston submitted a full kinetic theory, including the equal sharing of energy between light and heavy molecules, to the Royal Society; a referee called it "nothing but nonsense" and it lay unread in the archives until 1892. Rudolf Clausius published the theory in 1857. In 1860 James Clerk Maxwell went further and worked out the spread of molecular speeds around the average, one of the first laws of physics stated in terms of probability.
:::

::: note
Pressure is the average push of molecules striking the walls, and temperature is proportional to their average kinetic energy: ½m⟨v²⟩ = (3/2)kT.
:::

::: quiz
? A box holds helium and argon at the same temperature. Argon atoms are 10 times heavier. Compared with a helium atom, an average argon atom has
- 10 times the kinetic energy
+ the same kinetic energy, but lower speed
- the same speed
= ½m⟨v²⟩ = (3/2)kT depends only on T. With 10 times the mass, ⟨v²⟩ is 10 times smaller, so v_rms is √10 ≈ 3.2 times smaller.
? In the derivation, why does the time between hits on one wall equal 2L/v_x?
+ The molecule must travel to the opposite wall and back, and only its x-velocity carries it across
- Molecules collide with each other every 2L/v_x seconds
- The molecule's total speed is v_x
= Motion in y and z does not bring it closer to the wall; it covers 2L in the x-direction at speed v_x.
? The kelvin temperature of a gas is quadrupled. The rms speed of its molecules
- quadruples
+ doubles
- stays the same
= v_rms = √(3kT/m), so it grows as √T: √4 = 2.
:::
