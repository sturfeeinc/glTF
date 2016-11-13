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
// Geometry to be rendered with the given material.
type MeshPrimitive struct {
	GlTFProperty
	Attributes map[string]MeshPrimitiveAttribute `json:"attributes"`
	Indices    GlTFid                            `json:"indices,omitempty"`
	Material   GlTFid                            `json:"material,omitempty" validator:"required"`
	Mode       int                               `json:"mode"`
}
