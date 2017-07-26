package geohex_test

import (
	"fmt"

	geohex "github.com/cabify/go-geohex/v3"
)

func ExampleEncode() {
	code, _ := geohex.Encode(35.647401, 139.716911, 6)
	fmt.Println(code)

	// Output:
	// XM488541
}

func ExampleDecode() {
	ll, _ := geohex.Decode("XM488541")
	fmt.Println(ll.Lat, ll.Lon)

	// Output:
	// 35.63992106908978 139.72565157750344
}

func ExampleNeighbours() {
	tile, _ := geohex.DecodeTile("OX")
	fmt.Println(tile.Neighbours())

	// Output:
	// [Ob Oa OP OM OU OY]
}
