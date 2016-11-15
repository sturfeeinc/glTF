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
// A perspective camera containing properties to create a perspective projection matrix.
type CameraPerspective struct {
	GlTFProperty
	AspectRatio float64 `json:"aspectRatio,omitempty" validator:"gte=0"`
	Yfov        float64 `json:"yfov,omitempty" validator:"required, gte=0"`
	Zfar        float64 `json:"zfar,omitempty" validator:"required, gte=0"`
	Znear       float64 `json:"znear,omitempty" validator:"required, gte=0"`
}
