package obj2gltf

import (
	"os"
	"testing"
	"fmt"
)

func TestObjToGlTF(t *testing.T) {
	fileObj, err := os.Open("/home/petr/gocode/src/github.com/sturfeeinc/jlTF/obj2gltf/SF_MinnaFinal300.obj")
	if err != nil {
		t.Error(err)
	}
	raw, err := Convert(fileObj)
	if err != nil {
		t.Error(err)
	}
	if len(raw) == 0 {
		t.Error(err)
	}
}



func DEBUG(i interface{})  {
	fmt.Printf("%f\n", i)
	//fmt.Printf("%v\n", i)
}

