package main

import (
	"fmt"
	"math"
)

type Point struct {
	Val float64
	Ts  uint32
}

func Fix(in []Point, start, end, interval uint32) []Point {
	var result []Point = make([]Point, 0)
	for _, i := range in {

		index := shiftIndex(start, end, interval)
		if start < i.Ts {
			result = append(result, shiftLeading(index, start, interval)...)
		}
		if start == i.Ts {
			result = append(result, shiftLagging(i, index, start, interval)...)
		}
		start += interval
	}
	return result
}

func shiftIndex(start, end, delta uint32) int {
	return int((end - start) / delta)
}

func shiftLeading(index int, start, delta uint32) []Point {
	var result = make([]Point, 0)
	var offset = start

	for idx := 0; idx < index; idx++ {
		// create missing points
		result = append(result, Point{Val: math.NaN(), Ts: offset})
		offset += delta
	}
	return result
}

func shiftLagging(ref Point, index int, start, delta uint32) []Point {
	var result = make([]Point, 0)
	var offset = start
	for idx := 0; idx < index; idx++ {
		if offset == ref.Ts {
			result = append(result, ref)
		} else {
			result = append(result, Point{Val: math.NaN(), Ts: offset})
		}
		offset += delta
	}
	return result
}
