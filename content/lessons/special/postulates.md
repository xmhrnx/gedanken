In 1905 Albert Einstein, a 26-year-old technical examiner at the Swiss patent office in Bern, published a paper that rebuilt the ideas of space and time on which all of physics rested. He did it not with a new experiment or a new force, but by taking two statements that nearly every physicist already accepted, putting them side by side, and refusing to give up either one. Each statement on its own seems harmless. Together they appear to contradict each other, and the only way to keep both is to change what we mean by a length and a duration. This lesson sets out the two statements, called the postulates of special relativity, shows exactly where they clash, and checks how the new rule for adding speeds removes the clash. The lessons that follow work out what the postulates imply.

## What is a frame of reference?

Before the postulates can be stated, we need a word for a point of view. A **frame of reference** is a setting from which positions and times are measured: a laboratory, a train carriage, a spacecraft, each with its own rulers to measure distances and its own clocks to measure times. When we say a ball is moving at 10 m/s, we always mean relative to some frame, usually the ground, though we rarely say so.

Some frames are simpler to do physics in than others. An **inertial frame** is one that is not accelerating, which means one in which an object with no forces acting on it keeps moving in a straight line at constant speed, as Newton's first law says. A train gliding along a straight, smooth track at steady speed is very nearly an inertial frame: a ball placed on the floor stays put. A train that is braking is not: the same ball rolls forward with no visible force pushing it. Special relativity is about inertial frames, and the "special" in its name refers to this restriction. Accelerating frames, and gravity, needed Einstein's later general theory.

## What does the first postulate say?

The first postulate is the **principle of relativity**:

**The laws of physics are the same in every inertial frame.**

Its practical meaning is that no experiment carried out entirely inside a closed laboratory can tell you whether the laboratory is at rest or moving at a steady velocity. You can sense acceleration, as anyone does when an aircraft takes off. But once the aircraft is cruising smoothly, a cup of coffee pours, a dropped pen falls, and a ball tossed in the aisle moves exactly as they would on the ground.

The idea is old. Galileo made the point in 1632, in a book defending the view that Earth moves around the Sun. Critics had argued that if Earth moved, we would feel it, and falling stones would be left behind. Galileo answered by asking his readers to imagine themselves shut in a cabin below the deck of a large ship, with butterflies, a bowl of fish and a dripping bottle. Whether the ship is moored or sailing smoothly, the butterflies fly as freely in all directions, the fish swim with no more effort toward one side of their bowl than the other, and the drops fall straight into a vessel beneath. Nothing inside the cabin reveals the motion. Only motion relative to something else, the shore for instance, has any meaning.

For mechanics, this principle had been accepted for more than two centuries. Einstein's step was to insist that it apply to all of physics, including electricity, magnetism and light. That sounds like a small extension, but it is not, as we will see.

## What does the second postulate say?

The second postulate concerns light:

**Light in empty space travels at the same speed, c ≈ 3.00 × 10⁸ m/s, in every inertial frame.**

Einstein actually stated it more modestly: light in empty space always travels at a definite speed c, whatever the motion of the body that emits it. On its own, that is what any wave theory says. The speed of sound does not depend on how fast the trumpet is moving; it is fixed by the air. Light, a wave in the ether as most physicists then believed, would behave the same way.

The force of the postulate comes from combining it with the first. Maxwell's equations of electromagnetism predict that light travels at one definite speed, c, fixed by two constants of electricity and magnetism that can be measured on a laboratory bench. If, as the first postulate says, Maxwell's equations hold in every inertial frame, then every inertial observer must find that same speed c, whatever their own motion. There is no ether to measure it relative to, because no experiment could single out the ether's frame from any other. The speed of light is not relative to anything: it is the same for everyone.

## Where is the contradiction?

Everyday experience says velocities add. Walk toward the front of a train at 1 m/s while the train moves at 30 m/s, and someone standing beside the track sees you go past at 31 m/s. In Newton's physics this is exact:

::: equation Galilean velocity addition
u = u' + v
:::

Here u′ is your velocity as measured on the train, v is the train's velocity relative to the ground, and u is your velocity relative to the ground.

Now replace yourself with a pulse of light, sent forward from a lamp on the train. On the train, the second postulate says it moves at u′ = c. Then Newton's rule says that, from the ground, it moves at c + v. But the second postulate also applies on the ground, and says it moves at c. Both cannot be right. Measured from two frames moving relative to each other, the same light pulse cannot have the same speed if speeds simply add.

Einstein's way out was to ask what the addition rule actually assumes. To say that you cover 31 metres of track per second, the ground observer must measure distances and times, and so must you on the train. Newton's rule silently assumes that the two of you agree about both: that a metre on the train is a metre on the ground, and that a second on the train's clocks is a second on the ground's. Nobody had ever tested that at speeds near c. If it is false, the contradiction disappears, and the postulates can both be true. That was Einstein's conclusion: the assumption that time and distance are the same for every observer is the one to give up.

## What rule replaces simple addition?

