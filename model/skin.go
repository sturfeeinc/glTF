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
// Joints and matrices defining a skin.
type Skin struct {
	GlTFChildOfRootProperty
	BindShapeMatrix     []float64 `json:"bindShapeMatrix"`
	InverseBindMatrices GlTFid    `json:"inverseBindMatrices,omitempty" validator:"required"`
	JointNames          []GlTFid  `json:"jointNames,omitempty" validator:"required"`
}
