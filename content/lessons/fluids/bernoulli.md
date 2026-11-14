Put your thumb over the end of a garden hose and the water speeds up. Where it speeds up, its pressure drops. That link between speed and pressure in a moving fluid is Bernoulli's principle, and it follows from nothing more than conservation of mass and energy.

## Continuity: narrow pipes mean fast flow

Consider a liquid flowing steadily through a pipe that narrows from cross-section A₁ to A₂. Liquids barely compress, so in any second the same volume must pass every point along the pipe; otherwise liquid would pile up somewhere. In one second a length v of pipe empties past a point, a volume Av.

::: equation The continuity equation
A_1 v_1 = A_2 v_2
:::

Halve the area and the speed doubles. The fluid in the narrow section moves faster, so it has more kinetic energy per kilogram. Something must have pushed it to speed up, and in a fluid that something is a pressure difference: the pressure behind the fluid must be higher than the pressure ahead of it.

## Deriving Bernoulli's equation

Follow a small parcel of fluid, volume ΔV, as it travels from a point 1 (pressure P₁, speed v₁, height h₁) to a point 2 (P₂, v₂, h₂). Assume the flow is steady, the fluid incompressible, and friction (viscosity) negligible.

::: steps Bernoulli's equation from the work–energy theorem
$ W_1 = P_1 A_1 \Delta x_1 = P_1 \Delta V
The fluid behind pushes the parcel forward with force P₁A₁ through a distance Δx₁. Area times distance is the volume swept, ΔV.
---
$ W_2 = -P_2 \Delta V
At the far end the fluid ahead pushes back on the parcel as it moves forward, so this work is negative. The swept volume is the same ΔV, by continuity.
---
$ \Delta KE = \tfrac{1}{2}\rho \Delta V (v_2^2 - v_1^2), \quad \Delta PE = \rho \Delta V g (h_2 - h_1)
The parcel's mass is ρΔV. These are its changes in kinetic and gravitational potential energy.
---
$ (P_1 - P_2)\Delta V = \tfrac{1}{2}\rho \Delta V (v_2^2 - v_1^2) + \rho \Delta V g (h_2 - h_1)
With no friction, the net work done by pressure equals the gain in mechanical energy.
---
$ P_1 + \tfrac{1}{2}\rho v_1^2 + \rho g h_1 = P_2 + \tfrac{1}{2}\rho v_2^2 + \rho g h_2
Divide by ΔV and collect each point's terms on its own side. Since points 1 and 2 were arbitrary, the sum is the same everywhere along the flow.
:::

::: equation Bernoulli's equation
P + \tfrac{1}{2}\rho v^2 + \rho g h = \text{constant along a streamline}
:::

Every term is an energy per unit volume. When the fluid is still, it reduces to the previous lesson's P + ρgh = constant. At constant height it says: faster flow, lower pressure. The constant holds along a *streamline* (the path a fluid parcel follows); in steady flow that starts out uniform, it holds throughout.

## Two uses

**Draining a tank.** Water leaves a hole a depth h below the surface of an open tank. Both the surface and the jet are at atmospheric pressure, and the surface barely moves. Bernoulli gives ρgh = ½ρv², so v = √(2gh): the same speed as an object dropped from height h. Evangelista Torricelli found this experimentally in 1643.

**The Venturi meter.** In a pipe with a narrow throat, the throat's lower pressure is easy to measure, and with continuity it gives the flow speed. Carburettors, aircraft airspeed tubes and medical flowmeters use the same idea.

::: history 1738 · St Petersburg and Basel
Daniel Bernoulli, working at the St Petersburg Academy, poked open tubes into the walls of pipes and saw that water rose less high where the flow was faster. He published the result in *Hydrodynamica* (1738), reasoning from the conservation of "living force", what we call kinetic energy. The equation in its modern form came later, from Leonhard Euler in the 1750s, who wrote down the general equations of motion for a frictionless fluid.
:::

## How wings make lift, and a myth

A common textbook story runs: air splitting at the front of a wing must meet up again at the back; the top surface is longer, so air over the top must go faster; by Bernoulli its pressure is lower, and the wing is pushed up. This is the *equal-transit-time* explanation, and it is wrong. Nothing requires the two parcels to rejoin. Wind-tunnel photographs of smoke show air over the top arriving at the trailing edge well before the air underneath. And the story cannot explain how a thin flat sail generates lift, or how a plane flies upside down.

What actually happens: a wing, through its shape and its tilt (angle of attack), turns the oncoming air downward. By Newton's third law the air pushes the wing up. Air can only follow a curved path if pressure differs across it, so the pressure above the wing falls below atmospheric and the pressure beneath rises above it. That pressure difference, integrated over the wing, is the lift. The air above does move faster, much faster than equal transit would predict, and Bernoulli's equation correctly links that speed to the low pressure. But Bernoulli is a bookkeeping rule relating speed and pressure, not by itself the reason the air speeds up. The "downwash" picture and the pressure picture are the same physics described two ways.

::: note
Bernoulli's equation is energy conservation for a moving fluid: where the flow is faster, the pressure is lower. It relates speed and pressure; it does not on its own explain why the speed changes.
:::

::: quiz
? Water flows through a horizontal pipe that narrows to half its area. In the narrow part
- speed halves and pressure rises
+ speed doubles and pressure falls
- speed doubles and pressure rises
= Continuity doubles the speed; at the same height, the extra ½ρv² must come out of P.
? What is wrong with the equal-transit-time explanation of lift?
- Bernoulli's equation does not apply to air
+ There is no reason the air above and below must reach the trailing edge together, and in fact the upper air gets there first
- Wings produce lift only by deflecting air, so pressure plays no role
= Lift comes from the wing turning air downward, with lower pressure above; the upper air moves faster than equal transit predicts.
? Water leaks from a hole 5 m below the surface of an open tank. About how fast does it come out? (g ≈ 10 m/s²)
- 5 m/s
+ 10 m/s
- 50 m/s
= v = √(2gh) = √(2 × 10 × 5) = 10 m/s.
:::
