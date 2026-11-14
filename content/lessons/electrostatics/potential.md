A battery is labelled 1.5 V, a wall socket 230 V or 120 V, a lightning strike perhaps a hundred million volts. Voltage is the most quoted number in electricity, and it has a precise meaning: the energy each coulomb of charge gains or loses in moving from one place to another.

## Energy in a field

Lifting a book takes work against gravity, and that work is stored as gravitational potential energy, which comes back when the book falls. Pushing a positive charge against an electric field is exactly analogous. The electric force is conservative: the work done depends only on the start and end points, not the path. So every position can be given an electric potential energy U.

U depends on the charge being moved; a charge twice as big needs twice the work. Dividing by the charge gives a quantity that depends only on the location.

::: equation Electric potential
V = \frac{U}{q}, \qquad 1\,\text{V} = 1\,\text{J/C}
:::

V is the **electric potential**, and the difference ΔV between two points is the **potential difference**, or **voltage**. Only differences matter, just as only differences in height matter for a falling book; where you put zero is a choice. Moving a charge q across a voltage ΔV changes its energy by qΔV.

## Uniform fields

Between two large flat parallel plates with opposite charges, the field is nearly uniform: the same size and direction everywhere. This is the simplest case to work out.

::: steps Voltage across a uniform field
$ F = qE
A positive charge q between the plates feels a constant force qE, pointing from the positive plate toward the negative one.
---
$ W = Fd = qEd
To carry the charge a distance d from the negative plate to the positive one, you must push against this force. Work is force times distance when the force is constant and along the path.
---
$ \Delta U = qEd
The work you do is stored as potential energy, as with a lifted book.
---
$ \Delta V = \frac{\Delta U}{q} = Ed
Divide by q. The voltage between the plates is field times separation.
:::

Turned around, E = ΔV/d, so the field can be measured in volts per metre, the same unit as N/C. A 9 V battery connected across plates 1 mm apart gives a field of 9000 V/m. Dry air breaks down and sparks at about 3 × 10⁶ V/m, so a 1 cm spark needs roughly 30,000 V.

Field and potential are linked in general: the field points in the direction in which V falls fastest, and its size is how fast V falls per metre. A positive charge, left alone, moves from high potential to low, like a ball rolling downhill. Electrons, being negative, move from low to high.

## Point charges and equipotentials

For a single point charge Q, the potential at distance r (taking V = 0 infinitely far away) is

::: equation Potential of a point charge
V = k\frac{Q}{r}
:::

It falls as 1/r, more slowly than the field's 1/r², for the same reason gravitational potential energy goes as 1/r while the force goes as 1/r².

Surfaces of equal potential are called **equipotentials**. Moving a charge along one takes no work, so the field has no component along it: field lines always cross equipotentials at right angles. Around a point charge the equipotentials are spheres. The surface of a conductor in equilibrium is always an equipotential.

::: model field-lines
Field lines run downhill in potential, from high V to low. Picture the equipotentials as curves crossing every line at right angles, bunched where the lines are dense.
:::

## The electron-volt

On the atomic scale the joule is unwieldy. An electron accelerated through 1 V gains energy e × 1 V = 1.602 × 10⁻¹⁹ J, called one **electron-volt** (eV). Chemical bonds are a few eV; visible light photons carry 2–3 eV; the electron beam in an old television tube about 25,000 eV.

::: note
Voltage is energy per unit charge. A charge q crossing a potential difference ΔV gains or loses energy qΔV, whatever path it takes.
:::

::: history 1800 · Como, Italy
In the 1780s Luigi Galvani made dissected frogs' legs twitch by touching them with two different metals, and concluded that the electricity came from the animal. Alessandro Volta argued that it came from the contact between the dissimilar metals, with the moist tissue merely conducting. To prove it he stacked discs of zinc and copper (or silver), separated by cardboard soaked in brine, into a "pile". It produced a steady electrical push, and the more pairs in the stack, the stronger the shock. Announced in a letter to the Royal Society in 1800, the voltaic pile was the first battery. The unit of potential is named after him. Galvani was not entirely wrong: nerves and muscles do work electrically, though not in the way he thought.
:::

A battery is a device that keeps a fixed potential difference between its terminals by chemical reactions. In a 1.5 V cell, every coulomb that passes through the external circuit from one terminal to the other delivers 1.5 J.

::: quiz
? An electron moves through a potential difference of 100 V. Its energy changes by
- 100 J
+ 100 eV, about 1.6 × 10⁻¹⁷ J
- 1.6 × 10⁻¹⁹ J
= ΔU = qΔV = (1.6 × 10⁻¹⁹ C)(100 V) = 1.6 × 10⁻¹⁷ J, which is 100 eV by definition.
? Two plates 2 cm apart have 400 V between them. The field between them is
- 8 V/m
- 200 V/m
+ 20,000 V/m
= E = ΔV/d = 400 V / 0.02 m = 20,000 V/m.
? How much work does it take to move a charge along an equipotential surface?
+ None
- qV, where V is the potential of the surface
- It depends on the path
= On an equipotential V is the same everywhere, so ΔV = 0 and the energy change qΔV is zero.
:::
