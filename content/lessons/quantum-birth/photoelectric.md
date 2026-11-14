Planck quantized the energy of the oscillators in the walls of an oven, and hoped light itself could stay a classical wave. In 1905 Einstein went further: light itself comes in packets. His evidence was a small, odd effect in which light knocks electrons out of metal, and it was for this, not relativity, that he received the Nobel Prize.

## The effect

In 1887 Heinrich Hertz, while producing the radio waves that confirmed Maxwell's theory, noticed that sparks jumped a gap more easily when ultraviolet light shone on it. By 1899 J. J. Thomson had shown that light was knocking electrons out of the metal surface. Philipp Lenard then measured them in 1902, in a vacuum tube: light falls on a metal plate, electrons fly off, and a second electrode collects them. Making the collector negative repels the electrons; the voltage that just stops the most energetic ones, the **stopping voltage** V_s, measures their maximum kinetic energy, eV_s.

The results contradicted what a classical wave should do:

- **Threshold.** Below a certain frequency, no electrons come out at all, however bright the light. Red light never frees electrons from zinc, even at enormous intensity; weak ultraviolet does.
- **Energy independent of brightness.** Brighter light releases more electrons, but not faster ones. The maximum kinetic energy depends only on the frequency.
- **No delay.** Electrons appear immediately, even in very dim light.

A classical wave spreads its energy smoothly over the surface. Brighter light should shake electrons harder and give them more energy, any frequency should eventually work if you wait, and in dim light an electron should need to soak up energy for a long time, by classical estimates seconds or more, before it could leave.

## Einstein's photons

Einstein proposed that light of frequency f consists of localized quanta, later called **photons**, each with energy hf, and that a single photon gives all its energy to a single electron. Freeing the electron costs a minimum energy φ, the **work function** of the metal, which depends on the metal and its surface. Whatever is left becomes kinetic energy.

::: steps The photoelectric equation
$ E_{\text{photon}} = hf
Each photon carries Planck's quantum of energy.
---
$ hf = \phi + K
Energy is conserved in each event. The photon's energy pays the cost of escape φ, and the remainder appears as the electron's kinetic energy K. Electrons from deeper inside lose extra energy on the way out, so φ gives the maximum K.
---
$ K_{\max} = hf - \phi
Rearranged. If hf < φ, K would be negative: no single photon can free an electron, which is the threshold. Brightness means more photons, hence more electrons, but each photon still has energy hf, so K_max is unchanged.
---
$ eV_s = hf - \phi
The stopping voltage measures K_max directly.
---
$ V_s = \frac{h}{e}\,f - \frac{\phi}{e}
Plotted against frequency, the stopping voltage is a straight line. Its slope h/e is the same for every metal; only the intercept, set by φ, changes.
:::

Each feature of the experiment follows. The threshold frequency is f₀ = φ/h. There is no delay because one photon, arriving at once, does the job. And the prediction of a straight line with a universal slope was new and testable.

A useful number: hc ≈ 1240 eV·nm, so a photon of wavelength λ (in nanometres) has energy 1240/λ eV. Sodium has a work function of about 2.3 eV, so its threshold wavelength is about 540 nm, green. Red light at 650 nm (1.9 eV) frees nothing; ultraviolet at 250 nm (5.0 eV) frees electrons with up to 2.7 eV of kinetic energy, stopped by 2.7 V.

::: note
Light delivers energy in lumps of hf, absorbed one lump at a time. Frequency sets the size of each lump; brightness only sets how many arrive.
:::

::: history 1905–1916 · Bern and Chicago
Einstein's paper, sent to the journal in March 1905, three months before his first relativity paper, called the light quantum a "heuristic point of view", and almost nobody accepted it. Light's wave nature had been settled by interference experiments for a century. Robert Millikan at Chicago spent about ten years trying to prove the photoelectric equation wrong. He built an apparatus that shaved fresh surfaces of sodium, potassium and lithium with a knife inside a vacuum, to remove the oxide layers that had spoiled earlier measurements. By 1916 he had confirmed the straight line in every detail and measured h from its slope to within about half a percent of Planck's value. Even so he called the photon idea "bold, not to say reckless". Einstein's 1921 Nobel Prize cited "his discovery of the law of the photoelectric effect"; Millikan won in 1923.
:::

## How strong is the evidence?

The photoelectric effect convinced physicists that light energy is exchanged in quanta, but on its own it does not quite prove photons exist: a theory with quantized atoms and classical light waves can reproduce most of it. Stronger evidence came with Arthur Compton's 1923 experiment, in which X-rays scatter off electrons like billiard balls with momentum h/λ, and later with experiments detecting single photons one at a time. Light, it turned out, is neither a classical wave nor a classical particle, a puzzle that would not be resolved until quantum mechanics.

::: model bohr
Each jump in hydrogen emits a photon of one definite energy and colour. Which of these photons could free an electron from sodium, with its work function of 2.3 eV?
:::

::: quiz
? Doubling the intensity of light shining on a metal (frequency above threshold) does what?
- Doubles the maximum kinetic energy of the electrons
+ Doubles the number of electrons emitted, with the same maximum kinetic energy
- Has no effect
= More photons free more electrons, but each photon still carries hf, so K_max = hf − φ is unchanged.
? Light of 400 nm (3.1 eV) falls on a metal with work function 2.1 eV. What is the stopping voltage?
- 3.1 V
- 2.1 V
+ 1.0 V
= K_max = 3.1 eV − 2.1 eV = 1.0 eV, so a 1.0 V reverse voltage stops the fastest electrons.
? Millikan measured stopping voltage against frequency for several metals. What did the graphs have in common?
+ The same slope, h/e
- The same intercept
- The same threshold frequency
= The slope of eV_s = hf − φ is h for every metal; the work function only shifts the line.
:::
