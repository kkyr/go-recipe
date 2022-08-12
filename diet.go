//go:generate stringer -type=Diet -output=diet_string.go
package recipe

// Diet is a diet restricted to certain foods or preparations.
type Diet int

const (
	UnknownDiet Diet = iota
	DiabeticDiet
	GlutenFreeDiet
	HalalDiet
	HinduDiet
	KosherDiet
	LowCalorieDiet
	LowFatDiet
	LowLactoseDiet
	LowSaltDiet
	VeganDiet
	VegetarianDiet
)
