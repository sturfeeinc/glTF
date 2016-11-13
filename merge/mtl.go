package obj3d

import (
	"io"
	"bufio"
	"bytes"
)

/*
typedef struct {
  std::string name;

  float ambient[3];
  float diffuse[3];
  float specular[3];
  float transmittance[3];
  float emission[3];
  float shininess;
  float ior;       // index of refraction
  float dissolve;  // 1 == opaque; 0 == fully transparent
  // illumination model (see http://www.fileformat.info/format/material/)
  int illum;

  int dummy;  // Suppress padding warning.

  std::string ambient_texname;             // map_Ka
  std::string diffuse_texname;             // map_Kd
  std::string specular_texname;            // map_Ks
  std::string specular_highlight_texname;  // map_Ns
  std::string bump_texname;                // map_bump, bump
  std::string displacement_texname;        // disp
  std::string alpha_texname;               // map_d

  // PBR extension
  // http://exocortex.com/blog/extending_wavefront_mtl_to_support_pbr
  float roughness;                // [0, 1] default 0
  float metallic;                 // [0, 1] default 0
  float sheen;                    // [0, 1] default 0
  float clearcoat_thickness;      // [0, 1] default 0
  float clearcoat_roughness;      // [0, 1] default 0
  float anisotropy;               // aniso. [0, 1] default 0
  float anisotropy_rotation;      // anisor. [0, 1] default 0
  std::string roughness_texname;  // map_Pr
  std::string metallic_texname;   // map_Pm
  std::string sheen_texname;      // map_Ps
  std::string emissive_texname;   // map_Ke
  std::string normal_texname;     // norm. For normal mapping.

  std::map<std::string, std::string> unknown_parameter;
} material_t;
*/

type Mtl struct {
	Name                      string

	Ambient                   [3]float64
	Diffuse                   [3]float64
	Specular                  [3]float64
	Transmittance             [3]float64
	Emission                  [3]float64

	Shininess                 float64
	Ior                       float64 // index of refraction
	Dissolve                  float64 // 1 == opaque; 0 == fully transparent

	// illumination model (see http://www.fileformat.info/format/material/)
	Illum                     int

	AmbientTextName           string  // map_Ka
	DiffuseTextName           string  // map_Kd
	SpecularTextName          string  // map_Ks
	SpecularHighlightTextName string  // map_Ns
	BumpTextName              string  // map_bump, bump
	DisplacementTextName      string  // disp
	AlphaTextName             string  // map_d

	// PBR extension
	// http://exocortex.com/blog/extending_wavefront_mtl_to_support_pbr
	Roughness                 float64 // [0, 1] default 0
	Metallic                  float64 // [0, 1] default 0
	Sheen                     float64 // [0, 1] default 0
	ClearcoatThickness        float64 // [0, 1] default 0
	ClearcoatRoughness        float64 // [0, 1] default 0
	Anisotropy                float64 // aniso. [0, 1] default 0
	AnisotropyRotation        float64 // anisor. [0, 1] default 0
	RoughnessTextName         string  // map_Pr
	MetallicTextName          string  // map_Pm
	SheenTextName             string  // map_Ps
	EmissiveTextName          string  // map_Ke
	NormalTextName            string  // norm. For normal mapping.


	UnknownParameter          map[string]string
}

