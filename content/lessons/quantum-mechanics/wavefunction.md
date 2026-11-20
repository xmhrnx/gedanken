By the mid-1920s physicists had accepted something strange: electrons, which arrive at a detector as single, whole particles, also behave like waves. Louis de Broglie had given every particle a wavelength, λ = h/p, where h is Planck's constant and p the particle's momentum, and beams of electrons were soon seen to diffract off crystals just as X-rays do. That left a question nobody could answer at first. If an electron is a wave, a wave of *what*? This lesson is about the answer that quantum mechanics settled on, the wavefunction, and the single rule that ties it to anything we can measure. Almost everything else in the subject is built on top of it. By the end you should be able to say what ψ is and is not, turn a wavefunction into the odds of finding a particle somewhere, explain why two open slits can make a spot on a screen less likely to be hit than one, and see where the phase of the wave, which no single measurement shows, hides the particle's motion.

## A wave of what?

The evidence for electron waves was by then hard to argue with. In 1927 Clinton Davisson and Lester Germer, at Bell Labs in New York, fired slow electrons at a nickel crystal and found them bouncing off strongly at some angles and hardly at all at others, exactly as X-rays do when the waves scattered by neat rows of atoms reinforce in some directions and cancel in others. In the same year George Paget Thomson, the son of J. J. Thomson, sent faster electrons through thin metal films and photographed rings of the kind X-rays make. The angles matched de Broglie's wavelength. Something about an electron really does wave.

The obvious first guess is that the electron itself is spread out, its mass and charge smeared through space the way the energy of a sound wave is spread through air. That guess fails the moment you look at a detector. A screen hit by an electron beam does not glow faintly all over; it records sharp, separate dots, one per electron. A Geiger counter clicks once, not a little. Whatever is doing the waving, the electron is never found in pieces.

So the wave cannot be the electron's substance. It has to be something that governs where the whole electron turns up, without itself being the electron. Quantum mechanics calls this thing the **wavefunction**, written ψ (the Greek letter psi), and treats it as the complete description of the particle's state. Everything that can be predicted about the particle, where it will be found, how fast it is moving, what energy it has, is to be extracted from ψ by some rule. The rest of this lesson is about the first and most important of those rules.

## What the wavefunction is

The wavefunction ψ(x, t) assigns a number to every point x in space at every moment t. In that respect it is like the height of a water wave, which also has a value at each place and time. But the number is not an ordinary one. It is a **complex number**, which means it has two parts that are best pictured as a length and a direction on a dial. The length is called the size, |ψ|, and the direction is called the **phase**, an angle that says where in its cycle the wave is at that point.

The dial picture is worth holding on to, because it makes the arithmetic of complex numbers easy to follow. Multiplying two complex numbers multiplies their lengths and adds their angles. Adding them works like adding arrows: put the second arrow's tail on the first arrow's tip, and the sum runs from the start to the end. Two arrows of equal length pointing the same way add to an arrow twice as long; pointing in opposite directions they add to nothing. That last fact is the seed of everything strange in this lesson.

Neither part can be observed directly. There is no instrument that reads off ψ at a point the way a thermometer reads temperature. The wavefunction earns its place in physics only through a rule that turns it into something we can count.

::: equation The Born rule
P(x \text{ to } x + \mathrm{d}x) = |\psi(x,t)|^2\,\mathrm{d}x
:::

In words: the probability of finding the particle in a small interval of width dx around the point x is the square of the wavefunction's size there, times the width of the interval. That makes |ψ|² a **probability density**, a probability per unit length, or per unit volume in three dimensions. Where |ψ|² is large the particle is likely to be found; where it is small the particle rarely turns up; and where ψ is exactly zero the particle is never found at all. The phase drops out of |ψ|², which is why it cannot be seen in a single position measurement, though, as we will see shortly, it is far from irrelevant.

A density has units, and they are worth checking. In one dimension the probability P is a pure number and dx is a length, so |ψ|² has units of one over length, and ψ itself has units of one over the square root of a length. In three dimensions |ψ|² is a probability per cubic metre. This is a small point with a practical use: if a formula for ψ comes out with the wrong units, it is wrong.

