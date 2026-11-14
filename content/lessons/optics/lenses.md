A lens is a piece of glass shaped so that refraction at its two curved surfaces sends rays from one point of an object to one point of an image. Your eye, a phone camera and a telescope all do the same job: collect light spreading from each point of a scene and bring it back together somewhere else.

## Converging and diverging lenses

A lens thicker in the middle than at the edges is **converging**: rays arriving parallel to its axis bend inward and cross at a point, the **focal point**, a distance f from the lens called the **focal length**. A lens thinner in the middle is **diverging**: parallel rays spread out as if they came from a focal point on the near side, and f is counted as negative.

The more strongly curved the surfaces, and the higher the refractive index of the glass, the shorter the focal length. Opticians describe a lens by its **power**, 1/f in dioptres when f is in metres: a +2 D reading lens has f = 50 cm.

## Three easy rays

For a thin converging lens, three rays from the top of an object are easy to draw:

1. A ray parallel to the axis bends through the far focal point.
2. A ray through the centre of the lens goes straight on, because there the two surfaces are nearly parallel, like a thin flat pane.
3. A ray through the near focal point leaves parallel to the axis.

Where they meet is the image of the top of the object. Any two are enough.

## The thin lens equation

Put an object of height h_o at distance d_o from a converging lens of focal length f. Its image forms at distance d_i on the other side, upside down, with height h_i.

::: steps Deriving the thin lens equation
$ \frac{h_i}{h_o} = \frac{d_i}{d_o}
Follow the ray through the centre. It makes two similar triangles: one with the object (height h_o, base d_o) and one with the image (height h_i, base d_i), meeting at the centre of the lens.
---
$ \frac{h_i}{h_o} = \frac{d_i - f}{f}
Follow the ray that starts parallel to the axis. It leaves the lens at height h_o and passes through the focal point. Between lens and focal point it forms a triangle of height h_o and base f; beyond the focal point it continues to the image, forming a similar triangle of height h_i and base d_i − f.
---
$ \frac{d_i}{d_o} = \frac{d_i}{f} - 1
Both expressions equal h_i/h_o, so set them equal and split the right-hand fraction.
---
$ \frac{1}{d_o} = \frac{1}{f} - \frac{1}{d_i}
Divide every term by d_i.
---
$ \frac{1}{f} = \frac{1}{d_o} + \frac{1}{d_i}
Rearrange. This is the thin lens equation.
:::

The **magnification** follows from the first step: m = −d_i/d_o, the minus sign meaning the image is inverted.

A camera lens with f = 50 mm photographing a person 2 m away forms the image at 1/d_i = 1/50 − 1/2000 per mm, so d_i ≈ 51.3 mm. Focusing a camera means moving the lens that extra 1.3 mm. For a very distant object 1/d_o ≈ 0 and the image sits at the focal point.

## Real and virtual images

When d_i comes out positive, light really converges to the image, and a screen placed there shows it. That is a **real image**: the one on a camera sensor or a cinema screen.

Put the object closer than f, and 1/d_i = 1/f − 1/d_o is negative. The rays leaving the lens still diverge; they only appear to come from a point on the object's side. That is a **virtual image**: upright, magnified, and impossible to catch on a screen. A magnifying glass works this way.

## Eyes and telescopes

The eye focuses mainly with its curved cornea, which supplies about two-thirds of its roughly 60 dioptres of power. The flexible lens behind it adds the rest and changes shape to focus at different distances, forming a real, inverted image on the retina; the brain handles the inversion. In short sight, the eye is too strong for its length, so distant objects focus in front of the retina, and a diverging spectacle lens corrects it.

A simple telescope uses two converging lenses. The objective, with long focal length f_o, forms a real image of a distant object; the eyepiece, with short focal length f_e, acts as a magnifying glass on that image. The angular magnification is f_o/f_e: a 1000 mm objective with a 25 mm eyepiece magnifies 40 times.

::: note
A lens refracts every ray from one object point to one image point. The thin lens equation, 1/f = 1/d_o + 1/d_i, fixes where the image forms; its sign tells you whether it is real or virtual.
:::

::: history 1608–1611 · Middelburg, Padua, Prague
Spectacles for the long-sighted were in use in Italy by about 1290. In 1608 the Dutch spectacle-maker Hans Lipperhey applied for a patent on a device of two lenses that made distant things look near. Galileo heard of it in 1609, built his own, improved it to about 20 times, and in January 1610 saw four moons orbiting Jupiter. Galileo's design used a diverging eyepiece. In 1611 Johannes Kepler worked out the theory of lenses and proposed the two-converging-lens design described above, which gives an inverted image but a wider field, and became the standard for astronomy.
:::

::: quiz
? An object is placed 30 cm in front of a converging lens of focal length 10 cm. Where is the image?
- 10 cm behind the lens
+ 15 cm behind the lens
- 30 cm behind the lens
= 1/d_i = 1/10 − 1/30 = 2/30, so d_i = 15 cm.
? Why can you not project the image from a magnifying glass onto a screen?
- The image is too small
- The lens absorbs too much light
+ The image is virtual: the rays never actually meet there
= With the object inside the focal length, the rays diverge after the lens and only appear to come from the image.
? A telescope's objective has focal length 1.2 m. Which eyepiece gives the highest magnification?
+ 6 mm
- 25 mm
- 40 mm
= Angular magnification is f_o/f_e, so the shortest eyepiece focal length magnifies most: 1200/6 = 200 times.
:::
