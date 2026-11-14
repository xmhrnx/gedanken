Sit on a train moving at a steady 30 m/s and toss a coin straight up. To you, it rises, stops and falls straight back into your hand. To someone standing on the platform as the train goes by, the same coin traces a long, low arc, moving 30 metres forward every second as it rises and falls. Which description is right? Both are, and the fact that both are right is one of the deepest ideas in physics. Position and velocity have no meaning on their own; they always mean position and velocity relative to some chosen observer. This lesson is about how to move between observers, and about what all observers agree on even when they disagree about speeds.

## Measured from where?

When we measure where something is or how fast it is going, we always measure it from somewhere. A **frame of reference** is the viewpoint from which measurements are made: an origin to measure positions from, a set of directions, and a clock. Usually a frame is attached to some object, such as the ground, a train carriage or a ship, and everything is measured as it would be by someone riding along with that object.

Most of the time the frame goes unmentioned because it is obvious. "The car is doing 100 km/h" means 100 km/h relative to the road. But to the driver of another car overtaking it at 110 km/h, the first car is going backward at 10 km/h, and that is just as true a statement, made from a different frame.

It is tempting to think the ground gives the real velocity and every other frame gives a distorted one. But there is nothing special about the ground. It is attached to an Earth that spins on its axis, carrying a point on the equator around at about 460 m/s, and that orbits the Sun at about 30 km/s, and the Sun itself moves around the centre of the galaxy. There is no frame that is truly at rest, against which all others are measured. Physics has no preferred frame for velocity, and the ground is used only because it is convenient for everyday problems.

## How velocities combine

Since every velocity is relative to something, we need a way to convert between frames. Write v(A rel B) for the velocity of A as measured by B. Suppose we know how fast A moves relative to B, and how fast B moves relative to a third observer, C. Then the velocity of A relative to C is the sum of the two:

::: equation Relative velocity
v_{\text{A rel C}} = v_{\text{A rel B}} + v_{\text{B rel C}}
:::

For the walker on a train, A is the walker, B is the train and C is the ground: the walker's speed over the ground is her speed along the aisle plus the train's speed over the ground. The rule looks obvious, but it rests on an assumption that turned out, much later, to matter, so it is worth deriving properly.

::: steps Why velocities add
$ x_{\text{A rel C}} = x_{\text{A rel B}} + x_{\text{B rel C}}
Start with positions rather than velocities. To get from C to A, you can go first from C to B and then from B to A. Positions measured from an origin are vectors, and vectors add tip to tail, so A's position measured from C is the sum of B's position measured from C and A's position measured from B.
---
$ \frac{\mathrm{d}}{\mathrm{d}t}x_{\text{A rel C}} = \frac{\mathrm{d}}{\mathrm{d}t}x_{\text{A rel B}} + \frac{\mathrm{d}}{\mathrm{d}t}x_{\text{B rel C}}
The position equation holds at every moment, so the rates at which both sides change must also be equal. We therefore take the derivative of both sides with respect to time. The derivative of a sum is the sum of the derivatives, because the rate at which a total grows is the combined rate of its parts. Quietly, this step assumes that every observer measures the same time t, so that all three derivatives are taken against one shared clock.
---
$ v_{\text{A rel C}} = v_{\text{A rel B}} + v_{\text{B rel C}}
The derivative of a position with respect to time is a velocity, so each term becomes the matching velocity. Velocities therefore add in exactly the same way as positions do, which is the rule we wanted.
:::

Two more rules follow directly. First, reversing the viewpoint flips the sign: v(B rel A) = −v(A rel B). If you stand on a pavement and see a car approach you at 20 m/s, the driver sees you approaching at 20 m/s from the opposite direction. Second, in one dimension, the rate at which two objects close on each other is the difference of their velocities measured in a common frame. Two cars driving toward each other at 25 m/s each, one with velocity +25 m/s and the other −25 m/s, close the gap at 25 − (−25) = 50 m/s. That is why head-on collisions are so much more violent than rear-end ones at the same road speeds.

All of these are vector equations. When the velocities do not lie along one line, they must be added as vectors, by components, just as displacements were.

## Crossing a river

A classic example shows the vector version at work. A boat can move at 4 m/s relative to the water. It is crossing a river 80 m wide that flows east at 3 m/s, and the pilot points the boat straight north, directly at the opposite bank. Where does it land, and how long does it take?

The boat's velocity relative to the water is 4 m/s north, and the water's velocity relative to the ground is 3 m/s east. So, by the rule above, the boat's velocity relative to the ground is the vector sum of these two. They are at right angles, so the boat's speed over the ground is √(4² + 3²) = 5 m/s, in a direction about 37° east of north. The current does not slow the boat down; it pushes it sideways, and the boat actually moves faster over the ground than it does through the water.

