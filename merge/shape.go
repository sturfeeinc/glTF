package obj3d

/*
typedef struct {
  std::string name;
  mesh_t mesh;
} shape_t;
*/

type ShapeT struct {
	Name           string
	Mesh           MeshT
	SmoothingGroup float64
	Lines          *[]Line
}
