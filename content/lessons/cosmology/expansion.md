In 1915 most astronomers thought the universe and the Milky Way were essentially the same thing: a vast disc of stars, with nothing much beyond it, static and perhaps eternal. Within fifteen years both ideas were gone. Astronomers learned that the faint, fuzzy spirals seen through telescopes were galaxies like our own, millions of light-years away, and that almost all of them are moving away from us, faster the farther they are. This lesson follows how those two discoveries were made, what kind of measurement each needed, and what the result means. It turns out to mean something much stranger than galaxies flying apart through space.

## How do you tell that a galaxy is moving?

You cannot watch a galaxy move across the sky; it is too far away. But you can measure its motion toward or away from you using its light. When a source of waves moves away from you, each successive wave crest is emitted from a little farther away, so the crests arrive more spread out than they left. The wavelength you receive is longer. For light, longer wavelengths are redder, so this is called a **redshift**. A source moving toward you gives a blueshift instead. This is the Doppler effect, the same thing that lowers the pitch of a passing siren.

To use it, astronomers need to know what wavelength the light had when it left. Fortunately, starlight carries fingerprints. When light passes through the cooler outer layers of a star, atoms of each element absorb a few very specific wavelengths, leaving dark lines in the spectrum at positions that can be measured in a laboratory on Earth. Find the same pattern of lines in a galaxy's spectrum, shifted toward the red, and the size of the shift tells you the speed. The redshift z is defined as the fractional stretch of the wavelength.

::: equation Redshift, and speed for nearby galaxies
z = \frac{\lambda_{\mathrm{obs}} - \lambda_{\mathrm{emit}}}{\lambda_{\mathrm{emit}}}, \quad v \approx cz
:::

The first formula is a definition: the observed wavelength minus the emitted one, as a fraction of the emitted one. The second, which turns redshift into a speed by multiplying by the speed of light c, is an approximation that holds only when z is much smaller than 1. For larger redshifts, as we will see, "speed" is no longer even the right way to think about it.

::: history 1912–1917 · Flagstaff, Arizona
Vesto Slipher at Lowell Observatory was the first to measure the spectra of spiral nebulae, as the fuzzy spirals were then called. They were so faint that a single spectrum could need exposures lasting several nights, with the telescope kept pointed by hand. In 1912 he found that the great spiral in Andromeda was approaching the Sun at about 300 kilometres per second, already faster than almost any star. As he measured more, a pattern emerged. By 1917, of the 25 spirals he had measured, all but a few were moving away, some at more than 1,000 kilometres per second, speeds far beyond those of any star in the Milky Way. No one knew what to make of it, because no one knew how far away the spirals were.
:::

## How far away are they?

Distance is the hardest measurement in astronomy. A faint light might be a dim object nearby or a brilliant one far away, and without knowing which, you cannot tell how far it is. What astronomers needed was a **standard candle**, an object whose true brightness is known, so that how faint it looks gives its distance; brightness falls off as the square of distance.

The first good standard candle came from Henrietta Leavitt, who worked at Harvard College Observatory measuring stars on photographic plates. She studied **Cepheid variables**, stars that swell and shrink, brightening and dimming over a regular period of days to weeks. In 1912 she found that the brighter Cepheids pulse more slowly, with a precise relation between period and true brightness. Measure a Cepheid's period, and you know how bright it really is; compare that with how bright it looks, and you know its distance.

Whether the spirals lay inside or outside the Milky Way was still being argued in 1920, in a famous public debate between Harlow Shapley, who held that they belonged to the Milky Way's own system, and Heber Curtis, who held that they were distant galaxies. The question was settled in 1924, when Edwin Hubble, using the new 100-inch telescope on Mount Wilson in California, found Cepheids in the Andromeda spiral. Their periods and faintness put it far outside the Milky Way. The spirals were galaxies, and the universe was vastly larger than anyone had believed.

## A law of recession

::: history 1927 · Louvain, Belgium
Georges Lemaître was both a Catholic priest and a physicist trained in general relativity, Einstein's theory of gravity. He found, as Alexander Friedmann in Russia had already shown in 1922, that Einstein's equations allow a universe that expands. Lemaître went further and connected the theory to observation: in an expanding universe, galaxies should recede at speeds proportional to their distances, and using Slipher's speeds and rough distances he estimated the rate. He published in French, in a little-read Belgian journal, and the paper went almost unnoticed. When Lemaître met Einstein that year, Einstein, who still believed the universe was static, told him: "Your calculations are correct, but your physics is abominable."
:::

In 1929 Hubble, combining Slipher's redshifts with his own distances to 24 galaxies, published a plot of recession speed against distance. The points scattered widely, but the trend was clear: the farther the galaxy, the faster it receded, in a roughly straight line. Hubble was cautious about what this meant, but the result made the expanding universe impossible to ignore, and it was soon known as Hubble's law. In 2018 the International Astronomical Union recommended calling it the Hubble–Lemaître law, to recognise that Lemaître had predicted and measured it first.

::: equation The Hubble–Lemaître law
v = H_0 d
:::

The constant H₀, called the **Hubble constant**, is the expansion rate today. Astronomers quote it in an odd-looking unit, kilometres per second per megaparsec, where a megaparsec (Mpc) is about 3.26 million light-years. The unit says how much extra recession speed each extra megaparsec of distance brings. The modern value is about 70 km/s/Mpc, so a galaxy 100 Mpc away recedes at about 7,000 km/s.

Hubble's own value was about 500 km/s/Mpc, seven times too large. His redshifts were fine; the problem was distance. In 1952 Walter Baade discovered that there are two different kinds of Cepheid, and that Hubble had calibrated his distances with the wrong kind. All his galaxies were in fact more than twice as far away as he had thought, and further corrections, mostly by Allan Sandage, gradually brought H₀ down toward its present value.