There is one condition every wavefunction must meet. The particle exists, so if you search all of space you must find it somewhere, and the probabilities over all space must add up to exactly one:

::: equation Normalisation
\int_{-\infty}^{\infty} |\psi(x,t)|^2\,\mathrm{d}x = 1
:::

A wavefunction scaled to meet this condition is called normalised. The scaling is always possible for a wave that dies away at large distances: compute the integral, and divide ψ by its square root. What cannot be done is to normalise a wave that stays the same size forever in every direction, a point we will come back to under the limits of the idea.

## Probabilities, not outcomes

It is worth pausing on how much the Born rule gives up. It does not say where a particular electron will land. Two electrons prepared in exactly the same way, with identical wavefunctions, can arrive at quite different places. What the rule predicts is the pattern: repeat the experiment many times and the density of hits across the screen traces out |ψ|². The theory is precise about the distribution and silent about the individual event.

That was a sharp break with the physics before it. In Newton's mechanics, if you know where a particle is and how it is moving, you know where it will be. Uncertainty about the outcome meant only that you had not measured carefully enough. In quantum mechanics, as far as anyone has been able to tell, the probability is not a cover for missing information. The wavefunction is everything there is to know, and it still only gives odds.

From those odds come averages. If you prepare many identical particles and measure each one's position, the mean of all your results is called the **expectation value** of position. It is the position weighted by how likely it is, summed over every place:

::: equation Expectation value of position
\langle x \rangle = \int_{-\infty}^{\infty} x\,|\psi(x,t)|^2\,\mathrm{d}x
:::

The name is misleading in one respect. The expectation value need not be a result you ever get. For a particle whose |ψ|² has two humps, one on each side of the origin, the average position can be at the origin, exactly where the particle is least likely to be found. It is the centre of the distribution, not a prediction for any single trial.

::: history 1926 · Göttingen
Erwin Schrödinger had just written down the equation that tells ψ how to change in time, and he believed the wave was real in a direct sense: a cloud of electric charge, spread out in space. Max Born, in Göttingen, tested that idea on the problem of an electron bouncing off an atom. The scattered wave spreads out in every direction, like ripples from a stone. If it were the electron's charge, each electron would end up as a faint shell expanding outward. But each scattered electron is detected in one place, in one direction. Born concluded that the wave does not carry the electron; it carries the chance of finding it. In the first version of his paper he wrote that ψ gave the probability. While correcting the proofs he added a footnote: on more careful consideration, the probability is proportional to the *square* of ψ. That footnote became the interpretation the whole field uses. Schrödinger never accepted it. Born received the Nobel Prize for it in 1954, almost thirty years later.
:::

## Why square it?

A reasonable reader might ask why physics needs an intermediate quantity at all. Why not describe the electron directly with a probability that spreads out like a wave? The answer is that probabilities on their own cannot do what electrons do. Probabilities are never negative, so adding a second possible route can only make an outcome more likely, never less. Electrons break that rule, and the reason they can is that it is ψ, not the probability, that adds.

Take the classic test: a screen with two narrow slits, and a detector screen behind it. Call ψ₁ the wave that reaches a given point having come through slit 1, and ψ₂ the wave arriving through slit 2. Write each in terms of its size and phase as ψ = |ψ|e^{iφ}, where φ is the phase angle.

