Put your thumb partly over the end of a garden hose and the water shoots out much faster and further. That much is familiar. What is less obvious is that, wherever a flowing fluid speeds up like this, its pressure drops. The link between the speed of a moving fluid and its pressure is called **Bernoulli's principle**, and it explains things as different as how fast water drains from a tank and how engineers measure the flow in a pipe. It is also routinely misused to explain how wings lift aeroplanes. This lesson derives the principle from two ideas we already trust, the conservation of mass and the conservation of energy, and then looks honestly at what it can and cannot explain.

## Why a narrower pipe means faster flow

Start with the first idea. Picture a liquid flowing steadily through a pipe that narrows from a cross-sectional area A₁ to a smaller area A₂. "Steadily" means that the flow at each point does not change with time, even though the liquid itself is moving through.

Liquids barely compress, so a given mass of liquid always takes up the same volume. Now ask how much liquid passes a point in the wide part of the pipe in one second, and how much passes a point in the narrow part. These must be equal. If more entered a stretch of pipe than left it, liquid would have to pile up somewhere inside, and an incompressible liquid in a rigid pipe has nowhere to pile. So the volume flowing past each point per second is the same all along the pipe.

That volume is easy to write down. If the liquid moves at speed v, then in one second a column of liquid of length v slides past any given point, and the volume of that column is its cross-sectional area times its length, Av. Setting the flow equal at the two places gives the **continuity equation**:

::: equation The continuity equation
A_1 v_1 = A_2 v_2
:::

Halve the area of the pipe and the speed of the liquid doubles; the same amount of liquid has to get through a smaller opening in the same time, so it must move faster. That is exactly what your thumb does to the hose.

But this raises a question. If the liquid in the narrow section is moving faster, each kilogram of it has more **kinetic energy**, the energy of motion, ½mv² for a mass m at speed v. That extra energy has to come from somewhere, and something must have pushed the liquid to make it speed up. In a fluid, that push comes from a difference in pressure. A parcel of fluid speeds up only if the pressure behind it is greater than the pressure in front of it, so the pressure must fall as the fluid moves into the narrow, fast section. The second idea, conservation of energy, tells us by how much.

## Deriving Bernoulli's equation

Follow a small parcel of fluid, with volume ΔV, as it travels along the flow from a point 1, where the pressure is P₁, the speed v₁ and the height h₁, to a point 2, where the values are P₂, v₂ and h₂. We make three simplifying assumptions: the flow is steady, the fluid is incompressible, and **viscosity**, the internal friction that makes honey flow slowly, is small enough to ignore. Without friction, no energy is lost as heat, so the work done on the parcel must all go into changing its energy of motion and its energy of height.

::: steps Bernoulli's equation from the work–energy theorem
$ W_1 = P_1 A_1 \Delta x_1 = P_1 \Delta V
The fluid behind the parcel pushes it forward with a force equal to the pressure times the area, P₁A₁. As the parcel's rear end moves forward a short distance Δx₁, that force does work equal to force times distance. Area times distance is the volume the rear of the parcel sweeps through, which is the parcel's volume ΔV, so the work done is simply P₁ΔV.
---
$ W_2 = -P_2 A_2 \Delta x_2 = -P_2 \Delta V
At the front, the fluid ahead pushes back on the parcel with force P₂A₂ while the parcel moves forward a distance Δx₂. A force pointing against the motion does negative work, which is why there is a minus sign. By continuity, the volume swept at the front is the same ΔV as at the rear, because the parcel's volume does not change.
---
$ W_{\text{net}} = (P_1 - P_2)\,\Delta V
Adding the two gives the net work done on the parcel by the surrounding fluid. It is positive when the pressure behind is greater than the pressure ahead, which is exactly the condition for the fluid to be pushed forward and speed up.
---
$ \Delta KE = \tfrac{1}{2}\rho \Delta V (v_2^2 - v_1^2), \quad \Delta PE = \rho \Delta V g (h_2 - h_1)
Next we need the changes in the parcel's energy. Its mass is its density times its volume, ρΔV. Its kinetic energy changes from ½(ρΔV)v₁² to ½(ρΔV)v₂², and its gravitational potential energy, the energy it has because of its height, changes from (ρΔV)gh₁ to (ρΔV)gh₂.
---
$ (P_1 - P_2)\Delta V = \tfrac{1}{2}\rho \Delta V (v_2^2 - v_1^2) + \rho \Delta V g (h_2 - h_1)
With no friction to waste energy, the net work done on the parcel must equal the gain in its kinetic plus potential energy. This is the work–energy theorem applied to the parcel, and it is the step where conservation of energy enters.
---
$ P_1 + \tfrac{1}{2}\rho v_1^2 + \rho g h_1 = P_2 + \tfrac{1}{2}\rho v_2^2 + \rho g h_2
Every term contains ΔV, so dividing by it removes the parcel's size from the problem. Then we move all the quantities belonging to point 1 to the left and all those belonging to point 2 to the right. Because points 1 and 2 could have been any two points along the parcel's path, the combination on each side must take the same value everywhere along that path.
:::

