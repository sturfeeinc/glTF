/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type CameraPerspective struct {
	GlTFProperty
	AspectRatio *float64 `json:"aspectRatio" validator:"gte=0.0000000000"`
	Yfov        *float64 `json:"yfov" validator:"required, gte=0.0000000000"`
	Zfar        *float64 `json:"zfar" validator:"required, gte=0.0000000000"`
	Znear       *float64 `json:"znear" validator:"required, gte=0.0000000000"`
}
