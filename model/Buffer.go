/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Buffer struct {
	GlTFChildOfRootProperty
	Uri        *string `json:"uri" validator:"required"`
	ByteLength *int    `json:"byteLength" validator:"gte=0.0000000000"`
	Type       *string `json:"type"`
}
