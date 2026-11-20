In everyday speech "heat" and "temperature" mean much the same thing. In physics they are two different quantities, and for most of the eighteenth century the people studying them had no clear way to tell them apart. The difference is easy to feel once it is pointed out. A cup of boiling water is hotter than a bathtub of warm water, yet if you wanted to melt a large block of ice, the bathtub would melt far more of it. The cup has the higher temperature; the bathtub has far more energy to give up as it cools. This lesson separates the two ideas, shows how each is measured, and uses them to predict what happens when hot and cold things are put together. By the end you should be able to say why a thermometer can be trusted at all, why the kelvin scale starts where it does, how long a kettle ought to take to boil, why ice cools a drink so much better than cold water, and why a sea breeze keeps a coastal town mild.

## Hot, and how much

For most of history there was one word for the whole business: things were more or less "hot", and a hotter thing was assumed to hold more of whatever hotness was. The first thermometers, built in Italy in the early 1600s and refined by Daniel Fahrenheit in the 1710s, made the idea look precise. You could now give a number to how hot a body was, and it seemed natural to treat that number as a measure of its supply of heat.

Then the numbers started to misbehave. In the 1720s the Dutch physician Herman Boerhaave asked Fahrenheit to mix equal volumes of water and mercury, one hot and one cold. Boerhaave expected heat to be shared out by volume, so the mixture should have settled exactly halfway between the two starting temperatures. It did not. It settled much closer to the starting temperature of the water, whichever of the two had been the hot one. Mercury is more than thirteen times denser than water, so it was not a matter of "more stuff" either: a small volume of water seemed to push the mixture around far more than the same volume of a heavy metal.

So here is the puzzle. If a thermometer reading measured the amount of heat in a body, equal volumes at the same reading would hold equal amounts, and they would share it equally. They do not. Either the thermometer was measuring something other than quantity of heat, or heat was not one thing. It took about forty years before Joseph Black in Glasgow said clearly what was going on: there are two quantities here, not one. The rest of this lesson is the working out of that answer.

## What does a thermometer actually measure?

Start with something everyone has seen. Put a hot spoon into cold water and leave it. The spoon cools, the water warms, and after a while they stop changing. They have reached the same temperature, and once they are there, nothing more happens. This is the defining behaviour of **temperature**: it is the quantity that decides which way energy will flow when two bodies touch. Energy always passes from the body at higher temperature to the one at lower temperature, and the flow stops when the two temperatures are equal. Physicists call that final, unchanging condition **thermal equilibrium**.

A thermometer works because of a fact so ordinary it went unnamed until the 1930s. Suppose a thermometer is placed in body A and settles to some reading. It is then placed in body B and settles to the same reading. Then A and B, put in contact with each other, will exchange no energy at all: they are already in equilibrium. Without this rule, a thermometer reading would tell you only about the thermometer. With it, the reading tells you how the body will behave with every other body that gives the same reading. Because it was recognised after the first and second laws of thermodynamics had already been numbered, and because it logically comes before them, it was called the **zeroth law**.

To build a thermometer you need some property that changes steadily as a body warms. The length of a thread of mercury or alcohol in a narrow glass tube is the familiar choice, but the pressure of a gas sealed in a bulb, or the electrical resistance of a platinum wire, work just as well and are often more precise. Then you need a scale, which means choosing reference points. On the Celsius scale, pure water freezes at 0 °C and boils at 100 °C at normal sea-level pressure, and the interval between is divided into a hundred equal degrees.

@@@END

The Celsius zero is a convenience, tied to one substance. There is a more natural starting point. As the next lessons will show, the pressure of a gas held at fixed volume falls steadily as it cools, and if you extend the trend, it would reach zero at −273.15 °C. Better still, at low pressures every gas points to the same zero, so a gas thermometer measures something that belongs to temperature itself rather than to one material. Nothing can be colder than that, so it is called **absolute zero**. The **kelvin** scale keeps the Celsius size of degree but starts counting from absolute zero:

::: equation Kelvin and Celsius
T\,(\text{K}) = T\,(^\circ\text{C}) + 273.15
:::

::: model gas
Drag the temperature slider down and watch the pressure. Notice that it falls in proportion to the absolute temperature, so it would reach zero only at 0 K.
:::

