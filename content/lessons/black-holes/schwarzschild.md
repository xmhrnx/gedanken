In November 1915 Einstein completed general relativity. Its central claim is that gravity is not a force pulling through space but the curvature of spacetime itself: mass and energy bend spacetime, and freely falling objects follow the straightest possible paths through the bent geometry. The theory is summed up in one compact equation, but that equation is notoriously hard to solve, and Einstein himself had only found approximate solutions. Within weeks, Karl Schwarzschild, a German astronomer serving in the army on the Russian front, found the first exact one: the curved spacetime around a single round mass. Hidden in his solution, though nobody understood it for decades, was the black hole. This lesson follows his route step by step, from the question he asked to the answer he found, and ends at the two places where that answer seems to break.

## What do the field equations ask for?

Einstein's field equations are usually written in a single line:

::: equation Einstein's field equations
G_{\mu\nu} = \frac{8\pi G}{c^4}\,T_{\mu\nu}
:::

Read it as: geometry equals matter and energy. The right side, the stress–energy tensor T_μν, is a table that lists the density of energy, the flow of momentum, the pressure and the internal stresses of whatever matter is present at each point. The left side, the Einstein tensor G_μν, measures how spacetime curves at that point. It is built out of the **metric**, the rule that tells you the real distance or time interval between two nearby events in terms of their coordinates, together with the metric's first and second derivatives. The labels μ and ν each run over the four directions of spacetime, one of time and three of space. Because the table is symmetric, swapping μ and ν gives nothing new, so there are ten independent entries, and the single line is really ten equations, all coupled together.

That is what makes the equations hard. In Newton's theory, you are given where the mass is and you compute the gravitational field from it. In Einstein's, the unknown is the metric itself, the geometry in which everything else is measured, and it appears in the equations through its derivatives, multiplied together in complicated ways. There is no general method for solving ten such equations at once. The only hope is to choose a situation so symmetric that most of the unknowns disappear before you start.

## What can symmetry decide before any calculation?

Schwarzschild asked about the simplest situation that matters for astronomy: a single, round, non-rotating mass such as a star, sitting alone in space. He made three assumptions. First, the mass and its field are spherically symmetric, looking the same in every direction from the centre. Second, nothing changes with time. Third, and most important, he solved only for the empty space outside the mass, where there is no matter at all and so T_μν = 0. In that region the field equations become G_μν = 0.

It is worth pausing on that third point, because it is easy to misread. Setting G_μν to zero does not say that spacetime is flat outside the star. Empty space can still be curved; the tidal stretching around the Earth happens in a near vacuum. The vacuum equations only say that the curvature must arrange itself in a particular, restricted way. That restriction is what will let us solve them.

Before solving anything, symmetry alone narrows down the form the metric can take. The derivation below builds it up from the flat spacetime of special relativity.

::: steps The form of a static, spherical metric
$ \mathrm{d}s^2 = -c^2\,\mathrm{d}t^2 + \mathrm{d}r^2 + r^2\left(\mathrm{d}\theta^2 + \sin^2\theta\,\mathrm{d}\phi^2\right)
Start with flat spacetime, with no mass anywhere, written in spherical coordinates: a distance r from the centre and two angles, θ and φ, that give a direction, like latitude and longitude. The last term is the ordinary distance for a small step around a sphere of radius r. The interval ds² is the quantity every observer agrees on; for an object moving between two events, −ds²/c² is the square of the time its own clock records.
---
$ \mathrm{d}s^2 = {}&-A(r)\,c^2\,\mathrm{d}t^2 \\ &+ B(r)\,\mathrm{d}r^2 \\ &+ r^2\left(\mathrm{d}\theta^2 + \sin^2\theta\,\mathrm{d}\phi^2\right)
Now allow a mass at the centre and ask what can change. Because nothing changes with time, the coefficients cannot depend on t. Because every direction is alike, they cannot depend on θ or φ either, and the angular part must keep the shape of a sphere. That leaves just two unknown functions of r: A, which stretches or shrinks time, and B, which stretches or shrinks distances measured straight outward. The angular term is left as r², because we are free to define r as the label for which a sphere at r has area 4πr², and far from the mass A and B must both approach 1, so that flat spacetime returns.
:::

