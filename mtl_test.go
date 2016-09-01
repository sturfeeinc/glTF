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
		println("parseRGB() returns a bad result")
		println("We want")
		DEBUG(expectedMtl)
		println("But we have")
		DEBUG(have)
		t.Error("")
	}
}

var testMtl = `# Blender MTL File: 'None'
# Material Count: 1

newmtl None
Ns 0
Ka 0.000000 0.000000 0.000000
Kd 0.8 0.8 0.8
Ks 0.8 0.8 0.8
d 1
illum 2
`

func TestParseMtl(t *testing.T) {
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

	have := LoadMtl(bytes.NewBufferString(testMtl))

	if !reflect.DeepEqual(want, have) {
		println("parseRGB() returns a bad result")
		println("We want")
		DEBUG(want)
		println("But we have")
		DEBUG(have)
		t.Error("")
	}
}

