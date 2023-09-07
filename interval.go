package easy

import (
	"sort"
)

// Interval defines an interval with power.
type Interval[c Number, p Number] struct {
	// left and right endpoint
	Left, Right c
	// the power of the interval
	Power p
}

// MergeIntervals merges all the intervals, sums the powers and divides them into a new interval slice.
func MergeIntervals[c Number, p Number](intervals ...Interval[c, p]) []Interval[c, p] {
	if len(intervals) <= 1 {
		return intervals
	}

	// deduplicate the left and right endpoints.
	pointMapping := make(map[c]struct{})
	for _, interval := range intervals {
		pointMapping[interval.Left] = struct{}{}
		pointMapping[interval.Right] = struct{}{}
	}

	points := make([]c, 0, len(pointMapping))
	for point := range pointMapping {
		points = append(points, point)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i] < points[j]
	})

	// sums the power into a map whose key(s) is the left endpoint(s).
	powerMapping := make(map[c]p)
	for _, interval := range intervals {
		for _, point := range points {
			if interval.Right < point {
				break
			}

			if point >= interval.Left && point < interval.Right {
				powerMapping[point] += interval.Power
			}
		}
	}

	var result []Interval[c, p]
	for idx, point := range points {
		if _, ok := powerMapping[point]; ok {
			result = append(result, Interval[c, p]{
				Left:  point,
				Right: points[idx+1],
				Power: powerMapping[point],
			})
		}
	}
	return result
}
