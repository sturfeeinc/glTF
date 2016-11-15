[![GoDoc](https://godoc.org/github.com/sturfeeinc/glTF?status.svg)](https://godoc.org/github.com/sturfeeinc/glTF)

glTF
====

Struct for marshaling and unmarshaling [glTF](https://github.com/KhronosGroup/glTF)

> go get github.com/sturfeeinc/glTF

It's autogenerated code from official work group's specs. **Don't edit it**. Edit the **[generator](https://github.com/sturfeeinc/glTF/blob/master/glTFModelGenerator/main.go)**


Libraries
---------

* **[easyjson](https://github.com/mailru/easyjson)** for the fastest serialization.

* **[validator.v9](https://github.com/go-playground/validator)** for validation (via tags)

Validation
----------

There are not auto validation. You can use [validator.v9](https://github.com/go-playground/validator) explicitly before _json.Marshal_. 

```go 
package main

import (
	"encoding/json"
	"github.com/sturfeeinc/glTF"
	"gopkg.in/go-playground/validator.v9"
)

func main() {

	validate := validator.New()

	gltf := glTF{}

	// ...

	err := validate.Struct(gltf)
	if err != nil {
		// handling
	}
	res, err := json.Marshal(gltf)
	println(res)
}
```

I thought about realization of strict validation in _MarshalJSON_ and _UnmarshalJSON_ for every struct.
 But for now it will to much complex. 

Authors
-------

* Petr Lozhkin <im7mortal@gmail.com>