package easy

import (
	"sort"
)

type Period[n number] struct {
	Start, End int64
	Num        n
}

func MergePeriods[n number](periods ...Period[n]) []Period[n] {
	if len(periods) <= 1 {
		return periods
	}

	// 去重区间左右值
	pointMapping := make(map[int64]struct{})
	for _, period := range periods {
		pointMapping[period.Start] = struct{}{}
		pointMapping[period.End] = struct{}{}
	}

	var points []int64
	for point := range pointMapping {
		points = append(points, point)
	}

	// 排序获取一个区间左右值切片
	sort.Slice(points, func(i, j int) bool {
		return points[i] < points[j]
	})

	// 将区间的累加值存储到map，key值为区间左值
	nums := make(map[int64]n)
	for _, period := range periods {
		for _, point := range points {
			if period.End < point {
				break
			}

			if point >= period.Start && point < period.End {
				nums[point] += period.Num
			}
		}
	}

	// 将num值中取出，生成period，写入result
	var result []Period[n]
	for idx, point := range points {
		if _, ok := nums[point]; ok {
			result = append(result, Period[n]{
				Start: point,
				End:   points[idx+1],
				Num:   nums[point]},
			)
		}
	}
	return result
}
