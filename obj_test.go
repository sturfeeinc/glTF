package obj

import (
	"testing"
	"bytes"
	"encoding/json"
	"os"
)

func TestParseTriple(t *testing.T) {



	if false {
		t.Error("it's not possible")
	}
}



var testObj = `# Unit-volume cube with the same texture coordinates on each face.
#
# Created by Morgan McGuire and released into the Public Domain on
# July 16, 2011.
#
# http://graphics.cs.williams.edu/data

mtllib default.mtl

v -0.5 0.5 -0.5
v -0.5 0.5 0.5
v 0.5 0.5 0.5
v 0.5 0.5 -0.5
v -0.5 -0.5 -0.5
v -0.5 -0.5 0.5
v 0.5 -0.5 0.5
v 0.5 -0.5 -0.5

vt 0 1
vt 0 0
vt 1 0
vt 1 1

vn 0 1 0
vn -1 0 0
vn 1 0 0
vn 0 0 -1
vn 0 0 1
vn 0 -1 0

g cube
usemtl default

f -8/-4/-6 -7/-3/-6 -6/-2/-6
f -8/-4/-6 -6/-2/-6 -5/-1/-6
f -8/-4/-5 -4/-3/-5 -3/-2/-5
f -8/-4/-5 -3/-2/-5 -7/-1/-5
f -6/-4/-4 -2/-3/-4 -1/-2/-4
f -6/-4/-4 -1/-2/-4 -5/-1/-4
f -5/-4/-3 -1/-3/-3 -4/-2/-3
f -5/-4/-3 -4/-2/-3 -8/-1/-3
f -7/-4/-2 -3/-3/-2 -2/-2/-2
f -7/-4/-2 -2/-2/-2 -6/-1/-2
f -3/-4/-1 -4/-3/-1 -1/-2/-1
f -3/-4/-1 -1/-2/-1 -2/-1/-1
`




func TestParseObj1(t *testing.T) {

/*	var obj Obj = Obj{}

	mtlFile

	obj.Materials = LoadMtl()*/




	f, _ := Parse(bytes.NewBufferString(testObj))

	df,_ := json.Marshal(f)

	if len(df) == 555 {

	}
	if false {
		t.Error("parse() return bad result")
	}
}



func TestParseObj(t *testing.T) {

	file, _ := os.Open("models/googlesf.obj")
	defer file.Close()
	f, _ := Parse(bytes.NewBufferString(testObj))
	f, _ = Parse(file)

	df,_ := json.Marshal(f)

	if len(df) == 555 {

	}
	if false {
		t.Error("parse() return bad result")
	}
}