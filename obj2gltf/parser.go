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
//[0, 1, 2, 0, 2, 3, 0, 4, 5, 0, 5, 1, 3, 4, 0, 1, 5, 6, 1, 6, 2, 2, 6, 7, 2, 7, 3, 3, 7, 4, 5, 4, 7, 5, 7, 6]
//[1, 2, 3, 1, 3, 4, 1, 5, 6, 1, 6, 2, 3, 7, 8, 3, 8, 4, 4, 8, 5, 4, 5, 1, 2, 6, 7, 2, 7, 3, 6, 5, 8, 6, 8, 7]
//
//[-0.5, 0.5, -0.5, -0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, -0.5, -0.5, -0.5, -0.5, -0.5, -0.5, 0.5, 0.5, -0.5, 0.5, 0.5, -0.5, -0.5]
//[-0.5, 0.5, -0.5, -0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, -0.5, -0.5, -0.5, -0.5, -0.5, -0.5, 0.5, 0.5, -0.5, 0.5, 0.5, -0.5, -0.5]



var (
	WSS = []byte{' '} // whitespace separator
	WT = []byte{'\t'} // table separator
	SL = []byte{'/'} // table separator
)

func isSpace(b byte) bool {
	return b == WS || b == T
}

type Obj struct {
	Name      string
	Positions []float32
	Indices   []int16
	Fc, Vc    int
}

func SturfeeParser (r io.Reader) []Obj {
	file, _ := ioutil.TempFile("", "testCube.bin")
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
			obj.Indices = f
			obj.Positions = v
			objs = append(objs, obj)
			obj = Obj{}
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
			obj.Indices = f
			obj.Positions = v
			objs = append(objs, obj)

			println(len(f))
			println(len(v))

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

		default:
			//println(string(token))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

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

/*



x11,y11,z11 -> p11
x21,y21,z21 -> p21
x31,y31,z31 -> p31
x12,y12,z12 -> p12
x22,y22,z22 -> p22
x32,y32,z32 -> p32

x1,y1,z1 -> p1
x2,y2,z2 -> p2
x3,y3,z3 -> p3
x4,y4,z4 -> p4

f1 -> p1,p4,p3
f2 -> p1,p2,p3

*/