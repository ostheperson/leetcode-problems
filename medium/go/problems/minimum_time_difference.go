package problems

import (
	"sort"
	"strconv"
)

func FindMinDifference(timePoints []string) int {
	hash := make(map[string]int)
	for _, val := range timePoints {
		if _, exists := hash[val]; exists {
			return 0
		}
		hash[val] = convertMinute(val)
	}

	sort.SliceStable(timePoints, func(i, j int) bool {
		return timePoints[i] < timePoints[j]
	})
	var minimum int = 24 * 60

	for i := 1; i < len(timePoints); i++ {
	}
	return minimum
}

func convertMinute(i string) int {
	hourI, err1 := strconv.Atoi(i[:2])

	if err1 != nil {
		return 0
	}

	minuteI, err1 := strconv.Atoi(i[3:])

	if err1 != nil {
		return 0
	}

	return (hourI * 60) + minuteI
}

func calcDiff(i, j string) int {
	hourI, err1 := strconv.Atoi(i[:2])
	hourJ, err2 := strconv.Atoi(j[:2])

	if err1 != nil || err2 != nil {
		return 0
	}

	minuteI, err1 := strconv.Atoi(i[3:])
	minuteJ, err2 := strconv.Atoi(j[3:])

	if err1 != nil || err2 != nil {
		return 0
	}

	iMinutes := (60 * hourI) + minuteI
	jMinutes := (60 * hourJ) + minuteJ
	diff := abs(iMinutes - jMinutes)

	if diff > 12*60 {
		jMinutes = (60 * (hourJ + 24)) + jMinutes
		return abs(iMinutes - jMinutes)
	}
	return diff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
