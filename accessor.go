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
// A typed view into a bufferView.  A bufferView contains raw binary data.  An accessor provides a typed view into a bufferView or a subset of a bufferView similar to how WebGL's `vertexAttribPointer()` defines an attribute in a buffer.
type Accessor struct {
	GlTFChildOfRootProperty
	ByteStride    int       `json:"byteStride" validator:"gte=0.0000000000, lte=255.0000000000"`
	ComponentType int       `json:"componentType,omitempty" validator:"required"`
	Count         int       `json:"count,omitempty" validator:"required, gte=1.0000000000"`
	Type          string    `json:"type,omitempty" validator:"required"`
	Max           []float64 `json:"max,omitempty"`
	Min           []float64 `json:"min,omitempty"`
	BufferView    GlTFid    `json:"bufferView,omitempty" validator:"required"`
	ByteOffset    int       `json:"byteOffset,omitempty" validator:"required, gte=0.0000000000"`
}
