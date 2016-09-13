package obj2gltf

import (
	"io"
	"github.com/sturfeeinc/obj3d"
	"github.com/sturfeeinc/jlTF"
	"encoding/json"
	"os"
	"fmt"
	"encoding/binary"
)

func Convert(r io.Reader) ([]byte, error) {
	obj, err := obj3d.Parse(r)
	if err != nil {
		panic(err)
	}

	g, err := objToglTF(obj)
	if err != nil {

	}

	raw, err := json.Marshal(obj)
	if err != nil {

	}
println(string(raw))
	raw, err = json.Marshal(g)
	if err != nil {

	}
println(string(raw))
	return raw, nil
}

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