Working through what the postulates require, which the next lessons do one piece at a time, Einstein found that velocities along a single line combine by a different rule. Its derivation needs the Lorentz transformation, the relativistic recipe for converting the positions and times measured in one frame into those measured in another. Here we take the rule as given and check that it behaves as it must.

::: steps Checking the relativistic addition rule
$ u = \frac{u' + v}{1 + u'v/c^2}
This is the relativistic velocity addition formula, for an object moving at u′ inside a frame that itself moves at v, both along the same line. The top is Newton's sum. The bottom is a correction that depends on how large the two speeds are compared with c.
---
$ u \approx u' + v \quad \text{when } u', v \ll c
Check everyday speeds first. For a person walking at 1 m/s on a train moving at 30 m/s, the correction term u′v/c² is 30/(9 × 10¹⁶), about 3 × 10⁻¹⁶. The denominator differs from 1 by that tiny amount, and the formula agrees with Newton's to far better than any measurement could detect. This is why nobody noticed the correction before: it is real, but at ordinary speeds it is unimaginably small.
---
$ u = \frac{c + v}{1 + cv/c^2} = \frac{c + v}{1 + v/c}
Now check the case that caused the contradiction: light sent forward from the train, so u′ = c. Putting that into the formula, one factor of c cancels in the correction term, which becomes v/c.
---
$ u = \frac{c + v}{(c + v)/c} = c
Write the denominator 1 + v/c as a single fraction, (c + v)/c. The factor c + v now appears on the top and the bottom and cancels, leaving exactly c. The ground observer measures the light moving at c whatever the train's speed v, just as the second postulate demands.
:::

The rule has a striking consequence. Add 0.5c to 0.5c, as for a probe launched forward at half the speed of light from a ship already moving at half the speed of light, and the result is (0.5 + 0.5)c/(1 + 0.25) = 0.8c, not c. However many speeds below c you combine this way, the result stays below c. The speed of light is not merely the speed of light; it behaves as a limit built into the way motions combine.

The consequences come one after another, each derived from the two postulates with nothing else assumed. Moving clocks run slow. Moving objects are shorter along their motion. Two events that happen at the same moment for one observer happen at different moments for another. Mass is a form of energy. Each of these sounds absurd at first, and each has been measured.

::: note
Special relativity rests on two postulates: the laws of physics are the same in every inertial frame, and light's speed in empty space is c in every inertial frame. Velocities cannot simply add if both are true, so measurements of time and distance must depend on the motion of the observer.
:::

::: history 1895–1905 · Aarau and Bern
At sixteen, while at school in Aarau in Switzerland, Einstein asked himself what he would see if he could chase a beam of light at speed c. He would be riding alongside an electromagnetic wave, and should see it frozen in place, a pattern of electric and magnetic fields standing still in space. But Maxwell's equations allow no such thing: in them, changing fields keep each other going, and a frozen wave would not exist. Something was wrong either with the chase or with the idea that light's speed depends on the observer. Ten years later, his 1905 paper "On the Electrodynamics of Moving Bodies" opened with a different, quieter puzzle. Move a magnet near a stationary loop of wire and a current flows; hold the magnet still and move the wire instead, and exactly the same current flows. Yet the theory of the day explained the two cases in completely different ways, one with an electric field and one without. Einstein took this as a sign that only the relative motion matters, and that the theory was making a distinction nature did not. Lorentz and Henri Poincaré had already found much of the mathematics he would use, but they interpreted it as describing how matter behaves when it moves through the ether. Einstein made it a statement about space and time themselves, and called the ether "superfluous".
:::

## Has the second postulate been tested?

It has, directly and many times. In 1964 at CERN, the European particle physics laboratory near Geneva, a team led by Torsten Alväger produced short-lived particles called neutral pions moving at 99.975% of the speed of light. Each pion decays into two gamma rays, a form of very energetic light. If speeds simply added, gamma rays emitted forward should have travelled at nearly 2c. The team timed them between detectors tens of metres apart and found they travelled at c, to within about one part in ten thousand. The speed of the source made no difference.

::: quiz
? A spaceship moving at 0.9c relative to Earth fires a laser beam straight ahead. How fast does an observer on Earth measure the laser light to move?
- 1.9c
- 0.1c
+ c
= By the second postulate, light's speed in empty space is c in every inertial frame, whatever the motion of its source. The relativistic addition rule gives the same answer: (c + 0.9c)/(1 + 0.9) = c.
? According to the first postulate, what can a passenger in a windowless train moving smoothly at constant velocity learn about the train's speed from experiments inside?
+ Nothing at all
- The speed, by timing a pendulum
- The speed, by timing light beams
= The laws of physics are identical in every inertial frame, so every experiment inside gives the same result whether the train is moving steadily or at rest. Light beams are no exception; that was Einstein's extension of Galileo's principle.
? Two spaceships approach Earth from opposite sides, each moving at 0.5c relative to Earth. How fast does one ship measure the other to be approaching?
- c
+ 0.8c
- 0.5c
= With u = (u′ + v)/(1 + u′v/c²), the closing speed is (0.5 + 0.5)c/(1 + 0.25) = 0.8c. The correction in the denominator keeps the result below c.
:::
