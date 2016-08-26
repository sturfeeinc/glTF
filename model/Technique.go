/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Technique struct {
	GlTFChildOfRootProperty
	States     TechniqueStates                 `json:"states"`
	Parameters map[string]*TechniqueParameters `json:"parameters"`
	Attributes map[string]*TechniqueAttribute  `json:"attributes"`
	Program    GlTFid                          `json:"program" validator:"required"`
	Uniforms   map[string]*TechniqueUniform    `json:"uniforms"`
}
