Every measurement is a comparison. When you say a table is 2 metres long, you mean that a standard length, the metre, fits along it twice. Leave off the unit and "the table is 2" tells nobody anything, because they cannot know what was compared against what. This lesson is about those standards: where they came from, how physicists build every other unit out of a handful of them, and how treating units as algebra turns them into a tool for catching mistakes before they cost anything. By the end you should be able to convert any unit into any other without memorising a table, check whether a formula can possibly be right before trusting it, and even guess the shape of a law of nature from nothing more than the units of the quantities it involves.

## Why a measurement needs a unit

A number on its own describes nothing physical. The same table is 2 metres, 200 centimetres or about 6.6 feet long, and all three statements are true. What stays fixed is the length itself; what changes is the yardstick chosen to express it. So a measurement always has two parts, a number and a unit, and the number only means something once the unit says what it counts.

This sounds obvious, but it has a practical side. For most of history, units were local. A foot was the length of some ruler's foot, a pound differed from one market town to the next, and merchants and engineers lost time and money converting between them. Science made the problem sharper, because a result measured in one laboratory is only useful if another laboratory can check it. That requires everyone to agree on the same yardsticks, and to be able to reproduce them anywhere.

There is a deeper puzzle hiding here. Physics deals with a bewildering number of quantities: length, time, speed, mass, force, energy, pressure, power, charge, voltage, and dozens more. If each needed its own standard, kept in a vault and copied by hand, the system would be hopeless. Yet these quantities are not independent of one another. Speed is defined through length and time; force is defined through mass and acceleration; energy through force and distance. So the question becomes: how few standards do you actually need, what should they be tied to so that anyone anywhere can rebuild them, and what can the relationships between units tell you about the physics itself? The answers turn out to be "very few", "the constants of nature", and "a surprising amount".

::: history 1790s · Paris
During the French Revolution, the new government wanted measures that belonged to no king and no region. The Academy of Sciences proposed a length taken from the Earth itself: one ten-millionth of the distance from the North Pole to the equator, measured along the meridian through Paris. Two astronomers, Jean-Baptiste Delambre and Pierre Méchain, spent most of the 1790s surveying the arc from Dunkirk to Barcelona, through war and arrest, to fix that length. Méchain later found a small inconsistency in his own data and kept it secret, and it troubled him for the rest of his life. The metre they produced was slightly short of the intended fraction, but by then the bar made from it was the standard, and the idea of a shared, rational system had taken hold. Over the next two centuries the definitions moved from metal bars to physical constants. Since 1983 the metre has been the distance light travels in a vacuum in 1/299 792 458 of a second, and since 2019 the kilogram has been fixed by a value of Planck's constant rather than by a metal cylinder kept near Paris.
:::

## The SI base units

The system that grew from that project is the International System of Units, called SI after its French name, *Système international*. Its central idea is economy: pick a small number of **base units**, define each of them carefully, and build every other unit out of them. There are seven base units. Mechanics, the study of motion and force, needs only three:

| Quantity | Unit | Symbol |
|---|---|---|
| Length | metre | m |
| Mass | kilogram | kg |
| Time | second | s |

The other four are the ampere for electric current, the kelvin for temperature, the mole for amount of substance and the candela for brightness as the eye perceives it. They will appear later in the course, when those quantities do.

The modern definitions show what a good standard looks like. The second is exactly 9 192 631 770 oscillations of the radiation from a particular transition in the caesium-133 atom, and since every caesium atom is identical, any good laboratory can count the same second. The metre follows by fixing the speed of light, and the kilogram by fixing Planck's constant, whose units are kg·m²/s. No definition depends any longer on an object that could be scratched or lost: a standard is now a number assigned to a constant of nature.

