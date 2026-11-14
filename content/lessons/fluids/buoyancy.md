A steel bolt dropped into a harbour sinks straight to the bottom, yet a steel ship carrying thousands of tonnes of cargo floats in the same water. A helium balloon let go at a fair rises out of sight. These look like three different puzzles, but they have one explanation, and it follows directly from the result of the previous lesson: pressure in a fluid grows with depth. Any object immersed in a fluid has its bottom deeper than its top, so the fluid pushes up on the bottom harder than it pushes down on the top. The difference is an upward force called **buoyancy**. This lesson works out how big that force is, why the answer is so simple, and how it decides whether something sinks, floats or hovers.

## Where the upward push comes from

Start with the simplest possible shape, so that the geometry does not get in the way. Take a rectangular block of height h, whose top and bottom faces each have area A, and hold it under water with its top face at depth d. The surrounding water presses on every face of the block. We want the total force all those pressures add up to.

::: steps The upward force on a submerged block
$ P_{\text{top}} = P_0 + \rho_f g d, \quad P_{\text{bottom}} = P_0 + \rho_f g (d + h)
First find the pressure on the top and bottom faces from the depth formula, P = P₀ + ρgh. The top face is at depth d and the bottom face at depth d + h. Here ρ_f is the density of the fluid, not of the block, because it is the surrounding fluid whose weight creates the pressure; the block's own material has nothing to do with it.
---
$ F_{\text{net}} = P_{\text{bottom}} A - P_{\text{top}} A
Each pressure acts over a face of area A, so each face feels a force of pressure times area. The fluid pushes up on the bottom face and down on the top face, so the net upward force is the difference. The four side faces can be left out: each one is matched by an opposite face at the same depths, feeling equal pressures in opposite directions, so their forces cancel in pairs.
---
$ F_{\text{net}} = \rho_f g h A
Now subtract. The atmospheric pressure P₀ appears in both pressures and cancels, and so does the depth d. Only the difference in depth between the faces, h, survives. This means the buoyant force on a fully submerged object is the same however deep it is held, because sinking it further raises the push on its top and bottom by equal amounts.
---
$ F_b = \rho_f V g
The block's volume is V = hA, so the force is ρ_f V g. That expression has a meaning of its own. Density times volume is mass, so ρ_f V is the mass of fluid that would fill the space the block occupies, and multiplying by g turns it into that fluid's weight. The buoyant force equals the weight of the fluid the block has pushed aside, or **displaced**.
:::

## Why the shape does not matter

That derivation used a rectangular block, but the conclusion holds for an object of any shape, and there is an elegant argument that shows why without any further calculation.

Imagine any object submerged in still water, say a stone with an irregular shape. Now, in your mind, remove the stone and fill exactly the space it occupied with water instead. That blob of water, shaped like the stone, simply sits there, neither rising nor sinking, because it is just part of the still water. Its weight pulls it down, so the surrounding water must be pushing up on it with a force exactly equal to that weight.

Now put the stone back. The surrounding water has no way of knowing what is inside the boundary. The pressure at each point on the surface depends only on the depth of that point, so the water presses on the stone's surface exactly as it pressed on the water-shaped blob. The total upward push is therefore the same: the weight of a stone-shaped volume of water. This is **Archimedes' principle**:

::: equation Archimedes' principle
F_b = \rho_f V_{\text{disp}}\, g
:::

Here V_disp is the displaced volume, the part of the object below the surface. For a fully submerged object that is its whole volume; for a floating one it is only the part under water.

::: note
The buoyant force on an object equals the weight of the fluid it displaces. It depends on the density of the fluid and the volume below the surface, not on what the object is made of.
:::

## Sink, float or hover?

Whether something sinks or floats is a contest between two forces: the buoyant force pushing up and the object's weight pulling down. For a fully submerged object of volume V and density ρ_obj, the weight is ρ_obj V g and the buoyant force is ρ_f V g. Both contain V and g, so in the comparison those cancel, and what is left is a contest between two densities.

If the object is denser than the fluid, its weight is greater than the buoyant force, and it sinks. It still feels a buoyant force, which is why a stone lifted from the bottom of a pond seems lighter while it is under water, but the force is not enough to hold it up. If the object is less dense than the fluid, the buoyant force on it when fully submerged is greater than its weight, so it rises until part of it pokes out above the surface. And if the two densities are exactly equal, the forces balance at any depth and the object hovers wherever it is placed. Fish and submarines both exploit this. A fish adjusts the amount of gas in an internal bag called a swim bladder, and a submarine pumps water into or out of ballast tanks; either way, they set their average density to match the water's and stay at the depth they choose.

