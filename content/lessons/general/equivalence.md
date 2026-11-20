Special relativity, finished in 1905, rewrote the rules for space and time, but it left gravity out. That was not an oversight that could be patched. Newton's law of gravity says that the pull of the Sun on the Earth depends on where the Sun is right now, so if the Sun moved, the Earth would feel the change at once, across 150 million kilometres. Special relativity says no influence travels faster than light. One of the two had to give. Einstein spent eight years looking for a theory of gravity that fitted relativity, and the whole search grew out of a fact that everyone since Galileo had known and nobody had thought worth explaining: everything falls the same way. This lesson follows that fact from a curiosity to a principle. By the end you should be able to explain why a falling room cannot be told apart from one floating in space, why that forces light to bend and clocks to disagree, and why the principle only holds in a small region, which is the clue that leads to curved spacetime.

## Why should a feather and a cannonball fall alike?

Drop a lead ball and a wooden ball of the same size from a tower, and if air resistance is small they land together. Galileo argued this around 1600, and astronauts on the Moon confirmed it in 1971 with a hammer and a falcon feather, which hit the dusty ground at the same instant. It is so familiar that it seems obvious. It is not.

In Newton's physics, mass does two separate jobs. The first is resistance to being pushed: a loaded shopping trolley is harder to get moving than an empty one, and the quantity that measures this is called **inertial mass**. It is the m in Newton's second law, F = ma. The second job is responding to gravity: the Earth pulls harder on a heavy object than on a light one, and the quantity that measures this is called **gravitational mass**. It is the m in Newton's law of gravity.

Nothing in Newton's theory says these two quantities must be equal. Electric force offers a useful comparison. The electric pull on an object depends on its charge, and charge has nothing to do with inertia: a proton and a positron have the same charge, but the proton is nearly 2,000 times harder to accelerate, so in the same electric field the positron accelerates 2,000 times faster. If gravity worked like that, with its own "gravitational charge" unrelated to inertia, different materials would fall at different rates. The derivation below shows exactly where the equality enters.

::: steps Why everything falls alike
$ m_i\,a = F
Start with Newton's second law for a falling object, writing m_i for its inertial mass. This says how much acceleration a a given force F produces: the larger the inertial mass, the smaller the acceleration. It is a statement about the object's resistance to being pushed and says nothing yet about gravity.
---
$ F = \frac{GMm_g}{r^2} = m_g\,g
Now write the force itself. Newton's law of gravity says a body of mass M pulls on the object with a force proportional to the object's gravitational mass m_g. Near the Earth's surface the distance r barely changes over a fall, so GM/r² is a constant, the familiar g ≈ 9.8 m/s², and the force is simply m_g times g.
---
$ m_i\,a = m_g\,g
If gravity is the only force acting, as it is in free fall without air, then the force in the first line is the force in the second. Setting the two expressions equal links the object's acceleration to both of its masses.
---
$ a = \frac{m_g}{m_i}\,g
Divide both sides by the inertial mass m_i to get the acceleration on its own. This is the key line: the acceleration depends on the ratio of the two masses. If that ratio were slightly larger for lead than for wood, lead would fall slightly faster.
---
$ \frac{m_g}{m_i} = 1 \Rightarrow a = g
Experiment finds that every material falls with the same acceleration, so the ratio must be the same for everything. Since it is the same, we are free to choose units that make it exactly 1, and then the acceleration is g for every object. The mass of the falling object has disappeared from its motion entirely.
:::

For Newton this was a coincidence, to be measured rather than explained. The measurements have been pushed very far. Around 1900 the Hungarian physicist Loránd Eötvös hung weights of different materials from a fine fibre and looked for any twist that would show gravity and inertia acting differently on them. He found none, to a few parts in a billion. In 2022 the MICROSCOPE satellite, which compared titanium and platinum masses falling together in orbit around the Earth for months, brought the agreement to about one part in 10¹⁵. A coincidence that precise begins to look like a law.

The clever part of Eötvös's method is worth a sentence. A weight hanging at rest on the rotating Earth feels two things: the downward pull of gravity, which acts on its gravitational mass, and the tendency to fly outward from the Earth's axis, which comes from its inertia and so acts on its inertial mass. The thread hangs along the combination. If the two masses had a different ratio in copper than in wood, a copper weight and a wooden weight on opposite ends of a balance bar would hang in very slightly different directions, and the bar would twist. It did not twist.