Every other unit is **derived**, meaning it is made by multiplying and dividing base units. Speed is distance divided by time, so its unit is metres per second, m/s. Acceleration is the change in speed per unit of time, so its unit is metres per second per second, m/s². Force gets a name of its own, the newton (N), but the name is only shorthand: one newton is the force that gives one kilogram an acceleration of one metre per second per second, so 1 N = 1 kg·m/s². Nothing new has been introduced; the newton is the base units in disguise.

Because physical quantities range from the tiny to the enormous, SI also has **prefixes** that scale a unit by powers of ten. Kilo (k) means a thousand, milli (m) a thousandth and micro (μ) a millionth. A kilometre is 1000 m and a millisecond is 0.001 s. The prefixes change only how a quantity is written, never the quantity itself, which is why converting between them is a matter of moving a decimal point.

## Units behave like algebra

The most useful fact about units is that they obey the same rules as letters in an equation. They multiply, divide and cancel. A speed in metres per second, multiplied by a time in seconds, gives (m/s) × s, and the seconds cancel to leave metres, which is exactly what a distance should be. The units are doing bookkeeping on your behalf.

That makes conversion mechanical rather than a matter of memory. The trick is to multiply by a fraction whose top and bottom are the same quantity written in different units. Such a fraction equals 1, so multiplying by it cannot change the quantity, only its form. Here is a speed limit converted from kilometres per hour to metres per second.

::: steps Converting 90 km/h to m/s
$ 90\,\text{km/h} \times \frac{1000\,\text{m}}{1\,\text{km}}
We want metres on top instead of kilometres, so we multiply by a fraction with metres on top and kilometres underneath. Because 1000 m and 1 km are the same length, the fraction is equal to 1, and multiplying by it leaves the speed unchanged. The km on top of the original speed and the km underneath the fraction now cancel, as any common factor would.
---
$ = 90\,000\,\text{m/h} \times \frac{1\,\text{h}}{3600\,\text{s}}
The distance is now in metres, but the time is still in hours. One hour is 3600 seconds, so a fraction with hours on top and seconds underneath also equals 1. Multiplying by it puts hours on top to cancel the hours below, leaving seconds in their place.
---
$ = 25\,\text{m/s}
Only metres and seconds remain, which is the unit we wanted, and the arithmetic is 90 000 ÷ 3600 = 25. The fact that the leftover units came out right is itself a check: if we had divided where we should have multiplied, the units would have come out as m·s/h², and we would have known at once that something had gone wrong.
:::

So 90 km/h is 25 metres every second. That number is worth a moment's thought: a car at that speed crosses the length of a tennis court in about one second, which is a more vivid picture of the hazard than the figure on the sign.

The same method handles units that are squared or cubed, which is where conversions most often go wrong. Water has a density of 1 gram per cubic centimetre, and suppose you want that in the SI unit, kilograms per cubic metre. The mass part is easy: a gram is a thousandth of a kilogram. The volume part needs more care. A metre is 100 centimetres, but a cubic metre is a cube 100 centimetres on each side, so it holds 100 × 100 × 100 = 1 000 000 cubic centimetres. The conversion fraction has to be applied three times, once for each power.

::: steps Converting 1 g/cm³ to kg/m³
$ 1\,\frac{\text{g}}{\text{cm}^3} \times \frac{1\,\text{kg}}{1000\,\text{g}}
Grams are on top and we want kilograms there, so we multiply by a fraction equal to 1 with grams underneath. The grams cancel and the mass is now in kilograms. The quantity is 0.001 kg/cm³, which looks small only because a cubic centimetre is a small box.
---
$ \times \left( \frac{100\,\text{cm}}{1\,\text{m}} \right)^3
The unit cm³ means cm × cm × cm, so to cancel it we need three factors of centimetres on top. Cubing the fraction 100 cm / 1 m does exactly that, and since the fraction equals 1, its cube equals 1 too. The cube puts 100³ = 10⁶ in the numerator and m³ underneath.
---
$ = 1000\,\frac{\text{kg}}{\text{m}^3}
Multiplying 0.001 by 10⁶ gives 1000, and only kg and m³ remain. A cubic metre of water, a box one metre along each edge, has a mass of a whole tonne.
:::

