/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with glTFModelGenerator (https://github.com/sturfeeinc/glTFModelGenerator)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 * generated 2016-09-17 21:21:17.912222485 +0300 MSK
 */

package glTF

// http://json-schema.org/draft-03/schema
// An attribute or uniform input to a technique, and an optional semantic and value.
type TechniqueParameters struct {
	GlTFProperty
	Value    interface{} `json:"value,omitempty"`
	Count    int         `json:"count" validator:"gte=1"`
	Node     GlTFid      `json:"node,omitempty"`
	Type     int         `json:"type" validator:"required"`
	Semantic string      `json:"semantic,omitempty"`
}
