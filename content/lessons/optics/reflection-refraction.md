Put a straw in a glass of water and it looks broken where it enters the surface. Look down into a clear swimming pool and the bottom seems closer than it really is. Both illusions have the same cause, and it is a simple one: light travels more slowly in water than in air, and a wave that changes speed while crossing a boundary at an angle also changes direction. This lesson follows that idea from the plain rule for mirrors to the law of refraction, and then to two of its consequences, light trapped inside glass and white light split into colours. Almost every optical instrument, from a pair of spectacles to a fibre-optic cable, depends on what follows here.

## What happens when light bounces?

Start with the easier half of the story. When light meets a smooth surface such as polished metal or still water, some or all of it comes back. To describe where it goes, we need a way to measure angles that does not depend on how the surface happens to be tilted. The convention is to use the **normal**, an imaginary line standing perpendicular to the surface at the point where the light strikes. The angle between the incoming ray and the normal is the angle of incidence; the angle between the outgoing ray and the normal is the angle of reflection.

::: equation The law of reflection
\theta_r = \theta_i
:::

The two angles are equal, and the reflected ray lies in the same flat plane as the incoming ray and the normal, on the opposite side of the normal. This is what you would expect from symmetry: nothing about a flat, smooth surface prefers one side of the normal over the other, so light leaving at a different angle from the one it arrived at would need some reason to favour a direction, and there is none.

The same law explains why a mirror and a sheet of white paper look so different, even though both reflect most of the light that falls on them. Under a microscope, paper is a tangle of fibres, each with tiny faces pointing in every direction. Each small face obeys the law of reflection exactly, but because the faces are tilted at random, the reflected rays leave in all directions. That is called diffuse reflection, and it is why you see the paper from anywhere in the room but never see your face in it. A mirror is smooth on a scale smaller than the wavelength of light, so all its tiny faces share one normal, and an orderly bundle of rays arriving from your face leaves as an orderly bundle, carrying the image with it.

## Why does light slow down in glass?

Reflection does not require light to enter the material. Refraction, the bending of light as it passes from one material into another, does, and it begins with a change of speed.

Light is an electromagnetic wave: an electric field and a magnetic field oscillating together and travelling through space. When it enters glass or water, its electric field pushes on the electrons bound in the atoms and sets them oscillating at the same frequency. Oscillating electrons radiate waves of their own. Those new waves overlap with the original one, and the combined wave that results moves forward more slowly than light in empty space. Nothing in the light is literally being held up between atoms; the slower wave is the sum of the original and all the re-radiated ones.

The **refractive index** n of a material says how much slower:

::: equation Refractive index
n = \frac{c}{v}
:::

Here c is the speed of light in a vacuum, about 3.00 × 10⁸ m/s, and v is its speed in the material. Air has n ≈ 1.0003, so close to 1 that for most purposes air and vacuum are the same. Water has n ≈ 1.33, so light in water moves at about three-quarters of its vacuum speed. Ordinary glass has n ≈ 1.5, and diamond n ≈ 2.42, which means light crawls through diamond at well under half its usual speed.

One property of the wave does not change at the boundary: its frequency, the number of crests passing a point each second. The reason is that every crest arriving at the surface from the air produces exactly one crest leaving into the water. If more crests left than arrived, or fewer, crests would pile up or vanish at the surface, which they cannot do in a steady beam. With the frequency fixed and the speed reduced by the factor n, the wavelength (speed divided by frequency) must shrink by the same factor: λ_medium = λ/n. Red light of wavelength 650 nm in air has a wavelength of about 490 nm in water, though it is still red, because colour is set by frequency.

## Why does a change of speed bend the ray?

Here is the puzzle. Slowing down is a change in speed, yet what we see is a change in direction. How does one become the other?

