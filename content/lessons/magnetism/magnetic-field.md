Lodestone, a naturally magnetic iron ore that pulls on iron and swings to point north when hung on a thread, was known in ancient Greece and China. By about 1200, sailors in both China and Europe were steering by magnetized needles. For most of history magnetism looked like a force entirely separate from electricity, with its own materials, its own rules and its own mysteries. The central discovery of this chapter is that it is not separate at all: magnetic fields are made by moving electric charges. This first lesson sets out what a magnet does, how to describe its influence as a field, and where that field ultimately comes from.

## What are poles?

Every magnet has two ends where its pull is strongest. They are called its **north** and **south poles**, after the direction each end points when the magnet hangs freely. Bring two magnets together and you find that like poles repel each other and unlike poles attract. So far this sounds just like electric charge, where like charges repel and unlike charges attract, and it would be natural to guess that a north pole is a kind of magnetic charge.

But there is a crucial difference, and it shows up when you try to separate the poles. With electric charge you can easily have a lump that is purely positive. With a magnet you cannot. Cut a bar magnet in half, hoping to get a separate north end and south end, and you get two smaller magnets instead, each with its own north and south pole. Cut those in half and the same thing happens again. Keep going, down to single atoms, and you still never isolate one pole. A lone north or south pole would be called a magnetic **monopole**, and despite careful searches, in cosmic rays, in ancient rocks and in particle accelerators, none has ever been found.

::: history 1269 · Lucera, Italy
Petrus Peregrinus was a French scholar serving with the army of Charles of Anjou, which was besieging the Italian town of Lucera. With time on his hands, he wrote a long letter to a friend describing his experiments with a lodestone shaped into a sphere. He laid a small needle at many points on its surface, marked the direction it settled in each time, and found that all the lines he drew converged on two points at opposite sides of the sphere, which he named poles. He noted that like poles repel and unlike poles attract, and that a broken magnet becomes two complete magnets rather than two separate poles. His letter is one of the earliest accounts of a careful physical experiment in medieval Europe.
:::

## Why does a compass point north?

For centuries nobody knew why a magnetized needle should prefer north. Some thought it was drawn toward the Pole Star; others imagined a great magnetic mountain somewhere in the Arctic that pulled on every needle.

::: history 1600 · London
William Gilbert, a London physician who would later attend Queen Elizabeth I, spent years experimenting with magnets and in 1600 published his findings in a book called *De Magnete*. His key experiment used a lodestone ground into a sphere, which he called a *terrella*, a "little Earth". When he moved a small compass needle over its surface, the needle behaved just as compasses were known to behave on Earth: it pointed toward the sphere's poles, and near the poles it tilted to point down into the surface, an effect sailors had noticed and called "dip". Gilbert drew the bold conclusion that the Earth itself is a giant magnet. It was also one of the first arguments in physics built on a scale model: what holds for the little sphere holds for the big one.
:::

Gilbert's conclusion leads to a small puzzle about names. A compass needle's north pole is attracted toward the Earth's geographic north. Since unlike poles attract, the magnetic pole sitting in the Arctic must be, strictly speaking, a magnetic *south* pole. The labels are a historical accident that stuck. The Earth's magnetic poles also wander, by tens of kilometres a year, because the field is produced by moving molten iron deep in the Earth's core, and that motion changes over time.

## How do we describe magnetism everywhere at once?

As with electricity, it turns out to be cleanest to describe a magnet's influence by a **field**: a quantity that has a size and a direction at every point in space. The **magnetic field**, written B, points at each place in the direction a small compass needle would settle there. Its strength is measured in **teslas** (T), a unit defined in the next lesson by the force the field exerts on a moving charge. The tesla is a large unit, as these typical values show.

| Source | Field |
|---|---|
| Earth's surface | 25–65 μT |
| Fridge magnet | about 5 mT |
| Strong neodymium magnet, at its surface | about 1 T |
| Hospital MRI scanner | 1.5–3 T |

The field can be made visible. Sprinkle iron filings on a sheet of paper over a magnet and tap the paper: each filing becomes magnetized, behaves like a tiny compass, and lines up with the field where it lies. The filings form curving chains called **field lines**, which show the direction of the field at each point and crowd together where it is strongest.

Outside a bar magnet the field lines run from the north pole round to the south pole. Here the comparison with electricity breaks down again. Electric field lines begin on positive charges and end on negative ones. Magnetic field lines do not end at the south pole: they carry on through the body of the magnet, from south back to north, and so form closed loops with no beginning and no end. This is simply the field-line way of saying there are no monopoles. A monopole would be a place where lines start or stop, and there are none.

