Gravity pulls every piece of matter in the universe toward every other piece, so for most of the twentieth century astronomers took it for granted that the expansion of the universe must be slowing down. The only question was by how much. If the universe held enough matter, its gravity would eventually stop the expansion and pull everything back together; if not, the expansion would slow forever without quite stopping. Measuring the rate of slowing would decide the fate of the universe. In the 1990s two rival teams set out to measure it, and both found the opposite of what they were looking for: the expansion is speeding up. This lesson explains how they made the measurement, why the result was so surprising, how general relativity allows something to push the universe apart, and why the leading explanation leaves physicists with one of the worst mismatches between theory and observation in all of science.

## How do you measure the history of the expansion?

The light from a distant galaxy is a message from the past. Light from a galaxy billions of light-years away left it billions of years ago, so comparing distant galaxies with near ones shows how the universe has changed over time. Two measurements are needed for each galaxy. The first is its redshift z, which says how much the universe has stretched while the light was travelling: 1 + z is the ratio of the universe's size now to its size then. Redshifts are easy to measure from spectra. The second is the galaxy's distance, which tells how long the light has been travelling. Distances are the hard part.

The tool that made the measurement possible was the **Type Ia supernova**. This is the explosion of a **white dwarf**, the small, dense remnant of a star like the Sun, which has been gaining matter from a companion star until it becomes unstable and blows itself apart in a thermonuclear explosion. For a few weeks one can outshine an entire galaxy of a hundred billion stars, which means it can be seen across billions of light-years. Because white dwarfs explode under similar conditions, Type Ia supernovae are all roughly equally bright at their peak, but not exactly. In 1993 Mark Phillips showed that the variation follows a rule: the intrinsically brighter ones fade more slowly. So by measuring how fast a supernova fades, astronomers can correct for the difference and find its true brightness, well enough to give its distance to within about 10%. An object of known true brightness is called a **standard candle**: since apparent brightness falls off as the square of distance, how faint it looks tells you how far away it is.

Now consider what the measurement should show. Suppose the expansion has been slowing. Then in the past it was faster than it is now, so the universe reached its present size relatively quickly. For a supernova at a given redshift, meaning a given amount of stretching since its light left, that light would have been travelling for a relatively short time and would have covered a relatively short distance. The supernova would therefore look comparatively bright. If instead the expansion has been speeding up, it was slower in the past, the same amount of stretching took longer, and the light has come farther. The supernova would look fainter. Comparing how bright distant supernovae look with their redshifts therefore measures whether the expansion has been slowing or speeding up.

::: history 1998 · Berkeley and Mount Stromlo
Two teams competed to make this measurement. The Supernova Cosmology Project was led by Saul Perlmutter at the Lawrence Berkeley Laboratory in California. The High-z Supernova Search Team was led by Brian Schmidt at Mount Stromlo Observatory in Australia, with Adam Riess as lead author of its key paper. Finding distant supernovae was itself a challenge, because they are rare and unpredictable; both teams developed methods to image the same patches of sky weeks apart and catch new explosions while they were still brightening, so that larger telescopes could follow them up. Each team collected dozens of supernovae at redshifts around 0.5, when the universe was about two-thirds of its present size. Both found them about 20–25% fainter than expected in a universe whose expansion was being slowed by its matter, fainter even than they would be in an empty universe with no gravity at all. The expansion had been speeding up. Each team had hoped to beat the other, and each half expected to find a mistake in its own work. Getting the same unwelcome answer independently was what made the result convincing. Perlmutter, Schmidt and Riess shared the 2011 Nobel Prize in Physics.
:::

## Einstein's constant returns

::: history 1917 · Berlin
When Einstein first applied his new theory of gravity, general relativity, to the universe as a whole, he found that it could not stay still: gravity would make it collapse unless it was already expanding. Like nearly everyone at the time, Einstein believed the universe was static. So he added a term to his equations, the **cosmological constant**, written Λ (the Greek capital lambda), which acts as a uniform repulsion built into space itself and could balance gravity's pull. When the expansion of the universe was discovered in the late 1920s, the term was no longer needed, and Einstein abandoned it. George Gamow later recalled Einstein calling it the biggest blunder of his life, though historians doubt that Einstein ever used those words.
:::

The accelerating expansion is most simply explained by Einstein's Λ. Physicists today usually describe it in a different but equivalent way, as an energy that belongs to empty space itself, called **vacuum energy**. Its defining property is that its density does not change as space expands. Ordinary matter thins out as the universe grows, because the same amount of matter is spread through more volume. Vacuum energy does not: every new cubic metre of space comes with the same amount of it.

How can that make the expansion speed up? The answer lies in a feature of general relativity with no counterpart in Newton's gravity: pressure, not only mass, is a source of gravity. In everyday situations the effect of pressure is negligible, but on the scale of the universe it can dominate.