That tonne surprises most people, and it is why a full bath is far too heavy to lift. The general lesson is that a factor of 100 in length becomes 10 000 in area and a million in volume; forgetting the power is the commonest slip in conversion, and cancelling units one by one prevents it.

## Dimensions: the kind of quantity

Units can be swapped for one another, but some things cannot. A length can be measured in metres, feet or light-years, yet it remains a length; it can never be turned into a time. The word for this deeper property is **dimension**: the kind of quantity something is, regardless of the unit chosen to measure it.

Physicists write dimensions in square brackets, using L for length, M for mass and T for time. Speed has dimensions [L/T], because it is always some length divided by some time, whatever units those are in. Acceleration has dimensions [L/T²], and force, being mass times acceleration, has dimensions [M·L/T²].

Two rules follow from the idea, and both come from asking what an equation could possibly mean. The first is that you can only add, subtract or equate quantities of the same dimension. An equation claims that two things are the same amount of something; if one side is a length and the other a time, there is no sense in which they can be equal, any more than 3 apples can equal 3 hours. The second rule is that functions such as sin, log and eˣ need inputs with no dimension at all. The reason is that these functions can be written as sums of powers of their input, like x + x²/2 + x³/6 for eˣ; if x were a length, that sum would add metres to square metres to cubic metres, which breaks the first rule. So sin(3 metres) has no meaning, while sin of an angle, which is a ratio of two lengths and so has no dimension, is fine.

## Dimensional analysis

Checking that the dimensions of an equation agree is called **dimensional analysis**, and it is one of the cheapest ways to catch an error. Suppose you work out, or half-remember, that an object dropped from a height h reaches a speed v = 2gh, where g is the acceleration due to gravity. Check the dimensions of each side:

::: equation Checking v = 2gh
[g h] = \frac{L}{T^2} \cdot L = \frac{L^2}{T^2} \quad \text{but} \quad [v] = \frac{L}{T}
:::

The two sides do not match, so the formula cannot be right, whatever numbers you put into it. The mismatch even suggests the repair. The right side has the dimensions of a speed squared, L²/T², so the speed itself must involve a square root. The correct result, derived later in the course, is v = √(2gh).

Notice what the check could and could not do. It could not tell us the factor of 2 inside the root, because a pure number has no dimension and so slips past unseen. What it did pin down was how the quantities must combine: speed goes as the square root of g times h, not as their product. That is often the hardest part of a formula to remember, and dimensional analysis supplies it for free.

::: note
Units obey the rules of algebra: they multiply, divide and cancel, and the two sides of any true equation must have the same dimensions. If the dimensions do not match, the equation is wrong, however reasonable the numbers look.
:::

## Going deeper: laws from dimensions alone

Dimensional analysis can do more than check a formula you already have. Sometimes it can find one you do not. The method is to list the quantities you think the answer depends on, and then ask which combination of them has the right dimensions. Often there is only one.

Take a pendulum: a weight on a string of length L, swinging back and forth under gravity g. How long does one swing take? The period T might depend on the length of the string, on the mass m of the weight, and on the strength of gravity. Suppose the answer is some pure number times L, m and g each raised to some power. The dimensions on the two sides have to agree, and that turns out to be enough to fix the powers.