Think of the light not as a ray but as a wavefront, a line joining neighbouring crests, moving forward perpendicular to itself. When a straight wavefront arrives at the boundary at an angle, one end of it reaches the water first. That end immediately starts moving more slowly, while the rest of the wavefront is still in the air and still moving fast. The fast part gains on the slow part, and the whole front swings around until it is entirely in the water, where it moves on in a new direction. A line of soldiers marching across a field at an angle into a muddy patch would do the same: the first to reach the mud slow down, the others keep pace for a moment longer, and the line pivots. That picture breaks down only in that light has no soldiers to keep in step; it is the geometry of the front, not any communication along it, that does the turning.

The amount of pivoting can be worked out exactly with two right-angled triangles.

::: steps Snell's law from a wavefront crossing a boundary
$ \sin\theta_1 = \frac{v_1 t}{D}
Let one edge of a wavefront touch the boundary at a point A, while the other edge is still a distance v₁t from the boundary and will reach it at a point B after a time t. Call the distance from A to B along the surface D. The wavefront, the surface and that short path form a right triangle with hypotenuse D, and the angle at A equals θ₁, the angle between the incoming ray and the normal, because the wavefront is perpendicular to the ray just as the surface is perpendicular to the normal.
---
$ \sin\theta_2 = \frac{v_2 t}{D}
During that same time t, the light that entered at A has travelled a distance v₂t into the second medium at its slower speed. The new wavefront runs from B to the point that this light has reached, and it is tangent to the circle of radius v₂t around A. That gives a second right triangle with the same hypotenuse D, and its angle at B is θ₂, the angle between the refracted ray and the normal.
---
$ \frac{\sin\theta_1}{\sin\theta_2} = \frac{v_1}{v_2}
Divide the first equation by the second. The distance D and the time t appear in both and cancel, which is what we want, since neither has anything to do with the materials. What remains says that the ratio of the sines of the two angles is fixed by the ratio of the two speeds alone.
---
$ \frac{\sin\theta_1}{\sin\theta_2} = \frac{c/n_1}{c/n_2} = \frac{n_2}{n_1}
Each speed can be written using the definition of refractive index, v = c/n. The vacuum speed c cancels, and the ratio of speeds becomes the ratio of indices, turned upside down: a larger index means a slower speed.
---
$ n_1 \sin\theta_1 = n_2 \sin\theta_2
Multiply both sides by n₁ sin θ₂ to bring the quantities for each medium together on one side. This is Snell's law. Each side depends only on one medium and one angle, so the product n sin θ is the quantity that stays the same as light crosses the boundary.
:::

Now read the law. Going from air into water, n increases, so sin θ must decrease to keep the product the same: the ray bends toward the normal. Coming out of water into air, the opposite happens and the ray bends away from the normal. A ray that arrives straight along the normal, with θ₁ = 0, goes straight on, since sin 0 = 0 on both sides.

That explains the broken straw. Light from the submerged tip travels up to the surface and bends away from the normal as it leaves the water, heading toward your eye at a shallower angle than it started. Your brain assumes light travels in straight lines and traces the ray straight back, to a point higher than the real tip. The pool is the same effect seen from above: looking straight down, the bottom appears at about 1/1.33 of its true depth, or roughly three-quarters. A pool 2 m deep looks about 1.5 m deep, which is why people misjudge where to put their feet.

## What happens when light cannot get out?

Snell's law has a surprise in it. Go from a material with a high index into one with a lower index, say from water into air. The ray bends away from the normal, so the refracted angle is always larger than the incident angle. As you tilt the incoming ray further from the normal, the refracted ray swings toward the surface, and at some incident angle it reaches 90°, skimming along the boundary. Tilt the incoming ray further still and Snell's law asks for sin θ₂ greater than 1, which no angle has. There is no refracted ray at all. All the light is reflected back into the water. This is **total internal reflection**, and the incident angle at which it begins is the **critical angle** θ_c.

::: equation Critical angle, going from index n₁ to a smaller index n₂
\sin\theta_c = \frac{n_2}{n_1}
:::

