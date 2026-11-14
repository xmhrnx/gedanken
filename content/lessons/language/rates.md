Physics is mostly about change: how an object's position changes as time passes, how temperature changes along a heated rod, how a sample of radioactive atoms dwindles as they decay. To talk about change precisely we need a way of saying how fast one quantity changes as another one does. That idea is the **rate of change**, and its sharpest form, the rate at a single instant, is called the **derivative**. This lesson builds the derivative from the ordinary idea of average speed, and shows why it took a new kind of mathematics to make sense of something as everyday as a speedometer.

## Average rate: the slope of a line

Start with something familiar. If you drive 120 km in 2 hours, your average speed is 120 ÷ 2 = 60 km/h. You divided the change in one quantity, distance, by the change in another, time. The same recipe works for any pair of quantities. If a quantity y goes from y₁ to y₂ while another quantity x goes from x₁ to x₂, the average rate at which y changes with x is:

::: equation Average rate of change
\frac{\Delta y}{\Delta x} = \frac{y_2 - y_1}{x_2 - x_1}
:::

The symbol Δ, the Greek capital letter delta, is shorthand for "the change in", so Δy is how much y changed and Δx is how much x changed.

This ratio has a simple picture. Plot y against x on a graph and mark the two points. The straight line joining them rises by Δy while it runs across by Δx, so Δy/Δx is that line's slope, the familiar "rise over run". A steep line means y changes a lot for a little change in x; a flat line means y hardly changes at all; a line sloping downward means y decreases as x increases, and the slope is negative.

## Instantaneous rate: the puzzle of "now"

An average is useful, but it hides what happened in between. Over your two-hour drive you probably sped up, slowed down and stopped at junctions, and the 60 km/h average describes none of those moments. Your speedometer does something different. It shows how fast you are going at this instant.

That raises a real puzzle. An instant has no duration. In zero time you cover zero distance, so the recipe "distance divided by time" gives 0 ÷ 0, which is not a number at all. Yet the speedometer plainly shows something, and everyone understands that a car can be doing 60 km/h at a particular moment. What does that mean?

The answer is to stop trying to measure the rate at an instant directly, and instead to see what happens to the average as the interval shrinks. Measure the average speed over the next minute, then over the next second, then over the next thousandth of a second. Over shorter and shorter intervals the car has less and less time to change its speed, so the averages settle down toward one particular value. That value, the one the averages approach as the interval shrinks toward zero, is what we mean by the speed at that instant. The key step is that we never actually divide by zero; we look at where a sequence of perfectly ordinary divisions is heading.

On the graph, this has a clear meaning. Take a point on a curve and a second point a little way along it, and draw the straight line through them. As the second point slides toward the first, the line swings and then settles into one particular line that just touches the curve at that point without cutting across it. This is the **tangent line**. Its slope is the instantaneous rate of change of y with x at that point, and that slope is the derivative, written dy/dx. The notation is meant to look like Δy/Δx with the changes made infinitely small, which is exactly how it was arrived at.

::: model slope
Drag the point along the curve and watch the tangent line. Notice that its slope falls to zero at the top of the hump, where the curve is momentarily flat, and turns negative beyond it, where the curve heads downward.
:::

## Computing a derivative

The shrinking-interval idea can be turned into a calculation. Take the simplest curved graph, y = x², and find its slope at a general point x. The method is to compute the slope between two points a distance h apart and then see what happens as h shrinks.

::: steps The derivative of x²
$ \frac{\Delta y}{\Delta x} = \frac{(x + h)^2 - x^2}{h}
We compare the curve at the point x and at a nearby point x + h. The run between them is h, and the rise is the value of y at the second point minus its value at the first, which is (x + h)² − x². Dividing rise by run gives the slope of the line through the two points, an ordinary average rate.
---
$ = \frac{x^2 + 2xh + h^2 - x^2}{h}
Next we multiply out (x + h)² as (x + h) times (x + h), which gives x² + 2xh + h². This is ordinary algebra and assumes nothing new. Its purpose is to expose terms that will cancel.
---
$ = \frac{2xh + h^2}{h}
The x² at the start and the −x² at the end cancel each other. What is left on top is 2xh + h², and every term in it contains a factor of h. That is no accident: the rise between two points must vanish when the points coincide.
---
$ = 2x + h
Because every term on top contains h, we can divide it out. This is allowed because h, though small, is not zero; we have not yet let the points come together. The result, 2x + h, is the exact slope of the line through the two points, for any separation h.
---
$ \frac{\mathrm{d}y}{\mathrm{d}x} = 2x
Now we let h shrink toward zero. The expression 2x + h gets closer and closer to 2x, and that value is the slope of the tangent, the derivative. The division by zero that seemed unavoidable never happened: we divided first, while h was still a real number, and only then asked where the answer was heading.
:::

