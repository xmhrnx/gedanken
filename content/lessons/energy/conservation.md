Energy can change from one form to another, from motion to height, from height to heat, from chemical to electrical, but the total amount never changes. That principle, the conservation of energy, is one of the most far-reaching statements in physics. Within mechanics, it turns hard problems into easy ones: instead of following forces from moment to moment, you compare the energy at the start with the energy at the end, as if balancing an account. Beyond mechanics, it reaches into heat, chemistry, electricity, light and the nucleus, and no experiment has ever found it broken. This lesson derives the mechanical version from what we already have, shows how much work it saves, and then follows the energy that seems to go missing when friction acts.

## Why should kinetic plus potential energy stay constant?

Suppose the only forces doing work on a body are conservative ones, such as gravity and springs. Other forces may act, as long as they do no work: the tension in a pendulum string, for example, or the push of a smooth surface, both of which act at right angles to the motion. The last lesson gave two facts about work. Combining them gives a conservation law.

::: steps Conservation of mechanical energy
$ W = \Delta K
This is the work–energy theorem from the last lesson: the total work done on a body by all the forces acting on it equals the change in its kinetic energy. It holds for any forces at all, conservative or not.
---
$ W = -\Delta U
Now use the assumption that every force doing work is conservative. For a conservative force, the potential energy was defined so that its change equals minus the work the force does. Adding up over all the forces, the total work is minus the total change in potential energy.
---
$ \Delta K + \Delta U = 0
The two lines give two expressions for the same total work W, so they must be equal: ΔK = −ΔU. Moving ΔU to the left side gives this form, which says that any gain in kinetic energy is exactly matched by a loss in potential energy, and the other way round.
---
$ K_1 + U_1 = K_2 + U_2
Write the changes out in full for any two moments, labelled 1 and 2: ΔK is K₂ − K₁ and ΔU is U₂ − U₁. Rearranging puts everything at moment 1 on one side and everything at moment 2 on the other. The sum K + U, called the **mechanical energy**, has the same value at any two moments, so it never changes.
:::

So as a body falls, its potential energy decreases and its kinetic energy increases by exactly the same amount. As it rises, the exchange runs the other way. What makes this so useful is what the result leaves out. Neither the path the body follows nor the time it takes appears anywhere. Only the starting state and the ending state matter.

## How much work does the bookkeeping save?

Take a frictionless playground slide of height h. A child starts from rest at the top. How fast are they going at the bottom?

Try to answer with forces and the problem quickly becomes difficult. The slide curves, so the angle of the surface changes from point to point, and so does the push of the slide on the child. You would need the exact shape of the slide, and then you would have to add up the effect of gravity along every short stretch of it.

With energy, the answer takes one line. At the top, the child is at rest, so the kinetic energy is zero, and taking the bottom as the zero of height, the potential energy is mgh. At the bottom, the potential energy is zero and the kinetic energy is ½mv². The slide's push is always at right angles to the child's motion, so it does no work and does not enter the account. Setting the two totals equal gives ½mv² = mgh, and so v = √(2gh). This is exactly the speed the child would reach by falling straight down from the same height. The mass has cancelled, so a heavy child and a light one arrive at the same speed. And the shape has dropped out entirely: a steep straight slide and a long gentle twisting one give the same speed at the bottom, though they take different times to get there.

Numbers make this concrete. A roller coaster that starts near rest and drops 30 m reaches about √(2 × 9.8 × 30), roughly 24 m/s or 87 km/h, at the bottom of the drop, less whatever friction and air resistance take.

The same reasoning works upward. A ball thrown straight up at 15 m/s rises until it momentarily stops, at which point all its kinetic energy has become potential energy. Setting ½mv² = mgh gives h = v²/(2g), which is 225/19.6, about 11.5 m. The constant-acceleration equations from the motion chapter give the same answer, but energy gets there in one step.

## How does a pendulum trade energy back and forth?

A pendulum shows the exchange between the two forms continuously. Take a bob of mass m hanging on a string of length L, pull it aside until the string makes an angle θ with the vertical, and let it go. Two forces act on the bob: gravity and the tension in the string. The tension always points along the string, toward the pivot, while the bob moves along a circle around the pivot, so the tension is always at right angles to the motion and does no work. Only gravity does work, and gravity is conservative, so mechanical energy is conserved.

::: steps Speed at the bottom of a pendulum's swing
$ h = L - L\cos\theta = L(1 - \cos\theta)
First find how high the released bob is above the lowest point of its swing. When the string hangs straight down, the bob is a distance L below the pivot. When the string is pulled aside to angle θ, the vertical distance from the pivot down to the bob is the adjacent side of a right triangle with hypotenuse L, which is L cos θ. The difference between these two depths is the height the bob has been raised.
---
$ \tfrac{1}{2}mv^2 = mgL(1 - \cos\theta)
The bob is released from rest, so at the start it has potential energy mgh and no kinetic energy. At the bottom of the swing, measuring heights from there, it has no potential energy left and all the energy is kinetic. Conservation of mechanical energy sets the two equal, with h replaced by the expression from the first step.
---
$ v = \sqrt{2gL(1 - \cos\theta)}
Divide both sides by m, which cancels as it did on the slide, then multiply by 2 and take the square root. For a 1 m pendulum released at 60°, cos 60° is 0.5, so the bob starts 0.5 m above its lowest point and passes through the bottom at about 3.1 m/s.
:::

