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
// A vertex or fragment shader.
type Shader struct {
	GlTFChildOfRootProperty
	Uri  string `json:"uri,omitempty" validator:"required"`
	Type int    `json:"type,omitempty" validator:"required"`
}