::: steps Probability with two open slits
$ \psi = \psi_1 + \psi_2
With both slits open, the two waves arrive at the same point and superpose, which means they simply add. This is the key physical assumption, and it is the one experiment supports: in quantum mechanics it is amplitudes, the values of ψ, that add, not probabilities. Nothing has been squared yet.
---
$ |\psi|^2 = (\psi_1 + \psi_2)^*(\psi_1 + \psi_2)
To get the probability density we need the squared size of the total. For any complex number z, the squared size is |z|² = z*z, where z* is the complex conjugate: the number with the same size and the opposite phase. Multiplying a number by its conjugate cancels the phase and leaves a real, positive result, which is what a probability must be.
---
$ |\psi|^2 = \psi_1^*\psi_1 + \psi_2^*\psi_2 + \psi_1^*\psi_2 + \psi_1\psi_2^*
Multiply out the brackets term by term, as you would with ordinary numbers. The first two products are just |ψ₁|² and |ψ₂|². The last two mix the two waves together, and they are where everything interesting happens.
---
$ |\psi|^2 = |\psi_1|^2 + |\psi_2|^2 + 2\,\mathrm{Re}\left(\psi_1^*\psi_2\right)
The two mixed terms are complex conjugates of each other, so their imaginary parts cancel and their real parts double. Their sum is therefore twice the real part of ψ₁*ψ₂. This step has turned two awkward complex terms into a single real number that can be positive or negative.
---
$ |\psi|^2 = |\psi_1|^2 + |\psi_2|^2 + 2|\psi_1||\psi_2|\cos(\phi_2 - \phi_1)
Writing each wave as size times phase, ψ₁*ψ₂ = |ψ₁||ψ₂|e^{i(φ₂ − φ₁)}, whose real part is |ψ₁||ψ₂|cos(φ₂ − φ₁). So the mixed term depends on the difference between the two phases. The phases, which no single measurement can see, decide whether the probability at this point goes up or down.
:::

Look at what each piece means. The first two terms are what you would expect if electrons behaved like bullets: the chance of arriving through slit 1 plus the chance of arriving through slit 2. The third term, called the **interference term**, has no counterpart for bullets. Because the two paths have different lengths to different points on the screen, the phase difference changes as you move across it, and the cosine swings between +1 and −1. Where the waves arrive in step, the cosine is +1 and, if the two waves are equally strong, the probability is four times what one slit alone would give, not twice. Where they arrive exactly out of step, the cosine is −1 and the probability falls to zero. Opening a second slit makes some places on the screen *less* likely to be hit, something no theory of adding probabilities could produce.

Notice also that the total is not changed by the interference, only rearranged. Averaged across many fringes, the cosine is as often positive as negative, so the interference term averages to zero and the total number of electrons reaching the screen is the same as the bullet picture predicts. The electrons that are missing from the dark bands turn up in the bright ones. Normalisation is never violated; probability is moved around, not created.

::: model interference
Change the slit separation and watch the bright and dark fringes move closer together or further apart. For electrons, the bright bands are where |ψ|² is large: the places where many electrons land, each one arriving separately.
:::

## One electron at a time

A natural objection is that the electrons in a beam might be interfering with each other, jostling into a pattern the way crowds form lanes. The experiment has been done in a way that rules this out. Turn the source down until only one electron is in the apparatus at a time, with long gaps between them. Each electron makes a single dot on the screen, apparently at random. But after thousands of dots the same fringes appear. There was never a second electron to interfere with, so each electron's own wavefunction must have gone through both slits, and its own |ψ|², interference term included, set the odds of where it landed.

::: history 1961–1989 · Tübingen, Bologna, Tokyo
For decades the electron double slit was a thought experiment, because slits fine enough to diffract electrons were too hard to make. In 1961 Claus Jönsson in Tübingen managed it, cutting slits a fraction of a micrometre wide in a metal foil, and saw interference fringes from up to five of them. The next question was whether the pattern survives when electrons pass one at a time. In 1974 Pier Giorgio Merli, Gian Franco Missiroli and Giulio Pozzi in Bologna used an electron biprism, a fine charged wire that splits the electron wave in two, and recorded the pattern forming from single electrons. In 1989 Akira Tonomura's team at Hitachi near Tokyo made a film of it: scattered dots, apparently random, that slowly gather into fringes.
:::

There is a sequel. Add a detector that records which slit each electron passes through, and the fringes vanish. The hits now follow |ψ₁|² + |ψ₂|², the bullet pattern, as if the interference term had been deleted. Learning the path destroys the interference. Why that happens, and what exactly a measurement does to ψ, is the subject of the lesson on measurement.

## Waves that can be negative

The wavefunction can take negative values, or more generally any phase, while the probability |ψ|² never can. The simplest place to see this is a particle trapped in a box. Its allowed wavefunctions are standing waves, like the vibrations of a guitar string, with points called **nodes** where ψ crosses zero.

