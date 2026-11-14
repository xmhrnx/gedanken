Astronomers cannot put a galaxy on a scale. They weigh things in space by watching how they move, because gravity sets the motion and mass sets the gravity. A planet's orbit gives the mass of the Sun; the orbits of stars near the centre of the Milky Way revealed the mass of the black hole there. Applied to whole galaxies and clusters of galaxies, the same method gives a startling answer: they contain far more mass than all their stars and gas put together, several times more, in a form that gives off no light at all. This lesson follows how that conclusion was reached, why it took forty years to be believed, why astronomers now think it means a new kind of matter rather than a flaw in the law of gravity, and what that matter might be.

## Zwicky's missing mass

::: history 1933 · Pasadena, California
Fritz Zwicky, a Swiss astronomer at Caltech known for bold ideas and a combative temper, measured the speeds of galaxies in the Coma Cluster, a swarm of about a thousand galaxies some 100 megaparsecs away (a megaparsec is about 3.26 million light-years). The galaxies were moving around within the cluster at about 1,000 kilometres per second. That is fast enough that, if the cluster's only mass were its visible galaxies, their gravity could not hold them; the cluster should be flying apart. Zwicky calculated that holding it together needed hundreds of times more mass than the galaxies' light suggested. He called the missing ingredient *dunkle Materie*, dark matter. His factor was exaggerated, partly because he used the badly overestimated Hubble constant of the day, and decades later much of the shortfall was found as hot gas glowing in X-rays, invisible to his telescopes. But even after both corrections, a large gap remained.
:::

For about forty years the result was largely set aside. Clusters of galaxies are complicated systems, and no one was sure they were stable rather than slowly dispersing, in which case Zwicky's argument would not apply. The case that convinced most astronomers came instead from single galaxies, where the measurement is cleaner.

## Why should outer stars move slowly?

A spiral galaxy is a flat rotating disc of stars and gas. Its **rotation curve** is a graph of how fast material orbits the centre at each distance. To measure it, astronomers use the Doppler effect: on one side of a tilted disc, stars and gas are turning toward us and their light is blueshifted, and on the other side they are turning away and their light is redshifted. The size of the shift at each point gives the orbital speed there.

Before looking, it is worth asking what Newton's gravity predicts. The answer is simple enough to work out directly.

::: steps The flat-rotation-curve argument
$ \frac{v^2}{r} = \frac{G M(r)}{r^2}
A star on a circular orbit of radius r is constantly being turned toward the centre, which needs an acceleration of v²/r. Gravity supplies it. Here M(r) is the mass inside radius r, and G is Newton's gravitational constant. Only the inner mass counts because, for a roughly spherical distribution, Newton's shell theorem shows that the pull of mass outside the orbit cancels in all directions. A flat disc needs a correction of a few tens of percent, which changes the numbers but not the conclusion.
---
$ v = \sqrt{\frac{G M(r)}{r}}
Multiply both sides by r and take the square root, to get the orbital speed in terms of the enclosed mass. This is the general rule for circular orbits: speed depends on how much mass lies inside the orbit and how far out the orbit is.
---
$ M(r) \approx \text{constant} \Rightarrow v \propto \frac{1}{\sqrt{r}}
Now apply it to a galaxy whose mass follows its light. Most of a spiral's light comes from near its centre and fades steadily outward, so beyond the bright inner disc there is little more mass to add, and M(r) stops growing. With M(r) constant, speed should fall as one over the square root of distance, just as the outer planets move more slowly around the Sun than the inner ones: Neptune, about 80 times farther out than Mercury, orbits at about a ninth of its speed.
---
$ v \approx \text{constant} \Rightarrow M(r) = \frac{v^2 r}{G} \propto r
Instead, observations show that the speeds stay roughly constant far out, a **flat rotation curve**. Solve the first equation for M(r) and the meaning is clear: if v does not change, the enclosed mass must keep growing in direct proportion to r, far beyond the point where the light runs out. Something with mass, but without light, must extend well past the visible galaxy.
---
$ M = \frac{(2.2 \times 10^5\,\text{m/s})^2 \times 2.5 \times 10^{20}\,\text{m}}{6.67 \times 10^{-11}\,\text{m}^3\,\text{kg}^{-1}\,\text{s}^{-2}} \approx 1.8 \times 10^{41}\,\text{kg}
Put in numbers for our own galaxy. The Sun orbits at about 220 km/s at a distance from the centre of about 8 kiloparsecs, or 2.5 × 10²⁰ m. The mass inside the Sun's orbit comes out at about 1.8 × 10⁴¹ kg, or about 9 × 10¹⁰ times the mass of the Sun. Much of that inner mass is stars and gas, but the Milky Way's rotation curve stays flat far beyond the Sun, where there are few stars, so the total mass keeps climbing well past what the visible galaxy can supply.
:::

::: history 1970s · Washington, D.C.
Vera Rubin, working with the instrument builder Kent Ford at the Carnegie Institution, used a new and very sensitive spectrograph to measure the rotation of the Andromeda galaxy in 1970, and then of dozens of other spirals through the decade. She had expected to see the speeds fall off toward the edge of the disc. They did not; across galaxy after galaxy the curves stayed flat. At about the same time radio astronomers, notably Morton Roberts and Albert Bosma, used the radio emission of hydrogen gas, which extends far beyond the visible stars, to track rotation even farther out, and found the same. Theorists such as Jeremiah Ostriker and James Peebles were meanwhile showing that a flat disc of stars on its own would be unstable, and would be steadied by a massive surrounding halo. By the end of the decade most astronomers accepted that the missing mass was real.
:::

