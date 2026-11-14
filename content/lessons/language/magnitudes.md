Physics deals with numbers that span an absurd range. A proton is about 0.000 000 000 000 001 metres across; the observable universe is about 880 000 000 000 000 000 000 000 000 metres across. Writing out zeros is error-prone and hides what matters, so physicists use scientific notation.

## Scientific notation

A number in scientific notation is written as a number between 1 and 10, times a power of ten:

::: equation
a \times 10^n, \quad \text{with } 1 \leq a < 10 \text{ and } n \text{ a whole number}
:::

So 880 000 000 000 000 000 000 000 000 m becomes 8.8 × 10²⁶ m, and the proton becomes 1 × 10⁻¹⁵ m. A negative exponent means "divide by ten that many times": 10⁻³ = 1/1000.

The exponent does the heavy lifting. Multiplying numbers means adding exponents, and dividing means subtracting them:

::: steps How big is the universe in proton widths?
$ \frac{8.8 \times 10^{26}\,\text{m}}{1 \times 10^{-15}\,\text{m}}
Divide the two lengths. The metres cancel, so the answer is a pure number.
---
$ = \frac{8.8}{1} \times 10^{26 - (-15)}
Handle the leading numbers and the powers of ten separately. Dividing powers of ten subtracts the exponents.
---
$ = 8.8 \times 10^{41}
26 + 15 = 41. About 10⁴² protons laid end to end would span the universe.
:::

## Orders of magnitude

The *order of magnitude* of a number is its power of ten, rounded. 8.8 × 10²⁶ is "order 10²⁷" because 8.8 is closer to 10 than to 1 (on a logarithmic scale the dividing line is about 3, since √10 ≈ 3.16). Saying two things "differ by three orders of magnitude" means one is roughly a thousand times the other.

Thinking in orders of magnitude is useful because many questions only need that much precision. Will a bridge hold 10 tonnes or 10 000? Is an effect a 1% correction or a 1000% one? The first digit can wait.

## Estimation

Enrico Fermi was known for producing good answers to questions that seemed to have no data behind them. The method: break the unknown into factors you can guess within a factor of a few, then multiply. Errors tend to partly cancel, because some guesses are too high and others too low.

Example: how many litres of water does a city of a million people use at home per day?

- A shower uses about 10 litres per minute; call it 8 minutes: 80 L.
- Toilets: about 6 flushes at 6 L each: roughly 40 L.
- Cooking, drinking, washing up, laundry: perhaps 30 L.
- Total per person: about 150 L, or 1.5 × 10² L.
- Times 10⁶ people: 1.5 × 10⁸ L per day.

Published figures for European cities are in the range 100–200 L per person per day, so the estimate is right to within a factor of two. That is the goal: the right power of ten, quickly, with every assumption visible.

::: note
Get the power of ten right first. A result that is off by a factor of two is usually fine; one off by a factor of a thousand means a step is wrong.
:::

::: history 1945 · Trinity test, New Mexico
At the first nuclear test, Fermi dropped scraps of paper as the blast wave passed and watched how far they drifted, about 2.5 metres. From that one observation he estimated the yield at around 10 kilotons of TNT. Careful analysis later gave about 21 kilotons. His quick estimate had the right order of magnitude, which was what anyone standing in the desert needed to know.
:::

::: quiz
? What is (3 × 10⁸) × (2 × 10⁻³)?
- 6 × 10⁻²⁴
+ 6 × 10⁵
- 5 × 10⁵
= Multiply the leading numbers (3 × 2 = 6) and add the exponents (8 − 3 = 5).
? A cell is about 10⁻⁵ m across and a person about 1 m tall. By how many orders of magnitude do they differ?
- 4
+ 5
- 6
= 1 ÷ 10⁻⁵ = 10⁵, five powers of ten.
:::
