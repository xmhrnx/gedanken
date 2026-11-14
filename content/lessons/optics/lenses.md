Every point of a lit object sends light out in all directions. To form an image of it, on a screen, on film, or on the back of your eye, something has to gather a spread of rays leaving one point of the object and bring them back together at one point somewhere else. A lens does this with nothing but refraction: a piece of glass whose surfaces are curved so that each ray bends by just the right amount. This lesson works out where a lens forms its image, why some images can be caught on a screen and others cannot, and how the same rule describes a camera, the human eye and a telescope.

## How can curved glass bring rays together?

The previous lesson showed that light bends toward the normal when it enters glass and away from it when it leaves. On a flat pane of glass the two bends cancel, and a ray comes out parallel to the way it went in, only slightly shifted. A lens works by making the two surfaces non-parallel, and differently so at different heights.

Picture a lens that is thicker in the middle than at its edges. Near the centre, its two surfaces are nearly parallel, so a ray passing through there is hardly deflected. Farther from the centre, the surfaces are tilted toward each other like the faces of a thin prism, and a ray passing through is bent toward the middle. The farther out it passes, the steeper the tilt and the more it bends. With the right curvature, rays arriving parallel to the lens's axis, the line through its centre perpendicular to its faces, all bend just enough to cross at a single point. That is the **focal point**, and its distance from the lens is the **focal length** f. Such a lens is **converging**.

A lens thinner in the middle than at its edges does the opposite. Its outer parts act like prisms pointing the other way, so parallel rays spread apart after passing through, as if they were coming from a point on the near side of the lens. That lens is **diverging**, and by convention its focal length is counted as negative, which lets one equation serve for both kinds.

Two things make a focal length short. More strongly curved surfaces tilt more steeply, so rays bend more. And glass of higher refractive index bends light more at each surface. Opticians describe a lens by its **power**, defined as 1/f with f in metres; the unit is the dioptre. A reading lens marked +2 dioptres has a focal length of 1/2 m, or 50 cm. The advantage of power over focal length is that, for thin lenses placed close together, powers simply add, which is how an optometrist stacks trial lenses in front of your eye.

## How do we find the image without tracing every ray?

In principle you could follow thousands of rays through the lens with Snell's law. In practice, for a **thin lens** (one whose thickness is small compared with its focal length), three rays from any point of an object are easy to draw, and where they meet is the image of that point.

A ray arriving parallel to the axis leaves through the far focal point, because that is what the focal point means. A ray passing through the centre of the lens goes straight on, undeflected, because at the centre the two surfaces are parallel and the lens behaves like a thin flat pane. And a ray that passes through the near focal point on its way in leaves parallel to the axis; that is the first rule run backward, since a light path traced in reverse is also a valid light path.

Any two of these rays are enough to find the image, and the third serves as a check. But a drawing gives only one case at a time. To get a rule that covers every case, we turn the drawing into geometry.

## Where does the image form?

Put an object of height h_o at a distance d_o in front of a thin converging lens of focal length f. Its image forms at a distance d_i behind the lens, upside down, with height h_i. We want a relation between d_o, d_i and f.

::: steps Deriving the thin lens equation
$ \frac{h_i}{h_o} = \frac{d_i}{d_o}
Follow the ray from the top of the object through the centre of the lens. It goes straight, so it forms two triangles that meet at the centre: one with the object as its upright side (height h_o, base d_o) and one with the image (height h_i, base d_i). They have equal angles at the centre, where the ray crosses the axis, and both have a right angle at the axis, so they are similar. In similar triangles corresponding sides are in proportion, which gives this ratio.
---
$ \frac{h_i}{h_o} = \frac{d_i - f}{f}
Now follow the ray from the top of the object that travels parallel to the axis. It meets the lens at height h_o and then heads through the far focal point. Between the lens and the focal point it forms a right triangle with upright side h_o and base f. Beyond the focal point it carries on to the top of the image, forming a second right triangle with upright side h_i and base d_i − f. The two triangles meet at the focal point with equal angles, so they too are similar, and their sides give this second ratio.
---
$ \frac{d_i}{d_o} = \frac{d_i - f}{f} = \frac{d_i}{f} - 1
Both ratios equal the same thing, h_i/h_o, so they equal each other. That removes the heights, which we never needed: where the image forms cannot depend on how tall the object is. Splitting the right-hand fraction into two terms gets it ready for the next step.
---
$ \frac{1}{d_o} = \frac{1}{f} - \frac{1}{d_i}
Divide every term by d_i. This is allowed because d_i is not zero, and it turns each term into a reciprocal of a distance, which is the form in which the three lengths combine most simply.
---
$ \frac{1}{f} = \frac{1}{d_o} + \frac{1}{d_i}
Move 1/d_i to the left-hand side, adding it to both sides. This is the thin lens equation. It links the three distances in a symmetrical way: the object and image distances play the same role, which is another sign that light paths can be run backward.
:::

The first step also gives the **magnification**, the ratio of image height to object height. Taking into account that the image is inverted, it is written m = −d_i/d_o; the minus sign records the inversion.