A floating object comes to rest when the buoyant force on its submerged part exactly matches its whole weight. Writing V_sub for the submerged volume, that balance is ρ_f V_sub g = ρ_obj V g. Cancelling g and dividing both sides by ρ_f V gives the fraction of the object that sits below the surface:

::: equation Fraction of a floating object below the surface
\frac{V_{\text{sub}}}{V} = \frac{\rho_{\text{obj}}}{\rho_f}
:::

The fraction submerged is simply the ratio of the two densities. Ice has a density of about 917 kg/m³ and seawater about 1025 kg/m³, so an iceberg floats with about 917/1025 ≈ 89 percent of its volume below the waterline. The phrase "the tip of the iceberg" is, if anything, generous: only about a ninth of it shows.

## Why a steel ship floats

Steel is nearly eight times denser than water, so by the rule above a solid steel object must sink, and a bolt does. But a ship is not solid steel. Its hull is a thin steel shell enclosing a large volume of air, along with cargo, fuel and people. What decides whether the ship floats is not the density of steel but the average density of the whole ship, its total mass divided by the total volume enclosed by its hull below the deck. Because so much of that volume is air, the average is well below the density of water.

Another way to see it is to watch a ship being lowered into the water. As it settles, it pushes aside more and more water, and the buoyant force grows. It stops sinking at the depth where the water it has pushed aside weighs exactly as much as the ship. A large aircraft carrier of about 100 000 tonnes must therefore displace 100 000 tonnes of seawater, which at 1025 kg/m³ is about 98 000 cubic metres. Load more cargo and the ship settles lower, displacing more water to carry the extra weight. Hole the hull below the waterline and water flows in to replace the air, the average density climbs, and once it passes the density of seawater the ship sinks, steel and all.

## Floating in air

Nothing in the argument required the fluid to be a liquid. Air is a fluid too, with a density of about 1.2 kg/m³ at sea level, and anything in it feels a buoyant force equal to the weight of the air it displaces. For most objects that force is tiny compared with their weight, so we never notice it. But a balloon filled with helium, which is about seven times less dense than air, displaces air that weighs more than the helium and the rubber together, so it rises. A hot-air balloon works the same way without any special gas: heating the air inside the envelope makes it expand and so become less dense than the cooler air outside, and the difference in weight lifts the basket.

## A principle and a legend

::: history c. 250 BC · Syracuse
Archimedes set out the principle in a treatise called *On Floating Bodies*. He derived it from a single assumption about fluids, that a fluid at rest presses evenly and gives way wherever it is pushed harder on one side than another, which is close to the argument about pressure in this chapter. He then used it to work out when floating shapes are stable and when they tip over, a question of real importance for shipbuilders. The story everyone knows is different. King Hiero of Syracuse suspected that a goldsmith had mixed silver into a gold crown, and Archimedes, noticing the water rise as he stepped into his bath, saw how to test it and ran through the streets shouting "Eureka!", "I have found it!" That story comes from the Roman writer Vitruvius, about two centuries later. Vitruvius said Archimedes compared how much water the crown and an equal weight of pure gold made overflow from a full vessel. In 1586 the young Galileo pointed out that the difference in overflow would have been far too small to measure reliably. He suggested that Archimedes more probably hung the crown and an equal weight of gold from opposite ends of a balance and lowered both into water. A crown containing silver, which is less dense than gold, would have more volume for the same weight, so it would feel a larger buoyant force and the balance would tip. That method uses exactly the principle of this lesson, and it is sensitive enough to work.
:::

::: quiz
? A solid cube of aluminium and a solid cube of lead, the same size, are held fully under water. Which feels the greater buoyant force?
- The lead, because it is heavier
- The aluminium, because it is lighter
+ Neither: the buoyant forces are equal
= The two cubes displace the same volume of water, so the water pushes up on them equally. The lead sinks harder only because its weight is larger, not because its buoyant force is smaller.
? An ice cube floats in a glass filled to the brim with water. When the ice melts, what happens to the water level?
- It rises and the glass overflows
+ It stays the same
- It falls
= Floating ice displaces exactly its own weight of water. When it melts it becomes that same weight of water, which exactly fills the space it was displacing, so the level does not change.
? A wooden block floats in fresh water with 60% of its volume below the surface. What is its density?
- 400 kg/m³
+ 600 kg/m³
- 1600 kg/m³
= The fraction submerged equals the ratio of the block's density to the water's, so ρ_obj = 0.6 × 1000 kg/m³ = 600 kg/m³. The 400 kg/m³ option uses the fraction above the surface instead of below it.
:::
