In November 1915 Einstein finished general relativity. Its central claim is that gravity is not a force pulling through space but the curvature of spacetime itself: mass and energy bend spacetime, and freely falling objects follow the straightest possible paths through that bent geometry. Within weeks, Karl Schwarzschild found the first exact solution. Hidden in it was the black hole.

## What the field equations ask

Einstein's field equations are usually written in one line:

::: equation Einstein's field equations
G_{\mu\nu} = \frac{8\pi G}{c^4} T_{\mu\nu}
:::

Read it as: **geometry = matter and energy content**. The left side, the Einstein tensor G_μν, is built from the metric (the rule that gives distances and time intervals between nearby events) and measures how spacetime curves. The right side, the stress–energy tensor T_μν, lists the density of energy, momentum, pressure and stress at each point. The indices μ and ν each run over four directions (time and three of space), so this is really ten coupled equations, not one. The factor 8πG/c⁴ is tiny, which is why it takes a star's worth of mass to curve spacetime noticeably.

Solving them in general is hopeless. The trick is to impose enough symmetry that the problem shrinks.

## Schwarzschild's assumptions

Schwarzschild asked about the simplest possible case:

- **Spherical**: the mass and its field look the same in every direction.
- **Static and non-rotating**: nothing changes with time.
- **Vacuum outside**: we only want the field in the empty space around the mass, where T_μν = 0.

That last point matters. Outside a star the right-hand side is zero, so the equations become G_μν = 0. Empty space can still be curved; the condition only says curvature must arrange itself in a particular way.

::: steps From symmetry to the solution
$ ds^2 = {}&-A(r)\,c^2\,\mathrm{d}t^2 \\ &+ B(r)\,\mathrm{d}r^2 \\ &+ r^2\left(\mathrm{d}\theta^2 + \sin^2\theta\,\mathrm{d}\phi^2\right)
Symmetry alone fixes this form. ds² is the spacetime interval between two nearby events. A static, spherical geometry can only stretch time and radial distance by amounts depending on r; θ and φ are the usual angles on a sphere, and r is defined so a sphere at r has area 4πr².
---
$ G_{\mu\nu} = 0 \text{ gives } A(r) \times B(r) = 1
Feeding this form into the vacuum equations produces a few ordinary differential equations for A and B. One combination of them forces the product A·B to be constant, and the constant is 1 so that space is flat far away.
---
$ A(r) = 1 - \frac{k}{r}
The remaining equation is simple enough to solve exactly. k is a constant of integration: the field equations allow any value, just as Newton's laws allow a planet any mass.
---
$ A \to 1 \text{ as } r \to \infty
Far from the mass, spacetime is the flat spacetime of special relativity. The 1 in A guarantees this.
:::

The equations leave k open. Something physical must fix it, and that something is Newton.

## Matching to Newton fixes k

General relativity must reproduce Newtonian gravity where gravity is weak and speeds are low, because Newton works superbly for planets. That requirement pins k.

::: steps Weak-field limit of the time term
$ \mathrm{d}\tau = \mathrm{d}t\,\sqrt{A - \frac{v^2}{c^2}}
For a slowly moving particle, proper time τ (the time its own clock reads) follows from the metric. With dr²-type terms small, ds² ≈ −(A c² − v²) dt², and dτ² = −ds²/c².
---
$ \mathrm{d}\tau \approx \mathrm{d}t\left(1 - \tfrac{1}{2}(1 - A) - \frac{v^2}{2c^2}\right)
When 1 − A and v²/c² are both small, √(1 − x) ≈ 1 − x/2.
---
$ \text{maximize } \int \left(-\tfrac{1}{2}(1 - A)c^2 - \tfrac{1}{2}v^2\right)\mathrm{d}t
In relativity a free particle follows the path that makes its proper time largest. Multiplying by c² and dropping the constant term does not change which path wins.
---
$ \text{minimize } \int \left(\tfrac{1}{2}v^2 - \Phi_{\text{eff}}\right)\mathrm{d}t, \quad \text{with } \Phi_{\text{eff}} = -\tfrac{1}{2}(1 - A)c^2
Flipping sign turns "maximize" into "minimize". The integrand is kinetic minus potential energy per unit mass, the quantity Newtonian mechanics minimizes. So the particle moves as if in a Newtonian potential Φ_eff.
---
$ -\tfrac{1}{2}\frac{k}{r}c^2 = -\frac{GM}{r}
Newton's potential outside a mass M is Φ = −GM/r. Since 1 − A = k/r, matching the two gives this.
---
$ k = \frac{2GM}{c^2}
Solve for k. The constant is now fixed, and it is exactly Michell's radius.
:::

