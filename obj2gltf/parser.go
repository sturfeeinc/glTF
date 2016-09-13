package obj2gltf

import (
	"io"
	"io/ioutil"
	"bufio"
	"bytes"
	"strings"
	"strconv"
	"math"
	"encoding/binary"
	"time"
	"math/rand"
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

func isSpace(b byte) bool {
	return b == WS || b == T
}

type Obj struct {
	Name string
	V []float32
	F []int16
	Fc, Vc int
}

func SturfeeParser (r io.Reader) []Obj {
	file, _ := ioutil.TempFile("", "data.3d.bin")
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	v := []float32{}
	f := []int16{}
	obj := Obj{}
	objs := []Obj{}
	init := true
	var token []byte
	for scanner.Scan() {
		token = bytes.Trim(scanner.Bytes(), " ")
		switch {
		// empty
		case len(token) == 0:
		// vertices
		case token[0] == 'v' && isSpace(token[1]) :
			res := bytes.Split(token[2:], WSS)
			v = append(v,
				parseFloat32(string(res[0])),
				parseFloat32(string(res[1])),
				parseFloat32(string(res[2])),
			)
			obj.Vc +=3
		// face
		case token[0] == 'f' && isSpace(token[1]) :
			res := bytes.Split(token[2:], WSS)
			if len(res) < 3 {
				continue
			}
			f = append(f,
				parseInt16(string(res[0])),
				parseInt16(string(res[1])),
				parseInt16(string(res[2])),
			)
			obj.Fc +=3

		case token[0] == 'o' && isSpace(token[1]) :
			if init {
				init = false
				obj.Name = string(token[2:])
				continue
			}
			obj.F = f
			obj.V = v
			objs = append(objs, obj)
			obj := Obj{}
			obj.Name = string(token[2:])
			obj.Vc = 0
			obj.Fc = 0
			v = []float32{}
			f = []int16{}
		//str := string(token[2:])
			//println(str)
		default:
			//println(string(token))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}



	println(objs[11].Name)
	return objs
	err := binary.Write(file, binary.LittleEndian, v)
	if err != nil {
		println(err.Error())
	}

	err = binary.Write(file, binary.LittleEndian, f)
	if err != nil {
		println(err.Error())
	}
	file.Close()



	return []Obj{}
}



func Parse(r io.Reader) (int, error) {


	file, _ := ioutil.TempFile("", "data.3d.bin")
	BinWriter := binWriter{}
	BinWriter.w = bufio.NewWriter(file)
	BinWriter.currentTag = "notExist"
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var token []byte
	for scanner.Scan() {
		token = bytes.Trim(scanner.Bytes(), " ")
		switch {
		// empty
		case len(token) == 0:
		// vertices
		case token[0] == 'v' && isSpace(token[1]) :
			BinWriter.handlT(token[2:], "v")
		// normal
		case token[0] == 'v' && token[1] == 'n' && isSpace(token[2]) :
			BinWriter.handleP(token[3:])
		// texcoord
		case token[0] == 'v' && token[1] == 't' && isSpace(token[2]) :
			BinWriter.handleP(token[3:])
		// face
		/*case token[0] == 'f' && isSpace(token[1]) :
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
		// comment
		case token[0] == '#':
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
			name = &str*/
		default:
			//println(string(token))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	/*Shapes = append(Shapes, *shape)
	Attribute.Vertices = v
	Attribute.Texcoords = vt
	Attribute.Normals = vn
	obj := Obj{}
	obj.Name = name
	obj.Attribute = &Attribute
	obj.Materials = &Materials
	obj.Shapes = &Shapes*/

	return 1, nil
}

func proceedFloat32(s string) []byte {
	return Float32bytes(parseFloat32(s))
}



// alias for strconv.ParseFloat with exception
func parseFloat32(s string) float32 {
	f, err := strconv.ParseFloat(strings.Trim(s, " "), 32)
	if err != nil {
		panic(err)
	}
	return float32(f)
}

func Float32bytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytesBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytesBuf, bits)
	return bytesBuf
}

func (w *binWriter) handleP (token []byte) {
	res := bytes.Split(token, WSS)
	bytes_ := []byte{}
	for _, r := range res {
		bytes_ = append(bytes_, proceedFloat32(string(r))...)
	}
	if w.w.Available() < len(bytes_) {
		 w.w.Flush()
	}
	n, err := w.w.Write(bytes_)
	if err != nil {
		panic(err)
	}
	w.n += n
}


type binWriter struct {
	w *bufio.Writer
	n int
	offset int
	currentTag string
	currentID string
	ids []string
	metadata map[string]meta
}

type meta struct {
	length, offset int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (BinWriter *binWriter)  handlT(token []byte, tag string) {
	if BinWriter.currentTag != tag {
		if len(BinWriter.currentID) > 0 {
			Meta := meta{BinWriter.n, BinWriter.offset}
			BinWriter.metadata[BinWriter.currentID] = Meta
			BinWriter.offset += BinWriter.n
			BinWriter.n = 0
		}
		BinWriter.currentID = RandStringRunes(15)
		BinWriter.currentTag = tag
	}
	BinWriter.handleP(token)
}

// alias for strconv.ParseFloat with exception
func parseInt16(s string) int16 {
	f, err := strconv.ParseInt(strings.Trim(s, " "), 10, BS)
	if err != nil {
		panic(err)
	}
	return int16(f)
}