This comes straight from Snell's law with θ₂ set to 90°, where sin θ₂ = 1. For water into air, sin θ_c = 1/1.33, so θ_c ≈ 48.8°. A swimmer looking up from under water sees the whole sky squeezed into a bright circle overhead, about 97° across, and beyond its edge a silvery mirror reflecting the pool. For glass into air the critical angle is about 42°, and for diamond only about 24°.

The reflection is called total for a reason: an ordinary mirror absorbs a few per cent of the light on every bounce, but total internal reflection loses essentially nothing. That is what makes optical fibres possible. A fibre is a glass thread finer than a hair, with a core of slightly higher index than the glass around it. Light sent in at a shallow angle meets the boundary beyond the critical angle every time, and bounces along the thread through every bend, for tens of kilometres before it needs boosting. Most of the world's telephone and internet traffic travels this way. Diamond's small critical angle explains its sparkle: light that enters a well-cut stone strikes the back faces beyond 24° and is reflected around inside, leaving only through the top, so the stone throws back far more light than glass of the same shape.

## Why does a prism make colours?

The refractive index is not quite a single number. It depends slightly on the wavelength of the light, because the electrons in glass respond a little differently to fields oscillating at different frequencies. In ordinary glass, violet light meets a higher index than red, so violet slows more and bends more. This spreading of colours by refraction is called **dispersion**. A beam of white light, which contains all the visible wavelengths mixed together, enters a prism and leaves as a fan, red bent least and violet most. Newton showed in the 1660s that the colours were in the white light all along: a second prism, turned the other way, merges them back into white.

A rainbow is dispersion in raindrops. Sunlight enters each drop, reflects once off its back surface, and leaves, with each colour bent by a slightly different amount. Red light leaves at about 42° from the direction opposite the Sun and violet at about 40°, which is why the bow is a circle of that size around the shadow of your own head, red on the outside.

::: note
Refraction happens because light changes speed between materials, and a wavefront crossing a boundary at an angle pivots as one end slows before the other. The geometry gives Snell's law, n₁ sin θ₁ = n₂ sin θ₂, and from it follow the bending toward the normal, total internal reflection beyond the critical angle, and the splitting of colours by a prism.
:::

::: history 984–1850 · Alexandria, Baghdad, Paris
Ptolemy, working in Alexandria around 150 AD, measured how much light bends entering water at various angles, but he fitted his numbers with a rule that is close only for small angles. The correct law first appears in a treatise on burning mirrors and lenses written in Baghdad in 984 by Ibn Sahl, who used it to design lenses that focus sunlight, though his work was forgotten for centuries. In Europe the same law was found again by Thomas Harriot in 1602 and by Willebrord Snellius in 1621; neither published it, but Snellius's name stuck. René Descartes printed it in 1637 with an explanation that required light to travel faster in water than in air. Pierre de Fermat disliked that and in 1662 derived the same law from a different idea, that light takes the path of least time between two points, which works only if light is slower in water. The two explanations predicted opposite speeds, and nobody could measure the speed of light in water to decide. That waited until 1850, when Léon Foucault, using a rapidly spinning mirror, timed light through a tube of water and found it slower, just as Fermat's principle and the wave theory required.
:::

::: quiz
? A ray of light passes from air into glass at an angle to the normal. What happens to it?
- It speeds up and bends away from the normal
+ It slows down and bends toward the normal
- It keeps its direction but changes colour
= Glass has the larger refractive index, so light slows there. To keep n sin θ the same on both sides, sin θ must be smaller in the glass, so the ray bends toward the normal.
? When light passes from air into water, which property stays the same?
- Its speed
- Its wavelength
+ Its frequency
= Each crest arriving at the surface produces exactly one crest leaving it, so the frequency cannot change. The speed drops by the factor n, and so the wavelength shrinks by the same factor.
? Can total internal reflection happen for light going from air into water?
- Yes, for any angle above 48.8°
+ No, it needs light to go from a higher index into a lower one
- Yes, at every angle
= Entering a material with a higher index, the ray bends toward the normal, so Snell's law always has a solution. Only when light heads into a lower index can the refracted angle reach 90° and the light be trapped.
:::
