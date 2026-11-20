Put a straw in a glass of water and it looks broken where it enters the surface. Look down into a clear swimming pool and the bottom seems closer than it really is. Both illusions have the same cause, and it is a simple one: light travels more slowly in water than in air, and a wave that changes speed while crossing a boundary at an angle also changes direction. This lesson follows that idea from the plain rule for mirrors to the law of refraction, and then to two of its consequences, light trapped inside glass and white light split into colours. Almost every optical instrument, from a pair of spectacles to a fibre-optic cable, depends on what follows here.

## What was there to explain?

People have known since antiquity that light does two things at a surface. Some of it bounces back, which is how mirrors and still ponds work, and some of it goes in and changes direction, which is why an oar dipped in a lake looks bent. The first behaviour was understood early: Hero of Alexandria pointed out in the first century that a ray bouncing off a mirror takes the shortest possible path between source and eye, and that rule alone fixes where it goes. The second resisted for well over a thousand years. Ptolemy measured the bending carefully and tabulated it, but no simple rule fitted his angles. Double the angle at which light strikes water and the bending does not double, nor does it follow any obvious pattern in the angles themselves.

Behind the missing rule lay a deeper question: why should light bend at all? A beam crossing from air into water passes through no wall and meets no obstacle at the surface, yet it leaves in a new direction. Explaining that needs an idea about what light is and how fast it goes, and the right idea, that light is a wave which moves more slowly in water, was not settled until the nineteenth century. This lesson starts with the easy half and then builds the hard half from the speed of light in matter.

So, the easy half first. When light meets a smooth surface such as polished metal or still water, some or all of it comes back. To describe where it goes, we need a way to measure angles that does not depend on how the surface happens to be tilted. The convention is to use the **normal**, an imaginary line standing perpendicular to the surface at the point where the light strikes. The angle between the incoming ray and the normal is the angle of incidence; the angle between the outgoing ray and the normal is the angle of reflection.

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

The reflection is called total for a reason: an ordinary mirror absorbs a few per cent of the light on every bounce, but total internal reflection loses essentially nothing. That is what makes optical fibres possible. A fibre is a glass thread finer than a hair, with a core of slightly higher index than the glass around it, called the cladding. Light sent in at a shallow angle meets the boundary beyond the critical angle every time, and bounces along the thread through every bend, for tens of kilometres before it needs boosting. The cladding matters more than it might seem: a bare glass thread would also trap light, but every fingerprint, speck of dust or touching neighbour on its surface would change the index outside and let light leak away. With the guiding boundary buried inside the glass, the outer surface can be handled freely. Most of the world's telephone and internet traffic travels this way. Diamond's small critical angle explains its sparkle: light that enters a well-cut stone strikes the back faces beyond 24° and is reflected around inside, leaving only through the top, so the stone throws back far more light than glass of the same shape.

## Why does a prism make colours?

The refractive index is not quite a single number. It depends slightly on the wavelength of the light, because the electrons in glass respond a little differently to fields oscillating at different frequencies. In ordinary glass, violet light meets a higher index than red, so violet slows more and bends more. This spreading of colours by refraction is called **dispersion**. A beam of white light, which contains all the visible wavelengths mixed together, enters a prism and leaves as a fan, red bent least and violet most. Newton showed in the 1660s that the colours were in the white light all along: a second prism, turned the other way, merges them back into white.

A rainbow is dispersion in raindrops. Sunlight enters each drop, reflects once off its back surface, and leaves, with each colour bent by a slightly different amount. Red light leaves at about 42° from the direction opposite the Sun and violet at about 40°, which is why the bow is a circle of that size around the shadow of your own head, red on the outside.

The spread is small. In a common optical glass the index is about 1.514 for red light and about 1.522 for blue, a difference of half a per cent. That is enough to separate the colours over a few metres, and enough to blur the image in a simple lens, as the next lesson explains, but it is far too small to notice when you look through a window.

## Working through some numbers

Snell's law is easy to state, and the best way to get a feel for it is to put numbers into it. Three cases show three different sides of it: an ordinary bend, the design of a fibre, and the size of dispersion.

First, a ray of sunlight strikes a calm lake at 45° from the vertical. Air has n₁ ≈ 1.00 and water n₂ = 1.33, so sin θ₂ = sin 45°/1.33 = 0.707/1.33 ≈ 0.532, which gives θ₂ ≈ 32°. The ray has turned 13° toward the vertical. Had the lake been made of glass with n = 1.5, the angle would have been about 28°: the higher the index, the more sharply the light is pulled toward the normal. Notice that the bending is not proportional to the index in any simple way. It is the sines of the angles, not the angles, that keep a fixed ratio, which is exactly what defeated Ptolemy, whose measurements were good but whose guess was a rule in the angles themselves.

::: equation A ray entering water at 45°
\sin\theta_2 = \frac{n_1}{n_2}\sin\theta_1 = \frac{\sin 45^\circ}{1.33} \approx 0.532
:::