::: steps Why vacuum energy accelerates the expansion
$ \frac{1}{a}\frac{\mathrm{d}^2 a}{\mathrm{d}t^2} = -\frac{4\pi G}{3}\left(\rho + \frac{3p}{c^2}\right)
This is Friedmann's acceleration equation, which follows from general relativity for a uniform universe. Here a is the **scale factor**, a number that tracks the relative size of the universe, so its second derivative with time measures whether the expansion is speeding up or slowing down. On the right, ρ is the mass density and p is the pressure of whatever fills space. Note that pressure enters alongside density, and counts three times. A positive right-hand side means acceleration.
---
$ p \approx 0 \Rightarrow \frac{1}{a}\frac{\mathrm{d}^2 a}{\mathrm{d}t^2} < 0
For galaxies and dark matter, the pressure is effectively zero, because the particles move slowly compared with light. Then the right-hand side is just −(4πG/3)ρ, which is negative because density is positive. Ordinary matter, as expected, slows the expansion down.
---
$ \mathrm{d}E = -p\,\mathrm{d}V, \quad E = \rho c^2 V
To find the pressure of vacuum energy, use energy conservation for a region of space that is expanding along with the universe. When a region of volume V grows by dV, the energy inside changes by −p dV, the work done by its pressure on the surroundings, exactly as for a gas pushing a piston. The energy inside is the mass density times c² times the volume.
---
$ \rho = \text{constant} \Rightarrow \rho c^2\,\mathrm{d}V = -p\,\mathrm{d}V
For vacuum energy the density stays fixed as the region grows, so the energy inside increases by ρc² for every bit of added volume. Set that increase equal to −p dV, as energy conservation requires.
---
$ p = -\rho c^2
Cancel dV from both sides. The pressure of vacuum energy is negative, and equal in size to its energy density. A negative pressure is a tension, a pull rather than a push, like that of a stretched rubber sheet. That is the only way the energy in an expanding region can grow: the expansion does work against the tension, and that work appears as new vacuum energy.
---
$ \rho + \frac{3p}{c^2} = \rho - 3\rho = -2\rho
Now insert this pressure into the bracket on the right of Friedmann's equation. Because the pressure counts three times and is negative, it outweighs the density, and the bracket becomes negative.
---
$ \frac{1}{a}\frac{\mathrm{d}^2 a}{\mathrm{d}t^2} = +\frac{8\pi G}{3}\rho_{\mathrm{vac}} > 0
With a negative bracket, the minus sign in front turns the right-hand side positive. Vacuum energy makes the expansion accelerate. Writing ρ_vac = Λc²/(8πG) turns this into Einstein's term, Λc²/3, which shows that the cosmological constant and a constant vacuum energy are the same thing.
:::

The result is counterintuitive. Vacuum energy has positive energy density, so it does gravitate, but its tension gravitates even more strongly and with the opposite sign. The net effect is repulsive.

This also explains why the expansion slowed first and only later sped up. In the early universe, matter was densely packed, and its gravity dominated, so the expansion decelerated. But as space grew, matter thinned out, its density falling as one over the cube of the scale factor, while the density of vacuum energy stayed fixed. Eventually the vacuum term won. From the acceleration equation, that happens once the vacuum density exceeds half the matter density, and it happened about 6 billion years ago. Today vacuum energy, or whatever is playing its role, makes up about 68% of the energy of the universe.

Since 1998 the result has been confirmed by completely independent methods. The cosmic microwave background shows that space is flat, which in general relativity requires a particular total density; matter, ordinary and dark together, supplies only about a third of it, and the rest has to be something else. And the imprint of sound waves from the early universe, frozen into the distribution of galaxies at a characteristic separation, acts as a standard ruler whose apparent size at different distances traces the same history of expansion that the supernovae show.

## What is it?

The name **dark energy** is used for whatever is causing the acceleration, without committing to what it is. The cosmological constant fits nearly all the data, but it raises a severe theoretical problem. Quantum theory predicts that empty space should have an energy, because every quantum field has a minimum energy that cannot be removed. When physicists estimate that energy, the answer is larger than the observed dark energy density by something like 120 orders of magnitude, a factor of one followed by 120 zeros. Some unknown effect must cancel almost all of it, but not quite all, and no one knows what that effect is or why the remainder is so small but not zero.

Nor is it certain that dark energy is constant. Since 2024 the DESI survey, which has mapped the positions of millions of galaxies to trace the expansion history in detail, has reported hints that dark energy may be weakening over time, something a true cosmological constant cannot do. The hints grew stronger in its 2025 results but are not yet decisive, and confirming or ruling them out is a major goal of current surveys. If dark energy is constant, the expansion will keep accelerating forever, and galaxies beyond our own local group will eventually recede so fast that their light can no longer reach us. Astronomers in the very distant future would see an island of stars in an otherwise empty, dark universe.

::: note
Distant supernovae are fainter than they would be if the expansion were slowing, so the expansion is accelerating. In general relativity pressure is a source of gravity, and anything with a strongly negative pressure, such as a constant vacuum energy, drives the expansion faster. What dark energy really is, and why its density is so small, are unsolved.
:::

::: quiz
? If the expansion had been decelerating, how would distant supernovae appear compared with an accelerating universe?
+ Brighter, because their light would have travelled a shorter distance
- Fainter, because gravity absorbs light
- Exactly the same
= In a decelerating universe the expansion was faster in the past, so a given redshift is reached sooner and the light has a shorter distance to travel. The observed supernovae were fainter than that, showing acceleration.
? Why does matter's influence on the expansion fade while vacuum energy's grows?
- Matter is converted into vacuum energy
+ Matter's density falls as space expands, while vacuum energy's density stays constant
- Vacuum energy only appeared 6 billion years ago
= Matter is spread through ever more volume, so its density falls as 1/a³. Vacuum energy does not dilute, so it eventually outweighs matter in the acceleration equation.
? In the acceleration equation, why does vacuum energy push rather than pull?
- Its density is negative
+ Its negative pressure, counted three times, outweighs its positive density
- It does not interact through gravity
= With p = −ρc², the bracket ρ + 3p/c² becomes −2ρ, which is negative, so the acceleration of the expansion is positive.
:::
