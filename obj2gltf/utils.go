package obj2gltf

import (
	"encoding/binary"
	"math"
)

func fToB(a *[]float64) (*[]byte) {
	raw := []byte{}
	for _, f := range *a {
		raw = append(raw, Float64bytes(f)...)
	}
	return &raw
}

// i don't remember :-(
func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}
