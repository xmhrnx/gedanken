Drop a hammer and a feather side by side and the hammer lands first. For about two thousand years this everyday observation was taken as proof of what Aristotle had taught: heavier things fall faster, and a body ten times heavier falls ten times as fast. The conclusion is wrong, and finding out why it is wrong was one of the turning points of physics. The feather is slowed by the air it pushes through, not by being light. Take the air away and everything near the Earth's surface falls with exactly the same steady acceleration, whatever its mass or shape. This lesson is about that acceleration, how Galileo managed to measure a motion too fast for any clock of his day, and why mass drops out of the problem.

## What counts as free fall?

**Free fall** means motion under gravity alone, with every other influence, air resistance in particular, small enough to ignore. Strictly that happens only in a vacuum. But for dense, compact objects falling short distances, such as a stone dropped from a bridge or a ball tossed in the air, the air's effect is tiny compared with gravity's, and treating the motion as free fall gives answers that are very close to right. A feather, a sheet of paper or a skydiver after the first few seconds are different: for them, air resistance is a large part of the story, and this lesson's results do not apply.

## The acceleration of falling

The acceleration of a freely falling object is given its own letter, g:

::: equation Acceleration due to gravity near Earth's surface
g \approx 9.8\,\text{m/s}^2, \quad \text{directed downward}
:::

In words, every second of falling adds about 9.8 m/s to an object's downward velocity. A stone dropped from rest is moving at about 9.8 m/s after one second, 19.6 m/s after two and 29.4 m/s after three, and the pattern continues until it hits something.

The value is not quite the same everywhere. It ranges from about 9.78 m/s² at the equator to about 9.83 m/s² at the poles. Two effects cause this. The Earth spins, and a person at the equator is carried around a large circle once a day, so a little of gravity's pull goes into keeping them on that circle rather than into making them fall. And the Earth is slightly flattened, wider at the equator than from pole to pole, so a person at the equator is a little further from the centre and feels a slightly weaker pull. For most purposes 9.8 m/s² is accurate enough.

What matters most is that g is constant during a fall. That means the four equations from the lesson on constant acceleration apply directly, with no new work needed. If we take upward as the positive direction, which is the usual choice, then the acceleration is a = −g, negative because it points down.

::: model falling
The ball's position is marked every 0.1 seconds as it falls. Notice that the gaps between marks grow steadily, because the ball covers more ground in each interval as its speed increases.
:::

## A pattern in the gaps

The marks in the model follow a simple pattern, and it is the pattern Galileo used to establish the law of fall. Start an object from rest and measure the distance d it has fallen, counting downward. With v₀ = 0 and a downward acceleration g, the position equation from the previous lesson becomes d = ½gt². Now ask how far the object falls during each separate interval of time.

::: steps Distances in successive equal intervals
$ d(1) = \tfrac{1}{2}g, \quad d(2) = \tfrac{1}{2}g \cdot 4, \quad d(3) = \tfrac{1}{2}g \cdot 9
Measure time in equal units, whatever they are, and put t = 1, 2 and 3 into d = ½gt². The total distances fallen after one, two and three units are ½g times 1, 4 and 9. These are the square numbers, which is what a distance proportional to t² must give.
---
$ d(2) - d(1) = \tfrac{1}{2}g \cdot 3, \quad d(3) - d(2) = \tfrac{1}{2}g \cdot 5
The distance covered during the second interval alone is the total after two units minus the total after one, which is ½g times 4 − 1 = 3. Likewise, the third interval alone covers ½g times 9 − 4 = 5. So the separate intervals cover amounts in the ratio 1 : 3 : 5.
---
$ n^2 - (n - 1)^2 = n^2 - (n^2 - 2n + 1) = 2n - 1
To see that the pattern continues forever, take the general interval, the nth one. The distance it covers is proportional to n² − (n − 1)². Multiplying out (n − 1)² gives n² − 2n + 1, and subtracting that from n² leaves 2n − 1. For n = 1, 2, 3, 4 and onward, 2n − 1 gives 1, 3, 5, 7 and so on: the odd numbers, in order.
:::

So in successive equal intervals of time, a body falling from rest covers distances in the ratio 1 : 3 : 5 : 7, and so on through the odd numbers. The beauty of this result for an experimenter in the early 1600s is what it does not require. You do not need to know g, and you do not need a precise unit of time, only a way of marking off equal intervals. All you have to check is a ratio of distances. That is what made the law testable with the equipment Galileo had.

