/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type CameraOrthographic struct {
	GlTFProperty
	Znear *float64 `json:"znear" validator:"required, gte=0.0000000000"`
	Xmag  *float64 `json:"xmag" validator:"required"`
	Ymag  *float64 `json:"ymag" validator:"required"`
	Zfar  *float64 `json:"zfar" validator:"required, gte=0.0000000000"`
}
