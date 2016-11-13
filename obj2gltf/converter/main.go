package main

import (
	"github.com/sturfeeinc/glTF/obj2gltf"
	"os"
)

func main() {
	if len(os.Args) > 3 {
		println("necessary path to .obj and destination directory")
		return
	}
	fileObj, err := os.Open(os.Args[1])
	if err != nil {
		println(err)
		return
	}
	err = obj2gltf.Convert(fileObj, os.Args[2])
	if err != nil {
		println(err)
	}
}
