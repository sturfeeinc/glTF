package obj

import (
	"regexp"
	"os"
	"io"
	"bufio"
	"bytes"
	"math"
	"strconv"
)

const (
	// v float float float
	vertexPattern = `v( +[\d|\.|\+|\-|e|E]+)( +[\d|\.|\+|\-|e|E]+)( +[\d|\.|\+|\-|e|E]+)`
	// vn float float float
	normalPattern = `vn( +[\d|\.|\+|\-|e|E]+)( +[\d|\.|\+|\-|e|E]+)( +[\d|\.|\+|\-|e|E]+)`
	// vt float float
	uvPattern = `vt( +[\d|\.|\+|\-|e|E]+)( +[\d|\.|\+|\-|e|E]+)`
	// f vertex vertex vertex ...
	facePattern1 = `f( +-?\d+)\/?( +-?\d+)\/?( +-?\d+)\/?( +-?\d+)?\/?`
	// f vertex/uv vertex/uv vertex/uv ...
	facePattern2 = `f( +(-?\d+)\/(-?\d+)\/?)( +(-?\d+)\/(-?\d+)\/?)( +(-?\d+)\/(-?\d+)\/?)( +(-?\d+)\/(-?\d+)\/?)?`
	// f vertex/uv/normal vertex/uv/normal vertex/uv/normal ...
	facePattern3 = `f( +(-?\d+)\/(-?\d+)\/(-?\d+))( +(-?\d+)\/(-?\d+)\/(-?\d+))( +(-?\d+)\/(-?\d+)\/(-?\d+))( +(-?\d+)\/(-?\d+)\/(-?\d+))?`
	// f vertex//normal vertex//normal vertex//normal ...
	facePattern4 = `f( +(-?\d+)\/\/(-?\d+))( +(-?\d+)\/\/(-?\d+))( +(-?\d+)\/\/(-?\d+))( +(-?\d+)\/\/(-?\d+))?`
	// f vertex//normal vertex//normal vertex//normal ...
	usemtlPattern = `^usemtl `
)

var (
	vertexRegex = regexp.MustCompile(vertexPattern)
	normalRegex = regexp.MustCompile(normalPattern)
	uvRegex = regexp.MustCompile(uvPattern)
	face1Regex = regexp.MustCompile(facePattern1)
	face2Regex = regexp.MustCompile(facePattern2)
	face3Regex = regexp.MustCompile(facePattern3)
	face4Regex = regexp.MustCompile(facePattern4)
	usemtlRegex = regexp.MustCompile(usemtlPattern)
)

type Obj struct {
	VertexCount    int
	VertexArray    []float64
	PositionMin    []float64
	PositionMax    []float64
	HasUVs         bool
	HasNormals     bool
	MaterialGroups map[string]interface{}
	Materials      materials

	Info           info
	Images         images
	path           string
}

func Parse1() {

	file, _ := os.Open("")
	parse2(file)

}

