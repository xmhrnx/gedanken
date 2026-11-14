In September 1905, three months after his first relativity paper, Einstein sent a three-page sequel to the same journal. Its title asked a question: "Does the inertia of a body depend upon its energy content?" His answer was yes. A body that gives off energy E loses mass E/c². Mass is a form of energy, and c² is the exchange rate.

## Why Newton's momentum must change

In Newton's physics, momentum is p = mv, and it is conserved in collisions. Apply the postulates to a collision viewed from two frames and p = mv turns out to be conserved in one frame but not the other. The quantity that is conserved in every frame is the **relativistic momentum**:

::: equation Relativistic momentum
p = \gamma m v = \frac{mv}{\sqrt{1 - v^2/c^2}}
:::

Here m is the ordinary mass, the same in all frames. At everyday speeds γ ≈ 1 and this is Newton's mv. Near c, momentum grows without limit, which is why no amount of push brings a massive object to c.

## From momentum to energy

Kinetic energy is the work done to bring an object up to speed: force times distance, added up. With the new momentum, the result is no longer ½mv².

::: steps Kinetic energy with relativistic momentum
$ \mathrm{d}K = F\,\mathrm{d}x = \frac{\mathrm{d}p}{\mathrm{d}t}\,\mathrm{d}x = v\,\mathrm{d}p
The work done by a force F over a small distance dx raises the kinetic energy by F dx. Force is the rate of change of momentum, and dx/dt is the velocity v.
---
$ \frac{\mathrm{d}(\gamma v)}{\mathrm{d}v} = \gamma + v\,\frac{\mathrm{d}\gamma}{\mathrm{d}v} = \gamma + v \cdot \frac{\gamma^3 v}{c^2} = \gamma^3
With p = mγv, we need the derivative of γv. Differentiating γ = (1 − v²/c²)^(−1/2) gives γ³v/c². Then factor out γ³: γ³(1 − v²/c²) + γ³v²/c² = γ³, using 1/γ² = 1 − v²/c².
---
$ \mathrm{d}K = m v \gamma^3\,\mathrm{d}v
So dp = mγ³ dv, and dK = v dp.
---
$ \mathrm{d}\gamma = \frac{\gamma^3 v}{c^2}\,\mathrm{d}v \quad \Rightarrow \quad \mathrm{d}K = mc^2\,\mathrm{d}\gamma
The derivative of γ from the second step shows that mvγ³ dv is exactly mc² times the change in γ.
---
$ K = mc^2(\gamma - 1)
Add up from rest, where γ = 1 and K = 0, to speed v. The kinetic energy is mc² times the increase in γ.
---
$ \gamma mc^2 = K + mc^2
Rearrange. The quantity γmc² behaves as the total energy E, and it is made of two parts: the kinetic energy, and a part mc² that is there even when the object is at rest.
---
$ E_0 = mc^2
At rest, γ = 1 and the energy is mc². This is the **rest energy**.
:::

A check: for small v, γ ≈ 1 + v²/(2c²), so K ≈ mc² × v²/(2c²) = ½mv². Newton's kinetic energy is the low-speed approximation of the relativistic one.

This derivation shows that γmc² is the natural energy and that mc² sits inside it, but adding a constant to an energy usually changes nothing measurable. What makes mc² physical is the claim that it can be **converted**: that a body losing energy loses mass. That was the point of Einstein's argument.

::: history September 1905 · Bern
Einstein imagined a body at rest emitting two equal pulses of light in opposite directions, total energy L, so that it stays at rest. He then viewed the same emission from a frame moving past at speed v. Using his earlier result for how the energy of light changes between frames, the pulses carry total energy γL in the moving frame, more than L. Energy conservation in both frames then requires the body's kinetic energy in the moving frame to drop by L(γ − 1) ≈ ½(L/c²)v². Its speed has not changed, so its mass must have: it fell by L/c². He concluded that "the mass of a body is a measure of its energy content", and suggested radium salts might test it. His 1905 argument kept only the leading terms in v/c, and its rigour was debated for decades; the result was not.
:::

## The numbers

::: equation Energy in one gram of matter
E = (10^{-3}\,\text{kg})(3.00 \times 10^8\,\text{m/s})^2 = 9 \times 10^{13}\,\text{J}
:::

That is about 21 kilotons of TNT, more than the Hiroshima bomb, which converted roughly 0.7 g of mass to energy. The Sun converts about 4 million tonnes of mass into energy every second.

In chemistry the mass changes are real but tiny. Burning a kilogram of coal releases about 3 × 10⁷ J and loses about 0.3 micrograms, one part in three billion: no balance can detect it. In nuclear reactions the fraction is about a million times larger, up to about 0.1% in fission and 0.7% in hydrogen fusion, which is why nuclear energy dwarfs chemical energy.

E = mc² does not say mass "turns into" energy as if they were different substances. Mass is energy: the mass of a nucleus or an atom includes the energy of the fields and motion inside it. When energy leaves, mass goes with it. Older books speak of "relativistic mass" growing with speed; modern usage keeps m fixed and lets the energy γmc² grow.

## Confirmation

In 1932 John Cockcroft and Ernest Walton split lithium nuclei with protons, making two helium nuclei. The kinetic energy released, about 17 MeV, matched the measured loss of mass, the first direct check. In 2005 a comparison of nuclear masses with gamma-ray energies confirmed E = mc² to about four parts in ten million.

::: note
A body's energy content and its mass are the same thing measured in different units: E = mc². Any energy a body gains or loses changes its mass by E/c².
:::

::: quiz
? A hot cup of coffee cools down. What happens to its mass?
- Nothing; only nuclear reactions change mass
+ It decreases, by the heat lost divided by c²
- It increases, because the molecules slow down
= Any energy lost carries mass E/c² with it; for coffee the change is far too small to measure.
? Why can no massive object be accelerated to the speed of light?
- Air resistance becomes too large
+ Its energy γmc² and momentum γmv grow without limit as v approaches c
- Its mass becomes zero
= γ → ∞ as v → c, so reaching c would require infinite energy.
? For small speeds, what does K = mc²(γ − 1) reduce to?
- mc²
- mv
+ ½mv²
= With γ ≈ 1 + v²/(2c²), K ≈ ½mv²: Newton's formula is the low-speed limit.
:::
