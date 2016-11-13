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
// A keyframe animation.
type Animation struct {
	GlTFChildOfRootProperty
	Channels   []AnimationChannel            `json:"channels"`
	Parameters map[string]AnimationParameter `json:"parameters"`
	Samplers   map[string]AnimationSampler   `json:"samplers"`
}
