No one can say when a particular radioactive nucleus will decay. Yet a sample of a billion billion of them decays on a schedule so reliable that it can date a mummy or the Earth. The connection between the two is one line of calculus.

## A fixed chance per second

Rutherford found in 1900 that the activity of "thorium emanation", now known to be radon-220, fell by half about every minute, whatever amount he started with. The pattern holds for every radioactive substance, and it follows from a single assumption: each nucleus has the same probability λ of decaying in each second, no matter how long it has already existed. λ is the **decay constant**, measured in s⁻¹. A nucleus has no memory; one that has survived a million years is no more likely to decay next second than a new one.

If N nuclei are present, about λN of them decay each second. The number of nuclei falls at a rate proportional to how many are left.

::: steps Deriving the decay law
$ \frac{\mathrm{d}N}{\mathrm{d}t} = -\lambda N
The rate of change of N is minus the number decaying per second. The minus sign says N falls.
---
$ \frac{\mathrm{d}N}{N} = -\lambda\,\mathrm{d}t
Divide both sides by N and multiply by dt, separating the variables: everything involving N on the left, everything involving t on the right.
---
$ \int_{N_0}^{N} \frac{\mathrm{d}N'}{N'} = -\lambda \int_0^t \mathrm{d}t'
Add up the small changes from the start, when there are N₀ nuclei at time 0, to a later time t with N nuclei left. The primes mark the integration variables.
---
$ \ln N - \ln N_0 = -\lambda t
The integral of 1/N is the natural logarithm ln N; the integral of dt is t. λ is constant, so it comes outside.
---
$ \ln \frac{N}{N_0} = -\lambda t
The difference of two logarithms is the logarithm of their ratio.
---
$ N = N_0\,e^{-\lambda t}
Raise e to the power of each side, which undoes the logarithm.
:::

This is **exponential decay**: in each equal interval of time, the same fraction of the remaining nuclei decays.

## Half-life

The **half-life** t½ is the time for half of the nuclei to decay. It follows directly from the decay law.

::: steps Half-life from the decay constant
$ \frac{N_0}{2} = N_0\,e^{-\lambda t_{1/2}}
By definition, after one half-life N = N₀/2.
---
$ \tfrac{1}{2} = e^{-\lambda t_{1/2}}
Divide both sides by N₀. The starting amount drops out, so the half-life is the same for any sample size.
---
$ \ln 2 = \lambda\,t_{1/2}
Take the natural log of both sides: ln(½) = −ln 2, and the minus signs cancel.
---
$ t_{1/2} = \frac{\ln 2}{\lambda} \approx \frac{0.693}{\lambda}
Rearrange. A high decay probability means a short half-life.
:::

After two half-lives a quarter remains, after three an eighth, after ten about a thousandth. Half-lives range enormously: about 0.16 ms for polonium-214, 8 days for iodine-131, 5730 years for carbon-14, and 4.5 billion years for uranium-238, which is why about half the Earth's original uranium-238 is still here.

::: model decay
Watch the count. The individual decays are random, but the number left halves every half-life. Restart and the curve repeats.
:::

::: note
Each nucleus decays at random with a fixed probability per second. With huge numbers of nuclei that randomness averages into a precise law, N = N₀e^(−λt), and a fixed half-life t½ = ln 2/λ.
:::

## Carbon dating

Cosmic rays striking the upper atmosphere release neutrons, which convert nitrogen-14 into carbon-14. The carbon-14 forms carbon dioxide, mixes through the air and is taken up by plants and the animals that eat them. While an organism lives, it keeps the same carbon-14 share as the air, about one atom in a trillion. When it dies, intake stops and its carbon-14 decays with a half-life of 5730 years.

::: history 1946–1949 · Chicago
Willard Libby predicted that living matter should contain a measurable amount of carbon-14, built detectors sensitive enough to count its weak beta decays, and tested the method on samples of known age, including wood from an Egyptian tomb. The ages matched. He received the 1960 Nobel Prize in Chemistry.
:::

To date a sample, compare its carbon-14 per gram of carbon with that of living matter, and solve the decay law for t:

::: equation Age from the remaining fraction
t = \frac{t_{1/2}}{\ln 2}\,\ln\frac{N_0}{N}
:::

A bone with 40% of the living level has N₀/N = 2.5, so t = (5730 years ÷ 0.693) × ln 2.5 ≈ 8270 × 0.916 ≈ 7600 years. After about 50,000 years too little remains to measure. The atmosphere's carbon-14 level has varied over time, so real dates are corrected against tree rings and other records of known age.

::: quiz
? A sample's activity falls to one eighth in 30 days. What is its half-life?
- 3.75 days
+ 10 days
- 15 days
= One eighth is (½)³, so 30 days is three half-lives.
? A nucleus has survived for three half-lives. What is the chance it decays in the next half-life?
- One eighth
+ One half
- Almost certain
= Decay has no memory: each half-life, every surviving nucleus has a 50% chance of decaying.
? Why does carbon dating fail for a 200,000-year-old bone?
- Bones contain no carbon
- Carbon-14 decays faster in old samples
+ After about 35 half-lives, almost no carbon-14 is left to measure
= (½)³⁵ is about 3 × 10⁻¹¹, far below what any detector can distinguish.
:::
