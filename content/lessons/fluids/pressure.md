Push a drawing pin into a cork board. Your thumb presses on the flat head and the point presses on the cork, and since the pin is not accelerating, the force at both ends is essentially the same. Yet the head merely dents your thumb while the point drives into the board. The difference is the area over which the force is spread: tiny at the point, much larger at the head. The quantity that captures this, force divided by area, is called **pressure**, and it is the natural way to describe how **fluids**, meaning liquids and gases, anything that flows, push on the things around them. This lesson explains what pressure is, why it increases with depth, and how that one fact explains the barometer, the limit of a suction pump and the hydraulic jack.

## Force spread over an area

Pressure is defined as the force pressing straight onto a surface divided by the area of that surface:

::: equation Pressure
P = \frac{F}{A}
:::

Here F is the part of the force that acts perpendicular to the surface, and A is the surface's area. A force pushing along a surface, rather than into it, does not count toward pressure. The definition explains the drawing pin at once: the same F divided by the tiny area of the point gives an enormous pressure, enough to part the fibres of the cork, while divided by the area of the head it gives a pressure your thumb can bear.

The SI unit of pressure is the **pascal**, one newton per square metre, written 1 Pa = 1 N/m². It is a small unit, about the pressure a sheet of paper exerts lying flat on a table. The air around us presses on everything with about 101 000 Pa, or 101 kilopascals, a value also called one **atmosphere**. That is roughly the weight of 10 tonnes resting on every square metre. We do not feel it crushing us because the fluids inside our bodies are at the same pressure and push outward equally, so there is no net force on our skin. We notice it only when the balance is upset, as when our ears pop in a fast lift or a plane.

## A pressure with no direction

One feature of pressure in a fluid at rest is easy to miss: it does not point anywhere. Place a small surface at some point in still water and it feels the same push whether it faces up, down or sideways.

The reason comes from the fact that the fluid is at rest. A fluid, unlike a solid, cannot resist being pushed sideways; any unbalanced force on a small piece of it makes that piece flow. So imagine the pressure were stronger from one side than another at some point. Then a small parcel of fluid there would feel more push on one face than the opposite one, and it would start to move. In a fluid that is genuinely at rest, that cannot be happening, so the pressure at each point must be the same in every direction. That is why it makes sense to speak of "the pressure at a point" in still water without saying which way.

## Why pressure grows with depth

Anyone who has swum to the bottom of a deep pool has felt the pressure on their ears grow as they go down. The cause is weight. Each layer of water has to support all the water above it, so the deeper you go, the more water is pressing down from above, and the greater the pressure. We can make that precise by looking at one column of water and asking what holds it up.

::: steps Pressure at depth h
$ F_{\text{bottom}} = F_{\text{top}} + mg
Picture an imaginary vertical column of fluid, with cross-sectional area A, its top at the surface and its base at depth h. The column is at rest, so the forces on it balance. The fluid around it pushes in on its sides equally from all directions, so those sideways forces cancel. In the vertical direction, the upward push on its base must therefore support both the downward push on its top and the column's own weight, mg.
---
$ m = \rho A h
To find the weight we need the column's mass. The **density** of a substance, written ρ (the Greek letter rho), is its mass per unit volume, so the mass is density times volume, and the volume of the column is its area times its height, Ah. Liquids compress very little even under great pressure, so we can take ρ to be the same at every depth.
---
$ PA = P_0 A + \rho A h g
Now write each force in terms of pressure, using F = PA. The force on the top of the column is the surface pressure P₀ times the area, and P₀ is usually just atmospheric pressure. The force on the base is the unknown pressure P at depth h times the same area. The weight is the mass from the previous step times g.
---
$ P = P_0 + \rho g h
Every term contains the area A, so dividing both sides by A removes it. That cancellation is the important part: the result does not depend on how wide a column we chose, so it describes the pressure at depth h, not a property of our imaginary column. The pressure is the surface pressure plus an extra amount ρgh that grows steadily with depth.
:::

For water, with ρ = 1000 kg/m³ and g ≈ 9.81 m/s², the extra term ρgh comes to about 98 000 Pa for every 10 m of depth, almost exactly one more atmosphere. So a diver at 10 m feels about 2 atmospheres in total, one from the air and one from the water, and at 30 m about 4. That steady increase is why divers must breathe air delivered at the surrounding pressure, and why they must rise slowly at the end of a deep dive.

The formula contains a surprise: it says nothing about the shape or size of the container. A narrow tube filled with water to a height of 10 m produces the same pressure at its base as a lake 10 m deep. That seems wrong at first, because the lake holds vastly more water. But the lake's extra weight is spread over a correspondingly larger floor, so the force on each square metre is the same. Only the vertical depth below the surface matters.

For a gas the same argument works, with one change. Gases compress easily, so the air near the ground, squeezed by all the air above it, is denser than the air higher up. The density in our derivation is then not constant, and atmospheric pressure falls off more and more gently with height rather than in a straight line. It is about half its sea-level value at an altitude of 5.5 km, and about a third at the summit of Everest, which is why climbers there struggle for breath.

## Turning pressure into a height

