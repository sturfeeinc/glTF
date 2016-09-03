package obj

import (
	"io"
	"bufio"
	"bytes"
	"os"
)

const (
	WS = byte(' ')
	T = '\t'
	BS = 64 // bitSize for strconv.Parse
)

var (
	WSS = []byte{' '} // whitespace separator
	WT = []byte{'\t'} // table separator
)

// prototype
type Obj struct {
	Name *string
	Meshes    *[]MeshT
	Shapes    *[]ShapeT
	Materials *Materials
	Attribute *AttribT

	vertexIndex vertexIndex
}

func Parse(r io.Reader) (*Obj, error) {

	obj := Obj{}

	obj.Attribute = &AttribT{}
	obj.Materials = &[]Mtl{}
	obj.Meshes = &[]MeshT{}

	v := [][3]float64{}
	vn := [][3]float64{}
	vt := [][2]float64{}
	face := []face{}
	//tags := []TagT{}

	//var name string

	// material
	//materialMap := map[string]int{}
	//var material int = -1

	var shape ShapeT
	var mesh MeshT
	var currentMaterialId int
	var index IndexT


	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var token []byte
	var res [][]byte
	for scanner.Scan() {
		token = bytes.Trim(scanner.Bytes(), " ")

		switch {
		case len(token) == 0 || token[0] == '#':
		// vertex
		case token[0] == 'v' && isSpace(token[1]) :
			res = bytes.Split(token[2:], WSS)
			v = append(v, [3]float64{
				parseFloat(string(res[0])),
				parseFloat(string(res[1])),
				parseFloat(string(res[2])),
			})
		// normal
		case token[0] == 'v' && token[1] == 'n' && isSpace(token[2]) :
			res = bytes.Split(token[3:], WSS)
			vn = append(vn, [3]float64{
				parseFloat(string(res[0])),
				parseFloat(string(res[1])),
				parseFloat(string(res[2])),
			})
		// texcoord
		case token[0] == 'v' && token[1] == 't' && isSpace(token[2]) :
			res = bytes.Split(token[3:], WSS)
			vt = append(vt, [2]float64{
				parseFloat(string(res[0])),
				parseFloat(string(res[1])),
			})
		// face
		case token[0] == 'f' && isSpace(token[1]) :
			res = bytes.Split(token[2:], WSS)
			for _, a := range res {
				index = obj.parseTriple(a)
				mesh.indices = append(mesh.indices, index)
				mesh.material_ids = append(mesh.material_ids, currentMaterialId)
				mesh.num_face_vertices = append(mesh.num_face_vertices, uint(len(res)))
			}
		// use mtl
		case string(token[0:5]) == "usemtl" && isSpace(token[6]) :
			currentMaterialId = obj.Materials.getMaterialId(string(token[7:]))
			/*if shape != nil {
				obj.Shapes = append(obj.Shapes, shape)
			}
			shape = ShapeT{}
			shape.mesh.material_ids*/

		// load mtl
		case string(token[0:6]) == "mtllib" && isSpace(token[6]) :
			res := bytes.Split(token[7:], WSS)
			for _, name := range res {
				obj.Materials = mtllib(obj.Materials, string(name))
			}
		// object name/id
		case token[0] == 'o' && isSpace(token[1]) :
			name := string(token[2:])
			obj.Name = &name
		}



	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}


	return &obj, nil
}



/*

static vertex_index parseTriple(const char **token, int vsize, int vnsize,
                                int vtsize) {
  vertex_index vi(-1);

  vi.v_idx = fixIndex(atoi((*token)), vsize);
  (*token) += strcspn((*token), "/ \t\r");
  if ((*token)[0] != '/') {
    return vi;
  }
  (*token)++;

  // i//k
  if ((*token)[0] == '/') {
    (*token)++;
    vi.vn_idx = fixIndex(atoi((*token)), vnsize);
    (*token) += strcspn((*token), "/ \t\r");
    return vi;
  }

  // i/j/k or i/j
  vi.vt_idx = fixIndex(atoi((*token)), vtsize);
  (*token) += strcspn((*token), "/ \t\r");
  if ((*token)[0] != '/') {
    return vi;
  }

  // i/j/k
  (*token)++;  // skip '/'
  vi.vn_idx = fixIndex(atoi((*token)), vnsize);
  (*token) += strcspn((*token), "/ \t\r");
  return vi;
}*/


// Parse triples with index offsets: i, i/j/k, i//k, i/j
func (obj *Obj) parseTriple(token []byte) (f IndexT){
	/*complicated := bytes.ContainsRune(token, '/')
	if !complicated {
		ty := parseInt(string(token))
		f.v = &ty
		return
	}*/
	return IndexT{}
}

func mtllib(mtls *[]Mtl, mtlFileName string) *[]Mtl {
	mtlFile, err := os.Open(mtlFileName)
	if err != nil {
		panic("mtllib file doesn't exist")
	}
	defer mtlFile.Close()
	materials := LoadMtl(mtlFile)
	mtrls := append(*mtls, *materials...)
	return &mtrls
}