::: history c. 1604 · Padua
A dropped ball falls about 20 metres in two seconds, far too fast to time with the clocks of Galileo's day. His solution was to slow the fall down. He rolled a polished bronze ball down a long, smooth groove cut into a tilted wooden beam. He argued that rolling down a slope is falling diluted: the motion is of the same kind, only with a smaller acceleration, so whatever law it obeyed, free fall would obey too. To time each run, he let water flow from a large tank through a thin pipe into a cup while the ball rolled, then weighed the water; the weight measured the time. At every slope he tried, the distances grew as the square of the time, just as the odd-number rule requires. (A rolling ball's acceleration is somewhat less than a sliding one's, because some of its motion goes into spinning, but that changes only the size of the acceleration, not the square law.) He published the work in 1638 in *Two New Sciences*, and it stands as one of the first experiments designed to test a mathematical law of nature. The famous story of Galileo dropping balls from the Leaning Tower of Pisa comes from a biography written by his student Vincenzo Viviani many years later, and historians doubt it ever happened.
:::

## Why mass does not matter

Galileo did not rely on experiment alone. He also gave an argument, using nothing but reasoning, that Aristotle's view cannot be right. Suppose heavy things do fall faster than light ones. Now tie a heavy stone to a light one with a short cord and drop them together. The light stone, falling more slowly, should hold the heavy one back, so the pair should fall more slowly than the heavy stone alone. But the pair together is heavier than the heavy stone, so by the same assumption it should fall faster. The assumption predicts both a slower and a faster fall for the same object, which is impossible. The only way out is for every body to fall at the same rate.

That argument shows that the rate of fall cannot depend on weight, but it does not explain why. The explanation came from Newton, decades later. The gravitational pull on an object is proportional to its mass: twice the mass, twice the pull. But an object's resistance to being accelerated, its inertia, is also proportional to its mass: twice the mass needs twice the force for the same acceleration. In free fall these two effects exactly cancel. A heavier object is pulled harder, but it is harder to move by exactly the same factor, so it ends up with exactly the same acceleration. Why gravitational pull and inertia should be proportional to the same quantity was a puzzle that Newton noted and that Einstein later made the starting point of general relativity.

The most memorable demonstration took place far from Earth. In 1971 the Apollo 15 astronaut David Scott, standing on the airless surface of the Moon, held up a geological hammer and a falcon feather and let go of them together. With no air to slow the feather, they fell side by side and hit the dust at the same moment.

## A worked example

A stone is dropped from the top of a 20 m cliff. How long does it take to reach the ground, and how fast is it going when it gets there? Measuring distance downward from the top, the stone starts from rest and falls a distance d = 20 m with acceleration g.

::: steps A stone dropped from a 20 m cliff
$ 20 = \tfrac{1}{2}(9.8)\,t^2
For a fall from rest, the distance fallen is d = ½gt². We know the distance and g, so this equation contains only one unknown, the time t, which is what we want first.
---
$ t^2 = \frac{40}{9.8} \approx 4.08, \quad t \approx 2.0\,\text{s}
Multiplying both sides by 2 and dividing by 9.8 isolates t², and taking the square root gives the time. We keep the positive root because the stone lands after it is released, not before. The fall takes about two seconds.
---
$ v = \sqrt{2gd} = \sqrt{2 \times 9.8 \times 20} \approx 19.8\,\text{m/s}
For the speed we could use v = gt, but the time-free equation v² = v₀² + 2ad avoids relying on a rounded time. With v₀ = 0 and a = g, it gives v² = 2gd = 392, and so v ≈ 19.8 m/s.
:::

That final speed is about 71 km/h, the speed of a car on a main road, reached in two seconds from a modest cliff. It is a reminder of how quickly free fall builds up speed, and why even short falls are dangerous.

A ball thrown straight upward follows the same law, with the motion running in both directions. On the way up, gravity slows it by 9.8 m/s every second. At the top it stops for an instant. On the way down, gravity speeds it up by 9.8 m/s every second. The two halves are mirror images, so, without air resistance, the ball takes as long to come down as it took to go up and returns to your hand at the same speed it left.

::: note
Without air resistance, every object near the Earth's surface falls with the same constant acceleration, g ≈ 9.8 m/s², whatever its mass. Heavier objects are pulled harder but are harder to accelerate by exactly the same factor, so the two effects cancel.
:::

::: quiz
? A ball falls from rest. In the first second it falls about 4.9 m. How far does it fall during the third second alone?
- 9.8 m
+ 24.5 m
- 44.1 m
= By the odd-number rule, the distances covered in successive seconds are in the ratio 1 : 3 : 5, so the third second alone covers 5 × 4.9 = 24.5 m. The 44.1 m option is the total distance fallen after three seconds.
? A ball is thrown straight up. What is its acceleration at the highest point?
- Zero
+ 9.8 m/s² downward
- 9.8 m/s² upward
= Gravity keeps acting throughout the flight. The velocity is zero for an instant at the top, but it is still changing by 9.8 m/s every second, from upward to downward.
? In a vacuum, a 10 kg ball and a 1 kg ball are dropped together. What happens?
- The 10 kg ball lands first
+ They land together
- The 1 kg ball lands first
= The heavier ball feels ten times the gravitational pull but also has ten times the inertia, so both have the same acceleration. With no air resistance, nothing else distinguishes them.
:::