After passing the bottom, the bob climbs the other side, slowing as kinetic energy turns back into potential energy, until it stops at the same height it started from. Then it swings back. Twice in every complete cycle, all the energy is potential, at the two ends of the swing; twice, it is all kinetic, as the bob passes through the bottom.

::: model pendulum
Watch where the bob moves fastest and where it pauses. Its speed peaks at the bottom, where the potential energy is lowest; at each end of the swing the bob is momentarily still, and all of its energy is potential.
:::

## Where does the energy go when friction acts?

Real pendulums do not swing forever. Each swing is a little lower than the last, and eventually the bob hangs still. A real slide delivers a child more slowly than √(2gh) predicts. In both cases friction or air resistance is doing negative work, and mechanical energy is decreasing. Is energy being destroyed?

It is not, and the evidence is easy to feel. Rub your hands together briskly and they warm up. Brake hard in a car from motorway speed and the brake discs get hot enough to glow in the dark. The mechanical energy that seems to go missing reappears as **thermal energy**: the energy of the random jiggling of atoms and molecules, which we sense as temperature. At the microscopic level, friction takes the orderly motion of a whole sliding body and turns it into disorderly motion of its atoms. Once that thermal energy is counted, the account balances again:

::: equation Energy balance with friction
K_1 + U_1 = K_2 + U_2 + E_{\text{thermal}}
:::

Here E_thermal is the thermal energy produced between moments 1 and 2. The same principle extends to every other form of energy: chemical energy stored in fuel and food, electrical energy, the energy carried by light, and the energy locked in atomic nuclei. Counting all of them, the total energy of an isolated system, one with nothing flowing in or out, never changes.

The law has been tested hard. In the 1920s, measurements of radioactive beta decay, in which a nucleus emits an electron, seemed to show energy disappearing: the electrons came out with a spread of energies, and some always fell short. A few physicists, including Niels Bohr, were prepared to give up exact energy conservation. Wolfgang Pauli instead proposed in 1930 that an unseen, uncharged particle was carrying off the missing energy. That particle, the neutrino, was detected in 1956, and the law survived.

::: history 1842–1850 · Heilbronn, Manchester and Berlin
The general law of energy conservation was found in the 1840s by several people working independently, almost at once. Julius Robert Mayer, a German doctor, noticed while serving as a ship's physician in the tropics that the venous blood of sailors was unusually bright red, as if their bodies needed to burn less fuel to stay warm. That set him thinking about heat and work as two forms of one thing, and in 1842 he argued that they are interconvertible and estimated the exchange rate between them. In Manchester, James Joule, the son of a wealthy brewer, set out to measure that rate directly. In his most famous apparatus, falling weights turned a paddle wheel inside an insulated tank of water, and Joule recorded the tiny rise in the water's temperature with thermometers he could read to a small fraction of a degree. By 1850 his value was within about 1% of the modern figure of 4.184 joules per calorie. In 1847 Hermann von Helmholtz, then a young army surgeon, set out the principle as a law covering all of physics. Mayer, who was first but lacked standing and mathematical polish, was largely ignored for years before his priority was recognised.
:::

## Why is energy conserved at all?

For most of the nineteenth century, energy conservation was treated as a law confirmed by experiment, without any deeper explanation. In 1918 the German mathematician Emmy Noether found one. She proved a general theorem connecting symmetries of the laws of physics with conserved quantities. In particular, if the laws are the same at every moment of time, so that an experiment done today gives the same result as the identical experiment done tomorrow, then there must be a quantity that is conserved, and that quantity is energy. Similar arguments connect the sameness of the laws at every place with conservation of momentum, and their sameness in every direction with conservation of angular momentum, the subjects of the last two lessons in this chapter.

::: note
In an isolated system the total energy is constant. When only conservative forces do work, kinetic plus potential energy stays constant; friction converts mechanical energy into thermal energy, but none of it is destroyed.
:::

::: quiz
? Two frictionless slides start at the same height, one straight and steep, the other long and curving. At the bottom, the child on the curving slide is
- faster
- slower
+ the same speed
= The speed at the bottom depends only on the drop in height, v = √(2gh), because the slide's push does no work. The shape affects how long the trip takes, not the final speed.
? A pendulum released from rest 20 cm above its lowest point swings through. Ignoring friction, how high does it rise on the other side?
- Less than 20 cm
+ 20 cm
- More than 20 cm, because of its momentum
= Mechanical energy is conserved, so the bob rises until all its kinetic energy has become potential energy again. That happens at the same height it started from, 20 cm.
? A box slides across a floor and stops. What happened to its kinetic energy?
- It was destroyed by friction
- It became potential energy
+ It became thermal energy in the box and floor
= Friction converted the box's ordered motion into the random motion of atoms in the box and floor, which is thermal energy. The total energy is unchanged.
:::