func LoadMtl(r io.Reader) *[]Mtl {

	materials := []Mtl{}
	var currentMtl Mtl
	first := false
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var token []byte
	var res [][]byte
	for scanner.Scan() {
		token = bytes.Trim(scanner.Bytes(), " \t")
		switch {
		case len(token) == 0 || token[0] == '#':
		// new mtl
		case string(token[:6]) == "newmtl" && isSpace(token[6]) :
			if first {
				materials = append(materials, currentMtl)
			}
			first = true
			currentMtl = newMtl()
			currentMtl.Name = string(token[7:])


		// ambient
		case token[0] == 'K' && token[1] == 'a' && isSpace(token[2]) :
			res = bytes.Split(token[3:], WSS)
			currentMtl.Ambient = parseRGB(res)
		// diffuse
		case token[0] == 'K' && token[1] == 'd' && isSpace(token[2]) :
			res = bytes.Split(token[3:], WSS)
			currentMtl.Diffuse = parseRGB(res)
		// specular
		case token[0] == 'K' && token[1] == 's' && isSpace(token[2]) :
			res = bytes.Split(token[3:], WSS)
			currentMtl.Specular = parseRGB(res)
		// transmittance
		case token[0] == 'K' && token[1] == 't' && isSpace(token[2]) ||
			token[0] == 'T' && token[1] == 'f' && isSpace(token[2]) :
			res = bytes.Split(token[3:], WSS)
			currentMtl.Transmittance = parseRGB(res)
		// emission
		case token[0] == 'K' && token[1] == 'e' && isSpace(token[2]) :
			res = bytes.Split(token[3:], WSS)
			currentMtl.Emission = parseRGB(res)
		// ior(index of refraction)
		case token[0] == 'N' && token[1] == 'i' && isSpace(token[2]) :
			currentMtl.Ior = parseFloat(string(token[3:]))
		// shininess
		case token[0] == 'N' && token[1] == 's' && isSpace(token[2]) :
			currentMtl.Shininess = parseFloat(string(token[3:]))
		// illum model
		case string(token[0:5]) == "illum" && isSpace(token[5]) :
			currentMtl.Illum = parseInt(string(token[6:]))
		// dissolve
		case token[0] == 'd' && isSpace((token[1])) :
			currentMtl.Dissolve = parseFloat(string(token[2:]))
		case token[0] == 'T' && token[1] == 'r' && isSpace(token[2]) :
			// Invert value of Tr(assume Tr is in range [0, 1])
			currentMtl.Shininess = 1. - parseFloat(string(token[3:]))
		// PBR: roughness
		case token[0] == 'P' && token[1] == 'r' && isSpace((token[2])) :
			currentMtl.Roughness = parseFloat(string(token[3:]))
		// PBR: metallic
		case token[0] == 'P' && token[1] == 'm' && isSpace((token[2])) :
			currentMtl.Metallic = parseFloat(string(token[3:]))
		// PBR: sheen
		case token[0] == 'P' && token[1] == 's' && isSpace((token[2])) :
			currentMtl.Sheen = parseFloat(string(token[3:]))
		// PBR: clearcoat thickness
		case token[0] == 'P' && token[1] == 'c' && isSpace((token[2])) :
			currentMtl.ClearcoatThickness = parseFloat(string(token[3:]))
		// PBR: clearcoat roughness
		case string(token[:3]) == "Pcr" && isSpace(token[3]) :
			currentMtl.ClearcoatRoughness = parseFloat(string(token[4:]))
		// PBR: anisotropy
		case string(token[:5]) == "aniso" && isSpace(token[5]) :
			currentMtl.Anisotropy = parseFloat(string(token[6:]))
		// PBR: anisotropy rotation
		case string(token[:6]) == "anisor" && isSpace(token[6]) :
			currentMtl.AnisotropyRotation = parseFloat(string(token[7:]))
		// ambient texture
		case string(token[:6]) == "map_Ka" && isSpace(token[6]) :
			currentMtl.AmbientTextName = string(token[7:])
		// diffuse texture
		case string(token[:6]) == "map_Kd" && isSpace(token[6]) :
			currentMtl.DiffuseTextName = string(token[7:])
		// specular texture
		case string(token[:6]) == "map_Ks" && isSpace(token[6]) :
			currentMtl.SpecularTextName = string(token[7:])
		// specular highlight texture
		case string(token[:6]) == "map_Ns" && isSpace(token[6]) :
			currentMtl.SpecularHighlightTextName = string(token[7:])
		// alpha texture
		case string(token[:5]) == "map_d" && isSpace(token[5]) :
			currentMtl.AlphaTextName = string(token[6:])
		// bump texture
		case string(token[:8]) == "map_bump" && isSpace(token[8]) :
			currentMtl.BumpTextName = string(token[9:])
		case string(token[:4]) == "bump" && isSpace(token[4]) :
			currentMtl.BumpTextName = string(token[5:])
		// displacement texture
		case string(token[:4]) == "disp" && isSpace(token[4]) :
			currentMtl.DisplacementTextName = string(token[5:])
		// PBR: roughness texture
		case string(token[:6]) == "map_Pr" && isSpace(token[6]) :
			currentMtl.RoughnessTextName = string(token[7:])
		// PBR: metallic texture
		case string(token[:6]) == "map_Pm" && isSpace(token[6]) :
			currentMtl.MetallicTextName = string(token[7:])
		// PBR: sheen texture
		case string(token[:6]) == "map_Ps" && isSpace(token[6]) :
			currentMtl.SheenTextName = string(token[7:])
		// PBR: emissive texture
		case string(token[:6]) == "map_Ke" && isSpace(token[6]) :
			currentMtl.EmissiveTextName = string(token[7:])
		// PBR: normal map texture
		case string(token[:4]) == "norm" && isSpace(token[4]) :
			currentMtl.NormalTextName = string(token[5:])
		// unknown parameter
		default:
			if currentMtl.UnknownParameter == nil {
				currentMtl.UnknownParameter = map[string]string{}
			}
			i := bytes.IndexAny(token, " \t")
			// if it doesn't has parameters. put entire one
			if i == -1 {
				currentMtl.UnknownParameter[string(token)] = string(token)
			} else {
				currentMtl.UnknownParameter[string(token[:i])] = string(string(token[i + 1:]))
			}
		}
	}
	materials = append(materials, currentMtl)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return &materials
}

