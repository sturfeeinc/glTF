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
// A view into a buffer generally representing a subset of the buffer.
type BufferView struct {
	GlTFChildOfRootProperty
	Buffer     GlTFid `json:"buffer,omitempty" validator:"required"`
	ByteOffset int    `json:"byteOffset,omitempty" validator:"required, gte=0.0000000000"`
	ByteLength int    `json:"byteLength" validator:"gte=0.0000000000"`
	Target     int    `json:"target,omitempty"`
}
