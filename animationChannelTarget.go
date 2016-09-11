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
// The ID of the node and TRS property that an animation channel targets.
type AnimationChannelTarget struct {
	GlTFProperty
	Id   GlTFid `json:"id,omitempty" validator:"required"`
	Path string `json:"path,omitempty" validator:"required"`
}