func newMtl() Mtl {
	mtl := Mtl{}

	mtl.Name = ""

	mtl.AmbientTextName = ""
	mtl.DiffuseTextName = ""
	mtl.SpecularTextName = ""
	mtl.SpecularHighlightTextName = ""
	mtl.BumpTextName = ""
	mtl.DisplacementTextName = ""
	mtl.AlphaTextName = ""


		mtl.Ambient = [3]float64{}
		mtl.Diffuse = [3]float64{}
		mtl.Specular = [3]float64{}
		mtl.Transmittance = [3]float64{}
		mtl.Emission = [3]float64{}

	for i := 0; i < 3; i++ {
		mtl.Ambient[i] = 0.
		mtl.Diffuse[i] = 0.
		mtl.Specular[i] = 0.
		mtl.Transmittance[i] = 0.
		mtl.Emission[i] = 0.
	}

	mtl.Shininess = 1.
	mtl.Ior = 1.
	mtl.Dissolve = 1.
	mtl.Illum = 0.

	mtl.Roughness = 0.
	mtl.Metallic = 0.
	mtl.Sheen = 0.
	mtl.ClearcoatThickness = 0.
	mtl.ClearcoatRoughness = 0.
	mtl.Anisotropy = 0.
	mtl.AnisotropyRotation = 0.

	mtl.RoughnessTextName = ""
	mtl.MetallicTextName = ""
	mtl.SheenTextName = ""
	mtl.EmissiveTextName = ""
	mtl.NormalTextName = ""

	return mtl
}

func getMaterialId (m *[]Mtl, name string) (i int) {
	var mtl Mtl
	for i, mtl = range *m {
		if mtl.Name == name {
			return i
		}
	}
	if len(*m) == 0 {
		mtl = newMtl()
		mtls := []Mtl{mtl}
		m = &mtls
	}
	return 0
}

func defaultMtl() Mtl {
	mtl := newMtl()
	mtl.Name = "default"
	return mtl
}