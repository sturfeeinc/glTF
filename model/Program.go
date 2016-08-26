/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Program struct {
	GlTFChildOfRootProperty
	Attributes     []*string `json:"attributes"`
	FragmentShader GlTFid    `json:"fragmentShader" validator:"required"`
	VertexShader   GlTFid    `json:"vertexShader" validator:"required"`
}