## What can you learn inside a closed box?

Einstein's step was to stop treating the coincidence as a coincidence and to build on it. His argument uses a closed, windowless room, which later writers turned into an elevator.

Suppose first that the room sits on the surface of the Earth. You stand on the floor and feel it pressing up on your feet. You let go of a ball, and it falls to the floor with an acceleration of 9.8 m/s². Now suppose instead that the room is in deep space, far from any star or planet, and a rocket underneath pushes it "upward" with a steady acceleration of 9.8 m/s². You still feel the floor pressing on your feet, because the floor has to push you to make you accelerate along with it. When you let go of the ball, nothing pushes it any more, so it carries on at the speed it had while the floor speeds up beneath it. From inside, it looks as though the ball falls to the floor at 9.8 m/s².

In both rooms every ball falls with the same acceleration, whatever it is made of. In the rocket this is obvious: the balls are not doing anything, the floor is coming up to meet them, and the floor does not care what they are made of. On Earth it is true only because gravitational and inertial mass are equal. So the equality of the two masses is exactly what makes the two rooms impossible to tell apart.

The argument also runs the other way. A room falling freely near the Earth, with its cable cut, behaves just like a room floating in empty space: you, the balls and the walls all fall at the same rate, so nothing presses on anything and everything floats. This is why astronauts aboard the International Space Station are weightless. They and their station are falling around the Earth together, moving sideways fast enough that they keep missing it.

::: note
The equivalence principle says that in a small enough region, the effects of gravity are exactly the same as the effects of being in an accelerating frame of reference. No experiment done inside a closed room, whether with balls, light, clocks or anything else, can tell the two apart.
:::

The phrase "or anything else" is where Einstein went further than Newton. Newton's coincidence was about falling objects. Einstein assumed the equivalence holds for all of physics: electricity, light, the ticking of atomic clocks. That is a much stronger claim, and it immediately produces predictions that Newton's theory does not make.

::: history 1907 · Bern
In 1907 Einstein was still working as a technical examiner at the Swiss patent office in Bern when he was asked to write a review article on relativity. Gravity was the obvious gap, and he was stuck on how to fit it in. Then, as he recalled years later, he had "the happiest thought of my life": a person falling freely from the roof of a house would not feel their own weight. For someone falling, gravity has, in a sense, vanished. Einstein saw that if gravity can be made to disappear by choosing to fall, it might be understood as a property of the way we measure motion rather than as a force like any other. In that same review he stated the equivalence of gravity and acceleration and drew two predictions from it straight away: clocks run at different rates at different heights, and gravity bends light. It took him another eight years, and a great deal of mathematics he did not yet know, to turn the idea into a full theory.
:::

## What happens to light and clocks in the box?

Here is the first of those predictions. Go back to the rocket in deep space and shine a beam of light horizontally across the room, from one wall to the other. Light travels in a straight line, as seen by someone floating outside. But while the light is crossing, the room is accelerating upward, so by the time the light reaches the far wall, the wall has moved up a little and is moving faster than when the light set off. To someone inside, the light enters at one height and hits the far wall slightly lower, and its path curves downward, just as a thrown ball's would.

Now apply the principle. If the rocket and the room on Earth cannot be told apart by any experiment, then light must curve downward in the room on Earth too. Gravity bends light.

The second prediction concerns clocks. Put a lamp on the floor of the accelerating rocket, flashing at a steady rate, and a detector on the ceiling. Each flash takes a short time to climb to the ceiling, and during that time the ceiling has picked up a little extra speed, away from the light. So the detector is receding from where each flash was sent, and it receives the flashes slightly further apart than they were sent: the familiar Doppler shift, toward lower frequency. Seen from the ceiling, the floor clock runs slow. By the principle, the same must be true in a room standing on the Earth: a clock lower in a gravitational field runs slower than one higher up. Nothing about the clock is broken. Time itself passes at different rates at different heights, and the lesson Gravitational Time Dilation works this out in full.