Second, a telecommunications fibre has a core of index 1.48 and cladding of index 1.46. The critical angle at the core boundary is given by sin θ_c = 1.46/1.48 ≈ 0.986, so θ_c ≈ 80.6°. That is a very steep angle, measured from the normal to the wall, which means a ray is trapped only if it travels within about 9.4° of the fibre's axis. Light entering from air at the flat end bends toward the axis as it goes in, and working back through Snell's law, the widest cone of light the fibre can accept from outside has a half-angle of about 14°. Light arriving at a steeper slant leaks into the cladding and is lost. Even the trapped light pays a price for its zigzag: a ray at the edge of the cone travels about 1.4 per cent farther than one running straight down the axis, so over a kilometre the two arrive about 70 nanoseconds apart. Short pulses smear out, which limits how fast data can be sent. Long-distance fibres avoid this by making the core so narrow, about 9 micrometres, that only one way of travelling through it survives. At that size the ray picture no longer describes the light well, a hint of the limits discussed below.

Third, the time light spends in that fibre. Its speed in the core is c/1.48 ≈ 2.03 × 10⁸ m/s, so crossing 100 km takes about 0.49 milliseconds, against 0.33 milliseconds through vacuum. For traders sending orders between financial exchanges, that difference is worth money, and some routes have been replaced by microwave links through air, where n is almost exactly 1.

Finally, dispersion. Send white light into that common optical glass at 60°. Red light, meeting n = 1.514, refracts to about 34.9°; blue light, meeting n = 1.522, to about 34.7°. A fifth of a degree separates them. That is small, but a prism bends the light at two faces, both in the same sense, and a screen a few metres away turns a fifth of a degree into a spread of colour about a centimetre wide.

## Going deeper: the path of least time

The wavefront argument derives Snell's law from the picture of light as a wave. There is a second derivation that looks completely different and gives the same answer, and comparing the two is illuminating. It is the one Fermat used, and it rests on a single principle: of all the paths light could take between two points, it takes the one that needs the least time.

For reflection this reproduces Hero's rule, since in a single medium least time means shortest distance. For refraction it is more interesting. Imagine a lifeguard on the beach who must reach a swimmer out in the water. She runs faster than she swims, so the straight line is not the quickest route: it pays to run a little farther along the beach and swim a shorter distance. But running all the way to the point opposite the swimmer wastes too much time on sand. The best route is a compromise, and finding it is a problem in the lesson Rates of Change: write the total time as a function of where she enters the water, and find where its derivative is zero.

::: steps Snell's law from least time
$ T(x) = \frac{\sqrt{a^2 + x^2}}{v_1} + \frac{\sqrt{b^2 + (d - x)^2}}{v_2}
Put the source a height a above the boundary and the destination a depth b below it, a horizontal distance d further along. Let the light cross the boundary at a point a horizontal distance x from the source. Each leg is the hypotenuse of a right triangle, found by Pythagoras, and dividing each length by its speed gives the time spent on it. The total time T depends on the one free choice, x.
---
$ \frac{\mathrm{d}T}{\mathrm{d}x} = \frac{x}{v_1\sqrt{a^2 + x^2}} \\ \qquad - \frac{d - x}{v_2\sqrt{b^2 + (d - x)^2}} = 0
At the quickest path, moving the crossing point a little either way does not change the time to first order, so the derivative of T with respect to x is zero. Differentiating each square root gives the two terms; the second has a minus sign because moving the crossing point forward shortens the second leg.
---
$ \frac{\sin\theta_1}{v_1} = \frac{\sin\theta_2}{v_2}
Look at the triangles again. The fraction x/√(a² + x²) is the opposite side over the hypotenuse for the angle between the first leg and the normal, which is sin θ₁. In the same way, the second fraction is sin θ₂. Setting the derivative to zero therefore says that the sine of each angle divided by the speed on that side is the same.
---
$ n_1 \sin\theta_1 = n_2 \sin\theta_2
Multiply both sides by c, and use c/v = n on each side. This is Snell's law again, reached without mentioning wavefronts at all.
:::

That two such different arguments agree is no accident. In the wave picture, the paths close to the quickest one take almost the same time, so the waves along them arrive in step and reinforce each other, while waves along other paths arrive out of step and cancel. Least time is what the wave picture looks like when the wavelength is small. The same idea, that nature picks the path along which neighbouring paths agree, returns in quantum mechanics, where it governs the motion of electrons as well as light.

There is a second question the ray picture leaves open: how much light reflects and how much goes through? Maxwell's equations, applied at the boundary, answer it. For light arriving straight on, the fraction reflected is ((n₁ − n₂)/(n₁ + n₂))². For air and glass that is (0.5/2.5)² = 0.04, so a window reflects about 4 per cent of the light at each surface. In daylight that is lost against the bright scene beyond, but at night, with a dark garden outside and a lit room inside, the 4 per cent is all you see, and the window becomes a mirror. For water the figure is about 2 per cent, and for diamond about 17 per cent, which adds to its brilliance. At grazing angles the reflected fraction climbs toward 100 per cent for every material, which is why a lake seen at a low angle looks like a sheet of sky.

