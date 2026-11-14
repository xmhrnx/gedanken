Physics deals with numbers that span a range far beyond anything in daily life. A proton is about 0.000 000 000 000 001 7 metres across; the observable universe is about 880 000 000 000 000 000 000 000 000 metres across. Written out like that, the numbers are hard to read, easy to miscopy, and they hide the one thing that usually matters, which is how big they are. This lesson is about the notation physicists use instead, and about a habit of mind that goes with it: getting the size of an answer right before worrying about its details.

## Why the zeros get in the way

Try comparing the two lengths above by eye. To say how many times bigger the universe is than a proton, you would have to count the zeros in each, and a single miscount changes the answer tenfold. The digits that carry real information, the 17 and the 88, are lost in a sea of placeholders. What we want is a way of writing a number that separates its two jobs: saying roughly how large it is, and saying precisely what its leading digits are.

## Scientific notation

**Scientific notation** does exactly that. A number is written as a value between 1 and 10, multiplied by ten raised to some whole-number power:

::: equation
a \times 10^n, \quad \text{with } 1 \leq a < 10 \text{ and } n \text{ a whole number}
:::

The power of ten, n, called the **exponent**, says how many places the decimal point has been moved. So 880 000 000 000 000 000 000 000 000 m becomes 8.8 × 10²⁶ m, because the decimal point has to move 26 places to the left to leave a number between 1 and 10. The proton becomes 1.7 × 10⁻¹⁵ m. A negative exponent means "divide by ten that many times", so 10⁻³ is 1/1000; the minus sign is how the notation records that the decimal point moved to the right instead of the left.

The exponent carries most of the information. The value a only refines the estimate, while n says whether we are talking about atoms, people or galaxies. And the exponents make arithmetic easy, because of a simple fact about powers. Ten to the 3 is three tens multiplied together, and ten to the 2 is two tens; multiply them and you have five tens multiplied together. So multiplying powers of ten adds their exponents, and dividing subtracts them. Here is the comparison that was so awkward with zeros.

::: steps How many proton widths span the universe?
$ \frac{8.8 \times 10^{26}\,\text{m}}{1.7 \times 10^{-15}\,\text{m}}
To ask how many times one length fits into another, divide the larger by the smaller. Both are in metres, so the units cancel and the answer is a pure number, a count of proton widths. That is a useful check that we have set up the right question.
---
$ = \frac{8.8}{1.7} \times 10^{26 - (-15)}
Division can be done in any grouping, so we handle the leading numbers and the powers of ten separately. Dividing the powers of ten subtracts the exponents, for the reason given above. Subtracting a negative exponent is the same as adding it, so the exponent becomes 26 + 15 = 41.
---
$ \approx 5.2 \times 10^{41}
The leading numbers give 8.8 ÷ 1.7 ≈ 5.2, which already lies between 1 and 10, so no further adjustment is needed. About 5 × 10⁴¹ protons laid side by side would stretch across the observable universe.
:::

It is worth pausing on that result. The number 10⁴¹ is not something anyone can picture, but scientific notation lets us handle it with the same ease as 41. That is the point of the notation: it turns enormous ratios into small, manageable exponents.

## Orders of magnitude

Often the exponent alone is enough. The **order of magnitude** of a number is its power of ten after rounding. The universe's 8.8 × 10²⁶ m is "of order 10²⁷", because 8.8 is nearer to 10 than to 1. When people say two things "differ by three orders of magnitude", they mean one is roughly a thousand times the other.

There is a subtlety in that rounding. On an ordinary number line, the halfway point between 1 and 10 is 5.5. But orders of magnitude are about multiplying, not adding, so the fair halfway point is the number that is the same factor above 1 as it is below 10. That number is √10, about 3.16, because 3.16 × 3.16 ≈ 10. So a leading value below about 3 rounds down to the lower power of ten, and one above about 3 rounds up. On that rule the proton, at 1.7 × 10⁻¹⁵ m, is of order 10⁻¹⁵ m.

