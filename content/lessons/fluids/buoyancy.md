A steel bolt sinks; a steel ship carrying thousands of tonnes floats. A helium balloon rises through air. All three follow from the previous lesson's result, that pressure in a fluid grows with depth. The bottom of any submerged object is deeper than its top, so the fluid pushes up harder than it pushes down.

## The buoyant force

Take the simplest shape: a block of height h whose top and bottom each have area A, held under water with its top at depth d.

::: steps The upward force on a submerged block
$ P_{\text{top}} = P_0 + \rho_f g d, \quad P_{\text{bottom}} = P_0 + \rho_f g (d + h)
Pressure at each face from P = P₀ + ρgh, where ρ_f is the density of the fluid, not of the block.
---
$ F_{\text{net}} = P_{\text{bottom}} A - P_{\text{top}} A
The bottom face is pushed up, the top face down. The pressures on opposite side faces are equal and cancel, so only top and bottom count.
---
$ F_{\text{net}} = \rho_f g h A
Subtract. Atmospheric pressure and the depth d both drop out: the force is the same at any depth.
---
$ F_b = \rho_f V g
The block's volume is V = hA. And ρ_f V is the mass of fluid that would fill the block's space, so F_b is the weight of the fluid the block displaces.
:::

The result does not depend on the block's shape. Imagine any object, and replace it in your mind with fluid of exactly the same shape. That blob of fluid sits still, so the surrounding fluid must be pushing up on it with a force equal to its weight. The surrounding fluid cannot tell what is inside the boundary: the pressures on the surface are the same whether the space is filled by water, steel or wood.

::: equation Archimedes' principle
F_b = \rho_f V_{\text{disp}}\, g
:::

::: note
The buoyant force equals the weight of fluid displaced. It depends on the fluid's density and the submerged volume, not on what the object is made of.
:::

## Sink, float or hover

Compare the buoyant force on a fully submerged object with its weight, ρ_obj V g. The volume and g cancel, leaving a contest of densities:

- If ρ_obj > ρ_f the object sinks. A stone in water still feels a buoyant force, which is why it seems lighter, but not enough.
- If ρ_obj < ρ_f it rises until only part of it is under the surface.
- If they are equal it hovers at any depth. Fish and submarines adjust their average density, with a swim bladder or ballast tanks, to do this.

A floating object settles where the weight of displaced fluid equals its own weight: ρ_f V_sub g = ρ_obj V g. So the fraction submerged is

::: equation Fraction of a floating object below the surface
\frac{V_{\text{sub}}}{V} = \frac{\rho_{\text{obj}}}{\rho_f}
:::

Ice has density about 917 kg/m³ and seawater about 1025 kg/m³, so about 89% of an iceberg is underwater.

## Why ships float

Steel is nearly eight times denser than water, but a ship is not solid steel. Its hull encloses a large volume of air, and what matters is the average density of everything below the waterline. The ship sinks until it has pushed aside its own weight of water. A 100 000-tonne aircraft carrier displaces 100 000 tonnes of seawater, about 98 000 m³. Load it and it sits lower; hole the hull, let water replace the air, and its average density rises past that of water.

The principle works in any fluid, including air. A helium balloon floats because the air it displaces weighs more than the helium plus the rubber. A hot-air balloon works the same way: hot air is less dense than the cooler air around it.

::: history c. 250 BC · Syracuse
Archimedes set out the principle in *On Floating Bodies*, deriving it from the idea that fluid at rest pushes evenly, and used it to analyse when floating shapes are stable. The famous story, that he leapt from his bath shouting "Eureka!" after seeing how to test whether King Hiero's crown was pure gold, comes from Vitruvius two centuries later. Measuring the overflow of water, as Vitruvius described, would have been too imprecise for a crown; in 1586 the young Galileo suggested Archimedes more likely weighed the crown against gold while both hung in water, which uses exactly the buoyancy principle.
:::

::: quiz
? A solid cube of aluminium and one of lead, the same size, hang fully under water. Which feels the greater buoyant force?
- The lead, because it is heavier
- The aluminium, because it is lighter
+ Neither: the buoyant forces are equal
= Both displace the same volume of water, so the upward push is the same. The lead sinks harder only because its weight is larger.
? An ice cube floats in a full glass of water. When it melts, the water level
- rises and the glass overflows
+ stays the same
- falls
= Floating ice displaces exactly its own weight of water. When it melts it becomes that same weight of water, which fills exactly the space it was displacing.
? A wooden block floats with 60% of its volume underwater. What is its density?
- 400 kg/m³
+ 600 kg/m³
- 1600 kg/m³
= The fraction submerged equals ρ_obj/ρ_f, so ρ_obj = 0.6 × 1000 kg/m³.
:::
