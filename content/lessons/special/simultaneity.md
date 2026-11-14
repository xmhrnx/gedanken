Two events that happen at the same time for one observer can happen at different times for another observer moving relative to the first. Of all the consequences of Einstein's postulates, this is the hardest to accept, because the idea of a single "now" stretching across the whole universe seems built into the way we think. It is also the key that unlocks the others. Time dilation and length contraction both seem to lead to contradictions, since each of two observers finds the other's clocks slow and the other's rulers short. Once you see that they also disagree about which distant events happen together, the contradictions disappear. This lesson starts with what "at the same time" means, follows Einstein's train through a thunderstorm, works out how large the disagreement is, and shows why it never lets an effect come before its cause.

## What does "at the same time" mean?

For two events at the same place, there is no difficulty. Two balls strike the same spot on the floor together or they do not, and every observer, however they move, agrees, because a single collision either happens or it does not.

For events far apart, the question is harder than it looks. Suppose lightning strikes a tree in one town and a tower in another. To say the strikes were simultaneous, you need a clock at each place, and the two clocks must be synchronised, set to agree with each other. But how do you synchronise clocks that are far apart? You could carry one to the other and back, but carrying a clock is a journey, and the previous lesson showed that journeys change clock readings. You could send a signal from one to the other, but then you have to allow for the time the signal spends travelling, and to know that you need to know its speed, which you can only measure with synchronised clocks. The argument goes in a circle.

Einstein's way out, in 1905, was to treat the synchronisation of distant clocks as a definition, and to base it on the one speed that the second postulate fixes for everyone. Send a flash of light from a point exactly midway between two clocks. Light travels at c in both directions, so it covers the two equal distances in equal times, and arrives at both clocks together. Set each clock to zero when the flash arrives, and they are synchronised. The same idea works in reverse for judging events: two distant events are simultaneous if light from them reaches the point midway between them at the same moment.

This recipe is entirely reasonable, and every observer can use it. The surprise is that observers in relative motion, each applying it correctly, get different answers.

## What happens on Einstein's train?

A long train travels along a straight embankment at a steady speed v. Lightning strikes twice, once at the front of the train and once at the back, leaving scorch marks both on the train and on the embankment. An observer stands on the embankment exactly halfway between the two marks there. The flashes of light from the two strikes reach her at the same moment. Since she is at the midpoint and both flashes travelled at c, she concludes that the two strikes happened at the same time.

A passenger sits exactly halfway along the train. At the moment of the strikes, by the embankment observer's account, she is passing the embankment observer. But while the light is in flight, the train carries the passenger forward, toward the light coming from the front and away from the light coming from the back. So the flash from the front reaches her before the flash from the back.

That much is not in dispute. The arrival of two flashes at one person, in a definite order, is a fact about one place, and every observer agrees on it. The disagreement comes in interpreting it. The passenger is entitled, by the first postulate, to regard herself as at rest, with the embankment sliding backward beneath the train. In her frame, she sat at the midpoint of the train, between the two scorch marks on it. Both flashes travelled at c in her frame, by the second postulate, and covered equal distances. Yet the front one arrived first. So, by exactly the same reasoning the embankment observer used, the front strike happened first.

Neither observer has made a mistake. Each applied the same sensible rule in her own frame and reached a different conclusion. There is no experiment that could show that one of them is right and the other wrong, because the first postulate says neither frame is special. Simultaneity, for events at different places, is relative: it depends on the frame of the person judging it.

## How large is the disagreement?

To get a number, take events that are simultaneous in the train's frame, and ask how far apart in time they are from the ground. Inside a carriage whose proper length, its length measured on the train, is L₀, a lamp at the exact centre sends out a flash in both directions. In the train's frame the light travels equal distances at equal speeds, so it reaches the front and back walls at the same moment. Now work in the ground frame, where the carriage moves forward at speed v.

::: steps Timing the two arrivals from the ground
$ L = \frac{L_0}{\gamma}
From the ground, the moving carriage is contracted along its motion, so its length is L₀/γ. At the moment the flash leaves the lamp, each wall is a distance L/2 from it.
---
$ t_{\text{back}} = \frac{L/2}{c + v}
The light heading backward travels at c, by the second postulate, while the rear wall moves forward to meet it at v. The gap between them, initially L/2, closes at c + v, so this is how long the light takes to reach the back wall, measured from the ground.
---
$ t_{\text{front}} = \frac{L/2}{c - v}
The light heading forward also travels at c, but the front wall is running away from it at v. The light gains on the wall at only c − v, so it takes longer to close the same initial gap of L/2.
---
$ t_{\text{front}} - t_{\text{back}} = \frac{L}{2}\cdot\frac{(c + v) - (c - v)}{c^2 - v^2} = \frac{Lv}{c^2 - v^2}
Subtract the two times to find how much later the front wall is reached. Putting both fractions over the common denominator (c − v)(c + v) = c² − v², the numerator becomes (c + v) − (c − v) = 2v, and the 2 cancels with the one in L/2.
---
$ \frac{Lv}{c^2(1 - v^2/c^2)} = \frac{L\gamma^2 v}{c^2}
Take c² out of the denominator, leaving 1 − v²/c². That expression is exactly 1/γ², from the definition of γ, so dividing by it is the same as multiplying by γ². Writing the answer in terms of γ makes it easy to bring in the contracted length.
---
$ \frac{L\gamma^2 v}{c^2} = \frac{\gamma L_0 v}{c^2}
Finally, replace L with L₀/γ from the first step. One factor of γ cancels. The result is written entirely in terms of the carriage's proper length, which is the natural quantity to quote, since it is the separation of the two events in the frame where they are simultaneous.
:::