Why care about such rough figures? Because many real questions only need that much precision. Should a footbridge be designed to carry 10 tonnes or 10 000? Is a newly noticed effect a 1 percent correction or something a thousand times bigger? In each case the first decision is about the power of ten, and getting it wrong makes everything that follows wrong too, however carefully it is calculated. The digits after the first can wait until the order of magnitude is settled.

## Estimation

The same thinking makes it possible to answer questions that seem to have no data behind them. The method is to break the unknown quantity into a chain of factors, each of which you can guess to within a factor of two or three, and multiply them together.

This works better than you might expect, for a reason worth spelling out. When you multiply several rough guesses, their errors also multiply, but they are unlikely to all lean the same way. Some guesses will be too high and others too low, and those partly cancel. So a chain of five factors, each uncertain by a factor of two, usually gives an answer much closer than the worst case of 2⁵ = 32 times off.

Take a concrete question: how many litres of water does a city of a million people use at home each day? Start with one person. A shower uses about 10 litres per minute, and a typical shower lasts perhaps 8 minutes, which gives 80 litres. A toilet uses about 6 litres per flush, and six flushes a day adds roughly 40 litres. Cooking, drinking, washing up and laundry together might add another 30 litres. That makes about 150 litres per person, or 1.5 × 10² L. Multiply by 10⁶ people and the city uses about 1.5 × 10⁸ litres a day, which is roughly sixty Olympic swimming pools.

How good is that? Published figures for household use in European cities mostly fall between about 100 and 200 litres per person per day, so the estimate is right to well within a factor of two. That is the aim of the method: the right power of ten, reached quickly, with every assumption written down where someone can challenge it. If a figure later turns out wrong, you can see which guess was responsible and fix that one.

::: note
Get the power of ten right first. An answer off by a factor of two is usually good enough for a first look, but one off by a factor of a thousand means a step in the reasoning is wrong.
:::

## Fermi in the desert

This style of reasoning is often named after the physicist Enrico Fermi, who was famous among colleagues for producing sound answers from almost nothing, and for setting such problems to his students. One episode shows it at its most striking.

::: history 1945 · Trinity test, New Mexico
In July 1945 the first nuclear bomb was tested in the New Mexico desert, and no one knew in advance how much energy it would release. Fermi was at the base camp, roughly 16 kilometres from the explosion. As the blast wave arrived, about forty seconds after the flash, he let a handful of small pieces of paper fall from shoulder height. In still air they would have dropped straight down; instead, the passing wave carried them about 2.5 metres sideways. Fermi reasoned that the size of that push depended on how much energy had been released so far away, and from it he estimated a yield equivalent to about 10 kilotons of TNT. The careful analyses that came later put the true figure at around 20 kilotons. His estimate was low by about a factor of two, but it had the right order of magnitude, and it came within a minute of the blast, from a few scraps of paper.
:::

That factor of two is exactly the kind of error estimation tolerates. What mattered in the desert that morning was whether the explosion had been a fizzle of a few hundred tons, a success of tens of thousands, or something far larger, and Fermi's paper scraps answered that question.

::: quiz
? What is (3 × 10⁸) × (2 × 10⁻³)?
- 6 × 10⁻²⁴
+ 6 × 10⁵
- 5 × 10⁵
= Multiply the leading numbers to get 3 × 2 = 6, and add the exponents to get 8 + (−3) = 5. The first wrong answer comes from multiplying the exponents instead of adding them.
? A cell is about 10⁻⁵ m across and a person about 1 m tall. By how many orders of magnitude do they differ?
- 4
+ 5
- 6
= Dividing the person's height by the cell's size gives 1 ÷ 10⁻⁵ = 10⁵. A ratio of five powers of ten is a difference of five orders of magnitude.
? Why does a Fermi estimate built from several rough guesses usually come out closer than the worst case?
- Because each guess is secretly accurate
+ Because some guesses are too high and others too low, so their errors partly cancel
- Because multiplying numbers always reduces errors
= The errors of independent guesses rarely all lean the same way. When they are multiplied together, overestimates and underestimates partly offset each other.
:::
