package obj2gltf

import (
	"io"
	"github.com/sturfeeinc/obj3d"
	"github.com/sturfeeinc/jlTF"
	"encoding/json"
	"os"
	"encoding/binary"
)

func Convert(r io.Reader) ([]byte, error) {

	g, _ := og()
	raw, err := json.Marshal(g)
	if err != nil {

	}
	println(string(raw))
	return raw, nil
}

func og() (*glTF.GlTF, error) {

	fileObj, err := os.Open("/home/petr/gocode/src/github.com/sturfeeinc/jlTF/obj2gltf/SF_MinnaFinal300.obj")
	if err != nil {
		panic(err)
	}
	objs := SturfeeParser(fileObj)

	gltf := newjlTF()
	gltf.Scene = "defaultScenes"
	scene := glTF.Scene{}
	scene.Nodes = []glTF.GlTFid{}

	bigOne := []byte{}



	for _, o := range objs {



		node := glTF.Node{}
		node.Name = o.Name
		nodeID := "node_" + node.Name
		scene.Nodes = append(scene.Nodes, glTF.GlTFid(nodeID))
		node.Meshes = []glTF.GlTFid{}
		meshID := "mesh_" + node.Name
		node.Meshes = append(node.Meshes, glTF.GlTFid(meshID))
		mesh := glTF.Mesh{}
		mesh.Primitives = []glTF.MeshPrimitive{}
		primitive := glTF.MeshPrimitive{}
		primitive.Attributes = map[string]glTF.MeshPrimitiveAttribute{}
		accessoVID := "accessor_v_" + node.Name
		primitive.Attributes["POSITION"] = glTF.MeshPrimitiveAttribute(accessoVID)
		accessoFID := "accessor_i_" + node.Name
		primitive.Indices = glTF.GlTFid(accessoFID)
		mesh.Primitives = append(mesh.Primitives, primitive)

		bufferViewID := "bufferView_" + node.Name



		// bufferViews
		bufferView := glTF.BufferView{}
		bufferView.Buffer = glTF.GlTFid("buffer_id")
		bufferView.ByteOffset = len(bigOne)

		// v
		p := tob(o.V)
		lengthV := len(p)
		bigOne = append(bigOne, p...)
		// f
		p = tobi(o.F)
		lengthF := len(p)
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

		gltf.Nodes[nodeID] = node
		gltf.Meshes[meshID] = mesh
		gltf.BufferViews[bufferViewID] = bufferView
	}



















	// buffer
	buffer := glTF.Buffer{}
	buffer.ByteLength = len(bigOne)
	buffer.Uri = "verticies.bin"
	buffer.Type = "arraybuffer"

	gltf.Scenes["defaultScenes"] = scene
	gltf.Buffers["buffer_id"] = buffer
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
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint16(bs, uint16(oo))
		d = append(d, bs...)
	}
	return d
}







/*
func objToglTF(o *obj3d.Obj) (*glTF.GlTF, error) {



	fmt.Printf("%f\n", o)
	gltf := newjlTF()
	gltf.Scene = "defaultScenes"
	scene := glTF.Scene{}
	scene.Nodes = []glTF.GlTFid{glTF.GlTFid("node_1")}
	node := glTF.Node{}
	node.Name = *o.Name
	node.Meshes = []glTF.GlTFid{}
	node.Meshes = append(node.Meshes, glTF.GlTFid("mesh1_1"))
	mesh := glTF.Mesh{}
	mesh.Primitives = []glTF.MeshPrimitive{}
	primitive := glTF.MeshPrimitive{}
	primitive.Attributes = map[string]glTF.MeshPrimitiveAttribute{}
	primitive.Attributes["POSITION"] = glTF.MeshPrimitiveAttribute("accessor_id0")
	primitive.Indices = glTF.GlTFid("accessor_id1")
	mesh.Primitives = append(mesh.Primitives, primitive)











	file, err := os.OpenFile("verticies.bin", os.O_CREATE | os.O_RDWR, 0777)
	if err != nil {
		println(err.Error())
	}

	data := fToB(&o.Attribute.Vertices)
	n1, err := file.Write(*data)
	if err != nil {
		println(err.Error())
	}
	shape := *o.Shapes
	data1 := handle(shape[0].Mesh.Indices)
	binary.Write(file, binary.LittleEndian, data1)


	fileObj, err := os.Open("/home/petr/gocode/src/github.com/sturfeeinc/jlTF/obj2gltf/SF_MinnaFinal300.obj")
	if err != nil {
		panic(err)
	}
	_ = SturfeeParser(fileObj)


	// accessors
	accessor := glTF.Accessor{}
	accessor.BufferView = glTF.GlTFid("bufferViewWithVertices_id")
	accessor.ByteOffset = 0
	accessor.ComponentType = 5126
	accessor.Count = 1
	accessor.Type = "SCALAR"
	gltf.Accessors["accessor_id0"] = accessor
	accessor = glTF.Accessor{}
	accessor.BufferView = glTF.GlTFid("bufferViewWithVertices_id")
	accessor.ByteOffset = 0
	accessor.ComponentType = 5126
	accessor.Count = 1
	accessor.Type = "SCALAR"
	gltf.Accessors["accessor_id1"] = accessor

	// bufferViews
	bufferView := glTF.BufferView{}
	bufferView.Buffer = glTF.GlTFid("buffer_id")
	bufferView.ByteOffset = 0
	bufferView.ByteLength = n1

	// buffer
	buffer := glTF.Buffer{}
	buffer.ByteLength = n1
	buffer.Uri = "verticies.bin"
	buffer.Type = "arraybuffer"

	gltf.Scenes["defaultScenes"] = scene
	gltf.Nodes["node_1"] = node
	gltf.Meshes["mesh1_1"] = mesh
	gltf.BufferViews["bufferViewWithVertices_id"] = bufferView
	gltf.Buffers["buffer_id"] = buffer
	return &gltf, nil
}
*/

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

func handle(a []obj3d.IndexT) (b []int) {
	b = []int{}
	for _, u := range a {
		b = append(b, *u.VertexIndex)
	}
	return b
}