Neither prediction can be reached from Newton's law alone. Newton's gravity acts on mass, and light, as then understood, had none; and Newton's time is the same for everyone. The equivalence principle produces both results from nothing more than the picture of a rocket and the assumption that physics cannot tell it from the ground.

## Worked examples: how big are these effects?

The first example asks how far light falls while crossing a room. Take a room 3 m wide on the Earth's surface. Light crosses it in t = 3 m ÷ (3.0 × 10⁸ m/s) = 1.0 × 10⁻⁸ s. In that time anything falling from rest drops by ½gt² = ½ × 9.8 m/s² × (1.0 × 10⁻⁸ s)², which comes to about 5 × 10⁻¹⁶ m. That is less than the width of a single proton. The prediction is real but hopeless to see in a room; you need a stronger field and a much longer path, and near the Sun, where light grazes the surface and travels millions of kilometres through strong gravity, the bending becomes large enough to measure. The lesson Bending Light tells how that measurement made Einstein famous.

The second example asks how strong gravity is where the astronauts float. The International Space Station orbits about 400 km up, so its distance from the Earth's centre is about 6,371 km + 400 km = 6,771 km. Gravity weakens as the square of the distance, so it is weaker than at the surface by a factor (6,371/6,771)² ≈ 0.89. At the station, g is about 8.7 m/s², nearly 90 percent of its value on the ground. The astronauts are not beyond gravity at all. They feel nothing because they and everything around them fall with the same acceleration, so nothing presses on anything else. The station moves sideways at about 7.7 km/s, which is just fast enough that, as it falls, the curved Earth drops away beneath it at the same rate.

The third example asks how good a test MICROSCOPE was. Suppose titanium fell faster than platinum by one part in 10¹⁵ of g. In orbit the difference in acceleration would be about 8 × 10⁻¹⁵ m/s². At that rate, two masses starting side by side would drift apart by only about a ten-thousandth of a millimetre during a single orbit of about 100 minutes, a few hundred atoms' widths. The satellite's test masses were shielded inside a drag-free spacecraft, held in place by electric fields, and the force needed to keep them together was measured over many orbits. No difference was found at that level, so whatever gravity is, it treats titanium and platinum as identical to at least fifteen decimal places.

## Gravity as a force you can make disappear

The equivalence principle invites a change of view that is worth taking slowly, because the rest of general relativity rests on it.

Think of the forces you feel in a car that brakes hard. You lurch forward, as if pushed, but nothing pushes you. You are simply carrying on at the speed you had while the car slows around you. Physicists call such apparent pushes **inertial forces** (sometimes "fictitious forces"): they appear only because you describe motion from a frame that is itself accelerating. The centrifugal force on a merry-go-round is another. Inertial forces have a telling signature: they are always proportional to the mass of the object they act on, because they are really just inertia showing up in the wrong frame. Doubling the mass doubles the apparent force and leaves the acceleration unchanged.

Gravity has exactly this signature. It is proportional to mass and gives every object the same acceleration. Einstein drew the bold conclusion: perhaps gravity is an inertial force too, and the frame you stand in on the ground is the accelerating one. On this view, an apple falling from a tree is not being pushed by anything. It is the apple that is moving "naturally", and it is you, held up by the ground, who are being pushed. The ground presses upward on your feet with a force mg, and that force is what accelerates you away from the path you would follow if left alone.

A phone's accelerometer agrees. Lying on a table, it reports an upward acceleration of 9.8 m/s². Dropped onto a cushion, while falling, it reports zero. An accelerometer cannot sense gravity at all; it senses only the pushes of floors, chairs and rocket engines, which is to say the forces that stop you from falling freely. In Einstein's language, the freely falling frame is the true **inertial frame**, the one in which objects left alone keep moving in straight lines at steady speed, and the room standing on the ground is an accelerated frame.

This reframing has a price. If falling frames are the inertial ones, then there is no single inertial frame covering the whole Earth: a falling frame over London and a falling frame over Sydney accelerate in opposite directions. Each is good only nearby. Special relativity's single, universal set of inertial frames has been replaced by a patchwork of small local ones, and stitching them together is the job of geometry.

## Why must the box be small?

There is one condition in the principle that is easy to skip over: the region must be small. The reason is that real gravity comes from a body with a centre, and it points toward that centre.