## Common misconceptions

A popular explanation of why light slows in glass says that each photon is absorbed by an atom, held for a moment, and then emitted again, so the delays add up. That picture fails. An atom that absorbs and re-emits light sends it off in a random direction, so a beam treated this way would come out as a diffuse glow, not a sharp image. And real absorption happens only at particular frequencies, while glass slows all visible light by nearly the same amount. The slowing is the forward combination of the original wave and the waves from electrons driven in step with it, as described above. Between atoms, the field still moves at c.

A second common belief is that a mirror swaps left and right. Raise your right hand in front of a mirror and the image raises the hand on the same side of the picture; nothing has moved sideways. What the mirror reverses is front and back, the direction along the normal. We describe the image as left-right swapped only because we imagine walking round to stand where it is, which is a rotation that the mirror never made.

It is also easy to think that light curves gradually as it enters water. In a uniform material it travels in a straight line; all the turning happens at the boundary. Gradual curving does happen, but only where the index itself changes gradually, as in air over a hot road. Finally, total internal reflection does not need a silvered surface. It happens at a bare boundary between two transparent materials, and its completeness comes from Snell's law having no solution, not from any coating.

## Where the ray picture stops working

Everything in this lesson treats light as rays that travel in straight lines and bend at surfaces. That works when the objects involved are much larger than the wavelength of light, about half a micrometre. Shine light through a gap only a few wavelengths wide and it spreads out into the shadow, which no ray can do. Rays then have to give way to the full wave picture, which is the subject of the lessons Diffraction and The Double Slit.

Even total internal reflection is not quite what the ray picture says. The wave does not stop dead at the boundary: a weak field, called an evanescent wave, reaches about a wavelength into the outer material and fades rapidly. It carries no energy away, so the reflection is still total. But bring a second block of glass within a fraction of a micrometre and some light jumps the gap and carries on, an effect called frustrated total internal reflection. Some fingerprint scanners use it: ridges of skin touching the glass let light escape where valleys of air do not. It is the optical cousin of quantum tunnelling, met later in Tunneling.

Snell's law also assumes the material is the same in every direction. Some crystals, such as calcite, are not: they have two refractive indices depending on how the light's electric field is oriented, so a single ray entering them splits in two and you see a double image through them. And the index is not always greater than 1. For X-rays, whose frequency is far above the natural frequencies of the electrons in glass, n is very slightly less than 1, so X-rays striking glass at a grazing angle from air are totally reflected. X-ray telescopes are built from nested mirrors at grazing angles for exactly this reason.

## Where it shows up

Refraction is the working principle of every lens, from the cornea of your eye to a camera, and the next lesson, Lenses and Images, builds directly on Snell's law. Total internal reflection carries the internet through undersea fibres, folds the light path inside binoculars, where glass prisms replace mirrors because they lose nothing at each bounce, and lets the rain sensor on a car windscreen detect drops that spoil the reflection.

In nature, refraction by air whose index changes with temperature makes mirages. Near a hot road, the lowest air is warmest and has a slightly lower index, so light from the sky heading down bends gradually upward and reaches your eye from below, where it looks like a pool of water on the road. The whole atmosphere bends starlight in the same way, lifting objects near the horizon by about half a degree, roughly the width of the Sun. When you watch the Sun touch the horizon at sunset, it has already, geometrically, gone below it.

The ideas here also look back and forward. That light slows in matter follows from it being an electromagnetic wave, as in Electromagnetic Waves and Calculating c, and the dependence of the index on wavelength is what lets a prism lay out The Spectrum. Bending Light, in general relativity, describes how gravity deflects starlight, a different cause that produces lens-like effects on a cosmic scale.

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
? A ray inside a block of glass (n = 1.5) meets a boundary with water (n = 1.33) at 70° from the normal. What happens?
- It passes into the water, bending toward the normal
- It passes into the water, bending away from the normal
+ It is totally reflected back into the glass
= The critical angle from glass into water is given by sin θ_c = 1.33/1.5, so θ_c ≈ 62°. The ray arrives at 70°, beyond that, so Snell's law has no solution and all the light is reflected.
? Why does a window look like a mirror at night but not during the day?
- Glass reflects more light when it is cold
+ It always reflects about 4 per cent of the light, which only dominates when little light comes through from outside
- At night the glass becomes totally internally reflecting
= Each air-glass surface reflects ((n − 1)/(n + 1))² ≈ 4 per cent whatever the time of day. By day that weak reflection is swamped by bright light from outside; at night the dark scene beyond lets it show.
:::
