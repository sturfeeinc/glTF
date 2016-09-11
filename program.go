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
// A shader program, including its vertex and fragment shader, and names of vertex shader attributes.
type Program struct {
	GlTFChildOfRootProperty
	Attributes     []string `json:"attributes"`
	FragmentShader GlTFid   `json:"fragmentShader,omitempty" validator:"required"`
	VertexShader   GlTFid   `json:"vertexShader,omitempty" validator:"required"`
}
