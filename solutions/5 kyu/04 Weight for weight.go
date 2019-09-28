package kata

import "strings"

import "sort"

func OrderWeight(strng string) string {
	if strng == "" {
		return ""
	}

	sl := strings.Split(strng, " ")
	sort.Slice(sl, func(i, j int) bool {
		sumI, sumJ := 0, 0
		for _, c := range sl[i] {
			sumI += int(c - '0')
		}
		for _, c := range sl[j] {
			sumJ += int(c - '0')
		}
		if sumI == sumJ {
			return sl[i] < sl[j]
		}
		return sumI < sumJ
	})

	return strings.Join(sl, " ")
}
