Every measurement is a comparison. When you say a table is 2 metres long, you mean that a standard length, the metre, fits along it twice. Leave off the unit and "the table is 2" tells nobody anything, because they cannot know what was compared against what. This lesson is about those standards: where they came from, how physicists build every other unit out of a handful of them, and how treating units as algebra turns them into a tool for catching mistakes before they cost anything.

## Why a measurement needs a unit

A number on its own describes nothing physical. The same table is 2 metres, 200 centimetres or about 6.6 feet long, and all three statements are true. What stays fixed is the length itself; what changes is the yardstick chosen to express it. So a measurement always has two parts, a number and a unit, and the number only means something once the unit says what it counts.

This sounds obvious, but it has a practical side. For most of history, units were local. A foot was the length of some ruler's foot, a pound differed from one market town to the next, and merchants and engineers lost time and money converting between them. Science made the problem sharper, because a result measured in one laboratory is only useful if another laboratory can check it. That requires everyone to agree on the same yardsticks, and to be able to reproduce them anywhere.

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

## What happens when units are ignored

The rules can sound pedantic until you see what breaking them costs. In September 1999, NASA's Mars Climate Orbiter arrived at Mars after a nine-month journey and was lost. The investigation found that ground software written by the spacecraft's builder reported the push from small thruster firings in pound-force seconds, while the navigation software that used those figures expected newton-seconds. One pound-force second is about 4.45 newton-seconds, so every small push was being underestimated by that factor.

No single error was large, which is why it went unnoticed. But the thrusters fired many times during the cruise, and each firing nudged the predicted path a little further from the real one. When the orbiter reached Mars, it passed far lower through the atmosphere than planned, an estimated 57 kilometres up instead of the intended 150 or so, and it was destroyed. The spacecraft had cost about 125 million dollars. Nothing about the physics was hard; a number had simply been passed along without its unit, and the unit was the part that mattered.

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
:::