::: steps The period of a pendulum by dimensional analysis
$ T = k\, L^a\, m^b\, g^c
We guess that the period is a product of powers of the three quantities it could depend on, times a pure number k that dimensional analysis cannot see. The task is to find the exponents a, b and c. This guess is the real assumption in the method: we are betting that nothing else, such as air resistance, matters.
---
$ \mathrm{T}^1 = \mathrm{L}^a\, \mathrm{M}^b \left( \frac{\mathrm{L}}{\mathrm{T}^2} \right)^c
Replace each quantity by its dimensions: the period is a time, L is a length, m is a mass, and g is an acceleration, L/T². The number k has no dimensions and drops out. For the two sides to be the same kind of quantity, the powers of L, of M and of T must each match separately.
---
$ b = 0, \quad a + c = 0, \quad -2c = 1
Mass appears only through m, and there is no mass on the left, so b must be zero. Length appears in L and in g, and the total power of length on the right, a + c, must equal the zero on the left. Time appears only through g, giving a power of −2c, which must equal 1.
---
$ c = -\tfrac{1}{2}, \quad a = \tfrac{1}{2}
Solving these three simple equations gives c = −½ and therefore a = +½. The exponents are now fixed; there was only one way to make the dimensions agree.
---
$ T = k \sqrt{\frac{L}{g}}
Putting the powers back gives the period as a pure number times the square root of L over g. Two results have appeared without any knowledge of how pendulums work: a longer string swings more slowly, as the square root of its length, and the mass of the weight does not matter at all.
:::

A full analysis, in the lesson The Pendulum, gives k = 2π for small swings. For a one-metre string, √(1 / 9.81) is about 0.32 seconds, and 2π times that is about 2.0 seconds. That is why a pendulum about a metre long was used in clocks that ticked once a second: each one-way swing takes one second. Dimensional analysis got everything except the 2π.

::: history 1950 · London
In 1947 the United States released photographs of the first atomic explosion, the Trinity test of July 1945, taken with a high-speed camera and marked with a scale and the time after detonation. The energy of the blast was still a secret. The British physicist G. I. Taylor reasoned that the radius R of the expanding fireball should depend only on the energy released E, the time since the explosion t, and the density of the air ρ. The only combination of those three with the dimensions of length is (Et²/ρ)^⅕, so the energy should be roughly ρR⁵/t². A fireball with a radius of about 130 metres at 25 thousandths of a second, in air of density about 1.2 kg/m³, gives roughly 7 × 10¹³ joules. Taylor, who had worked out the theory in detail during the war, published an estimate equivalent to about 17 kilotons of TNT in 1950, close to the figure that was still classified.
:::

## Common misconceptions

The first misconception is that units are labels you attach at the end, after the real work is done with the numbers. Units carried through every line of a calculation will catch a misplaced division or a forgotten conversion at the moment it happens; units added at the end can only be guessed.

A second is that mass and weight are the same, because both are loosely called "weighing". In physics they are different quantities with different dimensions. Mass, in kilograms, measures how much matter there is and how hard it is to accelerate. Weight is the force gravity exerts on that mass, in newtons, and it changes from place to place. A 70 kg person has a weight of about 690 N on Earth and about 110 N on the Moon, but their mass is 70 kg in both places. A bathroom scale reads in kilograms only because it quietly divides the force by Earth's gravity.

A third concerns quantities with no dimension, such as angles. An angle in radians is an arc length divided by a radius, a length over a length, so it is dimensionless. But dimensionless does not mean that the unit is irrelevant. Degrees and radians are both dimensionless, yet sin(30) is 0.5 if 30 means degrees and about −0.99 if it means radians. Here the unit is a choice of scale, and it still has to be stated.

## What happens when units are ignored

The rules can sound pedantic until you see what breaking them costs. In September 1999, NASA's Mars Climate Orbiter arrived at Mars after a nine-month journey and was lost. The investigation found that ground software written by the spacecraft's builder reported the push from small thruster firings in pound-force seconds, while the navigation software that used those figures expected newton-seconds. One pound-force second is about 4.45 newton-seconds, so every small push was being underestimated by that factor.

No single error was large, which is why it went unnoticed. But the thrusters fired many times during the cruise, and each firing nudged the predicted path a little further from the real one. When the orbiter reached Mars, it passed far lower through the atmosphere than planned, an estimated 57 kilometres up instead of the intended 150 or so, and it was destroyed. The spacecraft had cost about 125 million dollars. Nothing about the physics was hard; a number had simply been passed along without its unit, and the unit was the part that mattered.