::: equation Bernoulli's equation
P + \tfrac{1}{2}\rho v^2 + \rho g h = \text{constant along a streamline}
:::

## What the equation says

Each term in Bernoulli's equation is an energy per unit volume. The middle term is the kinetic energy of a cubic metre of fluid, the last term is its potential energy, and the pressure P can be thought of as the work the surrounding fluid is ready to do on it. The equation says that, as a parcel moves along, these three can trade with one another, but their total stays fixed.

Two special cases show it is consistent with what we know. If the fluid is not moving at all, the speed term vanishes and the equation becomes P + ρgh = constant, which is the rule from the lesson on pressure: go deeper, lower h, and the pressure rises to compensate. And if the flow stays at one height, the height term is the same at both points, and what is left says that wherever the fluid moves faster, its pressure is lower. That is the precise version of the garden-hose observation.

There is one limitation to note. The constant holds along a **streamline**, the path traced by a single parcel of fluid. Different streamlines can in principle have different constants. In practice, for a steady flow that starts out uniform, such as air approaching a wing from far away or water entering a pipe from a large tank, every streamline begins with the same value, so the same constant applies throughout.

## Draining a tank

Here is the equation at work on a simple question. An open tank of water has a small hole in its side, a depth h below the water's surface. How fast does the water come out?

::: steps The speed of water leaving a tank
$ P_0 + \tfrac{1}{2}\rho v_{\text{s}}^2 + \rho g h = P_0 + \tfrac{1}{2}\rho v^2 + 0
Apply Bernoulli's equation between a point on the water surface and a point in the jet just outside the hole, measuring height from the hole. Both points are open to the air, so both are at atmospheric pressure, P₀. The surface is at height h and moves at a speed v_s; the jet is at height zero and moves at the speed v we want.
---
$ \rho g h = \tfrac{1}{2}\rho v^2
The atmospheric pressure appears on both sides and cancels. If the hole is small compared with the tank, the surface sinks extremely slowly, so its speed v_s is close to zero and its kinetic energy term can be dropped. What remains says that the potential energy lost by falling through height h becomes the jet's kinetic energy.
---
$ v = \sqrt{2gh}
Dividing both sides by ρ, multiplying by 2 and taking the square root gives the speed. The density has cancelled, so the result is the same for any liquid.
:::

This is the same speed that a stone reaches when dropped from a height h, a result we met in the lesson on free fall. That is not a coincidence: in both cases, energy of height becomes energy of motion, and nothing else happens. The result is called Torricelli's law, after Evangelista Torricelli, who found it experimentally in the 1640s, nearly a century before Bernoulli's equation existed to explain it.

## Measuring flow with a narrow throat

The equation also gives a practical way to measure how fast fluid flows through a pipe. Build a short narrow section, a **throat**, into the pipe. By continuity, the fluid moves faster in the throat; by Bernoulli, its pressure there is lower. The pressure difference between the wide part and the throat is easy to measure, for example by the difference in height of liquid in two small tubes connected to the pipe wall. From that difference and the two areas, the continuity and Bernoulli equations together give the flow speed. This device is called a **Venturi meter**. The same principle feeds fuel into older car engines through the carburettor, where fast air in a narrow throat draws petrol in through a jet, and related ideas are used in the pitot tubes that measure an aircraft's airspeed and in many medical flowmeters.

