/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Texture struct {
	GlTFChildOfRootProperty
	Type           *int   `json:"type"`
	Format         *int   `json:"format"`
	InternalFormat *int   `json:"internalFormat"`
	Sampler        GlTFid `json:"sampler" validator:"required"`
	Source         GlTFid `json:"source" validator:"required"`
	Target         *int   `json:"target"`
}