Imagine a large room falling toward the Earth, with one ball released near its left wall and another near its right wall. Each ball falls toward the centre of the Earth, so their paths are not quite parallel; they slope slightly toward each other. Over a long fall, the two balls drift together. Similarly, a ball released near the ceiling is slightly farther from the Earth than one near the floor, so it is pulled slightly less, and the two drift apart vertically. In a rocket accelerating uniformly in empty space, none of this happens: all the balls keep exactly the same spacing.

These small leftover effects are called **tidal effects**, because the same stretching along one direction and squeezing along the others is what raises the ocean tides. Their size can be worked out with a few lines of geometry.

::: steps The tidal accelerations in a falling room
$ \theta \approx \frac{d}{R}
Two balls a horizontal distance d apart, at a distance R from the Earth's centre, each fall toward that centre. Their two lines of fall meet at the centre, so they make a small angle θ with each other, equal (in radians) to the separation divided by the distance. For a room on the Earth, R is about 6.4 million metres, so θ is tiny.
---
$ a_{\mathrm{side}} \approx g\,\theta = \frac{g\,d}{R}
Each ball's acceleration g points along its own line of fall. Because the two lines are tilted toward each other by θ, each acceleration has a small sideways part, and together they give a relative acceleration of about gθ, pulling the balls together. This is the squeeze.
---
$ g(r) = \frac{GM}{r^2} \Rightarrow \frac{\mathrm{d}g}{\mathrm{d}r} = -\frac{2GM}{r^3} = -\frac{2g}{r}
Now take two balls one above the other. Gravity weakens with distance, and differentiating Newton's law tells us how fast: over a small height h, g falls by about 2gh/R. The factor 2 comes from the square in the inverse-square law.
---
$ a_{\mathrm{up}} \approx \frac{2g\,h}{R}
So the lower ball is pulled harder than the upper one, by this amount, and the two drift apart. This is the stretch. Squeeze across and stretch along the line to the centre: that pattern is the tidal signature, and nothing a uniformly accelerating rocket does can produce it.
:::

Put numbers in. In a room 10 m wide falling near the Earth's surface, the sideways relative acceleration is 9.8 × 10 / 6.4 × 10⁶ ≈ 1.5 × 10⁻⁵ m/s². Over a 10-second fall, two balls at opposite walls close in by ½ × 1.5 × 10⁻⁵ × 10² ≈ 0.8 mm. That is small, but measurable with care, and it does not go away however you move the room. Make the room smaller, or the observation shorter, and it shrinks. The equivalence principle is exact only in the limit of a very small region over a short time. What remains when you look over larger distances is, in Einstein's theory, the real content of gravity: the curvature of spacetime, which is the subject of the next lesson.

::: model falling
A ball dropped and marked every 0.1 s. Every object falls with this same pattern of spacing, whatever it is made of. Notice that someone falling alongside would see the ball hang motionless beside them.
:::

## Common misconceptions

The most common mistake is to think there is no gravity in orbit. As the station example shows, gravity there is almost as strong as on the ground. Weightlessness is the experience of falling freely, not of being far from the Earth; you could feel it for a few seconds in a lift whose cable had snapped.

A second mistake is to read the principle as saying that gravity is "not real", or that it is merely acceleration. The principle says something narrower: the effects of gravity can be removed in a small region by falling. Over larger regions they cannot, and the tidal effects that remain are as real as anything in physics. They raise the tides and they can tear a star apart near a black hole.

A third is to think heavier things fall faster because, in everyday life, a stone beats a leaf. The difference comes entirely from air resistance, which matters more for light, broad objects. In a vacuum, the leaf and the stone fall together, as the hammer and feather did on the Moon.

A fourth is to think the rocket in the thought experiment moves at steady speed. A rocket coasting at constant velocity is as good as floating; you would feel nothing and the ball would hang beside you. It is steady acceleration, a speed that keeps increasing, that mimics standing on the ground.

## Where does the principle stop?

Tidal effects are the first limit, and the most important. A second is that the principle concerns local experiments. Something spread across all of space, such as the electric field of a charge, is not a local experiment, and puzzles about whether a charge resting on a table radiates like an accelerated one dissolve once the whole field is taken into account.

