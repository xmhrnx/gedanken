A torch battery is labelled 1.5 V. A wall socket supplies 230 V in Europe and 120 V in North America. A lightning strike may involve a hundred million volts or more. Voltage is the most quoted number in electricity, and yet many people who use it every day would struggle to say what it measures. It has a precise meaning: voltage is the energy that each coulomb of charge gains or loses in moving from one place to another. This lesson builds up that idea from the work done against an electric field, connects it to the field itself, and shows why energy per charge is such a useful thing to know.

## Why can a position in a field store energy?

Start with a familiar case. Lifting a book onto a shelf takes work, because you must push upward against gravity. That work is not lost. It is stored as gravitational **potential energy**, energy the book has because of where it is, and it comes back as kinetic energy if the book falls. What makes this possible is a special property of gravity: the work needed to lift the book from the floor to the shelf depends only on the starting and ending heights, not on the route you take. Lift it straight up, or carry it round the room first, and the total work against gravity is the same. A force with this property is called **conservative**, and only for such a force can each position be given a definite potential energy.

The electric force between charges at rest is conservative too, for the same reason gravity is: it points along the line joining the two charges and its size depends only on the distance between them. Any work done moving sideways, around a charge, is zero, and only the change in distance counts, so the route cannot matter. So pushing a positive charge toward another positive charge, against their repulsion, stores energy in exactly the same way as lifting a book. Every position near a set of charges can be assigned an **electric potential energy** U, and the work done moving a charge from one point to another equals the change in U.

## What is voltage, then?

The potential energy depends on the charge being moved: a charge twice as large feels twice the force, so it takes twice the work to move it, and it gains twice the energy. That makes U a property of the charge and its position together. Dividing by the charge removes the dependence on the charge and leaves something that describes only the location.

::: equation Electric potential
V = \frac{U}{q}, \qquad 1\,\text{V} = 1\,\text{J/C}
:::

V is the **electric potential**: the potential energy per unit charge at a point. Its unit, the **volt**, is one joule per coulomb. The difference in potential between two points, ΔV, is called the **potential difference**, or the **voltage** between them. Moving a charge q from one point to another changes its energy by qΔV.

Only differences in potential have physical meaning. In the case of the book, it makes no difference whether you measure height from the floor or from sea level; the energy released by a fall depends only on how far the book drops. In the same way, where you choose V = 0 is a convention. What a battery's "1.5 V" means is that one terminal is 1.5 volts higher in potential than the other.

## What is the voltage across a uniform field?

The easiest arrangement to work out is a pair of large, flat, parallel metal plates with equal and opposite charges, close together. Between them the field is nearly **uniform**: the same size and direction at every point. That simplicity makes the work easy to calculate, because the force on a charge does not change as it moves.

::: steps Voltage across a uniform field
$ F = qE
A positive charge q placed between the plates feels a force of size qE, from the definition of the field. Because the field is uniform, the force is the same everywhere in the gap, and it points from the positive plate toward the negative one.
---
$ W = Fd = qEd
Now carry the charge a distance d from the negative plate to the positive one. You must push against the electric force all the way, and because the force is constant and points straight along the path, the work you do is simply force times distance. Had the force varied, we would have needed to add up small pieces, but here no such complication arises.
---
$ \Delta U = qEd
The electric force is conservative, so the work you did is not lost: it is stored as electric potential energy, just as the work of lifting a book is stored as gravitational potential energy. Let the charge go and it will be pushed back to the negative plate, gaining exactly this energy as kinetic energy.
---
$ \Delta V = \frac{\Delta U}{q} = Ed
Divide by the charge to turn energy into potential. The charge cancels, as it must, and the voltage between the plates is the field multiplied by the distance between them. This is the simplest link between field and voltage.
:::

Read the other way round, the result says E = ΔV/d. That shows the field can be measured in volts per metre, and since the field was defined in newtons per coulomb, the two units must be the same thing. Connect a 9 V battery across two plates 1 mm apart, and the field between them is 9 V / 0.001 m = 9000 V/m.

The same relation explains sparks. Dry air conducts very poorly until the field in it reaches about 3 × 10⁶ V/m. At that point, stray electrons are accelerated hard enough between collisions to knock electrons off air molecules, which are accelerated in turn, and an avalanche of charge carves a conducting path: a spark. So jumping a 1 cm gap between flat electrodes needs roughly 3 × 10⁶ × 0.01 = 30,000 V. The small spark you get from a door handle after walking across a carpet tells you that your body had charged to thousands of volts.

## How are field and potential related in general?

Between parallel plates, field and potential are related by a simple multiplication. In general the link is this: the field points in the direction in which the potential falls most steeply, and its size is how quickly the potential falls per metre in that direction. That is just what E = ΔV/d says for the uniform case, made local.

There is a helpful picture here, with one caution. Think of potential as height on a landscape. A positive charge released at rest moves from high potential to low, as a ball rolls downhill, gaining kinetic energy as it goes; the field at each point is the direction of steepest downhill slope. The picture breaks down for negative charges. An electron, having negative charge, is pushed the opposite way, so it moves from low potential toward high. That is why electrons in a wire flow toward the positive terminal of a battery.

## What is the potential around a point charge?

For a single point charge Q, the field is not uniform, so finding the potential requires adding up the work over a path along which the force keeps changing. The natural place to set the potential to zero is infinitely far away, where the charge has no influence.

