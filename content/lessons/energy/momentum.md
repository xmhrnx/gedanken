Fire a rifle and it slams back into your shoulder. Nothing outside the rifle pushed it backward; the only thing that happened was that it launched a bullet forward. Somehow, sending a small object one way sends the large one the other way, and the effect is precise enough to calculate. This lesson introduces the quantity behind the kick, **momentum**, and shows why it is conserved in every collision, explosion and push, however violent or complicated. It is the second great conservation law of mechanics, and it often succeeds where energy alone cannot, because in a crash energy turns into heat and noise, while momentum is never lost.

## What is momentum, and how does a force change it?

The momentum of a body is its mass times its velocity:

::: equation Momentum
\vec{p} = m\vec{v}
:::

Momentum is a vector, pointing in the same direction as the velocity, and its unit is the kilogram-metre per second, kg·m/s. It combines the two things that make a moving body hard to stop: how much of it there is, and how fast it is going. A slowly rolling truck and a fast bullet can have the same momentum, the truck through its mass and the bullet through its speed.

Momentum was in fact the quantity Newton had in mind in his second law, which he stated as a change of "motion" proportional to the force. The connection is short. For a body whose mass does not change, mass times acceleration is m times the rate of change of velocity, m dv/dt, and that is the same as the rate of change of mv. So the second law can be written F = dp/dt: the net force on a body is the rate at which its momentum changes. This form is the more general one. It still works when mass changes, as for a rocket burning fuel, where F = ma would need care.

Now ask what a force does over a stretch of time rather than at an instant. If a constant net force F acts for a time Δt, the momentum changes at the rate F for that whole time, so the total change is F times Δt:

::: equation Impulse
\vec{F}\,\Delta t = \Delta\vec{p}
:::

The product FΔt is called the **impulse** of the force. When the force varies, as it does in any real collision, the impulse is the area under a graph of force against time, and the same equation holds with that area in place of FΔt.

This equation explains every device designed to cushion a blow. Consider a 70 kg car occupant travelling at 15 m/s, about 54 km/h, when the car crashes. Whatever happens, the occupant ends up at rest, so their momentum must change by 70 × 15, which is 1,050 kg·m/s. That change is fixed by the speed and the mass; nothing can reduce it. What can be chosen is the time over which it happens. Stopped by an airbag that gives way over about 0.1 s, the average force on the occupant is 1,050 / 0.1, about 10,500 N, or 10 kN. Stopped by a rigid dashboard in about 0.01 s, the same change in momentum requires ten times the force, about 100 kN. The airbag does not reduce the change in momentum; it spreads it out in time, and that is what lowers the force. Bending your knees when you land from a jump, or drawing your hands back as you catch an egg, works the same way: the same change in momentum, a longer time, a smaller force.

## Why is momentum conserved?

Now consider two bodies that interact with each other, by colliding, by pushing apart, through a spring or a rope or gravity, while no outside forces act on them. The third law says the forces between them come in an equal and opposite pair. Combined with the second law in momentum form, that gives a conservation law directly.

::: steps Momentum is conserved in an isolated system
$ \vec{F}_{12} = -\vec{F}_{21}
Let F₁₂ be the force on body 1 exerted by body 2, and F₂₁ the force on body 2 exerted by body 1. By Newton's third law, these are equal in size and opposite in direction, at every instant of the interaction. This is true whatever kind of force it is and however it changes during the encounter.
---
$ \frac{\mathrm{d}\vec{p}_1}{\mathrm{d}t} = \vec{F}_{12}, \qquad \frac{\mathrm{d}\vec{p}_2}{\mathrm{d}t} = \vec{F}_{21}
Apply the second law in its momentum form, F = dp/dt, to each body separately. Because no outside forces act, the force each body feels from the other is the whole net force on it. So each body's momentum changes at a rate equal to the force from its partner.
---
$ \frac{\mathrm{d}}{\mathrm{d}t}\left(\vec{p}_1 + \vec{p}_2\right) = \vec{F}_{12} + \vec{F}_{21} = 0
Add the two equations. On the left, the rates of change of the two momenta add to give the rate of change of the total momentum. On the right, the two forces add to zero, because by the first step one is exactly minus the other. The total momentum is therefore not changing at any instant.
---
$ \vec{p}_1 + \vec{p}_2 = \text{constant}
A quantity whose rate of change is zero at every moment stays the same for all time. So the total momentum after the interaction equals the total before, however strong, brief or complicated the forces between the bodies were. We never needed to know anything about those forces except that they obey the third law.
:::

The argument extends to any number of bodies. In a system of many bodies, every internal force has its third-law partner inside the system, and they cancel in pairs, exactly as the action and reaction lesson showed. So the total momentum of an **isolated system**, one on which no outside forces act, is constant. In practice, outside forces are often present but small compared with the huge internal forces of a collision or explosion, which last only a fraction of a second, so momentum is very nearly conserved during the event itself.

## How fast does the rifle recoil?

Now the rifle can be explained exactly. Before firing, the rifle and the bullet are both at rest, so the total momentum of the system is zero. The forces of the exploding powder are internal to the system: they push the bullet forward and the rifle backward with equal and opposite forces. So the total momentum must still be zero after firing. Take a 10 g bullet leaving the barrel at 900 m/s from a 4.0 kg rifle.

