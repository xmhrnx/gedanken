package curriculum

// ModelIDs lists the interactive models lessons may reference. The UI
// provides one model per id.
var ModelIDs = []string{
	"slope", "area", "projectile", "falling", "escape", "gravity-well",
	"light-cones", "binding-curve", "chain-reaction", "critical-mass", "atom",
	"orbit", "pendulum", "wave", "interference", "gas", "field-lines",
	"light-clock", "bohr", "wavefunction", "decay",
}

// KnownModel reports whether id is in ModelIDs.
func KnownModel(id string) bool {
	for _, m := range ModelIDs {
		if m == id {
			return true
		}
	}
	return false
}
