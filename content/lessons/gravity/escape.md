Throw a ball upward and it comes back down. Throw it harder and it rises higher before returning. Is there a speed at which it never comes back at all? At first sight the answer might seem to be no: gravity reaches out without limit, so however far the ball climbs, it is still being pulled back. But gravity also weakens with distance, and it weakens fast enough that the total effort needed to climb all the way out is finite. A ball launched with enough speed can pay that cost and leave for good. This lesson works out what that speed is, and then what it does and does not mean.

## Why mgh is not enough

The usual way to work out the energy needed to lift something is to multiply its weight by the height: lifting a mass m through a height h takes work mgh, where g is the acceleration due to gravity. **Work** here means energy transferred by a force acting over a distance: force times distance moved in the direction of the force.

That formula assumes the force of gravity is the same all the way up, which is an excellent approximation for climbing a building and a poor one for a trip into space. At one Earth radius above the surface, twice as far from Earth's centre, gravity has dropped to a quarter of its surface strength. If we used mgh with h going to infinity, we would get an infinite amount of work, which is clearly wrong, because the force we are working against keeps getting weaker.

The right approach is to split the climb into many small steps, work out the force over each step, and add up force times distance for all of them. Adding up a quantity that changes continuously, over infinitely many tiny steps, is exactly what an integral does.

::: steps The work needed to escape from the surface
$ F(r) = \frac{GMm}{r^2}
This is the gravitational force on a mass m at a distance r from the centre of a planet of mass M, from Newton's law of gravitation. To lift the mass steadily you must push outward with a force this large. Because r is in the denominator, the force you need gets smaller as you rise.
---
$ W = \int_R^{\infty} \frac{GMm}{r^2}\,\mathrm{d}r
Over a short step dr the force hardly changes, so the work done on that step is F dr. The integral adds up all these small amounts, starting at the planet's surface, r = R, and continuing outward without limit. Writing the total this way is allowed because work over a path is just the sum of the work over its pieces.
---
$ W = GMm\left[-\frac{1}{r}\right]_R^{\infty} = GMm\left(0 + \frac{1}{R}\right)
To evaluate the integral we need a function whose slope is 1/r², and −1/r is one: as r increases, −1/r rises toward zero, and its rate of rise is exactly 1/r². The integral is then the value of −1/r at the far end minus its value at the near end. At infinity −1/r is zero, and at the surface it is −1/R, so the difference is 0 − (−1/R) = 1/R.
---
$ W = \frac{GMm}{R}
The total work is finite. The climb never ends, but most of the work is done close to the planet, where gravity is strongest, and the force falls off so quickly farther out that the infinite remainder of the journey adds only a finite amount. This is the energy needed to take a mass m from the surface to "infinity", which in practice means far enough that the planet's pull no longer matters.
:::

This result explains a convention that often looks strange at first. The gravitational **potential energy** of a mass m at distance r from a planet, the energy stored by its position, is written U = −GMm/r. The zero is placed infinitely far away, where the planet's pull vanishes, and at any finite distance the mass is GMm/r short of that zero: that is how much work it would take to free it. Potential energy is negative because a bound object has less energy than a free one. Close to the surface, raising the mass by a small height h changes U by very nearly mgh, so the familiar formula is still correct where it was always used.

## The escape velocity

Now suppose a projectile is launched from the surface and then left to coast, with no engine and no air resistance. After launch its only source of energy is the kinetic energy, the energy of motion, it was given at the start. As it climbs, that kinetic energy is converted into potential energy and it slows down. If the kinetic energy is less than GMm/R, the projectile runs out of speed at some finite height and falls back. If it is at least GMm/R, it never runs out. The dividing line is the escape velocity.

::: steps Escape velocity
$ \tfrac{1}{2}mv^2 = \frac{GMm}{R}
A projectile that only just escapes spends all its kinetic energy, ½mv², on the climb and arrives at infinity with zero speed. We can set the two sides equal because gravity conserves energy: the energy lost in climbing does not leak away as heat but is fully accounted for by the change in potential energy. The left side is what the projectile starts with; the right side is the price of escape.
---
$ v^2 = \frac{2GM}{R}
Divide both sides by m and multiply by 2. The projectile's mass has dropped out, because both its kinetic energy and the work needed to lift it are proportional to its mass. A heavier object needs more energy to escape, but it also carries more energy at the same speed, so the required speed is the same.
---
$ v_{\text{esc}} = \sqrt{\frac{2GM}{R}} = \sqrt{2gR}
Take the square root to find the speed. The second form comes from using surface gravity, g = GM/R², to replace GM with gR². It is convenient because g and R are both easy to measure, and it shows that escape velocity depends only on how strong gravity is at the surface and how big the planet is.
:::

::: note
To escape for good, a coasting object needs a speed of √(2GM/R), whatever its mass: about 11.2 km/s from Earth and 2.4 km/s from the Moon. The climb is endless, but the work it takes is finite, because gravity fades with distance fast enough.
:::

## Numbers for Earth and the Moon

For Earth, g = 9.81 m/s² and R = 6.37 × 10⁶ m, so √(2gR) = √(2 × 9.81 × 6.37 × 10⁶) ≈ 11.2 km/s. That is about 40,000 km/h, or more than thirty times the speed of sound in air.

