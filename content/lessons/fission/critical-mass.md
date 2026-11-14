Whether a chain reaction grows or dies depends on k, and k depends on how many neutrons escape before they can cause a fission. Escape depends on the size, shape and density of the material. This is the origin of **critical mass**: the smallest amount of a fissile material (one that can sustain a chain reaction) that can keep a chain going in a given arrangement.

## Neutrons are lost through the surface

A fast neutron in uranium metal travels several centimetres, on average, between interactions. In a small lump, many neutrons reach the surface and leave before they hit anything. In a large lump, a neutron born near the middle is likely to cause a fission before it gets out.

The key comparison is between where neutrons are produced, throughout the volume, and where they are lost, through the surface.

::: steps Why bigger lumps lose a smaller fraction
$ \text{production} \sim V = \frac{4}{3}\pi r^3
Fissions, and the neutrons they release, happen everywhere inside the material, so neutron production grows with volume.
---
$ \text{leakage} \sim A = 4\pi r^2
Neutrons leave only by crossing the surface, so the number escaping grows roughly with surface area.
---
$ \text{leakage} \div \text{production} \sim \frac{A}{V} = \frac{3}{r}
The fraction lost falls as the sphere grows. Double the radius and the lost fraction roughly halves, while the mass rises eightfold.
---
$ k \text{ rises with } r\text{, and } k = 1 \text{ at the critical radius}
Below some radius too many neutrons escape and k < 1. At the critical radius k = 1. The mass of that sphere is the critical mass.
:::

This is the same reason a large animal loses heat more slowly than a small one: heat is made in the volume and lost through the skin.

::: model critical-mass
Grow the sphere and watch the fraction of neutrons that escape fall until the chain sustains itself.
:::

## Three regimes

- **Subcritical** (k < 1): any chain started by a stray neutron fades out after a few generations. A subcritical piece of uranium is safe to handle, apart from its mild natural radioactivity.
- **Critical** (k = 1): a chain, once started, continues at a constant rate.
- **Supercritical** (k > 1): the chain grows exponentially. A weapon needs to be strongly supercritical, with k well above 1, so that many generations pass before the material blows apart.

Critical mass is not a fixed property of a material. It depends on the material's purity (enrichment), its density, its shape and what surrounds it.

## Shape

A sphere has the least surface area for a given volume, so it leaks the fewest neutrons and has the smallest critical mass. The same material spread into a thin sheet or a long rod can be far above the spherical critical mass and still be subcritical, because neutrons escape easily through its large surface. Conversely, reshaping a safe arrangement into a compact one can make it critical. Nuclear facilities control the shape of containers for exactly this reason.

## Density

Density matters even more. A neutron's chance of hitting a nucleus depends on how many nuclei lie along its path. The average distance it travels before an interaction, the **mean free path**, is inversely proportional to density.

::: steps How critical mass depends on density
$ \lambda = \frac{1}{n\sigma}
The mean free path λ is set by the number of nuclei per unit volume, n, and the effective target area each presents, σ.
---
$ R_\text{critical} \sim \lambda \sim \frac{1}{\rho}
All the relevant distances scale with the mean free path. Double the density ρ and every distance a neutron covers before interacting halves, so the critical radius halves too.
---
$ M_\text{critical} = \rho \times \frac{4}{3}\pi R^3 \sim \rho \times \left(\frac{1}{\rho}\right)^3 = \frac{1}{\rho^2}
The critical mass is density times volume. Doubling the density cuts it to about a quarter.
:::

This result is the physical basis of the implosion method described in the two-designs lesson: a piece of material that is subcritical at normal density can become supercritical without adding any material, simply by being squeezed.

::: note
Supercriticality can be reached in two ways: by bringing more material together into a compact shape, or by compressing a fixed amount of material to higher density. The two bombs of 1945 used one method each.
:::

## Reflectors

Surrounding the material with a layer of another substance that scatters neutrons well, called a **reflector** or tamper, sends some escaping neutrons back in and lowers the critical mass.

::: history 1946 · Los Alamos
Critical mass is unforgiving. On 21 August 1945 Harry Daghlian, and on 21 May 1946 Louis Slotin, were conducting experiments that brought pieces of the same plutonium core close to criticality by surrounding it with reflecting material. Each time, a slip made the assembly briefly supercritical. There was no explosion, because the heat and the operator's reaction pushed the system apart almost at once, but each received a fatal dose of radiation. The core became known as the "demon core". The accidents ended hands-on criticality experiments at Los Alamos.
:::

::: history 1940 · Birmingham, England
Otto Frisch and Rudolf Peierls made the first rough calculation of the critical mass of pure U-235 using fast neutrons. Earlier estimates, which assumed natural uranium, had suggested tons. Their much smaller figure showed that a weapon could be carried by aircraft, and it changed British and then American policy.
:::

::: quiz
? Why does a larger sphere of fissile material lose a smaller fraction of its neutrons?
- Neutrons move more slowly in larger spheres
+ Production grows with volume (r³) but loss grows with surface (r²)
- Larger spheres contain more U-238
= The fraction lost scales like A/V = 3/r, so it falls as r grows.
? If a subcritical sphere is compressed to twice its density, its critical mass becomes about
- twice as large
- half as large
+ a quarter as large
= Critical mass scales as 1/ρ², so doubling ρ divides it by four.
? Which shape has the smallest critical mass?
+ A sphere
- A thin flat sheet
- A long thin rod
= A sphere has the least surface per volume, so it leaks the fewest neutrons.
:::
