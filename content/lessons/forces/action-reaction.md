The first two laws are about a single body: what it does when no net force acts, and how it accelerates when one does. They say nothing about where forces come from. The third law fills that gap. It says that forces are never one-sided. Whenever one body pushes or pulls on another, the second pushes or pulls back on the first, equally hard and in the opposite direction. The law sounds simple, and it is easy to state, but it is also the most often misapplied of the three, and seeing why it does not lead to paradox is the best way to understand it.

## What does the law say?

If body A exerts a force on body B, then B exerts a force on A that is equal in size and opposite in direction.

::: equation Newton's third law
\vec{F}_{\text{on }A\text{ by }B} = -\vec{F}_{\text{on }B\text{ by }A}
:::

The two forces are traditionally called the action and the reaction, but the names suggest an order that does not exist. Neither force comes first and causes the other. They are two aspects of a single interaction between two bodies, and they arise and vanish together. When you lean on a wall, your hand pushes the wall and the wall pushes your hand at exactly the same moments.

Three features of the law are worth spelling out, because each one blocks a common mistake.

First, the two forces act on **different bodies**: one on A, one on B. That is why they never cancel each other. Forces cancel only when they act on the same body, because cancelling means their effects on that body's motion add up to nothing. A force on A and a force on B affect two different motions, so there is nothing for them to cancel in.

Second, the two forces are always of the **same kind**. If the Earth pulls the Moon by gravity, the Moon pulls the Earth by gravity. If your foot pushes the ground by contact, the ground pushes your foot by contact. A gravitational force is never paired with a contact force.

Third, the two forces are equal **at every instant**, not just on average and not just when things are still. During a collision lasting a thousandth of a second, the forces between the two bodies rise and fall together, matched at each moment.

## If the forces are equal, how does anything move?

There is an old puzzle that seems to show the law must be wrong. A horse pulls a cart. By the third law, the cart pulls back on the horse just as hard as the horse pulls forward on the cart. So the two forces are equal and opposite. How can the horse and cart ever start moving?

The puzzle dissolves as soon as you apply the method from the last lesson: pick one body at a time and list only the forces acting on it.

Take the cart first. Two horizontal forces act on it: the horse's pull forward, through the harness, and friction from the road and axles backward. If the forward pull is larger than the friction, the net force on the cart is forward and the cart accelerates. Nothing else matters to the cart.

Now take the horse. Two horizontal forces act on it too: the cart's pull backward, through the harness, and the push of the ground forward on its hooves. The horse drives its hooves backward against the ground, and by the third law the ground pushes forward on the horse. If that forward push is larger than the cart's backward pull, the horse accelerates.

The equal and opposite pair, horse-on-cart and cart-on-horse, appears once in each list: one force on the cart's diagram, the other on the horse's. Since they never appear together on the same diagram, they cannot stop either body. What actually decides whether the horse and cart move is a comparison between different forces: the ground's grip on the horse against the road's drag on the cart.

The same reasoning explains how you walk. Your leg muscles cannot push your body forward directly, because muscles pull one part of your body against another, and those forces stay inside you. Instead, your foot pushes backward against the ground. By the third law, the ground pushes forward on your foot, and that forward push, a friction force between sole and floor, is the external force that accelerates you. Try to start walking on perfectly smooth ice and you discover this the hard way: with no friction, your foot cannot push backward on the ground, so the ground cannot push forward on you, and you stay where you are.

## Are weight and the table's push a pair?

The most common error with the third law is to pair up two forces simply because they happen to be equal and opposite. A book resting on a table makes the point.

Two forces act on the book. One is its weight, the Earth's gravitational pull downward. The other is the upward push of the table, called the **normal force**, where "normal" means perpendicular to the surface. These two forces are equal in size and opposite in direction, so it is tempting to call them a third-law pair. They are not. They act on the same body, the book, and they are different kinds of force, one gravitational and one a contact push. They are equal for a different reason entirely: the book is not accelerating, so by the second law the net force on it must be zero. If you pressed down on the book, the table's push would grow while the weight stayed the same, and the two would no longer be equal. A true third-law pair is equal always, whatever else happens.

So where are the real partners? The weight is the Earth pulling the book by gravity, so its partner is the book pulling the Earth by gravity, upward, with the same strength. The table's push on the book is a contact force, so its partner is the book's push down on the table.

That first pairing has a striking consequence. The Earth really is pulled toward every object on it, with exactly the same force as the object's weight. It just responds very little, as a quick calculation shows.