For a box of width L with walls the particle cannot pass, ψ must vanish at both walls, and the standing waves that fit are sine curves with a whole number n of half-wavelengths across the box. The normalised versions are:

::: equation Particle in a box
\psi_n(x) = \sqrt{\frac{2}{L}}\,\sin\left(\frac{n\pi x}{L}\right)
:::

The factor √(2/L) is what normalisation demands. The average of sin² over whole half-cycles is exactly one half, so the integral of sin² across the box is L/2, and the factor 2/L in |ψ|² brings the total to one. The units come out right too: one over the square root of a length.

::: model wavefunction
A particle trapped in a box. Step through the levels and compare ψ with |ψ|². The wave ψ dips below zero; the probability never does. At the nodes, where ψ crosses zero, the particle is never found, even though it is found on both sides.
:::

That last point would be absurd for a classical particle bouncing between the walls, which must pass through every point to get from one side to the other. For a wave it is natural, and it is one more sign that the particle does not have a path in the ordinary sense.

## Worked examples

**Where in the box?** Take an electron in its lowest state, n = 1, in a box of width L = 1 nm, about the size of a small molecule. What is the chance of finding it in the middle third of the box, between x = L/3 and x = 2L/3? A classical particle bouncing back and forth at constant speed spends equal time everywhere, so its answer would be exactly one third. For the quantum particle we integrate |ψ₁|² over that region:

::: steps Probability in the middle third, n = 1
$ P = \frac{2}{L}\int_{L/3}^{2L/3} \sin^2\left(\frac{\pi x}{L}\right)\mathrm{d}x
The Born rule says the probability of finding the particle in a region is the integral of |ψ|² over it. Here |ψ₁|² is (2/L)sin²(πx/L), and the region runs from one third to two thirds of the way across.
---
$ \sin^2\theta = \tfrac{1}{2}\left(1 - \cos 2\theta\right)
This identity swaps the square of a sine, which is awkward to integrate, for a constant minus a cosine, each of which integrates in one line. It is the same trick that showed the normalisation factor is √(2/L).
---
$ P = \left[\frac{x}{L} - \frac{1}{2\pi}\sin\left(\frac{2\pi x}{L}\right)\right]_{L/3}^{2L/3}
Integrating term by term gives this expression, to be evaluated at the two ends of the region and subtracted. The constant part gives x/L, and the cosine part gives a sine with a factor 1/(2π) in front.
---
$ P = \frac{1}{3} + \frac{\sqrt{3}}{2\pi} \approx 0.61
The first term contributes 2/3 − 1/3 = 1/3, the classical answer. The sine term adds √3/(2π), about 0.28, because sin(4π/3) is −√3/2 and sin(2π/3) is +√3/2. The total is about 61 percent.
:::

The quantum electron in its lowest state crowds toward the middle, nearly twice as likely as a classical particle to be found in the central third. The same calculation for n = 2 gives 1/3 − √3/(4π), about 0.20, because the second level has a node right at the centre. Neither answer depends on the width L: shrink the box to an atom's size or stretch it to a centimetre, and the fractions stay the same. What L does set is the energy, which is the business of the lesson on the Schrödinger equation.

**How far apart are the fringes?** Accelerate electrons through 10,000 volts, so each carries 10 keV of kinetic energy. Its momentum is p = √(2mE), which works out at about 5.4 × 10⁻²³ kg m/s, and de Broglie's relation gives λ = h/p ≈ 1.2 × 10⁻¹¹ m, about 12 picometres. (At this energy the electron moves at about a fifth of the speed of light, and the relativistic correction changes the wavelength by only about half a percent.) Send these through two slits d = 1 µm apart onto a screen L = 1 m away. Bright fringes sit where the two paths differ by a whole number of wavelengths, which for small angles puts them a distance λL/d apart: 1.2 × 10⁻¹¹ × 1 / 10⁻⁶ ≈ 1.2 × 10⁻⁵ m, about 12 micrometres. That is far too fine to see on a phosphor screen by eye, which is why Jönsson had to magnify his pattern with electron lenses before he could photograph it. The same arithmetic explains why nobody sees a tennis ball interfere: its wavelength is around 10⁻³⁴ m, and its fringes would be closer together than anything in the universe can resolve.

