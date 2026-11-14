Throw a ball upward and it comes back. Throw it harder and it goes higher before returning. Is there a speed at which it never comes back at all? There is, because gravity weakens with distance: the total work needed to climb out of a planet's gravity is finite, even though the climb never quite ends.

## Why mgh is not enough

Near the ground, lifting a mass m through a height h takes work mgh. That formula assumes g is constant, which is fine for a building but not for a trip to space: at one Earth radius up, g has dropped to a quarter. To find the work needed to go all the way out, add up the force over each small step of the climb, which is exactly what an integral does.

::: steps The work needed to escape from the surface
$ F(r) = \frac{GMm}{r^2}
The force you must overcome at distance r from the centre of a planet of mass M. It weakens as you rise.
---
$ W = \int_R^{\infty} \frac{GMm}{r^2}\,\mathrm{d}r
Work is force times distance; with a changing force, add F dr over every small step dr from the surface, r = R, out to infinity.
---
$ W = GMm\left[-\frac{1}{r}\right]_R^{\infty} = GMm\left(0 + \frac{1}{R}\right)
The function −1/r has slope 1/r², so it is the antiderivative. Evaluate it at the two ends: −1/r is zero at infinity and −1/R at the surface.
---
$ W = \frac{GMm}{R}
The total is finite. Gravity is strongest near the surface and fades so fast farther out that the endless remainder of the climb adds up to a fixed amount.
:::

This is why the gravitational potential energy is written U = −GMm/r: it is zero infinitely far away, and at distance r you are GMm/r short of that. Close to the surface it changes by mgh per height h, matching the familiar formula.

## The escape velocity

A projectile launched from the surface and then left to coast (no engine, no air) has only its kinetic energy to pay for the climb. It escapes if that energy covers the work W.

::: steps Escape velocity
$ \tfrac{1}{2}mv^2 = \frac{GMm}{R}
The barely-escaping projectile spends all its kinetic energy on the climb and arrives "at infinity" with zero speed. Energy is conserved because gravity is conservative.
---
$ v^2 = \frac{2GM}{R}
Cancel m and multiply both sides by 2. The projectile's mass has dropped out.
---
$ v_{\text{esc}} = \sqrt{\frac{2GM}{R}} = \sqrt{2gR}
Take the square root. The second form uses g = GM/R², handy for a quick estimate.
:::

::: note
To escape for good, a coasting object needs v = √(2GM/R): about 11.2 km/s from Earth and 2.4 km/s from the Moon, whatever its mass.
:::

## Numbers for Earth and Moon

For Earth, √(2gR) = √(2 × 9.81 × 6.37 × 10⁶) ≈ 11.2 km/s, about 40,000 km/h. The Moon has 1/81 of Earth's mass and 0.27 of its radius (7.35 × 10²² kg, 1,737 km), which gives about 2.4 km/s. That difference shaped the Apollo missions: the Saturn V that left Earth stood 110 metres tall, while the lunar module's small ascent stage was enough to lift two astronauts off the Moon into lunar orbit.

Compare circular orbit speed, √(GM/r). Escape speed is √2 times larger, so a spacecraft already in low orbit at 7.7 km/s needs only another 3.2 km/s or so to leave. Apollo's translunar burn took the craft to about 10.8 km/s at a height of a few hundred kilometres, just below escape speed there: the Moon is still bound to Earth, so a long ellipse was enough to reach it.

## What escape velocity does not mean

Escape velocity is the speed for a *coasting* projectile, one given all its speed at launch. A rocket keeps pushing as it climbs, so in principle it could leave Earth at walking pace, given enough fuel. In practice rockets bound for the Moon or the planets still build up close to escape speed early on, because climbing slowly against gravity wastes fuel, but the speed is a property of the energy budget, not a barrier at a particular height.

Nor can you simply fire a payload at 11.2 km/s from a cannon on the ground. The air at that speed would heat and slow it violently, and the acceleration in the barrel would crush any passenger. Jules Verne's moon gun of 1865 fails on both counts.

Escape velocity also explains why the Moon has no atmosphere. Gas molecules move at random speeds that grow with temperature. At the Moon's daytime temperature, above 100 °C, the fastest molecules in the mix exceed 2.4 km/s, and over billions of years any atmosphere leaks away. Earth's higher escape speed holds nitrogen and oxygen, though it lets hydrogen and helium slowly escape.

::: history 1865–1903 · Paris and Kaluga
Verne's *From the Earth to the Moon* used a cannon with the right muzzle speed, about 11 km/s, and ignored the rest. Konstantin Tsiolkovsky, a Russian schoolteacher who read Verne as a boy, saw that a rocket, which carries its own propellant and needs no air to push against, could reach such speeds gradually. His 1903 paper worked out how a rocket's final speed depends on exhaust speed and fuel load, and why multistage rockets are needed to reach orbit and escape.
:::

The same formula reappears in the chapter on black holes. Push a mass into a small enough radius and its escape velocity reaches the speed of light; John Michell made exactly that argument in 1783.

::: quiz
? A spacecraft escapes Earth by coasting from a launch speed slightly above 11.2 km/s. A craft with ten times the mass needs a launch speed of
- about 35 km/s (√10 times more)
+ the same, about 11.2 km/s
- about 112 km/s
= Both kinetic energy and the work against gravity are proportional to m, so the mass cancels.
? Why can't you find the work to escape by using mgh with h going to infinity?
- Because g increases with height
+ Because g weakens with height, so mgh overestimates the work, and in fact the true total is finite
- Because mass changes as you rise
= The force GMm/r² falls off with distance, and its integral from R to infinity is the finite value GMm/R.
? A satellite circles in low orbit at speed v. Roughly what speed does it need to escape from that height?
- 2v
+ 1.41v
- v²
= Escape speed √(2GM/r) is √2 times circular speed √(GM/r) at the same radius.
:::
