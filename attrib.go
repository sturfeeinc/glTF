package obj3d

/*
// Vertex attributes
typedef struct {
  std::vector<float> vertices;   // 'v'
  std::vector<float> normals;    // 'vn'
  std::vector<float> texcoords;  // 'vt'
} attrib_t;
*/

type AttribT struct {
	Vertices  []float64	// 'v'
	Normals   []float64	// 'vn'
	Texcoords []float64	// 'vt'
}