**When one wave is weaker.** Suppose one slit is slightly blocked so that at some point on the screen |ψ₂| is only half of |ψ₁|. Measure probabilities in units of |ψ₁|². Without interference you would expect 1 + 0.25 = 1.25. With it, the bright fringes reach (1 + 0.5)² = 2.25 and the dark ones fall to (1 − 0.5)² = 0.25. The fringes are still there but no longer drop to zero. Physicists measure this with the **visibility**, (max − min)/(max + min), here 2.0/2.5 = 0.8. Visibility below one is a sign either of unequal amplitudes or, as the lesson on measurement explains, of the particle leaving partial information about its path behind.

## Going deeper: where the motion hides

If |ψ|² gives the odds of position and the phase drops out, it is fair to ask what the phase is for. The answer is that it carries the particle's motion. Consider the simplest travelling wave, one whose size is the same everywhere and whose phase turns steadily as you move along x:

::: equation A wave of definite momentum
\psi(x) = A\,e^{ikx}, \qquad p = \hbar k
:::

Here k, the **wavenumber**, is 2π divided by the wavelength, so it says how many radians the phase turns per metre, and ħ is Planck's constant divided by 2π. The relation p = ħk is de Broglie's λ = h/p in another form. Now look at |ψ|²: it is |A|² everywhere, perfectly flat. A position measurement can tell you nothing about which way this particle is moving or how fast. Yet the wave has a definite momentum, encoded entirely in how quickly the phase winds. Reverse the winding, e^{−ikx}, and the momentum reverses, with |ψ|² unchanged.

So the phase is not decoration. Two wavefunctions with the same |ψ|² everywhere can describe particles moving in opposite directions, and as soon as they are left alone they will evolve differently, so a later position measurement will tell them apart. The phase information is real; it is just stored in a form that position measurements read only indirectly, through how the wave develops and how its parts interfere.

There is one exception. Multiply the entire wavefunction by the same phase factor, e^{iα} with one fixed angle α everywhere, and nothing measurable changes at all. Every probability involves ψ times its conjugate, and the common factor cancels; every interference term involves a difference of phases, and a common shift drops out of the difference. This **global phase** is pure bookkeeping. Only relative phases, between one place and another or between one part of a superposition and another, have physical effects. This distinction will matter a great deal in the lesson on spin, where turning an electron through a full circle multiplies its wavefunction by −1.

A second point the simple picture hides concerns more than one particle. For a single particle, ψ lives in ordinary space: one number at each point. For two particles, the wavefunction is ψ(x₁, x₂), one complex number for each *pair* of positions, and |ψ(x₁, x₂)|² gives the odds of finding the first particle at x₁ and the second at x₂ together. For three particles in three dimensions, ψ is a function of nine coordinates. This was a large part of why Schrödinger's charge-cloud picture could not be rescued: a function on a nine-dimensional space is not a cloud in our three. It is also the root of entanglement, where the odds for one particle cannot be written separately from those for another.

## Common misconceptions

**"The electron is spread out over the wavefunction."** Every measurement of position finds a whole electron at one place, with its whole charge. Before the measurement, quantum mechanics does not say the electron has a position at all; it says what the odds are. A picture of a smeared-out electron predicts faint, partial hits, and no detector has ever recorded one.

**"The probability just reflects our ignorance of where the particle really is."** If the electron always went through one slit or the other and we simply did not know which, the hits would follow the bullet pattern |ψ₁|² + |ψ₂|². The interference term shows that both routes contribute at once, in a way no hidden choice of one route can mimic. The stronger form of this argument, ruling out a wide class of hidden-information theories, is Bell's theorem, described in the lesson on entanglement.

**"A large |ψ| at a point means the particle is probably exactly there."** |ψ|² is a density. The probability of any single exact point is zero; only regions have probabilities. A tall, narrow peak can carry less total probability than a low, broad hump, and what matters is the area under |ψ|² over the region you are asking about.