Room temperature, about 20 °C, is therefore about 293 K. The laws of gases, engines and radiation all take their simplest form when temperature is measured from absolute zero, so most formulas in these lessons need kelvin, and using Celsius in them gives wrong answers. A temperature *difference* is the same in both units, since the degrees are the same size: a rise of 10 °C is a rise of 10 K. It is only when a formula uses the temperature itself, as in a ratio, that the zero matters. Doubling the absolute temperature of a gas from 20 °C means going to 586 K, which is 313 °C, not 40 °C.

Until 2019 the kelvin was pinned to the triple point of water, the one condition at which ice, liquid and vapour coexist, defined as exactly 273.16 K. It is now defined by fixing the Boltzmann constant, k = 1.380649 × 10⁻²³ J/K, which links temperature to energy per molecule, so the unit rests on a constant of nature rather than on a sample of one substance.

What is a thermometer really detecting? The lesson on kinetic theory will answer that: temperature measures the average energy of the random, jostling motion of a body's molecules. A hotter body is one whose molecules are, on average, moving more vigorously.

## If temperature is not heat, what is heat?

**Heat** is energy that passes from one body to another because their temperatures differ. The important word is "passes". Heat is energy on the move, in the same way that work is energy on the move, and once it has arrived it is no longer heat. It has become part of the body's **internal energy**, the total energy of all its molecules: the kinetic energy of their motion and the potential energy stored in the forces between them.

This is why physicists avoid saying that a body "contains heat". The same rise in internal energy can be produced in two quite different ways. You can put your hands near a fire, and heat flows into them from the hotter flames. Or you can rub them together, and they warm with no hotter body anywhere nearby: the work done by your muscles against friction becomes internal energy directly. Once your hands are warm, there is no way to tell which method was used. A body has a temperature and an internal energy, but "how much heat it holds" is not a meaningful question.

## How much heat does it take to warm something?

If heat and temperature are different, the next question is how they are related. Experiment shows that the heat Q needed to raise a body's temperature by ΔT (the symbol Δ means "change in") is proportional to the temperature rise and to the mass m, and it depends on what the body is made of.

::: equation Heat and temperature change
Q = mc\,\Delta T
:::

The material enters through c, the **specific heat capacity**: the energy needed to warm one kilogram of the substance by one kelvin. Water's value is about 4180 J/(kg·K), which is unusually large. Iron's is about 450, and copper's about 385. So the energy that warms 1 kg of water by 1 K would warm 1 kg of copper by about 4180/385 ≈ 11 K. Water soaks up a great deal of energy for each degree it warms, and gives back a great deal for each degree it cools. That is why coastal climates are milder than inland ones: the sea warms slowly in summer and cools slowly in winter, damping the swings of the air above it. The scale of this is striking. The whole atmosphere has a mass of about 5 × 10¹⁸ kg, and air needs about 1000 J to warm a kilogram by a kelvin. Seawater needs about four times as much per kilogram and is about 800 times denser, so the top three or four metres of the oceans can absorb as much energy per degree as all the air above the planet.

Why should one substance need so much more energy than another? A clue came in 1819, when Pierre Dulong and Alexis Petit compared solids not kilogram for kilogram but atom for atom. A kilogram of copper contains more atoms than a kilogram of lead, and when they worked out the heat capacity per atom, nearly every solid element came out the same: about 25 joules per kelvin for a mole, the standard chemist's count of 6.02 × 10²³ atoms. Check it on iron, whose mole has a mass of about 0.056 kg: 450 × 0.056 ≈ 25 J/K. Copper gives 385 × 0.064 ≈ 25 J/K too. So the energy goes, roughly, into each atom equally, and a light element needs more energy per kilogram simply because a kilogram contains more atoms. Water is light (each molecule is three atoms with a combined mass of only 18 atomic units), and the hydrogen bonds that tie its molecules together soak up extra energy as well, which is why its value is so large. The kinetic theory lesson explains why energy is shared out per atom in this way.