The whole problem has collapsed from ten unknown functions of four coordinates to two unknown functions of one coordinate. What remains is to find out what the vacuum equations say about A and B.

## How do the vacuum equations fix A and B?

The Einstein tensor is calculated from the metric by a fixed recipe: take derivatives of A and B, combine them according to rules that do not depend on the problem, and collect the terms. For this metric the calculation is long but entirely mechanical, the kind of exercise students of relativity do once by hand. Rather than reproduce pages of algebra, the steps below quote the two results we need and show what each one forces. A prime means a derivative with respect to r, so A' = dA/dr.

::: steps Solving the vacuum equations
$ G_{\mu\nu} = 0 \;\Leftrightarrow\; R_{\mu\nu} = 0
The Einstein tensor is made from a simpler curvature measure, the Ricci tensor R_μν, and a single number formed from it. In empty space, taking the trace of G_μν = 0, a kind of sum over its diagonal, shows that this number must be zero, and then G_μν = 0 says exactly the same as R_μν = 0. So in vacuum we may work with the Ricci tensor, which is easier to compute.
---
$ \frac{B}{A}\,R_{tt} + R_{rr} = \frac{1}{r}\left(\frac{A'}{A} + \frac{B'}{B}\right)
Compute the time–time and radial–radial entries of the Ricci tensor for our metric, using ct as the time coordinate. Each one separately contains second derivatives of A, which makes it awkward. But in this particular combination, B/A times the first plus the second, the awkward terms cancel exactly, leaving only the simple expression on the right.
---
$ \frac{A'}{A} + \frac{B'}{B} = \frac{(AB)'}{AB} = 0
In vacuum both Ricci entries are zero, so the left side of the previous line is zero, and therefore so is the right side. Multiplying through by r leaves this sum. By the product rule, (AB)' = A'B + AB', so the sum is the derivative of the product AB divided by AB itself. If that derivative is zero, the product AB does not change with r.
---
$ A(r)\,B(r) = 1
The product AB is therefore a constant, the same at every radius. Far from the mass both A and B approach 1, so the constant must be 1. This is the first result: whatever stretching the mass does to time, it does the inverse to radial distance, and B = 1/A everywhere.
---
$ R_{\theta\theta} = 1 - \frac{\mathrm{d}}{\mathrm{d}r}\bigl(rA\bigr) = 0
Now compute one more entry of the Ricci tensor, the angular one, and substitute B = 1/A into it. It simplifies to this short expression, and the vacuum equations say it too must vanish. This is a single, simple equation for the single remaining unknown, A.
---
$ rA = r - k
The equation says that the derivative of rA with respect to r is exactly 1. The only functions whose derivative is 1 are r plus a constant, so rA must equal r minus some constant, which we call k. The minus sign is just a convenient choice of how to name the constant.
---
$ A(r) = 1 - \frac{k}{r}, \qquad B(r) = \frac{1}{1 - k/r}
Divide by r to get A, and use AB = 1 to get B. As r becomes large, k/r shrinks to zero and both functions approach 1, so spacetime becomes flat far away, as required. The vacuum equations are now completely solved, except that k can be any number at all: the equations allow it, just as Newton's laws allow a planet any mass.
:::

Pause on what has happened. The field equations, applied to empty space around a round body, have determined the geometry completely except for one number. That makes physical sense. Outside a round body, the only feature of the body that could matter is how much of it there is. So k must be related to the mass, and something outside the vacuum equations must tell us how.

There is a bonus hidden here, found by George Birkhoff in 1923. The same calculation goes through even if one does not assume that nothing changes with time: the only spherically symmetric vacuum solution is this one. So the field outside a round star is the same whether the star is quiet, pulsing in and out, or collapsing, as long as it stays round. That will matter when we come to collapsing stars.

## What fixes the constant k?

The answer is Newton. Newton's theory of gravity describes the motion of planets, moons and comets superbly, and wherever gravity is weak and speeds are low, general relativity must agree with it. Far from a star, k/r is small, and that is exactly a weak-field situation. So we can work out how a slowly moving particle behaves in this metric, compare it with how it would behave in Newton's theory, and read off k.

The link between the metric and motion is a principle mentioned in the previous chapter: a freely falling particle travels between two events along the path that makes the time on its own clock, its **proper time** τ, as large as possible.

::: steps Weak-field limit of the time term
$ \mathrm{d}\tau^2 = -\frac{\mathrm{d}s^2}{c^2} \approx \left(A - \frac{v^2}{c^2}\right)\mathrm{d}t^2
Take a particle moving slowly, at speed v, far from the mass. In a short time dt it moves a distance v dt, and the space terms of the metric contribute about v²dt². The radial coefficient B differs from 1 only by about k/r, so its correction to this is a product of two small quantities, which we drop. Dividing −ds² by c² then gives the square of the proper time.
---
$ \mathrm{d}\tau = \mathrm{d}t\,\sqrt{1 - (1 - A) - \frac{v^2}{c^2}}
Take the square root, and write A as 1 − (1 − A). This makes it plain that the quantity under the root is 1 minus two small numbers: 1 − A, which equals k/r, and v²/c², which is tiny for a planet or a thrown ball.
---
$ \mathrm{d}\tau \approx \mathrm{d}t\left(1 - \tfrac{1}{2}(1 - A) - \frac{v^2}{2c^2}\right)
For any small number x, √(1 − x) is very close to 1 − x/2; the neglected terms are of order x², far smaller still. Applying this with x equal to the sum of the two small numbers turns the square root into a simple sum, which is much easier to compare with Newton.
---
$ \tau \approx \int \left(1 - \frac{(1 - A)c^2 + v^2}{2c^2}\right)\mathrm{d}t
The total proper time along a path is the sum of these small pieces, an integral over the journey. The particle's actual path is the one that makes this total as large as possible, compared with all nearby paths between the same two events.
---
$ \text{maximize } \int \left(-\tfrac{1}{2}(1 - A)c^2 - \tfrac{1}{2}v^2\right)\mathrm{d}t
Multiply the integrand by c², which changes the size of every path's total by the same factor and so does not change which path wins. Then drop the constant c², which adds the same amount to every path between the same two events, for the same reason. What is left is the part that actually distinguishes one path from another.
---
$ \text{extremize } \int \left(\tfrac{1}{2}v^2 - \Phi_{\text{eff}}\right)\mathrm{d}t, \quad \Phi_{\text{eff}} = -\tfrac{1}{2}(1 - A)c^2
Flip the sign: making a quantity as large as possible is the same as making its negative as small as possible. The integrand is now ½v², the kinetic energy per kilogram, minus a function of position. In Newtonian mechanics, a particle in a potential Φ follows exactly the path that makes the integral of kinetic minus potential energy stationary; this is the principle of least action. So the particle moves as if it were in a Newtonian potential equal to Φ_eff.
---
$ -\tfrac{1}{2}\,\frac{k}{r}\,c^2 = -\frac{GM}{r}
Now compare. Since 1 − A = k/r, the effective potential is −½(k/r)c². Newton's potential outside a mass M is −GM/r. For general relativity to reproduce Newton's gravity, these must be the same function of r. Both fall off as 1/r, which is already a strong sign the solution is right; only the constants must be matched.
---
$ k = \frac{2GM}{c^2}
Cancel the 1/r and solve for k. The constant of integration is now fixed by the mass of the star, and it is exactly the critical radius that Michell found in 1783 by treating light as a cannonball.
:::

This constant is now called the **Schwarzschild radius**, written r_s = 2GM/c². Putting it into A and B gives the complete solution:

::: equation The Schwarzschild metric
\mathrm{d}s^2 = {}&-\left(1 - \frac{r_s}{r}\right)c^2\,\mathrm{d}t^2 \\ &+ \frac{\mathrm{d}r^2}{1 - r_s/r} \\ &+ r^2\left(\mathrm{d}\theta^2 + \sin^2\theta\,\mathrm{d}\phi^2\right)
:::

## What does each part of the solution say?

The first term governs time. Its factor, 1 − r_s/r, is less than 1 near the mass, so a clock held at rest close in ticks off less proper time than a clock far away during the same interval of t: clocks deeper down run slow. As the calculation just showed, this same term is what produces ordinary Newtonian gravity for slow objects.

The second term governs distance measured straight outward. Its factor, 1/(1 − r_s/r), is greater than 1, so the real distance between two nearby spheres, measured with a ruler, is greater than the difference between their r values. Space is stretched in the radial direction. This is the curvature of space that doubles the bending of starlight.

The third term is the ordinary distance around a sphere and is unchanged; that was built in by the way r was defined. And the coordinate t is the time kept by a clock at rest very far from the mass, where the factors are all 1. Every other observer's time has to be worked out from it using the metric.

For the Sun, r_s is about 3 km and its radius about 700,000 km, so r_s/r at its surface is about 4 × 10⁻⁶. The corrections to flat spacetime are a few parts in a million. Small as they are, they are measurable: they account for the extra swing of Mercury's orbit and for the bending of starlight measured in 1919.

## Where does the solution misbehave?

Now look at what happens at small r. At r = r_s, the factor 1 − r_s/r becomes zero. The time term vanishes, and the radial term, which divides by that factor, becomes infinite. At r = 0 things are worse still: r_s/r itself is infinite, and every term involving it breaks down.

For the Sun, neither problem is real. The solution was derived for the empty space outside the star. Inside the star there is matter, T_μν is not zero, and a different solution applies. The Sun's Schwarzschild radius of 3 km lies deep inside a star 700,000 km in radius, so the vacuum solution never reaches it. The same is true of every planet and every ordinary star.

But suppose a mass were ever smaller than its own Schwarzschild radius. Then both trouble spots, r = r_s and r = 0, would lie in empty space, in exactly the region the formula is supposed to describe. Either the solution would have to be telling us something real about those places, or it would be breaking down. Schwarzschild himself thought this case was physically impossible, and so, for many years, did nearly everyone else.

::: note
Nobody put a black hole into these equations. Schwarzschild asked the simplest question about gravity outside a round mass, the field equations fixed the answer up to one constant, and matching Newton fixed that constant as 2GM/c². The answer contained two places where the metric misbehaves, and working out whether they are real, and what they mean, took another half century.
:::

::: history 1915–16 · Russian front
Karl Schwarzschild was one of Germany's leading astronomers, director of the Astrophysical Observatory at Potsdam, when the First World War began. He volunteered for the army at the age of 40 and by late 1915 was serving as an artillery officer on the Russian front, calculating the paths of shells. There he read Einstein's November 1915 papers. Einstein had found the equations and had used approximate solutions to explain Mercury's orbit, but nobody expected an exact solution to exist. Schwarzschild found one within weeks and mailed it to Einstein on 22 December, adding that the war had treated him kindly enough, despite the heavy gunfire, to let him take "a walk in the land of your ideas". Einstein replied that he had not expected anyone could find the exact solution so simply, and presented the paper to the Prussian Academy of Sciences in January 1916. By then Schwarzschild had fallen ill with pemphigus, a painful autoimmune disease of the skin. He was sent home and died in May 1916, at 42.
:::

::: quiz
? Why is the Einstein tensor set to zero in Schwarzschild's problem?
- Because the star has no mass
+ Because he solved only for the empty space outside the star
- Because spacetime is flat there
= In vacuum there is no matter, so T_μν = 0 and therefore G_μν = 0. Spacetime can still be curved there; the equations only restrict how.
? What fixes the constant k = 2GM/c² in the solution?
- The vacuum field equations alone
+ Requiring the solution to match Newtonian gravity far from the mass
- Measurements of black holes
= The vacuum equations allow any value of k. Matching the time term to Newton's potential −GM/r, where gravity is weak, is what fixes it.
? At what radius does the radial term dr²/(1 − r_s/r) become infinite?
- Only at r = 0
- At r = 2r_s
+ At r = r_s
= The denominator 1 − r_s/r is zero when r = r_s, so the term divides by zero there. Whether that is a real feature of spacetime is the subject of the next lesson.
:::
