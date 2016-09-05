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
	SL = []byte{'/'} // table separator
)

// prototype
type Obj struct {
	Name *string
	Shapes    *[]ShapeT
	Materials *[]Mtl
	Attribute *AttribT
}

func Parse(r io.Reader) (*Obj, error) {

	var name *string
	Attribute := AttribT{}
	Materials := []Mtl{}
	Shapes := []ShapeT{}

	v := [][3]float64{}
	vn := [][3]float64{}
	vt := [][2]float64{}

	var shape *ShapeT
	var currentMaterialId int
	var index IndexT
	var check, check_ int


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
			check = 0 // nullify check flag
			for _, a := range res {
				index, check_ = parseTriple(a)
				if check != 0 && check != check_ {
					panic(`"f v//vn v/vt/vn v/vt/vn" error`)
				}
				check = check_
				shape1 := *shape
				shape1.mesh.indices = append(shape1.mesh.indices, index)
				shape1.mesh.material_ids = append(shape1.mesh.material_ids, currentMaterialId)
				shape1.mesh.num_face_vertices = append(shape1.mesh.num_face_vertices, uint(len(res)))
				shape = &shape1
			}
		// use mtl
		case string(token[0:6]) == "usemtl" && isSpace(token[6]) :
			if shape != nil {
				Shapes = append(Shapes, *shape)
			}
			currentMaterialId = getMaterialId(&Materials, string(token[7:]))
			shape1 := ShapeT{}
			shape1.mesh = MeshT{}
			shape1.mesh.indices = []IndexT{}
			shape1.mesh.material_ids = []int{}
			shape1.mesh.num_face_vertices = []uint{}
			shape = &shape1
		// load mtl
		case string(token[0:6]) == "mtllib" && isSpace(token[6]) :
			res := bytes.Split(token[7:], WSS)
			for _, name := range res {
				mtrls := mtllib(string(name))
				Materials = append(Materials, *mtrls...)
			}
		// object name/id
		case token[0] == 'o' && isSpace(token[1]) :
			str := string(token[2:])
			name = &str
		default:
			//println(string(token))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	Shapes = append(Shapes, *shape)
	Attribute.Vertices = v
	Attribute.Texcoords = vt
	Attribute.Normals = vn
	obj := Obj{}
	obj.Name = name
	obj.Attribute = &Attribute
	obj.Materials = &Materials
	obj.Shapes = &Shapes

	return &obj, nil
}

func mtllib(mtlFileName string) *[]Mtl {
	mtlFile, err := os.Open(mtlFileName)
	if err != nil {
		panic("mtllib file doesn't exist")
	}
	defer mtlFile.Close()
	return LoadMtl(mtlFile)
}
