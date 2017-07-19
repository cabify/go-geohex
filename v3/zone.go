package geohex

import (
	"fmt"
	"strconv"
)

// Error types
var (
	ErrLevelInvalid = fmt.Errorf("geohex: level invalid")
	ErrCodeInvalid  = fmt.Errorf("geohex: code invalid")
)

// A Zone represents a GeoHex tile
type Zone struct {
	Code string
	Pos  *Position
}

// String returns the zone code
func (z *Zone) String() string {
	return z.Code
}

// Level returns the level
func (z *Zone) Level() int {
	return len(z.Code) - 2
}

// Encode encodes a lat/lon/level into a Zone
func Encode(lat, lon float64, level int) (_ *Zone, err error) {
	zoom, ok := zooms[level]
	if !ok {
		return nil, ErrLevelInvalid
	}

	pnt := NewLL(lat, lon).Point() // Point at lat/lon
	pos := pnt.Position(zoom)      // Tile position

	x, y := float64(pos.X), float64(pos.Y)
	base, num, code := 0, 0, make([]byte, level+2)

	for i := 0; i < level+3; i++ {
		pow := pow3f[level+2-i]
		p2c := halfCeilPow3f[level+2-i]

		if x >= p2c {
			x -= pow
			num = 6
		} else if x <= -p2c {
			x += pow
			num = 0
		} else {
			num = 3
		}

		if y >= p2c {
			y -= pow
			num += 2
		} else if y <= -p2c {
			y += pow
			// num += 0
		} else {
			num += 1
		}

		if i >= 3 {
			code[i-1] = '0' + byte(num)
		} else if i == 2 {
			base += num
		} else if i == 1 {
			base += 10 * num
		} else {
			base += 100 * num
		}
	}

	code[0] = hChars[base/30]
	code[1] = hChars[base%30]

	return &Zone{Code: string(code), Pos: pos}, nil
}

// Decode decodes a string code into Point
func Decode(code string) (_ *LL, err error) {
	lnc := len(code)
	zoom, ok := zooms[lnc-2]
	if !ok {
		return nil, ErrCodeInvalid
	}

	var n1, n2 int
	if n1, ok = hIndex[code[0]]; !ok {
		return nil, ErrCodeInvalid
	} else if n2, ok = hIndex[code[1]]; !ok {
		return nil, ErrCodeInvalid
	}

	base := n1*30 + n2
	if base < 100 {
		code = "0" + strconv.Itoa(base) + code[2:]
	} else {
		code = strconv.Itoa(base) + code[2:]
	}

	pos := &Position{z: zoom}
	for i, digit := range code {
		n := int64(digit - '0')
		if n < 0 || n > 9 {
			err = fmt.Errorf("expected a digit, got '%s'", digit)
			return
		}

		pow := pow3i[lnc-i]
		c3x := n / 3
		c3y := n % 3
		switch c3x {
		case 0:
			pos.X -= pow
		case 2:
			pos.X += pow
		}
		switch c3y {
		case 0:
			pos.Y -= pow
		case 2:
			pos.Y += pow
		}
	}

	return pos.LL(), nil
}
