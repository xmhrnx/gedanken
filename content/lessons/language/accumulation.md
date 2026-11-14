The derivative answers: given how much of something there is at each moment, how fast is it changing? The integral answers the reverse: given how fast something is changing at each moment, how much has built up in total?

## Distance from a speed graph

If you drive at a steady 60 km/h for 2 hours, you cover 60 × 2 = 120 km. On a graph of speed against time, that journey is a horizontal line at height 60, and 120 is the area of the rectangle under it: height 60, width 2.

That is not a coincidence. Area on such a graph is height times width, which here is speed times time, which is distance. Whenever the vertical axis is a rate (something per unit of the horizontal axis), area under the graph is the total accumulated quantity.

If the speed varies, the region under the graph is not a rectangle. The way forward is to cut it into thin vertical strips. Over a short enough time, the speed barely changes, so each strip is almost a rectangle: height equal to the speed at that moment, width equal to the short time interval. Add up the strips and you get an approximate area. Use thinner strips and the approximation improves.

::: model area
Increase the number of rectangles with the slider. Their total area approaches the true area under the curve.
:::

## The integral

The *integral* is the value the sum approaches as the strips become infinitely thin. It is written:

::: equation The definite integral
\int_a^b f(x)\,\mathrm{d}x
:::

The ∫ is a stretched S, for "sum" (Leibniz's notation). f(x) is the height of a strip, dx is its width, and a and b are where the region starts and ends. Read it as "add up f(x) times dx for every x from a to b".

## The connection to derivatives

Computing areas by adding strips is tedious. The breakthrough of calculus is that you rarely have to.

::: steps Why the integral undoes the derivative
$ A(x) = \text{area under } f \text{ from } a \text{ to } x
Define a function A that gives the accumulated area up to any point x.
---
$ \Delta A \approx f(x) \cdot \Delta x
Move the right edge forward by a small Δx. The extra area is a thin strip of height f(x) and width Δx.
---
$ \frac{\Delta A}{\Delta x} \approx f(x)
Divide both sides by Δx. The approximation gets better as the strip gets thinner.
---
$ \frac{\mathrm{d}A}{\mathrm{d}x} = f(x)
In the limit, the rate at which area accumulates equals the height of the curve. So A is a function whose derivative is f.
:::

This is the *fundamental theorem of calculus*. To find the area under f, find a function whose derivative is f (an *antiderivative*), then take its value at b minus its value at a.

Example: the area under f(x) = 2x from 0 to 3. We know from the power rule that the derivative of x² is 2x, so x² is an antiderivative. Area = 3² − 0² = 9. Check with geometry: the region is a triangle of base 3 and height 6, area ½ × 3 × 6 = 9.

Reversing the power rule gives a general recipe: an antiderivative of xⁿ is xⁿ⁺¹/(n + 1), for any n except −1.

## In physics

If you know velocity over time, the area under the velocity graph is the displacement. If you know acceleration, the area under that graph is the change in velocity. Work, charge, energy, impulse: many quantities in this course are defined as the accumulation of some rate, and all of them are integrals.

::: note
The integral is the area under a graph: the total that accumulates from a rate. Finding it amounts to running a derivative backward.
:::

::: history c. 250 BC · Syracuse
Archimedes found the area of a parabolic segment by filling it with ever smaller triangles and showing their total approached exactly 4/3 of the largest triangle. This "method of exhaustion" is the strip idea, nearly 2000 years early. What Archimedes lacked, and Newton and Leibniz supplied in the 1660s–80s, was the link to derivatives that turns each new area problem from a feat of ingenuity into a routine calculation.
:::

::: quiz
? A car's velocity graph is a horizontal line at 20 m/s from t = 0 to t = 5 s. How far does it travel?
- 4 m
- 25 m
+ 100 m
= The area under the graph is a rectangle: 20 m/s × 5 s = 100 m.
? What is the area under f(x) = 3x² from x = 0 to x = 2?
- 4
+ 8
- 12
= x³ has derivative 3x², so the area is 2³ − 0³ = 8.
? What does the area under an acceleration–time graph give?
- Distance travelled
+ Change in velocity
- Average acceleration
= Acceleration is the rate of change of velocity, so accumulating it gives the change in velocity.
:::
