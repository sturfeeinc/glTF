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
		t.Error("parseRGB() returns a bad result")
	}
}

func TestParseFloat (t *testing.T) {
	have := parseFloat("0.8")
	want := 0.8
	if have != want {
		t.Error("parseFloat returns a bad result")
	}
}

func TestParseInt (t *testing.T) {
	have := parseInt("255")
	want := 255
	if have != want {
		t.Error("parseInt returns a bad result")
	}
}

func DEBUG(i interface{})  {
	fmt.Printf("%f\n", i)
	fmt.Printf("%v\n", i)
}
