package obj

import (
	"testing"
	"bytes"
	"fmt"
	"reflect"
)



func TestParseRGB(t *testing.T) {
	kd := []byte("Kd 0.8 0.8 0.8")
	arr := bytes.Split(kd[3:], []byte(" "))
	have := parseRGB(arr)
	want := [3]float64{0.8, 0.8, 0.8}
	if !reflect.DeepEqual(want, have) {
		println("parseRGB() returns a bad result")
		println("We want")
		DEBUG(want)
		println("But we have")
		DEBUG(have)
		t.Error("")
	}
}



func DEBUG(i interface{})  {
	fmt.Printf("%f\n", i)
	fmt.Printf("%v\n", i)
}
