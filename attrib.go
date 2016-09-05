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
	Vertices  [][3]float64	// 'v'
	Normals   [][3]float64	// 'vn'
	Texcoords [][2]float64	// 'vt'
}