::: history 1738 · St Petersburg and Basel
In the 1720s and early 1730s, Daniel Bernoulli, a member of a Swiss family that produced eight notable mathematicians over three generations, was working at the new Academy of Sciences in St Petersburg. He studied water flowing through pipes by poking small open tubes into the pipe walls and watching how high the water rose in each. The water stood lower where the flow was faster, showing that the pressure there was lower. He explained the effect by the conservation of what was then called "living force", essentially what we now call kinetic energy, and published the work in 1738 in a book called *Hydrodynamica*, which also gave fluid dynamics its name. His reasoning was sound, but he never wrote the equation in the clean form used today. That came in the 1750s from his friend Leonhard Euler, who wrote down the general equations of motion for a fluid without friction, from which Bernoulli's equation follows as a special case.
:::

## How wings make lift, and a myth

One explanation of how wings work appears in many textbooks and is wrong in an instructive way. It runs like this. Air meeting the front of a wing splits, some going over the top and some underneath. The two streams must meet again at the back edge at the same moment. The top surface is curved and so longer than the bottom, so air going over the top must travel faster to arrive in time. By Bernoulli, faster air means lower pressure, so the pressure on top is lower than underneath, and the wing is pushed up.

This is called the **equal-transit-time** explanation, and its central step has no basis. Nothing requires two parcels of air that separate at the front of a wing to meet again at the back. In fact they do not. Wind-tunnel photographs, in which pulses of smoke mark parcels of air, show the air over the top reaching the back edge well before the air underneath, so the upper air moves much faster than the story claims. The story also fails other tests. It cannot explain how a thin, flat sail or a paper aeroplane with no curved top produces lift, and it cannot explain how an aerobatic plane can fly upside down, with its curved surface underneath.

What actually happens is this. A wing, by its shape and by being tilted slightly upward into the oncoming air, at what is called the **angle of attack**, turns the air flowing past it downward. To turn the air downward, the wing must push down on it, and by Newton's third law the air pushes up on the wing with an equal force. That is the lift. The same story can be told in terms of pressure. Air follows a curved path only if the pressure is different on the two sides of the path, in the same way that a car turning a corner needs a sideways force. As air curves over and under the wing, the pressure above the wing drops below atmospheric and the pressure below rises above it, and that difference in pressure, added up over the whole wing, is the lift force. The downward turning of the air and the pressure difference on the wing are not two separate causes; they are one piece of physics described in two ways.

Where does Bernoulli fit? The air above the wing really does move faster, and Bernoulli's equation correctly links that higher speed to the lower pressure there. But the equation is a bookkeeping rule: it tells you that speed and pressure change together, not which one causes the other or why the flow takes the shape it does. To explain lift you need to explain why the air is turned downward, which depends on the wing's shape and angle, and Bernoulli's equation alone does not do that.

::: note
Bernoulli's equation is conservation of energy for a moving fluid: along a streamline, wherever the flow is faster, the pressure is lower. It relates speed and pressure to each other, but on its own it does not explain why the speed changes.
:::

::: quiz
? Water flows through a horizontal pipe that narrows to half its area. What happens in the narrow part?
- The speed halves and the pressure rises
+ The speed doubles and the pressure falls
- The speed doubles and the pressure rises
= By continuity, halving the area doubles the speed. With the height unchanged, the extra kinetic energy term ½ρv² must be balanced by a fall in pressure.
? What is wrong with the equal-transit-time explanation of lift?
- Bernoulli's equation does not apply to air
+ There is no reason the air above and below must reach the back edge together, and in fact the upper air gets there first
- Wings produce lift only by deflecting air, so pressure plays no role
= Nothing forces separated parcels of air to rejoin, and experiments show the upper air arriving first, moving faster than the story predicts. Lift comes from the wing turning air downward, which goes together with lower pressure above the wing and higher pressure below.
? Water leaks from a hole 5 m below the surface of an open tank. About how fast does it come out? (Take g ≈ 10 m/s².)
- 5 m/s
+ 10 m/s
- 50 m/s
= Torricelli's law gives v = √(2gh) = √(2 × 10 × 5) = √100 = 10 m/s. This is the same speed a stone would reach falling 5 m.
:::
