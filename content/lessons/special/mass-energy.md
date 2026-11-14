In September 1905, three months after his first paper on relativity, Einstein sent a short sequel, three pages long, to the same journal. Its title was a question: "Does the inertia of a body depend upon its energy content?" Inertia means resistance to being accelerated, which is what mass measures. His answer was yes. A body that gives off an amount of energy E loses an amount of mass E/c². Mass, in other words, is a form of energy, and the square of the speed of light is the rate of exchange between the units we use for the two. This lesson shows how the postulates force a change in the ideas of momentum and energy, derives the energy of a moving body, explains why the energy a body has at rest is more than a mathematical constant, and looks at what the enormous size of c² means in practice.

## Why must Newton's momentum change?

In Newton's physics, the momentum of a body is its mass times its velocity, p = mv, and the total momentum of a set of bodies is conserved: in any collision, the total before equals the total after. Conservation of momentum is one of the most useful laws in physics, and the first postulate says that whatever the correct laws are, they must hold in every inertial frame.

The trouble is that once velocities combine by the relativistic addition rule rather than by simple addition, p = mv no longer behaves. Take a collision in which momentum, defined as mv, is conserved as seen from one frame. Convert every velocity to the frame of an observer moving past, using the relativistic rule, and the totals before and after no longer match. A law that holds in one frame and fails in another would break the first postulate. So either momentum is not conserved, or mv is the wrong definition. Physicists chose to keep the conservation law and find the quantity that obeys it in every frame. It is the **relativistic momentum**:

::: equation Relativistic momentum
p = \gamma m v = \frac{mv}{\sqrt{1 - v^2/c^2}}
:::

Here m is the ordinary mass of the body, the same in every frame, and γ is the familiar Lorentz factor. At everyday speeds γ is almost exactly 1, and the formula becomes Newton's mv, which is why Newton's version worked for two centuries. Near the speed of light it is very different. As v approaches c, γ grows without limit, and so does the momentum. A steady force keeps adding momentum at a steady rate, as Newton's second law says, but each addition buys less and less extra speed. That is why no amount of pushing can bring a body with mass up to the speed of light: it would need infinite momentum.

## What is the energy of a moving body?

Kinetic energy is the energy a body has because it is moving. It can be defined as the work needed to bring the body from rest up to its speed, where work is force multiplied by the distance over which the force acts, added up over the whole push. In Newton's physics that gives the familiar ½mv². With the new momentum, the same calculation gives something else.

::: steps Kinetic energy from relativistic momentum
$ \mathrm{d}K = F\,\mathrm{d}x = \frac{\mathrm{d}p}{\mathrm{d}t}\,\mathrm{d}x = v\,\mathrm{d}p
Over a small distance dx, a force F does work F dx, and that raises the body's kinetic energy by the same amount. Newton's second law in its general form says that force is the rate of change of momentum, F = dp/dt, and this form still holds in relativity. Then dx/dt is the velocity v, so the small gain of energy is v times the small gain of momentum.
---
$ \frac{\mathrm{d}\gamma}{\mathrm{d}v} = \frac{\gamma^3 v}{c^2}
To go further we need to know how γ changes with speed. Writing γ = (1 − v²/c²)^(−1/2) and differentiating by the chain rule gives −½ times (1 − v²/c²)^(−3/2) times −2v/c². The two minus signs cancel and the factor (1 − v²/c²)^(−3/2) is γ³, which gives this result.
---
$ \frac{\mathrm{d}(\gamma v)}{\mathrm{d}v} = \gamma + v\,\frac{\mathrm{d}\gamma}{\mathrm{d}v} = \gamma + \frac{\gamma^3 v^2}{c^2} = \gamma^3
The momentum is mγv, so we need the rate of change of the product γv. The product rule gives γ plus v times dγ/dv, and the previous step supplies dγ/dv. To combine the two terms, write the first as γ³(1 − v²/c²), which is allowed because 1 − v²/c² = 1/γ². Then the v²/c² terms cancel and the whole thing is simply γ³.
---
$ \mathrm{d}K = v \cdot m\gamma^3\,\mathrm{d}v = mc^2 \cdot \frac{\gamma^3 v}{c^2}\,\mathrm{d}v = mc^2\,\mathrm{d}\gamma
The small change of momentum is therefore dp = mγ³ dv, and the first step says dK = v dp = mvγ³ dv. Multiplying and dividing by c² shows that this is mc² times γ³v/c² dv, which by the second step is exactly mc² times the small change dγ. Every small addition to the kinetic energy is mc² times the small increase in γ.
---
$ K = mc^2(\gamma - 1)
Add up all these small changes, from rest, where γ = 1 and the kinetic energy is zero, up to the final speed v. Since each step contributes mc² times its change in γ, the total is mc² times the total change in γ, which is γ − 1.
---
$ \gamma mc^2 = K + mc^2
Move the term −mc² to the other side. The left side, γmc², turns out to be the quantity that is conserved in collisions in every frame, along with the relativistic momentum, so it is the natural definition of the body's total energy E. The equation shows it splitting into two parts: the kinetic energy, which depends on the motion, and a part mc² that is present even when the body is not moving.
---
$ E_0 = mc^2
Set the speed to zero. Then γ = 1, the kinetic energy vanishes, and the total energy that remains is mc². This is called the **rest energy**: the energy a body has simply by having mass.
:::