**"The wavefunction is the particle's wave in the way a ripple is the water's."** A ripple is a displacement of a medium you can touch. ψ is complex, its global phase is unobservable, and for more than one particle it does not even live in ordinary space. It is better thought of as a catalogue of possibilities with a rule for weighing them.

## Limits of the picture

The wavefunction as described here is the description for one particle moving slowly compared with light, and several things change beyond that.

A wave of perfectly definite momentum, A e^{ikx}, cannot be normalised: its |ψ|² is the same out to infinity, so the integral is infinite whatever A is. Real particles are always described by **wave packets**, sums of many such waves with slightly different wavenumbers that add up to a lump of finite size. The price of building a lump is a spread of momenta, and the trade-off between the two spreads is the uncertainty principle.

For particles moving near the speed of light, the number of particles is no longer fixed: energy can turn into new electron and positron pairs. A wavefunction for a fixed number of particles cannot describe that, and the framework that replaces it is quantum field theory, sketched in the lessons on forces as particles and the Standard Model.

And for large, warm objects in contact with their surroundings, the relative phases between different parts of the wavefunction leak into the environment almost instantly. Interference becomes unobservable in practice, and the object behaves classically. That process, called decoherence, does not change the rules of this lesson, but it explains why we never see them at work on footballs.

## Where it shows up

Every chemistry diagram of a bond is a sketch of |ψ|², and the shapes in the lesson on orbitals and quantum numbers are the wavefunctions of an electron bound to a nucleus. The electron microscope works because electrons are waves with very short wavelengths, so they can resolve detail far below what light can. The scanning tunnelling microscope reads the tiny tail of an electron's wavefunction leaking across a gap, as the lesson on tunneling explains. Interferometers that send atoms or even large molecules down two paths at once use the interference term derived above to measure gravity and rotation with extraordinary precision.

Within this course the idea grows from matter waves, where de Broglie first gave electrons a wavelength, and the double slit, where interference was introduced for light. It leads directly to the Schrödinger equation, which says how ψ changes in time; to the uncertainty principle, which follows from building localised wave packets; and to measurement, which asks what becomes of ψ when someone looks.

::: note
The wavefunction is not a physical wave you can see. It is an amplitude: the square of its size gives the probability of finding the particle at each place. Because amplitudes add before they are squared, quantum probabilities can interfere and even cancel, and the phase that drops out of each probability is where the particle's motion is stored.
:::

::: quiz
? In one region ψ is twice as large as in another. How do the chances of finding the particle compare, per unit length?
- It is twice as likely to be found there
+ It is four times as likely to be found there
- It is equally likely, since only the phase matters
= The probability density is |ψ|², so doubling the size of ψ multiplies the density by four.
? With both slits open, some points on the screen receive fewer electrons than with one slit open. Why?
- Electrons from the two slits collide and knock each other away
+ The amplitudes from the two slits add and can cancel before being squared
- The second slit blocks part of the beam
= The interference term 2|ψ₁||ψ₂|cos(φ₂ − φ₁) is negative wherever the two waves arrive out of step, so the total probability there drops below the one-slit value.
? Electrons are sent through a double slit one at a time, hours apart. What builds up on the screen?
- Two bright bands, one behind each slit
+ An interference pattern, made of individual dots
- A smooth blur with no structure
= Each electron lands at a single point, but the odds of where it lands are set by its own |ψ|², which already contains the interference term, so the fringes emerge as dots accumulate.
? Two wavefunctions, A e^{ikx} and A e^{−ikx}, have exactly the same |ψ|² everywhere. What distinguishes them?
- Nothing: they describe the same state
+ Their phases wind in opposite directions, so they describe particles moving opposite ways
- One describes an electron and the other a positron
= Momentum is carried by how fast and which way the phase turns, p = ħk, so reversing the winding reverses the momentum even though no single position measurement can tell.
? At a point on a screen, the amplitude through one slit is three times the amplitude through the other. What is the ratio of the brightest to the darkest fringe there?
- 3 to 1
- 9 to 1
+ 4 to 1
= In units of the weaker amplitude, the bright fringe is (3 + 1)² = 16 and the dark one (3 − 1)² = 4, a ratio of 4. The fringes survive but no longer reach zero.
:::
