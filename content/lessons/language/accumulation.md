The previous lesson asked one kind of question: if you know how much of something there is at every moment, how fast is it changing? This lesson asks the reverse. If you know how fast something is changing at every moment, how much of it has built up in total? A car's speedometer reading, recorded over a journey, should be enough to tell you how far the car went; the flow rate through a pipe, recorded over a day, should tell you how much water passed. The tool that answers questions like these is the **integral**, and the surprise at the heart of calculus is that it turns out to be the derivative run backward.

## Distance from a speed graph

Begin with the simplest case. If you drive at a steady 60 km/h for 2 hours, you cover 60 × 2 = 120 km. Now draw that journey on a graph with speed on the vertical axis and time on the horizontal one. A steady speed is a horizontal line at height 60, running from time 0 to time 2. The region under that line is a rectangle, 60 high and 2 wide, and its area is 60 × 2 = 120, the same number as the distance.

That is no coincidence, and it is worth seeing why. The area of a rectangle is its height times its width. On this graph, height means speed and width means time, so the area is speed times time, and speed times time is distance. The units agree too: kilometres per hour times hours leaves kilometres. The argument does not depend on anything special about speed. Whenever the vertical axis of a graph is a **rate**, something per unit of whatever is on the horizontal axis, the area under the graph is the total amount that has accumulated.

## When the rate keeps changing

Real journeys are not that tidy. If the speed rises and falls, the region under the graph has a curved top, and there is no simple formula for its area. The way forward is an old one: cut the region into thin vertical strips.

The idea rests on one observation. Over a short enough stretch of time, the speed barely changes, so a thin strip of the region is almost exactly a rectangle, with a height equal to the speed at that moment and a width equal to the short time interval. The area of each strip is then easy, and it is approximately the distance covered during that short interval. Add up the areas of all the strips and you get an approximation to the total distance.

The approximation is not perfect, because each strip has a flat top where the true graph is curved, leaving a small sliver either missing or extra. But those slivers shrink faster than the strips do. Halve the width of every strip and each sliver becomes much smaller, so the total error falls. By making the strips thinner and thinner, you can make the approximation as close as you like to the true area.

::: model area
Increase the number of rectangles with the slider. Watch the gaps and overlaps along the top of the curve shrink, and the total area of the rectangles close in on the true area underneath.
:::

## The integral

The **integral** is the value that this sum of strips approaches as the strips become infinitely thin. It is written like this:

::: equation The definite integral
\int_a^b f(x)\,\mathrm{d}x
:::

Every part of the notation describes the strips. The tall ∫ sign is a stretched letter S, standing for "sum"; it was introduced by Leibniz, one of the two inventors of calculus. The f(x) is the height of a strip at position x, and dx stands for its width, made infinitely small. The numbers a and b mark where the region starts and ends. So the whole expression can be read aloud as "add up f(x) times dx for every x from a to b". Because it has definite starting and ending points, this is called a **definite integral**.

## Why the integral undoes the derivative

Adding up thousands of strips by hand is tedious, and for centuries every new area had to be found by its own clever argument. The breakthrough of calculus was the discovery that you almost never have to add strips at all. To see why, we look not at the area itself but at how fast the area grows as the region is extended.

::: steps Why the integral undoes the derivative
$ A(x) = \text{area under } f \text{ from } a \text{ to } x
We define a new function, A(x), which gives the area under the curve f from the fixed starting point a up to a movable right-hand edge at x. This turns a single area into something that varies, and anything that varies has a rate of change we can ask about.
---
$ \Delta A \approx f(x) \cdot \Delta x
Now move the right-hand edge forward by a small amount Δx. The area grows by one thin extra strip, whose height is about f(x), the height of the curve at the edge, and whose width is Δx. The approximation is good because, over a narrow strip, the curve hardly changes height.
---
$ \frac{\Delta A}{\Delta x} \approx f(x)
Dividing both sides by Δx turns this into a statement about rates: the extra area per unit of extra width is about f(x). This is allowed because Δx is small but not zero. The left side is now an average rate of change of A, exactly the kind of quantity the previous lesson studied.
---
$ \frac{\mathrm{d}A}{\mathrm{d}x} = f(x)
Finally, let Δx shrink toward zero. The strip becomes a better and better rectangle, so the approximation becomes exact, and the average rate of change of A becomes its derivative. The result says that the rate at which area accumulates equals the height of the curve at the edge. In other words, the area function A is a function whose derivative is f.
:::