::: steps How much does the Earth move when you jump?
$ F = mg = 70 \times 9.8 \approx 690\,\text{N}
While a 70 kg person is in the air after jumping, the only force on them is their weight. The Earth pulls them downward with about 690 newtons, and they accelerate down at g.
---
$ F_{\text{on Earth}} = 690\,\text{N} \text{ upward}
The weight is a gravitational force exerted by the Earth on the jumper. By the third law, the jumper exerts a gravitational force on the Earth, of the same size and in the opposite direction. So the Earth is pulled upward, toward the jumper, with about 690 newtons.
---
$ a_{\text{Earth}} = \frac{F}{M} = \frac{690}{5.97 \times 10^{24}} \approx 1.2 \times 10^{-22}\,\text{m/s}^2
Apply the second law to the Earth, whose mass is about 5.97 × 10²⁴ kg. The force is the same as on the jumper, but the mass is about 10²³ times larger, so the acceleration is about 10²³ times smaller. The Earth does move toward the jumper, but by far less than the width of an atomic nucleus during the jump.
:::

The forces are equal; the effects are not. That pattern, equal forces producing very different accelerations because the masses differ, recurs throughout mechanics, and it is the key to the first quiz question below.

## What happens when you add up a whole system?

The third law has a consequence that reaches well beyond pairs of bodies. Instead of one body, think of a group of bodies together and call it a **system**: a car with its passengers, a rocket with its fuel, the Earth and the Moon. Forces between members of the system are called internal; forces from anything outside are external.

Every internal force comes with its third-law partner, and the partner is internal too, since both bodies are inside the system. So when you add up all the forces on all the parts of the system, every internal force is matched by an equal and opposite one, and they cancel in pairs. Only the external forces remain. This means a system as a whole can accelerate only because of forces from outside it. You cannot lift yourself by pulling on your own bootstraps, and a person sitting in a car cannot make it move by pushing on the dashboard: the push on the car is cancelled by the dashboard's push on the person, both inside the same system. This idea, that internal forces cancel, is the root of the conservation of momentum, the subject of a later lesson.

A rocket might seem to break this rule, but it uses it. A rocket engine throws exhaust gas backward at high speed. To do that it must push on the gas, and by the third law the gas pushes forward on the rocket, just as hard. That forward push is what drives the rocket. Nothing outside is needed to push against, so a rocket works in the vacuum of space. In fact it works slightly better there than in air, because there is no surrounding air pressure pushing back against the exhaust as it leaves the nozzle.

::: history 1687 and 1920 · Cambridge and New York
Newton illustrated the third law in the *Principia* with everyday examples. If you press a stone with your finger, your finger is pressed by the stone. If a horse pulls a stone tied to a rope, the horse is pulled back toward the stone just as hard. He also tested the law by letting pendulum bobs swing into each other and measuring how they rebounded. More than two centuries later the law still caught out intelligent people. In January 1920 an editorial in the *New York Times* mocked the rocket pioneer Robert Goddard for claiming that a rocket could fly in space, reasoning that in a vacuum there would be nothing for it to push against, and suggesting he lacked "the knowledge ladled out daily in high schools". The mistake was the editor's: the rocket pushes against its own exhaust. On 17 July 1969, with the Apollo 11 astronauts on their way to the Moon, the paper published a brief correction, noting that the error was now regretted.
:::

The third law, in the form stated here, is not exact in every circumstance. Between moving electric charges, the forces each exerts on the other can be unequal for a time, because the electromagnetic field between them itself carries momentum. The deeper principle, conservation of momentum, still holds once the field's share is counted. For everything in this chapter, the law as Newton stated it is all that is needed.

::: note
Forces come in pairs: equal in size, opposite in direction, of the same kind, and acting on two different bodies. Because they act on different bodies they never cancel each other; they cancel only when you add up the forces on a whole system that contains both bodies.
:::

::: quiz
? A truck collides head-on with a small car. During the collision, which feels the larger force?
- The car
- The truck
+ Neither: the forces are equal
= The force of the truck on the car and the force of the car on the truck are a third-law pair, so they are equal in size. The car comes off worse because the same force gives its smaller mass a much larger acceleration.
? A book rests on a table. What is the third-law partner of the book's weight?
- The table's upward push on the book
+ The book's gravitational pull on the Earth
- The table's weight
= The weight is the Earth's gravitational pull on the book, so its partner is the book's gravitational pull on the Earth. The table's push acts on the same body and is a different kind of force; it balances the weight only because the book is not accelerating.
? What pushes a person forward when they start walking?
- Their leg muscles, directly
+ The ground, through friction on the foot
- Air pressure behind them
= The foot pushes backward on the ground, and the ground pushes forward on the foot with an equal friction force. The muscles produce internal forces, which cancel in pairs and cannot accelerate the body as a whole.
:::
