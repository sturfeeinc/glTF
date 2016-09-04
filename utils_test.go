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

func TestParseTriple (t *testing.T) {
	have, check := parseTriple([]byte("1234"))
	if check != 1 {
		print("parseTriple returns a bad result. We expected 1 (v) but got ")
		println(check)
		t.Error("parseTriple")
	}
	v := 1234
	want := IndexT{&v, nil, nil}
	if !reflect.DeepEqual(want, have) {
		t.Error("parseTriple returns a bad result for 1 (v)")
	}
	have, check = parseTriple([]byte("5432/4321"))
	if check != 2 {
		print("parseTriple returns a bad result. We expected 2 (v/vt) but got ")
		println(check)
		t.Error("parseTriple")
	}
	v = 5432
	vt := 4321
	want = IndexT{&v, &vt, nil}
	if !reflect.DeepEqual(want, have) {
		t.Error("parseTriple returns a bad result for 2 (v/vt)")
	}
	have, check = parseTriple([]byte("6543//3456"))
	if check != 3 {
		print("parseTriple returns a bad result. We expected 3 (v//vn) but got ")
		println(check)
		t.Error("parseTriple")
	}
	v = 6543
	vn := 3456
	want = IndexT{&v, nil, &vn}
	if !reflect.DeepEqual(want, have) {
		t.Error("parseTriple returns a bad result for 3 (v//vn)")
	}
	have, check = parseTriple([]byte("3456/6543/9876"))
	if check != 4 {
		print("parseTriple returns a bad result. We expected 4 (v/vt/vn) but got ")
		println(check)
		t.Error("parseTriple")
	}
	v = 3456
	vt = 6543
	vn = 9876
	want = IndexT{&v, &vt, &vn}
	if !reflect.DeepEqual(want, have) {
		t.Error("parseTriple returns a bad result for 4 (v/vt/vn)")
	}
}

func DEBUG(i interface{})  {
	fmt.Printf("%f\n", i)
	fmt.Printf("%v\n", i)
}