::: steps Potential of a point charge
$ U(r) = -\int_{\infty}^{r} k\frac{Qq}{r'^2}\,\mathrm{d}r'
Bring a test charge q in from very far away to a distance r from Q, slowly enough that it gains no kinetic energy. At each distance r' along the way, Q pushes on it with Coulomb's force kQq/r'², and to move it inward you must supply an equal and opposite force. The work you do, added up from infinity to r, is the potential energy at r; the minus sign is there because your force is opposite to Coulomb's force on the charge.
---
$ U(r) = -kQq\left(-\frac{1}{r'}\right)_{\infty}^{r} = k\frac{Qq}{r}
The integral of 1/r'² is −1/r'. Evaluate it at the two ends: at infinity the term is zero, and at r it is −1/r, so the two minus signs combine into a plus. For two like charges the result is positive, which makes sense, since work had to be done to push them together against their repulsion.
---
$ V = \frac{U}{q} = k\frac{Q}{r}
Divide by the test charge to get the potential. As always, q cancels, leaving a quantity set only by the source charge and the distance. The potential is positive near a positive charge and negative near a negative one.
:::

The potential falls off as 1/r, more slowly than the field's 1/r². That is a general pattern: the potential is built up by adding the field over distance, and adding a 1/r² quantity from infinity gives a 1/r result. Gravitational potential energy and force around a planet follow the same pair of laws for the same reason.

Surfaces on which the potential has the same value everywhere are called **equipotentials**. Moving a charge along an equipotential changes its energy by qΔV = 0, so no work is done, which means the field can have no component along the surface. Field lines therefore always cross equipotentials at right angles. Around a single point charge the equipotentials are spheres centred on the charge, and the field lines are the radial lines crossing them. The surface of a conductor in equilibrium is always an equipotential, because, as the previous lesson showed, the field just outside it points straight out from the surface.

::: model field-lines
Field lines run downhill in potential, from high V to low. Picture the equipotentials as curves crossing every line at right angles. They bunch together where the lines are dense, because there the potential changes quickly over a short distance.
:::

## Why use the electron-volt?

On the scale of atoms, the joule is an inconveniently large unit. An electron moved through a potential difference of 1 V changes its energy by e × 1 V = 1.602 × 10⁻¹⁹ J, and that amount of energy is called one **electron-volt** (eV). It is a natural unit because so many processes in atoms involve single electrons and voltages of a few volts. Breaking a typical chemical bond takes a few eV, which is why chemical batteries produce voltages of one to a few volts. Photons of visible light carry roughly 2 to 3 eV, and the electron beam in an old television tube was accelerated through about 25,000 V, giving each electron 25,000 eV.

::: note
Voltage is energy per unit charge. A charge q moving through a potential difference ΔV gains or loses energy qΔV, whatever path it takes, because the electric force between charges at rest is conservative.
:::

## How was the first steady voltage made?

::: history 1780s–1800 · Bologna and Como
In the 1780s Luigi Galvani, an anatomist in Bologna, found that the legs of a dissected frog twitched when touched with two different metals connected together. He concluded that the electricity came from the animal itself, a special "animal electricity" stored in the nerves and muscles. Alessandro Volta, a physicist at the University of Pavia, repeated the experiments and came to a different view: the electricity came from the contact between the two dissimilar metals, and the moist frog tissue merely conducted it and responded to it.

To prove his point Volta needed to produce the effect with no animal at all. Around 1799 he stacked discs of zinc and copper (or silver) in alternating layers, separating each pair with cardboard soaked in salt water. The "pile" produced a steady electrical push, and the more pairs he stacked, the stronger the shock it gave. He announced it in a letter from Como to the Royal Society in London in 1800. Before the pile, experimenters had only brief sparks from friction machines and stored charge; now they had a steady current, and within weeks others were using it to split water into hydrogen and oxygen. The voltaic pile was the first battery, and the volt is named after its inventor. Galvani was not wholly wrong: nerves and muscles do work electrically, though not in the way he imagined.
:::

A modern battery works on the same principle. Chemical reactions inside it keep one terminal at a higher potential than the other by a fixed amount. In a 1.5 V cell, every coulomb of charge that flows through the external circuit, from one terminal round to the other, delivers 1.5 J of energy to whatever the circuit is powering: a bulb, a motor, a radio.

::: quiz
? An electron moves through a potential difference of 100 V. By how much does its energy change?
- 100 J
+ 100 eV, about 1.6 × 10⁻¹⁷ J
- 1.6 × 10⁻¹⁹ J
= The energy change is qΔV = (1.6 × 10⁻¹⁹ C)(100 V) = 1.6 × 10⁻¹⁷ J. Since one electron-volt is the energy an electron gains across 1 V, this is 100 eV by definition.
? Two parallel plates 2 cm apart have 400 V between them. What is the field between them?
- 8 V/m
- 200 V/m
+ 20,000 V/m
= For a uniform field, E = ΔV/d. With the distance in metres, that is 400 V / 0.02 m = 20,000 V/m.
? How much work does it take to move a charge along an equipotential surface?
+ None
- qV, where V is the potential of the surface
- It depends on the path
= The potential is the same everywhere on an equipotential, so ΔV = 0 between any two points on it. The energy change qΔV is therefore zero, and no work is needed.
:::
