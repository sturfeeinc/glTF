package obj3d

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

	v := []float64{}
	vn := []float64{}
	vt := []float64{}

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
			v = append(v,
				parseFloat(string(res[0])),
				parseFloat(string(res[1])),
				parseFloat(string(res[2])),
			)
		// normal
		case token[0] == 'v' && token[1] == 'n' && isSpace(token[2]) :
			res = bytes.Split(token[3:], WSS)
			vn = append(vn,
				parseFloat(string(res[0])),
				parseFloat(string(res[1])),
				parseFloat(string(res[2])),
			)
		// texcoord
		case token[0] == 'v' && token[1] == 't' && isSpace(token[2]) :
			res = bytes.Split(token[3:], WSS)
			vt = append(vt,
				parseFloat(string(res[0])),
				parseFloat(string(res[1])),
			)
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
				shape1.Mesh.Indices = append(shape1.Mesh.Indices, index)
				shape1.Mesh.MaterialIds = append(shape1.Mesh.MaterialIds, currentMaterialId)
				shape1.Mesh.NumFaceVertices = append(shape1.Mesh.NumFaceVertices, uint(len(res)))
				shape = &shape1
			}
		// line
		case token[0] == 'l' && isSpace(token[1]) :
			res = bytes.Split(token[2:], WSS)
			check = 0 // nullify check flag
			line := Line{}
			line.Indices = []IndexT{}
			for _, a := range res {
				index, check_ = parseCouple(a)
				if check != 0 && check != check_ {
					panic(`"f v//vn v/vt/vn v/vt/vn" error`)
				}
				check = check_
				line.Indices = append(line.Indices, index)
			}
			shape1 := *shape
			if shape1.Lines == nil {
				lines := []Line{}
				shape1.Lines = &lines
			}
			lines := *shape1.Lines
			lines = append(lines, line)
			shape1.Lines = &lines
			shape = &shape1
		// use mtl
		case string(token[0:6]) == "usemtl" && isSpace(token[6]) :
			if shape != nil {
				Shapes = append(Shapes, *shape)
			}
			currentMaterialId = getMaterialId(&Materials, string(token[7:]))
			shape1 := ShapeT{}
			shape1.Mesh = MeshT{}
			shape1.Mesh.Indices = []IndexT{}
			shape1.Mesh.MaterialIds = []int{}
			shape1.Mesh.NumFaceVertices = []uint{}
			shape = &shape1
		// load mtl
		case string(token[0:6]) == "mtllib" && isSpace(token[6]) :
			res := bytes.Split(token[7:], WSS)
			for _, name := range res {
				mtrls := mtllib(string(name))
				Materials = append(Materials, *mtrls...)
			}
		// smoothing group
		case token[0] == 's' && isSpace(token[1]) :
			shape1 := *shape
			if token[2] == 'o' &&  token[3] == 'f' &&  token[4] == 'f' {
				shape1.SmoothingGroup = 0.
			} else {
				shape1.SmoothingGroup = parseFloat(string(token[2:]))
			}
			shape = &shape1
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
