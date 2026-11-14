Friction is the reason Aristotle believed that motion needs a push. It is the force that brings a sliding box to rest and a rolling ball to a stop, and for two thousand years it hid the first law from view. But friction is not only an obstacle. It is what lets you walk, what lets a car's tyres grip the road and its brakes stop it, and what keeps a knot from slipping undone. Its counterpart in air and water is called drag. Neither is one of the fundamental forces of nature; both are the combined effect of enormous numbers of electrical interactions between atoms. Yet both follow simple rules that work surprisingly well, and this lesson explains those rules, where they come from, and what they let you calculate.

## Why does a heavy box resist and then suddenly give way?

Push gently on a heavy box resting on the floor and it does not move. By the first law, the net force on it must be zero, so something is pushing back on the box exactly as hard as you push forward. That something is friction from the floor. Push a little harder and the box still does not move, so the friction must have grown to match your new push. It keeps matching you, up to a point. Then, past some limit, the box suddenly breaks free and slides. And once it is sliding, you usually find that a little less force than that limit is enough to keep it going.

So friction behaves in two distinct ways. While the surfaces are not sliding over each other, it is called **static friction**, and it takes whatever value is needed to prevent sliding, from zero up to some maximum. It is a responsive force, like the push of a table that supports however much weight is put on it. Once the surfaces slide, it is called **kinetic friction**, and it has a roughly fixed size, whatever you do.

Experiments show that both kinds are proportional to the **normal force** N, the force pressing the two surfaces together at right angles to them. For a box on a level floor, N is just the box's weight. Double the load in the box and you double both the maximum static friction and the kinetic friction. The rules are:

::: equation Friction laws
f_s \leq \mu_s N, \qquad f_k = \mu_k N
:::

The numbers μ_s and μ_k (the Greek letter mu) are called the **coefficients of static and kinetic friction**. They are pure numbers, with no units, because they are ratios of one force to another, and they depend on which two materials are in contact. The coefficient of kinetic friction is usually smaller than the static one, which is why the box is easier to keep moving than to start. Some rough values give a sense of scale. Rubber on dry asphalt has μ_s around 0.9 and μ_k around 0.7, which is why car tyres grip so well. Clean, dry steel on steel has μ_s of roughly 0.6 or more. A steel skate blade on ice has a kinetic coefficient of about 0.01 or less, which is why skating is possible at all.

Two features of these rules are surprising. The first is that friction does not depend on the apparent area of contact. A brick slides with the same friction whether it lies on its broad face or stands on its narrow edge, though intuition suggests that more surface should mean more grip. The second is that kinetic friction hardly depends on speed: over a wide range, a box sliding fast feels about the same friction as one sliding slowly.

## Where do these rules come from?

Both surprises have the same explanation, and it lies in what surfaces look like up close. Even a surface polished to a mirror finish is rough at the microscopic scale, covered with tiny hills and valleys. When two such surfaces are put together, they touch only where the tips of the highest hills meet. The true area of contact is a tiny fraction of the apparent area, often less than a thousandth of it.

Now press the surfaces together harder. The tips squash and flatten, and more of them come into contact, so the true contact area grows. It grows in proportion to the load. At each point of real contact the atoms of the two surfaces are close enough to bond to each other, and friction is the force needed to shear those little welded junctions apart so that the surfaces can slide. More load means more true contact, more junctions and more friction: proportional to the normal force, as observed. And turning the brick onto its edge changes nothing that matters. The same weight presses on a smaller apparent area, so each contact tip is squashed harder, and the true contact area comes out the same. Friction does not care about apparent area because the apparent area is not where the surfaces actually touch.

This picture was worked out by Frank Bowden and David Tabor in Cambridge around 1950. It is itself a simplification. Lubricants, surface films, wear and the way junctions form and break all complicate the story, and the friction between two surfaces remains an active field of research. But the simple rules hold well enough that engineers use them every day.

## How can you measure μ_s?

A neat way to find the coefficient of static friction needs only a board and a protractor. Put a block on the board and slowly tilt it until the block starts to slide. The angle at which it slips tells you μ_s directly, as the following argument shows.

::: steps The angle of slipping
$ W_{\parallel} = mg\sin\theta, \qquad W_{\perp} = mg\cos\theta
Tilt the board to an angle θ above the horizontal. The block's weight mg points straight down, but it is easier to work along and across the slope, so split the weight into two components, as in the vectors lesson. The part along the slope, mg sin θ, tries to drag the block downhill; the part perpendicular to the slope, mg cos θ, presses it into the board.
---
$ N = mg\cos\theta
The block does not sink into the board or lift off it, so it has no acceleration perpendicular to the slope. By the second law, the forces in that direction must balance, which means the normal force from the board equals the perpendicular component of the weight. On a steeper slope this component shrinks, so the board presses less hard.
---
$ mg\sin\theta \leq \mu_s mg\cos\theta
Along the slope, static friction points uphill and must balance the downhill component of weight if the block is to stay put. But static friction can be at most μ_s times the normal force. So the block stays at rest only as long as the downhill pull does not exceed that maximum, which is what this inequality says.
---
$ \tan\theta_{\max} = \mu_s
Divide both sides by mg cos θ, which is positive for any angle less than 90°. The mass and g cancel, and sin θ / cos θ is tan θ. The block holds until tan θ reaches μ_s and slips beyond it. Because the mass has cancelled, a heavy block and a light one of the same material slip at the same angle.
:::

