/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type AnimationChannelTarget struct {
	GlTFProperty
	Id   GlTFid  `json:"id" validator:"required"`
	Path *string `json:"path" validator:"required"`
}
