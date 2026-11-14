In everyday speech "heat" and "temperature" mean much the same thing. In physics they are two different quantities, and for most of the eighteenth century the people studying them had no clear way to tell them apart. The difference is easy to feel once it is pointed out. A cup of boiling water is hotter than a bathtub of warm water, yet if you wanted to melt a large block of ice, the bathtub would melt far more of it. The cup has the higher temperature; the bathtub has far more energy to give up as it cools. This lesson separates the two ideas, shows how each is measured, and uses them to predict what happens when hot and cold things are put together.

## What does a thermometer actually measure?

Start with something everyone has seen. Put a hot spoon into cold water and leave it. The spoon cools, the water warms, and after a while they stop changing. They have reached the same temperature, and once they are there, nothing more happens. This is the defining behaviour of **temperature**: it is the quantity that decides which way energy will flow when two bodies touch. Energy always passes from the body at higher temperature to the one at lower temperature, and the flow stops when the two temperatures are equal. Physicists call that final, unchanging condition **thermal equilibrium**.

A thermometer works because of a fact so ordinary it went unnamed until the 1930s. Suppose a thermometer is placed in body A and settles to some reading. It is then placed in body B and settles to the same reading. Then A and B, put in contact with each other, will exchange no energy at all: they are already in equilibrium. Without this rule, a thermometer reading would tell you only about the thermometer. With it, the reading tells you how the body will behave with every other body that gives the same reading. Because it was recognised after the first and second laws of thermodynamics had already been numbered, and because it logically comes before them, it was called the **zeroth law**.

To build a thermometer you need some property that changes steadily as a body warms. The length of a thread of mercury or alcohol in a narrow glass tube is the familiar choice, but the pressure of a gas sealed in a bulb, or the electrical resistance of a platinum wire, work just as well and are often more precise. Then you need a scale, which means choosing reference points. On the Celsius scale, pure water freezes at 0 °C and boils at 100 °C at normal sea-level pressure, and the interval between is divided into a hundred equal degrees.

The Celsius zero is a convenience, tied to one substance. There is a more natural starting point. As the next lessons will show, the pressure of a gas held at fixed volume falls steadily as it cools, and if you extend the trend, it would reach zero at −273.15 °C. Nothing can be colder than that, so it is called **absolute zero**. The **kelvin** scale keeps the Celsius size of degree but starts counting from absolute zero:

::: equation Kelvin and Celsius
T\,(\text{K}) = T\,(^\circ\text{C}) + 273.15
:::

Room temperature, about 20 °C, is therefore about 293 K. The laws of gases, engines and radiation all take their simplest form when temperature is measured from absolute zero, so most formulas in these lessons need kelvin, and using Celsius in them gives wrong answers.

What is a thermometer really detecting? The lesson on kinetic theory will answer that: temperature measures the average energy of the random, jostling motion of a body's molecules. A hotter body is one whose molecules are, on average, moving more vigorously.

## If temperature is not heat, what is heat?

**Heat** is energy that passes from one body to another because their temperatures differ. The important word is "passes". Heat is energy on the move, in the same way that work is energy on the move, and once it has arrived it is no longer heat. It has become part of the body's **internal energy**, the total energy of all its molecules: the kinetic energy of their motion and the potential energy stored in the forces between them.

This is why physicists avoid saying that a body "contains heat". The same rise in internal energy can be produced in two quite different ways. You can put your hands near a fire, and heat flows into them from the hotter flames. Or you can rub them together, and they warm with no hotter body anywhere nearby: the work done by your muscles against friction becomes internal energy directly. Once your hands are warm, there is no way to tell which method was used. A body has a temperature and an internal energy, but "how much heat it holds" is not a meaningful question.

## How much heat does it take to warm something?

If heat and temperature are different, the next question is how they are related. Experiment shows that the heat Q needed to raise a body's temperature by ΔT (the symbol Δ means "change in") is proportional to the temperature rise and to the mass m, and it depends on what the body is made of.

::: equation Heat and temperature change
Q = mc\,\Delta T
:::

The material enters through c, the **specific heat capacity**: the energy needed to warm one kilogram of the substance by one kelvin. Water's value is about 4180 J/(kg·K), which is unusually large. Iron's is about 450, and copper's about 385. So the energy that warms 1 kg of water by 1 K would warm 1 kg of copper by about 4180/385 ≈ 11 K. Water soaks up a great deal of energy for each degree it warms, and gives back a great deal for each degree it cools. That is why coastal climates are milder than inland ones: the sea warms slowly in summer and cools slowly in winter, damping the swings of the air above it.

Now the bathtub makes sense. Its water is only modestly warm, but its mass is enormous, so mc is huge, and every degree it cools releases a large amount of energy. The cup has a high temperature but a small mass, so it runs out of energy to give before it has melted much ice.

## What temperature do things end up at?

The rule that energy is conserved, combined with Q = mcΔT, lets you predict the outcome when two bodies at different temperatures are brought together. Suppose they are insulated from everything else, so no energy leaks to the surroundings. Then whatever energy the hotter body loses, the cooler one must gain.

