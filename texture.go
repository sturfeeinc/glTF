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
// A texture and its sampler.
type Texture struct {
	GlTFChildOfRootProperty
	Format         int    `json:"format"`
	InternalFormat int    `json:"internalFormat"`
	Sampler        GlTFid `json:"sampler,omitempty" validator:"required"`
	Source         GlTFid `json:"source,omitempty" validator:"required"`
	Target         int    `json:"target"`
	Type           int    `json:"type"`
}
