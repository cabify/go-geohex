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
	// Constant math stuff
	deg2rad = math.Pi / 360.0
	rad2deg = 360 / math.Pi
	pio4    = math.Pi / 4
)

var (
	// Precalculated math stuff
	pow3f         [MaxLevel + 4]float64
	pow3i         [MaxLevel + 4]int
	halfCeilPow3f [MaxLevel + 4]float64
)

// A zoom is a helper for level dimensions
type Zoom struct {
	level  int
	factor float64
}

// Cached zooms lookup
var zooms = make(map[int]*Zoom, MaxLevel)

// LL is a lat/lon tuple
type LL struct {
	Lat, Lon float64
}

// NewLL creates a new normalised LL
func NewLL(lat, lon float64) *LL {
	if lon <= -180 {
		lon += 360
	} else if lon > 180 {
		lon -= 360
	}
	return &LL{Lat: lat, Lon: lon}
}

// Point generates a grid point from a lat/lon
func (ll *LL) Point() *Point {
	e := ll.Lon / 180.0
	n := math.Log(math.Tan(ll.Lat*deg2rad+pio4)) / math.Pi

	return &Point{E: e, N: n}
}

// init precalculates frequently used things
func init() {
	for i := 0; i <= MaxLevel+3; i++ {
		pow3f[i] = math.Pow(3, float64(i))
		halfCeilPow3f[i] = pow3f[i] / 2
		pow3i[i] = int(math.Pow(3, float64(i)))
	}

	for level := 0; level <= MaxLevel; level++ {
		zooms[level] = &Zoom{level: level, factor: 6 / pow3f[level+3]}
	}

	for i, b := range hChars {
		hIndex[b] = i
	}
}
