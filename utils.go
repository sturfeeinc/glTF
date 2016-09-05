package obj

import (
	"strings"
	"strconv"
	"bytes"
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

// Parse triples with index offsets: i, i/j/k, i//k, i/j
func parseTriple(token []byte) (IndexT, int){
	count := bytes.Count(token, SL)
	i := IndexT{}
	if count == 0 {
		// v
		v := parseInt(string(token))
		i.VertexIndex = &v
		return i, 1
	}
	res := bytes.Split(token, SL)
	v := parseInt(string(res[0]))
	i.VertexIndex = &v
	if count == 1 {
		// v/vt
		vt := parseInt(string(res[1]))
		i.TexcoordIndex = &vt
		return i, 2
	}
	if len(res[1]) == 0 {
		// v//vn
		vn := parseInt(string(res[2]))
		i.NormalIndex = &vn
		return i, 3
	}
	// v/vt/vn
	vt := parseInt(string(res[1]))
	i.TexcoordIndex = &vt
	vn := parseInt(string(res[2]))
	i.NormalIndex = &vn
	return i, 4
}


// Parse triples with index offsets: i, i/j
func parseCouple(token []byte) (IndexT, int){
	count := bytes.Count(token, SL)
	i := IndexT{}

	if count == 0 {
		// v
		v := parseInt(string(token))
		i.VertexIndex = &v
		return i, 1
	}
	// v/vt
	res := bytes.Split(token, SL)
	v := parseInt(string(res[0]))
	i.VertexIndex = &v
	vt := parseInt(string(res[1]))
	i.TexcoordIndex = &vt
	return i, 2
}

func sacralSense(){
	complicated := bytes.ContainsRune([]byte{}, '/')
	if !complicated {
		return
	}
}