That last line is the **fundamental theorem of calculus**, and it changes everything about finding areas. It says that to find the area under f, you do not need strips at all. Instead you look for any function whose derivative is f, called an **antiderivative** of f, and the area from a to b is that function's value at b minus its value at a. The subtraction is needed because the area function starts counting at a, and any antiderivative may have some extra constant added to it, which the subtraction removes.

Here is an example. Find the area under f(x) = 2x between x = 0 and x = 3. From the previous lesson we know that the derivative of x² is 2x, so x² is an antiderivative. The area is therefore 3² − 0² = 9. We can check this with geometry, since the graph of 2x is a straight line and the region under it is a triangle with base 3 and height 2 × 3 = 6, whose area is ½ × 3 × 6 = 9. The two methods agree, but only one of them needed the shape to be simple.

Running the power rule backward gives a general recipe. Since differentiating xⁿ⁺¹ gives (n + 1)xⁿ, an antiderivative of xⁿ is xⁿ⁺¹/(n + 1). This works for every power except n = −1, where the formula would divide by zero, and that case needs a different function, the logarithm, which appears later in the course.

::: note
The integral is the area under a graph, and it gives the total that accumulates from a rate. The fundamental theorem of calculus says that finding it is the same as running a derivative backward: find a function whose rate of change is the given rate, and take the difference of its values at the two ends.
:::

## What integrals do in physics

The opening example generalises at once. If you know an object's velocity at every moment, the area under its velocity–time graph is its **displacement**, the change in its position. If you know its acceleration at every moment, the area under the acceleration–time graph is the change in its velocity. Each is the reverse of a derivative from the previous lesson, which is just what the fundamental theorem predicts.

Many other quantities in this course are built the same way. Work is force accumulated over distance, electric charge is current accumulated over time, and impulse is force accumulated over time. Whenever a quantity is the running total of some rate, it is an integral, and the fundamental theorem is how it gets calculated.

## An idea two thousand years in the making

The strip method is far older than calculus, and the story of how long it waited for its partner shows why the fundamental theorem mattered.

::: history c. 250 BC · Syracuse
Archimedes, working in the Greek city of Syracuse on Sicily, wanted the area of a parabolic segment, the region cut off a parabola by a straight line. He fitted the largest possible triangle inside it, then smaller triangles in the gaps left over, then smaller ones again, and showed that each new layer of triangles added a quarter of the area of the layer before. Adding up this endless sequence, and using a careful argument to rule out any other answer, he proved that the segment's area is exactly 4/3 of the first triangle's. This "method of exhaustion", which he inherited from the earlier mathematician Eudoxus, is the strip idea in all but name. But each new shape needed a new piece of ingenuity, and Archimedes had no general way to proceed. That came only in the 1660s and 1680s, when Newton and Leibniz, building on a geometric hint from Newton's teacher Isaac Barrow, recognised that finding areas and finding slopes are inverse problems. After that, an area problem that might once have occupied a great mathematician became a routine exercise.
:::

::: quiz
? A car's velocity graph is a horizontal line at 20 m/s from t = 0 to t = 5 s. How far does it travel?
- 4 m
- 25 m
+ 100 m
= The region under the graph is a rectangle 20 m/s high and 5 s wide, so its area, which is the distance, is 20 × 5 = 100 m. The wrong options come from dividing or adding the two numbers instead of multiplying them.
? What is the area under f(x) = 3x² from x = 0 to x = 2?
- 4
+ 8
- 12
= The function x³ has derivative 3x², so it is an antiderivative. The area is its value at 2 minus its value at 0, which is 8 − 0 = 8.
? What does the area under an acceleration–time graph give?
- The distance travelled
+ The change in velocity
- The average acceleration
= Acceleration is the rate at which velocity changes, so adding it up over time gives the total change in velocity. Distance would come from the area under a velocity–time graph instead.
:::