## Limits of dimensional reasoning

Dimensional analysis is powerful because it is cheap, and it is limited for the same reason: it knows only the kinds of quantities involved, not how they interact. Three limits matter in practice.

The first, already seen, is that it is blind to pure numbers. It found √(L/g) for the pendulum but could not supply the 2π, and it cannot tell ½mv² from mv² or from 3mv². When the missing number is close to 1, as in Taylor's estimate, the answer is good; when it is 2π, the answer is off by six times.

The second is that it cannot choose the variables for you. The pendulum argument worked because we guessed correctly that only L, m and g matter. Had we included air resistance, there would have been more than one way to make a time, and dimensions alone could not have chosen between them. In general, whenever two quantities in the list can be combined into a dimensionless ratio, such as the angle of swing of a pendulum, the answer may depend on that ratio in any way at all, and only real physics can say how. For the pendulum, the period does grow slowly with the angle of swing, which is why the 2π is exact only for small swings.

The third is that matching dimensions does not guarantee a correct equation, and two different quantities can share the same dimensions. Energy and torque, the turning effect of a force, are both measured in newton-metres, [M·L²/T²], yet they are physically different things, and adding one to the other is nonsense that dimensional analysis will not catch. A dimensional check is a necessary test, not a sufficient one: it can prove an equation wrong but never prove it right.

## Where it shows up

Every later lesson in this course leans on units, and most of them use dimensional analysis at least once, often silently. In Orders of Magnitude, the next lesson, units and powers of ten combine into the habit of estimating before calculating. In Constant Acceleration and Free Fall you will derive the formula v = √(2gh) that the dimensional check above anticipated. F = ma explains why the newton is a kilogram metre per second squared, and Work and Kinetic and Potential Energy introduce the joule, which is a newton-metre, and so a kilogram metre squared per second squared. The Pendulum supplies the 2π that dimensional analysis could not, and Calculating c shows how the speed of light, now built into the definition of the metre, was first worked out from electrical measurements.

Outside physics, the same habits matter wherever numbers carry consequences. Drug doses are calculated per kilogram of body weight, and confusing milligrams with micrograms, a factor of a thousand, is a known cause of medical error. Engineers test scale models of ships and aircraft by matching dimensionless ratios between model and full size, which is dimensional analysis put to work.

::: quiz
? A formula claims that the time for an object to fall a height h is t = h/g. Is it dimensionally correct?
- Yes, because h and g are both physical quantities
+ No: h/g has dimensions of time squared, not time
- Yes, because it gives a positive number
= Dividing a length by an acceleration gives L ÷ (L/T²) = T², which is not a time. The correct formula, t = √(2h/g), takes a square root to fix exactly this.
? What is 72 km/h in m/s?
- 7.2 m/s
+ 20 m/s
- 259 m/s
= Multiplying by 1000 m per km and by 1 h per 3600 s gives 72 × 1000 ÷ 3600 = 20 m/s. The wrong answers come from converting only one of the two units, or converting one of them upside down.
? If x is a length and t a time, which expression could be physically meaningful?
- x + t
- sin(x)
+ x/t²
= x/t² has the dimensions of an acceleration, a perfectly good quantity. Adding a length to a time breaks the rule that only like quantities add, and the sine of a length is meaningless because sine needs a dimensionless input.
? By dimensional analysis, the period of a mass bouncing on a spring depends on the mass m and the spring's stiffness k, which has units of N/m = kg/s². Which form must the period take?
- T ∝ √(k/m)
+ T ∝ √(m/k)
- T ∝ m/k
= m/k has dimensions kg ÷ (kg/s²) = s², so its square root is a time. √(k/m) would be one over a time, and m/k alone would be a time squared.
? How many square centimetres are in one square metre?
- 100
- 1000
+ 10 000
= A square metre is 100 cm on each side, so it holds 100 × 100 = 10 000 square centimetres. Each power of the unit needs its own conversion factor.
:::
