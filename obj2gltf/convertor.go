package obj2gltf

import (
	"io"
	"github.com/sturfeeinc/obj3d"
	"github.com/sturfeeinc/glTF"
	"encoding/json"
	"os"
	"encoding/binary"
	"io/ioutil"
	"fmt"
)

func Convert(r io.Reader, outputDir string) (error) {

	g, _ := og()
	raw, err := json.Marshal(g)
	if err != nil {

	}
	println(string(raw))
	return nil
}

func og() (*glTF.GlTF, error) {

	fileObj, err := os.Open("/home/petr/gocode/src/github.com/sturfeeinc/glTF/obj2gltf/cube.obj")
	if err != nil {
		panic(err)
	}
	objs := SturfeeParser(fileObj)

	gltf := newjlTF()
	gltf.Scene = "defaultScenes"
	scene := glTF.Scene{}
	scene.Nodes = []glTF.GlTFid{}

	bigOne := []byte{}

	gltf.Materials["defaultMtl"] = getDefMtl()

	for _, o := range objs {
		node := glTF.Node{}
		node.Name = o.Name
		nodeID := "node_" + node.Name
		scene.Nodes = append(scene.Nodes, glTF.GlTFid(nodeID))
		node.Matrix = []float64{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}
		node.Rotation = []float64{0,0,0,1}
		node.Scale = []float64{1,1,1}
		node.Translation = []float64{0,0,0}
		node.Meshes = []glTF.GlTFid{}
		meshID := "mesh_" + node.Name
		node.Meshes = append(node.Meshes, glTF.GlTFid(meshID))
		mesh := glTF.Mesh{}
		mesh.Primitives = []glTF.MeshPrimitive{}
		primitive := glTF.MeshPrimitive{}
		primitive.Attributes = map[string]glTF.MeshPrimitiveAttribute{}
		accessoVID := "accessor_v_" + node.Name
		primitive.Attributes["POSITION"] = glTF.MeshPrimitiveAttribute(accessoVID)
		primitive.Mode = 4
		accessoFID := "accessor_i_" + node.Name
		primitive.Indices = glTF.GlTFid(accessoFID)
		primitive.Material = glTF.GlTFid("defaultMtl")
		mesh.Primitives = append(mesh.Primitives, primitive)

		bufferViewID := "bufferView_v_" + node.Name



		// bufferViews
		bufferView := glTF.BufferView{}
		bufferView.Buffer = glTF.GlTFid("buffer_id")
		bufferView.ByteOffset = len(bigOne)

		// v
		p := tob(o.Positions)
		lengthV := len(p)
		bigOne = append(bigOne, p...)
		// f
		p = tobi(o.Indices)
		lengthF := len(p)
		fmt.Printf("%f\n", p)
		bigOne = append(bigOne, p...)

		bufferView.ByteLength = lengthF + lengthV
		// accessors

		accessor := glTF.Accessor{}
		accessor.BufferView = glTF.GlTFid(bufferViewID)
		accessor.ByteOffset = 0
		accessor.ComponentType = 5126
		accessor.Count = o.Vc
		accessor.Type = "SCALAR"
		gltf.Accessors[accessoVID] = accessor


		accessor = glTF.Accessor{}
		accessor.BufferView = glTF.GlTFid(bufferViewID)
		accessor.ByteOffset = lengthV
		accessor.ComponentType = 5123
		accessor.Count = o.Fc
		accessor.Type = "SCALAR"
		gltf.Accessors[accessoFID] = accessor
		gltf.BufferViews[bufferViewID] = bufferView

		gltf.Nodes[nodeID] = node
		gltf.Meshes[meshID] = mesh
	}



















	// buffer
	buffer := glTF.Buffer{}
	buffer.ByteLength = len(bigOne)
	buffer.Uri = "test.bin"
	buffer.Type = "arraybuffer"

	gltf.Scenes["defaultScenes"] = scene
	gltf.Buffers["buffer_id"] = buffer
	file, _ := ioutil.TempFile("", "test.bin")
	err = binary.Write(file, binary.LittleEndian, bigOne)
	if err != nil {
		println(err.Error())
	}
	file.Close()

	return &gltf, nil
}

func tob(b []float32) []byte {
	d := []byte{}
	for _, oo := range b {
		d = append(d, Float32bytes(oo)...)
	}
	return d
}

func tobi(b []int16) []byte {
	d := []byte{}

	for _, oo := range b {
		bs := make([]byte, 2)
		println(uint16(oo), oo)
		binary.LittleEndian.PutUint16(bs, uint16(oo))

		fmt.Printf("%f\n", bs)
		d = append(d, bs...)
	}
	return d
}

func newjlTF() (gltf glTF.GlTF) {
	gltf = glTF.GlTF{}
	gltf.Nodes = map[string]glTF.Node{}
	gltf.Techniques = map[string]glTF.Technique{}
	gltf.Buffers = map[string]glTF.Buffer{}
	gltf.Animations = map[string]glTF.Animation{}
	gltf.Cameras = map[string]glTF.Camera{}
	gltf.Materials = map[string]glTF.Material{}
	gltf.Meshes = map[string]glTF.Mesh{}
	gltf.Scenes = map[string]glTF.Scene{}
	gltf.Accessors = map[string]glTF.Accessor{}
	gltf.Samplers = map[string]glTF.Sampler{}
	gltf.Skins = map[string]glTF.Skin{}
	gltf.Images = map[string]glTF.Image{}
	gltf.Asset = glTF.Asset{}
	gltf.Asset.Version = "1.0.3"
	gltf.Asset.Profile = glTF.AssetProfile{}
	gltf.Asset.Profile.Api = "WebGL"
	gltf.Asset.Profile.Version = "1.0.3"
	gltf.BufferViews = map[string]glTF.BufferView{}
	gltf.Programs = map[string]glTF.Program{}
	gltf.Shaders = map[string]glTF.Shader{}
	gltf.Textures = map[string]glTF.Texture{}
	gltf.ExtensionsUsed = []string{}
	return gltf
}

func getDefMtl() (mtl glTF.Material) {
	mtl = glTF.Material{}
	mtl.Values = map[string]glTF.MaterialValues{}
	return
}

func handle(a []obj3d.IndexT) (b []int) {
	b = []int{}
	for _, u := range a {
		b = append(b, *u.VertexIndex)
	}
	return b
}

