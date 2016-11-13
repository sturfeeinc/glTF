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
// A typed view into a bufferView.  A bufferView contains raw binary data.  An accessor provides a typed view into a bufferView or a subset of a bufferView similar to how WebGL's `vertexAttribPointer()` defines an attribute in a buffer.
type Accessor struct {
	GlTFChildOfRootProperty
	ByteStride    int       `json:"byteStride" validator:"gte=0, lte=255"`
	ComponentType int       `json:"componentType" validator:"required"`
	Count         int       `json:"count" validator:"required, gte=1"`
	Type          string    `json:"type,omitempty" validator:"required"`
	Max           []float64 `json:"max,omitempty"`
	Min           []float64 `json:"min,omitempty"`
	BufferView    GlTFid    `json:"bufferView,omitempty" validator:"required"`
	ByteOffset    int       `json:"byteOffset" validator:"required, gte=0"`
}