The formula can also be read backward: a pressure can be measured as the height of a column of liquid that it will support. Fill a long glass tube, closed at one end, with mercury. Cover the open end, turn the tube upside down into a dish of mercury and uncover it. The mercury in the tube sinks a little and then stops, leaving an empty space at the top. The column stops where the pressure from its weight at the level of the dish exactly balances the air pressing down on the open surface of the dish. Since there is no air above the column, P₀ there is zero, and the column's height measures atmospheric pressure directly. This device is a **barometer**. Using mercury's density of 13 600 kg/m³:

::: equation Height of a mercury barometer
h = \frac{P_0}{\rho g} = \frac{101\,300}{13\,600 \times 9.81} \approx 0.76\,\text{m}
:::

That is about 76 centimetres, a convenient length for a glass tube. Mercury was chosen because it is so dense. The same calculation with water, which is 13.6 times less dense, gives a column about 10.3 m tall.

That number explains a puzzle that troubled engineers for a long time. A suction pump cannot lift water more than about 10 m, however powerful it is. The reason is that a suction pump does not really pull water upward at all. It removes air from above the water in a pipe, and the atmosphere, pressing on the open water below, pushes the water up the pipe. Even a perfect pump can only remove all the air, and then the atmosphere pushes as hard as it ever can, which is enough to hold up about 10 m of water and no more.

::: history 1643–1648 · Florence and Clermont-Ferrand
Galileo knew from the well-diggers of Florence that pumps could not raise water much more than about 10 m, but he blamed the limited strength of the vacuum, as if the empty space could only pull so hard before it gave way. In 1643 his former assistant Evangelista Torricelli turned the problem around. He made the mercury tube described above and argued that nothing was pulling at all: the column was pushed up by the weight of the air outside. "We live submerged at the bottom of an ocean of air," he wrote. In France, Blaise Pascal saw how to test this. If air has weight, there should be less of it pressing down at the top of a mountain than at the bottom, so the mercury should stand lower there. Pascal was in poor health, so in 1648 his brother-in-law Florin Périer carried a barometer up the Puy de Dôme, a volcanic peak about 1000 m above the town of Clermont, while a second barometer stayed below. At the summit the mercury stood about 8 cm lower, and it rose again on the way down. The ocean of air had a depth, and its weight could be measured.
:::

## Pascal's principle and the hydraulic jack

The depth formula has one more consequence. The pressure at any point in a still liquid is the surface pressure plus a term that depends only on depth. So if you raise the pressure at the surface by some amount, say by pushing on a piston, the pressure everywhere in the liquid rises by exactly that amount. This is **Pascal's principle**: in a confined liquid at rest, a change in pressure is passed on undiminished to every part of it. It is the principle behind car brakes and the hydraulic jack.

::: steps Force and distance in a hydraulic jack
$ \Delta P = \frac{F_1}{A_1}
A small piston of area A₁ presses on a sealed chamber of liquid with force F₁. By the definition of pressure, this raises the pressure in the liquid under it by F₁/A₁. Pascal's principle says the same increase appears everywhere in the liquid.
---
$ F_2 = \Delta P \cdot A_2 = F_1\frac{A_2}{A_1}
A second, larger piston of area A₂ is connected to the same liquid, so it feels the same pressure increase. The force on it is that pressure times its own, larger, area. The force is multiplied by the ratio of the areas: a piston with 100 times the area pushes out with 100 times the force.
---
$ A_1 d_1 = A_2 d_2
This is not something for nothing. Liquids barely compress, so the volume pushed out of the small cylinder must equal the volume that enters the large one. If the small piston moves down a distance d₁ and the large one rises d₂, those volumes are A₁d₁ and A₂d₂, and they must be equal. So the large piston moves A₁/A₂ times as far as the small one.
---
$ F_1 d_1 = F_2 d_2
Multiplying the force equation by the distance equation shows that force times distance, the work done, is the same at both pistons. The jack trades distance for force: you push gently over a long stroke, and the car rises with a large force over a short one. Energy is conserved, as it must be.
:::

::: note
In a fluid at rest, the pressure at a point is the same in every direction and depends only on the depth: P = P₀ + ρgh. It does not depend on the shape of the container or on how much fluid lies off to the side.
:::

::: quiz
? Two tanks hold water to the same depth. One is a narrow tube, the other a wide pool. How do the pressures at their bottoms compare?
- The pool's is higher because it holds more water
- The tube's is higher because the water is concentrated
+ They are equal
= The pressure P = P₀ + ρgh depends only on depth. The pool does hold more water, but its extra weight is spread over a correspondingly larger floor, so each square metre carries the same load.
? About how deep must a diver go in fresh water for the total pressure to reach 3 atmospheres?
- 10 m
+ 20 m
- 30 m
= The air above the surface already supplies 1 atmosphere, and each 10 m of water adds about one more. Two more atmospheres therefore need about 20 m of water.
? A hydraulic jack has pistons of area 1 cm² and 100 cm². Pushing the small piston down 10 cm raises the large one by how much?
- 10 cm
- 1 m
+ 1 mm
= The liquid volume leaving the small cylinder must enter the large one, so the large piston moves 100 times less far: 10 cm ÷ 100 = 1 mm. The hundredfold gain in force is paid for with a hundredfold loss in distance.
:::
