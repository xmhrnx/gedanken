A moving hammer can drive a nail into wood. A raised weight, as it slowly descends, can drive the gears of a clock for a week. A stretched spring can launch a dart across a room. In each case something has the ability to do work: to push on something else through a distance. That ability is called **energy**, and it comes in many forms. Mechanics deals with two of them: energy a body has because it is moving, and energy stored because of where things are. This lesson defines both, and shows that they are not arbitrary inventions but follow directly from Newton's laws and the definition of work.

## What does work do to a moving body?

The last lesson defined work but left a question open: when a net force does work on a body, what is the result? A natural guess is that it changes the body's speed, since a force along the motion speeds a body up and a force against it slows it down. To find out exactly how, take the simplest case: a constant net force F acting on a mass m along its direction of motion, over a distance d. Then combine the definition of work with what we already know about constant acceleration.

::: steps The work–energy theorem
$ W = Fd = mad
The net force points along the motion, so the work it does is simply force times distance. By the second law, the net force equals mass times acceleration, so we can replace F with ma. Now the work is written in terms of the body's mass and its motion, rather than in terms of the force.
---
$ v^2 = u^2 + 2ad
This is one of the constant-acceleration equations from the motion chapter, which links the speed u at the start of the distance d to the speed v at the end. It is the right one to use here because it contains a and d but not the time, and time does not appear in our expression for work.
---
$ ad = \tfrac{1}{2}\left(v^2 - u^2\right)
Subtract u² from both sides of that equation and divide by 2, so that the product ad stands alone. This is useful because ad is exactly the combination that appears in the work in the first step.
---
$ W = \tfrac{1}{2}mv^2 - \tfrac{1}{2}mu^2
Replace ad in the first step with this expression and multiply through by m. The force and the distance have disappeared. What remains says that the net work done on the body equals the change in a particular quantity, ½mv², between the start and the end.
:::

That quantity is the body's **kinetic energy**, its energy of motion:

::: equation Kinetic energy
K = \tfrac{1}{2}mv^2
:::

The result, net work equals change in kinetic energy, is called the **work–energy theorem**. We derived it for a constant force, but it holds for forces that vary too: split the path into short pieces over which the force is nearly constant, apply the result to each piece, and add them up. The kinetic energy changes piece by piece, and the total change is the total work.

The theorem gives a new way of looking at motion. A body's kinetic energy is the work that was needed to bring it from rest to its present speed, and it is also the work it can do on something else as it is brought back to rest. That is the sense in which a moving hammer stores the ability to drive a nail. Kinetic energy is measured in joules, like work. Because it depends on v², it depends on speed but not on direction, and it can never be negative.

The square makes speed far more important than mass. A 1000 kg car travelling at 20 m/s (72 km/h) has a kinetic energy of ½ × 1000 × 20², which is 200,000 J, or 200 kJ. At 40 m/s it has ½ × 1000 × 40², which is 800 kJ: twice the speed, four times the energy. When the driver brakes, the friction in the brakes must do enough negative work to remove all of that energy, and with a fixed maximum braking force, four times the energy means four times the distance. That is the deeper reason the friction lesson found braking distance growing with the square of speed.

## Where does the work go when you lift something?

Now consider a different situation. Lift a mass m through a height h at a steady speed. Since the speed does not change, the kinetic energy does not change either, and by the work–energy theorem the total work done on the mass is zero. That makes sense: you push up with a force mg, doing work mgh, while gravity pulls down with mg, doing work −mgh, and the two cancel.

But your work has not simply vanished. Let go of the mass and it falls back through the same height. On the way down, gravity does positive work mgh on it, and by the time it reaches its starting point it has gained exactly mgh of kinetic energy. The work you put in during the lift comes back in full as motion. In between, while the mass sat at the top, the energy was stored. It was not in the mass alone, nor in the Earth alone, but in their arrangement: in the fact that they had been pulled apart against the gravity attracting them. Energy stored in the arrangement of things that exert forces on each other is called **potential energy**. Near the Earth's surface, the gravitational potential energy of a mass m at height h is:

::: equation Gravitational potential energy near the Earth's surface
U = mgh
:::

The size of these energies is easier to grasp with a comparison. How high would you have to lift the 1000 kg car to give it the same 200 kJ it has at 72 km/h? Setting mgh = 200,000 J gives h = 200,000 / (1000 × 9.8), about 20 m, roughly the height of a six- or seven-storey building. A crash at 72 km/h delivers about as much energy as driving off the roof of such a building, which is a sobering way to think about an ordinary road speed.