::: steps Recoil velocity
$ 0 = m v + M V
Let m and v be the bullet's mass and velocity, and M and V the rifle's. The total momentum after firing is mv + MV, and it must equal the total before firing, which was zero. Taking the bullet's direction as positive, this equation says the two momenta must be equal and opposite.
---
$ V = -\frac{m v}{M}
Subtract mv from both sides and divide by the rifle's mass M to find the rifle's velocity. The minus sign shows it moves in the opposite direction to the bullet, and the ratio m/M shows that the recoil is smaller than the bullet's speed by the same factor that the rifle is heavier than the bullet.
---
$ V = -\frac{0.010 \times 900}{4.0} \approx -2.3\,\text{m/s}
Put in the numbers, with the bullet's mass converted to kilograms. The bullet's momentum is 9 kg·m/s forward, so the rifle must carry 9 kg·m/s backward, which for 4 kg means about 2.3 m/s. That is a firm shove against the shoulder, but nothing like the bullet's speed.
:::

There is a striking asymmetry hiding here. The momenta are equal and opposite, but the kinetic energies are not. The bullet carries ½ × 0.010 × 900², about 4,000 J. The rifle carries ½ × 4.0 × 2.3², about 10 J. Kinetic energy can be written as ½mv² = p²/(2m), so for two bodies with the same size of momentum, the energy is inversely proportional to the mass. The light body gets nearly all of it. That is why the bullet can kill while the kick merely bruises.

## What happens to energy in a collision?

Momentum is conserved in every collision between two bodies, as long as outside forces are small during the impact. Kinetic energy, in general, is not. During a collision, some kinetic energy usually goes into heat, sound, and the permanent bending of metal or squashing of material. Collisions are classified by how much is lost.

In an **elastic** collision, kinetic energy is conserved as well as momentum. No real collision between everyday objects is perfectly elastic, but collisions between hard steel balls or billiard balls come close. One special case is worth knowing. When two equal masses collide head-on elastically, and one was at rest, they exchange velocities: the moving ball stops dead and the struck ball moves off with the original speed. That is what happens in a Newton's cradle, the desk toy of hanging steel balls, where one ball swung in knocks exactly one ball out at the far end.

In a **perfectly inelastic** collision, the bodies stick together and move off as one. This loses as much kinetic energy as momentum conservation allows. Suppose a 1,000 kg car moving at 20 m/s runs into an identical car at rest and their bumpers lock. The total momentum before is 1,000 × 20, which is 20,000 kg·m/s. Afterwards, 2,000 kg moves together, so to keep the same momentum the pair must move at 10 m/s. Now check the energy. Before, it was ½ × 1,000 × 20², which is 200 kJ. After, it is ½ × 2,000 × 10², which is 100 kJ. Half the kinetic energy has gone into crumpling metal, heat and noise, but the momentum is exactly what it was. This is why momentum, not energy, is the tool for working out what happens in a crash.

Momentum conservation reaches far beyond colliding objects. Light carries momentum, enough to push a thin sail in space. In relativity and quantum mechanics the formula p = mv has to be replaced, but the law that total momentum is conserved survives intact.

::: history 1644–1669 · Amsterdam and London
René Descartes, in his *Principles of Philosophy* of 1644, proposed that God had created the universe with a certain total "quantity of motion", mass times speed, and kept it constant forever. It was the right kind of idea, but Descartes ignored direction, counting all motion as positive, and several of his rules for what happens in collisions were plainly contradicted by experiment. In 1668 the Royal Society in London invited papers on the laws of collision. Three answers arrived. John Wallis treated collisions in which the bodies stick together. Christopher Wren and Christiaan Huygens independently treated collisions between hard bodies that rebound. Huygens's analysis was the most complete. He showed that the quantity that is truly conserved is mass times velocity counted with its sign, positive in one direction and negative in the other, so that motions in opposite directions can cancel. And he showed that in collisions between hard bodies, the sum of each mass times its speed squared is conserved too. Momentum and kinetic energy appeared side by side, a century and a half before energy had a name.
:::

::: note
Momentum, p = mv, is conserved in any isolated system, because the forces the bodies exert on one another come in equal and opposite pairs. Kinetic energy is conserved only in elastic collisions; in all others some of it turns into heat and deformation, but the momentum is unchanged.
:::

::: quiz
? A 60 kg skater standing on ice throws a 3 kg ball forward at 10 m/s. What is her recoil speed?
+ 0.5 m/s
- 10 m/s
- 2 m/s
= The total momentum was zero and must stay zero, so the ball's 3 × 10 = 30 kg·m/s forward is matched by the skater's 60 × V backward. That gives V = 0.5 m/s.
? Why does an airbag reduce injury?
- It reduces the passenger's change in momentum
+ It spreads the same change in momentum over a longer time, so the force is smaller
- It absorbs momentum so that total momentum is not conserved
= The passenger must come to rest either way, so the change in momentum is fixed. Since FΔt = Δp, a longer stopping time means a smaller average force.
? Two cars of equal mass collide and stick together. Which is conserved?
- Kinetic energy only
+ Momentum only
- Both momentum and kinetic energy
= Momentum is conserved in every collision when outside forces are small. In a collision where the bodies stick together, some kinetic energy always turns into heat and deformation.
:::
