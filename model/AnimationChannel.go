/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type AnimationChannel struct {
	GlTFProperty
	Sampler GlTFid                 `json:"sampler" validator:"required"`
	Target  AnimationChannelTarget `json:"target" validator:"required"`
}