The Moon has about 1/81 of Earth's mass and 0.27 of its radius (7.35 × 10²² kg and 1,737 km), and putting those into √(2GM/R) gives about 2.4 km/s, less than a quarter of Earth's figure. Since the energy needed goes as the square of the speed, leaving the Moon takes only about one twentieth of the energy per kilogram that leaving Earth does. The Apollo missions show what that means. The Saturn V rocket that lifted the astronauts from Earth stood 110 metres tall and weighed nearly 3,000 tonnes at launch. The ascent stage that lifted two of them off the Moon and back into lunar orbit was a small cabin with a single engine, a few metres high.

It is useful to compare escape speed with the circular orbit speed from the previous lesson, √(GM/r). At the same distance from the planet, escape speed is larger by a factor of √2, about 1.41. A spacecraft already in low Earth orbit at 7.7 km/s therefore needs only about 3.2 km/s more to escape, a much smaller step than getting into orbit in the first place. The Apollo spacecraft did roughly this. Their engine burn toward the Moon, made from a parking orbit a few hundred kilometres up, raised their speed to about 10.8 km/s, just short of escape speed at that height. They did not need to escape completely, because the Moon itself is held by Earth's gravity: a long ellipse reaching out to the Moon's distance was enough.

## What escape velocity does not mean

It is easy to picture escape velocity as a barrier, a speed that anything leaving Earth must reach. That picture is wrong. Escape velocity is the launch speed needed by a coasting object, one that gets all its speed at the start and then receives no further push. A rocket is not like that. It keeps firing as it climbs, adding energy along the way, and in principle a rocket with enough fuel could leave Earth at walking pace. What has to be paid is the energy GMm/R per object, not a particular speed. In practice, rockets bound for the Moon or the planets still reach close to escape speed fairly early, because a rocket that climbs slowly spends its fuel holding itself up against gravity instead of gaining speed.

Nor could you launch a payload simply by firing it from a gun at 11.2 km/s. In the dense lower atmosphere, air resistance at that speed would heat the projectile violently and slow it down. And accelerating to that speed in the length of any practical barrel would subject the contents to accelerations tens of thousands of times stronger than gravity, which would crush any passenger.

::: history 1865–1903 · Paris and Kaluga
In *From the Earth to the Moon*, published in 1865, Jules Verne sent three men to the Moon in a shell fired from an enormous cannon buried in Florida. He took care to give the shell the right muzzle speed, about 11 km/s, and quietly ignored the air and the crushing acceleration. Among his readers was a deaf Russian boy named Konstantin Tsiolkovsky, who grew up to be a provincial schoolteacher in Kaluga and spent his spare time thinking about how space travel might really work. He saw that the answer was not a gun but a rocket, which carries its own propellant and pushes against its exhaust rather than against air, so it can build up speed gradually and keep working in a vacuum. In a paper published in 1903 he worked out how a rocket's final speed depends on the speed of its exhaust and on the fraction of its starting mass that is fuel. The result, now called the rocket equation, showed that a single rocket would struggle to carry enough fuel to reach orbit, let alone escape speed. In later work Tsiolkovsky proposed the way round it: rockets built in stages, each dropping its empty tanks as it went. Every orbital rocket since has followed that design.
:::

## Why the Moon has no air

Escape velocity also explains why some worlds have atmospheres and others do not. The molecules of a gas move at random, at a range of speeds whose average grows with temperature. At any moment a small fraction of them, in the fast tail of the distribution, are moving much faster than average. Any molecule near the top of the atmosphere that is moving upward faster than the escape speed is lost to space.

On the Moon the daytime surface temperature rises above 100 °C, and the escape speed is only 2.4 km/s. Molecules of nitrogen or oxygen at that temperature average a few hundred metres per second, but enough of the fastest ones exceed 2.4 km/s that, over billions of years, any atmosphere the Moon might have had has leaked away, helped by the solar wind stripping gas from the top. Earth's escape speed is almost five times higher, high enough to hold onto nitrogen and oxygen. It is not high enough to keep the lightest gases, hydrogen and helium, which move fastest at a given temperature: they escape slowly even from Earth, which is why the air holds so little helium even though radioactive rocks release it all the time.

The same formula will reappear in the chapter on black holes. Squeeze a fixed mass into a smaller radius and its escape velocity rises, because R is in the denominator. Squeeze it far enough and the escape velocity reaches the speed of light. In 1783 the English clergyman John Michell made exactly that argument and concluded that the universe might contain stars so compact that even light could not leave them.

::: quiz
? A spacecraft escapes Earth by coasting from a launch speed just above 11.2 km/s. A craft with ten times the mass needs a launch speed of
- about 35 km/s, √10 times more
+ the same, about 11.2 km/s
- about 112 km/s
= Both the kinetic energy and the work needed to climb out are proportional to the craft's mass, so the mass cancels. The heavier craft needs more energy but reaches it at the same speed.
? Why can't you find the work needed to escape by using mgh with h going to infinity?
- Because g increases with height
+ Because g weakens with height, so mgh overestimates the work; the true total is finite
- Because the object's mass changes as it rises
= The force GMm/r² falls off with distance, so the work per metre keeps shrinking. Adding it all up from R to infinity gives the finite value GMm/R.
? A satellite circles in low orbit at speed v. Roughly what speed does it need to escape from that height?
- 2v
+ 1.41v
- v²
= Escape speed √(2GM/r) is √2 times the circular orbit speed √(GM/r) at the same distance, and √2 is about 1.41.
:::