A check on the result is worth making. For speeds much smaller than c, γ is very close to 1 + v²/(2c²). Then K = mc²(γ − 1) ≈ mc² × v²/(2c²) = ½mv². Newton's kinetic energy is not wrong; it is what the relativistic formula becomes at low speeds, exactly as Newton's momentum was.

Still, the derivation by itself does not prove that the rest energy means anything. In ordinary physics, adding a constant to an energy changes nothing that can be measured, because only changes of energy matter: a ball's height can be measured from the floor or from the table, and its fall releases the same energy either way. So the mc² that appears here might be a bookkeeping constant, like the choice of floor or table. What makes it physical is a different claim: that rest energy can actually be converted into other forms, so that a body which loses energy loses mass. That was the claim of Einstein's three-page paper.

::: history September 1905 · Bern
Einstein's argument used no collisions and no forces. He imagined a body at rest emitting two equal pulses of light in opposite directions, with total energy L. Because the pulses carry equal and opposite momentum, the body stays at rest. Then he described the same emission from a frame moving past at speed v. In his first relativity paper he had worked out how the energy of light changes between frames, and from that result the two pulses together carry energy γL in the moving frame, slightly more than L. In both frames energy must be conserved. Comparing the two accounts, the body's kinetic energy as seen in the moving frame must drop, during the emission, by L(γ − 1), which for small v is ½(L/c²)v². But the body's speed in the moving frame has not changed: it was v before the emission and v after. A kinetic energy ½mv² that drops while v stays fixed can only mean that m has dropped, by L/c². He concluded that "the mass of a body is a measure of its energy content", and suggested that radium salts, which give off energy constantly, might one day show the effect. His argument kept only the leading terms in v/c, and its rigour was argued over for decades. The result never was.
:::

## How much energy is in a gram?

The factor c² is enormous in everyday units, so a small mass corresponds to a huge energy. Take one gram:

::: equation Rest energy of one gram of matter
E = (10^{-3}\,\text{kg})(3.00 \times 10^8\,\text{m/s})^2 = 9 \times 10^{13}\,\text{J}
:::

That is about the energy released by 21 thousand tonnes of TNT, more than the bomb that destroyed Hiroshima, which converted only about 0.7 g of mass into energy. The Sun shines by converting about 4 million tonnes of its mass into energy every second, and it is so large that it can keep doing so for billions of years.

Why, then, did nobody notice the effect before 1905? Because ordinary processes convert a minute fraction of mass. In chemistry the mass changes are real but tiny. Burning a kilogram of coal releases about 3 × 10⁷ J of heat, so the products of combustion, collected together, weigh about 0.3 micrograms less than the coal and oxygen did: one part in about three billion of the coal's mass, far below what any balance can detect. Chemists had confirmed to high precision that mass is conserved in reactions, and to the precision available they were right. In nuclear reactions the fraction is about a million times larger, around 0.1% of the mass in the fission of uranium and about 0.7% in the fusion of hydrogen into helium. That ratio is the reason a nuclear power station needs a few tens of tonnes of fuel a year where a coal station of the same output burns millions of tonnes.

It is tempting to say that in these processes mass "turns into" energy, as if they were two different substances. That is misleading. Mass is a measure of energy. The mass of an atom or a nucleus includes the energy of the fields and motions inside it, and when some of that energy leaves, as heat or light or the motion of fragments, the mass that goes with it leaves too. A hot cup of coffee weighs very slightly more than the same cup when cold. Older textbooks describe a "relativistic mass" γm that grows with speed; modern usage keeps the mass m fixed, the same in every frame, and lets the energy γmc² and momentum γmv do the growing.

## Has E = mc² been checked?

The first direct test came in 1932, when John Cockcroft and Ernest Walton in Cambridge fired protons at lithium nuclei and split each one into two helium nuclei. The fragments flew apart with about 17 MeV of kinetic energy (an MeV, a million electronvolts, is a unit of energy convenient for nuclear processes, about 1.6 × 10⁻¹³ J). The masses of the particles before and after had been measured separately, and the products were lighter by an amount that, multiplied by c², matched the released energy. In 2005 a team compared precisely measured masses of atoms before and after they absorbed a neutron with the energy of the gamma rays emitted, and confirmed E = mc² to about four parts in ten million.

::: note
A body's energy content and its mass are the same thing measured in different units, linked by E = mc². Any energy a body gains or loses, whether as heat, light or the motion of its fragments, changes its mass by that energy divided by c².
:::

::: quiz
? A hot cup of coffee cools down. What happens to its mass?
- Nothing, because only nuclear reactions change mass
+ It decreases, by the heat lost divided by c²
- It increases, because its molecules slow down
= Any energy that leaves a body takes mass E/c² with it, whatever form the energy takes. For a cup of coffee the change is real but billions of times too small to measure.
? Why can no body with mass be accelerated to the speed of light?
- Air resistance becomes too large
+ Its energy γmc² and momentum γmv grow without limit as its speed approaches c
- Its mass becomes zero
= As v approaches c, γ grows without limit, so reaching c would take an infinite amount of energy. Each push adds energy and momentum but less and less speed.
? For speeds much less than c, what does K = mc²(γ − 1) become?
- mc²
- mv
+ ½mv²
= For small v, γ ≈ 1 + v²/(2c²), so K ≈ mc² × v²/(2c²) = ½mv². Newton's kinetic energy is the low-speed form of the relativistic one.
:::
