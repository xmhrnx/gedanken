By the mid-1920s physicists had accepted something strange: electrons, which arrive at a detector as single, whole particles, also behave like waves. Louis de Broglie had given every particle a wavelength, λ = h/p, where h is Planck's constant and p the particle's momentum, and beams of electrons were soon seen to diffract off crystals just as X-rays do. That left a question nobody could answer at first. If an electron is a wave, a wave of *what*? This lesson is about the answer that quantum mechanics settled on, the wavefunction, and the single rule that ties it to anything we can measure. Almost everything else in the subject is built on top of it.

## A wave of what?

The obvious first guess is that the electron itself is spread out, its mass and charge smeared through space the way the energy of a sound wave is spread through air. That guess fails the moment you look at a detector. A screen hit by an electron beam does not glow faintly all over; it records sharp, separate dots, one per electron. A Geiger counter clicks once, not a little. Whatever is doing the waving, the electron is never found in pieces.

So the wave cannot be the electron's substance. It has to be something that governs where the whole electron turns up, without itself being the electron. Quantum mechanics calls this thing the **wavefunction**, written ψ (the Greek letter psi), and treats it as the complete description of the particle's state.

## What the wavefunction is

The wavefunction ψ(x, t) assigns a number to every point x in space at every moment t. In that respect it is like the height of a water wave, which also has a value at each place and time. But the number is not an ordinary one. It is a **complex number**, which means it has two parts that are best pictured as a length and a direction on a dial. The length is called the size, |ψ|, and the direction is called the **phase**, an angle that says where in its cycle the wave is at that point.

Neither part can be observed directly. There is no instrument that reads off ψ at a point the way a thermometer reads temperature. The wavefunction earns its place in physics only through a rule that turns it into something we can count.

::: equation The Born rule
P(x \text{ to } x + \mathrm{d}x) = |\psi(x,t)|^2\,\mathrm{d}x
:::

In words: the probability of finding the particle in a small interval of width dx around the point x is the square of the wavefunction's size there, times the width of the interval. That makes |ψ|² a **probability density**, a probability per unit length, or per unit volume in three dimensions. Where |ψ|² is large the particle is likely to be found; where it is small the particle rarely turns up; and where ψ is exactly zero the particle is never found at all. The phase drops out of |ψ|², which is why it cannot be seen in a single position measurement, though, as we will see shortly, it is far from irrelevant.

There is one condition every wavefunction must meet. The particle exists, so if you search all of space you must find it somewhere, and the probabilities over all space must add up to exactly one:

::: equation Normalisation
\int_{-\infty}^{\infty} |\psi(x,t)|^2\,\mathrm{d}x = 1
:::

A wavefunction scaled to meet this condition is called normalised.

## Probabilities, not outcomes

It is worth pausing on how much the Born rule gives up. It does not say where a particular electron will land. Two electrons prepared in exactly the same way, with identical wavefunctions, can arrive at quite different places. What the rule predicts is the pattern: repeat the experiment many times and the density of hits across the screen traces out |ψ|². The theory is precise about the distribution and silent about the individual event.

That was a sharp break with the physics before it. In Newton's mechanics, if you know where a particle is and how it is moving, you know where it will be. Uncertainty about the outcome meant only that you had not measured carefully enough. In quantum mechanics, as far as anyone has been able to tell, the probability is not a cover for missing information. The wavefunction is everything there is to know, and it still only gives odds.

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

::: model wavefunction
A particle trapped in a box. Step through the levels and compare ψ with |ψ|². The wave ψ dips below zero; the probability never does. At the nodes, where ψ crosses zero, the particle is never found, even though it is found on both sides.
:::

That last point would be absurd for a classical particle bouncing between the walls, which must pass through every point to get from one side to the other. For a wave it is natural, and it is one more sign that the particle does not have a path in the ordinary sense.

::: note
The wavefunction is not a physical wave you can see. It is an amplitude: the square of its size gives the probability of finding the particle at each place. Because amplitudes add before they are squared, quantum probabilities can interfere and even cancel.
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
:::
