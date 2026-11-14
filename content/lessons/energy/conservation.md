Energy changes form but its total never changes. That single principle, applied as bookkeeping, solves problems that would be painful with forces alone, and it reaches well beyond mechanics into heat, chemistry, electricity and nuclear physics.

## Mechanical energy is conserved

Suppose only conservative forces, such as gravity and springs, do work on a body. The last lesson gave two facts: net work equals the change in kinetic energy, and the work of a conservative force equals minus the change in its potential energy. Put them together.

::: steps Conservation of mechanical energy
$ W = \Delta K
The work–energy theorem: the total work done on the body equals its change in kinetic energy.
---
$ W = -\Delta U
If every force doing work is conservative, the total work is minus the total change in potential energy, by the definition of U.
---
$ \Delta K + \Delta U = 0
Equate the two expressions for W and move everything to one side.
---
$ K_1 + U_1 = K_2 + U_2
Written out for any two moments 1 and 2: the sum of kinetic and potential energy, the **mechanical energy**, stays constant.
:::

As a body falls, potential energy turns into kinetic; as it rises, the reverse. Neither the path nor the time it takes matters, only the start and end states.

## Bookkeeping instead of forces

Take a frictionless slide of height h. A child starts from rest at the top. How fast are they going at the bottom?

With forces, you would need the shape of the slide, the angle at every point and the normal force along the way. With energy: at the top, K = 0 and U = mgh; at the bottom, U = 0. So ½mv² = mgh, giving v = √(2gh), the same speed as falling straight down from that height. The mass cancels and so does the shape. A roller coaster that drops 30 m from near-rest reaches about √(2 × 9.8 × 30) ≈ 24 m/s at the bottom, minus what friction and air take.

The same works upward. A ball thrown up at 15 m/s rises until all its kinetic energy has become potential: h = v²/(2g) = 225/19.6 ≈ 11.5 m, the result the constant-acceleration equations give, in one line.

## The pendulum

A pendulum bob of mass m on a string of length L is pulled aside to angle θ and released. The string's tension is always perpendicular to the bob's motion, so it does no work; only gravity does.

::: steps Speed at the bottom of a pendulum's swing
$ h = L - L\cos\theta = L(1 - \cos\theta)
Height of the released bob above the lowest point. The bob hangs a vertical distance L cos θ below the pivot when displaced, and L below it at the bottom.
---
$ \tfrac{1}{2}mv^2 = mgL(1 - \cos\theta)
Released from rest, all its potential energy at the start has become kinetic energy at the bottom.
---
$ v = \sqrt{2gL(1 - \cos\theta)}
Cancel m and solve. For L = 1 m and θ = 60°, h = 0.5 m and v ≈ 3.1 m/s.
:::

The bob then climbs to the same height on the other side, and keeps swinging between the two, trading one form of energy for the other twice per cycle.

::: model pendulum
Watch where the bob moves fastest and where it pauses. Speed peaks at the bottom, where potential energy is lowest; at each end the bob is momentarily still and all the energy is potential.
:::

## When friction acts

A real pendulum slowly dies down and a real slide delivers a slower child. Friction and drag do negative work, and mechanical energy is lost. But it is not destroyed. Rub your hands together and they warm; brake hard and the brake discs can glow. The missing mechanical energy has become **thermal energy**, the random jiggling of atoms. Including it, the total is still constant:

::: equation Energy balance with friction
K_1 + U_1 = K_2 + U_2 + E_{\text{thermal}}
:::

Counting every form, kinetic, potential, thermal, chemical, electrical, radiant and nuclear, the energy of an isolated system never changes. No experiment has ever found an exception.

::: history 1842–1850 · Heilbronn, Manchester and Berlin
The general law was found in the 1840s by several people almost at once. Julius Robert Mayer, a German ship's doctor, argued in 1842 that heat and motion are interconvertible and estimated the exchange rate. In Manchester, James Joule, a brewer's son, measured it directly: falling weights turned a paddle wheel in an insulated tank of water, and he recorded the tiny temperature rise. By 1850 he had a value within about 1% of the modern 4.184 joules per calorie. Hermann von Helmholtz set out the principle for all of physics in 1847. Mayer, first and least credentialled, was largely ignored for years.
:::

Why should energy be conserved at all? In 1918 Emmy Noether proved that it follows from a symmetry: the laws of physics are the same today as tomorrow. Every such continuous symmetry gives a conserved quantity.

::: note
In an isolated system the total energy is constant. With only conservative forces, kinetic plus potential energy is constant; friction converts mechanical energy into thermal energy, but none is lost.
:::

::: quiz
? Two frictionless slides start at the same height, one straight and steep, the other long and curving. At the bottom, the child on the curving slide is
- faster
- slower
+ the same speed
= Speed at the bottom depends only on the drop in height: v = √(2gh). The shape affects the time taken, not the final speed.
? A pendulum released from rest at height 20 cm above its lowest point swings through. Ignoring friction, how high does it rise on the other side?
- Less than 20 cm
+ 20 cm
- More than 20 cm, because of its momentum
= Mechanical energy is conserved, so the bob returns to the same height, where its kinetic energy is again zero.
? A box slides across a floor and stops. What happened to its kinetic energy?
- It was destroyed by friction
- It became potential energy
+ It became thermal energy in the box and floor
= Friction converts mechanical energy into heat. The total energy is unchanged.
:::
