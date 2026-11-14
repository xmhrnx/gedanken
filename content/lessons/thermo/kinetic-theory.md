The ideal gas law has a curious feature: it is the same for every gas. Helium atoms and carbon dioxide molecules differ in mass, size and shape, yet at the same pressure and temperature a litre of either contains the same number of particles. A law that ignores what the particles are must depend on something they all share. Kinetic theory proposes what that is. It pictures a gas as an enormous number of tiny molecules flying about at random and bouncing off the walls of their container, and it calculates the pressure from those collisions. When the result is set beside the ideal gas law, it reveals what temperature actually is.

## What picture of a gas do we start from?

A model is useful when it leaves out everything that does not matter and keeps what does. Kinetic theory makes three assumptions. First, the gas consists of N identical molecules, each of mass m, moving in random directions with a spread of speeds. Second, the molecules are tiny compared with the distances between them, and they exert no forces on one another except during collisions; between collisions each moves in a straight line at constant speed. Third, when a molecule hits a wall it bounces off elastically, which means it leaves with the same speed it arrived with, so no energy is lost to the wall on average.

What about collisions between molecules? In a real gas they happen constantly, billions of times per second for each molecule. But every such collision conserves momentum and energy. It can speed one molecule up and slow another down, or send one molecule away from a wall just as it sends another toward it, but it cannot change the total momentum or the total energy, and it does not create a preferred direction. On average, then, these collisions only shuffle speeds among the molecules without changing the averages that determine the pressure, and we can leave them out of the calculation.

## Where does pressure come from?

Pressure is force per unit area. A gas pushes on a wall because molecules keep striking it, and each one that bounces back has had its momentum reversed. By Newton's laws, changing an object's momentum requires a force, and the wall must supply that force; by Newton's third law, the molecule pushes back on the wall equally. Each impact is tiny, but there are so many of them that their combined effect is a steady push. To calculate it, put the gas in a cubical box of side L, so its volume is V = L³, and start with a single molecule and a single wall: the one perpendicular to the x-direction.

::: steps Deriving the pressure of a gas
$ \Delta p_x = m v_x - (-m v_x) = 2 m v_x
Momentum is mass times velocity. A molecule moving toward the wall with x-velocity v_x bounces back elastically with x-velocity −v_x, while its velocities along y and z are untouched because the wall pushes only straight out. Its momentum in the x-direction has changed from mv_x to −mv_x, a change of size 2mv_x, and by Newton's third law the wall receives an equal impulse in the opposite direction.
---
$ \Delta t = \frac{2L}{v_x}
After bouncing, the molecule must cross to the far wall and come back before it can hit this wall again: a distance 2L in the x-direction. Only the x-part of its velocity carries it across, and bounces off the side walls reverse v_y or v_z but leave v_x alone. So the time between successive hits on our wall is 2L divided by v_x.
---
$ F_1 = \frac{\Delta p_x}{\Delta t} = \frac{2mv_x}{2L/v_x} = \frac{m v_x^2}{L}
Newton's second law says force is the rate of change of momentum. The molecule delivers momentum 2mv_x to the wall once every 2L/v_x seconds, so its average force on the wall is the ratio of the two. The factors of 2 cancel and v_x appears twice: a faster molecule hits harder and also hits more often.
---
$ F = \frac{m}{L}\left(v_{x1}^2 + v_{x2}^2 + \cdots + v_{xN}^2\right) = \frac{N m \langle v_x^2 \rangle}{L}
Each molecule contributes its own push of the same form, so the total force is the sum over all N molecules. Their x-velocities differ, but a sum of N values equals N times their average. The angle brackets ⟨ ⟩ mean "the average over all molecules", so ⟨v_x²⟩ is the mean of the squared x-velocities.
---
$ P = \frac{F}{L^2} = \frac{N m \langle v_x^2 \rangle}{L^3} = \frac{N m \langle v_x^2 \rangle}{V}
Pressure is force divided by area, and the wall's area is L². Dividing by L² turns the L in the denominator into L³, which is the volume of the box. The pressure is now expressed through the number of molecules, their mass and their motion.
---
$ v^2 = v_x^2 + v_y^2 + v_z^2
The pressure depends only on motion along x, but a molecule's speed v involves all three directions. Its velocity has three perpendicular parts, and by Pythagoras' theorem in three dimensions the square of the speed is the sum of their squares. This holds for each molecule individually.
---
$ \langle v^2 \rangle = \langle v_x^2 \rangle + \langle v_y^2 \rangle + \langle v_z^2 \rangle = 3\langle v_x^2 \rangle
Average the previous line over all molecules. Because the motion is random, with no preferred direction, the averages along x, y and z must be equal, so each is one third of ⟨v²⟩. This is the step that lets us replace a quantity tied to one wall with one that describes the whole gas.
---
$ PV = \tfrac{1}{3} N m \langle v^2 \rangle = \tfrac{2}{3} N \left( \tfrac{1}{2} m \langle v^2 \rangle \right)
Replace ⟨v_x²⟩ with ⟨v²⟩/3 in the pressure formula and multiply both sides by V. Rewriting the right side as two thirds of N times ½m⟨v²⟩ looks like a cosmetic change, but it picks out the average kinetic energy of one molecule, which is the quantity that turns out to matter.
:::

