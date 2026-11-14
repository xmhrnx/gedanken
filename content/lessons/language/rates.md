Physics is mostly about how things change: how position changes with time, how temperature changes along a rod, how a population of atoms shrinks as they decay. The tool for describing change is the *rate of change*, and its precise version is the *derivative*.

## Average rate: slope of a line

If you drive 120 km in 2 hours, your average speed is 120 ÷ 2 = 60 km/h. In general, if a quantity y goes from y₁ to y₂ while x goes from x₁ to x₂, the average rate of change is:

::: equation Average rate of change
\frac{\Delta y}{\Delta x} = \frac{y_2 - y_1}{x_2 - x_1}
:::

Δ (Greek capital delta) means "change in". On a graph of y against x, Δy/Δx is the slope of the straight line joining the two points: rise over run.

## Instantaneous rate: slope of a curve

Your speedometer does not show an average over two hours. It shows how fast you are going *now*. But "now" is a single instant, in which you travel zero distance in zero time, and 0/0 is not a number. How can a rate exist at an instant?

The answer is to take the average over a shorter and shorter interval and see what value it approaches. On the graph, the line through two nearby points tilts as the second point slides toward the first, and settles on one line that just touches the curve at that point: the *tangent line*. The slope of the tangent is the instantaneous rate of change. This value is the derivative, written dy/dx.

::: model slope
Drag the point along the curve. The tangent line shows the slope at that point; watch it go to zero at the top and turn negative beyond it.
:::

## Computing a derivative

Take y = x². Here is the derivative worked out from the definition.

::: steps The derivative of x²
$ \frac{\Delta y}{\Delta x} = \frac{(x + h)^2 - x^2}{h}
Compare the curve at x and at a nearby point x + h. The run is h; the rise is the difference in y.
---
$ = \frac{x^2 + 2xh + h^2 - x^2}{h}
Expand (x + h)². Nothing has been assumed yet; this is ordinary algebra.
---
$ = \frac{2xh + h^2}{h} = 2x + h
The x² terms cancel, and every remaining term has a factor of h, so we can divide it out. This is allowed because h is small but not zero.
---
$ \frac{\mathrm{d}y}{\mathrm{d}x} = 2x
Now let h shrink toward zero. The expression 2x + h approaches 2x. That limit is the derivative.
:::

So the slope of y = x² is 2x: zero at x = 0, 2 at x = 1, 6 at x = 3. The same method gives a general rule that covers most cases you will meet:

::: equation The power rule
\text{if } y = x^n, \text{ then } \frac{\mathrm{d}y}{\mathrm{d}x} = n x^{n-1}
:::

A constant has derivative zero (it does not change), and constant factors carry through: the derivative of 5x² is 10x.

## Why physics cares

Velocity is the derivative of position with respect to time: v = dx/dt. Acceleration is the derivative of velocity: a = dv/dt. The units follow from the definition: a rate is "units of y per unit of x", so if x is in metres and t in seconds, dx/dt is in m/s.

Graphically: where a position–time graph is steep, the object is moving fast; where it is flat, the object is momentarily at rest; where it slopes down, the object is moving backward.

::: note
The derivative is the slope of a graph at a single point, found by taking the slope between two points and letting them come together.
:::

::: history 1660s–1680s · Cambridge and Paris
Isaac Newton developed his "method of fluxions" in 1665–66, partly to handle motion and gravity, but published little of it for decades. Gottfried Wilhelm Leibniz developed the same ideas independently around 1675 and published first, in 1684. The dy/dx notation is Leibniz's, and it won because it makes calculations look like fractions that behave sensibly. The dispute over who deserved credit soured relations between English and continental mathematicians for a century.
:::

::: quiz
? What is the derivative of y = x³?
- x²
+ 3x²
- 3x³
= Power rule with n = 3: dy/dx = 3x³⁻¹ = 3x².
? On a position–time graph, the curve is flat at the top of a hump. What is the velocity there?
- Maximum
+ Zero
- Negative
= A flat tangent has zero slope, so the velocity at that instant is zero.
? Position is x = 4t² metres, with t in seconds. What is the velocity at t = 3 s?
- 36 m/s
+ 24 m/s
- 12 m/s
= v = dx/dt = 8t, so at t = 3 s, v = 24 m/s.
:::
