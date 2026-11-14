For two centuries after Gilbert, electricity and magnetism were studied side by side by the same people, often in the same rooms, without anyone finding a link between them. Static charges ignored magnets, and magnets ignored static charges. Several people suspected a connection, and some went looking for one, but they looked in the wrong place. The link appeared only after Volta's pile made steady currents possible, and when it came, a single chance observation was turned into a precise, general law within a few months. This lesson tells that story, and shows how it leads to the electromagnet, a magnet you can switch on and off.

## What did Ørsted see?

::: history 1820 · Copenhagen
Hans Christian Ørsted, a professor at the University of Copenhagen, believed on philosophical grounds that the forces of nature were all connected, and he had long suspected that electricity could affect a magnet. In the spring of 1820, during a lecture, he placed a wire carrying current from a voltaic pile near a compass, and the needle moved. The effect was small, and his audience was not much impressed, but Ørsted spent the following months investigating it carefully. He found that the needle did not swing toward the wire or away from it, as anyone would have expected, but turned *across* it, to stand at right angles to the wire. Reversing the current reversed the swing. Placing the needle above the wire instead of below it reversed it too. In July 1820 he announced his results in a four-page pamphlet written in Latin and sent it to scientists across Europe. Within weeks, the experiment was being repeated everywhere.
:::

Why had nobody found this before? Partly because earlier experimenters mostly used static charge, which does nothing to a magnet; it takes moving charge. And partly because they expected the wrong kind of effect. Every force known until then, whether gravitational, electric or magnetic, pulled or pushed along the straight line joining two objects. If a current affected a magnet, people expected it to attract or repel the poles. Ørsted's force did something else entirely: it went round.

We now describe this by saying that the magnetic field lines around a straight wire are circles centred on the wire, lying in planes at right angles to it. A compass needle lines up with the field, so it lies along one of these circles, across the wire, exactly as Ørsted saw. Another **right-hand rule** gives the direction: grip the wire with your right hand so that your thumb points along the current, and your fingers curl round the wire in the direction the field points.

## Can currents act on currents?

::: history 1820 · Paris
André-Marie Ampère, a mathematician in Paris, heard Ørsted's result described at the French Academy of Sciences on 11 September 1820. Within a week he had presented a discovery of his own. If a current makes a magnetic field, and a magnetic field pushes on a current, then two currents ought to push on each other with no magnet involved at all. Ampère showed that they do: two parallel wires carrying currents in the same direction attract each other, and carrying currents in opposite directions they repel. He went on to propose that all magnetism, including that of a lump of magnetized iron, is caused by electric currents, and he suggested that iron contains tiny circulating currents, all lined up. In the same weeks Jean-Baptiste Biot and Félix Savart measured how the field around a long wire depends on distance and found that it falls as 1/r.
:::

Ampère's guess about tiny currents inside iron was remarkably close to the truth, as the first lesson of this chapter described. The "currents" turned out to be the spin and orbital motion of electrons, which only quantum mechanics describes properly, but the idea that magnetism is always the work of moving charge was right.

## Is there a general rule linking field and current?

Biot and Savart had measured the field of one particular arrangement, a long straight wire. What physicists wanted was a rule that would work for any arrangement of currents, and Ampère's work led to one. It needs one new idea, which describes how much a field "goes round" a loop.

Draw any closed loop in space, of any shape. Walk round it in small steps, each of length dl. At each step, take the part of the magnetic field that points along your direction of travel, and multiply it by the step length. Then add up these products all the way round. The total is called the **circulation** of the field round the loop, and it is written ∮B·dl, where the circle on the integral sign means "all the way round a closed loop". For the circular field around a wire, the circulation is large, because the field points along the loop everywhere. **Ampère's law** says that the circulation depends only on the total current passing through the loop.

::: equation Ampère's law
\oint \vec{B} \cdot \mathrm{d}\vec{l} = \mu_0 I_{\text{enclosed}}
:::

Here I_enclosed is the total current threading through the loop, and μ₀ is the same constant that appeared in the field of a straight wire. Currents outside the loop do not count: their fields go round the loop partly with you and partly against you, and the contributions cancel. The law holds for any loop you care to draw, but it is most useful when the arrangement is symmetric enough that the field is the same size all round a well-chosen loop. Then it gives the field almost for free.

::: steps Field of a long straight wire from Ampère's law
$ \oint \vec{B} \cdot \mathrm{d}\vec{l} = B \oint \mathrm{d}l
Choose a circular loop of radius r, centred on the wire and at right angles to it. By symmetry, nothing distinguishes one point on this circle from another, so the field must have the same size B all the way round, and Ørsted's observation tells us it points along the circle. So at every step the part of B along the loop is just B, and since it is constant it can be taken outside the sum.
---
$ B \cdot 2\pi r = \mu_0 I
What is left inside the sum is just the step lengths added up, which is the total length of the loop: its circumference, 2πr. On the right-hand side, the only current passing through the loop is the wire's own current I.
---
$ B = \frac{\mu_0 I}{2\pi r}
Divide both sides by 2πr to find the field. This is exactly the 1/r law that Biot and Savart had measured, now derived from a general principle in three lines. The 1/r comes directly from the fact that the loop's length grows in proportion to its radius, so the same circulation is spread over a longer path.
:::