Here is the formula at work on a familiar job. A mug of tea needs about 0.25 kg of water, warmed from 20 °C to 100 °C, a rise of 80 K. The heat needed is Q = 0.25 × 4180 × 80 ≈ 84,000 J, or 84 kJ. A typical electric kettle delivers about 2,000 joules each second, so in principle the job takes about 42 seconds. In practice it takes closer to a minute, and the difference is worth noticing: some of the energy warms the kettle's own element and body, and some leaks to the air while the water heats. The formula gives the minimum, and the minimum is a useful benchmark for any heating gadget.

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

Look at what each body did: the iron cooled by 58 degrees and the water warmed by about 2, yet each exchanged the same energy, about 26 kJ. The body with the small heat capacity swings a long way; the one with the large heat capacity barely moves.

## When heat goes in and the temperature does not move

The rule Q = mcΔT has a striking exception, and it is the one that first convinced Black that heat and temperature are separate. Put a pan of ice and water on a steady flame and stir it. The thermometer sits at 0 °C, and stays there, for as long as any ice remains. Energy is pouring in the whole time, but the temperature does not rise until the last of the ice has melted. The same thing happens at the boiling point: a pan of water boiling hard on full heat stays at 100 °C, and all the energy goes into turning liquid into vapour.

Where does that energy go? Into breaking the bonds that hold the molecules in place. In ice, each water molecule is locked into a crystal by its neighbours. Melting does not make the molecules move faster on average, which is what a temperature rise would mean; it pulls them out of their fixed positions, raising their potential energy instead. The energy needed per kilogram is the **latent heat**, from the Latin for "hidden", because a thermometer does not see it.

::: equation Heat to change phase
Q = mL
:::

For water, the latent heat of melting is about 334,000 J/kg, and the latent heat of boiling is about 2,260,000 J/kg. Compare these with the 4180 J that warms a kilogram of water by one degree. Melting a kilogram of ice takes as much energy as warming the resulting water by 80 degrees. Boiling a kilogram of water away takes more than five times what it took to bring it from freezing to the boil. That is why a scald from steam is worse than one from boiling water: steam condensing on your skin hands over its latent heat as well.

Latent heat explains a familiar observation. Drop 50 g of ice at 0 °C into 300 g of a drink at 25 °C, treating the drink as water. Melting the ice takes 0.05 × 334,000 ≈ 16,700 J. Cooling the drink all the way to 0 °C would release 0.3 × 4180 × 25 ≈ 31,400 J, which is more than enough, so all the ice melts and the final temperature lies above zero. Set up the balance: the drink cooling from 25 °C to T_f gives out 0.3 × 4180 × (25 − T_f); the ice takes in 16,700 J to melt and then 0.05 × 4180 × T_f to warm the meltwater. Equating and solving, 31,350 − 1254 T_f = 16,700 + 209 T_f, so T_f = 14,650/1463 ≈ 10 °C.

Now repeat with 50 g of liquid water at 0 °C instead of ice. With no melting to pay for, the weighted average gives (0.3 × 25)/(0.35) ≈ 21 °C. The same mass at the same temperature cools the drink by 4 degrees as water and by 15 degrees as ice. Most of the cooling comes from the melting, not from the coldness of the ice.

## How the two ideas were separated

::: history 1760s–1850 · Glasgow, Munich, Manchester
In the 1760s Joseph Black, a professor of medicine and chemistry in Glasgow, noticed two things that a single idea of "hotness" could not explain. Equal masses of different substances needed different amounts of heat to warm by the same amount, which is the specific heat capacity of this lesson. And melting ice soaked up heat steadily without getting any warmer until it had all melted; Black called the hidden heat "latent heat". Both observations forced a distinction between how hot a body is and how much heat it has taken in. Black, like most scientists of his time, pictured heat as **caloric**, an invisible, weightless fluid that flowed from hot bodies to cold ones and could be neither created nor destroyed.

The caloric picture began to crack in 1798. Count Rumford, an American-born soldier running the arsenal in Munich, watched brass cannon being bored out by a drill turned by horses. With a deliberately blunt drill, the friction kept a tank of water boiling for as long as the horses kept walking. If heat were a conserved fluid squeezed out of the metal, the supply should have run out. It never did, and Rumford concluded that heat must be a form of motion. Many were unconvinced, because he could not say how much motion made how much heat.

