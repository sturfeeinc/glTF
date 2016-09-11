/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with glTFModelGenerator (https://github.com/sturfeeinc/glTFModelGenerator)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 * generated 2016-09-11 18:55:47.106260962 +0300 MSK
 */

package glTF

// http://json-schema.org/draft-03/schema
// An attribute or uniform input to a technique, and an optional semantic and value.
type TechniqueParameters struct {
	GlTFProperty
	Type     int         `json:"type,omitempty" validator:"required"`
	Semantic string      `json:"semantic,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Count    int         `json:"count,omitempty" validator:"gte=1.0000000000"`
	Node     GlTFid      `json:"node,omitempty"`
}
