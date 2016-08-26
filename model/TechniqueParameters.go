/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type TechniqueParameters struct {
	GlTFProperty
	Type     *int        `json:"type" validator:"required"`
	Semantic *string     `json:"semantic"`
	Value    interface{} `json:"value"`
	Count    *int        `json:"count" validator:"gte=1.0000000000"`
	Node     GlTFid      `json:"node"`
}