The cube was chosen only to make the geometry easy. The result holds for a container of any shape, because a gas at rest presses equally on every part of its walls; if it did not, the gas would flow until it did. Notice also what the derivation has achieved. It has explained Boyle's law: at fixed temperature, if the molecules' average kinetic energy stays the same, then PV is fixed. Squeeze the gas into half the volume and each molecule hits the walls twice as often, so the pressure doubles.

## What is temperature?

We now have two expressions for PV in the same gas. Kinetic theory gives it in terms of molecular motion. Experiment, through the ideal gas law written for N molecules, gives PV = NkT. Both describe the same thing, so they must be equal, and setting them equal tells us what temperature means at the level of molecules.

::: steps Temperature as molecular energy
$ \tfrac{2}{3} N \left( \tfrac{1}{2} m \langle v^2 \rangle \right) = NkT
Set the kinetic-theory expression for PV equal to the ideal gas law. The left side is a calculation from a model; the right side is a measured fact about real gases. Equating them is a hypothesis that the model is right, and it can be tested by its predictions.
---
$ \tfrac{1}{2} m \langle v^2 \rangle = \tfrac{3}{2} k T
Divide both sides by N, which cancels because both sides are proportional to the number of molecules, and then multiply both sides by 3/2. What remains links one molecule's average kinetic energy to the temperature, with nothing else involved.
:::

::: equation Average kinetic energy of a gas molecule
\tfrac{1}{2} m \langle v^2 \rangle = \tfrac{3}{2} k T
:::

This is the central result of kinetic theory. The absolute temperature of a gas is a direct measure of the average kinetic energy of its molecules' random motion: double the kelvin temperature and you double that energy. It explains why the kelvin scale is the natural one, since zero kelvin corresponds to zero motion in this picture. (Quantum mechanics says that even at absolute zero some unavoidable jiggling remains, but for gases at ordinary temperatures the classical result is excellent.)

Look at what is missing from the right-hand side: the mass of the molecule. At a given temperature, a heavy molecule and a light one have the same average kinetic energy. Since kinetic energy is ½mv², the lighter one must move faster to make up for its smaller mass. This also explains Avogadro's hypothesis. Each molecule, heavy or light, pushes on the walls with the same average energy, so at equal temperature and pressure, equal volumes must hold equal numbers of molecules.

## How fast are the molecules moving?

The formula lets us calculate a typical molecular speed. Solving for √⟨v²⟩, called the **root-mean-square speed** (the square root of the average of the squared speeds), gives v_rms = √(3kT/m).

