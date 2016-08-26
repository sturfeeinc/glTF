/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Asset struct {
	GlTFProperty
	Profile            AssetProfile `json:"profile"`
	Version            *string      `json:"version" validator:"required"`
	Copyright          *string      `json:"copyright"`
	Generator          *string      `json:"generator"`
	PremultipliedAlpha *bool        `json:"premultipliedAlpha"`
}