That number came from James Joule in Manchester during the 1840s. He let falling weights turn a paddle wheel inside an insulated can of water and measured the tiny rise in temperature. However he arranged the experiment, the same amount of work always produced the same amount of warming. His best value, published in 1850, was within about 1% of the modern figure of 4.184 joules per calorie (a calorie being the heat that warms one gram of water by one degree). Heat was not a substance after all. It was energy being transferred, and it could be traded against work at a fixed rate.
:::

## Common misconceptions

The commonest mistake is to treat cold as a thing that flows. "Close the door, you are letting the cold in" describes the situation backwards. Nothing called cold enters a house in winter; heat leaves it, from the warm inside to the cold outside. Insulation slows the outflow of energy; it does not block an inflow of anything.

A second mistake is to trust your skin as a thermometer. Your skin senses how fast energy is leaving or entering it, not the temperature of what it touches. That is why a metal bench feels colder than a wooden one in the same shade, and why a tiled bathroom floor feels colder than the bath mat beside it, even though both are at room temperature.@@@END


A third mistake is to assume that a hotter body always has more energy to give. The cup and the bathtub from the opening show otherwise: the energy a body releases in cooling depends on mc as well as on how far it cools. A spark from a sparkler is at over 1000 °C yet lands harmlessly on your hand, because it is so tiny that it carries only a fraction of a joule.
@@@END


## Where the simple rules stop working

Temperature is defined through equilibrium, so it is only meaningful for something that has settled down. A single molecule has a speed but no temperature. A gas just disturbed by an explosion does not have one temperature until collisions have shared the energy out. In a dense material that takes very little time; in the thin gas of space it can take so long that astronomers give the electrons and the ions in one gas different temperatures.

The formula Q = mcΔT is also an approximation. Specific heat capacities change with temperature, slowly near room temperature and dramatically at low temperatures. Dulong and Petit's rule, which works well at room temperature, fails badly in the cold: the heat capacity of diamond at room temperature is already well below the rule, and every solid's heat capacity falls toward zero as it approaches absolute zero. No classical theory could explain this. In 1907 Albert Einstein showed that it follows if the vibrational energy of atoms can only change in fixed steps, as Max Planck had proposed for light, which made the heat capacity of solids one of the first pieces of evidence for quantum theory.

Finally, the formula assumes no change of phase and no chemical reaction. Once ice melts, water boils or wood burns, energy goes in or comes out without the matching temperature change, and the latent heat or the reaction energy has to be added to the account by hand.

## Where it shows up

Every lesson in the rest of this chapter rests on the distinction made here. The Ideal Gas Law uses the kelvin scale to relate a gas's temperature to its pressure and volume. Kinetic Theory explains what temperature is at the molecular level and why heat capacities come out per atom. The First Law turns Joule's discovery, that heat and work are both ways of transferring energy, into an exact accounting rule. Heat Engines and Entropy and the Second Law ask how much of the heat flowing from hot to cold can be turned into work, and why heat never flows the other way of its own accord.

Outside the chapter, Blackbody Radiation shows how the temperature of a body sets the colour of the light it gives out, which is how astronomers take the temperature of stars and how a non-contact thermometer reads your forehead. In daily life, latent heat runs refrigerators, which pump energy by evaporating and condensing a fluid, and sweating cools you because each gram that evaporates carries off more than two thousand joules.

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
? A pan of water is boiling steadily at 100 °C. You turn the burner up. What happens?
- The water gets hotter than 100 °C
+ The water boils away faster but stays at 100 °C
- The water stops boiling
= At the boiling point, extra energy goes into latent heat, turning liquid into vapour. The temperature stays fixed until the water is gone; only the rate of boiling increases.
? A 1 kg block of aluminium (c ≈ 900 J/(kg·K)) and a 1 kg block of lead (c ≈ 130 J/(kg·K)) each absorb 9000 J. Which statement is right?
- Both warm by the same amount, because they have the same mass
- The aluminium warms more, because it has the larger c
+ The lead warms by about 69 K and the aluminium by 10 K
= ΔT = Q/(mc): 9000/900 = 10 K for aluminium and 9000/130 ≈ 69 K for lead. A large specific heat means a small temperature change for the same energy.
:::
