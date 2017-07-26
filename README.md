# GeoHex

[![Build Status](https://travis-ci.org/cabify/go-geohex.png)](https://travis-ci.org/cabify/go-geohex)
[![GoDoc](https://godoc.org/github.com/cabify/go-geohex?status.png)](http://godoc.org/github.com/cabify/go-geohex)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

[GeoHex](http://www.geohex.org/) implementation in Go.
Forked from [github.com/bsm/go-geohex](https://github.com/bsm/go-geohex)

## Quick Start

```go
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
```

## Efficiency
If you only need to simply convert Tiles to codes you can use `Encode` and `Decode`
functions, but if you need to calculate Neighbours of a cell, you should use the
`EncodeTile` and `DecodeTile` functions, avoiding this way to calculate the Lat/Lon
projection of the coordinates all the time.

## Running tests

You need to install Ginkgo & Gomega to run tests. Please see
http://onsi.github.io/ginkgo/ for more details.

    $ make testdeps

To run tests, call:

    $ make test

To run benchmarks, call:

    $ make bench

## Benchmarks for Encode/Decode

    BenchmarkEncodeLevel2-4        	 5000000	       285 ns/op	       8 B/op	       2 allocs/op
    BenchmarkEncodeLevel6-4        	 5000000	       356 ns/op	      16 B/op	       2 allocs/op
    BenchmarkEncodeLevel15-4       	 3000000	       521 ns/op	      64 B/op	       2 allocs/op
    BenchmarkDecodeLevel2-4        	 5000000	       358 ns/op	       3 B/op	       1 allocs/op
    BenchmarkDecodeLevel6-4        	 3000000	       426 ns/op	       3 B/op	       1 allocs/op
    BenchmarkDecodeLevel15-4       	 3000000	       575 ns/op	       3 B/op	       1 allocs/op
    BenchmarkDecodeTileLevel2-4    	 5000000	       237 ns/op	       3 B/op	       1 allocs/op
    BenchmarkDecodeTileLevel6-4    	 5000000	       306 ns/op	       3 B/op	       1 allocs/op
    BenchmarkDecodeTileLevel15-4   	 3000000	       440 ns/op	       3 B/op	       1 allocs/op