::: steps Final temperature of a mixture
$ m_1 c_1 (T_1 - T_f) = m_2 c_2 (T_f - T_2)
Body 1 starts hotter, at T₁, and cools to a final temperature T_f; body 2 starts at T₂ and warms to the same T_f. The left side is the heat body 1 gives out, and the right side is the heat body 2 takes in, each written as mcΔT with the temperature change taken as positive. Setting them equal is simply conservation of energy for an insulated pair.
---
$ m_1 c_1 T_1 - m_1 c_1 T_f = m_2 c_2 T_f - m_2 c_2 T_2
Multiply out both brackets. Nothing has changed in meaning; the products are now separated so that the unknown T_f can be gathered on one side.
---
$ m_1 c_1 T_1 + m_2 c_2 T_2 = (m_1 c_1 + m_2 c_2)\, T_f
Move the two terms containing T_f to the right and the two known terms to the left, adding the same quantities to both sides so the equation stays balanced. The unknown now appears once, multiplied by the total heat capacity of the pair.
---
$ T_f = \frac{m_1 c_1 T_1 + m_2 c_2 T_2}{m_1 c_1 + m_2 c_2}
Divide both sides by m₁c₁ + m₂c₂. The result is a weighted average of the two starting temperatures, where each body's weight is its heat capacity mc. The final temperature always lies between the two starting ones, and it sits closer to the temperature of whichever body has the larger heat capacity.
:::

Here is what that means in practice. Pour 1 kg of water at 80 °C into 3 kg of water at 20 °C. Since it is all water, c cancels from the top and bottom, and the formula gives (1 × 80 + 3 × 20)/(1 + 3) = 140/4 = 35 °C. The larger mass of cool water pulls the result toward its own temperature.

Now replace the hot water with a 1 kg block of iron, also at 80 °C. Iron's heat capacity per kilogram is about 450 J/K against water's 4180, a little less than a ninth. So the block carries far less energy per degree, and putting the numbers in, the 3 kg of water warms only to about 22 °C. The iron and the hot water started at the same temperature, but one changed the bath much more than the other. That is the distinction between temperature and heat in a single experiment.

## How the two ideas were separated

::: history 1760s–1850 · Glasgow, Munich, Manchester
In the 1760s Joseph Black, a professor of medicine and chemistry in Glasgow, noticed two things that a single idea of "hotness" could not explain. Equal masses of different substances needed different amounts of heat to warm by the same amount, which is the specific heat capacity of this lesson. And melting ice soaked up heat steadily without getting any warmer until it had all melted; Black called the hidden heat "latent heat". Both observations forced a distinction between how hot a body is and how much heat it has taken in. Black, like most scientists of his time, pictured heat as **caloric**, an invisible, weightless fluid that flowed from hot bodies to cold ones and could be neither created nor destroyed.

The caloric picture began to crack in 1798. Count Rumford, an American-born soldier running the arsenal in Munich, watched brass cannon being bored out by a drill turned by horses. With a deliberately blunt drill, the friction kept a tank of water boiling for as long as the horses kept walking. If heat were a conserved fluid squeezed out of the metal, the supply should have run out. It never did, and Rumford concluded that heat must be a form of motion. Many were unconvinced, because he could not say how much motion made how much heat.

That number came from James Joule in Manchester during the 1840s. He let falling weights turn a paddle wheel inside an insulated can of water and measured the tiny rise in temperature. However he arranged the experiment, the same amount of work always produced the same amount of warming. His best value, published in 1850, was within about 1% of the modern figure of 4.184 joules per calorie (a calorie being the heat that warms one gram of water by one degree). Heat was not a substance after all. It was energy being transferred, and it could be traded against work at a fixed rate.
:::

::: note
Temperature tells you which way energy will flow between two bodies, and heat is the energy that actually flows. A body has a temperature and an internal energy, but it does not contain heat, because the same internal energy can be reached by heating or by doing work.
:::

::: quiz
? A metal bench and a wooden bench have sat in the same shade all day. The metal one feels colder to the touch. Why?
- The metal is at a lower temperature
+ Metal conducts heat away from your hand much faster than wood
- Wood produces heat of its own
= Both benches are at air temperature. Your skin, which is warmer, loses energy into the metal much faster than into the wood, so your skin cools quickly and you feel cold.
? 2 kg of water at 10 °C is mixed with 2 kg of water at 50 °C, with no losses. What is the final temperature?
- 20 °C
+ 30 °C
- 40 °C
= The two bodies have equal mass and the same specific heat, so they have equal heat capacities and the weighted average becomes a simple average: (10 + 50)/2 = 30 °C.
? Which statement uses "heat" correctly in the physics sense?
- The kettle contains a lot of heat
+ 500 J of heat flowed from the kettle into the cup
- The heat of the room is 20 °C
= Heat is energy transferred because of a temperature difference. A kettle holds internal energy, and the room's hotness is described by its temperature.
:::
