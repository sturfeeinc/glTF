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
// A template for material appearances.
type Technique struct {
	GlTFChildOfRootProperty
	States     TechniqueStates                `json:"states"`
	Parameters map[string]TechniqueParameters `json:"parameters"`
	Attributes map[string]TechniqueAttribute  `json:"attributes"`
	Program    GlTFid                         `json:"program,omitempty" validator:"required"`
	Uniforms   map[string]TechniqueUniform    `json:"uniforms"`
}
