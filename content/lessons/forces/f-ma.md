The first law says that a net force changes a body's velocity, but it says nothing about how much. A gentle push and a hard shove both change velocity; so does the same push applied to a bicycle and to a lorry. To do physics with forces, rather than just talk about them, we need a rule that turns a force into a definite change of motion. That rule is Newton's second law. It is the most used equation in classical physics, and nearly every problem in mechanics, from a falling stone to a spacecraft's trajectory, begins by writing it down.

## How much does a force change motion?

Start with the experiments that anyone can do roughly with a supermarket trolley, and that can be done precisely with a low-friction track and a timer. Push the trolley with a steady force and it speeds up at a steady rate: it accelerates. Push twice as hard and it accelerates twice as fast. So the acceleration is proportional to the force.

Now load the trolley so that it has twice the mass and push it with the original force. The acceleration halves. Three times the mass gives a third of the acceleration. So for a given force, the acceleration is inversely proportional to the mass. Put the two findings together, choose units so that the constant of proportionality is one, and you have:

::: equation Newton's second law
\sum \vec{F} = m\vec{a}
:::

In words: the net force on a body equals its mass times its acceleration. The equation is short, and each part of it carries more than it seems to.

Start with the left side. The ΣF means the **net force**, the vector sum of every force acting on the body, not any single force. A trolley being pushed forward while friction drags it back accelerates according to the difference between the two. If you write down only the push and forget the friction, the equation will give a wrong answer, and nothing in the algebra will warn you.

Next, the equation is a relation between vectors, quantities with direction as well as size. The acceleration always points in the direction of the net force, which is not necessarily the direction the body is moving. A ball thrown upward is moving up but accelerating down, because the only force on it, gravity, points down. And because vectors can be split into components along chosen axes, the single vector equation is really one equation for each direction: ΣFₓ = maₓ horizontally and ΣF_y = ma_y vertically. This is the reason the motion chapter could treat a projectile's horizontal and vertical motion separately. Gravity has no horizontal component, so the horizontal acceleration is zero, while the vertical motion goes on as if the horizontal one did not exist.

Finally, the equation gives mass a precise meaning. Rearranged as a = F/m, it says that the larger the mass, the smaller the acceleration a given force produces. Mass is the measure of a body's resistance to changes in its velocity, which is exactly what the last lesson called inertia. The first law said inertia exists; the second law says how to measure it.

## What is a newton?

The equation also defines the unit of force. In the SI system, mass is in kilograms and acceleration in metres per second per second, so force comes out in kilogram-metres per second squared. That combination is named the **newton** (N): one newton is the force that gives a 1 kg mass an acceleration of 1 m/s².

To get a feel for it, a newton is about the force with which the Earth pulls on a small apple of roughly 100 g. Holding such an apple still in your palm means pushing up with about one newton. A 70 kg person standing on the floor pushes down with about 690 N.

## Why does everything fall at the same rate?

That last number involves a force we have not yet named. The force of gravity on a body near the Earth's surface is called its **weight**, and it is given by W = mg, pointing straight down, where g is about 9.8 m/s². Weight and mass are easily confused in everyday speech, but they are different quantities. Mass is an amount of inertia, measured in kilograms, and it is the same everywhere. Weight is a force, measured in newtons, and it depends on where you are. A 70 kg person weighs about 690 N on Earth and about 110 N on the Moon, where gravity is roughly a sixth as strong, but has a mass of 70 kg in both places.

With the second law in hand, we can now explain something the free-fall lesson could only report: that without air resistance, every body falls with the same acceleration, whatever its mass. Galileo established this by experiment, but it is far from obvious why it should be so. A bowling ball is pulled down far harder than a tennis ball, so why does it not fall faster?

::: steps Every body falls with the same acceleration
$ \sum F = W = mg
Consider a body falling with no air resistance, so that the only force acting on it is its weight. Taking downward as the positive direction, the net force is simply mg. Nothing else needs to be added, because nothing else touches the body.
---
$ ma = mg
Now apply the second law, which says the net force equals mass times acceleration. The left side is the second law's ma; the right side is the net force we just found. This equation is a statement about this one body and the one force on it.
---
$ a = g
Divide both sides by the mass m, which is allowed because m is not zero. The mass disappears entirely. A heavier body does feel a larger gravitational force, but it has proportionally more inertia resisting that force, and the two effects cancel exactly, leaving the same acceleration g for every body.
:::

