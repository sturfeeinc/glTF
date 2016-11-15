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
// Targets an animation's sampler at a node's property.
type AnimationChannel struct {
	GlTFProperty
	Target  AnimationChannelTarget `json:"target,omitempty" validator:"required"`
	Sampler GlTFid                 `json:"sampler,omitempty" validator:"required"`
}
