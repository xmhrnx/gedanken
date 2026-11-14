package physics

// Nuclide is a nucleus with its binding energy per nucleon.
type Nuclide struct {
	Name    string
	A       int     // mass number
	Binding float64 // MeV per nucleon
}

// BindingCurve lists measured binding energies per nucleon for nuclides
// along the stable valley, from light to heavy.
var BindingCurve = []Nuclide{
	{"Deuterium", 2, 1.112},
	{"Tritium", 3, 2.827},
	{"Helium-4", 4, 7.074},
	{"Lithium-6", 6, 5.332},
	{"Lithium-7", 7, 5.606},
	{"Beryllium-9", 9, 6.463},
	{"Boron-11", 11, 6.928},
	{"Carbon-12", 12, 7.680},
	{"Nitrogen-14", 14, 7.476},
	{"Oxygen-16", 16, 7.976},
	{"Neon-20", 20, 8.032},
	{"Magnesium-24", 24, 8.261},
	{"Silicon-28", 28, 8.448},
	{"Sulfur-32", 32, 8.493},
	{"Calcium-40", 40, 8.551},
	{"Chromium-52", 52, 8.776},
	{"Iron-56", 56, 8.790},
	{"Nickel-62", 62, 8.795},
	{"Zinc-64", 64, 8.736},
	{"Krypton-84", 84, 8.717},
	{"Zirconium-90", 90, 8.710},
	{"Tin-120", 120, 8.505},
	{"Barium-138", 138, 8.393},
	{"Neodymium-142", 142, 8.346},
	{"Gold-197", 197, 7.916},
	{"Lead-208", 208, 7.868},
	{"Uranium-235", 235, 7.591},
	{"Uranium-238", 238, 7.570},
}

// MostBound returns the nuclide with the highest binding per nucleon.
func MostBound() Nuclide {
	best := BindingCurve[0]
	for _, n := range BindingCurve {
		if n.Binding > best.Binding {
			best = n
		}
	}
	return best
}
