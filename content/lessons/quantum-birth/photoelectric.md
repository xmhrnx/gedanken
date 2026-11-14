Planck introduced energy lumps in 1900 as a property of the oscillators in the walls of a hot oven, and he hoped that light itself could remain an ordinary, continuous wave. In 1905 Einstein proposed something much more radical: that light itself travels and is absorbed in lumps. His evidence was a small and rather obscure effect, in which light shining on a metal knocks electrons out of it. The proposal contradicted a century of experiments showing that light is a wave, and almost nobody accepted it for years. It was nonetheless right, and it was for this work, not for relativity, that Einstein received the Nobel Prize. This lesson describes the effect, explains why a wave theory of light struggled with it, and shows how a single equation with light lumps accounts for everything.

## How was the effect discovered?

In 1887 Heinrich Hertz was producing the radio waves that confirmed Maxwell's theory of electromagnetism, using sparks jumping across a small gap as both transmitter and detector. He noticed, almost in passing, that a spark jumped more easily when ultraviolet light from another spark fell on the gap. Something about the light was helping electric charge escape from the metal. By 1899 J. J. Thomson, who two years earlier had identified the electron, showed that the escaping charges were electrons: light was knocking them out of the metal surface. This is the **photoelectric effect**.

In 1902 Philipp Lenard measured it in detail. His apparatus was a glass tube with the air pumped out, containing a metal plate and a second electrode facing it. Light falling on the plate freed electrons, some of which crossed to the second electrode and made a small current that could be measured. Lenard then made the second electrode negative, so that it pushed the electrons back. As he increased this opposing voltage, fewer and fewer electrons made it across, until at some voltage none did. The voltage that just stops even the fastest electrons is called the **stopping voltage** V_s. An electron of charge e climbing against a voltage V_s loses energy eV_s, so the stopping voltage measures the largest kinetic energy any of the freed electrons had: K_max = eV_s.

## Why was it hard to explain with waves?

In the classical picture, light is a continuous electromagnetic wave, and its energy is spread smoothly over the wavefront. Brighter light means a wave with a larger electric field, which shakes the electrons in the metal harder. From that picture follow three clear expectations. Brighter light should give the freed electrons more energy. Any frequency of light should free electrons eventually, if it is bright enough or shines long enough, since the energy just accumulates. And in very dim light there should be a delay, because an electron must soak up enough energy from the thin, spread-out wave before it can escape. Classical estimates of that delay, for very faint light, came to seconds or longer.

The experiments disagreed on every point. Lenard found that brighter light released more electrons, but not faster ones: the stopping voltage did not depend on the intensity at all. It did depend on the kind of light, with ultraviolet giving more energetic electrons than visible light. For each metal there was a threshold: below a certain frequency, no electrons came out at all, however bright the light. Red light, at any intensity, frees no electrons from zinc, while even weak ultraviolet does. And electrons appeared the instant the light was switched on, even when it was extremely dim; later experiments put any delay below a few billionths of a second.

## What did Einstein propose?

Einstein's paper of 1905 started not from the photoelectric effect but from a thermodynamic argument about blackbody radiation at high frequencies, which suggested that such radiation behaves, in some respects, like a gas of independent particles. He proposed that light of frequency f consists of localized lumps of energy, each carrying hf, Planck's quantum. These lumps were later named **photons**. A lump is absorbed whole or not at all, and in the photoelectric effect a single photon gives all its energy to a single electron.

Freeing an electron from a metal costs energy, because the electron is held in by the attraction of the positive ions around it. The minimum energy needed to pull an electron out through the surface is called the metal's **work function**, written φ (the Greek letter phi). It depends on the metal and on the condition of its surface; for most metals it is a few electronvolts, where an electronvolt, 1.60 × 10⁻¹⁹ J, is the energy an electron gains crossing one volt. Whatever energy the photon brings beyond the work function is left over as the electron's kinetic energy.

::: steps The photoelectric equation
$ E_{\text{photon}} = hf
Each photon of light with frequency f carries Planck's quantum of energy, hf. The frequency of the light therefore fixes the size of every lump, while the brightness of the light fixes only how many lumps arrive each second.
---
$ hf = \phi + K
Energy is conserved in each individual event, in which one photon is absorbed by one electron. The photon's energy is spent partly on the cost of escape, the work function φ, and the remainder appears as the electron's kinetic energy K as it leaves the metal.
---
$ K_{\max} = hf - \phi
Electrons from just below the surface pay only the minimum cost φ, but electrons from deeper inside lose extra energy in collisions on their way out and emerge slower. So the equation, rearranged to give the kinetic energy, describes the fastest electrons. The energy of each electron depends on the frequency of the light and on the metal, and on nothing else.
---
$ eV_s = hf - \phi
The stopping voltage measures exactly this maximum kinetic energy, since the fastest electrons are the last to be turned back. Replacing K_max with eV_s turns the equation into a prediction about something that can be measured with a voltmeter.
---
$ V_s = \frac{h}{e}\,f - \frac{\phi}{e}
Divide both sides by the electron's charge e. This says that if the stopping voltage is plotted against the frequency of the light, the graph is a straight line. Its slope is h/e, made only of constants of nature, so it must be the same for every metal; only the point where the line crosses the axis, set by the work function, differs from one metal to another.
:::