Try it with a camera. A lens of focal length 50 mm photographs a person standing 2 m (2000 mm) away. The equation gives 1/d_i = 1/50 − 1/2000 = 39/2000 per millimetre, so d_i ≈ 51.3 mm. The image forms 1.3 mm farther from the lens than it would for a very distant scene. That small number is what focusing means: turning the focus ring of a camera moves the lens a millimetre or so to put the sensor where the image is. For a very distant object, d_o is huge and 1/d_o is nearly zero, so d_i = f and the image sits at the focal point. That is why focal length is the number engraved on every camera lens.

## Why can some images be caught on a screen and others not?

When d_i comes out positive, the rays from each point of the object really do converge after the lens and cross at the image point. Put a screen there and a sharp picture appears on it. That is a **real image**, the kind formed on a camera sensor, a cinema screen or your retina. A single converging lens always makes real images upside down, as the drawing of crossing rays shows.

Now bring the object closer than the focal length. Then 1/d_o is larger than 1/f, so 1/d_i = 1/f − 1/d_o is negative, and so is d_i. What does a negative image distance mean? The lens is not strong enough to make the rays converge. They still spread apart after passing through, only less steeply than before. Traced backward, though, they all seem to come from a single point on the same side of the lens as the object, and farther away. Your eye, receiving those rays, cannot tell the difference: it sees an object there. That is a **virtual image**. It is upright and enlarged, and it cannot be caught on a screen, because no light actually passes through the place where it appears to be.

This is how a magnifying glass works. Hold a lens of focal length 10 cm about 8 cm above a page and the equation gives 1/d_i = 1/10 − 1/8 = −1/40 per centimetre, so the image is 40 cm behind the lens on the page's side, and the magnification −d_i/d_o is +5. You see the print upright and five times larger, at a comfortable reading distance.

## How do eyes and telescopes use lenses?

The eye is a camera built from living tissue. Most of its focusing is done not by the lens inside it but by the cornea, the clear curved front surface, where light passes from air into watery tissue and the change of refractive index is largest. The cornea supplies about two-thirds of the eye's total power of roughly 60 dioptres. The lens behind it supplies the rest, and it is flexible: small muscles change its shape to adjust its power, so that both distant mountains and the page in front of you can be focused onto the retina. The image there is real and inverted, and the brain learns to interpret it the right way up.

In short sight, the eye is too powerful for its length, usually because the eyeball is slightly too long. Light from distant objects comes to a focus in front of the retina and has spread out again by the time it arrives. A diverging spectacle lens, with negative power, weakens the combination just enough to move the focus back onto the retina. Long sight is the reverse and is corrected with a converging lens.

A telescope must deal with objects so distant that their light arrives in parallel rays, and so far away that their images are tiny. The simplest astronomical telescope uses two converging lenses. The first, the **objective**, has a long focal length f_o and forms a real image of the distant object near its focal point. The second, the **eyepiece**, has a short focal length f_e and is used as a magnifying glass to look at that image. What matters for a telescope is not the size of the image but the angle it fills in your view, and the two focal lengths fix how much that angle is increased: the **angular magnification** is f_o/f_e. A 1000 mm objective with a 25 mm eyepiece makes a distant object look 40 times larger in angle, so the Moon, which normally fills half a degree of sky, fills about 20 degrees.

::: note
A lens bends each ray from one point of an object so that the rays meet again at one point of an image. The thin lens equation, 1/f = 1/d_o + 1/d_i, says where that image forms, and its sign says what kind it is: positive for a real image that light actually reaches, negative for a virtual image the rays only appear to come from.
:::

::: history 1608–1611 · Middelburg, Padua, Prague
Lenses for reading spectacles were being ground in Italy by about 1290, but for three centuries nobody seems to have thought to look through two of them at once. In 1608 Hans Lipperhey, a spectacle-maker in the Dutch town of Middelburg, applied for a patent on a tube with a lens at each end that made distant things look near; it was refused, because others were already making similar devices. News reached Galileo in Padua in 1609. Without a sample to copy, he worked out a design for himself, a converging objective with a diverging eyepiece, ground his own lenses and pushed the magnification to about 20. In January 1610 he turned it on Jupiter and over several nights watched four small points of light shift from side to side of the planet, and concluded that they were moons orbiting it: proof that not everything in the sky circles the Earth. Galileo's instruments worked, but nobody yet understood exactly why. In 1611 Johannes Kepler, in Prague, published the first theory of how lenses form images, and in it he proposed the two-converging-lens design described above. It turns the image upside down, which hardly matters to an astronomer, and gives a much wider field of view, so it became the standard for astronomical telescopes.
:::

::: quiz
? An object is placed 30 cm in front of a converging lens of focal length 10 cm. Where does the image form?
- 10 cm behind the lens
+ 15 cm behind the lens
- 30 cm behind the lens
= The thin lens equation gives 1/d_i = 1/10 − 1/30 = 2/30 per centimetre, so d_i = 15 cm. The positive answer means a real, inverted image on the far side of the lens.
? Why can the image seen through a magnifying glass not be projected onto a screen?
- The image is too small to see on a screen
- The lens absorbs too much of the light
+ The image is virtual: the rays never actually meet where it appears to be
= With the object inside the focal length, the rays are still spreading apart after the lens. They only appear to come from the image point, so no light gathers there for a screen to catch.
? A telescope's objective has a focal length of 1.2 m. Which eyepiece gives the highest magnification?
+ 6 mm
- 25 mm
- 40 mm
= Angular magnification is f_o/f_e, so the eyepiece with the shortest focal length magnifies most. With 6 mm it is 1200/6 = 200 times.
:::
