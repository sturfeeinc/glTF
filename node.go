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
// A node in the node hierarchy.  A node can have either the `camera`, `meshes`, or `skeletons`/`skin`/`meshes` properties defined.  In the later case, all `primitives` in the referenced `meshes` contain `JOINT` and `WEIGHT` attributes and the referenced `material`/`technique` from each `primitive` has parameters with `JOINT` and `WEIGHT` semantics.  A node can have either a `matrix` or any combination of `translation`/`rotation`/`scale` (TRS) properties.  If none are provided, the transform is the identity.
type Node struct {
	GlTFChildOfRootProperty
	Meshes      []GlTFid  `json:"meshes,omitempty"`
	Rotation    []float64 `json:"rotation"`
	Scale       []float64 `json:"scale"`
	Translation []float64 `json:"translation"`
	Camera      GlTFid    `json:"camera,omitempty"`
	Skeletons   []GlTFid  `json:"skeletons,omitempty"`
	Skin        GlTFid    `json:"skin,omitempty"`
	JointName   GlTFid    `json:"jointName,omitempty"`
	Matrix      []float64 `json:"matrix"`
	Children    []GlTFid  `json:"children"`
}
