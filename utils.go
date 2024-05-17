package ucdt

import (
	"strings"

	"github.com/B9O2/canvas/containers"
)

func InterceptData[T string | []byte](data T, start, end int) T {
	var result T
	max := len(data)

	if end > 0 {
		if end < max {
			max = end
		}
	} else {
		max -= end
	}

	min := start
	if start < 0 {
		min = max - start
	}

	if min > max {
		return result
	} else {
		return data[min:max]
	}
}

func SpaceAllocate(ranges map[string]containers.Range, length uint) map[string]uint {
	rangeIndex := map[string]int{}
	rangeList := []containers.Range{}
	result := map[string]uint{}

	for k, r := range ranges {
		rangeIndex[k] = len(rangeList)
		rangeList = append(rangeList, r)
	}

	allocated, err := containers.SpaceAllocate(rangeList, length)
	if err != nil {
		if strings.Contains(err.Error(), "small") {
			for k, r := range ranges {
				result[k] = r.Max()
			}
		} else if strings.Contains(err.Error(), "large") {
			for k, r := range ranges {
				result[k] = r.Min()
			}
		} else {
			for k := range ranges {
				result[k] = 0
			}
		}
	} else {
		for k, i := range rangeIndex {
			result[k] = allocated[i]
		}
	}
	return result
}
