A straw in a glass of water looks broken at the surface. A pool looks shallower than it is. Both come from one fact: light travels more slowly in water than in air, and a wave that changes speed at an angle changes direction.

## Reflection

When light meets a smooth surface, some of it bounces back. Measure angles from the **normal**, the line perpendicular to the surface at the point where the light hits.

::: equation The law of reflection
\theta_{\text{reflected}} = \theta_{\text{incident}}
:::

The reflected ray also lies in the same plane as the incoming ray and the normal. A rough surface, like paper, obeys the same law at each tiny facet, but the facets point every which way, so the reflected light scatters in all directions. That is the difference between a mirror and a white wall.

## Light slows down in matter

In glass or water, light's electric field sets the electrons oscillating, and their re-radiated waves combine with the original to give a wave that moves more slowly than c. The **refractive index** n measures how much slower:

::: equation Refractive index
n = \frac{c}{v}
:::

Air has n ≈ 1.0003, water 1.33, ordinary glass about 1.5, diamond 2.42. The frequency of the wave does not change at the boundary (each crest arriving makes one crest leaving), so the wavelength shrinks by the same factor: λ_medium = λ/n.

## Snell's law

Picture a straight wavefront arriving at a flat boundary at an angle. One end of the wavefront reaches the boundary first and starts moving slowly in the new medium, while the rest is still moving fast in the old one. The front pivots, like a line of marchers when one end steps into mud.

::: steps Snell's law from wavefronts
$ \sin\theta_1 = \frac{v_1 t}{D}
Let one edge of a wavefront touch the boundary at point A while the other edge still has distance v₁t to go before reaching the boundary at point B, a distance D along the surface. The incoming wavefront makes the angle θ₁ with the surface, equal to the angle the ray makes with the normal, giving a right triangle with hypotenuse D.
---
$ \sin\theta_2 = \frac{v_2 t}{D}
In the same time t, the light that entered at A travels v₂t into the second medium. The new wavefront runs from B to that point, giving a second right triangle with the same hypotenuse D and the refracted angle θ₂.
---
$ \frac{\sin\theta_1}{\sin\theta_2} = \frac{v_1}{v_2}
Divide the two equations. D and t cancel.
---
$ \frac{\sin\theta_1}{\sin\theta_2} = \frac{c/n_1}{c/n_2} = \frac{n_2}{n_1}
Write each speed as c/n.
---
$ n_1 \sin\theta_1 = n_2 \sin\theta_2
Rearrange. This is Snell's law.
:::

Going into a slower medium (larger n), sin θ₂ is smaller than sin θ₁, so the ray bends **toward** the normal. Coming out, it bends away. Light from a straw's tip under water bends away from the normal as it leaves the surface; your eye traces it back in a straight line to a point higher than the real tip. Looking straight down, a pool appears about 1/1.33, or three-quarters, of its true depth.

## Total internal reflection

Going from a high index to a low one, the ray bends away from the normal, and the refracted angle reaches 90° before the incident angle does. Beyond that **critical angle**, Snell's law has no solution: no light gets out, and all of it reflects.

::: equation Critical angle, from n₁ to n₂ with n₁ > n₂
\sin\theta_c = \frac{n_2}{n_1}
:::

For water to air, θ_c = 48.8°; for glass to air, about 42°. Optical fibres trap light this way along thin glass threads, bouncing it off the walls for kilometres with little loss. Diamond's small critical angle, about 24°, makes light bounce around inside before escaping, which is why it sparkles.

## Colours

The index depends slightly on wavelength: in glass, violet light sees a higher n than red, so it bends more. A prism spreads white light into a spectrum this way, called **dispersion**, and raindrops do the same to make a rainbow.

::: note
Refraction happens because light changes speed between materials. Snell's law, n₁ sin θ₁ = n₂ sin θ₂, follows directly from a wavefront crossing the boundary at an angle.
:::

::: history 984–1850 · Baghdad to Paris
Ptolemy measured refraction angles around 150 AD but fitted them with the wrong rule. Ibn Sahl, in Baghdad in 984, gave the correct sine law in a treatise on burning mirrors and lenses. It was found again by Thomas Harriot in 1602 and Willebrord Snellius in 1621, neither of whom published, and first printed by René Descartes in 1637. Descartes' explanation assumed light travels faster in water. Pierre de Fermat derived the same law in 1662 from the principle that light takes the quickest path, which requires light to be slower in water. The question was settled only in 1850, when Léon Foucault measured light directly and found it slower in water, as the wave picture requires.
:::

::: quiz
? A ray of light passes from air into glass at an angle. What happens?
- It bends away from the normal and speeds up
+ It bends toward the normal and slows down
- It keeps its direction but changes colour
= Glass has the larger index, so the speed drops and n₁ sin θ₁ = n₂ sin θ₂ requires a smaller angle in glass.
? When light enters water, which of these stays the same?
- Its speed
- Its wavelength
+ Its frequency
= Each crest arriving at the surface produces one crest leaving it, so the frequency is fixed; speed and wavelength both drop by the factor n.
? Can total internal reflection happen for light going from air into water?
- Yes, at angles above 48.8°
+ No, it requires going from a higher index to a lower one
- Yes, at any angle
= Entering a higher index, the ray bends toward the normal, so a refracted ray always exists.
:::
