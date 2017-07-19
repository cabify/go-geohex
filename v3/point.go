package geohex

import (
	"math"
)

// Point implements geographic Cartesian coordinates
type Point struct {
	E, N float64
}

// Position returns the X/Y grid position of the Point
func (p *Point) Position(z *Zoom) *Position {
	x, y := (p.N/hK+p.E)/z.factor, (p.N/hK-p.E)/z.factor

	x0, y0 := math.Floor(x), math.Floor(y)
	xd, yd := x-x0, y-y0

	pos := Position{z: z}
	if yd > -xd+1 && yd < 2*xd && yd > 0.5*xd {
		pos.X, pos.Y = int(x0)+1, int(y0)+1
	} else if yd < -xd+1 && yd > 2*xd-1 && yd < 0.5*xd+0.5 {
		pos.X, pos.Y = int(x0), int(y0)
	} else {
		pos.X, pos.Y = int(math.Floor(x+0.499999)), int(math.Floor(y+0.499999))
	}

	return &pos
}

// LL returns LL coordinates of this point
func (p *Point) LL() *LL {
	lat := (math.Atan(math.Exp(p.N*math.Pi)) - pio4) * rad2deg
	lon := p.E * 180

	return NewLL(lat, lon)
}