::: equation Events simultaneous on the train and L₀ apart, as timed from the ground
\Delta t_{\text{ground}} = \frac{\gamma v L_0}{c^2}
:::

From the ground, the event at the rear of the train happens first, by this amount. The disagreement grows with the distance between the events and with the speed, and the c² in the denominator makes it tiny at everyday speeds. For a 100 m carriage moving at 100 m/s, about 360 km/h, it is about 10⁻¹³ seconds, a tenth of a trillionth of a second. That is why ordinary life gives no hint of it, and why the idea of a universal "now" works perfectly well for everything people do. For a carriage moving at 0.6c, where γ = 1.25, it becomes 0.25 microseconds for every 100 m of separation, during which light itself travels 75 m.

## Can the order of events be reversed?

If two events are far enough apart, the disagreement can be more than a question of how much time separates them. One observer may say event A happened first, another that the events were simultaneous, and a third that B happened first. That seems to threaten something we cannot do without: if one observer could see a window break before the ball was thrown, physics would make no sense.

It does not happen, and the reason is the speed of light. The order of two events can differ between observers only when they are so far apart, compared with the time between them, that not even light could travel from one to the other in that time. For such a pair, neither event can have any influence on the other, because no signal or object could get from one to the other in time. When one event can cause another, a signal travelling at c or slower can connect them, and then the formula never reverses their order: every observer, in every frame, agrees which came first. The relativity of simultaneity leaves cause and effect intact, but only on the condition that nothing, neither matter nor information, travels faster than light. That condition is one of the deepest reasons physicists believe faster-than-light signalling is impossible.

## How does this resolve the paradox of lengths?

The previous lesson left a puzzle: an observer on the ground measures a moving rod as shorter than its proper length, while the rod's owner measures the ground's rulers as shorter. How can both be right? The relativity of simultaneity answers it.

To measure the moving rod, the ground observer marks the positions of its front and back ends on the ground at the same moment, by ground clocks, and measures the distance between the marks. But the rod's frame does not agree that those two markings were simultaneous. In the rod's frame, the ground is moving backward, and by the result above the mark at the rod's front end was made first and the mark at its back end later. In between, the ground, carrying the first mark with it, slid backward along the rod. So the two marks end up closer together than the rod is long. Each observer, applying the same procedure in her own frame, gets a contracted length for the other, and each can explain, in her own terms, why the other's measurement came out short.

::: note
Whether two events at different places happen at the same time depends on the motion of the observer judging them. Events simultaneous in one frame and a distance L₀ apart are separated by γvL₀/c² in a frame moving at speed v along the line joining them, and the order of events can change only when no signal could pass between them.
:::

::: history 1898–1905 · Paris and Bern
Einstein spent years trying to reconcile Maxwell's equations, with their single speed for light, with the principle of relativity, and kept failing. In May 1905, after a long discussion of the problem with his friend Michele Besso, an engineer who also worked at the patent office, he went home still stuck. The next day he came back and said, in effect, thank you, I have solved it: time cannot be defined absolutely, and there is an inseparable link between time and the speed of signals. About five weeks later he submitted his paper, and in it Besso is the only person thanked. Einstein was not the first to see that distant simultaneity needs a definition. Henri Poincaré had argued in 1898 that simultaneity at a distance is a matter of convention, and in 1900 he described synchronising clocks with light signals. But Poincaré kept the ether as the frame in which the true time was defined, and treated the times shown by moving clocks as only apparent. Einstein abolished the privileged frame, so that every inertial observer's time was as true as any other's. The train and the lightning came later, in a popular book on relativity that Einstein wrote in 1916.
:::

::: quiz
? In the train example, the passenger sees the flash from the front first. Why does she conclude that the front strike happened first?
+ In her frame she sat at the midpoint and both flashes travelled at c, so the one that arrived first must have left first
- The front flash was brighter than the rear one
- Light from the front travelled faster because the train was moving toward it
= She is entitled to regard herself as at rest. Equal distances covered at the same speed c mean that the flash which arrived earlier was emitted earlier.
? Two events are simultaneous for an observer on the ground and are 300 km apart. Will an observer flying past at high speed also find them simultaneous?
- Yes, simultaneity is the same for everyone
+ Not if the observer moves along the line joining them; then one event comes first
- Only if the observer moves at less than half the speed of light
= The time difference γvL/c² is not zero for any motion along the line joining the events, so a moving observer finds that one of them happens first.
? Can a fast-moving observer see an effect happen before its cause?
- Yes, if they move fast enough
+ No; when one event can influence another, every observer agrees on their order
- Only if they move faster than light, which relativity allows
= The order of two events can differ between frames only when they are too far apart for any signal at speed c or less to link them. Events that can be cause and effect always keep their order.
:::
