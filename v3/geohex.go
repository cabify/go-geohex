package geohex

import (
	"math"
)

const VERSION = "3.0.0"

// MaxLevel is the maximum encoding level that this implementation supports
const MaxLevel = 20

var (
	hChars = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	hIndex = make(map[byte]int, len(hChars))
	hK     = math.Tan(math.Pi / 6.0)
)

const (
	hD2R  = math.Pi / 180.0
	hR2D  = 180 / math.Pi
	hBase = 20037508.34
	pio4  = math.Pi / 4
)

// A zoom is a helper for level dimensions
type Zoom struct {
	level  int
	size   float64
	factor float64
}

// Cached zooms lookup
var zooms = make(map[int]*Zoom, 20)

// LL is a lat/lon tuple
type LL struct {
	Lat, Lon float64
}

// NewLL creates a new normalised LL
func NewLL(lat, lon float64) *LL {
	if lon < -180 {
		lon += 360
	} else if lon > 180 {
		lon -= 360
	}
	return &LL{Lat: lat, Lon: lon}
}

// Point generates a grid point from a lat/lon
func (ll *LL) Point() *Point {
	e := ll.Lon * hBase / 180.0
	n := math.Log(math.Tan((90+ll.Lat)*hD2R/2)) / math.Pi * hBase

	return &Point{E: e, N: n}
}

// Init zooms
func init() {
	for level := 0; level <= MaxLevel; level++ {
		size := hBase / math.Pow(3, float64(level+3))
		zooms[level] = &Zoom{level: level, size: size, factor: 6 * size}
	}

	for i, b := range hChars {
		hIndex[b] = i
	}
}
