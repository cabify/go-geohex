package geohex

// Position implements a grid tile position
type Position struct {
	X, Y int
	z    *Zoom
}

// Centroid returns the centroid point of the tile
func (p *Position) Centroid() *Point {
	x, y := float64(p.X), float64(p.Y)

	return &Point{
		E: p.z.factor * (x - y) / 2,
		N: p.z.factor * (x + y) * nScaleFactor,
	}
}

// LL converts the position into a LL
func (p *Position) LL() *LL {
	return p.Centroid().LL()
}
