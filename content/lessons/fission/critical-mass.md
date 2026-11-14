The previous lesson showed that a chain reaction grows if k, the number of neutrons from each fission that go on to cause another, is greater than 1, and dies if it is less. But k is not a fixed property of a material. The same uranium can have k below 1 as a small lump and above 1 as a large one, because what happens to a neutron depends on whether it finds a nucleus before it finds the edge. This lesson explains why size matters, why shape and density matter just as much, and how this leads to the idea of a **critical mass**: the smallest amount of a fissile material (one that can sustain a chain reaction, such as uranium-235 or plutonium-239) that can keep a chain going in a given arrangement. It is the idea on which both the safety of handling fissile material and the design of the 1945 weapons depend.

## Neutrons are lost through the surface

Picture a neutron just released by a fission somewhere inside a lump of uranium metal. It travels in a straight line until it meets a nucleus, then either causes a fission, is captured, or bounces off in a new direction. A fast neutron in uranium metal travels several centimetres, on average, between such encounters. That distance is not small compared with a lump you could hold, and this is the heart of the matter.

In a small lump, many neutrons reach the surface and leave before they meet anything. They are lost to the chain. In a large lump, a neutron born near the middle has a long way to go before it reaches the edge, and it is likely to cause a fission on the way. So the fraction of neutrons lost depends on how big the lump is compared with the distance a neutron typically travels. The key comparison is between where neutrons are made and where they are lost.

::: steps Why bigger lumps lose a smaller fraction
$ \text{production} \propto V = \frac{4}{3}\pi r^3
Take a sphere of radius r. Fissions, and the neutrons they release, happen everywhere inside it, so the number of neutrons produced per second is proportional to the amount of material, which means to the volume. Doubling the radius multiplies the volume, and so the production, by eight.
---
$ \text{leakage} \propto S = 4\pi r^2
Neutrons can only leave by crossing the surface, so the number escaping per second grows roughly with the surface area, here written S. Doubling the radius multiplies the surface by only four. This is a simplification, since neutrons born near the surface are the most likely to escape, but it captures how the two quantities scale.
---
$ \frac{\text{leakage}}{\text{production}} \propto \frac{S}{V} = \frac{3}{r}
Divide one by the other to find the fraction of neutrons lost. The factors of 4π cancel and the ratio of surface to volume comes out as 3/r. As the sphere grows, this fraction shrinks: double the radius and the fraction lost roughly halves, while the mass of material rises eightfold.
---
$ k \text{ rises with } r\text{, and } k = 1 \text{ at the critical radius}
Since fewer neutrons are lost from a larger sphere, more of them cause fissions, and k rises steadily with size. Below some radius so many escape that k is less than 1. At one particular radius, the critical radius, k is exactly 1. The mass of a sphere of that radius is the critical mass.
:::

The same reasoning about surface and volume explains why a large animal keeps warm more easily than a small one. Heat is produced throughout the body's volume but lost through its skin, and a larger body has less skin for each kilogram. A mouse must eat constantly to stay warm; an elephant's problem is getting rid of heat. The analogy holds as long as one remembers that neutrons, unlike heat, do not spread smoothly: each one travels a definite distance before its fate is decided.

::: model critical-mass
Grow the sphere and watch the fraction of neutrons that escape fall. Notice that nothing changes about the material itself; only the size changes, until the chain sustains itself.
:::

## Three regimes

The value of k sorts any arrangement of fissile material into one of three states, and each has practical consequences.

A **subcritical** arrangement has k below 1. Stray neutrons are always present, from spontaneous fissions and cosmic rays, and occasionally one starts a short chain, but each generation is smaller than the last, and the chain fades out after a few steps. A subcritical piece of uranium can be handled safely, apart from its mild natural radioactivity.

A **critical** arrangement has k equal to 1. A chain, once started, keeps going at a steady rate, each fission replacing itself with exactly one more. That is how a reactor runs.

A **supercritical** arrangement has k above 1, and the chain grows exponentially. As the previous lesson showed, a weapon needs to be strongly supercritical, with k well above 1, so that many generations pass before the heat of the reaction blows the material apart and stops it.

Critical mass, then, is not a single number for a material, like its density or melting point. It depends on four things: the material's purity, meaning how much of it is the fissile isotope; its density; its shape; and what surrounds it. Each of these changes how many neutrons are lost, and so changes k.

## Shape

Of all shapes with a given volume, a sphere has the least surface area. It therefore leaks the fewest neutrons and has the smallest critical mass. Any other shape loses more.

That has a reassuring consequence. The same material spread out into a thin sheet or a long, narrow rod can be well above the spherical critical mass and still be safely subcritical, because neutrons escape easily through its large surface: from anywhere inside a thin sheet, the nearest surface is only a short distance away. It also has a dangerous one. Reshaping a safe arrangement into a more compact one, for instance by pouring a solution from a long thin pipe into a squat tank, can raise k past 1. Nuclear facilities control the shapes and sizes of containers for fissile material for exactly this reason, and several accidents have happened when a solution ended up in a container of the wrong shape.

## Density