Take nitrogen, which makes up most of the air. A nitrogen molecule has a mass of about 28 atomic mass units, and one atomic mass unit is 1.66 × 10⁻²⁷ kg, so m ≈ 4.65 × 10⁻²⁶ kg. At 300 K, about room temperature, the formula gives v_rms = √(3 × 1.381 × 10⁻²³ × 300 / 4.65 × 10⁻²⁶) ≈ 520 m/s. That is faster than a passenger jet, and faster than sound, which travels through air at about 340 m/s. The comparison makes sense: sound is a pressure disturbance passed from molecule to molecule by collisions, so it cannot outrun the molecules carrying it, and it travels at a fraction of their typical speed.

Hydrogen molecules are 14 times lighter than nitrogen molecules. With the same average kinetic energy, their speed must be larger by √14 ≈ 3.7, giving about 1900 m/s. That difference has large consequences. A small fraction of molecules in any gas move much faster than average, and for hydrogen and helium enough of them exceed Earth's escape speed over geological time that these light gases have largely leaked away from our atmosphere.

::: model gas
Raise the temperature and watch the molecules speed up. Their average kinetic energy grows in proportion to T, and the pressure on the walls grows with it.
:::

## Why did it take so long to accept?

::: history 1738–1860 · Basel, London, Zurich, Aberdeen
The core idea is old. In 1738 Daniel Bernoulli, in his book *Hydrodynamica*, derived Boyle's law from particles bouncing inside a cylinder. But at the time most scientists believed heat was a fluid, caloric, and a theory in which heat was motion had nowhere to fit. Bernoulli's argument was admired and then forgotten.

A century later John James Waterston, a Scottish engineer working in India, wrote a full kinetic theory, including the result that light and heavy molecules in a mixture share energy equally. He sent it to the Royal Society in London in 1845. A referee dismissed it as "nothing but nonsense", and the paper sat unread in the Society's archives until Lord Rayleigh found it in 1892, long after Waterston's death.

By the 1850s Joule had shown that heat was a form of energy, and the climate had changed. Rudolf Clausius, then in Zurich, published a clear kinetic theory in 1857, and the following year answered an obvious objection raised by a Dutch critic: if molecules move at hundreds of metres per second, why does a smell take minutes to cross a room? Because, he showed, each molecule travels only a tiny distance before colliding with another and changing direction. In 1860 James Clerk Maxwell, then a young professor in Aberdeen, went further and worked out how molecular speeds are spread around the average. It was one of the first laws of physics stated in terms of probability, and it pointed the way to the statistical view of entropy.
:::

::: note
The pressure of a gas is the average push of its molecules striking the walls, and its absolute temperature is proportional to their average kinetic energy: ½m⟨v²⟩ = (3/2)kT. The mass does not appear, so at the same temperature every kind of molecule has the same average energy, and lighter ones move faster.
:::

::: quiz
? A box holds helium and argon at the same temperature. Argon atoms are 10 times heavier. Compared with a helium atom, an average argon atom has
- 10 times the kinetic energy
+ the same kinetic energy, but lower speed
- the same speed
= Average kinetic energy depends only on temperature, so it is the same for both. With 10 times the mass, ⟨v²⟩ must be 10 times smaller, so the argon atom's rms speed is √10 ≈ 3.2 times lower.
? In the derivation, why does the time between hits on one wall equal 2L/v_x?
+ The molecule must travel to the opposite wall and back, and only its x-velocity carries it across
- Molecules collide with each other every 2L/v_x seconds
- The molecule's total speed is v_x
= Motion along y and z does not bring the molecule any closer to the wall. It must cover a distance 2L in the x-direction, and it does so at speed v_x.
? The kelvin temperature of a gas is quadrupled. What happens to the rms speed of its molecules?
- It quadruples
+ It doubles
- It stays the same
= Since v_rms = √(3kT/m), the speed grows as the square root of the temperature. Multiplying T by 4 multiplies the speed by √4 = 2.
:::