Here is how the law is used in practice, to turn a measured spectrum into a distance.

::: steps From a spectrum to a distance
$ \lambda_{\mathrm{emit}} = 656.3\,\text{nm}, \quad \lambda_{\mathrm{obs}} = 663.0\,\text{nm}
Start with a known spectral line. Hydrogen gas glows at 656.3 nanometres, in the red, and this line is easy to pick out in many galaxies. Suppose a galaxy's spectrum shows it at 663.0 nanometres instead.
---
$ z = \frac{663.0 - 656.3}{656.3} \approx 0.0102
Apply the definition of redshift: the shift of 6.7 nanometres, divided by the emitted wavelength. The light has been stretched by about 1%. Notice that the answer does not depend on which line we used; every line in the spectrum is stretched by the same fraction, which is how astronomers confirm they have identified the lines correctly.
---
$ v \approx cz = 3.00 \times 10^5\,\text{km/s} \times 0.0102 \approx 3060\,\text{km/s}
Because z is much smaller than 1, the simple approximation v ≈ cz is accurate here. Multiplying by the speed of light gives a recession speed of about 3,060 kilometres per second.
---
$ d = \frac{v}{H_0} \approx \frac{3060\,\text{km/s}}{70\,\text{km/s/Mpc}} \approx 44\,\text{Mpc}
Rearrange the Hubble–Lemaître law to give distance: divide the speed by H₀. The km/s cancels and leaves megaparsecs. The galaxy is about 44 Mpc, or about 140 million light-years, away.
---
$ \frac{\Delta v}{v} \approx \frac{300\,\text{km/s}}{3060\,\text{km/s}} \approx 10\%
One caution. Besides being carried along by the expansion, every galaxy also moves through its surroundings, pulled by its neighbours, at a few hundred kilometres per second. That motion adds to or subtracts from the measured redshift. For this galaxy it could make the distance wrong by up to about 10%, and for nearer galaxies the error is larger, which is why the law is most reliable at large distances.
:::

## What exactly is expanding?

At first sight the law seems to put us in a special place, at the centre of everything, with all galaxies fleeing from us. It does not. Suppose every distance in the universe grows by the same factor over the same time. Then a galaxy twice as far from you moves away twice as fast, simply because it has twice as much distance to grow. And the same is true from any galaxy, not just ours: every observer sees every other galaxy receding at a speed proportional to its distance. A uniform expansion looks the same from everywhere.

The one analogy worth making is a loaf of raisin bread rising in the oven. As the dough expands, every raisin moves away from every other, and from any raisin's point of view the others recede, those twice as far moving twice as fast. The analogy breaks down in one important way: a loaf has an edge and a centre, and as far as anyone can tell the universe has neither.

General relativity gives a deeper description. In it, galaxies are not flying through a fixed space; space itself is stretching, carrying the galaxies along, and the wavelength of light travelling through space is stretched with it. That is the real origin of the cosmological redshift, and it is why the Doppler formula holds only approximately and only nearby. The redshift is really a measure of how much the universe has grown while the light was in transit: 1 + z is the ratio of the size of the universe now to its size when the light set out. Light with z = 1 left its galaxy when the universe was half its present size. Within galaxies, and within stars and atoms, gravity and other forces hold things together firmly enough that they do not expand; it is only the distances between widely separated galaxies that grow.

## Do the measurements agree?

As measurements have become more precise, a puzzle has appeared. There are now two very different ways to find H₀. One uses the cosmic microwave background, the light left over from the early universe, together with the standard model of how the universe has evolved since, to predict the expansion rate today. It gives H₀ ≈ 67.4 ± 0.5 km/s/Mpc. The other, the **distance ladder**, builds outward step by step, using Cepheids to calibrate the brightness of exploding stars called Type Ia supernovae, which are bright enough to be seen far away. The leading team using this method gets about 73 ± 1 km/s/Mpc.

The two differ by about 8%, much more than their stated uncertainties allow. This **Hubble tension** has survived years of checking, including new observations of Cepheids with the James Webb Space Telescope, though some teams using other kinds of stars as calibrators get values near 70, between the two. Either one of the measurements has a hidden error that nobody has found, or the standard model of the universe is missing something, and which it is remains one of the open questions of cosmology.

::: note
Distant galaxies recede at speeds proportional to their distances, v = H₀d, with H₀ about 70 km/s/Mpc. This does not put us at the centre: every observer in a uniformly expanding universe sees the same law, because it is space itself that is stretching.
:::

::: quiz
? If every galaxy recedes from us with v ∝ d, are we at the centre of the universe?
- Yes, the law singles out our position
+ No, every observer in a uniformly expanding universe sees the same law
- Only if the universe is finite
= When every distance grows by the same factor, observers in every galaxy see all the others receding at speeds proportional to their distances. Every point looks like the centre, so none is.
? A galaxy has redshift z = 0.02. With H₀ = 70 km/s/Mpc, about how far away is it?
- 4 Mpc
+ 86 Mpc
- 860 Mpc
= The speed is v ≈ cz ≈ 6,000 km/s, and dividing by H₀ gives d ≈ 6,000/70 ≈ 86 Mpc.
? Why did Hubble's value of H₀ come out about seven times too large?
- He measured redshifts incorrectly
+ His distance calibration was wrong, so his galaxies seemed too close
- The universe was expanding faster in 1929
= H₀ is speed divided by distance, so underestimated distances make it too large. Hubble's redshifts were accurate; his Cepheid calibration was not.
:::
