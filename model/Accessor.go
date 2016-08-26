/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Accessor struct {
	GlTFChildOfRootProperty
	Min           []*float64 `json:"min"`
	BufferView    GlTFid     `json:"bufferView" validator:"required"`
	ByteOffset    *int       `json:"byteOffset" validator:"required, gte=0.0000000000"`
	ByteStride    *int       `json:"byteStride" validator:"gte=0.0000000000, lte=255.0000000000"`
	ComponentType *int       `json:"componentType" validator:"required"`
	Count         *int       `json:"count" validator:"required, gte=1.0000000000"`
	Type          *string    `json:"type" validator:"required"`
	Max           []*float64 `json:"max"`
}