The crossing time depends only on the northward part of the motion, since that is the part that carries the boat from one bank to the other. That part is 4 m/s, so the crossing takes 80 ÷ 4 = 20 s. Meanwhile the eastward part of the motion carries on independently, and in those 20 seconds the current sweeps the boat 3 × 20 = 60 m downstream. The boat reaches the far bank well east of the point it was aimed at.

To land directly opposite, the pilot must aim partly upstream, so that the westward part of the boat's velocity through the water exactly cancels the current's eastward 3 m/s. The boat's 4 m/s through the water is then the longest side of a right triangle whose other sides are 3 m/s west and the remaining northward speed. Pythagoras gives that northward speed as √(4² − 3²) = √7 ≈ 2.6 m/s. The crossing now takes 80 ÷ 2.6 ≈ 30 s, ten seconds longer, because part of the boat's effort goes into fighting the current rather than crossing.

## What every steady observer agrees on

Observers in different frames disagree about positions and velocities. Is there anything they agree on? There is, provided the frames move steadily relative to one another.

Suppose frame B moves at a constant velocity relative to frame C, as a train on a straight, smooth track does relative to the ground. Then v(B rel C) is a constant. Now take the velocity equation and differentiate it once more with respect to time. Each velocity becomes an acceleration, and the derivative of the constant v(B rel C) is zero, because a constant does not change. What remains is a(A rel C) = a(A rel B). Every observer moving at constant velocity relative to another measures the same acceleration for any object.

This is the reason the coin on the train behaves just as it would on the platform. Both observers see the coin accelerate downward at g; they differ only in the steady horizontal velocity they assign to it, which the passenger sees as zero and the platform observer sees as 30 m/s. And because, as the next chapter will show, forces are tied to accelerations, steadily moving observers agree on forces too. The laws of motion work the same way in each of their frames.

The agreement fails if a frame is speeding up, slowing down or turning. When a train brakes hard, a coin tossed inside lands ahead of the passenger's hand, and a bag on the seat slides forward without anything pushing it. That is the passenger's frame changing its velocity, and it is why the distinction between steady and accelerating frames will matter later.

::: note
Velocity is always measured relative to something, and there is no frame that is truly at rest. To change from one observer to another, add the velocity of one observer relative to the other, as vectors; observers moving steadily relative to each other then agree on every acceleration.
:::

## From a ship's cabin to Einstein

The idea that steady motion cannot be detected from inside first mattered in an argument about whether the Earth moves at all.

::: history 1632 · Florence
Opponents of the idea that the Earth spins and orbits the Sun had a strong-sounding objection: if the ground were rushing along, a stone dropped from a tower should land far behind its base, and we should feel a constant wind. Galileo's reply, in his *Dialogue Concerning the Two Chief World Systems*, was to ask the reader to imagine shutting themselves in the cabin below deck of a large ship, with some flies, a bowl of fish and a bottle dripping water into a jar. Watch carefully while the ship is at anchor, then again while it sails along smoothly at any speed. The flies fly in all directions equally easily, the fish swim as before, the drops fall straight into the jar, and nothing anyone can do in the cabin reveals whether the ship is moving. If steady motion cannot be detected, the objections to a moving Earth fall away. This principle of relativity became central to accepting the new astronomy, and it was built into Newton's mechanics. Nearly three centuries later, in 1905, Albert Einstein kept the principle but challenged the hidden assumption in the derivation above, that all observers share one universal time. Once that goes, velocities near the speed of light no longer simply add, and no combination of them ever exceeds the speed of light. At everyday speeds the simple rule of this lesson remains extremely accurate.
:::

::: quiz
? You walk forward at 1.5 m/s along the aisle of a train moving at 20 m/s. What is your speed relative to the ground?
- 18.5 m/s
+ 21.5 m/s
- 20 m/s
= Your velocity relative to the train and the train's velocity relative to the ground point the same way, so they add: 20 + 1.5 = 21.5 m/s. You would get 18.5 m/s only by walking toward the back of the train.
? Car A drives east at 25 m/s and car B drives west at 15 m/s. What is A's velocity relative to B?
- 10 m/s east
+ 40 m/s east
- 40 m/s west
= Taking east as positive, v(A rel B) = v(A rel ground) − v(B rel ground) = 25 − (−15) = 40 m/s, which is positive and so points east. From B's seat, A is rushing toward it at the sum of the two road speeds.
? A passenger on a smoothly moving train drops a ball. What path does she see it follow?
+ Straight down
- Curving backward
- Curving forward
= In the train's frame the ball starts with no horizontal velocity, and steadily moving frames agree on its downward acceleration, so it falls straight down. An observer on the ground sees the same ball follow a parabola, because to them it also has the train's forward velocity.
:::
