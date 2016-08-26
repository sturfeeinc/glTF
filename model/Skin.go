/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Skin struct {
	GlTFChildOfRootProperty
	BindShapeMatrix     []*float64 `json:"bindShapeMatrix"`
	InverseBindMatrices GlTFid     `json:"inverseBindMatrices" validator:"required"`
	JointNames          []*GlTFid  `json:"jointNames" validator:"required"`
}
