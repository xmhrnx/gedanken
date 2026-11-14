You sit on a train moving at 30 m/s and toss a coin straight up. To you, it rises and falls back into your hand. To someone on the platform, the coin traces a long arc and travels 30 m forward every second. Both descriptions are correct. Velocity, and position, only have meaning relative to some chosen observer.

## Frames of reference

A *frame of reference* is a viewpoint from which measurements are made: an origin, a set of directions, and a clock, usually attached to some object such as the ground, a train or a ship. "The car is doing 100 km/h" means 100 km/h relative to the road. Relative to another car overtaking it, it might be going backward.

Physics has no preferred frame for describing velocity. The ground is convenient for everyday problems, but the Earth spins at about 460 m/s at the equator and orbits the Sun at about 30 km/s.

## Adding velocities

Write v(A rel B) for the velocity of A as measured by B. If you know how fast A moves relative to B, and how fast B moves relative to C, you get A relative to C by adding:

::: equation Relative velocity
v_{\text{A rel C}} = v_{\text{A rel B}} + v_{\text{B rel C}}
:::

The derivation is just about positions:

::: steps Why velocities add
$ x_{\text{A rel C}} = x_{\text{A rel B}} + x_{\text{B rel C}}
To get from C to A, first go from C to B, then from B to A. Positions are vectors, and they add tip to tail.
---
$ \frac{\mathrm{d}}{\mathrm{d}t}x_{\text{A rel C}} = \frac{\mathrm{d}}{\mathrm{d}t}x_{\text{A rel B}} + \frac{\mathrm{d}}{\mathrm{d}t}x_{\text{B rel C}}
Take the rate of change of both sides. The derivative of a sum is the sum of the derivatives. This step assumes all observers agree on the time t.
---
$ v_{\text{A rel C}} = v_{\text{A rel B}} + v_{\text{B rel C}}
Each derivative of a position is the corresponding velocity.
:::

Two further rules follow. Reversing the viewpoint flips the sign: v(B rel A) = −v(A rel B). If you see a car approach at 20 m/s, the driver sees you approach at 20 m/s from the other direction. And in one dimension, the speed at which two objects approach each other is the difference of their velocities: two cars driving toward each other at 25 m/s each close at 50 m/s.

These are vector equations. When the velocities are not along one line, add them by components.

## Example: crossing a river

A boat can move at 4 m/s relative to the water. A river 80 m wide flows east at 3 m/s. The boat points straight north across.

- Velocity of the boat relative to the water: 4 m/s north.
- Velocity of the water relative to the ground: 3 m/s east.
- Velocity of the boat relative to the ground: their vector sum. Its magnitude is √(4² + 3²) = 5 m/s, directed about 37° east of north.

Time to cross depends only on the northward component: 80 ÷ 4 = 20 s. In that time the current carries the boat 3 × 20 = 60 m downstream. To land directly opposite, the pilot must aim partly upstream so that the eastward parts cancel; then the crossing speed is √(4² − 3²) = √7 ≈ 2.6 m/s, and the crossing takes longer.

## Accelerations agree

If frame B moves at constant velocity relative to frame C, then v(B rel C) is constant and its derivative is zero. Differentiating the velocity equation once more gives a(A rel C) = a(A rel B): every such observer measures the same acceleration. Frames moving at constant velocity relative to one another agree on accelerations, and therefore, as the next chapter will show, on forces. This is why the coin on a smoothly moving train behaves as it would on the ground.

::: note
Velocity is always measured relative to something. To change observer, add the velocity of one observer relative to the other, as vectors.
:::

::: history 1632 · Florence
To argue that the Earth could move without our noticing, Galileo described a ship's cabin with flies, fish in a bowl and dripping water. Whether the ship is at anchor or sailing smoothly, everything inside behaves identically; no experiment in the cabin reveals the motion. This principle of relativity was central to accepting a moving Earth. In 1905 Einstein kept it but rejected the hidden assumption in the derivation above, that all observers share one time. At speeds near light, velocities no longer simply add.
:::

::: quiz
? You walk forward at 1.5 m/s along the aisle of a train moving at 20 m/s. What is your speed relative to the ground?
- 18.5 m/s
+ 21.5 m/s
- 20 m/s
= Both velocities point the same way, so they add: 20 + 1.5 = 21.5 m/s.
? Car A drives east at 25 m/s; car B drives west at 15 m/s. What is A's velocity relative to B?
- 10 m/s east
+ 40 m/s east
- 40 m/s west
= v(A rel B) = v(A rel ground) − v(B rel ground) = 25 − (−15) = 40 m/s east.
? A passenger on a smoothly moving train drops a ball. What path does she see it follow?
+ Straight down
- Curving backward
- Curving forward
= In the train's frame the ball starts with zero horizontal velocity, so it falls straight down. A ground observer sees a parabola.
:::
