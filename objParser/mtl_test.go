package obj

import (
	"testing"
	"reflect"
	"bytes"
	"fmt"
)

var newMaterial map[string]interface{} = map[string]interface{}{
	"ambientColor": nil,
	"emissionColor": nil,
	"diffuseColor": nil,
	"specularColor": nil,
	"specularShininess": nil,
	"alpha": nil,
	"ambientColorMap": nil,
	"emissionColorMap": nil,
	"diffuseColorMap": nil,
	"specularColorMap": nil,
	"specularShininessMap": nil,
	"normalMap": nil,
	"alphaMap": nil,
}

func TestCreateMaterial(t *testing.T) {
	if !reflect.DeepEqual(createMaterial(), material(newMaterial)) {
		t.Error("createMaterial() return bad result")
	}
}

func TestGetDefaultMtl(t *testing.T) {
	newMaterial["diffuseColor"] = []float64{0.5, 0.5, 0.5, 1.0}
	if !reflect.DeepEqual(*getDefaultMtl(), material(newMaterial)) {
		t.Error("createMaterial() return bad result")
	}
}

/*
Result from https://github.com/AnalyticalGraphicsInc/obj2gltf/blob/6c0c39761b1fc734fd8b80f597f634d79ec73b46/lib/mtl.js
for default.mtl from http://graphics.cs.williams.edu/data/meshes.xml  
{ default:
   { ambientColor: [ 1, 1, 1, 1 ],
     emissionColor: [ 0, 0, 0, 1 ],
     diffuseColor: [ 1, 1, 1, 0 ],
     specularColor: [ 0.2, 0.2, 0.2, 1 ],
     specularShininess: 10,
     alpha: 0,
     ambientColorMap: undefined,
     emissionColorMap: undefined,
     diffuseColorMap: 'default.png',
     specularColorMap: undefined,
     specularShininessMap: undefined,
     normalMap: undefined,
     alphaMap: undefined
     }
}

*/

var testMaterial material = material{
	"ambientColor": []float64{1, 1, 1, 1 },
	"emissionColor": []float64{0, 0, 0, 1 },
	"diffuseColor": []float64{1, 1, 1, 0 },
	"specularColor": []float64{0.2, 0.2, 0.2, 1 },
	"specularShininess": 10.,
	"alpha": 0.,
	"ambientColorMap": nil,
	"emissionColorMap": nil,
	"diffuseColorMap": "default.png",
	"specularColorMap": nil,
	"specularShininessMap": nil,
	"normalMap": nil,
	"alphaMap": nil,
}

var defaultMtl = `# Default material file.  Created by Morgan McGuire and released into
# the Public Domain on July 16, 2011.
#
# http://graphics.cs.williams.edu/data

newmtl default
  Ns 10.0000
  Ni 1.5000
  Tr 0  illum 2
  Ka 1 1 1
  Kd 1 1 1
  Ks 0.2 0.2 0.2
  Ke 0 0 0
  map_Kd default.png
`

func TestParse(t *testing.T) {
	Materials := materials{}
	Materials["default"] = &testMaterial
	if !reflect.DeepEqual(Parse(bytes.NewBufferString(defaultMtl)), Materials) {
		t.Error("parse() return bad result")
	}
}

// bitSize in alias should be 64
// otherwise we have an artifact
func TestParseFloat(t *testing.T) {
	if ParseFloat("0.2") != 0.2 {
		t.Error("ParseFloat() return bad result")
	}
}

func DEBUG(i interface{})  {
	fmt.Printf("%f\n", i)
	fmt.Printf("%v\n", i)
}