So a block that starts sliding when the board reaches 30° has μ_s = tan 30° ≈ 0.58. The same reasoning explains why loose gravel and sand form heaps with a characteristic steepest slope, called the angle of repose: steeper than that, and the grains at the surface slide.

## How far does a car take to stop?

Friction sets the limit on how quickly a car can stop. When a driver brakes hard without skidding, the force slowing the car is static friction between the tyres and the road. (The tyres are rolling, so the rubber in contact with the road is momentarily at rest relative to it.) On a level road the normal force is the car's weight mg, so the largest friction force is μmg, and by the second law the largest deceleration is μmg/m = μg. The mass has cancelled again: a heavy lorry and a light car with the same tyres can, in principle, decelerate equally fast.

To turn that into a distance, use the constant-acceleration equation v² = u² + 2as from the motion chapter, with final speed v = 0 and acceleration −μg. Solving for the distance gives d = u²/(2μg). At 25 m/s (90 km/h) on dry road with μ = 0.7, that is 625/(2 × 0.7 × 9.8), about 46 m. On a wet road with μ = 0.4 it is about 80 m. Both figures leave out the distance the car covers while the driver is reacting, which at this speed adds another 25 m or so for each second of reaction time.

The formula's most important feature is the u². Stopping distance grows with the square of speed, so driving 40% faster roughly doubles the braking distance. That is why a modest increase in speed makes such a large difference in an emergency.

## What changes when the body moves through air?

A body moving through air or water must keep shoving the fluid out of its way, and the fluid pushes back. This resisting force is called **drag**. It always opposes the motion, like kinetic friction, but unlike sliding friction it depends strongly on speed. For most everyday objects moving at everyday speeds, it is well described by:

::: equation Drag force
F_D = \tfrac{1}{2}\rho C_D A v^2
:::

Here ρ (the Greek letter rho) is the density of the fluid, about 1.2 kg/m³ for air at sea level, A is the area of the body facing the flow, and C_D is the **drag coefficient**, a number set by the body's shape, about 0.5 for a sphere and 0.3 for a modern car. The v² is the important part. Double the speed and the drag quadruples. The reason, roughly, is that a faster body pushes aside more air each second and gives each bit of it more speed, so two factors of v multiply.

The speed dependence explains the feather in the free-fall lesson. A body dropped from rest starts with no drag at all, since it is not yet moving, so it accelerates at g. As it speeds up, drag grows, the net downward force shrinks, and the acceleration falls. Eventually the drag equals the weight, the net force is zero, and the body falls at a steady speed called its **terminal velocity**. Setting mg = ½ρC_DAv² and solving gives v_t = √(2mg/(ρC_DA)). A skydiver falling belly-down reaches about 55 m/s, close to 200 km/h. A feather, with tiny mass and a large area, reaches its much smaller terminal velocity almost at once, which is why it drifts down instead of falling.

::: history 1493–1781 · Milan and Paris
Leonardo da Vinci experimented with blocks dragged across tables in the 1490s and found the two central rules: friction is proportional to load and does not depend on the area of contact. But his notebooks were not published, and the findings were lost for two centuries. Guillaume Amontons, a French instrument maker, rediscovered both rules independently and presented them to the French Academy of Sciences in 1699. The independence of area struck the academicians as so implausible that one of them repeated the experiments before the result was accepted. The rules were finally put on a firm footing by Charles-Augustin de Coulomb, a military engineer who later became famous for the law of electric force. In a prize essay of 1781 he reported a long series of careful measurements under different loads, materials and sliding times, and drew the distinction between static and kinetic friction that is still used.
:::

::: note
Friction is proportional to the force pressing two surfaces together: f = μN. Static friction supplies only as much force as is needed to prevent sliding, up to μ_s N, while kinetic friction is roughly fixed at μ_k N. Drag in a fluid grows with the square of speed, and when it balances a falling body's weight it sets a terminal velocity.
:::

::: quiz
? A 10 kg box sits on a floor with μ_s = 0.5. You push it horizontally with 20 N. What is the friction force?
- 49 N
+ 20 N
- 0 N
= The maximum static friction is 0.5 × 98 ≈ 49 N, but static friction supplies only what is needed to prevent sliding. You push with 20 N, so friction is 20 N and the box stays still.
? A brick slides on its large face, then is turned onto its narrow edge. How does the kinetic friction change?
- It increases
- It decreases
+ It stays about the same
= Friction depends on the normal force, not on the apparent area of contact. The brick's weight is unchanged, so the true contact area and the friction are too.
? A car's speed doubles. If the braking friction stays the same, its braking distance
- doubles
+ quadruples
- stays the same
= The braking distance is d = u²/(2μg), which grows with the square of the initial speed. With the same deceleration, a car going twice as fast takes twice as long to stop and covers that time at twice the average speed, so it travels four times as far.
:::
