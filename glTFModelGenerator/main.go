package main

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"os/exec"
	"bytes"
	"strconv"
	"time"
)

var HEADER string

func init() {

	timeGenerated := time.Now().String()
	HEADER = `/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with glTFModelGenerator (https://github.com/sturfeeinc/glTFModelGenerator)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 * generated ` + timeGenerated + "\n" +
` */

`
}

//struct template

type schema struct {
	Schema               string `json:"$schema"`
	Description          string `json:"description"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	Extends              extends `json:"extends"`
	Properties           map[string]property `json:"properties"`
	// not used commonly

	// only in extension.schema.json
	AdditionalProperties interface{} `json:"additionalProperties"`
	// some time as description of function
	GltfWebgl            string `json:"gltf_webgl"`
	// some time as types
	Items                string `json:"items"`
	// some time as types
	Dependencies         string `json:"dependencies"`
	// some time as types
	MinLength            string `json:"minLength"`
}

type property struct {
	Default       interface{} `json:"default,omitempty"`
	Type          string `json:"type,omitempty"`
	Extends       extends `json:"extends,omitempty"`
	Items         items `json:"items,omitempty"`
	AddProperties extends `json:"additionalProperties,omitempty"`
	Ref           string `json:"$ref,omitempty"`
	Required      bool `json:"required,omitempty"`
	Min           *float64 `json:"minimum,omitempty"`
	Max           *float64 `json:"maximum,omitempty"`
	Description   string `json:"gltf_detailedDescription,omitempty"`
}

type items struct {
	Type string `json:"type"`
	Ref  string `json:"$ref"`
}

type extends struct {
	Ref string `json:"$ref"`
}

func main() {

	info, err := ioutil.ReadDir("../shemas/")
	if err != nil {
		panic(err)
	}

	for _, o := range info {
		generate(o.Name())
	}

	cmd := exec.Command("go", "fmt", "github.com/sturfeeinc/jlTF")
	var out bytes.Buffer
	cmd.Stdout = &out
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err = cmd.Run()
	if err != nil {
		panic("gofmt: " + err.Error())
	}
}

func publicField(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func getType(p property) string {

	if p.Ref != "" {
		return getTypeName(p.Ref)
	}
	switch p.Type {
	case "integer":
		return "int"
	case "array":
		return "[]" + getType2(p)
	case "object":
		return "map[string]" + getType2(p)
	case "boolean":
		return "bool"
	case "number":
		return "float64"
	case "string":
		return "string"
	case "": // TODO
		return getType2(p)
	}
	panic(p.Type + " : undefined type")
	return p.Type
}

func getType2(p property) string {
	if len(p.AddProperties.Ref) != 0 {
		return getTypeName(p.AddProperties.Ref)
	}
	if len(p.Extends.Ref) != 0 {
		return getTypeName(p.Extends.Ref)
	}
	if len(p.Items.Type) != 0 {
		switch p.Items.Type {
		case "integer":
			return "int"
		case "array":
			return "[]" + getType2(p)
		case "object":
			return "map[string]" + getType2(p)
		case "boolean":
			return "bool"
		case "number":
			return "float64"
		case "string":
			return "string"
		}
	}
	if len(p.Items.Ref) != 0 {
		return getTypeName(p.Items.Ref)
	}
	return "interface{}"
}

func generate(fileName string) {
	file, err := os.Open("../shemas/" + fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var Schema schema

	err = json.NewDecoder(file).Decode(&Schema)

	if err != nil {
		println(fileName)
		println(err)
		//panic(err)
	}

	p := HEADER + "package glTF\n"
	p += "// " + Schema.Schema + "\n"
	p += "// " + Schema.Description + "\n"
	p += "type " + getTypeName(fileName)

	if len(Schema.Properties) == 0 {
		if Schema.Type == "string" {
			p += ` string`
		} else {
			p += ` interface {}`
		}
	} else {
		p += ` struct {
	`

		if len(Schema.Extends.Ref) != 0 {
			p += getTypeName(Schema.Extends.Ref) + "\n"
		}

		if len(Schema.Properties) != 0 {
			for propertyName, pr := range Schema.Properties {

				p += publicField(propertyName) + " " + getType(pr) + " " + structTag(propertyName, pr) + "\n"
			}
		}
		p += "}"
	}



	fileName = getTypeName(fileName)
	fileName = "../" + strings.ToLower(string(fileName[0])) + fileName[1:] + ".go"

	ioutil.WriteFile(fileName, []byte(p), 0777)

	//easyJSON(fileName)
}

func getTypeName(filename string) string {

	eA := strings.Replace(filename, ".schema.json", "", 1)
	arr := strings.Split(eA, ".")
	for i, s := range arr {
		arr[i] = publicField(s)
	}
	return strings.Join(arr, "")
}

func easyJSON(filename string) {

	cmd := exec.Command("easyjson", "-all", filename)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		println(err.Error())
	}
}

func structTag(name string, p property) (s string) {

		s = " `json:\"" + name + "\""
	if p.Default == nil {
		s = " `json:\"" + name + ",omitempty\""
	}
	if p.Required ||
		p.Min != nil ||
		p.Max != nil {
		s += " validator:\""
		before := false
		if p.Required {
			s += "required"
			before = true
		}
		if p.Min != nil {
			if before {
				s += ", "
			}
			s += "gte=" + strconv.FormatFloat(*p.Min, 'f', 10, 64)
			before = true
		}
		if p.Max != nil {
			if before {
				s += ", "
			}
			s += "lte=" + strconv.FormatFloat(*p.Max, 'f', 10, 64)
			before = true
		}
		s += "\""
	}

	s += "`"
	return
}

