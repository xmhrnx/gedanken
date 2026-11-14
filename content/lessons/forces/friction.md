Friction is why Aristotle thought motion needs a push: it is the force that makes sliding things stop. It is also why you can walk, why brakes work and why a knot holds. Drag is its counterpart in air and water. Neither is a fundamental force; both are the summed effect of countless electrical interactions between atoms, described by simple rules that work surprisingly well.

## Static and kinetic friction

Push gently on a heavy box and it stays put. The floor is pushing back with a friction force exactly equal to your push. Push harder and the friction grows to match, up to a limit; past that the box breaks free and slides. Once moving, it usually takes a bit less force to keep it going.

So there are two regimes. **Static friction** acts between surfaces that are not sliding and takes whatever value is needed to prevent sliding, up to a maximum. **Kinetic friction** acts once they slide and has a roughly fixed size. Both are found to be proportional to the normal force N pressing the surfaces together:

::: equation Friction laws
f_s \leq \mu_s N, \qquad f_k = \mu_k N
:::

The **coefficients of friction** μ_s and μ_k are pure numbers that depend on the pair of materials, with μ_k usually smaller than μ_s. Rough typical values: rubber on dry asphalt, μ_s about 0.9 and μ_k about 0.7; steel on steel, about 0.6 static; a skate blade on ice, a few hundredths.

Two features are surprising. Friction does not depend on the apparent area of contact: a brick slides just as hard on its face as on its edge. And kinetic friction depends only weakly on speed.

## Why friction behaves this way

Even polished surfaces are rough at the microscopic scale, and two of them touch only at the tips of their highest bumps. The true contact area is a tiny fraction of the apparent area, and it grows in proportion to the load as the bumps squash. At those contact points the atoms bond; friction is the force needed to shear the bonds. Proportional to load, independent of apparent area: the laws follow. This picture, worked out by Frank Bowden and David Tabor around 1950, is itself simplified, and friction remains an active research field.

## Measuring μ_s with a ramp

Put a block on a board and slowly tilt it until the block starts to slide.

::: steps The angle of slipping
$ W_{\parallel} = mg\sin\theta, \qquad W_{\perp} = mg\cos\theta
Split the weight into components along the slope and perpendicular to it, as in the vectors lesson.
---
$ N = mg\cos\theta
The block does not accelerate into or out of the surface, so the normal force balances the perpendicular component of weight.
---
$ mg\sin\theta \leq \mu_s mg\cos\theta
To stay put, static friction must balance the component along the slope, and it can be at most μ_s N.
---
$ \tan\theta_{\max} = \mu_s
Divide both sides by mg cos θ. The block slips when tan θ exceeds μ_s, whatever its mass.
:::

A block that starts sliding at 30° has μ_s = tan 30° ≈ 0.58.

## Stopping a car

A car braking hard without skidding is slowed by friction between tyres and road. If the normal force is the car's weight, the largest friction force is μmg, so the largest deceleration is μmg/m = μg. From the constant-acceleration equation v² = u² + 2as with final speed zero, the stopping distance is d = u²/(2μg). At 25 m/s (90 km/h) on dry road with μ = 0.7, that gives 625/(2 × 0.7 × 9.8) ≈ 46 m, plus whatever distance the car covers while the driver reacts. On a wet road with μ = 0.4 it is about 80 m. Distance grows as the square of speed.

## Drag

A body moving through air or water must shove the fluid aside. The resulting **drag** force opposes the motion and, unlike sliding friction, depends strongly on speed. For most everyday objects at everyday speeds:

::: equation Drag force
F_D = \tfrac{1}{2}\rho C_D A v^2
:::

Here ρ is the fluid's density, A the cross-sectional area facing the flow and C_D a *drag coefficient* set by the shape, about 0.5 for a sphere and 0.3 for a modern car. Double the speed and the drag quadruples.

This explains the feather from the free-fall lesson. A falling body speeds up until drag equals its weight; then the net force is zero and it falls at a constant **terminal velocity**. Setting mg = ½ρC_DAv² gives v_t = √(2mg/(ρC_DA)). A skydiver falling belly-down reaches about 55 m/s, close to 200 km/h. A feather, with tiny mass and large area, reaches its terminal velocity almost at once.

::: history 1493–1781 · Milan, Paris
Leonardo da Vinci found in the 1490s that friction is proportional to load and independent of contact area, but his notebooks went unpublished. Guillaume Amontons rediscovered both rules and presented them to the French Academy in 1699, where they were met with disbelief. Charles-Augustin de Coulomb, an army engineer later famous for electrostatics, tested them exhaustively and distinguished static from kinetic friction in a prize essay of 1781.
:::

::: note
Friction is proportional to the force pressing surfaces together: f = μN. Static friction adjusts to prevent sliding, up to μ_s N. Drag grows with the square of speed and sets a terminal velocity.
:::

::: quiz
? A 10 kg box sits on a floor with μ_s = 0.5. You push it horizontally with 20 N. What is the friction force?
- 49 N
+ 20 N
- 0 N
= The maximum static friction is 0.5 × 98 ≈ 49 N, but static friction only supplies what is needed. You push with 20 N, so friction is 20 N and the box stays still.
? A brick slides on its large face, then is turned onto its narrow edge. Kinetic friction
- increases
- decreases
+ stays about the same
= Friction depends on the normal force, not the apparent area. The weight is unchanged, so the friction is too.
? A car's speed doubles. If braking friction stays the same, its braking distance
- doubles
+ quadruples
- stays the same
= d = u²/(2μg): distance scales with the square of the initial speed.
:::
