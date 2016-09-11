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
// Arguments for fixed-function rendering state functions other than `enable()`/`disable()`.
type TechniqueStatesFunctions struct {
	GlTFProperty
	BlendColor            []float64 `json:"blendColor"`
	BlendEquationSeparate []int     `json:"blendEquationSeparate"`
	BlendFuncSeparate     []int     `json:"blendFuncSeparate"`
	DepthRange            []float64 `json:"depthRange"`
	FrontFace             []int     `json:"frontFace"`
	LineWidth             []float64 `json:"lineWidth"`
	Scissor               []float64 `json:"scissor"`
	ColorMask             []bool    `json:"colorMask"`
	CullFace              []int     `json:"cullFace"`
	DepthFunc             []int     `json:"depthFunc"`
	DepthMask             []bool    `json:"depthMask"`
	PolygonOffset         []float64 `json:"polygonOffset"`
}
