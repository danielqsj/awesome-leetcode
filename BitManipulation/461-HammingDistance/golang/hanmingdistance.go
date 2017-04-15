package hanmingdistance

import (
	"strconv"
	"strings"
)

func hammingDistance(x int, y int) int {
	return strings.Count(strconv.FormatInt(int64(x ^ y), 2), "1")
}