The cancellation deserves a second look, because it is less innocent than the algebra makes it seem. The m in W = mg measures how strongly gravity pulls on a body; it is called **gravitational mass**. The m in F = ma measures how strongly a body resists being accelerated; it is called **inertial mass**. There is no obvious reason the two should be the same quantity. One could imagine a world in which gold was pulled harder by gravity, relative to its inertia, than wood was, and in that world gold would fall faster. Experiments have searched for exactly this difference with increasing precision, from Newton's own pendulums filled with different materials to a satellite experiment whose final results appeared in 2022, and found the two masses equal to about one part in 10¹⁵. Newton noticed the coincidence and could not explain it. Einstein took it as the starting point of general relativity, in which gravity is not a force at all but a feature of space and time that every body responds to in the same way.

## What do bathroom scales actually measure?

The second law becomes concrete in a lift. Stand on bathroom scales in a lift. The scales do not measure your mass or even your weight directly; they measure how hard they are pushing up on you. Call that upward push N. Two forces act on you: N upward from the scales and your weight mg downward. Whatever the lift is doing, the second law relates these to your acceleration, which is the same as the lift's.

::: steps What the scales read
$ N - mg = ma
Take upward as the positive direction. The net force on you is the upward push N minus your weight mg, and the second law sets this equal to your mass times your acceleration a, which is the lift's acceleration. When the lift is at rest or moving steadily, a is zero and N equals mg, which is why scales normally read your weight.
---
$ N = m(g + a)
Add mg to both sides to get N on its own, then take out the common factor m. This form shows directly how the reading depends on the lift's motion: an upward acceleration adds to g, and a downward one subtracts from it.
---
$ N = 70 \times (9.8 + 2.0) \approx 830\,\text{N}
Put in numbers for a 70 kg person in a lift accelerating upward at 2.0 m/s². The scales push up with about 830 N instead of the usual 690 N. Scales calibrated in kilograms divide by 9.8 and would show about 85 kg.
:::

When the lift starts downward, a is negative and the reading drops below your normal weight. Something remarkable happens if the cable snaps. Then the lift and everything in it fall freely, a = −g, and the formula gives N = 0: the scales read nothing. Gravity has not switched off, and your weight is just what it always was. The scales read zero because they are falling exactly as fast as you are, so they no longer need to push on you at all. The sensation of weightlessness is the absence of a supporting push, not the absence of gravity. That is exactly the situation of astronauts in orbit, as the circular motion lesson will show.

## How do you solve a mechanics problem?

The lift example follows a pattern that works for every problem in this chapter. First, decide which body you are applying the law to; the law is about one body at a time. Second, list every force acting on that body, and only on that body. The best way is a sketch of the body alone with an arrow for each force, called a **free-body diagram**. Third, choose directions for your axes and decide which way is positive. Fourth, write ΣF = ma along each axis and solve.

Almost all mistakes happen in the second step, not in the algebra. A force left out, such as friction, or a force counted that acts on some other body, leads to a confident wrong answer. The next lesson, on forces that come in pairs, shows how easy the second kind of mistake is to make.

::: history 1687–1750 · Cambridge and Berlin
Newton never wrote F = ma. His second law, in the *Principia* of 1687, said that the change of motion is proportional to the force impressed, and takes place along the line in which that force acts. By "motion" he meant the quantity of motion, mass times velocity, which we now call momentum, and he thought mostly in terms of sudden blows that change it at once. His proofs were geometric, drawn with lines and curves rather than written as equations. The form used today came from Leonhard Euler, the Swiss mathematician working at the Berlin Academy. In a paper written in 1750 he set out force equals mass times acceleration as a law to be applied separately along each of three perpendicular directions, and argued that it governs every part of any mechanical system, from a spinning top to a flowing fluid. That turned Newton's principle into a working tool, and it is Euler's version that appears in textbooks.
:::

::: note
The net force on a body equals its mass times its acceleration. Force decides how the velocity changes, and mass decides how hard it is to change; when gravity supplies the force, the mass cancels and every body falls alike.
:::

::: quiz
? A 2 kg box is pushed with 10 N to the right while friction pulls 4 N to the left. What is its acceleration?
- 5 m/s²
+ 3 m/s²
- 7 m/s²
= The net force is 10 − 4 = 6 N to the right, and the second law gives a = 6 N ÷ 2 kg = 3 m/s². Using the 10 N push alone ignores the friction, which also acts on the box.
? A lift moves upward at a constant 3 m/s. Compared with your weight, the scales read
- more
- less
+ the same
= At constant velocity the acceleration is zero, so N − mg = 0 and the scales read your ordinary weight. Only acceleration changes the reading, not speed.
? Why does a bowling ball not fall faster than a tennis ball, ignoring air?
- Gravity pulls equally on both
+ Gravity pulls harder on the bowling ball, but its greater mass resists acceleration equally more
- The Earth's acceleration cancels the difference
= The pull of gravity is mg and the resistance to acceleration is m, so a = mg/m = g for both. The larger force on the bowling ball is exactly offset by its larger inertia.
:::
