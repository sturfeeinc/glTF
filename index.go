package obj

/*
typedef struct {
  int vertex_index;
  int normal_index;
  int texcoord_index;
} index_t;
*/

type IndexT struct {
	VertexIndex   *int
	TexcoordIndex *int
	NormalIndex   *int
}
