package color

// Color represents a color type.
//
//go:generate stringer -type=Color
type Color int

const (
	Red Color = iota
	Green
	Blue
	Yellow
	Purple
)
