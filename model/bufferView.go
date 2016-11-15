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
// A view into a buffer generally representing a subset of the buffer.
type BufferView struct {
	GlTFChildOfRootProperty
	Buffer     GlTFid `json:"buffer,omitempty" validator:"required"`
	ByteOffset int    `json:"byteOffset" validator:"required, gte=0"`
	ByteLength int    `json:"byteLength" validator:"gte=0"`
	Target     int    `json:"target"`
}
