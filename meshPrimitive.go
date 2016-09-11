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
// Geometry to be rendered with the given material.
type MeshPrimitive struct {
	GlTFProperty
	Mode       int                               `json:"mode"`
	Attributes map[string]MeshPrimitiveAttribute `json:"attributes"`
	Indices    GlTFid                            `json:"indices,omitempty"`
	Material   GlTFid                            `json:"material,omitempty" validator:"required"`
}
