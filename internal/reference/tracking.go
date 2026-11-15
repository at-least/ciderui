package reference

// Tracking is ciderui's letter-spacing at one of apple.com's sizes: Em, in em, lands
// the kit's font stack on apple.com's widths at Size px. reference/tracking/table/
// holds one a size, which go run ./reference/tracking derive writes.
type Tracking struct {
	Size int
	Em   float64
}