Each puzzling feature now has a reason. The threshold exists because if hf is less than φ, no single photon has enough energy to free an electron, and since photons are absorbed one at a time, having more of them does not help. The threshold frequency is f₀ = φ/h. The electrons' energy does not depend on brightness, because brighter light means more photons, each of the same energy hf; more photons free more electrons, but each electron still gets one photon's worth. And there is no delay, because the energy does not need to accumulate: a single photon, arriving at a single spot, does the whole job at once. Most importantly, the equation made a new and specific prediction that no one had tested: a straight line, with the same slope for every metal.

A convenient number makes calculations easy. The product hc is about 1240 eV·nm, so a photon of wavelength λ, measured in nanometres, has energy 1240/λ electronvolts. Sodium has a work function of about 2.3 eV, so its threshold wavelength is 1240/2.3 ≈ 540 nm, in the green. Red light at 650 nm carries 1240/650 ≈ 1.9 eV per photon, not enough, so red light of any brightness frees nothing from sodium. Ultraviolet light at 250 nm carries about 5.0 eV per photon, and frees electrons with up to 5.0 − 2.3 = 2.7 eV of kinetic energy, which an opposing voltage of 2.7 V would stop.

::: note
Light delivers its energy in lumps of size hf, and each lump is absorbed whole by a single electron. The frequency of the light sets the size of each lump, and so the energy of each freed electron; the brightness only sets how many lumps arrive, and so how many electrons are freed.
:::

::: history 1905–1923 · Bern and Chicago
Einstein's paper reached the journal in March 1905, three months before his first paper on relativity, and he called it "a heuristic point of view", meaning a useful way of looking at things rather than a finished theory. In a letter to a friend he called it "very revolutionary", a description he gave to none of his other papers of that year. Almost no one agreed with it. The wave nature of light had been settled by interference experiments for a century, and Maxwell's theory described light as a continuous field with enormous success. Even when Planck and three colleagues recommended Einstein for the Prussian Academy in 1913, they asked that he not be blamed too much for having sometimes missed the target in his speculations, "as for example in his hypothesis of light quanta". Robert Millikan at the University of Chicago set out to test the photoelectric equation, expecting to show it wrong. Earlier measurements had been spoiled by thin films of oxide and contamination on the metal surfaces, which changed the work function. So Millikan built an apparatus in which a knife, worked by a magnet from outside the vacuum, shaved fresh clean surfaces of sodium, potassium and lithium just before each measurement. After about ten years of work he published in 1916: the straight line was exactly as Einstein had predicted, with the same slope for every metal, and from that slope he measured h to within about half a per cent of the value from blackbody radiation. Even so, Millikan called the light-quantum idea "bold, not to say reckless", and held that the equation was right but its explanation unproven. Einstein's Nobel Prize, the 1921 prize awarded in 1922, cited "his discovery of the law of the photoelectric effect". Millikan received his in 1923, partly for the same experiments.
:::

## How strong is the evidence for photons?

The photoelectric effect persuaded physicists that light exchanges energy with matter in lumps of hf. Strictly, though, it does not by itself prove that light travels as lumps. A theory in which light is a classical wave but the atoms absorbing it can only take energy in fixed amounts can reproduce most of the observations, and some physicists held that view for years. Stronger evidence came in 1923, when Arthur Compton scattered X-rays off electrons and found that the X-rays lost energy exactly as if each one were a particle carrying momentum h/λ, bouncing off an electron like one billiard ball off another. Later still, detectors sensitive enough to register single photons showed light arriving one indivisible click at a time. Yet the same light still produces interference fringes. Light, it turned out, is neither a classical wave nor a classical particle, and making sense of that took the full theory of quantum mechanics.

::: model bohr
Each jump between energy levels in hydrogen emits a photon of one definite energy and colour. Which of these photons could free an electron from sodium, whose work function is about 2.3 eV?
:::

::: quiz
? Light of a frequency above the threshold shines on a metal, and its intensity is doubled. What happens?
- The maximum kinetic energy of the electrons doubles
+ Twice as many electrons are emitted, with the same maximum kinetic energy
- Nothing changes
= Doubling the intensity doubles the number of photons, so twice as many electrons are freed. Each photon still carries hf, so the maximum kinetic energy hf − φ is unchanged.
? Light of wavelength 400 nm, with photons of 3.1 eV, falls on a metal whose work function is 2.1 eV. What is the stopping voltage?
- 3.1 V
- 2.1 V
+ 1.0 V
= The fastest electrons leave with K_max = 3.1 eV − 2.1 eV = 1.0 eV. An opposing voltage of 1.0 V takes exactly that much energy from each electron, so it stops them.
? Millikan plotted stopping voltage against frequency for several different metals. What did the graphs have in common?
+ The same slope, h/e
- The same intercept
- The same threshold frequency
= In V_s = (h/e)f − φ/e, the slope depends only on constants of nature, so it is the same for every metal. The work function differs between metals, which shifts each line sideways and changes its threshold.
:::
