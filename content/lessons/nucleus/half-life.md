No one can say when a particular radioactive nucleus will decay. It might be in the next second or in a million years, and nothing about the nucleus tells you which. Yet a sample containing a billion billion such nuclei decays on a schedule so reliable that it can date a wooden coffin or the Earth itself. This lesson explains how complete unpredictability at the level of one nucleus produces precise regularity for many, and the bridge between the two turns out to be a single assumption and a few lines of calculus. It then puts the result to work in carbon dating.

## A fixed chance per second

The first clue came from measurements. In 1900 Rutherford studied a radioactive gas given off by thorium, which he called "thorium emanation" and which is now known to be radon-220. He found that its activity, the number of decays per second, fell by half in about a minute. More striking, it fell by half in about a minute whatever amount he started with, and then by half again in the next minute, and again in the next. Every radioactive substance since has shown the same pattern, each with its own characteristic time.

That pattern follows from one simple assumption about individual nuclei: each nucleus has the same probability of decaying in any given second, and that probability never changes. It does not depend on how long the nucleus has existed, how hot it is, or what chemical compound it sits in. The probability per second is called the **decay constant**, written λ (the Greek letter lambda), and it is measured in units of inverse seconds, s⁻¹. A decay constant of 0.01 s⁻¹ means each nucleus has a 1% chance of decaying in the next second.

The phrase "never changes" carries a surprising consequence. A radioactive nucleus has no memory and does not age. One that has already survived a million years is no more likely to decay in the next second than one created a moment ago. That is quite unlike living things, or machines, which wear out. It is more like tossing a coin: having thrown ten tails in a row does not make heads any more likely on the next throw.

Now consider a large number N of identical nuclei. If each has a probability λ of decaying per second, then on average λN of them decay each second. The rate at which the sample shrinks is proportional to how much of it is left. When there are many nuclei the sample shrinks quickly; as they dwindle, it shrinks more slowly. That one statement is enough to work out how N changes with time.

::: steps Deriving the decay law
$ \frac{\mathrm{d}N}{\mathrm{d}t} = -\lambda N
Write the statement as an equation. The left side, dN/dt, is the rate of change of the number of nuclei, how fast N is changing at a given moment. The right side is the number decaying per second, λN, with a minus sign because decays make N smaller. An equation like this, linking a quantity to its own rate of change, is called a differential equation.
---
$ \frac{\mathrm{d}N}{N} = -\lambda\,\mathrm{d}t
Divide both sides by N and multiply both by dt. This is allowed because N is never zero while there is anything left to decay, and it is useful because it gathers everything involving N on the left and everything involving time on the right. Read this way, the equation says that in a short time dt, the fractional change in N, dN/N, is the same, −λ dt, however large N happens to be.
---
$ \int_{N_0}^{N} \frac{\mathrm{d}N'}{N'} = -\lambda \int_0^t \mathrm{d}t'
Now add up all the small changes, which is what integration does. Start at time 0, when there are N₀ nuclei, and stop at a later time t, when N are left; the same interval must be covered on both sides. The primes on N′ and t′ simply mark them as the variables being summed over, so they are not confused with the end points N and t. Since λ is a constant, it can be taken outside the integral.
---
$ \ln N - \ln N_0 = -\lambda t
Carry out both integrals. The integral of 1/N is the natural logarithm, ln N, evaluated between the two end points; this is a standard result of calculus. The integral of dt from 0 to t is just t. What was a statement about rates has become a statement about amounts.
---
$ \ln \frac{N}{N_0} = -\lambda t
Combine the two logarithms using the rule that the difference of two logarithms is the logarithm of their ratio. Now the equation involves only the fraction of nuclei remaining, N/N₀, which is the quantity we care about.
---
$ N = N_0\,e^{-\lambda t}
Finally, undo the logarithm by raising e to the power of each side, since e raised to ln x is x itself. Then multiply by N₀. This is the decay law: the number of nuclei remaining falls off exponentially with time, at a pace set entirely by λ.
:::

This is **exponential decay**, and its defining property is the one Rutherford saw. In any interval of time of a given length, the same fraction of the remaining nuclei decays, whether the sample is large or small, fresh or ancient. That is just what the second step said in a different form: a constant fractional loss in each moment, added up over time.

## Half-life

The decay constant is a probability per second, which is not the most intuitive way to describe a substance. More useful in practice is the time it takes for half of the nuclei to decay, the **half-life**, written t½. It follows directly from the decay law, and it is the quantity Rutherford actually measured.

::: steps Half-life from the decay constant
$ \frac{N_0}{2} = N_0\,e^{-\lambda t_{1/2}}
Use the definition. After one half-life, by definition, half of the original nuclei remain, so N = N₀/2 at time t = t½. Put both into the decay law.
---
$ \tfrac{1}{2} = e^{-\lambda t_{1/2}}
Divide both sides by N₀. The starting amount has dropped out entirely, which is the mathematical reason Rutherford found the same halving time whatever amount he started with. A half-life is a property of the kind of nucleus, not of the size of the sample.
---
$ \ln 2 = \lambda\,t_{1/2}
Take the natural logarithm of both sides. The left side becomes ln(½), which equals −ln 2, and the right side becomes −λt½. The two minus signs cancel, leaving a simple product.
---
$ t_{1/2} = \frac{\ln 2}{\lambda} \approx \frac{0.693}{\lambda}
Divide both sides by λ to get the half-life on its own. The relation is an inverse one: a large probability of decay per second means a short half-life, and a small one means a long half-life. The two numbers carry exactly the same information.
:::