The height h raises an obvious question: measured from where? The answer is that it does not matter. Only *changes* in potential energy ever enter a calculation, because only changes correspond to work being done. So the zero level can be chosen freely, at the floor, the table top or sea level, whichever makes the problem simplest, and the answers come out the same. The formula mgh itself, though, is valid only near the Earth's surface, where g can be treated as constant. Over large distances gravity weakens, and the general expression is U = −GMm/r, where r is the distance between the centres of the two bodies. That form, which the dark stars lesson used to find escape velocity, puts the zero at infinite separation, so the potential energy is negative everywhere else: you would have to add energy to separate the bodies completely.

Springs store potential energy too. The last lesson found that stretching a spring by x takes work ½kx². Release the spring and it does that same work as it returns to its natural length, so the work is stored, not lost:

::: equation Spring potential energy
U = \tfrac{1}{2}kx^2
:::

## Which forces can store energy?

Not every force can store energy in this way, and the difference is important.

Lift a book from the floor to a shelf 1.5 m up. You can lift it straight up, or carry it up a winding staircase and back down a bit before placing it. In every case the work done against gravity is the same, mgh with h = 1.5 m. The reason is that gravity points straight down, so only vertical movement matters to it. On a winding route, the extra climbing is cancelled by the extra descending, and horizontal movement involves no work against gravity at all. The work depends only on where the book starts and where it ends, not on the path between.

Forces with this property, whose work depends only on the starting and ending points, are called **conservative** forces. Gravity is one; so is the force of an ideal spring. For each conservative force we can define a potential energy, by saying that the change in potential energy equals minus the work done by the force. The minus sign makes the bookkeeping balance: when gravity does positive work on a falling book, the potential energy drops by the same amount that the kinetic energy rises.

Friction is not conservative. Push a box across the floor and back again. Friction opposes you on the way out and opposes you again on the way back, because friction always points against the motion. The work you do against it grows with the total length of the path, and when you return to where you started, none of it comes back to you as motion. So there is no "friction potential energy": there is no stored quantity that depends only on position and can be recovered. The work done against friction does not disappear, though. It turns into heat, as you can feel by rubbing your hands together, and the next lesson follows that energy.

::: history 1686–1740 · Leipzig, Leiden and Paris
For half a century, some of the best minds in Europe argued about how to measure the "force" of a moving body, by which they meant what we would call its capacity to produce effects. Descartes had held that the true measure was the quantity of motion, mass times speed. In 1686 Gottfried Leibniz argued instead for mass times speed squared, which he called *vis viva*, "living force". The dispute was partly philosophical and partly national, with Newton's English followers lining up against Leibniz's. Experiment eventually spoke. In the 1720s Willem 's Gravesande, in Leiden, dropped brass balls into soft clay from different heights. A ball dropped from four times the height arrived at twice the speed, and it made a dent four times as deep, not twice: the effect went with speed squared. Émilie du Châtelet, one of the few people in France who had mastered Newton's mathematics, used the result in her *Institutions de physique* of 1740 to argue for Leibniz's view. In the end both sides were partly right: mass times velocity measures momentum, mass times velocity squared measures energy, and each is conserved in its own way. The factor ½ came later, added by Coriolis in 1829 so that kinetic energy would equal the work needed to produce it.
:::

::: note
The net work done on a body equals its change in kinetic energy, K = ½mv². Work done against a conservative force such as gravity or a spring is not lost but stored as potential energy, and it can be fully recovered as motion.
:::

::: quiz
? A cyclist doubles her speed. Her kinetic energy
- doubles
+ quadruples
- stays the same, since her mass is unchanged
= Kinetic energy is ½mv², so doubling v multiplies it by four. Her brakes would have to remove four times as much energy to stop her.
? A 2 kg book is lifted from the floor to a 1.5 m shelf, once straight up and once along a curving path. How do the changes in potential energy compare?
- The curving path gives more, as it is longer
+ They are equal, about 29 J each
- The straight path gives more
= Gravity is conservative, so the change in potential energy depends only on the change in height: mgh = 2 × 9.8 × 1.5 ≈ 29 J whatever the route.
? Why can friction not be given a potential energy?
- It is too weak
+ The work it does depends on the path, not just the endpoints
- It acts only on moving bodies
= A potential energy requires work that depends only on where a body starts and ends. Friction always opposes the motion, so a longer path means more work against it, and none of that work is returned.
:::
