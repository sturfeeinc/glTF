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
// A buffer points to binary geometry, animation, or skins.
type Buffer struct {
	GlTFChildOfRootProperty
	Uri        string `json:"uri,omitempty" validator:"required"`
	ByteLength int    `json:"byteLength" validator:"gte=0.0000000000"`
	Type       string `json:"type"`
}
