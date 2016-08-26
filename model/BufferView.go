/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type BufferView struct {
	GlTFChildOfRootProperty
	Target     *int   `json:"target"`
	Buffer     GlTFid `json:"buffer" validator:"required"`
	ByteOffset *int   `json:"byteOffset" validator:"required, gte=0.0000000000"`
	ByteLength *int   `json:"byteLength" validator:"gte=0.0000000000"`
}
