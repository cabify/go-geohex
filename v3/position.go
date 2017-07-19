package geohex

// Position implements a grid tile position
type Position struct {
	X, Y int
	z    *Zoom
}

// Centroid returns the centroid point of the tile
func (p *Position) Centroid() *Point {
	x := float64(p.X)
	y := float64(p.Y)
	n := (x + y) * p.z.factor * hK / 2
	e := n/hK - y*p.z.factor
	return &Point{E: e, N: n}
}

// LL converts the position into a LL
func (p *Position) LL() *LL {
	return p.Centroid().LL()
}
