package obj2gltf

import (
	"os"
	"testing"
)

func TestP(t *testing.T) {
	fileObj, err := os.Open("/home/petr/gocode/src/github.com/sturfeeinc/jlTF/obj2gltf/SF_MinnaFinal300.obj")
	if err != nil {
		t.Error(err)
	}
	SturfeeParser(fileObj)
}