The picture that emerged is that each galaxy sits inside a much larger, roughly spherical cloud of unseen matter, called a **halo**, which outweighs its stars several times over. The visible galaxy is like the bright core of something much bigger.

## Can we see dark matter's gravity directly?

Rotation curves infer mass from motion, and they assume the law of gravity is right. A second method avoids motion entirely. General relativity predicts that mass bends the path of light, so a massive cluster of galaxies acts as a lens, distorting the images of more distant galaxies behind it into arcs and stretched shapes. Measuring those distortions maps the total mass of the cluster, visible or not, without needing to know how anything in the cluster moves. This is called **gravitational lensing**. Lensing masses agree with the masses found from motion: clusters contain about five to six times as much dark matter as ordinary matter.

::: history 2006 · the Bullet Cluster
Douglas Clowe and his colleagues studied an object catalogued as 1E 0657-56, nicknamed the Bullet Cluster: two clusters of galaxies that passed through each other about 150 million years ago. Most of the ordinary matter in a cluster is not in its galaxies but in hot gas between them, which glows in X-rays and was mapped by NASA's Chandra X-ray telescope. When the clusters collided, their gas clouds crashed into each other, were slowed down and were left behind in the middle. The galaxies, which are small and far apart, sailed through each other almost without touching. The lensing map then showed where the mass is: not in the middle with the gas, which holds most of the ordinary matter, but out with the two groups of galaxies, which had passed straight through.
:::

That separation is the important point. Some physicists have proposed that there is no dark matter, and that instead gravity behaves differently at very low accelerations; the best known of these theories is **MOND**, for Modified Newtonian Dynamics, put forward by Mordehai Milgrom in 1983. MOND can reproduce flat rotation curves well. But in such a theory the extra gravity should follow the ordinary matter, and in the Bullet Cluster most of the ordinary matter is in the gas while most of the gravity is elsewhere. That is very hard to explain by changing the law of gravity, and natural if most of the mass is made of particles that barely interact, with ordinary matter or even with each other, so that two clouds of it pass through each other unhindered.

## Why the early universe needs it too

A third line of evidence comes from the history of the universe as a whole. The cosmic microwave background shows that 380,000 years after the Big Bang, the matter in the universe was smooth to about one part in 100,000. Today it is gathered into galaxies, clusters and vast empty voids. Gravity turns small irregularities into large ones, because a region slightly denser than average pulls in more matter and becomes denser still, but the process takes time, and the question is whether there has been enough.

With ordinary matter alone, there has not. Before the universe became transparent, ordinary matter was tied to the light that filled it; the pressure of that light pushed back against any clump that began to form, so the irregularities in ordinary matter could not grow. Starting from one part in 100,000 at that point, they would not have had time to become galaxies by today. Dark matter does not interact with light, so its clumps could start growing much earlier, and when the ordinary matter was finally released it fell into the dark matter's ready-made wells. Computer simulations that start from the measured ripples and include dark matter produce a web of galaxies and clusters that looks very much like the real one. Without it, they produce a universe far smoother than the one we live in.

## What is it made of?

The first possibility to rule out is ordinary matter that happens to be dark: dim stars, cold gas, black holes formed from stars. Two independent measurements say this cannot be the answer. The amount of deuterium made in the first minutes after the Big Bang depends on the density of ordinary matter at that time, and the pattern of ripples in the cosmic microwave background depends on it too. Both give the same amount, and it is only about one sixth of the total matter the motions and lensing require. So most dark matter is not made of atoms at all.

Neutrinos, the one known particle that is both plentiful and barely interacting, were an early candidate. But they are so light that in the early universe they moved at nearly the speed of light, which would have smoothed out small clumps of matter instead of letting them grow into galaxies. The structure of the universe needs dark matter that was slow, or "cold", early on.

The leading candidates are therefore new particles not in the Standard Model. One family is **weakly interacting massive particles**, or WIMPs, heavy particles that would interact with ordinary matter only through the weak force or something similar. The other is **axions**, extremely light particles first proposed for a different reason, to solve a puzzle in the theory of the strong force. Large detectors deep underground, shielded from cosmic rays, have searched for decades for the rare collisions of WIMPs with atomic nuclei, and other experiments look for axions turning into light in strong magnetic fields. None has found anything yet. Each new search rules out more of the simplest possibilities, and the identity of most of the matter in the universe remains unknown.

::: note
Stars far from a galaxy's centre orbit as fast as those near it. By v = √(GM(r)/r), that means the enclosed mass keeps growing where there is no light. Gravitational lensing and the Bullet Cluster confirm that most of the matter in the universe is invisible, and measurements from the early universe show it is not made of ordinary atoms.
:::

::: quiz
? If all of a galaxy's mass sat in its bright centre, how would orbital speed change far out?
- It would stay constant
+ It would fall as 1/√r
- It would rise in proportion to r
= With the enclosed mass M constant, v = √(GM/r) falls as 1/√r, just as the outer planets orbit the Sun more slowly than the inner ones.
? A flat rotation curve, v constant, implies the enclosed mass grows
- as 1/r
- not at all
+ in proportion to r
= Solving for the mass gives M(r) = v²r/G. If v stays the same, M must grow in direct proportion to r.
? Why is the Bullet Cluster strong evidence for dark matter rather than modified gravity?
- Its galaxies move faster than light
+ The mass found by lensing is separated from the hot gas that holds most of the ordinary matter
- It emits no X-rays
= If gravity simply followed the visible matter, the lensing map would centre on the gas, where most of the ordinary matter is. Instead the mass lies with the galaxies, as expected if it is made of particles that passed through the collision without interacting.
:::