The same idea can be put precisely. Imagine any closed surface, like a balloon, anywhere in a magnetic field, and count field lines leaving it as positive and those entering as negative. Because every line is a closed loop, every line that enters the balloon must leave it again, so the net count is always zero. This statement is **Gauss's law for magnetism**, and in the compact notation of later lessons it is written ∇·B = 0. It is one of the four equations Maxwell eventually gathered together.

## Where does magnetic field come from?

In 1820 a Danish professor noticed that an electric current flowing in a wire made a nearby compass needle swing, a discovery told fully two lessons from now. Within months, careful measurements had found the field around a long straight wire. Its field lines turn out to be circles, centred on the wire, and the field's strength falls off steadily with distance.

::: equation Field of a long straight wire
B = \frac{\mu_0 I}{2\pi r}, \qquad \mu_0 \approx 4\pi \times 10^{-7}\,\text{T m/A}
:::

Here I is the current and r the distance from the wire. The constant μ₀ is called the **permeability of free space**; it sets how much magnetic field a given current produces, much as the constant in Coulomb's law sets how much electric field a given charge produces. Notice that the field falls as 1/r, not as 1/r² like the field of a point charge. That difference reflects the shape of the source: the wire is a long line, and the field spreads out around it in circles rather than in spheres.

To see what the formula means in practice, take a household current of 10 A, and stand a compass 1 cm from the wire. Then μ₀/2π is 2 × 10⁻⁷, and B = (2 × 10⁻⁷ × 10) ÷ 0.01 = 2 × 10⁻⁴ T. That is about four times the Earth's field, which is why a compass held near electrical wiring can point in quite the wrong direction.

The wire is not a special case. Every magnetic field ever measured is produced either by moving electric charge or by a property of particles that behaves in the same way. That raises an obvious question about the permanent magnet on a fridge door: no battery is connected to it, so where are its currents?

They are inside its atoms. Each electron behaves as a tiny magnet, partly because of its motion around the nucleus and mostly because of an intrinsic property called **spin**. Spin is sometimes pictured as the electron rotating on its axis, which would make it a little loop of moving charge, but that picture cannot be taken literally; spin is a quantum property with no exact everyday counterpart, and understanding it properly needs quantum mechanics. In most materials the atomic magnets point in random directions and cancel out. In iron, nickel and cobalt, neighbouring atoms tend to line up with one another, forming regions called **domains** in which millions of atomic magnets all point the same way. In an unmagnetized lump of iron the domains themselves point in different directions and cancel. Magnetizing the iron means getting its domains to line up, and then their fields add into one large field outside the magnet.

::: note
Magnetic fields come from moving charges, from currents in wires down to electrons in atoms. Their field lines always close on themselves, which is the same as saying that there are no magnetic monopoles: every magnet, however small, has both a north and a south pole.
:::

## Are electricity and magnetism really different?

The discovery that currents make magnetic fields was the first sign that electricity and magnetism are related. The full story came in 1905, when Albert Einstein showed that the split between them depends on how you are moving. Consider a charge sitting still in your laboratory. You measure only an electric field around it. Now let someone run past. To them the charge is moving, and a moving charge is a current, so they measure a magnetic field as well. Neither observer is wrong. What one calls a purely electric force, another calls partly magnetic. Electric and magnetic fields are two aspects of a single **electromagnetic field**, and which aspect you see depends on your motion. That deeper unity is where this part of the course is heading.

::: quiz
? A bar magnet is cut in half across its middle. What do you get?
- An isolated north pole and an isolated south pole
+ Two smaller magnets, each with a north and a south pole
- Two unmagnetized pieces
= Magnetic field lines form closed loops, so every piece of a magnet, however small, has both poles. No isolated magnetic pole has ever been found.
? At 2 cm from a long straight wire the field is 100 μT. What is it at 4 cm?
- 25 μT
+ 50 μT
- 100 μT
= The field of a straight wire is B = μ₀I/(2πr), which falls as 1/r rather than 1/r². Doubling the distance therefore halves the field.
? What ultimately produces the field of a permanent magnet?
- Stationary electric charges inside it
+ Aligned atomic magnetism, mainly from electron spin
- Magnetic monopoles packed into the poles
= Each electron acts as a tiny magnet, mostly because of its spin. In a magnetized piece of iron, vast numbers of these atomic magnets line up, so their fields add instead of cancelling.
:::