Because the same fraction goes in each equal interval, half-lives stack up neatly. After two half-lives a quarter of the nuclei remain, after three an eighth, and after ten about a thousandth, since 2¹⁰ is 1,024. The range of half-lives in nature is enormous. Polonium-214 has a half-life of about 0.16 milliseconds. Iodine-131, released in reactor accidents, has a half-life of 8 days, which is why the danger it poses fades within a few months. Carbon-14 has a half-life of 5,730 years. Uranium-238 has a half-life of about 4.5 billion years, close to the age of the Earth, which is why about half of the uranium-238 the Earth was born with is still here today.

::: model decay
Watch the count of remaining nuclei. The individual decays are random, but the number left halves once every half-life. Restart and the curve repeats, though the individual nuclei that decay first are different each time.
:::

The model shows the two levels at once. Watch a single nucleus and you see pure chance. Watch the total and you see a smooth curve. The regularity comes from numbers: with very many independent nuclei, the random fluctuations become tiny compared with the total, just as a coin tossed a million times comes up heads very close to half the time, even though no single toss can be predicted.

::: note
Each nucleus decays at random, with a fixed probability per second and no memory of its past. With huge numbers of nuclei that randomness averages out into a precise law, N = N₀e^(−λt), with a fixed half-life t½ = ln 2/λ that does not depend on how much material there is.
:::

## Carbon dating

The decay law becomes a clock as soon as you know how much of a radioactive substance a sample started with. For most materials that is impossible to know, but for anything that was once alive there is a natural way to find out, and it depends on a lucky chain of circumstances in the atmosphere.

Cosmic rays, fast particles from space, constantly strike the upper atmosphere and knock neutrons out of the nuclei there. When one of these neutrons hits a nitrogen-14 nucleus, it can convert it into carbon-14, which is radioactive. The new carbon-14 combines with oxygen to form carbon dioxide, mixes through the air, and is taken up by plants in photosynthesis and then by the animals that eat the plants. Production and decay balance out, so the air holds a nearly steady share of carbon-14: about one carbon atom in a trillion. While an organism lives, it keeps exchanging carbon with its surroundings, so it carries the same share. When it dies, the intake stops. From that moment its carbon-14 is no longer replenished and simply decays, halving every 5,730 years. The fraction remaining is a measure of how long ago it died.

::: history 1946–1949 · Chicago
Willard Libby, a chemist at the University of Chicago who had worked on the Manhattan Project, put the pieces together. He predicted that living matter should contain a small but measurable amount of carbon-14, and then faced the hard part: the decay of so little carbon-14 is extremely weak, easily swamped by background radiation from cosmic rays and the surroundings. He built detectors shielded by thick iron and surrounded by other counters that flagged and rejected cosmic-ray events, sensitive enough to count the faint signal. To test the method he needed samples whose ages were already known, and he used wood from ancient Egyptian tombs, dated from historical records, among others. The radiocarbon ages matched. Libby received the 1960 Nobel Prize in Chemistry, and archaeology gained its first absolute clock for the last fifty thousand years.
:::

To date a sample, measure its carbon-14 per gram of carbon and compare it with the level in living matter. That ratio is N/N₀. Solving the decay law for t gives the age directly:

::: equation Age from the remaining fraction
t = \frac{t_{1/2}}{\ln 2}\,\ln\frac{N_0}{N}
:::

As an example, suppose a bone has 40% of the carbon-14 level of living matter. Then N₀/N is 1/0.4, which is 2.5. The first factor, 5,730 years divided by 0.693, is about 8,270 years, and ln 2.5 is about 0.916, so the age is roughly 8,270 × 0.916, or about 7,600 years. That fits common sense: 40% is a little less than half, so the bone should be a little older than one half-life of 5,730 years, and it is.

The method has limits, and both come straight from the physics. The first is age. After about 50,000 years, close to nine half-lives, only about a four-hundredth of the original carbon-14 remains, and the signal becomes too faint to distinguish reliably from background. The second is the assumption that the atmosphere always held the same share of carbon-14. It did not: the rate of production varies with changes in the Sun's activity and the Earth's magnetic field, and since the nineteenth century, burning fossil fuels, which contain no carbon-14, and testing nuclear weapons have both disturbed it. So real radiocarbon dates are corrected against records of known age, above all the annual growth rings of very old trees, which can be counted back year by year and measured for their carbon-14.

::: quiz
? A sample's activity falls to one eighth in 30 days. What is its half-life?
- 3.75 days
+ 10 days
- 15 days
= One eighth is ½ × ½ × ½, so the activity has halved three times. Three half-lives in 30 days means each half-life is 10 days.
? A nucleus has survived for three half-lives. What is the chance it decays in the next half-life?
- One eighth
+ One half
- Almost certain
= A nucleus has no memory, so its chance of decaying does not change with age. Every surviving nucleus, old or new, has a 50% chance of decaying in the next half-life.
? Why does carbon dating fail for a 200,000-year-old bone?
- Bones contain no carbon
- Carbon-14 decays faster in old samples
+ After about 35 half-lives, almost no carbon-14 is left to measure
= 200,000 years is about 35 half-lives of carbon-14, and (½)³⁵ is about 3 × 10⁻¹¹. That is far too little to detect against the background.
:::
