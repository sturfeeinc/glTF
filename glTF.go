/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with glTFModelGenerator (https://github.com/sturfeeinc/glTFModelGenerator)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 * generated 2016-09-11 18:55:47.106260962 +0300 MSK
 */

package glTF

// http://json-schema.org/draft-03/schema
// The root object for a glTF asset.
type GlTF struct {
	GlTFProperty
	Nodes          map[string]Node       `json:"nodes"`
	Techniques     map[string]Technique  `json:"techniques"`
	Buffers        map[string]Buffer     `json:"buffers"`
	Animations     map[string]Animation  `json:"animations"`
	Cameras        map[string]Camera     `json:"cameras"`
	Materials      map[string]Material   `json:"materials"`
	Meshes         map[string]Mesh       `json:"meshes"`
	Scene          GlTFid                `json:"scene,omitempty"`
	Scenes         map[string]Scene      `json:"scenes"`
	Accessors      map[string]Accessor   `json:"accessors"`
	Samplers       map[string]Sampler    `json:"samplers"`
	Skins          map[string]Skin       `json:"skins"`
	Images         map[string]Image      `json:"images"`
	Asset          Asset                 `json:"asset"`
	BufferViews    map[string]BufferView `json:"bufferViews"`
	Programs       map[string]Program    `json:"programs"`
	Shaders        map[string]Shader     `json:"shaders"`
	Textures       map[string]Texture    `json:"textures"`
	ExtensionsUsed []string              `json:"extensionsUsed"`
}
