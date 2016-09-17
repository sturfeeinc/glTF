package main

import (
	"github.com/sturfeeinc/glTF/obj2gltf"
	"os"
)

func main() {
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
