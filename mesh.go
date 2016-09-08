package obj3d


/*
typedef struct {
  std::vector<index_t> indices;
  std::vector<unsigned char> num_face_vertices;  // The number of vertices per
                                                 // face. 3 = polygon, 4 = quad,
                                                 // ... Up to 255.
  std::vector<int> material_ids;                 // per-face material ID
  std::vector<tag_t> tags;                       // SubD tag
} mesh_t;
*/

type MeshT struct {
	Indices         []IndexT
	NumFaceVertices []uint // The number of vertices per
                           // face. 3 = polygon, 4 = quad,
                           // ... Up to 255.
	MaterialIds     []int  // per-face material ID
	Tags            []TagT
}