Call it the **Schwarzschild radius**, r_s = 2GM/c². The complete solution is:

::: equation The Schwarzschild metric
ds^2 = {}&-\left(1 - \frac{r_s}{r}\right)c^2\,\mathrm{d}t^2 \\ &+ \frac{\mathrm{d}r^2}{1 - \frac{r_s}{r}} \\ &+ r^2\left(\mathrm{d}\theta^2 + \sin^2\theta\,\mathrm{d}\phi^2\right)
:::

Term by term:

- **−(1 − r_s/r) c²dt²** governs time. The factor is less than 1 near the mass, so clocks lower down tick slower than distant clocks. This term also carries Newtonian gravity, as just shown.
- **dr² / (1 − r_s/r)** governs radial distance. The factor is greater than 1, so the actual distance between two spheres is more than the difference of their r values: space is stretched in the radial direction.
- **r²(dθ² + sin²θ dφ²)** is ordinary distance around a sphere of radius r, unchanged.
- **t** is time measured by a clock at rest very far away.

For the Sun, r_s/r at its surface is about 4 × 10⁻⁶, so the corrections to flat spacetime are tiny, yet measurable: they account for the extra precession of Mercury's orbit and for the bending of starlight measured in 1919.

## Where the math breaks

Now look at r = r_s. The time factor 1 − r_s/r becomes zero, and the radial term divides by zero and blows up. At r = 0 things are worse: r_s/r itself is infinite.

For the Sun this is harmless. The solution only applies outside the star, and r_s = 3 km lies deep inside a star 700,000 km in radius, where the vacuum assumption fails and a different solution takes over. But if a mass were ever smaller than its own Schwarzschild radius, both trouble spots would sit in empty space, in the region the formula describes.

::: note
Nobody put a black hole into these equations. Schwarzschild asked the simplest question about gravity outside a round mass, and the answer contained two places where the metric misbehaves. Whether they are real, and what they mean, took another fifty years to settle.
:::

::: history 1915–16 · Russian front
Schwarzschild, director of the Potsdam observatory, had volunteered for the German army at 40 and was serving as an artillery officer in Russia. He read Einstein's November 1915 papers there and mailed Einstein his exact solution on 22 December, joking that the war had treated him kindly enough to let him "take a walk in the land of your ideas." Einstein, who had only found approximate solutions, was surprised that an exact one existed and presented it to the Prussian Academy in January 1916. Schwarzschild fell ill with a skin disease at the front and died in May 1916.
:::

::: quiz
? Why is the Einstein tensor set to zero in Schwarzschild's problem?
- Because the star has no mass
+ Because he only solved for the empty space outside the star
- Because spacetime is flat there
= T_μν = 0 in vacuum, so G_μν = 0. Spacetime can still be curved there.
? What fixes the constant k = 2GM/c² in the solution?
- The field equations alone
+ Requiring the solution to match Newtonian gravity far from the mass
- Measurements of black holes
= The vacuum equations allow any k; matching the time term to Newton's potential −GM/r fixes it.
? At what radius does the radial term dr²/(1 − r_s/r) blow up?
- r = 0 only
- r = 2r_s
+ r = r_s
= The denominator 1 − r_s/r is zero when r = r_s.
:::
