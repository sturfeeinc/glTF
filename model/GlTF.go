/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type GlTF struct {
	GlTFProperty
	Scene          GlTFid                 `json:"scene"`
	Textures       map[string]*Texture    `json:"textures"`
	Cameras        map[string]*Camera     `json:"cameras"`
	Nodes          map[string]*Node       `json:"nodes"`
	Animations     map[string]*Animation  `json:"animations"`
	Buffers        map[string]*Buffer     `json:"buffers"`
	BufferViews    map[string]*BufferView `json:"bufferViews"`
	Images         map[string]*Image      `json:"images"`
	Materials      map[string]*Material   `json:"materials"`
	Samplers       map[string]*Sampler    `json:"samplers"`
	ExtensionsUsed []*string              `json:"extensionsUsed"`
	Accessors      map[string]*Accessor   `json:"accessors"`
	Scenes         map[string]*Scene      `json:"scenes"`
	Skins          map[string]*Skin       `json:"skins"`
	Techniques     map[string]*Technique  `json:"techniques"`
	Asset          Asset                  `json:"asset"`
	Programs       map[string]*Program    `json:"programs"`
	Meshes         map[string]*Mesh       `json:"meshes"`
	Shaders        map[string]*Shader     `json:"shaders"`
}
