package main

import (
	"math"
)

type Point struct {
	Val float64
	Ts  uint32
}

func Fix(in []Point, start, end, interval uint32) []Point {
	var result []Point = make([]Point, 0)
	for _, i := range in {
		if start < end {
			if start < i.Ts {
				index := offset(start, i.Ts, interval)
				result = append(result, shift(index, start, interval)...)
				for index != 0 {
					start += interval
					index--
				}

			} else if start == i.Ts {
				result = append(result, i)
				start += interval
			} else {
				continue
			}
		}
	}

	index := offset(start, end, interval)
	// are there missing points from last point to end
	if index > 1 {
		result = append(result, shift(index, start, interval)...)
	}
	return result
}

func offset(start, end, interval uint32) int {
	delta := int(end) - int(start)
	return delta / int(interval)
}

func shift(index int, start, delta uint32) []Point {
	var result = make([]Point, 0)

	for idx := 0; idx < index; idx++ {
		// create missing points
		result = append(result, Point{Val: math.NaN(), Ts: start})
		start += delta
	}
	return result
}

func roundToNearest10(delta uint32) uint32 {
	var ones float64 = float64(delta) / 10
	roundedFloat := math.Round(ones) * 10
	return math.Float32bits(float32(roundedFloat))
}
