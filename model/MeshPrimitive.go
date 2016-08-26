/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type MeshPrimitive struct {
	GlTFProperty
	Attributes map[string]*MeshPrimitiveAttribute `json:"attributes"`
	Indices    GlTFid                             `json:"indices"`
	Material   GlTFid                             `json:"material" validator:"required"`
	Mode       *int                               `json:"mode"`
}