## How do you make a strong magnetic field?

The field of a straight wire is weak, because it spreads out in all directions. To concentrate it, bend the wire into a loop. Now the field lines from every part of the loop pass through the middle in the same direction, and their contributions add. Wind many loops one after another into a long coil, called a **solenoid**, and the effect builds up further. Inside the coil the field becomes strong and nearly uniform, pointing along the coil's axis. Outside it is weak, and its pattern looks just like that of a bar magnet, with a north pole at one end and a south pole at the other. Ampère's law gives the field inside.

::: steps Field inside a long solenoid
$ \oint \vec{B} \cdot \mathrm{d}\vec{l} = B\,\ell
Choose a rectangular loop with one long side, of length ℓ, running inside the coil parallel to its axis, and the opposite side running outside. Inside, the field is uniform and points along that side, so that side contributes Bℓ. The two short sides cross the coil's wall at right angles to the field, so they contribute nothing. The side outside contributes almost nothing either, because the field outside a long solenoid is very weak.
---
$ I_{\text{enclosed}} = n\,\ell\,I
Now count the current passing through the rectangle. If the coil has n turns of wire per metre of its length, then nℓ turns pass through the rectangle, and each carries the current I, so the total current threading the loop is nℓI.
---
$ B\,\ell = \mu_0\, n\,\ell\,I
Ampère's law sets the circulation equal to μ₀ times the enclosed current. Both sides contain the length ℓ, which reflects the fact that a longer rectangle encloses more turns and also runs further along the field.
---
$ B = \mu_0 n I
Divide both sides by ℓ. The field inside depends only on the current and how tightly the coil is wound, not on the coil's width, and not on where inside it you measure, which is why the field is uniform.
:::

A coil wound with 1,000 turns per metre and carrying 2 A therefore has a field inside of B = 4π × 10⁻⁷ × 1000 × 2 ≈ 2.5 × 10⁻³ T, or 2.5 millitesla. That is about fifty times the Earth's field: strong for a simple coil, but still far short of a good permanent magnet.

The great improvement comes from putting a bar of iron inside the coil. The coil's field lines up the iron's magnetic domains, the regions of aligned atomic magnets described earlier, and their combined field adds to the coil's own. The result can be hundreds or even thousands of times stronger than the coil alone, until the iron's domains are all aligned and it can give no more. This is an **electromagnet**: a magnet that can be switched on and off, and whose strength is set by the current.

The first was built by William Sturgeon in England in 1825: a horseshoe of iron, varnished to insulate it, with a few turns of bare copper wire wound round it. Joseph Henry, a teacher in Albany, New York, saw that more turns would give a stronger field, and since the wire was bare the turns had to be kept apart; he insulated his wire by wrapping it in silk, which let him wind many layers tightly. In 1831 he built one for Yale College that could lift about a tonne.

Electromagnets are now everywhere: in scrapyard cranes that drop cars at the flick of a switch, in relays and electric door locks, in loudspeakers and in MRI scanners, whose coils are made of superconducting wire and carry hundreds of amperes with no resistance at all.

::: note
A current is surrounded by circular magnetic field lines, and Ampère's law says the circulation of the field round any loop is μ₀ times the current passing through it. Winding the wire into a coil concentrates the field, and an iron core multiplies it many times over, making an electromagnet.
:::

## How did Ampère's wires define the ampere?

The force between Ampère's parallel wires turned out to be useful in an unexpected way. Because it can be measured precisely, it was used from 1948 to 2019 as the official definition of the unit of current: one ampere was the current that, flowing in two long, thin, parallel wires one metre apart, makes them attract each other with a force of 2 × 10⁻⁷ newtons on each metre of length. That definition is the reason μ₀ was then exactly 4π × 10⁻⁷ T m/A; the number was fixed by the choice of unit rather than measured. Since 2019 the ampere has been defined instead by fixing the charge of a single electron at an exact value. As a result μ₀ is now a measured quantity, and it agrees with the old value to about one part in ten billion.

::: quiz
? Ørsted's compass needle, placed under a wire running north to south, swung toward east or west. Why was this surprising?
- Compasses were not expected to respond to anything but the Earth
+ The force pointed across the wire, not along a line toward or away from it
- The needle should have pointed at the battery
= Every force known at the time acted along the line between two objects. The magnetic field of a current instead circles around the wire, so the needle turned across it.
? Two parallel wires carry currents in opposite directions. What happens?
- They attract
+ They repel
- They feel no force
= Ampère found that parallel currents attract and opposite currents repel. Each wire sits in the other's magnetic field, and reversing one current reverses the force on it.
? A solenoid's current is doubled and its number of turns per metre is halved. What happens to the field inside?
- It doubles
+ It stays the same
- It halves
= The field inside is B = μ₀nI. Doubling I and halving n leave their product, and so the field, unchanged.
:::
