package obj

import (
	"testing"
	"reflect"
	"bytes"
)

var expectedMtl Mtl = Mtl{
	Name : "",

	AmbientTextName : "",
	DiffuseTextName : "",
	SpecularTextName : "",
	SpecularHighlightTextName : "",
	BumpTextName : "",
	DisplacementTextName : "",
	AlphaTextName : "",

		Ambient : [3]float64{0., 0., 0.},
		Diffuse : [3]float64{0., 0., 0.},
		Specular : [3]float64{0., 0., 0.},
		Transmittance : [3]float64{0., 0., 0.},
		Emission : [3]float64{0., 0., 0.},

	Shininess : 1.,
	Ior : 1.,
	Dissolve : 1.,
	Illum : 0.,

	Roughness : 0.,
	Metallic : 0.,
	Sheen : 0.,
	ClearcoatThickness : 0.,
	ClearcoatRoughness : 0.,
	Anisotropy : 0.,
	AnisotropyRotation : 0.,

	RoughnessTextName : "",
	MetallicTextName : "",
	SheenTextName : "",
	EmissiveTextName : "",
	NormalTextName : "",
}

func TestNewMtl(t *testing.T) {
	have := newMtl()
	if !reflect.DeepEqual(expectedMtl, have) {
		println("newMtl() returns a bad result")
		println("We want")
		DEBUG(expectedMtl)
		println("But we have")
		DEBUG(have)
		t.Error("newMtl() returns a bad result")
	}
}

var testMtl1 = `# Blender MTL File: 'None'
# Material Count: 1

newmtl None
	Ns 0
	Ka 0.000000 0.000000 0.000000
	Kd 0.8 0.8 0.8
	Ks 0.8 0.8 0.8
	d 1
	illum 2
`

func TestLoadMtl1(t *testing.T) {
	mtl := newMtl()
	mtl.Name = "None"
	mtl.Shininess = 0
	mtl.Ambient = [3]float64{0., 0., 0.}
	mtl.Diffuse = [3]float64{0.8, 0.8, 0.8}
	mtl.Specular = [3]float64{0.8, 0.8, 0.8}
	mtl.Dissolve = 1
	mtl.Illum = 2
	want_ := []Mtl{mtl}
	want := &want_

	have := LoadMtl(bytes.NewBufferString(testMtl1))

	if !reflect.DeepEqual(want, have) {
		println("parseRGB() returns a bad result")
		println("We want")
		DEBUG(want)
		println("But we have")
		DEBUG(have)
		t.Error("")
	}
}

var testMtl2 = `# Blender MTL File: 'None'
# Material Count: 1

newmtl TestOne

Ka 0.1 0.1 0.1
Kd 0.2 0.2 0.2
Ks 0.3 0.3 0.3
Kt 0.4 0.4 0.4
Ke 0.5 0.5 0.5

Ni 0.1
Ns 0.2
illum 5
d 0.3

Pr 0.1
Pm 0.2
Ps 0.3
Pc 0.4
Pcr 0.5
aniso 0.6
anisor 0.7

map_Ka AmbientTextName
map_Kd DiffuseTextName
map_Ks SpecularTextName
map_Ns SpecularHighlightTextName
map_bump BumpTextName
disp DisplacementTextName
map_d AlphaTextName

map_Pr RoughnessTextName
map_Pm MetallicTextName
map_Ps SheenTextName
map_Ke EmissiveTextName
norm NormalTextName

unknownParametr some unknown data
trash

newmtl TestSecond

Tr 0.34
bump BumpTextName
`
/*
That test check that every parameter will be handled. Include unknown parameters
*/
func TestLoadMtl2(t *testing.T) {
	mtl := Mtl{
	Name : "TestOne",

	AmbientTextName : "AmbientTextName",
	DiffuseTextName : "DiffuseTextName",
	SpecularTextName : "SpecularTextName",
	SpecularHighlightTextName : "SpecularHighlightTextName",
		DisplacementTextName : "DisplacementTextName",
		AlphaTextName : "AlphaTextName",
		BumpTextName : "BumpTextName",

		Ambient : [3]float64{0.1, 0.1, 0.1},
		Diffuse : [3]float64{0.2, 0.2, 0.2},
		Specular : [3]float64{0.3, 0.3, 0.3},
		Transmittance : [3]float64{0.4, 0.4, 0.4},
		Emission : [3]float64{0.5, 0.5, 0.5},

		Ior : 0.1,
		Shininess : 0.2,
	Dissolve : 0.3,
	Illum : 5,

	Roughness : 0.1,
	Metallic : 0.2,
	Sheen : 0.3,
	ClearcoatThickness : 0.4,
	ClearcoatRoughness : 0.5,
	Anisotropy : 0.6,
	AnisotropyRotation : 0.7,

	RoughnessTextName : "RoughnessTextName",
	MetallicTextName : "MetallicTextName",
	SheenTextName : "SheenTextName",
	EmissiveTextName : "EmissiveTextName",
	NormalTextName : "NormalTextName",
	UnknownParameter: map[string]string{
		"unknownParametr": "some unknown data",
		"trash": "trash",
	},
}

	want_ := []Mtl{mtl}

	mtl = newMtl()
	mtl.Name = "TestSecond"
	mtl.Shininess = 0.6599999999999999 // precision
	mtl.BumpTextName = "BumpTextName"
	want_ = append(want_, mtl)

	want := &want_

	have := LoadMtl(bytes.NewBufferString(testMtl2))

	if !reflect.DeepEqual(want, have) {
		println("TestLoadMtl2() returns a bad result")
		println("We want")
		DEBUG(want)
		println("But we have")
		DEBUG(have)
		t.Error("")
	}
}


func TestTrash (t *testing.T) {
	rawMtl := `
some unknown data before newmtl
I don't know how it should be handled for now just skip it

newmtl TestOne`


	str := ""
	rawMtl, str = str, rawMtl

}
