Push a drawing pin into a board and the same force that hurts your thumb at the flat end drives the sharp end into wood. What differs is the area the force is spread over. That ratio, force per area, is *pressure*, and in a fluid (a liquid or gas, anything that flows) it is the natural quantity to work with.

## Force per area

::: equation Pressure
P = \frac{F}{A}
:::

Here F is the force pressing perpendicular to a surface and A is the surface's area. The SI unit is the pascal: 1 Pa = 1 N/m², a small amount, roughly the pressure of a sheet of paper lying on a table. The air around you presses at about 101 kPa, or 1 atmosphere: about 10 tonnes-weight on every square metre. You do not feel it because the fluids in your body push back equally.

Pressure in a fluid at rest has no direction. A small surface placed anywhere in the fluid feels the same push whichever way it faces. If it did not, the fluid next to it would be pushed sideways unevenly and would start to flow, so a fluid at rest could not stay at rest.

## Why pressure grows with depth

A diver feels the pressure rise as she descends. The reason is weight: each layer of fluid has to hold up all the fluid above it.

::: steps Pressure at depth h
$ F_{\text{bottom}} = F_{\text{top}} + mg
Picture a vertical column of fluid with cross-section A and height h, its top at the surface. It is at rest, so the upward force on its base must balance the downward force on its top plus its weight. Sideways forces on its walls cancel.
---
$ m = \rho A h
The mass of the column is its density ρ (mass per volume) times its volume Ah. Liquids barely compress, so ρ is the same at every depth.
---
$ PA = P_0 A + \rho A h g
Write each force as pressure times area: P₀ is the pressure at the surface (usually atmospheric), P the pressure at depth h.
---
$ P = P_0 + \rho g h
Divide by A. The area has cancelled, so the answer holds for a column of any width.
:::

For water, ρ = 1000 kg/m³ and g = 9.81 m/s², so ρgh adds about 98 kPa for every 10 m. A diver at 10 m feels roughly 2 atmospheres, at 30 m roughly 4.

The result contains no reference to the container's shape. A thin tube of water 10 m tall exerts the same pressure on its base as a lake 10 m deep. Only the vertical depth below the surface matters. For a gas the same argument applies, but density falls with height, so atmospheric pressure drops gradually rather than linearly: it is about half its sea-level value at 5.5 km.

## Measuring the atmosphere

Turn the formula around and pressure becomes a height. Fill a long tube with mercury, close it, invert it into a dish of mercury and open the bottom. The column drops until the weight of mercury balances the air pressing on the dish, leaving vacuum at the top. Setting ρgh = P₀ with mercury's density, 13 600 kg/m³:

::: equation Height of a mercury barometer
h = \frac{P_0}{\rho g} = \frac{101\,300}{13\,600 \times 9.81} \approx 0.76\,\text{m}
:::

With water instead the column would be about 10.3 m tall. This is why a suction pump cannot lift water higher than about 10 m: the pump does not pull water up, it removes air above the water and lets the atmosphere push it up, and the atmosphere can only push so hard.

::: history 1643–1648 · Florence and Clermont
Galileo knew pump-makers could not raise water more than about 10 m but blamed the limited strength of a vacuum. In 1643 his pupil Evangelista Torricelli made the mercury tube above and argued instead that "we live submerged at the bottom of an ocean of air", whose weight holds the column up. Blaise Pascal saw the test: if air has weight, there is less of it above a mountain top. In 1648 his brother-in-law Florin Périer carried a barometer up the Puy de Dôme, about 1000 m above the town, and the mercury fell by roughly 8 cm.
:::

## Pascal's principle

Because pressure at a point depends only on the depth and the surface pressure, raising the pressure anywhere in a closed liquid raises it by the same amount everywhere. Push on a small piston of area A₁ with force F₁ and a large piston of area A₂ connected to the same liquid feels the same added pressure, so it pushes out with F₂ = F₁A₂/A₁. Car brakes and hydraulic jacks work this way. Energy is still conserved: the small piston must move A₂/A₁ times farther than the large one, so force times distance is the same at both ends.

::: note
Pressure in a still fluid depends only on depth: P = P₀ + ρgh. Not on direction, not on the container's shape, not on how much fluid lies to the side.
:::

::: quiz
? Two tanks have the same depth of water. One is a narrow tube, the other a wide pool. How do the pressures at their bottoms compare?
- The pool's is higher because it holds more water
- The tube's is higher because the water is concentrated
+ They are equal
= P = P₀ + ρgh depends only on depth. The extra weight of water in the pool is spread over a larger base.
? About how deep must a diver go in fresh water for the total pressure to reach 3 atmospheres?
- 10 m
+ 20 m
- 30 m
= The surface already supplies 1 atmosphere; each 10 m of water adds about one more.
? A hydraulic jack has pistons of area 1 cm² and 100 cm². Pushing the small one down 10 cm raises the large one by
- 10 cm
- 1 m
+ 1 mm
= The force is multiplied by 100, so the distance is divided by 100: work in equals work out.
:::
