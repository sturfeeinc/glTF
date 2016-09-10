package obj2gltf

import (
	"io"
	"github.com/sturfeeinc/obj3d"
	"github.com/sturfeeinc/jlTF/model"
	"encoding/json"
	"os"
	"fmt"
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

func objToglTF(o *obj3d.Obj) (*model.GlTF, error) {



	fmt.Printf("%f\n", o)
	gltf := model.GlTF{}
	gltf.Scene = "defaultScenes"
	scenes := map[string]*model.Scene{}
	gltf.Scenes = scenes
	scene := model.Scene{}
	scenes["defaultScenes"] = &scene
	scene.Nodes = []*model.GlTFid{}
	nodeName := model.GlTFid("node_1")
	scene.Nodes = append(scene.Nodes, &nodeName)
	nodes :=  map[string]*model.Node{}
	gltf.Nodes = nodes
	node := model.Node{}
	node.Name = o.Name
	meshes_ := []*model.GlTFid{}
	node.Meshes = meshes_

	nodes["node_1"] = &node

	meshName := model.GlTFid("mesh1_1")
	node.Meshes = append(meshes_, &meshName)
	mesh := model.Mesh{}

	primitives := []*model.MeshPrimitive{}
	mesh.Primitives = primitives
	primitive := model.MeshPrimitive{}
	primitive.Attributes = map[string]*model.MeshPrimitiveAttribute{}
	attr := model.MeshPrimitiveAttribute("accessor_id0")
	primitive.Attributes["POSITION"] = &attr
	meshes := map[string]*model.Mesh{}
	gltf.Meshes = meshes
	meshes["mesh1_1"] = &mesh

	mesh.Primitives = append(mesh.Primitives, &primitive)

	// accessors
	gltf.Accessors = map[string]*model.Accessor{}
	accessor := model.Accessor{}
	gltf.Accessors["accessor_id0"] = &accessor

	// bufferViews
	gltf.BufferViews = map[string]*model.BufferView{}
	bufferView := model.BufferView{}
	gltf.BufferViews["bufferViewWithVertices_id"] = &bufferView

	// buffer
	gltf.Buffers = map[string]*model.Buffer{}
	buffer := model.Buffer{}
	gltf.Buffers["buffer_id"] = &buffer


	accessor.BufferView = model.GlTFid("bufferViewWithVertices_id")


	file, err := os.OpenFile("verticies.bin", os.O_CREATE | os.O_RDWR, 0777)
	if err != nil {
		println(err.Error())
	}

	data := fToB(&o.Attribute.Vertices)
	n, err := file.Write(*data)
	if err != nil {
		println(err.Error())
	}
	offset := 0
	accessor.ByteOffset = &offset

	t := 5126
	accessor.ComponentType = &t

	t1 := 1
	accessor.Count = &t1
	ssss := "SCALAR"
	accessor.Type = &ssss

	bufferView.ByteOffset = &offset
	bufferView.ByteLength = &n
	bufM := model.GlTFid("buffer_id")
	bufferView.Buffer = bufM
	buffer.ByteLength = &n
	uri :=  "verticies.bin"
	buffer.Uri = &uri
	return &gltf, nil
}
