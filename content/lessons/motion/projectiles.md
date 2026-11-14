A *projectile* is anything thrown, kicked or fired that then moves under gravity alone: a football, a cannonball, water from a hose. Its path curves, which looks like it needs new mathematics. It does not. The key is to treat horizontal and vertical motion separately.

## Two independent motions

Ignoring air resistance, the only force on a projectile is gravity, which pulls straight down. So:

- **Horizontally** there is no acceleration. The horizontal velocity stays constant for the whole flight.
- **Vertically** the acceleration is g downward, exactly as for a dropped object.

The two motions do not interfere. A bullet fired horizontally and a bullet dropped from the same height at the same instant hit the ground together (on flat ground, without air). The fired one simply covers a lot of horizontal distance while it falls.

## Setting up

Launch from the origin with speed v₀ at angle θ above the horizontal. Split the launch velocity into components, as in the vectors lesson:

::: equation Launch velocity components
u &= v_0\cos\theta \quad \text{(horizontal)} \\ w &= v_0\sin\theta \quad \text{(vertical, upward)}
:::

Take x horizontal and y vertical, upward positive. Applying the constant-acceleration equations to each direction:

::: equation Position at time t
x &= ut \\ y &= wt - \tfrac{1}{2}gt^2
:::

::: model projectile
Change the launch angle and speed. Note which angle gives the longest range, and that pairs of angles adding to 90° land at the same spot.
:::

## The shape of the path

Eliminating t shows what curve this is.

::: steps The trajectory is a parabola
$ t = \frac{x}{u}
Solve the horizontal equation for t. Horizontal position is a direct measure of time elapsed.
---
$ y = w\left(\frac{x}{u}\right) - \tfrac{1}{2}g\left(\frac{x}{u}\right)^2
Substitute into the vertical equation.
---
$ y = x\tan\theta - \frac{g}{2v_0^2\cos^2\theta} \cdot x^2
Use w/u = sin θ / cos θ = tan θ, and u² = v₀² cos² θ. This is y = (constant)·x − (constant)·x²: a downward-opening parabola.
:::

## Time of flight, range and height

On level ground, the projectile lands when y returns to zero.

::: steps Range on level ground
$ wT - \tfrac{1}{2}gT^2 = 0
Set y = 0 at the landing time T.
---
$ T = \frac{2w}{g} = \frac{2v_0\sin\theta}{g}
Factor out T: T(w − ½gT) = 0. One solution is T = 0, the launch; the other is the landing.
---
$ R = uT = \frac{2v_0^2\sin\theta\cos\theta}{g}
The range R is the horizontal distance covered during the flight, at constant horizontal speed u.
---
$ R = \frac{v_0^2\sin(2\theta)}{g}
Use the identity 2 sin θ cos θ = sin(2θ).
:::

sin(2θ) is largest, equal to 1, when 2θ = 90°. So on level ground and without air resistance, **45° gives the maximum range**, R = v₀²/g. Also, since sin(2θ) = sin(180° − 2θ), angles θ and 90° − θ give the same range: 30° and 60° land at the same place, the 60° shot taking a higher, slower path.

The maximum height comes from the vertical motion alone. At the top, the vertical velocity is zero, so from v² = v₀² + 2aΔy with a = −g: 0 = w² − 2gH, giving H = w²/(2g).

## What air changes

Real projectiles lose speed to drag, a force opposing motion that grows rapidly with speed. Paths become lopsided, steeper coming down than going up, and the best angle for range drops below 45°. For a golf ball or a shot put, the simple parabola is a decent first approximation; for an artillery shell travelling at hundreds of metres per second, it is badly off.

::: note
Horizontal and vertical motion are independent. Horizontally the velocity is constant; vertically it is free fall. Solve each separately, linked only by time.
:::

::: history 1638 · Leiden
Medieval theory held that a cannonball flew straight until its "impetus" ran out, then dropped. In *Two New Sciences*, Galileo combined his uniform horizontal motion with his law of fall and showed the path must be a parabola. He also proved that 45° gives the greatest range and that angles equally above and below 45° give equal ranges, a fact gunners had noticed without explanation.
:::

::: quiz
? At the top of its path, what is a projectile's velocity?
- Zero
+ Horizontal, equal to v₀ cos θ
- Vertical, equal to v₀ sin θ
= The vertical component is zero at the top, but the horizontal component never changes.
? A ball launched at 20° lands 30 m away on level ground. At what other angle, with the same speed, would it land at the same spot?
- 40°
- 45°
+ 70°
= Angles θ and 90° − θ give the same range.
? Doubling the launch speed (same angle) changes the range by what factor?
- 2
+ 4
- √2
= R = v₀² sin(2θ)/g, so range goes as the square of the speed.
:::