So the slope of the curve y = x² at any point is 2x. At x = 0 the slope is zero, which matches the picture: the curve is flat at the bottom of its bowl. At x = 1 the slope is 2, and at x = 3 it is 6, so the curve gets steeper the further out you go, as anyone who has drawn a parabola will recognise. The derivative is not one number but a new function, giving the slope at every point.

The same method applied to other powers of x gives a general pattern, called the **power rule**, which covers most of the cases you will meet in this course:

::: equation The power rule
\text{if } y = x^n, \text{ then } \frac{\mathrm{d}y}{\mathrm{d}x} = n x^{n-1}
:::

Our result fits it: for n = 2 the rule gives 2x¹ = 2x. Two further facts follow from the definition. A constant has derivative zero, because a quantity that does not change has no rate of change. And a constant factor carries straight through, because scaling a graph vertically by 5 makes every slope 5 times steeper, so the derivative of 5x² is 5 × 2x = 10x.

::: note
The derivative is the slope of a graph at a single point. It is found by taking the slope between two nearby points and asking what value it approaches as the points come together, which gives a meaning to "the rate of change right now".
:::

## Why physics cares

The reason this matters so much in physics is that the most basic quantities of motion are derivatives. **Velocity** is the rate at which position changes with time, v = dx/dt. **Acceleration** is the rate at which velocity changes with time, a = dv/dt. The speedometer puzzle was really a question about derivatives, and the tangent line is its answer.

Units come along automatically. A rate is "so many units of y per unit of x", so if position x is measured in metres and time t in seconds, then dx/dt is in metres per second, and dv/dt is in metres per second per second.

The graphical picture is just as useful. On a graph of position against time, a steep stretch means the object is moving fast, because its position changes a lot in a little time. A flat stretch, or the flat top of a hump, means the object is momentarily at rest. A stretch sloping downward means the object is moving backward, toward smaller positions. Reading a motion graph is largely a matter of reading slopes.

## Who invented it, and the quarrel that followed

The ideas in this lesson were found twice, independently, and the question of who deserved the credit turned into one of the bitterest disputes in the history of science.

::: history 1665–1716 · Cambridge, Paris and London
In 1665 and 1666, with Cambridge closed by plague, the young Isaac Newton worked out what he called the "method of fluxions", a way of finding the rates at which flowing quantities change, which he needed for his work on motion and gravity. He showed it to a few colleagues in manuscript but published almost nothing of it for decades. About ten years later, in Paris, Gottfried Wilhelm Leibniz developed the same ideas by a different route, starting from the geometry of tangents, and in 1684 he published first. Leibniz's notation, dy/dx for the derivative, proved easier to use, because it lets calculations be carried out almost as if the d's were ordinary fractions, and mathematicians on the Continent ran with it. Years later Newton's supporters accused Leibniz of having copied him. In 1712 the Royal Society in London, of which Newton was president, issued a report finding in Newton's favour; Newton had written much of it himself. Today historians agree that the two men discovered calculus independently. The quarrel left English mathematicians loyal to Newton's clumsier notation, and they fell behind their continental colleagues for most of the following century.
:::

::: quiz
? What is the derivative of y = x³?
- x²
+ 3x²
- 3x³
= The power rule with n = 3 brings the 3 down in front and lowers the power by one, giving 3x². The other options either forget the factor of 3 or forget to lower the power.
? On a position–time graph, the curve is flat at the top of a hump. What is the velocity at that instant?
- At its maximum
+ Zero
- Negative
= Velocity is the slope of the position–time graph, and a flat tangent has zero slope. The object is momentarily at rest, as a ball is at the top of its flight before it starts to fall back.
? An object's position is x = 4t² metres, with t in seconds. What is its velocity at t = 3 s?
- 36 m/s
+ 24 m/s
- 12 m/s
= Differentiating gives v = dx/dt = 4 × 2t = 8t, so at t = 3 s the velocity is 24 m/s. The 36 m/s option is the position at that time, not the rate at which it is changing.
:::