A third limit is the one every classical idea in this course meets: quantum mechanics. The equivalence principle has been tested with individual atoms. In atom interferometers, clouds of rubidium atoms are dropped and their wave nature is used as a ruler, and two different isotopes have been found to fall alike to about one part in 10¹². So far, quantum matter obeys the principle. Whether it holds all the way down, at the scale where gravity itself would need a quantum description, is unknown, because no one yet has a tested quantum theory of gravity.

Physicists also distinguish versions of the principle by how much they claim. The weak version says only that all bodies fall alike. Einstein's version adds that all non-gravitational physics behaves in a falling room as it would in empty space. The strong version goes further still and includes bodies held together by their own gravity, such as planets, whose gravitational binding energy also has mass. Laser ranging to reflectors left on the Moon by the Apollo astronauts has shown that the Earth and Moon, which have very different amounts of gravitational binding, fall toward the Sun at the same rate to a few parts in 10¹³. General relativity obeys all three versions; many rival theories violate the strong one, which is why this test has ruled several of them out.

## Where it shows up

The equivalence principle is used every day without fanfare. Astronauts train for weightlessness in aircraft flying parabolic arcs: for about 20 to 25 seconds at a time, the aircraft follows the path a thrown ball would take, and everything inside floats. Drop towers do the same on a smaller scale; the tower in Bremen, Germany, gives about 4.7 seconds of free fall in an evacuated tube 110 m long, or about 9 seconds if the capsule is first catapulted upward. Experiments that need weightlessness, such as studies of flames or fluids without convection, are done in them.

The inertial navigation systems in aircraft and submarines rely on the other face of the principle: because accelerometers cannot feel gravity, the navigation computer must add in a model of the Earth's field to work out the true motion. And every satellite's orbit, from the Moon to a GPS satellite, is free fall.

Within this course, the principle builds on Free Fall, F = ma and Universal Gravitation, and the tidal effects connect to Tides. It leads directly into the rest of the chapter. Curved Spacetime takes the leftover tidal effects and turns them into geometry. Gravitational Time Dilation derives the clock effect predicted here, and Bending Light follows the light-bending prediction to its measurement in 1919. Later, in the Black Holes chapter, the same principle explains why someone falling through an event horizon feels nothing special at the moment of crossing.

::: quiz
? An astronaut floats inside the International Space Station, about 400 km above the ground. Why is she weightless?
- Gravity is nearly zero at that height
+ She and the station are in free fall together
- The station's rotation cancels gravity
= Gravity at that height is still about 90 percent of its strength at the surface. She and the station fall around the Earth at the same rate, so nothing presses on anything, and she feels no weight.
? Suppose gravitational and inertial mass had different ratios for different materials. What would the equivalence principle then fail to explain?
+ Why all objects fall with the same acceleration
- Why the Moon orbits the Earth
- Why heavier objects weigh more
= A falling object's acceleration is a = (m_g/m_i)g, so all objects fall alike only if that ratio is the same for everything. That shared acceleration is what makes a room in gravity indistinguishable from an accelerating rocket.
? Inside a sealed, falling room you see two balls released side by side slowly drift toward each other. What can you conclude?
- The room is accelerating in empty space
+ The room is falling in a real gravitational field, and you are seeing tidal effects
- The equivalence principle has been violated
= A uniformly accelerating room cannot make side-by-side balls converge. They converge because each falls toward the centre of a real body, and that tidal effect is the part of gravity the equivalence principle can only hide in a small enough region.
? A phone's accelerometer reads 9.8 m/s² upward while it lies on a table, and zero while it falls. What does this show?
- The accelerometer is faulty when falling
+ The table's push is a real acceleration; free fall is the unaccelerated state
- Gravity switches off during a fall
= An accelerometer measures departures from free fall. On the table, the table pushes the phone upward away from its free-fall path; falling, nothing pushes it, which is exactly what the equivalence principle says.
? A rocket in deep space accelerates upward. A lamp on the floor flashes once per second toward a detector on the ceiling. What does the detector see?
- Flashes slightly closer together than one per second
+ Flashes slightly further apart than one per second
- Exactly one flash per second
= While each flash climbs, the ceiling gains speed away from it, so the flashes arrive slightly stretched out. By the equivalence principle, a clock low in a gravitational field must likewise appear to run slow from above.
:::
