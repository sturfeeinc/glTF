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
// An orthographic camera containing properties to create an orthographic projection matrix.
type CameraOrthographic struct {
	GlTFProperty
	Xmag  float64 `json:"xmag,omitempty" validator:"required"`
	Ymag  float64 `json:"ymag,omitempty" validator:"required"`
	Zfar  float64 `json:"zfar,omitempty" validator:"required, gte=0.0000000000"`
	Znear float64 `json:"znear,omitempty" validator:"required, gte=0.0000000000"`
}
