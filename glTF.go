/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with glTFModelGenerator (https://github.com/sturfeeinc/glTFModelGenerator)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 * generated 2016-09-17 21:21:17.912222485 +0300 MSK
 */

package glTF

// http://json-schema.org/draft-03/schema
// The root object for a glTF asset.
type GlTF struct {
	GlTFProperty
	Animations     map[string]Animation  `json:"animations"`
	Asset          Asset                 `json:"asset"`
	Samplers       map[string]Sampler    `json:"samplers"`
	Scene          GlTFid                `json:"scene,omitempty"`
	Scenes         map[string]Scene      `json:"scenes"`
	ExtensionsUsed []string              `json:"extensionsUsed"`
	Buffers        map[string]Buffer     `json:"buffers"`
	Cameras        map[string]Camera     `json:"cameras"`
	Images         map[string]Image      `json:"images"`
	Nodes          map[string]Node       `json:"nodes"`
	Shaders        map[string]Shader     `json:"shaders"`
	Accessors      map[string]Accessor   `json:"accessors"`
	Materials      map[string]Material   `json:"materials"`
	Programs       map[string]Program    `json:"programs"`
	Skins          map[string]Skin       `json:"skins"`
	BufferViews    map[string]BufferView `json:"bufferViews"`
	Meshes         map[string]Mesh       `json:"meshes"`
	Techniques     map[string]Technique  `json:"techniques"`
	Textures       map[string]Texture    `json:"textures"`
}