func parse2(r io.Reader) flow {

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var line []byte
	var res [][][]byte
	var arr [][]byte
	f := flow{}
	f.currentIndexArray = &[]int{}
	f.vertexCache = []int{}
	f.positions = []float64{}
	f.normals = []float64{}
	f.uvs = []float64{}
	f.materialGroups = map[string][]int{}
	for scanner.Scan() {
		line = bytes.Trim(scanner.Bytes(), " ")
		switch {
		case len(line) == 0 || bytes.ContainsRune(line, '#'):
		case vertexRegex.Match(line) :
			res = vertexRegex.FindAllSubmatch(line, -1)
			// string() booooO!!!!
			f.positions = append(f.positions,
				ParseFloat(string(res[0][1])),
				ParseFloat(string(res[0][2])),
				ParseFloat(string(res[0][3])),
			)
		case normalRegex.Match(line) :
			res = normalRegex.FindAllSubmatch(line, -1)
			nx, ny, nz := Cartesian3Normalize(
				ParseFloat(string(res[0][1])),
				ParseFloat(string(res[0][2])),
				ParseFloat(string(res[0][3])),
			)
			f.normals = append(f.normals, nx, ny, nz)
		case uvRegex.Match(line) :
			res = uvRegex.FindAllSubmatch(line, -1)
			f.uvs = append(f.uvs,
				ParseFloat(string(res[0][1])),
				ParseFloat(string(res[0][2])),
			)
		case face1Regex.Match(line) :
			res = face1Regex.FindAllSubmatch(line, -1)
			arr = reformatSubArray(res)
			f.addFace(
				ParseInt(string(arr[0])), ParseInt(string(arr[0])), nil, nil,
				ParseInt(string(arr[1])), ParseInt(string(arr[1])), nil, nil,
				ParseInt(string(arr[2])), ParseInt(string(arr[2])), nil, nil,
				ParseInt(string(arr[3])), ParseInt(string(arr[3])), nil, nil,
			)
		case face2Regex.Match(line) :
			res = face2Regex.FindAllSubmatch(line, -1)
			arr = reformatSubArray(res)
			f.addFace(
				ParseInt(string(arr[0])), ParseInt(string(arr[1])), ParseInt(string(arr[2])), nil,
				ParseInt(string(arr[3])), ParseInt(string(arr[4])), ParseInt(string(arr[5])), nil,
				ParseInt(string(arr[6])), ParseInt(string(arr[7])), ParseInt(string(arr[8])), nil,
				ParseInt(string(arr[9])), ParseInt(string(arr[10])), ParseInt(string(arr[11])), nil,
			)
		case face3Regex.Match(line) :
			res = face3Regex.FindAllSubmatch(line, -1)
			arr = reformatSubArray(res)
			println(string(line))
			println(len(res[0]))
			f.addFace(
				ParseInt(string(arr[0])), ParseInt(string(arr[1])), ParseInt(string(arr[2])), ParseInt(string(arr[3])),
				ParseInt(string(arr[4])), ParseInt(string(arr[5])), ParseInt(string(arr[6])), ParseInt(string(arr[7])),
				ParseInt(string(arr[8])), ParseInt(string(arr[9])), ParseInt(string(arr[10])), ParseInt(string(arr[11])),
				ParseInt(string(arr[12])), ParseInt(string(arr[13])), ParseInt(string(arr[14])), ParseInt(string(arr[15])),
			)
		case face4Regex.Match(line) :
			res = face4Regex.FindAllSubmatch(line, -1)
			arr = reformatSubArray(res)
			f.addFace(
				ParseInt(string(arr[0])), ParseInt(string(arr[1])), nil, ParseInt(string(arr[2])),
				ParseInt(string(arr[3])), ParseInt(string(arr[4])), nil, ParseInt(string(arr[5])),
				ParseInt(string(arr[6])), ParseInt(string(arr[7])), nil, ParseInt(string(arr[8])),
				ParseInt(string(arr[9])), ParseInt(string(arr[10])), nil, ParseInt(string(arr[11])),
			)
		case usemtlRegex.Match(line) :
			line = bytes.Trim(line[7:], " ")
			f.useMaterial(string(line))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return f
}

type flow struct {
	currentIndexArray *[]int
	materials         materials
	vertexCache       []int
	positions         []float64
	normals           []float64
	uvs               []float64
	positionMin       []float64
	materialGroups    map[string][]int
	positionMax       []float64
	vertexArray       []float64
	vertexCount       *int
	hasNormals        bool
	hasUVs            bool
}

func (f *flow) getOffset(i *int, l int, components int) int {
	if *i < 0 {
		// Negative vertex indexes reference the vertices immediately above it
		return (l / components + *i) * components;
	}
	return (*i - 1) * components;
}

func (f *flow) createVertex(p, u, n *int) {
	// Positions
	pi := f.getOffset(p, len(f.positions), 3);

	var px = f.positions[pi + 0];
	var py = f.positions[pi + 1];
	var pz = f.positions[pi + 2];

	f.positionMin[0] = math.Min(px, f.positionMin[0]);
	f.positionMin[1] = math.Min(py, f.positionMin[1]);
	f.positionMin[2] = math.Min(pz, f.positionMin[2]);
	f.positionMax[0] = math.Max(px, f.positionMax[0]);
	f.positionMax[1] = math.Max(py, f.positionMax[1]);
	f.positionMax[2] = math.Max(pz, f.positionMax[2]);
	f.vertexArray = append(f.vertexArray, px, py, pz)

	// Normals
	if f.hasNormals {
		var ni = f.getOffset(n, len(f.normals), 3);
		var nx = f.normals[ni + 0];
		var ny = f.normals[ni + 1];
		var nz = f.normals[ni + 2];
		f.vertexArray = append(f.vertexArray, nx, ny, nz)
	}

	// UVs
	if f.hasUVs {
		if u != nil {
			var ui = f.getOffset(n, len(f.uvs), 3);
			var ux = f.uvs[ui + 0];
			var uy = f.uvs[ui + 1];
			// Flip y so 0.0 is the bottom of the image
			uy = 1.0 - uy;
			f.vertexArray = append(f.vertexArray, ux, uy)
		} else {
			// Some objects in the model may not have uvs, fill with 0's for consistency
			f.vertexArray = append(f.vertexArray, 0.0, 0.0)
		}
	}
}

func (f *flow)addVertex(v, p, u, n *int) int {
	var index int
	if *v >= 0 && *v < len(f.vertexCache) {
		index = f.vertexCache[*v];
	} else {
		index = *f.vertexCount
		*f.vertexCount++
		f.vertexCache[*v] = index;
		f.createVertex(p, u, n);

	}
	return index
}

func (f *flow)addFace(v1, p1, u1, n1, v2, p2, u2, n2, v3, p3, u3, n3, v4, p4, u4, n4 *int) {
	var index1 = f.addVertex(v1, p1, u1, n1);
	var index2 = f.addVertex(v2, p2, u2, n2);
	var index3 = f.addVertex(v3, p3, u3, n3);

	arr := append(*f.currentIndexArray, index1, index2, index3)
	f.currentIndexArray = &arr
	// Triangulate if the face is a quad
	if v4 != nil {
		var index4 = f.addVertex(v4, p4, u4, n4);
		arr := append(*f.currentIndexArray, index1, index3, index4)
		f.currentIndexArray = &arr
	}
}

func Cartesian3Normalize(x, y, z float64) (x_, y_, z_ float64) {
	magnitude := math.Sqrt(x * x + y * y + z * z)
	x_ = x / magnitude
	y_ = y / magnitude
	z_ = z / magnitude
	return
}

// alias for strconv.ParseFloat with exception
func ParseInt(s string) *int {
	f, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic(err)
	}
	f_ := int(f)
	return &f_
}

func (f *flow) useMaterial(mtl string) {
	if len(mtl) == 0 {
		f.useDefaultMaterial();
	} else {
		if _, ok := f.materials[mtl]; !ok {
			f.currentIndexArray = &[]int{}
			f.materialGroups[mtl] = *f.currentIndexArray;
		}
	}
}

func (f *flow) useDefaultMaterial() {
	defaultMaterial := "czmDefaultMat"
	if _, ok := f.materials[defaultMaterial]; !ok {
		f.materials[defaultMaterial] = getDefaultMtl()
	}
	f.useMaterial(defaultMaterial);
}

/*
f -8/-4/-6 -7/-3/-6 -6/-2/-6
 -8/-4/-6
-8
-4
-6
 -7/-3/-6
-7
-3
-6
 -6/-2/-6
-6
-2
-6
*/
func reformatSubArray(a [][][]byte) [][]byte {
	arr := [][]byte{}
	a = a[1:]

	for i := 0; i < len(a); i += 4 {
		arr = append(arr, arr[i + 1], arr[i + 2], arr[i + 3])
	}

	return arr
}

