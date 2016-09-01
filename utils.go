package obj

import (
	"strings"
	"strconv"
)

func parseRGB(arr [][]byte) [3]float64 {
	a := [3]float64{}
	a[0] = parseFloat(string(arr[0]))
	a[1] = parseFloat(string(arr[1]))
	a[2] = parseFloat(string(arr[2]))
	return a
}

// alias for strconv.ParseFloat with exception
func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(strings.Trim(s, " "), BS)
	if err != nil {
		panic(err)
	}
	return f
}

// alias for strconv.ParseFloat with exception
func parseInt(s string) int {
	f, err := strconv.ParseInt(strings.Trim(s, " "), 10, BS)
	if err != nil {
		panic(err)
	}
	return int(f)
}

func isSpace(b byte) bool {
	return b == WS || b == T
}


