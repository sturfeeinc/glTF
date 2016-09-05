package obj

import (
	"testing"
	"os"
	"reflect"
)

func TestParseObj(t *testing.T) {

	testObj, err := os.Open("models/default.obj")
	if err != nil {
		t.Error(err)
	}
	defer testObj.Close()

	indexes := []string{
		"-8/-4/-6", "-7/-3/-6", "-6/-2/-6",
		"-8/-4/-6", "-6/-2/-6", "-5/-1/-6",
		"-8/-4/-5", "-4/-3/-5", "-3/-2/-5",
		"-8/-4/-5", "-3/-2/-5", "-7/-1/-5",
		"-6/-4/-4", "-2/-3/-4", "-1/-2/-4",
		"-6/-4/-4", "-1/-2/-4", "-5/-1/-4",
		"-5/-4/-3", "-1/-3/-3", "-4/-2/-3",
		"-5/-4/-3", "-4/-2/-3", "-8/-1/-3",
		"-7/-4/-2", "-3/-3/-2", "-2/-2/-2",
		"-7/-4/-2", "-2/-2/-2", "-6/-1/-2",
		"-3/-4/-1", "-4/-3/-1", "-1/-2/-1",
		"-3/-4/-1", "-1/-2/-1", "-2/-1/-1",
	}
	mesh := MeshT{}
	n := uint(3)
	mesh.num_face_vertices = []uint{
		n, n, n, n, n, n, n, n, n, n, n, n,
		n, n, n, n, n, n, n, n, n, n, n, n,
		n, n, n, n, n, n, n, n, n, n, n, n,
	}
	mesh.material_ids = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	for _, triple := range indexes {
		index, _ := parseTriple([]byte(triple))
		mesh.indices = append(mesh.indices, index)
	}

	attr := AttribT{}
	attr.Vertices = [][3]float64{
		[3]float64{-0.5, 0.5, -0.5},
		[3]float64{-0.5, 0.5, 0.5},
		[3]float64{0.5, 0.5, 0.5},
		[3]float64{0.5, 0.5, -0.5},
		[3]float64{-0.5, -0.5, -0.5},
		[3]float64{-0.5, -0.5, 0.5},
		[3]float64{0.5, -0.5, 0.5},
		[3]float64{0.5, -0.5, -0.5},
	}
	attr.Texcoords = [][2]float64{
		[2]float64{0., 1.},
		[2]float64{0., 0.},
		[2]float64{1., 0.},
		[2]float64{1., 1.},
	}
	attr.Normals = [][3]float64{
		[3]float64{0, 1, 0},
		[3]float64{-1, 0, 0},
		[3]float64{1, 0, 0},
		[3]float64{0, 0, -1},
		[3]float64{0, 0, 1},
		[3]float64{0, -1, 0},
	}

	mtrls := []Mtl{}
	testMtl1, err := os.Open("models/default.mtl")
	if err != nil {
		t.Error(err)
	}
	defer testMtl1.Close()
	mtls := LoadMtl(testMtl1)
	mtrls = append(mtrls, *mtls...)
	testMtl2, err := os.Open("models/default2.mtl")
	if err != nil {
		t.Error(err)
	}
	defer testMtl2.Close()
	mtls = LoadMtl(testMtl2)
	mtrls = append(mtrls, *mtls...)

	shape := ShapeT{}
	shape.mesh = mesh
	shapes := []ShapeT{shape}

	want_ := Obj{}
	want_.Shapes = &shapes
	want_.Materials = &mtrls
	want_.Attribute = &attr
	want := &want_

	have, err := Parse(testObj)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(want, have) {
		t.Error("ParseObj() return bad result")
	}
}