Density matters even more, and the reason comes from thinking about the neutron's journey. Whether a neutron causes a fission before escaping depends on how many nuclei lie along its path. The average distance it travels before it meets a nucleus is called the **mean free path**. Pack the nuclei closer together and that distance shrinks; spread them out and it grows. Every length in the problem, including the critical radius, is measured against it.

::: steps How critical mass depends on density
$ \lambda = \frac{1}{n\sigma}
The mean free path λ depends on two things: n, the number of nuclei in each unit of volume, and σ, the effective target area that each nucleus presents to a passing neutron, called its cross-section. More nuclei, or bigger targets, mean a shorter distance between encounters. The relation is a simple inverse one.
---
$ R_\text{critical} \propto \lambda \propto \frac{1}{\rho}
The number of nuclei per unit volume is proportional to the density ρ, so the mean free path is inversely proportional to density. What makes a sphere critical is not its radius in centimetres but its radius measured in mean free paths, because that is what decides how likely a neutron is to escape. So if the density doubles, every mean free path halves, and a sphere with half the radius is now exactly as effective as the old one was.
---
$ M_\text{critical} = \rho \times \frac{4}{3}\pi R_\text{critical}^3
The critical mass is the density times the volume of the critical sphere. This expression has density in it twice: once directly, and once hidden inside the critical radius, which itself depends on density.
---
$ M_\text{critical} \propto \rho \times \left(\frac{1}{\rho}\right)^3 = \frac{1}{\rho^2}
Put in the result that the critical radius is proportional to 1/ρ. Cubing it gives 1/ρ³, and multiplying by ρ leaves 1/ρ². So doubling the density cuts the critical mass to about a quarter. The volume shrinks eightfold, which more than outweighs the doubled density.
:::

That result has a striking implication. A piece of material that is subcritical at its normal density can become supercritical without adding any material at all, simply by being squeezed. Its mass stays the same, but the critical mass falls below it. This is the physical basis of the implosion method described in the two-designs lesson.

::: note
A fissile material can be made supercritical in two ways: by bringing more of it together into a compact shape, or by compressing a fixed amount to a higher density. Both work by reducing the fraction of neutrons that escape. The two bombs of 1945 used one method each.
:::

## Reflectors

The last factor is what surrounds the material. If a fissile sphere is wrapped in a layer of another substance that scatters neutrons well, some of the neutrons that would have escaped bounce back in and get a second chance to cause a fission. Such a layer is called a **reflector**. It reduces the loss through the surface, so it lowers the critical mass. The effect also means that bringing almost any neutron-scattering material close to a nearly critical assembly, including a person's hands or body, can push it over the edge.

::: history 1945–1946 · Los Alamos
Critical mass is unforgiving, and two deaths at Los Alamos showed how unforgiving. Scientists there measured how close plutonium assemblies were to criticality by surrounding a core with reflecting material, a little at a time, and watching the neutron count rise. On 21 August 1945 Harry Daghlian, working late at night with only a security guard in the room, was stacking reflecting bricks around a plutonium core when one slipped from his hand and fell onto the pile. On 21 May 1946 Louis Slotin was demonstrating the same kind of measurement on the same core, holding a reflecting shell above it with a screwdriver, when the screwdriver slipped and the shell closed. Each time, the assembly became briefly supercritical and released an intense burst of neutrons and gamma rays. There was no explosion: the burst heated and expanded the core slightly, enough to drop k back down, and each man then pulled the assembly apart by hand. But each had received a lethal dose of radiation. Daghlian died 25 days later and Slotin after 9 days. The core became known as the "demon core", and after Slotin's death, hands-on criticality experiments at Los Alamos were stopped in favour of remotely controlled ones.
:::

The idea of a critical mass had already changed history once before those accidents, by showing that a weapon was possible at all.

::: history 1940 · Birmingham, England
Before 1940, estimates of the amount of uranium needed for an explosive chain reaction were huge, because they assumed natural uranium, where U-238 wastes most of the neutrons. Some put the figure at tons, too much for any aircraft to carry, and many physicists concluded that a bomb was impractical. Otto Frisch and Rudolf Peierls, two refugees working at the University of Birmingham, asked a different question: what if the material were pure U-235, with fast neutrons doing the work and no U-238 to absorb them? Their rough calculation of the critical mass gave a figure small enough for a bomb to be carried by a single aircraft. That one change of assumption turned the bomb from a fantasy into a possibility, and it changed British and then American policy.
:::

::: quiz
? Why does a larger sphere of fissile material lose a smaller fraction of its neutrons?
- Neutrons move more slowly in larger spheres
+ Production grows with volume (r³) but loss grows with surface (r²)
- Larger spheres contain more U-238
= Neutrons are made throughout the volume but escape only through the surface. The fraction lost scales like surface over volume, 3/r, so it falls as the sphere grows.
? If a subcritical sphere is compressed to twice its density, its critical mass becomes about
- twice as large
- half as large
+ a quarter as large
= Doubling the density halves the mean free path, so the critical radius halves and the critical volume falls eightfold. Multiplying by the doubled density gives a critical mass about a quarter as large, since it scales as 1/ρ².
? Which shape has the smallest critical mass?
+ A sphere
- A thin flat sheet
- A long thin rod
= A sphere has the least surface for its volume, so it leaks the fewest neutrons. Flattened or stretched shapes lose more neutrons and need more material to reach k = 1.
:::
