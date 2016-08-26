/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Node struct {
	GlTFChildOfRootProperty
	Camera      GlTFid     `json:"camera"`
	Children    []*GlTFid  `json:"children"`
	Skeletons   []*GlTFid  `json:"skeletons"`
	Skin        GlTFid     `json:"skin"`
	JointName   GlTFid     `json:"jointName"`
	Matrix      []*float64 `json:"matrix"`
	Meshes      []*GlTFid  `json:"meshes"`
	Rotation    []*float64 `json:"rotation"`
	Scale       []*float64 `json:"scale"`
	Translation []*float64 `json:"translation"`
}
