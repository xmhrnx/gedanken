A measurement is a number times a unit. "The table is 2" says nothing; "the table is 2 metres" says something you can check. The unit tells you what was compared against what: here, the table's length against a standard length called the metre.

## The SI base units

Science uses the International System of Units (SI, from the French *Système international*). It picks seven base units and builds every other unit out of them. Mechanics needs only three:

| Quantity | Unit | Symbol |
|---|---|---|
| Length | metre | m |
| Mass | kilogram | kg |
| Time | second | s |

The other four are the ampere (electric current), the kelvin (temperature), the mole (amount of substance) and the candela (brightness). Everything else is *derived*: speed is metres per second (m/s), acceleration is metres per second per second (m/s²), and force is the newton, N, which is shorthand for kg·m/s².

Prefixes scale a unit by powers of ten: kilo (k) means a thousand, milli (m) a thousandth, micro (μ) a millionth. A kilometre is 1000 m; a millisecond is 0.001 s.

## Units are algebra

Units multiply, divide and cancel exactly like symbols in an equation. That makes conversion mechanical: multiply by a fraction equal to 1.

::: steps Converting 90 km/h to m/s
$ 90\,\text{km/h} \times \frac{1000\,\text{m}}{1\,\text{km}}
1000 m and 1 km are the same length, so this fraction equals 1 and multiplying by it changes nothing but the form. The km cancels.
---
$ = 90\,000\,\text{m/h} \times \frac{1\,\text{h}}{3600\,\text{s}}
Same trick for time: 1 h equals 3600 s. Now the h cancels.
---
$ = 25\,\text{m/s}
90 000 ÷ 3600 = 25. Only m and s remain, which is what we wanted.
:::

## Dimensions

A *dimension* is the kind of quantity, independent of which unit you pick. Length is a dimension; metres, feet and light-years are units of it. Physicists write dimensions in square brackets with L for length, M for mass and T for time. So speed has dimensions [L/T], acceleration [L/T²], and force [M·L/T²].

Two rules follow, and they catch a large share of algebra mistakes:

1. You can only add, subtract or equate quantities with the same dimensions. Adding a length to a time is meaningless.
2. Functions like sin, log and eˣ need dimensionless inputs. sin(3 metres) has no meaning.

This check is called *dimensional analysis*. Suppose you derive that a falling object's speed after dropping a height h is v = 2gh, where g is an acceleration. Check it:

::: equation Checking v = 2gh
[g h] = \frac{L}{T^2} \cdot L = \frac{L^2}{T^2} \quad \text{but} \quad [v] = \frac{L}{T}
:::

The two sides disagree, so the formula is wrong. The dimensions suggest the fix: L²/T² is the square of a speed, so the right formula must involve a square root. Indeed, v = √(2gh). Dimensional analysis cannot find the pure number (the 2), but it pins down how the quantities combine.

::: note
Treat units as algebraic symbols. If the units on two sides of an equation do not match, the equation is wrong, whatever the numbers say.
:::

::: history 1999 · Mars orbit
NASA's Mars Climate Orbiter was lost because one team's software reported thruster impulse in pound-force seconds while the navigation software expected newton-seconds. The factor of 4.45 went unnoticed for months of small course corrections, and the spacecraft entered the Martian atmosphere far lower than planned. The 125 million dollar probe was destroyed by a missing unit conversion.
:::

::: quiz
? A formula claims the time for an object to fall a height h is t = h/g. Is it dimensionally correct?
- Yes, because h and g are both physical quantities
+ No: h/g has dimensions T², not T
- Yes, because it gives a positive number
= [h/g] = L ÷ (L/T²) = T². The correct formula is t = √(2h/g).
? What is 72 km/h in m/s?
- 7.2 m/s
+ 20 m/s
- 259 m/s
= 72 × 1000 ÷ 3600 = 20.
? Which expression could be physically meaningful if x is a length and t a time?
- x + t
- sin(x)
+ x/t²
= x/t² is an acceleration. Adding a length to a time, or taking the sine of a length, has no meaning.
:::
