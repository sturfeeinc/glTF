package obj

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